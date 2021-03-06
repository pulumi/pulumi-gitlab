// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # gitlab\_tag\_protection
 *
 * This resource allows you to protect a specific tag or wildcard by an access level so that the user with less access level cannot Create the tags.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const tagProtect = new gitlab.TagProtection("TagProtect", {
 *     createAccessLevel: "developer",
 *     project: "12345",
 *     tag: "TagProtected",
 * });
 * ```
 */
export class TagProtection extends pulumi.CustomResource {
    /**
     * Get an existing TagProtection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagProtectionState, opts?: pulumi.CustomResourceOptions): TagProtection {
        return new TagProtection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/tagProtection:TagProtection';

    /**
     * Returns true if the given object is an instance of TagProtection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagProtection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagProtection.__pulumiType;
    }

    /**
     * One of five levels of access to the project.
     */
    public readonly createAccessLevel!: pulumi.Output<string>;
    /**
     * The id of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Name of the tag or wildcard.
     */
    public readonly tag!: pulumi.Output<string>;

    /**
     * Create a TagProtection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagProtectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagProtectionArgs | TagProtectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagProtectionState | undefined;
            inputs["createAccessLevel"] = state ? state.createAccessLevel : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["tag"] = state ? state.tag : undefined;
        } else {
            const args = argsOrState as TagProtectionArgs | undefined;
            if ((!args || args.createAccessLevel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'createAccessLevel'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.tag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tag'");
            }
            inputs["createAccessLevel"] = args ? args.createAccessLevel : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["tag"] = args ? args.tag : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TagProtection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TagProtection resources.
 */
export interface TagProtectionState {
    /**
     * One of five levels of access to the project.
     */
    readonly createAccessLevel?: pulumi.Input<string>;
    /**
     * The id of the project.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Name of the tag or wildcard.
     */
    readonly tag?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TagProtection resource.
 */
export interface TagProtectionArgs {
    /**
     * One of five levels of access to the project.
     */
    readonly createAccessLevel: pulumi.Input<string>;
    /**
     * The id of the project.
     */
    readonly project: pulumi.Input<string>;
    /**
     * Name of the tag or wildcard.
     */
    readonly tag: pulumi.Input<string>;
}
