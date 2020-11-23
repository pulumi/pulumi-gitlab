// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ServiceGithub struct {
	pulumi.CustomResourceState

	Active    pulumi.BoolOutput   `pulumi:"active"`
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// ID of the project you want to activate integration on.
	Project       pulumi.StringOutput `pulumi:"project"`
	RepositoryUrl pulumi.StringOutput `pulumi:"repositoryUrl"`
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as _required_ in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext pulumi.BoolPtrOutput `pulumi:"staticContext"`
	Title         pulumi.StringOutput  `pulumi:"title"`
	// A GitHub personal access token with at least `repo:status` scope.
	Token     pulumi.StringOutput `pulumi:"token"`
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewServiceGithub registers a new resource with the given unique name, arguments, and options.
func NewServiceGithub(ctx *pulumi.Context,
	name string, args *ServiceGithubArgs, opts ...pulumi.ResourceOption) (*ServiceGithub, error) {
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil || args.RepositoryUrl == nil {
		return nil, errors.New("missing required argument 'RepositoryUrl'")
	}
	if args == nil || args.Token == nil {
		return nil, errors.New("missing required argument 'Token'")
	}
	if args == nil {
		args = &ServiceGithubArgs{}
	}
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
	Active    *bool   `pulumi:"active"`
	CreatedAt *string `pulumi:"createdAt"`
	// ID of the project you want to activate integration on.
	Project       *string `pulumi:"project"`
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as _required_ in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext *bool   `pulumi:"staticContext"`
	Title         *string `pulumi:"title"`
	// A GitHub personal access token with at least `repo:status` scope.
	Token     *string `pulumi:"token"`
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ServiceGithubState struct {
	Active    pulumi.BoolPtrInput
	CreatedAt pulumi.StringPtrInput
	// ID of the project you want to activate integration on.
	Project       pulumi.StringPtrInput
	RepositoryUrl pulumi.StringPtrInput
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as _required_ in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext pulumi.BoolPtrInput
	Title         pulumi.StringPtrInput
	// A GitHub personal access token with at least `repo:status` scope.
	Token     pulumi.StringPtrInput
	UpdatedAt pulumi.StringPtrInput
}

func (ServiceGithubState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceGithubState)(nil)).Elem()
}

type serviceGithubArgs struct {
	// ID of the project you want to activate integration on.
	Project       string `pulumi:"project"`
	RepositoryUrl string `pulumi:"repositoryUrl"`
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as _required_ in GitHub. See [Static / dynamic status check names] to learn more.
	StaticContext *bool `pulumi:"staticContext"`
	// A GitHub personal access token with at least `repo:status` scope.
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a ServiceGithub resource.
type ServiceGithubArgs struct {
	// ID of the project you want to activate integration on.
	Project       pulumi.StringInput
	RepositoryUrl pulumi.StringInput
	// Append instance name instead of branch to the status. Must enable to set a GitLab status check as _required_ in GitHub. See [Static / dynamic status check names] to learn more.
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

func (ServiceGithub) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceGithub)(nil)).Elem()
}

func (i ServiceGithub) ToServiceGithubOutput() ServiceGithubOutput {
	return i.ToServiceGithubOutputWithContext(context.Background())
}

func (i ServiceGithub) ToServiceGithubOutputWithContext(ctx context.Context) ServiceGithubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGithubOutput)
}

type ServiceGithubOutput struct {
	*pulumi.OutputState
}

func (ServiceGithubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceGithubOutput)(nil)).Elem()
}

func (o ServiceGithubOutput) ToServiceGithubOutput() ServiceGithubOutput {
	return o
}

func (o ServiceGithubOutput) ToServiceGithubOutputWithContext(ctx context.Context) ServiceGithubOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceGithubOutput{})
}
