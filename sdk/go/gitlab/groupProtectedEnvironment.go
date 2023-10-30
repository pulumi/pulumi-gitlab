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

// The `GroupProtectedEnvironment` resource allows to manage the lifecycle of a protected environment in a group.
//
// > In order to use a userId in the `deployAccessLevels` configuration,
//
//	you need to make sure that users have access to the group with Maintainer role or higher.
//	In order to use a groupId in the `deployAccessLevels` configuration,
//	the groupId must be a sub-group under the given group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_protected_environments.html)
//
// ## Import
//
// GitLab group protected environments can be imported using an id made up of `groupId:environmentName`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment bar 123:production
//
// ```
type GroupProtectedEnvironment struct {
	pulumi.CustomResourceState

	// Array of approval rules to deploy, with each described by a hash.
	ApprovalRules GroupProtectedEnvironmentApprovalRuleArrayOutput `pulumi:"approvalRules"`
	// Array of access levels allowed to deploy, with each described by a hash.
	DeployAccessLevels GroupProtectedEnvironmentDeployAccessLevelArrayOutput `pulumi:"deployAccessLevels"`
	// The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The ID or full path of the group which the protected environment is created against.
	Group pulumi.StringOutput `pulumi:"group"`
	// The number of approvals required to deploy to this environment.
	RequiredApprovalCount pulumi.IntOutput `pulumi:"requiredApprovalCount"`
}

// NewGroupProtectedEnvironment registers a new resource with the given unique name, arguments, and options.
func NewGroupProtectedEnvironment(ctx *pulumi.Context,
	name string, args *GroupProtectedEnvironmentArgs, opts ...pulumi.ResourceOption) (*GroupProtectedEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeployAccessLevels == nil {
		return nil, errors.New("invalid value for required argument 'DeployAccessLevels'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupProtectedEnvironment
	err := ctx.RegisterResource("gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupProtectedEnvironment gets an existing GroupProtectedEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupProtectedEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupProtectedEnvironmentState, opts ...pulumi.ResourceOption) (*GroupProtectedEnvironment, error) {
	var resource GroupProtectedEnvironment
	err := ctx.ReadResource("gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupProtectedEnvironment resources.
type groupProtectedEnvironmentState struct {
	// Array of approval rules to deploy, with each described by a hash.
	ApprovalRules []GroupProtectedEnvironmentApprovalRule `pulumi:"approvalRules"`
	// Array of access levels allowed to deploy, with each described by a hash.
	DeployAccessLevels []GroupProtectedEnvironmentDeployAccessLevel `pulumi:"deployAccessLevels"`
	// The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
	Environment *string `pulumi:"environment"`
	// The ID or full path of the group which the protected environment is created against.
	Group *string `pulumi:"group"`
	// The number of approvals required to deploy to this environment.
	RequiredApprovalCount *int `pulumi:"requiredApprovalCount"`
}

type GroupProtectedEnvironmentState struct {
	// Array of approval rules to deploy, with each described by a hash.
	ApprovalRules GroupProtectedEnvironmentApprovalRuleArrayInput
	// Array of access levels allowed to deploy, with each described by a hash.
	DeployAccessLevels GroupProtectedEnvironmentDeployAccessLevelArrayInput
	// The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
	Environment pulumi.StringPtrInput
	// The ID or full path of the group which the protected environment is created against.
	Group pulumi.StringPtrInput
	// The number of approvals required to deploy to this environment.
	RequiredApprovalCount pulumi.IntPtrInput
}

func (GroupProtectedEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupProtectedEnvironmentState)(nil)).Elem()
}

type groupProtectedEnvironmentArgs struct {
	// Array of approval rules to deploy, with each described by a hash.
	ApprovalRules []GroupProtectedEnvironmentApprovalRule `pulumi:"approvalRules"`
	// Array of access levels allowed to deploy, with each described by a hash.
	DeployAccessLevels []GroupProtectedEnvironmentDeployAccessLevel `pulumi:"deployAccessLevels"`
	// The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
	Environment string `pulumi:"environment"`
	// The ID or full path of the group which the protected environment is created against.
	Group string `pulumi:"group"`
	// The number of approvals required to deploy to this environment.
	RequiredApprovalCount *int `pulumi:"requiredApprovalCount"`
}

// The set of arguments for constructing a GroupProtectedEnvironment resource.
type GroupProtectedEnvironmentArgs struct {
	// Array of approval rules to deploy, with each described by a hash.
	ApprovalRules GroupProtectedEnvironmentApprovalRuleArrayInput
	// Array of access levels allowed to deploy, with each described by a hash.
	DeployAccessLevels GroupProtectedEnvironmentDeployAccessLevelArrayInput
	// The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
	Environment pulumi.StringInput
	// The ID or full path of the group which the protected environment is created against.
	Group pulumi.StringInput
	// The number of approvals required to deploy to this environment.
	RequiredApprovalCount pulumi.IntPtrInput
}

func (GroupProtectedEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupProtectedEnvironmentArgs)(nil)).Elem()
}

type GroupProtectedEnvironmentInput interface {
	pulumi.Input

	ToGroupProtectedEnvironmentOutput() GroupProtectedEnvironmentOutput
	ToGroupProtectedEnvironmentOutputWithContext(ctx context.Context) GroupProtectedEnvironmentOutput
}

func (*GroupProtectedEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupProtectedEnvironment)(nil)).Elem()
}

