// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/project.html.markdown.
type Project struct {
	s *pulumi.ResourceState
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOpt) (*Project, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["approvalsBeforeMerge"] = nil
		inputs["archived"] = nil
		inputs["containerRegistryEnabled"] = nil
		inputs["defaultBranch"] = nil
		inputs["description"] = nil
		inputs["initializeWithReadme"] = nil
		inputs["issuesEnabled"] = nil
		inputs["mergeMethod"] = nil
		inputs["mergeRequestsEnabled"] = nil
		inputs["name"] = nil
		inputs["namespaceId"] = nil
		inputs["onlyAllowMergeIfAllDiscussionsAreResolved"] = nil
		inputs["onlyAllowMergeIfPipelineSucceeds"] = nil
		inputs["path"] = nil
		inputs["sharedRunnersEnabled"] = nil
		inputs["sharedWithGroups"] = nil
		inputs["snippetsEnabled"] = nil
		inputs["tags"] = nil
		inputs["visibilityLevel"] = nil
		inputs["wikiEnabled"] = nil
	} else {
		inputs["approvalsBeforeMerge"] = args.ApprovalsBeforeMerge
		inputs["archived"] = args.Archived
		inputs["containerRegistryEnabled"] = args.ContainerRegistryEnabled
		inputs["defaultBranch"] = args.DefaultBranch
		inputs["description"] = args.Description
		inputs["initializeWithReadme"] = args.InitializeWithReadme
		inputs["issuesEnabled"] = args.IssuesEnabled
		inputs["mergeMethod"] = args.MergeMethod
		inputs["mergeRequestsEnabled"] = args.MergeRequestsEnabled
		inputs["name"] = args.Name
		inputs["namespaceId"] = args.NamespaceId
		inputs["onlyAllowMergeIfAllDiscussionsAreResolved"] = args.OnlyAllowMergeIfAllDiscussionsAreResolved
		inputs["onlyAllowMergeIfPipelineSucceeds"] = args.OnlyAllowMergeIfPipelineSucceeds
		inputs["path"] = args.Path
		inputs["sharedRunnersEnabled"] = args.SharedRunnersEnabled
		inputs["sharedWithGroups"] = args.SharedWithGroups
		inputs["snippetsEnabled"] = args.SnippetsEnabled
		inputs["tags"] = args.Tags
		inputs["visibilityLevel"] = args.VisibilityLevel
		inputs["wikiEnabled"] = args.WikiEnabled
	}
	inputs["httpUrlToRepo"] = nil
	inputs["runnersToken"] = nil
	inputs["sshUrlToRepo"] = nil
	inputs["webUrl"] = nil
	s, err := ctx.RegisterResource("gitlab:index/project:Project", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Project{s: s}, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ProjectState, opts ...pulumi.ResourceOpt) (*Project, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["approvalsBeforeMerge"] = state.ApprovalsBeforeMerge
		inputs["archived"] = state.Archived
		inputs["containerRegistryEnabled"] = state.ContainerRegistryEnabled
		inputs["defaultBranch"] = state.DefaultBranch
		inputs["description"] = state.Description
		inputs["httpUrlToRepo"] = state.HttpUrlToRepo
		inputs["initializeWithReadme"] = state.InitializeWithReadme
		inputs["issuesEnabled"] = state.IssuesEnabled
		inputs["mergeMethod"] = state.MergeMethod
		inputs["mergeRequestsEnabled"] = state.MergeRequestsEnabled
		inputs["name"] = state.Name
		inputs["namespaceId"] = state.NamespaceId
		inputs["onlyAllowMergeIfAllDiscussionsAreResolved"] = state.OnlyAllowMergeIfAllDiscussionsAreResolved
		inputs["onlyAllowMergeIfPipelineSucceeds"] = state.OnlyAllowMergeIfPipelineSucceeds
		inputs["path"] = state.Path
		inputs["runnersToken"] = state.RunnersToken
		inputs["sharedRunnersEnabled"] = state.SharedRunnersEnabled
		inputs["sharedWithGroups"] = state.SharedWithGroups
		inputs["snippetsEnabled"] = state.SnippetsEnabled
		inputs["sshUrlToRepo"] = state.SshUrlToRepo
		inputs["tags"] = state.Tags
		inputs["visibilityLevel"] = state.VisibilityLevel
		inputs["webUrl"] = state.WebUrl
		inputs["wikiEnabled"] = state.WikiEnabled
	}
	s, err := ctx.ReadResource("gitlab:index/project:Project", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Project{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Project) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Project) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Number of merge request approvals required for merging. Default is 0.
func (r *Project) ApprovalsBeforeMerge() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["approvalsBeforeMerge"])
}

// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
func (r *Project) Archived() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["archived"])
}

// Enable container registry for the project.
func (r *Project) ContainerRegistryEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["containerRegistryEnabled"])
}

// The default branch for the project.
func (r *Project) DefaultBranch() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["defaultBranch"])
}

