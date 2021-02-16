// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # gitlab\_deploy\_key\_enable
 *
 * This resource allows you to enable pre-existing deploy keys for your GitLab projects.
 *
 * **the GITLAB KEY_ID for the deploy key must be known**
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * // A repo to host the deployment key
 * const parentProject = new gitlab.Project("parent", {});
 * // A second repo to use the deployment key from the parent project
 * const fooProject = new gitlab.Project("foo", {});
 * // Upload a deployment key for the parent repo
 * const parentDeployKey = new gitlab.DeployKey("parent", {
 *     key: "ssh-rsa AAAA...",
 *     project: parentProject.id,
 *     title: "Example deploy key",
 * });
 * // Enable the deployment key on the second repo
 * const fooDeployKeyEnable = new gitlab.DeployKeyEnable("foo", {
 *     keyId: parentDeployKey.id,
 *     project: fooProject.id,
 * });
 * ```
 *
 * ## Import
 *
 * GitLab enabled deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example 12345:67890
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

    public readonly canPush!: pulumi.Output<boolean>;
    public readonly key!: pulumi.Output<string>;
    /**
     * The Gitlab key id for the pre-existing deploy key
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * The name or id of the project to add the deploy key to.
     */
    public readonly project!: pulumi.Output<string>;
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeployKeyEnableState | undefined;
            inputs["canPush"] = state ? state.canPush : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["keyId"] = state ? state.keyId : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as DeployKeyEnableArgs | undefined;
            if ((!args || args.keyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["canPush"] = args ? args.canPush : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["keyId"] = args ? args.keyId : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["title"] = args ? args.title : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DeployKeyEnable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeployKeyEnable resources.
 */
export interface DeployKeyEnableState {
    readonly canPush?: pulumi.Input<boolean>;
    readonly key?: pulumi.Input<string>;
    /**
     * The Gitlab key id for the pre-existing deploy key
     */
    readonly keyId?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the deploy key to.
     */
    readonly project?: pulumi.Input<string>;
    readonly title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeployKeyEnable resource.
 */
export interface DeployKeyEnableArgs {
    readonly canPush?: pulumi.Input<boolean>;
    readonly key?: pulumi.Input<string>;
    /**
     * The Gitlab key id for the pre-existing deploy key
     */
    readonly keyId: pulumi.Input<string>;
    /**
     * The name or id of the project to add the deploy key to.
     */
    readonly project: pulumi.Input<string>;
    readonly title?: pulumi.Input<string>;
}
