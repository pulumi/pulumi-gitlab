// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectAccessTokenArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectAccessTokenState;
import com.pulumi.gitlab.outputs.ProjectAccessTokenRotationConfiguration;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectAccessToken` resource allows to manage the lifecycle of a project access token.
 * 
 * &gt; Observability scopes are in beta and may not work on all instances. See more details in [the documentation](https://docs.gitlab.com/operations/tracing/)
 * 
 * &gt; Use `rotation_configuration` to automatically rotate tokens instead of using `timestamp()` as timestamp will cause changes with every plan. `pulumi up` must still be run to rotate the token.
 * 
 * &gt; Due to [Automatic reuse detection](https://docs.gitlab.com/api/project_access_tokens/#automatic-reuse-detection) it&#39;s possible that a new Project Access Token will immediately be revoked. Check if an old process using the old token is running if this happens.
 * 
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/project_access_tokens/)
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
 * import com.pulumi.gitlab.ProjectAccessToken;
 * import com.pulumi.gitlab.ProjectAccessTokenArgs;
 * import com.pulumi.gitlab.ProjectVariable;
 * import com.pulumi.gitlab.ProjectVariableArgs;
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
 *         var example = new ProjectAccessToken("example", ProjectAccessTokenArgs.builder()
 *             .project("25")
 *             .name("Example project access token")
 *             .expiresAt("2020-03-14")
 *             .accessLevel("reporter")
 *             .scopes("api")
 *             .build());
 * 
 *         var exampleProjectVariable = new ProjectVariable("exampleProjectVariable", ProjectVariableArgs.builder()
 *             .project(exampleGitlabProject.id())
 *             .key("pat")
 *             .value(example.token())
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_access_token`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_access_token.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * A GitLab Project Access Token can be imported using a key composed of `&lt;project-id&gt;:&lt;token-id&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectAccessToken:ProjectAccessToken example &#34;12345:1&#34;
 * ```
 * 
 * NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
 * 
 */
@ResourceType(type="gitlab:index/projectAccessToken:ProjectAccessToken")
public class ProjectAccessToken extends com.pulumi.resources.CustomResource {
    /**
     * The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
     * 
     */
    @Export(name="accessLevel", refs={String.class}, tree="[0]")
    private Output<String> accessLevel;

    /**
     * @return The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
     * 
     */
    public Output<String> accessLevel() {
        return this.accessLevel;
    }
    /**
     * True if the token is active.
     * 
     */
    @Export(name="active", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> active;

    /**
     * @return True if the token is active.
     * 
     */
    public Output<Boolean> active() {
        return this.active;
    }
    /**
     * Time the token has been created, RFC3339 format.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Time the token has been created, RFC3339 format.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The description of the project access token.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the project access token.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
     * 
     */
    @Export(name="expiresAt", refs={String.class}, tree="[0]")
    private Output<String> expiresAt;

    /**
     * @return When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
     * 
     */
    public Output<String> expiresAt() {
        return this.expiresAt;
    }
    /**
     * The name of the project access token.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the project access token.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID or full path of the project.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or full path of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * True if the token is revoked.
     * 
     */
    @Export(name="revoked", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> revoked;

    /**
     * @return True if the token is revoked.
     * 
     */
    public Output<Boolean> revoked() {
        return this.revoked;
    }
    /**
     * The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
     * 
     */
    @Export(name="rotationConfiguration", refs={ProjectAccessTokenRotationConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ ProjectAccessTokenRotationConfiguration> rotationConfiguration;

    /**
     * @return The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
     * 
     */
    public Output<Optional<ProjectAccessTokenRotationConfiguration>> rotationConfiguration() {
        return Codegen.optional(this.rotationConfiguration);
    }
    /**
     * The scopes of the project access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    @Export(name="scopes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> scopes;

    /**
     * @return The scopes of the project access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }
    /**
     * The token of the project access token. **Note**: the token is not available for imported resources.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return The token of the project access token. **Note**: the token is not available for imported resources.
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * The user_id associated to the token.
     * 
     */
    @Export(name="userId", refs={Integer.class}, tree="[0]")
    private Output<Integer> userId;

    /**
     * @return The user_id associated to the token.
     * 
     */
    public Output<Integer> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectAccessToken(java.lang.String name) {
        this(name, ProjectAccessTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectAccessToken(java.lang.String name, ProjectAccessTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectAccessToken(java.lang.String name, ProjectAccessTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectAccessToken:ProjectAccessToken", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ProjectAccessToken(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectAccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectAccessToken:ProjectAccessToken", name, state, makeResourceOptions(options, id), false);
    }

    private static ProjectAccessTokenArgs makeArgs(ProjectAccessTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProjectAccessTokenArgs.Empty : args;
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
    public static ProjectAccessToken get(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectAccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectAccessToken(name, id, state, options);
    }
}
