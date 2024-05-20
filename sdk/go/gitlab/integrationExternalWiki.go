// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v7/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `IntegrationExternalWiki` resource allows to manage the lifecycle of a project integration with External Wiki Service.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#external-wiki)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v7/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			awesomeProject, err := gitlab.NewProject(ctx, "awesome_project", &gitlab.ProjectArgs{
//				Name:            pulumi.String("awesome_project"),
//				Description:     pulumi.String("My awesome project."),
//				VisibilityLevel: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewIntegrationExternalWiki(ctx, "wiki", &gitlab.IntegrationExternalWikiArgs{
//				Project:         awesomeProject.ID(),
//				ExternalWikiUrl: pulumi.String("https://MyAwesomeExternalWikiURL.com"),
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
// You can import a gitlab_integration_external_wiki state using the project ID, e.g.
//
// ```sh
// $ pulumi import gitlab:index/integrationExternalWiki:IntegrationExternalWiki wiki 1
// ```
type IntegrationExternalWiki struct {
	pulumi.CustomResourceState

	// Whether the integration is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// The ISO8601 date/time that this integration was activated at in UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The URL of the external wiki.
	ExternalWikiUrl pulumi.StringOutput `pulumi:"externalWikiUrl"`
	// ID of the project you want to activate integration on.
	Project pulumi.StringOutput `pulumi:"project"`
	// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// Title of the integration.
	Title pulumi.StringOutput `pulumi:"title"`
	// The ISO8601 date/time that this integration was last updated at in UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewIntegrationExternalWiki registers a new resource with the given unique name, arguments, and options.
func NewIntegrationExternalWiki(ctx *pulumi.Context,
	name string, args *IntegrationExternalWikiArgs, opts ...pulumi.ResourceOption) (*IntegrationExternalWiki, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExternalWikiUrl == nil {
		return nil, errors.New("invalid value for required argument 'ExternalWikiUrl'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IntegrationExternalWiki
	err := ctx.RegisterResource("gitlab:index/integrationExternalWiki:IntegrationExternalWiki", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationExternalWiki gets an existing IntegrationExternalWiki resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationExternalWiki(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationExternalWikiState, opts ...pulumi.ResourceOption) (*IntegrationExternalWiki, error) {
	var resource IntegrationExternalWiki
	err := ctx.ReadResource("gitlab:index/integrationExternalWiki:IntegrationExternalWiki", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationExternalWiki resources.
type integrationExternalWikiState struct {
	// Whether the integration is active.
	Active *bool `pulumi:"active"`
	// The ISO8601 date/time that this integration was activated at in UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// The URL of the external wiki.
	ExternalWikiUrl *string `pulumi:"externalWikiUrl"`
	// ID of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug *string `pulumi:"slug"`
	// Title of the integration.
	Title *string `pulumi:"title"`
	// The ISO8601 date/time that this integration was last updated at in UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type IntegrationExternalWikiState struct {
	// Whether the integration is active.
	Active pulumi.BoolPtrInput
	// The ISO8601 date/time that this integration was activated at in UTC.
	CreatedAt pulumi.StringPtrInput
	// The URL of the external wiki.
	ExternalWikiUrl pulumi.StringPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug pulumi.StringPtrInput
	// Title of the integration.
	Title pulumi.StringPtrInput
	// The ISO8601 date/time that this integration was last updated at in UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (IntegrationExternalWikiState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationExternalWikiState)(nil)).Elem()
}

type integrationExternalWikiArgs struct {
	// The URL of the external wiki.
	ExternalWikiUrl string `pulumi:"externalWikiUrl"`
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a IntegrationExternalWiki resource.
type IntegrationExternalWikiArgs struct {
	// The URL of the external wiki.
	ExternalWikiUrl pulumi.StringInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
}

func (IntegrationExternalWikiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationExternalWikiArgs)(nil)).Elem()
}

type IntegrationExternalWikiInput interface {
	pulumi.Input

	ToIntegrationExternalWikiOutput() IntegrationExternalWikiOutput
	ToIntegrationExternalWikiOutputWithContext(ctx context.Context) IntegrationExternalWikiOutput
}

func (*IntegrationExternalWiki) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationExternalWiki)(nil)).Elem()
}

func (i *IntegrationExternalWiki) ToIntegrationExternalWikiOutput() IntegrationExternalWikiOutput {
	return i.ToIntegrationExternalWikiOutputWithContext(context.Background())
}

func (i *IntegrationExternalWiki) ToIntegrationExternalWikiOutputWithContext(ctx context.Context) IntegrationExternalWikiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationExternalWikiOutput)
}

// IntegrationExternalWikiArrayInput is an input type that accepts IntegrationExternalWikiArray and IntegrationExternalWikiArrayOutput values.
// You can construct a concrete instance of `IntegrationExternalWikiArrayInput` via:
//
//	IntegrationExternalWikiArray{ IntegrationExternalWikiArgs{...} }
type IntegrationExternalWikiArrayInput interface {
	pulumi.Input

	ToIntegrationExternalWikiArrayOutput() IntegrationExternalWikiArrayOutput
	ToIntegrationExternalWikiArrayOutputWithContext(context.Context) IntegrationExternalWikiArrayOutput
}

type IntegrationExternalWikiArray []IntegrationExternalWikiInput

func (IntegrationExternalWikiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationExternalWiki)(nil)).Elem()
}

