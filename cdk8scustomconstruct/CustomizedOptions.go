package cdk8scustomconstruct


type CustomizedOptions struct {
	ConfigmapOptions *CustomizedConfigmapOptions `field:"required" json:"configmapOptions" yaml:"configmapOptions"`
	DeploymentOptions *CustomizedDeploymentOptions `field:"required" json:"deploymentOptions" yaml:"deploymentOptions"`
}

