// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gitlab

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The provider type for the gitlab package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/index.html.markdown.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.BaseUrl == nil {
		args.BaseUrl = pulumi.StringPtr(getEnvOrDefault("", nil, "GITLAB_BASE_URL").(string))
	}
	if args.Token == nil {
		args.Token = pulumi.StringPtr(getEnvOrDefault("", nil, "GITLAB_TOKEN").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:gitlab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The GitLab Base API URL
	BaseUrl *string `pulumi:"baseUrl"`
	// A file containing the ca certificate to use in case ssl certificate is not from a standard chain
	CacertFile *string `pulumi:"cacertFile"`
	// Disable SSL verification of API calls
	Insecure *bool `pulumi:"insecure"`
	// The OAuth token used to connect to GitLab.
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The GitLab Base API URL
	BaseUrl pulumi.StringPtrInput
	// A file containing the ca certificate to use in case ssl certificate is not from a standard chain
	CacertFile pulumi.StringPtrInput
	// Disable SSL verification of API calls
	Insecure pulumi.BoolPtrInput
	// The OAuth token used to connect to GitLab.
	Token pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}