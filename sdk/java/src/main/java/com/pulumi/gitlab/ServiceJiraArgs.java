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


public final class ServiceJiraArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceJiraArgs Empty = new ServiceJiraArgs();

    /**
     * The base URL to the Jira instance API. Web URL value is used if not set. For example, [https://jira-api.example.com](https://jira-api.example.com).
     * 
     */
    @Import(name="apiUrl")
    private @Nullable Output<String> apiUrl;

    /**
     * @return The base URL to the Jira instance API. Web URL value is used if not set. For example, [https://jira-api.example.com](https://jira-api.example.com).
     * 
     */
    public Optional<Output<String>> apiUrl() {
        return Optional.ofNullable(this.apiUrl);
    }

    /**
     * Enable comments inside Jira issues on each GitLab event (commit / merge request)
     * 
     */
    @Import(name="commentOnEventEnabled")
    private @Nullable Output<Boolean> commentOnEventEnabled;

    /**
     * @return Enable comments inside Jira issues on each GitLab event (commit / merge request)
     * 
     */
    public Optional<Output<Boolean>> commentOnEventEnabled() {
        return Optional.ofNullable(this.commentOnEventEnabled);
    }

    /**
     * Enable notifications for commit events
     * 
     */
    @Import(name="commitEvents")
    private @Nullable Output<Boolean> commitEvents;

    /**
     * @return Enable notifications for commit events
     * 
     */
    public Optional<Output<Boolean>> commitEvents() {
        return Optional.ofNullable(this.commitEvents);
    }

    /**
     * Enable notifications for issues events.
     * 
     */
    @Import(name="issuesEvents")
    private @Nullable Output<Boolean> issuesEvents;

    /**
     * @return Enable notifications for issues events.
     * 
     */
    public Optional<Output<Boolean>> issuesEvents() {
        return Optional.ofNullable(this.issuesEvents);
    }

    /**
     * The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
     * 
     */
    @Import(name="jiraIssueTransitionId")
    private @Nullable Output<String> jiraIssueTransitionId;

    /**
     * @return The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
     * 
     */
    public Optional<Output<String>> jiraIssueTransitionId() {
        return Optional.ofNullable(this.jiraIssueTransitionId);
    }

    /**
     * Enable notifications for job events.
     * 
     */
    @Import(name="jobEvents")
    private @Nullable Output<Boolean> jobEvents;

    /**
     * @return Enable notifications for job events.
     * 
     */
    public Optional<Output<Boolean>> jobEvents() {
        return Optional.ofNullable(this.jobEvents);
    }

    /**
     * Enable notifications for merge request events
     * 
     */
    @Import(name="mergeRequestsEvents")
    private @Nullable Output<Boolean> mergeRequestsEvents;

    /**
     * @return Enable notifications for merge request events
     * 
     */
    public Optional<Output<Boolean>> mergeRequestsEvents() {
        return Optional.ofNullable(this.mergeRequestsEvents);
    }

    /**
     * Enable notifications for note events.
     * 
     */
    @Import(name="noteEvents")
    private @Nullable Output<Boolean> noteEvents;

    /**
     * @return Enable notifications for note events.
     * 
     */
    public Optional<Output<Boolean>> noteEvents() {
        return Optional.ofNullable(this.noteEvents);
    }

    /**
     * The password of the user created to be used with GitLab/JIRA.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The password of the user created to be used with GitLab/JIRA.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    /**
     * Enable notifications for pipeline events.
     * 
     */
    @Import(name="pipelineEvents")
    private @Nullable Output<Boolean> pipelineEvents;

    /**
     * @return Enable notifications for pipeline events.
     * 
     */
    public Optional<Output<Boolean>> pipelineEvents() {
        return Optional.ofNullable(this.pipelineEvents);
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
     * The short identifier for your JIRA project, all uppercase, e.g., PROJ.
     * 
     */
    @Import(name="projectKey")
    private @Nullable Output<String> projectKey;

    /**
     * @return The short identifier for your JIRA project, all uppercase, e.g., PROJ.
     * 
     */
    public Optional<Output<String>> projectKey() {
        return Optional.ofNullable(this.projectKey);
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
     * Enable notifications for tag_push events.
     * 
     */
    @Import(name="tagPushEvents")
    private @Nullable Output<Boolean> tagPushEvents;

    /**
     * @return Enable notifications for tag_push events.
     * 
     */
    public Optional<Output<Boolean>> tagPushEvents() {
        return Optional.ofNullable(this.tagPushEvents);
    }

    /**
     * The URL to the JIRA project which is being linked to this GitLab project. For example, [https://jira.example.com](https://jira.example.com).
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The URL to the JIRA project which is being linked to this GitLab project. For example, [https://jira.example.com](https://jira.example.com).
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * The username of the user created to be used with GitLab/JIRA.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username of the user created to be used with GitLab/JIRA.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private ServiceJiraArgs() {}

    private ServiceJiraArgs(ServiceJiraArgs $) {
        this.apiUrl = $.apiUrl;
        this.commentOnEventEnabled = $.commentOnEventEnabled;
        this.commitEvents = $.commitEvents;
        this.issuesEvents = $.issuesEvents;
        this.jiraIssueTransitionId = $.jiraIssueTransitionId;
        this.jobEvents = $.jobEvents;
        this.mergeRequestsEvents = $.mergeRequestsEvents;
        this.noteEvents = $.noteEvents;
        this.password = $.password;
        this.pipelineEvents = $.pipelineEvents;
        this.project = $.project;
        this.projectKey = $.projectKey;
        this.pushEvents = $.pushEvents;
        this.tagPushEvents = $.tagPushEvents;
        this.url = $.url;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceJiraArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceJiraArgs $;

        public Builder() {
            $ = new ServiceJiraArgs();
        }

        public Builder(ServiceJiraArgs defaults) {
            $ = new ServiceJiraArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiUrl The base URL to the Jira instance API. Web URL value is used if not set. For example, [https://jira-api.example.com](https://jira-api.example.com).
         * 
         * @return builder
         * 
         */
        public Builder apiUrl(@Nullable Output<String> apiUrl) {
            $.apiUrl = apiUrl;
            return this;
        }

        /**
         * @param apiUrl The base URL to the Jira instance API. Web URL value is used if not set. For example, [https://jira-api.example.com](https://jira-api.example.com).
         * 
         * @return builder
         * 
         */
        public Builder apiUrl(String apiUrl) {
            return apiUrl(Output.of(apiUrl));
        }

        /**
         * @param commentOnEventEnabled Enable comments inside Jira issues on each GitLab event (commit / merge request)
         * 
         * @return builder
         * 
         */
        public Builder commentOnEventEnabled(@Nullable Output<Boolean> commentOnEventEnabled) {
            $.commentOnEventEnabled = commentOnEventEnabled;
            return this;
        }

        /**
         * @param commentOnEventEnabled Enable comments inside Jira issues on each GitLab event (commit / merge request)
         * 
         * @return builder
         * 
         */
        public Builder commentOnEventEnabled(Boolean commentOnEventEnabled) {
            return commentOnEventEnabled(Output.of(commentOnEventEnabled));
        }

        /**
         * @param commitEvents Enable notifications for commit events
         * 
         * @return builder
         * 
         */
        public Builder commitEvents(@Nullable Output<Boolean> commitEvents) {
            $.commitEvents = commitEvents;
            return this;
        }

        /**
         * @param commitEvents Enable notifications for commit events
         * 
         * @return builder
         * 
         */
        public Builder commitEvents(Boolean commitEvents) {
            return commitEvents(Output.of(commitEvents));
        }

        /**
         * @param issuesEvents Enable notifications for issues events.
         * 
         * @return builder
         * 
         */
        public Builder issuesEvents(@Nullable Output<Boolean> issuesEvents) {
            $.issuesEvents = issuesEvents;
            return this;
        }

        /**
         * @param issuesEvents Enable notifications for issues events.
         * 
         * @return builder
         * 
         */
        public Builder issuesEvents(Boolean issuesEvents) {
            return issuesEvents(Output.of(issuesEvents));
        }

        /**
         * @param jiraIssueTransitionId The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
         * 
         * @return builder
         * 
         */
        public Builder jiraIssueTransitionId(@Nullable Output<String> jiraIssueTransitionId) {
            $.jiraIssueTransitionId = jiraIssueTransitionId;
            return this;
        }

        /**
         * @param jiraIssueTransitionId The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
         * 
         * @return builder
         * 
         */
        public Builder jiraIssueTransitionId(String jiraIssueTransitionId) {
            return jiraIssueTransitionId(Output.of(jiraIssueTransitionId));
        }

        /**
         * @param jobEvents Enable notifications for job events.
         * 
         * @return builder
         * 
         */
        public Builder jobEvents(@Nullable Output<Boolean> jobEvents) {
            $.jobEvents = jobEvents;
            return this;
        }

        /**
         * @param jobEvents Enable notifications for job events.
         * 
         * @return builder
         * 
         */
        public Builder jobEvents(Boolean jobEvents) {
            return jobEvents(Output.of(jobEvents));
        }

        /**
         * @param mergeRequestsEvents Enable notifications for merge request events
         * 
         * @return builder
         * 
         */
        public Builder mergeRequestsEvents(@Nullable Output<Boolean> mergeRequestsEvents) {
            $.mergeRequestsEvents = mergeRequestsEvents;
            return this;
        }

        /**
         * @param mergeRequestsEvents Enable notifications for merge request events
         * 
         * @return builder
         * 
         */
        public Builder mergeRequestsEvents(Boolean mergeRequestsEvents) {
            return mergeRequestsEvents(Output.of(mergeRequestsEvents));
        }

        /**
         * @param noteEvents Enable notifications for note events.
         * 
         * @return builder
         * 
         */
        public Builder noteEvents(@Nullable Output<Boolean> noteEvents) {
            $.noteEvents = noteEvents;
            return this;
        }

        /**
         * @param noteEvents Enable notifications for note events.
         * 
         * @return builder
         * 
         */
        public Builder noteEvents(Boolean noteEvents) {
            return noteEvents(Output.of(noteEvents));
        }

        /**
         * @param password The password of the user created to be used with GitLab/JIRA.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password of the user created to be used with GitLab/JIRA.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param pipelineEvents Enable notifications for pipeline events.
         * 
         * @return builder
         * 
         */
        public Builder pipelineEvents(@Nullable Output<Boolean> pipelineEvents) {
            $.pipelineEvents = pipelineEvents;
            return this;
        }

        /**
         * @param pipelineEvents Enable notifications for pipeline events.
         * 
         * @return builder
         * 
         */
        public Builder pipelineEvents(Boolean pipelineEvents) {
            return pipelineEvents(Output.of(pipelineEvents));
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
         * @param projectKey The short identifier for your JIRA project, all uppercase, e.g., PROJ.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(@Nullable Output<String> projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        /**
         * @param projectKey The short identifier for your JIRA project, all uppercase, e.g., PROJ.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(String projectKey) {
            return projectKey(Output.of(projectKey));
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
         * @param tagPushEvents Enable notifications for tag_push events.
         * 
         * @return builder
         * 
         */
        public Builder tagPushEvents(@Nullable Output<Boolean> tagPushEvents) {
            $.tagPushEvents = tagPushEvents;
            return this;
        }

        /**
         * @param tagPushEvents Enable notifications for tag_push events.
         * 
         * @return builder
         * 
         */
        public Builder tagPushEvents(Boolean tagPushEvents) {
            return tagPushEvents(Output.of(tagPushEvents));
        }

        /**
         * @param url The URL to the JIRA project which is being linked to this GitLab project. For example, [https://jira.example.com](https://jira.example.com).
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL to the JIRA project which is being linked to this GitLab project. For example, [https://jira.example.com](https://jira.example.com).
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param username The username of the user created to be used with GitLab/JIRA.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the user created to be used with GitLab/JIRA.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceJiraArgs build() {
            if ($.password == null) {
                throw new MissingRequiredPropertyException("ServiceJiraArgs", "password");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("ServiceJiraArgs", "project");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("ServiceJiraArgs", "url");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("ServiceJiraArgs", "username");
            }
            return $;
        }
    }

}
