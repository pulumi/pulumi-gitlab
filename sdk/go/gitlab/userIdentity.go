// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `UserIdentity` resource is for managing the lifecycle of a user's external identity.
//
// > the provider needs to be configured with admin-level access for this resource to work.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/users/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := gitlab.NewUser(ctx, "example", &gitlab.UserArgs{
//				Name:           pulumi.String("Example Foo"),
//				Username:       pulumi.String("example"),
//				Email:          pulumi.String("gitlab@user.create"),
//				IsAdmin:        pulumi.Bool(true),
//				ProjectsLimit:  pulumi.Int(4),
//				CanCreateGroup: pulumi.Bool(false),
//				IsExternal:     pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewUserIdentity(ctx, "example", &gitlab.UserIdentityArgs{
//				UserId:           example.ID(),
//				ExternalProvider: pulumi.String("google"),
//				ExternalUid:      pulumi.String("1234567890"),
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
// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_user_identity`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_user_identity.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Importing using the CLI is supported with the following syntax:
//
// ```sh
// $ pulumi import gitlab:index/userIdentity:UserIdentity You can import a user identity to terraform state using `<resource> <id>`.
// ```
//
// The `id` must be a string for the id of the user and identity provider you want to import,
//
// for example:
//
// ```sh
// $ pulumi import gitlab:index/userIdentity:UserIdentity example "42:google"
// ```
type UserIdentity struct {
	pulumi.CustomResourceState

	// The external provider name.
	ExternalProvider pulumi.StringOutput `pulumi:"externalProvider"`
	// A specific external authentication provider UID.
	ExternalUid pulumi.StringOutput `pulumi:"externalUid"`
	// The GitLab ID of the user.
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewUserIdentity registers a new resource with the given unique name, arguments, and options.
func NewUserIdentity(ctx *pulumi.Context,
	name string, args *UserIdentityArgs, opts ...pulumi.ResourceOption) (*UserIdentity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExternalProvider == nil {
		return nil, errors.New("invalid value for required argument 'ExternalProvider'")
	}
	if args.ExternalUid == nil {
		return nil, errors.New("invalid value for required argument 'ExternalUid'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserIdentity
	err := ctx.RegisterResource("gitlab:index/userIdentity:UserIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserIdentity gets an existing UserIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserIdentityState, opts ...pulumi.ResourceOption) (*UserIdentity, error) {
	var resource UserIdentity
	err := ctx.ReadResource("gitlab:index/userIdentity:UserIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserIdentity resources.
type userIdentityState struct {
	// The external provider name.
	ExternalProvider *string `pulumi:"externalProvider"`
	// A specific external authentication provider UID.
	ExternalUid *string `pulumi:"externalUid"`
	// The GitLab ID of the user.
	UserId *int `pulumi:"userId"`
}

type UserIdentityState struct {
	// The external provider name.
	ExternalProvider pulumi.StringPtrInput
	// A specific external authentication provider UID.
	ExternalUid pulumi.StringPtrInput
	// The GitLab ID of the user.
	UserId pulumi.IntPtrInput
}

func (UserIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*userIdentityState)(nil)).Elem()
}

type userIdentityArgs struct {
	// The external provider name.
	ExternalProvider string `pulumi:"externalProvider"`
	// A specific external authentication provider UID.
	ExternalUid string `pulumi:"externalUid"`
	// The GitLab ID of the user.
	UserId int `pulumi:"userId"`
}

// The set of arguments for constructing a UserIdentity resource.
type UserIdentityArgs struct {
	// The external provider name.
	ExternalProvider pulumi.StringInput
	// A specific external authentication provider UID.
	ExternalUid pulumi.StringInput
	// The GitLab ID of the user.
	UserId pulumi.IntInput
}

func (UserIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userIdentityArgs)(nil)).Elem()
}

type UserIdentityInput interface {
	pulumi.Input

	ToUserIdentityOutput() UserIdentityOutput
	ToUserIdentityOutputWithContext(ctx context.Context) UserIdentityOutput
}

func (*UserIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((**UserIdentity)(nil)).Elem()
}

func (i *UserIdentity) ToUserIdentityOutput() UserIdentityOutput {
	return i.ToUserIdentityOutputWithContext(context.Background())
}

func (i *UserIdentity) ToUserIdentityOutputWithContext(ctx context.Context) UserIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityOutput)
}

