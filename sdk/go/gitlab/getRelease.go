// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getRelease` data source retrieves information about a gitlab release for a project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/)
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
//			// By project ID and tag_name
//			_, err := gitlab.GetRelease(ctx, &gitlab.GetReleaseArgs{
//				ProjectId: "1234",
//				TagName:   "v1.0",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRelease(ctx *pulumi.Context, args *GetReleaseArgs, opts ...pulumi.InvokeOption) (*GetReleaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetReleaseResult
	err := ctx.Invoke("gitlab:index/getRelease:getRelease", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRelease.
type GetReleaseArgs struct {
	// The assets for a release
	Assets *GetReleaseAssets `pulumi:"assets"`
	// The ID or URL-encoded path of the project.
	ProjectId string `pulumi:"projectId"`
	// The Git tag the release is associated with.
	TagName string `pulumi:"tagName"`
}

// A collection of values returned by getRelease.
type GetReleaseResult struct {
	// The assets for a release
	Assets *GetReleaseAssets `pulumi:"assets"`
	// The date the release was created.
	CreatedAt string `pulumi:"createdAt"`
	// An HTML rendered description of the release.
	Description string `pulumi:"description"`
	Id          string `pulumi:"id"`
	// The name of the release.
	Name string `pulumi:"name"`
	// The ID or URL-encoded path of the project.
	ProjectId string `pulumi:"projectId"`
	// The date the release was created.
	ReleasedAt string `pulumi:"releasedAt"`
	// The Git tag the release is associated with.
	TagName string `pulumi:"tagName"`
}

func GetReleaseOutput(ctx *pulumi.Context, args GetReleaseOutputArgs, opts ...pulumi.InvokeOption) GetReleaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetReleaseResult, error) {
			args := v.(GetReleaseArgs)
			r, err := GetRelease(ctx, &args, opts...)
			var s GetReleaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetReleaseResultOutput)
}

// A collection of arguments for invoking getRelease.
type GetReleaseOutputArgs struct {
	// The assets for a release
	Assets GetReleaseAssetsPtrInput `pulumi:"assets"`
	// The ID or URL-encoded path of the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// The Git tag the release is associated with.
	TagName pulumi.StringInput `pulumi:"tagName"`
}

func (GetReleaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReleaseArgs)(nil)).Elem()
}

// A collection of values returned by getRelease.
type GetReleaseResultOutput struct{ *pulumi.OutputState }

func (GetReleaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReleaseResult)(nil)).Elem()
}

func (o GetReleaseResultOutput) ToGetReleaseResultOutput() GetReleaseResultOutput {
	return o
}

func (o GetReleaseResultOutput) ToGetReleaseResultOutputWithContext(ctx context.Context) GetReleaseResultOutput {
	return o
}

// The assets for a release
func (o GetReleaseResultOutput) Assets() GetReleaseAssetsPtrOutput {
	return o.ApplyT(func(v GetReleaseResult) *GetReleaseAssets { return v.Assets }).(GetReleaseAssetsPtrOutput)
}

// The date the release was created.
func (o GetReleaseResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetReleaseResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// An HTML rendered description of the release.
func (o GetReleaseResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetReleaseResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetReleaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetReleaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the release.
func (o GetReleaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetReleaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID or URL-encoded path of the project.
func (o GetReleaseResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetReleaseResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The date the release was created.
func (o GetReleaseResultOutput) ReleasedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetReleaseResult) string { return v.ReleasedAt }).(pulumi.StringOutput)
}

// The Git tag the release is associated with.
func (o GetReleaseResultOutput) TagName() pulumi.StringOutput {
	return o.ApplyT(func(v GetReleaseResult) string { return v.TagName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetReleaseResultOutput{})
}
