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

// The `ClusterAgent` resource allows to manage the lifecycle of a GitLab Agent for Kubernetes.
//
// > Note that this resource only registers the agent, but doesn't configure it.
//
//	The configuration needs to be manually added as described in
//	[the docs](https://docs.gitlab.com/ee/user/clusters/agent/install/index.html#create-an-agent-configuration-file).
//	However, a `RepositoryFile` resource may be used to achieve that.
//
// > Requires at least maintainer permissions on the project.
//
// > Requires at least GitLab 14.10
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/cluster_agents.html)
//
// ## Import
//
// GitLab Agent for Kubernetes can be imported with the following command and the id pattern `<project>:<agent-id>`
//
// ```sh
//
//	$ pulumi import gitlab:index/clusterAgent:ClusterAgent example '12345:42'
//
// ```
type ClusterAgent struct {
	pulumi.CustomResourceState

	// The ID of the agent.
	AgentId pulumi.IntOutput `pulumi:"agentId"`
	// The ISO8601 datetime when the agent was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the user who created the agent.
	CreatedByUserId pulumi.IntOutput `pulumi:"createdByUserId"`
	// The Name of the agent.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID or full path of the project maintained by the authenticated user.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewClusterAgent registers a new resource with the given unique name, arguments, and options.
func NewClusterAgent(ctx *pulumi.Context,
	name string, args *ClusterAgentArgs, opts ...pulumi.ResourceOption) (*ClusterAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterAgent
	err := ctx.RegisterResource("gitlab:index/clusterAgent:ClusterAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterAgent gets an existing ClusterAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterAgentState, opts ...pulumi.ResourceOption) (*ClusterAgent, error) {
	var resource ClusterAgent
	err := ctx.ReadResource("gitlab:index/clusterAgent:ClusterAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterAgent resources.
type clusterAgentState struct {
	// The ID of the agent.
	AgentId *int `pulumi:"agentId"`
	// The ISO8601 datetime when the agent was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the user who created the agent.
	CreatedByUserId *int `pulumi:"createdByUserId"`
	// The Name of the agent.
	Name *string `pulumi:"name"`
	// ID or full path of the project maintained by the authenticated user.
	Project *string `pulumi:"project"`
}

type ClusterAgentState struct {
	// The ID of the agent.
	AgentId pulumi.IntPtrInput
	// The ISO8601 datetime when the agent was created.
	CreatedAt pulumi.StringPtrInput
	// The ID of the user who created the agent.
	CreatedByUserId pulumi.IntPtrInput
	// The Name of the agent.
	Name pulumi.StringPtrInput
	// ID or full path of the project maintained by the authenticated user.
	Project pulumi.StringPtrInput
}

func (ClusterAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAgentState)(nil)).Elem()
}

type clusterAgentArgs struct {
	// The Name of the agent.
	Name *string `pulumi:"name"`
	// ID or full path of the project maintained by the authenticated user.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ClusterAgent resource.
type ClusterAgentArgs struct {
	// The Name of the agent.
	Name pulumi.StringPtrInput
	// ID or full path of the project maintained by the authenticated user.
	Project pulumi.StringInput
}

func (ClusterAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAgentArgs)(nil)).Elem()
}

type ClusterAgentInput interface {
	pulumi.Input

	ToClusterAgentOutput() ClusterAgentOutput
	ToClusterAgentOutputWithContext(ctx context.Context) ClusterAgentOutput
}

func (*ClusterAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAgent)(nil)).Elem()
}

func (i *ClusterAgent) ToClusterAgentOutput() ClusterAgentOutput {
	return i.ToClusterAgentOutputWithContext(context.Background())
}

func (i *ClusterAgent) ToClusterAgentOutputWithContext(ctx context.Context) ClusterAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAgentOutput)
}

func (i *ClusterAgent) ToOutput(ctx context.Context) pulumix.Output[*ClusterAgent] {
	return pulumix.Output[*ClusterAgent]{
		OutputState: i.ToClusterAgentOutputWithContext(ctx).OutputState,
	}
}

// ClusterAgentArrayInput is an input type that accepts ClusterAgentArray and ClusterAgentArrayOutput values.
// You can construct a concrete instance of `ClusterAgentArrayInput` via:
//
//	ClusterAgentArray{ ClusterAgentArgs{...} }
type ClusterAgentArrayInput interface {
	pulumi.Input

	ToClusterAgentArrayOutput() ClusterAgentArrayOutput
	ToClusterAgentArrayOutputWithContext(context.Context) ClusterAgentArrayOutput
}

type ClusterAgentArray []ClusterAgentInput

func (ClusterAgentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAgent)(nil)).Elem()
}

func (i ClusterAgentArray) ToClusterAgentArrayOutput() ClusterAgentArrayOutput {
	return i.ToClusterAgentArrayOutputWithContext(context.Background())
}

