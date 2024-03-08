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

// The `DeployToken` resource allows to manage the lifecycle of group and project deploy tokens.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_tokens.html)
//
// ## Import
//
// GitLab deploy tokens can be imported using an id made up of `{type}:{type_id}:{deploy_token_id}`, where type is one of: project, group.
//
// ```sh
// $ pulumi import gitlab:index/deployToken:DeployToken group_token group:1:3
// ```
//
// ```sh
// $ pulumi import gitlab:index/deployToken:DeployToken project_token project:1:4
// ```
//
// Note: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
type DeployToken struct {
	pulumi.CustomResourceState

	// The id of the deploy token.
	DeployTokenId pulumi.IntOutput `pulumi:"deployTokenId"`
	// Time the token will expire it, RFC3339 format. Will not expire per default.
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The name or id of the group to add the deploy token to.
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// A name to describe the deploy token with.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name or id of the project to add the deploy token to.
	Project pulumi.StringPtrOutput `pulumi:"project"`
	// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// The secret token. This is only populated when creating a new deploy token. **Note**: The token is not available for imported resources.
	Token pulumi.StringOutput `pulumi:"token"`
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewDeployToken registers a new resource with the given unique name, arguments, and options.
func NewDeployToken(ctx *pulumi.Context,
	name string, args *DeployTokenArgs, opts ...pulumi.ResourceOption) (*DeployToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DeployToken
	err := ctx.RegisterResource("gitlab:index/deployToken:DeployToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployToken gets an existing DeployToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployTokenState, opts ...pulumi.ResourceOption) (*DeployToken, error) {
	var resource DeployToken
	err := ctx.ReadResource("gitlab:index/deployToken:DeployToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeployToken resources.
type deployTokenState struct {
	// The id of the deploy token.
	DeployTokenId *int `pulumi:"deployTokenId"`
	// Time the token will expire it, RFC3339 format. Will not expire per default.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The name or id of the group to add the deploy token to.
	Group *string `pulumi:"group"`
	// A name to describe the deploy token with.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the deploy token to.
	Project *string `pulumi:"project"`
	// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
	Scopes []string `pulumi:"scopes"`
	// The secret token. This is only populated when creating a new deploy token. **Note**: The token is not available for imported resources.
	Token *string `pulumi:"token"`
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username *string `pulumi:"username"`
}

type DeployTokenState struct {
	// The id of the deploy token.
	DeployTokenId pulumi.IntPtrInput
	// Time the token will expire it, RFC3339 format. Will not expire per default.
	ExpiresAt pulumi.StringPtrInput
	// The name or id of the group to add the deploy token to.
	Group pulumi.StringPtrInput
	// A name to describe the deploy token with.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the deploy token to.
	Project pulumi.StringPtrInput
	// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
	Scopes pulumi.StringArrayInput
	// The secret token. This is only populated when creating a new deploy token. **Note**: The token is not available for imported resources.
	Token pulumi.StringPtrInput
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username pulumi.StringPtrInput
}

func (DeployTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployTokenState)(nil)).Elem()
}

type deployTokenArgs struct {
	// Time the token will expire it, RFC3339 format. Will not expire per default.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The name or id of the group to add the deploy token to.
	Group *string `pulumi:"group"`
	// A name to describe the deploy token with.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the deploy token to.
	Project *string `pulumi:"project"`
	// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
	Scopes []string `pulumi:"scopes"`
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a DeployToken resource.
type DeployTokenArgs struct {
	// Time the token will expire it, RFC3339 format. Will not expire per default.
	ExpiresAt pulumi.StringPtrInput
	// The name or id of the group to add the deploy token to.
	Group pulumi.StringPtrInput
	// A name to describe the deploy token with.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the deploy token to.
	Project pulumi.StringPtrInput
	// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
	Scopes pulumi.StringArrayInput
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username pulumi.StringPtrInput
}

func (DeployTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployTokenArgs)(nil)).Elem()
}

type DeployTokenInput interface {
	pulumi.Input

	ToDeployTokenOutput() DeployTokenOutput
	ToDeployTokenOutputWithContext(ctx context.Context) DeployTokenOutput
}

func (*DeployToken) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployToken)(nil)).Elem()
}

func (i *DeployToken) ToDeployTokenOutput() DeployTokenOutput {
	return i.ToDeployTokenOutputWithContext(context.Background())
}

func (i *DeployToken) ToDeployTokenOutputWithContext(ctx context.Context) DeployTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployTokenOutput)
}