func (i *GroupProtectedEnvironment) ToGroupProtectedEnvironmentOutput() GroupProtectedEnvironmentOutput {
	return i.ToGroupProtectedEnvironmentOutputWithContext(context.Background())
}

func (i *GroupProtectedEnvironment) ToGroupProtectedEnvironmentOutputWithContext(ctx context.Context) GroupProtectedEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupProtectedEnvironmentOutput)
}

func (i *GroupProtectedEnvironment) ToOutput(ctx context.Context) pulumix.Output[*GroupProtectedEnvironment] {
	return pulumix.Output[*GroupProtectedEnvironment]{
		OutputState: i.ToGroupProtectedEnvironmentOutputWithContext(ctx).OutputState,
	}
}

// GroupProtectedEnvironmentArrayInput is an input type that accepts GroupProtectedEnvironmentArray and GroupProtectedEnvironmentArrayOutput values.
// You can construct a concrete instance of `GroupProtectedEnvironmentArrayInput` via:
//
//	GroupProtectedEnvironmentArray{ GroupProtectedEnvironmentArgs{...} }
type GroupProtectedEnvironmentArrayInput interface {
	pulumi.Input

	ToGroupProtectedEnvironmentArrayOutput() GroupProtectedEnvironmentArrayOutput
	ToGroupProtectedEnvironmentArrayOutputWithContext(context.Context) GroupProtectedEnvironmentArrayOutput
}

type GroupProtectedEnvironmentArray []GroupProtectedEnvironmentInput

func (GroupProtectedEnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupProtectedEnvironment)(nil)).Elem()
}

func (i GroupProtectedEnvironmentArray) ToGroupProtectedEnvironmentArrayOutput() GroupProtectedEnvironmentArrayOutput {
	return i.ToGroupProtectedEnvironmentArrayOutputWithContext(context.Background())
}

func (i GroupProtectedEnvironmentArray) ToGroupProtectedEnvironmentArrayOutputWithContext(ctx context.Context) GroupProtectedEnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupProtectedEnvironmentArrayOutput)
}

func (i GroupProtectedEnvironmentArray) ToOutput(ctx context.Context) pulumix.Output[[]*GroupProtectedEnvironment] {
	return pulumix.Output[[]*GroupProtectedEnvironment]{
		OutputState: i.ToGroupProtectedEnvironmentArrayOutputWithContext(ctx).OutputState,
	}
}

// GroupProtectedEnvironmentMapInput is an input type that accepts GroupProtectedEnvironmentMap and GroupProtectedEnvironmentMapOutput values.
// You can construct a concrete instance of `GroupProtectedEnvironmentMapInput` via:
//
//	GroupProtectedEnvironmentMap{ "key": GroupProtectedEnvironmentArgs{...} }
type GroupProtectedEnvironmentMapInput interface {
	pulumi.Input

	ToGroupProtectedEnvironmentMapOutput() GroupProtectedEnvironmentMapOutput
	ToGroupProtectedEnvironmentMapOutputWithContext(context.Context) GroupProtectedEnvironmentMapOutput
}

type GroupProtectedEnvironmentMap map[string]GroupProtectedEnvironmentInput

func (GroupProtectedEnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupProtectedEnvironment)(nil)).Elem()
}

func (i GroupProtectedEnvironmentMap) ToGroupProtectedEnvironmentMapOutput() GroupProtectedEnvironmentMapOutput {
	return i.ToGroupProtectedEnvironmentMapOutputWithContext(context.Background())
}

func (i GroupProtectedEnvironmentMap) ToGroupProtectedEnvironmentMapOutputWithContext(ctx context.Context) GroupProtectedEnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupProtectedEnvironmentMapOutput)
}

func (i GroupProtectedEnvironmentMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*GroupProtectedEnvironment] {
	return pulumix.Output[map[string]*GroupProtectedEnvironment]{
		OutputState: i.ToGroupProtectedEnvironmentMapOutputWithContext(ctx).OutputState,
	}
}

type GroupProtectedEnvironmentOutput struct{ *pulumi.OutputState }

