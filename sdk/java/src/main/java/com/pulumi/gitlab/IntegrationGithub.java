// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.IntegrationGithubArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.IntegrationGithubState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.IntegrationGithub` resource allows to manage the lifecycle of a project integration with GitHub.
 * 
 * &gt; This resource requires a GitLab Enterprise instance.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#github)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectArgs;
 * import com.pulumi.gitlab.IntegrationGithub;
 * import com.pulumi.gitlab.IntegrationGithubArgs;
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
 *         var awesomeProject = new Project(&#34;awesomeProject&#34;, ProjectArgs.builder()        
 *             .description(&#34;My awesome project.&#34;)
 *             .visibilityLevel(&#34;public&#34;)
 *             .build());
 * 
 *         var github = new IntegrationGithub(&#34;github&#34;, IntegrationGithubArgs.builder()        
 *             .project(awesomeProject.id())
 *             .token(&#34;REDACTED&#34;)
 *             .repositoryUrl(&#34;https://github.com/gitlabhq/terraform-provider-gitlab&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import gitlab:index/integrationGithub:IntegrationGithub You can import a gitlab_integration_github state using `&lt;resource&gt; &lt;project_id&gt;`
 * ```
 * 
 * ```sh
 *  $ pulumi import gitlab:index/integrationGithub:IntegrationGithub github 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/integrationGithub:IntegrationGithub")
public class IntegrationGithub extends com.pulumi.resources.CustomResource {
    /**
     * Whether the integration is active.
     * 
     */
    @Export(name="active", type=Boolean.class, parameters={})
    private Output<Boolean> active;

    /**
     * @return Whether the integration is active.
     * 
     */
    public Output<Boolean> active() {
        return this.active;
    }
    /**
     * Create time.
     * 
     */
    @Export(name="createdAt", type=String.class, parameters={})
    private Output<String> createdAt;

    /**
     * @return Create time.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * ID of the project you want to activate integration on.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return ID of the project you want to activate integration on.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
     * 
     */
    @Export(name="repositoryUrl", type=String.class, parameters={})
    private Output<String> repositoryUrl;

    /**
     * @return The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
     * 
     */
    public Output<String> repositoryUrl() {
        return this.repositoryUrl;
    }
    /**
     * Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
     * 
     */
    @Export(name="staticContext", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> staticContext;

    /**
     * @return Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
     * 
     */
    public Output<Optional<Boolean>> staticContext() {
        return Codegen.optional(this.staticContext);
    }
    /**
     * Title.
     * 
     */
    @Export(name="title", type=String.class, parameters={})
    private Output<String> title;

    /**
     * @return Title.
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * A GitHub personal access token with at least `repo:status` scope.
     * 
     */
    @Export(name="token", type=String.class, parameters={})
    private Output<String> token;

    /**
     * @return A GitHub personal access token with at least `repo:status` scope.
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * Update time.
     * 
     */
    @Export(name="updatedAt", type=String.class, parameters={})
    private Output<String> updatedAt;

    /**
     * @return Update time.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationGithub(String name) {
        this(name, IntegrationGithubArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationGithub(String name, IntegrationGithubArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationGithub(String name, IntegrationGithubArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationGithub:IntegrationGithub", name, args == null ? IntegrationGithubArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IntegrationGithub(String name, Output<String> id, @Nullable IntegrationGithubState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationGithub:IntegrationGithub", name, state, makeResourceOptions(options, id));
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
    public static IntegrationGithub get(String name, Output<String> id, @Nullable IntegrationGithubState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationGithub(name, id, state, options);
    }
}
