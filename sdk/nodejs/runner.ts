// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.Runner` resource allows to manage the lifecycle of a runner.
 *
 * A runner can either be registered at an instance level or group level.
 * The runner will be registered at a group level if the token used is from a group, or at an instance level if the token used is for the instance.
 *
 * ~ > Using this resource will register a runner using the deprecated `registrationToken` flow. To use the new `authenticationToken` flow instead,
 * use the `gitlab.UserRunner` resource!
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/runners.html#register-a-new-runner)
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_runner`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_runner.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * A GitLab Runner can be imported using the runner's ID, eg
 *
 * ```sh
 * $ pulumi import gitlab:index/runner:Runner this 1
 * ```
 */
export class Runner extends pulumi.CustomResource {
    /**
     * Get an existing Runner resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RunnerState, opts?: pulumi.CustomResourceOptions): Runner {
        return new Runner(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/runner:Runner';

    /**
     * Returns true if the given object is an instance of Runner.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Runner {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Runner.__pulumiType;
    }

    /**
     * The accessLevel of the runner. Valid values are: `notProtected`, `refProtected`.
     */
    public readonly accessLevel!: pulumi.Output<string>;
    /**
     * The authentication token used for building a config.toml file. This value is not present when imported.
     */
    public /*out*/ readonly authenticationToken!: pulumi.Output<string>;
    /**
     * The runner's description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the runner should be locked for current project.
     */
    public readonly locked!: pulumi.Output<boolean>;
    /**
     * Free-form maintenance notes for the runner (1024 characters).
     */
    public readonly maintenanceNote!: pulumi.Output<string | undefined>;
    /**
     * Maximum timeout set when this runner handles the job.
     */
    public readonly maximumTimeout!: pulumi.Output<number | undefined>;
    /**
     * Whether the runner should ignore new jobs.
     */
    public readonly paused!: pulumi.Output<boolean>;
    /**
     * The registration token used to register the runner.
     */
    public readonly registrationToken!: pulumi.Output<string>;
    /**
     * Whether the runner should handle untagged jobs.
     */
    public readonly runUntagged!: pulumi.Output<boolean>;
    /**
     * The status of runners to show, one of: online and offline. active and paused are also possible values
     * 			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * List of runner’s tags.
     */
    public readonly tagLists!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Runner resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RunnerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RunnerArgs | RunnerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RunnerState | undefined;
            resourceInputs["accessLevel"] = state ? state.accessLevel : undefined;
            resourceInputs["authenticationToken"] = state ? state.authenticationToken : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["locked"] = state ? state.locked : undefined;
            resourceInputs["maintenanceNote"] = state ? state.maintenanceNote : undefined;
            resourceInputs["maximumTimeout"] = state ? state.maximumTimeout : undefined;
            resourceInputs["paused"] = state ? state.paused : undefined;
            resourceInputs["registrationToken"] = state ? state.registrationToken : undefined;
            resourceInputs["runUntagged"] = state ? state.runUntagged : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tagLists"] = state ? state.tagLists : undefined;
        } else {
            const args = argsOrState as RunnerArgs | undefined;
            if ((!args || args.registrationToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registrationToken'");
            }
            resourceInputs["accessLevel"] = args ? args.accessLevel : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["locked"] = args ? args.locked : undefined;
            resourceInputs["maintenanceNote"] = args ? args.maintenanceNote : undefined;
            resourceInputs["maximumTimeout"] = args ? args.maximumTimeout : undefined;
            resourceInputs["paused"] = args ? args.paused : undefined;
            resourceInputs["registrationToken"] = args?.registrationToken ? pulumi.secret(args.registrationToken) : undefined;
            resourceInputs["runUntagged"] = args ? args.runUntagged : undefined;
            resourceInputs["tagLists"] = args ? args.tagLists : undefined;
            resourceInputs["authenticationToken"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["authenticationToken", "registrationToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Runner.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Runner resources.
 */
export interface RunnerState {
    /**
     * The accessLevel of the runner. Valid values are: `notProtected`, `refProtected`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * The authentication token used for building a config.toml file. This value is not present when imported.
     */
    authenticationToken?: pulumi.Input<string>;
    /**
     * The runner's description.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the runner should be locked for current project.
     */
    locked?: pulumi.Input<boolean>;
    /**
     * Free-form maintenance notes for the runner (1024 characters).
     */
    maintenanceNote?: pulumi.Input<string>;
    /**
     * Maximum timeout set when this runner handles the job.
     */
    maximumTimeout?: pulumi.Input<number>;
    /**
     * Whether the runner should ignore new jobs.
     */
    paused?: pulumi.Input<boolean>;
    /**
     * The registration token used to register the runner.
     */
    registrationToken?: pulumi.Input<string>;
    /**
     * Whether the runner should handle untagged jobs.
     */
    runUntagged?: pulumi.Input<boolean>;
    /**
     * The status of runners to show, one of: online and offline. active and paused are also possible values
     * 			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
     */
    status?: pulumi.Input<string>;
    /**
     * List of runner’s tags.
     */
    tagLists?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Runner resource.
 */
export interface RunnerArgs {
    /**
     * The accessLevel of the runner. Valid values are: `notProtected`, `refProtected`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * The runner's description.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the runner should be locked for current project.
     */
    locked?: pulumi.Input<boolean>;
    /**
     * Free-form maintenance notes for the runner (1024 characters).
     */
    maintenanceNote?: pulumi.Input<string>;
    /**
     * Maximum timeout set when this runner handles the job.
     */
    maximumTimeout?: pulumi.Input<number>;
    /**
     * Whether the runner should ignore new jobs.
     */
    paused?: pulumi.Input<boolean>;
    /**
     * The registration token used to register the runner.
     */
    registrationToken: pulumi.Input<string>;
    /**
     * Whether the runner should handle untagged jobs.
     */
    runUntagged?: pulumi.Input<boolean>;
    /**
     * List of runner’s tags.
     */
    tagLists?: pulumi.Input<pulumi.Input<string>[]>;
}
