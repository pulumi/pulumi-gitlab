// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.PipelineScheduleArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.PipelineScheduleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.PipelineSchedule` resource allows to manage the lifecycle of a scheduled pipeline.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_schedules.html)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.PipelineSchedule;
 * import com.pulumi.gitlab.PipelineScheduleArgs;
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
 *         var example = new PipelineSchedule(&#34;example&#34;, PipelineScheduleArgs.builder()        
 *             .project(&#34;12345&#34;)
 *             .description(&#34;Used to schedule builds&#34;)
 *             .ref(&#34;master&#34;)
 *             .cron(&#34;0 1 * * *&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitLab pipeline schedules can be imported using an id made up of `{project_id}:{pipeline_schedule_id}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/pipelineSchedule:PipelineSchedule test 1:3
 * ```
 * 
 */
@ResourceType(type="gitlab:index/pipelineSchedule:PipelineSchedule")
public class PipelineSchedule extends com.pulumi.resources.CustomResource {
    /**
     * The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
     * 
     */
    @Export(name="active", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> active;

    /**
     * @return The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
     * 
     */
    public Output<Boolean> active() {
        return this.active;
    }
    /**
     * The cron (e.g. `0 1 * * *`).
     * 
     */
    @Export(name="cron", refs={String.class}, tree="[0]")
    private Output<String> cron;

    /**
     * @return The cron (e.g. `0 1 * * *`).
     * 
     */
    public Output<String> cron() {
        return this.cron;
    }
    /**
     * The timezone.
     * 
     */
    @Export(name="cronTimezone", refs={String.class}, tree="[0]")
    private Output<String> cronTimezone;

    /**
     * @return The timezone.
     * 
     */
    public Output<String> cronTimezone() {
        return this.cronTimezone;
    }
    /**
     * The description of the pipeline schedule.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the pipeline schedule.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The ID of the user that owns the pipeline schedule.
     * 
     */
    @Export(name="owner", refs={Integer.class}, tree="[0]")
    private Output<Integer> owner;

    /**
     * @return The ID of the user that owns the pipeline schedule.
     * 
     */
    public Output<Integer> owner() {
        return this.owner;
    }
    /**
     * The pipeline schedule id.
     * 
     */
    @Export(name="pipelineScheduleId", refs={Integer.class}, tree="[0]")
    private Output<Integer> pipelineScheduleId;

    /**
     * @return The pipeline schedule id.
     * 
     */
    public Output<Integer> pipelineScheduleId() {
        return this.pipelineScheduleId;
    }
    /**
     * The name or id of the project to add the schedule to.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The name or id of the project to add the schedule to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The branch/tag name to be triggered.
     * 
     */
    @Export(name="ref", refs={String.class}, tree="[0]")
    private Output<String> ref;

    /**
     * @return The branch/tag name to be triggered.
     * 
     */
    public Output<String> ref() {
        return this.ref;
    }
    @Export(name="takeOwnership", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> takeOwnership;

    public Output<Boolean> takeOwnership() {
        return this.takeOwnership;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PipelineSchedule(String name) {
        this(name, PipelineScheduleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PipelineSchedule(String name, PipelineScheduleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PipelineSchedule(String name, PipelineScheduleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/pipelineSchedule:PipelineSchedule", name, args == null ? PipelineScheduleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PipelineSchedule(String name, Output<String> id, @Nullable PipelineScheduleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/pipelineSchedule:PipelineSchedule", name, state, makeResourceOptions(options, id));
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
    public static PipelineSchedule get(String name, Output<String> id, @Nullable PipelineScheduleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PipelineSchedule(name, id, state, options);
    }
}
