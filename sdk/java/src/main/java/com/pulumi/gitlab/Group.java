// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupState;
import com.pulumi.gitlab.outputs.GroupDefaultBranchProtectionDefaults;
import com.pulumi.gitlab.outputs.GroupPushRules;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.Group` resource allows to manage the lifecycle of a group.
 * 
 * &gt; On GitLab SaaS, you must use the GitLab UI to create groups without a parent group. You cannot use this provider nor the API to do this.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.Group;
 * import com.pulumi.gitlab.GroupArgs;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectArgs;
 * import com.pulumi.gitlab.inputs.GroupPushRulesArgs;
 * import com.pulumi.gitlab.inputs.GroupDefaultBranchProtectionDefaultsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var example = new Group("example", GroupArgs.builder()
 *             .name("example")
 *             .path("example")
 *             .description("An example group")
 *             .build());
 * 
 *         // Create a project in the example group
 *         var exampleProject = new Project("exampleProject", ProjectArgs.builder()
 *             .name("example")
 *             .description("An example project")
 *             .namespaceId(example.id())
 *             .build());
 * 
 *         // Group with custom push rules
 *         var example_two = new Group("example-two", GroupArgs.builder()
 *             .name("example-two")
 *             .path("example-two")
 *             .description("An example group with push rules")
 *             .pushRules(GroupPushRulesArgs.builder()
 *                 .authorEmailRegex("}{@literal @}{@code example\\.com$")
 *                 .commitCommitterCheck(true)
 *                 .memberCheck(true)
 *                 .preventSecrets(true)
 *                 .build())
 *             .build());
 * 
 *         // Group with custom default branch protection defaults
 *         var example_three = new Group("example-three", GroupArgs.builder()
 *             .name("example-three")
 *             .path("example-three")
 *             .description("An example group with default branch protection defaults")
 *             .defaultBranchProtectionDefaults(GroupDefaultBranchProtectionDefaultsArgs.builder()
 *                 .allowedToPushes("developer")
 *                 .allowForcePush(true)
 *                 .allowedToMerges(                
 *                     "developer",
 *                     "maintainer")
 *                 .developerCanInitialPush(true)
 *                 .build())
 *             .build());
 * 
 *         // Group with custom default branch protection defaults
 *         var example_four = new Group("example-four", GroupArgs.builder()
 *             .name("example-four")
 *             .path("example-four")
 *             .description("An example group with default branch protection defaults")
 *             .defaultBranchProtectionDefaults(GroupDefaultBranchProtectionDefaultsArgs.builder()
 *                 .allowedToPushes("no one")
 *                 .allowForcePush(true)
 *                 .allowedToMerges("no one")
 *                 .developerCanInitialPush(true)
 *                 .build())
 *             .build());
 * 
 *         // Group with a default branch name specified
 *         var example_five = new Group("example-five", GroupArgs.builder()
 *             .name("example")
 *             .path("example")
 *             .defaultBranch("develop")
 *             .description("An example group with a default branch name")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_group.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * ```sh
 * $ pulumi import gitlab:index/group:Group You can import a group state using `&lt;resource&gt; &lt;id&gt;`. The
 * ```
 * 
 * `id` can be whatever the [details of a group][details_of_a_group] api takes for
 * 
 * its `:id` value, so for example:
 * 
 * ```sh
 * $ pulumi import gitlab:index/group:Group example example
 * ```
 * 
 */
@ResourceType(type="gitlab:index/group:Group")
public class Group extends com.pulumi.resources.CustomResource {
    /**
     * A list of email address domains to allow group access. Will be concatenated together into a comma separated string.
     * 
     */
    @Export(name="allowedEmailDomainsLists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedEmailDomainsLists;

    /**
     * @return A list of email address domains to allow group access. Will be concatenated together into a comma separated string.
     * 
     */
    public Output<Optional<List<String>>> allowedEmailDomainsLists() {
        return Codegen.optional(this.allowedEmailDomainsLists);
    }
    /**
     * Default to Auto DevOps pipeline for all projects within this group.
     * 
     */
    @Export(name="autoDevopsEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> autoDevopsEnabled;

    /**
     * @return Default to Auto DevOps pipeline for all projects within this group.
     * 
     */
    public Output<Boolean> autoDevopsEnabled() {
        return this.autoDevopsEnabled;
    }
    /**
     * A local path to the avatar image to upload. **Note**: not available for imported resources.
     * 
     */
    @Export(name="avatar", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> avatar;

    /**
     * @return A local path to the avatar image to upload. **Note**: not available for imported resources.
     * 
     */
    public Output<Optional<String>> avatar() {
        return Codegen.optional(this.avatar);
    }
    /**
     * The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
     * 
     */
    @Export(name="avatarHash", refs={String.class}, tree="[0]")
    private Output<String> avatarHash;

    /**
     * @return The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
     * 
     */
    public Output<String> avatarHash() {
        return this.avatarHash;
    }
    /**
     * The URL of the avatar image.
     * 
     */
    @Export(name="avatarUrl", refs={String.class}, tree="[0]")
    private Output<String> avatarUrl;

    /**
     * @return The URL of the avatar image.
     * 
     */
    public Output<String> avatarUrl() {
        return this.avatarUrl;
    }
    /**
     * Initial default branch name.
     * 
     */
    @Export(name="defaultBranch", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultBranch;

    /**
     * @return Initial default branch name.
     * 
     */
    public Output<Optional<String>> defaultBranch() {
        return Codegen.optional(this.defaultBranch);
    }
    /**
     * See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
     * 
     * @deprecated
     * Deprecated in GitLab 17.0. Use default_branch_protection_defaults instead.
     * 
     */
    @Deprecated /* Deprecated in GitLab 17.0. Use default_branch_protection_defaults instead. */
    @Export(name="defaultBranchProtection", refs={Integer.class}, tree="[0]")
    private Output<Integer> defaultBranchProtection;

    /**
     * @return See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
     * 
     */
    public Output<Integer> defaultBranchProtection() {
        return this.defaultBranchProtection;
    }
    /**
     * The default branch protection defaults
     * 
     */
    @Export(name="defaultBranchProtectionDefaults", refs={GroupDefaultBranchProtectionDefaults.class}, tree="[0]")
    private Output<GroupDefaultBranchProtectionDefaults> defaultBranchProtectionDefaults;

    /**
     * @return The default branch protection defaults
     * 
     */
    public Output<GroupDefaultBranchProtectionDefaults> defaultBranchProtectionDefaults() {
        return this.defaultBranchProtectionDefaults;
    }
    /**
     * The group&#39;s description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The group&#39;s description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Enable email notifications.
     * 
     */
    @Export(name="emailsEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> emailsEnabled;

    /**
     * @return Enable email notifications.
     * 
     */
    public Output<Boolean> emailsEnabled() {
        return this.emailsEnabled;
    }
    /**
     * Can be set by administrators only. Additional CI/CD minutes for this group.
     * 
     */
    @Export(name="extraSharedRunnersMinutesLimit", refs={Integer.class}, tree="[0]")
    private Output<Integer> extraSharedRunnersMinutesLimit;

    /**
     * @return Can be set by administrators only. Additional CI/CD minutes for this group.
     * 
     */
    public Output<Integer> extraSharedRunnersMinutesLimit() {
        return this.extraSharedRunnersMinutesLimit;
    }
    /**
     * The full name of the group.
     * 
     */
    @Export(name="fullName", refs={String.class}, tree="[0]")
    private Output<String> fullName;

    /**
     * @return The full name of the group.
     * 
     */
    public Output<String> fullName() {
        return this.fullName;
    }
    /**
     * The full path of the group.
     * 
     */
    @Export(name="fullPath", refs={String.class}, tree="[0]")
    private Output<String> fullPath;

    /**
     * @return The full path of the group.
     * 
     */
    public Output<String> fullPath() {
        return this.fullPath;
    }
    /**
     * A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
     * 
     */
    @Export(name="ipRestrictionRanges", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> ipRestrictionRanges;

    /**
     * @return A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
     * 
     */
    public Output<Optional<List<String>>> ipRestrictionRanges() {
        return Codegen.optional(this.ipRestrictionRanges);
    }
    /**
     * Enable/disable Large File Storage (LFS) for the projects in this group.
     * 
     */
    @Export(name="lfsEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> lfsEnabled;

    /**
     * @return Enable/disable Large File Storage (LFS) for the projects in this group.
     * 
     */
    public Output<Boolean> lfsEnabled() {
        return this.lfsEnabled;
    }
    /**
     * Users cannot be added to projects in this group.
     * 
     */
    @Export(name="membershipLock", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> membershipLock;

    /**
     * @return Users cannot be added to projects in this group.
     * 
     */
    public Output<Optional<Boolean>> membershipLock() {
        return Codegen.optional(this.membershipLock);
    }
    /**
     * Disable the capability of a group from getting mentioned.
     * 
     */
    @Export(name="mentionsDisabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> mentionsDisabled;

    /**
     * @return Disable the capability of a group from getting mentioned.
     * 
     */
    public Output<Boolean> mentionsDisabled() {
        return this.mentionsDisabled;
    }
    /**
     * The name of the group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Id of the parent group (creates a nested group).
     * 
     */
    @Export(name="parentId", refs={Integer.class}, tree="[0]")
    private Output<Integer> parentId;

    /**
     * @return Id of the parent group (creates a nested group).
     * 
     */
    public Output<Integer> parentId() {
        return this.parentId;
    }
    /**
     * The path of the group.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The path of the group.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
     * 
     */
    @Export(name="permanentlyRemoveOnDelete", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> permanentlyRemoveOnDelete;

    /**
     * @return Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
     * 
     */
    public Output<Optional<Boolean>> permanentlyRemoveOnDelete() {
        return Codegen.optional(this.permanentlyRemoveOnDelete);
    }
    /**
     * Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
     * 
     */
    @Export(name="preventForkingOutsideGroup", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> preventForkingOutsideGroup;

    /**
     * @return Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
     * 
     */
    public Output<Boolean> preventForkingOutsideGroup() {
        return this.preventForkingOutsideGroup;
    }
    /**
     * Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
     * 
     */
    @Export(name="projectCreationLevel", refs={String.class}, tree="[0]")
    private Output<String> projectCreationLevel;

    /**
     * @return Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
     * 
     */
    public Output<String> projectCreationLevel() {
        return this.projectCreationLevel;
    }
    /**
     * Push rules for the group.
     * 
     */
    @Export(name="pushRules", refs={GroupPushRules.class}, tree="[0]")
    private Output<GroupPushRules> pushRules;

    /**
     * @return Push rules for the group.
     * 
     */
    public Output<GroupPushRules> pushRules() {
        return this.pushRules;
    }
    /**
     * Allow users to request member access.
     * 
     */
    @Export(name="requestAccessEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> requestAccessEnabled;

    /**
     * @return Allow users to request member access.
     * 
     */
    public Output<Boolean> requestAccessEnabled() {
        return this.requestAccessEnabled;
    }
    /**
     * Require all users in this group to setup Two-factor authentication.
     * 
     */
    @Export(name="requireTwoFactorAuthentication", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> requireTwoFactorAuthentication;

    /**
     * @return Require all users in this group to setup Two-factor authentication.
     * 
     */
    public Output<Boolean> requireTwoFactorAuthentication() {
        return this.requireTwoFactorAuthentication;
    }
    /**
     * The group level registration token to use during runner setup.
     * 
     */
    @Export(name="runnersToken", refs={String.class}, tree="[0]")
    private Output<String> runnersToken;

    /**
     * @return The group level registration token to use during runner setup.
     * 
     */
    public Output<String> runnersToken() {
        return this.runnersToken;
    }
    /**
     * Prevent sharing a project with another group within this group.
     * 
     */
    @Export(name="shareWithGroupLock", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> shareWithGroupLock;

    /**
     * @return Prevent sharing a project with another group within this group.
     * 
     */
    public Output<Boolean> shareWithGroupLock() {
        return this.shareWithGroupLock;
    }
    /**
     * Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or &gt; 0.
     * 
     */
    @Export(name="sharedRunnersMinutesLimit", refs={Integer.class}, tree="[0]")
    private Output<Integer> sharedRunnersMinutesLimit;

    /**
     * @return Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or &gt; 0.
     * 
     */
    public Output<Integer> sharedRunnersMinutesLimit() {
        return this.sharedRunnersMinutesLimit;
    }
    /**
     * Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
     * 
     */
    @Export(name="sharedRunnersSetting", refs={String.class}, tree="[0]")
    private Output<String> sharedRunnersSetting;

    /**
     * @return Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
     * 
     */
    public Output<String> sharedRunnersSetting() {
        return this.sharedRunnersSetting;
    }
    /**
     * Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
     * 
     */
    @Export(name="subgroupCreationLevel", refs={String.class}, tree="[0]")
    private Output<String> subgroupCreationLevel;

    /**
     * @return Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
     * 
     */
    public Output<String> subgroupCreationLevel() {
        return this.subgroupCreationLevel;
    }
    /**
     * Defaults to 48. Time before Two-factor authentication is enforced (in hours).
     * 
     */
    @Export(name="twoFactorGracePeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> twoFactorGracePeriod;

    /**
     * @return Defaults to 48. Time before Two-factor authentication is enforced (in hours).
     * 
     */
    public Output<Integer> twoFactorGracePeriod() {
        return this.twoFactorGracePeriod;
    }
    /**
     * The group&#39;s visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
     * 
     */
    @Export(name="visibilityLevel", refs={String.class}, tree="[0]")
    private Output<String> visibilityLevel;

    /**
     * @return The group&#39;s visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
     * 
     */
    public Output<String> visibilityLevel() {
        return this.visibilityLevel;
    }
    /**
     * Web URL of the group.
     * 
     */
    @Export(name="webUrl", refs={String.class}, tree="[0]")
    private Output<String> webUrl;

    /**
     * @return Web URL of the group.
     * 
     */
    public Output<String> webUrl() {
        return this.webUrl;
    }
    /**
     * The group&#39;s wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
     * 
     */
    @Export(name="wikiAccessLevel", refs={String.class}, tree="[0]")
    private Output<String> wikiAccessLevel;

    /**
     * @return The group&#39;s wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
     * 
     */
    public Output<String> wikiAccessLevel() {
        return this.wikiAccessLevel;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Group(java.lang.String name) {
        this(name, GroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Group(java.lang.String name, GroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Group(java.lang.String name, GroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/group:Group", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Group(java.lang.String name, Output<java.lang.String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/group:Group", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupArgs makeArgs(GroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "runnersToken"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Group get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Group(name, id, state, options);
    }
}
