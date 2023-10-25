// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The `PagesDomain` resource allows connecting custom domains and TLS certificates in GitLab Pages.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pages_domains.html)
//
// ## Import
//
// GitLab pages domain can be imported using an id made up of `projectId:domain` _without_ the http protocol, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/pagesDomain:PagesDomain this 123:example.com
//
// ```
type PagesDomain struct {
	pulumi.CustomResourceState

	// Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
	AutoSslEnabled pulumi.BoolOutput `pulumi:"autoSslEnabled"`
	// The certificate in PEM format with intermediates following in most specific to least specific order.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The custom domain indicated by the user.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Whether the certificate is expired.
	Expired pulumi.BoolOutput `pulumi:"expired"`
	// The certificate key in PEM format.
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URL for the given domain.
	Url pulumi.StringOutput `pulumi:"url"`
	// The verification code for the domain.
	VerificationCode pulumi.StringOutput `pulumi:"verificationCode"`
	// The certificate data.
	Verified pulumi.BoolOutput `pulumi:"verified"`
}

// NewPagesDomain registers a new resource with the given unique name, arguments, and options.
func NewPagesDomain(ctx *pulumi.Context,
	name string, args *PagesDomainArgs, opts ...pulumi.ResourceOption) (*PagesDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"verificationCode",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PagesDomain
	err := ctx.RegisterResource("gitlab:index/pagesDomain:PagesDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPagesDomain gets an existing PagesDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPagesDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PagesDomainState, opts ...pulumi.ResourceOption) (*PagesDomain, error) {
	var resource PagesDomain
	err := ctx.ReadResource("gitlab:index/pagesDomain:PagesDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PagesDomain resources.
type pagesDomainState struct {
	// Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
	AutoSslEnabled *bool `pulumi:"autoSslEnabled"`
	// The certificate in PEM format with intermediates following in most specific to least specific order.
	Certificate *string `pulumi:"certificate"`
	// The custom domain indicated by the user.
	Domain *string `pulumi:"domain"`
	// Whether the certificate is expired.
	Expired *bool `pulumi:"expired"`
	// The certificate key in PEM format.
	Key *string `pulumi:"key"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
	Project *string `pulumi:"project"`
	// The URL for the given domain.
	Url *string `pulumi:"url"`
	// The verification code for the domain.
	VerificationCode *string `pulumi:"verificationCode"`
	// The certificate data.
	Verified *bool `pulumi:"verified"`
}

type PagesDomainState struct {
	// Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
	AutoSslEnabled pulumi.BoolPtrInput
	// The certificate in PEM format with intermediates following in most specific to least specific order.
	Certificate pulumi.StringPtrInput
	// The custom domain indicated by the user.
	Domain pulumi.StringPtrInput
	// Whether the certificate is expired.
	Expired pulumi.BoolPtrInput
	// The certificate key in PEM format.
	Key pulumi.StringPtrInput
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
	Project pulumi.StringPtrInput
	// The URL for the given domain.
	Url pulumi.StringPtrInput
	// The verification code for the domain.
	VerificationCode pulumi.StringPtrInput
	// The certificate data.
	Verified pulumi.BoolPtrInput
}

func (PagesDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*pagesDomainState)(nil)).Elem()
}

type pagesDomainArgs struct {
	// Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
	AutoSslEnabled *bool `pulumi:"autoSslEnabled"`
	// The certificate in PEM format with intermediates following in most specific to least specific order.
	Certificate *string `pulumi:"certificate"`
	// The custom domain indicated by the user.
	Domain string `pulumi:"domain"`
	// Whether the certificate is expired.
	Expired *bool `pulumi:"expired"`
	// The certificate key in PEM format.
	Key *string `pulumi:"key"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a PagesDomain resource.
type PagesDomainArgs struct {
	// Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
	AutoSslEnabled pulumi.BoolPtrInput
	// The certificate in PEM format with intermediates following in most specific to least specific order.
	Certificate pulumi.StringPtrInput
	// The custom domain indicated by the user.
	Domain pulumi.StringInput
	// Whether the certificate is expired.
	Expired pulumi.BoolPtrInput
	// The certificate key in PEM format.
	Key pulumi.StringPtrInput
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
	Project pulumi.StringInput
}

func (PagesDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pagesDomainArgs)(nil)).Elem()
}

type PagesDomainInput interface {
	pulumi.Input

	ToPagesDomainOutput() PagesDomainOutput
	ToPagesDomainOutputWithContext(ctx context.Context) PagesDomainOutput
}

func (*PagesDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**PagesDomain)(nil)).Elem()
}

func (i *PagesDomain) ToPagesDomainOutput() PagesDomainOutput {
	return i.ToPagesDomainOutputWithContext(context.Background())
}

func (i *PagesDomain) ToPagesDomainOutputWithContext(ctx context.Context) PagesDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PagesDomainOutput)
}

func (i *PagesDomain) ToOutput(ctx context.Context) pulumix.Output[*PagesDomain] {
	return pulumix.Output[*PagesDomain]{
		OutputState: i.ToPagesDomainOutputWithContext(ctx).OutputState,
	}
}

// PagesDomainArrayInput is an input type that accepts PagesDomainArray and PagesDomainArrayOutput values.
// You can construct a concrete instance of `PagesDomainArrayInput` via:
//
//	PagesDomainArray{ PagesDomainArgs{...} }
type PagesDomainArrayInput interface {
	pulumi.Input

	ToPagesDomainArrayOutput() PagesDomainArrayOutput
	ToPagesDomainArrayOutputWithContext(context.Context) PagesDomainArrayOutput
}

type PagesDomainArray []PagesDomainInput

func (PagesDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PagesDomain)(nil)).Elem()
}

