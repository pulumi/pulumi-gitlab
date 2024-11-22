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

// The `GroupServiceAccountAccessToken` resource allows to manage the lifecycle of a group service account access token.
//
// > Use of the `timestamp()` function with expiresAt will cause the resource to be re-created with every apply, it's recommended to use `plantimestamp()` or a static value instead.
//
// > Reading the access token status of a service account requires an admin token or a top-level group owner token on gitlab.com. As a result, this resource will ignore permission errors when attempting to read the token status, and will rely on the values in state instead. This can lead to apply-time failures if the token configured for the provider doesn't have permissions to rotate tokens for the service account.
//
// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/group_service_accounts.html#create-a-personal-access-token-for-a-service-account-user)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// This must be a top-level group
//			example, err := gitlab.NewGroup(ctx, "example", &gitlab.GroupArgs{
//				Name:        pulumi.String("example"),
//				Path:        pulumi.String("example"),
//				Description: pulumi.String("An example group"),
//			})
//			if err != nil {
//				return err
//			}
//			// The service account against the top-level group
//			exampleSa, err := gitlab.NewGroupServiceAccount(ctx, "example_sa", &gitlab.GroupServiceAccountArgs{
//				Group:    example.ID(),
//				Name:     pulumi.String("example-name"),
//				Username: pulumi.String("example-username"),
//			})
//			if err != nil {
//				return err
//			}
//			// To assign the service account to a group
//			_, err = gitlab.NewGroupMembership(ctx, "example_membership", &gitlab.GroupMembershipArgs{
//				GroupId:     example.ID(),
//				UserId:      exampleSa.ServiceAccountId,
//				AccessLevel: pulumi.String("developer"),
//				ExpiresAt:   pulumi.String("2020-03-14"),
//			})
//			if err != nil {
//				return err
//			}
//			// The service account access token
//			_, err = gitlab.NewGroupServiceAccountAccessToken(ctx, "example_sa_token", &gitlab.GroupServiceAccountAccessTokenArgs{
//				Group:     example.ID(),
//				UserId:    exampleSa.ServiceAccountId,
//				Name:      pulumi.String("Example service account access token"),
//				ExpiresAt: pulumi.String("2020-03-14"),
//				Scopes: pulumi.StringArray{
//					pulumi.String("api"),
//				},
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
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_service_account_access_token`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_group_service_account_access_token.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// ```sh
// $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken You can import a service account access token using `<resource> <id>`. The
// ```
//
// `id` is in the form of <group_id>:<service_account_id>:<access_token_id>
//
// Importing an access token does not import the access token value.
//
// ```sh
// $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken example 1:2:3
// ```
type GroupServiceAccountAccessToken struct {
	pulumi.CustomResourceState

	// True if the token is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// Time the token has been created, RFC3339 format.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
	ExpiresAt pulumi.StringOutput `pulumi:"expiresAt"`
	// The ID or URL-encoded path of the group containing the service account. Must be a top level group.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name of the personal access token.
	Name pulumi.StringOutput `pulumi:"name"`
	// True if the token is revoked.
	Revoked pulumi.BoolOutput `pulumi:"revoked"`
	// The scopes of the group service account access token. valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readObservability`, `writeObservability`
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// The token of the group service account access token. **Note**: the token is not available for imported resources.
	Token pulumi.StringOutput `pulumi:"token"`
	// The ID of a service account user.
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewGroupServiceAccountAccessToken registers a new resource with the given unique name, arguments, and options.
func NewGroupServiceAccountAccessToken(ctx *pulumi.Context,
	name string, args *GroupServiceAccountAccessTokenArgs, opts ...pulumi.ResourceOption) (*GroupServiceAccountAccessToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupServiceAccountAccessToken
	err := ctx.RegisterResource("gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupServiceAccountAccessToken gets an existing GroupServiceAccountAccessToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupServiceAccountAccessToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupServiceAccountAccessTokenState, opts ...pulumi.ResourceOption) (*GroupServiceAccountAccessToken, error) {
	var resource GroupServiceAccountAccessToken
	err := ctx.ReadResource("gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupServiceAccountAccessToken resources.
type groupServiceAccountAccessTokenState struct {
	// True if the token is active.
	Active *bool `pulumi:"active"`
	// Time the token has been created, RFC3339 format.
	CreatedAt *string `pulumi:"createdAt"`
	// The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The ID or URL-encoded path of the group containing the service account. Must be a top level group.
	Group *string `pulumi:"group"`
	// The name of the personal access token.
	Name *string `pulumi:"name"`
	// True if the token is revoked.
	Revoked *bool `pulumi:"revoked"`
	// The scopes of the group service account access token. valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readObservability`, `writeObservability`
	Scopes []string `pulumi:"scopes"`
	// The token of the group service account access token. **Note**: the token is not available for imported resources.
	Token *string `pulumi:"token"`
	// The ID of a service account user.
	UserId *int `pulumi:"userId"`
}

type GroupServiceAccountAccessTokenState struct {
	// True if the token is active.
	Active pulumi.BoolPtrInput
	// Time the token has been created, RFC3339 format.
	CreatedAt pulumi.StringPtrInput
	// The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
	ExpiresAt pulumi.StringPtrInput
	// The ID or URL-encoded path of the group containing the service account. Must be a top level group.
	Group pulumi.StringPtrInput
	// The name of the personal access token.
	Name pulumi.StringPtrInput
	// True if the token is revoked.
	Revoked pulumi.BoolPtrInput
	// The scopes of the group service account access token. valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readObservability`, `writeObservability`
	Scopes pulumi.StringArrayInput
	// The token of the group service account access token. **Note**: the token is not available for imported resources.
	Token pulumi.StringPtrInput
	// The ID of a service account user.
	UserId pulumi.IntPtrInput
}

func (GroupServiceAccountAccessTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupServiceAccountAccessTokenState)(nil)).Elem()
}

type groupServiceAccountAccessTokenArgs struct {
	// The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The ID or URL-encoded path of the group containing the service account. Must be a top level group.
	Group string `pulumi:"group"`
	// The name of the personal access token.
	Name *string `pulumi:"name"`
	// The scopes of the group service account access token. valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readObservability`, `writeObservability`
	Scopes []string `pulumi:"scopes"`
	// The ID of a service account user.
	UserId int `pulumi:"userId"`
}

// The set of arguments for constructing a GroupServiceAccountAccessToken resource.
type GroupServiceAccountAccessTokenArgs struct {
	// The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
	ExpiresAt pulumi.StringPtrInput
	// The ID or URL-encoded path of the group containing the service account. Must be a top level group.
	Group pulumi.StringInput
	// The name of the personal access token.
	Name pulumi.StringPtrInput
	// The scopes of the group service account access token. valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readObservability`, `writeObservability`
	Scopes pulumi.StringArrayInput
	// The ID of a service account user.
	UserId pulumi.IntInput
}

func (GroupServiceAccountAccessTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupServiceAccountAccessTokenArgs)(nil)).Elem()
}

