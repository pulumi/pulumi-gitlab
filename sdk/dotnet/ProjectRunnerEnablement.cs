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
    /// The `gitlab.ProjectRunnerEnablement` resource allows to enable a runner in a project.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/runners/#enable-a-runner-in-project)
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
    ///     var foo = new GitLab.ProjectRunnerEnablement("foo", new()
    ///     {
    ///         Project = "5",
    ///         RunnerId = 7,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_runner_enablement`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_project_runner_enablement.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// GitLab project runners can be imported using an id made up of `project:runner_id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement foo 5:7
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement")]
    public partial class ProjectRunnerEnablement : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The ID of a runner to enable for the project.
        /// </summary>
        [Output("runnerId")]
        public Output<int> RunnerId { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectRunnerEnablement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectRunnerEnablement(string name, ProjectRunnerEnablementArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement", name, args ?? new ProjectRunnerEnablementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectRunnerEnablement(string name, Input<string> id, ProjectRunnerEnablementState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectRunnerEnablement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectRunnerEnablement Get(string name, Input<string> id, ProjectRunnerEnablementState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectRunnerEnablement(name, id, state, options);
        }
    }

    public sealed class ProjectRunnerEnablementArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The ID of a runner to enable for the project.
        /// </summary>
        [Input("runnerId", required: true)]
        public Input<int> RunnerId { get; set; } = null!;

        public ProjectRunnerEnablementArgs()
        {
        }
        public static new ProjectRunnerEnablementArgs Empty => new ProjectRunnerEnablementArgs();
    }

    public sealed class ProjectRunnerEnablementState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The ID of a runner to enable for the project.
        /// </summary>
        [Input("runnerId")]
        public Input<int>? RunnerId { get; set; }

        public ProjectRunnerEnablementState()
        {
        }
        public static new ProjectRunnerEnablementState Empty => new ProjectRunnerEnablementState();
    }
}
