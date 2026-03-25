// Reusable CDK constructs for provisioning a small EC2-hosted Docker service with Nginx, EBS, KMS, and SSM access.
package cdkec2servicemodule

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-ec2-service-module.Ec2DockerService",
		reflect.TypeOf((*Ec2DockerService)(nil)).Elem(),
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
			j := jsiiProxy_Ec2DockerService{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.Ec2DockerServiceProps",
		reflect.TypeOf((*Ec2DockerServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.Ec2DockerServiceRuntimeProps",
		reflect.TypeOf((*Ec2DockerServiceRuntimeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.Ec2OperationalControls",
		reflect.TypeOf((*Ec2OperationalControls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.Ec2ServiceExposure",
		reflect.TypeOf((*Ec2ServiceExposure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.IngressRule",
		reflect.TypeOf((*IngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.NetworkAddressableServiceOutputs",
		reflect.TypeOf((*NetworkAddressableServiceOutputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.PlatformServiceIdentity",
		reflect.TypeOf((*PlatformServiceIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.PlatformServiceOutputs",
		reflect.TypeOf((*PlatformServiceOutputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.PlatformServiceProps",
		reflect.TypeOf((*PlatformServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-ec2-service-module.PrivateEc2DockerService",
		reflect.TypeOf((*PrivateEc2DockerService)(nil)).Elem(),
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
			j := jsiiProxy_PrivateEc2DockerService{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.PrivateEc2DockerServiceProps",
		reflect.TypeOf((*PrivateEc2DockerServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.ResolvedPlatformServiceIdentity",
		reflect.TypeOf((*ResolvedPlatformServiceIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-service-module.ServiceInfrastructureProps",
		reflect.TypeOf((*ServiceInfrastructureProps)(nil)).Elem(),
	)
}
