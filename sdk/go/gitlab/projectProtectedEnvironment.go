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

// The `ProjectProtectedEnvironment` resource allows to manage the lifecycle of a protected environment in a project.
//
// > In order to use a user or group in the `deployAccessLevels` configuration,
//
//	you need to make sure that users have access to the project and groups must have this project shared.
//	You may use the `ProjectMembership` and `gitlabProjectSharedGroup` resources to achieve this.
//	Unfortunately, the GitLab API does not complain about users and groups without access to the project and just ignores those.
//	In case this happens you will get perpetual state diffs.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_environments.html)
//
// ## Import
//
// GitLab protected environments can be imported using an id made up of `projectId:environmentName`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment bar 123:production
// ```
type ProjectProtectedEnvironment struct {
	pulumi.CustomResourceState

	// Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
	ApprovalRules ProjectProtectedEnvironmentApprovalRuleArrayOutput `pulumi:"approvalRules"`
	// Array of access levels allowed to deploy, with each described by a hash.  Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
	DeployAccessLevels ProjectProtectedEnvironmentDeployAccessLevelArrayOutput `pulumi:"deployAccessLevels"`
	// The name of the environment.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The ID or full path of the project which the protected environment is created against.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectProtectedEnvironment registers a new resource with the given unique name, arguments, and options.
func NewProjectProtectedEnvironment(ctx *pulumi.Context,
	name string, args *ProjectProtectedEnvironmentArgs, opts ...pulumi.ResourceOption) (*ProjectProtectedEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectProtectedEnvironment
	err := ctx.RegisterResource("gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectProtectedEnvironment gets an existing ProjectProtectedEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectProtectedEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectProtectedEnvironmentState, opts ...pulumi.ResourceOption) (*ProjectProtectedEnvironment, error) {
	var resource ProjectProtectedEnvironment
	err := ctx.ReadResource("gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectProtectedEnvironment resources.
type projectProtectedEnvironmentState struct {
	// Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
	ApprovalRules []ProjectProtectedEnvironmentApprovalRule `pulumi:"approvalRules"`
	// Array of access levels allowed to deploy, with each described by a hash.  Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
	DeployAccessLevels []ProjectProtectedEnvironmentDeployAccessLevel `pulumi:"deployAccessLevels"`
	// The name of the environment.
	Environment *string `pulumi:"environment"`
	// The ID or full path of the project which the protected environment is created against.
	Project *string `pulumi:"project"`
}

type ProjectProtectedEnvironmentState struct {
	// Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
	ApprovalRules ProjectProtectedEnvironmentApprovalRuleArrayInput
	// Array of access levels allowed to deploy, with each described by a hash.  Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
	DeployAccessLevels ProjectProtectedEnvironmentDeployAccessLevelArrayInput
	// The name of the environment.
	Environment pulumi.StringPtrInput
	// The ID or full path of the project which the protected environment is created against.
	Project pulumi.StringPtrInput
}

func (ProjectProtectedEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectProtectedEnvironmentState)(nil)).Elem()
}

type projectProtectedEnvironmentArgs struct {
	// Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
	ApprovalRules []ProjectProtectedEnvironmentApprovalRule `pulumi:"approvalRules"`
	// Array of access levels allowed to deploy, with each described by a hash.  Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
	DeployAccessLevels []ProjectProtectedEnvironmentDeployAccessLevel `pulumi:"deployAccessLevels"`
	// The name of the environment.
	Environment string `pulumi:"environment"`
	// The ID or full path of the project which the protected environment is created against.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectProtectedEnvironment resource.
type ProjectProtectedEnvironmentArgs struct {
	// Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
	ApprovalRules ProjectProtectedEnvironmentApprovalRuleArrayInput
	// Array of access levels allowed to deploy, with each described by a hash.  Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
	DeployAccessLevels ProjectProtectedEnvironmentDeployAccessLevelArrayInput
	// The name of the environment.
	Environment pulumi.StringInput
	// The ID or full path of the project which the protected environment is created against.
	Project pulumi.StringInput
}

func (ProjectProtectedEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectProtectedEnvironmentArgs)(nil)).Elem()
}

type ProjectProtectedEnvironmentInput interface {
	pulumi.Input

	ToProjectProtectedEnvironmentOutput() ProjectProtectedEnvironmentOutput
	ToProjectProtectedEnvironmentOutputWithContext(ctx context.Context) ProjectProtectedEnvironmentOutput
}

func (*ProjectProtectedEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectProtectedEnvironment)(nil)).Elem()
}

