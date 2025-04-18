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
    public sealed class GetProjectBranchesBranchResult
    {
        /// <summary>
        /// Bool, true if you can push to the branch.
        /// </summary>
        public readonly bool CanPush;
        /// <summary>
        /// The commit associated with this branch.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectBranchesBranchCommitResult> Commits;
        /// <summary>
        /// Bool, true if branch is the default branch for the project.
        /// </summary>
        public readonly bool Default;
        /// <summary>
        /// Bool, true if developer level access allows to merge branch.
        /// </summary>
        public readonly bool DevelopersCanMerge;
        /// <summary>
        /// Bool, true if developer level access allows git push.
        /// </summary>
        public readonly bool DevelopersCanPush;
        /// <summary>
        /// Bool, true if the branch has been merged into its parent.
        /// </summary>
        public readonly bool Merged;
        /// <summary>
        /// The name of the branch.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Bool, true if branch has branch protection.
        /// </summary>
        public readonly bool Protected;
        /// <summary>
        /// URL that can be used to find the branch in a browser.
        /// </summary>
        public readonly string WebUrl;

        [OutputConstructor]
        private GetProjectBranchesBranchResult(
            bool canPush,

            ImmutableArray<Outputs.GetProjectBranchesBranchCommitResult> commits,

            bool @default,

            bool developersCanMerge,

            bool developersCanPush,

            bool merged,

            string name,

            bool @protected,

            string webUrl)
        {
            CanPush = canPush;
            Commits = commits;
            Default = @default;
            DevelopersCanMerge = developersCanMerge;
            DevelopersCanPush = developersCanPush;
            Merged = merged;
            Name = name;
            Protected = @protected;
            WebUrl = webUrl;
        }
    }
}
