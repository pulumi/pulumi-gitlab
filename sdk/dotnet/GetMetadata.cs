// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetMetadata
    {
        /// <summary>
        /// The `gitlab.getMetadata` data source retrieves the metadata of the GitLab instance.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/metadata.html)
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
        ///     var @this = GitLab.GetMetadata.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetMetadataResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMetadataResult>("gitlab:index/getMetadata:getMetadata", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// The `gitlab.getMetadata` data source retrieves the metadata of the GitLab instance.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/metadata.html)
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
        ///     var @this = GitLab.GetMetadata.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetMetadataResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMetadataResult>("gitlab:index/getMetadata:getMetadata", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetMetadataResult
    {
        /// <summary>
        /// If the GitLab instance is an enterprise instance or not. Supported for GitLab 15.6 onwards.
        /// </summary>
        public readonly bool Enterprise;
        /// <summary>
        /// The id of the data source. It will always be `1`
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Metadata about the GitLab agent server for Kubernetes (KAS).
        /// </summary>
        public readonly Outputs.GetMetadataKasResult Kas;
        /// <summary>
        /// Revision of the GitLab instance.
        /// </summary>
        public readonly string Revision;
        /// <summary>
        /// Version of the GitLab instance.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetMetadataResult(
            bool enterprise,

            string id,

            Outputs.GetMetadataKasResult kas,

            string revision,

            string version)
        {
            Enterprise = enterprise;
            Id = id;
            Kas = kas;
            Revision = revision;
            Version = version;
        }
    }
}
