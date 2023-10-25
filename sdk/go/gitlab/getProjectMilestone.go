// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The `ProjectMilestone` data source allows get details of a project milestone.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)
func LookupProjectMilestone(ctx *pulumi.Context, args *LookupProjectMilestoneArgs, opts ...pulumi.InvokeOption) (*LookupProjectMilestoneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectMilestoneResult
	err := ctx.Invoke("gitlab:index/getProjectMilestone:getProjectMilestone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectMilestone.
type LookupProjectMilestoneArgs struct {
	// The instance-wide ID of the project’s milestone.
	MilestoneId int `pulumi:"milestoneId"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project string `pulumi:"project"`
}

// A collection of values returned by getProjectMilestone.
type LookupProjectMilestoneResult struct {
	// The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	CreatedAt string `pulumi:"createdAt"`
	// The description of the milestone.
	Description string `pulumi:"description"`
	// The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	DueDate string `pulumi:"dueDate"`
	// Bool, true if milestone expired.
	Expired bool `pulumi:"expired"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the project's milestone.
	Iid int `pulumi:"iid"`
	// The instance-wide ID of the project’s milestone.
	MilestoneId int `pulumi:"milestoneId"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project string `pulumi:"project"`
	// The project ID of milestone.
	ProjectId int `pulumi:"projectId"`
	// The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	StartDate string `pulumi:"startDate"`
	// The state of the milestone. Valid values are: `active`, `closed`.
	State string `pulumi:"state"`
	// The title of a milestone.
	Title string `pulumi:"title"`
	// The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	UpdatedAt string `pulumi:"updatedAt"`
	// The web URL of the milestone.
	WebUrl string `pulumi:"webUrl"`
}

func LookupProjectMilestoneOutput(ctx *pulumi.Context, args LookupProjectMilestoneOutputArgs, opts ...pulumi.InvokeOption) LookupProjectMilestoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectMilestoneResult, error) {
			args := v.(LookupProjectMilestoneArgs)
			r, err := LookupProjectMilestone(ctx, &args, opts...)
			var s LookupProjectMilestoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectMilestoneResultOutput)
}

// A collection of arguments for invoking getProjectMilestone.
type LookupProjectMilestoneOutputArgs struct {
	// The instance-wide ID of the project’s milestone.
	MilestoneId pulumi.IntInput `pulumi:"milestoneId"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project pulumi.StringInput `pulumi:"project"`
}

func (LookupProjectMilestoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectMilestoneArgs)(nil)).Elem()
}

// A collection of values returned by getProjectMilestone.
type LookupProjectMilestoneResultOutput struct{ *pulumi.OutputState }

func (LookupProjectMilestoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectMilestoneResult)(nil)).Elem()
}

func (o LookupProjectMilestoneResultOutput) ToLookupProjectMilestoneResultOutput() LookupProjectMilestoneResultOutput {
	return o
}

func (o LookupProjectMilestoneResultOutput) ToLookupProjectMilestoneResultOutputWithContext(ctx context.Context) LookupProjectMilestoneResultOutput {
	return o
}

func (o LookupProjectMilestoneResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupProjectMilestoneResult] {
	return pulumix.Output[LookupProjectMilestoneResult]{
		OutputState: o.OutputState,
	}
}

// The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
func (o LookupProjectMilestoneResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the milestone.
func (o LookupProjectMilestoneResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) string { return v.Description }).(pulumi.StringOutput)
}

// The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
func (o LookupProjectMilestoneResultOutput) DueDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) string { return v.DueDate }).(pulumi.StringOutput)
}

// Bool, true if milestone expired.
func (o LookupProjectMilestoneResultOutput) Expired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) bool { return v.Expired }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectMilestoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the project's milestone.
func (o LookupProjectMilestoneResultOutput) Iid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) int { return v.Iid }).(pulumi.IntOutput)
}

// The instance-wide ID of the project’s milestone.
func (o LookupProjectMilestoneResultOutput) MilestoneId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) int { return v.MilestoneId }).(pulumi.IntOutput)
}

// The ID or URL-encoded path of the project owned by the authenticated user.
func (o LookupProjectMilestoneResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) string { return v.Project }).(pulumi.StringOutput)
}

// The project ID of milestone.
func (o LookupProjectMilestoneResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

// The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
func (o LookupProjectMilestoneResultOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) string { return v.StartDate }).(pulumi.StringOutput)
}

// The state of the milestone. Valid values are: `active`, `closed`.
func (o LookupProjectMilestoneResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) string { return v.State }).(pulumi.StringOutput)
}

// The title of a milestone.
func (o LookupProjectMilestoneResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) string { return v.Title }).(pulumi.StringOutput)
}

// The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
func (o LookupProjectMilestoneResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The web URL of the milestone.
func (o LookupProjectMilestoneResultOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMilestoneResult) string { return v.WebUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectMilestoneResultOutput{})
}
