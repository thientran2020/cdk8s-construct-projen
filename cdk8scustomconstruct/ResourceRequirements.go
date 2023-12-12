package cdk8scustomconstruct


type ResourceRequirements struct {
	Limits *ResourceQuantity `field:"optional" json:"limits" yaml:"limits"`
	Requests *ResourceQuantity `field:"optional" json:"requests" yaml:"requests"`
}

