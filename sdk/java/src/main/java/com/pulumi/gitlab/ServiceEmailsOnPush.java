// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ServiceEmailsOnPushArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ServiceEmailsOnPushState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ServiceEmailsOnPush` resource allows to manage the lifecycle of a project integration with Emails on Push Service.
 * 
 * &gt; This resource is deprecated. Please use `gitlab.IntegrationEmailsOnPush` instead!
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#emails-on-push)
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
 * import com.pulumi.gitlab.ServiceEmailsOnPush;
 * import com.pulumi.gitlab.ServiceEmailsOnPushArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var awesomeProject = new Project("awesomeProject", ProjectArgs.builder()
 *             .name("awesome_project")
 *             .description("My awesome project.")
 *             .visibilityLevel("public")
 *             .build());
 * 
 *         var emails = new ServiceEmailsOnPush("emails", ServiceEmailsOnPushArgs.builder()
 *             .project(awesomeProject.id())
 *             .recipients("myrecipient}{@literal @}{@code example.com myotherrecipient}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_service_emails_on_push`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_service_emails_on_push.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * You can import a gitlab_service_emails_on_push state using the project ID, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush emails 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush")
public class ServiceEmailsOnPush extends com.pulumi.resources.CustomResource {
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
    public ServiceEmailsOnPush(java.lang.String name) {
        this(name, ServiceEmailsOnPushArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceEmailsOnPush(java.lang.String name, ServiceEmailsOnPushArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceEmailsOnPush(java.lang.String name, ServiceEmailsOnPushArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceEmailsOnPush(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceEmailsOnPushState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceEmailsOnPushArgs makeArgs(ServiceEmailsOnPushArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceEmailsOnPushArgs.Empty : args;
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
    public static ServiceEmailsOnPush get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceEmailsOnPushState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceEmailsOnPush(name, id, state, options);
    }
}
