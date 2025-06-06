// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProject
    {
        /// <summary>
        /// The `gitlab.Project` data source allows details of a project to be retrieved by either its ID or its path with namespace.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/projects/#get-a-single-project)
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("gitlab:index/getProject:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.Project` data source allows details of a project to be retrieved by either its ID or its path with namespace.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/projects/#get-a-single-project)
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectResult>("gitlab:index/getProject:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.Project` data source allows details of a project to be retrieved by either its ID or its path with namespace.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/projects/#get-a-single-project)
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectResult>("gitlab:index/getProject:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default number of revisions for shallow cloning.
        /// </summary>
        [Input("ciDefaultGitDepth")]
        public int? CiDefaultGitDepth { get; set; }

        [Input("ciIdTokenSubClaimComponents")]
        private List<string>? _ciIdTokenSubClaimComponents;

        /// <summary>
        /// Fields included in the sub claim of the ID Token. Accepts an array starting with project*path. The array might also include ref*type and ref. Defaults to ["project*path", "ref*type", "ref"]. Introduced in GitLab 17.10.
        /// </summary>
        public List<string> CiIdTokenSubClaimComponents
        {
            get => _ciIdTokenSubClaimComponents ?? (_ciIdTokenSubClaimComponents = new List<string>());
            set => _ciIdTokenSubClaimComponents = value;
        }

        /// <summary>
        /// The integer that uniquely identifies the project within the gitlab install.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The path of the repository with namespace.
        /// </summary>
        [Input("pathWithNamespace")]
        public string? PathWithNamespace { get; set; }

        /// <summary>
        /// If true, jobs can be viewed by non-project members.
        /// </summary>
        [Input("publicBuilds")]
        public bool? PublicBuilds { get; set; }

        public GetProjectArgs()
        {
        }
        public static new GetProjectArgs Empty => new GetProjectArgs();
    }

    public sealed class GetProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default number of revisions for shallow cloning.
        /// </summary>
        [Input("ciDefaultGitDepth")]
        public Input<int>? CiDefaultGitDepth { get; set; }

        [Input("ciIdTokenSubClaimComponents")]
        private InputList<string>? _ciIdTokenSubClaimComponents;

        /// <summary>
        /// Fields included in the sub claim of the ID Token. Accepts an array starting with project*path. The array might also include ref*type and ref. Defaults to ["project*path", "ref*type", "ref"]. Introduced in GitLab 17.10.
        /// </summary>
        public InputList<string> CiIdTokenSubClaimComponents
        {
            get => _ciIdTokenSubClaimComponents ?? (_ciIdTokenSubClaimComponents = new InputList<string>());
            set => _ciIdTokenSubClaimComponents = value;
        }

        /// <summary>
        /// The integer that uniquely identifies the project within the gitlab install.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The path of the repository with namespace.
        /// </summary>
        [Input("pathWithNamespace")]
        public Input<string>? PathWithNamespace { get; set; }

        /// <summary>
        /// If true, jobs can be viewed by non-project members.
        /// </summary>
        [Input("publicBuilds")]
        public Input<bool>? PublicBuilds { get; set; }

        public GetProjectInvokeArgs()
        {
        }
        public static new GetProjectInvokeArgs Empty => new GetProjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// Set whether or not a pipeline triggerer is allowed to approve deployments. Premium and Ultimate only.
        /// </summary>
        public readonly bool AllowPipelineTriggerApproveDeployment;
        /// <summary>
        /// Set the analytics access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string AnalyticsAccessLevel;
        /// <summary>
        /// Whether the project is in read-only mode (archived).
        /// </summary>
        public readonly bool Archived;
        /// <summary>
        /// Auto-cancel pending pipelines. This isn’t a boolean, but enabled/disabled.
        /// </summary>
        public readonly string AutoCancelPendingPipelines;
        /// <summary>
        /// Auto Deploy strategy. Valid values are `continuous`, `manual`, `timed_incremental`.
        /// </summary>
        public readonly string AutoDevopsDeployStrategy;
        /// <summary>
        /// Enable Auto DevOps for this project.
        /// </summary>
        public readonly bool AutoDevopsEnabled;
        /// <summary>
        /// Set whether auto-closing referenced issues on default branch.
        /// </summary>
        public readonly bool AutocloseReferencedIssues;
        /// <summary>
        /// The Git strategy. Defaults to fetch.
        /// </summary>
        public readonly string BuildGitStrategy;
        /// <summary>
        /// The maximum amount of time, in seconds, that a job can run.
        /// </summary>
        public readonly int BuildTimeout;
        /// <summary>
        /// Set the builds access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string BuildsAccessLevel;
        /// <summary>
        /// CI config file path for the project.
        /// </summary>
        public readonly string CiConfigPath;
        /// <summary>
        /// Default number of revisions for shallow cloning.
        /// </summary>
        public readonly int CiDefaultGitDepth;
        /// <summary>
        /// Pipelines older than the configured time are deleted.
        /// </summary>
        public readonly int CiDeletePipelinesInSeconds;
        /// <summary>
        /// Fields included in the sub claim of the ID Token. Accepts an array starting with project*path. The array might also include ref*type and ref. Defaults to ["project*path", "ref*type", "ref"]. Introduced in GitLab 17.10.
        /// </summary>
        public readonly ImmutableArray<string> CiIdTokenSubClaimComponents;
        /// <summary>
        /// The minimum role required to set variables when running pipelines and jobs. Introduced in GitLab 17.1. Valid values are `developer`, `maintainer`, `owner`, `no_one_allowed`
        /// </summary>
        public readonly string CiPipelineVariablesMinimumOverrideRole;
        /// <summary>
        /// The role required to cancel a pipeline or job. Premium and Ultimate only. Valid values are `developer`, `maintainer`, `no one`
        /// </summary>
        public readonly string CiRestrictPipelineCancellationRole;
        /// <summary>
        /// Use separate caches for protected branches.
        /// </summary>
        public readonly bool CiSeparatedCaches;
        /// <summary>
        /// Set the image cleanup policy for this project. **Note**: this field is sometimes named `container_expiration_policy_attributes` in the GitLab Upstream API.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectContainerExpirationPolicyResult> ContainerExpirationPolicies;
        /// <summary>
        /// Set visibility of container registry, for this project. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string ContainerRegistryAccessLevel;
        /// <summary>
        /// The default branch for the project.
        /// </summary>
        public readonly string DefaultBranch;
        /// <summary>
        /// A description of the project.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Enable email notifications.
        /// </summary>
        public readonly bool EmailsEnabled;
        /// <summary>
        /// Whether the project is empty.
        /// </summary>
        public readonly bool EmptyRepo;
        /// <summary>
        /// Set the environments access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string EnvironmentsAccessLevel;
        /// <summary>
        /// The classification label for the project.
        /// </summary>
        public readonly string ExternalAuthorizationClassificationLabel;
        /// <summary>
        /// Set the feature flags access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string FeatureFlagsAccessLevel;
        /// <summary>
        /// Set the forking access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string ForkingAccessLevel;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the
        /// </summary>
        public readonly string HttpUrlToRepo;
        /// <summary>
        /// The integer that uniquely identifies the project within the gitlab install.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// URL the project was imported from.
        /// </summary>
        public readonly string ImportUrl;
        /// <summary>
        /// Set the infrastructure access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string InfrastructureAccessLevel;
        /// <summary>
        /// Set the issues access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string IssuesAccessLevel;
        /// <summary>
        /// Enable issue tracking for the project.
        /// </summary>
        public readonly bool IssuesEnabled;
        /// <summary>
        /// Disable or enable the ability to keep the latest artifact for this project.
        /// </summary>
        public readonly bool KeepLatestArtifact;
        /// <summary>
        /// Enable LFS for the project.
        /// </summary>
        public readonly bool LfsEnabled;
        /// <summary>
        /// Template used to create merge commit message in merge requests.
        /// </summary>
        public readonly string MergeCommitTemplate;
        /// <summary>
        /// Enable or disable merge pipelines.
        /// </summary>
        public readonly bool MergePipelinesEnabled;
        /// <summary>
        /// Set the merge requests access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string MergeRequestsAccessLevel;
        /// <summary>
        /// Enable merge requests for the project.
        /// </summary>
        public readonly bool MergeRequestsEnabled;
        /// <summary>
        /// Enable or disable merge trains.
        /// </summary>
        public readonly bool MergeTrainsEnabled;
        /// <summary>
        /// The visibility of machine learning model experiments.
        /// </summary>
        public readonly string ModelExperimentsAccessLevel;
        /// <summary>
        /// The visibility of machine learning model registry.
        /// </summary>
        public readonly string ModelRegistryAccessLevel;
        /// <summary>
        /// Set the monitor access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string MonitorAccessLevel;
        /// <summary>
        /// The name of the project.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The namespace (group or user) of the project. Defaults to your user.
        /// </summary>
        public readonly int NamespaceId;
        /// <summary>
        /// The path of the repository.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The path of the repository with namespace.
        /// </summary>
        public readonly string PathWithNamespace;
        /// <summary>
        /// Enable pipelines for the project.
        /// </summary>
        public readonly bool PipelinesEnabled;
        /// <summary>
        /// Whether merge requests require an associated issue from Jira. Premium and Ultimate only.
        /// </summary>
        public readonly bool PreventMergeWithoutJiraIssue;
        /// <summary>
        /// Show link to create/view merge request when pushing from the command line
        /// </summary>
        public readonly bool PrintingMergeRequestLinkEnabled;
        /// <summary>
        /// If true, jobs can be viewed by non-project members.
        /// </summary>
        public readonly bool? PublicBuilds;
        /// <summary>
        /// Push rules for the project. Push rules are only available on Enterprise plans and if the authenticated has permissions to read them.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectPushRuleResult> PushRules;
        /// <summary>
        /// Set the releases access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string ReleasesAccessLevel;
        /// <summary>
        /// Enable `Delete source branch` option by default for all new merge requests
        /// </summary>
        public readonly bool RemoveSourceBranchAfterMerge;
        /// <summary>
        /// Set the repository access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string RepositoryAccessLevel;
        /// <summary>
        /// Which storage shard the repository is on. (administrator only)
        /// </summary>
        public readonly string RepositoryStorage;
        /// <summary>
        /// Allow users to request member access.
        /// </summary>
        public readonly bool RequestAccessEnabled;
        /// <summary>
        /// Set the requirements access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string RequirementsAccessLevel;
        /// <summary>
        /// Automatically resolve merge request diffs discussions on lines changed with a push.
        /// </summary>
        public readonly bool ResolveOutdatedDiffDiscussions;
        /// <summary>
        /// Allow only users with the Maintainer role to pass user-defined variables when triggering a pipeline.
        /// </summary>
        public readonly bool RestrictUserDefinedVariables;
        /// <summary>
        /// Registration token to use during runner setup.
        /// </summary>
        public readonly string RunnersToken;
        /// <summary>
        /// Set the security and compliance access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string SecurityAndComplianceAccessLevel;
        /// <summary>
        /// Describes groups which have access shared to this project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectSharedWithGroupResult> SharedWithGroups;
        /// <summary>
        /// Set the snippets access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string SnippetsAccessLevel;
        /// <summary>
        /// Enable snippets for the project.
        /// </summary>
        public readonly bool SnippetsEnabled;
        /// <summary>
        /// Template used to create squash commit message in merge requests.
        /// </summary>
        public readonly string SquashCommitTemplate;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the
        /// </summary>
        public readonly string SshUrlToRepo;
        /// <summary>
        /// The commit message used to apply merge request suggestions.
        /// </summary>
        public readonly string SuggestionCommitMessage;
        /// <summary>
        /// The list of topics for the project.
        /// </summary>
        public readonly ImmutableArray<string> Topics;
        /// <summary>
        /// Repositories are created as private by default.
        /// </summary>
        public readonly string VisibilityLevel;
        /// <summary>
        /// URL that can be used to find the project in a browser.
        /// </summary>
        public readonly string WebUrl;
        /// <summary>
        /// Set the wiki access level. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        public readonly string WikiAccessLevel;
        /// <summary>
        /// Enable wiki for the project.
        /// </summary>
        public readonly bool WikiEnabled;

        [OutputConstructor]
        private GetProjectResult(
            bool allowPipelineTriggerApproveDeployment,

            string analyticsAccessLevel,

            bool archived,

            string autoCancelPendingPipelines,

            string autoDevopsDeployStrategy,

            bool autoDevopsEnabled,

            bool autocloseReferencedIssues,

            string buildGitStrategy,

            int buildTimeout,

            string buildsAccessLevel,

            string ciConfigPath,

            int ciDefaultGitDepth,

            int ciDeletePipelinesInSeconds,

            ImmutableArray<string> ciIdTokenSubClaimComponents,

            string ciPipelineVariablesMinimumOverrideRole,

            string ciRestrictPipelineCancellationRole,

            bool ciSeparatedCaches,

            ImmutableArray<Outputs.GetProjectContainerExpirationPolicyResult> containerExpirationPolicies,

            string containerRegistryAccessLevel,

            string defaultBranch,

            string description,

            bool emailsEnabled,

            bool emptyRepo,

            string environmentsAccessLevel,

            string externalAuthorizationClassificationLabel,

            string featureFlagsAccessLevel,

            string forkingAccessLevel,

            string httpUrlToRepo,

            string id,

            string importUrl,

            string infrastructureAccessLevel,

            string issuesAccessLevel,

            bool issuesEnabled,

            bool keepLatestArtifact,

            bool lfsEnabled,

            string mergeCommitTemplate,

            bool mergePipelinesEnabled,

            string mergeRequestsAccessLevel,

            bool mergeRequestsEnabled,

            bool mergeTrainsEnabled,

            string modelExperimentsAccessLevel,

            string modelRegistryAccessLevel,

            string monitorAccessLevel,

            string name,

            int namespaceId,

            string path,

            string pathWithNamespace,

            bool pipelinesEnabled,

            bool preventMergeWithoutJiraIssue,

            bool printingMergeRequestLinkEnabled,

            bool? publicBuilds,

            ImmutableArray<Outputs.GetProjectPushRuleResult> pushRules,

            string releasesAccessLevel,

            bool removeSourceBranchAfterMerge,

            string repositoryAccessLevel,

            string repositoryStorage,

            bool requestAccessEnabled,

            string requirementsAccessLevel,

            bool resolveOutdatedDiffDiscussions,

            bool restrictUserDefinedVariables,

            string runnersToken,

            string securityAndComplianceAccessLevel,

            ImmutableArray<Outputs.GetProjectSharedWithGroupResult> sharedWithGroups,

            string snippetsAccessLevel,

            bool snippetsEnabled,

            string squashCommitTemplate,

            string sshUrlToRepo,

            string suggestionCommitMessage,

            ImmutableArray<string> topics,

            string visibilityLevel,

            string webUrl,

            string wikiAccessLevel,

            bool wikiEnabled)
        {
            AllowPipelineTriggerApproveDeployment = allowPipelineTriggerApproveDeployment;
            AnalyticsAccessLevel = analyticsAccessLevel;
            Archived = archived;
            AutoCancelPendingPipelines = autoCancelPendingPipelines;
            AutoDevopsDeployStrategy = autoDevopsDeployStrategy;
            AutoDevopsEnabled = autoDevopsEnabled;
            AutocloseReferencedIssues = autocloseReferencedIssues;
            BuildGitStrategy = buildGitStrategy;
            BuildTimeout = buildTimeout;
            BuildsAccessLevel = buildsAccessLevel;
            CiConfigPath = ciConfigPath;
            CiDefaultGitDepth = ciDefaultGitDepth;
            CiDeletePipelinesInSeconds = ciDeletePipelinesInSeconds;
            CiIdTokenSubClaimComponents = ciIdTokenSubClaimComponents;
            CiPipelineVariablesMinimumOverrideRole = ciPipelineVariablesMinimumOverrideRole;
            CiRestrictPipelineCancellationRole = ciRestrictPipelineCancellationRole;
            CiSeparatedCaches = ciSeparatedCaches;
            ContainerExpirationPolicies = containerExpirationPolicies;
            ContainerRegistryAccessLevel = containerRegistryAccessLevel;
            DefaultBranch = defaultBranch;
            Description = description;
            EmailsEnabled = emailsEnabled;
            EmptyRepo = emptyRepo;
            EnvironmentsAccessLevel = environmentsAccessLevel;
            ExternalAuthorizationClassificationLabel = externalAuthorizationClassificationLabel;
            FeatureFlagsAccessLevel = featureFlagsAccessLevel;
            ForkingAccessLevel = forkingAccessLevel;
            HttpUrlToRepo = httpUrlToRepo;
            Id = id;
            ImportUrl = importUrl;
            InfrastructureAccessLevel = infrastructureAccessLevel;
            IssuesAccessLevel = issuesAccessLevel;
            IssuesEnabled = issuesEnabled;
            KeepLatestArtifact = keepLatestArtifact;
            LfsEnabled = lfsEnabled;
            MergeCommitTemplate = mergeCommitTemplate;
            MergePipelinesEnabled = mergePipelinesEnabled;
            MergeRequestsAccessLevel = mergeRequestsAccessLevel;
            MergeRequestsEnabled = mergeRequestsEnabled;
            MergeTrainsEnabled = mergeTrainsEnabled;
            ModelExperimentsAccessLevel = modelExperimentsAccessLevel;
            ModelRegistryAccessLevel = modelRegistryAccessLevel;
            MonitorAccessLevel = monitorAccessLevel;
            Name = name;
            NamespaceId = namespaceId;
            Path = path;
            PathWithNamespace = pathWithNamespace;
            PipelinesEnabled = pipelinesEnabled;
            PreventMergeWithoutJiraIssue = preventMergeWithoutJiraIssue;
            PrintingMergeRequestLinkEnabled = printingMergeRequestLinkEnabled;
            PublicBuilds = publicBuilds;
            PushRules = pushRules;
            ReleasesAccessLevel = releasesAccessLevel;
            RemoveSourceBranchAfterMerge = removeSourceBranchAfterMerge;
            RepositoryAccessLevel = repositoryAccessLevel;
            RepositoryStorage = repositoryStorage;
            RequestAccessEnabled = requestAccessEnabled;
            RequirementsAccessLevel = requirementsAccessLevel;
            ResolveOutdatedDiffDiscussions = resolveOutdatedDiffDiscussions;
            RestrictUserDefinedVariables = restrictUserDefinedVariables;
            RunnersToken = runnersToken;
            SecurityAndComplianceAccessLevel = securityAndComplianceAccessLevel;
            SharedWithGroups = sharedWithGroups;
            SnippetsAccessLevel = snippetsAccessLevel;
            SnippetsEnabled = snippetsEnabled;
            SquashCommitTemplate = squashCommitTemplate;
            SshUrlToRepo = sshUrlToRepo;
            SuggestionCommitMessage = suggestionCommitMessage;
            Topics = topics;
            VisibilityLevel = visibilityLevel;
            WebUrl = webUrl;
            WikiAccessLevel = wikiAccessLevel;
            WikiEnabled = wikiEnabled;
        }
    }
}
