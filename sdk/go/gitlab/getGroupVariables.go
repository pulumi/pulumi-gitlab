// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getGroupVariables` data source allows to retrieve all group-level CI/CD variables.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_level_variables.html)
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
//			_, err := gitlab.GetGroupVariables(ctx, &gitlab.GetGroupVariablesArgs{
//				Group: "my/example/group",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.GetGroupVariables(ctx, &gitlab.GetGroupVariablesArgs{
//				EnvironmentScope: pulumi.StringRef("staging/*"),
//				Group:            "my/example/group",
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
func GetGroupVariables(ctx *pulumi.Context, args *GetGroupVariablesArgs, opts ...pulumi.InvokeOption) (*GetGroupVariablesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGroupVariablesResult
	err := ctx.Invoke("gitlab:index/getGroupVariables:getGroupVariables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupVariables.
type GetGroupVariablesArgs struct {
	EnvironmentScope *string `pulumi:"environmentScope"`
	Group            string  `pulumi:"group"`
}

// A collection of values returned by getGroupVariables.
type GetGroupVariablesResult struct {
	// The environment scope of the variable. Defaults to all environment (`*`).
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The name or id of the group.
	Group string `pulumi:"group"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of variables returned by the search
	Variables []GetGroupVariablesVariable `pulumi:"variables"`
}

func GetGroupVariablesOutput(ctx *pulumi.Context, args GetGroupVariablesOutputArgs, opts ...pulumi.InvokeOption) GetGroupVariablesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupVariablesResult, error) {
			args := v.(GetGroupVariablesArgs)
			r, err := GetGroupVariables(ctx, &args, opts...)
			var s GetGroupVariablesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupVariablesResultOutput)
}

// A collection of arguments for invoking getGroupVariables.
type GetGroupVariablesOutputArgs struct {
	EnvironmentScope pulumi.StringPtrInput `pulumi:"environmentScope"`
	Group            pulumi.StringInput    `pulumi:"group"`
}

func (GetGroupVariablesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupVariablesArgs)(nil)).Elem()
}

// A collection of values returned by getGroupVariables.
type GetGroupVariablesResultOutput struct{ *pulumi.OutputState }

func (GetGroupVariablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupVariablesResult)(nil)).Elem()
}

func (o GetGroupVariablesResultOutput) ToGetGroupVariablesResultOutput() GetGroupVariablesResultOutput {
	return o
}

func (o GetGroupVariablesResultOutput) ToGetGroupVariablesResultOutputWithContext(ctx context.Context) GetGroupVariablesResultOutput {
	return o
}

// The environment scope of the variable. Defaults to all environment (`*`).
func (o GetGroupVariablesResultOutput) EnvironmentScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupVariablesResult) *string { return v.EnvironmentScope }).(pulumi.StringPtrOutput)
}

// The name or id of the group.
func (o GetGroupVariablesResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupVariablesResult) string { return v.Group }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupVariablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupVariablesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of variables returned by the search
func (o GetGroupVariablesResultOutput) Variables() GetGroupVariablesVariableArrayOutput {
	return o.ApplyT(func(v GetGroupVariablesResult) []GetGroupVariablesVariable { return v.Variables }).(GetGroupVariablesVariableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupVariablesResultOutput{})
}
