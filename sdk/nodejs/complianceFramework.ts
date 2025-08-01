// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ComplianceFramework` resource allows to manage the lifecycle of a compliance framework on top-level groups.
 *
 * There can be only one `default` compliance framework. Of all the configured compliance frameworks marked as default, the last one applied will be the default compliance framework.
 *
 * > This resource requires a GitLab Enterprise instance with a Premium license to create the compliance framework.
 *
 * > This resource requires a GitLab Enterprise instance with an Ultimate license to specify a compliance pipeline configuration in the compliance framework.
 *
 * **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/api/graphql/reference/#mutationcreatecomplianceframework)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const sample = new gitlab.ComplianceFramework("sample", {
 *     namespacePath: "top-level-group",
 *     name: "HIPAA",
 *     description: "A HIPAA Compliance Framework",
 *     color: "#87BEEF",
 *     "default": false,
 *     pipelineConfigurationFullPath: ".hipaa.yml@top-level-group/compliance-frameworks",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_compliance_framework`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_compliance_framework.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * Gitlab compliance frameworks can be imported with a key composed of `<namespace_path>:<framework_id>`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/complianceFramework:ComplianceFramework sample "top-level-group:gid://gitlab/ComplianceManagement::Framework/12345"
 * ```
 */
export class ComplianceFramework extends pulumi.CustomResource {
    /**
     * Get an existing ComplianceFramework resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComplianceFrameworkState, opts?: pulumi.CustomResourceOptions): ComplianceFramework {
        return new ComplianceFramework(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/complianceFramework:ComplianceFramework';

    /**
     * Returns true if the given object is an instance of ComplianceFramework.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComplianceFramework {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComplianceFramework.__pulumiType;
    }

    /**
     * New color representation of the compliance framework in hex format. e.g. #FCA121.
     */
    public readonly color!: pulumi.Output<string>;
    /**
     * Set this compliance framework as the default framework for the group. Default: `false`
     */
    public readonly default!: pulumi.Output<boolean>;
    /**
     * Description for the compliance framework.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Globally unique ID of the compliance framework.
     */
    public /*out*/ readonly frameworkId!: pulumi.Output<string>;
    /**
     * Name for the compliance framework.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Full path of the namespace to add the compliance framework to.
     */
    public readonly namespacePath!: pulumi.Output<string>;
    /**
     * Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
     */
    public readonly pipelineConfigurationFullPath!: pulumi.Output<string | undefined>;

    /**
     * Create a ComplianceFramework resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComplianceFrameworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComplianceFrameworkArgs | ComplianceFrameworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComplianceFrameworkState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["default"] = state ? state.default : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["frameworkId"] = state ? state.frameworkId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespacePath"] = state ? state.namespacePath : undefined;
            resourceInputs["pipelineConfigurationFullPath"] = state ? state.pipelineConfigurationFullPath : undefined;
        } else {
            const args = argsOrState as ComplianceFrameworkArgs | undefined;
            if ((!args || args.color === undefined) && !opts.urn) {
                throw new Error("Missing required property 'color'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.namespacePath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespacePath'");
            }
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["default"] = args ? args.default : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespacePath"] = args ? args.namespacePath : undefined;
            resourceInputs["pipelineConfigurationFullPath"] = args ? args.pipelineConfigurationFullPath : undefined;
            resourceInputs["frameworkId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ComplianceFramework.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComplianceFramework resources.
 */
export interface ComplianceFrameworkState {
    /**
     * New color representation of the compliance framework in hex format. e.g. #FCA121.
     */
    color?: pulumi.Input<string>;
    /**
     * Set this compliance framework as the default framework for the group. Default: `false`
     */
    default?: pulumi.Input<boolean>;
    /**
     * Description for the compliance framework.
     */
    description?: pulumi.Input<string>;
    /**
     * Globally unique ID of the compliance framework.
     */
    frameworkId?: pulumi.Input<string>;
    /**
     * Name for the compliance framework.
     */
    name?: pulumi.Input<string>;
    /**
     * Full path of the namespace to add the compliance framework to.
     */
    namespacePath?: pulumi.Input<string>;
    /**
     * Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
     */
    pipelineConfigurationFullPath?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComplianceFramework resource.
 */
export interface ComplianceFrameworkArgs {
    /**
     * New color representation of the compliance framework in hex format. e.g. #FCA121.
     */
    color: pulumi.Input<string>;
    /**
     * Set this compliance framework as the default framework for the group. Default: `false`
     */
    default?: pulumi.Input<boolean>;
    /**
     * Description for the compliance framework.
     */
    description: pulumi.Input<string>;
    /**
     * Name for the compliance framework.
     */
    name?: pulumi.Input<string>;
    /**
     * Full path of the namespace to add the compliance framework to.
     */
    namespacePath: pulumi.Input<string>;
    /**
     * Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
     */
    pipelineConfigurationFullPath?: pulumi.Input<string>;
}
