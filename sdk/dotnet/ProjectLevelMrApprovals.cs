// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    [GitLabResourceType("gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals")]
    public partial class ProjectLevelMrApprovals : Pulumi.CustomResource
    {
        /// <summary>
        /// By default, users are able to edit the approval rules in merge requests. If set to true,
        /// the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
        /// </summary>
        [Output("disableOverridingApproversPerMergeRequest")]
        public Output<bool?> DisableOverridingApproversPerMergeRequest { get; private set; } = null!;

        /// <summary>
        /// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
        /// also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
        /// </summary>
        [Output("mergeRequestsAuthorApproval")]
        public Output<bool?> MergeRequestsAuthorApproval { get; private set; } = null!;

        /// <summary>
        /// Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
        /// </summary>
        [Output("mergeRequestsDisableCommittersApproval")]
        public Output<bool?> MergeRequestsDisableCommittersApproval { get; private set; } = null!;

        /// <summary>
        /// The ID of the project to change MR approval configuration.
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        /// </summary>
        [Output("resetApprovalsOnPush")]
        public Output<bool?> ResetApprovalsOnPush { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectLevelMrApprovals resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectLevelMrApprovals(string name, ProjectLevelMrApprovalsArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals", name, args ?? new ProjectLevelMrApprovalsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectLevelMrApprovals(string name, Input<string> id, ProjectLevelMrApprovalsState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectLevelMrApprovals resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectLevelMrApprovals Get(string name, Input<string> id, ProjectLevelMrApprovalsState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectLevelMrApprovals(name, id, state, options);
        }
    }

    public sealed class ProjectLevelMrApprovalsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// By default, users are able to edit the approval rules in merge requests. If set to true,
        /// the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
        /// </summary>
        [Input("disableOverridingApproversPerMergeRequest")]
        public Input<bool>? DisableOverridingApproversPerMergeRequest { get; set; }

        /// <summary>
        /// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
        /// also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
        /// </summary>
        [Input("mergeRequestsAuthorApproval")]
        public Input<bool>? MergeRequestsAuthorApproval { get; set; }

        /// <summary>
        /// Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
        /// </summary>
        [Input("mergeRequestsDisableCommittersApproval")]
        public Input<bool>? MergeRequestsDisableCommittersApproval { get; set; }

        /// <summary>
        /// The ID of the project to change MR approval configuration.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        /// </summary>
        [Input("resetApprovalsOnPush")]
        public Input<bool>? ResetApprovalsOnPush { get; set; }

        public ProjectLevelMrApprovalsArgs()
        {
        }
    }

    public sealed class ProjectLevelMrApprovalsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// By default, users are able to edit the approval rules in merge requests. If set to true,
        /// the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
        /// </summary>
        [Input("disableOverridingApproversPerMergeRequest")]
        public Input<bool>? DisableOverridingApproversPerMergeRequest { get; set; }

        /// <summary>
        /// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
        /// also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
        /// </summary>
        [Input("mergeRequestsAuthorApproval")]
        public Input<bool>? MergeRequestsAuthorApproval { get; set; }

        /// <summary>
        /// Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
        /// </summary>
        [Input("mergeRequestsDisableCommittersApproval")]
        public Input<bool>? MergeRequestsDisableCommittersApproval { get; set; }

        /// <summary>
        /// The ID of the project to change MR approval configuration.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        /// </summary>
        [Input("resetApprovalsOnPush")]
        public Input<bool>? ResetApprovalsOnPush { get; set; }

        public ProjectLevelMrApprovalsState()
        {
        }
    }
}
