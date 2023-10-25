// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.DeployKeyEnable` resource allows to enable an already existing deploy key (see `gitlab.DeployKey resource`) for a specific project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_keys.html#enable-a-deploy-key)
 *
 * ## Import
 *
 * GitLab enabled deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g. `project_id` can be whatever the [get single project api][get_single_project] takes for its `:id` value, so for example
 *
 * ```sh
 *  $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example 12345:67890
 * ```
 *
 * ```sh
 *  $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example richardc/example:67890
 * ```
 */
export class DeployKeyEnable extends pulumi.CustomResource {
    /**
     * Get an existing DeployKeyEnable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeployKeyEnableState, opts?: pulumi.CustomResourceOptions): DeployKeyEnable {
        return new DeployKeyEnable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/deployKeyEnable:DeployKeyEnable';

    /**
     * Returns true if the given object is an instance of DeployKeyEnable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeployKeyEnable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeployKeyEnable.__pulumiType;
    }

    /**
     * Can deploy key push to the project's repository.
     */
    public readonly canPush!: pulumi.Output<boolean | undefined>;
    /**
     * Deploy key.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The Gitlab key id for the pre-existing deploy key
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * The name or id of the project to add the deploy key to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Deploy key's title.
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a DeployKeyEnable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeployKeyEnableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeployKeyEnableArgs | DeployKeyEnableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeployKeyEnableState | undefined;
            resourceInputs["canPush"] = state ? state.canPush : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as DeployKeyEnableArgs | undefined;
            if ((!args || args.keyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["canPush"] = args ? args.canPush : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeployKeyEnable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeployKeyEnable resources.
 */
export interface DeployKeyEnableState {
    /**
     * Can deploy key push to the project's repository.
     */
    canPush?: pulumi.Input<boolean>;
    /**
     * Deploy key.
     */
    key?: pulumi.Input<string>;
    /**
     * The Gitlab key id for the pre-existing deploy key
     */
    keyId?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the deploy key to.
     */
    project?: pulumi.Input<string>;
    /**
     * Deploy key's title.
     */
    title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeployKeyEnable resource.
 */
export interface DeployKeyEnableArgs {
    /**
     * Can deploy key push to the project's repository.
     */
    canPush?: pulumi.Input<boolean>;
    /**
     * Deploy key.
     */
    key?: pulumi.Input<string>;
    /**
     * The Gitlab key id for the pre-existing deploy key
     */
    keyId: pulumi.Input<string>;
    /**
     * The name or id of the project to add the deploy key to.
     */
    project: pulumi.Input<string>;
    /**
     * Deploy key's title.
     */
    title?: pulumi.Input<string>;
}
