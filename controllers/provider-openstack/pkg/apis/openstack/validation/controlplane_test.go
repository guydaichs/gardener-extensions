// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package validation_test

import (
	apisopenstack "github.com/gardener/gardener-extensions/controllers/provider-openstack/pkg/apis/openstack"
	. "github.com/gardener/gardener-extensions/controllers/provider-openstack/pkg/apis/openstack/validation"

	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

var _ = Describe("ControlPlaneConfig validation", func() {
	var (
		region      = "foo"
		zone        = "some-zone"
		lbProvider1 = "foo"

		regions = []gardencorev1alpha1.Region{
			{
				Name: region,
				Zones: []gardencorev1alpha1.AvailabilityZone{
					{Name: zone},
				},
			},
		}

		constraints = apisopenstack.Constraints{
			LoadBalancerProviders: []apisopenstack.LoadBalancerProvider{
				{Name: lbProvider1},
			},
		}

		controlPlane *apisopenstack.ControlPlaneConfig
	)

	Describe("#ValidateControlPlaneConfig", func() {
		BeforeEach(func() {
			controlPlane = &apisopenstack.ControlPlaneConfig{
				LoadBalancerProvider: lbProvider1,
				Zone:                 "some-zone",
			}
		})

		It("should return no errors for a valid configuration", func() {
			Expect(ValidateControlPlaneConfig(controlPlane, region, regions, constraints)).To(BeEmpty())
		})

		It("should require the name of a load balancer provider", func() {
			controlPlane.LoadBalancerProvider = ""

			errorList := ValidateControlPlaneConfig(controlPlane, region, regions, constraints)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("loadBalancerProvider"),
			}))))
		})

		It("should require a name of a load balancer provider that is part of the constraints", func() {
			controlPlane.LoadBalancerProvider = "bar"

			errorList := ValidateControlPlaneConfig(controlPlane, region, regions, constraints)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeNotSupported),
				"Field": Equal("loadBalancerProvider"),
			}))))
		})

		It("should require the name of a zone", func() {
			controlPlane.Zone = ""

			errorList := ValidateControlPlaneConfig(controlPlane, region, regions, constraints)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("zone"),
			}))))
		})

		It("should require a name of a zone that is part of the regions", func() {
			controlPlane.Zone = "bar"

			errorList := ValidateControlPlaneConfig(controlPlane, region, regions, constraints)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeNotSupported),
				"Field": Equal("zone"),
			}))))
		})
	})

	Describe("#ValidateControlPlaneConfigUpdate", func() {
		It("should return no errors for an unchanged config", func() {
			Expect(ValidateControlPlaneConfigUpdate(controlPlane, controlPlane, region, regions, constraints)).To(BeEmpty())
		})

		It("should forbid changing the zone", func() {
			newControlPlane := controlPlane.DeepCopy()
			newControlPlane.Zone = "foo"

			errorList := ValidateControlPlaneConfigUpdate(controlPlane, newControlPlane, region, regions, constraints)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("zone"),
			}))))
		})
	})
})
