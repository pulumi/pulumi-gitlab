// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.ValueStreamAnalyticsArgs;
import com.pulumi.gitlab.inputs.ValueStreamAnalyticsState;
import com.pulumi.gitlab.outputs.ValueStreamAnalyticsStage;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ValueStreamAnalytics` resource allows to manage the lifecycle of value stream analytics.
 * 
 * &gt; This resource requires a GitLab Enterprise instance with a Premium license to create custom value stream analytics.
 * 
 * **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/api/graphql/reference/#mutationvaluestreamcreate)
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
 * import com.pulumi.gitlab.ValueStreamAnalytics;
 * import com.pulumi.gitlab.ValueStreamAnalyticsArgs;
 * import com.pulumi.gitlab.inputs.ValueStreamAnalyticsStageArgs;
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
 *         var project = new ValueStreamAnalytics("project", ValueStreamAnalyticsArgs.builder()
 *             .name("TEST")
 *             .projectFullPath("test/project")
 *             .stages(            
 *                 ValueStreamAnalyticsStageArgs.builder()
 *                     .name("Issue")
 *                     .custom(false)
 *                     .hidden(false)
 *                     .build(),
 *                 ValueStreamAnalyticsStageArgs.builder()
 *                     .name("Issue Labels")
 *                     .custom(true)
 *                     .hidden(false)
 *                     .start_event_identifier("ISSUE_LABEL_ADDED")
 *                     .start_event_label_id("gid://gitlab/ProjectLabel/0")
 *                     .end_event_identifier("ISSUE_LABEL_REMOVED")
 *                     .end_event_label_id("gid://gitlab/ProjectLabel/1")
 *                     .build())
 *             .build());
 * 
 *         var group = new ValueStreamAnalytics("group", ValueStreamAnalyticsArgs.builder()
 *             .name("TEST")
 *             .groupFullPath("test/group")
 *             .stages(            
 *                 ValueStreamAnalyticsStageArgs.builder()
 *                     .name("Issue")
 *                     .custom(false)
 *                     .hidden(false)
 *                     .build(),
 *                 ValueStreamAnalyticsStageArgs.builder()
 *                     .name("Issue Labels")
 *                     .custom(true)
 *                     .hidden(false)
 *                     .start_event_identifier("ISSUE_LABEL_ADDED")
 *                     .start_event_label_id("gid://gitlab/GroupLabel/0")
 *                     .end_event_identifier("ISSUE_LABEL_REMOVED")
 *                     .end_event_label_id("gid://gitlab/GroupLabel/1")
 *                     .build())
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_value_stream_analytics`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_value_stream_analytics.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * Gitlab value stream analytics can be imported with a key composed of `&lt;full_path_type&gt;:&lt;full_path&gt;:&lt;value_stream_id&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/valueStreamAnalytics:ValueStreamAnalytics group &#34;group:people/engineers:42&#34;
 * ```
 * 
 * ```sh
 * $ pulumi import gitlab:index/valueStreamAnalytics:ValueStreamAnalytics project &#34;project:projects/sample:43&#34;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/valueStreamAnalytics:ValueStreamAnalytics")
public class ValueStreamAnalytics extends com.pulumi.resources.CustomResource {
    /**
     * Full path of the group the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**
     * 
     */
    @Export(name="groupFullPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupFullPath;

    /**
     * @return Full path of the group the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**
     * 
     */
    public Output<Optional<String>> groupFullPath() {
        return Codegen.optional(this.groupFullPath);
    }
    /**
     * The name of the value stream
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the value stream
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Full path of the project the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**
     * 
     */
    @Export(name="projectFullPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectFullPath;

    /**
     * @return Full path of the project the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**
     * 
     */
    public Output<Optional<String>> projectFullPath() {
        return Codegen.optional(this.projectFullPath);
    }
    /**
     * Stages of the value stream
     * 
     */
    @Export(name="stages", refs={List.class,ValueStreamAnalyticsStage.class}, tree="[0,1]")
    private Output<List<ValueStreamAnalyticsStage>> stages;

    /**
     * @return Stages of the value stream
     * 
     */
    public Output<List<ValueStreamAnalyticsStage>> stages() {
        return this.stages;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ValueStreamAnalytics(String name) {
        this(name, ValueStreamAnalyticsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ValueStreamAnalytics(String name, ValueStreamAnalyticsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ValueStreamAnalytics(String name, ValueStreamAnalyticsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/valueStreamAnalytics:ValueStreamAnalytics", name, args == null ? ValueStreamAnalyticsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ValueStreamAnalytics(String name, Output<String> id, @Nullable ValueStreamAnalyticsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/valueStreamAnalytics:ValueStreamAnalytics", name, state, makeResourceOptions(options, id));
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
    public static ValueStreamAnalytics get(String name, Output<String> id, @Nullable ValueStreamAnalyticsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ValueStreamAnalytics(name, id, state, options);
    }
}
