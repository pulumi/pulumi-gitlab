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
    /// The `gitlab.GroupEpicBoard` resource allows to manage the lifecycle of a epic board in a group.
    /// 
    /// &gt; Multiple epic boards on one group requires a GitLab Premium or above License.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_boards.html)
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
    ///     var example = new GitLab.Group("example", new()
    ///     {
    ///         Name = "test_group",
    ///         Path = "test_group",
    ///         Description = "An example group",
    ///     });
    /// 
    ///     var label1 = new GitLab.GroupLabel("label_1", new()
    ///     {
    ///         Group = example.Id,
    ///         Color = "#FF0000",
    ///         Name = "red-label",
    ///     });
    /// 
    ///     var label3 = new GitLab.GroupLabel("label_3", new()
    ///     {
    ///         Group = example.Id,
    ///         Name = "label-3",
    ///         Color = "#003000",
    ///     });
    /// 
    ///     var epicBoard = new GitLab.GroupEpicBoard("epic_board", new()
    ///     {
    ///         Name = "epic board 6",
    ///         Group = example.Path,
    ///         Lists = new[]
    ///         {
    ///             new GitLab.Inputs.GroupEpicBoardListArgs
    ///             {
    ///                 LabelId = label1.LabelId,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// You can import this resource with an id made up of `{group-id}:{epic-board-id}`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/groupEpicBoard:GroupEpicBoard agile 70:156
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/groupEpicBoard:GroupEpicBoard")]
    public partial class GroupEpicBoard : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID or URL-encoded path of the group owned by the authenticated user.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The list of epic board lists.
        /// </summary>
        [Output("lists")]
        public Output<ImmutableArray<Outputs.GroupEpicBoardList>> Lists { get; private set; } = null!;

        /// <summary>
        /// The name of the board.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a GroupEpicBoard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupEpicBoard(string name, GroupEpicBoardArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupEpicBoard:GroupEpicBoard", name, args ?? new GroupEpicBoardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupEpicBoard(string name, Input<string> id, GroupEpicBoardState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupEpicBoard:GroupEpicBoard", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupEpicBoard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupEpicBoard Get(string name, Input<string> id, GroupEpicBoardState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupEpicBoard(name, id, state, options);
        }
    }

    public sealed class GroupEpicBoardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or URL-encoded path of the group owned by the authenticated user.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        [Input("lists")]
        private InputList<Inputs.GroupEpicBoardListArgs>? _lists;

        /// <summary>
        /// The list of epic board lists.
        /// </summary>
        public InputList<Inputs.GroupEpicBoardListArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.GroupEpicBoardListArgs>());
            set => _lists = value;
        }

        /// <summary>
        /// The name of the board.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupEpicBoardArgs()
        {
        }
        public static new GroupEpicBoardArgs Empty => new GroupEpicBoardArgs();
    }

    public sealed class GroupEpicBoardState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or URL-encoded path of the group owned by the authenticated user.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("lists")]
        private InputList<Inputs.GroupEpicBoardListGetArgs>? _lists;

        /// <summary>
        /// The list of epic board lists.
        /// </summary>
        public InputList<Inputs.GroupEpicBoardListGetArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.GroupEpicBoardListGetArgs>());
            set => _lists = value;
        }

        /// <summary>
        /// The name of the board.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupEpicBoardState()
        {
        }
        public static new GroupEpicBoardState Empty => new GroupEpicBoardState();
    }
}
