// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `PipelineScheduleVariable` resource allows to manage the lifecycle of a variable for a pipeline schedule.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_schedules.html#pipeline-schedule-variables)
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
//			examplePipelineSchedule, err := gitlab.NewPipelineSchedule(ctx, "examplePipelineSchedule", &gitlab.PipelineScheduleArgs{
//				Project:     pulumi.String("12345"),
//				Description: pulumi.String("Used to schedule builds"),
//				Ref:         pulumi.String("master"),
//				Cron:        pulumi.String("0 1 * * *"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewPipelineScheduleVariable(ctx, "examplePipelineScheduleVariable", &gitlab.PipelineScheduleVariableArgs{
//				Project:            examplePipelineSchedule.Project,
//				PipelineScheduleId: examplePipelineSchedule.PipelineScheduleId,
//				Key:                pulumi.String("EXAMPLE_KEY"),
//				Value:              pulumi.String("example"),
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
// Pipeline schedule variables can be imported using an id made up of `project_id:pipeline_schedule_id:key`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable example 123456789:13:mykey
//
// ```
type PipelineScheduleVariable struct {
	pulumi.CustomResourceState

	// Name of the variable.
	Key pulumi.StringOutput `pulumi:"key"`
	// The id of the pipeline schedule.
	PipelineScheduleId pulumi.IntOutput `pulumi:"pipelineScheduleId"`
	// The id of the project to add the schedule to.
	Project pulumi.StringOutput `pulumi:"project"`
	// Value of the variable.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewPipelineScheduleVariable registers a new resource with the given unique name, arguments, and options.
func NewPipelineScheduleVariable(ctx *pulumi.Context,
	name string, args *PipelineScheduleVariableArgs, opts ...pulumi.ResourceOption) (*PipelineScheduleVariable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.PipelineScheduleId == nil {
		return nil, errors.New("invalid value for required argument 'PipelineScheduleId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PipelineScheduleVariable
	err := ctx.RegisterResource("gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipelineScheduleVariable gets an existing PipelineScheduleVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineScheduleVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineScheduleVariableState, opts ...pulumi.ResourceOption) (*PipelineScheduleVariable, error) {
	var resource PipelineScheduleVariable
	err := ctx.ReadResource("gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipelineScheduleVariable resources.
type pipelineScheduleVariableState struct {
	// Name of the variable.
	Key *string `pulumi:"key"`
	// The id of the pipeline schedule.
	PipelineScheduleId *int `pulumi:"pipelineScheduleId"`
	// The id of the project to add the schedule to.
	Project *string `pulumi:"project"`
	// Value of the variable.
	Value *string `pulumi:"value"`
}

type PipelineScheduleVariableState struct {
	// Name of the variable.
	Key pulumi.StringPtrInput
	// The id of the pipeline schedule.
	PipelineScheduleId pulumi.IntPtrInput
	// The id of the project to add the schedule to.
	Project pulumi.StringPtrInput
	// Value of the variable.
	Value pulumi.StringPtrInput
}

func (PipelineScheduleVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineScheduleVariableState)(nil)).Elem()
}

type pipelineScheduleVariableArgs struct {
	// Name of the variable.
	Key string `pulumi:"key"`
	// The id of the pipeline schedule.
	PipelineScheduleId int `pulumi:"pipelineScheduleId"`
	// The id of the project to add the schedule to.
	Project string `pulumi:"project"`
	// Value of the variable.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a PipelineScheduleVariable resource.
type PipelineScheduleVariableArgs struct {
	// Name of the variable.
	Key pulumi.StringInput
	// The id of the pipeline schedule.
	PipelineScheduleId pulumi.IntInput
	// The id of the project to add the schedule to.
	Project pulumi.StringInput
	// Value of the variable.
	Value pulumi.StringInput
}

func (PipelineScheduleVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineScheduleVariableArgs)(nil)).Elem()
}

type PipelineScheduleVariableInput interface {
	pulumi.Input

	ToPipelineScheduleVariableOutput() PipelineScheduleVariableOutput
	ToPipelineScheduleVariableOutputWithContext(ctx context.Context) PipelineScheduleVariableOutput
}

func (*PipelineScheduleVariable) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineScheduleVariable)(nil)).Elem()
}

func (i *PipelineScheduleVariable) ToPipelineScheduleVariableOutput() PipelineScheduleVariableOutput {
	return i.ToPipelineScheduleVariableOutputWithContext(context.Background())
}

func (i *PipelineScheduleVariable) ToPipelineScheduleVariableOutputWithContext(ctx context.Context) PipelineScheduleVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineScheduleVariableOutput)
}

// PipelineScheduleVariableArrayInput is an input type that accepts PipelineScheduleVariableArray and PipelineScheduleVariableArrayOutput values.
// You can construct a concrete instance of `PipelineScheduleVariableArrayInput` via:
//
//	PipelineScheduleVariableArray{ PipelineScheduleVariableArgs{...} }
type PipelineScheduleVariableArrayInput interface {
	pulumi.Input

	ToPipelineScheduleVariableArrayOutput() PipelineScheduleVariableArrayOutput
	ToPipelineScheduleVariableArrayOutputWithContext(context.Context) PipelineScheduleVariableArrayOutput
}

