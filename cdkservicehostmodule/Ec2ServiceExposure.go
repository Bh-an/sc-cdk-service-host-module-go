package cdkservicehostmodule


type Ec2ServiceExposure struct {
	AssociatePublicIpAddress *bool `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	EnableElasticIp *bool `field:"optional" json:"enableElasticIp" yaml:"enableElasticIp"`
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

