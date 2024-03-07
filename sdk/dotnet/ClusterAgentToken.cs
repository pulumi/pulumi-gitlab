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
    /// The `gitlab.ClusterAgentToken` resource allows to manage the lifecycle of a token for a GitLab Agent for Kubernetes.
    /// 
    /// &gt; Requires at least maintainer permissions on the project.
    /// 
    /// &gt; Requires at least GitLab 15.0
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/cluster_agents.html#create-an-agent-token)
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// using Helm = Pulumi.Helm;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create token for an agent
    ///     var example = new GitLab.ClusterAgentToken("example", new()
    ///     {
    ///         Project = "12345",
    ///         AgentId = 42,
    ///         Description = "some token",
    ///     });
    /// 
    ///     var thisProject = GitLab.GetProject.Invoke(new()
    ///     {
    ///         PathWithNamespace = "my-org/example",
    ///     });
    /// 
    ///     var thisClusterAgent = new GitLab.ClusterAgent("thisClusterAgent", new()
    ///     {
    ///         Project = thisProject.Apply(getProjectResult =&gt; getProjectResult.Id),
    ///     });
    /// 
    ///     var thisClusterAgentToken = new GitLab.ClusterAgentToken("thisClusterAgentToken", new()
    ///     {
    ///         Project = thisProject.Apply(getProjectResult =&gt; getProjectResult.Id),
    ///         AgentId = thisClusterAgent.AgentId,
    ///         Description = "Token for the my-agent used with `gitlab-agent` Helm Chart",
    ///     });
    /// 
    ///     var gitlabAgent = new Helm.Index.Helm_release("gitlabAgent", new()
    ///     {
    ///         Name = "gitlab-agent",
    ///         Namespace = "gitlab-agent",
    ///         CreateNamespace = true,
    ///         Repository = "https://charts.gitlab.io",
    ///         Chart = "gitlab-agent",
    ///         Version = "1.2.0",
    ///         Set = new[]
    ///         {
    ///             
    ///             {
    ///                 { "name", "config.token" },
    ///                 { "value", thisClusterAgentToken.Token },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// A token for a GitLab Agent for Kubernetes can be imported with the following command and the id pattern `&lt;project&gt;:&lt;agent-id&gt;:&lt;token-id&gt;`:
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/clusterAgentToken:ClusterAgentToken example '12345:42:1'
    /// ```
    /// 
    /// ATTENTION: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
    /// </summary>
    [GitLabResourceType("gitlab:index/clusterAgentToken:ClusterAgentToken")]
    public partial class ClusterAgentToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the agent.
        /// </summary>
        [Output("agentId")]
        public Output<int> AgentId { get; private set; } = null!;

        /// <summary>
        /// The ISO8601 datetime when the agent was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The ID of the user who created the agent.
        /// </summary>
        [Output("createdByUserId")]
        public Output<int> CreatedByUserId { get; private set; } = null!;

        /// <summary>
        /// The Description for the agent.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ISO8601 datetime when the token was last used.
        /// </summary>
        [Output("lastUsedAt")]
        public Output<string> LastUsedAt { get; private set; } = null!;

        /// <summary>
        /// The Name of the agent.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID or full path of the project maintained by the authenticated user.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The status of the token. Valid values are `active`, `revoked`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The secret token for the agent. The `token` is not available in imported resources.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// The ID of the token.
        /// </summary>
        [Output("tokenId")]
        public Output<int> TokenId { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterAgentToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterAgentToken(string name, ClusterAgentTokenArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/clusterAgentToken:ClusterAgentToken", name, args ?? new ClusterAgentTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterAgentToken(string name, Input<string> id, ClusterAgentTokenState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/clusterAgentToken:ClusterAgentToken", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterAgentToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterAgentToken Get(string name, Input<string> id, ClusterAgentTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterAgentToken(name, id, state, options);
        }
    }

    public sealed class ClusterAgentTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the agent.
        /// </summary>
        [Input("agentId", required: true)]
        public Input<int> AgentId { get; set; } = null!;

        /// <summary>
        /// The Description for the agent.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Name of the agent.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID or full path of the project maintained by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public ClusterAgentTokenArgs()
        {
        }
        public static new ClusterAgentTokenArgs Empty => new ClusterAgentTokenArgs();
    }

    public sealed class ClusterAgentTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the agent.
        /// </summary>
        [Input("agentId")]
        public Input<int>? AgentId { get; set; }

        /// <summary>
        /// The ISO8601 datetime when the agent was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The ID of the user who created the agent.
        /// </summary>
        [Input("createdByUserId")]
        public Input<int>? CreatedByUserId { get; set; }

        /// <summary>
        /// The Description for the agent.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ISO8601 datetime when the token was last used.
        /// </summary>
        [Input("lastUsedAt")]
        public Input<string>? LastUsedAt { get; set; }

        /// <summary>
        /// The Name of the agent.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID or full path of the project maintained by the authenticated user.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The status of the token. Valid values are `active`, `revoked`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The secret token for the agent. The `token` is not available in imported resources.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The ID of the token.
        /// </summary>
        [Input("tokenId")]
        public Input<int>? TokenId { get; set; }

        public ClusterAgentTokenState()
        {
        }
        public static new ClusterAgentTokenState Empty => new ClusterAgentTokenState();
    }
}
