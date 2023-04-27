// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `UserGpgKey` resource allows to manage the lifecycle of a GPG key assigned to the current user or a specific user.
//
// > Managing GPG keys for arbitrary users requires admin privileges.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#get-a-specific-gpg-key)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleUser, err := gitlab.LookupUser(ctx, &gitlab.LookupUserArgs{
//				Username: pulumi.StringRef("example-user"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewUserGpgKey(ctx, "exampleUserGpgKey", &gitlab.UserGpgKeyArgs{
//				UserId: *pulumi.String(exampleUser.Id),
//				Key:    pulumi.String("-----BEGIN PGP PUBLIC KEY BLOCK-----\n...\n-----END PGP PUBLIC KEY BLOCK-----"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewUserGpgKey(ctx, "exampleUserUserGpgKey", &gitlab.UserGpgKeyArgs{
//				Key: pulumi.String("-----BEGIN PGP PUBLIC KEY BLOCK-----\n...\n-----END PGP PUBLIC KEY BLOCK-----"),
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
// You can import a GPG key for a specific user using an id made up of `{user-id}:{key}`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/userGpgKey:UserGpgKey example 42:1
//
// ```
//
//	Alternatively, you can import a GPG key for the current user using an id made up of `{key}`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/userGpgKey:UserGpgKey example_user 1
//
// ```
type UserGpgKey struct {
	pulumi.CustomResourceState

	// The time when this key was created in GitLab.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The armored GPG public key.
	Key pulumi.StringOutput `pulumi:"key"`
	// The ID of the GPG key.
	KeyId pulumi.IntOutput `pulumi:"keyId"`
	// The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the speicifed user, and an admin token is required.
	UserId pulumi.IntPtrOutput `pulumi:"userId"`
}

// NewUserGpgKey registers a new resource with the given unique name, arguments, and options.
func NewUserGpgKey(ctx *pulumi.Context,
	name string, args *UserGpgKeyArgs, opts ...pulumi.ResourceOption) (*UserGpgKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource UserGpgKey
	err := ctx.RegisterResource("gitlab:index/userGpgKey:UserGpgKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserGpgKey gets an existing UserGpgKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserGpgKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserGpgKeyState, opts ...pulumi.ResourceOption) (*UserGpgKey, error) {
	var resource UserGpgKey
	err := ctx.ReadResource("gitlab:index/userGpgKey:UserGpgKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserGpgKey resources.
type userGpgKeyState struct {
	// The time when this key was created in GitLab.
	CreatedAt *string `pulumi:"createdAt"`
	// The armored GPG public key.
	Key *string `pulumi:"key"`
	// The ID of the GPG key.
	KeyId *int `pulumi:"keyId"`
	// The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the speicifed user, and an admin token is required.
	UserId *int `pulumi:"userId"`
}

type UserGpgKeyState struct {
	// The time when this key was created in GitLab.
	CreatedAt pulumi.StringPtrInput
	// The armored GPG public key.
	Key pulumi.StringPtrInput
	// The ID of the GPG key.
	KeyId pulumi.IntPtrInput
	// The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the speicifed user, and an admin token is required.
	UserId pulumi.IntPtrInput
}

func (UserGpgKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*userGpgKeyState)(nil)).Elem()
}

type userGpgKeyArgs struct {
	// The armored GPG public key.
	Key string `pulumi:"key"`
	// The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the speicifed user, and an admin token is required.
	UserId *int `pulumi:"userId"`
}

// The set of arguments for constructing a UserGpgKey resource.
type UserGpgKeyArgs struct {
	// The armored GPG public key.
	Key pulumi.StringInput
	// The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the speicifed user, and an admin token is required.
	UserId pulumi.IntPtrInput
}

func (UserGpgKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userGpgKeyArgs)(nil)).Elem()
}

type UserGpgKeyInput interface {
	pulumi.Input

	ToUserGpgKeyOutput() UserGpgKeyOutput
	ToUserGpgKeyOutputWithContext(ctx context.Context) UserGpgKeyOutput
}

func (*UserGpgKey) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGpgKey)(nil)).Elem()
}

func (i *UserGpgKey) ToUserGpgKeyOutput() UserGpgKeyOutput {
	return i.ToUserGpgKeyOutputWithContext(context.Background())
}

func (i *UserGpgKey) ToUserGpgKeyOutputWithContext(ctx context.Context) UserGpgKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGpgKeyOutput)
}

