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
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleGroup, err := gitlab.NewGroup(ctx, "exampleGroup", &gitlab.GroupArgs{
//				Path:        pulumi.String("example"),
//				Description: pulumi.String("An example group"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewProject(ctx, "exampleProject", &gitlab.ProjectArgs{
//				Description: pulumi.String("An example project"),
//				NamespaceId: exampleGroup.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewGroup(ctx, "example-two", &gitlab.GroupArgs{
//				Path:        pulumi.String("example-two"),
//				Description: pulumi.String("An example group with push rules"),
//				PushRules: &gitlab.GroupPushRulesArgs{
//					AuthorEmailRegex:     pulumi.String("@example\\.com$"),
//					CommitCommitterCheck: pulumi.Bool(true),
//					MemberCheck:          pulumi.Bool(true),
//					PreventSecrets:       pulumi.Bool(true),
//				},
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
// ```sh
//
//	$ pulumi import gitlab:index/group:Group You can import a group state using `<resource> <id>`. The
//
// ```
//
//	`id` can be whatever the [details of a group][details_of_a_group] api takes for
//
//	its `:id` value, so for example:
//
// ```sh
// $ pulumi import gitlab:index/group:Group example example
// ```
type Group struct {
	pulumi.CustomResourceState

	// Default to Auto DevOps pipeline for all projects within this group.
	AutoDevopsEnabled pulumi.BoolOutput `pulumi:"autoDevopsEnabled"`
	// A local path to the avatar image to upload. **Note**: not available for imported resources.
	Avatar pulumi.StringPtrOutput `pulumi:"avatar"`
	// The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
	AvatarHash pulumi.StringOutput `pulumi:"avatarHash"`
	// The URL of the avatar image.
	AvatarUrl pulumi.StringOutput `pulumi:"avatarUrl"`
	// See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
	DefaultBranchProtection pulumi.IntOutput `pulumi:"defaultBranchProtection"`
	// The group's description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Disable email notifications.
	EmailsDisabled pulumi.BoolOutput `pulumi:"emailsDisabled"`
	// Can be set by administrators only. Additional CI/CD minutes for this group.
	ExtraSharedRunnersMinutesLimit pulumi.IntOutput `pulumi:"extraSharedRunnersMinutesLimit"`
	// The full name of the group.
	FullName pulumi.StringOutput `pulumi:"fullName"`
	// The full path of the group.
	FullPath pulumi.StringOutput `pulumi:"fullPath"`
	// A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
	IpRestrictionRanges pulumi.StringArrayOutput `pulumi:"ipRestrictionRanges"`
	// Enable/disable Large File Storage (LFS) for the projects in this group.
	LfsEnabled pulumi.BoolOutput `pulumi:"lfsEnabled"`
	// Users cannot be added to projects in this group.
	MembershipLock pulumi.BoolPtrOutput `pulumi:"membershipLock"`
	// Disable the capability of a group from getting mentioned.
	MentionsDisabled pulumi.BoolOutput `pulumi:"mentionsDisabled"`
	// The name of the group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Id of the parent group (creates a nested group).
	ParentId pulumi.IntOutput `pulumi:"parentId"`
	// The path of the group.
	Path pulumi.StringOutput `pulumi:"path"`
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup pulumi.BoolOutput `pulumi:"preventForkingOutsideGroup"`
	// Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
	ProjectCreationLevel pulumi.StringOutput `pulumi:"projectCreationLevel"`
	// Push rules for the group.
	PushRules GroupPushRulesOutput `pulumi:"pushRules"`
	// Allow users to request member access.
	RequestAccessEnabled pulumi.BoolOutput `pulumi:"requestAccessEnabled"`
	// Require all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication pulumi.BoolOutput `pulumi:"requireTwoFactorAuthentication"`
	// The group level registration token to use during runner setup.
	RunnersToken pulumi.StringOutput `pulumi:"runnersToken"`
	// Prevent sharing a project with another group within this group.
	ShareWithGroupLock pulumi.BoolOutput `pulumi:"shareWithGroupLock"`
	// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
	SharedRunnersMinutesLimit pulumi.IntOutput `pulumi:"sharedRunnersMinutesLimit"`
	// Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabledAndOverridable`, `disabledAndUnoverridable`, `disabledWithOverride`.
	SharedRunnersSetting pulumi.StringOutput `pulumi:"sharedRunnersSetting"`
	// Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
	SubgroupCreationLevel pulumi.StringOutput `pulumi:"subgroupCreationLevel"`
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod pulumi.IntOutput `pulumi:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
	VisibilityLevel pulumi.StringOutput `pulumi:"visibilityLevel"`
	// Web URL of the group.
	WebUrl pulumi.StringOutput `pulumi:"webUrl"`
	// The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
	WikiAccessLevel pulumi.StringOutput `pulumi:"wikiAccessLevel"`
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
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"runnersToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Default to Auto DevOps pipeline for all projects within this group.
	AutoDevopsEnabled *bool `pulumi:"autoDevopsEnabled"`
	// A local path to the avatar image to upload. **Note**: not available for imported resources.
	Avatar *string `pulumi:"avatar"`
	// The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
	AvatarHash *string `pulumi:"avatarHash"`
	// The URL of the avatar image.
	AvatarUrl *string `pulumi:"avatarUrl"`
	// See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
	DefaultBranchProtection *int `pulumi:"defaultBranchProtection"`
	// The group's description.
	Description *string `pulumi:"description"`
	// Disable email notifications.
	EmailsDisabled *bool `pulumi:"emailsDisabled"`
	// Can be set by administrators only. Additional CI/CD minutes for this group.
	ExtraSharedRunnersMinutesLimit *int `pulumi:"extraSharedRunnersMinutesLimit"`
	// The full name of the group.
	FullName *string `pulumi:"fullName"`
	// The full path of the group.
	FullPath *string `pulumi:"fullPath"`
	// A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
	IpRestrictionRanges []string `pulumi:"ipRestrictionRanges"`
	// Enable/disable Large File Storage (LFS) for the projects in this group.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// Users cannot be added to projects in this group.
	MembershipLock *bool `pulumi:"membershipLock"`
	// Disable the capability of a group from getting mentioned.
	MentionsDisabled *bool `pulumi:"mentionsDisabled"`
	// The name of the group.
	Name *string `pulumi:"name"`
	// Id of the parent group (creates a nested group).
	ParentId *int `pulumi:"parentId"`
	// The path of the group.
	Path *string `pulumi:"path"`
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup *bool `pulumi:"preventForkingOutsideGroup"`
	// Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
	ProjectCreationLevel *string `pulumi:"projectCreationLevel"`
	// Push rules for the group.
	PushRules *GroupPushRules `pulumi:"pushRules"`
	// Allow users to request member access.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Require all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication *bool `pulumi:"requireTwoFactorAuthentication"`
	// The group level registration token to use during runner setup.
	RunnersToken *string `pulumi:"runnersToken"`
	// Prevent sharing a project with another group within this group.
	ShareWithGroupLock *bool `pulumi:"shareWithGroupLock"`
	// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
	SharedRunnersMinutesLimit *int `pulumi:"sharedRunnersMinutesLimit"`
	// Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabledAndOverridable`, `disabledAndUnoverridable`, `disabledWithOverride`.
	SharedRunnersSetting *string `pulumi:"sharedRunnersSetting"`
	// Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
	SubgroupCreationLevel *string `pulumi:"subgroupCreationLevel"`
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod *int `pulumi:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
	// Web URL of the group.
	WebUrl *string `pulumi:"webUrl"`
	// The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
	WikiAccessLevel *string `pulumi:"wikiAccessLevel"`
}

type GroupState struct {
	// Default to Auto DevOps pipeline for all projects within this group.
	AutoDevopsEnabled pulumi.BoolPtrInput
	// A local path to the avatar image to upload. **Note**: not available for imported resources.
	Avatar pulumi.StringPtrInput
	// The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
	AvatarHash pulumi.StringPtrInput
	// The URL of the avatar image.
	AvatarUrl pulumi.StringPtrInput
	// See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
	DefaultBranchProtection pulumi.IntPtrInput
	// The group's description.
	Description pulumi.StringPtrInput
	// Disable email notifications.
	EmailsDisabled pulumi.BoolPtrInput
	// Can be set by administrators only. Additional CI/CD minutes for this group.
	ExtraSharedRunnersMinutesLimit pulumi.IntPtrInput
	// The full name of the group.
	FullName pulumi.StringPtrInput
	// The full path of the group.
	FullPath pulumi.StringPtrInput
	// A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
	IpRestrictionRanges pulumi.StringArrayInput
	// Enable/disable Large File Storage (LFS) for the projects in this group.
	LfsEnabled pulumi.BoolPtrInput
	// Users cannot be added to projects in this group.
	MembershipLock pulumi.BoolPtrInput
	// Disable the capability of a group from getting mentioned.
	MentionsDisabled pulumi.BoolPtrInput
	// The name of the group.
	Name pulumi.StringPtrInput
	// Id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrInput
	// The path of the group.
	Path pulumi.StringPtrInput
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup pulumi.BoolPtrInput
	// Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
	ProjectCreationLevel pulumi.StringPtrInput
	// Push rules for the group.
	PushRules GroupPushRulesPtrInput
	// Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Require all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication pulumi.BoolPtrInput
	// The group level registration token to use during runner setup.
	RunnersToken pulumi.StringPtrInput
	// Prevent sharing a project with another group within this group.
	ShareWithGroupLock pulumi.BoolPtrInput
	// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
	SharedRunnersMinutesLimit pulumi.IntPtrInput
	// Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabledAndOverridable`, `disabledAndUnoverridable`, `disabledWithOverride`.
	SharedRunnersSetting pulumi.StringPtrInput
	// Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
	SubgroupCreationLevel pulumi.StringPtrInput
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod pulumi.IntPtrInput
	// The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
	VisibilityLevel pulumi.StringPtrInput
	// Web URL of the group.
	WebUrl pulumi.StringPtrInput
	// The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
	WikiAccessLevel pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Default to Auto DevOps pipeline for all projects within this group.
	AutoDevopsEnabled *bool `pulumi:"autoDevopsEnabled"`
	// A local path to the avatar image to upload. **Note**: not available for imported resources.
	Avatar *string `pulumi:"avatar"`
	// The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
	AvatarHash *string `pulumi:"avatarHash"`
	// See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
	DefaultBranchProtection *int `pulumi:"defaultBranchProtection"`
	// The group's description.
	Description *string `pulumi:"description"`
	// Disable email notifications.
	EmailsDisabled *bool `pulumi:"emailsDisabled"`
	// Can be set by administrators only. Additional CI/CD minutes for this group.
	ExtraSharedRunnersMinutesLimit *int `pulumi:"extraSharedRunnersMinutesLimit"`
	// A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
	IpRestrictionRanges []string `pulumi:"ipRestrictionRanges"`
	// Enable/disable Large File Storage (LFS) for the projects in this group.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// Users cannot be added to projects in this group.
	MembershipLock *bool `pulumi:"membershipLock"`
	// Disable the capability of a group from getting mentioned.
	MentionsDisabled *bool `pulumi:"mentionsDisabled"`
	// The name of the group.
	Name *string `pulumi:"name"`
	// Id of the parent group (creates a nested group).
	ParentId *int `pulumi:"parentId"`
	// The path of the group.
	Path string `pulumi:"path"`
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup *bool `pulumi:"preventForkingOutsideGroup"`
	// Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
	ProjectCreationLevel *string `pulumi:"projectCreationLevel"`
	// Push rules for the group.
	PushRules *GroupPushRules `pulumi:"pushRules"`
	// Allow users to request member access.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Require all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication *bool `pulumi:"requireTwoFactorAuthentication"`
	// Prevent sharing a project with another group within this group.
	ShareWithGroupLock *bool `pulumi:"shareWithGroupLock"`
	// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
	SharedRunnersMinutesLimit *int `pulumi:"sharedRunnersMinutesLimit"`
	// Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabledAndOverridable`, `disabledAndUnoverridable`, `disabledWithOverride`.
	SharedRunnersSetting *string `pulumi:"sharedRunnersSetting"`
	// Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
	SubgroupCreationLevel *string `pulumi:"subgroupCreationLevel"`
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod *int `pulumi:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
	// The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
	WikiAccessLevel *string `pulumi:"wikiAccessLevel"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Default to Auto DevOps pipeline for all projects within this group.
	AutoDevopsEnabled pulumi.BoolPtrInput
	// A local path to the avatar image to upload. **Note**: not available for imported resources.
	Avatar pulumi.StringPtrInput
	// The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
	AvatarHash pulumi.StringPtrInput
	// See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
	DefaultBranchProtection pulumi.IntPtrInput
	// The group's description.
	Description pulumi.StringPtrInput
	// Disable email notifications.
	EmailsDisabled pulumi.BoolPtrInput
	// Can be set by administrators only. Additional CI/CD minutes for this group.
	ExtraSharedRunnersMinutesLimit pulumi.IntPtrInput
	// A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
	IpRestrictionRanges pulumi.StringArrayInput
	// Enable/disable Large File Storage (LFS) for the projects in this group.
	LfsEnabled pulumi.BoolPtrInput
	// Users cannot be added to projects in this group.
	MembershipLock pulumi.BoolPtrInput
	// Disable the capability of a group from getting mentioned.
	MentionsDisabled pulumi.BoolPtrInput
	// The name of the group.
	Name pulumi.StringPtrInput
	// Id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrInput
	// The path of the group.
	Path pulumi.StringInput
	// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup pulumi.BoolPtrInput
	// Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
	ProjectCreationLevel pulumi.StringPtrInput
	// Push rules for the group.
	PushRules GroupPushRulesPtrInput
	// Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Require all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication pulumi.BoolPtrInput
	// Prevent sharing a project with another group within this group.
	ShareWithGroupLock pulumi.BoolPtrInput
	// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
	SharedRunnersMinutesLimit pulumi.IntPtrInput
	// Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabledAndOverridable`, `disabledAndUnoverridable`, `disabledWithOverride`.
	SharedRunnersSetting pulumi.StringPtrInput
	// Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
	SubgroupCreationLevel pulumi.StringPtrInput
	// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod pulumi.IntPtrInput
	// The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
	VisibilityLevel pulumi.StringPtrInput
	// The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
	WikiAccessLevel pulumi.StringPtrInput
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
//	GroupArray{ GroupArgs{...} }
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
//	GroupMap{ "key": GroupArgs{...} }
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

// Default to Auto DevOps pipeline for all projects within this group.
func (o GroupOutput) AutoDevopsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolOutput { return v.AutoDevopsEnabled }).(pulumi.BoolOutput)
}

// A local path to the avatar image to upload. **Note**: not available for imported resources.
func (o GroupOutput) Avatar() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Avatar }).(pulumi.StringPtrOutput)
}

