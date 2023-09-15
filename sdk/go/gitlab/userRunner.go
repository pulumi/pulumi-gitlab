// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The `UserRunner` resource allows creating a GitLab runner using the new [GitLab Runner Registration Flow](https://docs.gitlab.com/ee/ci/runners/new_creation_workflow.html).
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#create-a-runner)
type UserRunner struct {
	pulumi.CustomResourceState

	// The access level of the runner. Valid values are: `notProtected`, `refProtected`.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// Description of the runner.
	Description pulumi.StringOutput `pulumi:"description"`
	// The ID of the group that the runner is created in. Required if runner*type is group*type.
	GroupId pulumi.IntPtrOutput `pulumi:"groupId"`
	// Specifies if the runner should be locked for the current project.
	Locked pulumi.BoolOutput `pulumi:"locked"`
	// Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
	MaximumTimeout pulumi.IntOutput `pulumi:"maximumTimeout"`
	// Specifies if the runner should ignore new jobs.
	Paused pulumi.BoolOutput `pulumi:"paused"`
	// The ID of the project that the runner is created in. Required if runner*type is project*type.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// The scope of the runner. Valid values are: `instanceType`, `groupType`, `projectType`.
	RunnerType pulumi.StringOutput `pulumi:"runnerType"`
	// A list of runner tags.
	TagLists pulumi.StringArrayOutput `pulumi:"tagLists"`
	// The authentication token to use when setting up a new runner with this configuration. This value cannot be imported.
	Token pulumi.StringOutput `pulumi:"token"`
	// Specifies if the runner should handle untagged jobs.
	Untagged pulumi.BoolOutput `pulumi:"untagged"`
}

// NewUserRunner registers a new resource with the given unique name, arguments, and options.
func NewUserRunner(ctx *pulumi.Context,
	name string, args *UserRunnerArgs, opts ...pulumi.ResourceOption) (*UserRunner, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RunnerType == nil {
		return nil, errors.New("invalid value for required argument 'RunnerType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserRunner
	err := ctx.RegisterResource("gitlab:index/userRunner:UserRunner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserRunner gets an existing UserRunner resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserRunner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRunnerState, opts ...pulumi.ResourceOption) (*UserRunner, error) {
	var resource UserRunner
	err := ctx.ReadResource("gitlab:index/userRunner:UserRunner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserRunner resources.
type userRunnerState struct {
	// The access level of the runner. Valid values are: `notProtected`, `refProtected`.
	AccessLevel *string `pulumi:"accessLevel"`
	// Description of the runner.
	Description *string `pulumi:"description"`
	// The ID of the group that the runner is created in. Required if runner*type is group*type.
	GroupId *int `pulumi:"groupId"`
	// Specifies if the runner should be locked for the current project.
	Locked *bool `pulumi:"locked"`
	// Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
	MaximumTimeout *int `pulumi:"maximumTimeout"`
	// Specifies if the runner should ignore new jobs.
	Paused *bool `pulumi:"paused"`
	// The ID of the project that the runner is created in. Required if runner*type is project*type.
	ProjectId *int `pulumi:"projectId"`
	// The scope of the runner. Valid values are: `instanceType`, `groupType`, `projectType`.
	RunnerType *string `pulumi:"runnerType"`
	// A list of runner tags.
	TagLists []string `pulumi:"tagLists"`
	// The authentication token to use when setting up a new runner with this configuration. This value cannot be imported.
	Token *string `pulumi:"token"`
	// Specifies if the runner should handle untagged jobs.
	Untagged *bool `pulumi:"untagged"`
}

type UserRunnerState struct {
	// The access level of the runner. Valid values are: `notProtected`, `refProtected`.
	AccessLevel pulumi.StringPtrInput
	// Description of the runner.
	Description pulumi.StringPtrInput
	// The ID of the group that the runner is created in. Required if runner*type is group*type.
	GroupId pulumi.IntPtrInput
	// Specifies if the runner should be locked for the current project.
	Locked pulumi.BoolPtrInput
	// Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
	MaximumTimeout pulumi.IntPtrInput
	// Specifies if the runner should ignore new jobs.
	Paused pulumi.BoolPtrInput
	// The ID of the project that the runner is created in. Required if runner*type is project*type.
	ProjectId pulumi.IntPtrInput
	// The scope of the runner. Valid values are: `instanceType`, `groupType`, `projectType`.
	RunnerType pulumi.StringPtrInput
	// A list of runner tags.
	TagLists pulumi.StringArrayInput
	// The authentication token to use when setting up a new runner with this configuration. This value cannot be imported.
	Token pulumi.StringPtrInput
	// Specifies if the runner should handle untagged jobs.
	Untagged pulumi.BoolPtrInput
}

func (UserRunnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRunnerState)(nil)).Elem()
}

type userRunnerArgs struct {
	// The access level of the runner. Valid values are: `notProtected`, `refProtected`.
	AccessLevel *string `pulumi:"accessLevel"`
	// Description of the runner.
	Description *string `pulumi:"description"`
	// The ID of the group that the runner is created in. Required if runner*type is group*type.
	GroupId *int `pulumi:"groupId"`
	// Specifies if the runner should be locked for the current project.
	Locked *bool `pulumi:"locked"`
	// Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
	MaximumTimeout *int `pulumi:"maximumTimeout"`
	// Specifies if the runner should ignore new jobs.
	Paused *bool `pulumi:"paused"`
	// The ID of the project that the runner is created in. Required if runner*type is project*type.
	ProjectId *int `pulumi:"projectId"`
	// The scope of the runner. Valid values are: `instanceType`, `groupType`, `projectType`.
	RunnerType string `pulumi:"runnerType"`
	// A list of runner tags.
	TagLists []string `pulumi:"tagLists"`
	// Specifies if the runner should handle untagged jobs.
	Untagged *bool `pulumi:"untagged"`
}

// The set of arguments for constructing a UserRunner resource.
type UserRunnerArgs struct {
	// The access level of the runner. Valid values are: `notProtected`, `refProtected`.
	AccessLevel pulumi.StringPtrInput
	// Description of the runner.
	Description pulumi.StringPtrInput
	// The ID of the group that the runner is created in. Required if runner*type is group*type.
	GroupId pulumi.IntPtrInput
	// Specifies if the runner should be locked for the current project.
	Locked pulumi.BoolPtrInput
	// Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
	MaximumTimeout pulumi.IntPtrInput
	// Specifies if the runner should ignore new jobs.
	Paused pulumi.BoolPtrInput
	// The ID of the project that the runner is created in. Required if runner*type is project*type.
	ProjectId pulumi.IntPtrInput
	// The scope of the runner. Valid values are: `instanceType`, `groupType`, `projectType`.
	RunnerType pulumi.StringInput
	// A list of runner tags.
	TagLists pulumi.StringArrayInput
	// Specifies if the runner should handle untagged jobs.
	Untagged pulumi.BoolPtrInput
}

func (UserRunnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRunnerArgs)(nil)).Elem()
}

type UserRunnerInput interface {
	pulumi.Input

	ToUserRunnerOutput() UserRunnerOutput
	ToUserRunnerOutputWithContext(ctx context.Context) UserRunnerOutput
}

func (*UserRunner) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRunner)(nil)).Elem()
}