func (i PagesDomainArray) ToPagesDomainArrayOutput() PagesDomainArrayOutput {
	return i.ToPagesDomainArrayOutputWithContext(context.Background())
}

func (i PagesDomainArray) ToPagesDomainArrayOutputWithContext(ctx context.Context) PagesDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PagesDomainArrayOutput)
}

func (i PagesDomainArray) ToOutput(ctx context.Context) pulumix.Output[[]*PagesDomain] {
	return pulumix.Output[[]*PagesDomain]{
		OutputState: i.ToPagesDomainArrayOutputWithContext(ctx).OutputState,
	}
}

// PagesDomainMapInput is an input type that accepts PagesDomainMap and PagesDomainMapOutput values.
// You can construct a concrete instance of `PagesDomainMapInput` via:
//
//	PagesDomainMap{ "key": PagesDomainArgs{...} }
type PagesDomainMapInput interface {
	pulumi.Input

	ToPagesDomainMapOutput() PagesDomainMapOutput
	ToPagesDomainMapOutputWithContext(context.Context) PagesDomainMapOutput
}

type PagesDomainMap map[string]PagesDomainInput

func (PagesDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PagesDomain)(nil)).Elem()
}

func (i PagesDomainMap) ToPagesDomainMapOutput() PagesDomainMapOutput {
	return i.ToPagesDomainMapOutputWithContext(context.Background())
}

func (i PagesDomainMap) ToPagesDomainMapOutputWithContext(ctx context.Context) PagesDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PagesDomainMapOutput)
}

func (i PagesDomainMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PagesDomain] {
	return pulumix.Output[map[string]*PagesDomain]{
		OutputState: i.ToPagesDomainMapOutputWithContext(ctx).OutputState,
	}
}

type PagesDomainOutput struct{ *pulumi.OutputState }

func (PagesDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PagesDomain)(nil)).Elem()
}

func (o PagesDomainOutput) ToPagesDomainOutput() PagesDomainOutput {
	return o
}

func (o PagesDomainOutput) ToPagesDomainOutputWithContext(ctx context.Context) PagesDomainOutput {
	return o
}

func (o PagesDomainOutput) ToOutput(ctx context.Context) pulumix.Output[*PagesDomain] {
	return pulumix.Output[*PagesDomain]{
		OutputState: o.OutputState,
	}
}

// Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
func (o PagesDomainOutput) AutoSslEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *PagesDomain) pulumi.BoolOutput { return v.AutoSslEnabled }).(pulumi.BoolOutput)
}

// The certificate in PEM format with intermediates following in most specific to least specific order.
func (o PagesDomainOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *PagesDomain) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// The custom domain indicated by the user.
func (o PagesDomainOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *PagesDomain) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Whether the certificate is expired.
func (o PagesDomainOutput) Expired() pulumi.BoolOutput {
	return o.ApplyT(func(v *PagesDomain) pulumi.BoolOutput { return v.Expired }).(pulumi.BoolOutput)
}

// The certificate key in PEM format.
func (o PagesDomainOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PagesDomain) pulumi.StringPtrOutput { return v.Key }).(pulumi.StringPtrOutput)
}

// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
func (o PagesDomainOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PagesDomain) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URL for the given domain.
func (o PagesDomainOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *PagesDomain) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The verification code for the domain.
func (o PagesDomainOutput) VerificationCode() pulumi.StringOutput {
	return o.ApplyT(func(v *PagesDomain) pulumi.StringOutput { return v.VerificationCode }).(pulumi.StringOutput)
}

// The certificate data.
func (o PagesDomainOutput) Verified() pulumi.BoolOutput {
	return o.ApplyT(func(v *PagesDomain) pulumi.BoolOutput { return v.Verified }).(pulumi.BoolOutput)
}

type PagesDomainArrayOutput struct{ *pulumi.OutputState }

func (PagesDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PagesDomain)(nil)).Elem()
}

func (o PagesDomainArrayOutput) ToPagesDomainArrayOutput() PagesDomainArrayOutput {
	return o
}

func (o PagesDomainArrayOutput) ToPagesDomainArrayOutputWithContext(ctx context.Context) PagesDomainArrayOutput {
	return o
}

func (o PagesDomainArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PagesDomain] {
	return pulumix.Output[[]*PagesDomain]{
		OutputState: o.OutputState,
	}
}

func (o PagesDomainArrayOutput) Index(i pulumi.IntInput) PagesDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PagesDomain {
		return vs[0].([]*PagesDomain)[vs[1].(int)]
	}).(PagesDomainOutput)
}

type PagesDomainMapOutput struct{ *pulumi.OutputState }

func (PagesDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PagesDomain)(nil)).Elem()
}

func (o PagesDomainMapOutput) ToPagesDomainMapOutput() PagesDomainMapOutput {
	return o
}

func (o PagesDomainMapOutput) ToPagesDomainMapOutputWithContext(ctx context.Context) PagesDomainMapOutput {
	return o
}

func (o PagesDomainMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PagesDomain] {
	return pulumix.Output[map[string]*PagesDomain]{
		OutputState: o.OutputState,
	}
}

func (o PagesDomainMapOutput) MapIndex(k pulumi.StringInput) PagesDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PagesDomain {
		return vs[0].(map[string]*PagesDomain)[vs[1].(string)]
	}).(PagesDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PagesDomainInput)(nil)).Elem(), &PagesDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*PagesDomainArrayInput)(nil)).Elem(), PagesDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PagesDomainMapInput)(nil)).Elem(), PagesDomainMap{})
	pulumi.RegisterOutputType(PagesDomainOutput{})
	pulumi.RegisterOutputType(PagesDomainArrayOutput{})
	pulumi.RegisterOutputType(PagesDomainMapOutput{})
}