type PipelineScheduleVariableArray []PipelineScheduleVariableInput

func (PipelineScheduleVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PipelineScheduleVariable)(nil)).Elem()
}

func (i PipelineScheduleVariableArray) ToPipelineScheduleVariableArrayOutput() PipelineScheduleVariableArrayOutput {
	return i.ToPipelineScheduleVariableArrayOutputWithContext(context.Background())
}

func (i PipelineScheduleVariableArray) ToPipelineScheduleVariableArrayOutputWithContext(ctx context.Context) PipelineScheduleVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineScheduleVariableArrayOutput)
}

// PipelineScheduleVariableMapInput is an input type that accepts PipelineScheduleVariableMap and PipelineScheduleVariableMapOutput values.
// You can construct a concrete instance of `PipelineScheduleVariableMapInput` via:
//
//	PipelineScheduleVariableMap{ "key": PipelineScheduleVariableArgs{...} }
type PipelineScheduleVariableMapInput interface {
	pulumi.Input

	ToPipelineScheduleVariableMapOutput() PipelineScheduleVariableMapOutput
	ToPipelineScheduleVariableMapOutputWithContext(context.Context) PipelineScheduleVariableMapOutput
}

type PipelineScheduleVariableMap map[string]PipelineScheduleVariableInput

func (PipelineScheduleVariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PipelineScheduleVariable)(nil)).Elem()
}

func (i PipelineScheduleVariableMap) ToPipelineScheduleVariableMapOutput() PipelineScheduleVariableMapOutput {
	return i.ToPipelineScheduleVariableMapOutputWithContext(context.Background())
}

func (i PipelineScheduleVariableMap) ToPipelineScheduleVariableMapOutputWithContext(ctx context.Context) PipelineScheduleVariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineScheduleVariableMapOutput)
}

type PipelineScheduleVariableOutput struct{ *pulumi.OutputState }

func (PipelineScheduleVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineScheduleVariable)(nil)).Elem()
}

func (o PipelineScheduleVariableOutput) ToPipelineScheduleVariableOutput() PipelineScheduleVariableOutput {
	return o
}

func (o PipelineScheduleVariableOutput) ToPipelineScheduleVariableOutputWithContext(ctx context.Context) PipelineScheduleVariableOutput {
	return o
}

// Name of the variable.
func (o PipelineScheduleVariableOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineScheduleVariable) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The id of the pipeline schedule.
func (o PipelineScheduleVariableOutput) PipelineScheduleId() pulumi.IntOutput {
	return o.ApplyT(func(v *PipelineScheduleVariable) pulumi.IntOutput { return v.PipelineScheduleId }).(pulumi.IntOutput)
}

// The id of the project to add the schedule to.
func (o PipelineScheduleVariableOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineScheduleVariable) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Value of the variable.
func (o PipelineScheduleVariableOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineScheduleVariable) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type PipelineScheduleVariableArrayOutput struct{ *pulumi.OutputState }

func (PipelineScheduleVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PipelineScheduleVariable)(nil)).Elem()
}

func (o PipelineScheduleVariableArrayOutput) ToPipelineScheduleVariableArrayOutput() PipelineScheduleVariableArrayOutput {
	return o
}

func (o PipelineScheduleVariableArrayOutput) ToPipelineScheduleVariableArrayOutputWithContext(ctx context.Context) PipelineScheduleVariableArrayOutput {
	return o
}

func (o PipelineScheduleVariableArrayOutput) Index(i pulumi.IntInput) PipelineScheduleVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PipelineScheduleVariable {
		return vs[0].([]*PipelineScheduleVariable)[vs[1].(int)]
	}).(PipelineScheduleVariableOutput)
}

type PipelineScheduleVariableMapOutput struct{ *pulumi.OutputState }

func (PipelineScheduleVariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PipelineScheduleVariable)(nil)).Elem()
}

func (o PipelineScheduleVariableMapOutput) ToPipelineScheduleVariableMapOutput() PipelineScheduleVariableMapOutput {
	return o
}

func (o PipelineScheduleVariableMapOutput) ToPipelineScheduleVariableMapOutputWithContext(ctx context.Context) PipelineScheduleVariableMapOutput {
	return o
}

func (o PipelineScheduleVariableMapOutput) MapIndex(k pulumi.StringInput) PipelineScheduleVariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PipelineScheduleVariable {
		return vs[0].(map[string]*PipelineScheduleVariable)[vs[1].(string)]
	}).(PipelineScheduleVariableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineScheduleVariableInput)(nil)).Elem(), &PipelineScheduleVariable{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineScheduleVariableArrayInput)(nil)).Elem(), PipelineScheduleVariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineScheduleVariableMapInput)(nil)).Elem(), PipelineScheduleVariableMap{})
	pulumi.RegisterOutputType(PipelineScheduleVariableOutput{})
	pulumi.RegisterOutputType(PipelineScheduleVariableArrayOutput{})
	pulumi.RegisterOutputType(PipelineScheduleVariableMapOutput{})
}
