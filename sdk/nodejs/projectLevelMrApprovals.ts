// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlabProjectLevelMrApprovalRule` resource allows to manage the lifecycle of a Merge Request-level approval rule.
 *
 * > This resource requires a GitLab Enterprise instance.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/merge_request_approvals.html#merge-request-level-mr-approvals)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const fooProject = new gitlab.Project("fooProject", {description: "My example project"});
 * const fooProjectLevelMrApprovals = new gitlab.ProjectLevelMrApprovals("fooProjectLevelMrApprovals", {
 *     project: fooProject.id,
 *     resetApprovalsOnPush: true,
 *     disableOverridingApproversPerMergeRequest: false,
 *     mergeRequestsAuthorApproval: false,
 *     mergeRequestsDisableCommittersApproval: true,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals You can import an approval configuration state using `<resource> <project_id>`.
 * ```
 *
 * # 
 *
 * For example:
 *
 * ```sh
 * $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals foo 1234
 * ```
 */
export class ProjectLevelMrApprovals extends pulumi.CustomResource {
    /**
     * Get an existing ProjectLevelMrApprovals resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectLevelMrApprovalsState, opts?: pulumi.CustomResourceOptions): ProjectLevelMrApprovals {
        return new ProjectLevelMrApprovals(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals';

    /**
     * Returns true if the given object is an instance of ProjectLevelMrApprovals.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectLevelMrApprovals {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectLevelMrApprovals.__pulumiType;
    }

    /**
     * Set to `true` to disable overriding approvers per merge request.
     */
    public readonly disableOverridingApproversPerMergeRequest!: pulumi.Output<boolean>;
    /**
     * Set to `true` to allow merge requests authors to approve their own merge requests.
     */
    public readonly mergeRequestsAuthorApproval!: pulumi.Output<boolean>;
    /**
     * Set to `true` to disable merge request committers from approving their own merge requests.
     */
    public readonly mergeRequestsDisableCommittersApproval!: pulumi.Output<boolean>;
    /**
     * The ID or URL-encoded path of a project to change MR approval configuration.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Set to `true` to require authentication to approve merge requests.
     */
    public readonly requirePasswordToApprove!: pulumi.Output<boolean>;
    /**
     * Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
     */
    public readonly resetApprovalsOnPush!: pulumi.Output<boolean>;
    /**
     * Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
     */
    public readonly selectiveCodeOwnerRemovals!: pulumi.Output<boolean>;

    /**
     * Create a ProjectLevelMrApprovals resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectLevelMrApprovalsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectLevelMrApprovalsArgs | ProjectLevelMrApprovalsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectLevelMrApprovalsState | undefined;
            resourceInputs["disableOverridingApproversPerMergeRequest"] = state ? state.disableOverridingApproversPerMergeRequest : undefined;
            resourceInputs["mergeRequestsAuthorApproval"] = state ? state.mergeRequestsAuthorApproval : undefined;
            resourceInputs["mergeRequestsDisableCommittersApproval"] = state ? state.mergeRequestsDisableCommittersApproval : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["requirePasswordToApprove"] = state ? state.requirePasswordToApprove : undefined;
            resourceInputs["resetApprovalsOnPush"] = state ? state.resetApprovalsOnPush : undefined;
            resourceInputs["selectiveCodeOwnerRemovals"] = state ? state.selectiveCodeOwnerRemovals : undefined;
        } else {
            const args = argsOrState as ProjectLevelMrApprovalsArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["disableOverridingApproversPerMergeRequest"] = args ? args.disableOverridingApproversPerMergeRequest : undefined;
            resourceInputs["mergeRequestsAuthorApproval"] = args ? args.mergeRequestsAuthorApproval : undefined;
            resourceInputs["mergeRequestsDisableCommittersApproval"] = args ? args.mergeRequestsDisableCommittersApproval : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requirePasswordToApprove"] = args ? args.requirePasswordToApprove : undefined;
            resourceInputs["resetApprovalsOnPush"] = args ? args.resetApprovalsOnPush : undefined;
            resourceInputs["selectiveCodeOwnerRemovals"] = args ? args.selectiveCodeOwnerRemovals : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectLevelMrApprovals.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectLevelMrApprovals resources.
 */
export interface ProjectLevelMrApprovalsState {
    /**
     * Set to `true` to disable overriding approvers per merge request.
     */
    disableOverridingApproversPerMergeRequest?: pulumi.Input<boolean>;
    /**
     * Set to `true` to allow merge requests authors to approve their own merge requests.
     */
    mergeRequestsAuthorApproval?: pulumi.Input<boolean>;
    /**
     * Set to `true` to disable merge request committers from approving their own merge requests.
     */
    mergeRequestsDisableCommittersApproval?: pulumi.Input<boolean>;
    /**
     * The ID or URL-encoded path of a project to change MR approval configuration.
     */
    project?: pulumi.Input<string>;
    /**
     * Set to `true` to require authentication to approve merge requests.
     */
    requirePasswordToApprove?: pulumi.Input<boolean>;
    /**
     * Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
     */
    resetApprovalsOnPush?: pulumi.Input<boolean>;
    /**
     * Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
     */
    selectiveCodeOwnerRemovals?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ProjectLevelMrApprovals resource.
 */
export interface ProjectLevelMrApprovalsArgs {
    /**
     * Set to `true` to disable overriding approvers per merge request.
     */
    disableOverridingApproversPerMergeRequest?: pulumi.Input<boolean>;
    /**
     * Set to `true` to allow merge requests authors to approve their own merge requests.
     */
    mergeRequestsAuthorApproval?: pulumi.Input<boolean>;
    /**
     * Set to `true` to disable merge request committers from approving their own merge requests.
     */
    mergeRequestsDisableCommittersApproval?: pulumi.Input<boolean>;
    /**
     * The ID or URL-encoded path of a project to change MR approval configuration.
     */
    project: pulumi.Input<string>;
    /**
     * Set to `true` to require authentication to approve merge requests.
     */
    requirePasswordToApprove?: pulumi.Input<boolean>;
    /**
     * Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
     */
    resetApprovalsOnPush?: pulumi.Input<boolean>;
    /**
     * Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
     */
    selectiveCodeOwnerRemovals?: pulumi.Input<boolean>;
}