func (i *ProjectProtectedEnvironment) ToProjectProtectedEnvironmentOutput() ProjectProtectedEnvironmentOutput {
	return i.ToProjectProtectedEnvironmentOutputWithContext(context.Background())
}

func (i *ProjectProtectedEnvironment) ToProjectProtectedEnvironmentOutputWithContext(ctx context.Context) ProjectProtectedEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectProtectedEnvironmentOutput)
}

// ProjectProtectedEnvironmentArrayInput is an input type that accepts ProjectProtectedEnvironmentArray and ProjectProtectedEnvironmentArrayOutput values.
// You can construct a concrete instance of `ProjectProtectedEnvironmentArrayInput` via:
//
//	ProjectProtectedEnvironmentArray{ ProjectProtectedEnvironmentArgs{...} }
type ProjectProtectedEnvironmentArrayInput interface {
	pulumi.Input

	ToProjectProtectedEnvironmentArrayOutput() ProjectProtectedEnvironmentArrayOutput
	ToProjectProtectedEnvironmentArrayOutputWithContext(context.Context) ProjectProtectedEnvironmentArrayOutput
}

type ProjectProtectedEnvironmentArray []ProjectProtectedEnvironmentInput

func (ProjectProtectedEnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectProtectedEnvironment)(nil)).Elem()
}

func (i ProjectProtectedEnvironmentArray) ToProjectProtectedEnvironmentArrayOutput() ProjectProtectedEnvironmentArrayOutput {
	return i.ToProjectProtectedEnvironmentArrayOutputWithContext(context.Background())
}

func (i ProjectProtectedEnvironmentArray) ToProjectProtectedEnvironmentArrayOutputWithContext(ctx context.Context) ProjectProtectedEnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectProtectedEnvironmentArrayOutput)
}

// ProjectProtectedEnvironmentMapInput is an input type that accepts ProjectProtectedEnvironmentMap and ProjectProtectedEnvironmentMapOutput values.
// You can construct a concrete instance of `ProjectProtectedEnvironmentMapInput` via:
//
//	ProjectProtectedEnvironmentMap{ "key": ProjectProtectedEnvironmentArgs{...} }
type ProjectProtectedEnvironmentMapInput interface {
	pulumi.Input

	ToProjectProtectedEnvironmentMapOutput() ProjectProtectedEnvironmentMapOutput
	ToProjectProtectedEnvironmentMapOutputWithContext(context.Context) ProjectProtectedEnvironmentMapOutput
}

type ProjectProtectedEnvironmentMap map[string]ProjectProtectedEnvironmentInput

func (ProjectProtectedEnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectProtectedEnvironment)(nil)).Elem()
}

func (i ProjectProtectedEnvironmentMap) ToProjectProtectedEnvironmentMapOutput() ProjectProtectedEnvironmentMapOutput {
	return i.ToProjectProtectedEnvironmentMapOutputWithContext(context.Background())
}

func (i ProjectProtectedEnvironmentMap) ToProjectProtectedEnvironmentMapOutputWithContext(ctx context.Context) ProjectProtectedEnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectProtectedEnvironmentMapOutput)
}

type ProjectProtectedEnvironmentOutput struct{ *pulumi.OutputState }

func (ProjectProtectedEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectProtectedEnvironment)(nil)).Elem()
}

