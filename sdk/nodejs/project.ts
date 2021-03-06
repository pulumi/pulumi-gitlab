// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * Number of merge request approvals required for merging. Default is 0.
     */
    public readonly approvalsBeforeMerge!: pulumi.Output<number | undefined>;
    /**
     * Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
     */
    public readonly archived!: pulumi.Output<boolean | undefined>;
    /**
     * Enable container registry for the project.
     */
    public readonly containerRegistryEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The default branch for the project.
     */
    public readonly defaultBranch!: pulumi.Output<string | undefined>;
    /**
     * A description of the project.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * For group-level custom templates, specifies ID of group from which all the custom project templates are sourced. Leave empty for instance-level templates. Requires useCustomTemplate to be true (enterprise edition).
     */
    public readonly groupWithProjectTemplatesId!: pulumi.Output<number | undefined>;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via HTTP.
     */
    public /*out*/ readonly httpUrlToRepo!: pulumi.Output<string>;
    /**
     * Git URL to a repository to be imported.
     */
    public readonly importUrl!: pulumi.Output<string | undefined>;
    /**
     * Create master branch with first commit containing a README.md file.
     */
    public readonly initializeWithReadme!: pulumi.Output<boolean | undefined>;
    /**
     * Enable issue tracking for the project.
     */
    public readonly issuesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Enable LFS for the project.
     */
    public readonly lfsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `ff` to create fast-forward merges
     * Valid values are `merge`, `rebaseMerge`, `ff`
     * Repositories are created with `merge` by default
     */
    public readonly mergeMethod!: pulumi.Output<string | undefined>;
    /**
     * Enable merge requests for the project.
     */
    public readonly mergeRequestsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Enables pull mirroring in a project. Default is `false`. For further information on mirroring,
     * consult the [gitlab documentation](https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring).
     */
    public readonly mirror!: pulumi.Output<boolean | undefined>;
    /**
     * Pull mirror overwrites diverged branches.
     */
    public readonly mirrorOverwritesDivergedBranches!: pulumi.Output<boolean | undefined>;
    /**
     * Pull mirroring triggers builds. Default is `false`.
     */
    public readonly mirrorTriggerBuilds!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace (group or user) of the project. Defaults to your user.
     * See `gitlab.Group` for an example.
     */
    public readonly namespaceId!: pulumi.Output<number>;
    /**
     * Set to true if you want allow merges only if all discussions are resolved.
     */
    public readonly onlyAllowMergeIfAllDiscussionsAreResolved!: pulumi.Output<boolean | undefined>;
    /**
     * Set to true if you want allow merges only if a pipeline succeeds.
     */
    public readonly onlyAllowMergeIfPipelineSucceeds!: pulumi.Output<boolean | undefined>;
    /**
     * Only mirror protected branches.
     */
    public readonly onlyMirrorProtectedBranches!: pulumi.Output<boolean | undefined>;
    /**
     * Enable packages repository for the project.
     */
    public readonly packagesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Enable pages access control
     * Valid values are `disabled`, `private`, `enabled`, `public`.
     * `private` is the default.
     */
    public readonly pagesAccessLevel!: pulumi.Output<string | undefined>;
    /**
     * The path of the repository.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The path of the repository with namespace.
     */
    public /*out*/ readonly pathWithNamespace!: pulumi.Output<string>;
    /**
     * Enable pipelines for the project.
     */
    public readonly pipelinesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Push rules for the project (documented below).
     */
    public readonly pushRules!: pulumi.Output<outputs.ProjectPushRules>;
    /**
     * Enable `Delete source branch` option by default for all new merge requests.
     */
    public readonly removeSourceBranchAfterMerge!: pulumi.Output<boolean | undefined>;
    /**
     * Allow users to request member access.
     */
    public readonly requestAccessEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Registration token to use during runner setup.
     */
    public /*out*/ readonly runnersToken!: pulumi.Output<string>;
    /**
     * Enable shared runners for this project.
     */
    public readonly sharedRunnersEnabled!: pulumi.Output<boolean>;
    /**
     * Enable snippets for the project.
     */
    public readonly snippetsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via SSH.
     */
    public /*out*/ readonly sshUrlToRepo!: pulumi.Output<string>;
    /**
     * Tags (topics) of the project.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * When used without use_custom_template, name of a built-in project template. When used with use_custom_template, name of a custom project template. This option is mutually exclusive with `templateProjectId`.
     */
    public readonly templateName!: pulumi.Output<string | undefined>;
    /**
     * When used with use_custom_template, project ID of a custom project template. This is preferable to using templateName since templateName may be ambiguous (enterprise edition). This option is mutually exclusive with `templateName`.
     */
    public readonly templateProjectId!: pulumi.Output<number | undefined>;
    /**
     * Use either custom instance or group (with group_with_project_templates_id) project template (enterprise edition).
     */
    public readonly useCustomTemplate!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `public` to create a public project.
     * Valid values are `private`, `internal`, `public`.
     * Repositories are created as private by default.
     */
    public readonly visibilityLevel!: pulumi.Output<string | undefined>;
    /**
     * URL that can be used to find the project in a browser.
     */
    public /*out*/ readonly webUrl!: pulumi.Output<string>;
    /**
     * Enable wiki for the project.
     */
    public readonly wikiEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["approvalsBeforeMerge"] = state ? state.approvalsBeforeMerge : undefined;
            inputs["archived"] = state ? state.archived : undefined;
            inputs["containerRegistryEnabled"] = state ? state.containerRegistryEnabled : undefined;
            inputs["defaultBranch"] = state ? state.defaultBranch : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["groupWithProjectTemplatesId"] = state ? state.groupWithProjectTemplatesId : undefined;
            inputs["httpUrlToRepo"] = state ? state.httpUrlToRepo : undefined;
            inputs["importUrl"] = state ? state.importUrl : undefined;
            inputs["initializeWithReadme"] = state ? state.initializeWithReadme : undefined;
            inputs["issuesEnabled"] = state ? state.issuesEnabled : undefined;
            inputs["lfsEnabled"] = state ? state.lfsEnabled : undefined;
            inputs["mergeMethod"] = state ? state.mergeMethod : undefined;
            inputs["mergeRequestsEnabled"] = state ? state.mergeRequestsEnabled : undefined;
            inputs["mirror"] = state ? state.mirror : undefined;
            inputs["mirrorOverwritesDivergedBranches"] = state ? state.mirrorOverwritesDivergedBranches : undefined;
            inputs["mirrorTriggerBuilds"] = state ? state.mirrorTriggerBuilds : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceId"] = state ? state.namespaceId : undefined;
            inputs["onlyAllowMergeIfAllDiscussionsAreResolved"] = state ? state.onlyAllowMergeIfAllDiscussionsAreResolved : undefined;
            inputs["onlyAllowMergeIfPipelineSucceeds"] = state ? state.onlyAllowMergeIfPipelineSucceeds : undefined;
            inputs["onlyMirrorProtectedBranches"] = state ? state.onlyMirrorProtectedBranches : undefined;
            inputs["packagesEnabled"] = state ? state.packagesEnabled : undefined;
            inputs["pagesAccessLevel"] = state ? state.pagesAccessLevel : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["pathWithNamespace"] = state ? state.pathWithNamespace : undefined;
            inputs["pipelinesEnabled"] = state ? state.pipelinesEnabled : undefined;
            inputs["pushRules"] = state ? state.pushRules : undefined;
            inputs["removeSourceBranchAfterMerge"] = state ? state.removeSourceBranchAfterMerge : undefined;
            inputs["requestAccessEnabled"] = state ? state.requestAccessEnabled : undefined;
            inputs["runnersToken"] = state ? state.runnersToken : undefined;
            inputs["sharedRunnersEnabled"] = state ? state.sharedRunnersEnabled : undefined;
            inputs["snippetsEnabled"] = state ? state.snippetsEnabled : undefined;
            inputs["sshUrlToRepo"] = state ? state.sshUrlToRepo : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["templateName"] = state ? state.templateName : undefined;
            inputs["templateProjectId"] = state ? state.templateProjectId : undefined;
            inputs["useCustomTemplate"] = state ? state.useCustomTemplate : undefined;
            inputs["visibilityLevel"] = state ? state.visibilityLevel : undefined;
            inputs["webUrl"] = state ? state.webUrl : undefined;
            inputs["wikiEnabled"] = state ? state.wikiEnabled : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            inputs["approvalsBeforeMerge"] = args ? args.approvalsBeforeMerge : undefined;
            inputs["archived"] = args ? args.archived : undefined;
            inputs["containerRegistryEnabled"] = args ? args.containerRegistryEnabled : undefined;
            inputs["defaultBranch"] = args ? args.defaultBranch : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["groupWithProjectTemplatesId"] = args ? args.groupWithProjectTemplatesId : undefined;
            inputs["importUrl"] = args ? args.importUrl : undefined;
            inputs["initializeWithReadme"] = args ? args.initializeWithReadme : undefined;
            inputs["issuesEnabled"] = args ? args.issuesEnabled : undefined;
            inputs["lfsEnabled"] = args ? args.lfsEnabled : undefined;
            inputs["mergeMethod"] = args ? args.mergeMethod : undefined;
            inputs["mergeRequestsEnabled"] = args ? args.mergeRequestsEnabled : undefined;
            inputs["mirror"] = args ? args.mirror : undefined;
            inputs["mirrorOverwritesDivergedBranches"] = args ? args.mirrorOverwritesDivergedBranches : undefined;
            inputs["mirrorTriggerBuilds"] = args ? args.mirrorTriggerBuilds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceId"] = args ? args.namespaceId : undefined;
            inputs["onlyAllowMergeIfAllDiscussionsAreResolved"] = args ? args.onlyAllowMergeIfAllDiscussionsAreResolved : undefined;
            inputs["onlyAllowMergeIfPipelineSucceeds"] = args ? args.onlyAllowMergeIfPipelineSucceeds : undefined;
            inputs["onlyMirrorProtectedBranches"] = args ? args.onlyMirrorProtectedBranches : undefined;
            inputs["packagesEnabled"] = args ? args.packagesEnabled : undefined;
            inputs["pagesAccessLevel"] = args ? args.pagesAccessLevel : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["pipelinesEnabled"] = args ? args.pipelinesEnabled : undefined;
            inputs["pushRules"] = args ? args.pushRules : undefined;
            inputs["removeSourceBranchAfterMerge"] = args ? args.removeSourceBranchAfterMerge : undefined;
            inputs["requestAccessEnabled"] = args ? args.requestAccessEnabled : undefined;
            inputs["sharedRunnersEnabled"] = args ? args.sharedRunnersEnabled : undefined;
            inputs["snippetsEnabled"] = args ? args.snippetsEnabled : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["templateName"] = args ? args.templateName : undefined;
            inputs["templateProjectId"] = args ? args.templateProjectId : undefined;
            inputs["useCustomTemplate"] = args ? args.useCustomTemplate : undefined;
            inputs["visibilityLevel"] = args ? args.visibilityLevel : undefined;
            inputs["wikiEnabled"] = args ? args.wikiEnabled : undefined;
            inputs["httpUrlToRepo"] = undefined /*out*/;
            inputs["pathWithNamespace"] = undefined /*out*/;
            inputs["runnersToken"] = undefined /*out*/;
            inputs["sshUrlToRepo"] = undefined /*out*/;
            inputs["webUrl"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * Number of merge request approvals required for merging. Default is 0.
     */
    readonly approvalsBeforeMerge?: pulumi.Input<number>;
    /**
     * Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
     */
    readonly archived?: pulumi.Input<boolean>;
    /**
     * Enable container registry for the project.
     */
    readonly containerRegistryEnabled?: pulumi.Input<boolean>;
    /**
     * The default branch for the project.
     */
    readonly defaultBranch?: pulumi.Input<string>;
    /**
     * A description of the project.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * For group-level custom templates, specifies ID of group from which all the custom project templates are sourced. Leave empty for instance-level templates. Requires useCustomTemplate to be true (enterprise edition).
     */
    readonly groupWithProjectTemplatesId?: pulumi.Input<number>;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via HTTP.
     */
    readonly httpUrlToRepo?: pulumi.Input<string>;
    /**
     * Git URL to a repository to be imported.
     */
    readonly importUrl?: pulumi.Input<string>;
    /**
     * Create master branch with first commit containing a README.md file.
     */
    readonly initializeWithReadme?: pulumi.Input<boolean>;
    /**
     * Enable issue tracking for the project.
     */
    readonly issuesEnabled?: pulumi.Input<boolean>;
    /**
     * Enable LFS for the project.
     */
    readonly lfsEnabled?: pulumi.Input<boolean>;
    /**
     * Set to `ff` to create fast-forward merges
     * Valid values are `merge`, `rebaseMerge`, `ff`
     * Repositories are created with `merge` by default
     */
    readonly mergeMethod?: pulumi.Input<string>;
    /**
     * Enable merge requests for the project.
     */
    readonly mergeRequestsEnabled?: pulumi.Input<boolean>;
    /**
     * Enables pull mirroring in a project. Default is `false`. For further information on mirroring,
     * consult the [gitlab documentation](https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring).
     */
    readonly mirror?: pulumi.Input<boolean>;
    /**
     * Pull mirror overwrites diverged branches.
     */
    readonly mirrorOverwritesDivergedBranches?: pulumi.Input<boolean>;
    /**
     * Pull mirroring triggers builds. Default is `false`.
     */
    readonly mirrorTriggerBuilds?: pulumi.Input<boolean>;
    /**
     * The name of the project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The namespace (group or user) of the project. Defaults to your user.
     * See `gitlab.Group` for an example.
     */
    readonly namespaceId?: pulumi.Input<number>;
    /**
     * Set to true if you want allow merges only if all discussions are resolved.
     */
    readonly onlyAllowMergeIfAllDiscussionsAreResolved?: pulumi.Input<boolean>;
    /**
     * Set to true if you want allow merges only if a pipeline succeeds.
     */
    readonly onlyAllowMergeIfPipelineSucceeds?: pulumi.Input<boolean>;
    /**
     * Only mirror protected branches.
     */
    readonly onlyMirrorProtectedBranches?: pulumi.Input<boolean>;
    /**
     * Enable packages repository for the project.
     */
    readonly packagesEnabled?: pulumi.Input<boolean>;
    /**
     * Enable pages access control
     * Valid values are `disabled`, `private`, `enabled`, `public`.
     * `private` is the default.
     */
    readonly pagesAccessLevel?: pulumi.Input<string>;
    /**
     * The path of the repository.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The path of the repository with namespace.
     */
    readonly pathWithNamespace?: pulumi.Input<string>;
    /**
     * Enable pipelines for the project.
     */
    readonly pipelinesEnabled?: pulumi.Input<boolean>;
    /**
     * Push rules for the project (documented below).
     */
    readonly pushRules?: pulumi.Input<inputs.ProjectPushRules>;
    /**
     * Enable `Delete source branch` option by default for all new merge requests.
     */
    readonly removeSourceBranchAfterMerge?: pulumi.Input<boolean>;
    /**
     * Allow users to request member access.
     */
    readonly requestAccessEnabled?: pulumi.Input<boolean>;
    /**
     * Registration token to use during runner setup.
     */
    readonly runnersToken?: pulumi.Input<string>;
    /**
     * Enable shared runners for this project.
     */
    readonly sharedRunnersEnabled?: pulumi.Input<boolean>;
    /**
     * Enable snippets for the project.
     */
    readonly snippetsEnabled?: pulumi.Input<boolean>;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via SSH.
     */
    readonly sshUrlToRepo?: pulumi.Input<string>;
    /**
     * Tags (topics) of the project.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When used without use_custom_template, name of a built-in project template. When used with use_custom_template, name of a custom project template. This option is mutually exclusive with `templateProjectId`.
     */
    readonly templateName?: pulumi.Input<string>;
    /**
     * When used with use_custom_template, project ID of a custom project template. This is preferable to using templateName since templateName may be ambiguous (enterprise edition). This option is mutually exclusive with `templateName`.
     */
    readonly templateProjectId?: pulumi.Input<number>;
    /**
     * Use either custom instance or group (with group_with_project_templates_id) project template (enterprise edition).
     */
    readonly useCustomTemplate?: pulumi.Input<boolean>;
    /**
     * Set to `public` to create a public project.
     * Valid values are `private`, `internal`, `public`.
     * Repositories are created as private by default.
     */
    readonly visibilityLevel?: pulumi.Input<string>;
    /**
     * URL that can be used to find the project in a browser.
     */
    readonly webUrl?: pulumi.Input<string>;
    /**
     * Enable wiki for the project.
     */
    readonly wikiEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * Number of merge request approvals required for merging. Default is 0.
     */
    readonly approvalsBeforeMerge?: pulumi.Input<number>;
    /**
     * Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
     */
    readonly archived?: pulumi.Input<boolean>;
    /**
     * Enable container registry for the project.
     */
    readonly containerRegistryEnabled?: pulumi.Input<boolean>;
    /**
     * The default branch for the project.
     */
    readonly defaultBranch?: pulumi.Input<string>;
    /**
     * A description of the project.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * For group-level custom templates, specifies ID of group from which all the custom project templates are sourced. Leave empty for instance-level templates. Requires useCustomTemplate to be true (enterprise edition).
     */
    readonly groupWithProjectTemplatesId?: pulumi.Input<number>;
    /**
     * Git URL to a repository to be imported.
     */
    readonly importUrl?: pulumi.Input<string>;
    /**
     * Create master branch with first commit containing a README.md file.
     */
    readonly initializeWithReadme?: pulumi.Input<boolean>;
    /**
     * Enable issue tracking for the project.
     */
    readonly issuesEnabled?: pulumi.Input<boolean>;
    /**
     * Enable LFS for the project.
     */
    readonly lfsEnabled?: pulumi.Input<boolean>;
    /**
     * Set to `ff` to create fast-forward merges
     * Valid values are `merge`, `rebaseMerge`, `ff`
     * Repositories are created with `merge` by default
     */
    readonly mergeMethod?: pulumi.Input<string>;
    /**
     * Enable merge requests for the project.
     */
    readonly mergeRequestsEnabled?: pulumi.Input<boolean>;
    /**
     * Enables pull mirroring in a project. Default is `false`. For further information on mirroring,
     * consult the [gitlab documentation](https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring).
     */
    readonly mirror?: pulumi.Input<boolean>;
    /**
     * Pull mirror overwrites diverged branches.
     */
    readonly mirrorOverwritesDivergedBranches?: pulumi.Input<boolean>;
    /**
     * Pull mirroring triggers builds. Default is `false`.
     */
    readonly mirrorTriggerBuilds?: pulumi.Input<boolean>;
    /**
     * The name of the project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The namespace (group or user) of the project. Defaults to your user.
     * See `gitlab.Group` for an example.
     */
    readonly namespaceId?: pulumi.Input<number>;
    /**
     * Set to true if you want allow merges only if all discussions are resolved.
     */
    readonly onlyAllowMergeIfAllDiscussionsAreResolved?: pulumi.Input<boolean>;
    /**
     * Set to true if you want allow merges only if a pipeline succeeds.
     */
    readonly onlyAllowMergeIfPipelineSucceeds?: pulumi.Input<boolean>;
    /**
     * Only mirror protected branches.
     */
    readonly onlyMirrorProtectedBranches?: pulumi.Input<boolean>;
    /**
     * Enable packages repository for the project.
     */
    readonly packagesEnabled?: pulumi.Input<boolean>;
    /**
     * Enable pages access control
     * Valid values are `disabled`, `private`, `enabled`, `public`.
     * `private` is the default.
     */
    readonly pagesAccessLevel?: pulumi.Input<string>;
    /**
     * The path of the repository.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Enable pipelines for the project.
     */
    readonly pipelinesEnabled?: pulumi.Input<boolean>;
    /**
     * Push rules for the project (documented below).
     */
    readonly pushRules?: pulumi.Input<inputs.ProjectPushRules>;
    /**
     * Enable `Delete source branch` option by default for all new merge requests.
     */
    readonly removeSourceBranchAfterMerge?: pulumi.Input<boolean>;
    /**
     * Allow users to request member access.
     */
    readonly requestAccessEnabled?: pulumi.Input<boolean>;
    /**
     * Enable shared runners for this project.
     */
    readonly sharedRunnersEnabled?: pulumi.Input<boolean>;
    /**
     * Enable snippets for the project.
     */
    readonly snippetsEnabled?: pulumi.Input<boolean>;
    /**
     * Tags (topics) of the project.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When used without use_custom_template, name of a built-in project template. When used with use_custom_template, name of a custom project template. This option is mutually exclusive with `templateProjectId`.
     */
    readonly templateName?: pulumi.Input<string>;
    /**
     * When used with use_custom_template, project ID of a custom project template. This is preferable to using templateName since templateName may be ambiguous (enterprise edition). This option is mutually exclusive with `templateName`.
     */
    readonly templateProjectId?: pulumi.Input<number>;
    /**
     * Use either custom instance or group (with group_with_project_templates_id) project template (enterprise edition).
     */
    readonly useCustomTemplate?: pulumi.Input<boolean>;
    /**
     * Set to `public` to create a public project.
     * Valid values are `private`, `internal`, `public`.
     * Repositories are created as private by default.
     */
    readonly visibilityLevel?: pulumi.Input<string>;
    /**
     * Enable wiki for the project.
     */
    readonly wikiEnabled?: pulumi.Input<boolean>;
}
