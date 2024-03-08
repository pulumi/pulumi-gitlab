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

// The `UserSshKey` resource allows to manage the lifecycle of an SSH key assigned to a user.
//
// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/users.html#single-ssh-key)
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			exampleUser, err := gitlab.LookupUser(ctx, &gitlab.LookupUserArgs{
//				Username: pulumi.StringRef("example-user"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewUserSshKey(ctx, "exampleUserSshKey", &gitlab.UserSshKeyArgs{
//				UserId:    *pulumi.String(exampleUser.Id),
//				Title:     pulumi.String("example-key"),
//				Key:       pulumi.String("ssh-ed25519 AAAA..."),
//				ExpiresAt: pulumi.String("2016-01-21T00:00:00.000Z"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// You can import a user ssh key using an id made up of `{user-id}:{key}`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/userSshKey:UserSshKey example 42:1
// ```
type UserSshKey struct {
	pulumi.CustomResourceState

	// The time when this key was created in GitLab.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
	Key pulumi.StringOutput `pulumi:"key"`
	// The ID of the ssh key.
	KeyId pulumi.IntOutput `pulumi:"keyId"`
	// The title of the ssh key.
	Title pulumi.StringOutput `pulumi:"title"`
	// The ID or username of the user. If this field is omitted, this resource manages a SSH key for the current user. Otherwise, this resource manages a SSH key for the specified user, and an admin token is required.
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewUserSshKey registers a new resource with the given unique name, arguments, and options.
func NewUserSshKey(ctx *pulumi.Context,
	name string, args *UserSshKeyArgs, opts ...pulumi.ResourceOption) (*UserSshKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserSshKey
	err := ctx.RegisterResource("gitlab:index/userSshKey:UserSshKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserSshKey gets an existing UserSshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserSshKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserSshKeyState, opts ...pulumi.ResourceOption) (*UserSshKey, error) {
	var resource UserSshKey
	err := ctx.ReadResource("gitlab:index/userSshKey:UserSshKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserSshKey resources.
type userSshKeyState struct {
	// The time when this key was created in GitLab.
	CreatedAt *string `pulumi:"createdAt"`
	// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
	ExpiresAt *string `pulumi:"expiresAt"`
	// The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
	Key *string `pulumi:"key"`
	// The ID of the ssh key.
	KeyId *int `pulumi:"keyId"`
	// The title of the ssh key.
	Title *string `pulumi:"title"`
	// The ID or username of the user. If this field is omitted, this resource manages a SSH key for the current user. Otherwise, this resource manages a SSH key for the specified user, and an admin token is required.
	UserId *int `pulumi:"userId"`
}

type UserSshKeyState struct {
	// The time when this key was created in GitLab.
	CreatedAt pulumi.StringPtrInput
	// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
	ExpiresAt pulumi.StringPtrInput
	// The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
	Key pulumi.StringPtrInput
	// The ID of the ssh key.
	KeyId pulumi.IntPtrInput
	// The title of the ssh key.
	Title pulumi.StringPtrInput
	// The ID or username of the user. If this field is omitted, this resource manages a SSH key for the current user. Otherwise, this resource manages a SSH key for the specified user, and an admin token is required.
	UserId pulumi.IntPtrInput
}

func (UserSshKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*userSshKeyState)(nil)).Elem()
}

type userSshKeyArgs struct {
	// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
	ExpiresAt *string `pulumi:"expiresAt"`
	// The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
	Key string `pulumi:"key"`
	// The title of the ssh key.
	Title string `pulumi:"title"`
	// The ID or username of the user. If this field is omitted, this resource manages a SSH key for the current user. Otherwise, this resource manages a SSH key for the specified user, and an admin token is required.
	UserId *int `pulumi:"userId"`
}

// The set of arguments for constructing a UserSshKey resource.
type UserSshKeyArgs struct {
	// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
	ExpiresAt pulumi.StringPtrInput
	// The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
	Key pulumi.StringInput
	// The title of the ssh key.
	Title pulumi.StringInput
	// The ID or username of the user. If this field is omitted, this resource manages a SSH key for the current user. Otherwise, this resource manages a SSH key for the specified user, and an admin token is required.
	UserId pulumi.IntPtrInput
}

func (UserSshKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userSshKeyArgs)(nil)).Elem()
}

type UserSshKeyInput interface {
	pulumi.Input

	ToUserSshKeyOutput() UserSshKeyOutput
	ToUserSshKeyOutputWithContext(ctx context.Context) UserSshKeyOutput
}

func (*UserSshKey) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSshKey)(nil)).Elem()
}

func (i *UserSshKey) ToUserSshKeyOutput() UserSshKeyOutput {
	return i.ToUserSshKeyOutputWithContext(context.Background())
}

func (i *UserSshKey) ToUserSshKeyOutputWithContext(ctx context.Context) UserSshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSshKeyOutput)
}

