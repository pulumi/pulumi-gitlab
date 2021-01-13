// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type User struct {
	pulumi.CustomResourceState

	// Boolean, defaults to false. Whether to allow the user to create groups.
	CanCreateGroup pulumi.BoolPtrOutput `pulumi:"canCreateGroup"`
	// The e-mail address of the user.
	Email pulumi.StringOutput `pulumi:"email"`
	// Boolean, defaults to false.  Whether to enable administrative priviledges
	// for the user.
	IsAdmin pulumi.BoolPtrOutput `pulumi:"isAdmin"`
	// Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
	IsExternal pulumi.BoolPtrOutput `pulumi:"isExternal"`
	// The name of the user.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password of the user.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Integer, defaults to 0.  Number of projects user can create.
	ProjectsLimit pulumi.IntPtrOutput `pulumi:"projectsLimit"`
	// Boolean, defaults to false. Send user password reset link.
	ResetPassword pulumi.BoolPtrOutput `pulumi:"resetPassword"`
	// Boolean, defaults to true. Whether to skip confirmation.
	SkipConfirmation pulumi.BoolPtrOutput `pulumi:"skipConfirmation"`
	// The username of the user.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource User
	err := ctx.RegisterResource("gitlab:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("gitlab:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Boolean, defaults to false. Whether to allow the user to create groups.
	CanCreateGroup *bool `pulumi:"canCreateGroup"`
	// The e-mail address of the user.
	Email *string `pulumi:"email"`
	// Boolean, defaults to false.  Whether to enable administrative priviledges
	// for the user.
	IsAdmin *bool `pulumi:"isAdmin"`
	// Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
	IsExternal *bool `pulumi:"isExternal"`
	// The name of the user.
	Name *string `pulumi:"name"`
	// The password of the user.
	Password *string `pulumi:"password"`
	// Integer, defaults to 0.  Number of projects user can create.
	ProjectsLimit *int `pulumi:"projectsLimit"`
	// Boolean, defaults to false. Send user password reset link.
	ResetPassword *bool `pulumi:"resetPassword"`
	// Boolean, defaults to true. Whether to skip confirmation.
	SkipConfirmation *bool `pulumi:"skipConfirmation"`
	// The username of the user.
	Username *string `pulumi:"username"`
}

type UserState struct {
	// Boolean, defaults to false. Whether to allow the user to create groups.
	CanCreateGroup pulumi.BoolPtrInput
	// The e-mail address of the user.
	Email pulumi.StringPtrInput
	// Boolean, defaults to false.  Whether to enable administrative priviledges
	// for the user.
	IsAdmin pulumi.BoolPtrInput
	// Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
	IsExternal pulumi.BoolPtrInput
	// The name of the user.
	Name pulumi.StringPtrInput
	// The password of the user.
	Password pulumi.StringPtrInput
	// Integer, defaults to 0.  Number of projects user can create.
	ProjectsLimit pulumi.IntPtrInput
	// Boolean, defaults to false. Send user password reset link.
	ResetPassword pulumi.BoolPtrInput
	// Boolean, defaults to true. Whether to skip confirmation.
	SkipConfirmation pulumi.BoolPtrInput
	// The username of the user.
	Username pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Boolean, defaults to false. Whether to allow the user to create groups.
	CanCreateGroup *bool `pulumi:"canCreateGroup"`
	// The e-mail address of the user.
	Email string `pulumi:"email"`
	// Boolean, defaults to false.  Whether to enable administrative priviledges
	// for the user.
	IsAdmin *bool `pulumi:"isAdmin"`
	// Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
	IsExternal *bool `pulumi:"isExternal"`
	// The name of the user.
	Name *string `pulumi:"name"`
	// The password of the user.
	Password *string `pulumi:"password"`
	// Integer, defaults to 0.  Number of projects user can create.
	ProjectsLimit *int `pulumi:"projectsLimit"`
	// Boolean, defaults to false. Send user password reset link.
	ResetPassword *bool `pulumi:"resetPassword"`
	// Boolean, defaults to true. Whether to skip confirmation.
	SkipConfirmation *bool `pulumi:"skipConfirmation"`
	// The username of the user.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Boolean, defaults to false. Whether to allow the user to create groups.
	CanCreateGroup pulumi.BoolPtrInput
	// The e-mail address of the user.
	Email pulumi.StringInput
	// Boolean, defaults to false.  Whether to enable administrative priviledges
	// for the user.
	IsAdmin pulumi.BoolPtrInput
	// Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
	IsExternal pulumi.BoolPtrInput
	// The name of the user.
	Name pulumi.StringPtrInput
	// The password of the user.
	Password pulumi.StringPtrInput
	// Integer, defaults to 0.  Number of projects user can create.
	ProjectsLimit pulumi.IntPtrInput
	// Boolean, defaults to false. Send user password reset link.
	ResetPassword pulumi.BoolPtrInput
	// Boolean, defaults to true. Whether to skip confirmation.
	SkipConfirmation pulumi.BoolPtrInput
	// The username of the user.
	Username pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (User) ElementType() reflect.Type {
	return reflect.TypeOf((*User)(nil)).Elem()
}

func (i User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

type UserOutput struct {
	*pulumi.OutputState
}

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserOutput)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
}