// The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
func (o GroupOutput) AvatarHash() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.AvatarHash }).(pulumi.StringOutput)
}

// The URL of the avatar image.
func (o GroupOutput) AvatarUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.AvatarUrl }).(pulumi.StringOutput)
}

// See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
func (o GroupOutput) DefaultBranchProtection() pulumi.IntOutput {
	return o.ApplyT(func(v *Group) pulumi.IntOutput { return v.DefaultBranchProtection }).(pulumi.IntOutput)
}

// The group's description.
func (o GroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Disable email notifications.
func (o GroupOutput) EmailsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolOutput { return v.EmailsDisabled }).(pulumi.BoolOutput)
}

// Can be set by administrators only. Additional CI/CD minutes for this group.
func (o GroupOutput) ExtraSharedRunnersMinutesLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Group) pulumi.IntOutput { return v.ExtraSharedRunnersMinutesLimit }).(pulumi.IntOutput)
}

// The full name of the group.
func (o GroupOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.FullName }).(pulumi.StringOutput)
}

// The full path of the group.
func (o GroupOutput) FullPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.FullPath }).(pulumi.StringOutput)
}

// A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
func (o GroupOutput) IpRestrictionRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.StringArrayOutput { return v.IpRestrictionRanges }).(pulumi.StringArrayOutput)
}

