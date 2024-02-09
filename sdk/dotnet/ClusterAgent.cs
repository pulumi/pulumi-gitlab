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
    /// The `gitlab.ClusterAgent` resource allows to manage the lifecycle of a GitLab Agent for Kubernetes.
    /// 
    /// &gt; Note that this resource only registers the agent, but doesn't configure it.
    ///    The configuration needs to be manually added as described in
    ///    [the docs](https://docs.gitlab.com/ee/user/clusters/agent/install/index.html#create-an-agent-configuration-file).
    ///    However, a `gitlab.RepositoryFile` resource may be used to achieve that.
    /// 
    /// &gt; Requires at least maintainer permissions on the project.
    /// 
    /// &gt; Requires at least GitLab 14.10
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/cluster_agents.html)
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
    ///     var example = new GitLab.ClusterAgent("example", new()
    ///     {
    ///         Project = "12345",
    ///     });
    /// 
    ///     // Optionally, configure the agent as described in
    ///     // https://docs.gitlab.com/ee/user/clusters/agent/install/index.html#create-an-agent-configuration-file
    ///     var exampleAgentConfig = new GitLab.RepositoryFile("exampleAgentConfig", new()
    ///     {
    ///         Project = example.Project,
    ///         Branch = "main",
    ///         FilePath = example.Name.Apply(name =&gt; $".gitlab/agents/{name}"),
    ///         Content = @"  gitops:
    ///     ...
    /// ",
    ///         AuthorEmail = "terraform@example.com",
    ///         AuthorName = "Terraform",
    ///         CommitMessage = example.Name.Apply(name =&gt; $"feature: add agent config for {name}"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitLab Agent for Kubernetes can be imported with the following command and the id pattern `&lt;project&gt;:&lt;agent-id&gt;`
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/clusterAgent:ClusterAgent example '12345:42'
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/clusterAgent:ClusterAgent")]
    public partial class ClusterAgent : global::Pulumi.CustomResource
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
        /// Create a ClusterAgent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterAgent(string name, ClusterAgentArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/clusterAgent:ClusterAgent", name, args ?? new ClusterAgentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterAgent(string name, Input<string> id, ClusterAgentState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/clusterAgent:ClusterAgent", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterAgent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterAgent Get(string name, Input<string> id, ClusterAgentState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterAgent(name, id, state, options);
        }
    }

    public sealed class ClusterAgentArgs : global::Pulumi.ResourceArgs
    {
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

        public ClusterAgentArgs()
        {
        }
        public static new ClusterAgentArgs Empty => new ClusterAgentArgs();
    }

    public sealed class ClusterAgentState : global::Pulumi.ResourceArgs
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
        /// The Name of the agent.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID or full path of the project maintained by the authenticated user.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ClusterAgentState()
        {
        }
        public static new ClusterAgentState Empty => new ClusterAgentState();
    }
}
