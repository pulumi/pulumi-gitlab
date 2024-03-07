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

// The `PipelineTrigger` resource allows to manage the lifecycle of a pipeline trigger.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_triggers.html)
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := gitlab.NewPipelineTrigger(ctx, "example", &gitlab.PipelineTriggerArgs{
//				Description: pulumi.String("Used to trigger builds"),
//				Project:     pulumi.String("12345"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// GitLab pipeline triggers can be imported using an id made up of `{project_id}:{pipeline_trigger_id}`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/pipelineTrigger:PipelineTrigger test 1:3
// ```
type PipelineTrigger struct {
	pulumi.CustomResourceState

	// The description of the pipeline trigger.
	Description pulumi.StringOutput `pulumi:"description"`
	// The pipeline trigger id.
	PipelineTriggerId pulumi.IntOutput `pulumi:"pipelineTriggerId"`
	// The name or id of the project to add the trigger to.
	Project pulumi.StringOutput `pulumi:"project"`
	// The pipeline trigger token.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewPipelineTrigger registers a new resource with the given unique name, arguments, and options.
func NewPipelineTrigger(ctx *pulumi.Context,
	name string, args *PipelineTriggerArgs, opts ...pulumi.ResourceOption) (*PipelineTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PipelineTrigger
	err := ctx.RegisterResource("gitlab:index/pipelineTrigger:PipelineTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipelineTrigger gets an existing PipelineTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineTriggerState, opts ...pulumi.ResourceOption) (*PipelineTrigger, error) {
	var resource PipelineTrigger
	err := ctx.ReadResource("gitlab:index/pipelineTrigger:PipelineTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipelineTrigger resources.
type pipelineTriggerState struct {
	// The description of the pipeline trigger.
	Description *string `pulumi:"description"`
	// The pipeline trigger id.
	PipelineTriggerId *int `pulumi:"pipelineTriggerId"`
	// The name or id of the project to add the trigger to.
	Project *string `pulumi:"project"`
	// The pipeline trigger token.
	Token *string `pulumi:"token"`
}

type PipelineTriggerState struct {
	// The description of the pipeline trigger.
	Description pulumi.StringPtrInput
	// The pipeline trigger id.
	PipelineTriggerId pulumi.IntPtrInput
	// The name or id of the project to add the trigger to.
	Project pulumi.StringPtrInput
	// The pipeline trigger token.
	Token pulumi.StringPtrInput
}

func (PipelineTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineTriggerState)(nil)).Elem()
}

type pipelineTriggerArgs struct {
	// The description of the pipeline trigger.
	Description string `pulumi:"description"`
	// The name or id of the project to add the trigger to.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a PipelineTrigger resource.
type PipelineTriggerArgs struct {
	// The description of the pipeline trigger.
	Description pulumi.StringInput
	// The name or id of the project to add the trigger to.
	Project pulumi.StringInput
}

func (PipelineTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineTriggerArgs)(nil)).Elem()
}

type PipelineTriggerInput interface {
	pulumi.Input

	ToPipelineTriggerOutput() PipelineTriggerOutput
	ToPipelineTriggerOutputWithContext(ctx context.Context) PipelineTriggerOutput
}

func (*PipelineTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTrigger)(nil)).Elem()
}

func (i *PipelineTrigger) ToPipelineTriggerOutput() PipelineTriggerOutput {
	return i.ToPipelineTriggerOutputWithContext(context.Background())
}

func (i *PipelineTrigger) ToPipelineTriggerOutputWithContext(ctx context.Context) PipelineTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerOutput)
}

// PipelineTriggerArrayInput is an input type that accepts PipelineTriggerArray and PipelineTriggerArrayOutput values.
// You can construct a concrete instance of `PipelineTriggerArrayInput` via:
//
//	PipelineTriggerArray{ PipelineTriggerArgs{...} }
type PipelineTriggerArrayInput interface {
	pulumi.Input

	ToPipelineTriggerArrayOutput() PipelineTriggerArrayOutput
	ToPipelineTriggerArrayOutputWithContext(context.Context) PipelineTriggerArrayOutput
}

type PipelineTriggerArray []PipelineTriggerInput

func (PipelineTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PipelineTrigger)(nil)).Elem()
}

