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

// The `ProjectAccessToken` resource allows to manage the lifecycle of a project access token.
//
// >  Use of the `timestamp()` function with expiresAt will cause the resource to be re-created with every apply, it's recommended to use `plantimestamp()` or a static value instead.
//
// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/project_access_tokens.html)
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
//			exampleProjectAccessToken, err := gitlab.NewProjectAccessToken(ctx, "exampleProjectAccessToken", &gitlab.ProjectAccessTokenArgs{
//				Project:     pulumi.String("25"),
//				ExpiresAt:   pulumi.String("2020-03-14"),
//				AccessLevel: pulumi.String("reporter"),
//				Scopes: pulumi.StringArray{
//					pulumi.String("api"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewProjectVariable(ctx, "exampleProjectVariable", &gitlab.ProjectVariableArgs{
//				Project: pulumi.Any(gitlab_project.Example.Id),
//				Key:     pulumi.String("pat"),
//				Value:   exampleProjectAccessToken.Token,
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
// A GitLab Project Access Token can be imported using a key composed of `<project-id>:<token-id>`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/projectAccessToken:ProjectAccessToken example "12345:1"
//
// ```
//
//	NOTEthe `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
type ProjectAccessToken struct {
	pulumi.CustomResourceState

	// The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
	AccessLevel pulumi.StringPtrOutput `pulumi:"accessLevel"`
	// True if the token is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// Time the token has been created, RFC3339 format.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Time the token will expire it, YYYY-MM-DD format.
	ExpiresAt pulumi.StringOutput `pulumi:"expiresAt"`
	// A name to describe the project access token.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the project to add the project access token to.
	Project pulumi.StringOutput `pulumi:"project"`
	// True if the token is revoked.
	Revoked pulumi.BoolOutput `pulumi:"revoked"`
	// The scope for the project access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// The secret token. **Note**: the token is not available for imported resources.
	Token pulumi.StringOutput `pulumi:"token"`
	// The userId associated to the token.
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewProjectAccessToken registers a new resource with the given unique name, arguments, and options.
func NewProjectAccessToken(ctx *pulumi.Context,
	name string, args *ProjectAccessTokenArgs, opts ...pulumi.ResourceOption) (*ProjectAccessToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExpiresAt == nil {
		return nil, errors.New("invalid value for required argument 'ExpiresAt'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectAccessToken
	err := ctx.RegisterResource("gitlab:index/projectAccessToken:ProjectAccessToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectAccessToken gets an existing ProjectAccessToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectAccessToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectAccessTokenState, opts ...pulumi.ResourceOption) (*ProjectAccessToken, error) {
	var resource ProjectAccessToken
	err := ctx.ReadResource("gitlab:index/projectAccessToken:ProjectAccessToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectAccessToken resources.
type projectAccessTokenState struct {
	// The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
	AccessLevel *string `pulumi:"accessLevel"`
	// True if the token is active.
	Active *bool `pulumi:"active"`
	// Time the token has been created, RFC3339 format.
	CreatedAt *string `pulumi:"createdAt"`
	// Time the token will expire it, YYYY-MM-DD format.
	ExpiresAt *string `pulumi:"expiresAt"`
	// A name to describe the project access token.
	Name *string `pulumi:"name"`
	// The id of the project to add the project access token to.
	Project *string `pulumi:"project"`
	// True if the token is revoked.
	Revoked *bool `pulumi:"revoked"`
	// The scope for the project access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
	Scopes []string `pulumi:"scopes"`
	// The secret token. **Note**: the token is not available for imported resources.
	Token *string `pulumi:"token"`
	// The userId associated to the token.
	UserId *int `pulumi:"userId"`
}

type ProjectAccessTokenState struct {
	// The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
	AccessLevel pulumi.StringPtrInput
	// True if the token is active.
	Active pulumi.BoolPtrInput
	// Time the token has been created, RFC3339 format.
	CreatedAt pulumi.StringPtrInput
	// Time the token will expire it, YYYY-MM-DD format.
	ExpiresAt pulumi.StringPtrInput
	// A name to describe the project access token.
	Name pulumi.StringPtrInput
	// The id of the project to add the project access token to.
	Project pulumi.StringPtrInput
	// True if the token is revoked.
	Revoked pulumi.BoolPtrInput
	// The scope for the project access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
	Scopes pulumi.StringArrayInput
	// The secret token. **Note**: the token is not available for imported resources.
	Token pulumi.StringPtrInput
	// The userId associated to the token.
	UserId pulumi.IntPtrInput
}

func (ProjectAccessTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectAccessTokenState)(nil)).Elem()
}

type projectAccessTokenArgs struct {
	// The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
	AccessLevel *string `pulumi:"accessLevel"`
	// Time the token will expire it, YYYY-MM-DD format.
	ExpiresAt string `pulumi:"expiresAt"`
	// A name to describe the project access token.
	Name *string `pulumi:"name"`
	// The id of the project to add the project access token to.
	Project string `pulumi:"project"`
	// The scope for the project access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a ProjectAccessToken resource.
type ProjectAccessTokenArgs struct {
	// The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
	AccessLevel pulumi.StringPtrInput
	// Time the token will expire it, YYYY-MM-DD format.
	ExpiresAt pulumi.StringInput
	// A name to describe the project access token.
	Name pulumi.StringPtrInput
	// The id of the project to add the project access token to.
	Project pulumi.StringInput
	// The scope for the project access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
	Scopes pulumi.StringArrayInput
}

func (ProjectAccessTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectAccessTokenArgs)(nil)).Elem()
}

type ProjectAccessTokenInput interface {
	pulumi.Input

	ToProjectAccessTokenOutput() ProjectAccessTokenOutput
	ToProjectAccessTokenOutputWithContext(ctx context.Context) ProjectAccessTokenOutput
}

func (*ProjectAccessToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectAccessToken)(nil)).Elem()
}

func (i *ProjectAccessToken) ToProjectAccessTokenOutput() ProjectAccessTokenOutput {
	return i.ToProjectAccessTokenOutputWithContext(context.Background())
}

func (i *ProjectAccessToken) ToProjectAccessTokenOutputWithContext(ctx context.Context) ProjectAccessTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectAccessTokenOutput)
}

