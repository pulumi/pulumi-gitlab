// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ProjectHook` resource allows to manage the lifecycle of a project hook.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#hooks)
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
//			_, err := gitlab.NewProjectHook(ctx, "example", &gitlab.ProjectHookArgs{
//				Project:             pulumi.String("example/hooked"),
//				Url:                 pulumi.String("https://example.com/hook/example"),
//				MergeRequestsEvents: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			// Using Custom Headers
//			// Values of headers can't be imported
//			_, err = gitlab.NewProjectHook(ctx, "custom_headers", &gitlab.ProjectHookArgs{
//				Project:             pulumi.String("example/hooked"),
//				Url:                 pulumi.String("https://example.com/hook/example"),
//				MergeRequestsEvents: pulumi.Bool(true),
//				CustomHeaders: gitlab.ProjectHookCustomHeaderArray{
//					&gitlab.ProjectHookCustomHeaderArgs{
//						Key:   pulumi.String("X-Custom-Header"),
//						Value: pulumi.String("example"),
//					},
//					&gitlab.ProjectHookCustomHeaderArgs{
//						Key:   pulumi.String("X-Custom-Header-Second"),
//						Value: pulumi.String("example-second"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_hook`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_project_hook.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// A GitLab Project Hook can be imported using a key composed of `<project-id>:<hook-id>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectHook:ProjectHook example "12345:1"
// ```
//
// NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
type ProjectHook struct {
	pulumi.CustomResourceState

	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolOutput `pulumi:"confidentialIssuesEvents"`
	// Invoke the hook for confidential note events.
	ConfidentialNoteEvents pulumi.BoolOutput `pulumi:"confidentialNoteEvents"`
	// Custom headers for the project webhook.
	CustomHeaders ProjectHookCustomHeaderArrayOutput `pulumi:"customHeaders"`
	// Custom webhook template.
	CustomWebhookTemplate pulumi.StringOutput `pulumi:"customWebhookTemplate"`
	// Invoke the hook for deployment events.
	DeploymentEvents pulumi.BoolOutput `pulumi:"deploymentEvents"`
	// Enable SSL verification when invoking the hook.
	EnableSslVerification pulumi.BoolOutput `pulumi:"enableSslVerification"`
	// The id of the project hook.
	HookId pulumi.IntOutput `pulumi:"hookId"`
	// Invoke the hook for issues events.
	IssuesEvents pulumi.BoolOutput `pulumi:"issuesEvents"`
	// Invoke the hook for job events.
	JobEvents pulumi.BoolOutput `pulumi:"jobEvents"`
	// Invoke the hook for merge requests events.
	MergeRequestsEvents pulumi.BoolOutput `pulumi:"mergeRequestsEvents"`
	// Invoke the hook for note events.
	NoteEvents pulumi.BoolOutput `pulumi:"noteEvents"`
	// Invoke the hook for pipeline events.
	PipelineEvents pulumi.BoolOutput `pulumi:"pipelineEvents"`
	// The name or id of the project to add the hook to.
	Project pulumi.StringOutput `pulumi:"project"`
	// The id of the project for the hook.
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Invoke the hook for push events.
	PushEvents pulumi.BoolOutput `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter pulumi.StringOutput `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for release events.
	ReleasesEvents pulumi.BoolOutput `pulumi:"releasesEvents"`
	// Invoke the hook for tag push events.
	TagPushEvents pulumi.BoolOutput `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token pulumi.StringOutput `pulumi:"token"`
	// The url of the hook to invoke. Forces re-creation to preserve `token`.
	Url pulumi.StringOutput `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents pulumi.BoolOutput `pulumi:"wikiPageEvents"`
}

