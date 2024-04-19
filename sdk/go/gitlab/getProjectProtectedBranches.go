// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getProjectProtectedBranches` data source allows details of the protected branches of a given project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_branches.html#list-protected-branches)
func GetProjectProtectedBranches(ctx *pulumi.Context, args *GetProjectProtectedBranchesArgs, opts ...pulumi.InvokeOption) (*GetProjectProtectedBranchesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectProtectedBranchesResult
	err := ctx.Invoke("gitlab:index/getProjectProtectedBranches:getProjectProtectedBranches", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectProtectedBranches.
type GetProjectProtectedBranchesArgs struct {
	// The integer or path with namespace that uniquely identifies the project.
	ProjectId string `pulumi:"projectId"`
	// A list of protected branches, as defined below.
	ProtectedBranches []GetProjectProtectedBranchesProtectedBranch `pulumi:"protectedBranches"`
}

// A collection of values returned by getProjectProtectedBranches.
type GetProjectProtectedBranchesResult struct {
	// The ID of this resource.
	Id int `pulumi:"id"`
	// The integer or path with namespace that uniquely identifies the project.
	ProjectId string `pulumi:"projectId"`
	// A list of protected branches, as defined below.
	ProtectedBranches []GetProjectProtectedBranchesProtectedBranch `pulumi:"protectedBranches"`
}

func GetProjectProtectedBranchesOutput(ctx *pulumi.Context, args GetProjectProtectedBranchesOutputArgs, opts ...pulumi.InvokeOption) GetProjectProtectedBranchesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectProtectedBranchesResult, error) {
			args := v.(GetProjectProtectedBranchesArgs)
			r, err := GetProjectProtectedBranches(ctx, &args, opts...)
			var s GetProjectProtectedBranchesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectProtectedBranchesResultOutput)
}

// A collection of arguments for invoking getProjectProtectedBranches.
type GetProjectProtectedBranchesOutputArgs struct {
	// The integer or path with namespace that uniquely identifies the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// A list of protected branches, as defined below.
	ProtectedBranches GetProjectProtectedBranchesProtectedBranchArrayInput `pulumi:"protectedBranches"`
}

func (GetProjectProtectedBranchesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectProtectedBranchesArgs)(nil)).Elem()
}

// A collection of values returned by getProjectProtectedBranches.
type GetProjectProtectedBranchesResultOutput struct{ *pulumi.OutputState }

func (GetProjectProtectedBranchesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectProtectedBranchesResult)(nil)).Elem()
}

func (o GetProjectProtectedBranchesResultOutput) ToGetProjectProtectedBranchesResultOutput() GetProjectProtectedBranchesResultOutput {
	return o
}

func (o GetProjectProtectedBranchesResultOutput) ToGetProjectProtectedBranchesResultOutputWithContext(ctx context.Context) GetProjectProtectedBranchesResultOutput {
	return o
}

// The ID of this resource.
func (o GetProjectProtectedBranchesResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetProjectProtectedBranchesResult) int { return v.Id }).(pulumi.IntOutput)
}

// The integer or path with namespace that uniquely identifies the project.
func (o GetProjectProtectedBranchesResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectProtectedBranchesResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// A list of protected branches, as defined below.
func (o GetProjectProtectedBranchesResultOutput) ProtectedBranches() GetProjectProtectedBranchesProtectedBranchArrayOutput {
	return o.ApplyT(func(v GetProjectProtectedBranchesResult) []GetProjectProtectedBranchesProtectedBranch {
		return v.ProtectedBranches
	}).(GetProjectProtectedBranchesProtectedBranchArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectProtectedBranchesResultOutput{})
}
