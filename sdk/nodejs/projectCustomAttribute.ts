// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectCustomAttribute` resource allows to manage custom attributes for a project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/custom_attributes.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const attr = new gitlab.ProjectCustomAttribute("attr", {
 *     project: 42,
 *     key: "location",
 *     value: "Greenland",
 * });
 * ```
 *
 * ## Import
 *
 * You can import a project custom attribute using an id made up of `{project-id}:{key}`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectCustomAttribute:ProjectCustomAttribute attr 42:location
 * ```
 */
export class ProjectCustomAttribute extends pulumi.CustomResource {
    /**
     * Get an existing ProjectCustomAttribute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectCustomAttributeState, opts?: pulumi.CustomResourceOptions): ProjectCustomAttribute {
        return new ProjectCustomAttribute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectCustomAttribute:ProjectCustomAttribute';

    /**
     * Returns true if the given object is an instance of ProjectCustomAttribute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectCustomAttribute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectCustomAttribute.__pulumiType;
    }

    /**
     * Key for the Custom Attribute.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The id of the project.
     */
    public readonly project!: pulumi.Output<number>;
    /**
     * Value for the Custom Attribute.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a ProjectCustomAttribute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectCustomAttributeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectCustomAttributeArgs | ProjectCustomAttributeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectCustomAttributeState | undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ProjectCustomAttributeArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectCustomAttribute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectCustomAttribute resources.
 */
export interface ProjectCustomAttributeState {
    /**
     * Key for the Custom Attribute.
     */
    key?: pulumi.Input<string>;
    /**
     * The id of the project.
     */
    project?: pulumi.Input<number>;
    /**
     * Value for the Custom Attribute.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectCustomAttribute resource.
 */
export interface ProjectCustomAttributeArgs {
    /**
     * Key for the Custom Attribute.
     */
    key: pulumi.Input<string>;
    /**
     * The id of the project.
     */
    project: pulumi.Input<number>;
    /**
     * Value for the Custom Attribute.
     */
    value: pulumi.Input<string>;
}
