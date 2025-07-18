// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupShareGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupShareGroupArgs Empty = new GroupShareGroupArgs();

    /**
     * Share expiration date. Format: `YYYY-MM-DD`
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return Share expiration date. Format: `YYYY-MM-DD`
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     */
    @Import(name="groupAccess", required=true)
    private Output<String> groupAccess;

    /**
     * @return The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     */
    public Output<String> groupAccess() {
        return this.groupAccess;
    }

    /**
     * The id of the main group to be shared.
     * 
     */
    @Import(name="groupId", required=true)
    private Output<String> groupId;

    /**
     * @return The id of the main group to be shared.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }

    /**
     * The ID of a custom member role. Only available for Ultimate instances. If `member_role_id` is removed from the config, the group share will revert to a base role.
     * 
     */
    @Import(name="memberRoleId")
    private @Nullable Output<Integer> memberRoleId;

    /**
     * @return The ID of a custom member role. Only available for Ultimate instances. If `member_role_id` is removed from the config, the group share will revert to a base role.
     * 
     */
    public Optional<Output<Integer>> memberRoleId() {
        return Optional.ofNullable(this.memberRoleId);
    }

    /**
     * The id of the additional group with which the main group will be shared.
     * 
     */
    @Import(name="shareGroupId", required=true)
    private Output<Integer> shareGroupId;

    /**
     * @return The id of the additional group with which the main group will be shared.
     * 
     */
    public Output<Integer> shareGroupId() {
        return this.shareGroupId;
    }

    private GroupShareGroupArgs() {}

    private GroupShareGroupArgs(GroupShareGroupArgs $) {
        this.expiresAt = $.expiresAt;
        this.groupAccess = $.groupAccess;
        this.groupId = $.groupId;
        this.memberRoleId = $.memberRoleId;
        this.shareGroupId = $.shareGroupId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupShareGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupShareGroupArgs $;

        public Builder() {
            $ = new GroupShareGroupArgs();
        }

        public Builder(GroupShareGroupArgs defaults) {
            $ = new GroupShareGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param expiresAt Share expiration date. Format: `YYYY-MM-DD`
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt Share expiration date. Format: `YYYY-MM-DD`
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param groupAccess The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
         * 
         * @return builder
         * 
         */
        public Builder groupAccess(Output<String> groupAccess) {
            $.groupAccess = groupAccess;
            return this;
        }

        /**
         * @param groupAccess The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
         * 
         * @return builder
         * 
         */
        public Builder groupAccess(String groupAccess) {
            return groupAccess(Output.of(groupAccess));
        }

        /**
         * @param groupId The id of the main group to be shared.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The id of the main group to be shared.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param memberRoleId The ID of a custom member role. Only available for Ultimate instances. If `member_role_id` is removed from the config, the group share will revert to a base role.
         * 
         * @return builder
         * 
         */
        public Builder memberRoleId(@Nullable Output<Integer> memberRoleId) {
            $.memberRoleId = memberRoleId;
            return this;
        }

        /**
         * @param memberRoleId The ID of a custom member role. Only available for Ultimate instances. If `member_role_id` is removed from the config, the group share will revert to a base role.
         * 
         * @return builder
         * 
         */
        public Builder memberRoleId(Integer memberRoleId) {
            return memberRoleId(Output.of(memberRoleId));
        }

        /**
         * @param shareGroupId The id of the additional group with which the main group will be shared.
         * 
         * @return builder
         * 
         */
        public Builder shareGroupId(Output<Integer> shareGroupId) {
            $.shareGroupId = shareGroupId;
            return this;
        }

        /**
         * @param shareGroupId The id of the additional group with which the main group will be shared.
         * 
         * @return builder
         * 
         */
        public Builder shareGroupId(Integer shareGroupId) {
            return shareGroupId(Output.of(shareGroupId));
        }

        public GroupShareGroupArgs build() {
            if ($.groupAccess == null) {
                throw new MissingRequiredPropertyException("GroupShareGroupArgs", "groupAccess");
            }
            if ($.groupId == null) {
                throw new MissingRequiredPropertyException("GroupShareGroupArgs", "groupId");
            }
            if ($.shareGroupId == null) {
                throw new MissingRequiredPropertyException("GroupShareGroupArgs", "shareGroupId");
            }
            return $;
        }
    }

}
