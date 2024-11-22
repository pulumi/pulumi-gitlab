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

// The `ServiceSlack` resource allows to manage the lifecycle of a project integration with Slack.
//
// > This resource is deprecated. use `IntegrationSlack`instead!
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#slack-notifications)
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
//			awesomeProject, err := gitlab.NewProject(ctx, "awesome_project", &gitlab.ProjectArgs{
//				Name:            pulumi.String("awesome_project"),
//				Description:     pulumi.String("My awesome project."),
//				VisibilityLevel: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewServiceSlack(ctx, "slack", &gitlab.ServiceSlackArgs{
//				Project:     awesomeProject.ID(),
//				Webhook:     pulumi.String("https://webhook.com"),
//				Username:    pulumi.String("myuser"),
//				PushEvents:  pulumi.Bool(true),
//				PushChannel: pulumi.String("push_chan"),
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
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_service_slack`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_service_slack.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// You can import a gitlab_service_slack.slack state using the project ID, e.g.
//
// ```sh
// $ pulumi import gitlab:index/serviceSlack:ServiceSlack email 1
// ```
type ServiceSlack struct {
	pulumi.CustomResourceState

	// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
	BranchesToBeNotified pulumi.StringOutput `pulumi:"branchesToBeNotified"`
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel pulumi.StringPtrOutput `pulumi:"confidentialIssueChannel"`
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolOutput `pulumi:"confidentialIssuesEvents"`
	// The name of the channel to receive confidential note events notifications.
	ConfidentialNoteChannel pulumi.StringOutput `pulumi:"confidentialNoteChannel"`
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents pulumi.BoolOutput `pulumi:"confidentialNoteEvents"`
	// The name of the channel to receive issue events notifications.
	IssueChannel pulumi.StringPtrOutput `pulumi:"issueChannel"`
	// Enable notifications for issues events.
	IssuesEvents pulumi.BoolOutput `pulumi:"issuesEvents"`
	// Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to https://github.com/xanzy/go-gitlab/issues/1354.
	JobEvents pulumi.BoolOutput `pulumi:"jobEvents"`
	// The name of the channel to receive merge request events notifications.
	MergeRequestChannel pulumi.StringPtrOutput `pulumi:"mergeRequestChannel"`
	// Enable notifications for merge requests events.
	MergeRequestsEvents pulumi.BoolOutput `pulumi:"mergeRequestsEvents"`
	// The name of the channel to receive note events notifications.
	NoteChannel pulumi.StringPtrOutput `pulumi:"noteChannel"`
	// Enable notifications for note events.
	NoteEvents pulumi.BoolOutput `pulumi:"noteEvents"`
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines pulumi.BoolOutput `pulumi:"notifyOnlyBrokenPipelines"`
	// This parameter has been replaced with `branchesToBeNotified`.
	//
	// Deprecated: use 'branches_to_be_notified' argument instead
	NotifyOnlyDefaultBranch pulumi.BoolOutput `pulumi:"notifyOnlyDefaultBranch"`
	// The name of the channel to receive pipeline events notifications.
	PipelineChannel pulumi.StringPtrOutput `pulumi:"pipelineChannel"`
	// Enable notifications for pipeline events.
	PipelineEvents pulumi.BoolOutput `pulumi:"pipelineEvents"`
	// ID of the project you want to activate integration on.
	Project pulumi.StringOutput `pulumi:"project"`
	// The name of the channel to receive push events notifications.
	PushChannel pulumi.StringPtrOutput `pulumi:"pushChannel"`
	// Enable notifications for push events.
	PushEvents pulumi.BoolOutput `pulumi:"pushEvents"`
	// The name of the channel to receive tag push events notifications.
	TagPushChannel pulumi.StringPtrOutput `pulumi:"tagPushChannel"`
	// Enable notifications for tag push events.
	TagPushEvents pulumi.BoolOutput `pulumi:"tagPushEvents"`
	// Username to use.
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// Webhook URL (Example, https://hooks.slack.com/services/...). This value cannot be imported.
	Webhook pulumi.StringOutput `pulumi:"webhook"`
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel pulumi.StringPtrOutput `pulumi:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	WikiPageEvents pulumi.BoolOutput `pulumi:"wikiPageEvents"`
}

// NewServiceSlack registers a new resource with the given unique name, arguments, and options.
func NewServiceSlack(ctx *pulumi.Context,
	name string, args *ServiceSlackArgs, opts ...pulumi.ResourceOption) (*ServiceSlack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Webhook == nil {
		return nil, errors.New("invalid value for required argument 'Webhook'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceSlack
	err := ctx.RegisterResource("gitlab:index/serviceSlack:ServiceSlack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceSlack gets an existing ServiceSlack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceSlack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceSlackState, opts ...pulumi.ResourceOption) (*ServiceSlack, error) {
	var resource ServiceSlack
	err := ctx.ReadResource("gitlab:index/serviceSlack:ServiceSlack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceSlack resources.
type serviceSlackState struct {
	// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel *string `pulumi:"confidentialIssueChannel"`
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// The name of the channel to receive confidential note events notifications.
	ConfidentialNoteChannel *string `pulumi:"confidentialNoteChannel"`
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// The name of the channel to receive issue events notifications.
	IssueChannel *string `pulumi:"issueChannel"`
	// Enable notifications for issues events.
	IssuesEvents *bool `pulumi:"issuesEvents"`
	// Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to https://github.com/xanzy/go-gitlab/issues/1354.
	JobEvents *bool `pulumi:"jobEvents"`
	// The name of the channel to receive merge request events notifications.
	MergeRequestChannel *string `pulumi:"mergeRequestChannel"`
	// Enable notifications for merge requests events.
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// The name of the channel to receive note events notifications.
	NoteChannel *string `pulumi:"noteChannel"`
	// Enable notifications for note events.
	NoteEvents *bool `pulumi:"noteEvents"`
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// This parameter has been replaced with `branchesToBeNotified`.
	//
	// Deprecated: use 'branches_to_be_notified' argument instead
	NotifyOnlyDefaultBranch *bool `pulumi:"notifyOnlyDefaultBranch"`
	// The name of the channel to receive pipeline events notifications.
	PipelineChannel *string `pulumi:"pipelineChannel"`
	// Enable notifications for pipeline events.
	PipelineEvents *bool `pulumi:"pipelineEvents"`
	// ID of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// The name of the channel to receive push events notifications.
	PushChannel *string `pulumi:"pushChannel"`
	// Enable notifications for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// The name of the channel to receive tag push events notifications.
	TagPushChannel *string `pulumi:"tagPushChannel"`
	// Enable notifications for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// Username to use.
	Username *string `pulumi:"username"`
	// Webhook URL (Example, https://hooks.slack.com/services/...). This value cannot be imported.
	Webhook *string `pulumi:"webhook"`
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel *string `pulumi:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

type ServiceSlackState struct {
	// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
	BranchesToBeNotified pulumi.StringPtrInput
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel pulumi.StringPtrInput
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// The name of the channel to receive confidential note events notifications.
	ConfidentialNoteChannel pulumi.StringPtrInput
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// The name of the channel to receive issue events notifications.
	IssueChannel pulumi.StringPtrInput
	// Enable notifications for issues events.
	IssuesEvents pulumi.BoolPtrInput
	// Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to https://github.com/xanzy/go-gitlab/issues/1354.
	JobEvents pulumi.BoolPtrInput
	// The name of the channel to receive merge request events notifications.
	MergeRequestChannel pulumi.StringPtrInput
	// Enable notifications for merge requests events.
	MergeRequestsEvents pulumi.BoolPtrInput
	// The name of the channel to receive note events notifications.
	NoteChannel pulumi.StringPtrInput
	// Enable notifications for note events.
	NoteEvents pulumi.BoolPtrInput
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// This parameter has been replaced with `branchesToBeNotified`.
	//
	// Deprecated: use 'branches_to_be_notified' argument instead
	NotifyOnlyDefaultBranch pulumi.BoolPtrInput
	// The name of the channel to receive pipeline events notifications.
	PipelineChannel pulumi.StringPtrInput
	// Enable notifications for pipeline events.
	PipelineEvents pulumi.BoolPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// The name of the channel to receive push events notifications.
	PushChannel pulumi.StringPtrInput
	// Enable notifications for push events.
	PushEvents pulumi.BoolPtrInput
	// The name of the channel to receive tag push events notifications.
	TagPushChannel pulumi.StringPtrInput
	// Enable notifications for tag push events.
	TagPushEvents pulumi.BoolPtrInput
	// Username to use.
	Username pulumi.StringPtrInput
	// Webhook URL (Example, https://hooks.slack.com/services/...). This value cannot be imported.
	Webhook pulumi.StringPtrInput
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel pulumi.StringPtrInput
	// Enable notifications for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (ServiceSlackState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceSlackState)(nil)).Elem()
}

type serviceSlackArgs struct {
	// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel *string `pulumi:"confidentialIssueChannel"`
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// The name of the channel to receive confidential note events notifications.
	ConfidentialNoteChannel *string `pulumi:"confidentialNoteChannel"`
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// The name of the channel to receive issue events notifications.
	IssueChannel *string `pulumi:"issueChannel"`
	// Enable notifications for issues events.
	IssuesEvents *bool `pulumi:"issuesEvents"`
	// The name of the channel to receive merge request events notifications.
	MergeRequestChannel *string `pulumi:"mergeRequestChannel"`
	// Enable notifications for merge requests events.
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// The name of the channel to receive note events notifications.
	NoteChannel *string `pulumi:"noteChannel"`
	// Enable notifications for note events.
	NoteEvents *bool `pulumi:"noteEvents"`
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// This parameter has been replaced with `branchesToBeNotified`.
	//
	// Deprecated: use 'branches_to_be_notified' argument instead
	NotifyOnlyDefaultBranch *bool `pulumi:"notifyOnlyDefaultBranch"`
	// The name of the channel to receive pipeline events notifications.
	PipelineChannel *string `pulumi:"pipelineChannel"`
	// Enable notifications for pipeline events.
	PipelineEvents *bool `pulumi:"pipelineEvents"`
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// The name of the channel to receive push events notifications.
	PushChannel *string `pulumi:"pushChannel"`
	// Enable notifications for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// The name of the channel to receive tag push events notifications.
	TagPushChannel *string `pulumi:"tagPushChannel"`
	// Enable notifications for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// Username to use.
	Username *string `pulumi:"username"`
	// Webhook URL (Example, https://hooks.slack.com/services/...). This value cannot be imported.
	Webhook string `pulumi:"webhook"`
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel *string `pulumi:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

// The set of arguments for constructing a ServiceSlack resource.
type ServiceSlackArgs struct {
	// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
	BranchesToBeNotified pulumi.StringPtrInput
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel pulumi.StringPtrInput
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// The name of the channel to receive confidential note events notifications.
	ConfidentialNoteChannel pulumi.StringPtrInput
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// The name of the channel to receive issue events notifications.
	IssueChannel pulumi.StringPtrInput
	// Enable notifications for issues events.
	IssuesEvents pulumi.BoolPtrInput
	// The name of the channel to receive merge request events notifications.
	MergeRequestChannel pulumi.StringPtrInput
	// Enable notifications for merge requests events.
	MergeRequestsEvents pulumi.BoolPtrInput
	// The name of the channel to receive note events notifications.
	NoteChannel pulumi.StringPtrInput
	// Enable notifications for note events.
	NoteEvents pulumi.BoolPtrInput
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// This parameter has been replaced with `branchesToBeNotified`.
	//
	// Deprecated: use 'branches_to_be_notified' argument instead
	NotifyOnlyDefaultBranch pulumi.BoolPtrInput
	// The name of the channel to receive pipeline events notifications.
	PipelineChannel pulumi.StringPtrInput
	// Enable notifications for pipeline events.
	PipelineEvents pulumi.BoolPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
	// The name of the channel to receive push events notifications.
	PushChannel pulumi.StringPtrInput
	// Enable notifications for push events.
	PushEvents pulumi.BoolPtrInput
	// The name of the channel to receive tag push events notifications.
	TagPushChannel pulumi.StringPtrInput
	// Enable notifications for tag push events.
	TagPushEvents pulumi.BoolPtrInput
	// Username to use.
	Username pulumi.StringPtrInput
	// Webhook URL (Example, https://hooks.slack.com/services/...). This value cannot be imported.
	Webhook pulumi.StringInput
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel pulumi.StringPtrInput
	// Enable notifications for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (ServiceSlackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceSlackArgs)(nil)).Elem()
}

type ServiceSlackInput interface {
	pulumi.Input

	ToServiceSlackOutput() ServiceSlackOutput
	ToServiceSlackOutputWithContext(ctx context.Context) ServiceSlackOutput
}

func (*ServiceSlack) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceSlack)(nil)).Elem()
}

func (i *ServiceSlack) ToServiceSlackOutput() ServiceSlackOutput {
	return i.ToServiceSlackOutputWithContext(context.Background())
}

func (i *ServiceSlack) ToServiceSlackOutputWithContext(ctx context.Context) ServiceSlackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSlackOutput)
}

// ServiceSlackArrayInput is an input type that accepts ServiceSlackArray and ServiceSlackArrayOutput values.
// You can construct a concrete instance of `ServiceSlackArrayInput` via:
//
//	ServiceSlackArray{ ServiceSlackArgs{...} }
type ServiceSlackArrayInput interface {
	pulumi.Input

	ToServiceSlackArrayOutput() ServiceSlackArrayOutput
	ToServiceSlackArrayOutputWithContext(context.Context) ServiceSlackArrayOutput
}

type ServiceSlackArray []ServiceSlackInput

func (ServiceSlackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceSlack)(nil)).Elem()
}

func (i ServiceSlackArray) ToServiceSlackArrayOutput() ServiceSlackArrayOutput {
	return i.ToServiceSlackArrayOutputWithContext(context.Background())
}

func (i ServiceSlackArray) ToServiceSlackArrayOutputWithContext(ctx context.Context) ServiceSlackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSlackArrayOutput)
}

// ServiceSlackMapInput is an input type that accepts ServiceSlackMap and ServiceSlackMapOutput values.
// You can construct a concrete instance of `ServiceSlackMapInput` via:
//
//	ServiceSlackMap{ "key": ServiceSlackArgs{...} }
type ServiceSlackMapInput interface {
	pulumi.Input

	ToServiceSlackMapOutput() ServiceSlackMapOutput
	ToServiceSlackMapOutputWithContext(context.Context) ServiceSlackMapOutput
}

type ServiceSlackMap map[string]ServiceSlackInput

func (ServiceSlackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceSlack)(nil)).Elem()
}

func (i ServiceSlackMap) ToServiceSlackMapOutput() ServiceSlackMapOutput {
	return i.ToServiceSlackMapOutputWithContext(context.Background())
}

func (i ServiceSlackMap) ToServiceSlackMapOutputWithContext(ctx context.Context) ServiceSlackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSlackMapOutput)
}

type ServiceSlackOutput struct{ *pulumi.OutputState }

func (ServiceSlackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceSlack)(nil)).Elem()
}

func (o ServiceSlackOutput) ToServiceSlackOutput() ServiceSlackOutput {
	return o
}

func (o ServiceSlackOutput) ToServiceSlackOutputWithContext(ctx context.Context) ServiceSlackOutput {
	return o
}

// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
func (o ServiceSlackOutput) BranchesToBeNotified() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringOutput { return v.BranchesToBeNotified }).(pulumi.StringOutput)
}

// The name of the channel to receive confidential issue events notifications.
func (o ServiceSlackOutput) ConfidentialIssueChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringPtrOutput { return v.ConfidentialIssueChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for confidential issues events.
func (o ServiceSlackOutput) ConfidentialIssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.ConfidentialIssuesEvents }).(pulumi.BoolOutput)
}

// The name of the channel to receive confidential note events notifications.
func (o ServiceSlackOutput) ConfidentialNoteChannel() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringOutput { return v.ConfidentialNoteChannel }).(pulumi.StringOutput)
}

// Enable notifications for confidential note events.
func (o ServiceSlackOutput) ConfidentialNoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.ConfidentialNoteEvents }).(pulumi.BoolOutput)
}

// The name of the channel to receive issue events notifications.
func (o ServiceSlackOutput) IssueChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringPtrOutput { return v.IssueChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for issues events.
func (o ServiceSlackOutput) IssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.IssuesEvents }).(pulumi.BoolOutput)
}

// Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to https://github.com/xanzy/go-gitlab/issues/1354.
func (o ServiceSlackOutput) JobEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.JobEvents }).(pulumi.BoolOutput)
}

// The name of the channel to receive merge request events notifications.
func (o ServiceSlackOutput) MergeRequestChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringPtrOutput { return v.MergeRequestChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for merge requests events.
func (o ServiceSlackOutput) MergeRequestsEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.MergeRequestsEvents }).(pulumi.BoolOutput)
}

// The name of the channel to receive note events notifications.
func (o ServiceSlackOutput) NoteChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringPtrOutput { return v.NoteChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for note events.
func (o ServiceSlackOutput) NoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.NoteEvents }).(pulumi.BoolOutput)
}

// Send notifications for broken pipelines.
func (o ServiceSlackOutput) NotifyOnlyBrokenPipelines() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.NotifyOnlyBrokenPipelines }).(pulumi.BoolOutput)
}

// This parameter has been replaced with `branchesToBeNotified`.
//
// Deprecated: use 'branches_to_be_notified' argument instead
func (o ServiceSlackOutput) NotifyOnlyDefaultBranch() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.NotifyOnlyDefaultBranch }).(pulumi.BoolOutput)
}

// The name of the channel to receive pipeline events notifications.
func (o ServiceSlackOutput) PipelineChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringPtrOutput { return v.PipelineChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for pipeline events.
func (o ServiceSlackOutput) PipelineEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.PipelineEvents }).(pulumi.BoolOutput)
}

// ID of the project you want to activate integration on.
func (o ServiceSlackOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name of the channel to receive push events notifications.
func (o ServiceSlackOutput) PushChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringPtrOutput { return v.PushChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for push events.
func (o ServiceSlackOutput) PushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.PushEvents }).(pulumi.BoolOutput)
}

// The name of the channel to receive tag push events notifications.
func (o ServiceSlackOutput) TagPushChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringPtrOutput { return v.TagPushChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for tag push events.
func (o ServiceSlackOutput) TagPushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.TagPushEvents }).(pulumi.BoolOutput)
}

// Username to use.
func (o ServiceSlackOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

// Webhook URL (Example, https://hooks.slack.com/services/...). This value cannot be imported.
func (o ServiceSlackOutput) Webhook() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringOutput { return v.Webhook }).(pulumi.StringOutput)
}

// The name of the channel to receive wiki page events notifications.
func (o ServiceSlackOutput) WikiPageChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.StringPtrOutput { return v.WikiPageChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for wiki page events.
func (o ServiceSlackOutput) WikiPageEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceSlack) pulumi.BoolOutput { return v.WikiPageEvents }).(pulumi.BoolOutput)
}

type ServiceSlackArrayOutput struct{ *pulumi.OutputState }

func (ServiceSlackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceSlack)(nil)).Elem()
}

func (o ServiceSlackArrayOutput) ToServiceSlackArrayOutput() ServiceSlackArrayOutput {
	return o
}

func (o ServiceSlackArrayOutput) ToServiceSlackArrayOutputWithContext(ctx context.Context) ServiceSlackArrayOutput {
	return o
}

func (o ServiceSlackArrayOutput) Index(i pulumi.IntInput) ServiceSlackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceSlack {
		return vs[0].([]*ServiceSlack)[vs[1].(int)]
	}).(ServiceSlackOutput)
}

type ServiceSlackMapOutput struct{ *pulumi.OutputState }

func (ServiceSlackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceSlack)(nil)).Elem()
}

func (o ServiceSlackMapOutput) ToServiceSlackMapOutput() ServiceSlackMapOutput {
	return o
}

func (o ServiceSlackMapOutput) ToServiceSlackMapOutputWithContext(ctx context.Context) ServiceSlackMapOutput {
	return o
}

func (o ServiceSlackMapOutput) MapIndex(k pulumi.StringInput) ServiceSlackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceSlack {
		return vs[0].(map[string]*ServiceSlack)[vs[1].(string)]
	}).(ServiceSlackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceSlackInput)(nil)).Elem(), &ServiceSlack{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceSlackArrayInput)(nil)).Elem(), ServiceSlackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceSlackMapInput)(nil)).Elem(), ServiceSlackMap{})
	pulumi.RegisterOutputType(ServiceSlackOutput{})
	pulumi.RegisterOutputType(ServiceSlackArrayOutput{})
	pulumi.RegisterOutputType(ServiceSlackMapOutput{})
}
