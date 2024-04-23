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
    /// The `gitlab.ComplianceFramework` resource allows to manage the lifecycle of a compliance framework on top-level groups.
    /// 
    /// There can be only one `default` compliance framework. Of all the configured compliance frameworks marked as default, the last one applied will be the default compliance framework.
    /// 
    /// &gt; This resource requires a GitLab Enterprise instance with a Premium license to create the compliance framework.
    /// 
    /// &gt; This resource requires a GitLab Enterprise instance with an Ultimate license to specify a compliance pipeline configuration in the compliance framework.
    /// 
    /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#mutationcreatecomplianceframework)
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
    ///     var sample = new GitLab.ComplianceFramework("sample", new()
    ///     {
    ///         NamespacePath = "top-level-group",
    ///         Name = "HIPAA",
    ///         Description = "A HIPAA Compliance Framework",
    ///         Color = "#87BEEF",
    ///         Default = false,
    ///         PipelineConfigurationFullPath = ".hipaa.yml@top-level-group/compliance-frameworks",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Gitlab compliance frameworks can be imported with a key composed of `&lt;namespace_path&gt;:&lt;framework_id&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/complianceFramework:ComplianceFramework sample "top-level-group:gid://gitlab/ComplianceManagement::Framework/12345"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/complianceFramework:ComplianceFramework")]
    public partial class ComplianceFramework : global::Pulumi.CustomResource
    {
        /// <summary>
        /// New color representation of the compliance framework in hex format. e.g. #FCA121.
        /// </summary>
        [Output("color")]
        public Output<string> Color { get; private set; } = null!;

        /// <summary>
        /// Set this compliance framework as the default framework for the group. Default: `false`
        /// </summary>
        [Output("default")]
        public Output<bool> Default { get; private set; } = null!;

        /// <summary>
        /// Description for the compliance framework.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Globally unique ID of the compliance framework.
        /// </summary>
        [Output("frameworkId")]
        public Output<string> FrameworkId { get; private set; } = null!;

        /// <summary>
        /// Name for the compliance framework.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Full path of the namespace to add the compliance framework to.
        /// </summary>
        [Output("namespacePath")]
        public Output<string> NamespacePath { get; private set; } = null!;

        /// <summary>
        /// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
        /// </summary>
        [Output("pipelineConfigurationFullPath")]
        public Output<string?> PipelineConfigurationFullPath { get; private set; } = null!;


        /// <summary>
        /// Create a ComplianceFramework resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComplianceFramework(string name, ComplianceFrameworkArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/complianceFramework:ComplianceFramework", name, args ?? new ComplianceFrameworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ComplianceFramework(string name, Input<string> id, ComplianceFrameworkState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/complianceFramework:ComplianceFramework", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ComplianceFramework resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComplianceFramework Get(string name, Input<string> id, ComplianceFrameworkState? state = null, CustomResourceOptions? options = null)
        {
            return new ComplianceFramework(name, id, state, options);
        }
    }

    public sealed class ComplianceFrameworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// New color representation of the compliance framework in hex format. e.g. #FCA121.
        /// </summary>
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        /// <summary>
        /// Set this compliance framework as the default framework for the group. Default: `false`
        /// </summary>
        [Input("default")]
        public Input<bool>? Default { get; set; }

        /// <summary>
        /// Description for the compliance framework.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Name for the compliance framework.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Full path of the namespace to add the compliance framework to.
        /// </summary>
        [Input("namespacePath", required: true)]
        public Input<string> NamespacePath { get; set; } = null!;

        /// <summary>
        /// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
        /// </summary>
        [Input("pipelineConfigurationFullPath")]
        public Input<string>? PipelineConfigurationFullPath { get; set; }

        public ComplianceFrameworkArgs()
        {
        }
        public static new ComplianceFrameworkArgs Empty => new ComplianceFrameworkArgs();
    }

    public sealed class ComplianceFrameworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// New color representation of the compliance framework in hex format. e.g. #FCA121.
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// Set this compliance framework as the default framework for the group. Default: `false`
        /// </summary>
        [Input("default")]
        public Input<bool>? Default { get; set; }

        /// <summary>
        /// Description for the compliance framework.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Globally unique ID of the compliance framework.
        /// </summary>
        [Input("frameworkId")]
        public Input<string>? FrameworkId { get; set; }

        /// <summary>
        /// Name for the compliance framework.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Full path of the namespace to add the compliance framework to.
        /// </summary>
        [Input("namespacePath")]
        public Input<string>? NamespacePath { get; set; }

        /// <summary>
        /// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
        /// </summary>
        [Input("pipelineConfigurationFullPath")]
        public Input<string>? PipelineConfigurationFullPath { get; set; }

        public ComplianceFrameworkState()
        {
        }
        public static new ComplianceFrameworkState Empty => new ComplianceFrameworkState();
    }
}
