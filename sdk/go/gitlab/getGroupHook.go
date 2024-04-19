// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GroupHook` data source allows to retrieve details about a hook in a group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#get-group-hook)
func LookupGroupHook(ctx *pulumi.Context, args *LookupGroupHookArgs, opts ...pulumi.InvokeOption) (*LookupGroupHookResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupHookResult
	err := ctx.Invoke("gitlab:index/getGroupHook:getGroupHook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupHook.
type LookupGroupHookArgs struct {
	// The ID or full path of the group.
	Group string `pulumi:"group"`
	// The id of the group hook.
	HookId int `pulumi:"hookId"`
}

// A collection of values returned by getGroupHook.
type LookupGroupHookResult struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents bool `pulumi:"confidentialIssuesEvents"`
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents bool `pulumi:"confidentialNoteEvents"`
	// Set a custom webhook template.
	CustomWebhookTemplate string `pulumi:"customWebhookTemplate"`
	// Invoke the hook for deployment events.
	DeploymentEvents bool `pulumi:"deploymentEvents"`
	// Enable ssl verification when invoking the hook.
	EnableSslVerification bool `pulumi:"enableSslVerification"`
	// The ID or full path of the group.
	Group string `pulumi:"group"`
	// The id of the group for the hook.
	GroupId int `pulumi:"groupId"`
	// The id of the group hook.
	HookId int `pulumi:"hookId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Invoke the hook for issues events.
	IssuesEvents bool `pulumi:"issuesEvents"`
	// Invoke the hook for job events.
	JobEvents bool `pulumi:"jobEvents"`
	// Invoke the hook for merge requests.
	MergeRequestsEvents bool `pulumi:"mergeRequestsEvents"`
	// Invoke the hook for notes events.
	NoteEvents bool `pulumi:"noteEvents"`
	// Invoke the hook for pipeline events.
	PipelineEvents bool `pulumi:"pipelineEvents"`
	// Invoke the hook for push events.
	PushEvents bool `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter string `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for releases events.
	ReleasesEvents bool `pulumi:"releasesEvents"`
	// Invoke the hook for subgroup events.
	SubgroupEvents bool `pulumi:"subgroupEvents"`
	// Invoke the hook for tag push events.
	TagPushEvents bool `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token string `pulumi:"token"`
	// The url of the hook to invoke.
	Url string `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents bool `pulumi:"wikiPageEvents"`
}

func LookupGroupHookOutput(ctx *pulumi.Context, args LookupGroupHookOutputArgs, opts ...pulumi.InvokeOption) LookupGroupHookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupHookResult, error) {
			args := v.(LookupGroupHookArgs)
			r, err := LookupGroupHook(ctx, &args, opts...)
			var s LookupGroupHookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupHookResultOutput)
}

// A collection of arguments for invoking getGroupHook.
type LookupGroupHookOutputArgs struct {
	// The ID or full path of the group.
	Group pulumi.StringInput `pulumi:"group"`
	// The id of the group hook.
	HookId pulumi.IntInput `pulumi:"hookId"`
}

func (LookupGroupHookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupHookArgs)(nil)).Elem()
}

// A collection of values returned by getGroupHook.
type LookupGroupHookResultOutput struct{ *pulumi.OutputState }

func (LookupGroupHookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupHookResult)(nil)).Elem()
}

func (o LookupGroupHookResultOutput) ToLookupGroupHookResultOutput() LookupGroupHookResultOutput {
	return o
}

func (o LookupGroupHookResultOutput) ToLookupGroupHookResultOutputWithContext(ctx context.Context) LookupGroupHookResultOutput {
	return o
}

// Invoke the hook for confidential issues events.
func (o LookupGroupHookResultOutput) ConfidentialIssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.ConfidentialIssuesEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for confidential notes events.
func (o LookupGroupHookResultOutput) ConfidentialNoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.ConfidentialNoteEvents }).(pulumi.BoolOutput)
}

// Set a custom webhook template.
func (o LookupGroupHookResultOutput) CustomWebhookTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupHookResult) string { return v.CustomWebhookTemplate }).(pulumi.StringOutput)
}

// Invoke the hook for deployment events.
func (o LookupGroupHookResultOutput) DeploymentEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.DeploymentEvents }).(pulumi.BoolOutput)
}

// Enable ssl verification when invoking the hook.
func (o LookupGroupHookResultOutput) EnableSslVerification() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.EnableSslVerification }).(pulumi.BoolOutput)
}

// The ID or full path of the group.
func (o LookupGroupHookResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupHookResult) string { return v.Group }).(pulumi.StringOutput)
}

// The id of the group for the hook.
func (o LookupGroupHookResultOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupHookResult) int { return v.GroupId }).(pulumi.IntOutput)
}

// The id of the group hook.
func (o LookupGroupHookResultOutput) HookId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupHookResult) int { return v.HookId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupHookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupHookResult) string { return v.Id }).(pulumi.StringOutput)
}

// Invoke the hook for issues events.
func (o LookupGroupHookResultOutput) IssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.IssuesEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for job events.
func (o LookupGroupHookResultOutput) JobEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.JobEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for merge requests.
func (o LookupGroupHookResultOutput) MergeRequestsEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.MergeRequestsEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for notes events.
func (o LookupGroupHookResultOutput) NoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.NoteEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for pipeline events.
func (o LookupGroupHookResultOutput) PipelineEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.PipelineEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for push events.
func (o LookupGroupHookResultOutput) PushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.PushEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for push events on matching branches only.
func (o LookupGroupHookResultOutput) PushEventsBranchFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupHookResult) string { return v.PushEventsBranchFilter }).(pulumi.StringOutput)
}

// Invoke the hook for releases events.
func (o LookupGroupHookResultOutput) ReleasesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.ReleasesEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for subgroup events.
func (o LookupGroupHookResultOutput) SubgroupEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.SubgroupEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for tag push events.
func (o LookupGroupHookResultOutput) TagPushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.TagPushEvents }).(pulumi.BoolOutput)
}

// A token to present when invoking the hook. The token is not available for imported resources.
func (o LookupGroupHookResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupHookResult) string { return v.Token }).(pulumi.StringOutput)
}

// The url of the hook to invoke.
func (o LookupGroupHookResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupHookResult) string { return v.Url }).(pulumi.StringOutput)
}

// Invoke the hook for wiki page events.
func (o LookupGroupHookResultOutput) WikiPageEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupHookResult) bool { return v.WikiPageEvents }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupHookResultOutput{})
}
