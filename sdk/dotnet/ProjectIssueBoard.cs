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
    /// The `gitlab.ProjectIssueBoard` resource allows to manage the lifecycle of a Project Issue Board.
    /// 
    /// &gt; **NOTE:** If the board lists are changed all lists will be recreated.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/boards.html)
    /// 
    /// ## Import
    /// 
    /// You can import this resource with an id made up of `{project-id}:{issue-board-id}`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/projectIssueBoard:ProjectIssueBoard kanban 42:1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectIssueBoard:ProjectIssueBoard")]
    public partial class ProjectIssueBoard : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The assignee the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        [Output("assigneeId")]
        public Output<int?> AssigneeId { get; private set; } = null!;

        /// <summary>
        /// The list of label names which the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The list of issue board lists
        /// </summary>
        [Output("lists")]
        public Output<ImmutableArray<Outputs.ProjectIssueBoardList>> Lists { get; private set; } = null!;

        /// <summary>
        /// The milestone the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        [Output("milestoneId")]
        public Output<int?> MilestoneId { get; private set; } = null!;

        /// <summary>
        /// The name of the board.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID or full path of the project maintained by the authenticated user.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectIssueBoard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectIssueBoard(string name, ProjectIssueBoardArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectIssueBoard:ProjectIssueBoard", name, args ?? new ProjectIssueBoardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectIssueBoard(string name, Input<string> id, ProjectIssueBoardState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectIssueBoard:ProjectIssueBoard", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectIssueBoard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectIssueBoard Get(string name, Input<string> id, ProjectIssueBoardState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectIssueBoard(name, id, state, options);
        }
    }

    public sealed class ProjectIssueBoardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The assignee the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        [Input("assigneeId")]
        public Input<int>? AssigneeId { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// The list of label names which the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        [Input("lists")]
        private InputList<Inputs.ProjectIssueBoardListArgs>? _lists;

        /// <summary>
        /// The list of issue board lists
        /// </summary>
        public InputList<Inputs.ProjectIssueBoardListArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.ProjectIssueBoardListArgs>());
            set => _lists = value;
        }

        /// <summary>
        /// The milestone the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        [Input("milestoneId")]
        public Input<int>? MilestoneId { get; set; }

        /// <summary>
        /// The name of the board.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project maintained by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ProjectIssueBoardArgs()
        {
        }
        public static new ProjectIssueBoardArgs Empty => new ProjectIssueBoardArgs();
    }

    public sealed class ProjectIssueBoardState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The assignee the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        [Input("assigneeId")]
        public Input<int>? AssigneeId { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// The list of label names which the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        [Input("lists")]
        private InputList<Inputs.ProjectIssueBoardListGetArgs>? _lists;

        /// <summary>
        /// The list of issue board lists
        /// </summary>
        public InputList<Inputs.ProjectIssueBoardListGetArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.ProjectIssueBoardListGetArgs>());
            set => _lists = value;
        }

        /// <summary>
        /// The milestone the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        [Input("milestoneId")]
        public Input<int>? MilestoneId { get; set; }

        /// <summary>
        /// The name of the board.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project maintained by the authenticated user.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ProjectIssueBoardState()
        {
        }
        public static new ProjectIssueBoardState Empty => new ProjectIssueBoardState();
    }
}
