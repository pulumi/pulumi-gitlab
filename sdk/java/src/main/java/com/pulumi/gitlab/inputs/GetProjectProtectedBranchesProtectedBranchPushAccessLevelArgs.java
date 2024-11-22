// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs Empty = new GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs();

    /**
     * Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     * 
     */
    @Import(name="accessLevel", required=true)
    private Output<String> accessLevel;

    /**
     * @return Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     * 
     */
    public Output<String> accessLevel() {
        return this.accessLevel;
    }

    /**
     * Readable description of access level.
     * 
     */
    @Import(name="accessLevelDescription", required=true)
    private Output<String> accessLevelDescription;

    /**
     * @return Readable description of access level.
     * 
     */
    public Output<String> accessLevelDescription() {
        return this.accessLevelDescription;
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

    private GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs() {}

    private GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs(GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs $) {
        this.accessLevel = $.accessLevel;
        this.accessLevelDescription = $.accessLevelDescription;
        this.deployKeyId = $.deployKeyId;
        this.groupId = $.groupId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs $;

        public Builder() {
            $ = new GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs();
        }

        public Builder(GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs defaults) {
            $ = new GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
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
        public Builder accessLevelDescription(Output<String> accessLevelDescription) {
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

        public GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs build() {
            if ($.accessLevel == null) {
                throw new MissingRequiredPropertyException("GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs", "accessLevel");
            }
            if ($.accessLevelDescription == null) {
                throw new MissingRequiredPropertyException("GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs", "accessLevelDescription");
            }
            return $;
        }
    }

}
