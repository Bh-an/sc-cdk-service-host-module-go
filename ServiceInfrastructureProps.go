package cdkec2servicemodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

type ServiceInfrastructureProps struct {
	SubnetSelection *awsec2.SubnetSelection `field:"required" json:"subnetSelection" yaml:"subnetSelection"`
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	KeyPair awsec2.IKeyPair `field:"optional" json:"keyPair" yaml:"keyPair"`
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	SharedTags *map[string]*string `field:"optional" json:"sharedTags" yaml:"sharedTags"`
}