func (o ProjectProtectedEnvironmentOutput) ToProjectProtectedEnvironmentOutput() ProjectProtectedEnvironmentOutput {
	return o
}

func (o ProjectProtectedEnvironmentOutput) ToProjectProtectedEnvironmentOutputWithContext(ctx context.Context) ProjectProtectedEnvironmentOutput {
	return o
}

// Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
func (o ProjectProtectedEnvironmentOutput) ApprovalRules() ProjectProtectedEnvironmentApprovalRuleArrayOutput {
	return o.ApplyT(func(v *ProjectProtectedEnvironment) ProjectProtectedEnvironmentApprovalRuleArrayOutput {
		return v.ApprovalRules
	}).(ProjectProtectedEnvironmentApprovalRuleArrayOutput)
}

// Array of access levels allowed to deploy, with each described by a hash.  Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
func (o ProjectProtectedEnvironmentOutput) DeployAccessLevels() ProjectProtectedEnvironmentDeployAccessLevelArrayOutput {
	return o.ApplyT(func(v *ProjectProtectedEnvironment) ProjectProtectedEnvironmentDeployAccessLevelArrayOutput {
		return v.DeployAccessLevels
	}).(ProjectProtectedEnvironmentDeployAccessLevelArrayOutput)
}

// The name of the environment.
func (o ProjectProtectedEnvironmentOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectProtectedEnvironment) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The ID or full path of the project which the protected environment is created against.
func (o ProjectProtectedEnvironmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectProtectedEnvironment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ProjectProtectedEnvironmentArrayOutput struct{ *pulumi.OutputState }

func (ProjectProtectedEnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectProtectedEnvironment)(nil)).Elem()
}

func (o ProjectProtectedEnvironmentArrayOutput) ToProjectProtectedEnvironmentArrayOutput() ProjectProtectedEnvironmentArrayOutput {
	return o
}

func (o ProjectProtectedEnvironmentArrayOutput) ToProjectProtectedEnvironmentArrayOutputWithContext(ctx context.Context) ProjectProtectedEnvironmentArrayOutput {
	return o
}

func (o ProjectProtectedEnvironmentArrayOutput) Index(i pulumi.IntInput) ProjectProtectedEnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectProtectedEnvironment {
		return vs[0].([]*ProjectProtectedEnvironment)[vs[1].(int)]
	}).(ProjectProtectedEnvironmentOutput)
}

type ProjectProtectedEnvironmentMapOutput struct{ *pulumi.OutputState }

func (ProjectProtectedEnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectProtectedEnvironment)(nil)).Elem()
}

func (o ProjectProtectedEnvironmentMapOutput) ToProjectProtectedEnvironmentMapOutput() ProjectProtectedEnvironmentMapOutput {
	return o
}

func (o ProjectProtectedEnvironmentMapOutput) ToProjectProtectedEnvironmentMapOutputWithContext(ctx context.Context) ProjectProtectedEnvironmentMapOutput {
	return o
}

func (o ProjectProtectedEnvironmentMapOutput) MapIndex(k pulumi.StringInput) ProjectProtectedEnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectProtectedEnvironment {
		return vs[0].(map[string]*ProjectProtectedEnvironment)[vs[1].(string)]
	}).(ProjectProtectedEnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectProtectedEnvironmentInput)(nil)).Elem(), &ProjectProtectedEnvironment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectProtectedEnvironmentArrayInput)(nil)).Elem(), ProjectProtectedEnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectProtectedEnvironmentMapInput)(nil)).Elem(), ProjectProtectedEnvironmentMap{})
	pulumi.RegisterOutputType(ProjectProtectedEnvironmentOutput{})
	pulumi.RegisterOutputType(ProjectProtectedEnvironmentArrayOutput{})
	pulumi.RegisterOutputType(ProjectProtectedEnvironmentMapOutput{})
}
