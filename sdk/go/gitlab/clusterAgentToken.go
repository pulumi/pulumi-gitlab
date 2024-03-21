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

// The `ClusterAgentToken` resource allows to manage the lifecycle of a token for a GitLab Agent for Kubernetes.
//
// > Requires at least maintainer permissions on the project.
//
// > Requires at least GitLab 15.0
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/cluster_agents.html#create-an-agent-token)
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
//	"github.com/pulumi/pulumi-helm/sdk/v1/go/helm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Create token for an agent
//			_, err := gitlab.NewClusterAgentToken(ctx, "example", &gitlab.ClusterAgentTokenArgs{
//				Project:     pulumi.String("12345"),
//				AgentId:     pulumi.Int(42),
//				Description: pulumi.String("some token"),
//			})
//			if err != nil {
//				return err
//			}
//			thisProject, err := gitlab.LookupProject(ctx, &gitlab.LookupProjectArgs{
//				PathWithNamespace: pulumi.StringRef("my-org/example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			thisClusterAgent, err := gitlab.NewClusterAgent(ctx, "thisClusterAgent", &gitlab.ClusterAgentArgs{
//				Project: pulumi.String(thisProject.Id),
//			})
//			if err != nil {
//				return err
//			}
//			thisClusterAgentToken, err := gitlab.NewClusterAgentToken(ctx, "thisClusterAgentToken", &gitlab.ClusterAgentTokenArgs{
//				Project:     pulumi.String(thisProject.Id),
//				AgentId:     thisClusterAgent.AgentId,
//				Description: pulumi.String("Token for the my-agent used with `gitlab-agent` Helm Chart"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = index.NewHelm_release(ctx, "gitlabAgent", &index.Helm_releaseArgs{
//				Name:            "gitlab-agent",
//				Namespace:       "gitlab-agent",
//				CreateNamespace: true,
//				Repository:      "https://charts.gitlab.io",
//				Chart:           "gitlab-agent",
//				Version:         "1.2.0",
//				Set: []map[string]interface{}{
//					map[string]interface{}{
//						"name":  "config.token",
//						"value": thisClusterAgentToken.Token,
//					},
//				},
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
// A token for a GitLab Agent for Kubernetes can be imported with the following command and the id pattern `<project>:<agent-id>:<token-id>`:
//
// ```sh
// $ pulumi import gitlab:index/clusterAgentToken:ClusterAgentToken example '12345:42:1'
// ```
//
// ATTENTION: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
type ClusterAgentToken struct {
	pulumi.CustomResourceState

	// The ID of the agent.
	AgentId pulumi.IntOutput `pulumi:"agentId"`
	// The ISO8601 datetime when the agent was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the user who created the agent.
	CreatedByUserId pulumi.IntOutput `pulumi:"createdByUserId"`
	// The Description for the agent.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ISO8601 datetime when the token was last used.
	LastUsedAt pulumi.StringOutput `pulumi:"lastUsedAt"`
	// The Name of the agent.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID or full path of the project maintained by the authenticated user.
	Project pulumi.StringOutput `pulumi:"project"`
	// The status of the token. Valid values are `active`, `revoked`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The secret token for the agent. The `token` is not available in imported resources.
	Token pulumi.StringOutput `pulumi:"token"`
	// The ID of the token.
	TokenId pulumi.IntOutput `pulumi:"tokenId"`
}

