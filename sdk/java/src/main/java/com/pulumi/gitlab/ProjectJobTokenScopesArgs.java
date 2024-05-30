// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;


public final class ProjectJobTokenScopesArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectJobTokenScopesArgs Empty = new ProjectJobTokenScopesArgs();

    /**
     * The ID of the project.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<Integer> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Output<Integer> projectId() {
        return this.projectId;
    }

    /**
     * A set of project IDs that are in the CI/CD job token inbound allowlist.
     * 
     */
    @Import(name="targetProjectIds", required=true)
    private Output<List<Integer>> targetProjectIds;

    /**
     * @return A set of project IDs that are in the CI/CD job token inbound allowlist.
     * 
     */
    public Output<List<Integer>> targetProjectIds() {
        return this.targetProjectIds;
    }

    private ProjectJobTokenScopesArgs() {}

    private ProjectJobTokenScopesArgs(ProjectJobTokenScopesArgs $) {
        this.projectId = $.projectId;
        this.targetProjectIds = $.targetProjectIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectJobTokenScopesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectJobTokenScopesArgs $;

        public Builder() {
            $ = new ProjectJobTokenScopesArgs();
        }

        public Builder(ProjectJobTokenScopesArgs defaults) {
            $ = new ProjectJobTokenScopesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<Integer> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Integer projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param targetProjectIds A set of project IDs that are in the CI/CD job token inbound allowlist.
         * 
         * @return builder
         * 
         */
        public Builder targetProjectIds(Output<List<Integer>> targetProjectIds) {
            $.targetProjectIds = targetProjectIds;
            return this;
        }

        /**
         * @param targetProjectIds A set of project IDs that are in the CI/CD job token inbound allowlist.
         * 
         * @return builder
         * 
         */
        public Builder targetProjectIds(List<Integer> targetProjectIds) {
            return targetProjectIds(Output.of(targetProjectIds));
        }

        /**
         * @param targetProjectIds A set of project IDs that are in the CI/CD job token inbound allowlist.
         * 
         * @return builder
         * 
         */
        public Builder targetProjectIds(Integer... targetProjectIds) {
            return targetProjectIds(List.of(targetProjectIds));
        }

        public ProjectJobTokenScopesArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ProjectJobTokenScopesArgs", "projectId");
            }
            if ($.targetProjectIds == null) {
                throw new MissingRequiredPropertyException("ProjectJobTokenScopesArgs", "targetProjectIds");
            }
            return $;
        }
    }

}
