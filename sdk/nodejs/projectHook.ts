// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectHook` resource allows to manage the lifecycle of a project hook.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#hooks)
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
 *     mergeRequestsEvents: true,
 * });
 * ```
 *
 * ## Import
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
    public readonly confidentialIssuesEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for confidential notes events.
     */
    public readonly confidentialNoteEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Set a custom webhook template.
     */
    public readonly customWebhookTemplate!: pulumi.Output<string | undefined>;
    /**
     * Invoke the hook for deployment events.
     */
    public readonly deploymentEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Enable ssl verification when invoking the hook.
     */
    public readonly enableSslVerification!: pulumi.Output<boolean | undefined>;
    /**
     * The id of the project hook.
     */
    public /*out*/ readonly hookId!: pulumi.Output<number>;
    /**
     * Invoke the hook for issues events.
     */
    public readonly issuesEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for job events.
     */
    public readonly jobEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for merge requests.
     */
    public readonly mergeRequestsEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for notes events.
     */
    public readonly noteEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for pipeline events.
     */
    public readonly pipelineEvents!: pulumi.Output<boolean | undefined>;
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
    public readonly pushEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for push events on matching branches only.
     */
    public readonly pushEventsBranchFilter!: pulumi.Output<string | undefined>;
    /**
     * Invoke the hook for releases events.
     */
    public readonly releasesEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for tag push events.
     */
    public readonly tagPushEvents!: pulumi.Output<boolean | undefined>;
    /**
     * A token to present when invoking the hook. The token is not available for imported resources.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * The url of the hook to invoke. Forces re-creation to preserve `token`.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Invoke the hook for wiki page events.
     */
    public readonly wikiPageEvents!: pulumi.Output<boolean | undefined>;

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
            resourceInputs["customWebhookTemplate"] = state ? state.customWebhookTemplate : undefined;
            resourceInputs["deploymentEvents"] = state ? state.deploymentEvents : undefined;
            resourceInputs["enableSslVerification"] = state ? state.enableSslVerification : undefined;
            resourceInputs["hookId"] = state ? state.hookId : undefined;
            resourceInputs["issuesEvents"] = state ? state.issuesEvents : undefined;
            resourceInputs["jobEvents"] = state ? state.jobEvents : undefined;
            resourceInputs["mergeRequestsEvents"] = state ? state.mergeRequestsEvents : undefined;
            resourceInputs["noteEvents"] = state ? state.noteEvents : undefined;
            resourceInputs["pipelineEvents"] = state ? state.pipelineEvents : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["pushEvents"] = state ? state.pushEvents : undefined;
            resourceInputs["pushEventsBranchFilter"] = state ? state.pushEventsBranchFilter : undefined;
            resourceInputs["releasesEvents"] = state ? state.releasesEvents : undefined;
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
            resourceInputs["customWebhookTemplate"] = args ? args.customWebhookTemplate : undefined;
            resourceInputs["deploymentEvents"] = args ? args.deploymentEvents : undefined;
            resourceInputs["enableSslVerification"] = args ? args.enableSslVerification : undefined;
            resourceInputs["issuesEvents"] = args ? args.issuesEvents : undefined;
            resourceInputs["jobEvents"] = args ? args.jobEvents : undefined;
            resourceInputs["mergeRequestsEvents"] = args ? args.mergeRequestsEvents : undefined;
            resourceInputs["noteEvents"] = args ? args.noteEvents : undefined;
            resourceInputs["pipelineEvents"] = args ? args.pipelineEvents : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pushEvents"] = args ? args.pushEvents : undefined;
            resourceInputs["pushEventsBranchFilter"] = args ? args.pushEventsBranchFilter : undefined;
            resourceInputs["releasesEvents"] = args ? args.releasesEvents : undefined;
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
     * Invoke the hook for confidential notes events.
     */
    confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * Set a custom webhook template.
     */
    customWebhookTemplate?: pulumi.Input<string>;
    /**
     * Invoke the hook for deployment events.
     */
    deploymentEvents?: pulumi.Input<boolean>;
    /**
     * Enable ssl verification when invoking the hook.
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
     * Invoke the hook for merge requests.
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for notes events.
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
     * Invoke the hook for releases events.
     */
    releasesEvents?: pulumi.Input<boolean>;
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
     * Invoke the hook for confidential notes events.
     */
    confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * Set a custom webhook template.
     */
    customWebhookTemplate?: pulumi.Input<string>;
    /**
     * Invoke the hook for deployment events.
     */
    deploymentEvents?: pulumi.Input<boolean>;
    /**
     * Enable ssl verification when invoking the hook.
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
     * Invoke the hook for merge requests.
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for notes events.
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
     * Invoke the hook for releases events.
     */
    releasesEvents?: pulumi.Input<boolean>;
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
