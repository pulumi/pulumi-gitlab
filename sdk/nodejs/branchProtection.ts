// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to protect a specific branch by an access level so that the user with less access level cannot Merge/Push to the branch. GitLab EE features to protect by group or user are not supported.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const branchProtect = new gitlab.BranchProtection("BranchProtect", {
 *     branch: "BranchProtected",
 *     mergeAccessLevel: "developer",
 *     project: "12345",
 *     pushAccessLevel: "developer",
 * });
 * ```
 */
export class BranchProtection extends pulumi.CustomResource {
    /**
     * Get an existing BranchProtection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BranchProtectionState, opts?: pulumi.CustomResourceOptions): BranchProtection {
        return new BranchProtection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/branchProtection:BranchProtection';

    /**
     * Returns true if the given object is an instance of BranchProtection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BranchProtection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BranchProtection.__pulumiType;
    }

    /**
     * Name of the branch.
     */
    public readonly branch!: pulumi.Output<string>;
    /**
     * One of five levels of access to the project.
     */
    public readonly mergeAccessLevel!: pulumi.Output<string>;
    /**
     * The id of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * One of five levels of access to the project.
     */
    public readonly pushAccessLevel!: pulumi.Output<string>;

    /**
     * Create a BranchProtection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BranchProtectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BranchProtectionArgs | BranchProtectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BranchProtectionState | undefined;
            inputs["branch"] = state ? state.branch : undefined;
            inputs["mergeAccessLevel"] = state ? state.mergeAccessLevel : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pushAccessLevel"] = state ? state.pushAccessLevel : undefined;
        } else {
            const args = argsOrState as BranchProtectionArgs | undefined;
            if (!args || args.branch === undefined) {
                throw new Error("Missing required property 'branch'");
            }
            if (!args || args.mergeAccessLevel === undefined) {
                throw new Error("Missing required property 'mergeAccessLevel'");
            }
            if (!args || args.project === undefined) {
                throw new Error("Missing required property 'project'");
            }
            if (!args || args.pushAccessLevel === undefined) {
                throw new Error("Missing required property 'pushAccessLevel'");
            }
            inputs["branch"] = args ? args.branch : undefined;
            inputs["mergeAccessLevel"] = args ? args.mergeAccessLevel : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pushAccessLevel"] = args ? args.pushAccessLevel : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BranchProtection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BranchProtection resources.
 */
export interface BranchProtectionState {
    /**
     * Name of the branch.
     */
    readonly branch?: pulumi.Input<string>;
    /**
     * One of five levels of access to the project.
     */
    readonly mergeAccessLevel?: pulumi.Input<string>;
    /**
     * The id of the project.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * One of five levels of access to the project.
     */
    readonly pushAccessLevel?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BranchProtection resource.
 */
export interface BranchProtectionArgs {
    /**
     * Name of the branch.
     */
    readonly branch: pulumi.Input<string>;
    /**
     * One of five levels of access to the project.
     */
    readonly mergeAccessLevel: pulumi.Input<string>;
    /**
     * The id of the project.
     */
    readonly project: pulumi.Input<string>;
    /**
     * One of five levels of access to the project.
     */
    readonly pushAccessLevel: pulumi.Input<string>;
}
