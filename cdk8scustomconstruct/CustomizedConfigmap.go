package cdk8scustomconstruct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/thientran2020/cdk8s-construct-projen/cdk8scustomconstruct/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/thientran2020/cdk8s-construct-projen/cdk8scustomconstruct/internal"
)

type CustomizedConfigmap interface {
	constructs.Construct
	Name() *string
	Namespace() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CustomizedConfigmap
type jsiiProxy_CustomizedConfigmap struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CustomizedConfigmap) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedConfigmap) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedConfigmap) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCustomizedConfigmap(scope constructs.Construct, name *string, opts *CustomizedConfigmapOptions) CustomizedConfigmap {
	_init_.Initialize()

	if err := validateNewCustomizedConfigmapParameters(scope, name, opts); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomizedConfigmap{}

	_jsii_.Create(
		"cdk8s-custom-construct.CustomizedConfigmap",
		[]interface{}{scope, name, opts},
		&j,
	)

	return &j
}

func NewCustomizedConfigmap_Override(c CustomizedConfigmap, scope constructs.Construct, name *string, opts *CustomizedConfigmapOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-custom-construct.CustomizedConfigmap",
		[]interface{}{scope, name, opts},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CustomizedConfigmap_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomizedConfigmap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-custom-construct.CustomizedConfigmap",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedConfigmap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

