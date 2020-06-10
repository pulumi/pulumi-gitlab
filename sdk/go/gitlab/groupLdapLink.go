// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to add an LDAP link to an existing GitLab group.
type GroupLdapLink struct {
	pulumi.CustomResourceState

	// Acceptable values are: guest, reporter, developer, maintainer, owner.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// The CN of the LDAP group to link with.
	Cn    pulumi.StringOutput  `pulumi:"cn"`
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The id of the GitLab group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the LDAP provider as stored in the GitLab database.
	LdapProvider pulumi.StringOutput `pulumi:"ldapProvider"`
}

// NewGroupLdapLink registers a new resource with the given unique name, arguments, and options.
func NewGroupLdapLink(ctx *pulumi.Context,
	name string, args *GroupLdapLinkArgs, opts ...pulumi.ResourceOption) (*GroupLdapLink, error) {
	if args == nil || args.AccessLevel == nil {
		return nil, errors.New("missing required argument 'AccessLevel'")
	}
	if args == nil || args.Cn == nil {
		return nil, errors.New("missing required argument 'Cn'")
	}
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	if args == nil || args.LdapProvider == nil {
		return nil, errors.New("missing required argument 'LdapProvider'")
	}
	if args == nil {
		args = &GroupLdapLinkArgs{}
	}
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
	// Acceptable values are: guest, reporter, developer, maintainer, owner.
	AccessLevel *string `pulumi:"accessLevel"`
	// The CN of the LDAP group to link with.
	Cn    *string `pulumi:"cn"`
	Force *bool   `pulumi:"force"`
	// The id of the GitLab group.
	GroupId *string `pulumi:"groupId"`
	// The name of the LDAP provider as stored in the GitLab database.
	LdapProvider *string `pulumi:"ldapProvider"`
}

type GroupLdapLinkState struct {
	// Acceptable values are: guest, reporter, developer, maintainer, owner.
	AccessLevel pulumi.StringPtrInput
	// The CN of the LDAP group to link with.
	Cn    pulumi.StringPtrInput
	Force pulumi.BoolPtrInput
	// The id of the GitLab group.
	GroupId pulumi.StringPtrInput
	// The name of the LDAP provider as stored in the GitLab database.
	LdapProvider pulumi.StringPtrInput
}

func (GroupLdapLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupLdapLinkState)(nil)).Elem()
}

type groupLdapLinkArgs struct {
	// Acceptable values are: guest, reporter, developer, maintainer, owner.
	AccessLevel string `pulumi:"accessLevel"`
	// The CN of the LDAP group to link with.
	Cn    string `pulumi:"cn"`
	Force *bool  `pulumi:"force"`
	// The id of the GitLab group.
	GroupId string `pulumi:"groupId"`
	// The name of the LDAP provider as stored in the GitLab database.
	LdapProvider string `pulumi:"ldapProvider"`
}

// The set of arguments for constructing a GroupLdapLink resource.
type GroupLdapLinkArgs struct {
	// Acceptable values are: guest, reporter, developer, maintainer, owner.
	AccessLevel pulumi.StringInput
	// The CN of the LDAP group to link with.
	Cn    pulumi.StringInput
	Force pulumi.BoolPtrInput
	// The id of the GitLab group.
	GroupId pulumi.StringInput
	// The name of the LDAP provider as stored in the GitLab database.
	LdapProvider pulumi.StringInput
}

func (GroupLdapLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupLdapLinkArgs)(nil)).Elem()
}