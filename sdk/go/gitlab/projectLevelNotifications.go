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

// The `ProjectLevelNotifications` resource allows to manage notifications for a project.
//
// > While the API supports both groups and projects, this resource only supports projects currently.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/notification_settings.html#group--project-level-notification-settings)
type ProjectLevelNotifications struct {
	pulumi.CustomResourceState

	// Enable notifications for closed issues. Can only be used when `level` is `custom`.
	CloseIssue pulumi.BoolOutput `pulumi:"closeIssue"`
	// Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
	CloseMergeRequest pulumi.BoolOutput `pulumi:"closeMergeRequest"`
	// Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
	FailedPipeline pulumi.BoolOutput `pulumi:"failedPipeline"`
	// Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
	FixedPipeline pulumi.BoolOutput `pulumi:"fixedPipeline"`
	// Enable notifications for due issues. Can only be used when `level` is `custom`.
	IssueDue pulumi.BoolOutput `pulumi:"issueDue"`
	// The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
	Level pulumi.StringOutput `pulumi:"level"`
	// Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
	MergeMergeRequest pulumi.BoolOutput `pulumi:"mergeMergeRequest"`
	// Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
	MergeWhenPipelineSucceeds pulumi.BoolOutput `pulumi:"mergeWhenPipelineSucceeds"`
	// Enable notifications for moved projects. Can only be used when `level` is `custom`.
	MovedProject pulumi.BoolOutput `pulumi:"movedProject"`
	// Enable notifications for new issues. Can only be used when `level` is `custom`.
	NewIssue pulumi.BoolOutput `pulumi:"newIssue"`
	// Enable notifications for new merge requests. Can only be used when `level` is `custom`.
	NewMergeRequest pulumi.BoolOutput `pulumi:"newMergeRequest"`
	// Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
	NewNote pulumi.BoolOutput `pulumi:"newNote"`
	// The ID or URL-encoded path of a project where notifications will be configured.
	Project pulumi.StringOutput `pulumi:"project"`
	// Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
	PushToMergeRequest pulumi.BoolOutput `pulumi:"pushToMergeRequest"`
	// Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
	ReassignIssue pulumi.BoolOutput `pulumi:"reassignIssue"`
	// Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
	ReassignMergeRequest pulumi.BoolOutput `pulumi:"reassignMergeRequest"`
	// Enable notifications for reopened issues. Can only be used when `level` is `custom`.
	ReopenIssue pulumi.BoolOutput `pulumi:"reopenIssue"`
	// Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
	ReopenMergeRequest pulumi.BoolOutput `pulumi:"reopenMergeRequest"`
	// Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
	SuccessPipeline pulumi.BoolOutput `pulumi:"successPipeline"`
}

