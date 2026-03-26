package cdkservicehostmodule


type ResolvedPlatformServiceIdentity struct {
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	ResourcePrefix *string `field:"required" json:"resourcePrefix" yaml:"resourcePrefix"`
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	Tags *map[string]*string `field:"required" json:"tags" yaml:"tags"`
}

