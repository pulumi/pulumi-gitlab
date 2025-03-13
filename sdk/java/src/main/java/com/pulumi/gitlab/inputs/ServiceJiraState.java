// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceJiraState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceJiraState Empty = new ServiceJiraState();

    /**
     * Whether the integration is active.
     * 
     */
    @Import(name="active")
    private @Nullable Output<Boolean> active;

    /**
     * @return Whether the integration is active.
     * 
     */
    public Optional<Output<Boolean>> active() {
        return Optional.ofNullable(this.active);
    }

    /**
     * The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
     * 
     */
    @Import(name="apiUrl")
    private @Nullable Output<String> apiUrl;

    /**
     * @return The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
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
     * Create time.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Create time.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Enable viewing Jira issues in GitLab.
     * 
     */
    @Import(name="issuesEnabled")
    private @Nullable Output<Boolean> issuesEnabled;

    /**
     * @return Enable viewing Jira issues in GitLab.
     * 
     */
    public Optional<Output<Boolean>> issuesEnabled() {
        return Optional.ofNullable(this.issuesEnabled);
    }

    /**
     * The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
     * 
     */
    @Import(name="jiraAuthType")
    private @Nullable Output<Integer> jiraAuthType;

    /**
     * @return The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
     * 
     */
    public Optional<Output<Integer>> jiraAuthType() {
        return Optional.ofNullable(this.jiraAuthType);
    }

    /**
     * Prefix to match Jira issue keys.
     * 
     */
    @Import(name="jiraIssuePrefix")
    private @Nullable Output<String> jiraIssuePrefix;

    /**
     * @return Prefix to match Jira issue keys.
     * 
     */
    public Optional<Output<String>> jiraIssuePrefix() {
        return Optional.ofNullable(this.jiraIssuePrefix);
    }

    /**
     * Regular expression to match Jira issue keys.
     * 
     */
    @Import(name="jiraIssueRegex")
    private @Nullable Output<String> jiraIssueRegex;

    /**
     * @return Regular expression to match Jira issue keys.
     * 
     */
    public Optional<Output<String>> jiraIssueRegex() {
        return Optional.ofNullable(this.jiraIssueRegex);
    }

    @Import(name="jiraIssueTransitionAutomatic")
    private @Nullable Output<Boolean> jiraIssueTransitionAutomatic;

    public Optional<Output<Boolean>> jiraIssueTransitionAutomatic() {
        return Optional.ofNullable(this.jiraIssueTransitionAutomatic);
    }

    /**
     * The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
     * 
     */
    @Import(name="jiraIssueTransitionId")
    private @Nullable Output<String> jiraIssueTransitionId;

    /**
     * @return The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
     * 
     */
    public Optional<Output<String>> jiraIssueTransitionId() {
        return Optional.ofNullable(this.jiraIssueTransitionId);
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
     * The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
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
     * Keys of Jira projects. When issues_enabled is true, this setting specifies which Jira projects to view issues from in GitLab.
     * 
     */
    @Import(name="projectKeys")
    private @Nullable Output<List<String>> projectKeys;

    /**
     * @return Keys of Jira projects. When issues_enabled is true, this setting specifies which Jira projects to view issues from in GitLab.
     * 
     */
    public Optional<Output<List<String>>> projectKeys() {
        return Optional.ofNullable(this.projectKeys);
    }

    /**
     * Title.
     * 
     */
    @Import(name="title")
    private @Nullable Output<String> title;

    /**
     * @return Title.
     * 
     */
    public Optional<Output<String>> title() {
        return Optional.ofNullable(this.title);
    }

    /**
     * Update time.
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return Update time.
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    /**
     * The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
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

    /**
     * The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ServiceJiraState() {}

    private ServiceJiraState(ServiceJiraState $) {
        this.active = $.active;
        this.apiUrl = $.apiUrl;
        this.commentOnEventEnabled = $.commentOnEventEnabled;
        this.commitEvents = $.commitEvents;
        this.createdAt = $.createdAt;
        this.issuesEnabled = $.issuesEnabled;
        this.jiraAuthType = $.jiraAuthType;
        this.jiraIssuePrefix = $.jiraIssuePrefix;
        this.jiraIssueRegex = $.jiraIssueRegex;
        this.jiraIssueTransitionAutomatic = $.jiraIssueTransitionAutomatic;
        this.jiraIssueTransitionId = $.jiraIssueTransitionId;
        this.mergeRequestsEvents = $.mergeRequestsEvents;
        this.password = $.password;
        this.project = $.project;
        this.projectKey = $.projectKey;
        this.projectKeys = $.projectKeys;
        this.title = $.title;
        this.updatedAt = $.updatedAt;
        this.url = $.url;
        this.useInheritedSettings = $.useInheritedSettings;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceJiraState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceJiraState $;

        public Builder() {
            $ = new ServiceJiraState();
        }

        public Builder(ServiceJiraState defaults) {
            $ = new ServiceJiraState(Objects.requireNonNull(defaults));
        }

        /**
         * @param active Whether the integration is active.
         * 
         * @return builder
         * 
         */
        public Builder active(@Nullable Output<Boolean> active) {
            $.active = active;
            return this;
        }

        /**
         * @param active Whether the integration is active.
         * 
         * @return builder
         * 
         */
        public Builder active(Boolean active) {
            return active(Output.of(active));
        }

        /**
         * @param apiUrl The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
         * 
         * @return builder
         * 
         */
        public Builder apiUrl(@Nullable Output<String> apiUrl) {
            $.apiUrl = apiUrl;
            return this;
        }

        /**
         * @param apiUrl The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
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
         * @param createdAt Create time.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Create time.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param issuesEnabled Enable viewing Jira issues in GitLab.
         * 
         * @return builder
         * 
         */
        public Builder issuesEnabled(@Nullable Output<Boolean> issuesEnabled) {
            $.issuesEnabled = issuesEnabled;
            return this;
        }

        /**
         * @param issuesEnabled Enable viewing Jira issues in GitLab.
         * 
         * @return builder
         * 
         */
        public Builder issuesEnabled(Boolean issuesEnabled) {
            return issuesEnabled(Output.of(issuesEnabled));
        }

        /**
         * @param jiraAuthType The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
         * 
         * @return builder
         * 
         */
        public Builder jiraAuthType(@Nullable Output<Integer> jiraAuthType) {
            $.jiraAuthType = jiraAuthType;
            return this;
        }

        /**
         * @param jiraAuthType The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
         * 
         * @return builder
         * 
         */
        public Builder jiraAuthType(Integer jiraAuthType) {
            return jiraAuthType(Output.of(jiraAuthType));
        }

        /**
         * @param jiraIssuePrefix Prefix to match Jira issue keys.
         * 
         * @return builder
         * 
         */
        public Builder jiraIssuePrefix(@Nullable Output<String> jiraIssuePrefix) {
            $.jiraIssuePrefix = jiraIssuePrefix;
            return this;
        }

        /**
         * @param jiraIssuePrefix Prefix to match Jira issue keys.
         * 
         * @return builder
         * 
         */
        public Builder jiraIssuePrefix(String jiraIssuePrefix) {
            return jiraIssuePrefix(Output.of(jiraIssuePrefix));
        }

        /**
         * @param jiraIssueRegex Regular expression to match Jira issue keys.
         * 
         * @return builder
         * 
         */
        public Builder jiraIssueRegex(@Nullable Output<String> jiraIssueRegex) {
            $.jiraIssueRegex = jiraIssueRegex;
            return this;
        }

        /**
         * @param jiraIssueRegex Regular expression to match Jira issue keys.
         * 
         * @return builder
         * 
         */
        public Builder jiraIssueRegex(String jiraIssueRegex) {
            return jiraIssueRegex(Output.of(jiraIssueRegex));
        }

        public Builder jiraIssueTransitionAutomatic(@Nullable Output<Boolean> jiraIssueTransitionAutomatic) {
            $.jiraIssueTransitionAutomatic = jiraIssueTransitionAutomatic;
            return this;
        }

        public Builder jiraIssueTransitionAutomatic(Boolean jiraIssueTransitionAutomatic) {
            return jiraIssueTransitionAutomatic(Output.of(jiraIssueTransitionAutomatic));
        }

        /**
         * @param jiraIssueTransitionId The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
         * 
         * @return builder
         * 
         */
        public Builder jiraIssueTransitionId(@Nullable Output<String> jiraIssueTransitionId) {
            $.jiraIssueTransitionId = jiraIssueTransitionId;
            return this;
        }

        /**
         * @param jiraIssueTransitionId The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
         * 
         * @return builder
         * 
         */
        public Builder jiraIssueTransitionId(String jiraIssueTransitionId) {
            return jiraIssueTransitionId(Output.of(jiraIssueTransitionId));
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
         * @param password The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
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
         * @param projectKeys Keys of Jira projects. When issues_enabled is true, this setting specifies which Jira projects to view issues from in GitLab.
         * 
         * @return builder
         * 
         */
        public Builder projectKeys(@Nullable Output<List<String>> projectKeys) {
            $.projectKeys = projectKeys;
            return this;
        }

        /**
         * @param projectKeys Keys of Jira projects. When issues_enabled is true, this setting specifies which Jira projects to view issues from in GitLab.
         * 
         * @return builder
         * 
         */
        public Builder projectKeys(List<String> projectKeys) {
            return projectKeys(Output.of(projectKeys));
        }

        /**
         * @param projectKeys Keys of Jira projects. When issues_enabled is true, this setting specifies which Jira projects to view issues from in GitLab.
         * 
         * @return builder
         * 
         */
        public Builder projectKeys(String... projectKeys) {
            return projectKeys(List.of(projectKeys));
        }

        /**
         * @param title Title.
         * 
         * @return builder
         * 
         */
        public Builder title(@Nullable Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title Title.
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        /**
         * @param updatedAt Update time.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt Update time.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        /**
         * @param url The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
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

        /**
         * @param username The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ServiceJiraState build() {
            return $;
        }
    }

}
