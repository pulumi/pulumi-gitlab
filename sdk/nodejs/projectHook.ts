// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectHook` resource allows to manage the lifecycle of a project hook.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/projects/#hooks)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.ProjectHook("example", {
 *     project: "example/hooked",
 *     url: "https://example.com/hook/example",
 *     name: "example",
 *     description: "Example hook",
 *     mergeRequestsEvents: true,
 * });
 * // Using Custom Headers
 * // Values of headers can't be imported
 * const customHeaders = new gitlab.ProjectHook("custom_headers", {
 *     project: "example/hooked",
 *     url: "https://example.com/hook/example",
 *     mergeRequestsEvents: true,
 *     customHeaders: [
 *         {
 *             key: "X-Custom-Header",
 *             value: "example",
 *         },
 *         {
 *             key: "X-Custom-Header-Second",
 *             value: "example-second",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_hook`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_project_hook.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * A GitLab Project Hook can be imported using a key composed of `<project-id>:<hook-id>`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectHook:ProjectHook example "12345:1"
 * ```
 *
 * NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
 */
export class ProjectHook extends pulumi.CustomResource {
    /**
     * Get an existing ProjectHook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectHookState, opts?: pulumi.CustomResourceOptions): ProjectHook {
        return new ProjectHook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectHook:ProjectHook';

    /**
     * Returns true if the given object is an instance of ProjectHook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectHook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectHook.__pulumiType;
    }

    /**
     * Invoke the hook for confidential issues events.
     */
    public readonly confidentialIssuesEvents!: pulumi.Output<boolean>;
    /**
     * Invoke the hook for confidential note events.
     */
    public readonly confidentialNoteEvents!: pulumi.Output<boolean>;
    /**
     * Custom headers for the project webhook.
     */
    public readonly customHeaders!: pulumi.Output<outputs.ProjectHookCustomHeader[] | undefined>;
    /**
     * Custom webhook template.
     */
    public readonly customWebhookTemplate!: pulumi.Output<string>;
    /**
     * Invoke the hook for deployment events.
     */
    public readonly deploymentEvents!: pulumi.Output<boolean>;
    /**
     * Description of the webhook.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Enable SSL verification when invoking the hook.
     */
    public readonly enableSslVerification!: pulumi.Output<boolean>;
    /**
     * The id of the project hook.
     */
    public /*out*/ readonly hookId!: pulumi.Output<number>;
    /**
     * Invoke the hook for issues events.
     */
    public readonly issuesEvents!: pulumi.Output<boolean>;
    /**
     * Invoke the hook for job events.
     */
    public readonly jobEvents!: pulumi.Output<boolean>;
    /**
     * Invoke the hook for merge requests events.
     */
    public readonly mergeRequestsEvents!: pulumi.Output<boolean>;
    /**
     * Name of the project webhook.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Invoke the hook for note events.
     */
    public readonly noteEvents!: pulumi.Output<boolean>;
    /**
     * Invoke the hook for pipeline events.
     */
    public readonly pipelineEvents!: pulumi.Output<boolean>;
    /**
     * The name or id of the project to add the hook to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The id of the project for the hook.
     */
    public /*out*/ readonly projectId!: pulumi.Output<number>;
    /**
     * Invoke the hook for push events.
     */
    public readonly pushEvents!: pulumi.Output<boolean>;
    /**
     * Invoke the hook for push events on matching branches only.
     */
    public readonly pushEventsBranchFilter!: pulumi.Output<string>;
    /**
     * Invoke the hook for release events.
     */
    public readonly releasesEvents!: pulumi.Output<boolean>;
    /**
     * Invoke the hook for project access token expiry events.
     */
    public readonly resourceAccessTokenEvents!: pulumi.Output<boolean>;
    /**
     * Invoke the hook for tag push events.
     */
    public readonly tagPushEvents!: pulumi.Output<boolean>;
    /**
     * A token to present when invoking the hook. The token is not available for imported resources.
     */
    public readonly token!: pulumi.Output<string>;
    /**
     * The url of the hook to invoke. Forces re-creation to preserve `token`.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Invoke the hook for wiki page events.
     */
    public readonly wikiPageEvents!: pulumi.Output<boolean>;

    /**
     * Create a ProjectHook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectHookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectHookArgs | ProjectHookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectHookState | undefined;
            resourceInputs["confidentialIssuesEvents"] = state ? state.confidentialIssuesEvents : undefined;
            resourceInputs["confidentialNoteEvents"] = state ? state.confidentialNoteEvents : undefined;
            resourceInputs["customHeaders"] = state ? state.customHeaders : undefined;
            resourceInputs["customWebhookTemplate"] = state ? state.customWebhookTemplate : undefined;
            resourceInputs["deploymentEvents"] = state ? state.deploymentEvents : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enableSslVerification"] = state ? state.enableSslVerification : undefined;
            resourceInputs["hookId"] = state ? state.hookId : undefined;
            resourceInputs["issuesEvents"] = state ? state.issuesEvents : undefined;
            resourceInputs["jobEvents"] = state ? state.jobEvents : undefined;
            resourceInputs["mergeRequestsEvents"] = state ? state.mergeRequestsEvents : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["noteEvents"] = state ? state.noteEvents : undefined;
            resourceInputs["pipelineEvents"] = state ? state.pipelineEvents : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["pushEvents"] = state ? state.pushEvents : undefined;
            resourceInputs["pushEventsBranchFilter"] = state ? state.pushEventsBranchFilter : undefined;
            resourceInputs["releasesEvents"] = state ? state.releasesEvents : undefined;
            resourceInputs["resourceAccessTokenEvents"] = state ? state.resourceAccessTokenEvents : undefined;
            resourceInputs["tagPushEvents"] = state ? state.tagPushEvents : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["wikiPageEvents"] = state ? state.wikiPageEvents : undefined;
        } else {
            const args = argsOrState as ProjectHookArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["confidentialIssuesEvents"] = args ? args.confidentialIssuesEvents : undefined;
            resourceInputs["confidentialNoteEvents"] = args ? args.confidentialNoteEvents : undefined;
            resourceInputs["customHeaders"] = args ? args.customHeaders : undefined;
            resourceInputs["customWebhookTemplate"] = args ? args.customWebhookTemplate : undefined;
            resourceInputs["deploymentEvents"] = args ? args.deploymentEvents : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableSslVerification"] = args ? args.enableSslVerification : undefined;
            resourceInputs["issuesEvents"] = args ? args.issuesEvents : undefined;
            resourceInputs["jobEvents"] = args ? args.jobEvents : undefined;
            resourceInputs["mergeRequestsEvents"] = args ? args.mergeRequestsEvents : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["noteEvents"] = args ? args.noteEvents : undefined;
            resourceInputs["pipelineEvents"] = args ? args.pipelineEvents : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pushEvents"] = args ? args.pushEvents : undefined;
            resourceInputs["pushEventsBranchFilter"] = args ? args.pushEventsBranchFilter : undefined;
            resourceInputs["releasesEvents"] = args ? args.releasesEvents : undefined;
            resourceInputs["resourceAccessTokenEvents"] = args ? args.resourceAccessTokenEvents : undefined;
            resourceInputs["tagPushEvents"] = args ? args.tagPushEvents : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["wikiPageEvents"] = args ? args.wikiPageEvents : undefined;
            resourceInputs["hookId"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ProjectHook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectHook resources.
 */
export interface ProjectHookState {
    /**
     * Invoke the hook for confidential issues events.
     */
    confidentialIssuesEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for confidential note events.
     */
    confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * Custom headers for the project webhook.
     */
    customHeaders?: pulumi.Input<pulumi.Input<inputs.ProjectHookCustomHeader>[]>;
    /**
     * Custom webhook template.
     */
    customWebhookTemplate?: pulumi.Input<string>;
    /**
     * Invoke the hook for deployment events.
     */
    deploymentEvents?: pulumi.Input<boolean>;
    /**
     * Description of the webhook.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable SSL verification when invoking the hook.
     */
    enableSslVerification?: pulumi.Input<boolean>;
    /**
     * The id of the project hook.
     */
    hookId?: pulumi.Input<number>;
    /**
     * Invoke the hook for issues events.
     */
    issuesEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for job events.
     */
    jobEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for merge requests events.
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * Name of the project webhook.
     */
    name?: pulumi.Input<string>;
    /**
     * Invoke the hook for note events.
     */
    noteEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for pipeline events.
     */
    pipelineEvents?: pulumi.Input<boolean>;
    /**
     * The name or id of the project to add the hook to.
     */
    project?: pulumi.Input<string>;
    /**
     * The id of the project for the hook.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Invoke the hook for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for push events on matching branches only.
     */
    pushEventsBranchFilter?: pulumi.Input<string>;
    /**
     * Invoke the hook for release events.
     */
    releasesEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for project access token expiry events.
     */
    resourceAccessTokenEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for tag push events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * A token to present when invoking the hook. The token is not available for imported resources.
     */
    token?: pulumi.Input<string>;
    /**
     * The url of the hook to invoke. Forces re-creation to preserve `token`.
     */
    url?: pulumi.Input<string>;
    /**
     * Invoke the hook for wiki page events.
     */
    wikiPageEvents?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ProjectHook resource.
 */
export interface ProjectHookArgs {
    /**
     * Invoke the hook for confidential issues events.
     */
    confidentialIssuesEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for confidential note events.
     */
    confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * Custom headers for the project webhook.
     */
    customHeaders?: pulumi.Input<pulumi.Input<inputs.ProjectHookCustomHeader>[]>;
    /**
     * Custom webhook template.
     */
    customWebhookTemplate?: pulumi.Input<string>;
    /**
     * Invoke the hook for deployment events.
     */
    deploymentEvents?: pulumi.Input<boolean>;
    /**
     * Description of the webhook.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable SSL verification when invoking the hook.
     */
    enableSslVerification?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for issues events.
     */
    issuesEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for job events.
     */
    jobEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for merge requests events.
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * Name of the project webhook.
     */
    name?: pulumi.Input<string>;
    /**
     * Invoke the hook for note events.
     */
    noteEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for pipeline events.
     */
    pipelineEvents?: pulumi.Input<boolean>;
    /**
     * The name or id of the project to add the hook to.
     */
    project: pulumi.Input<string>;
    /**
     * Invoke the hook for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for push events on matching branches only.
     */
    pushEventsBranchFilter?: pulumi.Input<string>;
    /**
     * Invoke the hook for release events.
     */
    releasesEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for project access token expiry events.
     */
    resourceAccessTokenEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for tag push events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * A token to present when invoking the hook. The token is not available for imported resources.
     */
    token?: pulumi.Input<string>;
    /**
     * The url of the hook to invoke. Forces re-creation to preserve `token`.
     */
    url: pulumi.Input<string>;
    /**
     * Invoke the hook for wiki page events.
     */
    wikiPageEvents?: pulumi.Input<boolean>;
}
