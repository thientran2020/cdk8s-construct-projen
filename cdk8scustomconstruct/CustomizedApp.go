package cdk8scustomconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/thientran2020/cdk8s-construct-projen/cdk8scustomconstruct/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/thientran2020/cdk8s-construct-projen/cdk8scustomconstruct/internal"
)

type CustomizedApp interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CustomizedApp
type jsiiProxy_CustomizedApp struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CustomizedApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CustomizedApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomizedApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-custom-construct.CustomizedApp",
		"isConstruct",
		[]interface{}{x},
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

