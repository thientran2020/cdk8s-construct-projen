package cdk8scustomconstruct


type ResourceQuantity struct {
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

