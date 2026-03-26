package cdkservicehostmodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type PlatformServiceOutputs struct {
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	SecurityGroup awsec2.ISecurityGroup `field:"required" json:"securityGroup" yaml:"securityGroup"`
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	Tags *map[string]*string `field:"required" json:"tags" yaml:"tags"`
}

