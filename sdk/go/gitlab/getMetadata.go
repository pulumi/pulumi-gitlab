// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getMetadata` data source retrieves the metadata of the GitLab instance.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/metadata/)
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
//			_, err := gitlab.GetMetadata(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMetadata(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetMetadataResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMetadataResult
	err := ctx.Invoke("gitlab:index/getMetadata:getMetadata", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getMetadata.
type GetMetadataResult struct {
	// If the GitLab instance is an enterprise instance or not.
	Enterprise bool `pulumi:"enterprise"`
	// The id of the data source. It will always be `1`
	Id string `pulumi:"id"`
	// Metadata about the GitLab agent server for Kubernetes (KAS).
	Kas GetMetadataKas `pulumi:"kas"`
	// Revision of the GitLab instance.
	Revision string `pulumi:"revision"`
	// Version of the GitLab instance.
	Version string `pulumi:"version"`
}

func GetMetadataOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetMetadataResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetMetadataResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("gitlab:index/getMetadata:getMetadata", nil, GetMetadataResultOutput{}, options).(GetMetadataResultOutput), nil
	}).(GetMetadataResultOutput)
}

// A collection of values returned by getMetadata.
type GetMetadataResultOutput struct{ *pulumi.OutputState }

func (GetMetadataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMetadataResult)(nil)).Elem()
}

func (o GetMetadataResultOutput) ToGetMetadataResultOutput() GetMetadataResultOutput {
	return o
}

func (o GetMetadataResultOutput) ToGetMetadataResultOutputWithContext(ctx context.Context) GetMetadataResultOutput {
	return o
}

// If the GitLab instance is an enterprise instance or not.
func (o GetMetadataResultOutput) Enterprise() pulumi.BoolOutput {
	return o.ApplyT(func(v GetMetadataResult) bool { return v.Enterprise }).(pulumi.BoolOutput)
}

// The id of the data source. It will always be `1`
func (o GetMetadataResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetadataResult) string { return v.Id }).(pulumi.StringOutput)
}

// Metadata about the GitLab agent server for Kubernetes (KAS).
func (o GetMetadataResultOutput) Kas() GetMetadataKasOutput {
	return o.ApplyT(func(v GetMetadataResult) GetMetadataKas { return v.Kas }).(GetMetadataKasOutput)
}

// Revision of the GitLab instance.
func (o GetMetadataResultOutput) Revision() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetadataResult) string { return v.Revision }).(pulumi.StringOutput)
}

// Version of the GitLab instance.
func (o GetMetadataResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetadataResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMetadataResultOutput{})
}
