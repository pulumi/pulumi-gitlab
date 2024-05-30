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

// The `GroupLdapLink` resource allows to manage the lifecycle of an LDAP integration with a group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#ldap-group-links)
//
// ## Import
//
// GitLab group ldap links can be imported using an id made up of `group_id:ldap_provider:cn:filter`. CN and Filter are mutually exclusive, so one will be missing.
//
// If using the CN for the group link, the ID will end with a blank filter (":"). e.g.,
//
// ```sh
// $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain:testcn:"
// ```
//
// If using the Filter for the group link, the ID will have two "::" in the middle due to having a blank CN. e.g.,
//
// ```sh
// $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain::testfilter"
// ```
type GroupLdapLink struct {
	pulumi.CustomResourceState

	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
	//
	// Deprecated: Use `groupAccess` instead of the `accessLevel` attribute.
	AccessLevel pulumi.StringPtrOutput `pulumi:"accessLevel"`
	// The CN of the LDAP group to link with. Required if `filter` is not provided.
	Cn pulumi.StringOutput `pulumi:"cn"`
	// The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The ID or URL-encoded path of the group
	Group pulumi.StringOutput `pulumi:"group"`
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
	GroupAccess pulumi.StringPtrOutput `pulumi:"groupAccess"`
	// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	LdapProvider pulumi.StringOutput `pulumi:"ldapProvider"`
}

// NewGroupLdapLink registers a new resource with the given unique name, arguments, and options.
func NewGroupLdapLink(ctx *pulumi.Context,
	name string, args *GroupLdapLinkArgs, opts ...pulumi.ResourceOption) (*GroupLdapLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.LdapProvider == nil {
		return nil, errors.New("invalid value for required argument 'LdapProvider'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupLdapLink
	err := ctx.RegisterResource("gitlab:index/groupLdapLink:GroupLdapLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupLdapLink gets an existing GroupLdapLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupLdapLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupLdapLinkState, opts ...pulumi.ResourceOption) (*GroupLdapLink, error) {
	var resource GroupLdapLink
	err := ctx.ReadResource("gitlab:index/groupLdapLink:GroupLdapLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupLdapLink resources.
type groupLdapLinkState struct {
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
	//
	// Deprecated: Use `groupAccess` instead of the `accessLevel` attribute.
	AccessLevel *string `pulumi:"accessLevel"`
	// The CN of the LDAP group to link with. Required if `filter` is not provided.
	Cn *string `pulumi:"cn"`
	// The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
	Filter *string `pulumi:"filter"`
	// If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
	Force *bool `pulumi:"force"`
	// The ID or URL-encoded path of the group
	Group *string `pulumi:"group"`
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
	GroupAccess *string `pulumi:"groupAccess"`
	// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	LdapProvider *string `pulumi:"ldapProvider"`
}

type GroupLdapLinkState struct {
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
	//
	// Deprecated: Use `groupAccess` instead of the `accessLevel` attribute.
	AccessLevel pulumi.StringPtrInput
	// The CN of the LDAP group to link with. Required if `filter` is not provided.
	Cn pulumi.StringPtrInput
	// The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
	Filter pulumi.StringPtrInput
	// If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
	Force pulumi.BoolPtrInput
	// The ID or URL-encoded path of the group
	Group pulumi.StringPtrInput
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
	GroupAccess pulumi.StringPtrInput
	// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	LdapProvider pulumi.StringPtrInput
}

func (GroupLdapLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupLdapLinkState)(nil)).Elem()
}

type groupLdapLinkArgs struct {
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
	//
	// Deprecated: Use `groupAccess` instead of the `accessLevel` attribute.
	AccessLevel *string `pulumi:"accessLevel"`
	// The CN of the LDAP group to link with. Required if `filter` is not provided.
	Cn *string `pulumi:"cn"`
	// The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
	Filter *string `pulumi:"filter"`
	// If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
	Force *bool `pulumi:"force"`
	// The ID or URL-encoded path of the group
	Group string `pulumi:"group"`
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
	GroupAccess *string `pulumi:"groupAccess"`
	// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	LdapProvider string `pulumi:"ldapProvider"`
}

// The set of arguments for constructing a GroupLdapLink resource.
type GroupLdapLinkArgs struct {
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
	//
	// Deprecated: Use `groupAccess` instead of the `accessLevel` attribute.
	AccessLevel pulumi.StringPtrInput
	// The CN of the LDAP group to link with. Required if `filter` is not provided.
	Cn pulumi.StringPtrInput
	// The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
	Filter pulumi.StringPtrInput
	// If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
	Force pulumi.BoolPtrInput
	// The ID or URL-encoded path of the group
	Group pulumi.StringInput
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
	GroupAccess pulumi.StringPtrInput
	// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	LdapProvider pulumi.StringInput
}

func (GroupLdapLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupLdapLinkArgs)(nil)).Elem()
}

type GroupLdapLinkInput interface {
	pulumi.Input

	ToGroupLdapLinkOutput() GroupLdapLinkOutput
	ToGroupLdapLinkOutputWithContext(ctx context.Context) GroupLdapLinkOutput
}

func (*GroupLdapLink) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupLdapLink)(nil)).Elem()
}

func (i *GroupLdapLink) ToGroupLdapLinkOutput() GroupLdapLinkOutput {
	return i.ToGroupLdapLinkOutputWithContext(context.Background())
}

func (i *GroupLdapLink) ToGroupLdapLinkOutputWithContext(ctx context.Context) GroupLdapLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLdapLinkOutput)
}