// UserGpgKeyArrayInput is an input type that accepts UserGpgKeyArray and UserGpgKeyArrayOutput values.
// You can construct a concrete instance of `UserGpgKeyArrayInput` via:
//
//	UserGpgKeyArray{ UserGpgKeyArgs{...} }
type UserGpgKeyArrayInput interface {
	pulumi.Input

	ToUserGpgKeyArrayOutput() UserGpgKeyArrayOutput
	ToUserGpgKeyArrayOutputWithContext(context.Context) UserGpgKeyArrayOutput
}

type UserGpgKeyArray []UserGpgKeyInput

func (UserGpgKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGpgKey)(nil)).Elem()
}

func (i UserGpgKeyArray) ToUserGpgKeyArrayOutput() UserGpgKeyArrayOutput {
	return i.ToUserGpgKeyArrayOutputWithContext(context.Background())
}

func (i UserGpgKeyArray) ToUserGpgKeyArrayOutputWithContext(ctx context.Context) UserGpgKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGpgKeyArrayOutput)
}

// UserGpgKeyMapInput is an input type that accepts UserGpgKeyMap and UserGpgKeyMapOutput values.
// You can construct a concrete instance of `UserGpgKeyMapInput` via:
//
//	UserGpgKeyMap{ "key": UserGpgKeyArgs{...} }
type UserGpgKeyMapInput interface {
	pulumi.Input

	ToUserGpgKeyMapOutput() UserGpgKeyMapOutput
	ToUserGpgKeyMapOutputWithContext(context.Context) UserGpgKeyMapOutput
}

type UserGpgKeyMap map[string]UserGpgKeyInput

func (UserGpgKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGpgKey)(nil)).Elem()
}

func (i UserGpgKeyMap) ToUserGpgKeyMapOutput() UserGpgKeyMapOutput {
	return i.ToUserGpgKeyMapOutputWithContext(context.Background())
}

func (i UserGpgKeyMap) ToUserGpgKeyMapOutputWithContext(ctx context.Context) UserGpgKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGpgKeyMapOutput)
}

type UserGpgKeyOutput struct{ *pulumi.OutputState }

func (UserGpgKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGpgKey)(nil)).Elem()
}

func (o UserGpgKeyOutput) ToUserGpgKeyOutput() UserGpgKeyOutput {
	return o
}

func (o UserGpgKeyOutput) ToUserGpgKeyOutputWithContext(ctx context.Context) UserGpgKeyOutput {
	return o
}

// The time when this key was created in GitLab.
func (o UserGpgKeyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGpgKey) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The armored GPG public key.
func (o UserGpgKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGpgKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The ID of the GPG key.
func (o UserGpgKeyOutput) KeyId() pulumi.IntOutput {
	return o.ApplyT(func(v *UserGpgKey) pulumi.IntOutput { return v.KeyId }).(pulumi.IntOutput)
}

// The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the speicifed user, and an admin token is required.
func (o UserGpgKeyOutput) UserId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserGpgKey) pulumi.IntPtrOutput { return v.UserId }).(pulumi.IntPtrOutput)
}

type UserGpgKeyArrayOutput struct{ *pulumi.OutputState }

func (UserGpgKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGpgKey)(nil)).Elem()
}

func (o UserGpgKeyArrayOutput) ToUserGpgKeyArrayOutput() UserGpgKeyArrayOutput {
	return o
}

func (o UserGpgKeyArrayOutput) ToUserGpgKeyArrayOutputWithContext(ctx context.Context) UserGpgKeyArrayOutput {
	return o
}

func (o UserGpgKeyArrayOutput) Index(i pulumi.IntInput) UserGpgKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserGpgKey {
		return vs[0].([]*UserGpgKey)[vs[1].(int)]
	}).(UserGpgKeyOutput)
}

type UserGpgKeyMapOutput struct{ *pulumi.OutputState }

func (UserGpgKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGpgKey)(nil)).Elem()
}

func (o UserGpgKeyMapOutput) ToUserGpgKeyMapOutput() UserGpgKeyMapOutput {
	return o
}

func (o UserGpgKeyMapOutput) ToUserGpgKeyMapOutputWithContext(ctx context.Context) UserGpgKeyMapOutput {
	return o
}

func (o UserGpgKeyMapOutput) MapIndex(k pulumi.StringInput) UserGpgKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserGpgKey {
		return vs[0].(map[string]*UserGpgKey)[vs[1].(string)]
	}).(UserGpgKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserGpgKeyInput)(nil)).Elem(), &UserGpgKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGpgKeyArrayInput)(nil)).Elem(), UserGpgKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGpgKeyMapInput)(nil)).Elem(), UserGpgKeyMap{})
	pulumi.RegisterOutputType(UserGpgKeyOutput{})
	pulumi.RegisterOutputType(UserGpgKeyArrayOutput{})
	pulumi.RegisterOutputType(UserGpgKeyMapOutput{})
}
