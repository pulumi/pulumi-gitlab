// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Group` resource allows to manage the lifecycle of a group.
//
// > On GitLab SaaS, you must use the GitLab UI to create groups without a parent group. You cannot use this provider nor the API to do this.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html)
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
// 		exampleGroup, err := gitlab.NewGroup(ctx, "exampleGroup", &gitlab.GroupArgs{
// 			Path:        pulumi.String("example"),
// 			Description: pulumi.String("An example group"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewProject(ctx, "exampleProject", &gitlab.ProjectArgs{
// 			Description: pulumi.String("An example project"),
// 			NamespaceId: exampleGroup.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ```sh
//  $ pulumi import gitlab:index/group:Group # You can import a group state using `<resource> <id>`. The
// ```
//
// # `id` can be whatever the [details of a group][details_of_a_group] api takes for # its `:id` value, so for example
//
// ```sh
//  $ pulumi import gitlab:index/group:Group example example
// ```
type Group struct {
	pulumi.CustomResourceState

	// Defaults to false. Default to Auto DevOps pipeline for all projects within this group.
	AutoDevopsEnabled pulumi.BoolPtrOutput `pulumi:"autoDevopsEnabled"`
	// Defaults to 2. See https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection
	DefaultBranchProtection pulumi.IntPtrOutput `pulumi:"defaultBranchProtection"`
	// The description of the group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Defaults to false. Disable email notifications.
	EmailsDisabled pulumi.BoolPtrOutput `pulumi:"emailsDisabled"`
	// The full name of the group.
	FullName pulumi.StringOutput `pulumi:"fullName"`
	// The full path of the group.
	FullPath pulumi.StringOutput `pulumi:"fullPath"`
	// Defaults to true. Enable/disable Large File Storage (LFS) for the projects in this group.
	LfsEnabled pulumi.BoolPtrOutput `pulumi:"lfsEnabled"`
	// Defaults to false. Disable the capability of a group from getting mentioned.
	MentionsDisabled pulumi.BoolPtrOutput `pulumi:"mentionsDisabled"`
	// The name of this group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrOutput `pulumi:"parentId"`
	// The path of the group.
	Path pulumi.StringOutput `pulumi:"path"`
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup pulumi.BoolPtrOutput `pulumi:"preventForkingOutsideGroup"`
	// Defaults to maintainer. Determine if developers can create projects in the group.
	ProjectCreationLevel pulumi.StringPtrOutput `pulumi:"projectCreationLevel"`
	// Defaults to false. Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrOutput `pulumi:"requestAccessEnabled"`
	// Defaults to false. Require all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication pulumi.BoolPtrOutput `pulumi:"requireTwoFactorAuthentication"`
	// The group level registration token to use during runner setup.
	RunnersToken pulumi.StringOutput `pulumi:"runnersToken"`
	// Defaults to false. Prevent sharing a project with another group within this group.
	ShareWithGroupLock pulumi.BoolPtrOutput `pulumi:"shareWithGroupLock"`
	// Defaults to owner. Allowed to create subgroups.
	SubgroupCreationLevel pulumi.StringPtrOutput `pulumi:"subgroupCreationLevel"`
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod pulumi.IntPtrOutput `pulumi:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`.
	VisibilityLevel pulumi.StringOutput `pulumi:"visibilityLevel"`
	// Web URL of the group.
	WebUrl pulumi.StringOutput `pulumi:"webUrl"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	var resource Group
	err := ctx.RegisterResource("gitlab:index/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("gitlab:index/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// Defaults to false. Default to Auto DevOps pipeline for all projects within this group.
	AutoDevopsEnabled *bool `pulumi:"autoDevopsEnabled"`
	// Defaults to 2. See https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection
	DefaultBranchProtection *int `pulumi:"defaultBranchProtection"`
	// The description of the group.
	Description *string `pulumi:"description"`
	// Defaults to false. Disable email notifications.
	EmailsDisabled *bool `pulumi:"emailsDisabled"`
	// The full name of the group.
	FullName *string `pulumi:"fullName"`
	// The full path of the group.
	FullPath *string `pulumi:"fullPath"`
	// Defaults to true. Enable/disable Large File Storage (LFS) for the projects in this group.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// Defaults to false. Disable the capability of a group from getting mentioned.
	MentionsDisabled *bool `pulumi:"mentionsDisabled"`
	// The name of this group.
	Name *string `pulumi:"name"`
	// Id of the parent group (creates a nested group).
	ParentId *int `pulumi:"parentId"`
	// The path of the group.
	Path *string `pulumi:"path"`
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup *bool `pulumi:"preventForkingOutsideGroup"`
	// Defaults to maintainer. Determine if developers can create projects in the group.
	ProjectCreationLevel *string `pulumi:"projectCreationLevel"`
	// Defaults to false. Allow users to request member access.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Defaults to false. Require all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication *bool `pulumi:"requireTwoFactorAuthentication"`
	// The group level registration token to use during runner setup.
	RunnersToken *string `pulumi:"runnersToken"`
	// Defaults to false. Prevent sharing a project with another group within this group.
	ShareWithGroupLock *bool `pulumi:"shareWithGroupLock"`
	// Defaults to owner. Allowed to create subgroups.
	SubgroupCreationLevel *string `pulumi:"subgroupCreationLevel"`
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod *int `pulumi:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
	// Web URL of the group.
	WebUrl *string `pulumi:"webUrl"`
}

