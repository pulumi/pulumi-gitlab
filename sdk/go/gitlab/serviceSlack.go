// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ServiceSlack struct {
	pulumi.CustomResourceState

	// Branches to send notifications for. Valid options are "all", "default", "protected", and "defaultAndProtected".
	BranchesToBeNotified pulumi.StringOutput `pulumi:"branchesToBeNotified"`
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel pulumi.StringPtrOutput `pulumi:"confidentialIssueChannel"`
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolOutput `pulumi:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents pulumi.BoolOutput `pulumi:"confidentialNoteEvents"`
	// The name of the channel to receive issue events notifications.
	IssueChannel pulumi.StringPtrOutput `pulumi:"issueChannel"`
	// Enable notifications for issues events.
	IssuesEvents pulumi.BoolOutput `pulumi:"issuesEvents"`
	JobEvents    pulumi.BoolOutput `pulumi:"jobEvents"`
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
	// DEPRECATED: This parameter has been replaced with `branchesToBeNotified`.
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
	// Webhook URL (ex.: https://hooks.slack.com/services/...)
	Webhook pulumi.StringOutput `pulumi:"webhook"`
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel pulumi.StringPtrOutput `pulumi:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	WikiPageEvents pulumi.BoolOutput `pulumi:"wikiPageEvents"`
}

// NewServiceSlack registers a new resource with the given unique name, arguments, and options.
func NewServiceSlack(ctx *pulumi.Context,
	name string, args *ServiceSlackArgs, opts ...pulumi.ResourceOption) (*ServiceSlack, error) {
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil || args.Webhook == nil {
		return nil, errors.New("missing required argument 'Webhook'")
	}
	if args == nil {
		args = &ServiceSlackArgs{}
	}
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
	// Branches to send notifications for. Valid options are "all", "default", "protected", and "defaultAndProtected".
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel *string `pulumi:"confidentialIssueChannel"`
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// The name of the channel to receive issue events notifications.
	IssueChannel *string `pulumi:"issueChannel"`
	// Enable notifications for issues events.
	IssuesEvents *bool `pulumi:"issuesEvents"`
	JobEvents    *bool `pulumi:"jobEvents"`
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
	// DEPRECATED: This parameter has been replaced with `branchesToBeNotified`.
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
	// Webhook URL (ex.: https://hooks.slack.com/services/...)
	Webhook *string `pulumi:"webhook"`
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel *string `pulumi:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

type ServiceSlackState struct {
	// Branches to send notifications for. Valid options are "all", "default", "protected", and "defaultAndProtected".
	BranchesToBeNotified pulumi.StringPtrInput
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel pulumi.StringPtrInput
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// The name of the channel to receive issue events notifications.
	IssueChannel pulumi.StringPtrInput
	// Enable notifications for issues events.
	IssuesEvents pulumi.BoolPtrInput
	JobEvents    pulumi.BoolPtrInput
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
	// DEPRECATED: This parameter has been replaced with `branchesToBeNotified`.
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
	// Webhook URL (ex.: https://hooks.slack.com/services/...)
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
	// Branches to send notifications for. Valid options are "all", "default", "protected", and "defaultAndProtected".
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel *string `pulumi:"confidentialIssueChannel"`
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
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
	// DEPRECATED: This parameter has been replaced with `branchesToBeNotified`.
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
	// Webhook URL (ex.: https://hooks.slack.com/services/...)
	Webhook string `pulumi:"webhook"`
	// The name of the channel to receive wiki page events notifications.
	WikiPageChannel *string `pulumi:"wikiPageChannel"`
	// Enable notifications for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

// The set of arguments for constructing a ServiceSlack resource.
type ServiceSlackArgs struct {
	// Branches to send notifications for. Valid options are "all", "default", "protected", and "defaultAndProtected".
	BranchesToBeNotified pulumi.StringPtrInput
	// The name of the channel to receive confidential issue events notifications.
	ConfidentialIssueChannel pulumi.StringPtrInput
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
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
	// DEPRECATED: This parameter has been replaced with `branchesToBeNotified`.
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
	// Webhook URL (ex.: https://hooks.slack.com/services/...)
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

func (ServiceSlack) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceSlack)(nil)).Elem()
}

func (i ServiceSlack) ToServiceSlackOutput() ServiceSlackOutput {
	return i.ToServiceSlackOutputWithContext(context.Background())
}

func (i ServiceSlack) ToServiceSlackOutputWithContext(ctx context.Context) ServiceSlackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSlackOutput)
}

type ServiceSlackOutput struct {
	*pulumi.OutputState
}

func (ServiceSlackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceSlackOutput)(nil)).Elem()
}

func (o ServiceSlackOutput) ToServiceSlackOutput() ServiceSlackOutput {
	return o
}

func (o ServiceSlackOutput) ToServiceSlackOutputWithContext(ctx context.Context) ServiceSlackOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceSlackOutput{})
}
