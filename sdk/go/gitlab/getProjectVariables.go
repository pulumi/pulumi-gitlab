// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getProjectVariables` data source allows to retrieve all project-level CI/CD variables.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.GetProjectVariables(ctx, &gitlab.GetProjectVariablesArgs{
//				Project: "my/example/project",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Using an environment scope
//			_, err = gitlab.GetProjectVariables(ctx, &gitlab.GetProjectVariablesArgs{
//				Project:          "my/example/project",
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
func GetProjectVariables(ctx *pulumi.Context, args *GetProjectVariablesArgs, opts ...pulumi.InvokeOption) (*GetProjectVariablesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectVariablesResult
	err := ctx.Invoke("gitlab:index/getProjectVariables:getProjectVariables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectVariables.
type GetProjectVariablesArgs struct {
	// The environment scope of the variable. Defaults to all environment (`*`).
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The name or id of the project.
	Project string `pulumi:"project"`
}

// A collection of values returned by getProjectVariables.
type GetProjectVariablesResult struct {
	// The environment scope of the variable. Defaults to all environment (`*`).
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name or id of the project.
	Project string `pulumi:"project"`
	// The list of variables returned by the search
	Variables []GetProjectVariablesVariable `pulumi:"variables"`
}

func GetProjectVariablesOutput(ctx *pulumi.Context, args GetProjectVariablesOutputArgs, opts ...pulumi.InvokeOption) GetProjectVariablesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectVariablesResultOutput, error) {
			args := v.(GetProjectVariablesArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetProjectVariablesResult
			secret, err := ctx.InvokePackageRaw("gitlab:index/getProjectVariables:getProjectVariables", args, &rv, "", opts...)
			if err != nil {
				return GetProjectVariablesResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetProjectVariablesResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetProjectVariablesResultOutput), nil
			}
			return output, nil
		}).(GetProjectVariablesResultOutput)
}

// A collection of arguments for invoking getProjectVariables.
type GetProjectVariablesOutputArgs struct {
	// The environment scope of the variable. Defaults to all environment (`*`).
	EnvironmentScope pulumi.StringPtrInput `pulumi:"environmentScope"`
	// The name or id of the project.
	Project pulumi.StringInput `pulumi:"project"`
}

func (GetProjectVariablesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectVariablesArgs)(nil)).Elem()
}

// A collection of values returned by getProjectVariables.
type GetProjectVariablesResultOutput struct{ *pulumi.OutputState }

func (GetProjectVariablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectVariablesResult)(nil)).Elem()
}

func (o GetProjectVariablesResultOutput) ToGetProjectVariablesResultOutput() GetProjectVariablesResultOutput {
	return o
}

func (o GetProjectVariablesResultOutput) ToGetProjectVariablesResultOutputWithContext(ctx context.Context) GetProjectVariablesResultOutput {
	return o
}

// The environment scope of the variable. Defaults to all environment (`*`).
func (o GetProjectVariablesResultOutput) EnvironmentScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectVariablesResult) *string { return v.EnvironmentScope }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectVariablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectVariablesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name or id of the project.
func (o GetProjectVariablesResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectVariablesResult) string { return v.Project }).(pulumi.StringOutput)
}

// The list of variables returned by the search
func (o GetProjectVariablesResultOutput) Variables() GetProjectVariablesVariableArrayOutput {
	return o.ApplyT(func(v GetProjectVariablesResult) []GetProjectVariablesVariable { return v.Variables }).(GetProjectVariablesVariableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectVariablesResultOutput{})
}
