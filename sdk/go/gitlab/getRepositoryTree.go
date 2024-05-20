// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v7/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getRepositoryTree` data source allows details of directories and files in a repository to be retrieved.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/repositories.html#list-repository-tree)
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
//			_, err := gitlab.GetRepositoryTree(ctx, &gitlab.GetRepositoryTreeArgs{
//				Project:   "example",
//				Ref:       "main",
//				Path:      pulumi.StringRef("ExampleSubFolder"),
//				Recursive: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRepositoryTree(ctx *pulumi.Context, args *GetRepositoryTreeArgs, opts ...pulumi.InvokeOption) (*GetRepositoryTreeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRepositoryTreeResult
	err := ctx.Invoke("gitlab:index/getRepositoryTree:getRepositoryTree", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositoryTree.
type GetRepositoryTreeArgs struct {
	// The path inside repository. Used to get content of subdirectories.
	Path *string `pulumi:"path"`
	// The ID or full path of the project owned by the authenticated user.
	Project string `pulumi:"project"`
	// Boolean value used to get a recursive tree (false by default).
	Recursive *bool `pulumi:"recursive"`
	// The name of a repository branch or tag.
	Ref string `pulumi:"ref"`
}

// A collection of values returned by getRepositoryTree.
type GetRepositoryTreeResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The path inside repository. Used to get content of subdirectories.
	Path *string `pulumi:"path"`
	// The ID or full path of the project owned by the authenticated user.
	Project string `pulumi:"project"`
	// Boolean value used to get a recursive tree (false by default).
	Recursive *bool `pulumi:"recursive"`
	// The name of a repository branch or tag.
	Ref string `pulumi:"ref"`
	// The list of files/directories returned by the search
	Trees []GetRepositoryTreeTree `pulumi:"trees"`
}

func GetRepositoryTreeOutput(ctx *pulumi.Context, args GetRepositoryTreeOutputArgs, opts ...pulumi.InvokeOption) GetRepositoryTreeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRepositoryTreeResult, error) {
			args := v.(GetRepositoryTreeArgs)
			r, err := GetRepositoryTree(ctx, &args, opts...)
			var s GetRepositoryTreeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRepositoryTreeResultOutput)
}

// A collection of arguments for invoking getRepositoryTree.
type GetRepositoryTreeOutputArgs struct {
	// The path inside repository. Used to get content of subdirectories.
	Path pulumi.StringPtrInput `pulumi:"path"`
	// The ID or full path of the project owned by the authenticated user.
	Project pulumi.StringInput `pulumi:"project"`
	// Boolean value used to get a recursive tree (false by default).
	Recursive pulumi.BoolPtrInput `pulumi:"recursive"`
	// The name of a repository branch or tag.
	Ref pulumi.StringInput `pulumi:"ref"`
}

func (GetRepositoryTreeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoryTreeArgs)(nil)).Elem()
}

// A collection of values returned by getRepositoryTree.
type GetRepositoryTreeResultOutput struct{ *pulumi.OutputState }

func (GetRepositoryTreeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoryTreeResult)(nil)).Elem()
}

func (o GetRepositoryTreeResultOutput) ToGetRepositoryTreeResultOutput() GetRepositoryTreeResultOutput {
	return o
}

func (o GetRepositoryTreeResultOutput) ToGetRepositoryTreeResultOutputWithContext(ctx context.Context) GetRepositoryTreeResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRepositoryTreeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoryTreeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The path inside repository. Used to get content of subdirectories.
func (o GetRepositoryTreeResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRepositoryTreeResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// The ID or full path of the project owned by the authenticated user.
func (o GetRepositoryTreeResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoryTreeResult) string { return v.Project }).(pulumi.StringOutput)
}

// Boolean value used to get a recursive tree (false by default).
func (o GetRepositoryTreeResultOutput) Recursive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetRepositoryTreeResult) *bool { return v.Recursive }).(pulumi.BoolPtrOutput)
}

// The name of a repository branch or tag.
func (o GetRepositoryTreeResultOutput) Ref() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoryTreeResult) string { return v.Ref }).(pulumi.StringOutput)
}

// The list of files/directories returned by the search
func (o GetRepositoryTreeResultOutput) Trees() GetRepositoryTreeTreeArrayOutput {
	return o.ApplyT(func(v GetRepositoryTreeResult) []GetRepositoryTreeTree { return v.Trees }).(GetRepositoryTreeTreeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRepositoryTreeResultOutput{})
}
