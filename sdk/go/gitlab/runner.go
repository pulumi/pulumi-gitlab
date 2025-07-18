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

// The `Runner` resource allows to manage the lifecycle of a runner.
//
// A runner can either be registered at an instance level or group level.
// The runner will be registered at a group level if the token used is from a group, or at an instance level if the token used is for the instance.
//
// ~ > Using this resource will register a runner using the deprecated `registrationToken` flow. To use the new `authenticationToken` flow instead,
// use the `UserRunner` resource!
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/runners/#register-a-new-runner)
//
// ## Import
//
// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_runner`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_runner.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Importing using the CLI is supported with the following syntax:
//
// # A GitLab Runner can be imported using the runner's ID, eg
//
// ```sh
// $ pulumi import gitlab:index/runner:Runner this 1
// ```
type Runner struct {
	pulumi.CustomResourceState

	// The accessLevel of the runner. Valid values are: `notProtected`, `refProtected`.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// The authentication token used for building a config.toml file. This value is not present when imported.
	AuthenticationToken pulumi.StringOutput `pulumi:"authenticationToken"`
	// The runner's description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the runner should be locked for current project.
	Locked pulumi.BoolOutput `pulumi:"locked"`
	// Free-form maintenance notes for the runner (1024 characters).
	MaintenanceNote pulumi.StringPtrOutput `pulumi:"maintenanceNote"`
	// Maximum timeout set when this runner handles the job.
	MaximumTimeout pulumi.IntPtrOutput `pulumi:"maximumTimeout"`
	// Whether the runner should ignore new jobs.
	Paused pulumi.BoolOutput `pulumi:"paused"`
	// The registration token used to register the runner.
	RegistrationToken pulumi.StringOutput `pulumi:"registrationToken"`
	// Whether the runner should handle untagged jobs.
	RunUntagged pulumi.BoolOutput `pulumi:"runUntagged"`
	// The status of runners to show, one of: online and offline. active and paused are also possible values
	// 			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
	Status pulumi.StringOutput `pulumi:"status"`
	// List of runner’s tags.
	TagLists pulumi.StringArrayOutput `pulumi:"tagLists"`
}

