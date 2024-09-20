// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getProjectProtectedTags` data source allows details of the protected tags of a given project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_tags.html#list-protected-tags)
func GetProjectProtectedTags(ctx *pulumi.Context, args *GetProjectProtectedTagsArgs, opts ...pulumi.InvokeOption) (*GetProjectProtectedTagsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectProtectedTagsResult
	err := ctx.Invoke("gitlab:index/getProjectProtectedTags:getProjectProtectedTags", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectProtectedTags.
type GetProjectProtectedTagsArgs struct {
	// The integer or path with namespace that uniquely identifies the project.
	Project string `pulumi:"project"`
}

// A collection of values returned by getProjectProtectedTags.
type GetProjectProtectedTagsResult struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// The integer or path with namespace that uniquely identifies the project.
	Project string `pulumi:"project"`
	// A list of protected tags, as defined below.
	ProtectedTags []GetProjectProtectedTagsProtectedTag `pulumi:"protectedTags"`
}

func GetProjectProtectedTagsOutput(ctx *pulumi.Context, args GetProjectProtectedTagsOutputArgs, opts ...pulumi.InvokeOption) GetProjectProtectedTagsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectProtectedTagsResult, error) {
			args := v.(GetProjectProtectedTagsArgs)
			r, err := GetProjectProtectedTags(ctx, &args, opts...)
			var s GetProjectProtectedTagsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectProtectedTagsResultOutput)
}

// A collection of arguments for invoking getProjectProtectedTags.
type GetProjectProtectedTagsOutputArgs struct {
	// The integer or path with namespace that uniquely identifies the project.
	Project pulumi.StringInput `pulumi:"project"`
}

func (GetProjectProtectedTagsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectProtectedTagsArgs)(nil)).Elem()
}

// A collection of values returned by getProjectProtectedTags.
type GetProjectProtectedTagsResultOutput struct{ *pulumi.OutputState }

func (GetProjectProtectedTagsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectProtectedTagsResult)(nil)).Elem()
}

func (o GetProjectProtectedTagsResultOutput) ToGetProjectProtectedTagsResultOutput() GetProjectProtectedTagsResultOutput {
	return o
}

func (o GetProjectProtectedTagsResultOutput) ToGetProjectProtectedTagsResultOutputWithContext(ctx context.Context) GetProjectProtectedTagsResultOutput {
	return o
}

// The ID of this resource.
func (o GetProjectProtectedTagsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectProtectedTagsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The integer or path with namespace that uniquely identifies the project.
func (o GetProjectProtectedTagsResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectProtectedTagsResult) string { return v.Project }).(pulumi.StringOutput)
}

// A list of protected tags, as defined below.
func (o GetProjectProtectedTagsResultOutput) ProtectedTags() GetProjectProtectedTagsProtectedTagArrayOutput {
	return o.ApplyT(func(v GetProjectProtectedTagsResult) []GetProjectProtectedTagsProtectedTag { return v.ProtectedTags }).(GetProjectProtectedTagsProtectedTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectProtectedTagsResultOutput{})
}