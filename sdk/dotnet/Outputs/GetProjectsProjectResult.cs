// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Outputs
{

    [OutputType]
    public sealed class GetProjectsProjectResult
    {
        public readonly ImmutableDictionary<string, string> _links;
        public readonly bool AllowMergeOnSkippedPipeline;
        public readonly int ApprovalsBeforeMerge;
        /// <summary>
        /// Limit by archived status.
        /// </summary>
        public readonly bool Archived;
        public readonly string AvatarUrl;
        public readonly string BuildCoverageRegex;
        public readonly string CiConfigPath;
        public readonly bool ContainerRegistryEnabled;
        public readonly string CreatedAt;
        public readonly int CreatorId;
        public readonly ImmutableArray<ImmutableDictionary<string, object>> CustomAttributes;
        public readonly string DefaultBranch;
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetProjectsProjectForkedFromProjectResult> ForkedFromProjects;
        public readonly int ForksCount;
        public readonly string HttpUrlToRepo;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly int Id;
        public readonly string ImportError;
        public readonly string ImportStatus;
        public readonly bool IssuesEnabled;
        public readonly bool JobsEnabled;
        public readonly string LastActivityAt;
        public readonly bool LfsEnabled;
        public readonly string MergeMethod;
        public readonly bool MergeRequestsEnabled;
        public readonly bool Mirror;
        public readonly bool MirrorOverwritesDivergedBranches;
        public readonly bool MirrorTriggerBuilds;
        public readonly int MirrorUserId;
        public readonly string Name;
        public readonly string NameWithNamespace;
        public readonly ImmutableArray<Outputs.GetProjectsProjectNamespaceResult> Namespaces;
        public readonly bool OnlyAllowMergeIfAllDiscussionsAreResolved;
        public readonly bool OnlyAllowMergeIfPipelineSucceeds;
        public readonly bool OnlyMirrorProtectedBranches;
        public readonly int OpenIssuesCount;
        public readonly ImmutableArray<Outputs.GetProjectsProjectOwnerResult> Owners;
        public readonly bool PackagesEnabled;
        public readonly string Path;
        public readonly string PathWithNamespace;
        public readonly ImmutableArray<Outputs.GetProjectsProjectPermissionResult> Permissions;
        public readonly bool Public;
        public readonly bool PublicBuilds;
        public readonly string ReadmeUrl;
        public readonly bool RequestAccessEnabled;
        public readonly bool ResolveOutdatedDiffDiscussions;
        public readonly string RunnersToken;
        public readonly bool SharedRunnersEnabled;
        public readonly ImmutableArray<Outputs.GetProjectsProjectSharedWithGroupResult> SharedWithGroups;
        public readonly bool SnippetsEnabled;
        public readonly string SshUrlToRepo;
        public readonly int StarCount;
        /// <summary>
        /// Include project statistics. Cannot be used with `group_id`.
        /// </summary>
        public readonly ImmutableDictionary<string, int> Statistics;
        public readonly ImmutableArray<string> TagLists;
        /// <summary>
        /// Limit by visibility `public`, `internal`, or `private`.
        /// </summary>
        public readonly string Visibility;
        public readonly string WebUrl;
        public readonly bool WikiEnabled;

        [OutputConstructor]
        private GetProjectsProjectResult(
            ImmutableDictionary<string, string> _links,

            bool allowMergeOnSkippedPipeline,

            int approvalsBeforeMerge,

            bool archived,

            string avatarUrl,

            string buildCoverageRegex,

            string ciConfigPath,

            bool containerRegistryEnabled,

            string createdAt,

            int creatorId,

            ImmutableArray<ImmutableDictionary<string, object>> customAttributes,

            string defaultBranch,

            string description,

            ImmutableArray<Outputs.GetProjectsProjectForkedFromProjectResult> forkedFromProjects,

            int forksCount,

            string httpUrlToRepo,

            int id,

            string importError,

            string importStatus,

            bool issuesEnabled,

            bool jobsEnabled,

            string lastActivityAt,

            bool lfsEnabled,

            string mergeMethod,

            bool mergeRequestsEnabled,

            bool mirror,

            bool mirrorOverwritesDivergedBranches,

            bool mirrorTriggerBuilds,

            int mirrorUserId,

            string name,

            string nameWithNamespace,

            ImmutableArray<Outputs.GetProjectsProjectNamespaceResult> namespaces,

            bool onlyAllowMergeIfAllDiscussionsAreResolved,

            bool onlyAllowMergeIfPipelineSucceeds,

            bool onlyMirrorProtectedBranches,

            int openIssuesCount,

            ImmutableArray<Outputs.GetProjectsProjectOwnerResult> owners,

            bool packagesEnabled,

            string path,

            string pathWithNamespace,

            ImmutableArray<Outputs.GetProjectsProjectPermissionResult> permissions,

            bool @public,

            bool publicBuilds,

            string readmeUrl,

            bool requestAccessEnabled,

            bool resolveOutdatedDiffDiscussions,

            string runnersToken,

            bool sharedRunnersEnabled,

            ImmutableArray<Outputs.GetProjectsProjectSharedWithGroupResult> sharedWithGroups,

            bool snippetsEnabled,

            string sshUrlToRepo,

            int starCount,

            ImmutableDictionary<string, int> statistics,

            ImmutableArray<string> tagLists,

            string visibility,

            string webUrl,

            bool wikiEnabled)
        {
            this._links = _links;
            AllowMergeOnSkippedPipeline = allowMergeOnSkippedPipeline;
            ApprovalsBeforeMerge = approvalsBeforeMerge;
            Archived = archived;
            AvatarUrl = avatarUrl;
            BuildCoverageRegex = buildCoverageRegex;
            CiConfigPath = ciConfigPath;
            ContainerRegistryEnabled = containerRegistryEnabled;
            CreatedAt = createdAt;
            CreatorId = creatorId;
            CustomAttributes = customAttributes;
            DefaultBranch = defaultBranch;
            Description = description;
            ForkedFromProjects = forkedFromProjects;
            ForksCount = forksCount;
            HttpUrlToRepo = httpUrlToRepo;
            Id = id;
            ImportError = importError;
            ImportStatus = importStatus;
            IssuesEnabled = issuesEnabled;
            JobsEnabled = jobsEnabled;
            LastActivityAt = lastActivityAt;
            LfsEnabled = lfsEnabled;
            MergeMethod = mergeMethod;
            MergeRequestsEnabled = mergeRequestsEnabled;
            Mirror = mirror;
            MirrorOverwritesDivergedBranches = mirrorOverwritesDivergedBranches;
            MirrorTriggerBuilds = mirrorTriggerBuilds;
            MirrorUserId = mirrorUserId;
            Name = name;
            NameWithNamespace = nameWithNamespace;
            Namespaces = namespaces;
            OnlyAllowMergeIfAllDiscussionsAreResolved = onlyAllowMergeIfAllDiscussionsAreResolved;
            OnlyAllowMergeIfPipelineSucceeds = onlyAllowMergeIfPipelineSucceeds;
            OnlyMirrorProtectedBranches = onlyMirrorProtectedBranches;
            OpenIssuesCount = openIssuesCount;
            Owners = owners;
            PackagesEnabled = packagesEnabled;
            Path = path;
            PathWithNamespace = pathWithNamespace;
            Permissions = permissions;
            Public = @public;
            PublicBuilds = publicBuilds;
            ReadmeUrl = readmeUrl;
            RequestAccessEnabled = requestAccessEnabled;
            ResolveOutdatedDiffDiscussions = resolveOutdatedDiffDiscussions;
            RunnersToken = runnersToken;
            SharedRunnersEnabled = sharedRunnersEnabled;
            SharedWithGroups = sharedWithGroups;
            SnippetsEnabled = snippetsEnabled;
            SshUrlToRepo = sshUrlToRepo;
            StarCount = starCount;
            Statistics = statistics;
            TagLists = tagLists;
            Visibility = visibility;
            WebUrl = webUrl;
            WikiEnabled = wikiEnabled;
        }
    }
}
