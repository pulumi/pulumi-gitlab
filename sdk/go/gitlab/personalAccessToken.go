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

// The `PersonalAccessToken` resource allows to manage the lifecycle of a personal access token.
//
// > This resource requires administration privileges.
//
// > Use of the `timestamp()` function with expiresAt will cause the resource to be re-created with every apply, it's recommended to use `plantimestamp()` or a static value instead.
//
// > Observability scopes are in beta and may not work on all instances. See more details in [the documentation](https://docs.gitlab.com/ee/operations/tracing.html)
//
// > Due to [Automatic reuse detection](https://docs.gitlab.com/ee/api/personal_access_tokens.html#automatic-reuse-detection) it's possible that a new Personal Access Token will immediately be revoked. Check if an old process using the old token is running if this happens.
//
// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/personal_access_tokens.html)
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
//			example, err := gitlab.NewPersonalAccessToken(ctx, "example", &gitlab.PersonalAccessTokenArgs{
//				UserId:    pulumi.Int(25),
//				Name:      pulumi.String("Example personal access token"),
//				ExpiresAt: pulumi.String("2020-03-14"),
//				Scopes: pulumi.StringArray{
//					pulumi.String("api"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewProjectVariable(ctx, "example", &gitlab.ProjectVariableArgs{
//				Project: pulumi.Any(exampleGitlabProject.Id),
//				Key:     pulumi.String("pat"),
//				Value:   example.Token,
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
// A GitLab Personal Access Token can be imported using a key composed of `<user-id>:<token-id>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/personalAccessToken:PersonalAccessToken example "12345:1"
// ```
//
// NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
type PersonalAccessToken struct {
	pulumi.CustomResourceState

	// True if the token is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// Time the token has been created, RFC3339 format.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// When the token will expire, YYYY-MM-DD format.
	ExpiresAt pulumi.StringOutput `pulumi:"expiresAt"`
	// The name of the personal access token.
	Name pulumi.StringOutput `pulumi:"name"`
	// True if the token is revoked.
	Revoked pulumi.BoolOutput `pulumi:"revoked"`
	// The scopes of the personal access token. valid values are: `api`, `readUser`, `readApi`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readServicePing`
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// The token of the personal access token. **Note**: the token is not available for imported resources.
	Token pulumi.StringOutput `pulumi:"token"`
	// The ID of the user.
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewPersonalAccessToken registers a new resource with the given unique name, arguments, and options.
func NewPersonalAccessToken(ctx *pulumi.Context,
	name string, args *PersonalAccessTokenArgs, opts ...pulumi.ResourceOption) (*PersonalAccessToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource PersonalAccessToken
	err := ctx.RegisterResource("gitlab:index/personalAccessToken:PersonalAccessToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersonalAccessToken gets an existing PersonalAccessToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersonalAccessToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersonalAccessTokenState, opts ...pulumi.ResourceOption) (*PersonalAccessToken, error) {
	var resource PersonalAccessToken
	err := ctx.ReadResource("gitlab:index/personalAccessToken:PersonalAccessToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersonalAccessToken resources.
type personalAccessTokenState struct {
	// True if the token is active.
	Active *bool `pulumi:"active"`
	// Time the token has been created, RFC3339 format.
	CreatedAt *string `pulumi:"createdAt"`
	// When the token will expire, YYYY-MM-DD format.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The name of the personal access token.
	Name *string `pulumi:"name"`
	// True if the token is revoked.
	Revoked *bool `pulumi:"revoked"`
	// The scopes of the personal access token. valid values are: `api`, `readUser`, `readApi`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readServicePing`
	Scopes []string `pulumi:"scopes"`
	// The token of the personal access token. **Note**: the token is not available for imported resources.
	Token *string `pulumi:"token"`
	// The ID of the user.
	UserId *int `pulumi:"userId"`
}

type PersonalAccessTokenState struct {
	// True if the token is active.
	Active pulumi.BoolPtrInput
	// Time the token has been created, RFC3339 format.
	CreatedAt pulumi.StringPtrInput
	// When the token will expire, YYYY-MM-DD format.
	ExpiresAt pulumi.StringPtrInput
	// The name of the personal access token.
	Name pulumi.StringPtrInput
	// True if the token is revoked.
	Revoked pulumi.BoolPtrInput
	// The scopes of the personal access token. valid values are: `api`, `readUser`, `readApi`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readServicePing`
	Scopes pulumi.StringArrayInput
	// The token of the personal access token. **Note**: the token is not available for imported resources.
	Token pulumi.StringPtrInput
	// The ID of the user.
	UserId pulumi.IntPtrInput
}

func (PersonalAccessTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*personalAccessTokenState)(nil)).Elem()
}

type personalAccessTokenArgs struct {
	// When the token will expire, YYYY-MM-DD format.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The name of the personal access token.
	Name *string `pulumi:"name"`
	// The scopes of the personal access token. valid values are: `api`, `readUser`, `readApi`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readServicePing`
	Scopes []string `pulumi:"scopes"`
	// The ID of the user.
	UserId int `pulumi:"userId"`
}

// The set of arguments for constructing a PersonalAccessToken resource.
type PersonalAccessTokenArgs struct {
	// When the token will expire, YYYY-MM-DD format.
	ExpiresAt pulumi.StringPtrInput
	// The name of the personal access token.
	Name pulumi.StringPtrInput
	// The scopes of the personal access token. valid values are: `api`, `readUser`, `readApi`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readServicePing`
	Scopes pulumi.StringArrayInput
	// The ID of the user.
	UserId pulumi.IntInput
}

func (PersonalAccessTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*personalAccessTokenArgs)(nil)).Elem()
}

