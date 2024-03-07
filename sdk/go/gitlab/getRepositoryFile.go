// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `RepositoryFile` data source allows details of a file in a repository to be retrieved.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/repository_files.html)
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.LookupRepositoryFile(ctx, &gitlab.LookupRepositoryFileArgs{
//				FilePath: "README.md",
//				Project:  "example",
//				Ref:      "main",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupRepositoryFile(ctx *pulumi.Context, args *LookupRepositoryFileArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryFileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRepositoryFileResult
	err := ctx.Invoke("gitlab:index/getRepositoryFile:getRepositoryFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositoryFile.
type LookupRepositoryFileArgs struct {
	// The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
	FilePath string `pulumi:"filePath"`
	// The name or ID of the project.
	Project string `pulumi:"project"`
	// The name of branch, tag or commit.
	Ref string `pulumi:"ref"`
}

// A collection of values returned by getRepositoryFile.
type LookupRepositoryFileResult struct {
	// The blob id.
	BlobId string `pulumi:"blobId"`
	// The commit id.
	CommitId string `pulumi:"commitId"`
	// File content.
	Content string `pulumi:"content"`
	// File content sha256 digest.
	ContentSha256 string `pulumi:"contentSha256"`
	// The file content encoding.
	Encoding string `pulumi:"encoding"`
	// Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
	ExecuteFilemode bool `pulumi:"executeFilemode"`
	// The filename.
	FileName string `pulumi:"fileName"`
	// The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
	FilePath string `pulumi:"filePath"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The last known commit id.
	LastCommitId string `pulumi:"lastCommitId"`
	// The name or ID of the project.
	Project string `pulumi:"project"`
	// The name of branch, tag or commit.
	Ref string `pulumi:"ref"`
	// The file size.
	Size int `pulumi:"size"`
}

func LookupRepositoryFileOutput(ctx *pulumi.Context, args LookupRepositoryFileOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryFileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRepositoryFileResult, error) {
			args := v.(LookupRepositoryFileArgs)
			r, err := LookupRepositoryFile(ctx, &args, opts...)
			var s LookupRepositoryFileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRepositoryFileResultOutput)
}

// A collection of arguments for invoking getRepositoryFile.
type LookupRepositoryFileOutputArgs struct {
	// The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
	FilePath pulumi.StringInput `pulumi:"filePath"`
	// The name or ID of the project.
	Project pulumi.StringInput `pulumi:"project"`
	// The name of branch, tag or commit.
	Ref pulumi.StringInput `pulumi:"ref"`
}

func (LookupRepositoryFileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryFileArgs)(nil)).Elem()
}

// A collection of values returned by getRepositoryFile.
type LookupRepositoryFileResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryFileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryFileResult)(nil)).Elem()
}

func (o LookupRepositoryFileResultOutput) ToLookupRepositoryFileResultOutput() LookupRepositoryFileResultOutput {
	return o
}

func (o LookupRepositoryFileResultOutput) ToLookupRepositoryFileResultOutputWithContext(ctx context.Context) LookupRepositoryFileResultOutput {
	return o
}

// The blob id.
func (o LookupRepositoryFileResultOutput) BlobId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.BlobId }).(pulumi.StringOutput)
}

// The commit id.
func (o LookupRepositoryFileResultOutput) CommitId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.CommitId }).(pulumi.StringOutput)
}

// File content.
func (o LookupRepositoryFileResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.Content }).(pulumi.StringOutput)
}

// File content sha256 digest.
func (o LookupRepositoryFileResultOutput) ContentSha256() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.ContentSha256 }).(pulumi.StringOutput)
}

// The file content encoding.
func (o LookupRepositoryFileResultOutput) Encoding() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.Encoding }).(pulumi.StringOutput)
}

// Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
func (o LookupRepositoryFileResultOutput) ExecuteFilemode() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) bool { return v.ExecuteFilemode }).(pulumi.BoolOutput)
}

// The filename.
func (o LookupRepositoryFileResultOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.FileName }).(pulumi.StringOutput)
}

// The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
func (o LookupRepositoryFileResultOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.FilePath }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRepositoryFileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last known commit id.
func (o LookupRepositoryFileResultOutput) LastCommitId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.LastCommitId }).(pulumi.StringOutput)
}

// The name or ID of the project.
func (o LookupRepositoryFileResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.Project }).(pulumi.StringOutput)
}

// The name of branch, tag or commit.
func (o LookupRepositoryFileResultOutput) Ref() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) string { return v.Ref }).(pulumi.StringOutput)
}

// The file size.
func (o LookupRepositoryFileResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRepositoryFileResult) int { return v.Size }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryFileResultOutput{})
}