// NewProjectLevelNotifications registers a new resource with the given unique name, arguments, and options.
func NewProjectLevelNotifications(ctx *pulumi.Context,
	name string, args *ProjectLevelNotificationsArgs, opts ...pulumi.ResourceOption) (*ProjectLevelNotifications, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectLevelNotifications
	err := ctx.RegisterResource("gitlab:index/projectLevelNotifications:ProjectLevelNotifications", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectLevelNotifications gets an existing ProjectLevelNotifications resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectLevelNotifications(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectLevelNotificationsState, opts ...pulumi.ResourceOption) (*ProjectLevelNotifications, error) {
	var resource ProjectLevelNotifications
	err := ctx.ReadResource("gitlab:index/projectLevelNotifications:ProjectLevelNotifications", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectLevelNotifications resources.
type projectLevelNotificationsState struct {
	// Enable notifications for closed issues. Can only be used when `level` is `custom`.
	CloseIssue *bool `pulumi:"closeIssue"`
	// Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
	CloseMergeRequest *bool `pulumi:"closeMergeRequest"`
	// Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
	FailedPipeline *bool `pulumi:"failedPipeline"`
	// Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
	FixedPipeline *bool `pulumi:"fixedPipeline"`
	// Enable notifications for due issues. Can only be used when `level` is `custom`.
	IssueDue *bool `pulumi:"issueDue"`
	// The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
	Level *string `pulumi:"level"`
	// Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
	MergeMergeRequest *bool `pulumi:"mergeMergeRequest"`
	// Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
	MergeWhenPipelineSucceeds *bool `pulumi:"mergeWhenPipelineSucceeds"`
	// Enable notifications for moved projects. Can only be used when `level` is `custom`.
	MovedProject *bool `pulumi:"movedProject"`
	// Enable notifications for new issues. Can only be used when `level` is `custom`.
	NewIssue *bool `pulumi:"newIssue"`
	// Enable notifications for new merge requests. Can only be used when `level` is `custom`.
	NewMergeRequest *bool `pulumi:"newMergeRequest"`
	// Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
	NewNote *bool `pulumi:"newNote"`
	// The ID or URL-encoded path of a project where notifications will be configured.
	Project *string `pulumi:"project"`
	// Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
	PushToMergeRequest *bool `pulumi:"pushToMergeRequest"`
	// Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
	ReassignIssue *bool `pulumi:"reassignIssue"`
	// Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
	ReassignMergeRequest *bool `pulumi:"reassignMergeRequest"`
	// Enable notifications for reopened issues. Can only be used when `level` is `custom`.
	ReopenIssue *bool `pulumi:"reopenIssue"`
	// Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
	ReopenMergeRequest *bool `pulumi:"reopenMergeRequest"`
	// Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
	SuccessPipeline *bool `pulumi:"successPipeline"`
}

type ProjectLevelNotificationsState struct {
	// Enable notifications for closed issues. Can only be used when `level` is `custom`.
	CloseIssue pulumi.BoolPtrInput
	// Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
	CloseMergeRequest pulumi.BoolPtrInput
	// Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
	FailedPipeline pulumi.BoolPtrInput
	// Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
	FixedPipeline pulumi.BoolPtrInput
	// Enable notifications for due issues. Can only be used when `level` is `custom`.
	IssueDue pulumi.BoolPtrInput
	// The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
	Level pulumi.StringPtrInput
	// Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
	MergeMergeRequest pulumi.BoolPtrInput
	// Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
	MergeWhenPipelineSucceeds pulumi.BoolPtrInput
	// Enable notifications for moved projects. Can only be used when `level` is `custom`.
	MovedProject pulumi.BoolPtrInput
	// Enable notifications for new issues. Can only be used when `level` is `custom`.
	NewIssue pulumi.BoolPtrInput
	// Enable notifications for new merge requests. Can only be used when `level` is `custom`.
	NewMergeRequest pulumi.BoolPtrInput
	// Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
	NewNote pulumi.BoolPtrInput
	// The ID or URL-encoded path of a project where notifications will be configured.
	Project pulumi.StringPtrInput
	// Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
	PushToMergeRequest pulumi.BoolPtrInput
	// Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
	ReassignIssue pulumi.BoolPtrInput
	// Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
	ReassignMergeRequest pulumi.BoolPtrInput
	// Enable notifications for reopened issues. Can only be used when `level` is `custom`.
	ReopenIssue pulumi.BoolPtrInput
	// Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
	ReopenMergeRequest pulumi.BoolPtrInput
	// Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
	SuccessPipeline pulumi.BoolPtrInput
}

func (ProjectLevelNotificationsState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLevelNotificationsState)(nil)).Elem()
}

type projectLevelNotificationsArgs struct {
	// Enable notifications for closed issues. Can only be used when `level` is `custom`.
	CloseIssue *bool `pulumi:"closeIssue"`
	// Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
	CloseMergeRequest *bool `pulumi:"closeMergeRequest"`
	// Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
	FailedPipeline *bool `pulumi:"failedPipeline"`
	// Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
	FixedPipeline *bool `pulumi:"fixedPipeline"`
	// Enable notifications for due issues. Can only be used when `level` is `custom`.
	IssueDue *bool `pulumi:"issueDue"`
	// The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
	Level *string `pulumi:"level"`
	// Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
	MergeMergeRequest *bool `pulumi:"mergeMergeRequest"`
	// Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
	MergeWhenPipelineSucceeds *bool `pulumi:"mergeWhenPipelineSucceeds"`
	// Enable notifications for moved projects. Can only be used when `level` is `custom`.
	MovedProject *bool `pulumi:"movedProject"`
	// Enable notifications for new issues. Can only be used when `level` is `custom`.
	NewIssue *bool `pulumi:"newIssue"`
	// Enable notifications for new merge requests. Can only be used when `level` is `custom`.
	NewMergeRequest *bool `pulumi:"newMergeRequest"`
	// Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
	NewNote *bool `pulumi:"newNote"`
	// The ID or URL-encoded path of a project where notifications will be configured.
	Project string `pulumi:"project"`
	// Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
	PushToMergeRequest *bool `pulumi:"pushToMergeRequest"`
	// Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
	ReassignIssue *bool `pulumi:"reassignIssue"`
	// Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
	ReassignMergeRequest *bool `pulumi:"reassignMergeRequest"`
	// Enable notifications for reopened issues. Can only be used when `level` is `custom`.
	ReopenIssue *bool `pulumi:"reopenIssue"`
	// Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
	ReopenMergeRequest *bool `pulumi:"reopenMergeRequest"`
	// Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
	SuccessPipeline *bool `pulumi:"successPipeline"`
}

// The set of arguments for constructing a ProjectLevelNotifications resource.
type ProjectLevelNotificationsArgs struct {
	// Enable notifications for closed issues. Can only be used when `level` is `custom`.
	CloseIssue pulumi.BoolPtrInput
	// Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
	CloseMergeRequest pulumi.BoolPtrInput
	// Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
	FailedPipeline pulumi.BoolPtrInput
	// Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
	FixedPipeline pulumi.BoolPtrInput
	// Enable notifications for due issues. Can only be used when `level` is `custom`.
	IssueDue pulumi.BoolPtrInput
	// The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
	Level pulumi.StringPtrInput
	// Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
	MergeMergeRequest pulumi.BoolPtrInput
	// Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
	MergeWhenPipelineSucceeds pulumi.BoolPtrInput
	// Enable notifications for moved projects. Can only be used when `level` is `custom`.
	MovedProject pulumi.BoolPtrInput
	// Enable notifications for new issues. Can only be used when `level` is `custom`.
	NewIssue pulumi.BoolPtrInput
	// Enable notifications for new merge requests. Can only be used when `level` is `custom`.
	NewMergeRequest pulumi.BoolPtrInput
	// Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
	NewNote pulumi.BoolPtrInput
	// The ID or URL-encoded path of a project where notifications will be configured.
	Project pulumi.StringInput
	// Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
	PushToMergeRequest pulumi.BoolPtrInput
	// Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
	ReassignIssue pulumi.BoolPtrInput
	// Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
	ReassignMergeRequest pulumi.BoolPtrInput
	// Enable notifications for reopened issues. Can only be used when `level` is `custom`.
	ReopenIssue pulumi.BoolPtrInput
	// Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
	ReopenMergeRequest pulumi.BoolPtrInput
	// Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
	SuccessPipeline pulumi.BoolPtrInput
}

func (ProjectLevelNotificationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLevelNotificationsArgs)(nil)).Elem()
}

