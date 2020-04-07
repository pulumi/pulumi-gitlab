// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gitlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource allows you to add a current user to an existing project with a set access level.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/project_membership.html.markdown.
type ProjectMembership struct {
	pulumi.CustomResourceState

	// One of five levels of access to the project.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// The id of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The id of the user.
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewProjectMembership registers a new resource with the given unique name, arguments, and options.
func NewProjectMembership(ctx *pulumi.Context,
	name string, args *ProjectMembershipArgs, opts ...pulumi.ResourceOption) (*ProjectMembership, error) {
	if args == nil || args.AccessLevel == nil {
		return nil, errors.New("missing required argument 'AccessLevel'")
	}
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil || args.UserId == nil {
		return nil, errors.New("missing required argument 'UserId'")
	}
	if args == nil {
		args = &ProjectMembershipArgs{}
	}
	var resource ProjectMembership
	err := ctx.RegisterResource("gitlab:index/projectMembership:ProjectMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectMembership gets an existing ProjectMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectMembershipState, opts ...pulumi.ResourceOption) (*ProjectMembership, error) {
	var resource ProjectMembership
	err := ctx.ReadResource("gitlab:index/projectMembership:ProjectMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectMembership resources.
type projectMembershipState struct {
	// One of five levels of access to the project.
	AccessLevel *string `pulumi:"accessLevel"`
	// The id of the project.
	ProjectId *string `pulumi:"projectId"`
	// The id of the user.
	UserId *int `pulumi:"userId"`
}

type ProjectMembershipState struct {
	// One of five levels of access to the project.
	AccessLevel pulumi.StringPtrInput
	// The id of the project.
	ProjectId pulumi.StringPtrInput
	// The id of the user.
	UserId pulumi.IntPtrInput
}

func (ProjectMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMembershipState)(nil)).Elem()
}

type projectMembershipArgs struct {
	// One of five levels of access to the project.
	AccessLevel string `pulumi:"accessLevel"`
	// The id of the project.
	ProjectId string `pulumi:"projectId"`
	// The id of the user.
	UserId int `pulumi:"userId"`
}

// The set of arguments for constructing a ProjectMembership resource.
type ProjectMembershipArgs struct {
	// One of five levels of access to the project.
	AccessLevel pulumi.StringInput
	// The id of the project.
	ProjectId pulumi.StringInput
	// The id of the user.
	UserId pulumi.IntInput
}

func (ProjectMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMembershipArgs)(nil)).Elem()
}
