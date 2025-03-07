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
    /// The `gitlab.GroupServiceAccount` resource allows creating a GitLab group service account.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/group_service_accounts/)
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
    ///     // This must be a top-level group
    ///     var example = new GitLab.Group("example", new()
    ///     {
    ///         Name = "example",
    ///         Path = "example",
    ///         Description = "An example group",
    ///     });
    /// 
    ///     // The service account against the top-level group
    ///     var exampleSa = new GitLab.GroupServiceAccount("example_sa", new()
    ///     {
    ///         Group = example.Id,
    ///         Name = "example-name",
    ///         Username = "example-username",
    ///     });
    /// 
    ///     // Group to assign the service account to. Can be the same top-level group resource as above, or a subgroup of that group.
    ///     var exampleSubgroup = new GitLab.Group("example_subgroup", new()
    ///     {
    ///         Name = "subgroup",
    ///         Path = "example/subgroup",
    ///         Description = "An example subgroup",
    ///     });
    /// 
    ///     // To assign the service account to a group
    ///     var exampleMembership = new GitLab.GroupMembership("example_membership", new()
    ///     {
    ///         GroupId = exampleSubgroup.Id,
    ///         UserId = exampleSa.ServiceAccountId,
    ///         AccessLevel = "developer",
    ///         ExpiresAt = "2020-03-14",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_service_account`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_group_service_account.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/groupServiceAccount:GroupServiceAccount You can import a group service account using `&lt;resource&gt; &lt;id&gt;`. The
    /// ```
    /// 
    /// `id` is in the form of &lt;group_id&gt;:&lt;service_account_id&gt;
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/groupServiceAccount:GroupServiceAccount example example
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/groupServiceAccount:GroupServiceAccount")]
    public partial class GroupServiceAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID or URL-encoded path of the group that the service account is created in. Must be a top level group.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The name of the user. If not specified, the default Service account user name is used.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The service account id.
        /// </summary>
        [Output("serviceAccountId")]
        public Output<string> ServiceAccountId { get; private set; } = null!;

        /// <summary>
        /// The username of the user. If not specified, it’s automatically generated.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a GroupServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupServiceAccount(string name, GroupServiceAccountArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupServiceAccount:GroupServiceAccount", name, args ?? new GroupServiceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupServiceAccount(string name, Input<string> id, GroupServiceAccountState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupServiceAccount:GroupServiceAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupServiceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupServiceAccount Get(string name, Input<string> id, GroupServiceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupServiceAccount(name, id, state, options);
        }
    }

    public sealed class GroupServiceAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or URL-encoded path of the group that the service account is created in. Must be a top level group.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The name of the user. If not specified, the default Service account user name is used.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The username of the user. If not specified, it’s automatically generated.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public GroupServiceAccountArgs()
        {
        }
        public static new GroupServiceAccountArgs Empty => new GroupServiceAccountArgs();
    }

    public sealed class GroupServiceAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or URL-encoded path of the group that the service account is created in. Must be a top level group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The name of the user. If not specified, the default Service account user name is used.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The service account id.
        /// </summary>
        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        /// <summary>
        /// The username of the user. If not specified, it’s automatically generated.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public GroupServiceAccountState()
        {
        }
        public static new GroupServiceAccountState Empty => new GroupServiceAccountState();
    }
}
