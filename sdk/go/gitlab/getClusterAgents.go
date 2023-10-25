// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The `getClusterAgents` data source allows details of GitLab Agents for Kubernetes in a project.
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
//	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.GetClusterAgents(ctx, &gitlab.GetClusterAgentsArgs{
//				Project: "12345",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetClusterAgents(ctx *pulumi.Context, args *GetClusterAgentsArgs, opts ...pulumi.InvokeOption) (*GetClusterAgentsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetClusterAgentsResult
	err := ctx.Invoke("gitlab:index/getClusterAgents:getClusterAgents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterAgents.
type GetClusterAgentsArgs struct {
	Project string `pulumi:"project"`
}

// A collection of values returned by getClusterAgents.
type GetClusterAgentsResult struct {
	// List of the registered agents.
	ClusterAgents []GetClusterAgentsClusterAgent `pulumi:"clusterAgents"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID or full path of the project owned by the authenticated user.
	Project string `pulumi:"project"`
}

func GetClusterAgentsOutput(ctx *pulumi.Context, args GetClusterAgentsOutputArgs, opts ...pulumi.InvokeOption) GetClusterAgentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClusterAgentsResult, error) {
			args := v.(GetClusterAgentsArgs)
			r, err := GetClusterAgents(ctx, &args, opts...)
			var s GetClusterAgentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClusterAgentsResultOutput)
}

// A collection of arguments for invoking getClusterAgents.
type GetClusterAgentsOutputArgs struct {
	Project pulumi.StringInput `pulumi:"project"`
}

func (GetClusterAgentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterAgentsArgs)(nil)).Elem()
}

// A collection of values returned by getClusterAgents.
type GetClusterAgentsResultOutput struct{ *pulumi.OutputState }

func (GetClusterAgentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterAgentsResult)(nil)).Elem()
}

func (o GetClusterAgentsResultOutput) ToGetClusterAgentsResultOutput() GetClusterAgentsResultOutput {
	return o
}

func (o GetClusterAgentsResultOutput) ToGetClusterAgentsResultOutputWithContext(ctx context.Context) GetClusterAgentsResultOutput {
	return o
}

func (o GetClusterAgentsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetClusterAgentsResult] {
	return pulumix.Output[GetClusterAgentsResult]{
		OutputState: o.OutputState,
	}
}

// List of the registered agents.
func (o GetClusterAgentsResultOutput) ClusterAgents() GetClusterAgentsClusterAgentArrayOutput {
	return o.ApplyT(func(v GetClusterAgentsResult) []GetClusterAgentsClusterAgent { return v.ClusterAgents }).(GetClusterAgentsClusterAgentArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClusterAgentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterAgentsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID or full path of the project owned by the authenticated user.
func (o GetClusterAgentsResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterAgentsResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterAgentsResultOutput{})
}
