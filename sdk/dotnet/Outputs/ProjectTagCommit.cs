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
    public sealed class ProjectTagCommit
    {
        /// <summary>
        /// The email of the author.
        /// </summary>
        public readonly string? AuthorEmail;
        /// <summary>
        /// The name of the author.
        /// </summary>
        public readonly string? AuthorName;
        /// <summary>
        /// The date which the commit was authored (format: yyyy-MM-ddTHH:mm:ssZ).
        /// </summary>
        public readonly string? AuthoredDate;
        /// <summary>
        /// The date at which the commit was pushed (format: yyyy-MM-ddTHH:mm:ssZ).
        /// </summary>
        public readonly string? CommittedDate;
        /// <summary>
        /// The email of the user that committed.
        /// </summary>
        public readonly string? CommitterEmail;
        /// <summary>
        /// The name of the user that committed.
        /// </summary>
        public readonly string? CommitterName;
        /// <summary>
        /// The unique id assigned to the commit by Gitlab.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The commit message
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// The id of the parents of the commit
        /// </summary>
        public readonly ImmutableArray<string> ParentIds;
        /// <summary>
        /// The short id assigned to the commit by Gitlab.
        /// </summary>
        public readonly string? ShortId;
        /// <summary>
        /// The title of the commit
        /// </summary>
        public readonly string? Title;

        [OutputConstructor]
        private ProjectTagCommit(
            string? authorEmail,

            string? authorName,

            string? authoredDate,

            string? committedDate,

            string? committerEmail,

            string? committerName,

            string? id,

            string? message,

            ImmutableArray<string> parentIds,

            string? shortId,

            string? title)
        {
            AuthorEmail = authorEmail;
            AuthorName = authorName;
            AuthoredDate = authoredDate;
            CommittedDate = committedDate;
            CommitterEmail = committerEmail;
            CommitterName = committerName;
            Id = id;
            Message = message;
            ParentIds = parentIds;
            ShortId = shortId;
            Title = title;
        }
    }
}
