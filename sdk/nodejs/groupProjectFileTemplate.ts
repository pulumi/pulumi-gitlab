// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.GroupProjectFileTemplate` resource allows setting a project from which
 * custom file templates will be loaded. In order to use this resource, the project selected must be a direct child of
 * the group selected. After the resource has run, `gitlab_project_template.template_project_id` is available for use.
 * For more information about which file types are available as templates, view
 * [GitLab's documentation](https://docs.gitlab.com/user/group/custom_project_templates/)
 *
 * > This resource requires a GitLab Enterprise instance with a Premium license.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#update-group-attributes)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = new gitlab.Group("foo", {
 *     name: "group",
 *     path: "group",
 *     description: "An example group",
 * });
 * const bar = new gitlab.Project("bar", {
 *     name: "template project",
 *     description: "contains file templates",
 *     visibilityLevel: "public",
 *     namespaceId: foo.id,
 * });
 * const templateLink = new gitlab.GroupProjectFileTemplate("template_link", {
 *     groupId: foo.id,
 *     fileTemplateProjectId: bar.id,
 * });
 * ```
 */
export class GroupProjectFileTemplate extends pulumi.CustomResource {
    /**
     * Get an existing GroupProjectFileTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupProjectFileTemplateState, opts?: pulumi.CustomResourceOptions): GroupProjectFileTemplate {
        return new GroupProjectFileTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate';

    /**
     * Returns true if the given object is an instance of GroupProjectFileTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupProjectFileTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupProjectFileTemplate.__pulumiType;
    }

    /**
     * The ID of the project that will be used for file templates. This project must be the direct
     * 			child of the project defined by the group_id
     */
    public readonly fileTemplateProjectId!: pulumi.Output<number>;
    /**
     * The ID of the group that will use the file template project. This group must be the direct
     *             parent of the project defined by project_id
     */
    public readonly groupId!: pulumi.Output<number>;

    /**
     * Create a GroupProjectFileTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupProjectFileTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupProjectFileTemplateArgs | GroupProjectFileTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupProjectFileTemplateState | undefined;
            resourceInputs["fileTemplateProjectId"] = state ? state.fileTemplateProjectId : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
        } else {
            const args = argsOrState as GroupProjectFileTemplateArgs | undefined;
            if ((!args || args.fileTemplateProjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileTemplateProjectId'");
            }
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            resourceInputs["fileTemplateProjectId"] = args ? args.fileTemplateProjectId : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupProjectFileTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupProjectFileTemplate resources.
 */
export interface GroupProjectFileTemplateState {
    /**
     * The ID of the project that will be used for file templates. This project must be the direct
     * 			child of the project defined by the group_id
     */
    fileTemplateProjectId?: pulumi.Input<number>;
    /**
     * The ID of the group that will use the file template project. This group must be the direct
     *             parent of the project defined by project_id
     */
    groupId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GroupProjectFileTemplate resource.
 */
export interface GroupProjectFileTemplateArgs {
    /**
     * The ID of the project that will be used for file templates. This project must be the direct
     * 			child of the project defined by the group_id
     */
    fileTemplateProjectId: pulumi.Input<number>;
    /**
     * The ID of the group that will use the file template project. This group must be the direct
     *             parent of the project defined by project_id
     */
    groupId: pulumi.Input<number>;
}
