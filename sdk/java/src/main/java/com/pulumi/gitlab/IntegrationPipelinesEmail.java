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
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#pipeline-status-emails)
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
 * import com.pulumi.gitlab.IntegrationPipelinesEmail;
 * import com.pulumi.gitlab.IntegrationPipelinesEmailArgs;
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
 *         var email = new IntegrationPipelinesEmail("email", IntegrationPipelinesEmailArgs.builder()
 *             .project(awesomeProject.id())
 *             .recipients("gitlab}{@literal @}{@code user.create")
 *             .notifyOnlyBrokenPipelines(true)
 *             .branchesToBeNotified("all")
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_pipelines_email`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_integration_pipelines_email.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
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
    public IntegrationPipelinesEmail(java.lang.String name) {
        this(name, IntegrationPipelinesEmailArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationPipelinesEmail(java.lang.String name, IntegrationPipelinesEmailArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationPipelinesEmail(java.lang.String name, IntegrationPipelinesEmailArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IntegrationPipelinesEmail(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationPipelinesEmailState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail", name, state, makeResourceOptions(options, id), false);
    }

    private static IntegrationPipelinesEmailArgs makeArgs(IntegrationPipelinesEmailArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IntegrationPipelinesEmailArgs.Empty : args;
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
    public static IntegrationPipelinesEmail get(java.lang.String name, Output<java.lang.String> id, @Nullable IntegrationPipelinesEmailState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationPipelinesEmail(name, id, state, options);
    }
}
