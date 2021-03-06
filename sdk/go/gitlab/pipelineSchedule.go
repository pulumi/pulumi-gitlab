// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # gitlab\_pipeline\_schedule
//
// This resource allows you to create and manage pipeline schedules.
// For further information on clusters, consult the [gitlab
// documentation](https://docs.gitlab.com/ce/user/project/pipelines/schedules.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gitlab.NewPipelineSchedule(ctx, "example", &gitlab.PipelineScheduleArgs{
// 			Cron:        pulumi.String("0 1 * * *"),
// 			Description: pulumi.String("Used to schedule builds"),
// 			Project:     pulumi.String("12345"),
// 			Ref:         pulumi.String("master"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PipelineSchedule struct {
	pulumi.CustomResourceState

	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// The cron (e.g. `0 1 * * *`).
	Cron pulumi.StringOutput `pulumi:"cron"`
	// The timezone.
	CronTimezone pulumi.StringPtrOutput `pulumi:"cronTimezone"`
	// The description of the pipeline schedule.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name or id of the project to add the schedule to.
	Project pulumi.StringOutput `pulumi:"project"`
	// The branch/tag name to be triggered.
	Ref pulumi.StringOutput `pulumi:"ref"`
}

// NewPipelineSchedule registers a new resource with the given unique name, arguments, and options.
func NewPipelineSchedule(ctx *pulumi.Context,
	name string, args *PipelineScheduleArgs, opts ...pulumi.ResourceOption) (*PipelineSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cron == nil {
		return nil, errors.New("invalid value for required argument 'Cron'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Ref == nil {
		return nil, errors.New("invalid value for required argument 'Ref'")
	}
	var resource PipelineSchedule
	err := ctx.RegisterResource("gitlab:index/pipelineSchedule:PipelineSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipelineSchedule gets an existing PipelineSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineScheduleState, opts ...pulumi.ResourceOption) (*PipelineSchedule, error) {
	var resource PipelineSchedule
	err := ctx.ReadResource("gitlab:index/pipelineSchedule:PipelineSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipelineSchedule resources.
type pipelineScheduleState struct {
	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	Active *bool `pulumi:"active"`
	// The cron (e.g. `0 1 * * *`).
	Cron *string `pulumi:"cron"`
	// The timezone.
	CronTimezone *string `pulumi:"cronTimezone"`
	// The description of the pipeline schedule.
	Description *string `pulumi:"description"`
	// The name or id of the project to add the schedule to.
	Project *string `pulumi:"project"`
	// The branch/tag name to be triggered.
	Ref *string `pulumi:"ref"`
}

type PipelineScheduleState struct {
	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	Active pulumi.BoolPtrInput
	// The cron (e.g. `0 1 * * *`).
	Cron pulumi.StringPtrInput
	// The timezone.
	CronTimezone pulumi.StringPtrInput
	// The description of the pipeline schedule.
	Description pulumi.StringPtrInput
	// The name or id of the project to add the schedule to.
	Project pulumi.StringPtrInput
	// The branch/tag name to be triggered.
	Ref pulumi.StringPtrInput
}

func (PipelineScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineScheduleState)(nil)).Elem()
}

type pipelineScheduleArgs struct {
	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	Active *bool `pulumi:"active"`
	// The cron (e.g. `0 1 * * *`).
	Cron string `pulumi:"cron"`
	// The timezone.
	CronTimezone *string `pulumi:"cronTimezone"`
	// The description of the pipeline schedule.
	Description string `pulumi:"description"`
	// The name or id of the project to add the schedule to.
	Project string `pulumi:"project"`
	// The branch/tag name to be triggered.
	Ref string `pulumi:"ref"`
}

// The set of arguments for constructing a PipelineSchedule resource.
type PipelineScheduleArgs struct {
	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	Active pulumi.BoolPtrInput
	// The cron (e.g. `0 1 * * *`).
	Cron pulumi.StringInput
	// The timezone.
	CronTimezone pulumi.StringPtrInput
	// The description of the pipeline schedule.
	Description pulumi.StringInput
	// The name or id of the project to add the schedule to.
	Project pulumi.StringInput
	// The branch/tag name to be triggered.
	Ref pulumi.StringInput
}

func (PipelineScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineScheduleArgs)(nil)).Elem()
}

type PipelineScheduleInput interface {
	pulumi.Input

	ToPipelineScheduleOutput() PipelineScheduleOutput
	ToPipelineScheduleOutputWithContext(ctx context.Context) PipelineScheduleOutput
}

func (*PipelineSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSchedule)(nil))
}

func (i *PipelineSchedule) ToPipelineScheduleOutput() PipelineScheduleOutput {
	return i.ToPipelineScheduleOutputWithContext(context.Background())
}

func (i *PipelineSchedule) ToPipelineScheduleOutputWithContext(ctx context.Context) PipelineScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineScheduleOutput)
}

func (i *PipelineSchedule) ToPipelineSchedulePtrOutput() PipelineSchedulePtrOutput {
	return i.ToPipelineSchedulePtrOutputWithContext(context.Background())
}

func (i *PipelineSchedule) ToPipelineSchedulePtrOutputWithContext(ctx context.Context) PipelineSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSchedulePtrOutput)
}

type PipelineSchedulePtrInput interface {
	pulumi.Input

	ToPipelineSchedulePtrOutput() PipelineSchedulePtrOutput
	ToPipelineSchedulePtrOutputWithContext(ctx context.Context) PipelineSchedulePtrOutput
}

type pipelineSchedulePtrType PipelineScheduleArgs

