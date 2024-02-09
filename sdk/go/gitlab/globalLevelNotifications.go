// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GlobalLevelNotifications` resource allows to manage global notifications.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/notification_settings.html#group--project-level-notification-settings)
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
//			_, err := gitlab.NewGlobalLevelNotifications(ctx, "foo", &gitlab.GlobalLevelNotificationsArgs{
//				Level:           pulumi.String("custom"),
//				NewMergeRequest: pulumi.Bool(true),
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
// Note: You can import a global notification state using "gitlab" as the ID.
//
//	The ID will always be gitlab, because the global notificatio only exists
//
//	once per user
//
// ```sh
// $ pulumi import gitlab:index/globalLevelNotifications:GlobalLevelNotifications example gitlab
// ```
type GlobalLevelNotifications struct {
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

// NewGlobalLevelNotifications registers a new resource with the given unique name, arguments, and options.
func NewGlobalLevelNotifications(ctx *pulumi.Context,
	name string, args *GlobalLevelNotificationsArgs, opts ...pulumi.ResourceOption) (*GlobalLevelNotifications, error) {
	if args == nil {
		args = &GlobalLevelNotificationsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GlobalLevelNotifications
	err := ctx.RegisterResource("gitlab:index/globalLevelNotifications:GlobalLevelNotifications", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalLevelNotifications gets an existing GlobalLevelNotifications resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalLevelNotifications(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalLevelNotificationsState, opts ...pulumi.ResourceOption) (*GlobalLevelNotifications, error) {
	var resource GlobalLevelNotifications
	err := ctx.ReadResource("gitlab:index/globalLevelNotifications:GlobalLevelNotifications", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalLevelNotifications resources.
type globalLevelNotificationsState struct {
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

type GlobalLevelNotificationsState struct {
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

func (GlobalLevelNotificationsState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalLevelNotificationsState)(nil)).Elem()
}

type globalLevelNotificationsArgs struct {
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

// The set of arguments for constructing a GlobalLevelNotifications resource.
type GlobalLevelNotificationsArgs struct {
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

func (GlobalLevelNotificationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalLevelNotificationsArgs)(nil)).Elem()
}

type GlobalLevelNotificationsInput interface {
	pulumi.Input

	ToGlobalLevelNotificationsOutput() GlobalLevelNotificationsOutput
	ToGlobalLevelNotificationsOutputWithContext(ctx context.Context) GlobalLevelNotificationsOutput
}

func (*GlobalLevelNotifications) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalLevelNotifications)(nil)).Elem()
}

func (i *GlobalLevelNotifications) ToGlobalLevelNotificationsOutput() GlobalLevelNotificationsOutput {
	return i.ToGlobalLevelNotificationsOutputWithContext(context.Background())
}

func (i *GlobalLevelNotifications) ToGlobalLevelNotificationsOutputWithContext(ctx context.Context) GlobalLevelNotificationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalLevelNotificationsOutput)
}

// GlobalLevelNotificationsArrayInput is an input type that accepts GlobalLevelNotificationsArray and GlobalLevelNotificationsArrayOutput values.
// You can construct a concrete instance of `GlobalLevelNotificationsArrayInput` via:
//
//	GlobalLevelNotificationsArray{ GlobalLevelNotificationsArgs{...} }
type GlobalLevelNotificationsArrayInput interface {
	pulumi.Input

	ToGlobalLevelNotificationsArrayOutput() GlobalLevelNotificationsArrayOutput
	ToGlobalLevelNotificationsArrayOutputWithContext(context.Context) GlobalLevelNotificationsArrayOutput
}

type GlobalLevelNotificationsArray []GlobalLevelNotificationsInput

func (GlobalLevelNotificationsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalLevelNotifications)(nil)).Elem()
}

func (i GlobalLevelNotificationsArray) ToGlobalLevelNotificationsArrayOutput() GlobalLevelNotificationsArrayOutput {
	return i.ToGlobalLevelNotificationsArrayOutputWithContext(context.Background())
}

func (i GlobalLevelNotificationsArray) ToGlobalLevelNotificationsArrayOutputWithContext(ctx context.Context) GlobalLevelNotificationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalLevelNotificationsArrayOutput)
}

// GlobalLevelNotificationsMapInput is an input type that accepts GlobalLevelNotificationsMap and GlobalLevelNotificationsMapOutput values.
// You can construct a concrete instance of `GlobalLevelNotificationsMapInput` via:
//
//	GlobalLevelNotificationsMap{ "key": GlobalLevelNotificationsArgs{...} }
type GlobalLevelNotificationsMapInput interface {
	pulumi.Input

	ToGlobalLevelNotificationsMapOutput() GlobalLevelNotificationsMapOutput
	ToGlobalLevelNotificationsMapOutputWithContext(context.Context) GlobalLevelNotificationsMapOutput
}

type GlobalLevelNotificationsMap map[string]GlobalLevelNotificationsInput

