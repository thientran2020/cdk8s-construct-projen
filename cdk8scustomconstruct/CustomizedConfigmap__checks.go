//go:build !no_runtime_type_checking

package cdk8scustomconstruct

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

func (c *jsiiProxy_CustomizedConfigmap) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateNewCustomizedConfigmapParameters(scope constructs.Construct, name *string, opts *CustomizedConfigmapOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if opts == nil {
		return fmt.Errorf("parameter opts is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

