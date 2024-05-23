// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.PipelineScheduleVariableArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.PipelineScheduleVariableState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.PipelineScheduleVariable` resource allows to manage the lifecycle of a variable for a pipeline schedule.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_schedules.html#pipeline-schedule-variables)
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
 * import com.pulumi.gitlab.PipelineSchedule;
 * import com.pulumi.gitlab.PipelineScheduleArgs;
 * import com.pulumi.gitlab.PipelineScheduleVariable;
 * import com.pulumi.gitlab.PipelineScheduleVariableArgs;
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
 *         var example = new PipelineSchedule("example", PipelineScheduleArgs.builder()
 *             .project("12345")
 *             .description("Used to schedule builds")
 *             .ref("master")
 *             .cron("0 1 * * *")
 *             .build());
 * 
 *         var examplePipelineScheduleVariable = new PipelineScheduleVariable("examplePipelineScheduleVariable", PipelineScheduleVariableArgs.builder()
 *             .project(example.project())
 *             .pipelineScheduleId(example.pipelineScheduleId())
 *             .key("EXAMPLE_KEY")
 *             .value("example")
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
 * Pipeline schedule variables can be imported using an id made up of `project_id:pipeline_schedule_id:key`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable example 123456789:13:mykey
 * ```
 * 
 */
@ResourceType(type="gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable")
public class PipelineScheduleVariable extends com.pulumi.resources.CustomResource {
    /**
     * Name of the variable.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return Name of the variable.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The id of the pipeline schedule.
     * 
     */
    @Export(name="pipelineScheduleId", refs={Integer.class}, tree="[0]")
    private Output<Integer> pipelineScheduleId;

    /**
     * @return The id of the pipeline schedule.
     * 
     */
    public Output<Integer> pipelineScheduleId() {
        return this.pipelineScheduleId;
    }
    /**
     * The id of the project to add the schedule to.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The id of the project to add the schedule to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Value of the variable.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return Value of the variable.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PipelineScheduleVariable(String name) {
        this(name, PipelineScheduleVariableArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PipelineScheduleVariable(String name, PipelineScheduleVariableArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PipelineScheduleVariable(String name, PipelineScheduleVariableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable", name, args == null ? PipelineScheduleVariableArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PipelineScheduleVariable(String name, Output<String> id, @Nullable PipelineScheduleVariableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable", name, state, makeResourceOptions(options, id));
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
    public static PipelineScheduleVariable get(String name, Output<String> id, @Nullable PipelineScheduleVariableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PipelineScheduleVariable(name, id, state, options);
    }
}