// DeployTokenArrayInput is an input type that accepts DeployTokenArray and DeployTokenArrayOutput values.
// You can construct a concrete instance of `DeployTokenArrayInput` via:
//
//	DeployTokenArray{ DeployTokenArgs{...} }
type DeployTokenArrayInput interface {
	pulumi.Input

	ToDeployTokenArrayOutput() DeployTokenArrayOutput
	ToDeployTokenArrayOutputWithContext(context.Context) DeployTokenArrayOutput
}

type DeployTokenArray []DeployTokenInput

func (DeployTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeployToken)(nil)).Elem()
}

func (i DeployTokenArray) ToDeployTokenArrayOutput() DeployTokenArrayOutput {
	return i.ToDeployTokenArrayOutputWithContext(context.Background())
}

func (i DeployTokenArray) ToDeployTokenArrayOutputWithContext(ctx context.Context) DeployTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployTokenArrayOutput)
}

// DeployTokenMapInput is an input type that accepts DeployTokenMap and DeployTokenMapOutput values.
// You can construct a concrete instance of `DeployTokenMapInput` via:
//
//	DeployTokenMap{ "key": DeployTokenArgs{...} }
type DeployTokenMapInput interface {
	pulumi.Input

	ToDeployTokenMapOutput() DeployTokenMapOutput
	ToDeployTokenMapOutputWithContext(context.Context) DeployTokenMapOutput
}

type DeployTokenMap map[string]DeployTokenInput

func (DeployTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeployToken)(nil)).Elem()
}

func (i DeployTokenMap) ToDeployTokenMapOutput() DeployTokenMapOutput {
	return i.ToDeployTokenMapOutputWithContext(context.Background())
}

func (i DeployTokenMap) ToDeployTokenMapOutputWithContext(ctx context.Context) DeployTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployTokenMapOutput)
}

type DeployTokenOutput struct{ *pulumi.OutputState }

func (DeployTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployToken)(nil)).Elem()
}

func (o DeployTokenOutput) ToDeployTokenOutput() DeployTokenOutput {
	return o
}

func (o DeployTokenOutput) ToDeployTokenOutputWithContext(ctx context.Context) DeployTokenOutput {
	return o
}

// The id of the deploy token.
func (o DeployTokenOutput) DeployTokenId() pulumi.IntOutput {
	return o.ApplyT(func(v *DeployToken) pulumi.IntOutput { return v.DeployTokenId }).(pulumi.IntOutput)
}

// Time the token will expire it, RFC3339 format. Will not expire per default.
func (o DeployTokenOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeployToken) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// The name or id of the group to add the deploy token to.
func (o DeployTokenOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeployToken) pulumi.StringPtrOutput { return v.Group }).(pulumi.StringPtrOutput)
}

// A name to describe the deploy token with.
func (o DeployTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeployToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name or id of the project to add the deploy token to.
func (o DeployTokenOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeployToken) pulumi.StringPtrOutput { return v.Project }).(pulumi.StringPtrOutput)
}

// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
func (o DeployTokenOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DeployToken) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// The secret token. This is only populated when creating a new deploy token. **Note**: The token is not available for imported resources.
func (o DeployTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *DeployToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
func (o DeployTokenOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *DeployToken) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type DeployTokenArrayOutput struct{ *pulumi.OutputState }

func (DeployTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeployToken)(nil)).Elem()
}

func (o DeployTokenArrayOutput) ToDeployTokenArrayOutput() DeployTokenArrayOutput {
	return o
}

func (o DeployTokenArrayOutput) ToDeployTokenArrayOutputWithContext(ctx context.Context) DeployTokenArrayOutput {
	return o
}

func (o DeployTokenArrayOutput) Index(i pulumi.IntInput) DeployTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeployToken {
		return vs[0].([]*DeployToken)[vs[1].(int)]
	}).(DeployTokenOutput)
}

type DeployTokenMapOutput struct{ *pulumi.OutputState }

func (DeployTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeployToken)(nil)).Elem()
}

func (o DeployTokenMapOutput) ToDeployTokenMapOutput() DeployTokenMapOutput {
	return o
}

func (o DeployTokenMapOutput) ToDeployTokenMapOutputWithContext(ctx context.Context) DeployTokenMapOutput {
	return o
}

func (o DeployTokenMapOutput) MapIndex(k pulumi.StringInput) DeployTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeployToken {
		return vs[0].(map[string]*DeployToken)[vs[1].(string)]
	}).(DeployTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeployTokenInput)(nil)).Elem(), &DeployToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeployTokenArrayInput)(nil)).Elem(), DeployTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeployTokenMapInput)(nil)).Elem(), DeployTokenMap{})
	pulumi.RegisterOutputType(DeployTokenOutput{})
	pulumi.RegisterOutputType(DeployTokenArrayOutput{})
	pulumi.RegisterOutputType(DeployTokenMapOutput{})
}