// NewClusterAgentToken registers a new resource with the given unique name, arguments, and options.
func NewClusterAgentToken(ctx *pulumi.Context,
	name string, args *ClusterAgentTokenArgs, opts ...pulumi.ResourceOption) (*ClusterAgentToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterAgentToken
	err := ctx.RegisterResource("gitlab:index/clusterAgentToken:ClusterAgentToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterAgentToken gets an existing ClusterAgentToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterAgentToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterAgentTokenState, opts ...pulumi.ResourceOption) (*ClusterAgentToken, error) {
	var resource ClusterAgentToken
	err := ctx.ReadResource("gitlab:index/clusterAgentToken:ClusterAgentToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterAgentToken resources.
type clusterAgentTokenState struct {
	// The ID of the agent.
	AgentId *int `pulumi:"agentId"`
	// The ISO8601 datetime when the agent was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the user who created the agent.
	CreatedByUserId *int `pulumi:"createdByUserId"`
	// The Description for the agent.
	Description *string `pulumi:"description"`
	// The ISO8601 datetime when the token was last used.
	LastUsedAt *string `pulumi:"lastUsedAt"`
	// The Name of the agent.
	Name *string `pulumi:"name"`
	// ID or full path of the project maintained by the authenticated user.
	Project *string `pulumi:"project"`
	// The status of the token. Valid values are `active`, `revoked`.
	Status *string `pulumi:"status"`
	// The secret token for the agent. The `token` is not available in imported resources.
	Token *string `pulumi:"token"`
	// The ID of the token.
	TokenId *int `pulumi:"tokenId"`
}

type ClusterAgentTokenState struct {
	// The ID of the agent.
	AgentId pulumi.IntPtrInput
	// The ISO8601 datetime when the agent was created.
	CreatedAt pulumi.StringPtrInput
	// The ID of the user who created the agent.
	CreatedByUserId pulumi.IntPtrInput
	// The Description for the agent.
	Description pulumi.StringPtrInput
	// The ISO8601 datetime when the token was last used.
	LastUsedAt pulumi.StringPtrInput
	// The Name of the agent.
	Name pulumi.StringPtrInput
	// ID or full path of the project maintained by the authenticated user.
	Project pulumi.StringPtrInput
	// The status of the token. Valid values are `active`, `revoked`.
	Status pulumi.StringPtrInput
	// The secret token for the agent. The `token` is not available in imported resources.
	Token pulumi.StringPtrInput
	// The ID of the token.
	TokenId pulumi.IntPtrInput
}

func (ClusterAgentTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAgentTokenState)(nil)).Elem()
}

type clusterAgentTokenArgs struct {
	// The ID of the agent.
	AgentId int `pulumi:"agentId"`
	// The Description for the agent.
	Description *string `pulumi:"description"`
	// The Name of the agent.
	Name *string `pulumi:"name"`
	// ID or full path of the project maintained by the authenticated user.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ClusterAgentToken resource.
type ClusterAgentTokenArgs struct {
	// The ID of the agent.
	AgentId pulumi.IntInput
	// The Description for the agent.
	Description pulumi.StringPtrInput
	// The Name of the agent.
	Name pulumi.StringPtrInput
	// ID or full path of the project maintained by the authenticated user.
	Project pulumi.StringInput
}

func (ClusterAgentTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAgentTokenArgs)(nil)).Elem()
}

type ClusterAgentTokenInput interface {
	pulumi.Input

	ToClusterAgentTokenOutput() ClusterAgentTokenOutput
	ToClusterAgentTokenOutputWithContext(ctx context.Context) ClusterAgentTokenOutput
}

func (*ClusterAgentToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAgentToken)(nil)).Elem()
}

func (i *ClusterAgentToken) ToClusterAgentTokenOutput() ClusterAgentTokenOutput {
	return i.ToClusterAgentTokenOutputWithContext(context.Background())
}

func (i *ClusterAgentToken) ToClusterAgentTokenOutputWithContext(ctx context.Context) ClusterAgentTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAgentTokenOutput)
}

// ClusterAgentTokenArrayInput is an input type that accepts ClusterAgentTokenArray and ClusterAgentTokenArrayOutput values.
// You can construct a concrete instance of `ClusterAgentTokenArrayInput` via:
//
//	ClusterAgentTokenArray{ ClusterAgentTokenArgs{...} }
type ClusterAgentTokenArrayInput interface {
	pulumi.Input

	ToClusterAgentTokenArrayOutput() ClusterAgentTokenArrayOutput
	ToClusterAgentTokenArrayOutputWithContext(context.Context) ClusterAgentTokenArrayOutput
}

type ClusterAgentTokenArray []ClusterAgentTokenInput

func (ClusterAgentTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAgentToken)(nil)).Elem()
}

func (i ClusterAgentTokenArray) ToClusterAgentTokenArrayOutput() ClusterAgentTokenArrayOutput {
	return i.ToClusterAgentTokenArrayOutputWithContext(context.Background())
}

func (i ClusterAgentTokenArray) ToClusterAgentTokenArrayOutputWithContext(ctx context.Context) ClusterAgentTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAgentTokenArrayOutput)
}

