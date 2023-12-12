// cdk8s-custom-construct
package cdk8scustomconstruct

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk8s-custom-construct.CustomizedApp",
		reflect.TypeOf((*CustomizedApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustomizedApp{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-custom-construct.CustomizedConfigmap",
		reflect.TypeOf((*CustomizedConfigmap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustomizedConfigmap{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-custom-construct.CustomizedConfigmapOptions",
		reflect.TypeOf((*CustomizedConfigmapOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-custom-construct.CustomizedDeployment",
		reflect.TypeOf((*CustomizedDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustomizedDeployment{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-custom-construct.CustomizedDeploymentOptions",
		reflect.TypeOf((*CustomizedDeploymentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-custom-construct.CustomizedOptions",
		reflect.TypeOf((*CustomizedOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-custom-construct.Hello",
		reflect.TypeOf((*Hello)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "sayHello", GoMethod: "SayHello"},
		},
		func() interface{} {
			return &jsiiProxy_Hello{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-custom-construct.ResourceQuantity",
		reflect.TypeOf((*ResourceQuantity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-custom-construct.ResourceRequirements",
		reflect.TypeOf((*ResourceRequirements)(nil)).Elem(),
	)
}
