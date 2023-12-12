//go:build no_runtime_type_checking

package cdk8scustomconstruct

// Building without runtime type checking enabled, so all the below just return nil

func validateCustomizedApp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCustomizedAppParameters(scope constructs.Construct, name *string, opts *CustomizedOptions) error {
	return nil
}

