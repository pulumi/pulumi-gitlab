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
    /// The `gitlab.ProjectSecurityPolicyAttachment` resource allows to attach a security policy project to a project.
    /// 
    /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/index.html#mutationsecuritypolicyprojectassign)
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
    ///     // This resource can be used to attach a security policy to a pre-existing project
    ///     var foo = new GitLab.ProjectSecurityPolicyAttachment("foo", new()
    ///     {
    ///         Project = "1234",
    ///         PolicyProject = "4567",
    ///     });
    /// 
    ///     // Or you can use Terraform to create a new project, add a policy to that project,
    ///     // then attach that policy project to other projects.
    ///     var my_policy_project = new GitLab.Project("my-policy-project", new()
    ///     {
    ///         Name = "security-policy-project",
    ///     });
    /// 
    ///     var policy_yml = new GitLab.RepositoryFile("policy-yml", new()
    ///     {
    ///         Project = my_policy_project.Id,
    ///         FilePath = ".gitlab/security-policies/my-policy.yml",
    ///         Branch = "master",
    ///         Encoding = "text",
    ///         Content = @"---
    /// approval_policy:
    /// - name: test
    /// description: test
    /// enabled: true
    /// rules:
    /// - type: any_merge_request
    ///     branch_type: protected
    ///     commits: any
    /// approval_settings:
    ///     block_branch_modification: true
    ///     prevent_pushing_and_force_pushing: true
    ///     prevent_approval_by_author: true
    ///     prevent_approval_by_commit_author: true
    ///     remove_approvals_with_new_commit: true
    ///     require_password_to_approve: false
    /// fallback_behavior:
    ///     fail: closed
    /// actions:
    /// - type: send_bot_message
    ///     enabled: true
    /// ",
    ///     });
    /// 
    ///     var my_policy = new GitLab.Index.ProjectSecurityPolicy("my-policy", new()
    ///     {
    ///         Project = 1234,
    ///         PolicyProject = my_policy_project.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitLab project security policy attachments can be imported using an id made up of `project:policy_project_id` where the policy project ID is the project ID of the policy project, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectSecurityPolicyAttachment:ProjectSecurityPolicyAttachment foo 1:2
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectSecurityPolicyAttachment:ProjectSecurityPolicyAttachment")]
    public partial class ProjectSecurityPolicyAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID or Full Path of the security policy project.
        /// </summary>
        [Output("policyProject")]
        public Output<string> PolicyProject { get; private set; } = null!;

        /// <summary>
        /// The GraphQL ID of the security policy project.
        /// </summary>
        [Output("policyProjectGraphqlId")]
        public Output<string> PolicyProjectGraphqlId { get; private set; } = null!;

        /// <summary>
        /// The ID or Full Path of the project which will have the security policy project assigned to it.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The GraphQL ID of the project to which the security policty project will be attached.
        /// </summary>
        [Output("projectGraphqlId")]
        public Output<string> ProjectGraphqlId { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectSecurityPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectSecurityPolicyAttachment(string name, ProjectSecurityPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectSecurityPolicyAttachment:ProjectSecurityPolicyAttachment", name, args ?? new ProjectSecurityPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectSecurityPolicyAttachment(string name, Input<string> id, ProjectSecurityPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectSecurityPolicyAttachment:ProjectSecurityPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectSecurityPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectSecurityPolicyAttachment Get(string name, Input<string> id, ProjectSecurityPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectSecurityPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class ProjectSecurityPolicyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or Full Path of the security policy project.
        /// </summary>
        [Input("policyProject", required: true)]
        public Input<string> PolicyProject { get; set; } = null!;

        /// <summary>
        /// The ID or Full Path of the project which will have the security policy project assigned to it.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public ProjectSecurityPolicyAttachmentArgs()
        {
        }
        public static new ProjectSecurityPolicyAttachmentArgs Empty => new ProjectSecurityPolicyAttachmentArgs();
    }

    public sealed class ProjectSecurityPolicyAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or Full Path of the security policy project.
        /// </summary>
        [Input("policyProject")]
        public Input<string>? PolicyProject { get; set; }

        /// <summary>
        /// The GraphQL ID of the security policy project.
        /// </summary>
        [Input("policyProjectGraphqlId")]
        public Input<string>? PolicyProjectGraphqlId { get; set; }

        /// <summary>
        /// The ID or Full Path of the project which will have the security policy project assigned to it.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The GraphQL ID of the project to which the security policty project will be attached.
        /// </summary>
        [Input("projectGraphqlId")]
        public Input<string>? ProjectGraphqlId { get; set; }

        public ProjectSecurityPolicyAttachmentState()
        {
        }
        public static new ProjectSecurityPolicyAttachmentState Empty => new ProjectSecurityPolicyAttachmentState();
    }
}