// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ProjectHook` data source allows to retrieve details about a hook in a project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_webhooks/#get-a-project-webhook)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
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
//			_, err = gitlab.LookupProjectHook(ctx, &gitlab.LookupProjectHookArgs{
//				Project: example.Id,
//				HookId:  1,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupProjectHook(ctx *pulumi.Context, args *LookupProjectHookArgs, opts ...pulumi.InvokeOption) (*LookupProjectHookResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectHookResult
	err := ctx.Invoke("gitlab:index/getProjectHook:getProjectHook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectHook.
type LookupProjectHookArgs struct {
	// The id of the project hook.
	HookId int `pulumi:"hookId"`
	// The name or id of the project to add the hook to.
	Project string `pulumi:"project"`
}

// A collection of values returned by getProjectHook.
type LookupProjectHookResult struct {
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
	// The id of the project hook.
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
	// The name or id of the project to add the hook to.
	Project string `pulumi:"project"`
	// The id of the project for the hook.
	ProjectId int `pulumi:"projectId"`
	// Invoke the hook for push events.
	PushEvents bool `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter string `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for releases events.
	ReleasesEvents bool `pulumi:"releasesEvents"`
	// Invoke the hook for tag push events.
	TagPushEvents bool `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token string `pulumi:"token"`
	// The url of the hook to invoke.
	Url string `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents bool `pulumi:"wikiPageEvents"`
}

func LookupProjectHookOutput(ctx *pulumi.Context, args LookupProjectHookOutputArgs, opts ...pulumi.InvokeOption) LookupProjectHookResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectHookResultOutput, error) {
			args := v.(LookupProjectHookArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("gitlab:index/getProjectHook:getProjectHook", args, LookupProjectHookResultOutput{}, options).(LookupProjectHookResultOutput), nil
		}).(LookupProjectHookResultOutput)
}

// A collection of arguments for invoking getProjectHook.
type LookupProjectHookOutputArgs struct {
	// The id of the project hook.
	HookId pulumi.IntInput `pulumi:"hookId"`
	// The name or id of the project to add the hook to.
	Project pulumi.StringInput `pulumi:"project"`
}

func (LookupProjectHookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectHookArgs)(nil)).Elem()
}

// A collection of values returned by getProjectHook.
type LookupProjectHookResultOutput struct{ *pulumi.OutputState }

func (LookupProjectHookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectHookResult)(nil)).Elem()
}

func (o LookupProjectHookResultOutput) ToLookupProjectHookResultOutput() LookupProjectHookResultOutput {
	return o
}

func (o LookupProjectHookResultOutput) ToLookupProjectHookResultOutputWithContext(ctx context.Context) LookupProjectHookResultOutput {
	return o
}

// Invoke the hook for confidential issues events.
func (o LookupProjectHookResultOutput) ConfidentialIssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.ConfidentialIssuesEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for confidential notes events.
func (o LookupProjectHookResultOutput) ConfidentialNoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.ConfidentialNoteEvents }).(pulumi.BoolOutput)
}

// Set a custom webhook template.
func (o LookupProjectHookResultOutput) CustomWebhookTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectHookResult) string { return v.CustomWebhookTemplate }).(pulumi.StringOutput)
}

// Invoke the hook for deployment events.
func (o LookupProjectHookResultOutput) DeploymentEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.DeploymentEvents }).(pulumi.BoolOutput)
}

// Enable ssl verification when invoking the hook.
func (o LookupProjectHookResultOutput) EnableSslVerification() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.EnableSslVerification }).(pulumi.BoolOutput)
}

// The id of the project hook.
func (o LookupProjectHookResultOutput) HookId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectHookResult) int { return v.HookId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectHookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectHookResult) string { return v.Id }).(pulumi.StringOutput)
}

// Invoke the hook for issues events.
func (o LookupProjectHookResultOutput) IssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.IssuesEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for job events.
func (o LookupProjectHookResultOutput) JobEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.JobEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for merge requests.
func (o LookupProjectHookResultOutput) MergeRequestsEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.MergeRequestsEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for notes events.
func (o LookupProjectHookResultOutput) NoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.NoteEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for pipeline events.
func (o LookupProjectHookResultOutput) PipelineEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.PipelineEvents }).(pulumi.BoolOutput)
}

// The name or id of the project to add the hook to.
func (o LookupProjectHookResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectHookResult) string { return v.Project }).(pulumi.StringOutput)
}

// The id of the project for the hook.
func (o LookupProjectHookResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectHookResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

// Invoke the hook for push events.
func (o LookupProjectHookResultOutput) PushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.PushEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for push events on matching branches only.
func (o LookupProjectHookResultOutput) PushEventsBranchFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectHookResult) string { return v.PushEventsBranchFilter }).(pulumi.StringOutput)
}

// Invoke the hook for releases events.
func (o LookupProjectHookResultOutput) ReleasesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.ReleasesEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for tag push events.
func (o LookupProjectHookResultOutput) TagPushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.TagPushEvents }).(pulumi.BoolOutput)
}

// A token to present when invoking the hook. The token is not available for imported resources.
func (o LookupProjectHookResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectHookResult) string { return v.Token }).(pulumi.StringOutput)
}

// The url of the hook to invoke.
func (o LookupProjectHookResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectHookResult) string { return v.Url }).(pulumi.StringOutput)
}

// Invoke the hook for wiki page events.
func (o LookupProjectHookResultOutput) WikiPageEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectHookResult) bool { return v.WikiPageEvents }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectHookResultOutput{})
}