// ClusterAgentTokenMapInput is an input type that accepts ClusterAgentTokenMap and ClusterAgentTokenMapOutput values.
// You can construct a concrete instance of `ClusterAgentTokenMapInput` via:
//
//	ClusterAgentTokenMap{ "key": ClusterAgentTokenArgs{...} }
type ClusterAgentTokenMapInput interface {
	pulumi.Input

	ToClusterAgentTokenMapOutput() ClusterAgentTokenMapOutput
	ToClusterAgentTokenMapOutputWithContext(context.Context) ClusterAgentTokenMapOutput
}

type ClusterAgentTokenMap map[string]ClusterAgentTokenInput

func (ClusterAgentTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAgentToken)(nil)).Elem()
}

func (i ClusterAgentTokenMap) ToClusterAgentTokenMapOutput() ClusterAgentTokenMapOutput {
	return i.ToClusterAgentTokenMapOutputWithContext(context.Background())
}

func (i ClusterAgentTokenMap) ToClusterAgentTokenMapOutputWithContext(ctx context.Context) ClusterAgentTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAgentTokenMapOutput)
}

type ClusterAgentTokenOutput struct{ *pulumi.OutputState }

func (ClusterAgentTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAgentToken)(nil)).Elem()
}

func (o ClusterAgentTokenOutput) ToClusterAgentTokenOutput() ClusterAgentTokenOutput {
	return o
}

func (o ClusterAgentTokenOutput) ToClusterAgentTokenOutputWithContext(ctx context.Context) ClusterAgentTokenOutput {
	return o
}

// The ID of the agent.
func (o ClusterAgentTokenOutput) AgentId() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.IntOutput { return v.AgentId }).(pulumi.IntOutput)
}

// The ISO8601 datetime when the agent was created.
func (o ClusterAgentTokenOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the user who created the agent.
func (o ClusterAgentTokenOutput) CreatedByUserId() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.IntOutput { return v.CreatedByUserId }).(pulumi.IntOutput)
}

// The Description for the agent.
func (o ClusterAgentTokenOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ISO8601 datetime when the token was last used.
func (o ClusterAgentTokenOutput) LastUsedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.LastUsedAt }).(pulumi.StringOutput)
}

// The Name of the agent.
func (o ClusterAgentTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID or full path of the project maintained by the authenticated user.
func (o ClusterAgentTokenOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The status of the token. Valid values are `active`, `revoked`.
func (o ClusterAgentTokenOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The secret token for the agent. The `token` is not available in imported resources.
func (o ClusterAgentTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The ID of the token.
func (o ClusterAgentTokenOutput) TokenId() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.IntOutput { return v.TokenId }).(pulumi.IntOutput)
}

type ClusterAgentTokenArrayOutput struct{ *pulumi.OutputState }

func (ClusterAgentTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAgentToken)(nil)).Elem()
}

func (o ClusterAgentTokenArrayOutput) ToClusterAgentTokenArrayOutput() ClusterAgentTokenArrayOutput {
	return o
}

func (o ClusterAgentTokenArrayOutput) ToClusterAgentTokenArrayOutputWithContext(ctx context.Context) ClusterAgentTokenArrayOutput {
	return o
}

func (o ClusterAgentTokenArrayOutput) Index(i pulumi.IntInput) ClusterAgentTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterAgentToken {
		return vs[0].([]*ClusterAgentToken)[vs[1].(int)]
	}).(ClusterAgentTokenOutput)
}

type ClusterAgentTokenMapOutput struct{ *pulumi.OutputState }

func (ClusterAgentTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAgentToken)(nil)).Elem()
}

func (o ClusterAgentTokenMapOutput) ToClusterAgentTokenMapOutput() ClusterAgentTokenMapOutput {
	return o
}

func (o ClusterAgentTokenMapOutput) ToClusterAgentTokenMapOutputWithContext(ctx context.Context) ClusterAgentTokenMapOutput {
	return o
}

func (o ClusterAgentTokenMapOutput) MapIndex(k pulumi.StringInput) ClusterAgentTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterAgentToken {
		return vs[0].(map[string]*ClusterAgentToken)[vs[1].(string)]
	}).(ClusterAgentTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAgentTokenInput)(nil)).Elem(), &ClusterAgentToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAgentTokenArrayInput)(nil)).Elem(), ClusterAgentTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAgentTokenMapInput)(nil)).Elem(), ClusterAgentTokenMap{})
	pulumi.RegisterOutputType(ClusterAgentTokenOutput{})
	pulumi.RegisterOutputType(ClusterAgentTokenArrayOutput{})
	pulumi.RegisterOutputType(ClusterAgentTokenMapOutput{})
}
