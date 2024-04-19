// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetRelease
    {
        /// <summary>
        /// The `gitlab.getRelease` data source retrieves information about a gitlab release for a project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/)
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // By project ID and tag_name
        ///     var example = GitLab.GetRelease.Invoke(new()
        ///     {
        ///         ProjectId = "1234",
        ///         TagName = "v1.0",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetReleaseResult> InvokeAsync(GetReleaseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReleaseResult>("gitlab:index/getRelease:getRelease", args ?? new GetReleaseArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getRelease` data source retrieves information about a gitlab release for a project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/)
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // By project ID and tag_name
        ///     var example = GitLab.GetRelease.Invoke(new()
        ///     {
        ///         ProjectId = "1234",
        ///         TagName = "v1.0",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetReleaseResult> Invoke(GetReleaseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReleaseResult>("gitlab:index/getRelease:getRelease", args ?? new GetReleaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReleaseArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The assets for a release
        /// </summary>
        [Input("assets")]
        public Inputs.GetReleaseAssetsArgs? Assets { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        /// <summary>
        /// The Git tag the release is associated with.
        /// </summary>
        [Input("tagName", required: true)]
        public string TagName { get; set; } = null!;

        public GetReleaseArgs()
        {
        }
        public static new GetReleaseArgs Empty => new GetReleaseArgs();
    }

    public sealed class GetReleaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The assets for a release
        /// </summary>
        [Input("assets")]
        public Input<Inputs.GetReleaseAssetsInputArgs>? Assets { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The Git tag the release is associated with.
        /// </summary>
        [Input("tagName", required: true)]
        public Input<string> TagName { get; set; } = null!;

        public GetReleaseInvokeArgs()
        {
        }
        public static new GetReleaseInvokeArgs Empty => new GetReleaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetReleaseResult
    {
        /// <summary>
        /// The assets for a release
        /// </summary>
        public readonly Outputs.GetReleaseAssetsResult? Assets;
        /// <summary>
        /// The date the release was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// An HTML rendered description of the release.
        /// </summary>
        public readonly string Description;
        public readonly string Id;
        /// <summary>
        /// The name of the release.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID or URL-encoded path of the project.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The date the release was created.
        /// </summary>
        public readonly string ReleasedAt;
        /// <summary>
        /// The Git tag the release is associated with.
        /// </summary>
        public readonly string TagName;

        [OutputConstructor]
        private GetReleaseResult(
            Outputs.GetReleaseAssetsResult? assets,

            string createdAt,

            string description,

            string id,

            string name,

            string projectId,

            string releasedAt,

            string tagName)
        {
            Assets = assets;
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Name = name;
            ProjectId = projectId;
            ReleasedAt = releasedAt;
            TagName = tagName;
        }
    }
}