type PersonalAccessTokenInput interface {
	pulumi.Input

	ToPersonalAccessTokenOutput() PersonalAccessTokenOutput
	ToPersonalAccessTokenOutputWithContext(ctx context.Context) PersonalAccessTokenOutput
}

func (*PersonalAccessToken) ElementType() reflect.Type {
	return reflect.TypeOf((**PersonalAccessToken)(nil)).Elem()
}

func (i *PersonalAccessToken) ToPersonalAccessTokenOutput() PersonalAccessTokenOutput {
	return i.ToPersonalAccessTokenOutputWithContext(context.Background())
}

func (i *PersonalAccessToken) ToPersonalAccessTokenOutputWithContext(ctx context.Context) PersonalAccessTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalAccessTokenOutput)
}

// PersonalAccessTokenArrayInput is an input type that accepts PersonalAccessTokenArray and PersonalAccessTokenArrayOutput values.
// You can construct a concrete instance of `PersonalAccessTokenArrayInput` via:
//
//	PersonalAccessTokenArray{ PersonalAccessTokenArgs{...} }
type PersonalAccessTokenArrayInput interface {
	pulumi.Input

	ToPersonalAccessTokenArrayOutput() PersonalAccessTokenArrayOutput
	ToPersonalAccessTokenArrayOutputWithContext(context.Context) PersonalAccessTokenArrayOutput
}

type PersonalAccessTokenArray []PersonalAccessTokenInput

func (PersonalAccessTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersonalAccessToken)(nil)).Elem()
}

func (i PersonalAccessTokenArray) ToPersonalAccessTokenArrayOutput() PersonalAccessTokenArrayOutput {
	return i.ToPersonalAccessTokenArrayOutputWithContext(context.Background())
}

func (i PersonalAccessTokenArray) ToPersonalAccessTokenArrayOutputWithContext(ctx context.Context) PersonalAccessTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalAccessTokenArrayOutput)
}

// PersonalAccessTokenMapInput is an input type that accepts PersonalAccessTokenMap and PersonalAccessTokenMapOutput values.
// You can construct a concrete instance of `PersonalAccessTokenMapInput` via:
//
//	PersonalAccessTokenMap{ "key": PersonalAccessTokenArgs{...} }
type PersonalAccessTokenMapInput interface {
	pulumi.Input

	ToPersonalAccessTokenMapOutput() PersonalAccessTokenMapOutput
	ToPersonalAccessTokenMapOutputWithContext(context.Context) PersonalAccessTokenMapOutput
}

