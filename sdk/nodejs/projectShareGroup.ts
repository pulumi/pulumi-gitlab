// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to share a project with a group
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const test = new gitlab.ProjectShareGroup("test", {
 *     groupAccess: "guest",
 *     groupId: 1337,
 *     projectId: "12345",
 * });
 * ```
 *
 * ## Import
 *
 * # GitLab project group shares can be imported using an id made up of `projectid:groupid`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/projectShareGroup:ProjectShareGroup test 12345:1337
 * ```
 */
export class ProjectShareGroup extends pulumi.CustomResource {
    /**
     * Get an existing ProjectShareGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectShareGroupState, opts?: pulumi.CustomResourceOptions): ProjectShareGroup {
        return new ProjectShareGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectShareGroup:ProjectShareGroup';

    /**
     * Returns true if the given object is an instance of ProjectShareGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectShareGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectShareGroup.__pulumiType;
    }

    /**
     * The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `master`
     *
     * @deprecated Use `group_access` instead of the `access_level` attribute.
     */
    public readonly accessLevel!: pulumi.Output<string | undefined>;
    /**
     * The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `master`
     */
    public readonly groupAccess!: pulumi.Output<string | undefined>;
    /**
     * The id of the group.
     */
    public readonly groupId!: pulumi.Output<number>;
    /**
     * The id of the project.
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a ProjectShareGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectShareGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectShareGroupArgs | ProjectShareGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectShareGroupState | undefined;
            resourceInputs["accessLevel"] = state ? state.accessLevel : undefined;
            resourceInputs["groupAccess"] = state ? state.groupAccess : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as ProjectShareGroupArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["accessLevel"] = args ? args.accessLevel : undefined;
            resourceInputs["groupAccess"] = args ? args.groupAccess : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectShareGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectShareGroup resources.
 */
export interface ProjectShareGroupState {
    /**
     * The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `master`
     *
     * @deprecated Use `group_access` instead of the `access_level` attribute.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `master`
     */
    groupAccess?: pulumi.Input<string>;
    /**
     * The id of the group.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The id of the project.
     */
    projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectShareGroup resource.
 */
export interface ProjectShareGroupArgs {
    /**
     * The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `master`
     *
     * @deprecated Use `group_access` instead of the `access_level` attribute.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `master`
     */
    groupAccess?: pulumi.Input<string>;
    /**
     * The id of the group.
     */
    groupId: pulumi.Input<number>;
    /**
     * The id of the project.
     */
    projectId: pulumi.Input<string>;
}