// UserIdentityArrayInput is an input type that accepts UserIdentityArray and UserIdentityArrayOutput values.
// You can construct a concrete instance of `UserIdentityArrayInput` via:
//
//	UserIdentityArray{ UserIdentityArgs{...} }
type UserIdentityArrayInput interface {
	pulumi.Input

	ToUserIdentityArrayOutput() UserIdentityArrayOutput
	ToUserIdentityArrayOutputWithContext(context.Context) UserIdentityArrayOutput
}

type UserIdentityArray []UserIdentityInput

func (UserIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserIdentity)(nil)).Elem()
}

func (i UserIdentityArray) ToUserIdentityArrayOutput() UserIdentityArrayOutput {
	return i.ToUserIdentityArrayOutputWithContext(context.Background())
}

func (i UserIdentityArray) ToUserIdentityArrayOutputWithContext(ctx context.Context) UserIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityArrayOutput)
}

// UserIdentityMapInput is an input type that accepts UserIdentityMap and UserIdentityMapOutput values.
// You can construct a concrete instance of `UserIdentityMapInput` via:
//
//	UserIdentityMap{ "key": UserIdentityArgs{...} }
type UserIdentityMapInput interface {
	pulumi.Input

	ToUserIdentityMapOutput() UserIdentityMapOutput
	ToUserIdentityMapOutputWithContext(context.Context) UserIdentityMapOutput
}

type UserIdentityMap map[string]UserIdentityInput

func (UserIdentityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserIdentity)(nil)).Elem()
}

func (i UserIdentityMap) ToUserIdentityMapOutput() UserIdentityMapOutput {
	return i.ToUserIdentityMapOutputWithContext(context.Background())
}

func (i UserIdentityMap) ToUserIdentityMapOutputWithContext(ctx context.Context) UserIdentityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityMapOutput)
}

type UserIdentityOutput struct{ *pulumi.OutputState }

func (UserIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserIdentity)(nil)).Elem()
}

func (o UserIdentityOutput) ToUserIdentityOutput() UserIdentityOutput {
	return o
}

func (o UserIdentityOutput) ToUserIdentityOutputWithContext(ctx context.Context) UserIdentityOutput {
	return o
}

// The external provider name.
func (o UserIdentityOutput) ExternalProvider() pulumi.StringOutput {
	return o.ApplyT(func(v *UserIdentity) pulumi.StringOutput { return v.ExternalProvider }).(pulumi.StringOutput)
}

// A specific external authentication provider UID.
func (o UserIdentityOutput) ExternalUid() pulumi.StringOutput {
	return o.ApplyT(func(v *UserIdentity) pulumi.StringOutput { return v.ExternalUid }).(pulumi.StringOutput)
}

// The GitLab ID of the user.
func (o UserIdentityOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *UserIdentity) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type UserIdentityArrayOutput struct{ *pulumi.OutputState }

func (UserIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserIdentity)(nil)).Elem()
}

func (o UserIdentityArrayOutput) ToUserIdentityArrayOutput() UserIdentityArrayOutput {
	return o
}

func (o UserIdentityArrayOutput) ToUserIdentityArrayOutputWithContext(ctx context.Context) UserIdentityArrayOutput {
	return o
}

func (o UserIdentityArrayOutput) Index(i pulumi.IntInput) UserIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserIdentity {
		return vs[0].([]*UserIdentity)[vs[1].(int)]
	}).(UserIdentityOutput)
}

type UserIdentityMapOutput struct{ *pulumi.OutputState }

func (UserIdentityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserIdentity)(nil)).Elem()
}

func (o UserIdentityMapOutput) ToUserIdentityMapOutput() UserIdentityMapOutput {
	return o
}

func (o UserIdentityMapOutput) ToUserIdentityMapOutputWithContext(ctx context.Context) UserIdentityMapOutput {
	return o
}

func (o UserIdentityMapOutput) MapIndex(k pulumi.StringInput) UserIdentityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserIdentity {
		return vs[0].(map[string]*UserIdentity)[vs[1].(string)]
	}).(UserIdentityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserIdentityInput)(nil)).Elem(), &UserIdentity{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserIdentityArrayInput)(nil)).Elem(), UserIdentityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserIdentityMapInput)(nil)).Elem(), UserIdentityMap{})
	pulumi.RegisterOutputType(UserIdentityOutput{})
	pulumi.RegisterOutputType(UserIdentityArrayOutput{})
	pulumi.RegisterOutputType(UserIdentityMapOutput{})
}