type PersonalAccessTokenMap map[string]PersonalAccessTokenInput

func (PersonalAccessTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersonalAccessToken)(nil)).Elem()
}

func (i PersonalAccessTokenMap) ToPersonalAccessTokenMapOutput() PersonalAccessTokenMapOutput {
	return i.ToPersonalAccessTokenMapOutputWithContext(context.Background())
}

func (i PersonalAccessTokenMap) ToPersonalAccessTokenMapOutputWithContext(ctx context.Context) PersonalAccessTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalAccessTokenMapOutput)
}

type PersonalAccessTokenOutput struct{ *pulumi.OutputState }

func (PersonalAccessTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersonalAccessToken)(nil)).Elem()
}

func (o PersonalAccessTokenOutput) ToPersonalAccessTokenOutput() PersonalAccessTokenOutput {
	return o
}

func (o PersonalAccessTokenOutput) ToPersonalAccessTokenOutputWithContext(ctx context.Context) PersonalAccessTokenOutput {
	return o
}

// True if the token is active.
func (o PersonalAccessTokenOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// Time the token has been created, RFC3339 format.
func (o PersonalAccessTokenOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// When the token will expire, YYYY-MM-DD format.
func (o PersonalAccessTokenOutput) ExpiresAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.StringOutput { return v.ExpiresAt }).(pulumi.StringOutput)
}

// The name of the personal access token.
func (o PersonalAccessTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// True if the token is revoked.
func (o PersonalAccessTokenOutput) Revoked() pulumi.BoolOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.BoolOutput { return v.Revoked }).(pulumi.BoolOutput)
}

// The scopes of the personal access token. valid values are: `api`, `readUser`, `readApi`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readServicePing`
func (o PersonalAccessTokenOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// The token of the personal access token. **Note**: the token is not available for imported resources.
func (o PersonalAccessTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The ID of the user.
func (o PersonalAccessTokenOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *PersonalAccessToken) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type PersonalAccessTokenArrayOutput struct{ *pulumi.OutputState }

func (PersonalAccessTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersonalAccessToken)(nil)).Elem()
}

func (o PersonalAccessTokenArrayOutput) ToPersonalAccessTokenArrayOutput() PersonalAccessTokenArrayOutput {
	return o
}

func (o PersonalAccessTokenArrayOutput) ToPersonalAccessTokenArrayOutputWithContext(ctx context.Context) PersonalAccessTokenArrayOutput {
	return o
}

func (o PersonalAccessTokenArrayOutput) Index(i pulumi.IntInput) PersonalAccessTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PersonalAccessToken {
		return vs[0].([]*PersonalAccessToken)[vs[1].(int)]
	}).(PersonalAccessTokenOutput)
}

type PersonalAccessTokenMapOutput struct{ *pulumi.OutputState }

func (PersonalAccessTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersonalAccessToken)(nil)).Elem()
}

func (o PersonalAccessTokenMapOutput) ToPersonalAccessTokenMapOutput() PersonalAccessTokenMapOutput {
	return o
}

func (o PersonalAccessTokenMapOutput) ToPersonalAccessTokenMapOutputWithContext(ctx context.Context) PersonalAccessTokenMapOutput {
	return o
}

func (o PersonalAccessTokenMapOutput) MapIndex(k pulumi.StringInput) PersonalAccessTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PersonalAccessToken {
		return vs[0].(map[string]*PersonalAccessToken)[vs[1].(string)]
	}).(PersonalAccessTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PersonalAccessTokenInput)(nil)).Elem(), &PersonalAccessToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersonalAccessTokenArrayInput)(nil)).Elem(), PersonalAccessTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersonalAccessTokenMapInput)(nil)).Elem(), PersonalAccessTokenMap{})
	pulumi.RegisterOutputType(PersonalAccessTokenOutput{})
	pulumi.RegisterOutputType(PersonalAccessTokenArrayOutput{})
	pulumi.RegisterOutputType(PersonalAccessTokenMapOutput{})
}
