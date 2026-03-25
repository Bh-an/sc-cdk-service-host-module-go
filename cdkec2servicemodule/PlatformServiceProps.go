package cdkec2servicemodule


type PlatformServiceProps struct {
	Infrastructure *ServiceInfrastructureProps `field:"required" json:"infrastructure" yaml:"infrastructure"`
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	AdditionalTags *map[string]*string `field:"optional" json:"additionalTags" yaml:"additionalTags"`
	Identity *PlatformServiceIdentity `field:"optional" json:"identity" yaml:"identity"`
}

