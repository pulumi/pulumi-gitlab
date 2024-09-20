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

// The `MemberRole` resource allows to manage the lifecycle of a custom member role.
//
// Custom roles allow an organization to create user roles with the precise privileges and permissions required for that organization’s needs.
//
// > This resource requires an Ultimate license.
//
// > Most custom roles are considered billable users that use a seat. [Custom roles billing and seat usage](https://docs.gitlab.com/ee/user/custom_roles.html#billing-and-seat-usage)
//
// > There can be only 10 custom roles on your instance or namespace. See [issue 450929](https://gitlab.com/gitlab-org/gitlab/-/issues/450929) for more details.
//
// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#mutationmemberrolecreate)
type MemberRole struct {
	pulumi.CustomResourceState

	// The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
	BaseAccessLevel pulumi.StringOutput `pulumi:"baseAccessLevel"`
	// Timestamp of when the member role was created. Only available with GitLab version 17.3 or higher.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description for the member role.
	Description pulumi.StringOutput `pulumi:"description"`
	// The Web UI path to edit the member role
	EditPath pulumi.StringOutput `pulumi:"editPath"`
	// All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
	EnabledPermissions pulumi.StringArrayOutput `pulumi:"enabledPermissions"`
	// Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
	GroupPath pulumi.StringOutput `pulumi:"groupPath"`
	// The id integer value extracted from the `id` attribute
	Iid pulumi.IntOutput `pulumi:"iid"`
	// Name for the member role.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewMemberRole registers a new resource with the given unique name, arguments, and options.
func NewMemberRole(ctx *pulumi.Context,
	name string, args *MemberRoleArgs, opts ...pulumi.ResourceOption) (*MemberRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseAccessLevel == nil {
		return nil, errors.New("invalid value for required argument 'BaseAccessLevel'")
	}
	if args.EnabledPermissions == nil {
		return nil, errors.New("invalid value for required argument 'EnabledPermissions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MemberRole
	err := ctx.RegisterResource("gitlab:index/memberRole:MemberRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMemberRole gets an existing MemberRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMemberRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MemberRoleState, opts ...pulumi.ResourceOption) (*MemberRole, error) {
	var resource MemberRole
	err := ctx.ReadResource("gitlab:index/memberRole:MemberRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MemberRole resources.
type memberRoleState struct {
	// The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
	BaseAccessLevel *string `pulumi:"baseAccessLevel"`
	// Timestamp of when the member role was created. Only available with GitLab version 17.3 or higher.
	CreatedAt *string `pulumi:"createdAt"`
	// Description for the member role.
	Description *string `pulumi:"description"`
	// The Web UI path to edit the member role
	EditPath *string `pulumi:"editPath"`
	// All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
	EnabledPermissions []string `pulumi:"enabledPermissions"`
	// Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
	GroupPath *string `pulumi:"groupPath"`
	// The id integer value extracted from the `id` attribute
	Iid *int `pulumi:"iid"`
	// Name for the member role.
	Name *string `pulumi:"name"`
}

type MemberRoleState struct {
	// The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
	BaseAccessLevel pulumi.StringPtrInput
	// Timestamp of when the member role was created. Only available with GitLab version 17.3 or higher.
	CreatedAt pulumi.StringPtrInput
	// Description for the member role.
	Description pulumi.StringPtrInput
	// The Web UI path to edit the member role
	EditPath pulumi.StringPtrInput
	// All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
	EnabledPermissions pulumi.StringArrayInput
	// Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
	GroupPath pulumi.StringPtrInput
	// The id integer value extracted from the `id` attribute
	Iid pulumi.IntPtrInput
	// Name for the member role.
	Name pulumi.StringPtrInput
}

func (MemberRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*memberRoleState)(nil)).Elem()
}

type memberRoleArgs struct {
	// The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
	BaseAccessLevel string `pulumi:"baseAccessLevel"`
	// Description for the member role.
	Description *string `pulumi:"description"`
	// All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
	EnabledPermissions []string `pulumi:"enabledPermissions"`
	// Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
	GroupPath *string `pulumi:"groupPath"`
	// Name for the member role.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a MemberRole resource.
type MemberRoleArgs struct {
	// The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
	BaseAccessLevel pulumi.StringInput
	// Description for the member role.
	Description pulumi.StringPtrInput
	// All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
	EnabledPermissions pulumi.StringArrayInput
	// Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
	GroupPath pulumi.StringPtrInput
	// Name for the member role.
	Name pulumi.StringPtrInput
}

func (MemberRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*memberRoleArgs)(nil)).Elem()
}

type MemberRoleInput interface {
	pulumi.Input

	ToMemberRoleOutput() MemberRoleOutput
	ToMemberRoleOutputWithContext(ctx context.Context) MemberRoleOutput
}

func (*MemberRole) ElementType() reflect.Type {
	return reflect.TypeOf((**MemberRole)(nil)).Elem()
}

func (i *MemberRole) ToMemberRoleOutput() MemberRoleOutput {
	return i.ToMemberRoleOutputWithContext(context.Background())
}

func (i *MemberRole) ToMemberRoleOutputWithContext(ctx context.Context) MemberRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberRoleOutput)
}

// MemberRoleArrayInput is an input type that accepts MemberRoleArray and MemberRoleArrayOutput values.
// You can construct a concrete instance of `MemberRoleArrayInput` via:
//
//	MemberRoleArray{ MemberRoleArgs{...} }
type MemberRoleArrayInput interface {
	pulumi.Input

	ToMemberRoleArrayOutput() MemberRoleArrayOutput
	ToMemberRoleArrayOutputWithContext(context.Context) MemberRoleArrayOutput
}

type MemberRoleArray []MemberRoleInput

func (MemberRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MemberRole)(nil)).Elem()
}

