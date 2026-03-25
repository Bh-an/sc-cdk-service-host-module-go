package cdkec2servicemodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/Bh-an/sc-cdk-ec2-service-module-go/cdkec2servicemodule/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/Bh-an/sc-cdk-ec2-service-module-go/cdkec2servicemodule/internal"
)

type PrivateEc2DockerService interface {
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

// The jsii proxy struct for PrivateEc2DockerService
type jsiiProxy_PrivateEc2DockerService struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PrivateEc2DockerService) DataKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"dataKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEc2DockerService) DataMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEc2DockerService) DataVolumeDeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataVolumeDeviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEc2DockerService) ElasticIp() awsec2.CfnEIP {
	var returns awsec2.CfnEIP
	_jsii_.Get(
		j,
		"elasticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEc2DockerService) Instance() awsec2.Instance {
	var returns awsec2.Instance
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEc2DockerService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEc2DockerService) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEc2DockerService) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEc2DockerService) ServiceIdentity() *ResolvedPlatformServiceIdentity {
	var returns *ResolvedPlatformServiceIdentity
	_jsii_.Get(
		j,
		"serviceIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateEc2DockerService) ServiceOutputs() *NetworkAddressableServiceOutputs {
	var returns *NetworkAddressableServiceOutputs
	_jsii_.Get(
		j,
		"serviceOutputs",
		&returns,
	)
	return returns
}


func NewPrivateEc2DockerService(scope constructs.Construct, id *string, props *PrivateEc2DockerServiceProps) PrivateEc2DockerService {
	_init_.Initialize()

	if err := validateNewPrivateEc2DockerServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivateEc2DockerService{}

	_jsii_.Create(
		"cdk-ec2-service-module.PrivateEc2DockerService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPrivateEc2DockerService_Override(p PrivateEc2DockerService, scope constructs.Construct, id *string, props *PrivateEc2DockerServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-ec2-service-module.PrivateEc2DockerService",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func PrivateEc2DockerService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivateEc2DockerService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-ec2-service-module.PrivateEc2DockerService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateEc2DockerService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

