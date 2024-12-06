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
    public sealed class GetGroupBillableMemberMembershipsMembershipResult
    {
        /// <summary>
        /// Access-level of the member. For details see: https://docs.gitlab.com/ee/api/access_requests.html#valid-access-levels
        /// </summary>
        public readonly string AccessLevel;
        /// <summary>
        /// Datetime when the membership was first added.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Date when the membership will end.
        /// </summary>
        public readonly string ExpiresAt;
        /// <summary>
        /// The id of the membership.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Breadcrumb-style, full display-name of the group or project.
        /// </summary>
        public readonly string SourceFullName;
        /// <summary>
        /// The id of the group or project, the user is a (direct) member of.
        /// </summary>
        public readonly int SourceId;
        /// <summary>
        /// URL to the members-page of the group or project.
        /// </summary>
        public readonly string SourceMembersUrl;

        [OutputConstructor]
        private GetGroupBillableMemberMembershipsMembershipResult(
            string accessLevel,

            string createdAt,

            string expiresAt,

            int id,

            string sourceFullName,

            int sourceId,

            string sourceMembersUrl)
        {
            AccessLevel = accessLevel;
            CreatedAt = createdAt;
            ExpiresAt = expiresAt;
            Id = id;
            SourceFullName = sourceFullName;
            SourceId = sourceId;
            SourceMembersUrl = sourceMembersUrl;
        }
    }
}