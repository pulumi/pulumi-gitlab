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

// The `IntegrationGithub` resource allows to manage the lifecycle of a project integration with GitHub.
//
// > This resource requires a GitLab Enterprise instance.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#github)
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
//			awesomeProject, err := gitlab.NewProject(ctx, "awesomeProject", &gitlab.ProjectArgs{
//				Description:     pulumi.String("My awesome project."),
//				VisibilityLevel: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewIntegrationGithub(ctx, "github", &gitlab.IntegrationGithubArgs{
//				Project:       awesomeProject.ID(),
//				Token:         pulumi.String("REDACTED"),
//				RepositoryUrl: pulumi.String("https://github.com/gitlabhq/terraform-provider-gitlab"),
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
// ```sh
//
//	$ pulumi import gitlab:index/integrationGithub:IntegrationGithub You can import a gitlab_integration_github state using `<resource> <project_id>`:
//
// ```
//
// ```sh
// $ pulumi import gitlab:index/integrationGithub:IntegrationGithub github 1
// ```
type IntegrationGithub struct {
	pulumi.CustomResourceState

	// Whether the integration is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// Create time.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// ID of the project you want to activate integration on.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
	RepositoryUrl pulumi.StringOutput `pulumi:"repositoryUrl"`
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext pulumi.BoolPtrOutput `pulumi:"staticContext"`
	// Title.
	Title pulumi.StringOutput `pulumi:"title"`
	// A GitHub personal access token with at least `repo:status` scope.
	Token pulumi.StringOutput `pulumi:"token"`
	// Update time.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewIntegrationGithub registers a new resource with the given unique name, arguments, and options.
func NewIntegrationGithub(ctx *pulumi.Context,
	name string, args *IntegrationGithubArgs, opts ...pulumi.ResourceOption) (*IntegrationGithub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.RepositoryUrl == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryUrl'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IntegrationGithub
	err := ctx.RegisterResource("gitlab:index/integrationGithub:IntegrationGithub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationGithub gets an existing IntegrationGithub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationGithub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationGithubState, opts ...pulumi.ResourceOption) (*IntegrationGithub, error) {
	var resource IntegrationGithub
	err := ctx.ReadResource("gitlab:index/integrationGithub:IntegrationGithub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationGithub resources.
type integrationGithubState struct {
	// Whether the integration is active.
	Active *bool `pulumi:"active"`
	// Create time.
	CreatedAt *string `pulumi:"createdAt"`
	// ID of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext *bool `pulumi:"staticContext"`
	// Title.
	Title *string `pulumi:"title"`
	// A GitHub personal access token with at least `repo:status` scope.
	Token *string `pulumi:"token"`
	// Update time.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type IntegrationGithubState struct {
	// Whether the integration is active.
	Active pulumi.BoolPtrInput
	// Create time.
	CreatedAt pulumi.StringPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
	RepositoryUrl pulumi.StringPtrInput
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext pulumi.BoolPtrInput
	// Title.
	Title pulumi.StringPtrInput
	// A GitHub personal access token with at least `repo:status` scope.
	Token pulumi.StringPtrInput
	// Update time.
	UpdatedAt pulumi.StringPtrInput
}

func (IntegrationGithubState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationGithubState)(nil)).Elem()
}

type integrationGithubArgs struct {
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
	RepositoryUrl string `pulumi:"repositoryUrl"`
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext *bool `pulumi:"staticContext"`
	// A GitHub personal access token with at least `repo:status` scope.
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a IntegrationGithub resource.
type IntegrationGithubArgs struct {
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
	// The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
	RepositoryUrl pulumi.StringInput
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext pulumi.BoolPtrInput
	// A GitHub personal access token with at least `repo:status` scope.
	Token pulumi.StringInput
}

func (IntegrationGithubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationGithubArgs)(nil)).Elem()
}

type IntegrationGithubInput interface {
	pulumi.Input

	ToIntegrationGithubOutput() IntegrationGithubOutput
	ToIntegrationGithubOutputWithContext(ctx context.Context) IntegrationGithubOutput
}

func (*IntegrationGithub) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationGithub)(nil)).Elem()
}

func (i *IntegrationGithub) ToIntegrationGithubOutput() IntegrationGithubOutput {
	return i.ToIntegrationGithubOutputWithContext(context.Background())
}

func (i *IntegrationGithub) ToIntegrationGithubOutputWithContext(ctx context.Context) IntegrationGithubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationGithubOutput)
}

// IntegrationGithubArrayInput is an input type that accepts IntegrationGithubArray and IntegrationGithubArrayOutput values.
// You can construct a concrete instance of `IntegrationGithubArrayInput` via:
//
//	IntegrationGithubArray{ IntegrationGithubArgs{...} }
type IntegrationGithubArrayInput interface {
	pulumi.Input

	ToIntegrationGithubArrayOutput() IntegrationGithubArrayOutput
	ToIntegrationGithubArrayOutputWithContext(context.Context) IntegrationGithubArrayOutput
}

