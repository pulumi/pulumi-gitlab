// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getGroupHooks` data source allows to retrieve details about hooks in a group.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/group_webhooks/#list-group-hooks)
 */
export function getGroupHooks(args: GetGroupHooksArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupHooksResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getGroupHooks:getGroupHooks", {
        "group": args.group,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupHooks.
 */
export interface GetGroupHooksArgs {
    /**
     * The ID or full path of the group.
     */
    group: string;
}

/**
 * A collection of values returned by getGroupHooks.
 */
export interface GetGroupHooksResult {
    /**
     * The ID or full path of the group.
     */
    readonly group: string;
    /**
     * The list of hooks.
     */
    readonly hooks: outputs.GetGroupHooksHook[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
/**
 * The `gitlab.getGroupHooks` data source allows to retrieve details about hooks in a group.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/group_webhooks/#list-group-hooks)
 */
export function getGroupHooksOutput(args: GetGroupHooksOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupHooksResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getGroupHooks:getGroupHooks", {
        "group": args.group,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupHooks.
 */
export interface GetGroupHooksOutputArgs {
    /**
     * The ID or full path of the group.
     */
    group: pulumi.Input<string>;
}
