// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ServiceJiraArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ServiceJiraState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ServiceJira` resource allows to manage the lifecycle of a project integration with Jira.
 * 
 * &gt; This resource is deprecated. use `gitlab.IntegrationJira`instead!
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/services.html#jira)
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
 * import com.pulumi.gitlab.ServiceJira;
 * import com.pulumi.gitlab.ServiceJiraArgs;
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
 *         var jira = new ServiceJira("jira", ServiceJiraArgs.builder()
 *             .project(awesomeProject.id())
 *             .url("https://jira.example.com")
 *             .username("user")
 *             .password("mypass")
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
 * You can import a gitlab_service_jira state using the project ID, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/serviceJira:ServiceJira jira 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/serviceJira:ServiceJira")
public class ServiceJira extends com.pulumi.resources.CustomResource {
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
     * The base URL to the Jira instance API. Web URL value is used if not set. For example, [https://jira-api.example.com](https://jira-api.example.com).
     * 
     */
    @Export(name="apiUrl", refs={String.class}, tree="[0]")
    private Output<String> apiUrl;

    /**
     * @return The base URL to the Jira instance API. Web URL value is used if not set. For example, [https://jira-api.example.com](https://jira-api.example.com).
     * 
     */
    public Output<String> apiUrl() {
        return this.apiUrl;
    }
    /**
     * Enable comments inside Jira issues on each GitLab event (commit / merge request)
     * 
     */
    @Export(name="commentOnEventEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> commentOnEventEnabled;

    /**
     * @return Enable comments inside Jira issues on each GitLab event (commit / merge request)
     * 
     */
    public Output<Boolean> commentOnEventEnabled() {
        return this.commentOnEventEnabled;
    }
    /**
     * Enable notifications for commit events
     * 
     */
    @Export(name="commitEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> commitEvents;

    /**
     * @return Enable notifications for commit events
     * 
     */
    public Output<Boolean> commitEvents() {
        return this.commitEvents;
    }
    /**
     * Create time.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Create time.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Enable notifications for issues events.
     * 
     */
    @Export(name="issuesEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> issuesEvents;

    /**
     * @return Enable notifications for issues events.
     * 
     */
    public Output<Boolean> issuesEvents() {
        return this.issuesEvents;
    }
    /**
     * The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
     * 
     */
    @Export(name="jiraIssueTransitionId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> jiraIssueTransitionId;

    /**
     * @return The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
     * 
     */
    public Output<Optional<String>> jiraIssueTransitionId() {
        return Codegen.optional(this.jiraIssueTransitionId);
    }
    /**
     * Enable notifications for job events.
     * 
     */
    @Export(name="jobEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> jobEvents;

    /**
     * @return Enable notifications for job events.
     * 
     */
    public Output<Boolean> jobEvents() {
        return this.jobEvents;
    }
    /**
     * Enable notifications for merge request events
     * 
     */
    @Export(name="mergeRequestsEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> mergeRequestsEvents;

    /**
     * @return Enable notifications for merge request events
     * 
     */
    public Output<Boolean> mergeRequestsEvents() {
        return this.mergeRequestsEvents;
    }
    /**
     * Enable notifications for note events.
     * 
     */
    @Export(name="noteEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> noteEvents;

    /**
     * @return Enable notifications for note events.
     * 
     */
    public Output<Boolean> noteEvents() {
        return this.noteEvents;
    }
    /**
     * The password of the user created to be used with GitLab/JIRA.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return The password of the user created to be used with GitLab/JIRA.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * Enable notifications for pipeline events.
     * 
     */
    @Export(name="pipelineEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> pipelineEvents;

    /**
     * @return Enable notifications for pipeline events.
     * 
     */
    public Output<Boolean> pipelineEvents() {
        return this.pipelineEvents;
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
     * The short identifier for your JIRA project, all uppercase, e.g., PROJ.
     * 
     */
    @Export(name="projectKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return The short identifier for your JIRA project, all uppercase, e.g., PROJ.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
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
     * Enable notifications for tag_push events.
     * 
     */
    @Export(name="tagPushEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> tagPushEvents;

    /**
     * @return Enable notifications for tag_push events.
     * 
     */
    public Output<Boolean> tagPushEvents() {
        return this.tagPushEvents;
    }
    /**
     * Title.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return Title.
     * 
     */
    public Output<String> title() {
        return this.title;
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
     * The URL to the JIRA project which is being linked to this GitLab project. For example, [https://jira.example.com](https://jira.example.com).
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL to the JIRA project which is being linked to this GitLab project. For example, [https://jira.example.com](https://jira.example.com).
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * The username of the user created to be used with GitLab/JIRA.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return The username of the user created to be used with GitLab/JIRA.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceJira(java.lang.String name) {
        this(name, ServiceJiraArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceJira(java.lang.String name, ServiceJiraArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceJira(java.lang.String name, ServiceJiraArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/serviceJira:ServiceJira", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceJira(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceJiraState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/serviceJira:ServiceJira", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceJiraArgs makeArgs(ServiceJiraArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceJiraArgs.Empty : args;
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
    public static ServiceJira get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceJiraState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceJira(name, id, state, options);
    }
}
