// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ReleaseLink` data source allows get details of a release link.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)
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
//			_, err := gitlab.LookupReleaseLink(ctx, &gitlab.LookupReleaseLinkArgs{
//				LinkId:  11,
//				Project: "foo/bar",
//				TagName: "v1.0.1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupReleaseLink(ctx *pulumi.Context, args *LookupReleaseLinkArgs, opts ...pulumi.InvokeOption) (*LookupReleaseLinkResult, error) {
	var rv LookupReleaseLinkResult
	err := ctx.Invoke("gitlab:index/getReleaseLink:getReleaseLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReleaseLink.
type LookupReleaseLinkArgs struct {
	// The ID of the link.
	LinkId int `pulumi:"linkId"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
	Project string `pulumi:"project"`
	// The tag associated with the Release.
	TagName string `pulumi:"tagName"`
}

// A collection of values returned by getReleaseLink.
type LookupReleaseLinkResult struct {
	// Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
	DirectAssetUrl string `pulumi:"directAssetUrl"`
	// External or internal link.
	External bool `pulumi:"external"`
	// Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
	Filepath string `pulumi:"filepath"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the link.
	LinkId int `pulumi:"linkId"`
	// The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
	LinkType string `pulumi:"linkType"`
	// The name of the link. Link names must be unique within the release.
	Name string `pulumi:"name"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
	Project string `pulumi:"project"`
	// The tag associated with the Release.
	TagName string `pulumi:"tagName"`
	// The URL of the link. Link URLs must be unique within the release.
	Url string `pulumi:"url"`
}

func LookupReleaseLinkOutput(ctx *pulumi.Context, args LookupReleaseLinkOutputArgs, opts ...pulumi.InvokeOption) LookupReleaseLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReleaseLinkResult, error) {
			args := v.(LookupReleaseLinkArgs)
			r, err := LookupReleaseLink(ctx, &args, opts...)
			var s LookupReleaseLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReleaseLinkResultOutput)
}

// A collection of arguments for invoking getReleaseLink.
type LookupReleaseLinkOutputArgs struct {
	// The ID of the link.
	LinkId pulumi.IntInput `pulumi:"linkId"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
	Project pulumi.StringInput `pulumi:"project"`
	// The tag associated with the Release.
	TagName pulumi.StringInput `pulumi:"tagName"`
}

func (LookupReleaseLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReleaseLinkArgs)(nil)).Elem()
}

// A collection of values returned by getReleaseLink.
type LookupReleaseLinkResultOutput struct{ *pulumi.OutputState }

func (LookupReleaseLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReleaseLinkResult)(nil)).Elem()
}

func (o LookupReleaseLinkResultOutput) ToLookupReleaseLinkResultOutput() LookupReleaseLinkResultOutput {
	return o
}

func (o LookupReleaseLinkResultOutput) ToLookupReleaseLinkResultOutputWithContext(ctx context.Context) LookupReleaseLinkResultOutput {
	return o
}

// Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
func (o LookupReleaseLinkResultOutput) DirectAssetUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseLinkResult) string { return v.DirectAssetUrl }).(pulumi.StringOutput)
}

// External or internal link.
func (o LookupReleaseLinkResultOutput) External() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupReleaseLinkResult) bool { return v.External }).(pulumi.BoolOutput)
}

// Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
func (o LookupReleaseLinkResultOutput) Filepath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseLinkResult) string { return v.Filepath }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupReleaseLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the link.
func (o LookupReleaseLinkResultOutput) LinkId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupReleaseLinkResult) int { return v.LinkId }).(pulumi.IntOutput)
}

// The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
func (o LookupReleaseLinkResultOutput) LinkType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseLinkResult) string { return v.LinkType }).(pulumi.StringOutput)
}

// The name of the link. Link names must be unique within the release.
func (o LookupReleaseLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
func (o LookupReleaseLinkResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseLinkResult) string { return v.Project }).(pulumi.StringOutput)
}

// The tag associated with the Release.
func (o LookupReleaseLinkResultOutput) TagName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseLinkResult) string { return v.TagName }).(pulumi.StringOutput)
}

// The URL of the link. Link URLs must be unique within the release.
func (o LookupReleaseLinkResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseLinkResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReleaseLinkResultOutput{})
}
