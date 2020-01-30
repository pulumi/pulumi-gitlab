// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gitlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/group.html.markdown.
type Group struct {
	pulumi.CustomResourceState

	// The description of the group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The full name of the group.
	FullName pulumi.StringOutput `pulumi:"fullName"`
	// The full path of the group.
	FullPath pulumi.StringOutput `pulumi:"fullPath"`
	// Boolean, defaults to true.  Whether to enable LFS
	// support for projects in this group.
	LfsEnabled pulumi.BoolPtrOutput `pulumi:"lfsEnabled"`
	// The name of this group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Integer, id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrOutput `pulumi:"parentId"`
	// The path of the group.
	Path pulumi.StringOutput `pulumi:"path"`
	// Boolean, defaults to false.  Whether to
	// enable users to request access to the group.
	RequestAccessEnabled pulumi.BoolPtrOutput `pulumi:"requestAccessEnabled"`
	// The group level registration token to use during runner setup.
	RunnersToken pulumi.StringOutput `pulumi:"runnersToken"`
	// Set to `public` to create a public group.
	// Valid values are `private`, `internal`, `public`.
	// Groups are created as private by default.
	VisibilityLevel pulumi.StringOutput `pulumi:"visibilityLevel"`
	// Web URL of the group.
	WebUrl pulumi.StringOutput `pulumi:"webUrl"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil {
		args = &GroupArgs{}
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
	// The description of the group.
	Description *string `pulumi:"description"`
	// The full name of the group.
	FullName *string `pulumi:"fullName"`
	// The full path of the group.
	FullPath *string `pulumi:"fullPath"`
	// Boolean, defaults to true.  Whether to enable LFS
	// support for projects in this group.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// The name of this group.
	Name *string `pulumi:"name"`
	// Integer, id of the parent group (creates a nested group).
	ParentId *int `pulumi:"parentId"`
	// The path of the group.
	Path *string `pulumi:"path"`
	// Boolean, defaults to false.  Whether to
	// enable users to request access to the group.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// The group level registration token to use during runner setup.
	RunnersToken *string `pulumi:"runnersToken"`
	// Set to `public` to create a public group.
	// Valid values are `private`, `internal`, `public`.
	// Groups are created as private by default.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
	// Web URL of the group.
	WebUrl *string `pulumi:"webUrl"`
}

type GroupState struct {
	// The description of the group.
	Description pulumi.StringPtrInput
	// The full name of the group.
	FullName pulumi.StringPtrInput
	// The full path of the group.
	FullPath pulumi.StringPtrInput
	// Boolean, defaults to true.  Whether to enable LFS
	// support for projects in this group.
	LfsEnabled pulumi.BoolPtrInput
	// The name of this group.
	Name pulumi.StringPtrInput
	// Integer, id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrInput
	// The path of the group.
	Path pulumi.StringPtrInput
	// Boolean, defaults to false.  Whether to
	// enable users to request access to the group.
	RequestAccessEnabled pulumi.BoolPtrInput
	// The group level registration token to use during runner setup.
	RunnersToken pulumi.StringPtrInput
	// Set to `public` to create a public group.
	// Valid values are `private`, `internal`, `public`.
	// Groups are created as private by default.
	VisibilityLevel pulumi.StringPtrInput
	// Web URL of the group.
	WebUrl pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The description of the group.
	Description *string `pulumi:"description"`
	// Boolean, defaults to true.  Whether to enable LFS
	// support for projects in this group.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// The name of this group.
	Name *string `pulumi:"name"`
	// Integer, id of the parent group (creates a nested group).
	ParentId *int `pulumi:"parentId"`
	// The path of the group.
	Path string `pulumi:"path"`
	// Boolean, defaults to false.  Whether to
	// enable users to request access to the group.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Set to `public` to create a public group.
	// Valid values are `private`, `internal`, `public`.
	// Groups are created as private by default.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The description of the group.
	Description pulumi.StringPtrInput
	// Boolean, defaults to true.  Whether to enable LFS
	// support for projects in this group.
	LfsEnabled pulumi.BoolPtrInput
	// The name of this group.
	Name pulumi.StringPtrInput
	// Integer, id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrInput
	// The path of the group.
	Path pulumi.StringInput
	// Boolean, defaults to false.  Whether to
	// enable users to request access to the group.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Set to `public` to create a public group.
	// Valid values are `private`, `internal`, `public`.
	// Groups are created as private by default.
	VisibilityLevel pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}
