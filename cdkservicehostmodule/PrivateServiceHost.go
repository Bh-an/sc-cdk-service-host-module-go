package cdkservicehostmodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/Bh-an/sc-cdk-service-host-module-go/cdkservicehostmodule/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/Bh-an/sc-cdk-service-host-module-go/cdkservicehostmodule/internal"
)

type PrivateServiceHost interface {
	constructs.Construct
	DataKey() awskms.IKey
	DataMountPath() *string
	DataVolumeDeviceName() *string
	ElasticIp() awsec2.CfnEIP
	Instance() awsec2.Instance
	// The tree node.
	Node() constructs.Node
	Role() awsiam.IRole
	SecurityGroup() awsec2.ISecurityGroup
	ServiceIdentity() *ResolvedPlatformServiceIdentity
	ServiceOutputs() *NetworkAddressableServiceOutputs
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for PrivateServiceHost
type jsiiProxy_PrivateServiceHost struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PrivateServiceHost) DataKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"dataKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateServiceHost) DataMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateServiceHost) DataVolumeDeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataVolumeDeviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateServiceHost) ElasticIp() awsec2.CfnEIP {
	var returns awsec2.CfnEIP
	_jsii_.Get(
		j,
		"elasticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateServiceHost) Instance() awsec2.Instance {
	var returns awsec2.Instance
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateServiceHost) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateServiceHost) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateServiceHost) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateServiceHost) ServiceIdentity() *ResolvedPlatformServiceIdentity {
	var returns *ResolvedPlatformServiceIdentity
	_jsii_.Get(
		j,
		"serviceIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateServiceHost) ServiceOutputs() *NetworkAddressableServiceOutputs {
	var returns *NetworkAddressableServiceOutputs
	_jsii_.Get(
		j,
		"serviceOutputs",
		&returns,
	)
	return returns
}


func NewPrivateServiceHost(scope constructs.Construct, id *string, props *PrivateServiceHostProps) PrivateServiceHost {
	_init_.Initialize()

	if err := validateNewPrivateServiceHostParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivateServiceHost{}

	_jsii_.Create(
		"cdk-service-host-module.PrivateServiceHost",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPrivateServiceHost_Override(p PrivateServiceHost, scope constructs.Construct, id *string, props *PrivateServiceHostProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-service-host-module.PrivateServiceHost",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func PrivateServiceHost_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivateServiceHost_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-service-host-module.PrivateServiceHost",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateServiceHost) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