func (i IntegrationExternalWikiArray) ToIntegrationExternalWikiArrayOutput() IntegrationExternalWikiArrayOutput {
	return i.ToIntegrationExternalWikiArrayOutputWithContext(context.Background())
}

func (i IntegrationExternalWikiArray) ToIntegrationExternalWikiArrayOutputWithContext(ctx context.Context) IntegrationExternalWikiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationExternalWikiArrayOutput)
}

// IntegrationExternalWikiMapInput is an input type that accepts IntegrationExternalWikiMap and IntegrationExternalWikiMapOutput values.
// You can construct a concrete instance of `IntegrationExternalWikiMapInput` via:
//
//	IntegrationExternalWikiMap{ "key": IntegrationExternalWikiArgs{...} }
type IntegrationExternalWikiMapInput interface {
	pulumi.Input

	ToIntegrationExternalWikiMapOutput() IntegrationExternalWikiMapOutput
	ToIntegrationExternalWikiMapOutputWithContext(context.Context) IntegrationExternalWikiMapOutput
}

type IntegrationExternalWikiMap map[string]IntegrationExternalWikiInput

func (IntegrationExternalWikiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationExternalWiki)(nil)).Elem()
}

func (i IntegrationExternalWikiMap) ToIntegrationExternalWikiMapOutput() IntegrationExternalWikiMapOutput {
	return i.ToIntegrationExternalWikiMapOutputWithContext(context.Background())
}

func (i IntegrationExternalWikiMap) ToIntegrationExternalWikiMapOutputWithContext(ctx context.Context) IntegrationExternalWikiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationExternalWikiMapOutput)
}

type IntegrationExternalWikiOutput struct{ *pulumi.OutputState }

func (IntegrationExternalWikiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationExternalWiki)(nil)).Elem()
}

func (o IntegrationExternalWikiOutput) ToIntegrationExternalWikiOutput() IntegrationExternalWikiOutput {
	return o
}

func (o IntegrationExternalWikiOutput) ToIntegrationExternalWikiOutputWithContext(ctx context.Context) IntegrationExternalWikiOutput {
	return o
}

// Whether the integration is active.
func (o IntegrationExternalWikiOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationExternalWiki) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// The ISO8601 date/time that this integration was activated at in UTC.
func (o IntegrationExternalWikiOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationExternalWiki) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The URL of the external wiki.
func (o IntegrationExternalWikiOutput) ExternalWikiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationExternalWiki) pulumi.StringOutput { return v.ExternalWikiUrl }).(pulumi.StringOutput)
}

// ID of the project you want to activate integration on.
func (o IntegrationExternalWikiOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationExternalWiki) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
func (o IntegrationExternalWikiOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationExternalWiki) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// Title of the integration.
func (o IntegrationExternalWikiOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationExternalWiki) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The ISO8601 date/time that this integration was last updated at in UTC.
func (o IntegrationExternalWikiOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationExternalWiki) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type IntegrationExternalWikiArrayOutput struct{ *pulumi.OutputState }

func (IntegrationExternalWikiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationExternalWiki)(nil)).Elem()
}

func (o IntegrationExternalWikiArrayOutput) ToIntegrationExternalWikiArrayOutput() IntegrationExternalWikiArrayOutput {
	return o
}

func (o IntegrationExternalWikiArrayOutput) ToIntegrationExternalWikiArrayOutputWithContext(ctx context.Context) IntegrationExternalWikiArrayOutput {
	return o
}

func (o IntegrationExternalWikiArrayOutput) Index(i pulumi.IntInput) IntegrationExternalWikiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationExternalWiki {
		return vs[0].([]*IntegrationExternalWiki)[vs[1].(int)]
	}).(IntegrationExternalWikiOutput)
}

type IntegrationExternalWikiMapOutput struct{ *pulumi.OutputState }

func (IntegrationExternalWikiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationExternalWiki)(nil)).Elem()
}

func (o IntegrationExternalWikiMapOutput) ToIntegrationExternalWikiMapOutput() IntegrationExternalWikiMapOutput {
	return o
}

func (o IntegrationExternalWikiMapOutput) ToIntegrationExternalWikiMapOutputWithContext(ctx context.Context) IntegrationExternalWikiMapOutput {
	return o
}

func (o IntegrationExternalWikiMapOutput) MapIndex(k pulumi.StringInput) IntegrationExternalWikiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationExternalWiki {
		return vs[0].(map[string]*IntegrationExternalWiki)[vs[1].(string)]
	}).(IntegrationExternalWikiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationExternalWikiInput)(nil)).Elem(), &IntegrationExternalWiki{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationExternalWikiArrayInput)(nil)).Elem(), IntegrationExternalWikiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationExternalWikiMapInput)(nil)).Elem(), IntegrationExternalWikiMap{})
	pulumi.RegisterOutputType(IntegrationExternalWikiOutput{})
	pulumi.RegisterOutputType(IntegrationExternalWikiArrayOutput{})
	pulumi.RegisterOutputType(IntegrationExternalWikiMapOutput{})
}
