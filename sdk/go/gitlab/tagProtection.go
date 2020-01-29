// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gitlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource allows you to protect a specific tag or wildcard by an access level so that the user with less access level cannot Create the tags.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/tag_protection.html.markdown.
type TagProtection struct {
	pulumi.CustomResourceState

	// One of five levels of access to the project.
	CreateAccessLevel pulumi.StringOutput `pulumi:"createAccessLevel"`
	// The id of the project.
	Project pulumi.StringOutput `pulumi:"project"`
	// Name of the tag or wildcard.
	Tag pulumi.StringOutput `pulumi:"tag"`
}

// NewTagProtection registers a new resource with the given unique name, arguments, and options.
func NewTagProtection(ctx *pulumi.Context,
	name string, args *TagProtectionArgs, opts ...pulumi.ResourceOption) (*TagProtection, error) {
	if args == nil || args.CreateAccessLevel == nil {
		return nil, errors.New("missing required argument 'CreateAccessLevel'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil || args.Tag == nil {
		return nil, errors.New("missing required argument 'Tag'")
	}
	if args == nil {
		args = &TagProtectionArgs{}
	}
	var resource TagProtection
	err := ctx.RegisterResource("gitlab:index/tagProtection:TagProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagProtection gets an existing TagProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagProtectionState, opts ...pulumi.ResourceOption) (*TagProtection, error) {
	var resource TagProtection
	err := ctx.ReadResource("gitlab:index/tagProtection:TagProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagProtection resources.
type tagProtectionState struct {
	// One of five levels of access to the project.
	CreateAccessLevel *string `pulumi:"createAccessLevel"`
	// The id of the project.
	Project *string `pulumi:"project"`
	// Name of the tag or wildcard.
	Tag *string `pulumi:"tag"`
}

type TagProtectionState struct {
	// One of five levels of access to the project.
	CreateAccessLevel pulumi.StringPtrInput
	// The id of the project.
	Project pulumi.StringPtrInput
	// Name of the tag or wildcard.
	Tag pulumi.StringPtrInput
}

func (TagProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagProtectionState)(nil)).Elem()
}

type tagProtectionArgs struct {
	// One of five levels of access to the project.
	CreateAccessLevel string `pulumi:"createAccessLevel"`
	// The id of the project.
	Project string `pulumi:"project"`
	// Name of the tag or wildcard.
	Tag string `pulumi:"tag"`
}

// The set of arguments for constructing a TagProtection resource.
type TagProtectionArgs struct {
	// One of five levels of access to the project.
	CreateAccessLevel pulumi.StringInput
	// The id of the project.
	Project pulumi.StringInput
	// Name of the tag or wildcard.
	Tag pulumi.StringInput
}

func (TagProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagProtectionArgs)(nil)).Elem()
}

