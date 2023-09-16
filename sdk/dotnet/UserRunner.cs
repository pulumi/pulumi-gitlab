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
    /// The `gitlab.UserRunner` resource allows creating a GitLab runner using the new [GitLab Runner Registration Flow](https://docs.gitlab.com/ee/ci/runners/new_creation_workflow.html).
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#create-a-runner)
    /// </summary>
    [GitLabResourceType("gitlab:index/userRunner:UserRunner")]
    public partial class UserRunner : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access level of the runner. Valid values are: `not_protected`, `ref_protected`.
        /// </summary>
        [Output("accessLevel")]
        public Output<string> AccessLevel { get; private set; } = null!;

        /// <summary>
        /// Description of the runner.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the group that the runner is created in. Required if runner*type is group*type.
        /// </summary>
        [Output("groupId")]
        public Output<int?> GroupId { get; private set; } = null!;

        /// <summary>
        /// Specifies if the runner should be locked for the current project.
        /// </summary>
        [Output("locked")]
        public Output<bool> Locked { get; private set; } = null!;

        /// <summary>
        /// Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
        /// </summary>
        [Output("maximumTimeout")]
        public Output<int> MaximumTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies if the runner should ignore new jobs.
        /// </summary>
        [Output("paused")]
        public Output<bool> Paused { get; private set; } = null!;

        /// <summary>
        /// The ID of the project that the runner is created in. Required if runner*type is project*type.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The scope of the runner. Valid values are: `instance_type`, `group_type`, `project_type`.
        /// </summary>
        [Output("runnerType")]
        public Output<string> RunnerType { get; private set; } = null!;

        /// <summary>
        /// A list of runner tags.
        /// </summary>
        [Output("tagLists")]
        public Output<ImmutableArray<string>> TagLists { get; private set; } = null!;

        /// <summary>
        /// The authentication token to use when setting up a new runner with this configuration. This value cannot be imported.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// Specifies if the runner should handle untagged jobs.
        /// </summary>
        [Output("untagged")]
        public Output<bool> Untagged { get; private set; } = null!;


        /// <summary>
        /// Create a UserRunner resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserRunner(string name, UserRunnerArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/userRunner:UserRunner", name, args ?? new UserRunnerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserRunner(string name, Input<string> id, UserRunnerState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/userRunner:UserRunner", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserRunner resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserRunner Get(string name, Input<string> id, UserRunnerState? state = null, CustomResourceOptions? options = null)
        {
            return new UserRunner(name, id, state, options);
        }
    }

    public sealed class UserRunnerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level of the runner. Valid values are: `not_protected`, `ref_protected`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// Description of the runner.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the group that the runner is created in. Required if runner*type is group*type.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// Specifies if the runner should be locked for the current project.
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        /// <summary>
        /// Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
        /// </summary>
        [Input("maximumTimeout")]
        public Input<int>? MaximumTimeout { get; set; }

        /// <summary>
        /// Specifies if the runner should ignore new jobs.
        /// </summary>
        [Input("paused")]
        public Input<bool>? Paused { get; set; }

        /// <summary>
        /// The ID of the project that the runner is created in. Required if runner*type is project*type.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// The scope of the runner. Valid values are: `instance_type`, `group_type`, `project_type`.
        /// </summary>
        [Input("runnerType", required: true)]
        public Input<string> RunnerType { get; set; } = null!;

        [Input("tagLists")]
        private InputList<string>? _tagLists;

        /// <summary>
        /// A list of runner tags.
        /// </summary>
        public InputList<string> TagLists
        {
            get => _tagLists ?? (_tagLists = new InputList<string>());
            set => _tagLists = value;
        }

        /// <summary>
        /// Specifies if the runner should handle untagged jobs.
        /// </summary>
        [Input("untagged")]
        public Input<bool>? Untagged { get; set; }

        public UserRunnerArgs()
        {
        }
        public static new UserRunnerArgs Empty => new UserRunnerArgs();
    }

    public sealed class UserRunnerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level of the runner. Valid values are: `not_protected`, `ref_protected`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// Description of the runner.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the group that the runner is created in. Required if runner*type is group*type.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// Specifies if the runner should be locked for the current project.
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        /// <summary>
        /// Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
        /// </summary>
        [Input("maximumTimeout")]
        public Input<int>? MaximumTimeout { get; set; }

        /// <summary>
        /// Specifies if the runner should ignore new jobs.
        /// </summary>
        [Input("paused")]
        public Input<bool>? Paused { get; set; }

        /// <summary>
        /// The ID of the project that the runner is created in. Required if runner*type is project*type.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// The scope of the runner. Valid values are: `instance_type`, `group_type`, `project_type`.
        /// </summary>
        [Input("runnerType")]
        public Input<string>? RunnerType { get; set; }

        [Input("tagLists")]
        private InputList<string>? _tagLists;

        /// <summary>
        /// A list of runner tags.
        /// </summary>
        public InputList<string> TagLists
        {
            get => _tagLists ?? (_tagLists = new InputList<string>());
            set => _tagLists = value;
        }

        /// <summary>
        /// The authentication token to use when setting up a new runner with this configuration. This value cannot be imported.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// Specifies if the runner should handle untagged jobs.
        /// </summary>
        [Input("untagged")]
        public Input<bool>? Untagged { get; set; }

        public UserRunnerState()
        {
        }
        public static new UserRunnerState Empty => new UserRunnerState();
    }
}