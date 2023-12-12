//go:build no_runtime_type_checking

package cdk8scustomconstruct

// Building without runtime type checking enabled, so all the below just return nil

func validateCustomizedDeployment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCustomizedDeploymentParameters(scope constructs.Construct, name *string, selector *map[string]*string, opts *CustomizedDeploymentOptions) error {
	return nil
}

