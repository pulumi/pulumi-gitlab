// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide details about a specific project in the gitlab provider. The results include the name of the project, path, description, default branch, etc.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gitlab.LookupProject(ctx, &GetProjectArgs{
// 			Id: pulumi.StringRef("foo/bar/baz"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("gitlab:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// The integer or path with namespace that uniquely identifies the project within the gitlab install.
	Id *string `pulumi:"id"`
	// The path of the repository with namespace.
	PathWithNamespace *string `pulumi:"pathWithNamespace"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// Whether the project is in read-only mode (archived).
	Archived bool `pulumi:"archived"`
	// The default branch for the project.
	DefaultBranch string `pulumi:"defaultBranch"`
	// A description of the project.
	Description string `pulumi:"description"`
	// URL that can be provided to `git clone` to clone the
	HttpUrlToRepo string `pulumi:"httpUrlToRepo"`
	// The integer or path with namespace that uniquely identifies the project within the gitlab install.
	Id string `pulumi:"id"`
	// Enable issue tracking for the project.
	IssuesEnabled bool `pulumi:"issuesEnabled"`
	// Enable LFS for the project.
	LfsEnabled bool `pulumi:"lfsEnabled"`
	// Enable merge requests for the project.
	MergeRequestsEnabled bool `pulumi:"mergeRequestsEnabled"`
	// The name of the project.
	Name string `pulumi:"name"`
	// The namespace (group or user) of the project. Defaults to your user.
	NamespaceId int `pulumi:"namespaceId"`
	// The path of the repository.
	Path string `pulumi:"path"`
	// The path of the repository with namespace.
	PathWithNamespace string `pulumi:"pathWithNamespace"`
	// Enable pipelines for the project.
	PipelinesEnabled bool `pulumi:"pipelinesEnabled"`
	// Push rules for the project.
	PushRules GetProjectPushRules `pulumi:"pushRules"`
	// Enable `Delete source branch` option by default for all new merge requests
	RemoveSourceBranchAfterMerge bool `pulumi:"removeSourceBranchAfterMerge"`
	// Allow users to request member access.
	RequestAccessEnabled bool `pulumi:"requestAccessEnabled"`
	// Registration token to use during runner setup.
	RunnersToken string `pulumi:"runnersToken"`
	// Enable snippets for the project.
	SnippetsEnabled bool `pulumi:"snippetsEnabled"`
	// URL that can be provided to `git clone` to clone the
	SshUrlToRepo string `pulumi:"sshUrlToRepo"`
	// Repositories are created as private by default.
	VisibilityLevel string `pulumi:"visibilityLevel"`
	// URL that can be used to find the project in a browser.
	WebUrl string `pulumi:"webUrl"`
	// Enable wiki for the project.
	WikiEnabled bool `pulumi:"wikiEnabled"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			return *r, err
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// The integer or path with namespace that uniquely identifies the project within the gitlab install.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The path of the repository with namespace.
	PathWithNamespace pulumi.StringPtrInput `pulumi:"pathWithNamespace"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// Whether the project is in read-only mode (archived).
func (o LookupProjectResultOutput) Archived() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.Archived }).(pulumi.BoolOutput)
}

// The default branch for the project.
func (o LookupProjectResultOutput) DefaultBranch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.DefaultBranch }).(pulumi.StringOutput)
}

// A description of the project.
func (o LookupProjectResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Description }).(pulumi.StringOutput)
}

// URL that can be provided to `git clone` to clone the
func (o LookupProjectResultOutput) HttpUrlToRepo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.HttpUrlToRepo }).(pulumi.StringOutput)
}

// The integer or path with namespace that uniquely identifies the project within the gitlab install.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable issue tracking for the project.
func (o LookupProjectResultOutput) IssuesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.IssuesEnabled }).(pulumi.BoolOutput)
}

// Enable LFS for the project.
func (o LookupProjectResultOutput) LfsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.LfsEnabled }).(pulumi.BoolOutput)
}

// Enable merge requests for the project.
func (o LookupProjectResultOutput) MergeRequestsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.MergeRequestsEnabled }).(pulumi.BoolOutput)
}

// The name of the project.
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// The namespace (group or user) of the project. Defaults to your user.
func (o LookupProjectResultOutput) NamespaceId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.NamespaceId }).(pulumi.IntOutput)
}

// The path of the repository.
func (o LookupProjectResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Path }).(pulumi.StringOutput)
}

// The path of the repository with namespace.
func (o LookupProjectResultOutput) PathWithNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.PathWithNamespace }).(pulumi.StringOutput)
}

// Enable pipelines for the project.
func (o LookupProjectResultOutput) PipelinesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.PipelinesEnabled }).(pulumi.BoolOutput)
}

// Push rules for the project.
func (o LookupProjectResultOutput) PushRules() GetProjectPushRulesOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectPushRules { return v.PushRules }).(GetProjectPushRulesOutput)
}

// Enable `Delete source branch` option by default for all new merge requests
func (o LookupProjectResultOutput) RemoveSourceBranchAfterMerge() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.RemoveSourceBranchAfterMerge }).(pulumi.BoolOutput)
}

// Allow users to request member access.
func (o LookupProjectResultOutput) RequestAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.RequestAccessEnabled }).(pulumi.BoolOutput)
}

// Registration token to use during runner setup.
func (o LookupProjectResultOutput) RunnersToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.RunnersToken }).(pulumi.StringOutput)
}

// Enable snippets for the project.
func (o LookupProjectResultOutput) SnippetsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.SnippetsEnabled }).(pulumi.BoolOutput)
}

// URL that can be provided to `git clone` to clone the
func (o LookupProjectResultOutput) SshUrlToRepo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.SshUrlToRepo }).(pulumi.StringOutput)
}

// Repositories are created as private by default.
func (o LookupProjectResultOutput) VisibilityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.VisibilityLevel }).(pulumi.StringOutput)
}

// URL that can be used to find the project in a browser.
func (o LookupProjectResultOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.WebUrl }).(pulumi.StringOutput)
}

// Enable wiki for the project.
func (o LookupProjectResultOutput) WikiEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.WikiEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