// NewRunner registers a new resource with the given unique name, arguments, and options.
func NewRunner(ctx *pulumi.Context,
	name string, args *RunnerArgs, opts ...pulumi.ResourceOption) (*Runner, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistrationToken == nil {
		return nil, errors.New("invalid value for required argument 'RegistrationToken'")
	}
	if args.RegistrationToken != nil {
		args.RegistrationToken = pulumi.ToSecret(args.RegistrationToken).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authenticationToken",
		"registrationToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Runner
	err := ctx.RegisterResource("gitlab:index/runner:Runner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRunner gets an existing Runner resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRunner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RunnerState, opts ...pulumi.ResourceOption) (*Runner, error) {
	var resource Runner
	err := ctx.ReadResource("gitlab:index/runner:Runner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Runner resources.
type runnerState struct {
	// The accessLevel of the runner. Valid values are: `notProtected`, `refProtected`.
	AccessLevel *string `pulumi:"accessLevel"`
	// The authentication token used for building a config.toml file. This value is not present when imported.
	AuthenticationToken *string `pulumi:"authenticationToken"`
	// The runner's description.
	Description *string `pulumi:"description"`
	// Whether the runner should be locked for current project.
	Locked *bool `pulumi:"locked"`
	// Free-form maintenance notes for the runner (1024 characters).
	MaintenanceNote *string `pulumi:"maintenanceNote"`
	// Maximum timeout set when this runner handles the job.
	MaximumTimeout *int `pulumi:"maximumTimeout"`
	// Whether the runner should ignore new jobs.
	Paused *bool `pulumi:"paused"`
	// The registration token used to register the runner.
	RegistrationToken *string `pulumi:"registrationToken"`
	// Whether the runner should handle untagged jobs.
	RunUntagged *bool `pulumi:"runUntagged"`
	// The status of runners to show, one of: online and offline. active and paused are also possible values
	// 			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
	Status *string `pulumi:"status"`
	// List of runner’s tags.
	TagLists []string `pulumi:"tagLists"`
}

type RunnerState struct {
	// The accessLevel of the runner. Valid values are: `notProtected`, `refProtected`.
	AccessLevel pulumi.StringPtrInput
	// The authentication token used for building a config.toml file. This value is not present when imported.
	AuthenticationToken pulumi.StringPtrInput
	// The runner's description.
	Description pulumi.StringPtrInput
	// Whether the runner should be locked for current project.
	Locked pulumi.BoolPtrInput
	// Free-form maintenance notes for the runner (1024 characters).
	MaintenanceNote pulumi.StringPtrInput
	// Maximum timeout set when this runner handles the job.
	MaximumTimeout pulumi.IntPtrInput
	// Whether the runner should ignore new jobs.
	Paused pulumi.BoolPtrInput
	// The registration token used to register the runner.
	RegistrationToken pulumi.StringPtrInput
	// Whether the runner should handle untagged jobs.
	RunUntagged pulumi.BoolPtrInput
	// The status of runners to show, one of: online and offline. active and paused are also possible values
	// 			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
	Status pulumi.StringPtrInput
	// List of runner’s tags.
	TagLists pulumi.StringArrayInput
}

func (RunnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*runnerState)(nil)).Elem()
}

type runnerArgs struct {
	// The accessLevel of the runner. Valid values are: `notProtected`, `refProtected`.
	AccessLevel *string `pulumi:"accessLevel"`
	// The runner's description.
	Description *string `pulumi:"description"`
	// Whether the runner should be locked for current project.
	Locked *bool `pulumi:"locked"`
	// Free-form maintenance notes for the runner (1024 characters).
	MaintenanceNote *string `pulumi:"maintenanceNote"`
	// Maximum timeout set when this runner handles the job.
	MaximumTimeout *int `pulumi:"maximumTimeout"`
	// Whether the runner should ignore new jobs.
	Paused *bool `pulumi:"paused"`
	// The registration token used to register the runner.
	RegistrationToken string `pulumi:"registrationToken"`
	// Whether the runner should handle untagged jobs.
	RunUntagged *bool `pulumi:"runUntagged"`
	// List of runner’s tags.
	TagLists []string `pulumi:"tagLists"`
}

// The set of arguments for constructing a Runner resource.
type RunnerArgs struct {
	// The accessLevel of the runner. Valid values are: `notProtected`, `refProtected`.
	AccessLevel pulumi.StringPtrInput
	// The runner's description.
	Description pulumi.StringPtrInput
	// Whether the runner should be locked for current project.
	Locked pulumi.BoolPtrInput
	// Free-form maintenance notes for the runner (1024 characters).
	MaintenanceNote pulumi.StringPtrInput
	// Maximum timeout set when this runner handles the job.
	MaximumTimeout pulumi.IntPtrInput
	// Whether the runner should ignore new jobs.
	Paused pulumi.BoolPtrInput
	// The registration token used to register the runner.
	RegistrationToken pulumi.StringInput
	// Whether the runner should handle untagged jobs.
	RunUntagged pulumi.BoolPtrInput
	// List of runner’s tags.
	TagLists pulumi.StringArrayInput
}

func (RunnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runnerArgs)(nil)).Elem()
}

type RunnerInput interface {
	pulumi.Input

	ToRunnerOutput() RunnerOutput
	ToRunnerOutputWithContext(ctx context.Context) RunnerOutput
}

func (*Runner) ElementType() reflect.Type {
	return reflect.TypeOf((**Runner)(nil)).Elem()
}

func (i *Runner) ToRunnerOutput() RunnerOutput {
	return i.ToRunnerOutputWithContext(context.Background())
}

func (i *Runner) ToRunnerOutputWithContext(ctx context.Context) RunnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunnerOutput)
}

// RunnerArrayInput is an input type that accepts RunnerArray and RunnerArrayOutput values.
// You can construct a concrete instance of `RunnerArrayInput` via:
//
//	RunnerArray{ RunnerArgs{...} }
type RunnerArrayInput interface {
	pulumi.Input

	ToRunnerArrayOutput() RunnerArrayOutput
	ToRunnerArrayOutputWithContext(context.Context) RunnerArrayOutput
}

type RunnerArray []RunnerInput

func (RunnerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Runner)(nil)).Elem()
}

func (i RunnerArray) ToRunnerArrayOutput() RunnerArrayOutput {
	return i.ToRunnerArrayOutputWithContext(context.Background())
}

func (i RunnerArray) ToRunnerArrayOutputWithContext(ctx context.Context) RunnerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunnerArrayOutput)
}

// RunnerMapInput is an input type that accepts RunnerMap and RunnerMapOutput values.
// You can construct a concrete instance of `RunnerMapInput` via:
//
//	RunnerMap{ "key": RunnerArgs{...} }
type RunnerMapInput interface {
	pulumi.Input

	ToRunnerMapOutput() RunnerMapOutput
	ToRunnerMapOutputWithContext(context.Context) RunnerMapOutput
}

