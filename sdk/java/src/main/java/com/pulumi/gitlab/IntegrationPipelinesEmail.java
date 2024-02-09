// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.IntegrationPipelinesEmailArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.IntegrationPipelinesEmailState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.IntegrationPipelinesEmail` resource allows to manage the lifecycle of a project integration with Pipeline Emails Service.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#pipeline-emails)
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
 * import com.pulumi.gitlab.IntegrationPipelinesEmail;
 * import com.pulumi.gitlab.IntegrationPipelinesEmailArgs;
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
 *         var email = new IntegrationPipelinesEmail(&#34;email&#34;, IntegrationPipelinesEmailArgs.builder()        
 *             .project(awesomeProject.id())
 *             .recipients(&#34;gitlab@user.create&#34;)
 *             .notifyOnlyBrokenPipelines(true)
 *             .branchesToBeNotified(&#34;all&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * You can import a gitlab_integration_pipelines_email state using the project ID, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail email 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail")
public class IntegrationPipelinesEmail extends com.pulumi.resources.CustomResource {
    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
     * 
     */
    @Export(name="branchesToBeNotified", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> branchesToBeNotified;

    /**
     * @return Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
     * 
     */
    public Output<Optional<String>> branchesToBeNotified() {
        return Codegen.optional(this.branchesToBeNotified);
    }
    /**
     * Notify only broken pipelines. Default is true.
     * 
     */
    @Export(name="notifyOnlyBrokenPipelines", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> notifyOnlyBrokenPipelines;

    /**
     * @return Notify only broken pipelines. Default is true.
     * 
     */
    public Output<Optional<Boolean>> notifyOnlyBrokenPipelines() {
        return Codegen.optional(this.notifyOnlyBrokenPipelines);
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
     * ) email addresses where notifications are sent.
     * 
     */
    @Export(name="recipients", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> recipients;

    /**
     * @return ) email addresses where notifications are sent.
     * 
     */
    public Output<List<String>> recipients() {
        return this.recipients;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationPipelinesEmail(String name) {
        this(name, IntegrationPipelinesEmailArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationPipelinesEmail(String name, IntegrationPipelinesEmailArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationPipelinesEmail(String name, IntegrationPipelinesEmailArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail", name, args == null ? IntegrationPipelinesEmailArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IntegrationPipelinesEmail(String name, Output<String> id, @Nullable IntegrationPipelinesEmailState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail", name, state, makeResourceOptions(options, id));
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
    public static IntegrationPipelinesEmail get(String name, Output<String> id, @Nullable IntegrationPipelinesEmailState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationPipelinesEmail(name, id, state, options);
    }
}