type GroupServiceAccountAccessTokenInput interface {
	pulumi.Input

	ToGroupServiceAccountAccessTokenOutput() GroupServiceAccountAccessTokenOutput
	ToGroupServiceAccountAccessTokenOutputWithContext(ctx context.Context) GroupServiceAccountAccessTokenOutput
}

func (*GroupServiceAccountAccessToken) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupServiceAccountAccessToken)(nil)).Elem()
}

func (i *GroupServiceAccountAccessToken) ToGroupServiceAccountAccessTokenOutput() GroupServiceAccountAccessTokenOutput {
	return i.ToGroupServiceAccountAccessTokenOutputWithContext(context.Background())
}

func (i *GroupServiceAccountAccessToken) ToGroupServiceAccountAccessTokenOutputWithContext(ctx context.Context) GroupServiceAccountAccessTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupServiceAccountAccessTokenOutput)
}

// GroupServiceAccountAccessTokenArrayInput is an input type that accepts GroupServiceAccountAccessTokenArray and GroupServiceAccountAccessTokenArrayOutput values.
// You can construct a concrete instance of `GroupServiceAccountAccessTokenArrayInput` via:
//
//	GroupServiceAccountAccessTokenArray{ GroupServiceAccountAccessTokenArgs{...} }
type GroupServiceAccountAccessTokenArrayInput interface {
	pulumi.Input

	ToGroupServiceAccountAccessTokenArrayOutput() GroupServiceAccountAccessTokenArrayOutput
	ToGroupServiceAccountAccessTokenArrayOutputWithContext(context.Context) GroupServiceAccountAccessTokenArrayOutput
}

type GroupServiceAccountAccessTokenArray []GroupServiceAccountAccessTokenInput

func (GroupServiceAccountAccessTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupServiceAccountAccessToken)(nil)).Elem()
}

func (i GroupServiceAccountAccessTokenArray) ToGroupServiceAccountAccessTokenArrayOutput() GroupServiceAccountAccessTokenArrayOutput {
	return i.ToGroupServiceAccountAccessTokenArrayOutputWithContext(context.Background())
}

func (i GroupServiceAccountAccessTokenArray) ToGroupServiceAccountAccessTokenArrayOutputWithContext(ctx context.Context) GroupServiceAccountAccessTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupServiceAccountAccessTokenArrayOutput)
}

// GroupServiceAccountAccessTokenMapInput is an input type that accepts GroupServiceAccountAccessTokenMap and GroupServiceAccountAccessTokenMapOutput values.
// You can construct a concrete instance of `GroupServiceAccountAccessTokenMapInput` via:
//
//	GroupServiceAccountAccessTokenMap{ "key": GroupServiceAccountAccessTokenArgs{...} }
type GroupServiceAccountAccessTokenMapInput interface {
	pulumi.Input

	ToGroupServiceAccountAccessTokenMapOutput() GroupServiceAccountAccessTokenMapOutput
	ToGroupServiceAccountAccessTokenMapOutputWithContext(context.Context) GroupServiceAccountAccessTokenMapOutput
}

type GroupServiceAccountAccessTokenMap map[string]GroupServiceAccountAccessTokenInput

func (GroupServiceAccountAccessTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupServiceAccountAccessToken)(nil)).Elem()
}

func (i GroupServiceAccountAccessTokenMap) ToGroupServiceAccountAccessTokenMapOutput() GroupServiceAccountAccessTokenMapOutput {
	return i.ToGroupServiceAccountAccessTokenMapOutputWithContext(context.Background())
}

func (i GroupServiceAccountAccessTokenMap) ToGroupServiceAccountAccessTokenMapOutputWithContext(ctx context.Context) GroupServiceAccountAccessTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupServiceAccountAccessTokenMapOutput)
}

type GroupServiceAccountAccessTokenOutput struct{ *pulumi.OutputState }

func (GroupServiceAccountAccessTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupServiceAccountAccessToken)(nil)).Elem()
}

func (o GroupServiceAccountAccessTokenOutput) ToGroupServiceAccountAccessTokenOutput() GroupServiceAccountAccessTokenOutput {
	return o
}

func (o GroupServiceAccountAccessTokenOutput) ToGroupServiceAccountAccessTokenOutputWithContext(ctx context.Context) GroupServiceAccountAccessTokenOutput {
	return o
}

// True if the token is active.
func (o GroupServiceAccountAccessTokenOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *GroupServiceAccountAccessToken) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// Time the token has been created, RFC3339 format.
func (o GroupServiceAccountAccessTokenOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupServiceAccountAccessToken) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
func (o GroupServiceAccountAccessTokenOutput) ExpiresAt() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupServiceAccountAccessToken) pulumi.StringOutput { return v.ExpiresAt }).(pulumi.StringOutput)
}

// The ID or URL-encoded path of the group containing the service account. Must be a top level group.
func (o GroupServiceAccountAccessTokenOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupServiceAccountAccessToken) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The name of the personal access token.
func (o GroupServiceAccountAccessTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupServiceAccountAccessToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// True if the token is revoked.
func (o GroupServiceAccountAccessTokenOutput) Revoked() pulumi.BoolOutput {
	return o.ApplyT(func(v *GroupServiceAccountAccessToken) pulumi.BoolOutput { return v.Revoked }).(pulumi.BoolOutput)
}

// The scopes of the group service account access token. valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readObservability`, `writeObservability`
func (o GroupServiceAccountAccessTokenOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupServiceAccountAccessToken) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// The token of the group service account access token. **Note**: the token is not available for imported resources.
func (o GroupServiceAccountAccessTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupServiceAccountAccessToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The ID of a service account user.
func (o GroupServiceAccountAccessTokenOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupServiceAccountAccessToken) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type GroupServiceAccountAccessTokenArrayOutput struct{ *pulumi.OutputState }

func (GroupServiceAccountAccessTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupServiceAccountAccessToken)(nil)).Elem()
}

func (o GroupServiceAccountAccessTokenArrayOutput) ToGroupServiceAccountAccessTokenArrayOutput() GroupServiceAccountAccessTokenArrayOutput {
	return o
}

func (o GroupServiceAccountAccessTokenArrayOutput) ToGroupServiceAccountAccessTokenArrayOutputWithContext(ctx context.Context) GroupServiceAccountAccessTokenArrayOutput {
	return o
}

func (o GroupServiceAccountAccessTokenArrayOutput) Index(i pulumi.IntInput) GroupServiceAccountAccessTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupServiceAccountAccessToken {
		return vs[0].([]*GroupServiceAccountAccessToken)[vs[1].(int)]
	}).(GroupServiceAccountAccessTokenOutput)
}

type GroupServiceAccountAccessTokenMapOutput struct{ *pulumi.OutputState }

func (GroupServiceAccountAccessTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupServiceAccountAccessToken)(nil)).Elem()
}

func (o GroupServiceAccountAccessTokenMapOutput) ToGroupServiceAccountAccessTokenMapOutput() GroupServiceAccountAccessTokenMapOutput {
	return o
}

func (o GroupServiceAccountAccessTokenMapOutput) ToGroupServiceAccountAccessTokenMapOutputWithContext(ctx context.Context) GroupServiceAccountAccessTokenMapOutput {
	return o
}

func (o GroupServiceAccountAccessTokenMapOutput) MapIndex(k pulumi.StringInput) GroupServiceAccountAccessTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupServiceAccountAccessToken {
		return vs[0].(map[string]*GroupServiceAccountAccessToken)[vs[1].(string)]
	}).(GroupServiceAccountAccessTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupServiceAccountAccessTokenInput)(nil)).Elem(), &GroupServiceAccountAccessToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupServiceAccountAccessTokenArrayInput)(nil)).Elem(), GroupServiceAccountAccessTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupServiceAccountAccessTokenMapInput)(nil)).Elem(), GroupServiceAccountAccessTokenMap{})
	pulumi.RegisterOutputType(GroupServiceAccountAccessTokenOutput{})
	pulumi.RegisterOutputType(GroupServiceAccountAccessTokenArrayOutput{})
	pulumi.RegisterOutputType(GroupServiceAccountAccessTokenMapOutput{})
}