func (i *UserRunner) ToUserRunnerOutput() UserRunnerOutput {
	return i.ToUserRunnerOutputWithContext(context.Background())
}

func (i *UserRunner) ToUserRunnerOutputWithContext(ctx context.Context) UserRunnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRunnerOutput)
}

func (i *UserRunner) ToOutput(ctx context.Context) pulumix.Output[*UserRunner] {
	return pulumix.Output[*UserRunner]{
		OutputState: i.ToUserRunnerOutputWithContext(ctx).OutputState,
	}
}

// UserRunnerArrayInput is an input type that accepts UserRunnerArray and UserRunnerArrayOutput values.
// You can construct a concrete instance of `UserRunnerArrayInput` via:
//
//	UserRunnerArray{ UserRunnerArgs{...} }
type UserRunnerArrayInput interface {
	pulumi.Input

	ToUserRunnerArrayOutput() UserRunnerArrayOutput
	ToUserRunnerArrayOutputWithContext(context.Context) UserRunnerArrayOutput
}

type UserRunnerArray []UserRunnerInput

func (UserRunnerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserRunner)(nil)).Elem()
}

func (i UserRunnerArray) ToUserRunnerArrayOutput() UserRunnerArrayOutput {
	return i.ToUserRunnerArrayOutputWithContext(context.Background())
}

func (i UserRunnerArray) ToUserRunnerArrayOutputWithContext(ctx context.Context) UserRunnerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRunnerArrayOutput)
}

func (i UserRunnerArray) ToOutput(ctx context.Context) pulumix.Output[[]*UserRunner] {
	return pulumix.Output[[]*UserRunner]{
		OutputState: i.ToUserRunnerArrayOutputWithContext(ctx).OutputState,
	}
}

// UserRunnerMapInput is an input type that accepts UserRunnerMap and UserRunnerMapOutput values.
// You can construct a concrete instance of `UserRunnerMapInput` via:
//
//	UserRunnerMap{ "key": UserRunnerArgs{...} }
type UserRunnerMapInput interface {
	pulumi.Input

	ToUserRunnerMapOutput() UserRunnerMapOutput
	ToUserRunnerMapOutputWithContext(context.Context) UserRunnerMapOutput
}

type UserRunnerMap map[string]UserRunnerInput

func (UserRunnerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserRunner)(nil)).Elem()
}