type GroupState struct {
	// Defaults to false. Default to Auto DevOps pipeline for all projects within this group.
	AutoDevopsEnabled pulumi.BoolPtrInput
	// Defaults to 2. See https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection
	DefaultBranchProtection pulumi.IntPtrInput
	// The description of the group.
	Description pulumi.StringPtrInput
	// Defaults to false. Disable email notifications.
	EmailsDisabled pulumi.BoolPtrInput
	// The full name of the group.
	FullName pulumi.StringPtrInput
	// The full path of the group.
	FullPath pulumi.StringPtrInput
	// Defaults to true. Enable/disable Large File Storage (LFS) for the projects in this group.
	LfsEnabled pulumi.BoolPtrInput
	// Defaults to false. Disable the capability of a group from getting mentioned.
	MentionsDisabled pulumi.BoolPtrInput
	// The name of this group.
	Name pulumi.StringPtrInput
	// Id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrInput
	// The path of the group.
	Path pulumi.StringPtrInput
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup pulumi.BoolPtrInput
	// Defaults to maintainer. Determine if developers can create projects in the group.
	ProjectCreationLevel pulumi.StringPtrInput
	// Defaults to false. Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Defaults to false. Require all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication pulumi.BoolPtrInput
	// The group level registration token to use during runner setup.
	RunnersToken pulumi.StringPtrInput
	// Defaults to false. Prevent sharing a project with another group within this group.
	ShareWithGroupLock pulumi.BoolPtrInput
	// Defaults to owner. Allowed to create subgroups.
	SubgroupCreationLevel pulumi.StringPtrInput
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod pulumi.IntPtrInput
	// The group's visibility. Can be `private`, `internal`, or `public`.
	VisibilityLevel pulumi.StringPtrInput
	// Web URL of the group.
	WebUrl pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Defaults to false. Default to Auto DevOps pipeline for all projects within this group.
	AutoDevopsEnabled *bool `pulumi:"autoDevopsEnabled"`
	// Defaults to 2. See https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection
	DefaultBranchProtection *int `pulumi:"defaultBranchProtection"`
	// The description of the group.
	Description *string `pulumi:"description"`
	// Defaults to false. Disable email notifications.
	EmailsDisabled *bool `pulumi:"emailsDisabled"`
	// Defaults to true. Enable/disable Large File Storage (LFS) for the projects in this group.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// Defaults to false. Disable the capability of a group from getting mentioned.
	MentionsDisabled *bool `pulumi:"mentionsDisabled"`
	// The name of this group.
	Name *string `pulumi:"name"`
	// Id of the parent group (creates a nested group).
	ParentId *int `pulumi:"parentId"`
	// The path of the group.
	Path string `pulumi:"path"`
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup *bool `pulumi:"preventForkingOutsideGroup"`
	// Defaults to maintainer. Determine if developers can create projects in the group.
	ProjectCreationLevel *string `pulumi:"projectCreationLevel"`
	// Defaults to false. Allow users to request member access.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Defaults to false. Require all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication *bool `pulumi:"requireTwoFactorAuthentication"`
	// Defaults to false. Prevent sharing a project with another group within this group.
	ShareWithGroupLock *bool `pulumi:"shareWithGroupLock"`
	// Defaults to owner. Allowed to create subgroups.
	SubgroupCreationLevel *string `pulumi:"subgroupCreationLevel"`
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod *int `pulumi:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Defaults to false. Default to Auto DevOps pipeline for all projects within this group.
	AutoDevopsEnabled pulumi.BoolPtrInput
	// Defaults to 2. See https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection
	DefaultBranchProtection pulumi.IntPtrInput
	// The description of the group.
	Description pulumi.StringPtrInput
	// Defaults to false. Disable email notifications.
	EmailsDisabled pulumi.BoolPtrInput
	// Defaults to true. Enable/disable Large File Storage (LFS) for the projects in this group.
	LfsEnabled pulumi.BoolPtrInput
	// Defaults to false. Disable the capability of a group from getting mentioned.
	MentionsDisabled pulumi.BoolPtrInput
	// The name of this group.
	Name pulumi.StringPtrInput
	// Id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrInput
	// The path of the group.
	Path pulumi.StringInput
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup pulumi.BoolPtrInput
	// Defaults to maintainer. Determine if developers can create projects in the group.
	ProjectCreationLevel pulumi.StringPtrInput
	// Defaults to false. Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Defaults to false. Require all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication pulumi.BoolPtrInput
	// Defaults to false. Prevent sharing a project with another group within this group.
	ShareWithGroupLock pulumi.BoolPtrInput
	// Defaults to owner. Allowed to create subgroups.
	SubgroupCreationLevel pulumi.StringPtrInput
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod pulumi.IntPtrInput
	// The group's visibility. Can be `private`, `internal`, or `public`.
	VisibilityLevel pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//          GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//          GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
