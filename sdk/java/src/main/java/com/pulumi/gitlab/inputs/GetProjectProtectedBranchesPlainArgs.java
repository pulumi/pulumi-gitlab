// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.inputs.GetProjectProtectedBranchesProtectedBranch;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetProjectProtectedBranchesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectProtectedBranchesPlainArgs Empty = new GetProjectProtectedBranchesPlainArgs();

    /**
     * The integer or path with namespace that uniquely identifies the project.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The integer or path with namespace that uniquely identifies the project.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    /**
     * A list of protected branches, as defined below.
     * 
     */
    @Import(name="protectedBranches")
    private @Nullable List<GetProjectProtectedBranchesProtectedBranch> protectedBranches;

    /**
     * @return A list of protected branches, as defined below.
     * 
     */
    public Optional<List<GetProjectProtectedBranchesProtectedBranch>> protectedBranches() {
        return Optional.ofNullable(this.protectedBranches);
    }

    private GetProjectProtectedBranchesPlainArgs() {}

    private GetProjectProtectedBranchesPlainArgs(GetProjectProtectedBranchesPlainArgs $) {
        this.projectId = $.projectId;
        this.protectedBranches = $.protectedBranches;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectProtectedBranchesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectProtectedBranchesPlainArgs $;

        public Builder() {
            $ = new GetProjectProtectedBranchesPlainArgs();
        }

        public Builder(GetProjectProtectedBranchesPlainArgs defaults) {
            $ = new GetProjectProtectedBranchesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The integer or path with namespace that uniquely identifies the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param protectedBranches A list of protected branches, as defined below.
         * 
         * @return builder
         * 
         */
        public Builder protectedBranches(@Nullable List<GetProjectProtectedBranchesProtectedBranch> protectedBranches) {
            $.protectedBranches = protectedBranches;
            return this;
        }

        /**
         * @param protectedBranches A list of protected branches, as defined below.
         * 
         * @return builder
         * 
         */
        public Builder protectedBranches(GetProjectProtectedBranchesProtectedBranch... protectedBranches) {
            return protectedBranches(List.of(protectedBranches));
        }

        public GetProjectProtectedBranchesPlainArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetProjectProtectedBranchesPlainArgs", "projectId");
            }
            return $;
        }
    }

}
