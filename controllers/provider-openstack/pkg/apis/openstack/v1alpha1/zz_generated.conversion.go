// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	openstack "github.com/gardener/gardener-extensions/controllers/provider-openstack/pkg/apis/openstack"
	corev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*InfrastructureConfig)(nil), (*openstack.InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureConfig_To_openstack_InfrastructureConfig(a.(*InfrastructureConfig), b.(*openstack.InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*openstack.InfrastructureConfig)(nil), (*InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_openstack_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(a.(*openstack.InfrastructureConfig), b.(*InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureStatus)(nil), (*openstack.InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureStatus_To_openstack_InfrastructureStatus(a.(*InfrastructureStatus), b.(*openstack.InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*openstack.InfrastructureStatus)(nil), (*InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_openstack_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(a.(*openstack.InfrastructureStatus), b.(*InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkStatus)(nil), (*openstack.NetworkStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NetworkStatus_To_openstack_NetworkStatus(a.(*NetworkStatus), b.(*openstack.NetworkStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*openstack.NetworkStatus)(nil), (*NetworkStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_openstack_NetworkStatus_To_v1alpha1_NetworkStatus(a.(*openstack.NetworkStatus), b.(*NetworkStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Networks)(nil), (*openstack.Networks)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Networks_To_openstack_Networks(a.(*Networks), b.(*openstack.Networks), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*openstack.Networks)(nil), (*Networks)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_openstack_Networks_To_v1alpha1_Networks(a.(*openstack.Networks), b.(*Networks), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeStatus)(nil), (*openstack.NodeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NodeStatus_To_openstack_NodeStatus(a.(*NodeStatus), b.(*openstack.NodeStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*openstack.NodeStatus)(nil), (*NodeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_openstack_NodeStatus_To_v1alpha1_NodeStatus(a.(*openstack.NodeStatus), b.(*NodeStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Router)(nil), (*openstack.Router)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Router_To_openstack_Router(a.(*Router), b.(*openstack.Router), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*openstack.Router)(nil), (*Router)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_openstack_Router_To_v1alpha1_Router(a.(*openstack.Router), b.(*Router), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RouterStatus)(nil), (*openstack.RouterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RouterStatus_To_openstack_RouterStatus(a.(*RouterStatus), b.(*openstack.RouterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*openstack.RouterStatus)(nil), (*RouterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_openstack_RouterStatus_To_v1alpha1_RouterStatus(a.(*openstack.RouterStatus), b.(*RouterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SecurityGroup)(nil), (*openstack.SecurityGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SecurityGroup_To_openstack_SecurityGroup(a.(*SecurityGroup), b.(*openstack.SecurityGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*openstack.SecurityGroup)(nil), (*SecurityGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_openstack_SecurityGroup_To_v1alpha1_SecurityGroup(a.(*openstack.SecurityGroup), b.(*SecurityGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Subnet)(nil), (*openstack.Subnet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Subnet_To_openstack_Subnet(a.(*Subnet), b.(*openstack.Subnet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*openstack.Subnet)(nil), (*Subnet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_openstack_Subnet_To_v1alpha1_Subnet(a.(*openstack.Subnet), b.(*Subnet), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_InfrastructureConfig_To_openstack_InfrastructureConfig(in *InfrastructureConfig, out *openstack.InfrastructureConfig, s conversion.Scope) error {
	out.FloatingPoolName = in.FloatingPoolName
	if err := Convert_v1alpha1_Networks_To_openstack_Networks(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_InfrastructureConfig_To_openstack_InfrastructureConfig is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureConfig_To_openstack_InfrastructureConfig(in *InfrastructureConfig, out *openstack.InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureConfig_To_openstack_InfrastructureConfig(in, out, s)
}

func autoConvert_openstack_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *openstack.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	out.FloatingPoolName = in.FloatingPoolName
	if err := Convert_openstack_Networks_To_v1alpha1_Networks(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	return nil
}

// Convert_openstack_InfrastructureConfig_To_v1alpha1_InfrastructureConfig is an autogenerated conversion function.
func Convert_openstack_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *openstack.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_openstack_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureStatus_To_openstack_InfrastructureStatus(in *InfrastructureStatus, out *openstack.InfrastructureStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_NetworkStatus_To_openstack_NetworkStatus(&in.Network, &out.Network, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_RouterStatus_To_openstack_RouterStatus(&in.Router, &out.Router, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_NodeStatus_To_openstack_NodeStatus(&in.Node, &out.Node, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_InfrastructureStatus_To_openstack_InfrastructureStatus is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureStatus_To_openstack_InfrastructureStatus(in *InfrastructureStatus, out *openstack.InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureStatus_To_openstack_InfrastructureStatus(in, out, s)
}

func autoConvert_openstack_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *openstack.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	if err := Convert_openstack_NetworkStatus_To_v1alpha1_NetworkStatus(&in.Network, &out.Network, s); err != nil {
		return err
	}
	if err := Convert_openstack_RouterStatus_To_v1alpha1_RouterStatus(&in.Router, &out.Router, s); err != nil {
		return err
	}
	if err := Convert_openstack_NodeStatus_To_v1alpha1_NodeStatus(&in.Node, &out.Node, s); err != nil {
		return err
	}
	return nil
}

// Convert_openstack_InfrastructureStatus_To_v1alpha1_InfrastructureStatus is an autogenerated conversion function.
func Convert_openstack_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *openstack.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_openstack_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in, out, s)
}

func autoConvert_v1alpha1_NetworkStatus_To_openstack_NetworkStatus(in *NetworkStatus, out *openstack.NetworkStatus, s conversion.Scope) error {
	out.ID = in.ID
	out.Subnets = *(*[]openstack.Subnet)(unsafe.Pointer(&in.Subnets))
	out.SecurityGroups = *(*[]openstack.SecurityGroup)(unsafe.Pointer(&in.SecurityGroups))
	return nil
}

// Convert_v1alpha1_NetworkStatus_To_openstack_NetworkStatus is an autogenerated conversion function.
func Convert_v1alpha1_NetworkStatus_To_openstack_NetworkStatus(in *NetworkStatus, out *openstack.NetworkStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_NetworkStatus_To_openstack_NetworkStatus(in, out, s)
}

func autoConvert_openstack_NetworkStatus_To_v1alpha1_NetworkStatus(in *openstack.NetworkStatus, out *NetworkStatus, s conversion.Scope) error {
	out.ID = in.ID
	out.Subnets = *(*[]Subnet)(unsafe.Pointer(&in.Subnets))
	out.SecurityGroups = *(*[]SecurityGroup)(unsafe.Pointer(&in.SecurityGroups))
	return nil
}

// Convert_openstack_NetworkStatus_To_v1alpha1_NetworkStatus is an autogenerated conversion function.
func Convert_openstack_NetworkStatus_To_v1alpha1_NetworkStatus(in *openstack.NetworkStatus, out *NetworkStatus, s conversion.Scope) error {
	return autoConvert_openstack_NetworkStatus_To_v1alpha1_NetworkStatus(in, out, s)
}

func autoConvert_v1alpha1_Networks_To_openstack_Networks(in *Networks, out *openstack.Networks, s conversion.Scope) error {
	out.Router = (*openstack.Router)(unsafe.Pointer(in.Router))
	out.Worker = corev1alpha1.CIDR(in.Worker)
	return nil
}

// Convert_v1alpha1_Networks_To_openstack_Networks is an autogenerated conversion function.
func Convert_v1alpha1_Networks_To_openstack_Networks(in *Networks, out *openstack.Networks, s conversion.Scope) error {
	return autoConvert_v1alpha1_Networks_To_openstack_Networks(in, out, s)
}

func autoConvert_openstack_Networks_To_v1alpha1_Networks(in *openstack.Networks, out *Networks, s conversion.Scope) error {
	out.Router = (*Router)(unsafe.Pointer(in.Router))
	out.Worker = corev1alpha1.CIDR(in.Worker)
	return nil
}

// Convert_openstack_Networks_To_v1alpha1_Networks is an autogenerated conversion function.
func Convert_openstack_Networks_To_v1alpha1_Networks(in *openstack.Networks, out *Networks, s conversion.Scope) error {
	return autoConvert_openstack_Networks_To_v1alpha1_Networks(in, out, s)
}

func autoConvert_v1alpha1_NodeStatus_To_openstack_NodeStatus(in *NodeStatus, out *openstack.NodeStatus, s conversion.Scope) error {
	out.KeyName = in.KeyName
	return nil
}

// Convert_v1alpha1_NodeStatus_To_openstack_NodeStatus is an autogenerated conversion function.
func Convert_v1alpha1_NodeStatus_To_openstack_NodeStatus(in *NodeStatus, out *openstack.NodeStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeStatus_To_openstack_NodeStatus(in, out, s)
}

func autoConvert_openstack_NodeStatus_To_v1alpha1_NodeStatus(in *openstack.NodeStatus, out *NodeStatus, s conversion.Scope) error {
	out.KeyName = in.KeyName
	return nil
}

// Convert_openstack_NodeStatus_To_v1alpha1_NodeStatus is an autogenerated conversion function.
func Convert_openstack_NodeStatus_To_v1alpha1_NodeStatus(in *openstack.NodeStatus, out *NodeStatus, s conversion.Scope) error {
	return autoConvert_openstack_NodeStatus_To_v1alpha1_NodeStatus(in, out, s)
}

func autoConvert_v1alpha1_Router_To_openstack_Router(in *Router, out *openstack.Router, s conversion.Scope) error {
	out.ID = in.ID
	return nil
}

// Convert_v1alpha1_Router_To_openstack_Router is an autogenerated conversion function.
func Convert_v1alpha1_Router_To_openstack_Router(in *Router, out *openstack.Router, s conversion.Scope) error {
	return autoConvert_v1alpha1_Router_To_openstack_Router(in, out, s)
}

func autoConvert_openstack_Router_To_v1alpha1_Router(in *openstack.Router, out *Router, s conversion.Scope) error {
	out.ID = in.ID
	return nil
}

// Convert_openstack_Router_To_v1alpha1_Router is an autogenerated conversion function.
func Convert_openstack_Router_To_v1alpha1_Router(in *openstack.Router, out *Router, s conversion.Scope) error {
	return autoConvert_openstack_Router_To_v1alpha1_Router(in, out, s)
}

func autoConvert_v1alpha1_RouterStatus_To_openstack_RouterStatus(in *RouterStatus, out *openstack.RouterStatus, s conversion.Scope) error {
	out.ID = in.ID
	out.Subnets = *(*[]openstack.Subnet)(unsafe.Pointer(&in.Subnets))
	return nil
}

// Convert_v1alpha1_RouterStatus_To_openstack_RouterStatus is an autogenerated conversion function.
func Convert_v1alpha1_RouterStatus_To_openstack_RouterStatus(in *RouterStatus, out *openstack.RouterStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_RouterStatus_To_openstack_RouterStatus(in, out, s)
}

func autoConvert_openstack_RouterStatus_To_v1alpha1_RouterStatus(in *openstack.RouterStatus, out *RouterStatus, s conversion.Scope) error {
	out.ID = in.ID
	out.Subnets = *(*[]Subnet)(unsafe.Pointer(&in.Subnets))
	return nil
}

// Convert_openstack_RouterStatus_To_v1alpha1_RouterStatus is an autogenerated conversion function.
func Convert_openstack_RouterStatus_To_v1alpha1_RouterStatus(in *openstack.RouterStatus, out *RouterStatus, s conversion.Scope) error {
	return autoConvert_openstack_RouterStatus_To_v1alpha1_RouterStatus(in, out, s)
}

func autoConvert_v1alpha1_SecurityGroup_To_openstack_SecurityGroup(in *SecurityGroup, out *openstack.SecurityGroup, s conversion.Scope) error {
	out.Purpose = in.Purpose
	out.ID = in.ID
	return nil
}

// Convert_v1alpha1_SecurityGroup_To_openstack_SecurityGroup is an autogenerated conversion function.
func Convert_v1alpha1_SecurityGroup_To_openstack_SecurityGroup(in *SecurityGroup, out *openstack.SecurityGroup, s conversion.Scope) error {
	return autoConvert_v1alpha1_SecurityGroup_To_openstack_SecurityGroup(in, out, s)
}

func autoConvert_openstack_SecurityGroup_To_v1alpha1_SecurityGroup(in *openstack.SecurityGroup, out *SecurityGroup, s conversion.Scope) error {
	out.Purpose = in.Purpose
	out.ID = in.ID
	return nil
}

// Convert_openstack_SecurityGroup_To_v1alpha1_SecurityGroup is an autogenerated conversion function.
func Convert_openstack_SecurityGroup_To_v1alpha1_SecurityGroup(in *openstack.SecurityGroup, out *SecurityGroup, s conversion.Scope) error {
	return autoConvert_openstack_SecurityGroup_To_v1alpha1_SecurityGroup(in, out, s)
}

func autoConvert_v1alpha1_Subnet_To_openstack_Subnet(in *Subnet, out *openstack.Subnet, s conversion.Scope) error {
	out.Purpose = openstack.SubnetPurpose(in.Purpose)
	out.ID = in.ID
	return nil
}

// Convert_v1alpha1_Subnet_To_openstack_Subnet is an autogenerated conversion function.
func Convert_v1alpha1_Subnet_To_openstack_Subnet(in *Subnet, out *openstack.Subnet, s conversion.Scope) error {
	return autoConvert_v1alpha1_Subnet_To_openstack_Subnet(in, out, s)
}

func autoConvert_openstack_Subnet_To_v1alpha1_Subnet(in *openstack.Subnet, out *Subnet, s conversion.Scope) error {
	out.Purpose = SubnetPurpose(in.Purpose)
	out.ID = in.ID
	return nil
}

// Convert_openstack_Subnet_To_v1alpha1_Subnet is an autogenerated conversion function.
func Convert_openstack_Subnet_To_v1alpha1_Subnet(in *openstack.Subnet, out *Subnet, s conversion.Scope) error {
	return autoConvert_openstack_Subnet_To_v1alpha1_Subnet(in, out, s)
}
