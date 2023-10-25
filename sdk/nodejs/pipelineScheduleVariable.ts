// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.PipelineScheduleVariable` resource allows to manage the lifecycle of a variable for a pipeline schedule.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_schedules.html#pipeline-schedule-variables)
 *
 * ## Import
 *
 * Pipeline schedule variables can be imported using an id made up of `project_id:pipeline_schedule_id:key`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable example 123456789:13:mykey
 * ```
 */
export class PipelineScheduleVariable extends pulumi.CustomResource {
    /**
     * Get an existing PipelineScheduleVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineScheduleVariableState, opts?: pulumi.CustomResourceOptions): PipelineScheduleVariable {
        return new PipelineScheduleVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable';

    /**
     * Returns true if the given object is an instance of PipelineScheduleVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PipelineScheduleVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PipelineScheduleVariable.__pulumiType;
    }

    /**
     * Name of the variable.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The id of the pipeline schedule.
     */
    public readonly pipelineScheduleId!: pulumi.Output<number>;
    /**
     * The id of the project to add the schedule to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Value of the variable.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a PipelineScheduleVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineScheduleVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineScheduleVariableArgs | PipelineScheduleVariableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PipelineScheduleVariableState | undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["pipelineScheduleId"] = state ? state.pipelineScheduleId : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as PipelineScheduleVariableArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.pipelineScheduleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pipelineScheduleId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["pipelineScheduleId"] = args ? args.pipelineScheduleId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PipelineScheduleVariable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PipelineScheduleVariable resources.
 */
export interface PipelineScheduleVariableState {
    /**
     * Name of the variable.
     */
    key?: pulumi.Input<string>;
    /**
     * The id of the pipeline schedule.
     */
    pipelineScheduleId?: pulumi.Input<number>;
    /**
     * The id of the project to add the schedule to.
     */
    project?: pulumi.Input<string>;
    /**
     * Value of the variable.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PipelineScheduleVariable resource.
 */
export interface PipelineScheduleVariableArgs {
    /**
     * Name of the variable.
     */
    key: pulumi.Input<string>;
    /**
     * The id of the pipeline schedule.
     */
    pipelineScheduleId: pulumi.Input<number>;
    /**
     * The id of the project to add the schedule to.
     */
    project: pulumi.Input<string>;
    /**
     * Value of the variable.
     */
    value: pulumi.Input<string>;
}
