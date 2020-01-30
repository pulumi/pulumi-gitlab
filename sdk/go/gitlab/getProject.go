// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gitlab

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides details about a specific project in the gitlab provider. The results include the name of the project, path, description, default branch, etc.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/d/project.html.markdown.
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
	Archived      *bool   `pulumi:"archived"`
	DefaultBranch *string `pulumi:"defaultBranch"`
	Description   *string `pulumi:"description"`
	HttpUrlToRepo *string `pulumi:"httpUrlToRepo"`
	// The integer that uniquely identifies the project within the gitlab install.
	Id                   int     `pulumi:"id"`
	IssuesEnabled        *bool   `pulumi:"issuesEnabled"`
	LfsEnabled           *bool   `pulumi:"lfsEnabled"`
	MergeRequestsEnabled *bool   `pulumi:"mergeRequestsEnabled"`
	Name                 *string `pulumi:"name"`
	NamespaceId          *int    `pulumi:"namespaceId"`
	Path                 *string `pulumi:"path"`
	PipelinesEnabled     *bool   `pulumi:"pipelinesEnabled"`
	RequestAccessEnabled *bool   `pulumi:"requestAccessEnabled"`
	RunnersToken         *string `pulumi:"runnersToken"`
	SnippetsEnabled      *bool   `pulumi:"snippetsEnabled"`
	SshUrlToRepo         *string `pulumi:"sshUrlToRepo"`
	VisibilityLevel      *string `pulumi:"visibilityLevel"`
	WebUrl               *string `pulumi:"webUrl"`
	WikiEnabled          *bool   `pulumi:"wikiEnabled"`
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
	// repository via HTTP.
	HttpUrlToRepo string `pulumi:"httpUrlToRepo"`
	// Integer that uniquely identifies the project within the gitlab install.
	Id int `pulumi:"id"`
	// Enable issue tracking for the project.
	IssuesEnabled bool `pulumi:"issuesEnabled"`
	// Enable LFS for the project.
	LfsEnabled bool `pulumi:"lfsEnabled"`
	// Enable merge requests for the project.
	MergeRequestsEnabled bool   `pulumi:"mergeRequestsEnabled"`
	Name                 string `pulumi:"name"`
	// The namespace (group or user) of the project. Defaults to your user.
	// See `.Group` for an example.
	NamespaceId int `pulumi:"namespaceId"`
	// The path of the repository.
	Path string `pulumi:"path"`
	// Enable pipelines for the project.
	PipelinesEnabled bool `pulumi:"pipelinesEnabled"`
	// Allow users to request member access.
	RequestAccessEnabled bool `pulumi:"requestAccessEnabled"`
	// Registration token to use during runner setup.
	RunnersToken string `pulumi:"runnersToken"`
	// Enable snippets for the project.
	SnippetsEnabled bool `pulumi:"snippetsEnabled"`
	// URL that can be provided to `git clone` to clone the
	// repository via SSH.
	SshUrlToRepo string `pulumi:"sshUrlToRepo"`
	// Repositories are created as private by default.
	VisibilityLevel string `pulumi:"visibilityLevel"`
	// URL that can be used to find the project in a browser.
	WebUrl string `pulumi:"webUrl"`
	// Enable wiki for the project.
	WikiEnabled bool `pulumi:"wikiEnabled"`
}
