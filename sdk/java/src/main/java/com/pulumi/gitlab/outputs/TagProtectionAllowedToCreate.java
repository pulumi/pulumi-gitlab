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
public final class TagProtectionAllowedToCreate {
    /**
     * @return Level of access.
     * 
     */
    private @Nullable String accessLevel;
    /**
     * @return Readable description of level of access.
     * 
     */
    private @Nullable String accessLevelDescription;
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

    private TagProtectionAllowedToCreate() {}
    /**
     * @return Level of access.
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

    public static Builder builder(TagProtectionAllowedToCreate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String accessLevel;
        private @Nullable String accessLevelDescription;
        private @Nullable Integer groupId;
        private @Nullable Integer userId;
        public Builder() {}
        public Builder(TagProtectionAllowedToCreate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessLevel = defaults.accessLevel;
    	      this.accessLevelDescription = defaults.accessLevelDescription;
    	      this.groupId = defaults.groupId;
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
        public Builder userId(@Nullable Integer userId) {
            this.userId = userId;
            return this;
        }
        public TagProtectionAllowedToCreate build() {
            final var _resultValue = new TagProtectionAllowedToCreate();
            _resultValue.accessLevel = accessLevel;
            _resultValue.accessLevelDescription = accessLevelDescription;
            _resultValue.groupId = groupId;
            _resultValue.userId = userId;
            return _resultValue;
        }
    }
}
