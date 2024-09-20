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
    public sealed class ProjectPushRules
    {
        /// <summary>
        /// All commit author emails must match this regex, e.g. `@my-company.com$`.
        /// </summary>
        public readonly string? AuthorEmailRegex;
        /// <summary>
        /// All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
        /// </summary>
        public readonly string? BranchNameRegex;
        /// <summary>
        /// Users can only push commits to this repository that were committed with one of their own verified emails.
        /// </summary>
        public readonly bool? CommitCommitterCheck;
        /// <summary>
        /// Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
        /// </summary>
        public readonly bool? CommitCommitterNameCheck;
        /// <summary>
        /// No commit message is allowed to match this regex, e.g. `ssh\:\/\/`.
        /// </summary>
        public readonly string? CommitMessageNegativeRegex;
        /// <summary>
        /// All commit messages must match this regex, e.g. `Fixed \d+\..*`.
        /// </summary>
        public readonly string? CommitMessageRegex;
        /// <summary>
        /// Deny deleting a tag.
        /// </summary>
        public readonly bool? DenyDeleteTag;
        /// <summary>
        /// All committed filenames must not match this regex, e.g. `(jar|exe)$`.
        /// </summary>
        public readonly string? FileNameRegex;
        /// <summary>
        /// Maximum file size (MB).
        /// </summary>
        public readonly int? MaxFileSize;
        /// <summary>
        /// Restrict commits by author (email) to existing GitLab users.
        /// </summary>
        public readonly bool? MemberCheck;
        /// <summary>
        /// GitLab will reject any files that are likely to contain secrets.
        /// </summary>
        public readonly bool? PreventSecrets;
        /// <summary>
        /// Reject commit when it’s not DCO certified.
        /// </summary>
        public readonly bool? RejectNonDcoCommits;
        /// <summary>
        /// Reject commit when it’s not signed through GPG.
        /// </summary>
        public readonly bool? RejectUnsignedCommits;

        [OutputConstructor]
        private ProjectPushRules(
            string? authorEmailRegex,

            string? branchNameRegex,

            bool? commitCommitterCheck,

            bool? commitCommitterNameCheck,

            string? commitMessageNegativeRegex,

            string? commitMessageRegex,

            bool? denyDeleteTag,

            string? fileNameRegex,

            int? maxFileSize,

            bool? memberCheck,

            bool? preventSecrets,

            bool? rejectNonDcoCommits,

            bool? rejectUnsignedCommits)
        {
            AuthorEmailRegex = authorEmailRegex;
            BranchNameRegex = branchNameRegex;
            CommitCommitterCheck = commitCommitterCheck;
            CommitCommitterNameCheck = commitCommitterNameCheck;
            CommitMessageNegativeRegex = commitMessageNegativeRegex;
            CommitMessageRegex = commitMessageRegex;
            DenyDeleteTag = denyDeleteTag;
            FileNameRegex = fileNameRegex;
            MaxFileSize = maxFileSize;
            MemberCheck = memberCheck;
            PreventSecrets = preventSecrets;
            RejectNonDcoCommits = rejectNonDcoCommits;
            RejectUnsignedCommits = rejectUnsignedCommits;
        }
    }
}
