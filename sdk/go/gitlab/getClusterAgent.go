// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ClusterAgent` data source allows to retrieve details about a GitLab Agent for Kubernetes.
//
// > Requires at least GitLab 14.10
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/cluster_agents.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.LookupClusterAgent(ctx, &gitlab.LookupClusterAgentArgs{
//				Project: "12345",
//				AgentId: 1,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupClusterAgent(ctx *pulumi.Context, args *LookupClusterAgentArgs, opts ...pulumi.InvokeOption) (*LookupClusterAgentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterAgentResult
	err := ctx.Invoke("gitlab:index/getClusterAgent:getClusterAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterAgent.
type LookupClusterAgentArgs struct {
	// The ID of the agent.
	AgentId int `pulumi:"agentId"`
	// ID or full path of the project maintained by the authenticated user.
	Project string `pulumi:"project"`
}

// A collection of values returned by getClusterAgent.
type LookupClusterAgentResult struct {
	// The ID of the agent.
	AgentId int `pulumi:"agentId"`
	// The ISO8601 datetime when the agent was created.
	CreatedAt string `pulumi:"createdAt"`
	// The ID of the user who created the agent.
	CreatedByUserId int `pulumi:"createdByUserId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Name of the agent.
	Name string `pulumi:"name"`
	// ID or full path of the project maintained by the authenticated user.
	Project string `pulumi:"project"`
}

func LookupClusterAgentOutput(ctx *pulumi.Context, args LookupClusterAgentOutputArgs, opts ...pulumi.InvokeOption) LookupClusterAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterAgentResultOutput, error) {
			args := v.(LookupClusterAgentArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupClusterAgentResult
			secret, err := ctx.InvokePackageRaw("gitlab:index/getClusterAgent:getClusterAgent", args, &rv, "", opts...)
			if err != nil {
				return LookupClusterAgentResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupClusterAgentResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupClusterAgentResultOutput), nil
			}
			return output, nil
		}).(LookupClusterAgentResultOutput)
}

// A collection of arguments for invoking getClusterAgent.
type LookupClusterAgentOutputArgs struct {
	// The ID of the agent.
	AgentId pulumi.IntInput `pulumi:"agentId"`
	// ID or full path of the project maintained by the authenticated user.
	Project pulumi.StringInput `pulumi:"project"`
}

func (LookupClusterAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterAgentArgs)(nil)).Elem()
}

// A collection of values returned by getClusterAgent.
type LookupClusterAgentResultOutput struct{ *pulumi.OutputState }

func (LookupClusterAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterAgentResult)(nil)).Elem()
}

func (o LookupClusterAgentResultOutput) ToLookupClusterAgentResultOutput() LookupClusterAgentResultOutput {
	return o
}

func (o LookupClusterAgentResultOutput) ToLookupClusterAgentResultOutputWithContext(ctx context.Context) LookupClusterAgentResultOutput {
	return o
}

// The ID of the agent.
func (o LookupClusterAgentResultOutput) AgentId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterAgentResult) int { return v.AgentId }).(pulumi.IntOutput)
}

// The ISO8601 datetime when the agent was created.
func (o LookupClusterAgentResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterAgentResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the user who created the agent.
func (o LookupClusterAgentResultOutput) CreatedByUserId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterAgentResult) int { return v.CreatedByUserId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClusterAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Name of the agent.
func (o LookupClusterAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID or full path of the project maintained by the authenticated user.
func (o LookupClusterAgentResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterAgentResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterAgentResultOutput{})
}
