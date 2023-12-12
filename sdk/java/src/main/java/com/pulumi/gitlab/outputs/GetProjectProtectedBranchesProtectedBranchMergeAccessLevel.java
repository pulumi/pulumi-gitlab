// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectProtectedBranchesProtectedBranchMergeAccessLevel {
    private String accessLevel;
    private String accessLevelDescription;
    private Integer groupId;
    private Integer userId;

    private GetProjectProtectedBranchesProtectedBranchMergeAccessLevel() {}
    public String accessLevel() {
        return this.accessLevel;
    }
    public String accessLevelDescription() {
        return this.accessLevelDescription;
    }
    public Integer groupId() {
        return this.groupId;
    }
    public Integer userId() {
        return this.userId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectProtectedBranchesProtectedBranchMergeAccessLevel defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessLevel;
        private String accessLevelDescription;
        private Integer groupId;
        private Integer userId;
        public Builder() {}
        public Builder(GetProjectProtectedBranchesProtectedBranchMergeAccessLevel defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessLevel = defaults.accessLevel;
    	      this.accessLevelDescription = defaults.accessLevelDescription;
    	      this.groupId = defaults.groupId;
    	      this.userId = defaults.userId;
        }

        @CustomType.Setter
        public Builder accessLevel(String accessLevel) {
            this.accessLevel = Objects.requireNonNull(accessLevel);
            return this;
        }
        @CustomType.Setter
        public Builder accessLevelDescription(String accessLevelDescription) {
            this.accessLevelDescription = Objects.requireNonNull(accessLevelDescription);
            return this;
        }
        @CustomType.Setter
        public Builder groupId(Integer groupId) {
            this.groupId = Objects.requireNonNull(groupId);
            return this;
        }
        @CustomType.Setter
        public Builder userId(Integer userId) {
            this.userId = Objects.requireNonNull(userId);
            return this;
        }
        public GetProjectProtectedBranchesProtectedBranchMergeAccessLevel build() {
            final var _resultValue = new GetProjectProtectedBranchesProtectedBranchMergeAccessLevel();
            _resultValue.accessLevel = accessLevel;
            _resultValue.accessLevelDescription = accessLevelDescription;
            _resultValue.groupId = groupId;
            _resultValue.userId = userId;
            return _resultValue;
        }
    }
}
