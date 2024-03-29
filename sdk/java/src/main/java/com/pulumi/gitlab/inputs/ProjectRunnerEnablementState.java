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


public final class ProjectRunnerEnablementState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectRunnerEnablementState Empty = new ProjectRunnerEnablementState();

    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The ID of a runner to enable for the project.
     * 
     */
    @Import(name="runnerId")
    private @Nullable Output<Integer> runnerId;

    /**
     * @return The ID of a runner to enable for the project.
     * 
     */
    public Optional<Output<Integer>> runnerId() {
        return Optional.ofNullable(this.runnerId);
    }

    private ProjectRunnerEnablementState() {}

    private ProjectRunnerEnablementState(ProjectRunnerEnablementState $) {
        this.project = $.project;
        this.runnerId = $.runnerId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectRunnerEnablementState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectRunnerEnablementState $;

        public Builder() {
            $ = new ProjectRunnerEnablementState();
        }

        public Builder(ProjectRunnerEnablementState defaults) {
            $ = new ProjectRunnerEnablementState(Objects.requireNonNull(defaults));
        }

        /**
         * @param project The ID or URL-encoded path of the project owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or URL-encoded path of the project owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param runnerId The ID of a runner to enable for the project.
         * 
         * @return builder
         * 
         */
        public Builder runnerId(@Nullable Output<Integer> runnerId) {
            $.runnerId = runnerId;
            return this;
        }

        /**
         * @param runnerId The ID of a runner to enable for the project.
         * 
         * @return builder
         * 
         */
        public Builder runnerId(Integer runnerId) {
            return runnerId(Output.of(runnerId));
        }

        public ProjectRunnerEnablementState build() {
            return $;
        }
    }

}
