// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMembershipState extends com.pulumi.resources.ResourceArgs {

    public static final GroupMembershipState Empty = new GroupMembershipState();

    /**
     * Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    public Optional<Output<String>> accessLevel() {
        return Optional.ofNullable(this.accessLevel);
    }

    /**
     * Expiration date for the group membership. Format: `YYYY-MM-DD`
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return Expiration date for the group membership. Format: `YYYY-MM-DD`
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * The ID of the group.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<Integer> groupId;

    /**
     * @return The ID of the group.
     * 
     */
    public Optional<Output<Integer>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * The ID of a custom member role. Only available for Ultimate instances.
     * 
     */
    @Import(name="memberRoleId")
    private @Nullable Output<Integer> memberRoleId;

    /**
     * @return The ID of a custom member role. Only available for Ultimate instances.
     * 
     */
    public Optional<Output<Integer>> memberRoleId() {
        return Optional.ofNullable(this.memberRoleId);
    }

    /**
     * Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
     * 
     */
    @Import(name="skipSubresourcesOnDestroy")
    private @Nullable Output<Boolean> skipSubresourcesOnDestroy;

    /**
     * @return Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
     * 
     */
    public Optional<Output<Boolean>> skipSubresourcesOnDestroy() {
        return Optional.ofNullable(this.skipSubresourcesOnDestroy);
    }

    /**
     * Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
     * 
     */
    @Import(name="unassignIssuablesOnDestroy")
    private @Nullable Output<Boolean> unassignIssuablesOnDestroy;

    /**
     * @return Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
     * 
     */
    public Optional<Output<Boolean>> unassignIssuablesOnDestroy() {
        return Optional.ofNullable(this.unassignIssuablesOnDestroy);
    }

    /**
     * The ID of the user.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<Integer> userId;

    /**
     * @return The ID of the user.
     * 
     */
    public Optional<Output<Integer>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private GroupMembershipState() {}

    private GroupMembershipState(GroupMembershipState $) {
        this.accessLevel = $.accessLevel;
        this.expiresAt = $.expiresAt;
        this.groupId = $.groupId;
        this.memberRoleId = $.memberRoleId;
        this.skipSubresourcesOnDestroy = $.skipSubresourcesOnDestroy;
        this.unassignIssuablesOnDestroy = $.unassignIssuablesOnDestroy;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMembershipState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMembershipState $;

        public Builder() {
            $ = new GroupMembershipState();
        }

        public Builder(GroupMembershipState defaults) {
            $ = new GroupMembershipState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
        }

        /**
         * @param expiresAt Expiration date for the group membership. Format: `YYYY-MM-DD`
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt Expiration date for the group membership. Format: `YYYY-MM-DD`
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param groupId The ID of the group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<Integer> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of the group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Integer groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param memberRoleId The ID of a custom member role. Only available for Ultimate instances.
         * 
         * @return builder
         * 
         */
        public Builder memberRoleId(@Nullable Output<Integer> memberRoleId) {
            $.memberRoleId = memberRoleId;
            return this;
        }

        /**
         * @param memberRoleId The ID of a custom member role. Only available for Ultimate instances.
         * 
         * @return builder
         * 
         */
        public Builder memberRoleId(Integer memberRoleId) {
            return memberRoleId(Output.of(memberRoleId));
        }

        /**
         * @param skipSubresourcesOnDestroy Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
         * 
         * @return builder
         * 
         */
        public Builder skipSubresourcesOnDestroy(@Nullable Output<Boolean> skipSubresourcesOnDestroy) {
            $.skipSubresourcesOnDestroy = skipSubresourcesOnDestroy;
            return this;
        }

        /**
         * @param skipSubresourcesOnDestroy Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
         * 
         * @return builder
         * 
         */
        public Builder skipSubresourcesOnDestroy(Boolean skipSubresourcesOnDestroy) {
            return skipSubresourcesOnDestroy(Output.of(skipSubresourcesOnDestroy));
        }

        /**
         * @param unassignIssuablesOnDestroy Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
         * 
         * @return builder
         * 
         */
        public Builder unassignIssuablesOnDestroy(@Nullable Output<Boolean> unassignIssuablesOnDestroy) {
            $.unassignIssuablesOnDestroy = unassignIssuablesOnDestroy;
            return this;
        }

        /**
         * @param unassignIssuablesOnDestroy Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
         * 
         * @return builder
         * 
         */
        public Builder unassignIssuablesOnDestroy(Boolean unassignIssuablesOnDestroy) {
            return unassignIssuablesOnDestroy(Output.of(unassignIssuablesOnDestroy));
        }

        /**
         * @param userId The ID of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        public GroupMembershipState build() {
            return $;
        }
    }

}
