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

type PublicServiceHost interface {
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

// The jsii proxy struct for PublicServiceHost
type jsiiProxy_PublicServiceHost struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PublicServiceHost) DataKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"dataKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicServiceHost) DataMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicServiceHost) DataVolumeDeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataVolumeDeviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicServiceHost) ElasticIp() awsec2.CfnEIP {
	var returns awsec2.CfnEIP
	_jsii_.Get(
		j,
		"elasticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicServiceHost) Instance() awsec2.Instance {
	var returns awsec2.Instance
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicServiceHost) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicServiceHost) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicServiceHost) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicServiceHost) ServiceIdentity() *ResolvedPlatformServiceIdentity {
	var returns *ResolvedPlatformServiceIdentity
	_jsii_.Get(
		j,
		"serviceIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicServiceHost) ServiceOutputs() *NetworkAddressableServiceOutputs {
	var returns *NetworkAddressableServiceOutputs
	_jsii_.Get(
		j,
		"serviceOutputs",
		&returns,
	)
	return returns
}


func NewPublicServiceHost(scope constructs.Construct, id *string, props *PublicServiceHostProps) PublicServiceHost {
	_init_.Initialize()

	if err := validateNewPublicServiceHostParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PublicServiceHost{}

	_jsii_.Create(
		"cdk-service-host-module.PublicServiceHost",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPublicServiceHost_Override(p PublicServiceHost, scope constructs.Construct, id *string, props *PublicServiceHostProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-service-host-module.PublicServiceHost",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func PublicServiceHost_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePublicServiceHost_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-service-host-module.PublicServiceHost",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicServiceHost) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