// UserSshKeyArrayInput is an input type that accepts UserSshKeyArray and UserSshKeyArrayOutput values.
// You can construct a concrete instance of `UserSshKeyArrayInput` via:
//
//	UserSshKeyArray{ UserSshKeyArgs{...} }
type UserSshKeyArrayInput interface {
	pulumi.Input

	ToUserSshKeyArrayOutput() UserSshKeyArrayOutput
	ToUserSshKeyArrayOutputWithContext(context.Context) UserSshKeyArrayOutput
}

type UserSshKeyArray []UserSshKeyInput

func (UserSshKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserSshKey)(nil)).Elem()
}

func (i UserSshKeyArray) ToUserSshKeyArrayOutput() UserSshKeyArrayOutput {
	return i.ToUserSshKeyArrayOutputWithContext(context.Background())
}

func (i UserSshKeyArray) ToUserSshKeyArrayOutputWithContext(ctx context.Context) UserSshKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSshKeyArrayOutput)
}

// UserSshKeyMapInput is an input type that accepts UserSshKeyMap and UserSshKeyMapOutput values.
// You can construct a concrete instance of `UserSshKeyMapInput` via:
//
//	UserSshKeyMap{ "key": UserSshKeyArgs{...} }
type UserSshKeyMapInput interface {
	pulumi.Input

	ToUserSshKeyMapOutput() UserSshKeyMapOutput
	ToUserSshKeyMapOutputWithContext(context.Context) UserSshKeyMapOutput
}

type UserSshKeyMap map[string]UserSshKeyInput

func (UserSshKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserSshKey)(nil)).Elem()
}

func (i UserSshKeyMap) ToUserSshKeyMapOutput() UserSshKeyMapOutput {
	return i.ToUserSshKeyMapOutputWithContext(context.Background())
}

func (i UserSshKeyMap) ToUserSshKeyMapOutputWithContext(ctx context.Context) UserSshKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSshKeyMapOutput)
}

type UserSshKeyOutput struct{ *pulumi.OutputState }

func (UserSshKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSshKey)(nil)).Elem()
}

func (o UserSshKeyOutput) ToUserSshKeyOutput() UserSshKeyOutput {
	return o
}

func (o UserSshKeyOutput) ToUserSshKeyOutputWithContext(ctx context.Context) UserSshKeyOutput {
	return o
}

// The time when this key was created in GitLab.
func (o UserSshKeyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *UserSshKey) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
func (o UserSshKeyOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSshKey) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
func (o UserSshKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *UserSshKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The ID of the ssh key.
func (o UserSshKeyOutput) KeyId() pulumi.IntOutput {
	return o.ApplyT(func(v *UserSshKey) pulumi.IntOutput { return v.KeyId }).(pulumi.IntOutput)
}

// The title of the ssh key.
func (o UserSshKeyOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *UserSshKey) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The ID or username of the user. If this field is omitted, this resource manages a SSH key for the current user. Otherwise, this resource manages a SSH key for the specified user, and an admin token is required.
func (o UserSshKeyOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *UserSshKey) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type UserSshKeyArrayOutput struct{ *pulumi.OutputState }

func (UserSshKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserSshKey)(nil)).Elem()
}

func (o UserSshKeyArrayOutput) ToUserSshKeyArrayOutput() UserSshKeyArrayOutput {
	return o
}

func (o UserSshKeyArrayOutput) ToUserSshKeyArrayOutputWithContext(ctx context.Context) UserSshKeyArrayOutput {
	return o
}

func (o UserSshKeyArrayOutput) Index(i pulumi.IntInput) UserSshKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserSshKey {
		return vs[0].([]*UserSshKey)[vs[1].(int)]
	}).(UserSshKeyOutput)
}

type UserSshKeyMapOutput struct{ *pulumi.OutputState }

func (UserSshKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserSshKey)(nil)).Elem()
}

func (o UserSshKeyMapOutput) ToUserSshKeyMapOutput() UserSshKeyMapOutput {
	return o
}

func (o UserSshKeyMapOutput) ToUserSshKeyMapOutputWithContext(ctx context.Context) UserSshKeyMapOutput {
	return o
}

func (o UserSshKeyMapOutput) MapIndex(k pulumi.StringInput) UserSshKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserSshKey {
		return vs[0].(map[string]*UserSshKey)[vs[1].(string)]
	}).(UserSshKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserSshKeyInput)(nil)).Elem(), &UserSshKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserSshKeyArrayInput)(nil)).Elem(), UserSshKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserSshKeyMapInput)(nil)).Elem(), UserSshKeyMap{})
	pulumi.RegisterOutputType(UserSshKeyOutput{})
	pulumi.RegisterOutputType(UserSshKeyArrayOutput{})
	pulumi.RegisterOutputType(UserSshKeyMapOutput{})
}
