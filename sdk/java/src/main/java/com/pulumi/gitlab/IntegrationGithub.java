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
 * The `gitlab.IntegrationGithub` resource manages the lifecycle of a project integration with GitHub.
 * 
 * &gt; This resource requires a GitLab Enterprise instance.
 * 
 * &gt; This resource is deprecated and will be removed in 19.0. Use `gitlab.ProjectIntegrationGithub` instead.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#github)
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
 *         var awesomeProject = new Project("awesomeProject", ProjectArgs.builder()
 *             .name("awesome_project")
 *             .description("My awesome project.")
 *             .visibilityLevel("public")
 *             .build());
 * 
 *         var github = new IntegrationGithub("github", IntegrationGithubArgs.builder()
 *             .project(awesomeProject.id())
 *             .token("REDACTED")
 *             .repositoryUrl("https://github.com/gitlabhq/terraform-provider-gitlab")
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
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_integration_github`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_integration_github.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Importing using the CLI is supported with the following syntax:
 * 
 * ```sh
 * $ pulumi import gitlab:index/integrationGithub:IntegrationGithub You can import a gitlab_integration_github state using `&lt;resource&gt; &lt;project_id&gt;`:
 * ```
 * 
 * ```sh
 * $ pulumi import gitlab:index/integrationGithub:IntegrationGithub github 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/integrationGithub:IntegrationGithub")
public class IntegrationGithub extends com.pulumi.resources.CustomResource {
    /**
     * Whether the integration is active.
     * 
     */
    @Export(name="active", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> active;

    /**
     * @return Whether the integration is active.
     * 
     */
    public Output<Boolean> active() {
        return this.active;
    }
    /**
     * Creation time.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Creation time.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * ID of the project you want to activate the integration on.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return ID of the project you want to activate the integration on.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    @Export(name="repositoryUrl", refs={String.class}, tree="[0]")
    private Output<String> repositoryUrl;

    public Output<String> repositoryUrl() {
        return this.repositoryUrl;
    }
    /**
     * Append the instance name instead of the branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
     * 
     */
    @Export(name="staticContext", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> staticContext;

    /**
     * @return Append the instance name instead of the branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
     * 
     */
    public Output<Optional<Boolean>> staticContext() {
        return Codegen.optional(this.staticContext);
    }
    /**
     * The title of this resource.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return The title of this resource.
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * A GitHub personal access token with at least the `repo:status` scope.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return A GitHub personal access token with at least the `repo:status` scope.
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * Update time.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
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
    public IntegrationGithub(java.lang.String name) {
        this(name, IntegrationGithubArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationGithub(java.lang.String name, IntegrationGithubArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationGithub(java.lang.String name, IntegrationGithubArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationGithub:IntegrationGithub", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IntegrationGithub(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationGithubState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationGithub:IntegrationGithub", name, state, makeResourceOptions(options, id), false);
    }

    private static IntegrationGithubArgs makeArgs(IntegrationGithubArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IntegrationGithubArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static IntegrationGithub get(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationGithubState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationGithub(name, id, state, options);
    }
}
