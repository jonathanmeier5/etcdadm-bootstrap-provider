// +build !ignore_autogenerated_etcd_bootstrap

/*


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

package v1alpha3

import (
	unsafe "unsafe"

	v1beta1 "github.com/mrajashree/etcdadm-bootstrap-provider/api/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	clusterapiapiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	clusterapiapiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	apiv1alpha3 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha3"
	apiv1beta1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*BottlerocketConfig)(nil), (*v1beta1.BottlerocketConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_BottlerocketConfig_To_v1beta1_BottlerocketConfig(a.(*BottlerocketConfig), b.(*v1beta1.BottlerocketConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.BottlerocketConfig)(nil), (*BottlerocketConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_BottlerocketConfig_To_v1alpha3_BottlerocketConfig(a.(*v1beta1.BottlerocketConfig), b.(*BottlerocketConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CloudInitConfig)(nil), (*v1beta1.CloudInitConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_CloudInitConfig_To_v1beta1_CloudInitConfig(a.(*CloudInitConfig), b.(*v1beta1.CloudInitConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CloudInitConfig)(nil), (*CloudInitConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CloudInitConfig_To_v1alpha3_CloudInitConfig(a.(*v1beta1.CloudInitConfig), b.(*CloudInitConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EtcdadmConfig)(nil), (*v1beta1.EtcdadmConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_EtcdadmConfig_To_v1beta1_EtcdadmConfig(a.(*EtcdadmConfig), b.(*v1beta1.EtcdadmConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EtcdadmConfig)(nil), (*EtcdadmConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EtcdadmConfig_To_v1alpha3_EtcdadmConfig(a.(*v1beta1.EtcdadmConfig), b.(*EtcdadmConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EtcdadmConfigList)(nil), (*v1beta1.EtcdadmConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_EtcdadmConfigList_To_v1beta1_EtcdadmConfigList(a.(*EtcdadmConfigList), b.(*v1beta1.EtcdadmConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EtcdadmConfigList)(nil), (*EtcdadmConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EtcdadmConfigList_To_v1alpha3_EtcdadmConfigList(a.(*v1beta1.EtcdadmConfigList), b.(*EtcdadmConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EtcdadmConfigSpec)(nil), (*v1beta1.EtcdadmConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_EtcdadmConfigSpec_To_v1beta1_EtcdadmConfigSpec(a.(*EtcdadmConfigSpec), b.(*v1beta1.EtcdadmConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EtcdadmConfigSpec)(nil), (*EtcdadmConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EtcdadmConfigSpec_To_v1alpha3_EtcdadmConfigSpec(a.(*v1beta1.EtcdadmConfigSpec), b.(*EtcdadmConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EtcdadmConfigStatus)(nil), (*v1beta1.EtcdadmConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_EtcdadmConfigStatus_To_v1beta1_EtcdadmConfigStatus(a.(*EtcdadmConfigStatus), b.(*v1beta1.EtcdadmConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EtcdadmConfigStatus)(nil), (*EtcdadmConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EtcdadmConfigStatus_To_v1alpha3_EtcdadmConfigStatus(a.(*v1beta1.EtcdadmConfigStatus), b.(*EtcdadmConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProxyConfiguration)(nil), (*v1beta1.ProxyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ProxyConfiguration_To_v1beta1_ProxyConfiguration(a.(*ProxyConfiguration), b.(*v1beta1.ProxyConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ProxyConfiguration)(nil), (*ProxyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ProxyConfiguration_To_v1alpha3_ProxyConfiguration(a.(*v1beta1.ProxyConfiguration), b.(*ProxyConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RegistryMirrorConfiguration)(nil), (*v1beta1.RegistryMirrorConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_RegistryMirrorConfiguration_To_v1beta1_RegistryMirrorConfiguration(a.(*RegistryMirrorConfiguration), b.(*v1beta1.RegistryMirrorConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RegistryMirrorConfiguration)(nil), (*RegistryMirrorConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RegistryMirrorConfiguration_To_v1alpha3_RegistryMirrorConfiguration(a.(*v1beta1.RegistryMirrorConfiguration), b.(*RegistryMirrorConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_BottlerocketConfig_To_v1beta1_BottlerocketConfig(in *BottlerocketConfig, out *v1beta1.BottlerocketConfig, s conversion.Scope) error {
	out.EtcdImage = in.EtcdImage
	out.BootstrapImage = in.BootstrapImage
	out.PauseImage = in.PauseImage
	return nil
}

// Convert_v1alpha3_BottlerocketConfig_To_v1beta1_BottlerocketConfig is an autogenerated conversion function.
func Convert_v1alpha3_BottlerocketConfig_To_v1beta1_BottlerocketConfig(in *BottlerocketConfig, out *v1beta1.BottlerocketConfig, s conversion.Scope) error {
	return autoConvert_v1alpha3_BottlerocketConfig_To_v1beta1_BottlerocketConfig(in, out, s)
}

func autoConvert_v1beta1_BottlerocketConfig_To_v1alpha3_BottlerocketConfig(in *v1beta1.BottlerocketConfig, out *BottlerocketConfig, s conversion.Scope) error {
	out.EtcdImage = in.EtcdImage
	out.BootstrapImage = in.BootstrapImage
	out.PauseImage = in.PauseImage
	return nil
}

// Convert_v1beta1_BottlerocketConfig_To_v1alpha3_BottlerocketConfig is an autogenerated conversion function.
func Convert_v1beta1_BottlerocketConfig_To_v1alpha3_BottlerocketConfig(in *v1beta1.BottlerocketConfig, out *BottlerocketConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_BottlerocketConfig_To_v1alpha3_BottlerocketConfig(in, out, s)
}

func autoConvert_v1alpha3_CloudInitConfig_To_v1beta1_CloudInitConfig(in *CloudInitConfig, out *v1beta1.CloudInitConfig, s conversion.Scope) error {
	out.Version = in.Version
	out.EtcdReleaseURL = in.EtcdReleaseURL
	out.InstallDir = in.InstallDir
	return nil
}

// Convert_v1alpha3_CloudInitConfig_To_v1beta1_CloudInitConfig is an autogenerated conversion function.
func Convert_v1alpha3_CloudInitConfig_To_v1beta1_CloudInitConfig(in *CloudInitConfig, out *v1beta1.CloudInitConfig, s conversion.Scope) error {
	return autoConvert_v1alpha3_CloudInitConfig_To_v1beta1_CloudInitConfig(in, out, s)
}

func autoConvert_v1beta1_CloudInitConfig_To_v1alpha3_CloudInitConfig(in *v1beta1.CloudInitConfig, out *CloudInitConfig, s conversion.Scope) error {
	out.Version = in.Version
	out.EtcdReleaseURL = in.EtcdReleaseURL
	out.InstallDir = in.InstallDir
	return nil
}

// Convert_v1beta1_CloudInitConfig_To_v1alpha3_CloudInitConfig is an autogenerated conversion function.
func Convert_v1beta1_CloudInitConfig_To_v1alpha3_CloudInitConfig(in *v1beta1.CloudInitConfig, out *CloudInitConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_CloudInitConfig_To_v1alpha3_CloudInitConfig(in, out, s)
}

func autoConvert_v1alpha3_EtcdadmConfig_To_v1beta1_EtcdadmConfig(in *EtcdadmConfig, out *v1beta1.EtcdadmConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_EtcdadmConfigSpec_To_v1beta1_EtcdadmConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_EtcdadmConfigStatus_To_v1beta1_EtcdadmConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_EtcdadmConfig_To_v1beta1_EtcdadmConfig is an autogenerated conversion function.
func Convert_v1alpha3_EtcdadmConfig_To_v1beta1_EtcdadmConfig(in *EtcdadmConfig, out *v1beta1.EtcdadmConfig, s conversion.Scope) error {
	return autoConvert_v1alpha3_EtcdadmConfig_To_v1beta1_EtcdadmConfig(in, out, s)
}

func autoConvert_v1beta1_EtcdadmConfig_To_v1alpha3_EtcdadmConfig(in *v1beta1.EtcdadmConfig, out *EtcdadmConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_EtcdadmConfigSpec_To_v1alpha3_EtcdadmConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_EtcdadmConfigStatus_To_v1alpha3_EtcdadmConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_EtcdadmConfig_To_v1alpha3_EtcdadmConfig is an autogenerated conversion function.
func Convert_v1beta1_EtcdadmConfig_To_v1alpha3_EtcdadmConfig(in *v1beta1.EtcdadmConfig, out *EtcdadmConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_EtcdadmConfig_To_v1alpha3_EtcdadmConfig(in, out, s)
}

func autoConvert_v1alpha3_EtcdadmConfigList_To_v1beta1_EtcdadmConfigList(in *EtcdadmConfigList, out *v1beta1.EtcdadmConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.EtcdadmConfig)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha3_EtcdadmConfigList_To_v1beta1_EtcdadmConfigList is an autogenerated conversion function.
func Convert_v1alpha3_EtcdadmConfigList_To_v1beta1_EtcdadmConfigList(in *EtcdadmConfigList, out *v1beta1.EtcdadmConfigList, s conversion.Scope) error {
	return autoConvert_v1alpha3_EtcdadmConfigList_To_v1beta1_EtcdadmConfigList(in, out, s)
}

func autoConvert_v1beta1_EtcdadmConfigList_To_v1alpha3_EtcdadmConfigList(in *v1beta1.EtcdadmConfigList, out *EtcdadmConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]EtcdadmConfig)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_EtcdadmConfigList_To_v1alpha3_EtcdadmConfigList is an autogenerated conversion function.
func Convert_v1beta1_EtcdadmConfigList_To_v1alpha3_EtcdadmConfigList(in *v1beta1.EtcdadmConfigList, out *EtcdadmConfigList, s conversion.Scope) error {
	return autoConvert_v1beta1_EtcdadmConfigList_To_v1alpha3_EtcdadmConfigList(in, out, s)
}

func autoConvert_v1alpha3_EtcdadmConfigSpec_To_v1beta1_EtcdadmConfigSpec(in *EtcdadmConfigSpec, out *v1beta1.EtcdadmConfigSpec, s conversion.Scope) error {
	out.Users = *(*[]apiv1beta1.User)(unsafe.Pointer(&in.Users))
	out.EtcdadmBuiltin = in.EtcdadmBuiltin
	out.EtcdadmInstallCommands = *(*[]string)(unsafe.Pointer(&in.EtcdadmInstallCommands))
	out.PreEtcdadmCommands = *(*[]string)(unsafe.Pointer(&in.PreEtcdadmCommands))
	out.PostEtcdadmCommands = *(*[]string)(unsafe.Pointer(&in.PostEtcdadmCommands))
	out.Format = v1beta1.Format(in.Format)
	out.BottlerocketConfig = (*v1beta1.BottlerocketConfig)(unsafe.Pointer(in.BottlerocketConfig))
	out.CloudInitConfig = (*v1beta1.CloudInitConfig)(unsafe.Pointer(in.CloudInitConfig))
	out.Files = *(*[]apiv1beta1.File)(unsafe.Pointer(&in.Files))
	out.Proxy = (*v1beta1.ProxyConfiguration)(unsafe.Pointer(in.Proxy))
	out.RegistryMirror = (*v1beta1.RegistryMirrorConfiguration)(unsafe.Pointer(in.RegistryMirror))
	out.CipherSuites = in.CipherSuites
	return nil
}

// Convert_v1alpha3_EtcdadmConfigSpec_To_v1beta1_EtcdadmConfigSpec is an autogenerated conversion function.
func Convert_v1alpha3_EtcdadmConfigSpec_To_v1beta1_EtcdadmConfigSpec(in *EtcdadmConfigSpec, out *v1beta1.EtcdadmConfigSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_EtcdadmConfigSpec_To_v1beta1_EtcdadmConfigSpec(in, out, s)
}

func autoConvert_v1beta1_EtcdadmConfigSpec_To_v1alpha3_EtcdadmConfigSpec(in *v1beta1.EtcdadmConfigSpec, out *EtcdadmConfigSpec, s conversion.Scope) error {
	out.Users = *(*[]apiv1alpha3.User)(unsafe.Pointer(&in.Users))
	out.EtcdadmBuiltin = in.EtcdadmBuiltin
	out.EtcdadmInstallCommands = *(*[]string)(unsafe.Pointer(&in.EtcdadmInstallCommands))
	out.PreEtcdadmCommands = *(*[]string)(unsafe.Pointer(&in.PreEtcdadmCommands))
	out.PostEtcdadmCommands = *(*[]string)(unsafe.Pointer(&in.PostEtcdadmCommands))
	out.Format = Format(in.Format)
	out.BottlerocketConfig = (*BottlerocketConfig)(unsafe.Pointer(in.BottlerocketConfig))
	out.CloudInitConfig = (*CloudInitConfig)(unsafe.Pointer(in.CloudInitConfig))
	out.Files = *(*[]apiv1alpha3.File)(unsafe.Pointer(&in.Files))
	out.Proxy = (*ProxyConfiguration)(unsafe.Pointer(in.Proxy))
	out.RegistryMirror = (*RegistryMirrorConfiguration)(unsafe.Pointer(in.RegistryMirror))
	out.CipherSuites = in.CipherSuites
	return nil
}

// Convert_v1beta1_EtcdadmConfigSpec_To_v1alpha3_EtcdadmConfigSpec is an autogenerated conversion function.
func Convert_v1beta1_EtcdadmConfigSpec_To_v1alpha3_EtcdadmConfigSpec(in *v1beta1.EtcdadmConfigSpec, out *EtcdadmConfigSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_EtcdadmConfigSpec_To_v1alpha3_EtcdadmConfigSpec(in, out, s)
}

func autoConvert_v1alpha3_EtcdadmConfigStatus_To_v1beta1_EtcdadmConfigStatus(in *EtcdadmConfigStatus, out *v1beta1.EtcdadmConfigStatus, s conversion.Scope) error {
	out.Conditions = *(*clusterapiapiv1beta1.Conditions)(unsafe.Pointer(&in.Conditions))
	out.DataSecretName = (*string)(unsafe.Pointer(in.DataSecretName))
	out.Ready = in.Ready
	return nil
}

// Convert_v1alpha3_EtcdadmConfigStatus_To_v1beta1_EtcdadmConfigStatus is an autogenerated conversion function.
func Convert_v1alpha3_EtcdadmConfigStatus_To_v1beta1_EtcdadmConfigStatus(in *EtcdadmConfigStatus, out *v1beta1.EtcdadmConfigStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_EtcdadmConfigStatus_To_v1beta1_EtcdadmConfigStatus(in, out, s)
}

func autoConvert_v1beta1_EtcdadmConfigStatus_To_v1alpha3_EtcdadmConfigStatus(in *v1beta1.EtcdadmConfigStatus, out *EtcdadmConfigStatus, s conversion.Scope) error {
	out.Conditions = *(*clusterapiapiv1alpha3.Conditions)(unsafe.Pointer(&in.Conditions))
	out.DataSecretName = (*string)(unsafe.Pointer(in.DataSecretName))
	out.Ready = in.Ready
	return nil
}

// Convert_v1beta1_EtcdadmConfigStatus_To_v1alpha3_EtcdadmConfigStatus is an autogenerated conversion function.
func Convert_v1beta1_EtcdadmConfigStatus_To_v1alpha3_EtcdadmConfigStatus(in *v1beta1.EtcdadmConfigStatus, out *EtcdadmConfigStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_EtcdadmConfigStatus_To_v1alpha3_EtcdadmConfigStatus(in, out, s)
}

func autoConvert_v1alpha3_ProxyConfiguration_To_v1beta1_ProxyConfiguration(in *ProxyConfiguration, out *v1beta1.ProxyConfiguration, s conversion.Scope) error {
	out.HTTPProxy = in.HTTPProxy
	out.HTTPSProxy = in.HTTPSProxy
	out.NoProxy = *(*[]string)(unsafe.Pointer(&in.NoProxy))
	return nil
}

// Convert_v1alpha3_ProxyConfiguration_To_v1beta1_ProxyConfiguration is an autogenerated conversion function.
func Convert_v1alpha3_ProxyConfiguration_To_v1beta1_ProxyConfiguration(in *ProxyConfiguration, out *v1beta1.ProxyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha3_ProxyConfiguration_To_v1beta1_ProxyConfiguration(in, out, s)
}

func autoConvert_v1beta1_ProxyConfiguration_To_v1alpha3_ProxyConfiguration(in *v1beta1.ProxyConfiguration, out *ProxyConfiguration, s conversion.Scope) error {
	out.HTTPProxy = in.HTTPProxy
	out.HTTPSProxy = in.HTTPSProxy
	out.NoProxy = *(*[]string)(unsafe.Pointer(&in.NoProxy))
	return nil
}

// Convert_v1beta1_ProxyConfiguration_To_v1alpha3_ProxyConfiguration is an autogenerated conversion function.
func Convert_v1beta1_ProxyConfiguration_To_v1alpha3_ProxyConfiguration(in *v1beta1.ProxyConfiguration, out *ProxyConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_ProxyConfiguration_To_v1alpha3_ProxyConfiguration(in, out, s)
}

func autoConvert_v1alpha3_RegistryMirrorConfiguration_To_v1beta1_RegistryMirrorConfiguration(in *RegistryMirrorConfiguration, out *v1beta1.RegistryMirrorConfiguration, s conversion.Scope) error {
	out.Endpoint = in.Endpoint
	out.CACert = in.CACert
	return nil
}

// Convert_v1alpha3_RegistryMirrorConfiguration_To_v1beta1_RegistryMirrorConfiguration is an autogenerated conversion function.
func Convert_v1alpha3_RegistryMirrorConfiguration_To_v1beta1_RegistryMirrorConfiguration(in *RegistryMirrorConfiguration, out *v1beta1.RegistryMirrorConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha3_RegistryMirrorConfiguration_To_v1beta1_RegistryMirrorConfiguration(in, out, s)
}

func autoConvert_v1beta1_RegistryMirrorConfiguration_To_v1alpha3_RegistryMirrorConfiguration(in *v1beta1.RegistryMirrorConfiguration, out *RegistryMirrorConfiguration, s conversion.Scope) error {
	out.Endpoint = in.Endpoint
	out.CACert = in.CACert
	return nil
}

// Convert_v1beta1_RegistryMirrorConfiguration_To_v1alpha3_RegistryMirrorConfiguration is an autogenerated conversion function.
func Convert_v1beta1_RegistryMirrorConfiguration_To_v1alpha3_RegistryMirrorConfiguration(in *v1beta1.RegistryMirrorConfiguration, out *RegistryMirrorConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_RegistryMirrorConfiguration_To_v1alpha3_RegistryMirrorConfiguration(in, out, s)
}
