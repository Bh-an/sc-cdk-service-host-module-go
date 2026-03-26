// Reusable CDK constructs for provisioning a small EC2-hosted Docker service with Nginx, EBS, KMS, and SSM access.
package cdkservicehostmodule

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk-service-host-module.IngressRule",
		reflect.TypeOf((*IngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.NetworkAddressableServiceOutputs",
		reflect.TypeOf((*NetworkAddressableServiceOutputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.PlatformServiceIdentity",
		reflect.TypeOf((*PlatformServiceIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.PlatformServiceOutputs",
		reflect.TypeOf((*PlatformServiceOutputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.PlatformServiceProps",
		reflect.TypeOf((*PlatformServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-service-host-module.PrivateServiceHost",
		reflect.TypeOf((*PrivateServiceHost)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dataKey", GoGetter: "DataKey"},
			_jsii_.MemberProperty{JsiiProperty: "dataMountPath", GoGetter: "DataMountPath"},
			_jsii_.MemberProperty{JsiiProperty: "dataVolumeDeviceName", GoGetter: "DataVolumeDeviceName"},
			_jsii_.MemberProperty{JsiiProperty: "elasticIp", GoGetter: "ElasticIp"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "serviceIdentity", GoGetter: "ServiceIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "serviceOutputs", GoGetter: "ServiceOutputs"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PrivateServiceHost{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.PrivateServiceHostProps",
		reflect.TypeOf((*PrivateServiceHostProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-service-host-module.PublicServiceHost",
		reflect.TypeOf((*PublicServiceHost)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dataKey", GoGetter: "DataKey"},
			_jsii_.MemberProperty{JsiiProperty: "dataMountPath", GoGetter: "DataMountPath"},
			_jsii_.MemberProperty{JsiiProperty: "dataVolumeDeviceName", GoGetter: "DataVolumeDeviceName"},
			_jsii_.MemberProperty{JsiiProperty: "elasticIp", GoGetter: "ElasticIp"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "serviceIdentity", GoGetter: "ServiceIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "serviceOutputs", GoGetter: "ServiceOutputs"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PublicServiceHost{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.PublicServiceHostProps",
		reflect.TypeOf((*PublicServiceHostProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.ResolvedPlatformServiceIdentity",
		reflect.TypeOf((*ResolvedPlatformServiceIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.ServiceHostExposure",
		reflect.TypeOf((*ServiceHostExposure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.ServiceHostOperationalControls",
		reflect.TypeOf((*ServiceHostOperationalControls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.ServiceHostRuntimeProps",
		reflect.TypeOf((*ServiceHostRuntimeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-service-host-module.ServiceInfrastructureProps",
		reflect.TypeOf((*ServiceInfrastructureProps)(nil)).Elem(),
	)
}
