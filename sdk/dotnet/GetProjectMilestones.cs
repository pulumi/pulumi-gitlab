// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectMilestones
    {
        /// <summary>
        /// The `gitlab.getProjectMilestones` data source allows get details of a project milestones.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = GitLab.GetProjectMilestones.Invoke(new()
        ///     {
        ///         Project = "foo/bar",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectMilestonesResult> InvokeAsync(GetProjectMilestonesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectMilestonesResult>("gitlab:index/getProjectMilestones:getProjectMilestones", args ?? new GetProjectMilestonesArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectMilestones` data source allows get details of a project milestones.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = GitLab.GetProjectMilestones.Invoke(new()
        ///     {
        ///         Project = "foo/bar",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectMilestonesResult> Invoke(GetProjectMilestonesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectMilestonesResult>("gitlab:index/getProjectMilestones:getProjectMilestones", args ?? new GetProjectMilestonesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectMilestonesArgs : global::Pulumi.InvokeArgs
    {
        [Input("iids")]
        private List<int>? _iids;

        /// <summary>
        /// Return only the milestones having the given `iid` (Note: ignored if `include_parent_milestones` is set as `true`).
        /// </summary>
        public List<int> Iids
        {
            get => _iids ?? (_iids = new List<int>());
            set => _iids = value;
        }

        /// <summary>
        /// Include group milestones from parent group and its ancestors. Introduced in GitLab 13.4.
        /// </summary>
        [Input("includeParentMilestones")]
        public bool? IncludeParentMilestones { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        /// <summary>
        /// Return only milestones with a title or description matching the provided string.
        /// </summary>
        [Input("search")]
        public string? Search { get; set; }

        /// <summary>
        /// Return only `active` or `closed` milestones.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        /// <summary>
        /// Return only the milestones having the given `title`.
        /// </summary>
        [Input("title")]
        public string? Title { get; set; }

        public GetProjectMilestonesArgs()
        {
        }
        public static new GetProjectMilestonesArgs Empty => new GetProjectMilestonesArgs();
    }

    public sealed class GetProjectMilestonesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("iids")]
        private InputList<int>? _iids;

        /// <summary>
        /// Return only the milestones having the given `iid` (Note: ignored if `include_parent_milestones` is set as `true`).
        /// </summary>
        public InputList<int> Iids
        {
            get => _iids ?? (_iids = new InputList<int>());
            set => _iids = value;
        }

        /// <summary>
        /// Include group milestones from parent group and its ancestors. Introduced in GitLab 13.4.
        /// </summary>
        [Input("includeParentMilestones")]
        public Input<bool>? IncludeParentMilestones { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Return only milestones with a title or description matching the provided string.
        /// </summary>
        [Input("search")]
        public Input<string>? Search { get; set; }

        /// <summary>
        /// Return only `active` or `closed` milestones.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Return only the milestones having the given `title`.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public GetProjectMilestonesInvokeArgs()
        {
        }
        public static new GetProjectMilestonesInvokeArgs Empty => new GetProjectMilestonesInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectMilestonesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Return only the milestones having the given `iid` (Note: ignored if `include_parent_milestones` is set as `true`).
        /// </summary>
        public readonly ImmutableArray<int> Iids;
        /// <summary>
        /// Include group milestones from parent group and its ancestors. Introduced in GitLab 13.4.
        /// </summary>
        public readonly bool? IncludeParentMilestones;
        /// <summary>
        /// List of milestones from a project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectMilestonesMilestoneResult> Milestones;
        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Return only milestones with a title or description matching the provided string.
        /// </summary>
        public readonly string? Search;
        /// <summary>
        /// Return only `active` or `closed` milestones.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Return only the milestones having the given `title`.
        /// </summary>
        public readonly string? Title;

        [OutputConstructor]
        private GetProjectMilestonesResult(
            string id,

            ImmutableArray<int> iids,

            bool? includeParentMilestones,

            ImmutableArray<Outputs.GetProjectMilestonesMilestoneResult> milestones,

            string project,

            string? search,

            string? state,

            string? title)
        {
            Id = id;
            Iids = iids;
            IncludeParentMilestones = includeParentMilestones;
            Milestones = milestones;
            Project = project;
            Search = search;
            State = state;
            Title = title;
        }
    }
}
