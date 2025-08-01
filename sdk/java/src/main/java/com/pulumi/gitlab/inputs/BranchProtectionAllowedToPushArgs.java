// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchProtectionAllowedToPushArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchProtectionAllowedToPushArgs Empty = new BranchProtectionAllowedToPushArgs();

    /**
     * Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
     * 
     */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
     * 
     */
    public Optional<Output<String>> accessLevel() {
        return Optional.ofNullable(this.accessLevel);
    }

    /**
     * Readable description of access level.
     * 
     */
    @Import(name="accessLevelDescription")
    private @Nullable Output<String> accessLevelDescription;

    /**
     * @return Readable description of access level.
     * 
     */
    public Optional<Output<String>> accessLevelDescription() {
        return Optional.ofNullable(this.accessLevelDescription);
    }

    /**
     * The ID of a GitLab deploy key allowed to perform the relevant action. Mutually exclusive with `group_id` and `user_id`. This field is read-only until Gitlab 17.5.
     * 
     */
    @Import(name="deployKeyId")
    private @Nullable Output<Integer> deployKeyId;

    /**
     * @return The ID of a GitLab deploy key allowed to perform the relevant action. Mutually exclusive with `group_id` and `user_id`. This field is read-only until Gitlab 17.5.
     * 
     */
    public Optional<Output<Integer>> deployKeyId() {
        return Optional.ofNullable(this.deployKeyId);
    }

    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `user_id`.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<Integer> groupId;

    /**
     * @return The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `user_id`.
     * 
     */
    public Optional<Output<Integer>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `group_id`.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<Integer> userId;

    /**
     * @return The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `group_id`.
     * 
     */
    public Optional<Output<Integer>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private BranchProtectionAllowedToPushArgs() {}

    private BranchProtectionAllowedToPushArgs(BranchProtectionAllowedToPushArgs $) {
        this.accessLevel = $.accessLevel;
        this.accessLevelDescription = $.accessLevelDescription;
        this.deployKeyId = $.deployKeyId;
        this.groupId = $.groupId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchProtectionAllowedToPushArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchProtectionAllowedToPushArgs $;

        public Builder() {
            $ = new BranchProtectionAllowedToPushArgs();
        }

        public Builder(BranchProtectionAllowedToPushArgs defaults) {
            $ = new BranchProtectionAllowedToPushArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
        }

        /**
         * @param accessLevelDescription Readable description of access level.
         * 
         * @return builder
         * 
         */
        public Builder accessLevelDescription(@Nullable Output<String> accessLevelDescription) {
            $.accessLevelDescription = accessLevelDescription;
            return this;
        }

        /**
         * @param accessLevelDescription Readable description of access level.
         * 
         * @return builder
         * 
         */
        public Builder accessLevelDescription(String accessLevelDescription) {
            return accessLevelDescription(Output.of(accessLevelDescription));
        }

        /**
         * @param deployKeyId The ID of a GitLab deploy key allowed to perform the relevant action. Mutually exclusive with `group_id` and `user_id`. This field is read-only until Gitlab 17.5.
         * 
         * @return builder
         * 
         */
        public Builder deployKeyId(@Nullable Output<Integer> deployKeyId) {
            $.deployKeyId = deployKeyId;
            return this;
        }

        /**
         * @param deployKeyId The ID of a GitLab deploy key allowed to perform the relevant action. Mutually exclusive with `group_id` and `user_id`. This field is read-only until Gitlab 17.5.
         * 
         * @return builder
         * 
         */
        public Builder deployKeyId(Integer deployKeyId) {
            return deployKeyId(Output.of(deployKeyId));
        }

        /**
         * @param groupId The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `user_id`.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<Integer> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `user_id`.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Integer groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param userId The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        public BranchProtectionAllowedToPushArgs build() {
            return $;
        }
    }

}
