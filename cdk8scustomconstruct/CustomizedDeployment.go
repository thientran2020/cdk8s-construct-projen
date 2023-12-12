package cdk8scustomconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/thientran2020/cdk8s-construct-projen/cdk8scustomconstruct/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/thientran2020/cdk8s-construct-projen/cdk8scustomconstruct/internal"
)

type CustomizedDeployment interface {
	constructs.Construct
	Name() *string
	Namespace() *string
	// The tree node.
	Node() constructs.Node
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

func (j *jsiiProxy_CustomizedDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CustomizedDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomizedDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-custom-construct.CustomizedDeployment",
		"isConstruct",
		[]interface{}{x},
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