func (i UserRunnerMap) ToUserRunnerMapOutput() UserRunnerMapOutput {
	return i.ToUserRunnerMapOutputWithContext(context.Background())
}

func (i UserRunnerMap) ToUserRunnerMapOutputWithContext(ctx context.Context) UserRunnerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRunnerMapOutput)
}

func (i UserRunnerMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserRunner] {
	return pulumix.Output[map[string]*UserRunner]{
		OutputState: i.ToUserRunnerMapOutputWithContext(ctx).OutputState,
	}
}

type UserRunnerOutput struct{ *pulumi.OutputState }

func (UserRunnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRunner)(nil)).Elem()
}

func (o UserRunnerOutput) ToUserRunnerOutput() UserRunnerOutput {
	return o
}

func (o UserRunnerOutput) ToUserRunnerOutputWithContext(ctx context.Context) UserRunnerOutput {
	return o
}

func (o UserRunnerOutput) ToOutput(ctx context.Context) pulumix.Output[*UserRunner] {
	return pulumix.Output[*UserRunner]{
		OutputState: o.OutputState,
	}
}

// The access level of the runner. Valid values are: `notProtected`, `refProtected`.
func (o UserRunnerOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.StringOutput { return v.AccessLevel }).(pulumi.StringOutput)
}

// Description of the runner.
func (o UserRunnerOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The ID of the group that the runner is created in. Required if runner*type is group*type.
func (o UserRunnerOutput) GroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.IntPtrOutput { return v.GroupId }).(pulumi.IntPtrOutput)
}

// Specifies if the runner should be locked for the current project.
func (o UserRunnerOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.BoolOutput { return v.Locked }).(pulumi.BoolOutput)
}

// Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
func (o UserRunnerOutput) MaximumTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.IntOutput { return v.MaximumTimeout }).(pulumi.IntOutput)
}

// Specifies if the runner should ignore new jobs.
func (o UserRunnerOutput) Paused() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.BoolOutput { return v.Paused }).(pulumi.BoolOutput)
}

// The ID of the project that the runner is created in. Required if runner*type is project*type.
func (o UserRunnerOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// The scope of the runner. Valid values are: `instanceType`, `groupType`, `projectType`.
func (o UserRunnerOutput) RunnerType() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.StringOutput { return v.RunnerType }).(pulumi.StringOutput)
}

// A list of runner tags.
func (o UserRunnerOutput) TagLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.StringArrayOutput { return v.TagLists }).(pulumi.StringArrayOutput)
}

// The authentication token to use when setting up a new runner with this configuration. This value cannot be imported.
func (o UserRunnerOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// Specifies if the runner should handle untagged jobs.
func (o UserRunnerOutput) Untagged() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserRunner) pulumi.BoolOutput { return v.Untagged }).(pulumi.BoolOutput)
}

type UserRunnerArrayOutput struct{ *pulumi.OutputState }

func (UserRunnerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserRunner)(nil)).Elem()
}

func (o UserRunnerArrayOutput) ToUserRunnerArrayOutput() UserRunnerArrayOutput {
	return o
}

func (o UserRunnerArrayOutput) ToUserRunnerArrayOutputWithContext(ctx context.Context) UserRunnerArrayOutput {
	return o
}

func (o UserRunnerArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*UserRunner] {
	return pulumix.Output[[]*UserRunner]{
		OutputState: o.OutputState,
	}
}

func (o UserRunnerArrayOutput) Index(i pulumi.IntInput) UserRunnerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserRunner {
		return vs[0].([]*UserRunner)[vs[1].(int)]
	}).(UserRunnerOutput)
}

type UserRunnerMapOutput struct{ *pulumi.OutputState }

func (UserRunnerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserRunner)(nil)).Elem()
}

func (o UserRunnerMapOutput) ToUserRunnerMapOutput() UserRunnerMapOutput {
	return o
}

func (o UserRunnerMapOutput) ToUserRunnerMapOutputWithContext(ctx context.Context) UserRunnerMapOutput {
	return o
}

func (o UserRunnerMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserRunner] {
	return pulumix.Output[map[string]*UserRunner]{
		OutputState: o.OutputState,
	}
}

func (o UserRunnerMapOutput) MapIndex(k pulumi.StringInput) UserRunnerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserRunner {
		return vs[0].(map[string]*UserRunner)[vs[1].(string)]
	}).(UserRunnerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserRunnerInput)(nil)).Elem(), &UserRunner{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRunnerArrayInput)(nil)).Elem(), UserRunnerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRunnerMapInput)(nil)).Elem(), UserRunnerMap{})
	pulumi.RegisterOutputType(UserRunnerOutput{})
	pulumi.RegisterOutputType(UserRunnerArrayOutput{})
	pulumi.RegisterOutputType(UserRunnerMapOutput{})
}
