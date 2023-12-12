package cdk8scustomconstruct


type CustomizedDeploymentOptions struct {
	Annotations *map[string]*string `field:"required" json:"annotations" yaml:"annotations"`
	Image *string `field:"required" json:"image" yaml:"image"`
	Labels *map[string]*string `field:"required" json:"labels" yaml:"labels"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	Resources *ResourceRequirements `field:"optional" json:"resources" yaml:"resources"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
}

