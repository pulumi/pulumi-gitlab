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
    /// The `gitlab.ProjectJobTokenScope` resource allows to manage the CI/CD Job Token scope in a project.
    /// Any projects added to the CI/CD Job Token scope outside of TF will be untouched by the resource.
    /// 
    /// &gt; Conflicts with the use of `gitlab.ProjectJobTokenScopes` when used on the same project. Use one or the other to ensure the desired state.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_job_token_scopes/)
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_job_token_scope`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_project_job_token_scope.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// GitLab project job token scopes can be imported using an id made up of `projectId:targetProjectId`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectJobTokenScope:ProjectJobTokenScope bar 123:321
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectJobTokenScope:ProjectJobTokenScope")]
    public partial class ProjectJobTokenScope : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The ID of the group that is in the CI/CD job token inbound allowlist.
        /// </summary>
        [Output("targetGroupId")]
        public Output<int?> TargetGroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the project that is in the CI/CD job token inbound allowlist.
        /// </summary>
        [Output("targetProjectId")]
        public Output<int?> TargetProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectJobTokenScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectJobTokenScope(string name, ProjectJobTokenScopeArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectJobTokenScope:ProjectJobTokenScope", name, args ?? new ProjectJobTokenScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectJobTokenScope(string name, Input<string> id, ProjectJobTokenScopeState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectJobTokenScope:ProjectJobTokenScope", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectJobTokenScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectJobTokenScope Get(string name, Input<string> id, ProjectJobTokenScopeState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectJobTokenScope(name, id, state, options);
        }
    }

    public sealed class ProjectJobTokenScopeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The ID of the group that is in the CI/CD job token inbound allowlist.
        /// </summary>
        [Input("targetGroupId")]
        public Input<int>? TargetGroupId { get; set; }

        /// <summary>
        /// The ID of the project that is in the CI/CD job token inbound allowlist.
        /// </summary>
        [Input("targetProjectId")]
        public Input<int>? TargetProjectId { get; set; }

        public ProjectJobTokenScopeArgs()
        {
        }
        public static new ProjectJobTokenScopeArgs Empty => new ProjectJobTokenScopeArgs();
    }

    public sealed class ProjectJobTokenScopeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The ID of the group that is in the CI/CD job token inbound allowlist.
        /// </summary>
        [Input("targetGroupId")]
        public Input<int>? TargetGroupId { get; set; }

        /// <summary>
        /// The ID of the project that is in the CI/CD job token inbound allowlist.
        /// </summary>
        [Input("targetProjectId")]
        public Input<int>? TargetProjectId { get; set; }

        public ProjectJobTokenScopeState()
        {
        }
        public static new ProjectJobTokenScopeState Empty => new ProjectJobTokenScopeState();
    }
}
