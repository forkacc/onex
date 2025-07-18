//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	config "github.com/onexstack/onex/internal/controller/job/apis/config"
	configv1beta1 "github.com/onexstack/onex/pkg/config/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*JobControllerConfiguration)(nil), (*config.JobControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_JobControllerConfiguration_To_config_JobControllerConfiguration(a.(*JobControllerConfiguration), b.(*config.JobControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.JobControllerConfiguration)(nil), (*JobControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_JobControllerConfiguration_To_v1beta1_JobControllerConfiguration(a.(*config.JobControllerConfiguration), b.(*JobControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_JobControllerConfiguration_To_config_JobControllerConfiguration(in *JobControllerConfiguration, out *config.JobControllerConfiguration, s conversion.Scope) error {
	if err := configv1beta1.Convert_v1beta1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(&in.Generic, &out.Generic, s); err != nil {
		return err
	}
	out.ConcurrentCronJobSyncs = in.ConcurrentCronJobSyncs
	out.ConcurrentJobSyncs = in.ConcurrentJobSyncs
	return nil
}

// Convert_v1beta1_JobControllerConfiguration_To_config_JobControllerConfiguration is an autogenerated conversion function.
func Convert_v1beta1_JobControllerConfiguration_To_config_JobControllerConfiguration(in *JobControllerConfiguration, out *config.JobControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_JobControllerConfiguration_To_config_JobControllerConfiguration(in, out, s)
}

func autoConvert_config_JobControllerConfiguration_To_v1beta1_JobControllerConfiguration(in *config.JobControllerConfiguration, out *JobControllerConfiguration, s conversion.Scope) error {
	if err := configv1beta1.Convert_config_GenericControllerManagerConfiguration_To_v1beta1_GenericControllerManagerConfiguration(&in.Generic, &out.Generic, s); err != nil {
		return err
	}
	out.ConcurrentCronJobSyncs = in.ConcurrentCronJobSyncs
	out.ConcurrentJobSyncs = in.ConcurrentJobSyncs
	return nil
}

// Convert_config_JobControllerConfiguration_To_v1beta1_JobControllerConfiguration is an autogenerated conversion function.
func Convert_config_JobControllerConfiguration_To_v1beta1_JobControllerConfiguration(in *config.JobControllerConfiguration, out *JobControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_JobControllerConfiguration_To_v1beta1_JobControllerConfiguration(in, out, s)
}
