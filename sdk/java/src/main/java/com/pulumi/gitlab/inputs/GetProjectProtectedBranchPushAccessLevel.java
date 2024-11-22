// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetProjectProtectedBranchPushAccessLevel extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectProtectedBranchPushAccessLevel Empty = new GetProjectProtectedBranchPushAccessLevel();

    /**
     * Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     * 
     */
    @Import(name="accessLevel", required=true)
    private String accessLevel;

    /**
     * @return Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     * 
     */
    public String accessLevel() {
        return this.accessLevel;
    }

    /**
     * Readable description of access level.
     * 
     */
    @Import(name="accessLevelDescription", required=true)
    private String accessLevelDescription;

    /**
     * @return Readable description of access level.
     * 
     */
    public String accessLevelDescription() {
        return this.accessLevelDescription;
    }

    /**
     * The ID of a GitLab deploy key allowed to perform the relevant action. Mutually exclusive with `group_id` and `user_id`. This field is read-only until Gitlab 17.5.
     * 
     */
    @Import(name="deployKeyId")
    private @Nullable Integer deployKeyId;

    /**
     * @return The ID of a GitLab deploy key allowed to perform the relevant action. Mutually exclusive with `group_id` and `user_id`. This field is read-only until Gitlab 17.5.
     * 
     */
    public Optional<Integer> deployKeyId() {
        return Optional.ofNullable(this.deployKeyId);
    }

    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `user_id`.
     * 
     */
    @Import(name="groupId")
    private @Nullable Integer groupId;

    /**
     * @return The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `user_id`.
     * 
     */
    public Optional<Integer> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `group_id`.
     * 
     */
    @Import(name="userId")
    private @Nullable Integer userId;

    /**
     * @return The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `group_id`.
     * 
     */
    public Optional<Integer> userId() {
        return Optional.ofNullable(this.userId);
    }

    private GetProjectProtectedBranchPushAccessLevel() {}

    private GetProjectProtectedBranchPushAccessLevel(GetProjectProtectedBranchPushAccessLevel $) {
        this.accessLevel = $.accessLevel;
        this.accessLevelDescription = $.accessLevelDescription;
        this.deployKeyId = $.deployKeyId;
        this.groupId = $.groupId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectProtectedBranchPushAccessLevel defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectProtectedBranchPushAccessLevel $;

        public Builder() {
            $ = new GetProjectProtectedBranchPushAccessLevel();
        }

        public Builder(GetProjectProtectedBranchPushAccessLevel defaults) {
            $ = new GetProjectProtectedBranchPushAccessLevel(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevelDescription Readable description of access level.
         * 
         * @return builder
         * 
         */
        public Builder accessLevelDescription(String accessLevelDescription) {
            $.accessLevelDescription = accessLevelDescription;
            return this;
        }

        /**
         * @param deployKeyId The ID of a GitLab deploy key allowed to perform the relevant action. Mutually exclusive with `group_id` and `user_id`. This field is read-only until Gitlab 17.5.
         * 
         * @return builder
         * 
         */
        public Builder deployKeyId(@Nullable Integer deployKeyId) {
            $.deployKeyId = deployKeyId;
            return this;
        }

        /**
         * @param groupId The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `user_id`.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Integer groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param userId The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Integer userId) {
            $.userId = userId;
            return this;
        }

        public GetProjectProtectedBranchPushAccessLevel build() {
            if ($.accessLevel == null) {
                throw new MissingRequiredPropertyException("GetProjectProtectedBranchPushAccessLevel", "accessLevel");
            }
            if ($.accessLevelDescription == null) {
                throw new MissingRequiredPropertyException("GetProjectProtectedBranchPushAccessLevel", "accessLevelDescription");
            }
            return $;
        }
    }

}