func (GlobalLevelNotificationsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalLevelNotifications)(nil)).Elem()
}

func (i GlobalLevelNotificationsMap) ToGlobalLevelNotificationsMapOutput() GlobalLevelNotificationsMapOutput {
	return i.ToGlobalLevelNotificationsMapOutputWithContext(context.Background())
}

func (i GlobalLevelNotificationsMap) ToGlobalLevelNotificationsMapOutputWithContext(ctx context.Context) GlobalLevelNotificationsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalLevelNotificationsMapOutput)
}

type GlobalLevelNotificationsOutput struct{ *pulumi.OutputState }

func (GlobalLevelNotificationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalLevelNotifications)(nil)).Elem()
}

func (o GlobalLevelNotificationsOutput) ToGlobalLevelNotificationsOutput() GlobalLevelNotificationsOutput {
	return o
}

func (o GlobalLevelNotificationsOutput) ToGlobalLevelNotificationsOutputWithContext(ctx context.Context) GlobalLevelNotificationsOutput {
	return o
}

// Enable notifications for closed issues. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) CloseIssue() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.CloseIssue }).(pulumi.BoolOutput)
}

// Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) CloseMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.CloseMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) FailedPipeline() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.FailedPipeline }).(pulumi.BoolOutput)
}

// Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) FixedPipeline() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.FixedPipeline }).(pulumi.BoolOutput)
}

// Enable notifications for due issues. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) IssueDue() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.IssueDue }).(pulumi.BoolOutput)
}

// The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
func (o GlobalLevelNotificationsOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.StringOutput { return v.Level }).(pulumi.StringOutput)
}

// Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) MergeMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.MergeMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) MergeWhenPipelineSucceeds() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.MergeWhenPipelineSucceeds }).(pulumi.BoolOutput)
}

// Enable notifications for moved projects. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) MovedProject() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.MovedProject }).(pulumi.BoolOutput)
}

// Enable notifications for new issues. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) NewIssue() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.NewIssue }).(pulumi.BoolOutput)
}

// Enable notifications for new merge requests. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) NewMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.NewMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) NewNote() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.NewNote }).(pulumi.BoolOutput)
}

// Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) PushToMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.PushToMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) ReassignIssue() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.ReassignIssue }).(pulumi.BoolOutput)
}

// Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) ReassignMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.ReassignMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for reopened issues. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) ReopenIssue() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.ReopenIssue }).(pulumi.BoolOutput)
}

// Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) ReopenMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.ReopenMergeRequest }).(pulumi.BoolOutput)
}

// Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
func (o GlobalLevelNotificationsOutput) SuccessPipeline() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalLevelNotifications) pulumi.BoolOutput { return v.SuccessPipeline }).(pulumi.BoolOutput)
}

type GlobalLevelNotificationsArrayOutput struct{ *pulumi.OutputState }

func (GlobalLevelNotificationsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalLevelNotifications)(nil)).Elem()
}

func (o GlobalLevelNotificationsArrayOutput) ToGlobalLevelNotificationsArrayOutput() GlobalLevelNotificationsArrayOutput {
	return o
}

func (o GlobalLevelNotificationsArrayOutput) ToGlobalLevelNotificationsArrayOutputWithContext(ctx context.Context) GlobalLevelNotificationsArrayOutput {
	return o
}

func (o GlobalLevelNotificationsArrayOutput) Index(i pulumi.IntInput) GlobalLevelNotificationsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GlobalLevelNotifications {
		return vs[0].([]*GlobalLevelNotifications)[vs[1].(int)]
	}).(GlobalLevelNotificationsOutput)
}

type GlobalLevelNotificationsMapOutput struct{ *pulumi.OutputState }

func (GlobalLevelNotificationsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalLevelNotifications)(nil)).Elem()
}

func (o GlobalLevelNotificationsMapOutput) ToGlobalLevelNotificationsMapOutput() GlobalLevelNotificationsMapOutput {
	return o
}

func (o GlobalLevelNotificationsMapOutput) ToGlobalLevelNotificationsMapOutputWithContext(ctx context.Context) GlobalLevelNotificationsMapOutput {
	return o
}

func (o GlobalLevelNotificationsMapOutput) MapIndex(k pulumi.StringInput) GlobalLevelNotificationsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GlobalLevelNotifications {
		return vs[0].(map[string]*GlobalLevelNotifications)[vs[1].(string)]
	}).(GlobalLevelNotificationsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalLevelNotificationsInput)(nil)).Elem(), &GlobalLevelNotifications{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalLevelNotificationsArrayInput)(nil)).Elem(), GlobalLevelNotificationsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalLevelNotificationsMapInput)(nil)).Elem(), GlobalLevelNotificationsMap{})
	pulumi.RegisterOutputType(GlobalLevelNotificationsOutput{})
	pulumi.RegisterOutputType(GlobalLevelNotificationsArrayOutput{})
	pulumi.RegisterOutputType(GlobalLevelNotificationsMapOutput{})
}
