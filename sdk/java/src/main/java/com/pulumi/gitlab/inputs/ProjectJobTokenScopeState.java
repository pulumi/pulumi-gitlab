// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectJobTokenScopeState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectJobTokenScopeState Empty = new ProjectJobTokenScopeState();

    /**
     * The ID or full path of the project.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID or full path of the project.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The ID of the project that is in the CI/CD job token inbound allowlist.
     * 
     */
    @Import(name="targetProjectId")
    private @Nullable Output<Integer> targetProjectId;

    /**
     * @return The ID of the project that is in the CI/CD job token inbound allowlist.
     * 
     */
    public Optional<Output<Integer>> targetProjectId() {
        return Optional.ofNullable(this.targetProjectId);
    }

    private ProjectJobTokenScopeState() {}

    private ProjectJobTokenScopeState(ProjectJobTokenScopeState $) {
        this.project = $.project;
        this.targetProjectId = $.targetProjectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectJobTokenScopeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectJobTokenScopeState $;

        public Builder() {
            $ = new ProjectJobTokenScopeState();
        }

        public Builder(ProjectJobTokenScopeState defaults) {
            $ = new ProjectJobTokenScopeState(Objects.requireNonNull(defaults));
        }

        /**
         * @param project The ID or full path of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or full path of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param targetProjectId The ID of the project that is in the CI/CD job token inbound allowlist.
         * 
         * @return builder
         * 
         */
        public Builder targetProjectId(@Nullable Output<Integer> targetProjectId) {
            $.targetProjectId = targetProjectId;
            return this;
        }

        /**
         * @param targetProjectId The ID of the project that is in the CI/CD job token inbound allowlist.
         * 
         * @return builder
         * 
         */
        public Builder targetProjectId(Integer targetProjectId) {
            return targetProjectId(Output.of(targetProjectId));
        }

        public ProjectJobTokenScopeState build() {
            return $;
        }
    }

}