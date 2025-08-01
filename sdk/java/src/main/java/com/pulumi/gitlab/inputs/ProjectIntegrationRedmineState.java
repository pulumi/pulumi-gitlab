// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectIntegrationRedmineState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectIntegrationRedmineState Empty = new ProjectIntegrationRedmineState();

    /**
     * The URL to the Redmine project issue to link to this GitLab project.
     * 
     */
    @Import(name="issuesUrl")
    private @Nullable Output<String> issuesUrl;

    /**
     * @return The URL to the Redmine project issue to link to this GitLab project.
     * 
     */
    public Optional<Output<String>> issuesUrl() {
        return Optional.ofNullable(this.issuesUrl);
    }

    /**
     * The URL to use to create a new issue in the Redmine project linked to this GitLab project.
     * 
     */
    @Import(name="newIssueUrl")
    private @Nullable Output<String> newIssueUrl;

    /**
     * @return The URL to use to create a new issue in the Redmine project linked to this GitLab project.
     * 
     */
    public Optional<Output<String>> newIssueUrl() {
        return Optional.ofNullable(this.newIssueUrl);
    }

    /**
     * ID of the project you want to activate integration on.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return ID of the project you want to activate integration on.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The URL to the Redmine project to link to this GitLab project.
     * 
     */
    @Import(name="projectUrl")
    private @Nullable Output<String> projectUrl;

    /**
     * @return The URL to the Redmine project to link to this GitLab project.
     * 
     */
    public Optional<Output<String>> projectUrl() {
        return Optional.ofNullable(this.projectUrl);
    }

    /**
     * Indicates whether or not to inherit default settings. Defaults to false.
     * 
     */
    @Import(name="useInheritedSettings")
    private @Nullable Output<Boolean> useInheritedSettings;

    /**
     * @return Indicates whether or not to inherit default settings. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> useInheritedSettings() {
        return Optional.ofNullable(this.useInheritedSettings);
    }

    private ProjectIntegrationRedmineState() {}

    private ProjectIntegrationRedmineState(ProjectIntegrationRedmineState $) {
        this.issuesUrl = $.issuesUrl;
        this.newIssueUrl = $.newIssueUrl;
        this.project = $.project;
        this.projectUrl = $.projectUrl;
        this.useInheritedSettings = $.useInheritedSettings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectIntegrationRedmineState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectIntegrationRedmineState $;

        public Builder() {
            $ = new ProjectIntegrationRedmineState();
        }

        public Builder(ProjectIntegrationRedmineState defaults) {
            $ = new ProjectIntegrationRedmineState(Objects.requireNonNull(defaults));
        }

        /**
         * @param issuesUrl The URL to the Redmine project issue to link to this GitLab project.
         * 
         * @return builder
         * 
         */
        public Builder issuesUrl(@Nullable Output<String> issuesUrl) {
            $.issuesUrl = issuesUrl;
            return this;
        }

        /**
         * @param issuesUrl The URL to the Redmine project issue to link to this GitLab project.
         * 
         * @return builder
         * 
         */
        public Builder issuesUrl(String issuesUrl) {
            return issuesUrl(Output.of(issuesUrl));
        }

        /**
         * @param newIssueUrl The URL to use to create a new issue in the Redmine project linked to this GitLab project.
         * 
         * @return builder
         * 
         */
        public Builder newIssueUrl(@Nullable Output<String> newIssueUrl) {
            $.newIssueUrl = newIssueUrl;
            return this;
        }

        /**
         * @param newIssueUrl The URL to use to create a new issue in the Redmine project linked to this GitLab project.
         * 
         * @return builder
         * 
         */
        public Builder newIssueUrl(String newIssueUrl) {
            return newIssueUrl(Output.of(newIssueUrl));
        }

        /**
         * @param project ID of the project you want to activate integration on.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project ID of the project you want to activate integration on.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param projectUrl The URL to the Redmine project to link to this GitLab project.
         * 
         * @return builder
         * 
         */
        public Builder projectUrl(@Nullable Output<String> projectUrl) {
            $.projectUrl = projectUrl;
            return this;
        }

        /**
         * @param projectUrl The URL to the Redmine project to link to this GitLab project.
         * 
         * @return builder
         * 
         */
        public Builder projectUrl(String projectUrl) {
            return projectUrl(Output.of(projectUrl));
        }

        /**
         * @param useInheritedSettings Indicates whether or not to inherit default settings. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder useInheritedSettings(@Nullable Output<Boolean> useInheritedSettings) {
            $.useInheritedSettings = useInheritedSettings;
            return this;
        }

        /**
         * @param useInheritedSettings Indicates whether or not to inherit default settings. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder useInheritedSettings(Boolean useInheritedSettings) {
            return useInheritedSettings(Output.of(useInheritedSettings));
        }

        public ProjectIntegrationRedmineState build() {
            return $;
        }
    }

}
