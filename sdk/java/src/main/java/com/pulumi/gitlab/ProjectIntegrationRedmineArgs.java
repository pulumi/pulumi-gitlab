// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectIntegrationRedmineArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectIntegrationRedmineArgs Empty = new ProjectIntegrationRedmineArgs();

    /**
     * The URL to the Redmine project issue to link to this GitLab project.
     * 
     */
    @Import(name="issuesUrl", required=true)
    private Output<String> issuesUrl;

    /**
     * @return The URL to the Redmine project issue to link to this GitLab project.
     * 
     */
    public Output<String> issuesUrl() {
        return this.issuesUrl;
    }

    /**
     * The URL to use to create a new issue in the Redmine project linked to this GitLab project.
     * 
     */
    @Import(name="newIssueUrl", required=true)
    private Output<String> newIssueUrl;

    /**
     * @return The URL to use to create a new issue in the Redmine project linked to this GitLab project.
     * 
     */
    public Output<String> newIssueUrl() {
        return this.newIssueUrl;
    }

    /**
     * ID of the project you want to activate integration on.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return ID of the project you want to activate integration on.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * The URL to the Redmine project to link to this GitLab project.
     * 
     */
    @Import(name="projectUrl", required=true)
    private Output<String> projectUrl;

    /**
     * @return The URL to the Redmine project to link to this GitLab project.
     * 
     */
    public Output<String> projectUrl() {
        return this.projectUrl;
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

    private ProjectIntegrationRedmineArgs() {}

    private ProjectIntegrationRedmineArgs(ProjectIntegrationRedmineArgs $) {
        this.issuesUrl = $.issuesUrl;
        this.newIssueUrl = $.newIssueUrl;
        this.project = $.project;
        this.projectUrl = $.projectUrl;
        this.useInheritedSettings = $.useInheritedSettings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectIntegrationRedmineArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectIntegrationRedmineArgs $;

        public Builder() {
            $ = new ProjectIntegrationRedmineArgs();
        }

        public Builder(ProjectIntegrationRedmineArgs defaults) {
            $ = new ProjectIntegrationRedmineArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param issuesUrl The URL to the Redmine project issue to link to this GitLab project.
         * 
         * @return builder
         * 
         */
        public Builder issuesUrl(Output<String> issuesUrl) {
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
        public Builder newIssueUrl(Output<String> newIssueUrl) {
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
        public Builder project(Output<String> project) {
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
        public Builder projectUrl(Output<String> projectUrl) {
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

        public ProjectIntegrationRedmineArgs build() {
            if ($.issuesUrl == null) {
                throw new MissingRequiredPropertyException("ProjectIntegrationRedmineArgs", "issuesUrl");
            }
            if ($.newIssueUrl == null) {
                throw new MissingRequiredPropertyException("ProjectIntegrationRedmineArgs", "newIssueUrl");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("ProjectIntegrationRedmineArgs", "project");
            }
            if ($.projectUrl == null) {
                throw new MissingRequiredPropertyException("ProjectIntegrationRedmineArgs", "projectUrl");
            }
            return $;
        }
    }

}
