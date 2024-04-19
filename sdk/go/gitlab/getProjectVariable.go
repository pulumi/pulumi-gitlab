// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ProjectVariable` data source allows to retrieve details about a project-level CI/CD variable.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := gitlab.LookupProjectVariable(ctx, &gitlab.LookupProjectVariableArgs{
//				Project: "my/example/project",
//				Key:     "foo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Using an environment scope
//			_, err = gitlab.LookupProjectVariable(ctx, &gitlab.LookupProjectVariableArgs{
//				Project:          "my/example/project",
//				Key:              "bar",
//				EnvironmentScope: pulumi.StringRef("staging/*"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupProjectVariable(ctx *pulumi.Context, args *LookupProjectVariableArgs, opts ...pulumi.InvokeOption) (*LookupProjectVariableResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectVariableResult
	err := ctx.Invoke("gitlab:index/getProjectVariable:getProjectVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectVariable.
type LookupProjectVariableArgs struct {
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The name of the variable.
	Key string `pulumi:"key"`
	// The name or id of the project.
	Project string `pulumi:"project"`
}

// A collection of values returned by getProjectVariable.
type LookupProjectVariableResult struct {
	// The description of the variable.
	Description string `pulumi:"description"`
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope string `pulumi:"environmentScope"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the variable.
	Key string `pulumi:"key"`
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked bool `pulumi:"masked"`
	// The name or id of the project.
	Project string `pulumi:"project"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected bool `pulumi:"protected"`
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw bool `pulumi:"raw"`
	// The value of the variable.
	Value string `pulumi:"value"`
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType string `pulumi:"variableType"`
}

func LookupProjectVariableOutput(ctx *pulumi.Context, args LookupProjectVariableOutputArgs, opts ...pulumi.InvokeOption) LookupProjectVariableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectVariableResult, error) {
			args := v.(LookupProjectVariableArgs)
			r, err := LookupProjectVariable(ctx, &args, opts...)
			var s LookupProjectVariableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectVariableResultOutput)
}

// A collection of arguments for invoking getProjectVariable.
type LookupProjectVariableOutputArgs struct {
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope pulumi.StringPtrInput `pulumi:"environmentScope"`
	// The name of the variable.
	Key pulumi.StringInput `pulumi:"key"`
	// The name or id of the project.
	Project pulumi.StringInput `pulumi:"project"`
}

func (LookupProjectVariableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectVariableArgs)(nil)).Elem()
}

// A collection of values returned by getProjectVariable.
type LookupProjectVariableResultOutput struct{ *pulumi.OutputState }

func (LookupProjectVariableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectVariableResult)(nil)).Elem()
}

func (o LookupProjectVariableResultOutput) ToLookupProjectVariableResultOutput() LookupProjectVariableResultOutput {
	return o
}

func (o LookupProjectVariableResultOutput) ToLookupProjectVariableResultOutputWithContext(ctx context.Context) LookupProjectVariableResultOutput {
	return o
}

// The description of the variable.
func (o LookupProjectVariableResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectVariableResult) string { return v.Description }).(pulumi.StringOutput)
}

// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
func (o LookupProjectVariableResultOutput) EnvironmentScope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectVariableResult) string { return v.EnvironmentScope }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectVariableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectVariableResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the variable.
func (o LookupProjectVariableResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectVariableResult) string { return v.Key }).(pulumi.StringOutput)
}

// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
func (o LookupProjectVariableResultOutput) Masked() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectVariableResult) bool { return v.Masked }).(pulumi.BoolOutput)
}

// The name or id of the project.
func (o LookupProjectVariableResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectVariableResult) string { return v.Project }).(pulumi.StringOutput)
}

// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
func (o LookupProjectVariableResultOutput) Protected() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectVariableResult) bool { return v.Protected }).(pulumi.BoolOutput)
}

// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
func (o LookupProjectVariableResultOutput) Raw() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectVariableResult) bool { return v.Raw }).(pulumi.BoolOutput)
}

// The value of the variable.
func (o LookupProjectVariableResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectVariableResult) string { return v.Value }).(pulumi.StringOutput)
}

// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
func (o LookupProjectVariableResultOutput) VariableType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectVariableResult) string { return v.VariableType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectVariableResultOutput{})
}
