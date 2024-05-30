// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getProjectHooks` data source allows to retrieve details about hooks in a project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#list-project-hooks)
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
//			example, err := gitlab.LookupProject(ctx, &gitlab.LookupProjectArgs{
//				Id: pulumi.StringRef("foo/bar/baz"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.GetProjectHooks(ctx, &gitlab.GetProjectHooksArgs{
//				Project: example.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProjectHooks(ctx *pulumi.Context, args *GetProjectHooksArgs, opts ...pulumi.InvokeOption) (*GetProjectHooksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectHooksResult
	err := ctx.Invoke("gitlab:index/getProjectHooks:getProjectHooks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectHooks.
type GetProjectHooksArgs struct {
	// The name or id of the project.
	Project string `pulumi:"project"`
}

// A collection of values returned by getProjectHooks.
type GetProjectHooksResult struct {
	// The list of hooks.
	Hooks []GetProjectHooksHook `pulumi:"hooks"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name or id of the project.
	Project string `pulumi:"project"`
}

func GetProjectHooksOutput(ctx *pulumi.Context, args GetProjectHooksOutputArgs, opts ...pulumi.InvokeOption) GetProjectHooksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectHooksResult, error) {
			args := v.(GetProjectHooksArgs)
			r, err := GetProjectHooks(ctx, &args, opts...)
			var s GetProjectHooksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectHooksResultOutput)
}

// A collection of arguments for invoking getProjectHooks.
type GetProjectHooksOutputArgs struct {
	// The name or id of the project.
	Project pulumi.StringInput `pulumi:"project"`
}

func (GetProjectHooksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectHooksArgs)(nil)).Elem()
}

// A collection of values returned by getProjectHooks.
type GetProjectHooksResultOutput struct{ *pulumi.OutputState }

func (GetProjectHooksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectHooksResult)(nil)).Elem()
}

func (o GetProjectHooksResultOutput) ToGetProjectHooksResultOutput() GetProjectHooksResultOutput {
	return o
}

func (o GetProjectHooksResultOutput) ToGetProjectHooksResultOutputWithContext(ctx context.Context) GetProjectHooksResultOutput {
	return o
}

// The list of hooks.
func (o GetProjectHooksResultOutput) Hooks() GetProjectHooksHookArrayOutput {
	return o.ApplyT(func(v GetProjectHooksResult) []GetProjectHooksHook { return v.Hooks }).(GetProjectHooksHookArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectHooksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectHooksResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name or id of the project.
func (o GetProjectHooksResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectHooksResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectHooksResultOutput{})
}