// Enable/disable Large File Storage (LFS) for the projects in this group.
func (o GroupOutput) LfsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolOutput { return v.LfsEnabled }).(pulumi.BoolOutput)
}

// Users cannot be added to projects in this group.
func (o GroupOutput) MembershipLock() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolPtrOutput { return v.MembershipLock }).(pulumi.BoolPtrOutput)
}

// Disable the capability of a group from getting mentioned.
func (o GroupOutput) MentionsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolOutput { return v.MentionsDisabled }).(pulumi.BoolOutput)
}

// The name of the group.
func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Id of the parent group (creates a nested group).
func (o GroupOutput) ParentId() pulumi.IntOutput {
	return o.ApplyT(func(v *Group) pulumi.IntOutput { return v.ParentId }).(pulumi.IntOutput)
}

// The path of the group.
func (o GroupOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
func (o GroupOutput) PreventForkingOutsideGroup() pulumi.BoolOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolOutput { return v.PreventForkingOutsideGroup }).(pulumi.BoolOutput)
}

// Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
func (o GroupOutput) ProjectCreationLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.ProjectCreationLevel }).(pulumi.StringOutput)
}

// Push rules for the group.
func (o GroupOutput) PushRules() GroupPushRulesOutput {
	return o.ApplyT(func(v *Group) GroupPushRulesOutput { return v.PushRules }).(GroupPushRulesOutput)
}

