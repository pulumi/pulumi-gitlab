// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gitlab.outputs.GetProjectBranchesBranch;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetProjectBranchesResult {
    /**
     * @return The list of branches of the project, as defined below.
     * 
     */
    private List<GetProjectBranchesBranch> branches;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    private String project;

    private GetProjectBranchesResult() {}
    /**
     * @return The list of branches of the project, as defined below.
     * 
     */
    public List<GetProjectBranchesBranch> branches() {
        return this.branches;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    public String project() {
        return this.project;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectBranchesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetProjectBranchesBranch> branches;
        private String id;
        private String project;
        public Builder() {}
        public Builder(GetProjectBranchesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.branches = defaults.branches;
    	      this.id = defaults.id;
    	      this.project = defaults.project;
        }

        @CustomType.Setter
        public Builder branches(List<GetProjectBranchesBranch> branches) {
            this.branches = Objects.requireNonNull(branches);
            return this;
        }
        public Builder branches(GetProjectBranchesBranch... branches) {
            return branches(List.of(branches));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            this.project = Objects.requireNonNull(project);
            return this;
        }
        public GetProjectBranchesResult build() {
            final var o = new GetProjectBranchesResult();
            o.branches = branches;
            o.id = id;
            o.project = project;
            return o;
        }
    }
}