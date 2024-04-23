// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getReleaseLinks` data source allows get details of release links.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)
func GetReleaseLinks(ctx *pulumi.Context, args *GetReleaseLinksArgs, opts ...pulumi.InvokeOption) (*GetReleaseLinksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetReleaseLinksResult
	err := ctx.Invoke("gitlab:index/getReleaseLinks:getReleaseLinks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReleaseLinks.
type GetReleaseLinksArgs struct {
	// The ID or full path to the project.
	Project string `pulumi:"project"`
	// The tag associated with the Release.
	TagName string `pulumi:"tagName"`
}

// A collection of values returned by getReleaseLinks.
type GetReleaseLinksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID or full path to the project.
	Project string `pulumi:"project"`
	// List of release links
	ReleaseLinks []GetReleaseLinksReleaseLink `pulumi:"releaseLinks"`
	// The tag associated with the Release.
	TagName string `pulumi:"tagName"`
}

func GetReleaseLinksOutput(ctx *pulumi.Context, args GetReleaseLinksOutputArgs, opts ...pulumi.InvokeOption) GetReleaseLinksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetReleaseLinksResult, error) {
			args := v.(GetReleaseLinksArgs)
			r, err := GetReleaseLinks(ctx, &args, opts...)
			var s GetReleaseLinksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetReleaseLinksResultOutput)
}

// A collection of arguments for invoking getReleaseLinks.
type GetReleaseLinksOutputArgs struct {
	// The ID or full path to the project.
	Project pulumi.StringInput `pulumi:"project"`
	// The tag associated with the Release.
	TagName pulumi.StringInput `pulumi:"tagName"`
}

func (GetReleaseLinksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReleaseLinksArgs)(nil)).Elem()
}

// A collection of values returned by getReleaseLinks.
type GetReleaseLinksResultOutput struct{ *pulumi.OutputState }

func (GetReleaseLinksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReleaseLinksResult)(nil)).Elem()
}

func (o GetReleaseLinksResultOutput) ToGetReleaseLinksResultOutput() GetReleaseLinksResultOutput {
	return o
}

func (o GetReleaseLinksResultOutput) ToGetReleaseLinksResultOutputWithContext(ctx context.Context) GetReleaseLinksResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetReleaseLinksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetReleaseLinksResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID or full path to the project.
func (o GetReleaseLinksResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetReleaseLinksResult) string { return v.Project }).(pulumi.StringOutput)
}

// List of release links
func (o GetReleaseLinksResultOutput) ReleaseLinks() GetReleaseLinksReleaseLinkArrayOutput {
	return o.ApplyT(func(v GetReleaseLinksResult) []GetReleaseLinksReleaseLink { return v.ReleaseLinks }).(GetReleaseLinksReleaseLinkArrayOutput)
}

// The tag associated with the Release.
func (o GetReleaseLinksResultOutput) TagName() pulumi.StringOutput {
	return o.ApplyT(func(v GetReleaseLinksResult) string { return v.TagName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetReleaseLinksResultOutput{})
}