func (*pipelineSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSchedule)(nil))
}

func (i *pipelineSchedulePtrType) ToPipelineSchedulePtrOutput() PipelineSchedulePtrOutput {
	return i.ToPipelineSchedulePtrOutputWithContext(context.Background())
}

func (i *pipelineSchedulePtrType) ToPipelineSchedulePtrOutputWithContext(ctx context.Context) PipelineSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSchedulePtrOutput)
}

// PipelineScheduleArrayInput is an input type that accepts PipelineScheduleArray and PipelineScheduleArrayOutput values.
// You can construct a concrete instance of `PipelineScheduleArrayInput` via:
//
//          PipelineScheduleArray{ PipelineScheduleArgs{...} }
type PipelineScheduleArrayInput interface {
	pulumi.Input

	ToPipelineScheduleArrayOutput() PipelineScheduleArrayOutput
	ToPipelineScheduleArrayOutputWithContext(context.Context) PipelineScheduleArrayOutput
}

type PipelineScheduleArray []PipelineScheduleInput

func (PipelineScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PipelineSchedule)(nil))
}

func (i PipelineScheduleArray) ToPipelineScheduleArrayOutput() PipelineScheduleArrayOutput {
	return i.ToPipelineScheduleArrayOutputWithContext(context.Background())
}

func (i PipelineScheduleArray) ToPipelineScheduleArrayOutputWithContext(ctx context.Context) PipelineScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineScheduleArrayOutput)
}

// PipelineScheduleMapInput is an input type that accepts PipelineScheduleMap and PipelineScheduleMapOutput values.
// You can construct a concrete instance of `PipelineScheduleMapInput` via:
//
//          PipelineScheduleMap{ "key": PipelineScheduleArgs{...} }
type PipelineScheduleMapInput interface {
	pulumi.Input

	ToPipelineScheduleMapOutput() PipelineScheduleMapOutput
	ToPipelineScheduleMapOutputWithContext(context.Context) PipelineScheduleMapOutput
}

type PipelineScheduleMap map[string]PipelineScheduleInput

func (PipelineScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PipelineSchedule)(nil))
}

func (i PipelineScheduleMap) ToPipelineScheduleMapOutput() PipelineScheduleMapOutput {
	return i.ToPipelineScheduleMapOutputWithContext(context.Background())
}

func (i PipelineScheduleMap) ToPipelineScheduleMapOutputWithContext(ctx context.Context) PipelineScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineScheduleMapOutput)
}

type PipelineScheduleOutput struct {
	*pulumi.OutputState
}

func (PipelineScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSchedule)(nil))
}

func (o PipelineScheduleOutput) ToPipelineScheduleOutput() PipelineScheduleOutput {
	return o
}

func (o PipelineScheduleOutput) ToPipelineScheduleOutputWithContext(ctx context.Context) PipelineScheduleOutput {
	return o
}

func (o PipelineScheduleOutput) ToPipelineSchedulePtrOutput() PipelineSchedulePtrOutput {
	return o.ToPipelineSchedulePtrOutputWithContext(context.Background())
}

func (o PipelineScheduleOutput) ToPipelineSchedulePtrOutputWithContext(ctx context.Context) PipelineSchedulePtrOutput {
	return o.ApplyT(func(v PipelineSchedule) *PipelineSchedule {
		return &v
	}).(PipelineSchedulePtrOutput)
}

type PipelineSchedulePtrOutput struct {
	*pulumi.OutputState
}

func (PipelineSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSchedule)(nil))
}

func (o PipelineSchedulePtrOutput) ToPipelineSchedulePtrOutput() PipelineSchedulePtrOutput {
	return o
}

func (o PipelineSchedulePtrOutput) ToPipelineSchedulePtrOutputWithContext(ctx context.Context) PipelineSchedulePtrOutput {
	return o
}

type PipelineScheduleArrayOutput struct{ *pulumi.OutputState }

func (PipelineScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PipelineSchedule)(nil))
}

func (o PipelineScheduleArrayOutput) ToPipelineScheduleArrayOutput() PipelineScheduleArrayOutput {
	return o
}

func (o PipelineScheduleArrayOutput) ToPipelineScheduleArrayOutputWithContext(ctx context.Context) PipelineScheduleArrayOutput {
	return o
}

func (o PipelineScheduleArrayOutput) Index(i pulumi.IntInput) PipelineScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PipelineSchedule {
		return vs[0].([]PipelineSchedule)[vs[1].(int)]
	}).(PipelineScheduleOutput)
}

type PipelineScheduleMapOutput struct{ *pulumi.OutputState }

func (PipelineScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PipelineSchedule)(nil))
}

func (o PipelineScheduleMapOutput) ToPipelineScheduleMapOutput() PipelineScheduleMapOutput {
	return o
}

func (o PipelineScheduleMapOutput) ToPipelineScheduleMapOutputWithContext(ctx context.Context) PipelineScheduleMapOutput {
	return o
}

func (o PipelineScheduleMapOutput) MapIndex(k pulumi.StringInput) PipelineScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PipelineSchedule {
		return vs[0].(map[string]PipelineSchedule)[vs[1].(string)]
	}).(PipelineScheduleOutput)
}

func init() {
	pulumi.RegisterOutputType(PipelineScheduleOutput{})
	pulumi.RegisterOutputType(PipelineSchedulePtrOutput{})
	pulumi.RegisterOutputType(PipelineScheduleArrayOutput{})
	pulumi.RegisterOutputType(PipelineScheduleMapOutput{})
}
