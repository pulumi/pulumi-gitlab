// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GroupVariable` data source allows to retrieve details about a group-level CI/CD variable.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_level_variables.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.LookupGroupVariable(ctx, &gitlab.LookupGroupVariableArgs{
//				Group: "my/example/group",
//				Key:   "foo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.LookupGroupVariable(ctx, &gitlab.LookupGroupVariableArgs{
//				EnvironmentScope: pulumi.StringRef("staging/*"),
//				Group:            "my/example/group",
//				Key:              "bar",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupGroupVariable(ctx *pulumi.Context, args *LookupGroupVariableArgs, opts ...pulumi.InvokeOption) (*LookupGroupVariableResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupVariableResult
	err := ctx.Invoke("gitlab:index/getGroupVariable:getGroupVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupVariable.
type LookupGroupVariableArgs struct {
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The name or id of the group.
	Group string `pulumi:"group"`
	// The name of the variable.
	Key string `pulumi:"key"`
}

// A collection of values returned by getGroupVariable.
type LookupGroupVariableResult struct {
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope string `pulumi:"environmentScope"`
	// The name or id of the group.
	Group string `pulumi:"group"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the variable.
	Key string `pulumi:"key"`
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked bool `pulumi:"masked"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected bool `pulumi:"protected"`
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw bool `pulumi:"raw"`
	// The value of the variable.
	Value string `pulumi:"value"`
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType string `pulumi:"variableType"`
}

func LookupGroupVariableOutput(ctx *pulumi.Context, args LookupGroupVariableOutputArgs, opts ...pulumi.InvokeOption) LookupGroupVariableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupVariableResult, error) {
			args := v.(LookupGroupVariableArgs)
			r, err := LookupGroupVariable(ctx, &args, opts...)
			var s LookupGroupVariableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupVariableResultOutput)
}

// A collection of arguments for invoking getGroupVariable.
type LookupGroupVariableOutputArgs struct {
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope pulumi.StringPtrInput `pulumi:"environmentScope"`
	// The name or id of the group.
	Group pulumi.StringInput `pulumi:"group"`
	// The name of the variable.
	Key pulumi.StringInput `pulumi:"key"`
}

func (LookupGroupVariableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupVariableArgs)(nil)).Elem()
}

// A collection of values returned by getGroupVariable.
type LookupGroupVariableResultOutput struct{ *pulumi.OutputState }

func (LookupGroupVariableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupVariableResult)(nil)).Elem()
}

func (o LookupGroupVariableResultOutput) ToLookupGroupVariableResultOutput() LookupGroupVariableResultOutput {
	return o
}

func (o LookupGroupVariableResultOutput) ToLookupGroupVariableResultOutputWithContext(ctx context.Context) LookupGroupVariableResultOutput {
	return o
}

// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
func (o LookupGroupVariableResultOutput) EnvironmentScope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupVariableResult) string { return v.EnvironmentScope }).(pulumi.StringOutput)
}

// The name or id of the group.
func (o LookupGroupVariableResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupVariableResult) string { return v.Group }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupVariableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupVariableResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the variable.
func (o LookupGroupVariableResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupVariableResult) string { return v.Key }).(pulumi.StringOutput)
}

// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
func (o LookupGroupVariableResultOutput) Masked() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupVariableResult) bool { return v.Masked }).(pulumi.BoolOutput)
}

// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
func (o LookupGroupVariableResultOutput) Protected() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupVariableResult) bool { return v.Protected }).(pulumi.BoolOutput)
}

// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
func (o LookupGroupVariableResultOutput) Raw() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupVariableResult) bool { return v.Raw }).(pulumi.BoolOutput)
}

// The value of the variable.
func (o LookupGroupVariableResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupVariableResult) string { return v.Value }).(pulumi.StringOutput)
}

// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
func (o LookupGroupVariableResultOutput) VariableType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupVariableResult) string { return v.VariableType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupVariableResultOutput{})
}
