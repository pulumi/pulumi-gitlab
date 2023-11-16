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
    public sealed class GroupPushRules
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
        /// Only commits pushed using verified emails are allowed.  **Note** This attribute is only supported in GitLab versions &gt;= 16.4.
        /// </summary>
        public readonly bool? CommitCommitterCheck;
        /// <summary>
        /// No commit message is allowed to match this regex, for example `ssh\:\/\/`.
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
        /// Filenames matching the regular expression provided in this attribute are not allowed, for example, `(jar|exe)$`.
        /// </summary>
        public readonly string? FileNameRegex;
        /// <summary>
        /// Maximum file size (MB) allowed.
        /// </summary>
        public readonly int? MaxFileSize;
        /// <summary>
        /// Allows only GitLab users to author commits.
        /// </summary>
        public readonly bool? MemberCheck;
        /// <summary>
        /// GitLab will reject any files that are likely to contain secrets.
        /// </summary>
        public readonly bool? PreventSecrets;
        /// <summary>
        /// Only commits signed through GPG are allowed.  **Note** This attribute is only supported in GitLab versions &gt;= 16.4.
        /// </summary>
        public readonly bool? RejectUnsignedCommits;

        [OutputConstructor]
        private GroupPushRules(
            string? authorEmailRegex,

            string? branchNameRegex,

            bool? commitCommitterCheck,

            string? commitMessageNegativeRegex,

            string? commitMessageRegex,

            bool? denyDeleteTag,

            string? fileNameRegex,

            int? maxFileSize,

            bool? memberCheck,

            bool? preventSecrets,

            bool? rejectUnsignedCommits)
        {
            AuthorEmailRegex = authorEmailRegex;
            BranchNameRegex = branchNameRegex;
            CommitCommitterCheck = commitCommitterCheck;
            CommitMessageNegativeRegex = commitMessageNegativeRegex;
            CommitMessageRegex = commitMessageRegex;
            DenyDeleteTag = denyDeleteTag;
            FileNameRegex = fileNameRegex;
            MaxFileSize = maxFileSize;
            MemberCheck = memberCheck;
            PreventSecrets = preventSecrets;
            RejectUnsignedCommits = rejectUnsignedCommits;
        }
    }
}
