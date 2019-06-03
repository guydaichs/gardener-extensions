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

package controlplanebackup

import (
	"context"
	"path"

	"github.com/gardener/gardener-extensions/controllers/provider-gcp/pkg/apis/config"
	"github.com/gardener/gardener-extensions/controllers/provider-gcp/pkg/gcp"
	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	"github.com/gardener/gardener-extensions/pkg/webhook/controlplane"
	"github.com/gardener/gardener-extensions/pkg/webhook/controlplane/genericmutator"

	"github.com/gardener/gardener/pkg/operation/common"
	"github.com/gardener/gardener/pkg/utils/imagevector"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// NewEnsurer creates a new controlplaneexposure ensurer.
func NewEnsurer(etcdBackup *config.ETCDBackup, imageVector imagevector.ImageVector, logger logr.Logger) genericmutator.Ensurer {
	return &ensurer{
		etcdBackup:  etcdBackup,
		imageVector: imageVector,
		logger:      logger.WithName("gcp-controlplanebackup-ensurer"),
	}
}

type ensurer struct {
	genericmutator.NoopEnsurer
	etcdBackup  *config.ETCDBackup
	imageVector imagevector.ImageVector
	logger      logr.Logger
}

// EnsureETCDStatefulSet ensures that the etcd stateful sets conform to the provider requirements.
func (e *ensurer) EnsureETCDStatefulSet(ctx context.Context, ss *appsv1.StatefulSet, cluster *extensionscontroller.Cluster) error {
	if err := e.ensureContainers(&ss.Spec.Template.Spec, ss.Name, cluster); err != nil {
		return err
	}
	e.ensureVolumes(&ss.Spec.Template.Spec, ss.Name)
	return nil
}

func (e *ensurer) ensureContainers(ps *corev1.PodSpec, name string, cluster *extensionscontroller.Cluster) error {
	c, err := e.getBackupRestoreContainer(name, cluster)
	if err != nil {
		return err
	}
	ps.Containers = controlplane.EnsureContainerWithName(ps.Containers, *c)
	return nil
}

func (e *ensurer) getBackupRestoreContainer(name string, cluster *extensionscontroller.Cluster) (*corev1.Container, error) {
	// Find etcd-backup-restore image
	// TODO Get seed version from clientset when it's possible to inject it
	image, err := e.imageVector.FindImage(gcp.ETCDBackupRestoreImageName, "", cluster.Shoot.Spec.Kubernetes.Version)
	if err != nil {
		return nil, errors.Wrapf(err, "could not find image %s", gcp.ETCDBackupRestoreImageName)
	}

	const (
		mountPath       = "/root/.gcp/"
		defaultSchedule = "0 */24 * * *"
	)

	// Determine provider, container env variables, and volume mounts
	// They are only specified for the etcd-main stateful set (backup is enabled)
	var (
		provider     string
		env          []corev1.EnvVar
		volumeMounts []corev1.VolumeMount
	)
	if name == common.EtcdMainStatefulSetName {
		provider = gcp.StorageProviderName
		env = []corev1.EnvVar{
			{
				Name: "STORAGE_CONTAINER",
				// The bucket name is written to the backup secret by Gardener as a temporary solution.
				// TODO In the future, the bucket name should come from a BackupBucket resource (see https://github.com/gardener/gardener/blob/master/docs/proposals/02-backupinfra.md)
				ValueFrom: &corev1.EnvVarSource{
					SecretKeyRef: &corev1.SecretKeySelector{
						Key:                  gcp.BucketName,
						LocalObjectReference: corev1.LocalObjectReference{Name: gcp.BackupSecretName},
					},
				},
			},
			{
				Name:  "GOOGLE_APPLICATION_CREDENTIALS",
				Value: path.Join(mountPath, gcp.ServiceAccountJSONField),
			},
		}
		volumeMounts = []corev1.VolumeMount{
			{
				Name:      gcp.BackupSecretName,
				MountPath: mountPath,
			},
		}
	}

	// Determine schedule
	var schedule = defaultSchedule
	if e.etcdBackup.Schedule != nil {
		schedule = defaultSchedule
	}

	return controlplane.GetBackupRestoreContainer(name, schedule, provider, image.String(), nil, env, volumeMounts), nil
}

func (e *ensurer) ensureVolumes(ps *corev1.PodSpec, name string) {
	if name == common.EtcdMainStatefulSetName {
		etcdBackupSecretVolume := corev1.Volume{
			Name: gcp.BackupSecretName,
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: gcp.BackupSecretName,
				},
			},
		}
		ps.Volumes = controlplane.EnsureVolumeWithName(ps.Volumes, etcdBackupSecretVolume)
	}
}
