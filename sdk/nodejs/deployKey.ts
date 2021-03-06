// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # gitlab\_deploy\_key
 *
 * This resource allows you to create and manage deploy keys for your GitLab projects.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.DeployKey("example", {
 *     key: "ssh-rsa AAAA...",
 *     project: "example/deploying",
 *     title: "Example deploy key",
 * });
 * ```
 *
 * ## Import
 *
 * GitLab deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/deployKey:DeployKey test 1:3
 * ```
 */
export class DeployKey extends pulumi.CustomResource {
    /**
     * Get an existing DeployKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeployKeyState, opts?: pulumi.CustomResourceOptions): DeployKey {
        return new DeployKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/deployKey:DeployKey';

    /**
     * Returns true if the given object is an instance of DeployKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeployKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeployKey.__pulumiType;
    }

    /**
     * Allow this deploy key to be used to push changes to the project.  Defaults to `false`. **NOTE::** this cannot currently be managed.
     */
    public readonly canPush!: pulumi.Output<boolean | undefined>;
    /**
     * The public ssh key body.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The name or id of the project to add the deploy key to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A title to describe the deploy key with.
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a DeployKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeployKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeployKeyArgs | DeployKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeployKeyState | undefined;
            inputs["canPush"] = state ? state.canPush : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as DeployKeyArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            inputs["canPush"] = args ? args.canPush : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["title"] = args ? args.title : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DeployKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeployKey resources.
 */
export interface DeployKeyState {
    /**
     * Allow this deploy key to be used to push changes to the project.  Defaults to `false`. **NOTE::** this cannot currently be managed.
     */
    readonly canPush?: pulumi.Input<boolean>;
    /**
     * The public ssh key body.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the deploy key to.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A title to describe the deploy key with.
     */
    readonly title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeployKey resource.
 */
export interface DeployKeyArgs {
    /**
     * Allow this deploy key to be used to push changes to the project.  Defaults to `false`. **NOTE::** this cannot currently be managed.
     */
    readonly canPush?: pulumi.Input<boolean>;
    /**
     * The public ssh key body.
     */
    readonly key: pulumi.Input<string>;
    /**
     * The name or id of the project to add the deploy key to.
     */
    readonly project: pulumi.Input<string>;
    /**
     * A title to describe the deploy key with.
     */
    readonly title: pulumi.Input<string>;
}
