// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getProjectProtectedBranch` data source allows details of a protected branch to be retrieved by its name and the project it belongs to.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/protected_branches/#get-a-single-protected-branch-or-wildcard-protected-branch)
func GetProjectProtectedBranch(ctx *pulumi.Context, args *GetProjectProtectedBranchArgs, opts ...pulumi.InvokeOption) (*GetProjectProtectedBranchResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectProtectedBranchResult
	err := ctx.Invoke("gitlab:index/getProjectProtectedBranch:getProjectProtectedBranch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectProtectedBranch.
type GetProjectProtectedBranchArgs struct {
	// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
	MergeAccessLevels []GetProjectProtectedBranchMergeAccessLevel `pulumi:"mergeAccessLevels"`
	// The name of the protected branch.
	Name string `pulumi:"name"`
	// The integer or path with namespace that uniquely identifies the project.
	ProjectId string `pulumi:"projectId"`
	// Array of access levels and user(s)/group(s) allowed to push to protected branch.
	PushAccessLevels []GetProjectProtectedBranchPushAccessLevel `pulumi:"pushAccessLevels"`
}

// A collection of values returned by getProjectProtectedBranch.
type GetProjectProtectedBranchResult struct {
	// Whether force push is allowed.
	AllowForcePush bool `pulumi:"allowForcePush"`
	// Reject code pushes that change files listed in the CODEOWNERS file.
	CodeOwnerApprovalRequired bool `pulumi:"codeOwnerApprovalRequired"`
	// The ID of this resource.
	Id int `pulumi:"id"`
	// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
	MergeAccessLevels []GetProjectProtectedBranchMergeAccessLevel `pulumi:"mergeAccessLevels"`
	// The name of the protected branch.
	Name string `pulumi:"name"`
	// The integer or path with namespace that uniquely identifies the project.
	ProjectId string `pulumi:"projectId"`
	// Array of access levels and user(s)/group(s) allowed to push to protected branch.
	PushAccessLevels []GetProjectProtectedBranchPushAccessLevel `pulumi:"pushAccessLevels"`
}

func GetProjectProtectedBranchOutput(ctx *pulumi.Context, args GetProjectProtectedBranchOutputArgs, opts ...pulumi.InvokeOption) GetProjectProtectedBranchResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetProjectProtectedBranchResultOutput, error) {
			args := v.(GetProjectProtectedBranchArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("gitlab:index/getProjectProtectedBranch:getProjectProtectedBranch", args, GetProjectProtectedBranchResultOutput{}, options).(GetProjectProtectedBranchResultOutput), nil
		}).(GetProjectProtectedBranchResultOutput)
}

// A collection of arguments for invoking getProjectProtectedBranch.
type GetProjectProtectedBranchOutputArgs struct {
	// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
	MergeAccessLevels GetProjectProtectedBranchMergeAccessLevelArrayInput `pulumi:"mergeAccessLevels"`
	// The name of the protected branch.
	Name pulumi.StringInput `pulumi:"name"`
	// The integer or path with namespace that uniquely identifies the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// Array of access levels and user(s)/group(s) allowed to push to protected branch.
	PushAccessLevels GetProjectProtectedBranchPushAccessLevelArrayInput `pulumi:"pushAccessLevels"`
}

func (GetProjectProtectedBranchOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectProtectedBranchArgs)(nil)).Elem()
}

// A collection of values returned by getProjectProtectedBranch.
type GetProjectProtectedBranchResultOutput struct{ *pulumi.OutputState }

func (GetProjectProtectedBranchResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectProtectedBranchResult)(nil)).Elem()
}

func (o GetProjectProtectedBranchResultOutput) ToGetProjectProtectedBranchResultOutput() GetProjectProtectedBranchResultOutput {
	return o
}

func (o GetProjectProtectedBranchResultOutput) ToGetProjectProtectedBranchResultOutputWithContext(ctx context.Context) GetProjectProtectedBranchResultOutput {
	return o
}

// Whether force push is allowed.
func (o GetProjectProtectedBranchResultOutput) AllowForcePush() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProjectProtectedBranchResult) bool { return v.AllowForcePush }).(pulumi.BoolOutput)
}

// Reject code pushes that change files listed in the CODEOWNERS file.
func (o GetProjectProtectedBranchResultOutput) CodeOwnerApprovalRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProjectProtectedBranchResult) bool { return v.CodeOwnerApprovalRequired }).(pulumi.BoolOutput)
}

// The ID of this resource.
func (o GetProjectProtectedBranchResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetProjectProtectedBranchResult) int { return v.Id }).(pulumi.IntOutput)
}

// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
func (o GetProjectProtectedBranchResultOutput) MergeAccessLevels() GetProjectProtectedBranchMergeAccessLevelArrayOutput {
	return o.ApplyT(func(v GetProjectProtectedBranchResult) []GetProjectProtectedBranchMergeAccessLevel {
		return v.MergeAccessLevels
	}).(GetProjectProtectedBranchMergeAccessLevelArrayOutput)
}

// The name of the protected branch.
func (o GetProjectProtectedBranchResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectProtectedBranchResult) string { return v.Name }).(pulumi.StringOutput)
}

// The integer or path with namespace that uniquely identifies the project.
func (o GetProjectProtectedBranchResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectProtectedBranchResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Array of access levels and user(s)/group(s) allowed to push to protected branch.
func (o GetProjectProtectedBranchResultOutput) PushAccessLevels() GetProjectProtectedBranchPushAccessLevelArrayOutput {
	return o.ApplyT(func(v GetProjectProtectedBranchResult) []GetProjectProtectedBranchPushAccessLevel {
		return v.PushAccessLevels
	}).(GetProjectProtectedBranchPushAccessLevelArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectProtectedBranchResultOutput{})
}
