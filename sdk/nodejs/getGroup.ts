// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.Group` data source allows details of a group to be retrieved by its id or full path.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#details-of-a-group)
 */
export function getGroup(args?: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
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
     * The default branch of the group.
     */
    readonly defaultBranch: string;
    /**
     * Whether developers and maintainers can push to the applicable default branch.
     */
    readonly defaultBranchProtection: number;
    /**
     * The description of the group.
     */
    readonly description: string;
    /**
     * Can be set by administrators only. Additional CI/CD minutes for this group.
     */
    readonly extraSharedRunnersMinutesLimit: number;
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
     * Users cannot be added to projects in this group.
     */
    readonly membershipLock: boolean;
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
     * When enabled, users can not fork projects from this group to external namespaces.
     */
    readonly preventForkingOutsideGroup: boolean;
    /**
     * Boolean, is request for access enabled to the group.
     */
    readonly requestAccessEnabled: boolean;
    /**
     * The group level registration token to use during runner setup.
     */
    readonly runnersToken: string;
    /**
     * Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
     */
    readonly sharedRunnersMinutesLimit: number;
    /**
     * Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabledAndOverridable`, `disabledAndUnoverridable`, `disabledWithOverride`.
     */
    readonly sharedRunnersSetting: string;
    /**
     * Describes groups which have access shared to this group.
     */
    readonly sharedWithGroups: outputs.GetGroupSharedWithGroup[];
    /**
     * Visibility level of the group. Possible values are `private`, `internal`, `public`.
     */
    readonly visibilityLevel: string;
    /**
     * Web URL of the group.
     */
    readonly webUrl: string;
    /**
     * The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly wikiAccessLevel: string;
}
/**
 * The `gitlab.Group` data source allows details of a group to be retrieved by its id or full path.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#details-of-a-group)
 */
export function getGroupOutput(args?: GetGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getGroup:getGroup", {
        "fullPath": args.fullPath,
        "groupId": args.groupId,
    }, opts);
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
