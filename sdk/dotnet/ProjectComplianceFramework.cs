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
    /// The `gitlab.ProjectComplianceFramework` resource allows to manage the lifecycle of a compliance framework on a project.
    /// 
    /// &gt; This resource requires a GitLab Enterprise instance with a Premium license to set the compliance framework on a project.
    /// 
    /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#mutationprojectsetcomplianceframework)
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
    ///     var sampleComplianceFramework = new GitLab.ComplianceFramework("sampleComplianceFramework", new()
    ///     {
    ///         NamespacePath = "top-level-group",
    ///         Description = "A HIPAA Compliance Framework",
    ///         Color = "#87BEEF",
    ///         Default = false,
    ///         PipelineConfigurationFullPath = ".hipaa.yml@top-level-group/compliance-frameworks",
    ///     });
    /// 
    ///     var sampleProjectComplianceFramework = new GitLab.ProjectComplianceFramework("sampleProjectComplianceFramework", new()
    ///     {
    ///         ComplianceFrameworkId = sampleComplianceFramework.FrameworkId,
    ///         Project = "12345678",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Gitlab project compliance frameworks can be imported with a key composed of `&lt;project_id&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/projectComplianceFramework:ProjectComplianceFramework sample "42"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectComplianceFramework:ProjectComplianceFramework")]
    public partial class ProjectComplianceFramework : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Globally unique ID of the compliance framework to assign to the project.
        /// </summary>
        [Output("complianceFrameworkId")]
        public Output<string> ComplianceFrameworkId { get; private set; } = null!;

        /// <summary>
        /// The ID or full path of the project to change the compliance framework of.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectComplianceFramework resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectComplianceFramework(string name, ProjectComplianceFrameworkArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectComplianceFramework:ProjectComplianceFramework", name, args ?? new ProjectComplianceFrameworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectComplianceFramework(string name, Input<string> id, ProjectComplianceFrameworkState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectComplianceFramework:ProjectComplianceFramework", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectComplianceFramework resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectComplianceFramework Get(string name, Input<string> id, ProjectComplianceFrameworkState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectComplianceFramework(name, id, state, options);
        }
    }

    public sealed class ProjectComplianceFrameworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Globally unique ID of the compliance framework to assign to the project.
        /// </summary>
        [Input("complianceFrameworkId", required: true)]
        public Input<string> ComplianceFrameworkId { get; set; } = null!;

        /// <summary>
        /// The ID or full path of the project to change the compliance framework of.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public ProjectComplianceFrameworkArgs()
        {
        }
        public static new ProjectComplianceFrameworkArgs Empty => new ProjectComplianceFrameworkArgs();
    }

    public sealed class ProjectComplianceFrameworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Globally unique ID of the compliance framework to assign to the project.
        /// </summary>
        [Input("complianceFrameworkId")]
        public Input<string>? ComplianceFrameworkId { get; set; }

        /// <summary>
        /// The ID or full path of the project to change the compliance framework of.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ProjectComplianceFrameworkState()
        {
        }
        public static new ProjectComplianceFrameworkState Empty => new ProjectComplianceFrameworkState();
    }
}