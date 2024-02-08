// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.IntegrationEmailsOnPushArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.IntegrationEmailsOnPushState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.IntegrationEmailsOnPush` resource allows to manage the lifecycle of a project integration with Emails on Push Service.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#emails-on-push)
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
 * import com.pulumi.gitlab.IntegrationEmailsOnPush;
 * import com.pulumi.gitlab.IntegrationEmailsOnPushArgs;
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
 *         var emails = new IntegrationEmailsOnPush(&#34;emails&#34;, IntegrationEmailsOnPushArgs.builder()        
 *             .project(awesomeProject.id())
 *             .recipients(&#34;myrecipient@example.com myotherrecipient@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * You can import a gitlab_integration_emails_on_push state using the project ID, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush emails 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush")
public class IntegrationEmailsOnPush extends com.pulumi.resources.CustomResource {
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
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
     * 
     */
    @Export(name="branchesToBeNotified", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> branchesToBeNotified;

    /**
     * @return Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
     * 
     */
    public Output<Optional<String>> branchesToBeNotified() {
        return Codegen.optional(this.branchesToBeNotified);
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
     * Disable code diffs.
     * 
     */
    @Export(name="disableDiffs", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableDiffs;

    /**
     * @return Disable code diffs.
     * 
     */
    public Output<Optional<Boolean>> disableDiffs() {
        return Codegen.optional(this.disableDiffs);
    }
    /**
     * ID or full-path of the project you want to activate integration on.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return ID or full-path of the project you want to activate integration on.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Enable notifications for push events.
     * 
     */
    @Export(name="pushEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> pushEvents;

    /**
     * @return Enable notifications for push events.
     * 
     */
    public Output<Optional<Boolean>> pushEvents() {
        return Codegen.optional(this.pushEvents);
    }
    /**
     * Emails separated by whitespace.
     * 
     */
    @Export(name="recipients", refs={String.class}, tree="[0]")
    private Output<String> recipients;

    /**
     * @return Emails separated by whitespace.
     * 
     */
    public Output<String> recipients() {
        return this.recipients;
    }
    /**
     * Send from committer.
     * 
     */
    @Export(name="sendFromCommitterEmail", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> sendFromCommitterEmail;

    /**
     * @return Send from committer.
     * 
     */
    public Output<Optional<Boolean>> sendFromCommitterEmail() {
        return Codegen.optional(this.sendFromCommitterEmail);
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
     * Enable notifications for tag push events.
     * 
     */
    @Export(name="tagPushEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> tagPushEvents;

    /**
     * @return Enable notifications for tag push events.
     * 
     */
    public Output<Optional<Boolean>> tagPushEvents() {
        return Codegen.optional(this.tagPushEvents);
    }
    /**
     * Title of the integration.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return Title of the integration.
     * 
     */
    public Output<String> title() {
        return this.title;
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
    public IntegrationEmailsOnPush(String name) {
        this(name, IntegrationEmailsOnPushArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationEmailsOnPush(String name, IntegrationEmailsOnPushArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationEmailsOnPush(String name, IntegrationEmailsOnPushArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush", name, args == null ? IntegrationEmailsOnPushArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IntegrationEmailsOnPush(String name, Output<String> id, @Nullable IntegrationEmailsOnPushState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush", name, state, makeResourceOptions(options, id));
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
    public static IntegrationEmailsOnPush get(String name, Output<String> id, @Nullable IntegrationEmailsOnPushState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationEmailsOnPush(name, id, state, options);
    }
}
