package cdkservicehostmodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type NetworkAddressableServiceOutputs struct {
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	SecurityGroup awsec2.ISecurityGroup `field:"required" json:"securityGroup" yaml:"securityGroup"`
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	Tags *map[string]*string `field:"required" json:"tags" yaml:"tags"`
	ExposureKind *string `field:"required" json:"exposureKind" yaml:"exposureKind"`
	HasPublicEndpoint *bool `field:"required" json:"hasPublicEndpoint" yaml:"hasPublicEndpoint"`
	ListenerPort *float64 `field:"required" json:"listenerPort" yaml:"listenerPort"`
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

