// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    /// <summary>
    /// This resource allows you to add a user to an existing group.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/group_membership.html.markdown.
    /// </summary>
    public partial class GroupMembership : Pulumi.CustomResource
    {
        /// <summary>
        /// Acceptable values are: guest, reporter, developer, maintainer, owner.
        /// </summary>
        [Output("accessLevel")]
        public Output<string> AccessLevel { get; private set; } = null!;

        /// <summary>
        /// Expiration date for the group membership. Format: `YYYY-MM-DD`
        /// </summary>
        [Output("expiresAt")]
        public Output<string?> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The id of the group.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The id of the user.
        /// </summary>
        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMembership(string name, GroupMembershipArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupMembership:GroupMembership", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GroupMembership(string name, Input<string> id, GroupMembershipState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupMembership:GroupMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMembership Get(string name, Input<string> id, GroupMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMembership(name, id, state, options);
        }
    }

    public sealed class GroupMembershipArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Acceptable values are: guest, reporter, developer, maintainer, owner.
        /// </summary>
        [Input("accessLevel", required: true)]
        public Input<string> AccessLevel { get; set; } = null!;

        /// <summary>
        /// Expiration date for the group membership. Format: `YYYY-MM-DD`
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The id of the group.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The id of the user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public GroupMembershipArgs()
        {
        }
    }

    public sealed class GroupMembershipState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Acceptable values are: guest, reporter, developer, maintainer, owner.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// Expiration date for the group membership. Format: `YYYY-MM-DD`
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The id of the group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The id of the user.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public GroupMembershipState()
        {
        }
    }
}
