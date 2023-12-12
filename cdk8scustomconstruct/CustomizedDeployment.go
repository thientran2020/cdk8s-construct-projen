package cdk8scustomconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/thientran2020/cdk8s-construct-projen/cdk8scustomconstruct/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/thientran2020/cdk8s-construct-projen/cdk8scustomconstruct/internal"
)

type CustomizedDeployment interface {
	constructs.Construct
	Name() *string
	Namespace() *string
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

// The jsii proxy struct for CustomizedDeployment
type jsiiProxy_CustomizedDeployment struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CustomizedDeployment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedDeployment) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}


func NewCustomizedDeployment(scope constructs.Construct, name *string, selector *map[string]*string, opts *CustomizedDeploymentOptions) CustomizedDeployment {
	_init_.Initialize()

	if err := validateNewCustomizedDeploymentParameters(scope, name, selector, opts); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomizedDeployment{}

	_jsii_.Create(
		"cdk8s-custom-construct.CustomizedDeployment",
		[]interface{}{scope, name, selector, opts},
		&j,
	)

	return &j
}

func NewCustomizedDeployment_Override(c CustomizedDeployment, scope constructs.Construct, name *string, selector *map[string]*string, opts *CustomizedDeploymentOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-custom-construct.CustomizedDeployment",
		[]interface{}{scope, name, selector, opts},
		c,
	)
}

func (c *jsiiProxy_CustomizedDeployment) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedDeployment) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CustomizedDeployment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

