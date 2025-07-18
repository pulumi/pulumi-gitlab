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


public final class GetProjectProtectedBranchesProtectedBranchMergeAccessLevel extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectProtectedBranchesProtectedBranchMergeAccessLevel Empty = new GetProjectProtectedBranchesProtectedBranchMergeAccessLevel();

    /**
     * Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
     * 
     */
    @Import(name="accessLevel", required=true)
    private String accessLevel;

    /**
     * @return Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
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
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
     * 
     */
    @Import(name="groupId")
    private @Nullable Integer groupId;

    /**
     * @return The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
     * 
     */
    public Optional<Integer> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
     * 
     */
    @Import(name="userId")
    private @Nullable Integer userId;

    /**
     * @return The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
     * 
     */
    public Optional<Integer> userId() {
        return Optional.ofNullable(this.userId);
    }

    private GetProjectProtectedBranchesProtectedBranchMergeAccessLevel() {}

    private GetProjectProtectedBranchesProtectedBranchMergeAccessLevel(GetProjectProtectedBranchesProtectedBranchMergeAccessLevel $) {
        this.accessLevel = $.accessLevel;
        this.accessLevelDescription = $.accessLevelDescription;
        this.groupId = $.groupId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectProtectedBranchesProtectedBranchMergeAccessLevel defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectProtectedBranchesProtectedBranchMergeAccessLevel $;

        public Builder() {
            $ = new GetProjectProtectedBranchesProtectedBranchMergeAccessLevel();
        }

        public Builder(GetProjectProtectedBranchesProtectedBranchMergeAccessLevel defaults) {
            $ = new GetProjectProtectedBranchesProtectedBranchMergeAccessLevel(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
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
         * @param groupId The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Integer groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param userId The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Integer userId) {
            $.userId = userId;
            return this;
        }

        public GetProjectProtectedBranchesProtectedBranchMergeAccessLevel build() {
            if ($.accessLevel == null) {
                throw new MissingRequiredPropertyException("GetProjectProtectedBranchesProtectedBranchMergeAccessLevel", "accessLevel");
            }
            if ($.accessLevelDescription == null) {
                throw new MissingRequiredPropertyException("GetProjectProtectedBranchesProtectedBranchMergeAccessLevel", "accessLevelDescription");
            }
            return $;
        }
    }

}
