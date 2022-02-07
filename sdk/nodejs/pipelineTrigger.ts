// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage pipeline triggers
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.PipelineTrigger("example", {
 *     description: "Used to trigger builds",
 *     project: "12345",
 * });
 * ```
 *
 * ## Import
 *
 * # GitLab pipeline triggers can be imported using an id made up of `{project_id}:{pipeline_trigger_id}`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/pipelineTrigger:PipelineTrigger test 1:3
 * ```
 */
export class PipelineTrigger extends pulumi.CustomResource {
    /**
     * Get an existing PipelineTrigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineTriggerState, opts?: pulumi.CustomResourceOptions): PipelineTrigger {
        return new PipelineTrigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/pipelineTrigger:PipelineTrigger';

    /**
     * Returns true if the given object is an instance of PipelineTrigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PipelineTrigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PipelineTrigger.__pulumiType;
    }

    /**
     * The description of the pipeline trigger.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The name or id of the project to add the trigger to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The pipeline trigger token.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;

    /**
     * Create a PipelineTrigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineTriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineTriggerArgs | PipelineTriggerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PipelineTriggerState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
        } else {
            const args = argsOrState as PipelineTriggerArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["token"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PipelineTrigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PipelineTrigger resources.
 */
export interface PipelineTriggerState {
    /**
     * The description of the pipeline trigger.
     */
    description?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the trigger to.
     */
    project?: pulumi.Input<string>;
    /**
     * The pipeline trigger token.
     */
    token?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PipelineTrigger resource.
 */
export interface PipelineTriggerArgs {
    /**
     * The description of the pipeline trigger.
     */
    description: pulumi.Input<string>;
    /**
     * The name or id of the project to add the trigger to.
     */
    project: pulumi.Input<string>;
}
