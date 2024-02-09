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

// The `gitlabGroupAccess`token resource allows to manage the lifecycle of a group access token.
//
// > Group Access Token were introduced in GitLab 14.7
//
// **Upstream API**: [GitLab REST API](https://docs.gitlab.com/ee/api/group_access_tokens.html)
//
// ## Example Usage
//
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
//			exampleGroupAccessToken, err := gitlab.NewGroupAccessToken(ctx, "exampleGroupAccessToken", &gitlab.GroupAccessTokenArgs{
//				Group:       pulumi.String("25"),
//				ExpiresAt:   pulumi.String("2020-03-14"),
//				AccessLevel: pulumi.String("developer"),
//				Scopes: pulumi.StringArray{
//					pulumi.String("api"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewGroupVariable(ctx, "exampleGroupVariable", &gitlab.GroupVariableArgs{
//				Group: pulumi.String("25"),
//				Key:   pulumi.String("gat"),
//				Value: exampleGroupAccessToken.Token,
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
// A GitLab Group Access Token can be imported using a key composed of `<group-id>:<token-id>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/groupAccessToken:GroupAccessToken example "12345:1"
// ```
//
//	ATTENTION: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
type GroupAccessToken struct {
	pulumi.CustomResourceState

	// The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel pulumi.StringPtrOutput `pulumi:"accessLevel"`
	// True if the token is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// Time the token has been created, RFC3339 format.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
	ExpiresAt pulumi.StringOutput `pulumi:"expiresAt"`
	// The ID or path of the group to add the group access token to.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name of the group access token.
	Name pulumi.StringOutput `pulumi:"name"`
	// True if the token is revoked.
	Revoked pulumi.BoolOutput `pulumi:"revoked"`
	// The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// The group access token. This is only populated when creating a new group access token. This attribute is not available for imported resources.
	Token pulumi.StringOutput `pulumi:"token"`
	// The user id associated to the token.
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewGroupAccessToken registers a new resource with the given unique name, arguments, and options.
func NewGroupAccessToken(ctx *pulumi.Context,
	name string, args *GroupAccessTokenArgs, opts ...pulumi.ResourceOption) (*GroupAccessToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExpiresAt == nil {
		return nil, errors.New("invalid value for required argument 'ExpiresAt'")
	}
	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupAccessToken
	err := ctx.RegisterResource("gitlab:index/groupAccessToken:GroupAccessToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupAccessToken gets an existing GroupAccessToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupAccessToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupAccessTokenState, opts ...pulumi.ResourceOption) (*GroupAccessToken, error) {
	var resource GroupAccessToken
	err := ctx.ReadResource("gitlab:index/groupAccessToken:GroupAccessToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupAccessToken resources.
type groupAccessTokenState struct {
	// The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel *string `pulumi:"accessLevel"`
	// True if the token is active.
	Active *bool `pulumi:"active"`
	// Time the token has been created, RFC3339 format.
	CreatedAt *string `pulumi:"createdAt"`
	// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The ID or path of the group to add the group access token to.
	Group *string `pulumi:"group"`
	// The name of the group access token.
	Name *string `pulumi:"name"`
	// True if the token is revoked.
	Revoked *bool `pulumi:"revoked"`
	// The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
	Scopes []string `pulumi:"scopes"`
	// The group access token. This is only populated when creating a new group access token. This attribute is not available for imported resources.
	Token *string `pulumi:"token"`
	// The user id associated to the token.
	UserId *int `pulumi:"userId"`
}

type GroupAccessTokenState struct {
	// The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel pulumi.StringPtrInput
	// True if the token is active.
	Active pulumi.BoolPtrInput
	// Time the token has been created, RFC3339 format.
	CreatedAt pulumi.StringPtrInput
	// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
	ExpiresAt pulumi.StringPtrInput
	// The ID or path of the group to add the group access token to.
	Group pulumi.StringPtrInput
	// The name of the group access token.
	Name pulumi.StringPtrInput
	// True if the token is revoked.
	Revoked pulumi.BoolPtrInput
	// The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
	Scopes pulumi.StringArrayInput
	// The group access token. This is only populated when creating a new group access token. This attribute is not available for imported resources.
	Token pulumi.StringPtrInput
	// The user id associated to the token.
	UserId pulumi.IntPtrInput
}

func (GroupAccessTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupAccessTokenState)(nil)).Elem()
}

type groupAccessTokenArgs struct {
	// The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel *string `pulumi:"accessLevel"`
	// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
	ExpiresAt string `pulumi:"expiresAt"`
	// The ID or path of the group to add the group access token to.
	Group string `pulumi:"group"`
	// The name of the group access token.
	Name *string `pulumi:"name"`
	// The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a GroupAccessToken resource.
type GroupAccessTokenArgs struct {
	// The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel pulumi.StringPtrInput
	// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
	ExpiresAt pulumi.StringInput
	// The ID or path of the group to add the group access token to.
	Group pulumi.StringInput
	// The name of the group access token.
	Name pulumi.StringPtrInput
	// The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
	Scopes pulumi.StringArrayInput
}

func (GroupAccessTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupAccessTokenArgs)(nil)).Elem()
}

type GroupAccessTokenInput interface {
	pulumi.Input

	ToGroupAccessTokenOutput() GroupAccessTokenOutput
	ToGroupAccessTokenOutputWithContext(ctx context.Context) GroupAccessTokenOutput
}

func (*GroupAccessToken) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupAccessToken)(nil)).Elem()
}

