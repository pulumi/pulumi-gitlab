// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v7/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ComplianceFramework` data source allows details of a compliance framework to be retrieved by its name and the namespace it belongs to.
//
// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#querynamespace)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v7/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.LookupComplianceFramework(ctx, &gitlab.LookupComplianceFrameworkArgs{
//				NamespacePath: "top-level-group",
//				Name:          "HIPAA",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupComplianceFramework(ctx *pulumi.Context, args *LookupComplianceFrameworkArgs, opts ...pulumi.InvokeOption) (*LookupComplianceFrameworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupComplianceFrameworkResult
	err := ctx.Invoke("gitlab:index/getComplianceFramework:getComplianceFramework", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComplianceFramework.
type LookupComplianceFrameworkArgs struct {
	// Name for the compliance framework.
	Name string `pulumi:"name"`
	// Full path of the namespace to where the compliance framework is.
	NamespacePath string `pulumi:"namespacePath"`
}

// A collection of values returned by getComplianceFramework.
type LookupComplianceFrameworkResult struct {
	// Color representation of the compliance framework in hex format. e.g. #FCA121.
	Color string `pulumi:"color"`
	// Is the compliance framework the default framework for the group.
	Default bool `pulumi:"default"`
	// Description for the compliance framework.
	Description string `pulumi:"description"`
	// Globally unique ID of the compliance framework.
	FrameworkId string `pulumi:"frameworkId"`
	Id          string `pulumi:"id"`
	// Name for the compliance framework.
	Name string `pulumi:"name"`
	// Full path of the namespace to where the compliance framework is.
	NamespacePath string `pulumi:"namespacePath"`
	// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
	PipelineConfigurationFullPath string `pulumi:"pipelineConfigurationFullPath"`
}

func LookupComplianceFrameworkOutput(ctx *pulumi.Context, args LookupComplianceFrameworkOutputArgs, opts ...pulumi.InvokeOption) LookupComplianceFrameworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComplianceFrameworkResult, error) {
			args := v.(LookupComplianceFrameworkArgs)
			r, err := LookupComplianceFramework(ctx, &args, opts...)
			var s LookupComplianceFrameworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComplianceFrameworkResultOutput)
}

// A collection of arguments for invoking getComplianceFramework.
type LookupComplianceFrameworkOutputArgs struct {
	// Name for the compliance framework.
	Name pulumi.StringInput `pulumi:"name"`
	// Full path of the namespace to where the compliance framework is.
	NamespacePath pulumi.StringInput `pulumi:"namespacePath"`
}

func (LookupComplianceFrameworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComplianceFrameworkArgs)(nil)).Elem()
}

// A collection of values returned by getComplianceFramework.
type LookupComplianceFrameworkResultOutput struct{ *pulumi.OutputState }

func (LookupComplianceFrameworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComplianceFrameworkResult)(nil)).Elem()
}

func (o LookupComplianceFrameworkResultOutput) ToLookupComplianceFrameworkResultOutput() LookupComplianceFrameworkResultOutput {
	return o
}

func (o LookupComplianceFrameworkResultOutput) ToLookupComplianceFrameworkResultOutputWithContext(ctx context.Context) LookupComplianceFrameworkResultOutput {
	return o
}

// Color representation of the compliance framework in hex format. e.g. #FCA121.
func (o LookupComplianceFrameworkResultOutput) Color() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComplianceFrameworkResult) string { return v.Color }).(pulumi.StringOutput)
}

// Is the compliance framework the default framework for the group.
func (o LookupComplianceFrameworkResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupComplianceFrameworkResult) bool { return v.Default }).(pulumi.BoolOutput)
}

// Description for the compliance framework.
func (o LookupComplianceFrameworkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComplianceFrameworkResult) string { return v.Description }).(pulumi.StringOutput)
}

// Globally unique ID of the compliance framework.
func (o LookupComplianceFrameworkResultOutput) FrameworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComplianceFrameworkResult) string { return v.FrameworkId }).(pulumi.StringOutput)
}

func (o LookupComplianceFrameworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComplianceFrameworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name for the compliance framework.
func (o LookupComplianceFrameworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComplianceFrameworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Full path of the namespace to where the compliance framework is.
func (o LookupComplianceFrameworkResultOutput) NamespacePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComplianceFrameworkResult) string { return v.NamespacePath }).(pulumi.StringOutput)
}

// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
func (o LookupComplianceFrameworkResultOutput) PipelineConfigurationFullPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComplianceFrameworkResult) string { return v.PipelineConfigurationFullPath }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComplianceFrameworkResultOutput{})
}