func (i ClusterAgentArray) ToClusterAgentArrayOutputWithContext(ctx context.Context) ClusterAgentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAgentArrayOutput)
}

func (i ClusterAgentArray) ToOutput(ctx context.Context) pulumix.Output[[]*ClusterAgent] {
	return pulumix.Output[[]*ClusterAgent]{
		OutputState: i.ToClusterAgentArrayOutputWithContext(ctx).OutputState,
	}
}

// ClusterAgentMapInput is an input type that accepts ClusterAgentMap and ClusterAgentMapOutput values.
// You can construct a concrete instance of `ClusterAgentMapInput` via:
//
//	ClusterAgentMap{ "key": ClusterAgentArgs{...} }
type ClusterAgentMapInput interface {
	pulumi.Input

	ToClusterAgentMapOutput() ClusterAgentMapOutput
	ToClusterAgentMapOutputWithContext(context.Context) ClusterAgentMapOutput
}

type ClusterAgentMap map[string]ClusterAgentInput

func (ClusterAgentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAgent)(nil)).Elem()
}

func (i ClusterAgentMap) ToClusterAgentMapOutput() ClusterAgentMapOutput {
	return i.ToClusterAgentMapOutputWithContext(context.Background())
}

func (i ClusterAgentMap) ToClusterAgentMapOutputWithContext(ctx context.Context) ClusterAgentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAgentMapOutput)
}

func (i ClusterAgentMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ClusterAgent] {
	return pulumix.Output[map[string]*ClusterAgent]{
		OutputState: i.ToClusterAgentMapOutputWithContext(ctx).OutputState,
	}
}

type ClusterAgentOutput struct{ *pulumi.OutputState }

func (ClusterAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAgent)(nil)).Elem()
}

func (o ClusterAgentOutput) ToClusterAgentOutput() ClusterAgentOutput {
	return o
}

func (o ClusterAgentOutput) ToClusterAgentOutputWithContext(ctx context.Context) ClusterAgentOutput {
	return o
}

func (o ClusterAgentOutput) ToOutput(ctx context.Context) pulumix.Output[*ClusterAgent] {
	return pulumix.Output[*ClusterAgent]{
		OutputState: o.OutputState,
	}
}

// The ID of the agent.
func (o ClusterAgentOutput) AgentId() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterAgent) pulumi.IntOutput { return v.AgentId }).(pulumi.IntOutput)
}

// The ISO8601 datetime when the agent was created.
func (o ClusterAgentOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgent) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the user who created the agent.
func (o ClusterAgentOutput) CreatedByUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterAgent) pulumi.IntOutput { return v.CreatedByUserId }).(pulumi.IntOutput)
}

// The Name of the agent.
func (o ClusterAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID or full path of the project maintained by the authenticated user.
func (o ClusterAgentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgent) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ClusterAgentArrayOutput struct{ *pulumi.OutputState }

func (ClusterAgentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAgent)(nil)).Elem()
}

func (o ClusterAgentArrayOutput) ToClusterAgentArrayOutput() ClusterAgentArrayOutput {
	return o
}

func (o ClusterAgentArrayOutput) ToClusterAgentArrayOutputWithContext(ctx context.Context) ClusterAgentArrayOutput {
	return o
}

func (o ClusterAgentArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ClusterAgent] {
	return pulumix.Output[[]*ClusterAgent]{
		OutputState: o.OutputState,
	}
}

func (o ClusterAgentArrayOutput) Index(i pulumi.IntInput) ClusterAgentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterAgent {
		return vs[0].([]*ClusterAgent)[vs[1].(int)]
	}).(ClusterAgentOutput)
}

type ClusterAgentMapOutput struct{ *pulumi.OutputState }

func (ClusterAgentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAgent)(nil)).Elem()
}

func (o ClusterAgentMapOutput) ToClusterAgentMapOutput() ClusterAgentMapOutput {
	return o
}

func (o ClusterAgentMapOutput) ToClusterAgentMapOutputWithContext(ctx context.Context) ClusterAgentMapOutput {
	return o
}

func (o ClusterAgentMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ClusterAgent] {
	return pulumix.Output[map[string]*ClusterAgent]{
		OutputState: o.OutputState,
	}
}

func (o ClusterAgentMapOutput) MapIndex(k pulumi.StringInput) ClusterAgentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterAgent {
		return vs[0].(map[string]*ClusterAgent)[vs[1].(string)]
	}).(ClusterAgentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAgentInput)(nil)).Elem(), &ClusterAgent{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAgentArrayInput)(nil)).Elem(), ClusterAgentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAgentMapInput)(nil)).Elem(), ClusterAgentMap{})
	pulumi.RegisterOutputType(ClusterAgentOutput{})
	pulumi.RegisterOutputType(ClusterAgentArrayOutput{})
	pulumi.RegisterOutputType(ClusterAgentMapOutput{})
}
