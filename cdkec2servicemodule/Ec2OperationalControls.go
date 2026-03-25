package cdkec2servicemodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type Ec2OperationalControls struct {
	AdditionalManagedPolicies *[]awsiam.IManagedPolicy `field:"optional" json:"additionalManagedPolicies" yaml:"additionalManagedPolicies"`
	AdditionalRolePolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"additionalRolePolicyStatements" yaml:"additionalRolePolicyStatements"`
	EnableDetailedMonitoring *bool `field:"optional" json:"enableDetailedMonitoring" yaml:"enableDetailedMonitoring"`
	PostBootstrapCommands *[]*string `field:"optional" json:"postBootstrapCommands" yaml:"postBootstrapCommands"`
	PreBootstrapCommands *[]*string `field:"optional" json:"preBootstrapCommands" yaml:"preBootstrapCommands"`
	SkipSystemPackagesUpdate *bool `field:"optional" json:"skipSystemPackagesUpdate" yaml:"skipSystemPackagesUpdate"`
}

