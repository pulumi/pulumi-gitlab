// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetReleaseLinks
    {
        /// <summary>
        /// The `gitlab.getReleaseLinks` data source allows get details of release links.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)
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
        ///     var example = GitLab.GetReleaseLinks.Invoke(new()
        ///     {
        ///         Project = "foo/bar",
        ///         TagName = "v1.0.1",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetReleaseLinksResult> InvokeAsync(GetReleaseLinksArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReleaseLinksResult>("gitlab:index/getReleaseLinks:getReleaseLinks", args ?? new GetReleaseLinksArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getReleaseLinks` data source allows get details of release links.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)
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
        ///     var example = GitLab.GetReleaseLinks.Invoke(new()
        ///     {
        ///         Project = "foo/bar",
        ///         TagName = "v1.0.1",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetReleaseLinksResult> Invoke(GetReleaseLinksInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReleaseLinksResult>("gitlab:index/getReleaseLinks:getReleaseLinks", args ?? new GetReleaseLinksInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReleaseLinksArgs : global::Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("tagName", required: true)]
        public string TagName { get; set; } = null!;

        public GetReleaseLinksArgs()
        {
        }
        public static new GetReleaseLinksArgs Empty => new GetReleaseLinksArgs();
    }

    public sealed class GetReleaseLinksInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("tagName", required: true)]
        public Input<string> TagName { get; set; } = null!;

        public GetReleaseLinksInvokeArgs()
        {
        }
        public static new GetReleaseLinksInvokeArgs Empty => new GetReleaseLinksInvokeArgs();
    }


    [OutputType]
    public sealed class GetReleaseLinksResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID or full path to the project.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// List of release links
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReleaseLinksReleaseLinkResult> ReleaseLinks;
        /// <summary>
        /// The tag associated with the Release.
        /// </summary>
        public readonly string TagName;

        [OutputConstructor]
        private GetReleaseLinksResult(
            string id,

            string project,

            ImmutableArray<Outputs.GetReleaseLinksReleaseLinkResult> releaseLinks,

            string tagName)
        {
            Id = id;
            Project = project;
            ReleaseLinks = releaseLinks;
            TagName = tagName;
        }
    }
}
