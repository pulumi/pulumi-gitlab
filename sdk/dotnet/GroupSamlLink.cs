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
    /// The `gitlab.GroupSamlLink` resource allows to manage the lifecycle of an SAML integration with a group.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#saml-group-links)
    /// 
    /// ## Import
    /// 
    /// GitLab group saml links can be imported using an id made up of `group_id:saml_group_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/groupSamlLink:GroupSamlLink test "12345:samlgroupname1"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/groupSamlLink:GroupSamlLink")]
    public partial class GroupSamlLink : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        /// </summary>
        [Output("accessLevel")]
        public Output<string> AccessLevel { get; private set; } = null!;

        /// <summary>
        /// The ID or path of the group to add the SAML Group Link to.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The name of the SAML group.
        /// </summary>
        [Output("samlGroupName")]
        public Output<string> SamlGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a GroupSamlLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupSamlLink(string name, GroupSamlLinkArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupSamlLink:GroupSamlLink", name, args ?? new GroupSamlLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupSamlLink(string name, Input<string> id, GroupSamlLinkState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupSamlLink:GroupSamlLink", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupSamlLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupSamlLink Get(string name, Input<string> id, GroupSamlLinkState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupSamlLink(name, id, state, options);
        }
    }

    public sealed class GroupSamlLinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        /// </summary>
        [Input("accessLevel", required: true)]
        public Input<string> AccessLevel { get; set; } = null!;

        /// <summary>
        /// The ID or path of the group to add the SAML Group Link to.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The name of the SAML group.
        /// </summary>
        [Input("samlGroupName", required: true)]
        public Input<string> SamlGroupName { get; set; } = null!;

        public GroupSamlLinkArgs()
        {
        }
        public static new GroupSamlLinkArgs Empty => new GroupSamlLinkArgs();
    }

    public sealed class GroupSamlLinkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// The ID or path of the group to add the SAML Group Link to.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The name of the SAML group.
        /// </summary>
        [Input("samlGroupName")]
        public Input<string>? SamlGroupName { get; set; }

        public GroupSamlLinkState()
        {
        }
        public static new GroupSamlLinkState Empty => new GroupSamlLinkState();
    }
}