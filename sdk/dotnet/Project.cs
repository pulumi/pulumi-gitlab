// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gitlab
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/project.html.markdown.
    /// </summary>
    public partial class Project : Pulumi.CustomResource
    {
        /// <summary>
        /// Number of merge request approvals required for merging. Default is 0.
        /// </summary>
        [Output("approvalsBeforeMerge")]
        public Output<int?> ApprovalsBeforeMerge { get; private set; } = null!;

        /// <summary>
        /// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
        /// </summary>
        [Output("archived")]
        public Output<bool?> Archived { get; private set; } = null!;

        /// <summary>
        /// Enable container registry for the project.
        /// </summary>
        [Output("containerRegistryEnabled")]
        public Output<bool?> ContainerRegistryEnabled { get; private set; } = null!;

        /// <summary>
        /// The default branch for the project.
        /// </summary>
        [Output("defaultBranch")]
        public Output<string?> DefaultBranch { get; private set; } = null!;

        /// <summary>
        /// A description of the project.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// URL that can be provided to `git clone` to clone the
        /// repository via HTTP.
        /// </summary>
        [Output("httpUrlToRepo")]
        public Output<string> HttpUrlToRepo { get; private set; } = null!;

        /// <summary>
        /// Create master branch with first commit containing a README.md file.
        /// </summary>
        [Output("initializeWithReadme")]
        public Output<bool?> InitializeWithReadme { get; private set; } = null!;

        /// <summary>
        /// Enable issue tracking for the project.
        /// </summary>
        [Output("issuesEnabled")]
        public Output<bool?> IssuesEnabled { get; private set; } = null!;

        /// <summary>
        /// Set to `ff` to create fast-forward merges
        /// Valid values are `merge`, `rebase_merge`, `ff`
        /// Repositories are created with `merge` by default
        /// </summary>
        [Output("mergeMethod")]
        public Output<string?> MergeMethod { get; private set; } = null!;

        /// <summary>
        /// Enable merge requests for the project.
        /// </summary>
        [Output("mergeRequestsEnabled")]
        public Output<bool?> MergeRequestsEnabled { get; private set; } = null!;

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace (group or user) of the project. Defaults to your user.
        /// See `gitlab..Group` for an example.
        /// </summary>
        [Output("namespaceId")]
        public Output<int> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// Set to true if you want allow merges only if all discussions are resolved.
        /// </summary>
        [Output("onlyAllowMergeIfAllDiscussionsAreResolved")]
        public Output<bool?> OnlyAllowMergeIfAllDiscussionsAreResolved { get; private set; } = null!;

        /// <summary>
        /// Set to true if you want allow merges only if a pipeline succeeds.
        /// </summary>
        [Output("onlyAllowMergeIfPipelineSucceeds")]
        public Output<bool?> OnlyAllowMergeIfPipelineSucceeds { get; private set; } = null!;

        /// <summary>
        /// The path of the repository.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// Registration token to use during runner setup.
        /// </summary>
        [Output("runnersToken")]
        public Output<string> RunnersToken { get; private set; } = null!;

        /// <summary>
        /// Enable shared runners for this project.
        /// </summary>
        [Output("sharedRunnersEnabled")]
        public Output<bool> SharedRunnersEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable sharing the project with a list of groups (maps).
        /// </summary>
        [Output("sharedWithGroups")]
        public Output<ImmutableArray<Outputs.ProjectSharedWithGroups>> SharedWithGroups { get; private set; } = null!;

        /// <summary>
        /// Enable snippets for the project.
        /// </summary>
        [Output("snippetsEnabled")]
        public Output<bool?> SnippetsEnabled { get; private set; } = null!;

        /// <summary>
        /// URL that can be provided to `git clone` to clone the
        /// repository via SSH.
        /// </summary>
        [Output("sshUrlToRepo")]
        public Output<string> SshUrlToRepo { get; private set; } = null!;

        /// <summary>
        /// Tags (topics) of the project.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Set to `public` to create a public project.
        /// Valid values are `private`, `internal`, `public`.
        /// Repositories are created as private by default.
        /// </summary>
        [Output("visibilityLevel")]
        public Output<string?> VisibilityLevel { get; private set; } = null!;

        /// <summary>
        /// URL that can be used to find the project in a browser.
        /// </summary>
        [Output("webUrl")]
        public Output<string> WebUrl { get; private set; } = null!;

        /// <summary>
        /// Enable wiki for the project.
        /// </summary>
        [Output("wikiEnabled")]
        public Output<bool?> WikiEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs? args = null, CustomResourceOptions? options = null)
            : base("gitlab:index/project:Project", name, args, MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/project:Project", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of merge request approvals required for merging. Default is 0.
        /// </summary>
        [Input("approvalsBeforeMerge")]
        public Input<int>? ApprovalsBeforeMerge { get; set; }

        /// <summary>
        /// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
        /// </summary>
        [Input("archived")]
        public Input<bool>? Archived { get; set; }

        /// <summary>
        /// Enable container registry for the project.
        /// </summary>
        [Input("containerRegistryEnabled")]
        public Input<bool>? ContainerRegistryEnabled { get; set; }

        /// <summary>
        /// The default branch for the project.
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        /// <summary>
        /// A description of the project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Create master branch with first commit containing a README.md file.
        /// </summary>
        [Input("initializeWithReadme")]
        public Input<bool>? InitializeWithReadme { get; set; }

        /// <summary>
        /// Enable issue tracking for the project.
        /// </summary>
        [Input("issuesEnabled")]
        public Input<bool>? IssuesEnabled { get; set; }

        /// <summary>
        /// Set to `ff` to create fast-forward merges
        /// Valid values are `merge`, `rebase_merge`, `ff`
        /// Repositories are created with `merge` by default
        /// </summary>
        [Input("mergeMethod")]
        public Input<string>? MergeMethod { get; set; }

        /// <summary>
        /// Enable merge requests for the project.
        /// </summary>
        [Input("mergeRequestsEnabled")]
        public Input<bool>? MergeRequestsEnabled { get; set; }

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace (group or user) of the project. Defaults to your user.
        /// See `gitlab..Group` for an example.
        /// </summary>
        [Input("namespaceId")]
        public Input<int>? NamespaceId { get; set; }

        /// <summary>
        /// Set to true if you want allow merges only if all discussions are resolved.
        /// </summary>
        [Input("onlyAllowMergeIfAllDiscussionsAreResolved")]
        public Input<bool>? OnlyAllowMergeIfAllDiscussionsAreResolved { get; set; }

        /// <summary>
        /// Set to true if you want allow merges only if a pipeline succeeds.
        /// </summary>
        [Input("onlyAllowMergeIfPipelineSucceeds")]
        public Input<bool>? OnlyAllowMergeIfPipelineSucceeds { get; set; }

        /// <summary>
        /// The path of the repository.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Enable shared runners for this project.
        /// </summary>
        [Input("sharedRunnersEnabled")]
        public Input<bool>? SharedRunnersEnabled { get; set; }

        [Input("sharedWithGroups")]
        private InputList<Inputs.ProjectSharedWithGroupsArgs>? _sharedWithGroups;

        /// <summary>
        /// Enable sharing the project with a list of groups (maps).
        /// </summary>
        public InputList<Inputs.ProjectSharedWithGroupsArgs> SharedWithGroups
        {
            get => _sharedWithGroups ?? (_sharedWithGroups = new InputList<Inputs.ProjectSharedWithGroupsArgs>());
            set => _sharedWithGroups = value;
        }

        /// <summary>
        /// Enable snippets for the project.
        /// </summary>
        [Input("snippetsEnabled")]
        public Input<bool>? SnippetsEnabled { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags (topics) of the project.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Set to `public` to create a public project.
        /// Valid values are `private`, `internal`, `public`.
        /// Repositories are created as private by default.
        /// </summary>
        [Input("visibilityLevel")]
        public Input<string>? VisibilityLevel { get; set; }

        /// <summary>
        /// Enable wiki for the project.
        /// </summary>
        [Input("wikiEnabled")]
        public Input<bool>? WikiEnabled { get; set; }

        public ProjectArgs()
        {
        }
    }

    public sealed class ProjectState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of merge request approvals required for merging. Default is 0.
        /// </summary>
        [Input("approvalsBeforeMerge")]
        public Input<int>? ApprovalsBeforeMerge { get; set; }

        /// <summary>
        /// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
        /// </summary>
        [Input("archived")]
        public Input<bool>? Archived { get; set; }

        /// <summary>
        /// Enable container registry for the project.
        /// </summary>
        [Input("containerRegistryEnabled")]
        public Input<bool>? ContainerRegistryEnabled { get; set; }

        /// <summary>
        /// The default branch for the project.
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        /// <summary>
        /// A description of the project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// URL that can be provided to `git clone` to clone the
        /// repository via HTTP.
        /// </summary>
        [Input("httpUrlToRepo")]
        public Input<string>? HttpUrlToRepo { get; set; }

        /// <summary>
        /// Create master branch with first commit containing a README.md file.
        /// </summary>
        [Input("initializeWithReadme")]
        public Input<bool>? InitializeWithReadme { get; set; }

        /// <summary>
        /// Enable issue tracking for the project.
        /// </summary>
        [Input("issuesEnabled")]
        public Input<bool>? IssuesEnabled { get; set; }

        /// <summary>
        /// Set to `ff` to create fast-forward merges
        /// Valid values are `merge`, `rebase_merge`, `ff`
        /// Repositories are created with `merge` by default
        /// </summary>
        [Input("mergeMethod")]
        public Input<string>? MergeMethod { get; set; }

        /// <summary>
        /// Enable merge requests for the project.
        /// </summary>
        [Input("mergeRequestsEnabled")]
        public Input<bool>? MergeRequestsEnabled { get; set; }

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace (group or user) of the project. Defaults to your user.
        /// See `gitlab..Group` for an example.
        /// </summary>
        [Input("namespaceId")]
        public Input<int>? NamespaceId { get; set; }

        /// <summary>
        /// Set to true if you want allow merges only if all discussions are resolved.
        /// </summary>
        [Input("onlyAllowMergeIfAllDiscussionsAreResolved")]
        public Input<bool>? OnlyAllowMergeIfAllDiscussionsAreResolved { get; set; }

        /// <summary>
        /// Set to true if you want allow merges only if a pipeline succeeds.
        /// </summary>
        [Input("onlyAllowMergeIfPipelineSucceeds")]
        public Input<bool>? OnlyAllowMergeIfPipelineSucceeds { get; set; }

        /// <summary>
        /// The path of the repository.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Registration token to use during runner setup.
        /// </summary>
        [Input("runnersToken")]
        public Input<string>? RunnersToken { get; set; }

        /// <summary>
        /// Enable shared runners for this project.
        /// </summary>
        [Input("sharedRunnersEnabled")]
        public Input<bool>? SharedRunnersEnabled { get; set; }

        [Input("sharedWithGroups")]
        private InputList<Inputs.ProjectSharedWithGroupsGetArgs>? _sharedWithGroups;

        /// <summary>
        /// Enable sharing the project with a list of groups (maps).
        /// </summary>
        public InputList<Inputs.ProjectSharedWithGroupsGetArgs> SharedWithGroups
        {
            get => _sharedWithGroups ?? (_sharedWithGroups = new InputList<Inputs.ProjectSharedWithGroupsGetArgs>());
            set => _sharedWithGroups = value;
        }

        /// <summary>
        /// Enable snippets for the project.
        /// </summary>
        [Input("snippetsEnabled")]
        public Input<bool>? SnippetsEnabled { get; set; }

        /// <summary>
        /// URL that can be provided to `git clone` to clone the
        /// repository via SSH.
        /// </summary>
        [Input("sshUrlToRepo")]
        public Input<string>? SshUrlToRepo { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags (topics) of the project.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Set to `public` to create a public project.
        /// Valid values are `private`, `internal`, `public`.
        /// Repositories are created as private by default.
        /// </summary>
        [Input("visibilityLevel")]
        public Input<string>? VisibilityLevel { get; set; }

        /// <summary>
        /// URL that can be used to find the project in a browser.
        /// </summary>
        [Input("webUrl")]
        public Input<string>? WebUrl { get; set; }

        /// <summary>
        /// Enable wiki for the project.
        /// </summary>
        [Input("wikiEnabled")]
        public Input<bool>? WikiEnabled { get; set; }

        public ProjectState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ProjectSharedWithGroupsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group's sharing permissions. See [group members permission][group_members_permissions] for more info.
        /// Valid values are `guest`, `reporter`, `developer`, `master`.
        /// </summary>
        [Input("groupAccessLevel", required: true)]
        public Input<string> GroupAccessLevel { get; set; } = null!;

        /// <summary>
        /// Group id of the group you want to share the project with.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        /// <summary>
        /// Group's name.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        public ProjectSharedWithGroupsArgs()
        {
        }
    }

    public sealed class ProjectSharedWithGroupsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group's sharing permissions. See [group members permission][group_members_permissions] for more info.
        /// Valid values are `guest`, `reporter`, `developer`, `master`.
        /// </summary>
        [Input("groupAccessLevel", required: true)]
        public Input<string> GroupAccessLevel { get; set; } = null!;

        /// <summary>
        /// Group id of the group you want to share the project with.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        /// <summary>
        /// Group's name.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        public ProjectSharedWithGroupsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ProjectSharedWithGroups
    {
        /// <summary>
        /// Group's sharing permissions. See [group members permission][group_members_permissions] for more info.
        /// Valid values are `guest`, `reporter`, `developer`, `master`.
        /// </summary>
        public readonly string GroupAccessLevel;
        /// <summary>
        /// Group id of the group you want to share the project with.
        /// </summary>
        public readonly int GroupId;
        /// <summary>
        /// Group's name.
        /// </summary>
        public readonly string GroupName;

        [OutputConstructor]
        private ProjectSharedWithGroups(
            string groupAccessLevel,
            int groupId,
            string groupName)
        {
            GroupAccessLevel = groupAccessLevel;
            GroupId = groupId;
            GroupName = groupName;
        }
    }
    }
}
