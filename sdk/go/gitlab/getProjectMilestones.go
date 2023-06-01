// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getProjectMilestones` data source allows get details of a project milestones.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)
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
//			_, err := gitlab.GetProjectMilestones(ctx, &gitlab.GetProjectMilestonesArgs{
//				Project: "foo/bar",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProjectMilestones(ctx *pulumi.Context, args *GetProjectMilestonesArgs, opts ...pulumi.InvokeOption) (*GetProjectMilestonesResult, error) {
	var rv GetProjectMilestonesResult
	err := ctx.Invoke("gitlab:index/getProjectMilestones:getProjectMilestones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectMilestones.
type GetProjectMilestonesArgs struct {
	// Return only the milestones having the given `iid` (Note: ignored if `includeParentMilestones` is set as `true`).
	Iids []int `pulumi:"iids"`
	// Include group milestones from parent group and its ancestors. Introduced in GitLab 13.4.
	IncludeParentMilestones *bool `pulumi:"includeParentMilestones"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project string `pulumi:"project"`
	// Return only milestones with a title or description matching the provided string.
	Search *string `pulumi:"search"`
	// Return only `active` or `closed` milestones.
	State *string `pulumi:"state"`
	// Return only the milestones having the given `title`.
	Title *string `pulumi:"title"`
}

// A collection of values returned by getProjectMilestones.
type GetProjectMilestonesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Return only the milestones having the given `iid` (Note: ignored if `includeParentMilestones` is set as `true`).
	Iids []int `pulumi:"iids"`
	// Include group milestones from parent group and its ancestors. Introduced in GitLab 13.4.
	IncludeParentMilestones *bool `pulumi:"includeParentMilestones"`
	// List of milestones from a project.
	Milestones []GetProjectMilestonesMilestone `pulumi:"milestones"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project string `pulumi:"project"`
	// Return only milestones with a title or description matching the provided string.
	Search *string `pulumi:"search"`
	// Return only `active` or `closed` milestones.
	State *string `pulumi:"state"`
	// Return only the milestones having the given `title`.
	Title *string `pulumi:"title"`
}

func GetProjectMilestonesOutput(ctx *pulumi.Context, args GetProjectMilestonesOutputArgs, opts ...pulumi.InvokeOption) GetProjectMilestonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectMilestonesResult, error) {
			args := v.(GetProjectMilestonesArgs)
			r, err := GetProjectMilestones(ctx, &args, opts...)
			var s GetProjectMilestonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectMilestonesResultOutput)
}

// A collection of arguments for invoking getProjectMilestones.
type GetProjectMilestonesOutputArgs struct {
	// Return only the milestones having the given `iid` (Note: ignored if `includeParentMilestones` is set as `true`).
	Iids pulumi.IntArrayInput `pulumi:"iids"`
	// Include group milestones from parent group and its ancestors. Introduced in GitLab 13.4.
	IncludeParentMilestones pulumi.BoolPtrInput `pulumi:"includeParentMilestones"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project pulumi.StringInput `pulumi:"project"`
	// Return only milestones with a title or description matching the provided string.
	Search pulumi.StringPtrInput `pulumi:"search"`
	// Return only `active` or `closed` milestones.
	State pulumi.StringPtrInput `pulumi:"state"`
	// Return only the milestones having the given `title`.
	Title pulumi.StringPtrInput `pulumi:"title"`
}

func (GetProjectMilestonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectMilestonesArgs)(nil)).Elem()
}

// A collection of values returned by getProjectMilestones.
type GetProjectMilestonesResultOutput struct{ *pulumi.OutputState }

func (GetProjectMilestonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectMilestonesResult)(nil)).Elem()
}

func (o GetProjectMilestonesResultOutput) ToGetProjectMilestonesResultOutput() GetProjectMilestonesResultOutput {
	return o
}

func (o GetProjectMilestonesResultOutput) ToGetProjectMilestonesResultOutputWithContext(ctx context.Context) GetProjectMilestonesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectMilestonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectMilestonesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Return only the milestones having the given `iid` (Note: ignored if `includeParentMilestones` is set as `true`).
func (o GetProjectMilestonesResultOutput) Iids() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetProjectMilestonesResult) []int { return v.Iids }).(pulumi.IntArrayOutput)
}

// Include group milestones from parent group and its ancestors. Introduced in GitLab 13.4.
func (o GetProjectMilestonesResultOutput) IncludeParentMilestones() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectMilestonesResult) *bool { return v.IncludeParentMilestones }).(pulumi.BoolPtrOutput)
}

// List of milestones from a project.
func (o GetProjectMilestonesResultOutput) Milestones() GetProjectMilestonesMilestoneArrayOutput {
	return o.ApplyT(func(v GetProjectMilestonesResult) []GetProjectMilestonesMilestone { return v.Milestones }).(GetProjectMilestonesMilestoneArrayOutput)
}

// The ID or URL-encoded path of the project owned by the authenticated user.
func (o GetProjectMilestonesResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectMilestonesResult) string { return v.Project }).(pulumi.StringOutput)
}

// Return only milestones with a title or description matching the provided string.
func (o GetProjectMilestonesResultOutput) Search() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectMilestonesResult) *string { return v.Search }).(pulumi.StringPtrOutput)
}

// Return only `active` or `closed` milestones.
func (o GetProjectMilestonesResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectMilestonesResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Return only the milestones having the given `title`.
func (o GetProjectMilestonesResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectMilestonesResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectMilestonesResultOutput{})
}
