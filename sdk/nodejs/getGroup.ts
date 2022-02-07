// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provide details about a specific group in the gitlab provider.
 *
 * > **Note**: exactly one of groupId or fullPath must be provided.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * // By group's full path
 * const foo = pulumi.output(gitlab.getGroup({
 *     fullPath: "foo/bar",
 * }));
 * ```
 */
export function getGroup(args?: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gitlab:index/getGroup:getGroup", {
        "fullPath": args.fullPath,
        "groupId": args.groupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * The full path of the group.
     */
    fullPath?: string;
    /**
     * The ID of the group.
     */
    groupId?: number;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * Whether developers and maintainers can push to the applicable default branch.
     */
    readonly defaultBranchProtection: number;
    /**
     * The description of the group.
     */
    readonly description: string;
    /**
     * The full name of the group.
     */
    readonly fullName: string;
    /**
     * The full path of the group.
     */
    readonly fullPath: string;
    /**
     * The ID of the group.
     */
    readonly groupId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Boolean, is LFS enabled for projects in this group.
     */
    readonly lfsEnabled: boolean;
    /**
     * The name of this group.
     */
    readonly name: string;
    /**
     * Integer, ID of the parent group.
     */
    readonly parentId: number;
    /**
     * The path of the group.
     */
    readonly path: string;
    /**
     * Boolean, is request for access enabled to the group.
     */
    readonly requestAccessEnabled: boolean;
    /**
     * The group level registration token to use during runner setup.
     */
    readonly runnersToken: string;
    /**
     * Visibility level of the group. Possible values are `private`, `internal`, `public`.
     */
    readonly visibilityLevel: string;
    /**
     * Web URL of the group.
     */
    readonly webUrl: string;
}

export function getGroupOutput(args?: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    return pulumi.output(args).apply(a => getGroup(a, opts))
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupOutputArgs {
    /**
     * The full path of the group.
     */
    fullPath?: pulumi.Input<string>;
    /**
     * The ID of the group.
     */
    groupId?: pulumi.Input<number>;
}
