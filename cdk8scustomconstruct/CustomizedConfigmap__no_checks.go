//go:build no_runtime_type_checking

package cdk8scustomconstruct

// Building without runtime type checking enabled, so all the below just return nil

func validateCustomizedConfigmap_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCustomizedConfigmapParameters(scope constructs.Construct, name *string, opts *CustomizedConfigmapOptions) error {
	return nil
}

