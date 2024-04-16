// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.PipelineSchedule` resource allows to manage the lifecycle of a scheduled pipeline.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_schedules.html)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.PipelineSchedule("example", {
 *     project: "12345",
 *     description: "Used to schedule builds",
 *     ref: "master",
 *     cron: "0 1 * * *",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * GitLab pipeline schedules can be imported using an id made up of `{project_id}:{pipeline_schedule_id}`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/pipelineSchedule:PipelineSchedule test 1:3
 * ```
 */
export class PipelineSchedule extends pulumi.CustomResource {
    /**
     * Get an existing PipelineSchedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineScheduleState, opts?: pulumi.CustomResourceOptions): PipelineSchedule {
        return new PipelineSchedule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/pipelineSchedule:PipelineSchedule';

    /**
     * Returns true if the given object is an instance of PipelineSchedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PipelineSchedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PipelineSchedule.__pulumiType;
    }

    /**
     * The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
     */
    public readonly active!: pulumi.Output<boolean>;
    /**
     * The cron (e.g. `0 1 * * *`).
     */
    public readonly cron!: pulumi.Output<string>;
    /**
     * The timezone.
     */
    public readonly cronTimezone!: pulumi.Output<string>;
    /**
     * The description of the pipeline schedule.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The ID of the user that owns the pipeline schedule.
     */
    public /*out*/ readonly owner!: pulumi.Output<number>;
    /**
     * The pipeline schedule id.
     */
    public /*out*/ readonly pipelineScheduleId!: pulumi.Output<number>;
    /**
     * The name or id of the project to add the schedule to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The branch/tag name to be triggered.
     */
    public readonly ref!: pulumi.Output<string>;
    /**
     * When set to `true`, the user represented by the token running Terraform will take ownership of the scheduled pipeline
     * prior to editing it. This can help when managing scheduled pipeline drift when other users are making changes outside
     * Terraform.
     */
    public readonly takeOwnership!: pulumi.Output<boolean>;

    /**
     * Create a PipelineSchedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineScheduleArgs | PipelineScheduleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PipelineScheduleState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["cron"] = state ? state.cron : undefined;
            resourceInputs["cronTimezone"] = state ? state.cronTimezone : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["pipelineScheduleId"] = state ? state.pipelineScheduleId : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["ref"] = state ? state.ref : undefined;
            resourceInputs["takeOwnership"] = state ? state.takeOwnership : undefined;
        } else {
            const args = argsOrState as PipelineScheduleArgs | undefined;
            if ((!args || args.cron === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cron'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.ref === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ref'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["cron"] = args ? args.cron : undefined;
            resourceInputs["cronTimezone"] = args ? args.cronTimezone : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["ref"] = args ? args.ref : undefined;
            resourceInputs["takeOwnership"] = args ? args.takeOwnership : undefined;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["pipelineScheduleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PipelineSchedule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PipelineSchedule resources.
 */
export interface PipelineScheduleState {
    /**
     * The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The cron (e.g. `0 1 * * *`).
     */
    cron?: pulumi.Input<string>;
    /**
     * The timezone.
     */
    cronTimezone?: pulumi.Input<string>;
    /**
     * The description of the pipeline schedule.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the user that owns the pipeline schedule.
     */
    owner?: pulumi.Input<number>;
    /**
     * The pipeline schedule id.
     */
    pipelineScheduleId?: pulumi.Input<number>;
    /**
     * The name or id of the project to add the schedule to.
     */
    project?: pulumi.Input<string>;
    /**
     * The branch/tag name to be triggered.
     */
    ref?: pulumi.Input<string>;
    /**
     * When set to `true`, the user represented by the token running Terraform will take ownership of the scheduled pipeline
     * prior to editing it. This can help when managing scheduled pipeline drift when other users are making changes outside
     * Terraform.
     */
    takeOwnership?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a PipelineSchedule resource.
 */
export interface PipelineScheduleArgs {
    /**
     * The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The cron (e.g. `0 1 * * *`).
     */
    cron: pulumi.Input<string>;
    /**
     * The timezone.
     */
    cronTimezone?: pulumi.Input<string>;
    /**
     * The description of the pipeline schedule.
     */
    description: pulumi.Input<string>;
    /**
     * The name or id of the project to add the schedule to.
     */
    project: pulumi.Input<string>;
    /**
     * The branch/tag name to be triggered.
     */
    ref: pulumi.Input<string>;
    /**
     * When set to `true`, the user represented by the token running Terraform will take ownership of the scheduled pipeline
     * prior to editing it. This can help when managing scheduled pipeline drift when other users are making changes outside
     * Terraform.
     */
    takeOwnership?: pulumi.Input<boolean>;
}
