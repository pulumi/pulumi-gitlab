// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.IntegrationMattermostArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.IntegrationMattermostState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.IntegrationMattermost` resource allows to manage the lifecycle of a project integration with Mattermost.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#mattermost-notifications)
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
 * import com.pulumi.gitlab.IntegrationMattermost;
 * import com.pulumi.gitlab.IntegrationMattermostArgs;
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
 *         var mattermost = new IntegrationMattermost("mattermost", IntegrationMattermostArgs.builder()
 *             .project(awesomeProject.id())
 *             .webhook("https://webhook.com")
 *             .username("myuser")
 *             .pushEvents(true)
 *             .pushChannel("push_chan")
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_mattermost`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_integration_mattermost.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * You can import a gitlab_integration_mattermost.mattermost state using the project ID, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/integrationMattermost:IntegrationMattermost mattermost 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/integrationMattermost:IntegrationMattermost")
public class IntegrationMattermost extends com.pulumi.resources.CustomResource {
    /**
     * Branches to send notifications for. Valid options are &#34;all&#34;, &#34;default&#34;, &#34;protected&#34;, and &#34;default*and*protected&#34;.
     * 
     */
    @Export(name="branchesToBeNotified", refs={String.class}, tree="[0]")
    private Output<String> branchesToBeNotified;

    /**
     * @return Branches to send notifications for. Valid options are &#34;all&#34;, &#34;default&#34;, &#34;protected&#34;, and &#34;default*and*protected&#34;.
     * 
     */
    public Output<String> branchesToBeNotified() {
        return this.branchesToBeNotified;
    }
    /**
     * The name of the channel to receive confidential issue events notifications.
     * 
     */
    @Export(name="confidentialIssueChannel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> confidentialIssueChannel;

    /**
     * @return The name of the channel to receive confidential issue events notifications.
     * 
     */
    public Output<Optional<String>> confidentialIssueChannel() {
        return Codegen.optional(this.confidentialIssueChannel);
    }
    /**
     * Enable notifications for confidential issues events.
     * 
     */
    @Export(name="confidentialIssuesEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> confidentialIssuesEvents;

    /**
     * @return Enable notifications for confidential issues events.
     * 
     */
    public Output<Boolean> confidentialIssuesEvents() {
        return this.confidentialIssuesEvents;
    }
    /**
     * The name of the channel to receive confidential note events notifications.
     * 
     */
    @Export(name="confidentialNoteChannel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> confidentialNoteChannel;

    /**
     * @return The name of the channel to receive confidential note events notifications.
     * 
     */
    public Output<Optional<String>> confidentialNoteChannel() {
        return Codegen.optional(this.confidentialNoteChannel);
    }
    /**
     * Enable notifications for confidential note events.
     * 
     */
    @Export(name="confidentialNoteEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> confidentialNoteEvents;

    /**
     * @return Enable notifications for confidential note events.
     * 
     */
    public Output<Boolean> confidentialNoteEvents() {
        return this.confidentialNoteEvents;
    }
    /**
     * The name of the channel to receive issue events notifications.
     * 
     */
    @Export(name="issueChannel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> issueChannel;

    /**
     * @return The name of the channel to receive issue events notifications.
     * 
     */
    public Output<Optional<String>> issueChannel() {
        return Codegen.optional(this.issueChannel);
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
     * The name of the channel to receive merge request events notifications.
     * 
     */
    @Export(name="mergeRequestChannel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mergeRequestChannel;

    /**
     * @return The name of the channel to receive merge request events notifications.
     * 
     */
    public Output<Optional<String>> mergeRequestChannel() {
        return Codegen.optional(this.mergeRequestChannel);
    }
    /**
     * Enable notifications for merge requests events.
     * 
     */
    @Export(name="mergeRequestsEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> mergeRequestsEvents;

    /**
     * @return Enable notifications for merge requests events.
     * 
     */
    public Output<Boolean> mergeRequestsEvents() {
        return this.mergeRequestsEvents;
    }
    /**
     * The name of the channel to receive note events notifications.
     * 
     */
    @Export(name="noteChannel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> noteChannel;

    /**
     * @return The name of the channel to receive note events notifications.
     * 
     */
    public Output<Optional<String>> noteChannel() {
        return Codegen.optional(this.noteChannel);
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
     * Send notifications for broken pipelines.
     * 
     */
    @Export(name="notifyOnlyBrokenPipelines", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> notifyOnlyBrokenPipelines;

    /**
     * @return Send notifications for broken pipelines.
     * 
     */
    public Output<Boolean> notifyOnlyBrokenPipelines() {
        return this.notifyOnlyBrokenPipelines;
    }
    /**
     * The name of the channel to receive pipeline events notifications.
     * 
     */
    @Export(name="pipelineChannel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pipelineChannel;

    /**
     * @return The name of the channel to receive pipeline events notifications.
     * 
     */
    public Output<Optional<String>> pipelineChannel() {
        return Codegen.optional(this.pipelineChannel);
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
     * The name of the channel to receive push events notifications.
     * 
     */
    @Export(name="pushChannel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pushChannel;

    /**
     * @return The name of the channel to receive push events notifications.
     * 
     */
    public Output<Optional<String>> pushChannel() {
        return Codegen.optional(this.pushChannel);
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
     * The name of the channel to receive tag push events notifications.
     * 
     */
    @Export(name="tagPushChannel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tagPushChannel;

    /**
     * @return The name of the channel to receive tag push events notifications.
     * 
     */
    public Output<Optional<String>> tagPushChannel() {
        return Codegen.optional(this.tagPushChannel);
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
     * Username to use.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> username;

    /**
     * @return Username to use.
     * 
     */
    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }
    /**
     * Webhook URL (Example, https://mattermost.yourdomain.com/hooks/...). This value cannot be imported.
     * 
     */
    @Export(name="webhook", refs={String.class}, tree="[0]")
    private Output<String> webhook;

    /**
     * @return Webhook URL (Example, https://mattermost.yourdomain.com/hooks/...). This value cannot be imported.
     * 
     */
    public Output<String> webhook() {
        return this.webhook;
    }
    /**
     * The name of the channel to receive wiki page events notifications.
     * 
     */
    @Export(name="wikiPageChannel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> wikiPageChannel;

    /**
     * @return The name of the channel to receive wiki page events notifications.
     * 
     */
    public Output<Optional<String>> wikiPageChannel() {
        return Codegen.optional(this.wikiPageChannel);
    }
    /**
     * Enable notifications for wiki page events.
     * 
     */
    @Export(name="wikiPageEvents", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> wikiPageEvents;

    /**
     * @return Enable notifications for wiki page events.
     * 
     */
    public Output<Boolean> wikiPageEvents() {
        return this.wikiPageEvents;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationMattermost(java.lang.String name) {
        this(name, IntegrationMattermostArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationMattermost(java.lang.String name, IntegrationMattermostArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationMattermost(java.lang.String name, IntegrationMattermostArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationMattermost:IntegrationMattermost", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IntegrationMattermost(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationMattermostState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationMattermost:IntegrationMattermost", name, state, makeResourceOptions(options, id), false);
    }

    private static IntegrationMattermostArgs makeArgs(IntegrationMattermostArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IntegrationMattermostArgs.Empty : args;
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
    public static IntegrationMattermost get(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationMattermostState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationMattermost(name, id, state, options);
    }
}
