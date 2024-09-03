// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.inputs.GroupPushRulesArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupArgs Empty = new GroupArgs();

    /**
     * Default to Auto DevOps pipeline for all projects within this group.
     * 
     */
    @Import(name="autoDevopsEnabled")
    private @Nullable Output<Boolean> autoDevopsEnabled;

    /**
     * @return Default to Auto DevOps pipeline for all projects within this group.
     * 
     */
    public Optional<Output<Boolean>> autoDevopsEnabled() {
        return Optional.ofNullable(this.autoDevopsEnabled);
    }

    /**
     * A local path to the avatar image to upload. **Note**: not available for imported resources.
     * 
     */
    @Import(name="avatar")
    private @Nullable Output<String> avatar;

    /**
     * @return A local path to the avatar image to upload. **Note**: not available for imported resources.
     * 
     */
    public Optional<Output<String>> avatar() {
        return Optional.ofNullable(this.avatar);
    }

    /**
     * The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
     * 
     */
    @Import(name="avatarHash")
    private @Nullable Output<String> avatarHash;

    /**
     * @return The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
     * 
     */
    public Optional<Output<String>> avatarHash() {
        return Optional.ofNullable(this.avatarHash);
    }

    /**
     * See [https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection](https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection). Valid values are: `0`, `1`, `2`, `3`, `4`.
     * 
     */
    @Import(name="defaultBranchProtection")
    private @Nullable Output<Integer> defaultBranchProtection;

    /**
     * @return See [https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection](https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection). Valid values are: `0`, `1`, `2`, `3`, `4`.
     * 
     */
    public Optional<Output<Integer>> defaultBranchProtection() {
        return Optional.ofNullable(this.defaultBranchProtection);
    }

    /**
     * The group&#39;s description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The group&#39;s description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Enable email notifications.
     * 
     */
    @Import(name="emailsEnabled")
    private @Nullable Output<Boolean> emailsEnabled;

    /**
     * @return Enable email notifications.
     * 
     */
    public Optional<Output<Boolean>> emailsEnabled() {
        return Optional.ofNullable(this.emailsEnabled);
    }

    /**
     * Can be set by administrators only. Additional CI/CD minutes for this group.
     * 
     */
    @Import(name="extraSharedRunnersMinutesLimit")
    private @Nullable Output<Integer> extraSharedRunnersMinutesLimit;

    /**
     * @return Can be set by administrators only. Additional CI/CD minutes for this group.
     * 
     */
    public Optional<Output<Integer>> extraSharedRunnersMinutesLimit() {
        return Optional.ofNullable(this.extraSharedRunnersMinutesLimit);
    }

    /**
     * A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
     * 
     */
    @Import(name="ipRestrictionRanges")
    private @Nullable Output<List<String>> ipRestrictionRanges;

    /**
     * @return A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
     * 
     */
    public Optional<Output<List<String>>> ipRestrictionRanges() {
        return Optional.ofNullable(this.ipRestrictionRanges);
    }

    /**
     * Enable/disable Large File Storage (LFS) for the projects in this group.
     * 
     */
    @Import(name="lfsEnabled")
    private @Nullable Output<Boolean> lfsEnabled;

    /**
     * @return Enable/disable Large File Storage (LFS) for the projects in this group.
     * 
     */
    public Optional<Output<Boolean>> lfsEnabled() {
        return Optional.ofNullable(this.lfsEnabled);
    }

    /**
     * Users cannot be added to projects in this group.
     * 
     */
    @Import(name="membershipLock")
    private @Nullable Output<Boolean> membershipLock;

    /**
     * @return Users cannot be added to projects in this group.
     * 
     */
    public Optional<Output<Boolean>> membershipLock() {
        return Optional.ofNullable(this.membershipLock);
    }

    /**
     * Disable the capability of a group from getting mentioned.
     * 
     */
    @Import(name="mentionsDisabled")
    private @Nullable Output<Boolean> mentionsDisabled;

    /**
     * @return Disable the capability of a group from getting mentioned.
     * 
     */
    public Optional<Output<Boolean>> mentionsDisabled() {
        return Optional.ofNullable(this.mentionsDisabled);
    }

    /**
     * The name of the group.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the group.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Id of the parent group (creates a nested group).
     * 
     */
    @Import(name="parentId")
    private @Nullable Output<Integer> parentId;

    /**
     * @return Id of the parent group (creates a nested group).
     * 
     */
    public Optional<Output<Integer>> parentId() {
        return Optional.ofNullable(this.parentId);
    }

    /**
     * The path of the group.
     * 
     */
    @Import(name="path", required=true)
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
    @Import(name="permanentlyRemoveOnDelete")
    private @Nullable Output<Boolean> permanentlyRemoveOnDelete;

    /**
     * @return Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
     * 
     */
    public Optional<Output<Boolean>> permanentlyRemoveOnDelete() {
        return Optional.ofNullable(this.permanentlyRemoveOnDelete);
    }

    /**
     * Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
     * 
     */
    @Import(name="preventForkingOutsideGroup")
    private @Nullable Output<Boolean> preventForkingOutsideGroup;

    /**
     * @return Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
     * 
     */
    public Optional<Output<Boolean>> preventForkingOutsideGroup() {
        return Optional.ofNullable(this.preventForkingOutsideGroup);
    }

    /**
     * Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
     * 
     */
    @Import(name="projectCreationLevel")
    private @Nullable Output<String> projectCreationLevel;

    /**
     * @return Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
     * 
     */
    public Optional<Output<String>> projectCreationLevel() {
        return Optional.ofNullable(this.projectCreationLevel);
    }

    /**
     * Push rules for the group.
     * 
     */
    @Import(name="pushRules")
    private @Nullable Output<GroupPushRulesArgs> pushRules;

    /**
     * @return Push rules for the group.
     * 
     */
    public Optional<Output<GroupPushRulesArgs>> pushRules() {
        return Optional.ofNullable(this.pushRules);
    }

    /**
     * Allow users to request member access.
     * 
     */
    @Import(name="requestAccessEnabled")
    private @Nullable Output<Boolean> requestAccessEnabled;

    /**
     * @return Allow users to request member access.
     * 
     */
    public Optional<Output<Boolean>> requestAccessEnabled() {
        return Optional.ofNullable(this.requestAccessEnabled);
    }

    /**
     * Require all users in this group to setup Two-factor authentication.
     * 
     */
    @Import(name="requireTwoFactorAuthentication")
    private @Nullable Output<Boolean> requireTwoFactorAuthentication;

    /**
     * @return Require all users in this group to setup Two-factor authentication.
     * 
     */
    public Optional<Output<Boolean>> requireTwoFactorAuthentication() {
        return Optional.ofNullable(this.requireTwoFactorAuthentication);
    }

    /**
     * Prevent sharing a project with another group within this group.
     * 
     */
    @Import(name="shareWithGroupLock")
    private @Nullable Output<Boolean> shareWithGroupLock;

    /**
     * @return Prevent sharing a project with another group within this group.
     * 
     */
    public Optional<Output<Boolean>> shareWithGroupLock() {
        return Optional.ofNullable(this.shareWithGroupLock);
    }

    /**
     * Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or &gt; 0.
     * 
     */
    @Import(name="sharedRunnersMinutesLimit")
    private @Nullable Output<Integer> sharedRunnersMinutesLimit;

    /**
     * @return Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or &gt; 0.
     * 
     */
    public Optional<Output<Integer>> sharedRunnersMinutesLimit() {
        return Optional.ofNullable(this.sharedRunnersMinutesLimit);
    }

    /**
     * Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
     * 
     */
    @Import(name="sharedRunnersSetting")
    private @Nullable Output<String> sharedRunnersSetting;

    /**
     * @return Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
     * 
     */
    public Optional<Output<String>> sharedRunnersSetting() {
        return Optional.ofNullable(this.sharedRunnersSetting);
    }

    /**
     * Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
     * 
     */
    @Import(name="subgroupCreationLevel")
    private @Nullable Output<String> subgroupCreationLevel;

    /**
     * @return Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
     * 
     */
    public Optional<Output<String>> subgroupCreationLevel() {
        return Optional.ofNullable(this.subgroupCreationLevel);
    }

    /**
     * Defaults to 48. Time before Two-factor authentication is enforced (in hours).
     * 
     */
    @Import(name="twoFactorGracePeriod")
    private @Nullable Output<Integer> twoFactorGracePeriod;

    /**
     * @return Defaults to 48. Time before Two-factor authentication is enforced (in hours).
     * 
     */
    public Optional<Output<Integer>> twoFactorGracePeriod() {
        return Optional.ofNullable(this.twoFactorGracePeriod);
    }

    /**
     * The group&#39;s visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
     * 
     */
    @Import(name="visibilityLevel")
    private @Nullable Output<String> visibilityLevel;

    /**
     * @return The group&#39;s visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
     * 
     */
    public Optional<Output<String>> visibilityLevel() {
        return Optional.ofNullable(this.visibilityLevel);
    }

    /**
     * The group&#39;s wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
     * 
     */
    @Import(name="wikiAccessLevel")
    private @Nullable Output<String> wikiAccessLevel;

    /**
     * @return The group&#39;s wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
     * 
     */
    public Optional<Output<String>> wikiAccessLevel() {
        return Optional.ofNullable(this.wikiAccessLevel);
    }

    private GroupArgs() {}

    private GroupArgs(GroupArgs $) {
        this.autoDevopsEnabled = $.autoDevopsEnabled;
        this.avatar = $.avatar;
        this.avatarHash = $.avatarHash;
        this.defaultBranchProtection = $.defaultBranchProtection;
        this.description = $.description;
        this.emailsEnabled = $.emailsEnabled;
        this.extraSharedRunnersMinutesLimit = $.extraSharedRunnersMinutesLimit;
        this.ipRestrictionRanges = $.ipRestrictionRanges;
        this.lfsEnabled = $.lfsEnabled;
        this.membershipLock = $.membershipLock;
        this.mentionsDisabled = $.mentionsDisabled;
        this.name = $.name;
        this.parentId = $.parentId;
        this.path = $.path;
        this.permanentlyRemoveOnDelete = $.permanentlyRemoveOnDelete;
        this.preventForkingOutsideGroup = $.preventForkingOutsideGroup;
        this.projectCreationLevel = $.projectCreationLevel;
        this.pushRules = $.pushRules;
        this.requestAccessEnabled = $.requestAccessEnabled;
        this.requireTwoFactorAuthentication = $.requireTwoFactorAuthentication;
        this.shareWithGroupLock = $.shareWithGroupLock;
        this.sharedRunnersMinutesLimit = $.sharedRunnersMinutesLimit;
        this.sharedRunnersSetting = $.sharedRunnersSetting;
        this.subgroupCreationLevel = $.subgroupCreationLevel;
        this.twoFactorGracePeriod = $.twoFactorGracePeriod;
        this.visibilityLevel = $.visibilityLevel;
        this.wikiAccessLevel = $.wikiAccessLevel;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupArgs $;

        public Builder() {
            $ = new GroupArgs();
        }

        public Builder(GroupArgs defaults) {
            $ = new GroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoDevopsEnabled Default to Auto DevOps pipeline for all projects within this group.
         * 
         * @return builder
         * 
         */
        public Builder autoDevopsEnabled(@Nullable Output<Boolean> autoDevopsEnabled) {
            $.autoDevopsEnabled = autoDevopsEnabled;
            return this;
        }

        /**
         * @param autoDevopsEnabled Default to Auto DevOps pipeline for all projects within this group.
         * 
         * @return builder
         * 
         */
        public Builder autoDevopsEnabled(Boolean autoDevopsEnabled) {
            return autoDevopsEnabled(Output.of(autoDevopsEnabled));
        }

        /**
         * @param avatar A local path to the avatar image to upload. **Note**: not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder avatar(@Nullable Output<String> avatar) {
            $.avatar = avatar;
            return this;
        }

        /**
         * @param avatar A local path to the avatar image to upload. **Note**: not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder avatar(String avatar) {
            return avatar(Output.of(avatar));
        }

        /**
         * @param avatarHash The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
         * 
         * @return builder
         * 
         */
        public Builder avatarHash(@Nullable Output<String> avatarHash) {
            $.avatarHash = avatarHash;
            return this;
        }

        /**
         * @param avatarHash The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
         * 
         * @return builder
         * 
         */
        public Builder avatarHash(String avatarHash) {
            return avatarHash(Output.of(avatarHash));
        }

        /**
         * @param defaultBranchProtection See [https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection](https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection). Valid values are: `0`, `1`, `2`, `3`, `4`.
         * 
         * @return builder
         * 
         */
        public Builder defaultBranchProtection(@Nullable Output<Integer> defaultBranchProtection) {
            $.defaultBranchProtection = defaultBranchProtection;
            return this;
        }

        /**
         * @param defaultBranchProtection See [https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection](https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection). Valid values are: `0`, `1`, `2`, `3`, `4`.
         * 
         * @return builder
         * 
         */
        public Builder defaultBranchProtection(Integer defaultBranchProtection) {
            return defaultBranchProtection(Output.of(defaultBranchProtection));
        }

        /**
         * @param description The group&#39;s description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The group&#39;s description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param emailsEnabled Enable email notifications.
         * 
         * @return builder
         * 
         */
        public Builder emailsEnabled(@Nullable Output<Boolean> emailsEnabled) {
            $.emailsEnabled = emailsEnabled;
            return this;
        }

        /**
         * @param emailsEnabled Enable email notifications.
         * 
         * @return builder
         * 
         */
        public Builder emailsEnabled(Boolean emailsEnabled) {
            return emailsEnabled(Output.of(emailsEnabled));
        }

        /**
         * @param extraSharedRunnersMinutesLimit Can be set by administrators only. Additional CI/CD minutes for this group.
         * 
         * @return builder
         * 
         */
        public Builder extraSharedRunnersMinutesLimit(@Nullable Output<Integer> extraSharedRunnersMinutesLimit) {
            $.extraSharedRunnersMinutesLimit = extraSharedRunnersMinutesLimit;
            return this;
        }

        /**
         * @param extraSharedRunnersMinutesLimit Can be set by administrators only. Additional CI/CD minutes for this group.
         * 
         * @return builder
         * 
         */
        public Builder extraSharedRunnersMinutesLimit(Integer extraSharedRunnersMinutesLimit) {
            return extraSharedRunnersMinutesLimit(Output.of(extraSharedRunnersMinutesLimit));
        }

        /**
         * @param ipRestrictionRanges A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictionRanges(@Nullable Output<List<String>> ipRestrictionRanges) {
            $.ipRestrictionRanges = ipRestrictionRanges;
            return this;
        }

        /**
         * @param ipRestrictionRanges A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictionRanges(List<String> ipRestrictionRanges) {
            return ipRestrictionRanges(Output.of(ipRestrictionRanges));
        }

        /**
         * @param ipRestrictionRanges A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictionRanges(String... ipRestrictionRanges) {
            return ipRestrictionRanges(List.of(ipRestrictionRanges));
        }

        /**
         * @param lfsEnabled Enable/disable Large File Storage (LFS) for the projects in this group.
         * 
         * @return builder
         * 
         */
        public Builder lfsEnabled(@Nullable Output<Boolean> lfsEnabled) {
            $.lfsEnabled = lfsEnabled;
            return this;
        }

        /**
         * @param lfsEnabled Enable/disable Large File Storage (LFS) for the projects in this group.
         * 
         * @return builder
         * 
         */
        public Builder lfsEnabled(Boolean lfsEnabled) {
            return lfsEnabled(Output.of(lfsEnabled));
        }

        /**
         * @param membershipLock Users cannot be added to projects in this group.
         * 
         * @return builder
         * 
         */
        public Builder membershipLock(@Nullable Output<Boolean> membershipLock) {
            $.membershipLock = membershipLock;
            return this;
        }

        /**
         * @param membershipLock Users cannot be added to projects in this group.
         * 
         * @return builder
         * 
         */
        public Builder membershipLock(Boolean membershipLock) {
            return membershipLock(Output.of(membershipLock));
        }

        /**
         * @param mentionsDisabled Disable the capability of a group from getting mentioned.
         * 
         * @return builder
         * 
         */
        public Builder mentionsDisabled(@Nullable Output<Boolean> mentionsDisabled) {
            $.mentionsDisabled = mentionsDisabled;
            return this;
        }

        /**
         * @param mentionsDisabled Disable the capability of a group from getting mentioned.
         * 
         * @return builder
         * 
         */
        public Builder mentionsDisabled(Boolean mentionsDisabled) {
            return mentionsDisabled(Output.of(mentionsDisabled));
        }

        /**
         * @param name The name of the group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parentId Id of the parent group (creates a nested group).
         * 
         * @return builder
         * 
         */
        public Builder parentId(@Nullable Output<Integer> parentId) {
            $.parentId = parentId;
            return this;
        }

        /**
         * @param parentId Id of the parent group (creates a nested group).
         * 
         * @return builder
         * 
         */
        public Builder parentId(Integer parentId) {
            return parentId(Output.of(parentId));
        }

        /**
         * @param path The path of the group.
         * 
         * @return builder
         * 
         */
        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path of the group.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param permanentlyRemoveOnDelete Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
         * 
         * @return builder
         * 
         */
        public Builder permanentlyRemoveOnDelete(@Nullable Output<Boolean> permanentlyRemoveOnDelete) {
            $.permanentlyRemoveOnDelete = permanentlyRemoveOnDelete;
            return this;
        }

        /**
         * @param permanentlyRemoveOnDelete Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
         * 
         * @return builder
         * 
         */
        public Builder permanentlyRemoveOnDelete(Boolean permanentlyRemoveOnDelete) {
            return permanentlyRemoveOnDelete(Output.of(permanentlyRemoveOnDelete));
        }

        /**
         * @param preventForkingOutsideGroup Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
         * 
         * @return builder
         * 
         */
        public Builder preventForkingOutsideGroup(@Nullable Output<Boolean> preventForkingOutsideGroup) {
            $.preventForkingOutsideGroup = preventForkingOutsideGroup;
            return this;
        }

        /**
         * @param preventForkingOutsideGroup Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
         * 
         * @return builder
         * 
         */
        public Builder preventForkingOutsideGroup(Boolean preventForkingOutsideGroup) {
            return preventForkingOutsideGroup(Output.of(preventForkingOutsideGroup));
        }

        /**
         * @param projectCreationLevel Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
         * 
         * @return builder
         * 
         */
        public Builder projectCreationLevel(@Nullable Output<String> projectCreationLevel) {
            $.projectCreationLevel = projectCreationLevel;
            return this;
        }

        /**
         * @param projectCreationLevel Determine if developers can create projects in the group. Valid values are: `noone`, `maintainer`, `developer`
         * 
         * @return builder
         * 
         */
        public Builder projectCreationLevel(String projectCreationLevel) {
            return projectCreationLevel(Output.of(projectCreationLevel));
        }

        /**
         * @param pushRules Push rules for the group.
         * 
         * @return builder
         * 
         */
        public Builder pushRules(@Nullable Output<GroupPushRulesArgs> pushRules) {
            $.pushRules = pushRules;
            return this;
        }

        /**
         * @param pushRules Push rules for the group.
         * 
         * @return builder
         * 
         */
        public Builder pushRules(GroupPushRulesArgs pushRules) {
            return pushRules(Output.of(pushRules));
        }

        /**
         * @param requestAccessEnabled Allow users to request member access.
         * 
         * @return builder
         * 
         */
        public Builder requestAccessEnabled(@Nullable Output<Boolean> requestAccessEnabled) {
            $.requestAccessEnabled = requestAccessEnabled;
            return this;
        }

        /**
         * @param requestAccessEnabled Allow users to request member access.
         * 
         * @return builder
         * 
         */
        public Builder requestAccessEnabled(Boolean requestAccessEnabled) {
            return requestAccessEnabled(Output.of(requestAccessEnabled));
        }

        /**
         * @param requireTwoFactorAuthentication Require all users in this group to setup Two-factor authentication.
         * 
         * @return builder
         * 
         */
        public Builder requireTwoFactorAuthentication(@Nullable Output<Boolean> requireTwoFactorAuthentication) {
            $.requireTwoFactorAuthentication = requireTwoFactorAuthentication;
            return this;
        }

        /**
         * @param requireTwoFactorAuthentication Require all users in this group to setup Two-factor authentication.
         * 
         * @return builder
         * 
         */
        public Builder requireTwoFactorAuthentication(Boolean requireTwoFactorAuthentication) {
            return requireTwoFactorAuthentication(Output.of(requireTwoFactorAuthentication));
        }

        /**
         * @param shareWithGroupLock Prevent sharing a project with another group within this group.
         * 
         * @return builder
         * 
         */
        public Builder shareWithGroupLock(@Nullable Output<Boolean> shareWithGroupLock) {
            $.shareWithGroupLock = shareWithGroupLock;
            return this;
        }

        /**
         * @param shareWithGroupLock Prevent sharing a project with another group within this group.
         * 
         * @return builder
         * 
         */
        public Builder shareWithGroupLock(Boolean shareWithGroupLock) {
            return shareWithGroupLock(Output.of(shareWithGroupLock));
        }

        /**
         * @param sharedRunnersMinutesLimit Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or &gt; 0.
         * 
         * @return builder
         * 
         */
        public Builder sharedRunnersMinutesLimit(@Nullable Output<Integer> sharedRunnersMinutesLimit) {
            $.sharedRunnersMinutesLimit = sharedRunnersMinutesLimit;
            return this;
        }

        /**
         * @param sharedRunnersMinutesLimit Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or &gt; 0.
         * 
         * @return builder
         * 
         */
        public Builder sharedRunnersMinutesLimit(Integer sharedRunnersMinutesLimit) {
            return sharedRunnersMinutesLimit(Output.of(sharedRunnersMinutesLimit));
        }

        /**
         * @param sharedRunnersSetting Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
         * 
         * @return builder
         * 
         */
        public Builder sharedRunnersSetting(@Nullable Output<String> sharedRunnersSetting) {
            $.sharedRunnersSetting = sharedRunnersSetting;
            return this;
        }

        /**
         * @param sharedRunnersSetting Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
         * 
         * @return builder
         * 
         */
        public Builder sharedRunnersSetting(String sharedRunnersSetting) {
            return sharedRunnersSetting(Output.of(sharedRunnersSetting));
        }

        /**
         * @param subgroupCreationLevel Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder subgroupCreationLevel(@Nullable Output<String> subgroupCreationLevel) {
            $.subgroupCreationLevel = subgroupCreationLevel;
            return this;
        }

        /**
         * @param subgroupCreationLevel Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder subgroupCreationLevel(String subgroupCreationLevel) {
            return subgroupCreationLevel(Output.of(subgroupCreationLevel));
        }

        /**
         * @param twoFactorGracePeriod Defaults to 48. Time before Two-factor authentication is enforced (in hours).
         * 
         * @return builder
         * 
         */
        public Builder twoFactorGracePeriod(@Nullable Output<Integer> twoFactorGracePeriod) {
            $.twoFactorGracePeriod = twoFactorGracePeriod;
            return this;
        }

        /**
         * @param twoFactorGracePeriod Defaults to 48. Time before Two-factor authentication is enforced (in hours).
         * 
         * @return builder
         * 
         */
        public Builder twoFactorGracePeriod(Integer twoFactorGracePeriod) {
            return twoFactorGracePeriod(Output.of(twoFactorGracePeriod));
        }

        /**
         * @param visibilityLevel The group&#39;s visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
         * 
         * @return builder
         * 
         */
        public Builder visibilityLevel(@Nullable Output<String> visibilityLevel) {
            $.visibilityLevel = visibilityLevel;
            return this;
        }

        /**
         * @param visibilityLevel The group&#39;s visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
         * 
         * @return builder
         * 
         */
        public Builder visibilityLevel(String visibilityLevel) {
            return visibilityLevel(Output.of(visibilityLevel));
        }

        /**
         * @param wikiAccessLevel The group&#39;s wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
         * 
         * @return builder
         * 
         */
        public Builder wikiAccessLevel(@Nullable Output<String> wikiAccessLevel) {
            $.wikiAccessLevel = wikiAccessLevel;
            return this;
        }

        /**
         * @param wikiAccessLevel The group&#39;s wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
         * 
         * @return builder
         * 
         */
        public Builder wikiAccessLevel(String wikiAccessLevel) {
            return wikiAccessLevel(Output.of(wikiAccessLevel));
        }

        public GroupArgs build() {
            if ($.path == null) {
                throw new MissingRequiredPropertyException("GroupArgs", "path");
            }
            return $;
        }
    }

}