type RunnerMap map[string]RunnerInput

func (RunnerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Runner)(nil)).Elem()
}

func (i RunnerMap) ToRunnerMapOutput() RunnerMapOutput {
	return i.ToRunnerMapOutputWithContext(context.Background())
}

func (i RunnerMap) ToRunnerMapOutputWithContext(ctx context.Context) RunnerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunnerMapOutput)
}

type RunnerOutput struct{ *pulumi.OutputState }

func (RunnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Runner)(nil)).Elem()
}

func (o RunnerOutput) ToRunnerOutput() RunnerOutput {
	return o
}

func (o RunnerOutput) ToRunnerOutputWithContext(ctx context.Context) RunnerOutput {
	return o
}

// The accessLevel of the runner. Valid values are: `notProtected`, `refProtected`.
func (o RunnerOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Runner) pulumi.StringOutput { return v.AccessLevel }).(pulumi.StringOutput)
}

// The authentication token used for building a config.toml file. This value is not present when imported.
func (o RunnerOutput) AuthenticationToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Runner) pulumi.StringOutput { return v.AuthenticationToken }).(pulumi.StringOutput)
}

// The runner's description.
func (o RunnerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runner) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether the runner should be locked for current project.
func (o RunnerOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v *Runner) pulumi.BoolOutput { return v.Locked }).(pulumi.BoolOutput)
}

// Free-form maintenance notes for the runner (1024 characters).
func (o RunnerOutput) MaintenanceNote() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runner) pulumi.StringPtrOutput { return v.MaintenanceNote }).(pulumi.StringPtrOutput)
}

// Maximum timeout set when this runner handles the job.
func (o RunnerOutput) MaximumTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Runner) pulumi.IntPtrOutput { return v.MaximumTimeout }).(pulumi.IntPtrOutput)
}

// Whether the runner should ignore new jobs.
func (o RunnerOutput) Paused() pulumi.BoolOutput {
	return o.ApplyT(func(v *Runner) pulumi.BoolOutput { return v.Paused }).(pulumi.BoolOutput)
}

// The registration token used to register the runner.
func (o RunnerOutput) RegistrationToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Runner) pulumi.StringOutput { return v.RegistrationToken }).(pulumi.StringOutput)
}

// Whether the runner should handle untagged jobs.
func (o RunnerOutput) RunUntagged() pulumi.BoolOutput {
	return o.ApplyT(func(v *Runner) pulumi.BoolOutput { return v.RunUntagged }).(pulumi.BoolOutput)
}

// The status of runners to show, one of: online and offline. active and paused are also possible values
//
//	which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
func (o RunnerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Runner) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// List of runner’s tags.
func (o RunnerOutput) TagLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Runner) pulumi.StringArrayOutput { return v.TagLists }).(pulumi.StringArrayOutput)
}

type RunnerArrayOutput struct{ *pulumi.OutputState }

func (RunnerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Runner)(nil)).Elem()
}

func (o RunnerArrayOutput) ToRunnerArrayOutput() RunnerArrayOutput {
	return o
}

func (o RunnerArrayOutput) ToRunnerArrayOutputWithContext(ctx context.Context) RunnerArrayOutput {
	return o
}

func (o RunnerArrayOutput) Index(i pulumi.IntInput) RunnerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Runner {
		return vs[0].([]*Runner)[vs[1].(int)]
	}).(RunnerOutput)
}

type RunnerMapOutput struct{ *pulumi.OutputState }

func (RunnerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Runner)(nil)).Elem()
}

func (o RunnerMapOutput) ToRunnerMapOutput() RunnerMapOutput {
	return o
}

func (o RunnerMapOutput) ToRunnerMapOutputWithContext(ctx context.Context) RunnerMapOutput {
	return o
}

func (o RunnerMapOutput) MapIndex(k pulumi.StringInput) RunnerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Runner {
		return vs[0].(map[string]*Runner)[vs[1].(string)]
	}).(RunnerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RunnerInput)(nil)).Elem(), &Runner{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunnerArrayInput)(nil)).Elem(), RunnerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunnerMapInput)(nil)).Elem(), RunnerMap{})
	pulumi.RegisterOutputType(RunnerOutput{})
	pulumi.RegisterOutputType(RunnerArrayOutput{})
	pulumi.RegisterOutputType(RunnerMapOutput{})
}
