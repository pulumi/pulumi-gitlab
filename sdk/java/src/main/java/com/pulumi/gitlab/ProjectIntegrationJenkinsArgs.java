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


public final class ProjectIntegrationJenkinsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectIntegrationJenkinsArgs Empty = new ProjectIntegrationJenkinsArgs();

    /**
     * Enable SSL verification. Defaults to `true` (enabled).
     * 
     */
    @Import(name="enableSslVerification")
    private @Nullable Output<Boolean> enableSslVerification;

    /**
     * @return Enable SSL verification. Defaults to `true` (enabled).
     * 
     */
    public Optional<Output<Boolean>> enableSslVerification() {
        return Optional.ofNullable(this.enableSslVerification);
    }

    /**
     * Jenkins URL like `http://jenkins.example.com`
     * 
     */
    @Import(name="jenkinsUrl", required=true)
    private Output<String> jenkinsUrl;

    /**
     * @return Jenkins URL like `http://jenkins.example.com`
     * 
     */
    public Output<String> jenkinsUrl() {
        return this.jenkinsUrl;
    }

    /**
     * Enable notifications for merge request events.
     * 
     */
    @Import(name="mergeRequestEvents")
    private @Nullable Output<Boolean> mergeRequestEvents;

    /**
     * @return Enable notifications for merge request events.
     * 
     */
    public Optional<Output<Boolean>> mergeRequestEvents() {
        return Optional.ofNullable(this.mergeRequestEvents);
    }

    /**
     * Password for authentication with the Jenkins server, if authentication is required by the server.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Password for authentication with the Jenkins server, if authentication is required by the server.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
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
     * The URL-friendly project name. Example: `my_project_name`.
     * 
     */
    @Import(name="projectName", required=true)
    private Output<String> projectName;

    /**
     * @return The URL-friendly project name. Example: `my_project_name`.
     * 
     */
    public Output<String> projectName() {
        return this.projectName;
    }

    /**
     * Enable notifications for push events.
     * 
     */
    @Import(name="pushEvents")
    private @Nullable Output<Boolean> pushEvents;

    /**
     * @return Enable notifications for push events.
     * 
     */
    public Optional<Output<Boolean>> pushEvents() {
        return Optional.ofNullable(this.pushEvents);
    }

    /**
     * Enable notifications for tag push events.
     * 
     */
    @Import(name="tagPushEvents")
    private @Nullable Output<Boolean> tagPushEvents;

    /**
     * @return Enable notifications for tag push events.
     * 
     */
    public Optional<Output<Boolean>> tagPushEvents() {
        return Optional.ofNullable(this.tagPushEvents);
    }

    /**
     * Username for authentication with the Jenkins server, if authentication is required by the server.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return Username for authentication with the Jenkins server, if authentication is required by the server.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ProjectIntegrationJenkinsArgs() {}

    private ProjectIntegrationJenkinsArgs(ProjectIntegrationJenkinsArgs $) {
        this.enableSslVerification = $.enableSslVerification;
        this.jenkinsUrl = $.jenkinsUrl;
        this.mergeRequestEvents = $.mergeRequestEvents;
        this.password = $.password;
        this.project = $.project;
        this.projectName = $.projectName;
        this.pushEvents = $.pushEvents;
        this.tagPushEvents = $.tagPushEvents;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectIntegrationJenkinsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectIntegrationJenkinsArgs $;

        public Builder() {
            $ = new ProjectIntegrationJenkinsArgs();
        }

        public Builder(ProjectIntegrationJenkinsArgs defaults) {
            $ = new ProjectIntegrationJenkinsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableSslVerification Enable SSL verification. Defaults to `true` (enabled).
         * 
         * @return builder
         * 
         */
        public Builder enableSslVerification(@Nullable Output<Boolean> enableSslVerification) {
            $.enableSslVerification = enableSslVerification;
            return this;
        }

        /**
         * @param enableSslVerification Enable SSL verification. Defaults to `true` (enabled).
         * 
         * @return builder
         * 
         */
        public Builder enableSslVerification(Boolean enableSslVerification) {
            return enableSslVerification(Output.of(enableSslVerification));
        }

        /**
         * @param jenkinsUrl Jenkins URL like `http://jenkins.example.com`
         * 
         * @return builder
         * 
         */
        public Builder jenkinsUrl(Output<String> jenkinsUrl) {
            $.jenkinsUrl = jenkinsUrl;
            return this;
        }

        /**
         * @param jenkinsUrl Jenkins URL like `http://jenkins.example.com`
         * 
         * @return builder
         * 
         */
        public Builder jenkinsUrl(String jenkinsUrl) {
            return jenkinsUrl(Output.of(jenkinsUrl));
        }

        /**
         * @param mergeRequestEvents Enable notifications for merge request events.
         * 
         * @return builder
         * 
         */
        public Builder mergeRequestEvents(@Nullable Output<Boolean> mergeRequestEvents) {
            $.mergeRequestEvents = mergeRequestEvents;
            return this;
        }

        /**
         * @param mergeRequestEvents Enable notifications for merge request events.
         * 
         * @return builder
         * 
         */
        public Builder mergeRequestEvents(Boolean mergeRequestEvents) {
            return mergeRequestEvents(Output.of(mergeRequestEvents));
        }

        /**
         * @param password Password for authentication with the Jenkins server, if authentication is required by the server.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Password for authentication with the Jenkins server, if authentication is required by the server.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
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
         * @param projectName The URL-friendly project name. Example: `my_project_name`.
         * 
         * @return builder
         * 
         */
        public Builder projectName(Output<String> projectName) {
            $.projectName = projectName;
            return this;
        }

        /**
         * @param projectName The URL-friendly project name. Example: `my_project_name`.
         * 
         * @return builder
         * 
         */
        public Builder projectName(String projectName) {
            return projectName(Output.of(projectName));
        }

        /**
         * @param pushEvents Enable notifications for push events.
         * 
         * @return builder
         * 
         */
        public Builder pushEvents(@Nullable Output<Boolean> pushEvents) {
            $.pushEvents = pushEvents;
            return this;
        }

        /**
         * @param pushEvents Enable notifications for push events.
         * 
         * @return builder
         * 
         */
        public Builder pushEvents(Boolean pushEvents) {
            return pushEvents(Output.of(pushEvents));
        }

        /**
         * @param tagPushEvents Enable notifications for tag push events.
         * 
         * @return builder
         * 
         */
        public Builder tagPushEvents(@Nullable Output<Boolean> tagPushEvents) {
            $.tagPushEvents = tagPushEvents;
            return this;
        }

        /**
         * @param tagPushEvents Enable notifications for tag push events.
         * 
         * @return builder
         * 
         */
        public Builder tagPushEvents(Boolean tagPushEvents) {
            return tagPushEvents(Output.of(tagPushEvents));
        }

        /**
         * @param username Username for authentication with the Jenkins server, if authentication is required by the server.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Username for authentication with the Jenkins server, if authentication is required by the server.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ProjectIntegrationJenkinsArgs build() {
            if ($.jenkinsUrl == null) {
                throw new MissingRequiredPropertyException("ProjectIntegrationJenkinsArgs", "jenkinsUrl");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("ProjectIntegrationJenkinsArgs", "project");
            }
            if ($.projectName == null) {
                throw new MissingRequiredPropertyException("ProjectIntegrationJenkinsArgs", "projectName");
            }
            return $;
        }
    }

}
