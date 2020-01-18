// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package worker

import (
	"context"
	"fmt"
	"path/filepath"

	apisaws "github.com/gardener/gardener-extensions/controllers/provider-aws/pkg/apis/aws"
	awsapi "github.com/gardener/gardener-extensions/controllers/provider-aws/pkg/apis/aws"
	awsapihelper "github.com/gardener/gardener-extensions/controllers/provider-aws/pkg/apis/aws/helper"
	"github.com/gardener/gardener-extensions/controllers/provider-aws/pkg/aws"
	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	"github.com/gardener/gardener-extensions/pkg/controller/worker"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"

	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

// MachineClassKind yields the name of the AWS machine class.
func (w *workerDelegate) MachineClassKind() string {
	return "AWSMachineClass"
}

// MachineClassList yields a newly initialized AWSMachineClassList object.
func (w *workerDelegate) MachineClassList() runtime.Object {
	return &machinev1alpha1.AWSMachineClassList{}
}

// DeployMachineClasses generates and creates the AWS specific machine classes.
func (w *workerDelegate) DeployMachineClasses(ctx context.Context) error {
	if w.machineClasses == nil {
		if err := w.generateMachineConfig(ctx); err != nil {
			return err
		}
	}
	return w.seedChartApplier.ApplyChart(ctx, filepath.Join(aws.InternalChartsPath, "machineclass"), w.worker.Namespace, "machineclass", map[string]interface{}{"machineClasses": w.machineClasses}, nil)
}

// GenerateMachineDeployments generates the configuration for the desired machine deployments.
func (w *workerDelegate) GenerateMachineDeployments(ctx context.Context) (worker.MachineDeployments, error) {
	if w.machineDeployments == nil {
		if err := w.generateMachineConfig(ctx); err != nil {
			return nil, err
		}
	}
	return w.machineDeployments, nil
}

func (w *workerDelegate) generateMachineClassSecretData(ctx context.Context) (map[string][]byte, error) {
	secret, err := extensionscontroller.GetSecretByReference(ctx, w.client, &w.worker.Spec.SecretRef)
	if err != nil {
		return nil, err
	}

	credentials, err := aws.ReadCredentialsSecret(secret)
	if err != nil {
		return nil, err
	}

	return map[string][]byte{
		machinev1alpha1.AWSAccessKeyID:     credentials.AccessKeyID,
		machinev1alpha1.AWSSecretAccessKey: credentials.SecretAccessKey,
	}, nil
}

func (w *workerDelegate) generateMachineConfig(ctx context.Context) error {
	var (
		machineDeployments = worker.MachineDeployments{}
		machineClasses     []map[string]interface{}
		machineImages      []apisaws.MachineImage
	)

	machineClassSecretData, err := w.generateMachineClassSecretData(ctx)
	if err != nil {
		return err
	}

	infrastructureStatus := &awsapi.InfrastructureStatus{}
	if _, _, err := w.decoder.Decode(w.worker.Spec.InfrastructureProviderStatus.Raw, nil, infrastructureStatus); err != nil {
		return err
	}

	nodesInstanceProfile, err := awsapihelper.FindInstanceProfileForPurpose(infrastructureStatus.IAM.InstanceProfiles, awsapi.PurposeNodes)
	if err != nil {
		return err
	}
	nodesSecurityGroup, err := awsapihelper.FindSecurityGroupForPurpose(infrastructureStatus.VPC.SecurityGroups, awsapi.PurposeNodes)
	if err != nil {
		return err
	}

	for _, pool := range w.worker.Spec.Pools {
		zoneLen := len(pool.Zones)

		workerPoolHash, err := worker.WorkerPoolHash(pool, w.cluster)
		if err != nil {
			return err
		}

		ami, err := w.findMachineImage(pool.MachineImage.Name, pool.MachineImage.Version, w.worker.Spec.Region)
		if err != nil {
			return err
		}
		machineImages = appendMachineImage(machineImages, apisaws.MachineImage{
			Name:    pool.MachineImage.Name,
			Version: pool.MachineImage.Version,
			AMI:     ami,
		})

		workerConfig := &awsapi.WorkerConfig{}
		if pool.ProviderConfig != nil && pool.ProviderConfig.Raw != nil {
			if _, _, err := w.decoder.Decode(pool.ProviderConfig.Raw, nil, workerConfig); err != nil {
				return fmt.Errorf("could not decode provider config: %+v", err)
			}
		}

		volumeSize, err := worker.DiskSize(pool.Volume.Size)
		if err != nil {
			return err
		}
		ebs := map[string]interface{}{
			"volumeSize": volumeSize,
		}
		if pool.Volume.Type != nil {
			ebs["volumeType"] = *pool.Volume.Type
		}
		if workerConfig.Volume != nil && workerConfig.Volume.IOPS != nil {
			ebs["iops"] = *workerConfig.Volume.IOPS
		}
		ebs["deleteOnTermination"] = true

		rootDevice := map[string]interface{}{
			"deviceName" : "/root",
			"ebs" : ebs,
		}

		var blockDevices []map[string]interface{}
		blockDevices = append(blockDevices,  rootDevice)

		for _, vol := range pool.DataVolumes {
			volumeSize, err := worker.DiskSize(vol.Size)
			if err != nil {
				return err
			}
			ebs := map[string]interface{}{
				"volumeSize": volumeSize,
			}
			if vol.Type != nil {
				ebs["volumeType"] = vol.Type
			}
			ebs["encrypted"] = vol.Encrypted
			ebs["deleteOnTermination"] = true
			device := map[string]interface{}{
				"deviceName" : vol.Name,
				"ebs" : ebs,
			}
			blockDevices = append(blockDevices, device)
		}

		for zoneIndex, zone := range pool.Zones {
			nodesSubnet, err := awsapihelper.FindSubnetForPurposeAndZone(infrastructureStatus.VPC.Subnets, awsapi.PurposeNodes, zone)
			if err != nil {
				return err
			}

			machineClassSpec := map[string]interface{}{
				"ami":                ami,
				"region":             w.worker.Spec.Region,
				"machineType":        pool.MachineType,
				"iamInstanceProfile": nodesInstanceProfile.Name,
				"keyName":            infrastructureStatus.EC2.KeyName,
				"networkInterfaces": []map[string]interface{}{
					{
						"subnetID":         nodesSubnet.ID,
						"securityGroupIDs": []string{nodesSecurityGroup.ID},
					},
				},
				"tags": map[string]string{
					fmt.Sprintf("kubernetes.io/cluster/%s", w.worker.Namespace): "1",
					"kubernetes.io/role/node":                                   "1",
				},
				"secret": map[string]interface{}{
					"cloudConfig": string(pool.UserData),
				},
				"blockDevices": blockDevices,
				//"blockDevices": []map[string]interface{}{
				//	{
				//		"ebs": ebs,
				//	},
				//},
			}

			var (
				deploymentName = fmt.Sprintf("%s-%s-z%d", w.worker.Namespace, pool.Name, zoneIndex+1)
				className      = fmt.Sprintf("%s-%s", deploymentName, workerPoolHash)
			)

			machineDeployments = append(machineDeployments, worker.MachineDeployment{
				Name:           deploymentName,
				ClassName:      className,
				SecretName:     className,
				Minimum:        worker.DistributeOverZones(zoneIndex, pool.Minimum, zoneLen),
				Maximum:        worker.DistributeOverZones(zoneIndex, pool.Maximum, zoneLen),
				MaxSurge:       worker.DistributePositiveIntOrPercent(zoneIndex, pool.MaxSurge, zoneLen, pool.Maximum),
				MaxUnavailable: worker.DistributePositiveIntOrPercent(zoneIndex, pool.MaxUnavailable, zoneLen, pool.Minimum),
				Labels:         pool.Labels,
				Annotations:    pool.Annotations,
				Taints:         pool.Taints,
			})

			machineClassSpec["name"] = className
			machineClassSpec["labels"] = map[string]string{
				v1beta1constants.GardenPurpose: v1beta1constants.GardenPurposeMachineClass,
			}
			machineClassSpec["secret"].(map[string]interface{})[aws.AccessKeyID] = string(machineClassSecretData[machinev1alpha1.AWSAccessKeyID])
			machineClassSpec["secret"].(map[string]interface{})[aws.SecretAccessKey] = string(machineClassSecretData[machinev1alpha1.AWSSecretAccessKey])

			machineClasses = append(machineClasses, machineClassSpec)
		}
	}

	w.machineDeployments = machineDeployments
	w.machineClasses = machineClasses
	w.machineImages = machineImages

	return nil
}
