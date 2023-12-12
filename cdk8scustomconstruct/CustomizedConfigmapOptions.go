package cdk8scustomconstruct


type CustomizedConfigmapOptions struct {
	Data *map[string]*string `field:"required" json:"data" yaml:"data"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

