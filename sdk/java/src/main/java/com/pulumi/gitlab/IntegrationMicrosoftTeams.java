// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.IntegrationMicrosoftTeamsArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.IntegrationMicrosoftTeamsState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.IntegrationMicrosoftTeams` resource allows to manage the lifecycle of a project integration with Microsoft Teams.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#microsoft-teams-notifications)
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
 * import com.pulumi.gitlab.IntegrationMicrosoftTeams;
 * import com.pulumi.gitlab.IntegrationMicrosoftTeamsArgs;
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
 *         var teams = new IntegrationMicrosoftTeams("teams", IntegrationMicrosoftTeamsArgs.builder()
 *             .project(awesomeProject.id())
 *             .webhook("https://testurl.com/?token=XYZ")
 *             .pushEvents(true)
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_microsoft_teams`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_integration_microsoft_teams.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * You can import a gitlab_integration_microsoft_teams state using the project ID, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/integrationMicrosoftTeams:IntegrationMicrosoftTeams teams 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/integrationMicrosoftTeams:IntegrationMicrosoftTeams")
public class IntegrationMicrosoftTeams extends com.pulumi.resources.CustomResource {
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
     * Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
     * 
     */
    @Export(name="branchesToBeNotified", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> branchesToBeNotified;

    /**
     * @return Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
     * 
     */
    public Output<Optional<String>> branchesToBeNotified() {
        return Codegen.optional(this.branchesToBeNotified);
    }
    /**
     * Enable notifications for confidential issue events
     * 
     */
    @Export(name="confidentialIssuesEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> confidentialIssuesEvents;

    /**
     * @return Enable notifications for confidential issue events
     * 
     */
    public Output<Optional<Boolean>> confidentialIssuesEvents() {
        return Codegen.optional(this.confidentialIssuesEvents);
    }
    /**
     * Enable notifications for confidential note events
     * 
     */
    @Export(name="confidentialNoteEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> confidentialNoteEvents;

    /**
     * @return Enable notifications for confidential note events
     * 
     */
    public Output<Optional<Boolean>> confidentialNoteEvents() {
        return Codegen.optional(this.confidentialNoteEvents);
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
     * Enable notifications for issue events
     * 
     */
    @Export(name="issuesEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> issuesEvents;

    /**
     * @return Enable notifications for issue events
     * 
     */
    public Output<Optional<Boolean>> issuesEvents() {
        return Codegen.optional(this.issuesEvents);
    }
    /**
     * Enable notifications for merge request events
     * 
     */
    @Export(name="mergeRequestsEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mergeRequestsEvents;

    /**
     * @return Enable notifications for merge request events
     * 
     */
    public Output<Optional<Boolean>> mergeRequestsEvents() {
        return Codegen.optional(this.mergeRequestsEvents);
    }
    /**
     * Enable notifications for note events
     * 
     */
    @Export(name="noteEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> noteEvents;

    /**
     * @return Enable notifications for note events
     * 
     */
    public Output<Optional<Boolean>> noteEvents() {
        return Codegen.optional(this.noteEvents);
    }
    /**
     * Send notifications for broken pipelines
     * 
     */
    @Export(name="notifyOnlyBrokenPipelines", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> notifyOnlyBrokenPipelines;

    /**
     * @return Send notifications for broken pipelines
     * 
     */
    public Output<Optional<Boolean>> notifyOnlyBrokenPipelines() {
        return Codegen.optional(this.notifyOnlyBrokenPipelines);
    }
    /**
     * Enable notifications for pipeline events
     * 
     */
    @Export(name="pipelineEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> pipelineEvents;

    /**
     * @return Enable notifications for pipeline events
     * 
     */
    public Output<Optional<Boolean>> pipelineEvents() {
        return Codegen.optional(this.pipelineEvents);
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
     * Enable notifications for push events
     * 
     */
    @Export(name="pushEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> pushEvents;

    /**
     * @return Enable notifications for push events
     * 
     */
    public Output<Optional<Boolean>> pushEvents() {
        return Codegen.optional(this.pushEvents);
    }
    /**
     * Enable notifications for tag push events
     * 
     */
    @Export(name="tagPushEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> tagPushEvents;

    /**
     * @return Enable notifications for tag push events
     * 
     */
    public Output<Optional<Boolean>> tagPushEvents() {
        return Codegen.optional(this.tagPushEvents);
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
     * The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
     * 
     */
    @Export(name="webhook", refs={String.class}, tree="[0]")
    private Output<String> webhook;

    /**
     * @return The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
     * 
     */
    public Output<String> webhook() {
        return this.webhook;
    }
    /**
     * Enable notifications for wiki page events
     * 
     */
    @Export(name="wikiPageEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> wikiPageEvents;

    /**
     * @return Enable notifications for wiki page events
     * 
     */
    public Output<Optional<Boolean>> wikiPageEvents() {
        return Codegen.optional(this.wikiPageEvents);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationMicrosoftTeams(java.lang.String name) {
        this(name, IntegrationMicrosoftTeamsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationMicrosoftTeams(java.lang.String name, IntegrationMicrosoftTeamsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationMicrosoftTeams(java.lang.String name, IntegrationMicrosoftTeamsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationMicrosoftTeams:IntegrationMicrosoftTeams", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IntegrationMicrosoftTeams(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationMicrosoftTeamsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationMicrosoftTeams:IntegrationMicrosoftTeams", name, state, makeResourceOptions(options, id), false);
    }

    private static IntegrationMicrosoftTeamsArgs makeArgs(IntegrationMicrosoftTeamsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IntegrationMicrosoftTeamsArgs.Empty : args;
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
    public static IntegrationMicrosoftTeams get(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationMicrosoftTeamsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationMicrosoftTeams(name, id, state, options);
    }
}