type ProjectLevelNotificationsInput interface {
	pulumi.Input

	ToProjectLevelNotificationsOutput() ProjectLevelNotificationsOutput
	ToProjectLevelNotificationsOutputWithContext(ctx context.Context) ProjectLevelNotificationsOutput
}

func (*ProjectLevelNotifications) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectLevelNotifications)(nil)).Elem()
}

func (i *ProjectLevelNotifications) ToProjectLevelNotificationsOutput() ProjectLevelNotificationsOutput {
	return i.ToProjectLevelNotificationsOutputWithContext(context.Background())
}

func (i *ProjectLevelNotifications) ToProjectLevelNotificationsOutputWithContext(ctx context.Context) ProjectLevelNotificationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectLevelNotificationsOutput)
}

// ProjectLevelNotificationsArrayInput is an input type that accepts ProjectLevelNotificationsArray and ProjectLevelNotificationsArrayOutput values.
// You can construct a concrete instance of `ProjectLevelNotificationsArrayInput` via:
//
//	ProjectLevelNotificationsArray{ ProjectLevelNotificationsArgs{...} }
type ProjectLevelNotificationsArrayInput interface {
	pulumi.Input

	ToProjectLevelNotificationsArrayOutput() ProjectLevelNotificationsArrayOutput
	ToProjectLevelNotificationsArrayOutputWithContext(context.Context) ProjectLevelNotificationsArrayOutput
}

type ProjectLevelNotificationsArray []ProjectLevelNotificationsInput

func (ProjectLevelNotificationsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectLevelNotifications)(nil)).Elem()
}

func (i ProjectLevelNotificationsArray) ToProjectLevelNotificationsArrayOutput() ProjectLevelNotificationsArrayOutput {
	return i.ToProjectLevelNotificationsArrayOutputWithContext(context.Background())
}

func (i ProjectLevelNotificationsArray) ToProjectLevelNotificationsArrayOutputWithContext(ctx context.Context) ProjectLevelNotificationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectLevelNotificationsArrayOutput)
}

// ProjectLevelNotificationsMapInput is an input type that accepts ProjectLevelNotificationsMap and ProjectLevelNotificationsMapOutput values.
// You can construct a concrete instance of `ProjectLevelNotificationsMapInput` via:
//
//	ProjectLevelNotificationsMap{ "key": ProjectLevelNotificationsArgs{...} }
type ProjectLevelNotificationsMapInput interface {
	pulumi.Input

	ToProjectLevelNotificationsMapOutput() ProjectLevelNotificationsMapOutput
	ToProjectLevelNotificationsMapOutputWithContext(context.Context) ProjectLevelNotificationsMapOutput
}

type ProjectLevelNotificationsMap map[string]ProjectLevelNotificationsInput

func (ProjectLevelNotificationsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectLevelNotifications)(nil)).Elem()
}

func (i ProjectLevelNotificationsMap) ToProjectLevelNotificationsMapOutput() ProjectLevelNotificationsMapOutput {
	return i.ToProjectLevelNotificationsMapOutputWithContext(context.Background())
}

