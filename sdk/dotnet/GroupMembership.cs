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
    /// The `gitlab.GroupMembership` resource allows to manage the lifecycle of a users group membership.
    /// 
    /// &gt; If a group should grant membership to another group use the `gitlab.GroupShareGroup` resource instead.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new GitLab.GroupMembership("test", new()
    ///     {
    ///         AccessLevel = "guest",
    ///         ExpiresAt = "2020-12-31",
    ///         GroupId = "12345",
    ///         UserId = 1337,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitLab group membership can be imported using an id made up of `group_id:user_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/groupMembership:GroupMembership test "12345:1337"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/groupMembership:GroupMembership")]
    public partial class GroupMembership : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
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
        /// Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
        /// </summary>
        [Output("skipSubresourcesOnDestroy")]
        public Output<bool?> SkipSubresourcesOnDestroy { get; private set; } = null!;

        /// <summary>
        /// Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
        /// </summary>
        [Output("unassignIssuablesOnDestroy")]
        public Output<bool?> UnassignIssuablesOnDestroy { get; private set; } = null!;

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
            : base("gitlab:index/groupMembership:GroupMembership", name, args ?? new GroupMembershipArgs(), MakeResourceOptions(options, ""))
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

    public sealed class GroupMembershipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
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
        /// Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
        /// </summary>
        [Input("skipSubresourcesOnDestroy")]
        public Input<bool>? SkipSubresourcesOnDestroy { get; set; }

        /// <summary>
        /// Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
        /// </summary>
        [Input("unassignIssuablesOnDestroy")]
        public Input<bool>? UnassignIssuablesOnDestroy { get; set; }

        /// <summary>
        /// The id of the user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public GroupMembershipArgs()
        {
        }
        public static new GroupMembershipArgs Empty => new GroupMembershipArgs();
    }

    public sealed class GroupMembershipState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
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
        /// Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
        /// </summary>
        [Input("skipSubresourcesOnDestroy")]
        public Input<bool>? SkipSubresourcesOnDestroy { get; set; }

        /// <summary>
        /// Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
        /// </summary>
        [Input("unassignIssuablesOnDestroy")]
        public Input<bool>? UnassignIssuablesOnDestroy { get; set; }

        /// <summary>
        /// The id of the user.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public GroupMembershipState()
        {
        }
        public static new GroupMembershipState Empty => new GroupMembershipState();
    }
}
