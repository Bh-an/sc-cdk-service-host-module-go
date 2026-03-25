package cdkec2servicemodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type IngressRule struct {
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	SourceSecurityGroup awsec2.ISecurityGroup `field:"optional" json:"sourceSecurityGroup" yaml:"sourceSecurityGroup"`
}

