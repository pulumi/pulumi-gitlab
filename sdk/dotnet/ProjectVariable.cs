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
    /// This resource allows you to create and manage CI/CD variables for your GitLab projects.
    /// For further information on variables, consult the [gitlab
    /// documentation](https://docs.gitlab.com/ce/ci/variables/README.html#variables).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new GitLab.ProjectVariable("example", new GitLab.ProjectVariableArgs
    ///         {
    ///             Key = "project_variable_key",
    ///             Project = "12345",
    ///             Protected = false,
    ///             Value = "project_variable_value",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ProjectVariable : Pulumi.CustomResource
    {
        /// <summary>
        /// The environment_scope of the variable
        /// </summary>
        [Output("environmentScope")]
        public Output<string?> EnvironmentScope { get; private set; } = null!;

        /// <summary>
        /// The name of the variable.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the variable will be masked if it would have been written to the logs. Defaults to `false`.
        /// </summary>
        [Output("masked")]
        public Output<bool?> Masked { get; private set; } = null!;

        /// <summary>
        /// The name or id of the project to add the hook to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        /// </summary>
        [Output("protected")]
        public Output<bool?> Protected { get; private set; } = null!;

        /// <summary>
        /// The value of the variable.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// The type of a variable. Available types are: env_var (default) and file.
        /// </summary>
        [Output("variableType")]
        public Output<string?> VariableType { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectVariable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectVariable(string name, ProjectVariableArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectVariable:ProjectVariable", name, args ?? new ProjectVariableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectVariable(string name, Input<string> id, ProjectVariableState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectVariable:ProjectVariable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectVariable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectVariable Get(string name, Input<string> id, ProjectVariableState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectVariable(name, id, state, options);
        }
    }

    public sealed class ProjectVariableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The environment_scope of the variable
        /// </summary>
        [Input("environmentScope")]
        public Input<string>? EnvironmentScope { get; set; }

        /// <summary>
        /// The name of the variable.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// If set to `true`, the variable will be masked if it would have been written to the logs. Defaults to `false`.
        /// </summary>
        [Input("masked")]
        public Input<bool>? Masked { get; set; }

        /// <summary>
        /// The name or id of the project to add the hook to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        /// <summary>
        /// The value of the variable.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// The type of a variable. Available types are: env_var (default) and file.
        /// </summary>
        [Input("variableType")]
        public Input<string>? VariableType { get; set; }

        public ProjectVariableArgs()
        {
        }
    }

    public sealed class ProjectVariableState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The environment_scope of the variable
        /// </summary>
        [Input("environmentScope")]
        public Input<string>? EnvironmentScope { get; set; }

        /// <summary>
        /// The name of the variable.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// If set to `true`, the variable will be masked if it would have been written to the logs. Defaults to `false`.
        /// </summary>
        [Input("masked")]
        public Input<bool>? Masked { get; set; }

        /// <summary>
        /// The name or id of the project to add the hook to.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        /// <summary>
        /// The value of the variable.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The type of a variable. Available types are: env_var (default) and file.
        /// </summary>
        [Input("variableType")]
        public Input<string>? VariableType { get; set; }

        public ProjectVariableState()
        {
        }
    }
}
