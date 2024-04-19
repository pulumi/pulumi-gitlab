// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.SystemHook` resource allows to manage the lifecycle of a system hook.
 *
 * > This resource requires GitLab 14.9
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/system_hooks.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.SystemHook("example", {
 *     enableSslVerification: true,
 *     mergeRequestsEvents: true,
 *     pushEvents: true,
 *     repositoryUpdateEvents: true,
 *     tagPushEvents: true,
 *     token: "secret-token",
 *     url: "https://example.com/hook-%d",
 * });
 * ```
 *
 * ## Import
 *
 * You can import a system hook using the hook id `{hook-id}`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/systemHook:SystemHook example 42
 * ```
 *
 * NOTE: the `token` attribute won't be available for imported resources.
 */
export class SystemHook extends pulumi.CustomResource {
    /**
     * Get an existing SystemHook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemHookState, opts?: pulumi.CustomResourceOptions): SystemHook {
        return new SystemHook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/systemHook:SystemHook';

    /**
     * Returns true if the given object is an instance of SystemHook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemHook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemHook.__pulumiType;
    }

    /**
     * The date and time the hook was created in ISO8601 format.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Do SSL verification when triggering the hook.
     */
    public readonly enableSslVerification!: pulumi.Output<boolean | undefined>;
    /**
     * Trigger hook on merge requests events.
     */
    public readonly mergeRequestsEvents!: pulumi.Output<boolean | undefined>;
    /**
     * When true, the hook fires on push events.
     */
    public readonly pushEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Trigger hook on repository update events.
     */
    public readonly repositoryUpdateEvents!: pulumi.Output<boolean | undefined>;
    /**
     * When true, the hook fires on new tags being pushed.
     */
    public readonly tagPushEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * The hook URL.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a SystemHook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemHookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemHookArgs | SystemHookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemHookState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["enableSslVerification"] = state ? state.enableSslVerification : undefined;
            resourceInputs["mergeRequestsEvents"] = state ? state.mergeRequestsEvents : undefined;
            resourceInputs["pushEvents"] = state ? state.pushEvents : undefined;
            resourceInputs["repositoryUpdateEvents"] = state ? state.repositoryUpdateEvents : undefined;
            resourceInputs["tagPushEvents"] = state ? state.tagPushEvents : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as SystemHookArgs | undefined;
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["enableSslVerification"] = args ? args.enableSslVerification : undefined;
            resourceInputs["mergeRequestsEvents"] = args ? args.mergeRequestsEvents : undefined;
            resourceInputs["pushEvents"] = args ? args.pushEvents : undefined;
            resourceInputs["repositoryUpdateEvents"] = args ? args.repositoryUpdateEvents : undefined;
            resourceInputs["tagPushEvents"] = args ? args.tagPushEvents : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SystemHook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemHook resources.
 */
export interface SystemHookState {
    /**
     * The date and time the hook was created in ISO8601 format.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Do SSL verification when triggering the hook.
     */
    enableSslVerification?: pulumi.Input<boolean>;
    /**
     * Trigger hook on merge requests events.
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * When true, the hook fires on push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Trigger hook on repository update events.
     */
    repositoryUpdateEvents?: pulumi.Input<boolean>;
    /**
     * When true, the hook fires on new tags being pushed.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
     */
    token?: pulumi.Input<string>;
    /**
     * The hook URL.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemHook resource.
 */
export interface SystemHookArgs {
    /**
     * Do SSL verification when triggering the hook.
     */
    enableSslVerification?: pulumi.Input<boolean>;
    /**
     * Trigger hook on merge requests events.
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * When true, the hook fires on push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Trigger hook on repository update events.
     */
    repositoryUpdateEvents?: pulumi.Input<boolean>;
    /**
     * When true, the hook fires on new tags being pushed.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
     */
    token?: pulumi.Input<string>;
    /**
     * The hook URL.
     */
    url: pulumi.Input<string>;
}
