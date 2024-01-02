// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GroupProtectedEnvironmentDeployAccessLevel {
    /**
     * @return Levels of access required to deploy to this protected environment. Valid values are `developer`, `maintainer`.
     * 
     */
    private @Nullable String accessLevel;
    /**
     * @return Readable description of level of access.
     * 
     */
    private @Nullable String accessLevelDescription;
    /**
     * @return The ID of the group allowed to deploy to this protected environment. The group must be a sub-group under the given group.
     * 
     */
    private @Nullable Integer groupId;
    /**
     * @return The unique ID of the Deploy Access Level object.
     * 
     */
    private @Nullable Integer id;
    /**
     * @return The ID of the user allowed to deploy to this protected environment. The user must be a member of the group with Maintainer role or higher.
     * 
     */
    private @Nullable Integer userId;

    private GroupProtectedEnvironmentDeployAccessLevel() {}
    /**
     * @return Levels of access required to deploy to this protected environment. Valid values are `developer`, `maintainer`.
     * 
     */
    public Optional<String> accessLevel() {
        return Optional.ofNullable(this.accessLevel);
    }
    /**
     * @return Readable description of level of access.
     * 
     */
    public Optional<String> accessLevelDescription() {
        return Optional.ofNullable(this.accessLevelDescription);
    }
    /**
     * @return The ID of the group allowed to deploy to this protected environment. The group must be a sub-group under the given group.
     * 
     */
    public Optional<Integer> groupId() {
        return Optional.ofNullable(this.groupId);
    }
    /**
     * @return The unique ID of the Deploy Access Level object.
     * 
     */
    public Optional<Integer> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The ID of the user allowed to deploy to this protected environment. The user must be a member of the group with Maintainer role or higher.
     * 
     */
    public Optional<Integer> userId() {
        return Optional.ofNullable(this.userId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GroupProtectedEnvironmentDeployAccessLevel defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String accessLevel;
        private @Nullable String accessLevelDescription;
        private @Nullable Integer groupId;
        private @Nullable Integer id;
        private @Nullable Integer userId;
        public Builder() {}
        public Builder(GroupProtectedEnvironmentDeployAccessLevel defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessLevel = defaults.accessLevel;
    	      this.accessLevelDescription = defaults.accessLevelDescription;
    	      this.groupId = defaults.groupId;
    	      this.id = defaults.id;
    	      this.userId = defaults.userId;
        }

        @CustomType.Setter
        public Builder accessLevel(@Nullable String accessLevel) {

            this.accessLevel = accessLevel;
            return this;
        }
        @CustomType.Setter
        public Builder accessLevelDescription(@Nullable String accessLevelDescription) {

            this.accessLevelDescription = accessLevelDescription;
            return this;
        }
        @CustomType.Setter
        public Builder groupId(@Nullable Integer groupId) {

            this.groupId = groupId;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable Integer id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder userId(@Nullable Integer userId) {

            this.userId = userId;
            return this;
        }
        public GroupProtectedEnvironmentDeployAccessLevel build() {
            final var _resultValue = new GroupProtectedEnvironmentDeployAccessLevel();
            _resultValue.accessLevel = accessLevel;
            _resultValue.accessLevelDescription = accessLevelDescription;
            _resultValue.groupId = groupId;
            _resultValue.id = id;
            _resultValue.userId = userId;
            return _resultValue;
        }
    }
}
