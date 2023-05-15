// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ServiceGithub` resource allows to manage the lifecycle of a project integration with GitHub.
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
//	"github.com/pulumi/pulumi-gitlab/sdk/v5/go/gitlab"
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
//			_, err = gitlab.NewServiceGithub(ctx, "github", &gitlab.ServiceGithubArgs{
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
//	$ pulumi import gitlab:index/serviceGithub:ServiceGithub You can import a service_github state using `<resource> <project_id>`
//
// ```
//
// ```sh
//
//	$ pulumi import gitlab:index/serviceGithub:ServiceGithub github 1
//
// ```
type ServiceGithub struct {
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

// NewServiceGithub registers a new resource with the given unique name, arguments, and options.
func NewServiceGithub(ctx *pulumi.Context,
	name string, args *ServiceGithubArgs, opts ...pulumi.ResourceOption) (*ServiceGithub, error) {
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
	var resource ServiceGithub
	err := ctx.RegisterResource("gitlab:index/serviceGithub:ServiceGithub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceGithub gets an existing ServiceGithub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceGithub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceGithubState, opts ...pulumi.ResourceOption) (*ServiceGithub, error) {
	var resource ServiceGithub
	err := ctx.ReadResource("gitlab:index/serviceGithub:ServiceGithub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceGithub resources.
type serviceGithubState struct {
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

type ServiceGithubState struct {
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

func (ServiceGithubState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceGithubState)(nil)).Elem()
}

type serviceGithubArgs struct {
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
	RepositoryUrl string `pulumi:"repositoryUrl"`
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext *bool `pulumi:"staticContext"`
	// A GitHub personal access token with at least `repo:status` scope.
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a ServiceGithub resource.
type ServiceGithubArgs struct {
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
	// The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
	RepositoryUrl pulumi.StringInput
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext pulumi.BoolPtrInput
	// A GitHub personal access token with at least `repo:status` scope.
	Token pulumi.StringInput
}

func (ServiceGithubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceGithubArgs)(nil)).Elem()
}

type ServiceGithubInput interface {
	pulumi.Input

	ToServiceGithubOutput() ServiceGithubOutput
	ToServiceGithubOutputWithContext(ctx context.Context) ServiceGithubOutput
}

func (*ServiceGithub) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceGithub)(nil)).Elem()
}

func (i *ServiceGithub) ToServiceGithubOutput() ServiceGithubOutput {
	return i.ToServiceGithubOutputWithContext(context.Background())
}

func (i *ServiceGithub) ToServiceGithubOutputWithContext(ctx context.Context) ServiceGithubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGithubOutput)
}

// ServiceGithubArrayInput is an input type that accepts ServiceGithubArray and ServiceGithubArrayOutput values.
// You can construct a concrete instance of `ServiceGithubArrayInput` via:
//
//	ServiceGithubArray{ ServiceGithubArgs{...} }
type ServiceGithubArrayInput interface {
	pulumi.Input

	ToServiceGithubArrayOutput() ServiceGithubArrayOutput
	ToServiceGithubArrayOutputWithContext(context.Context) ServiceGithubArrayOutput
}

type ServiceGithubArray []ServiceGithubInput

func (ServiceGithubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceGithub)(nil)).Elem()
}

func (i ServiceGithubArray) ToServiceGithubArrayOutput() ServiceGithubArrayOutput {
	return i.ToServiceGithubArrayOutputWithContext(context.Background())
}

func (i ServiceGithubArray) ToServiceGithubArrayOutputWithContext(ctx context.Context) ServiceGithubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGithubArrayOutput)
}

// ServiceGithubMapInput is an input type that accepts ServiceGithubMap and ServiceGithubMapOutput values.
// You can construct a concrete instance of `ServiceGithubMapInput` via:
//
//	ServiceGithubMap{ "key": ServiceGithubArgs{...} }
type ServiceGithubMapInput interface {
	pulumi.Input

	ToServiceGithubMapOutput() ServiceGithubMapOutput
	ToServiceGithubMapOutputWithContext(context.Context) ServiceGithubMapOutput
}

type ServiceGithubMap map[string]ServiceGithubInput

func (ServiceGithubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceGithub)(nil)).Elem()
}

func (i ServiceGithubMap) ToServiceGithubMapOutput() ServiceGithubMapOutput {
	return i.ToServiceGithubMapOutputWithContext(context.Background())
}

func (i ServiceGithubMap) ToServiceGithubMapOutputWithContext(ctx context.Context) ServiceGithubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGithubMapOutput)
}

type ServiceGithubOutput struct{ *pulumi.OutputState }

func (ServiceGithubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceGithub)(nil)).Elem()
}

func (o ServiceGithubOutput) ToServiceGithubOutput() ServiceGithubOutput {
	return o
}

func (o ServiceGithubOutput) ToServiceGithubOutputWithContext(ctx context.Context) ServiceGithubOutput {
	return o
}

// Whether the integration is active.
func (o ServiceGithubOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceGithub) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// Create time.
func (o ServiceGithubOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceGithub) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// ID of the project you want to activate integration on.
func (o ServiceGithubOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceGithub) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
func (o ServiceGithubOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceGithub) pulumi.StringOutput { return v.RepositoryUrl }).(pulumi.StringOutput)
}

// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
func (o ServiceGithubOutput) StaticContext() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceGithub) pulumi.BoolPtrOutput { return v.StaticContext }).(pulumi.BoolPtrOutput)
}

// Title.
func (o ServiceGithubOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceGithub) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// A GitHub personal access token with at least `repo:status` scope.
func (o ServiceGithubOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceGithub) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// Update time.
func (o ServiceGithubOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceGithub) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ServiceGithubArrayOutput struct{ *pulumi.OutputState }

func (ServiceGithubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceGithub)(nil)).Elem()
}

func (o ServiceGithubArrayOutput) ToServiceGithubArrayOutput() ServiceGithubArrayOutput {
	return o
}

func (o ServiceGithubArrayOutput) ToServiceGithubArrayOutputWithContext(ctx context.Context) ServiceGithubArrayOutput {
	return o
}

func (o ServiceGithubArrayOutput) Index(i pulumi.IntInput) ServiceGithubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceGithub {
		return vs[0].([]*ServiceGithub)[vs[1].(int)]
	}).(ServiceGithubOutput)
}

type ServiceGithubMapOutput struct{ *pulumi.OutputState }

func (ServiceGithubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceGithub)(nil)).Elem()
}

func (o ServiceGithubMapOutput) ToServiceGithubMapOutput() ServiceGithubMapOutput {
	return o
}

func (o ServiceGithubMapOutput) ToServiceGithubMapOutputWithContext(ctx context.Context) ServiceGithubMapOutput {
	return o
}

func (o ServiceGithubMapOutput) MapIndex(k pulumi.StringInput) ServiceGithubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceGithub {
		return vs[0].(map[string]*ServiceGithub)[vs[1].(string)]
	}).(ServiceGithubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGithubInput)(nil)).Elem(), &ServiceGithub{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGithubArrayInput)(nil)).Elem(), ServiceGithubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGithubMapInput)(nil)).Elem(), ServiceGithubMap{})
	pulumi.RegisterOutputType(ServiceGithubOutput{})
	pulumi.RegisterOutputType(ServiceGithubArrayOutput{})
	pulumi.RegisterOutputType(ServiceGithubMapOutput{})
}