// ProjectAccessTokenArrayInput is an input type that accepts ProjectAccessTokenArray and ProjectAccessTokenArrayOutput values.
// You can construct a concrete instance of `ProjectAccessTokenArrayInput` via:
//
//	ProjectAccessTokenArray{ ProjectAccessTokenArgs{...} }
type ProjectAccessTokenArrayInput interface {
	pulumi.Input

	ToProjectAccessTokenArrayOutput() ProjectAccessTokenArrayOutput
	ToProjectAccessTokenArrayOutputWithContext(context.Context) ProjectAccessTokenArrayOutput
}

type ProjectAccessTokenArray []ProjectAccessTokenInput

func (ProjectAccessTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectAccessToken)(nil)).Elem()
}

func (i ProjectAccessTokenArray) ToProjectAccessTokenArrayOutput() ProjectAccessTokenArrayOutput {
	return i.ToProjectAccessTokenArrayOutputWithContext(context.Background())
}

func (i ProjectAccessTokenArray) ToProjectAccessTokenArrayOutputWithContext(ctx context.Context) ProjectAccessTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectAccessTokenArrayOutput)
}

// ProjectAccessTokenMapInput is an input type that accepts ProjectAccessTokenMap and ProjectAccessTokenMapOutput values.
// You can construct a concrete instance of `ProjectAccessTokenMapInput` via:
//
//	ProjectAccessTokenMap{ "key": ProjectAccessTokenArgs{...} }
type ProjectAccessTokenMapInput interface {
	pulumi.Input

	ToProjectAccessTokenMapOutput() ProjectAccessTokenMapOutput
	ToProjectAccessTokenMapOutputWithContext(context.Context) ProjectAccessTokenMapOutput
}

type ProjectAccessTokenMap map[string]ProjectAccessTokenInput

func (ProjectAccessTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectAccessToken)(nil)).Elem()
}

