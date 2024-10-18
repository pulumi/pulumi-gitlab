// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.GroupHook` resource allows to manage the lifecycle of a group hook.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#hooks)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.GroupHook("example", {
 *     group: "example/hooked",
 *     url: "https://example.com/hook/example",
 *     mergeRequestsEvents: true,
 * });
 * // Setting all attributes
 * const allAttributes = new gitlab.GroupHook("all_attributes", {
 *     group: "1",
 *     url: "http://example.com",
 *     token: "supersecret",
 *     enableSslVerification: false,
 *     pushEvents: true,
 *     pushEventsBranchFilter: "devel",
 *     issuesEvents: false,
 *     confidentialIssuesEvents: false,
 *     mergeRequestsEvents: true,
 *     tagPushEvents: true,
 *     noteEvents: true,
 *     confidentialNoteEvents: true,
 *     jobEvents: true,
 *     pipelineEvents: true,
 *     wikiPageEvents: true,
 *     deploymentEvents: true,
 *     releasesEvents: true,
 *     subgroupEvents: true,
 * });
 * ```
 *
 * ## Import
 *
 * A GitLab Group Hook can be imported using a key composed of `<group-id>:<hook-id>`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/groupHook:GroupHook example "12345:1"
 * ```
 *
 * NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
 */
export class GroupHook extends pulumi.CustomResource {
    /**
     * Get an existing GroupHook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupHookState, opts?: pulumi.CustomResourceOptions): GroupHook {
        return new GroupHook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupHook:GroupHook';

    /**
     * Returns true if the given object is an instance of GroupHook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupHook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupHook.__pulumiType;
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
     * Custom webhook template.
     */
    public readonly customWebhookTemplate!: pulumi.Output<string>;
    /**
     * Invoke the hook for deployment events.
     */
    public readonly deploymentEvents!: pulumi.Output<boolean>;
    /**
     * Enable SSL verification when invoking the hook.
     */
    public readonly enableSslVerification!: pulumi.Output<boolean>;
    /**
     * The full path or id of the group to add the hook to.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * The id of the group for the hook.
     */
    public /*out*/ readonly groupId!: pulumi.Output<number>;
    /**
     * The id of the group hook.
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
     * Invoke the hook for note events.
     */
    public readonly noteEvents!: pulumi.Output<boolean>;
    /**
     * Invoke the hook for pipeline events.
     */
    public readonly pipelineEvents!: pulumi.Output<boolean>;
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
     * Invoke the hook for subgroup events.
     */
    public readonly subgroupEvents!: pulumi.Output<boolean>;
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
     * Create a GroupHook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupHookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupHookArgs | GroupHookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupHookState | undefined;
            resourceInputs["confidentialIssuesEvents"] = state ? state.confidentialIssuesEvents : undefined;
            resourceInputs["confidentialNoteEvents"] = state ? state.confidentialNoteEvents : undefined;
            resourceInputs["customWebhookTemplate"] = state ? state.customWebhookTemplate : undefined;
            resourceInputs["deploymentEvents"] = state ? state.deploymentEvents : undefined;
            resourceInputs["enableSslVerification"] = state ? state.enableSslVerification : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["hookId"] = state ? state.hookId : undefined;
            resourceInputs["issuesEvents"] = state ? state.issuesEvents : undefined;
            resourceInputs["jobEvents"] = state ? state.jobEvents : undefined;
            resourceInputs["mergeRequestsEvents"] = state ? state.mergeRequestsEvents : undefined;
            resourceInputs["noteEvents"] = state ? state.noteEvents : undefined;
            resourceInputs["pipelineEvents"] = state ? state.pipelineEvents : undefined;
            resourceInputs["pushEvents"] = state ? state.pushEvents : undefined;
            resourceInputs["pushEventsBranchFilter"] = state ? state.pushEventsBranchFilter : undefined;
            resourceInputs["releasesEvents"] = state ? state.releasesEvents : undefined;
            resourceInputs["subgroupEvents"] = state ? state.subgroupEvents : undefined;
            resourceInputs["tagPushEvents"] = state ? state.tagPushEvents : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["wikiPageEvents"] = state ? state.wikiPageEvents : undefined;
        } else {
            const args = argsOrState as GroupHookArgs | undefined;
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["confidentialIssuesEvents"] = args ? args.confidentialIssuesEvents : undefined;
            resourceInputs["confidentialNoteEvents"] = args ? args.confidentialNoteEvents : undefined;
            resourceInputs["customWebhookTemplate"] = args ? args.customWebhookTemplate : undefined;
            resourceInputs["deploymentEvents"] = args ? args.deploymentEvents : undefined;
            resourceInputs["enableSslVerification"] = args ? args.enableSslVerification : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["issuesEvents"] = args ? args.issuesEvents : undefined;
            resourceInputs["jobEvents"] = args ? args.jobEvents : undefined;
            resourceInputs["mergeRequestsEvents"] = args ? args.mergeRequestsEvents : undefined;
            resourceInputs["noteEvents"] = args ? args.noteEvents : undefined;
            resourceInputs["pipelineEvents"] = args ? args.pipelineEvents : undefined;
            resourceInputs["pushEvents"] = args ? args.pushEvents : undefined;
            resourceInputs["pushEventsBranchFilter"] = args ? args.pushEventsBranchFilter : undefined;
            resourceInputs["releasesEvents"] = args ? args.releasesEvents : undefined;
            resourceInputs["subgroupEvents"] = args ? args.subgroupEvents : undefined;
            resourceInputs["tagPushEvents"] = args ? args.tagPushEvents : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["wikiPageEvents"] = args ? args.wikiPageEvents : undefined;
            resourceInputs["groupId"] = undefined /*out*/;
            resourceInputs["hookId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(GroupHook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupHook resources.
 */
export interface GroupHookState {
    /**
     * Invoke the hook for confidential issues events.
     */
    confidentialIssuesEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for confidential note events.
     */
    confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * Custom webhook template.
     */
    customWebhookTemplate?: pulumi.Input<string>;
    /**
     * Invoke the hook for deployment events.
     */
    deploymentEvents?: pulumi.Input<boolean>;
    /**
     * Enable SSL verification when invoking the hook.
     */
    enableSslVerification?: pulumi.Input<boolean>;
    /**
     * The full path or id of the group to add the hook to.
     */
    group?: pulumi.Input<string>;
    /**
     * The id of the group for the hook.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The id of the group hook.
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
     * Invoke the hook for note events.
     */
    noteEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for pipeline events.
     */
    pipelineEvents?: pulumi.Input<boolean>;
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
     * Invoke the hook for subgroup events.
     */
    subgroupEvents?: pulumi.Input<boolean>;
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
 * The set of arguments for constructing a GroupHook resource.
 */
export interface GroupHookArgs {
    /**
     * Invoke the hook for confidential issues events.
     */
    confidentialIssuesEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for confidential note events.
     */
    confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * Custom webhook template.
     */
    customWebhookTemplate?: pulumi.Input<string>;
    /**
     * Invoke the hook for deployment events.
     */
    deploymentEvents?: pulumi.Input<boolean>;
    /**
     * Enable SSL verification when invoking the hook.
     */
    enableSslVerification?: pulumi.Input<boolean>;
    /**
     * The full path or id of the group to add the hook to.
     */
    group: pulumi.Input<string>;
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
     * Invoke the hook for note events.
     */
    noteEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for pipeline events.
     */
    pipelineEvents?: pulumi.Input<boolean>;
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
     * Invoke the hook for subgroup events.
     */
    subgroupEvents?: pulumi.Input<boolean>;
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
