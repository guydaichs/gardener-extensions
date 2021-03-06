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

package internal

import (
	"fmt"

	"github.com/gardener/gardener-extensions/controllers/provider-azure/pkg/apis/azure/install"
	azurev1alpha1 "github.com/gardener/gardener-extensions/controllers/provider-azure/pkg/apis/azure/v1alpha1"

	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var (
	// Scheme is a scheme with the types relevant for Azure actuators.
	Scheme *runtime.Scheme

	decoder runtime.Decoder
)

func init() {
	Scheme = runtime.NewScheme()
	utilruntime.Must(install.AddToScheme(Scheme))

	decoder = serializer.NewCodecFactory(Scheme).UniversalDecoder()
}

// InfrastructureConfigFromInfrastructure extracts the InfrastructureConfig from the
// ProviderConfig section of the given Infrastructure.
func InfrastructureConfigFromInfrastructure(infra *extensionsv1alpha1.Infrastructure) (*azurev1alpha1.InfrastructureConfig, error) {
	config := &azurev1alpha1.InfrastructureConfig{}
	if infra.Spec.ProviderConfig != nil && infra.Spec.ProviderConfig.Raw != nil {
		if _, _, err := decoder.Decode(infra.Spec.ProviderConfig.Raw, nil, config); err != nil {
			return nil, err
		}

		return config, nil
	}
	return nil, fmt.Errorf("infrastructure config is not set on the infrastructure resource")
}

// CloudProfileConfigFromCloudProfile extracts the CloudProfileConfig from the
// ProviderConfig section of the given CloudProfile.
func CloudProfileConfigFromCloudProfile(cloudProfile *gardencorev1alpha1.CloudProfile) (*azurev1alpha1.CloudProfileConfig, error) {
	config := &azurev1alpha1.CloudProfileConfig{}
	if cloudProfile.Spec.ProviderConfig != nil && cloudProfile.Spec.ProviderConfig.Raw != nil {
		if _, _, err := decoder.Decode(cloudProfile.Spec.ProviderConfig.Raw, nil, config); err != nil {
			return nil, err
		}

		return config, nil
	}
	return nil, fmt.Errorf("cloud profile config is not set on the cloudprofile resource")
}