func (i PipelineTriggerArray) ToPipelineTriggerArrayOutput() PipelineTriggerArrayOutput {
	return i.ToPipelineTriggerArrayOutputWithContext(context.Background())
}

func (i PipelineTriggerArray) ToPipelineTriggerArrayOutputWithContext(ctx context.Context) PipelineTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerArrayOutput)
}

// PipelineTriggerMapInput is an input type that accepts PipelineTriggerMap and PipelineTriggerMapOutput values.
// You can construct a concrete instance of `PipelineTriggerMapInput` via:
//
//	PipelineTriggerMap{ "key": PipelineTriggerArgs{...} }
type PipelineTriggerMapInput interface {
	pulumi.Input

	ToPipelineTriggerMapOutput() PipelineTriggerMapOutput
	ToPipelineTriggerMapOutputWithContext(context.Context) PipelineTriggerMapOutput
}

type PipelineTriggerMap map[string]PipelineTriggerInput

func (PipelineTriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PipelineTrigger)(nil)).Elem()
}

func (i PipelineTriggerMap) ToPipelineTriggerMapOutput() PipelineTriggerMapOutput {
	return i.ToPipelineTriggerMapOutputWithContext(context.Background())
}

func (i PipelineTriggerMap) ToPipelineTriggerMapOutputWithContext(ctx context.Context) PipelineTriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerMapOutput)
}

type PipelineTriggerOutput struct{ *pulumi.OutputState }

func (PipelineTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTrigger)(nil)).Elem()
}

func (o PipelineTriggerOutput) ToPipelineTriggerOutput() PipelineTriggerOutput {
	return o
}

func (o PipelineTriggerOutput) ToPipelineTriggerOutputWithContext(ctx context.Context) PipelineTriggerOutput {
	return o
}

// The description of the pipeline trigger.
func (o PipelineTriggerOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineTrigger) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The pipeline trigger id.
func (o PipelineTriggerOutput) PipelineTriggerId() pulumi.IntOutput {
	return o.ApplyT(func(v *PipelineTrigger) pulumi.IntOutput { return v.PipelineTriggerId }).(pulumi.IntOutput)
}

// The name or id of the project to add the trigger to.
func (o PipelineTriggerOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineTrigger) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The pipeline trigger token.
func (o PipelineTriggerOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineTrigger) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type PipelineTriggerArrayOutput struct{ *pulumi.OutputState }

func (PipelineTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PipelineTrigger)(nil)).Elem()
}

func (o PipelineTriggerArrayOutput) ToPipelineTriggerArrayOutput() PipelineTriggerArrayOutput {
	return o
}

func (o PipelineTriggerArrayOutput) ToPipelineTriggerArrayOutputWithContext(ctx context.Context) PipelineTriggerArrayOutput {
	return o
}

func (o PipelineTriggerArrayOutput) Index(i pulumi.IntInput) PipelineTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PipelineTrigger {
		return vs[0].([]*PipelineTrigger)[vs[1].(int)]
	}).(PipelineTriggerOutput)
}

type PipelineTriggerMapOutput struct{ *pulumi.OutputState }

func (PipelineTriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PipelineTrigger)(nil)).Elem()
}

func (o PipelineTriggerMapOutput) ToPipelineTriggerMapOutput() PipelineTriggerMapOutput {
	return o
}

func (o PipelineTriggerMapOutput) ToPipelineTriggerMapOutputWithContext(ctx context.Context) PipelineTriggerMapOutput {
	return o
}

func (o PipelineTriggerMapOutput) MapIndex(k pulumi.StringInput) PipelineTriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PipelineTrigger {
		return vs[0].(map[string]*PipelineTrigger)[vs[1].(string)]
	}).(PipelineTriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineTriggerInput)(nil)).Elem(), &PipelineTrigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineTriggerArrayInput)(nil)).Elem(), PipelineTriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineTriggerMapInput)(nil)).Elem(), PipelineTriggerMap{})
	pulumi.RegisterOutputType(PipelineTriggerOutput{})
	pulumi.RegisterOutputType(PipelineTriggerArrayOutput{})
	pulumi.RegisterOutputType(PipelineTriggerMapOutput{})
}
