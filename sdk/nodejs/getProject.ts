// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.Project` data source allows details of a project to be retrieved by either its ID or its path with namespace.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#get-single-project)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getProject({
 *     id: "foo/bar/baz",
 * });
 * ```
 */
export function getProject(args?: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProject:getProject", {
        "ciDefaultGitDepth": args.ciDefaultGitDepth,
        "id": args.id,
        "pathWithNamespace": args.pathWithNamespace,
        "publicBuilds": args.publicBuilds,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * Default number of revisions for shallow cloning.
     */
    ciDefaultGitDepth?: number;
    /**
     * The integer or path with namespace that uniquely identifies the project within the gitlab install.
     */
    id?: string;
    /**
     * The path of the repository with namespace.
     */
    pathWithNamespace?: string;
    /**
     * If true, jobs can be viewed by non-project members.
     */
    publicBuilds?: boolean;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * Set the analytics access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly analyticsAccessLevel: string;
    /**
     * Whether the project is in read-only mode (archived).
     */
    readonly archived: boolean;
    /**
     * Auto-cancel pending pipelines. This isn’t a boolean, but enabled/disabled.
     */
    readonly autoCancelPendingPipelines: string;
    /**
     * Auto Deploy strategy. Valid values are `continuous`, `manual`, `timedIncremental`.
     */
    readonly autoDevopsDeployStrategy: string;
    /**
     * Enable Auto DevOps for this project.
     */
    readonly autoDevopsEnabled: boolean;
    /**
     * Set whether auto-closing referenced issues on default branch.
     */
    readonly autocloseReferencedIssues: boolean;
    /**
     * The Git strategy. Defaults to fetch.
     */
    readonly buildGitStrategy: string;
    /**
     * The maximum amount of time, in seconds, that a job can run.
     */
    readonly buildTimeout: number;
    /**
     * Set the builds access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly buildsAccessLevel: string;
    /**
     * CI config file path for the project.
     */
    readonly ciConfigPath: string;
    /**
     * Default number of revisions for shallow cloning.
     */
    readonly ciDefaultGitDepth: number;
    /**
     * Set the image cleanup policy for this project. **Note**: this field is sometimes named `containerExpirationPolicyAttributes` in the GitLab Upstream API.
     */
    readonly containerExpirationPolicies: outputs.GetProjectContainerExpirationPolicy[];
    /**
     * Set visibility of container registry, for this project. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly containerRegistryAccessLevel: string;
    /**
     * The default branch for the project.
     */
    readonly defaultBranch: string;
    /**
     * A description of the project.
     */
    readonly description: string;
    /**
     * Disable email notifications.
     */
    readonly emailsDisabled: boolean;
    /**
     * The classification label for the project.
     */
    readonly externalAuthorizationClassificationLabel: string;
    /**
     * Set the forking access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly forkingAccessLevel: string;
    /**
     * URL that can be provided to `git clone` to clone the
     */
    readonly httpUrlToRepo: string;
    /**
     * The integer or path with namespace that uniquely identifies the project within the gitlab install.
     */
    readonly id: string;
    /**
     * Set the issues access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly issuesAccessLevel: string;
    /**
     * Enable issue tracking for the project.
     */
    readonly issuesEnabled: boolean;
    /**
     * Enable LFS for the project.
     */
    readonly lfsEnabled: boolean;
    /**
     * Template used to create merge commit message in merge requests. (Introduced in GitLab 14.5.)
     */
    readonly mergeCommitTemplate: string;
    /**
     * Enable or disable merge pipelines.
     */
    readonly mergePipelinesEnabled: boolean;
    /**
     * Set the merge requests access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly mergeRequestsAccessLevel: string;
    /**
     * Enable merge requests for the project.
     */
    readonly mergeRequestsEnabled: boolean;
    /**
     * Enable or disable merge trains.
     */
    readonly mergeTrainsEnabled: boolean;
    /**
     * The name of the project.
     */
    readonly name: string;
    /**
     * The namespace (group or user) of the project. Defaults to your user.
     */
    readonly namespaceId: number;
    /**
     * Set the operations access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly operationsAccessLevel: string;
    /**
     * The path of the repository.
     */
    readonly path: string;
    /**
     * The path of the repository with namespace.
     */
    readonly pathWithNamespace: string;
    /**
     * Enable pipelines for the project.
     */
    readonly pipelinesEnabled: boolean;
    /**
     * Show link to create/view merge request when pushing from the command line
     */
    readonly printingMergeRequestLinkEnabled: boolean;
    /**
     * If true, jobs can be viewed by non-project members.
     */
    readonly publicBuilds?: boolean;
    /**
     * Push rules for the project.
     */
    readonly pushRules: outputs.GetProjectPushRules;
    /**
     * Enable `Delete source branch` option by default for all new merge requests
     */
    readonly removeSourceBranchAfterMerge: boolean;
    /**
     * Set the repository access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly repositoryAccessLevel: string;
    /**
     * Which storage shard the repository is on. (administrator only)
     */
    readonly repositoryStorage: string;
    /**
     * Allow users to request member access.
     */
    readonly requestAccessEnabled: boolean;
    /**
     * Set the requirements access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly requirementsAccessLevel: string;
    /**
     * Automatically resolve merge request diffs discussions on lines changed with a push.
     */
    readonly resolveOutdatedDiffDiscussions: boolean;
    /**
     * Registration token to use during runner setup.
     */
    readonly runnersToken: string;
    /**
     * Set the security and compliance access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly securityAndComplianceAccessLevel: string;
    /**
     * Set the snippets access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly snippetsAccessLevel: string;
    /**
     * Enable snippets for the project.
     */
    readonly snippetsEnabled: boolean;
    /**
     * Template used to create squash commit message in merge requests. (Introduced in GitLab 14.6.)
     */
    readonly squashCommitTemplate: string;
    /**
     * URL that can be provided to `git clone` to clone the
     */
    readonly sshUrlToRepo: string;
    /**
     * The commit message used to apply merge request suggestions.
     */
    readonly suggestionCommitMessage: string;
    /**
     * The list of topics for the project.
     */
    readonly topics: string[];
    /**
     * Repositories are created as private by default.
     */
    readonly visibilityLevel: string;
    /**
     * URL that can be used to find the project in a browser.
     */
    readonly webUrl: string;
    /**
     * Set the wiki access level. Valid values are `disabled`, `private`, `enabled`.
     */
    readonly wikiAccessLevel: string;
    /**
     * Enable wiki for the project.
     */
    readonly wikiEnabled: boolean;
}
/**
 * The `gitlab.Project` data source allows details of a project to be retrieved by either its ID or its path with namespace.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#get-single-project)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getProject({
 *     id: "foo/bar/baz",
 * });
 * ```
 */
export function getProjectOutput(args?: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply((a: any) => getProject(a, opts))
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectOutputArgs {
    /**
     * Default number of revisions for shallow cloning.
     */
    ciDefaultGitDepth?: pulumi.Input<number>;
    /**
     * The integer or path with namespace that uniquely identifies the project within the gitlab install.
     */
    id?: pulumi.Input<string>;
    /**
     * The path of the repository with namespace.
     */
    pathWithNamespace?: pulumi.Input<string>;
    /**
     * If true, jobs can be viewed by non-project members.
     */
    publicBuilds?: pulumi.Input<boolean>;
}
