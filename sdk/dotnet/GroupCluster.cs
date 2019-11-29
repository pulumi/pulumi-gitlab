// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gitlab
{
    /// <summary>
    /// This resource allows you to create and manage group clusters for your GitLab groups.
    /// For further information on clusters, consult the [gitlab
    /// documentation](https://docs.gitlab.com/ce/user/group/clusters/index.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/group_cluster.html.markdown.
    /// </summary>
    public partial class GroupCluster : Pulumi.CustomResource
    {
        [Output("clusterType")]
        public Output<string> ClusterType { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The base domain of the cluster.
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The associated environment to the cluster. Defaults to `*`.
        /// </summary>
        [Output("environmentScope")]
        public Output<string?> EnvironmentScope { get; private set; } = null!;

        /// <summary>
        /// The id of the group to add the cluster to.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The URL to access the Kubernetes API.
        /// </summary>
        [Output("kubernetesApiUrl")]
        public Output<string> KubernetesApiUrl { get; private set; } = null!;

        /// <summary>
        /// The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        /// </summary>
        [Output("kubernetesAuthorizationType")]
        public Output<string?> KubernetesAuthorizationType { get; private set; } = null!;

        /// <summary>
        /// TLS certificate (needed if API is using a self-signed TLS certificate).
        /// </summary>
        [Output("kubernetesCaCert")]
        public Output<string?> KubernetesCaCert { get; private set; } = null!;

        /// <summary>
        /// The token to authenticate against Kubernetes.
        /// </summary>
        [Output("kubernetesToken")]
        public Output<string> KubernetesToken { get; private set; } = null!;

        /// <summary>
        /// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Output("managed")]
        public Output<bool?> Managed { get; private set; } = null!;

        /// <summary>
        /// The name of cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("platformType")]
        public Output<string> PlatformType { get; private set; } = null!;

        [Output("providerType")]
        public Output<string> ProviderType { get; private set; } = null!;


        /// <summary>
        /// Create a GroupCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupCluster(string name, GroupClusterArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupCluster:GroupCluster", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GroupCluster(string name, Input<string> id, GroupClusterState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupCluster:GroupCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupCluster Get(string name, Input<string> id, GroupClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupCluster(name, id, state, options);
        }
    }

    public sealed class GroupClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base domain of the cluster.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The associated environment to the cluster. Defaults to `*`.
        /// </summary>
        [Input("environmentScope")]
        public Input<string>? EnvironmentScope { get; set; }

        /// <summary>
        /// The id of the group to add the cluster to.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The URL to access the Kubernetes API.
        /// </summary>
        [Input("kubernetesApiUrl", required: true)]
        public Input<string> KubernetesApiUrl { get; set; } = null!;

        /// <summary>
        /// The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        /// </summary>
        [Input("kubernetesAuthorizationType")]
        public Input<string>? KubernetesAuthorizationType { get; set; }

        /// <summary>
        /// TLS certificate (needed if API is using a self-signed TLS certificate).
        /// </summary>
        [Input("kubernetesCaCert")]
        public Input<string>? KubernetesCaCert { get; set; }

        /// <summary>
        /// The token to authenticate against Kubernetes.
        /// </summary>
        [Input("kubernetesToken", required: true)]
        public Input<string> KubernetesToken { get; set; } = null!;

        /// <summary>
        /// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Input("managed")]
        public Input<bool>? Managed { get; set; }

        /// <summary>
        /// The name of cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupClusterArgs()
        {
        }
    }

    public sealed class GroupClusterState : Pulumi.ResourceArgs
    {
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The base domain of the cluster.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The associated environment to the cluster. Defaults to `*`.
        /// </summary>
        [Input("environmentScope")]
        public Input<string>? EnvironmentScope { get; set; }

        /// <summary>
        /// The id of the group to add the cluster to.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The URL to access the Kubernetes API.
        /// </summary>
        [Input("kubernetesApiUrl")]
        public Input<string>? KubernetesApiUrl { get; set; }

        /// <summary>
        /// The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        /// </summary>
        [Input("kubernetesAuthorizationType")]
        public Input<string>? KubernetesAuthorizationType { get; set; }

        /// <summary>
        /// TLS certificate (needed if API is using a self-signed TLS certificate).
        /// </summary>
        [Input("kubernetesCaCert")]
        public Input<string>? KubernetesCaCert { get; set; }

        /// <summary>
        /// The token to authenticate against Kubernetes.
        /// </summary>
        [Input("kubernetesToken")]
        public Input<string>? KubernetesToken { get; set; }

        /// <summary>
        /// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Input("managed")]
        public Input<bool>? Managed { get; set; }

        /// <summary>
        /// The name of cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platformType")]
        public Input<string>? PlatformType { get; set; }

        [Input("providerType")]
        public Input<string>? ProviderType { get; set; }

        public GroupClusterState()
        {
        }
    }
}
