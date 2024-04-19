// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetInstanceDeployKeys
    {
        /// <summary>
        /// The `gitlab.getInstanceDeployKeys` data source allows to retrieve a list of deploy keys for a GitLab instance.
        /// 
        /// &gt; This data source requires administration privileges.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_keys.html#list-all-deploy-keys)
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
        ///     var example = GitLab.GetInstanceDeployKeys.Invoke(new()
        ///     {
        ///         Public = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetInstanceDeployKeysResult> InvokeAsync(GetInstanceDeployKeysArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceDeployKeysResult>("gitlab:index/getInstanceDeployKeys:getInstanceDeployKeys", args ?? new GetInstanceDeployKeysArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getInstanceDeployKeys` data source allows to retrieve a list of deploy keys for a GitLab instance.
        /// 
        /// &gt; This data source requires administration privileges.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_keys.html#list-all-deploy-keys)
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
        ///     var example = GitLab.GetInstanceDeployKeys.Invoke(new()
        ///     {
        ///         Public = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetInstanceDeployKeysResult> Invoke(GetInstanceDeployKeysInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceDeployKeysResult>("gitlab:index/getInstanceDeployKeys:getInstanceDeployKeys", args ?? new GetInstanceDeployKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceDeployKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Only return deploy keys that are public.
        /// </summary>
        [Input("public")]
        public bool? Public { get; set; }

        public GetInstanceDeployKeysArgs()
        {
        }
        public static new GetInstanceDeployKeysArgs Empty => new GetInstanceDeployKeysArgs();
    }

    public sealed class GetInstanceDeployKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Only return deploy keys that are public.
        /// </summary>
        [Input("public")]
        public Input<bool>? Public { get; set; }

        public GetInstanceDeployKeysInvokeArgs()
        {
        }
        public static new GetInstanceDeployKeysInvokeArgs Empty => new GetInstanceDeployKeysInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceDeployKeysResult
    {
        /// <summary>
        /// The list of all deploy keys across all projects of the GitLab instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceDeployKeysDeployKeyResult> DeployKeys;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Only return deploy keys that are public.
        /// </summary>
        public readonly bool? Public;

        [OutputConstructor]
        private GetInstanceDeployKeysResult(
            ImmutableArray<Outputs.GetInstanceDeployKeysDeployKeyResult> deployKeys,

            string id,

            bool? @public)
        {
            DeployKeys = deployKeys;
            Id = id;
            Public = @public;
        }
    }
}
