package cdkservicehostmodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type ServiceHostRuntimeProps struct {
	DockerImage *string `field:"required" json:"dockerImage" yaml:"dockerImage"`
	AllowedIngress *[]*IngressRule `field:"optional" json:"allowedIngress" yaml:"allowedIngress"`
	BridgeCidr *string `field:"optional" json:"bridgeCidr" yaml:"bridgeCidr"`
	BridgeIp *string `field:"optional" json:"bridgeIp" yaml:"bridgeIp"`
	BridgeNetworkName *string `field:"optional" json:"bridgeNetworkName" yaml:"bridgeNetworkName"`
	DataMountPath *string `field:"optional" json:"dataMountPath" yaml:"dataMountPath"`
	DataVolumeDeviceName *string `field:"optional" json:"dataVolumeDeviceName" yaml:"dataVolumeDeviceName"`
	DataVolumeSizeGiB *float64 `field:"optional" json:"dataVolumeSizeGiB" yaml:"dataVolumeSizeGiB"`
	Exposure *ServiceHostExposure `field:"optional" json:"exposure" yaml:"exposure"`
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	MachineImage awsec2.IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
	NginxMainConfig *string `field:"optional" json:"nginxMainConfig" yaml:"nginxMainConfig"`
	NginxRoutesConfig *string `field:"optional" json:"nginxRoutesConfig" yaml:"nginxRoutesConfig"`
	Operations *ServiceHostOperationalControls `field:"optional" json:"operations" yaml:"operations"`
	PublicPort *float64 `field:"optional" json:"publicPort" yaml:"publicPort"`
	RootVolumeSizeGiB *float64 `field:"optional" json:"rootVolumeSizeGiB" yaml:"rootVolumeSizeGiB"`
	ServicePort *float64 `field:"optional" json:"servicePort" yaml:"servicePort"`
}