func (i ProjectAccessTokenMap) ToProjectAccessTokenMapOutput() ProjectAccessTokenMapOutput {
	return i.ToProjectAccessTokenMapOutputWithContext(context.Background())
}

func (i ProjectAccessTokenMap) ToProjectAccessTokenMapOutputWithContext(ctx context.Context) ProjectAccessTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectAccessTokenMapOutput)
}

type ProjectAccessTokenOutput struct{ *pulumi.OutputState }

func (ProjectAccessTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectAccessToken)(nil)).Elem()
}

func (o ProjectAccessTokenOutput) ToProjectAccessTokenOutput() ProjectAccessTokenOutput {
	return o
}

func (o ProjectAccessTokenOutput) ToProjectAccessTokenOutputWithContext(ctx context.Context) ProjectAccessTokenOutput {
	return o
}

// The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
func (o ProjectAccessTokenOutput) AccessLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectAccessToken) pulumi.StringPtrOutput { return v.AccessLevel }).(pulumi.StringPtrOutput)
}

// True if the token is active.
func (o ProjectAccessTokenOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectAccessToken) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// Time the token has been created, RFC3339 format.
func (o ProjectAccessTokenOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectAccessToken) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Time the token will expire it, YYYY-MM-DD format.
func (o ProjectAccessTokenOutput) ExpiresAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectAccessToken) pulumi.StringOutput { return v.ExpiresAt }).(pulumi.StringOutput)
}

// A name to describe the project access token.
func (o ProjectAccessTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectAccessToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the project to add the project access token to.
func (o ProjectAccessTokenOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectAccessToken) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// True if the token is revoked.
func (o ProjectAccessTokenOutput) Revoked() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectAccessToken) pulumi.BoolOutput { return v.Revoked }).(pulumi.BoolOutput)
}

// The scope for the project access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`.
func (o ProjectAccessTokenOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectAccessToken) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// The secret token. **Note**: the token is not available for imported resources.
func (o ProjectAccessTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectAccessToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The userId associated to the token.
func (o ProjectAccessTokenOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectAccessToken) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type ProjectAccessTokenArrayOutput struct{ *pulumi.OutputState }

func (ProjectAccessTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectAccessToken)(nil)).Elem()
}

func (o ProjectAccessTokenArrayOutput) ToProjectAccessTokenArrayOutput() ProjectAccessTokenArrayOutput {
	return o
}

func (o ProjectAccessTokenArrayOutput) ToProjectAccessTokenArrayOutputWithContext(ctx context.Context) ProjectAccessTokenArrayOutput {
	return o
}

func (o ProjectAccessTokenArrayOutput) Index(i pulumi.IntInput) ProjectAccessTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectAccessToken {
		return vs[0].([]*ProjectAccessToken)[vs[1].(int)]
	}).(ProjectAccessTokenOutput)
}

type ProjectAccessTokenMapOutput struct{ *pulumi.OutputState }

func (ProjectAccessTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectAccessToken)(nil)).Elem()
}

func (o ProjectAccessTokenMapOutput) ToProjectAccessTokenMapOutput() ProjectAccessTokenMapOutput {
	return o
}

func (o ProjectAccessTokenMapOutput) ToProjectAccessTokenMapOutputWithContext(ctx context.Context) ProjectAccessTokenMapOutput {
	return o
}

func (o ProjectAccessTokenMapOutput) MapIndex(k pulumi.StringInput) ProjectAccessTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectAccessToken {
		return vs[0].(map[string]*ProjectAccessToken)[vs[1].(string)]
	}).(ProjectAccessTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectAccessTokenInput)(nil)).Elem(), &ProjectAccessToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectAccessTokenArrayInput)(nil)).Elem(), ProjectAccessTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectAccessTokenMapInput)(nil)).Elem(), ProjectAccessTokenMap{})
	pulumi.RegisterOutputType(ProjectAccessTokenOutput{})
	pulumi.RegisterOutputType(ProjectAccessTokenArrayOutput{})
	pulumi.RegisterOutputType(ProjectAccessTokenMapOutput{})
}
