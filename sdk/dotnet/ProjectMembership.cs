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
    /// The `gitlab.ProjectMembership` resource allows to manage the lifecycle of a users project membership.
    /// 
    /// &gt; If a project should grant membership to an entire group use the `gitlab.ProjectShareGroup` resource instead.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/members/)
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
    ///     var test = new GitLab.ProjectMembership("test", new()
    ///     {
    ///         Project = "12345",
    ///         UserId = 1337,
    ///         AccessLevel = "guest",
    ///     });
    /// 
    ///     var example = new GitLab.ProjectMembership("example", new()
    ///     {
    ///         Project = "67890",
    ///         UserId = 1234,
    ///         AccessLevel = "guest",
    ///         ExpiresAt = "2022-12-31",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_membership`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_project_membership.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// GitLab project membership can be imported using an id made up of `project_id:user_id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectMembership:ProjectMembership test "12345:1337"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectMembership:ProjectMembership")]
    public partial class ProjectMembership : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
        /// </summary>
        [Output("accessLevel")]
        public Output<string> AccessLevel { get; private set; } = null!;

        /// <summary>
        /// Expiration date for the project membership. Format: `YYYY-MM-DD`
        /// </summary>
        [Output("expiresAt")]
        public Output<string?> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The ID of a custom member role. Only available for Ultimate instances.
        /// </summary>
        [Output("memberRoleId")]
        public Output<int?> MemberRoleId { get; private set; } = null!;

        /// <summary>
        /// The ID or URL-encoded path of the project.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The id of the user.
        /// </summary>
        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectMembership(string name, ProjectMembershipArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectMembership:ProjectMembership", name, args ?? new ProjectMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectMembership(string name, Input<string> id, ProjectMembershipState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectMembership:ProjectMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectMembership Get(string name, Input<string> id, ProjectMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectMembership(name, id, state, options);
        }
    }

    public sealed class ProjectMembershipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
        /// </summary>
        [Input("accessLevel", required: true)]
        public Input<string> AccessLevel { get; set; } = null!;

        /// <summary>
        /// Expiration date for the project membership. Format: `YYYY-MM-DD`
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ID of a custom member role. Only available for Ultimate instances.
        /// </summary>
        [Input("memberRoleId")]
        public Input<int>? MemberRoleId { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The id of the user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public ProjectMembershipArgs()
        {
        }
        public static new ProjectMembershipArgs Empty => new ProjectMembershipArgs();
    }

    public sealed class ProjectMembershipState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// Expiration date for the project membership. Format: `YYYY-MM-DD`
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ID of a custom member role. Only available for Ultimate instances.
        /// </summary>
        [Input("memberRoleId")]
        public Input<int>? MemberRoleId { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The id of the user.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public ProjectMembershipState()
        {
        }
        public static new ProjectMembershipState Empty => new ProjectMembershipState();
    }
}
