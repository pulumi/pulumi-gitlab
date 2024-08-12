// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.IntegrationCustomIssueTrackerArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.IntegrationCustomIssueTrackerState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.IntegrationCustomIssueTracker` resource allows to manage the lifecycle of a project integration with Custom Issue Tracker.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#custom-issue-tracker)
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
 * import com.pulumi.gitlab.IntegrationCustomIssueTracker;
 * import com.pulumi.gitlab.IntegrationCustomIssueTrackerArgs;
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
 *         var tracker = new IntegrationCustomIssueTracker("tracker", IntegrationCustomIssueTrackerArgs.builder()
 *             .project(awesomeProject.id())
 *             .projectUrl("https://customtracker.com/issues")
 *             .issuesUrl("https://customtracker.com/TEST-:id")
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
 * You can import a gitlab_integration_custom_issue_tracker state using the project ID, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/integrationCustomIssueTracker:IntegrationCustomIssueTracker tracker 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/integrationCustomIssueTracker:IntegrationCustomIssueTracker")
public class IntegrationCustomIssueTracker extends com.pulumi.resources.CustomResource {
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
     * The ISO8601 date/time that this integration was activated at in UTC.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The ISO8601 date/time that this integration was activated at in UTC.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The URL to view an issue in the external issue tracker. Must contain :id.
     * 
     */
    @Export(name="issuesUrl", refs={String.class}, tree="[0]")
    private Output<String> issuesUrl;

    /**
     * @return The URL to view an issue in the external issue tracker. Must contain :id.
     * 
     */
    public Output<String> issuesUrl() {
        return this.issuesUrl;
    }
    /**
     * The ID or full path of the project for the custom issue tracker.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or full path of the project for the custom issue tracker.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The URL to the project in the external issue tracker.
     * 
     */
    @Export(name="projectUrl", refs={String.class}, tree="[0]")
    private Output<String> projectUrl;

    /**
     * @return The URL to the project in the external issue tracker.
     * 
     */
    public Output<String> projectUrl() {
        return this.projectUrl;
    }
    /**
     * The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     * 
     */
    @Export(name="slug", refs={String.class}, tree="[0]")
    private Output<String> slug;

    /**
     * @return The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     * 
     */
    public Output<String> slug() {
        return this.slug;
    }
    /**
     * The ISO8601 date/time that this integration was last updated at in UTC.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The ISO8601 date/time that this integration was last updated at in UTC.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationCustomIssueTracker(java.lang.String name) {
        this(name, IntegrationCustomIssueTrackerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationCustomIssueTracker(java.lang.String name, IntegrationCustomIssueTrackerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationCustomIssueTracker(java.lang.String name, IntegrationCustomIssueTrackerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationCustomIssueTracker:IntegrationCustomIssueTracker", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IntegrationCustomIssueTracker(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationCustomIssueTrackerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationCustomIssueTracker:IntegrationCustomIssueTracker", name, state, makeResourceOptions(options, id), false);
    }

    private static IntegrationCustomIssueTrackerArgs makeArgs(IntegrationCustomIssueTrackerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IntegrationCustomIssueTrackerArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static IntegrationCustomIssueTracker get(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationCustomIssueTrackerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationCustomIssueTracker(name, id, state, options);
    }
}
