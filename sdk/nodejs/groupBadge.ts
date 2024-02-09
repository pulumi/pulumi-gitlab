// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.GroupBadge` resource allows to manage the lifecycle of group badges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/badges.html#group-badges)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = new gitlab.Group("foo", {});
 * const example = new gitlab.GroupBadge("example", {
 *     group: foo.id,
 *     linkUrl: "https://example.com/badge-123",
 *     imageUrl: "https://example.com/badge-123.svg",
 * });
 * // Pipeline status badges with placeholders will be enabled for each project
 * const gitlabPipeline = new gitlab.GroupBadge("gitlabPipeline", {
 *     group: foo.id,
 *     linkUrl: "https://gitlab.example.com/%{project_path}/-/pipelines?ref=%{default_branch}",
 *     imageUrl: "https://gitlab.example.com/%{project_path}/badges/%{default_branch}/pipeline.svg",
 * });
 * // Test coverage report badges with placeholders will be enabled for each project
 * const gitlabCoverage = new gitlab.GroupBadge("gitlabCoverage", {
 *     group: foo.id,
 *     linkUrl: "https://gitlab.example.com/%{project_path}/-/jobs",
 *     imageUrl: "https://gitlab.example.com/%{project_path}/badges/%{default_branch}/coverage.svg",
 * });
 * // Latest release badges with placeholders will be enabled for each project
 * const gitlabRelease = new gitlab.GroupBadge("gitlabRelease", {
 *     group: foo.id,
 *     linkUrl: "https://gitlab.example.com/%{project_path}/-/releases",
 *     imageUrl: "https://gitlab.example.com/%{project_path}/-/badges/release.svg",
 * });
 * ```
 *
 * ## Import
 *
 * GitLab group badges can be imported using an id made up of `{group_id}:{badge_id}`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/groupBadge:GroupBadge foo 1:3
 * ```
 */
export class GroupBadge extends pulumi.CustomResource {
    /**
     * Get an existing GroupBadge resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupBadgeState, opts?: pulumi.CustomResourceOptions): GroupBadge {
        return new GroupBadge(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupBadge:GroupBadge';

    /**
     * Returns true if the given object is an instance of GroupBadge.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupBadge {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupBadge.__pulumiType;
    }

    /**
     * The id of the group to add the badge to.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * The image url which will be presented on group overview.
     */
    public readonly imageUrl!: pulumi.Output<string>;
    /**
     * The url linked with the badge.
     */
    public readonly linkUrl!: pulumi.Output<string>;
    /**
     * The name of the badge.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The imageUrl argument rendered (in case of use of placeholders).
     */
    public /*out*/ readonly renderedImageUrl!: pulumi.Output<string>;
    /**
     * The linkUrl argument rendered (in case of use of placeholders).
     */
    public /*out*/ readonly renderedLinkUrl!: pulumi.Output<string>;

    /**
     * Create a GroupBadge resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupBadgeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupBadgeArgs | GroupBadgeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupBadgeState | undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["imageUrl"] = state ? state.imageUrl : undefined;
            resourceInputs["linkUrl"] = state ? state.linkUrl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["renderedImageUrl"] = state ? state.renderedImageUrl : undefined;
            resourceInputs["renderedLinkUrl"] = state ? state.renderedLinkUrl : undefined;
        } else {
            const args = argsOrState as GroupBadgeArgs | undefined;
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.imageUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageUrl'");
            }
            if ((!args || args.linkUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'linkUrl'");
            }
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["imageUrl"] = args ? args.imageUrl : undefined;
            resourceInputs["linkUrl"] = args ? args.linkUrl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["renderedImageUrl"] = undefined /*out*/;
            resourceInputs["renderedLinkUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupBadge.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupBadge resources.
 */
export interface GroupBadgeState {
    /**
     * The id of the group to add the badge to.
     */
    group?: pulumi.Input<string>;
    /**
     * The image url which will be presented on group overview.
     */
    imageUrl?: pulumi.Input<string>;
    /**
     * The url linked with the badge.
     */
    linkUrl?: pulumi.Input<string>;
    /**
     * The name of the badge.
     */
    name?: pulumi.Input<string>;
    /**
     * The imageUrl argument rendered (in case of use of placeholders).
     */
    renderedImageUrl?: pulumi.Input<string>;
    /**
     * The linkUrl argument rendered (in case of use of placeholders).
     */
    renderedLinkUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupBadge resource.
 */
export interface GroupBadgeArgs {
    /**
     * The id of the group to add the badge to.
     */
    group: pulumi.Input<string>;
    /**
     * The image url which will be presented on group overview.
     */
    imageUrl: pulumi.Input<string>;
    /**
     * The url linked with the badge.
     */
    linkUrl: pulumi.Input<string>;
    /**
     * The name of the badge.
     */
    name?: pulumi.Input<string>;
}
