// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetProjectProtectedBranchMergeAccessLevel {
    /**
     * @return Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
     * 
     */
    private String accessLevel;
    /**
     * @return Readable description of access level.
     * 
     */
    private String accessLevelDescription;
    /**
     * @return The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
     * 
     */
    private @Nullable Integer groupId;
    /**
     * @return The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
     * 
     */
    private @Nullable Integer userId;

    private GetProjectProtectedBranchMergeAccessLevel() {}
    /**
     * @return Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
     * 
     */
    public String accessLevel() {
        return this.accessLevel;
    }
    /**
     * @return Readable description of access level.
     * 
     */
    public String accessLevelDescription() {
        return this.accessLevelDescription;
    }
    /**
     * @return The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
     * 
     */
    public Optional<Integer> groupId() {
        return Optional.ofNullable(this.groupId);
    }
    /**
     * @return The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
     * 
     */
    public Optional<Integer> userId() {
        return Optional.ofNullable(this.userId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectProtectedBranchMergeAccessLevel defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessLevel;
        private String accessLevelDescription;
        private @Nullable Integer groupId;
        private @Nullable Integer userId;
        public Builder() {}
        public Builder(GetProjectProtectedBranchMergeAccessLevel defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessLevel = defaults.accessLevel;
    	      this.accessLevelDescription = defaults.accessLevelDescription;
    	      this.groupId = defaults.groupId;
    	      this.userId = defaults.userId;
        }

        @CustomType.Setter
        public Builder accessLevel(String accessLevel) {
            if (accessLevel == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedBranchMergeAccessLevel", "accessLevel");
            }
            this.accessLevel = accessLevel;
            return this;
        }
        @CustomType.Setter
        public Builder accessLevelDescription(String accessLevelDescription) {
            if (accessLevelDescription == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedBranchMergeAccessLevel", "accessLevelDescription");
            }
            this.accessLevelDescription = accessLevelDescription;
            return this;
        }
        @CustomType.Setter
        public Builder groupId(@Nullable Integer groupId) {

            this.groupId = groupId;
            return this;
        }
        @CustomType.Setter
        public Builder userId(@Nullable Integer userId) {

            this.userId = userId;
            return this;
        }
        public GetProjectProtectedBranchMergeAccessLevel build() {
            final var _resultValue = new GetProjectProtectedBranchMergeAccessLevel();
            _resultValue.accessLevel = accessLevel;
            _resultValue.accessLevelDescription = accessLevelDescription;
            _resultValue.groupId = groupId;
            _resultValue.userId = userId;
            return _resultValue;
        }
    }
}