// NewProjectHook registers a new resource with the given unique name, arguments, and options.
func NewProjectHook(ctx *pulumi.Context,
	name string, args *ProjectHookArgs, opts ...pulumi.ResourceOption) (*ProjectHook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectHook
	err := ctx.RegisterResource("gitlab:index/projectHook:ProjectHook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectHook gets an existing ProjectHook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectHook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectHookState, opts ...pulumi.ResourceOption) (*ProjectHook, error) {
	var resource ProjectHook
	err := ctx.ReadResource("gitlab:index/projectHook:ProjectHook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectHook resources.
type projectHookState struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// Invoke the hook for confidential note events.
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// Custom headers for the project webhook.
	CustomHeaders []ProjectHookCustomHeader `pulumi:"customHeaders"`
	// Custom webhook template.
	CustomWebhookTemplate *string `pulumi:"customWebhookTemplate"`
	// Invoke the hook for deployment events.
	DeploymentEvents *bool `pulumi:"deploymentEvents"`
	// Enable SSL verification when invoking the hook.
	EnableSslVerification *bool `pulumi:"enableSslVerification"`
	// The id of the project hook.
	HookId *int `pulumi:"hookId"`
	// Invoke the hook for issues events.
	IssuesEvents *bool `pulumi:"issuesEvents"`
	// Invoke the hook for job events.
	JobEvents *bool `pulumi:"jobEvents"`
	// Invoke the hook for merge requests events.
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// Invoke the hook for note events.
	NoteEvents *bool `pulumi:"noteEvents"`
	// Invoke the hook for pipeline events.
	PipelineEvents *bool `pulumi:"pipelineEvents"`
	// The name or id of the project to add the hook to.
	Project *string `pulumi:"project"`
	// The id of the project for the hook.
	ProjectId *int `pulumi:"projectId"`
	// Invoke the hook for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter *string `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for release events.
	ReleasesEvents *bool `pulumi:"releasesEvents"`
	// Invoke the hook for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token *string `pulumi:"token"`
	// The url of the hook to invoke. Forces re-creation to preserve `token`.
	Url *string `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

type ProjectHookState struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// Invoke the hook for confidential note events.
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// Custom headers for the project webhook.
	CustomHeaders ProjectHookCustomHeaderArrayInput
	// Custom webhook template.
	CustomWebhookTemplate pulumi.StringPtrInput
	// Invoke the hook for deployment events.
	DeploymentEvents pulumi.BoolPtrInput
	// Enable SSL verification when invoking the hook.
	EnableSslVerification pulumi.BoolPtrInput
	// The id of the project hook.
	HookId pulumi.IntPtrInput
	// Invoke the hook for issues events.
	IssuesEvents pulumi.BoolPtrInput
	// Invoke the hook for job events.
	JobEvents pulumi.BoolPtrInput
	// Invoke the hook for merge requests events.
	MergeRequestsEvents pulumi.BoolPtrInput
	// Invoke the hook for note events.
	NoteEvents pulumi.BoolPtrInput
	// Invoke the hook for pipeline events.
	PipelineEvents pulumi.BoolPtrInput
	// The name or id of the project to add the hook to.
	Project pulumi.StringPtrInput
	// The id of the project for the hook.
	ProjectId pulumi.IntPtrInput
	// Invoke the hook for push events.
	PushEvents pulumi.BoolPtrInput
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter pulumi.StringPtrInput
	// Invoke the hook for release events.
	ReleasesEvents pulumi.BoolPtrInput
	// Invoke the hook for tag push events.
	TagPushEvents pulumi.BoolPtrInput
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token pulumi.StringPtrInput
	// The url of the hook to invoke. Forces re-creation to preserve `token`.
	Url pulumi.StringPtrInput
	// Invoke the hook for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (ProjectHookState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectHookState)(nil)).Elem()
}

type projectHookArgs struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// Invoke the hook for confidential note events.
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// Custom headers for the project webhook.
	CustomHeaders []ProjectHookCustomHeader `pulumi:"customHeaders"`
	// Custom webhook template.
	CustomWebhookTemplate *string `pulumi:"customWebhookTemplate"`
	// Invoke the hook for deployment events.
	DeploymentEvents *bool `pulumi:"deploymentEvents"`
	// Enable SSL verification when invoking the hook.
	EnableSslVerification *bool `pulumi:"enableSslVerification"`
	// Invoke the hook for issues events.
	IssuesEvents *bool `pulumi:"issuesEvents"`
	// Invoke the hook for job events.
	JobEvents *bool `pulumi:"jobEvents"`
	// Invoke the hook for merge requests events.
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// Invoke the hook for note events.
	NoteEvents *bool `pulumi:"noteEvents"`
	// Invoke the hook for pipeline events.
	PipelineEvents *bool `pulumi:"pipelineEvents"`
	// The name or id of the project to add the hook to.
	Project string `pulumi:"project"`
	// Invoke the hook for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter *string `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for release events.
	ReleasesEvents *bool `pulumi:"releasesEvents"`
	// Invoke the hook for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token *string `pulumi:"token"`
	// The url of the hook to invoke. Forces re-creation to preserve `token`.
	Url string `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

// The set of arguments for constructing a ProjectHook resource.
type ProjectHookArgs struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// Invoke the hook for confidential note events.
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// Custom headers for the project webhook.
	CustomHeaders ProjectHookCustomHeaderArrayInput
	// Custom webhook template.
	CustomWebhookTemplate pulumi.StringPtrInput
	// Invoke the hook for deployment events.
	DeploymentEvents pulumi.BoolPtrInput
	// Enable SSL verification when invoking the hook.
	EnableSslVerification pulumi.BoolPtrInput
	// Invoke the hook for issues events.
	IssuesEvents pulumi.BoolPtrInput
	// Invoke the hook for job events.
	JobEvents pulumi.BoolPtrInput
	// Invoke the hook for merge requests events.
	MergeRequestsEvents pulumi.BoolPtrInput
	// Invoke the hook for note events.
	NoteEvents pulumi.BoolPtrInput
	// Invoke the hook for pipeline events.
	PipelineEvents pulumi.BoolPtrInput
	// The name or id of the project to add the hook to.
	Project pulumi.StringInput
	// Invoke the hook for push events.
	PushEvents pulumi.BoolPtrInput
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter pulumi.StringPtrInput
	// Invoke the hook for release events.
	ReleasesEvents pulumi.BoolPtrInput
	// Invoke the hook for tag push events.
	TagPushEvents pulumi.BoolPtrInput
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token pulumi.StringPtrInput
	// The url of the hook to invoke. Forces re-creation to preserve `token`.
	Url pulumi.StringInput
	// Invoke the hook for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (ProjectHookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectHookArgs)(nil)).Elem()
}

type ProjectHookInput interface {
	pulumi.Input

	ToProjectHookOutput() ProjectHookOutput
	ToProjectHookOutputWithContext(ctx context.Context) ProjectHookOutput
}

func (*ProjectHook) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectHook)(nil)).Elem()
}

func (i *ProjectHook) ToProjectHookOutput() ProjectHookOutput {
	return i.ToProjectHookOutputWithContext(context.Background())
}

func (i *ProjectHook) ToProjectHookOutputWithContext(ctx context.Context) ProjectHookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectHookOutput)
}

// ProjectHookArrayInput is an input type that accepts ProjectHookArray and ProjectHookArrayOutput values.
// You can construct a concrete instance of `ProjectHookArrayInput` via:
//
//	ProjectHookArray{ ProjectHookArgs{...} }
type ProjectHookArrayInput interface {
	pulumi.Input

	ToProjectHookArrayOutput() ProjectHookArrayOutput
	ToProjectHookArrayOutputWithContext(context.Context) ProjectHookArrayOutput
}

type ProjectHookArray []ProjectHookInput

func (ProjectHookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectHook)(nil)).Elem()
}

func (i ProjectHookArray) ToProjectHookArrayOutput() ProjectHookArrayOutput {
	return i.ToProjectHookArrayOutputWithContext(context.Background())
}

func (i ProjectHookArray) ToProjectHookArrayOutputWithContext(ctx context.Context) ProjectHookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectHookArrayOutput)
}

// ProjectHookMapInput is an input type that accepts ProjectHookMap and ProjectHookMapOutput values.
// You can construct a concrete instance of `ProjectHookMapInput` via:
//
//	ProjectHookMap{ "key": ProjectHookArgs{...} }
type ProjectHookMapInput interface {
	pulumi.Input

	ToProjectHookMapOutput() ProjectHookMapOutput
	ToProjectHookMapOutputWithContext(context.Context) ProjectHookMapOutput
}

type ProjectHookMap map[string]ProjectHookInput

func (ProjectHookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectHook)(nil)).Elem()
}

func (i ProjectHookMap) ToProjectHookMapOutput() ProjectHookMapOutput {
	return i.ToProjectHookMapOutputWithContext(context.Background())
}

func (i ProjectHookMap) ToProjectHookMapOutputWithContext(ctx context.Context) ProjectHookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectHookMapOutput)
}

type ProjectHookOutput struct{ *pulumi.OutputState }

func (ProjectHookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectHook)(nil)).Elem()
}

func (o ProjectHookOutput) ToProjectHookOutput() ProjectHookOutput {
	return o
}

func (o ProjectHookOutput) ToProjectHookOutputWithContext(ctx context.Context) ProjectHookOutput {
	return o
}

// Invoke the hook for confidential issues events.
func (o ProjectHookOutput) ConfidentialIssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.ConfidentialIssuesEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for confidential note events.
func (o ProjectHookOutput) ConfidentialNoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.ConfidentialNoteEvents }).(pulumi.BoolOutput)
}

// Custom headers for the project webhook.
func (o ProjectHookOutput) CustomHeaders() ProjectHookCustomHeaderArrayOutput {
	return o.ApplyT(func(v *ProjectHook) ProjectHookCustomHeaderArrayOutput { return v.CustomHeaders }).(ProjectHookCustomHeaderArrayOutput)
}

// Custom webhook template.
func (o ProjectHookOutput) CustomWebhookTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.StringOutput { return v.CustomWebhookTemplate }).(pulumi.StringOutput)
}

// Invoke the hook for deployment events.
func (o ProjectHookOutput) DeploymentEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.DeploymentEvents }).(pulumi.BoolOutput)
}

// Enable SSL verification when invoking the hook.
func (o ProjectHookOutput) EnableSslVerification() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.EnableSslVerification }).(pulumi.BoolOutput)
}

// The id of the project hook.
func (o ProjectHookOutput) HookId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.IntOutput { return v.HookId }).(pulumi.IntOutput)
}

// Invoke the hook for issues events.
func (o ProjectHookOutput) IssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.IssuesEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for job events.
func (o ProjectHookOutput) JobEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.JobEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for merge requests events.
func (o ProjectHookOutput) MergeRequestsEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.MergeRequestsEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for note events.
func (o ProjectHookOutput) NoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.NoteEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for pipeline events.
func (o ProjectHookOutput) PipelineEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.PipelineEvents }).(pulumi.BoolOutput)
}

// The name or id of the project to add the hook to.
func (o ProjectHookOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The id of the project for the hook.
func (o ProjectHookOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Invoke the hook for push events.
func (o ProjectHookOutput) PushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.PushEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for push events on matching branches only.
func (o ProjectHookOutput) PushEventsBranchFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.StringOutput { return v.PushEventsBranchFilter }).(pulumi.StringOutput)
}

// Invoke the hook for release events.
func (o ProjectHookOutput) ReleasesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.ReleasesEvents }).(pulumi.BoolOutput)
}

// Invoke the hook for tag push events.
func (o ProjectHookOutput) TagPushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.TagPushEvents }).(pulumi.BoolOutput)
}

// A token to present when invoking the hook. The token is not available for imported resources.
func (o ProjectHookOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The url of the hook to invoke. Forces re-creation to preserve `token`.
func (o ProjectHookOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Invoke the hook for wiki page events.
func (o ProjectHookOutput) WikiPageEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectHook) pulumi.BoolOutput { return v.WikiPageEvents }).(pulumi.BoolOutput)
}

type ProjectHookArrayOutput struct{ *pulumi.OutputState }

func (ProjectHookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectHook)(nil)).Elem()
}

func (o ProjectHookArrayOutput) ToProjectHookArrayOutput() ProjectHookArrayOutput {
	return o
}

func (o ProjectHookArrayOutput) ToProjectHookArrayOutputWithContext(ctx context.Context) ProjectHookArrayOutput {
	return o
}

func (o ProjectHookArrayOutput) Index(i pulumi.IntInput) ProjectHookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectHook {
		return vs[0].([]*ProjectHook)[vs[1].(int)]
	}).(ProjectHookOutput)
}

type ProjectHookMapOutput struct{ *pulumi.OutputState }

func (ProjectHookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectHook)(nil)).Elem()
}

func (o ProjectHookMapOutput) ToProjectHookMapOutput() ProjectHookMapOutput {
	return o
}

func (o ProjectHookMapOutput) ToProjectHookMapOutputWithContext(ctx context.Context) ProjectHookMapOutput {
	return o
}

func (o ProjectHookMapOutput) MapIndex(k pulumi.StringInput) ProjectHookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectHook {
		return vs[0].(map[string]*ProjectHook)[vs[1].(string)]
	}).(ProjectHookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectHookInput)(nil)).Elem(), &ProjectHook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectHookArrayInput)(nil)).Elem(), ProjectHookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectHookMapInput)(nil)).Elem(), ProjectHookMap{})
	pulumi.RegisterOutputType(ProjectHookOutput{})
	pulumi.RegisterOutputType(ProjectHookArrayOutput{})
	pulumi.RegisterOutputType(ProjectHookMapOutput{})
}