func (i MemberRoleArray) ToMemberRoleArrayOutput() MemberRoleArrayOutput {
	return i.ToMemberRoleArrayOutputWithContext(context.Background())
}

func (i MemberRoleArray) ToMemberRoleArrayOutputWithContext(ctx context.Context) MemberRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberRoleArrayOutput)
}

// MemberRoleMapInput is an input type that accepts MemberRoleMap and MemberRoleMapOutput values.
// You can construct a concrete instance of `MemberRoleMapInput` via:
//
//	MemberRoleMap{ "key": MemberRoleArgs{...} }
type MemberRoleMapInput interface {
	pulumi.Input

	ToMemberRoleMapOutput() MemberRoleMapOutput
	ToMemberRoleMapOutputWithContext(context.Context) MemberRoleMapOutput
}

type MemberRoleMap map[string]MemberRoleInput

func (MemberRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MemberRole)(nil)).Elem()
}

func (i MemberRoleMap) ToMemberRoleMapOutput() MemberRoleMapOutput {
	return i.ToMemberRoleMapOutputWithContext(context.Background())
}

func (i MemberRoleMap) ToMemberRoleMapOutputWithContext(ctx context.Context) MemberRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberRoleMapOutput)
}

type MemberRoleOutput struct{ *pulumi.OutputState }

func (MemberRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MemberRole)(nil)).Elem()
}

func (o MemberRoleOutput) ToMemberRoleOutput() MemberRoleOutput {
	return o
}

func (o MemberRoleOutput) ToMemberRoleOutputWithContext(ctx context.Context) MemberRoleOutput {
	return o
}

// The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
func (o MemberRoleOutput) BaseAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *MemberRole) pulumi.StringOutput { return v.BaseAccessLevel }).(pulumi.StringOutput)
}

// Timestamp of when the member role was created. Only available with GitLab version 17.3 or higher.
func (o MemberRoleOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *MemberRole) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description for the member role.
func (o MemberRoleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *MemberRole) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The Web UI path to edit the member role
func (o MemberRoleOutput) EditPath() pulumi.StringOutput {
	return o.ApplyT(func(v *MemberRole) pulumi.StringOutput { return v.EditPath }).(pulumi.StringOutput)
}

// All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
func (o MemberRoleOutput) EnabledPermissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MemberRole) pulumi.StringArrayOutput { return v.EnabledPermissions }).(pulumi.StringArrayOutput)
}

// Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
func (o MemberRoleOutput) GroupPath() pulumi.StringOutput {
	return o.ApplyT(func(v *MemberRole) pulumi.StringOutput { return v.GroupPath }).(pulumi.StringOutput)
}

// The id integer value extracted from the `id` attribute
func (o MemberRoleOutput) Iid() pulumi.IntOutput {
	return o.ApplyT(func(v *MemberRole) pulumi.IntOutput { return v.Iid }).(pulumi.IntOutput)
}

// Name for the member role.
func (o MemberRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MemberRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type MemberRoleArrayOutput struct{ *pulumi.OutputState }

func (MemberRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MemberRole)(nil)).Elem()
}

func (o MemberRoleArrayOutput) ToMemberRoleArrayOutput() MemberRoleArrayOutput {
	return o
}

func (o MemberRoleArrayOutput) ToMemberRoleArrayOutputWithContext(ctx context.Context) MemberRoleArrayOutput {
	return o
}

func (o MemberRoleArrayOutput) Index(i pulumi.IntInput) MemberRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MemberRole {
		return vs[0].([]*MemberRole)[vs[1].(int)]
	}).(MemberRoleOutput)
}

type MemberRoleMapOutput struct{ *pulumi.OutputState }

func (MemberRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MemberRole)(nil)).Elem()
}

func (o MemberRoleMapOutput) ToMemberRoleMapOutput() MemberRoleMapOutput {
	return o
}

func (o MemberRoleMapOutput) ToMemberRoleMapOutputWithContext(ctx context.Context) MemberRoleMapOutput {
	return o
}

func (o MemberRoleMapOutput) MapIndex(k pulumi.StringInput) MemberRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MemberRole {
		return vs[0].(map[string]*MemberRole)[vs[1].(string)]
	}).(MemberRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MemberRoleInput)(nil)).Elem(), &MemberRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberRoleArrayInput)(nil)).Elem(), MemberRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberRoleMapInput)(nil)).Elem(), MemberRoleMap{})
	pulumi.RegisterOutputType(MemberRoleOutput{})
	pulumi.RegisterOutputType(MemberRoleArrayOutput{})
	pulumi.RegisterOutputType(MemberRoleMapOutput{})
}
