// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const branchProtect = new gitlab.BranchProtection("BranchProtect", {
 *     project: "12345",
 *     branch: "BranchProtected",
 *     pushAccessLevel: "developer",
 *     mergeAccessLevel: "developer",
 *     unprotectAccessLevel: "developer",
 *     allowForcePush: true,
 *     codeOwnerApprovalRequired: true,
 *     allowedToPushes: [
 *         {
 *             userId: 5,
 *         },
 *         {
 *             userId: 521,
 *         },
 *     ],
 *     allowedToMerges: [
 *         {
 *             userId: 15,
 *         },
 *         {
 *             userId: 37,
 *         },
 *     ],
 *     allowedToUnprotects: [
 *         {
 *             userId: 15,
 *         },
 *         {
 *             groupId: 42,
 *         },
 *     ],
 * });
 * // Example using dynamic block
 * const main = new gitlab.BranchProtection("main", {
 *     allowedToPushes: [
 *         50,
 *         55,
 *         60,
 *     ].map((v, k) => ({key: k, value: v})).map(entry => ({
 *         userId: entry.value,
 *     })),
 *     project: "12345",
 *     branch: "main",
 *     pushAccessLevel: "maintainer",
 *     mergeAccessLevel: "maintainer",
 *     unprotectAccessLevel: "maintainer",
 * });
 * ```
 *
 * ## Import
 *
 * Gitlab protected branches can be imported with a key composed of `<project_id>:<branch>`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/branchProtection:BranchProtection BranchProtect "12345:main"
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
     * Can be set to true to allow users with push access to force push.
     */
    public readonly allowForcePush!: pulumi.Output<boolean>;
    /**
     * Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     */
    public readonly allowedToMerges!: pulumi.Output<outputs.BranchProtectionAllowedToMerge[] | undefined>;
    /**
     * Array of access levels and user(s)/group(s) allowed to push to protected branch.
     */
    public readonly allowedToPushes!: pulumi.Output<outputs.BranchProtectionAllowedToPush[] | undefined>;
    /**
     * Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
     */
    public readonly allowedToUnprotects!: pulumi.Output<outputs.BranchProtectionAllowedToUnprotect[] | undefined>;
    /**
     * Name of the branch.
     */
    public readonly branch!: pulumi.Output<string>;
    /**
     * The ID of the branch protection (not the branch name).
     */
    public /*out*/ readonly branchProtectionId!: pulumi.Output<number>;
    /**
     * Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
     */
    public readonly codeOwnerApprovalRequired!: pulumi.Output<boolean>;
    /**
     * Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
     */
    public readonly mergeAccessLevel!: pulumi.Output<string>;
    /**
     * The id of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
     */
    public readonly pushAccessLevel!: pulumi.Output<string>;
    /**
     * Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
     */
    public readonly unprotectAccessLevel!: pulumi.Output<string>;

    /**
     * Create a BranchProtection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BranchProtectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BranchProtectionArgs | BranchProtectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BranchProtectionState | undefined;
            resourceInputs["allowForcePush"] = state ? state.allowForcePush : undefined;
            resourceInputs["allowedToMerges"] = state ? state.allowedToMerges : undefined;
            resourceInputs["allowedToPushes"] = state ? state.allowedToPushes : undefined;
            resourceInputs["allowedToUnprotects"] = state ? state.allowedToUnprotects : undefined;
            resourceInputs["branch"] = state ? state.branch : undefined;
            resourceInputs["branchProtectionId"] = state ? state.branchProtectionId : undefined;
            resourceInputs["codeOwnerApprovalRequired"] = state ? state.codeOwnerApprovalRequired : undefined;
            resourceInputs["mergeAccessLevel"] = state ? state.mergeAccessLevel : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pushAccessLevel"] = state ? state.pushAccessLevel : undefined;
            resourceInputs["unprotectAccessLevel"] = state ? state.unprotectAccessLevel : undefined;
        } else {
            const args = argsOrState as BranchProtectionArgs | undefined;
            if ((!args || args.branch === undefined) && !opts.urn) {
                throw new Error("Missing required property 'branch'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["allowForcePush"] = args ? args.allowForcePush : undefined;
            resourceInputs["allowedToMerges"] = args ? args.allowedToMerges : undefined;
            resourceInputs["allowedToPushes"] = args ? args.allowedToPushes : undefined;
            resourceInputs["allowedToUnprotects"] = args ? args.allowedToUnprotects : undefined;
            resourceInputs["branch"] = args ? args.branch : undefined;
            resourceInputs["codeOwnerApprovalRequired"] = args ? args.codeOwnerApprovalRequired : undefined;
            resourceInputs["mergeAccessLevel"] = args ? args.mergeAccessLevel : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pushAccessLevel"] = args ? args.pushAccessLevel : undefined;
            resourceInputs["unprotectAccessLevel"] = args ? args.unprotectAccessLevel : undefined;
            resourceInputs["branchProtectionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BranchProtection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BranchProtection resources.
 */
export interface BranchProtectionState {
    /**
     * Can be set to true to allow users with push access to force push.
     */
    allowForcePush?: pulumi.Input<boolean>;
    /**
     * Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     */
    allowedToMerges?: pulumi.Input<pulumi.Input<inputs.BranchProtectionAllowedToMerge>[]>;
    /**
     * Array of access levels and user(s)/group(s) allowed to push to protected branch.
     */
    allowedToPushes?: pulumi.Input<pulumi.Input<inputs.BranchProtectionAllowedToPush>[]>;
    /**
     * Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
     */
    allowedToUnprotects?: pulumi.Input<pulumi.Input<inputs.BranchProtectionAllowedToUnprotect>[]>;
    /**
     * Name of the branch.
     */
    branch?: pulumi.Input<string>;
    /**
     * The ID of the branch protection (not the branch name).
     */
    branchProtectionId?: pulumi.Input<number>;
    /**
     * Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
     */
    codeOwnerApprovalRequired?: pulumi.Input<boolean>;
    /**
     * Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
     */
    mergeAccessLevel?: pulumi.Input<string>;
    /**
     * The id of the project.
     */
    project?: pulumi.Input<string>;
    /**
     * Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
     */
    pushAccessLevel?: pulumi.Input<string>;
    /**
     * Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
     */
    unprotectAccessLevel?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BranchProtection resource.
 */
export interface BranchProtectionArgs {
    /**
     * Can be set to true to allow users with push access to force push.
     */
    allowForcePush?: pulumi.Input<boolean>;
    /**
     * Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     */
    allowedToMerges?: pulumi.Input<pulumi.Input<inputs.BranchProtectionAllowedToMerge>[]>;
    /**
     * Array of access levels and user(s)/group(s) allowed to push to protected branch.
     */
    allowedToPushes?: pulumi.Input<pulumi.Input<inputs.BranchProtectionAllowedToPush>[]>;
    /**
     * Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
     */
    allowedToUnprotects?: pulumi.Input<pulumi.Input<inputs.BranchProtectionAllowedToUnprotect>[]>;
    /**
     * Name of the branch.
     */
    branch: pulumi.Input<string>;
    /**
     * Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
     */
    codeOwnerApprovalRequired?: pulumi.Input<boolean>;
    /**
     * Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
     */
    mergeAccessLevel?: pulumi.Input<string>;
    /**
     * The id of the project.
     */
    project: pulumi.Input<string>;
    /**
     * Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
     */
    pushAccessLevel?: pulumi.Input<string>;
    /**
     * Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
     */
    unprotectAccessLevel?: pulumi.Input<string>;
}
