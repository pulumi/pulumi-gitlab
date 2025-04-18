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
    /// The `gitlab.GroupShareGroup` resource allows managing the lifecycle of a group shared with another group.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#share-groups-with-groups)
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
    ///     var test = new GitLab.GroupShareGroup("test", new()
    ///     {
    ///         GroupId = foo.Id,
    ///         ShareGroupId = bar.Id,
    ///         GroupAccess = "guest",
    ///         ExpiresAt = "2099-01-01",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_share_group`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_group_share_group.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// GitLab group shares can be imported using an id made up of `mainGroupId:shareGroupId`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/groupShareGroup:GroupShareGroup test 12345:1337
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/groupShareGroup:GroupShareGroup")]
    public partial class GroupShareGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Share expiration date. Format: `YYYY-MM-DD`
        /// </summary>
        [Output("expiresAt")]
        public Output<string?> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
        /// </summary>
        [Output("groupAccess")]
        public Output<string> GroupAccess { get; private set; } = null!;

        /// <summary>
        /// The id of the main group to be shared.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of a custom member role. Only available for Ultimate instances.
        /// </summary>
        [Output("memberRoleId")]
        public Output<int> MemberRoleId { get; private set; } = null!;

        /// <summary>
        /// The id of the additional group with which the main group will be shared.
        /// </summary>
        [Output("shareGroupId")]
        public Output<int> ShareGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupShareGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupShareGroup(string name, GroupShareGroupArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupShareGroup:GroupShareGroup", name, args ?? new GroupShareGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupShareGroup(string name, Input<string> id, GroupShareGroupState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupShareGroup:GroupShareGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupShareGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupShareGroup Get(string name, Input<string> id, GroupShareGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupShareGroup(name, id, state, options);
        }
    }

    public sealed class GroupShareGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Share expiration date. Format: `YYYY-MM-DD`
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
        /// </summary>
        [Input("groupAccess", required: true)]
        public Input<string> GroupAccess { get; set; } = null!;

        /// <summary>
        /// The id of the main group to be shared.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The ID of a custom member role. Only available for Ultimate instances.
        /// </summary>
        [Input("memberRoleId")]
        public Input<int>? MemberRoleId { get; set; }

        /// <summary>
        /// The id of the additional group with which the main group will be shared.
        /// </summary>
        [Input("shareGroupId", required: true)]
        public Input<int> ShareGroupId { get; set; } = null!;

        public GroupShareGroupArgs()
        {
        }
        public static new GroupShareGroupArgs Empty => new GroupShareGroupArgs();
    }

    public sealed class GroupShareGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Share expiration date. Format: `YYYY-MM-DD`
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
        /// </summary>
        [Input("groupAccess")]
        public Input<string>? GroupAccess { get; set; }

        /// <summary>
        /// The id of the main group to be shared.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The ID of a custom member role. Only available for Ultimate instances.
        /// </summary>
        [Input("memberRoleId")]
        public Input<int>? MemberRoleId { get; set; }

        /// <summary>
        /// The id of the additional group with which the main group will be shared.
        /// </summary>
        [Input("shareGroupId")]
        public Input<int>? ShareGroupId { get; set; }

        public GroupShareGroupState()
        {
        }
        public static new GroupShareGroupState Empty => new GroupShareGroupState();
    }
}
