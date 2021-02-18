// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # gitlab.ProjectFreezePeriod
 *
 * This resource allows you to create and manage freeze periods. For further information on freeze periods, consult the [gitlab documentation](https://docs.gitlab.com/ee/api/freeze_periods.html#create-a-freeze-period).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const schedule = new gitlab.ProjectFreezePeriod("schedule", {
 *     projectId: gitlab_project.foo.id,
 *     freezeStart: "0 23 * * 5",
 *     freezeEnd: "0 7 * * 1",
 *     cronTimezone: "UTC",
 * });
 * ```
 *
 * ## Import
 *
 * GitLab project freeze periods can be imported using an id made up of `project_id:freeze_period_id`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/projectFreezePeriod:ProjectFreezePeriod schedule "12345:1337"
 * ```
 */
export class ProjectFreezePeriod extends pulumi.CustomResource {
    /**
     * Get an existing ProjectFreezePeriod resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectFreezePeriodState, opts?: pulumi.CustomResourceOptions): ProjectFreezePeriod {
        return new ProjectFreezePeriod(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectFreezePeriod:ProjectFreezePeriod';

    /**
     * Returns true if the given object is an instance of ProjectFreezePeriod.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectFreezePeriod {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectFreezePeriod.__pulumiType;
    }

    /**
     * The timezone.
     */
    public readonly cronTimezone!: pulumi.Output<string | undefined>;
    /**
     * End of the Freeze Period in cron format (e.g. `0 2 * * *`).
     */
    public readonly freezeEnd!: pulumi.Output<string>;
    /**
     * Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
     */
    public readonly freezeStart!: pulumi.Output<string>;
    /**
     * The id of the project to add the schedule to.
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a ProjectFreezePeriod resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectFreezePeriodArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectFreezePeriodArgs | ProjectFreezePeriodState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectFreezePeriodState | undefined;
            inputs["cronTimezone"] = state ? state.cronTimezone : undefined;
            inputs["freezeEnd"] = state ? state.freezeEnd : undefined;
            inputs["freezeStart"] = state ? state.freezeStart : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as ProjectFreezePeriodArgs | undefined;
            if ((!args || args.freezeEnd === undefined) && !opts.urn) {
                throw new Error("Missing required property 'freezeEnd'");
            }
            if ((!args || args.freezeStart === undefined) && !opts.urn) {
                throw new Error("Missing required property 'freezeStart'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["cronTimezone"] = args ? args.cronTimezone : undefined;
            inputs["freezeEnd"] = args ? args.freezeEnd : undefined;
            inputs["freezeStart"] = args ? args.freezeStart : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ProjectFreezePeriod.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectFreezePeriod resources.
 */
export interface ProjectFreezePeriodState {
    /**
     * The timezone.
     */
    readonly cronTimezone?: pulumi.Input<string>;
    /**
     * End of the Freeze Period in cron format (e.g. `0 2 * * *`).
     */
    readonly freezeEnd?: pulumi.Input<string>;
    /**
     * Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
     */
    readonly freezeStart?: pulumi.Input<string>;
    /**
     * The id of the project to add the schedule to.
     */
    readonly projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectFreezePeriod resource.
 */
export interface ProjectFreezePeriodArgs {
    /**
     * The timezone.
     */
    readonly cronTimezone?: pulumi.Input<string>;
    /**
     * End of the Freeze Period in cron format (e.g. `0 2 * * *`).
     */
    readonly freezeEnd: pulumi.Input<string>;
    /**
     * Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
     */
    readonly freezeStart: pulumi.Input<string>;
    /**
     * The id of the project to add the schedule to.
     */
    readonly projectId: pulumi.Input<string>;
}
