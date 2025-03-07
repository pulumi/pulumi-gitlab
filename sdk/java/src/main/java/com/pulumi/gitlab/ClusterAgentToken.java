// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ClusterAgentTokenArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ClusterAgentTokenState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ClusterAgentToken` resource allows to manage the lifecycle of a token for a GitLab Agent for Kubernetes.
 * 
 * &gt; Requires at least maintainer permissions on the project.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/cluster_agents/#create-an-agent-token)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.ClusterAgentToken;
 * import com.pulumi.gitlab.ClusterAgentTokenArgs;
 * import com.pulumi.gitlab.GitlabFunctions;
 * import com.pulumi.gitlab.inputs.GetProjectArgs;
 * import com.pulumi.gitlab.ClusterAgent;
 * import com.pulumi.gitlab.ClusterAgentArgs;
 * import com.pulumi.helm.release;
 * import com.pulumi.helm.ReleaseArgs;
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
 *         // Create token for an agent
 *         var example = new ClusterAgentToken("example", ClusterAgentTokenArgs.builder()
 *             .project("12345")
 *             .agentId(42)
 *             .name("some-token")
 *             .description("some token")
 *             .build());
 * 
 *         // The following example creates a GitLab Agent for Kubernetes in a given project,
 *         // creates a token and install the `gitlab-agent` Helm Chart.
 *         // (see https://gitlab.com/gitlab-org/charts/gitlab-agent)
 *         final var this = GitlabFunctions.getProject(GetProjectArgs.builder()
 *             .pathWithNamespace("my-org/example")
 *             .build());
 * 
 *         var thisClusterAgent = new ClusterAgent("thisClusterAgent", ClusterAgentArgs.builder()
 *             .project(this_.id())
 *             .name("my-agent")
 *             .build());
 * 
 *         var thisClusterAgentToken = new ClusterAgentToken("thisClusterAgentToken", ClusterAgentTokenArgs.builder()
 *             .project(this_.id())
 *             .agentId(thisClusterAgent.agentId())
 *             .name("my-agent-token")
 *             .description("Token for the my-agent used with `gitlab-agent` Helm Chart")
 *             .build());
 * 
 *         var gitlabAgent = new Release("gitlabAgent", ReleaseArgs.builder()
 *             .name("gitlab-agent")
 *             .namespace("gitlab-agent")
 *             .createNamespace(true)
 *             .repository("https://charts.gitlab.io")
 *             .chart("gitlab-agent")
 *             .version("1.2.0")
 *             .set(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_cluster_agent_token`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_cluster_agent_token.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * A token for a GitLab Agent for Kubernetes can be imported with the following command and the id pattern `&lt;project&gt;:&lt;agent-id&gt;:&lt;token-id&gt;`:
 * 
 * ```sh
 * $ pulumi import gitlab:index/clusterAgentToken:ClusterAgentToken example &#39;12345:42:1&#39;
 * ```
 * 
 * ATTENTION: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
 * 
 */
@ResourceType(type="gitlab:index/clusterAgentToken:ClusterAgentToken")
public class ClusterAgentToken extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the agent.
     * 
     */
    @Export(name="agentId", refs={Integer.class}, tree="[0]")
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
    @Export(name="createdAt", refs={String.class}, tree="[0]")
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
    @Export(name="createdByUserId", refs={Integer.class}, tree="[0]")
    private Output<Integer> createdByUserId;

    /**
     * @return The ID of the user who created the agent.
     * 
     */
    public Output<Integer> createdByUserId() {
        return this.createdByUserId;
    }
    /**
     * The Description for the agent.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The Description for the agent.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The ISO8601 datetime when the token was last used.
     * 
     */
    @Export(name="lastUsedAt", refs={String.class}, tree="[0]")
    private Output<String> lastUsedAt;

    /**
     * @return The ISO8601 datetime when the token was last used.
     * 
     */
    public Output<String> lastUsedAt() {
        return this.lastUsedAt;
    }
    /**
     * The Name of the agent.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
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
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return ID or full path of the project maintained by the authenticated user.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The status of the token. Valid values are `active`, `revoked`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the token. Valid values are `active`, `revoked`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The secret token for the agent. The `token` is not available in imported resources.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return The secret token for the agent. The `token` is not available in imported resources.
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * The ID of the token.
     * 
     */
    @Export(name="tokenId", refs={Integer.class}, tree="[0]")
    private Output<Integer> tokenId;

    /**
     * @return The ID of the token.
     * 
     */
    public Output<Integer> tokenId() {
        return this.tokenId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClusterAgentToken(String name) {
        this(name, ClusterAgentTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClusterAgentToken(String name, ClusterAgentTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClusterAgentToken(String name, ClusterAgentTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/clusterAgentToken:ClusterAgentToken", name, args == null ? ClusterAgentTokenArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClusterAgentToken(String name, Output<String> id, @Nullable ClusterAgentTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/clusterAgentToken:ClusterAgentToken", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "token"
            ))
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
    public static ClusterAgentToken get(String name, Output<String> id, @Nullable ClusterAgentTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClusterAgentToken(name, id, state, options);
    }
}
