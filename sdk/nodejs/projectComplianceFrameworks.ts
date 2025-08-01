// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectComplianceFrameworks` resource allows to manage the lifecycle of compliance frameworks on a project.
 *
 * > This resource requires a GitLab Enterprise instance with a Premium license to set the compliance frameworks on a project.
 *
 * **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/api/graphql/reference/#mutationprojectupdatecomplianceframeworks)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const alpha = new gitlab.ComplianceFramework("alpha", {
 *     namespacePath: "top-level-group",
 *     name: "HIPAA",
 *     description: "A HIPAA Compliance Framework",
 *     color: "#87BEEF",
 *     "default": false,
 * });
 * const beta = new gitlab.ComplianceFramework("beta", {
 *     namespacePath: "top-level-group",
 *     name: "SOC",
 *     description: "A SOC Compliance Framework",
 *     color: "#223344",
 *     "default": false,
 * });
 * const sample = new gitlab.ProjectComplianceFrameworks("sample", {
 *     complianceFrameworkIds: [
 *         alpha.frameworkId,
 *         beta.frameworkId,
 *     ],
 *     project: "12345678",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_compliance_frameworks`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_project_compliance_frameworks.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * Gitlab project compliance frameworks can be imported with a key composed of `<project_id>`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectComplianceFrameworks:ProjectComplianceFrameworks sample "42"
 * ```
 */
export class ProjectComplianceFrameworks extends pulumi.CustomResource {
    /**
     * Get an existing ProjectComplianceFrameworks resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectComplianceFrameworksState, opts?: pulumi.CustomResourceOptions): ProjectComplianceFrameworks {
        return new ProjectComplianceFrameworks(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectComplianceFrameworks:ProjectComplianceFrameworks';

    /**
     * Returns true if the given object is an instance of ProjectComplianceFrameworks.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectComplianceFrameworks {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectComplianceFrameworks.__pulumiType;
    }

    /**
     * Globally unique IDs of the compliance frameworks to assign to the project.
     */
    public readonly complianceFrameworkIds!: pulumi.Output<string[]>;
    /**
     * The ID or full path of the project to change the compliance frameworks of.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a ProjectComplianceFrameworks resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectComplianceFrameworksArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectComplianceFrameworksArgs | ProjectComplianceFrameworksState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectComplianceFrameworksState | undefined;
            resourceInputs["complianceFrameworkIds"] = state ? state.complianceFrameworkIds : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ProjectComplianceFrameworksArgs | undefined;
            if ((!args || args.complianceFrameworkIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'complianceFrameworkIds'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["complianceFrameworkIds"] = args ? args.complianceFrameworkIds : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectComplianceFrameworks.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectComplianceFrameworks resources.
 */
export interface ProjectComplianceFrameworksState {
    /**
     * Globally unique IDs of the compliance frameworks to assign to the project.
     */
    complianceFrameworkIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID or full path of the project to change the compliance frameworks of.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectComplianceFrameworks resource.
 */
export interface ProjectComplianceFrameworksArgs {
    /**
     * Globally unique IDs of the compliance frameworks to assign to the project.
     */
    complianceFrameworkIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID or full path of the project to change the compliance frameworks of.
     */
    project: pulumi.Input<string>;
}
