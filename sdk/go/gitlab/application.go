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

// The `Application` resource allows to manage the lifecycle of applications in gitlab.
//
// > In order to use a user for a user to create an application, they must have admin privileges at the instance level.
// To create an OIDC application, a scope of "openid".
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/applications.html)
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
//			_, err := gitlab.NewApplication(ctx, "oidc", &gitlab.ApplicationArgs{
//				Confidential: pulumi.Bool(true),
//				Scopes: pulumi.StringArray{
//					pulumi.String("openid"),
//				},
//				Name:        pulumi.String("company_oidc"),
//				RedirectUrl: pulumi.String("https://mycompany.com"),
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
// Gitlab applications can be imported with their id, e.g.
//
// ```sh
// $ pulumi import gitlab:index/application:Application example "1"
// ```
//
// NOTE: the secret and scopes cannot be imported
type Application struct {
	pulumi.CustomResourceState

	// Internal name of the application.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
	Confidential pulumi.BoolOutput `pulumi:"confidential"`
	// Name of the application.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL gitlab should send the user to after authentication.
	RedirectUrl pulumi.StringOutput `pulumi:"redirectUrl"`
	// Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `readApi`, `readUser`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `openid`, `profile`, `email`.
	// This is only populated when creating a new application. This attribute is not available for imported resources
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
	Secret pulumi.StringOutput `pulumi:"secret"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RedirectUrl == nil {
		return nil, errors.New("invalid value for required argument 'RedirectUrl'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("gitlab:index/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("gitlab:index/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// Internal name of the application.
	ApplicationId *string `pulumi:"applicationId"`
	// The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
	Confidential *bool `pulumi:"confidential"`
	// Name of the application.
	Name *string `pulumi:"name"`
	// The URL gitlab should send the user to after authentication.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `readApi`, `readUser`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `openid`, `profile`, `email`.
	// This is only populated when creating a new application. This attribute is not available for imported resources
	Scopes []string `pulumi:"scopes"`
	// Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
	Secret *string `pulumi:"secret"`
}

type ApplicationState struct {
	// Internal name of the application.
	ApplicationId pulumi.StringPtrInput
	// The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
	Confidential pulumi.BoolPtrInput
	// Name of the application.
	Name pulumi.StringPtrInput
	// The URL gitlab should send the user to after authentication.
	RedirectUrl pulumi.StringPtrInput
	// Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `readApi`, `readUser`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `openid`, `profile`, `email`.
	// This is only populated when creating a new application. This attribute is not available for imported resources
	Scopes pulumi.StringArrayInput
	// Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
	Secret pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
	Confidential *bool `pulumi:"confidential"`
	// Name of the application.
	Name *string `pulumi:"name"`
	// The URL gitlab should send the user to after authentication.
	RedirectUrl string `pulumi:"redirectUrl"`
	// Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `readApi`, `readUser`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `openid`, `profile`, `email`.
	// This is only populated when creating a new application. This attribute is not available for imported resources
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
	Confidential pulumi.BoolPtrInput
	// Name of the application.
	Name pulumi.StringPtrInput
	// The URL gitlab should send the user to after authentication.
	RedirectUrl pulumi.StringInput
	// Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `readApi`, `readUser`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `openid`, `profile`, `email`.
	// This is only populated when creating a new application. This attribute is not available for imported resources
	Scopes pulumi.StringArrayInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

// ApplicationArrayInput is an input type that accepts ApplicationArray and ApplicationArrayOutput values.
// You can construct a concrete instance of `ApplicationArrayInput` via:
//
//	ApplicationArray{ ApplicationArgs{...} }
type ApplicationArrayInput interface {
	pulumi.Input

	ToApplicationArrayOutput() ApplicationArrayOutput
	ToApplicationArrayOutputWithContext(context.Context) ApplicationArrayOutput
}

type ApplicationArray []ApplicationInput

func (ApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (i ApplicationArray) ToApplicationArrayOutput() ApplicationArrayOutput {
	return i.ToApplicationArrayOutputWithContext(context.Background())
}

func (i ApplicationArray) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArrayOutput)
}

// ApplicationMapInput is an input type that accepts ApplicationMap and ApplicationMapOutput values.
// You can construct a concrete instance of `ApplicationMapInput` via:
//
//	ApplicationMap{ "key": ApplicationArgs{...} }
type ApplicationMapInput interface {
	pulumi.Input

	ToApplicationMapOutput() ApplicationMapOutput
	ToApplicationMapOutputWithContext(context.Context) ApplicationMapOutput
}

type ApplicationMap map[string]ApplicationInput

func (ApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (i ApplicationMap) ToApplicationMapOutput() ApplicationMapOutput {
	return i.ToApplicationMapOutputWithContext(context.Background())
}

func (i ApplicationMap) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMapOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

// Internal name of the application.
func (o ApplicationOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
func (o ApplicationOutput) Confidential() pulumi.BoolOutput {
	return o.ApplyT(func(v *Application) pulumi.BoolOutput { return v.Confidential }).(pulumi.BoolOutput)
}

// Name of the application.
func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The URL gitlab should send the user to after authentication.
func (o ApplicationOutput) RedirectUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.RedirectUrl }).(pulumi.StringOutput)
}

// Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `readApi`, `readUser`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `openid`, `profile`, `email`.
// This is only populated when creating a new application. This attribute is not available for imported resources
func (o ApplicationOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Application) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
func (o ApplicationOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

type ApplicationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (o ApplicationArrayOutput) ToApplicationArrayOutput() ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) Index(i pulumi.IntInput) ApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Application {
		return vs[0].([]*Application)[vs[1].(int)]
	}).(ApplicationOutput)
}

type ApplicationMapOutput struct{ *pulumi.OutputState }

func (ApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (o ApplicationMapOutput) ToApplicationMapOutput() ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) MapIndex(k pulumi.StringInput) ApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Application {
		return vs[0].(map[string]*Application)[vs[1].(string)]
	}).(ApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArrayInput)(nil)).Elem(), ApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationMapInput)(nil)).Elem(), ApplicationMap{})
	pulumi.RegisterOutputType(ApplicationOutput{})
	pulumi.RegisterOutputType(ApplicationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationMapOutput{})
}