func (i *GroupAccessToken) ToGroupAccessTokenOutput() GroupAccessTokenOutput {
	return i.ToGroupAccessTokenOutputWithContext(context.Background())
}

func (i *GroupAccessToken) ToGroupAccessTokenOutputWithContext(ctx context.Context) GroupAccessTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupAccessTokenOutput)
}

// GroupAccessTokenArrayInput is an input type that accepts GroupAccessTokenArray and GroupAccessTokenArrayOutput values.
// You can construct a concrete instance of `GroupAccessTokenArrayInput` via:
//
//	GroupAccessTokenArray{ GroupAccessTokenArgs{...} }
type GroupAccessTokenArrayInput interface {
	pulumi.Input

	ToGroupAccessTokenArrayOutput() GroupAccessTokenArrayOutput
	ToGroupAccessTokenArrayOutputWithContext(context.Context) GroupAccessTokenArrayOutput
}

type GroupAccessTokenArray []GroupAccessTokenInput

func (GroupAccessTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupAccessToken)(nil)).Elem()
}

func (i GroupAccessTokenArray) ToGroupAccessTokenArrayOutput() GroupAccessTokenArrayOutput {
	return i.ToGroupAccessTokenArrayOutputWithContext(context.Background())
}

func (i GroupAccessTokenArray) ToGroupAccessTokenArrayOutputWithContext(ctx context.Context) GroupAccessTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupAccessTokenArrayOutput)
}

// GroupAccessTokenMapInput is an input type that accepts GroupAccessTokenMap and GroupAccessTokenMapOutput values.
// You can construct a concrete instance of `GroupAccessTokenMapInput` via:
//
//	GroupAccessTokenMap{ "key": GroupAccessTokenArgs{...} }
type GroupAccessTokenMapInput interface {
	pulumi.Input

	ToGroupAccessTokenMapOutput() GroupAccessTokenMapOutput
	ToGroupAccessTokenMapOutputWithContext(context.Context) GroupAccessTokenMapOutput
}

type GroupAccessTokenMap map[string]GroupAccessTokenInput

func (GroupAccessTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupAccessToken)(nil)).Elem()
}

func (i GroupAccessTokenMap) ToGroupAccessTokenMapOutput() GroupAccessTokenMapOutput {
	return i.ToGroupAccessTokenMapOutputWithContext(context.Background())
}

