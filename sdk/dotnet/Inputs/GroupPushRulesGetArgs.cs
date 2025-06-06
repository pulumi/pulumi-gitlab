// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class GroupPushRulesGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// All commit author emails must match this regex, e.g. `@my-company.com$`.
        /// </summary>
        [Input("authorEmailRegex")]
        public Input<string>? AuthorEmailRegex { get; set; }

        /// <summary>
        /// All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
        /// </summary>
        [Input("branchNameRegex")]
        public Input<string>? BranchNameRegex { get; set; }

        /// <summary>
        /// Only commits pushed using verified emails are allowed.
        /// </summary>
        [Input("commitCommitterCheck")]
        public Input<bool>? CommitCommitterCheck { get; set; }

        /// <summary>
        /// Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
        /// </summary>
        [Input("commitCommitterNameCheck")]
        public Input<bool>? CommitCommitterNameCheck { get; set; }

        /// <summary>
        /// No commit message is allowed to match this regex, for example `ssh\:\/\/`.
        /// </summary>
        [Input("commitMessageNegativeRegex")]
        public Input<string>? CommitMessageNegativeRegex { get; set; }

        /// <summary>
        /// All commit messages must match this regex, e.g. `Fixed \d+\..*`.
        /// </summary>
        [Input("commitMessageRegex")]
        public Input<string>? CommitMessageRegex { get; set; }

        /// <summary>
        /// Deny deleting a tag.
        /// </summary>
        [Input("denyDeleteTag")]
        public Input<bool>? DenyDeleteTag { get; set; }

        /// <summary>
        /// Filenames matching the regular expression provided in this attribute are not allowed, for example, `(jar|exe)$`.
        /// </summary>
        [Input("fileNameRegex")]
        public Input<string>? FileNameRegex { get; set; }

        /// <summary>
        /// Maximum file size (MB) allowed.
        /// </summary>
        [Input("maxFileSize")]
        public Input<int>? MaxFileSize { get; set; }

        /// <summary>
        /// Allows only GitLab users to author commits.
        /// </summary>
        [Input("memberCheck")]
        public Input<bool>? MemberCheck { get; set; }

        /// <summary>
        /// GitLab will reject any files that are likely to contain secrets.
        /// </summary>
        [Input("preventSecrets")]
        public Input<bool>? PreventSecrets { get; set; }

        /// <summary>
        /// Reject commit when it’s not DCO certified.
        /// </summary>
        [Input("rejectNonDcoCommits")]
        public Input<bool>? RejectNonDcoCommits { get; set; }

        /// <summary>
        /// Only commits signed through GPG are allowed.
        /// </summary>
        [Input("rejectUnsignedCommits")]
        public Input<bool>? RejectUnsignedCommits { get; set; }

        public GroupPushRulesGetArgs()
        {
        }
        public static new GroupPushRulesGetArgs Empty => new GroupPushRulesGetArgs();
    }
}
