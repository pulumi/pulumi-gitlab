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
import javax.annotation.Nullable;

@CustomType
public final class GetProjectProtectedBranchesProtectedBranch {
    /**
     * @return Whether force push is allowed.
     * 
     */
    private Boolean allowForcePush;
    /**
     * @return Reject code pushes that change files listed in the CODEOWNERS file.
     * 
     */
    private Boolean codeOwnerApprovalRequired;
    /**
     * @return The ID of this resource.
     * 
     */
    private Integer id;
    /**
     * @return Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     * 
     */
    private @Nullable List<GetProjectProtectedBranchesProtectedBranchMergeAccessLevel> mergeAccessLevels;
    /**
     * @return The name of the protected branch.
     * 
     */
    private String name;
    /**
     * @return Array of access levels and user(s)/group(s) allowed to push to protected branch.
     * 
     */
    private @Nullable List<GetProjectProtectedBranchesProtectedBranchPushAccessLevel> pushAccessLevels;

    private GetProjectProtectedBranchesProtectedBranch() {}
    /**
     * @return Whether force push is allowed.
     * 
     */
    public Boolean allowForcePush() {
        return this.allowForcePush;
    }
    /**
     * @return Reject code pushes that change files listed in the CODEOWNERS file.
     * 
     */
    public Boolean codeOwnerApprovalRequired() {
        return this.codeOwnerApprovalRequired;
    }
    /**
     * @return The ID of this resource.
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     * 
     */
    public List<GetProjectProtectedBranchesProtectedBranchMergeAccessLevel> mergeAccessLevels() {
        return this.mergeAccessLevels == null ? List.of() : this.mergeAccessLevels;
    }
    /**
     * @return The name of the protected branch.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Array of access levels and user(s)/group(s) allowed to push to protected branch.
     * 
     */
    public List<GetProjectProtectedBranchesProtectedBranchPushAccessLevel> pushAccessLevels() {
        return this.pushAccessLevels == null ? List.of() : this.pushAccessLevels;
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
        private @Nullable List<GetProjectProtectedBranchesProtectedBranchMergeAccessLevel> mergeAccessLevels;
        private String name;
        private @Nullable List<GetProjectProtectedBranchesProtectedBranchPushAccessLevel> pushAccessLevels;
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
        public Builder mergeAccessLevels(@Nullable List<GetProjectProtectedBranchesProtectedBranchMergeAccessLevel> mergeAccessLevels) {

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
        public Builder pushAccessLevels(@Nullable List<GetProjectProtectedBranchesProtectedBranchPushAccessLevel> pushAccessLevels) {

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
