// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetReleaseLink
    {
        /// <summary>
        /// The `gitlab.ReleaseLink` data source allows get details of a release link.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)
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
        ///     var example = GitLab.GetReleaseLink.Invoke(new()
        ///     {
        ///         LinkId = 11,
        ///         Project = "foo/bar",
        ///         TagName = "v1.0.1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetReleaseLinkResult> InvokeAsync(GetReleaseLinkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReleaseLinkResult>("gitlab:index/getReleaseLink:getReleaseLink", args ?? new GetReleaseLinkArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.ReleaseLink` data source allows get details of a release link.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)
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
        ///     var example = GitLab.GetReleaseLink.Invoke(new()
        ///     {
        ///         LinkId = 11,
        ///         Project = "foo/bar",
        ///         TagName = "v1.0.1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetReleaseLinkResult> Invoke(GetReleaseLinkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReleaseLinkResult>("gitlab:index/getReleaseLink:getReleaseLink", args ?? new GetReleaseLinkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReleaseLinkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the link.
        /// </summary>
        [Input("linkId", required: true)]
        public int LinkId { get; set; }

        /// <summary>
        /// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        /// <summary>
        /// The tag associated with the Release.
        /// </summary>
        [Input("tagName", required: true)]
        public string TagName { get; set; } = null!;

        public GetReleaseLinkArgs()
        {
        }
        public static new GetReleaseLinkArgs Empty => new GetReleaseLinkArgs();
    }

    public sealed class GetReleaseLinkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the link.
        /// </summary>
        [Input("linkId", required: true)]
        public Input<int> LinkId { get; set; } = null!;

        /// <summary>
        /// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The tag associated with the Release.
        /// </summary>
        [Input("tagName", required: true)]
        public Input<string> TagName { get; set; } = null!;

        public GetReleaseLinkInvokeArgs()
        {
        }
        public static new GetReleaseLinkInvokeArgs Empty => new GetReleaseLinkInvokeArgs();
    }


    [OutputType]
    public sealed class GetReleaseLinkResult
    {
        /// <summary>
        /// Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        /// </summary>
        public readonly string DirectAssetUrl;
        /// <summary>
        /// External or internal link.
        /// </summary>
        public readonly bool External;
        /// <summary>
        /// Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        /// </summary>
        public readonly string Filepath;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the link.
        /// </summary>
        public readonly int LinkId;
        /// <summary>
        /// The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
        /// </summary>
        public readonly string LinkType;
        /// <summary>
        /// The name of the link. Link names must be unique within the release.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The tag associated with the Release.
        /// </summary>
        public readonly string TagName;
        /// <summary>
        /// The URL of the link. Link URLs must be unique within the release.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetReleaseLinkResult(
            string directAssetUrl,

            bool external,

            string filepath,

            string id,

            int linkId,

            string linkType,

            string name,

            string project,

            string tagName,

            string url)
        {
            DirectAssetUrl = directAssetUrl;
            External = external;
            Filepath = filepath;
            Id = id;
            LinkId = linkId;
            LinkType = linkType;
            Name = name;
            Project = project;
            TagName = tagName;
            Url = url;
        }
    }
}