func (i GroupAccessTokenMap) ToGroupAccessTokenMapOutputWithContext(ctx context.Context) GroupAccessTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupAccessTokenMapOutput)
}

type GroupAccessTokenOutput struct{ *pulumi.OutputState }

func (GroupAccessTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupAccessToken)(nil)).Elem()
}

func (o GroupAccessTokenOutput) ToGroupAccessTokenOutput() GroupAccessTokenOutput {
	return o
}

func (o GroupAccessTokenOutput) ToGroupAccessTokenOutputWithContext(ctx context.Context) GroupAccessTokenOutput {
	return o
}

// The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
func (o GroupAccessTokenOutput) AccessLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupAccessToken) pulumi.StringPtrOutput { return v.AccessLevel }).(pulumi.StringPtrOutput)
}

// True if the token is active.
func (o GroupAccessTokenOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *GroupAccessToken) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// Time the token has been created, RFC3339 format.
func (o GroupAccessTokenOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupAccessToken) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
func (o GroupAccessTokenOutput) ExpiresAt() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupAccessToken) pulumi.StringOutput { return v.ExpiresAt }).(pulumi.StringOutput)
}

// The ID or path of the group to add the group access token to.
func (o GroupAccessTokenOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupAccessToken) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The name of the group access token.
func (o GroupAccessTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupAccessToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// True if the token is revoked.
func (o GroupAccessTokenOutput) Revoked() pulumi.BoolOutput {
	return o.ApplyT(func(v *GroupAccessToken) pulumi.BoolOutput { return v.Revoked }).(pulumi.BoolOutput)
}

// The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
func (o GroupAccessTokenOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupAccessToken) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// The group access token. This is only populated when creating a new group access token. This attribute is not available for imported resources.
func (o GroupAccessTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupAccessToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The user id associated to the token.
func (o GroupAccessTokenOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupAccessToken) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type GroupAccessTokenArrayOutput struct{ *pulumi.OutputState }

func (GroupAccessTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupAccessToken)(nil)).Elem()
}

func (o GroupAccessTokenArrayOutput) ToGroupAccessTokenArrayOutput() GroupAccessTokenArrayOutput {
	return o
}

func (o GroupAccessTokenArrayOutput) ToGroupAccessTokenArrayOutputWithContext(ctx context.Context) GroupAccessTokenArrayOutput {
	return o
}

func (o GroupAccessTokenArrayOutput) Index(i pulumi.IntInput) GroupAccessTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupAccessToken {
		return vs[0].([]*GroupAccessToken)[vs[1].(int)]
	}).(GroupAccessTokenOutput)
}

type GroupAccessTokenMapOutput struct{ *pulumi.OutputState }

func (GroupAccessTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupAccessToken)(nil)).Elem()
}

func (o GroupAccessTokenMapOutput) ToGroupAccessTokenMapOutput() GroupAccessTokenMapOutput {
	return o
}

func (o GroupAccessTokenMapOutput) ToGroupAccessTokenMapOutputWithContext(ctx context.Context) GroupAccessTokenMapOutput {
	return o
}

func (o GroupAccessTokenMapOutput) MapIndex(k pulumi.StringInput) GroupAccessTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupAccessToken {
		return vs[0].(map[string]*GroupAccessToken)[vs[1].(string)]
	}).(GroupAccessTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupAccessTokenInput)(nil)).Elem(), &GroupAccessToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupAccessTokenArrayInput)(nil)).Elem(), GroupAccessTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupAccessTokenMapInput)(nil)).Elem(), GroupAccessTokenMap{})
	pulumi.RegisterOutputType(GroupAccessTokenOutput{})
	pulumi.RegisterOutputType(GroupAccessTokenArrayOutput{})
	pulumi.RegisterOutputType(GroupAccessTokenMapOutput{})
}
