// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectBadge` resource allows to mange the lifecycle of project badges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/badges.html#project-badges)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = new gitlab.Project("foo", {});
 * const example = new gitlab.ProjectBadge("example", {
 *     project: foo.id,
 *     linkUrl: "https://example.com/badge-123",
 *     imageUrl: "https://example.com/badge-123.svg",
 * });
 * ```
 *
 * ## Import
 *
 * # GitLab project badges can be imported using an id made up of `{project_id}:{badge_id}`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/projectBadge:ProjectBadge foo 1:3
 * ```
 */
export class ProjectBadge extends pulumi.CustomResource {
    /**
     * Get an existing ProjectBadge resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectBadgeState, opts?: pulumi.CustomResourceOptions): ProjectBadge {
        return new ProjectBadge(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectBadge:ProjectBadge';

    /**
     * Returns true if the given object is an instance of ProjectBadge.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectBadge {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectBadge.__pulumiType;
    }

    /**
     * The image url which will be presented on project overview.
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
     * The id of the project to add the badge to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The image_url argument rendered (in case of use of placeholders).
     */
    public /*out*/ readonly renderedImageUrl!: pulumi.Output<string>;
    /**
     * The link_url argument rendered (in case of use of placeholders).
     */
    public /*out*/ readonly renderedLinkUrl!: pulumi.Output<string>;

    /**
     * Create a ProjectBadge resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectBadgeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectBadgeArgs | ProjectBadgeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectBadgeState | undefined;
            resourceInputs["imageUrl"] = state ? state.imageUrl : undefined;
            resourceInputs["linkUrl"] = state ? state.linkUrl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["renderedImageUrl"] = state ? state.renderedImageUrl : undefined;
            resourceInputs["renderedLinkUrl"] = state ? state.renderedLinkUrl : undefined;
        } else {
            const args = argsOrState as ProjectBadgeArgs | undefined;
            if ((!args || args.imageUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageUrl'");
            }
            if ((!args || args.linkUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'linkUrl'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["imageUrl"] = args ? args.imageUrl : undefined;
            resourceInputs["linkUrl"] = args ? args.linkUrl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["renderedImageUrl"] = undefined /*out*/;
            resourceInputs["renderedLinkUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectBadge.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectBadge resources.
 */
export interface ProjectBadgeState {
    /**
     * The image url which will be presented on project overview.
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
     * The id of the project to add the badge to.
     */
    project?: pulumi.Input<string>;
    /**
     * The image_url argument rendered (in case of use of placeholders).
     */
    renderedImageUrl?: pulumi.Input<string>;
    /**
     * The link_url argument rendered (in case of use of placeholders).
     */
    renderedLinkUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectBadge resource.
 */
export interface ProjectBadgeArgs {
    /**
     * The image url which will be presented on project overview.
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
    /**
     * The id of the project to add the badge to.
     */
    project: pulumi.Input<string>;
}