// Allow users to request member access.
func (o GroupOutput) RequestAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolOutput { return v.RequestAccessEnabled }).(pulumi.BoolOutput)
}

// Require all users in this group to setup Two-factor authentication.
func (o GroupOutput) RequireTwoFactorAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolOutput { return v.RequireTwoFactorAuthentication }).(pulumi.BoolOutput)
}

// The group level registration token to use during runner setup.
func (o GroupOutput) RunnersToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.RunnersToken }).(pulumi.StringOutput)
}

// Prevent sharing a project with another group within this group.
func (o GroupOutput) ShareWithGroupLock() pulumi.BoolOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolOutput { return v.ShareWithGroupLock }).(pulumi.BoolOutput)
}

// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
func (o GroupOutput) SharedRunnersMinutesLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Group) pulumi.IntOutput { return v.SharedRunnersMinutesLimit }).(pulumi.IntOutput)
}

// Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabledAndOverridable`, `disabledAndUnoverridable`, `disabledWithOverride`.
func (o GroupOutput) SharedRunnersSetting() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.SharedRunnersSetting }).(pulumi.StringOutput)
}

// Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
func (o GroupOutput) SubgroupCreationLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.SubgroupCreationLevel }).(pulumi.StringOutput)
}

// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
func (o GroupOutput) TwoFactorGracePeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Group) pulumi.IntOutput { return v.TwoFactorGracePeriod }).(pulumi.IntOutput)
}

// The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
func (o GroupOutput) VisibilityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.VisibilityLevel }).(pulumi.StringOutput)
}

// Web URL of the group.
func (o GroupOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.WebUrl }).(pulumi.StringOutput)
}

// The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
func (o GroupOutput) WikiAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.WikiAccessLevel }).(pulumi.StringOutput)
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