// A description of the project.
func (r *Project) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// URL that can be provided to `git clone` to clone the
// repository via HTTP.
func (r *Project) HttpUrlToRepo() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["httpUrlToRepo"])
}

// Create master branch with first commit containing a README.md file.
func (r *Project) InitializeWithReadme() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["initializeWithReadme"])
}

// Enable issue tracking for the project.
func (r *Project) IssuesEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["issuesEnabled"])
}

// Set to `ff` to create fast-forward merges
// Valid values are `merge`, `rebaseMerge`, `ff`
// Repositories are created with `merge` by default
func (r *Project) MergeMethod() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["mergeMethod"])
}

// Enable merge requests for the project.
func (r *Project) MergeRequestsEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["mergeRequestsEnabled"])
}

// The name of the project.
func (r *Project) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The namespace (group or user) of the project. Defaults to your user.
// See `.Group` for an example.
func (r *Project) NamespaceId() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["namespaceId"])
}

// Set to true if you want allow merges only if all discussions are resolved.
func (r *Project) OnlyAllowMergeIfAllDiscussionsAreResolved() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["onlyAllowMergeIfAllDiscussionsAreResolved"])
}

// Set to true if you want allow merges only if a pipeline succeeds.
func (r *Project) OnlyAllowMergeIfPipelineSucceeds() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["onlyAllowMergeIfPipelineSucceeds"])
}

// The path of the repository.
func (r *Project) Path() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["path"])
}

// Registration token to use during runner setup.
func (r *Project) RunnersToken() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["runnersToken"])
}

// Enable shared runners for this project.
func (r *Project) SharedRunnersEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["sharedRunnersEnabled"])
}

// Enable sharing the project with a list of groups (maps).
func (r *Project) SharedWithGroups() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["sharedWithGroups"])
}

// Enable snippets for the project.
func (r *Project) SnippetsEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["snippetsEnabled"])
}

// URL that can be provided to `git clone` to clone the
// repository via SSH.
func (r *Project) SshUrlToRepo() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sshUrlToRepo"])
}

// Tags (topics) of the project.
func (r *Project) Tags() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["tags"])
}

// Set to `public` to create a public project.
// Valid values are `private`, `internal`, `public`.
// Repositories are created as private by default.
func (r *Project) VisibilityLevel() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["visibilityLevel"])
}

// URL that can be used to find the project in a browser.
func (r *Project) WebUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["webUrl"])
}

// Enable wiki for the project.
func (r *Project) WikiEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["wikiEnabled"])
}

// Input properties used for looking up and filtering Project resources.
type ProjectState struct {
	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge interface{}
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived interface{}
	// Enable container registry for the project.
	ContainerRegistryEnabled interface{}
	// The default branch for the project.
	DefaultBranch interface{}
	// A description of the project.
	Description interface{}
	// URL that can be provided to `git clone` to clone the
	// repository via HTTP.
	HttpUrlToRepo interface{}
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme interface{}
	// Enable issue tracking for the project.
	IssuesEnabled interface{}
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod interface{}
	// Enable merge requests for the project.
	MergeRequestsEnabled interface{}
	// The name of the project.
	Name interface{}
	// The namespace (group or user) of the project. Defaults to your user.
	// See `.Group` for an example.
	NamespaceId interface{}
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved interface{}
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds interface{}
	// The path of the repository.
	Path interface{}
	// Registration token to use during runner setup.
	RunnersToken interface{}
	// Enable shared runners for this project.
	SharedRunnersEnabled interface{}
	// Enable sharing the project with a list of groups (maps).
	SharedWithGroups interface{}
	// Enable snippets for the project.
	SnippetsEnabled interface{}
	// URL that can be provided to `git clone` to clone the
	// repository via SSH.
	SshUrlToRepo interface{}
	// Tags (topics) of the project.
	Tags interface{}
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel interface{}
	// URL that can be used to find the project in a browser.
	WebUrl interface{}
	// Enable wiki for the project.
	WikiEnabled interface{}
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge interface{}
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived interface{}
	// Enable container registry for the project.
	ContainerRegistryEnabled interface{}
	// The default branch for the project.
	DefaultBranch interface{}
	// A description of the project.
	Description interface{}
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme interface{}
	// Enable issue tracking for the project.
	IssuesEnabled interface{}
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod interface{}
	// Enable merge requests for the project.
	MergeRequestsEnabled interface{}
	// The name of the project.
	Name interface{}
	// The namespace (group or user) of the project. Defaults to your user.
	// See `.Group` for an example.
	NamespaceId interface{}
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved interface{}
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds interface{}
	// The path of the repository.
	Path interface{}
	// Enable shared runners for this project.
	SharedRunnersEnabled interface{}
	// Enable sharing the project with a list of groups (maps).
	SharedWithGroups interface{}
	// Enable snippets for the project.
	SnippetsEnabled interface{}
	// Tags (topics) of the project.
	Tags interface{}
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel interface{}
	// Enable wiki for the project.
	WikiEnabled interface{}
}
