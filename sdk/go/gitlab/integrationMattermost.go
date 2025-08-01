// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `IntegrationMattermost` resource manages the lifecycle of a project integration with Mattermost.
//
// > This resource is deprecated and will be removed in 19.0. Use `ProjectIntegrationMattermost`instead!
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#mattermost-notifications)
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
//			awesomeProject, err := gitlab.NewProject(ctx, "awesome_project", &gitlab.ProjectArgs{
//				Name:            pulumi.String("awesome_project"),
//				Description:     pulumi.String("My awesome project."),
//				VisibilityLevel: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewIntegrationMattermost(ctx, "mattermost", &gitlab.IntegrationMattermostArgs{
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
// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_integration_mattermost`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_integration_mattermost.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Importing using the CLI is supported with the following syntax:
//
// You can import a gitlab_integration_mattermost.mattermost state using the project ID, e.g.
//
// ```sh
// $ pulumi import gitlab:index/integrationMattermost:IntegrationMattermost mattermost 1
// ```
type IntegrationMattermost struct {
	pulumi.CustomResourceState

	// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
	BranchesToBeNotified pulumi.StringOutput `pulumi:"branchesToBeNotified"`
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel pulumi.StringPtrOutput `pulumi:"confidentialIssueChannel"`
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolOutput `pulumi:"confidentialIssuesEvents"`
	// The name of the channel to receive confidential note events notifications.
	ConfidentialNoteChannel pulumi.StringPtrOutput `pulumi:"confidentialNoteChannel"`
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents pulumi.BoolOutput `pulumi:"confidentialNoteEvents"`
	// The name of the channel to receive issue events notifications.
	IssueChannel pulumi.StringPtrOutput `pulumi:"issueChannel"`
	// Enable notifications for issues events.
	IssuesEvents pulumi.BoolOutput `pulumi:"issuesEvents"`
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
	// Webhook URL (Example, https://mattermost.yourdomain.com/hooks/...). This value cannot be imported.
	Webhook pulumi.StringOutput `pulumi:"webhook"`
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel pulumi.StringPtrOutput `pulumi:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	WikiPageEvents pulumi.BoolOutput `pulumi:"wikiPageEvents"`
}

// NewIntegrationMattermost registers a new resource with the given unique name, arguments, and options.
func NewIntegrationMattermost(ctx *pulumi.Context,
	name string, args *IntegrationMattermostArgs, opts ...pulumi.ResourceOption) (*IntegrationMattermost, error) {
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
	var resource IntegrationMattermost
	err := ctx.RegisterResource("gitlab:index/integrationMattermost:IntegrationMattermost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationMattermost gets an existing IntegrationMattermost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationMattermost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationMattermostState, opts ...pulumi.ResourceOption) (*IntegrationMattermost, error) {
	var resource IntegrationMattermost
	err := ctx.ReadResource("gitlab:index/integrationMattermost:IntegrationMattermost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationMattermost resources.
type integrationMattermostState struct {
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
	// Webhook URL (Example, https://mattermost.yourdomain.com/hooks/...). This value cannot be imported.
	Webhook *string `pulumi:"webhook"`
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel *string `pulumi:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

type IntegrationMattermostState struct {
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
	// Webhook URL (Example, https://mattermost.yourdomain.com/hooks/...). This value cannot be imported.
	Webhook pulumi.StringPtrInput
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel pulumi.StringPtrInput
	// Enable notifications for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (IntegrationMattermostState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationMattermostState)(nil)).Elem()
}

type integrationMattermostArgs struct {
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
	// Webhook URL (Example, https://mattermost.yourdomain.com/hooks/...). This value cannot be imported.
	Webhook string `pulumi:"webhook"`
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel *string `pulumi:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

// The set of arguments for constructing a IntegrationMattermost resource.
type IntegrationMattermostArgs struct {
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
	// Webhook URL (Example, https://mattermost.yourdomain.com/hooks/...). This value cannot be imported.
	Webhook pulumi.StringInput
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel pulumi.StringPtrInput
	// Enable notifications for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (IntegrationMattermostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationMattermostArgs)(nil)).Elem()
}

type IntegrationMattermostInput interface {
	pulumi.Input

	ToIntegrationMattermostOutput() IntegrationMattermostOutput
	ToIntegrationMattermostOutputWithContext(ctx context.Context) IntegrationMattermostOutput
}

func (*IntegrationMattermost) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationMattermost)(nil)).Elem()
}

func (i *IntegrationMattermost) ToIntegrationMattermostOutput() IntegrationMattermostOutput {
	return i.ToIntegrationMattermostOutputWithContext(context.Background())
}

func (i *IntegrationMattermost) ToIntegrationMattermostOutputWithContext(ctx context.Context) IntegrationMattermostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMattermostOutput)
}

// IntegrationMattermostArrayInput is an input type that accepts IntegrationMattermostArray and IntegrationMattermostArrayOutput values.
// You can construct a concrete instance of `IntegrationMattermostArrayInput` via:
//
//	IntegrationMattermostArray{ IntegrationMattermostArgs{...} }
type IntegrationMattermostArrayInput interface {
	pulumi.Input

	ToIntegrationMattermostArrayOutput() IntegrationMattermostArrayOutput
	ToIntegrationMattermostArrayOutputWithContext(context.Context) IntegrationMattermostArrayOutput
}

type IntegrationMattermostArray []IntegrationMattermostInput

func (IntegrationMattermostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationMattermost)(nil)).Elem()
}

func (i IntegrationMattermostArray) ToIntegrationMattermostArrayOutput() IntegrationMattermostArrayOutput {
	return i.ToIntegrationMattermostArrayOutputWithContext(context.Background())
}

func (i IntegrationMattermostArray) ToIntegrationMattermostArrayOutputWithContext(ctx context.Context) IntegrationMattermostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMattermostArrayOutput)
}

// IntegrationMattermostMapInput is an input type that accepts IntegrationMattermostMap and IntegrationMattermostMapOutput values.
// You can construct a concrete instance of `IntegrationMattermostMapInput` via:
//
//	IntegrationMattermostMap{ "key": IntegrationMattermostArgs{...} }
type IntegrationMattermostMapInput interface {
	pulumi.Input

	ToIntegrationMattermostMapOutput() IntegrationMattermostMapOutput
	ToIntegrationMattermostMapOutputWithContext(context.Context) IntegrationMattermostMapOutput
}

type IntegrationMattermostMap map[string]IntegrationMattermostInput

func (IntegrationMattermostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationMattermost)(nil)).Elem()
}

func (i IntegrationMattermostMap) ToIntegrationMattermostMapOutput() IntegrationMattermostMapOutput {
	return i.ToIntegrationMattermostMapOutputWithContext(context.Background())
}

func (i IntegrationMattermostMap) ToIntegrationMattermostMapOutputWithContext(ctx context.Context) IntegrationMattermostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMattermostMapOutput)
}

type IntegrationMattermostOutput struct{ *pulumi.OutputState }

func (IntegrationMattermostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationMattermost)(nil)).Elem()
}

func (o IntegrationMattermostOutput) ToIntegrationMattermostOutput() IntegrationMattermostOutput {
	return o
}

func (o IntegrationMattermostOutput) ToIntegrationMattermostOutputWithContext(ctx context.Context) IntegrationMattermostOutput {
	return o
}

// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
func (o IntegrationMattermostOutput) BranchesToBeNotified() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringOutput { return v.BranchesToBeNotified }).(pulumi.StringOutput)
}

// The name of the channel to receive confidential issue events notifications.
func (o IntegrationMattermostOutput) ConfidentialIssueChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringPtrOutput { return v.ConfidentialIssueChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for confidential issues events.
func (o IntegrationMattermostOutput) ConfidentialIssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.BoolOutput { return v.ConfidentialIssuesEvents }).(pulumi.BoolOutput)
}

// The name of the channel to receive confidential note events notifications.
func (o IntegrationMattermostOutput) ConfidentialNoteChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringPtrOutput { return v.ConfidentialNoteChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for confidential note events.
func (o IntegrationMattermostOutput) ConfidentialNoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.BoolOutput { return v.ConfidentialNoteEvents }).(pulumi.BoolOutput)
}

// The name of the channel to receive issue events notifications.
func (o IntegrationMattermostOutput) IssueChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringPtrOutput { return v.IssueChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for issues events.
func (o IntegrationMattermostOutput) IssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.BoolOutput { return v.IssuesEvents }).(pulumi.BoolOutput)
}

// The name of the channel to receive merge request events notifications.
func (o IntegrationMattermostOutput) MergeRequestChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringPtrOutput { return v.MergeRequestChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for merge requests events.
func (o IntegrationMattermostOutput) MergeRequestsEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.BoolOutput { return v.MergeRequestsEvents }).(pulumi.BoolOutput)
}

// The name of the channel to receive note events notifications.
func (o IntegrationMattermostOutput) NoteChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringPtrOutput { return v.NoteChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for note events.
func (o IntegrationMattermostOutput) NoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.BoolOutput { return v.NoteEvents }).(pulumi.BoolOutput)
}

// Send notifications for broken pipelines.
func (o IntegrationMattermostOutput) NotifyOnlyBrokenPipelines() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.BoolOutput { return v.NotifyOnlyBrokenPipelines }).(pulumi.BoolOutput)
}

// The name of the channel to receive pipeline events notifications.
func (o IntegrationMattermostOutput) PipelineChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringPtrOutput { return v.PipelineChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for pipeline events.
func (o IntegrationMattermostOutput) PipelineEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.BoolOutput { return v.PipelineEvents }).(pulumi.BoolOutput)
}

// ID of the project you want to activate integration on.
func (o IntegrationMattermostOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name of the channel to receive push events notifications.
func (o IntegrationMattermostOutput) PushChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringPtrOutput { return v.PushChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for push events.
func (o IntegrationMattermostOutput) PushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.BoolOutput { return v.PushEvents }).(pulumi.BoolOutput)
}

// The name of the channel to receive tag push events notifications.
func (o IntegrationMattermostOutput) TagPushChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringPtrOutput { return v.TagPushChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for tag push events.
func (o IntegrationMattermostOutput) TagPushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.BoolOutput { return v.TagPushEvents }).(pulumi.BoolOutput)
}

// Username to use.
func (o IntegrationMattermostOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

// Webhook URL (Example, https://mattermost.yourdomain.com/hooks/...). This value cannot be imported.
func (o IntegrationMattermostOutput) Webhook() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringOutput { return v.Webhook }).(pulumi.StringOutput)
}

// The name of the channel to receive wiki page events notifications.
func (o IntegrationMattermostOutput) WikiPageChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.StringPtrOutput { return v.WikiPageChannel }).(pulumi.StringPtrOutput)
}

// Enable notifications for wiki page events.
func (o IntegrationMattermostOutput) WikiPageEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMattermost) pulumi.BoolOutput { return v.WikiPageEvents }).(pulumi.BoolOutput)
}

type IntegrationMattermostArrayOutput struct{ *pulumi.OutputState }

func (IntegrationMattermostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationMattermost)(nil)).Elem()
}

func (o IntegrationMattermostArrayOutput) ToIntegrationMattermostArrayOutput() IntegrationMattermostArrayOutput {
	return o
}

func (o IntegrationMattermostArrayOutput) ToIntegrationMattermostArrayOutputWithContext(ctx context.Context) IntegrationMattermostArrayOutput {
	return o
}

func (o IntegrationMattermostArrayOutput) Index(i pulumi.IntInput) IntegrationMattermostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationMattermost {
		return vs[0].([]*IntegrationMattermost)[vs[1].(int)]
	}).(IntegrationMattermostOutput)
}

type IntegrationMattermostMapOutput struct{ *pulumi.OutputState }

func (IntegrationMattermostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationMattermost)(nil)).Elem()
}

func (o IntegrationMattermostMapOutput) ToIntegrationMattermostMapOutput() IntegrationMattermostMapOutput {
	return o
}

func (o IntegrationMattermostMapOutput) ToIntegrationMattermostMapOutputWithContext(ctx context.Context) IntegrationMattermostMapOutput {
	return o
}

func (o IntegrationMattermostMapOutput) MapIndex(k pulumi.StringInput) IntegrationMattermostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationMattermost {
		return vs[0].(map[string]*IntegrationMattermost)[vs[1].(string)]
	}).(IntegrationMattermostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMattermostInput)(nil)).Elem(), &IntegrationMattermost{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMattermostArrayInput)(nil)).Elem(), IntegrationMattermostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMattermostMapInput)(nil)).Elem(), IntegrationMattermostMap{})
	pulumi.RegisterOutputType(IntegrationMattermostOutput{})
	pulumi.RegisterOutputType(IntegrationMattermostArrayOutput{})
	pulumi.RegisterOutputType(IntegrationMattermostMapOutput{})
}
