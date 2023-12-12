package cdk8scustomconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/thientran2020/cdk8s-construct-projen/cdk8scustomconstruct/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/thientran2020/cdk8s-construct-projen/cdk8scustomconstruct/internal"
)

type CustomizedApp interface {
	constructs.Construct
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CustomizedApp
type jsiiProxy_CustomizedApp struct {
	internal.Type__constructsConstruct
}

func NewCustomizedApp(scope constructs.Construct, name *string, opts *CustomizedOptions) CustomizedApp {
	_init_.Initialize()

	if err := validateNewCustomizedAppParameters(scope, name, opts); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomizedApp{}

	_jsii_.Create(
		"cdk8s-custom-construct.CustomizedApp",
		[]interface{}{scope, name, opts},
		&j,
	)

	return &j
}

func NewCustomizedApp_Override(c CustomizedApp, scope constructs.Construct, name *string, opts *CustomizedOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-custom-construct.CustomizedApp",
		[]interface{}{scope, name, opts},
		c,
	)
}

func (c *jsiiProxy_CustomizedApp) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedApp) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CustomizedApp) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