func (i ProjectLevelNotificationsMap) ToProjectLevelNotificationsMapOutputWithContext(ctx context.Context) ProjectLevelNotificationsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectLevelNotificationsMapOutput)
}

type ProjectLevelNotificationsOutput struct{ *pulumi.OutputState }

func (ProjectLevelNotificationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectLevelNotifications)(nil)).Elem()
}

func (o ProjectLevelNotificationsOutput) ToProjectLevelNotificationsOutput() ProjectLevelNotificationsOutput {
	return o
}

func (o ProjectLevelNotificationsOutput) ToProjectLevelNotificationsOutputWithContext(ctx context.Context) ProjectLevelNotificationsOutput {
	return o
}

// Enable notifications for closed issues. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) CloseIssue() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.CloseIssue }).(pulumi.BoolOutput)
}

// Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) CloseMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.CloseMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) FailedPipeline() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.FailedPipeline }).(pulumi.BoolOutput)
}

// Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) FixedPipeline() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.FixedPipeline }).(pulumi.BoolOutput)
}

// Enable notifications for due issues. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) IssueDue() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.IssueDue }).(pulumi.BoolOutput)
}

// The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
func (o ProjectLevelNotificationsOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.StringOutput { return v.Level }).(pulumi.StringOutput)
}

// Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) MergeMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.MergeMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) MergeWhenPipelineSucceeds() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.MergeWhenPipelineSucceeds }).(pulumi.BoolOutput)
}

// Enable notifications for moved projects. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) MovedProject() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.MovedProject }).(pulumi.BoolOutput)
}

// Enable notifications for new issues. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) NewIssue() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.NewIssue }).(pulumi.BoolOutput)
}

// Enable notifications for new merge requests. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) NewMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.NewMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) NewNote() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.NewNote }).(pulumi.BoolOutput)
}

// The ID or URL-encoded path of a project where notifications will be configured.
func (o ProjectLevelNotificationsOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) PushToMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.PushToMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) ReassignIssue() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.ReassignIssue }).(pulumi.BoolOutput)
}

// Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) ReassignMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.ReassignMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for reopened issues. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) ReopenIssue() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.ReopenIssue }).(pulumi.BoolOutput)
}

// Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) ReopenMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.ReopenMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
func (o ProjectLevelNotificationsOutput) SuccessPipeline() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelNotifications) pulumi.BoolOutput { return v.SuccessPipeline }).(pulumi.BoolOutput)
}

type ProjectLevelNotificationsArrayOutput struct{ *pulumi.OutputState }

func (ProjectLevelNotificationsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectLevelNotifications)(nil)).Elem()
}

func (o ProjectLevelNotificationsArrayOutput) ToProjectLevelNotificationsArrayOutput() ProjectLevelNotificationsArrayOutput {
	return o
}

func (o ProjectLevelNotificationsArrayOutput) ToProjectLevelNotificationsArrayOutputWithContext(ctx context.Context) ProjectLevelNotificationsArrayOutput {
	return o
}

func (o ProjectLevelNotificationsArrayOutput) Index(i pulumi.IntInput) ProjectLevelNotificationsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectLevelNotifications {
		return vs[0].([]*ProjectLevelNotifications)[vs[1].(int)]
	}).(ProjectLevelNotificationsOutput)
}

type ProjectLevelNotificationsMapOutput struct{ *pulumi.OutputState }

func (ProjectLevelNotificationsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectLevelNotifications)(nil)).Elem()
}

func (o ProjectLevelNotificationsMapOutput) ToProjectLevelNotificationsMapOutput() ProjectLevelNotificationsMapOutput {
	return o
}

func (o ProjectLevelNotificationsMapOutput) ToProjectLevelNotificationsMapOutputWithContext(ctx context.Context) ProjectLevelNotificationsMapOutput {
	return o
}

func (o ProjectLevelNotificationsMapOutput) MapIndex(k pulumi.StringInput) ProjectLevelNotificationsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectLevelNotifications {
		return vs[0].(map[string]*ProjectLevelNotifications)[vs[1].(string)]
	}).(ProjectLevelNotificationsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLevelNotificationsInput)(nil)).Elem(), &ProjectLevelNotifications{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLevelNotificationsArrayInput)(nil)).Elem(), ProjectLevelNotificationsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLevelNotificationsMapInput)(nil)).Elem(), ProjectLevelNotificationsMap{})
	pulumi.RegisterOutputType(ProjectLevelNotificationsOutput{})
	pulumi.RegisterOutputType(ProjectLevelNotificationsArrayOutput{})
	pulumi.RegisterOutputType(ProjectLevelNotificationsMapOutput{})
}
