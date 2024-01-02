// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.outputs.GetProjectProtectedBranchesProtectedBranchMergeAccessLevel;
import com.pulumi.gitlab.outputs.GetProjectProtectedBranchesProtectedBranchPushAccessLevel;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetProjectProtectedBranchesProtectedBranch {
    private Boolean allowForcePush;
    private Boolean codeOwnerApprovalRequired;
    private Integer id;
    private List<GetProjectProtectedBranchesProtectedBranchMergeAccessLevel> mergeAccessLevels;
    private String name;
    private List<GetProjectProtectedBranchesProtectedBranchPushAccessLevel> pushAccessLevels;

    private GetProjectProtectedBranchesProtectedBranch() {}
    public Boolean allowForcePush() {
        return this.allowForcePush;
    }
    public Boolean codeOwnerApprovalRequired() {
        return this.codeOwnerApprovalRequired;
    }
    public Integer id() {
        return this.id;
    }
    public List<GetProjectProtectedBranchesProtectedBranchMergeAccessLevel> mergeAccessLevels() {
        return this.mergeAccessLevels;
    }
    public String name() {
        return this.name;
    }
    public List<GetProjectProtectedBranchesProtectedBranchPushAccessLevel> pushAccessLevels() {
        return this.pushAccessLevels;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectProtectedBranchesProtectedBranch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean allowForcePush;
        private Boolean codeOwnerApprovalRequired;
        private Integer id;
        private List<GetProjectProtectedBranchesProtectedBranchMergeAccessLevel> mergeAccessLevels;
        private String name;
        private List<GetProjectProtectedBranchesProtectedBranchPushAccessLevel> pushAccessLevels;
        public Builder() {}
        public Builder(GetProjectProtectedBranchesProtectedBranch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowForcePush = defaults.allowForcePush;
    	      this.codeOwnerApprovalRequired = defaults.codeOwnerApprovalRequired;
    	      this.id = defaults.id;
    	      this.mergeAccessLevels = defaults.mergeAccessLevels;
    	      this.name = defaults.name;
    	      this.pushAccessLevels = defaults.pushAccessLevels;
        }

        @CustomType.Setter
        public Builder allowForcePush(Boolean allowForcePush) {
            if (allowForcePush == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedBranchesProtectedBranch", "allowForcePush");
            }
            this.allowForcePush = allowForcePush;
            return this;
        }
        @CustomType.Setter
        public Builder codeOwnerApprovalRequired(Boolean codeOwnerApprovalRequired) {
            if (codeOwnerApprovalRequired == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedBranchesProtectedBranch", "codeOwnerApprovalRequired");
            }
            this.codeOwnerApprovalRequired = codeOwnerApprovalRequired;
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedBranchesProtectedBranch", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder mergeAccessLevels(List<GetProjectProtectedBranchesProtectedBranchMergeAccessLevel> mergeAccessLevels) {
            if (mergeAccessLevels == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedBranchesProtectedBranch", "mergeAccessLevels");
            }
            this.mergeAccessLevels = mergeAccessLevels;
            return this;
        }
        public Builder mergeAccessLevels(GetProjectProtectedBranchesProtectedBranchMergeAccessLevel... mergeAccessLevels) {
            return mergeAccessLevels(List.of(mergeAccessLevels));
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedBranchesProtectedBranch", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder pushAccessLevels(List<GetProjectProtectedBranchesProtectedBranchPushAccessLevel> pushAccessLevels) {
            if (pushAccessLevels == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedBranchesProtectedBranch", "pushAccessLevels");
            }
            this.pushAccessLevels = pushAccessLevels;
            return this;
        }
        public Builder pushAccessLevels(GetProjectProtectedBranchesProtectedBranchPushAccessLevel... pushAccessLevels) {
            return pushAccessLevels(List.of(pushAccessLevels));
        }
        public GetProjectProtectedBranchesProtectedBranch build() {
            final var _resultValue = new GetProjectProtectedBranchesProtectedBranch();
            _resultValue.allowForcePush = allowForcePush;
            _resultValue.codeOwnerApprovalRequired = codeOwnerApprovalRequired;
            _resultValue.id = id;
            _resultValue.mergeAccessLevels = mergeAccessLevels;
            _resultValue.name = name;
            _resultValue.pushAccessLevels = pushAccessLevels;
            return _resultValue;
        }
    }
}
