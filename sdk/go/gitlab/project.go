// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gitlab

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/project.html.markdown.
type Project struct {
	pulumi.CustomResourceState

	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge pulumi.IntPtrOutput `pulumi:"approvalsBeforeMerge"`
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived pulumi.BoolPtrOutput `pulumi:"archived"`
	// Enable container registry for the project.
	ContainerRegistryEnabled pulumi.BoolPtrOutput `pulumi:"containerRegistryEnabled"`
	// The default branch for the project.
	DefaultBranch pulumi.StringPtrOutput `pulumi:"defaultBranch"`
	// A description of the project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// URL that can be provided to `git clone` to clone the
	// repository via HTTP.
	HttpUrlToRepo pulumi.StringOutput `pulumi:"httpUrlToRepo"`
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme pulumi.BoolPtrOutput `pulumi:"initializeWithReadme"`
	// Enable issue tracking for the project.
	IssuesEnabled pulumi.BoolPtrOutput `pulumi:"issuesEnabled"`
	// Enable LFS for the project.
	LfsEnabled pulumi.BoolPtrOutput `pulumi:"lfsEnabled"`
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod pulumi.StringPtrOutput `pulumi:"mergeMethod"`
	// Enable merge requests for the project.
	MergeRequestsEnabled pulumi.BoolPtrOutput `pulumi:"mergeRequestsEnabled"`
	// The name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace (group or user) of the project. Defaults to your user.
	// See `.Group` for an example.
	NamespaceId pulumi.IntOutput `pulumi:"namespaceId"`
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved pulumi.BoolPtrOutput `pulumi:"onlyAllowMergeIfAllDiscussionsAreResolved"`
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds pulumi.BoolPtrOutput `pulumi:"onlyAllowMergeIfPipelineSucceeds"`
	// The path of the repository.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Enable pipelines for the project.
	PipelinesEnabled pulumi.BoolPtrOutput `pulumi:"pipelinesEnabled"`
	// Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrOutput `pulumi:"requestAccessEnabled"`
	// Registration token to use during runner setup.
	RunnersToken pulumi.StringOutput `pulumi:"runnersToken"`
	// Enable shared runners for this project.
	SharedRunnersEnabled pulumi.BoolOutput `pulumi:"sharedRunnersEnabled"`
	// Enable sharing the project with a list of groups (maps).
	SharedWithGroups ProjectSharedWithGroupArrayOutput `pulumi:"sharedWithGroups"`
	// Enable snippets for the project.
	SnippetsEnabled pulumi.BoolPtrOutput `pulumi:"snippetsEnabled"`
	// URL that can be provided to `git clone` to clone the
	// repository via SSH.
	SshUrlToRepo pulumi.StringOutput `pulumi:"sshUrlToRepo"`
	// Tags (topics) of the project.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel pulumi.StringPtrOutput `pulumi:"visibilityLevel"`
	// URL that can be used to find the project in a browser.
	WebUrl pulumi.StringOutput `pulumi:"webUrl"`
	// Enable wiki for the project.
	WikiEnabled pulumi.BoolPtrOutput `pulumi:"wikiEnabled"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}
	var resource Project
	err := ctx.RegisterResource("gitlab:index/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("gitlab:index/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge *int `pulumi:"approvalsBeforeMerge"`
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived *bool `pulumi:"archived"`
	// Enable container registry for the project.
	ContainerRegistryEnabled *bool `pulumi:"containerRegistryEnabled"`
	// The default branch for the project.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// A description of the project.
	Description *string `pulumi:"description"`
	// URL that can be provided to `git clone` to clone the
	// repository via HTTP.
	HttpUrlToRepo *string `pulumi:"httpUrlToRepo"`
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme *bool `pulumi:"initializeWithReadme"`
	// Enable issue tracking for the project.
	IssuesEnabled *bool `pulumi:"issuesEnabled"`
	// Enable LFS for the project.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod *string `pulumi:"mergeMethod"`
	// Enable merge requests for the project.
	MergeRequestsEnabled *bool `pulumi:"mergeRequestsEnabled"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The namespace (group or user) of the project. Defaults to your user.
	// See `.Group` for an example.
	NamespaceId *int `pulumi:"namespaceId"`
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool `pulumi:"onlyAllowMergeIfAllDiscussionsAreResolved"`
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds *bool `pulumi:"onlyAllowMergeIfPipelineSucceeds"`
	// The path of the repository.
	Path *string `pulumi:"path"`
	// Enable pipelines for the project.
	PipelinesEnabled *bool `pulumi:"pipelinesEnabled"`
	// Allow users to request member access.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Registration token to use during runner setup.
	RunnersToken *string `pulumi:"runnersToken"`
	// Enable shared runners for this project.
	SharedRunnersEnabled *bool `pulumi:"sharedRunnersEnabled"`
	// Enable sharing the project with a list of groups (maps).
	SharedWithGroups []ProjectSharedWithGroup `pulumi:"sharedWithGroups"`
	// Enable snippets for the project.
	SnippetsEnabled *bool `pulumi:"snippetsEnabled"`
	// URL that can be provided to `git clone` to clone the
	// repository via SSH.
	SshUrlToRepo *string `pulumi:"sshUrlToRepo"`
	// Tags (topics) of the project.
	Tags []string `pulumi:"tags"`
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
	// URL that can be used to find the project in a browser.
	WebUrl *string `pulumi:"webUrl"`
	// Enable wiki for the project.
	WikiEnabled *bool `pulumi:"wikiEnabled"`
}

