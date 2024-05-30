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
    /// The `gitlab.InstanceVariable` resource allows to manage the lifecycle of an instance-level CI/CD variable.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/instance_level_ci_variables.html)
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
    ///     var example = new GitLab.InstanceVariable("example", new()
    ///     {
    ///         Key = "instance_variable_key",
    ///         Value = "instance_variable_value",
    ///         Protected = false,
    ///         Masked = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitLab instance variables can be imported using an id made up of `variablename`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/instanceVariable:InstanceVariable example instance_variable_key
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/instanceVariable:InstanceVariable")]
    public partial class InstanceVariable : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the variable. Maximum of 255 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the variable.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        /// </summary>
        [Output("masked")]
        public Output<bool?> Masked { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        /// </summary>
        [Output("protected")]
        public Output<bool?> Protected { get; private set; } = null!;

        /// <summary>
        /// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        /// </summary>
        [Output("raw")]
        public Output<bool?> Raw { get; private set; } = null!;

        /// <summary>
        /// The value of the variable.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        /// </summary>
        [Output("variableType")]
        public Output<string?> VariableType { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceVariable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceVariable(string name, InstanceVariableArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/instanceVariable:InstanceVariable", name, args ?? new InstanceVariableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceVariable(string name, Input<string> id, InstanceVariableState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/instanceVariable:InstanceVariable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceVariable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceVariable Get(string name, Input<string> id, InstanceVariableState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceVariable(name, id, state, options);
        }
    }

    public sealed class InstanceVariableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the variable. Maximum of 255 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the variable.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        /// </summary>
        [Input("masked")]
        public Input<bool>? Masked { get; set; }

        /// <summary>
        /// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        /// <summary>
        /// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        /// </summary>
        [Input("raw")]
        public Input<bool>? Raw { get; set; }

        /// <summary>
        /// The value of the variable.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        /// </summary>
        [Input("variableType")]
        public Input<string>? VariableType { get; set; }

        public InstanceVariableArgs()
        {
        }
        public static new InstanceVariableArgs Empty => new InstanceVariableArgs();
    }

    public sealed class InstanceVariableState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the variable. Maximum of 255 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the variable.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        /// </summary>
        [Input("masked")]
        public Input<bool>? Masked { get; set; }

        /// <summary>
        /// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        /// <summary>
        /// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        /// </summary>
        [Input("raw")]
        public Input<bool>? Raw { get; set; }

        /// <summary>
        /// The value of the variable.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        /// </summary>
        [Input("variableType")]
        public Input<string>? VariableType { get; set; }

        public InstanceVariableState()
        {
        }
        public static new InstanceVariableState Empty => new InstanceVariableState();
    }
}
