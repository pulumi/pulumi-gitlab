// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.Group("example", {
 *     name: "example",
 *     path: "example",
 *     description: "An example group",
 * });
 * // Create a project in the example group
 * const exampleProject = new gitlab.Project("example", {
 *     name: "example",
 *     description: "An example project",
 *     namespaceId: example.id,
 * });
 * // Group with custom push rules
 * const example_two = new gitlab.Group("example-two", {
 *     name: "example-two",
 *     path: "example-two",
 *     description: "An example group with push rules",
 *     pushRules: {
 *         authorEmailRegex: "@example\\.com$",
 *         commitCommitterCheck: true,
 *         memberCheck: true,
 *         preventSecrets: true,
 *     },
 * });
 * // Group with custom default branch protection defaults
 * const example_three = new gitlab.Group("example-three", {
 *     name: "example-three",
 *     path: "example-three",
 *     description: "An example group with default branch protection defaults",
 *     defaultBranchProtectionDefaults: {
 *         allowedToPushes: ["developer"],
 *         allowForcePush: true,
 *         allowedToMerges: [
 *             "developer",
 *             "maintainer",
 *         ],
 *         developerCanInitialPush: true,
 *     },
 * });
 * // Group with custom default branch protection defaults
 * const example_four = new gitlab.Group("example-four", {
 *     name: "example-four",
 *     path: "example-four",
 *     description: "An example group with default branch protection defaults",
 *     defaultBranchProtectionDefaults: {
 *         allowedToPushes: ["no one"],
 *         allowForcePush: true,
 *         allowedToMerges: ["no one"],
 *         developerCanInitialPush: true,
 *     },
 * });
 * // Group with a default branch name specified
 * const example_five = new gitlab.Group("example-five", {
 *     name: "example",
 *     path: "example",
 *     defaultBranch: "develop",
 *     description: "An example group with a default branch name",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_group`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_group.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * ```sh
 * $ pulumi import gitlab:index/group:Group You can import a group state using `<resource> <id>`. The
 * ```
 *
 * `id` can be whatever the [details of a group][details_of_a_group] api takes for
 *
 * its `:id` value, so for example:
 *
 * ```sh
 * $ pulumi import gitlab:index/group:Group example example
 * ```
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * A list of email address domains to allow group access. Will be concatenated together into a comma separated string.
     */
    public readonly allowedEmailDomainsLists!: pulumi.Output<string[]>;
    /**
     * Default to Auto DevOps pipeline for all projects within this group.
     */
    public readonly autoDevopsEnabled!: pulumi.Output<boolean>;
    /**
     * A local path to the avatar image to upload. **Note**: not available for imported resources.
     */
    public readonly avatar!: pulumi.Output<string | undefined>;
    /**
     * The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
     */
    public readonly avatarHash!: pulumi.Output<string>;
    /**
     * The URL of the avatar image.
     */
    public /*out*/ readonly avatarUrl!: pulumi.Output<string>;
    /**
     * Initial default branch name.
     */
    public readonly defaultBranch!: pulumi.Output<string | undefined>;
    /**
     * See https://docs.gitlab.com/api/groups/#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
     *
     * @deprecated Deprecated in GitLab 17.0, due for removal in v5 of the API. Use defaultBranchProtectionDefaults instead.
     */
    public readonly defaultBranchProtection!: pulumi.Output<number>;
    /**
     * The default branch protection defaults
     */
    public readonly defaultBranchProtectionDefaults!: pulumi.Output<outputs.GroupDefaultBranchProtectionDefaults>;
    /**
     * The group's description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Enable email notifications.
     */
    public readonly emailsEnabled!: pulumi.Output<boolean>;
    /**
     * Can be set by administrators only. Additional CI/CD minutes for this group.
     */
    public readonly extraSharedRunnersMinutesLimit!: pulumi.Output<number>;
    /**
     * The full name of the group.
     */
    public /*out*/ readonly fullName!: pulumi.Output<string>;
    /**
     * The full path of the group.
     */
    public /*out*/ readonly fullPath!: pulumi.Output<string>;
    /**
     * A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
     */
    public readonly ipRestrictionRanges!: pulumi.Output<string[] | undefined>;
    /**
     * Enable/disable Large File Storage (LFS) for the projects in this group.
     */
    public readonly lfsEnabled!: pulumi.Output<boolean>;
    /**
     * Users cannot be added to projects in this group.
     */
    public readonly membershipLock!: pulumi.Output<boolean>;
    /**
     * Disable the capability of a group from getting mentioned.
     */
    public readonly mentionsDisabled!: pulumi.Output<boolean>;
    /**
     * The name of the group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Id of the parent group (creates a nested group).
     */
    public readonly parentId!: pulumi.Output<number>;
    /**
     * The path of the group.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
     */
    public readonly permanentlyRemoveOnDelete!: pulumi.Output<boolean | undefined>;
    /**
     * Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
     */
    public readonly preventForkingOutsideGroup!: pulumi.Output<boolean>;
    /**
     * Determine if developers can create projects in the group. Valid values are: `noone`, `owner`, `maintainer`, `developer`
     */
    public readonly projectCreationLevel!: pulumi.Output<string>;
    /**
     * Push rules for the group.
     */
    public readonly pushRules!: pulumi.Output<outputs.GroupPushRules>;
    /**
     * Allow users to request member access.
     */
    public readonly requestAccessEnabled!: pulumi.Output<boolean>;
    /**
     * Require all users in this group to setup Two-factor authentication.
     */
    public readonly requireTwoFactorAuthentication!: pulumi.Output<boolean>;
    /**
     * The group level registration token to use during runner setup.
     */
    public /*out*/ readonly runnersToken!: pulumi.Output<string>;
    /**
     * Prevent sharing a project with another group within this group.
     */
    public readonly shareWithGroupLock!: pulumi.Output<boolean>;
    /**
     * Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
     */
    public readonly sharedRunnersMinutesLimit!: pulumi.Output<number>;
    /**
     * Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabledAndOverridable`, `disabledAndUnoverridable`, `disabledWithOverride`.
     */
    public readonly sharedRunnersSetting!: pulumi.Output<string>;
    /**
     * Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
     */
    public readonly subgroupCreationLevel!: pulumi.Output<string>;
    /**
     * Defaults to 48. Time before Two-factor authentication is enforced (in hours).
     */
    public readonly twoFactorGracePeriod!: pulumi.Output<number>;
    /**
     * The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
     */
    public readonly visibilityLevel!: pulumi.Output<string>;
    /**
     * Web URL of the group.
     */
    public /*out*/ readonly webUrl!: pulumi.Output<string>;
    /**
     * The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
     */
    public readonly wikiAccessLevel!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["allowedEmailDomainsLists"] = state ? state.allowedEmailDomainsLists : undefined;
            resourceInputs["autoDevopsEnabled"] = state ? state.autoDevopsEnabled : undefined;
            resourceInputs["avatar"] = state ? state.avatar : undefined;
            resourceInputs["avatarHash"] = state ? state.avatarHash : undefined;
            resourceInputs["avatarUrl"] = state ? state.avatarUrl : undefined;
            resourceInputs["defaultBranch"] = state ? state.defaultBranch : undefined;
            resourceInputs["defaultBranchProtection"] = state ? state.defaultBranchProtection : undefined;
            resourceInputs["defaultBranchProtectionDefaults"] = state ? state.defaultBranchProtectionDefaults : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["emailsEnabled"] = state ? state.emailsEnabled : undefined;
            resourceInputs["extraSharedRunnersMinutesLimit"] = state ? state.extraSharedRunnersMinutesLimit : undefined;
            resourceInputs["fullName"] = state ? state.fullName : undefined;
            resourceInputs["fullPath"] = state ? state.fullPath : undefined;
            resourceInputs["ipRestrictionRanges"] = state ? state.ipRestrictionRanges : undefined;
            resourceInputs["lfsEnabled"] = state ? state.lfsEnabled : undefined;
            resourceInputs["membershipLock"] = state ? state.membershipLock : undefined;
            resourceInputs["mentionsDisabled"] = state ? state.mentionsDisabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentId"] = state ? state.parentId : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["permanentlyRemoveOnDelete"] = state ? state.permanentlyRemoveOnDelete : undefined;
            resourceInputs["preventForkingOutsideGroup"] = state ? state.preventForkingOutsideGroup : undefined;
            resourceInputs["projectCreationLevel"] = state ? state.projectCreationLevel : undefined;
            resourceInputs["pushRules"] = state ? state.pushRules : undefined;
            resourceInputs["requestAccessEnabled"] = state ? state.requestAccessEnabled : undefined;
            resourceInputs["requireTwoFactorAuthentication"] = state ? state.requireTwoFactorAuthentication : undefined;
            resourceInputs["runnersToken"] = state ? state.runnersToken : undefined;
            resourceInputs["shareWithGroupLock"] = state ? state.shareWithGroupLock : undefined;
            resourceInputs["sharedRunnersMinutesLimit"] = state ? state.sharedRunnersMinutesLimit : undefined;
            resourceInputs["sharedRunnersSetting"] = state ? state.sharedRunnersSetting : undefined;
            resourceInputs["subgroupCreationLevel"] = state ? state.subgroupCreationLevel : undefined;
            resourceInputs["twoFactorGracePeriod"] = state ? state.twoFactorGracePeriod : undefined;
            resourceInputs["visibilityLevel"] = state ? state.visibilityLevel : undefined;
            resourceInputs["webUrl"] = state ? state.webUrl : undefined;
            resourceInputs["wikiAccessLevel"] = state ? state.wikiAccessLevel : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["allowedEmailDomainsLists"] = args ? args.allowedEmailDomainsLists : undefined;
            resourceInputs["autoDevopsEnabled"] = args ? args.autoDevopsEnabled : undefined;
            resourceInputs["avatar"] = args ? args.avatar : undefined;
            resourceInputs["avatarHash"] = args ? args.avatarHash : undefined;
            resourceInputs["defaultBranch"] = args ? args.defaultBranch : undefined;
            resourceInputs["defaultBranchProtection"] = args ? args.defaultBranchProtection : undefined;
            resourceInputs["defaultBranchProtectionDefaults"] = args ? args.defaultBranchProtectionDefaults : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["emailsEnabled"] = args ? args.emailsEnabled : undefined;
            resourceInputs["extraSharedRunnersMinutesLimit"] = args ? args.extraSharedRunnersMinutesLimit : undefined;
            resourceInputs["ipRestrictionRanges"] = args ? args.ipRestrictionRanges : undefined;
            resourceInputs["lfsEnabled"] = args ? args.lfsEnabled : undefined;
            resourceInputs["membershipLock"] = args ? args.membershipLock : undefined;
            resourceInputs["mentionsDisabled"] = args ? args.mentionsDisabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentId"] = args ? args.parentId : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["permanentlyRemoveOnDelete"] = args ? args.permanentlyRemoveOnDelete : undefined;
            resourceInputs["preventForkingOutsideGroup"] = args ? args.preventForkingOutsideGroup : undefined;
            resourceInputs["projectCreationLevel"] = args ? args.projectCreationLevel : undefined;
            resourceInputs["pushRules"] = args ? args.pushRules : undefined;
            resourceInputs["requestAccessEnabled"] = args ? args.requestAccessEnabled : undefined;
            resourceInputs["requireTwoFactorAuthentication"] = args ? args.requireTwoFactorAuthentication : undefined;
            resourceInputs["shareWithGroupLock"] = args ? args.shareWithGroupLock : undefined;
            resourceInputs["sharedRunnersMinutesLimit"] = args ? args.sharedRunnersMinutesLimit : undefined;
            resourceInputs["sharedRunnersSetting"] = args ? args.sharedRunnersSetting : undefined;
            resourceInputs["subgroupCreationLevel"] = args ? args.subgroupCreationLevel : undefined;
            resourceInputs["twoFactorGracePeriod"] = args ? args.twoFactorGracePeriod : undefined;
            resourceInputs["visibilityLevel"] = args ? args.visibilityLevel : undefined;
            resourceInputs["wikiAccessLevel"] = args ? args.wikiAccessLevel : undefined;
            resourceInputs["avatarUrl"] = undefined /*out*/;
            resourceInputs["fullName"] = undefined /*out*/;
            resourceInputs["fullPath"] = undefined /*out*/;
            resourceInputs["runnersToken"] = undefined /*out*/;
            resourceInputs["webUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["runnersToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * A list of email address domains to allow group access. Will be concatenated together into a comma separated string.
     */
    allowedEmailDomainsLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default to Auto DevOps pipeline for all projects within this group.
     */
    autoDevopsEnabled?: pulumi.Input<boolean>;
    /**
     * A local path to the avatar image to upload. **Note**: not available for imported resources.
     */
    avatar?: pulumi.Input<string>;
    /**
     * The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
     */
    avatarHash?: pulumi.Input<string>;
    /**
     * The URL of the avatar image.
     */
    avatarUrl?: pulumi.Input<string>;
    /**
     * Initial default branch name.
     */
    defaultBranch?: pulumi.Input<string>;
    /**
     * See https://docs.gitlab.com/api/groups/#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
     *
     * @deprecated Deprecated in GitLab 17.0, due for removal in v5 of the API. Use defaultBranchProtectionDefaults instead.
     */
    defaultBranchProtection?: pulumi.Input<number>;
    /**
     * The default branch protection defaults
     */
    defaultBranchProtectionDefaults?: pulumi.Input<inputs.GroupDefaultBranchProtectionDefaults>;
    /**
     * The group's description.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable email notifications.
     */
    emailsEnabled?: pulumi.Input<boolean>;
    /**
     * Can be set by administrators only. Additional CI/CD minutes for this group.
     */
    extraSharedRunnersMinutesLimit?: pulumi.Input<number>;
    /**
     * The full name of the group.
     */
    fullName?: pulumi.Input<string>;
    /**
     * The full path of the group.
     */
    fullPath?: pulumi.Input<string>;
    /**
     * A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
     */
    ipRestrictionRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable/disable Large File Storage (LFS) for the projects in this group.
     */
    lfsEnabled?: pulumi.Input<boolean>;
    /**
     * Users cannot be added to projects in this group.
     */
    membershipLock?: pulumi.Input<boolean>;
    /**
     * Disable the capability of a group from getting mentioned.
     */
    mentionsDisabled?: pulumi.Input<boolean>;
    /**
     * The name of the group.
     */
    name?: pulumi.Input<string>;
    /**
     * Id of the parent group (creates a nested group).
     */
    parentId?: pulumi.Input<number>;
    /**
     * The path of the group.
     */
    path?: pulumi.Input<string>;
    /**
     * Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
     */
    permanentlyRemoveOnDelete?: pulumi.Input<boolean>;
    /**
     * Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
     */
    preventForkingOutsideGroup?: pulumi.Input<boolean>;
    /**
     * Determine if developers can create projects in the group. Valid values are: `noone`, `owner`, `maintainer`, `developer`
     */
    projectCreationLevel?: pulumi.Input<string>;
    /**
     * Push rules for the group.
     */
    pushRules?: pulumi.Input<inputs.GroupPushRules>;
    /**
     * Allow users to request member access.
     */
    requestAccessEnabled?: pulumi.Input<boolean>;
    /**
     * Require all users in this group to setup Two-factor authentication.
     */
    requireTwoFactorAuthentication?: pulumi.Input<boolean>;
    /**
     * The group level registration token to use during runner setup.
     */
    runnersToken?: pulumi.Input<string>;
    /**
     * Prevent sharing a project with another group within this group.
     */
    shareWithGroupLock?: pulumi.Input<boolean>;
    /**
     * Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
     */
    sharedRunnersMinutesLimit?: pulumi.Input<number>;
    /**
     * Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabledAndOverridable`, `disabledAndUnoverridable`, `disabledWithOverride`.
     */
    sharedRunnersSetting?: pulumi.Input<string>;
    /**
     * Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
     */
    subgroupCreationLevel?: pulumi.Input<string>;
    /**
     * Defaults to 48. Time before Two-factor authentication is enforced (in hours).
     */
    twoFactorGracePeriod?: pulumi.Input<number>;
    /**
     * The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
     */
    visibilityLevel?: pulumi.Input<string>;
    /**
     * Web URL of the group.
     */
    webUrl?: pulumi.Input<string>;
    /**
     * The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
     */
    wikiAccessLevel?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * A list of email address domains to allow group access. Will be concatenated together into a comma separated string.
     */
    allowedEmailDomainsLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default to Auto DevOps pipeline for all projects within this group.
     */
    autoDevopsEnabled?: pulumi.Input<boolean>;
    /**
     * A local path to the avatar image to upload. **Note**: not available for imported resources.
     */
    avatar?: pulumi.Input<string>;
    /**
     * The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
     */
    avatarHash?: pulumi.Input<string>;
    /**
     * Initial default branch name.
     */
    defaultBranch?: pulumi.Input<string>;
    /**
     * See https://docs.gitlab.com/api/groups/#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
     *
     * @deprecated Deprecated in GitLab 17.0, due for removal in v5 of the API. Use defaultBranchProtectionDefaults instead.
     */
    defaultBranchProtection?: pulumi.Input<number>;
    /**
     * The default branch protection defaults
     */
    defaultBranchProtectionDefaults?: pulumi.Input<inputs.GroupDefaultBranchProtectionDefaults>;
    /**
     * The group's description.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable email notifications.
     */
    emailsEnabled?: pulumi.Input<boolean>;
    /**
     * Can be set by administrators only. Additional CI/CD minutes for this group.
     */
    extraSharedRunnersMinutesLimit?: pulumi.Input<number>;
    /**
     * A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
     */
    ipRestrictionRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable/disable Large File Storage (LFS) for the projects in this group.
     */
    lfsEnabled?: pulumi.Input<boolean>;
    /**
     * Users cannot be added to projects in this group.
     */
    membershipLock?: pulumi.Input<boolean>;
    /**
     * Disable the capability of a group from getting mentioned.
     */
    mentionsDisabled?: pulumi.Input<boolean>;
    /**
     * The name of the group.
     */
    name?: pulumi.Input<string>;
    /**
     * Id of the parent group (creates a nested group).
     */
    parentId?: pulumi.Input<number>;
    /**
     * The path of the group.
     */
    path: pulumi.Input<string>;
    /**
     * Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
     */
    permanentlyRemoveOnDelete?: pulumi.Input<boolean>;
    /**
     * Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
     */
    preventForkingOutsideGroup?: pulumi.Input<boolean>;
    /**
     * Determine if developers can create projects in the group. Valid values are: `noone`, `owner`, `maintainer`, `developer`
     */
    projectCreationLevel?: pulumi.Input<string>;
    /**
     * Push rules for the group.
     */
    pushRules?: pulumi.Input<inputs.GroupPushRules>;
    /**
     * Allow users to request member access.
     */
    requestAccessEnabled?: pulumi.Input<boolean>;
    /**
     * Require all users in this group to setup Two-factor authentication.
     */
    requireTwoFactorAuthentication?: pulumi.Input<boolean>;
    /**
     * Prevent sharing a project with another group within this group.
     */
    shareWithGroupLock?: pulumi.Input<boolean>;
    /**
     * Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
     */
    sharedRunnersMinutesLimit?: pulumi.Input<number>;
    /**
     * Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabledAndOverridable`, `disabledAndUnoverridable`, `disabledWithOverride`.
     */
    sharedRunnersSetting?: pulumi.Input<string>;
    /**
     * Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
     */
    subgroupCreationLevel?: pulumi.Input<string>;
    /**
     * Defaults to 48. Time before Two-factor authentication is enforced (in hours).
     */
    twoFactorGracePeriod?: pulumi.Input<number>;
    /**
     * The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
     */
    visibilityLevel?: pulumi.Input<string>;
    /**
     * The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
     */
    wikiAccessLevel?: pulumi.Input<string>;
}
