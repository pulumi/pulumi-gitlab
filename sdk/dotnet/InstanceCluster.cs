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
    /// ## # gitlab\_instance\_cluster
    /// 
    /// This resource allows you to create and manage instance clusters for your GitLab instances.
    /// For further information on clusters, consult the [gitlab
    /// documentation](https://docs.gitlab.com/ee/user/instance/clusters/).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bar = new GitLab.InstanceCluster("bar", new GitLab.InstanceClusterArgs
    ///         {
    ///             Domain = "example.com",
    ///             Enabled = true,
    ///             EnvironmentScope = "*",
    ///             KubernetesApiUrl = "https://124.124.124",
    ///             KubernetesAuthorizationType = "rbac",
    ///             KubernetesCaCert = "some-cert",
    ///             KubernetesNamespace = "namespace",
    ///             KubernetesToken = "some-token",
    ///             ManagementProjectId = "123456",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class InstanceCluster : Pulumi.CustomResource
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
        /// The unique namespace related to the instance.
        /// </summary>
        [Output("kubernetesNamespace")]
        public Output<string?> KubernetesNamespace { get; private set; } = null!;

        /// <summary>
        /// The token to authenticate against Kubernetes. This attribute cannot be read.
        /// </summary>
        [Output("kubernetesToken")]
        public Output<string> KubernetesToken { get; private set; } = null!;

        /// <summary>
        /// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Output("managed")]
        public Output<bool?> Managed { get; private set; } = null!;

        /// <summary>
        /// The ID of the management project for the cluster.
        /// </summary>
        [Output("managementProjectId")]
        public Output<string?> ManagementProjectId { get; private set; } = null!;

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
        /// Create a InstanceCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceCluster(string name, InstanceClusterArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/instanceCluster:InstanceCluster", name, args ?? new InstanceClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceCluster(string name, Input<string> id, InstanceClusterState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/instanceCluster:InstanceCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceCluster Get(string name, Input<string> id, InstanceClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceCluster(name, id, state, options);
        }
    }

    public sealed class InstanceClusterArgs : Pulumi.ResourceArgs
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
        /// The unique namespace related to the instance.
        /// </summary>
        [Input("kubernetesNamespace")]
        public Input<string>? KubernetesNamespace { get; set; }

        /// <summary>
        /// The token to authenticate against Kubernetes. This attribute cannot be read.
        /// </summary>
        [Input("kubernetesToken", required: true)]
        public Input<string> KubernetesToken { get; set; } = null!;

        /// <summary>
        /// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Input("managed")]
        public Input<bool>? Managed { get; set; }

        /// <summary>
        /// The ID of the management project for the cluster.
        /// </summary>
        [Input("managementProjectId")]
        public Input<string>? ManagementProjectId { get; set; }

        /// <summary>
        /// The name of cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public InstanceClusterArgs()
        {
        }
    }

    public sealed class InstanceClusterState : Pulumi.ResourceArgs
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
        /// The unique namespace related to the instance.
        /// </summary>
        [Input("kubernetesNamespace")]
        public Input<string>? KubernetesNamespace { get; set; }

        /// <summary>
        /// The token to authenticate against Kubernetes. This attribute cannot be read.
        /// </summary>
        [Input("kubernetesToken")]
        public Input<string>? KubernetesToken { get; set; }

        /// <summary>
        /// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Input("managed")]
        public Input<bool>? Managed { get; set; }

        /// <summary>
        /// The ID of the management project for the cluster.
        /// </summary>
        [Input("managementProjectId")]
        public Input<string>? ManagementProjectId { get; set; }

        /// <summary>
        /// The name of cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platformType")]
        public Input<string>? PlatformType { get; set; }

        [Input("providerType")]
        public Input<string>? ProviderType { get; set; }

        public InstanceClusterState()
        {
        }
    }
}