func (GroupProtectedEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupProtectedEnvironment)(nil)).Elem()
}

func (o GroupProtectedEnvironmentOutput) ToGroupProtectedEnvironmentOutput() GroupProtectedEnvironmentOutput {
	return o
}

func (o GroupProtectedEnvironmentOutput) ToGroupProtectedEnvironmentOutputWithContext(ctx context.Context) GroupProtectedEnvironmentOutput {
	return o
}

func (o GroupProtectedEnvironmentOutput) ToOutput(ctx context.Context) pulumix.Output[*GroupProtectedEnvironment] {
	return pulumix.Output[*GroupProtectedEnvironment]{
		OutputState: o.OutputState,
	}
}

// Array of approval rules to deploy, with each described by a hash.
func (o GroupProtectedEnvironmentOutput) ApprovalRules() GroupProtectedEnvironmentApprovalRuleArrayOutput {
	return o.ApplyT(func(v *GroupProtectedEnvironment) GroupProtectedEnvironmentApprovalRuleArrayOutput {
		return v.ApprovalRules
	}).(GroupProtectedEnvironmentApprovalRuleArrayOutput)
}

// Array of access levels allowed to deploy, with each described by a hash.
func (o GroupProtectedEnvironmentOutput) DeployAccessLevels() GroupProtectedEnvironmentDeployAccessLevelArrayOutput {
	return o.ApplyT(func(v *GroupProtectedEnvironment) GroupProtectedEnvironmentDeployAccessLevelArrayOutput {
		return v.DeployAccessLevels
	}).(GroupProtectedEnvironmentDeployAccessLevelArrayOutput)
}

// The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
func (o GroupProtectedEnvironmentOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupProtectedEnvironment) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The ID or full path of the group which the protected environment is created against.
func (o GroupProtectedEnvironmentOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupProtectedEnvironment) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The number of approvals required to deploy to this environment.
func (o GroupProtectedEnvironmentOutput) RequiredApprovalCount() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupProtectedEnvironment) pulumi.IntOutput { return v.RequiredApprovalCount }).(pulumi.IntOutput)
}

type GroupProtectedEnvironmentArrayOutput struct{ *pulumi.OutputState }

func (GroupProtectedEnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupProtectedEnvironment)(nil)).Elem()
}

func (o GroupProtectedEnvironmentArrayOutput) ToGroupProtectedEnvironmentArrayOutput() GroupProtectedEnvironmentArrayOutput {
	return o
}

func (o GroupProtectedEnvironmentArrayOutput) ToGroupProtectedEnvironmentArrayOutputWithContext(ctx context.Context) GroupProtectedEnvironmentArrayOutput {
	return o
}

func (o GroupProtectedEnvironmentArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*GroupProtectedEnvironment] {
	return pulumix.Output[[]*GroupProtectedEnvironment]{
		OutputState: o.OutputState,
	}
}

func (o GroupProtectedEnvironmentArrayOutput) Index(i pulumi.IntInput) GroupProtectedEnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupProtectedEnvironment {
		return vs[0].([]*GroupProtectedEnvironment)[vs[1].(int)]
	}).(GroupProtectedEnvironmentOutput)
}

type GroupProtectedEnvironmentMapOutput struct{ *pulumi.OutputState }

func (GroupProtectedEnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupProtectedEnvironment)(nil)).Elem()
}

func (o GroupProtectedEnvironmentMapOutput) ToGroupProtectedEnvironmentMapOutput() GroupProtectedEnvironmentMapOutput {
	return o
}

func (o GroupProtectedEnvironmentMapOutput) ToGroupProtectedEnvironmentMapOutputWithContext(ctx context.Context) GroupProtectedEnvironmentMapOutput {
	return o
}

func (o GroupProtectedEnvironmentMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*GroupProtectedEnvironment] {
	return pulumix.Output[map[string]*GroupProtectedEnvironment]{
		OutputState: o.OutputState,
	}
}

func (o GroupProtectedEnvironmentMapOutput) MapIndex(k pulumi.StringInput) GroupProtectedEnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupProtectedEnvironment {
		return vs[0].(map[string]*GroupProtectedEnvironment)[vs[1].(string)]
	}).(GroupProtectedEnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupProtectedEnvironmentInput)(nil)).Elem(), &GroupProtectedEnvironment{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupProtectedEnvironmentArrayInput)(nil)).Elem(), GroupProtectedEnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupProtectedEnvironmentMapInput)(nil)).Elem(), GroupProtectedEnvironmentMap{})
	pulumi.RegisterOutputType(GroupProtectedEnvironmentOutput{})
	pulumi.RegisterOutputType(GroupProtectedEnvironmentArrayOutput{})
	pulumi.RegisterOutputType(GroupProtectedEnvironmentMapOutput{})
}