type IntegrationGithubArray []IntegrationGithubInput

func (IntegrationGithubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationGithub)(nil)).Elem()
}

func (i IntegrationGithubArray) ToIntegrationGithubArrayOutput() IntegrationGithubArrayOutput {
	return i.ToIntegrationGithubArrayOutputWithContext(context.Background())
}

func (i IntegrationGithubArray) ToIntegrationGithubArrayOutputWithContext(ctx context.Context) IntegrationGithubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationGithubArrayOutput)
}

// IntegrationGithubMapInput is an input type that accepts IntegrationGithubMap and IntegrationGithubMapOutput values.
// You can construct a concrete instance of `IntegrationGithubMapInput` via:
//
//	IntegrationGithubMap{ "key": IntegrationGithubArgs{...} }
type IntegrationGithubMapInput interface {
	pulumi.Input

	ToIntegrationGithubMapOutput() IntegrationGithubMapOutput
	ToIntegrationGithubMapOutputWithContext(context.Context) IntegrationGithubMapOutput
}

type IntegrationGithubMap map[string]IntegrationGithubInput

func (IntegrationGithubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationGithub)(nil)).Elem()
}

func (i IntegrationGithubMap) ToIntegrationGithubMapOutput() IntegrationGithubMapOutput {
	return i.ToIntegrationGithubMapOutputWithContext(context.Background())
}

func (i IntegrationGithubMap) ToIntegrationGithubMapOutputWithContext(ctx context.Context) IntegrationGithubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationGithubMapOutput)
}

type IntegrationGithubOutput struct{ *pulumi.OutputState }

func (IntegrationGithubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationGithub)(nil)).Elem()
}

func (o IntegrationGithubOutput) ToIntegrationGithubOutput() IntegrationGithubOutput {
	return o
}

func (o IntegrationGithubOutput) ToIntegrationGithubOutputWithContext(ctx context.Context) IntegrationGithubOutput {
	return o
}

// Whether the integration is active.
func (o IntegrationGithubOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationGithub) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// Create time.
func (o IntegrationGithubOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationGithub) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// ID of the project you want to activate integration on.
func (o IntegrationGithubOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationGithub) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
func (o IntegrationGithubOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationGithub) pulumi.StringOutput { return v.RepositoryUrl }).(pulumi.StringOutput)
}

// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
func (o IntegrationGithubOutput) StaticContext() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationGithub) pulumi.BoolPtrOutput { return v.StaticContext }).(pulumi.BoolPtrOutput)
}

// Title.
func (o IntegrationGithubOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationGithub) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// A GitHub personal access token with at least `repo:status` scope.
func (o IntegrationGithubOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationGithub) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// Update time.
func (o IntegrationGithubOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationGithub) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type IntegrationGithubArrayOutput struct{ *pulumi.OutputState }

func (IntegrationGithubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationGithub)(nil)).Elem()
}

func (o IntegrationGithubArrayOutput) ToIntegrationGithubArrayOutput() IntegrationGithubArrayOutput {
	return o
}

func (o IntegrationGithubArrayOutput) ToIntegrationGithubArrayOutputWithContext(ctx context.Context) IntegrationGithubArrayOutput {
	return o
}

func (o IntegrationGithubArrayOutput) Index(i pulumi.IntInput) IntegrationGithubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationGithub {
		return vs[0].([]*IntegrationGithub)[vs[1].(int)]
	}).(IntegrationGithubOutput)
}

type IntegrationGithubMapOutput struct{ *pulumi.OutputState }

func (IntegrationGithubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationGithub)(nil)).Elem()
}

func (o IntegrationGithubMapOutput) ToIntegrationGithubMapOutput() IntegrationGithubMapOutput {
	return o
}

func (o IntegrationGithubMapOutput) ToIntegrationGithubMapOutputWithContext(ctx context.Context) IntegrationGithubMapOutput {
	return o
}

func (o IntegrationGithubMapOutput) MapIndex(k pulumi.StringInput) IntegrationGithubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationGithub {
		return vs[0].(map[string]*IntegrationGithub)[vs[1].(string)]
	}).(IntegrationGithubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationGithubInput)(nil)).Elem(), &IntegrationGithub{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationGithubArrayInput)(nil)).Elem(), IntegrationGithubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationGithubMapInput)(nil)).Elem(), IntegrationGithubMap{})
	pulumi.RegisterOutputType(IntegrationGithubOutput{})
	pulumi.RegisterOutputType(IntegrationGithubArrayOutput{})
	pulumi.RegisterOutputType(IntegrationGithubMapOutput{})
}
