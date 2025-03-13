// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.IntegrationJenkinsArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.IntegrationJenkinsState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.IntegrationJenkins` resource allows to manage the lifecycle of a project integration with Jenkins.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/integrations/#jenkins)
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
 * import com.pulumi.gitlab.IntegrationJenkins;
 * import com.pulumi.gitlab.IntegrationJenkinsArgs;
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
 *         var jenkins = new IntegrationJenkins("jenkins", IntegrationJenkinsArgs.builder()
 *             .project(awesomeProject.id())
 *             .jenkinsUrl("http://jenkins.example.com")
 *             .projectName("my_project_name")
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_jenkins`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_integration_jenkins.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * ```sh
 * $ pulumi import gitlab:index/integrationJenkins:IntegrationJenkins You can import a gitlab_integration_jenkins state using `&lt;resource&gt; &lt;project_id&gt;`:
 * ```
 * 
 * ```sh
 * $ pulumi import gitlab:index/integrationJenkins:IntegrationJenkins jenkins 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/integrationJenkins:IntegrationJenkins")
public class IntegrationJenkins extends com.pulumi.resources.CustomResource {
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
     * Enable SSL verification. Defaults to `true` (enabled).
     * 
     */
    @Export(name="enableSslVerification", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableSslVerification;

    /**
     * @return Enable SSL verification. Defaults to `true` (enabled).
     * 
     */
    public Output<Boolean> enableSslVerification() {
        return this.enableSslVerification;
    }
    /**
     * Jenkins URL like `http://jenkins.example.com`
     * 
     */
    @Export(name="jenkinsUrl", refs={String.class}, tree="[0]")
    private Output<String> jenkinsUrl;

    /**
     * @return Jenkins URL like `http://jenkins.example.com`
     * 
     */
    public Output<String> jenkinsUrl() {
        return this.jenkinsUrl;
    }
    /**
     * Enable notifications for merge request events.
     * 
     */
    @Export(name="mergeRequestEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> mergeRequestEvents;

    /**
     * @return Enable notifications for merge request events.
     * 
     */
    public Output<Boolean> mergeRequestEvents() {
        return this.mergeRequestEvents;
    }
    /**
     * Password for authentication with the Jenkins server, if authentication is required by the server.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return Password for authentication with the Jenkins server, if authentication is required by the server.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * ID of the project you want to activate integration on.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return ID of the project you want to activate integration on.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The URL-friendly project name. Example: `my_project_name`.
     * 
     */
    @Export(name="projectName", refs={String.class}, tree="[0]")
    private Output<String> projectName;

    /**
     * @return The URL-friendly project name. Example: `my_project_name`.
     * 
     */
    public Output<String> projectName() {
        return this.projectName;
    }
    /**
     * Enable notifications for push events.
     * 
     */
    @Export(name="pushEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> pushEvents;

    /**
     * @return Enable notifications for push events.
     * 
     */
    public Output<Boolean> pushEvents() {
        return this.pushEvents;
    }
    /**
     * Enable notifications for tag push events.
     * 
     */
    @Export(name="tagPushEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> tagPushEvents;

    /**
     * @return Enable notifications for tag push events.
     * 
     */
    public Output<Boolean> tagPushEvents() {
        return this.tagPushEvents;
    }
    /**
     * Username for authentication with the Jenkins server, if authentication is required by the server.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> username;

    /**
     * @return Username for authentication with the Jenkins server, if authentication is required by the server.
     * 
     */
    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationJenkins(java.lang.String name) {
        this(name, IntegrationJenkinsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationJenkins(java.lang.String name, IntegrationJenkinsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationJenkins(java.lang.String name, IntegrationJenkinsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationJenkins:IntegrationJenkins", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IntegrationJenkins(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationJenkinsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationJenkins:IntegrationJenkins", name, state, makeResourceOptions(options, id), false);
    }

    private static IntegrationJenkinsArgs makeArgs(IntegrationJenkinsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IntegrationJenkinsArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static IntegrationJenkins get(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationJenkinsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationJenkins(name, id, state, options);
    }
}