// GroupLdapLinkArrayInput is an input type that accepts GroupLdapLinkArray and GroupLdapLinkArrayOutput values.
// You can construct a concrete instance of `GroupLdapLinkArrayInput` via:
//
//	GroupLdapLinkArray{ GroupLdapLinkArgs{...} }
type GroupLdapLinkArrayInput interface {
	pulumi.Input

	ToGroupLdapLinkArrayOutput() GroupLdapLinkArrayOutput
	ToGroupLdapLinkArrayOutputWithContext(context.Context) GroupLdapLinkArrayOutput
}

type GroupLdapLinkArray []GroupLdapLinkInput

func (GroupLdapLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupLdapLink)(nil)).Elem()
}

func (i GroupLdapLinkArray) ToGroupLdapLinkArrayOutput() GroupLdapLinkArrayOutput {
	return i.ToGroupLdapLinkArrayOutputWithContext(context.Background())
}

func (i GroupLdapLinkArray) ToGroupLdapLinkArrayOutputWithContext(ctx context.Context) GroupLdapLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLdapLinkArrayOutput)
}

// GroupLdapLinkMapInput is an input type that accepts GroupLdapLinkMap and GroupLdapLinkMapOutput values.
// You can construct a concrete instance of `GroupLdapLinkMapInput` via:
//
//	GroupLdapLinkMap{ "key": GroupLdapLinkArgs{...} }
type GroupLdapLinkMapInput interface {
	pulumi.Input

	ToGroupLdapLinkMapOutput() GroupLdapLinkMapOutput
	ToGroupLdapLinkMapOutputWithContext(context.Context) GroupLdapLinkMapOutput
}

type GroupLdapLinkMap map[string]GroupLdapLinkInput

func (GroupLdapLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupLdapLink)(nil)).Elem()
}

func (i GroupLdapLinkMap) ToGroupLdapLinkMapOutput() GroupLdapLinkMapOutput {
	return i.ToGroupLdapLinkMapOutputWithContext(context.Background())
}

func (i GroupLdapLinkMap) ToGroupLdapLinkMapOutputWithContext(ctx context.Context) GroupLdapLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLdapLinkMapOutput)
}

type GroupLdapLinkOutput struct{ *pulumi.OutputState }

func (GroupLdapLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupLdapLink)(nil)).Elem()
}

func (o GroupLdapLinkOutput) ToGroupLdapLinkOutput() GroupLdapLinkOutput {
	return o
}

func (o GroupLdapLinkOutput) ToGroupLdapLinkOutputWithContext(ctx context.Context) GroupLdapLinkOutput {
	return o
}

// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
//
// Deprecated: Use `groupAccess` instead of the `accessLevel` attribute.
func (o GroupLdapLinkOutput) AccessLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringPtrOutput { return v.AccessLevel }).(pulumi.StringPtrOutput)
}

// The CN of the LDAP group to link with. Required if `filter` is not provided.
func (o GroupLdapLinkOutput) Cn() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringOutput { return v.Cn }).(pulumi.StringOutput)
}

// The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
func (o GroupLdapLinkOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
func (o GroupLdapLinkOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// The ID or URL-encoded path of the group
func (o GroupLdapLinkOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
func (o GroupLdapLinkOutput) GroupAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringPtrOutput { return v.GroupAccess }).(pulumi.StringPtrOutput)
}

// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
func (o GroupLdapLinkOutput) LdapProvider() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringOutput { return v.LdapProvider }).(pulumi.StringOutput)
}

type GroupLdapLinkArrayOutput struct{ *pulumi.OutputState }

func (GroupLdapLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupLdapLink)(nil)).Elem()
}

func (o GroupLdapLinkArrayOutput) ToGroupLdapLinkArrayOutput() GroupLdapLinkArrayOutput {
	return o
}

func (o GroupLdapLinkArrayOutput) ToGroupLdapLinkArrayOutputWithContext(ctx context.Context) GroupLdapLinkArrayOutput {
	return o
}

func (o GroupLdapLinkArrayOutput) Index(i pulumi.IntInput) GroupLdapLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupLdapLink {
		return vs[0].([]*GroupLdapLink)[vs[1].(int)]
	}).(GroupLdapLinkOutput)
}

type GroupLdapLinkMapOutput struct{ *pulumi.OutputState }

func (GroupLdapLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupLdapLink)(nil)).Elem()
}

func (o GroupLdapLinkMapOutput) ToGroupLdapLinkMapOutput() GroupLdapLinkMapOutput {
	return o
}

func (o GroupLdapLinkMapOutput) ToGroupLdapLinkMapOutputWithContext(ctx context.Context) GroupLdapLinkMapOutput {
	return o
}

func (o GroupLdapLinkMapOutput) MapIndex(k pulumi.StringInput) GroupLdapLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupLdapLink {
		return vs[0].(map[string]*GroupLdapLink)[vs[1].(string)]
	}).(GroupLdapLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupLdapLinkInput)(nil)).Elem(), &GroupLdapLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupLdapLinkArrayInput)(nil)).Elem(), GroupLdapLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupLdapLinkMapInput)(nil)).Elem(), GroupLdapLinkMap{})
	pulumi.RegisterOutputType(GroupLdapLinkOutput{})
	pulumi.RegisterOutputType(GroupLdapLinkArrayOutput{})
	pulumi.RegisterOutputType(GroupLdapLinkMapOutput{})
}