type ProjectState struct {
	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge pulumi.IntPtrInput
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived pulumi.BoolPtrInput
	// Enable container registry for the project.
	ContainerRegistryEnabled pulumi.BoolPtrInput
	// The default branch for the project.
	DefaultBranch pulumi.StringPtrInput
	// A description of the project.
	Description pulumi.StringPtrInput
	// URL that can be provided to `git clone` to clone the
	// repository via HTTP.
	HttpUrlToRepo pulumi.StringPtrInput
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme pulumi.BoolPtrInput
	// Enable issue tracking for the project.
	IssuesEnabled pulumi.BoolPtrInput
	// Enable LFS for the project.
	LfsEnabled pulumi.BoolPtrInput
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod pulumi.StringPtrInput
	// Enable merge requests for the project.
	MergeRequestsEnabled pulumi.BoolPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// The namespace (group or user) of the project. Defaults to your user.
	// See `.Group` for an example.
	NamespaceId pulumi.IntPtrInput
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved pulumi.BoolPtrInput
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds pulumi.BoolPtrInput
	// The path of the repository.
	Path pulumi.StringPtrInput
	// Enable pipelines for the project.
	PipelinesEnabled pulumi.BoolPtrInput
	// Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Registration token to use during runner setup.
	RunnersToken pulumi.StringPtrInput
	// Enable shared runners for this project.
	SharedRunnersEnabled pulumi.BoolPtrInput
	// Enable sharing the project with a list of groups (maps).
	SharedWithGroups ProjectSharedWithGroupArrayInput
	// Enable snippets for the project.
	SnippetsEnabled pulumi.BoolPtrInput
	// URL that can be provided to `git clone` to clone the
	// repository via SSH.
	SshUrlToRepo pulumi.StringPtrInput
	// Tags (topics) of the project.
	Tags pulumi.StringArrayInput
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel pulumi.StringPtrInput
	// URL that can be used to find the project in a browser.
	WebUrl pulumi.StringPtrInput
	// Enable wiki for the project.
	WikiEnabled pulumi.BoolPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge *int `pulumi:"approvalsBeforeMerge"`
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived *bool `pulumi:"archived"`
	// Enable container registry for the project.
	ContainerRegistryEnabled *bool `pulumi:"containerRegistryEnabled"`
	// The default branch for the project.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// A description of the project.
	Description *string `pulumi:"description"`
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme *bool `pulumi:"initializeWithReadme"`
	// Enable issue tracking for the project.
	IssuesEnabled *bool `pulumi:"issuesEnabled"`
	// Enable LFS for the project.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod *string `pulumi:"mergeMethod"`
	// Enable merge requests for the project.
	MergeRequestsEnabled *bool `pulumi:"mergeRequestsEnabled"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The namespace (group or user) of the project. Defaults to your user.
	// See `.Group` for an example.
	NamespaceId *int `pulumi:"namespaceId"`
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool `pulumi:"onlyAllowMergeIfAllDiscussionsAreResolved"`
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds *bool `pulumi:"onlyAllowMergeIfPipelineSucceeds"`
	// The path of the repository.
	Path *string `pulumi:"path"`
	// Enable pipelines for the project.
	PipelinesEnabled *bool `pulumi:"pipelinesEnabled"`
	// Allow users to request member access.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Enable shared runners for this project.
	SharedRunnersEnabled *bool `pulumi:"sharedRunnersEnabled"`
	// Enable sharing the project with a list of groups (maps).
	SharedWithGroups []ProjectSharedWithGroup `pulumi:"sharedWithGroups"`
	// Enable snippets for the project.
	SnippetsEnabled *bool `pulumi:"snippetsEnabled"`
	// Tags (topics) of the project.
	Tags []string `pulumi:"tags"`
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
	// Enable wiki for the project.
	WikiEnabled *bool `pulumi:"wikiEnabled"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge pulumi.IntPtrInput
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived pulumi.BoolPtrInput
	// Enable container registry for the project.
	ContainerRegistryEnabled pulumi.BoolPtrInput
	// The default branch for the project.
	DefaultBranch pulumi.StringPtrInput
	// A description of the project.
	Description pulumi.StringPtrInput
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme pulumi.BoolPtrInput
	// Enable issue tracking for the project.
	IssuesEnabled pulumi.BoolPtrInput
	// Enable LFS for the project.
	LfsEnabled pulumi.BoolPtrInput
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod pulumi.StringPtrInput
	// Enable merge requests for the project.
	MergeRequestsEnabled pulumi.BoolPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// The namespace (group or user) of the project. Defaults to your user.
	// See `.Group` for an example.
	NamespaceId pulumi.IntPtrInput
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved pulumi.BoolPtrInput
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds pulumi.BoolPtrInput
	// The path of the repository.
	Path pulumi.StringPtrInput
	// Enable pipelines for the project.
	PipelinesEnabled pulumi.BoolPtrInput
	// Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Enable shared runners for this project.
	SharedRunnersEnabled pulumi.BoolPtrInput
	// Enable sharing the project with a list of groups (maps).
	SharedWithGroups ProjectSharedWithGroupArrayInput
	// Enable snippets for the project.
	SnippetsEnabled pulumi.BoolPtrInput
	// Tags (topics) of the project.
	Tags pulumi.StringArrayInput
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel pulumi.StringPtrInput
	// Enable wiki for the project.
	WikiEnabled pulumi.BoolPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}
