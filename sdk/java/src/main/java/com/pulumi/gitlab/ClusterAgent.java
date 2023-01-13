// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ClusterAgentArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ClusterAgentState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.ClusterAgent` resource allows to manage the lifecycle of a GitLab Agent for Kubernetes.
 * 
 * &gt; Note that this resource only registers the agent, but doesn&#39;t configure it.
 *    The configuration needs to be manually added as described in
 *    [the docs](https://docs.gitlab.com/ee/user/clusters/agent/install/index.html#create-an-agent-configuration-file).
 *    However, a `gitlab.RepositoryFile` resource may be used to achieve that.
 * 
 * &gt; Requires at least maintainer permissions on the project.
 * 
 * &gt; Requires at least GitLab 14.10
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/cluster_agents.html)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.ClusterAgent;
 * import com.pulumi.gitlab.ClusterAgentArgs;
 * import com.pulumi.gitlab.RepositoryFile;
 * import com.pulumi.gitlab.RepositoryFileArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new ClusterAgent(&#34;example&#34;, ClusterAgentArgs.builder()        
 *             .project(&#34;12345&#34;)
 *             .build());
 * 
 *         var exampleAgentConfig = new RepositoryFile(&#34;exampleAgentConfig&#34;, RepositoryFileArgs.builder()        
 *             .project(example.project())
 *             .branch(&#34;main&#34;)
 *             .filePath(example.name().applyValue(name -&gt; String.format(&#34;.gitlab/agents/%s&#34;, name)))
 *             .content(&#34;&#34;&#34;
 *   gitops:
 *     ...
 *             &#34;&#34;&#34;)
 *             .authorEmail(&#34;terraform@example.com&#34;)
 *             .authorName(&#34;Terraform&#34;)
 *             .commitMessage(example.name().applyValue(name -&gt; String.format(&#34;feature: add agent config for %s&#34;, name)))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitLab Agent for Kubernetes can be imported with the following command and the id pattern `&lt;project&gt;:&lt;agent-id&gt;`
 * 
 * ```sh
 *  $ pulumi import gitlab:index/clusterAgent:ClusterAgent example &#39;12345:42&#39;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/clusterAgent:ClusterAgent")
public class ClusterAgent extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the agent.
     * 
     */
    @Export(name="agentId", type=Integer.class, parameters={})
    private Output<Integer> agentId;

    /**
     * @return The ID of the agent.
     * 
     */
    public Output<Integer> agentId() {
        return this.agentId;
    }
    /**
     * The ISO8601 datetime when the agent was created.
     * 
     */
    @Export(name="createdAt", type=String.class, parameters={})
    private Output<String> createdAt;

    /**
     * @return The ISO8601 datetime when the agent was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The ID of the user who created the agent.
     * 
     */
    @Export(name="createdByUserId", type=Integer.class, parameters={})
    private Output<Integer> createdByUserId;

    /**
     * @return The ID of the user who created the agent.
     * 
     */
    public Output<Integer> createdByUserId() {
        return this.createdByUserId;
    }
    /**
     * The Name of the agent.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The Name of the agent.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * ID or full path of the project maintained by the authenticated user.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return ID or full path of the project maintained by the authenticated user.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClusterAgent(String name) {
        this(name, ClusterAgentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClusterAgent(String name, ClusterAgentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClusterAgent(String name, ClusterAgentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/clusterAgent:ClusterAgent", name, args == null ? ClusterAgentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClusterAgent(String name, Output<String> id, @Nullable ClusterAgentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/clusterAgent:ClusterAgent", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ClusterAgent get(String name, Output<String> id, @Nullable ClusterAgentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClusterAgent(name, id, state, options);
    }
}