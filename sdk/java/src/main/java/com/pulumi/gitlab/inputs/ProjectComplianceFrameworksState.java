// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectComplianceFrameworksState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectComplianceFrameworksState Empty = new ProjectComplianceFrameworksState();

    /**
     * Globally unique IDs of the compliance frameworks to assign to the project.
     * 
     */
    @Import(name="complianceFrameworkIds")
    private @Nullable Output<List<String>> complianceFrameworkIds;

    /**
     * @return Globally unique IDs of the compliance frameworks to assign to the project.
     * 
     */
    public Optional<Output<List<String>>> complianceFrameworkIds() {
        return Optional.ofNullable(this.complianceFrameworkIds);
    }

    /**
     * The ID or full path of the project to change the compliance frameworks of.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID or full path of the project to change the compliance frameworks of.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    private ProjectComplianceFrameworksState() {}

    private ProjectComplianceFrameworksState(ProjectComplianceFrameworksState $) {
        this.complianceFrameworkIds = $.complianceFrameworkIds;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectComplianceFrameworksState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectComplianceFrameworksState $;

        public Builder() {
            $ = new ProjectComplianceFrameworksState();
        }

        public Builder(ProjectComplianceFrameworksState defaults) {
            $ = new ProjectComplianceFrameworksState(Objects.requireNonNull(defaults));
        }

        /**
         * @param complianceFrameworkIds Globally unique IDs of the compliance frameworks to assign to the project.
         * 
         * @return builder
         * 
         */
        public Builder complianceFrameworkIds(@Nullable Output<List<String>> complianceFrameworkIds) {
            $.complianceFrameworkIds = complianceFrameworkIds;
            return this;
        }

        /**
         * @param complianceFrameworkIds Globally unique IDs of the compliance frameworks to assign to the project.
         * 
         * @return builder
         * 
         */
        public Builder complianceFrameworkIds(List<String> complianceFrameworkIds) {
            return complianceFrameworkIds(Output.of(complianceFrameworkIds));
        }

        /**
         * @param complianceFrameworkIds Globally unique IDs of the compliance frameworks to assign to the project.
         * 
         * @return builder
         * 
         */
        public Builder complianceFrameworkIds(String... complianceFrameworkIds) {
            return complianceFrameworkIds(List.of(complianceFrameworkIds));
        }

        /**
         * @param project The ID or full path of the project to change the compliance frameworks of.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or full path of the project to change the compliance frameworks of.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public ProjectComplianceFrameworksState build() {
            return $;
        }
    }

}
