// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.GroupHook` data source allows to retrieve details about a hook in a group.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/group_webhooks/#get-a-group-hook)
 */
export function getGroupHook(args: GetGroupHookArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupHookResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getGroupHook:getGroupHook", {
        "group": args.group,
        "hookId": args.hookId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupHook.
 */
export interface GetGroupHookArgs {
    /**
     * The ID or full path of the group.
     */
    group: string;
    /**
     * The id of the group hook.
     */
    hookId: number;
}

/**
 * A collection of values returned by getGroupHook.
 */
export interface GetGroupHookResult {
    /**
     * Invoke the hook for confidential issues events.
     */
    readonly confidentialIssuesEvents: boolean;
    /**
     * Invoke the hook for confidential notes events.
     */
    readonly confidentialNoteEvents: boolean;
    /**
     * Set a custom webhook template.
     */
    readonly customWebhookTemplate: string;
    /**
     * Invoke the hook for deployment events.
     */
    readonly deploymentEvents: boolean;
    /**
     * Invoke the hook for emoji events.
     */
    readonly emojiEvents: boolean;
    /**
     * Enable ssl verification when invoking the hook.
     */
    readonly enableSslVerification: boolean;
    /**
     * The ID or full path of the group.
     */
    readonly group: string;
    /**
     * The id of the group for the hook.
     */
    readonly groupId: number;
    /**
     * The id of the group hook.
     */
    readonly hookId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Invoke the hook for issues events.
     */
    readonly issuesEvents: boolean;
    /**
     * Invoke the hook for job events.
     */
    readonly jobEvents: boolean;
    /**
     * Invoke the hook for merge requests.
     */
    readonly mergeRequestsEvents: boolean;
    /**
     * Invoke the hook for notes events.
     */
    readonly noteEvents: boolean;
    /**
     * Invoke the hook for pipeline events.
     */
    readonly pipelineEvents: boolean;
    /**
     * Invoke the hook for push events.
     */
    readonly pushEvents: boolean;
    /**
     * Invoke the hook for push events on matching branches only.
     */
    readonly pushEventsBranchFilter: string;
    /**
     * Invoke the hook for releases events.
     */
    readonly releasesEvents: boolean;
    /**
     * Invoke the hook for subgroup events.
     */
    readonly subgroupEvents: boolean;
    /**
     * Invoke the hook for tag push events.
     */
    readonly tagPushEvents: boolean;
    /**
     * A token to present when invoking the hook. The token is not available for imported resources.
     */
    readonly token: string;
    /**
     * The url of the hook to invoke.
     */
    readonly url: string;
    /**
     * Invoke the hook for wiki page events.
     */
    readonly wikiPageEvents: boolean;
}
/**
 * The `gitlab.GroupHook` data source allows to retrieve details about a hook in a group.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/group_webhooks/#get-a-group-hook)
 */
export function getGroupHookOutput(args: GetGroupHookOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupHookResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getGroupHook:getGroupHook", {
        "group": args.group,
        "hookId": args.hookId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupHook.
 */
export interface GetGroupHookOutputArgs {
    /**
     * The ID or full path of the group.
     */
    group: pulumi.Input<string>;
    /**
     * The id of the group hook.
     */
    hookId: pulumi.Input<number>;
}
