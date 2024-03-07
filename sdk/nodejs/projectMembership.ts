// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectMembership` resource allows to manage the lifecycle of a users project membership.
 *
 * > If a project should grant membership to an entire group use the `gitlab.ProjectShareGroup` resource instead.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const test = new gitlab.ProjectMembership("test", {
 *     accessLevel: "guest",
 *     project: "12345",
 *     userId: 1337,
 * });
 * const example = new gitlab.ProjectMembership("example", {
 *     accessLevel: "guest",
 *     expiresAt: "2022-12-31",
 *     project: "67890",
 *     userId: 1234,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * GitLab project membership can be imported using an id made up of `project_id:user_id`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectMembership:ProjectMembership test "12345:1337"
 * ```
 */
export class ProjectMembership extends pulumi.CustomResource {
    /**
     * Get an existing ProjectMembership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectMembershipState, opts?: pulumi.CustomResourceOptions): ProjectMembership {
        return new ProjectMembership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectMembership:ProjectMembership';

    /**
     * Returns true if the given object is an instance of ProjectMembership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectMembership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectMembership.__pulumiType;
    }

    /**
     * The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     */
    public readonly accessLevel!: pulumi.Output<string>;
    /**
     * Expiration date for the project membership. Format: `YYYY-MM-DD`
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * The ID or URL-encoded path of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The id of the user.
     */
    public readonly userId!: pulumi.Output<number>;

    /**
     * Create a ProjectMembership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectMembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectMembershipArgs | ProjectMembershipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectMembershipState | undefined;
            resourceInputs["accessLevel"] = state ? state.accessLevel : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as ProjectMembershipArgs | undefined;
            if ((!args || args.accessLevel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessLevel'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["accessLevel"] = args ? args.accessLevel : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectMembership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectMembership resources.
 */
export interface ProjectMembershipState {
    /**
     * The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Expiration date for the project membership. Format: `YYYY-MM-DD`
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID or URL-encoded path of the project.
     */
    project?: pulumi.Input<string>;
    /**
     * The id of the user.
     */
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProjectMembership resource.
 */
export interface ProjectMembershipArgs {
    /**
     * The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     */
    accessLevel: pulumi.Input<string>;
    /**
     * Expiration date for the project membership. Format: `YYYY-MM-DD`
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID or URL-encoded path of the project.
     */
    project: pulumi.Input<string>;
    /**
     * The id of the user.
     */
    userId: pulumi.Input<number>;
}
