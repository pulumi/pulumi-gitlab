// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectHook` data source allows to retrieve details about a hook in a project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#get-project-hook)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getProject({
 *     id: "foo/bar/baz",
 * });
 * const exampleGetProjectHook = example.then(example => gitlab.getProjectHook({
 *     project: example.id,
 *     hookId: 1,
 * }));
 * ```
 */
export function getProjectHook(args: GetProjectHookArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectHookResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProjectHook:getProjectHook", {
        "hookId": args.hookId,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectHook.
 */
export interface GetProjectHookArgs {
    /**
     * The id of the project hook.
     */
    hookId: number;
    /**
     * The name or id of the project to add the hook to.
     */
    project: string;
}

/**
 * A collection of values returned by getProjectHook.
 */
export interface GetProjectHookResult {
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
     * Enable ssl verification when invoking the hook.
     */
    readonly enableSslVerification: boolean;
    /**
     * The id of the project hook.
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
     * The name or id of the project to add the hook to.
     */
    readonly project: string;
    /**
     * The id of the project for the hook.
     */
    readonly projectId: number;
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
 * The `gitlab.ProjectHook` data source allows to retrieve details about a hook in a project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#get-project-hook)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getProject({
 *     id: "foo/bar/baz",
 * });
 * const exampleGetProjectHook = example.then(example => gitlab.getProjectHook({
 *     project: example.id,
 *     hookId: 1,
 * }));
 * ```
 */
export function getProjectHookOutput(args: GetProjectHookOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectHookResult> {
    return pulumi.output(args).apply((a: any) => getProjectHook(a, opts))
}

/**
 * A collection of arguments for invoking getProjectHook.
 */
export interface GetProjectHookOutputArgs {
    /**
     * The id of the project hook.
     */
    hookId: pulumi.Input<number>;
    /**
     * The name or id of the project to add the hook to.
     */
    project: pulumi.Input<string>;
}
