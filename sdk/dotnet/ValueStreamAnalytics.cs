// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    /// <summary>
    /// The `gitlab.ValueStreamAnalytics` resource allows to manage the lifecycle of value stream analytics.
    /// 
    /// &gt; This resource requires a GitLab Enterprise instance with a Premium license to create custom value stream analytics.
    /// 
    /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/api/graphql/reference/#mutationvaluestreamcreate)
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_value_stream_analytics`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_value_stream_analytics.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Importing using the CLI is supported with the following syntax:
    /// 
    /// Gitlab value stream analytics can be imported with a key composed of `&lt;full_path_type&gt;:&lt;full_path&gt;:&lt;value_stream_id&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/valueStreamAnalytics:ValueStreamAnalytics group "group:people/engineers:42"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/valueStreamAnalytics:ValueStreamAnalytics project "project:projects/sample:43"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/valueStreamAnalytics:ValueStreamAnalytics")]
    public partial class ValueStreamAnalytics : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Full path of the group the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**
        /// </summary>
        [Output("groupFullPath")]
        public Output<string?> GroupFullPath { get; private set; } = null!;

        /// <summary>
        /// The name of the value stream
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Full path of the project the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**
        /// </summary>
        [Output("projectFullPath")]
        public Output<string?> ProjectFullPath { get; private set; } = null!;

        /// <summary>
        /// Stages of the value stream
        /// </summary>
        [Output("stages")]
        public Output<ImmutableArray<Outputs.ValueStreamAnalyticsStage>> Stages { get; private set; } = null!;


        /// <summary>
        /// Create a ValueStreamAnalytics resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ValueStreamAnalytics(string name, ValueStreamAnalyticsArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/valueStreamAnalytics:ValueStreamAnalytics", name, args ?? new ValueStreamAnalyticsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ValueStreamAnalytics(string name, Input<string> id, ValueStreamAnalyticsState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/valueStreamAnalytics:ValueStreamAnalytics", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ValueStreamAnalytics resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ValueStreamAnalytics Get(string name, Input<string> id, ValueStreamAnalyticsState? state = null, CustomResourceOptions? options = null)
        {
            return new ValueStreamAnalytics(name, id, state, options);
        }
    }

    public sealed class ValueStreamAnalyticsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Full path of the group the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**
        /// </summary>
        [Input("groupFullPath")]
        public Input<string>? GroupFullPath { get; set; }

        /// <summary>
        /// The name of the value stream
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Full path of the project the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**
        /// </summary>
        [Input("projectFullPath")]
        public Input<string>? ProjectFullPath { get; set; }

        [Input("stages", required: true)]
        private InputList<Inputs.ValueStreamAnalyticsStageArgs>? _stages;

        /// <summary>
        /// Stages of the value stream
        /// </summary>
        public InputList<Inputs.ValueStreamAnalyticsStageArgs> Stages
        {
            get => _stages ?? (_stages = new InputList<Inputs.ValueStreamAnalyticsStageArgs>());
            set => _stages = value;
        }

        public ValueStreamAnalyticsArgs()
        {
        }
        public static new ValueStreamAnalyticsArgs Empty => new ValueStreamAnalyticsArgs();
    }

    public sealed class ValueStreamAnalyticsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Full path of the group the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**
        /// </summary>
        [Input("groupFullPath")]
        public Input<string>? GroupFullPath { get; set; }

        /// <summary>
        /// The name of the value stream
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Full path of the project the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**
        /// </summary>
        [Input("projectFullPath")]
        public Input<string>? ProjectFullPath { get; set; }

        [Input("stages")]
        private InputList<Inputs.ValueStreamAnalyticsStageGetArgs>? _stages;

        /// <summary>
        /// Stages of the value stream
        /// </summary>
        public InputList<Inputs.ValueStreamAnalyticsStageGetArgs> Stages
        {
            get => _stages ?? (_stages = new InputList<Inputs.ValueStreamAnalyticsStageGetArgs>());
            set => _stages = value;
        }

        public ValueStreamAnalyticsState()
        {
        }
        public static new ValueStreamAnalyticsState Empty => new ValueStreamAnalyticsState();
    }
}
