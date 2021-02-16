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

func (*ServiceGithub) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceGithub)(nil))
}

func (i *ServiceGithub) ToServiceGithubOutput() ServiceGithubOutput {
	return i.ToServiceGithubOutputWithContext(context.Background())
}

func (i *ServiceGithub) ToServiceGithubOutputWithContext(ctx context.Context) ServiceGithubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGithubOutput)
}

func (i *ServiceGithub) ToServiceGithubPtrOutput() ServiceGithubPtrOutput {
	return i.ToServiceGithubPtrOutputWithContext(context.Background())
}

func (i *ServiceGithub) ToServiceGithubPtrOutputWithContext(ctx context.Context) ServiceGithubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGithubPtrOutput)
}

type ServiceGithubPtrInput interface {
	pulumi.Input

	ToServiceGithubPtrOutput() ServiceGithubPtrOutput
	ToServiceGithubPtrOutputWithContext(ctx context.Context) ServiceGithubPtrOutput
}

type serviceGithubPtrType ServiceGithubArgs

func (*serviceGithubPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceGithub)(nil))
}

func (i *serviceGithubPtrType) ToServiceGithubPtrOutput() ServiceGithubPtrOutput {
	return i.ToServiceGithubPtrOutputWithContext(context.Background())
}

func (i *serviceGithubPtrType) ToServiceGithubPtrOutputWithContext(ctx context.Context) ServiceGithubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGithubPtrOutput)
}

// ServiceGithubArrayInput is an input type that accepts ServiceGithubArray and ServiceGithubArrayOutput values.
// You can construct a concrete instance of `ServiceGithubArrayInput` via:
//
//          ServiceGithubArray{ ServiceGithubArgs{...} }
type ServiceGithubArrayInput interface {
	pulumi.Input

	ToServiceGithubArrayOutput() ServiceGithubArrayOutput
	ToServiceGithubArrayOutputWithContext(context.Context) ServiceGithubArrayOutput
}

type ServiceGithubArray []ServiceGithubInput

func (ServiceGithubArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServiceGithub)(nil))
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
//          ServiceGithubMap{ "key": ServiceGithubArgs{...} }
type ServiceGithubMapInput interface {
	pulumi.Input

	ToServiceGithubMapOutput() ServiceGithubMapOutput
	ToServiceGithubMapOutputWithContext(context.Context) ServiceGithubMapOutput
}

type ServiceGithubMap map[string]ServiceGithubInput

func (ServiceGithubMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServiceGithub)(nil))
}

func (i ServiceGithubMap) ToServiceGithubMapOutput() ServiceGithubMapOutput {
	return i.ToServiceGithubMapOutputWithContext(context.Background())
}

func (i ServiceGithubMap) ToServiceGithubMapOutputWithContext(ctx context.Context) ServiceGithubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGithubMapOutput)
}

type ServiceGithubOutput struct {
	*pulumi.OutputState
}

func (ServiceGithubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceGithub)(nil))
}

func (o ServiceGithubOutput) ToServiceGithubOutput() ServiceGithubOutput {
	return o
}

func (o ServiceGithubOutput) ToServiceGithubOutputWithContext(ctx context.Context) ServiceGithubOutput {
	return o
}

func (o ServiceGithubOutput) ToServiceGithubPtrOutput() ServiceGithubPtrOutput {
	return o.ToServiceGithubPtrOutputWithContext(context.Background())
}

func (o ServiceGithubOutput) ToServiceGithubPtrOutputWithContext(ctx context.Context) ServiceGithubPtrOutput {
	return o.ApplyT(func(v ServiceGithub) *ServiceGithub {
		return &v
	}).(ServiceGithubPtrOutput)
}

type ServiceGithubPtrOutput struct {
	*pulumi.OutputState
}

func (ServiceGithubPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceGithub)(nil))
}

func (o ServiceGithubPtrOutput) ToServiceGithubPtrOutput() ServiceGithubPtrOutput {
	return o
}

func (o ServiceGithubPtrOutput) ToServiceGithubPtrOutputWithContext(ctx context.Context) ServiceGithubPtrOutput {
	return o
}

type ServiceGithubArrayOutput struct{ *pulumi.OutputState }

func (ServiceGithubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceGithub)(nil))
}

func (o ServiceGithubArrayOutput) ToServiceGithubArrayOutput() ServiceGithubArrayOutput {
	return o
}

func (o ServiceGithubArrayOutput) ToServiceGithubArrayOutputWithContext(ctx context.Context) ServiceGithubArrayOutput {
	return o
}

func (o ServiceGithubArrayOutput) Index(i pulumi.IntInput) ServiceGithubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceGithub {
		return vs[0].([]ServiceGithub)[vs[1].(int)]
	}).(ServiceGithubOutput)
}

type ServiceGithubMapOutput struct{ *pulumi.OutputState }

func (ServiceGithubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceGithub)(nil))
}

func (o ServiceGithubMapOutput) ToServiceGithubMapOutput() ServiceGithubMapOutput {
	return o
}

func (o ServiceGithubMapOutput) ToServiceGithubMapOutputWithContext(ctx context.Context) ServiceGithubMapOutput {
	return o
}

func (o ServiceGithubMapOutput) MapIndex(k pulumi.StringInput) ServiceGithubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceGithub {
		return vs[0].(map[string]ServiceGithub)[vs[1].(string)]
	}).(ServiceGithubOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceGithubOutput{})
	pulumi.RegisterOutputType(ServiceGithubPtrOutput{})
	pulumi.RegisterOutputType(ServiceGithubArrayOutput{})
	pulumi.RegisterOutputType(ServiceGithubMapOutput{})
}
