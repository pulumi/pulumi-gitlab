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


public final class IntegrationSlackState extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationSlackState Empty = new IntegrationSlackState();

    /**
     * Branches to send notifications for. Valid options are &#34;all&#34;, &#34;default&#34;, &#34;protected&#34;, and &#34;default*and*protected&#34;.
     * 
     */
    @Import(name="branchesToBeNotified")
    private @Nullable Output<String> branchesToBeNotified;

    /**
     * @return Branches to send notifications for. Valid options are &#34;all&#34;, &#34;default&#34;, &#34;protected&#34;, and &#34;default*and*protected&#34;.
     * 
     */
    public Optional<Output<String>> branchesToBeNotified() {
        return Optional.ofNullable(this.branchesToBeNotified);
    }

    /**
     * The name of the channel to receive confidential issue events notifications.
     * 
     */
    @Import(name="confidentialIssueChannel")
    private @Nullable Output<String> confidentialIssueChannel;

    /**
     * @return The name of the channel to receive confidential issue events notifications.
     * 
     */
    public Optional<Output<String>> confidentialIssueChannel() {
        return Optional.ofNullable(this.confidentialIssueChannel);
    }

    /**
     * Enable notifications for confidential issues events.
     * 
     */
    @Import(name="confidentialIssuesEvents")
    private @Nullable Output<Boolean> confidentialIssuesEvents;

    /**
     * @return Enable notifications for confidential issues events.
     * 
     */
    public Optional<Output<Boolean>> confidentialIssuesEvents() {
        return Optional.ofNullable(this.confidentialIssuesEvents);
    }

    /**
     * The name of the channel to receive confidential note events notifications.
     * 
     */
    @Import(name="confidentialNoteChannel")
    private @Nullable Output<String> confidentialNoteChannel;

    /**
     * @return The name of the channel to receive confidential note events notifications.
     * 
     */
    public Optional<Output<String>> confidentialNoteChannel() {
        return Optional.ofNullable(this.confidentialNoteChannel);
    }

    /**
     * Enable notifications for confidential note events.
     * 
     */
    @Import(name="confidentialNoteEvents")
    private @Nullable Output<Boolean> confidentialNoteEvents;

    /**
     * @return Enable notifications for confidential note events.
     * 
     */
    public Optional<Output<Boolean>> confidentialNoteEvents() {
        return Optional.ofNullable(this.confidentialNoteEvents);
    }

    /**
     * The name of the channel to receive issue events notifications.
     * 
     */
    @Import(name="issueChannel")
    private @Nullable Output<String> issueChannel;

    /**
     * @return The name of the channel to receive issue events notifications.
     * 
     */
    public Optional<Output<String>> issueChannel() {
        return Optional.ofNullable(this.issueChannel);
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
     * Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to [https://github.com/xanzy/go-gitlab/issues/1354](https://github.com/xanzy/go-gitlab/issues/1354).
     * 
     */
    @Import(name="jobEvents")
    private @Nullable Output<Boolean> jobEvents;

    /**
     * @return Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to [https://github.com/xanzy/go-gitlab/issues/1354](https://github.com/xanzy/go-gitlab/issues/1354).
     * 
     */
    public Optional<Output<Boolean>> jobEvents() {
        return Optional.ofNullable(this.jobEvents);
    }

    /**
     * The name of the channel to receive merge request events notifications.
     * 
     */
    @Import(name="mergeRequestChannel")
    private @Nullable Output<String> mergeRequestChannel;

    /**
     * @return The name of the channel to receive merge request events notifications.
     * 
     */
    public Optional<Output<String>> mergeRequestChannel() {
        return Optional.ofNullable(this.mergeRequestChannel);
    }

    /**
     * Enable notifications for merge requests events.
     * 
     */
    @Import(name="mergeRequestsEvents")
    private @Nullable Output<Boolean> mergeRequestsEvents;

    /**
     * @return Enable notifications for merge requests events.
     * 
     */
    public Optional<Output<Boolean>> mergeRequestsEvents() {
        return Optional.ofNullable(this.mergeRequestsEvents);
    }

    /**
     * The name of the channel to receive note events notifications.
     * 
     */
    @Import(name="noteChannel")
    private @Nullable Output<String> noteChannel;

    /**
     * @return The name of the channel to receive note events notifications.
     * 
     */
    public Optional<Output<String>> noteChannel() {
        return Optional.ofNullable(this.noteChannel);
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
     * Send notifications for broken pipelines.
     * 
     */
    @Import(name="notifyOnlyBrokenPipelines")
    private @Nullable Output<Boolean> notifyOnlyBrokenPipelines;

    /**
     * @return Send notifications for broken pipelines.
     * 
     */
    public Optional<Output<Boolean>> notifyOnlyBrokenPipelines() {
        return Optional.ofNullable(this.notifyOnlyBrokenPipelines);
    }

    /**
     * This parameter has been replaced with `branches_to_be_notified`.
     * 
     * @deprecated
     * use &#39;branches_to_be_notified&#39; argument instead
     * 
     */
    @Deprecated /* use 'branches_to_be_notified' argument instead */
    @Import(name="notifyOnlyDefaultBranch")
    private @Nullable Output<Boolean> notifyOnlyDefaultBranch;

    /**
     * @return This parameter has been replaced with `branches_to_be_notified`.
     * 
     * @deprecated
     * use &#39;branches_to_be_notified&#39; argument instead
     * 
     */
    @Deprecated /* use 'branches_to_be_notified' argument instead */
    public Optional<Output<Boolean>> notifyOnlyDefaultBranch() {
        return Optional.ofNullable(this.notifyOnlyDefaultBranch);
    }

    /**
     * The name of the channel to receive pipeline events notifications.
     * 
     */
    @Import(name="pipelineChannel")
    private @Nullable Output<String> pipelineChannel;

    /**
     * @return The name of the channel to receive pipeline events notifications.
     * 
     */
    public Optional<Output<String>> pipelineChannel() {
        return Optional.ofNullable(this.pipelineChannel);
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
     * The name of the channel to receive push events notifications.
     * 
     */
    @Import(name="pushChannel")
    private @Nullable Output<String> pushChannel;

    /**
     * @return The name of the channel to receive push events notifications.
     * 
     */
    public Optional<Output<String>> pushChannel() {
        return Optional.ofNullable(this.pushChannel);
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
     * The name of the channel to receive tag push events notifications.
     * 
     */
    @Import(name="tagPushChannel")
    private @Nullable Output<String> tagPushChannel;

    /**
     * @return The name of the channel to receive tag push events notifications.
     * 
     */
    public Optional<Output<String>> tagPushChannel() {
        return Optional.ofNullable(this.tagPushChannel);
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
     * Username to use.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return Username to use.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    /**
     * Webhook URL (Example, [https://hooks.slack.com/services/...)](https://hooks.slack.com/services/...)). This value cannot be imported.
     * 
     */
    @Import(name="webhook")
    private @Nullable Output<String> webhook;

    /**
     * @return Webhook URL (Example, [https://hooks.slack.com/services/...)](https://hooks.slack.com/services/...)). This value cannot be imported.
     * 
     */
    public Optional<Output<String>> webhook() {
        return Optional.ofNullable(this.webhook);
    }

    /**
     * The name of the channel to receive wiki page events notifications.
     * 
     */
    @Import(name="wikiPageChannel")
    private @Nullable Output<String> wikiPageChannel;

    /**
     * @return The name of the channel to receive wiki page events notifications.
     * 
     */
    public Optional<Output<String>> wikiPageChannel() {
        return Optional.ofNullable(this.wikiPageChannel);
    }

    /**
     * Enable notifications for wiki page events.
     * 
     */
    @Import(name="wikiPageEvents")
    private @Nullable Output<Boolean> wikiPageEvents;

    /**
     * @return Enable notifications for wiki page events.
     * 
     */
    public Optional<Output<Boolean>> wikiPageEvents() {
        return Optional.ofNullable(this.wikiPageEvents);
    }

    private IntegrationSlackState() {}

    private IntegrationSlackState(IntegrationSlackState $) {
        this.branchesToBeNotified = $.branchesToBeNotified;
        this.confidentialIssueChannel = $.confidentialIssueChannel;
        this.confidentialIssuesEvents = $.confidentialIssuesEvents;
        this.confidentialNoteChannel = $.confidentialNoteChannel;
        this.confidentialNoteEvents = $.confidentialNoteEvents;
        this.issueChannel = $.issueChannel;
        this.issuesEvents = $.issuesEvents;
        this.jobEvents = $.jobEvents;
        this.mergeRequestChannel = $.mergeRequestChannel;
        this.mergeRequestsEvents = $.mergeRequestsEvents;
        this.noteChannel = $.noteChannel;
        this.noteEvents = $.noteEvents;
        this.notifyOnlyBrokenPipelines = $.notifyOnlyBrokenPipelines;
        this.notifyOnlyDefaultBranch = $.notifyOnlyDefaultBranch;
        this.pipelineChannel = $.pipelineChannel;
        this.pipelineEvents = $.pipelineEvents;
        this.project = $.project;
        this.pushChannel = $.pushChannel;
        this.pushEvents = $.pushEvents;
        this.tagPushChannel = $.tagPushChannel;
        this.tagPushEvents = $.tagPushEvents;
        this.username = $.username;
        this.webhook = $.webhook;
        this.wikiPageChannel = $.wikiPageChannel;
        this.wikiPageEvents = $.wikiPageEvents;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationSlackState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationSlackState $;

        public Builder() {
            $ = new IntegrationSlackState();
        }

        public Builder(IntegrationSlackState defaults) {
            $ = new IntegrationSlackState(Objects.requireNonNull(defaults));
        }

        /**
         * @param branchesToBeNotified Branches to send notifications for. Valid options are &#34;all&#34;, &#34;default&#34;, &#34;protected&#34;, and &#34;default*and*protected&#34;.
         * 
         * @return builder
         * 
         */
        public Builder branchesToBeNotified(@Nullable Output<String> branchesToBeNotified) {
            $.branchesToBeNotified = branchesToBeNotified;
            return this;
        }

        /**
         * @param branchesToBeNotified Branches to send notifications for. Valid options are &#34;all&#34;, &#34;default&#34;, &#34;protected&#34;, and &#34;default*and*protected&#34;.
         * 
         * @return builder
         * 
         */
        public Builder branchesToBeNotified(String branchesToBeNotified) {
            return branchesToBeNotified(Output.of(branchesToBeNotified));
        }

        /**
         * @param confidentialIssueChannel The name of the channel to receive confidential issue events notifications.
         * 
         * @return builder
         * 
         */
        public Builder confidentialIssueChannel(@Nullable Output<String> confidentialIssueChannel) {
            $.confidentialIssueChannel = confidentialIssueChannel;
            return this;
        }

        /**
         * @param confidentialIssueChannel The name of the channel to receive confidential issue events notifications.
         * 
         * @return builder
         * 
         */
        public Builder confidentialIssueChannel(String confidentialIssueChannel) {
            return confidentialIssueChannel(Output.of(confidentialIssueChannel));
        }

        /**
         * @param confidentialIssuesEvents Enable notifications for confidential issues events.
         * 
         * @return builder
         * 
         */
        public Builder confidentialIssuesEvents(@Nullable Output<Boolean> confidentialIssuesEvents) {
            $.confidentialIssuesEvents = confidentialIssuesEvents;
            return this;
        }

        /**
         * @param confidentialIssuesEvents Enable notifications for confidential issues events.
         * 
         * @return builder
         * 
         */
        public Builder confidentialIssuesEvents(Boolean confidentialIssuesEvents) {
            return confidentialIssuesEvents(Output.of(confidentialIssuesEvents));
        }

        /**
         * @param confidentialNoteChannel The name of the channel to receive confidential note events notifications.
         * 
         * @return builder
         * 
         */
        public Builder confidentialNoteChannel(@Nullable Output<String> confidentialNoteChannel) {
            $.confidentialNoteChannel = confidentialNoteChannel;
            return this;
        }

        /**
         * @param confidentialNoteChannel The name of the channel to receive confidential note events notifications.
         * 
         * @return builder
         * 
         */
        public Builder confidentialNoteChannel(String confidentialNoteChannel) {
            return confidentialNoteChannel(Output.of(confidentialNoteChannel));
        }

        /**
         * @param confidentialNoteEvents Enable notifications for confidential note events.
         * 
         * @return builder
         * 
         */
        public Builder confidentialNoteEvents(@Nullable Output<Boolean> confidentialNoteEvents) {
            $.confidentialNoteEvents = confidentialNoteEvents;
            return this;
        }

        /**
         * @param confidentialNoteEvents Enable notifications for confidential note events.
         * 
         * @return builder
         * 
         */
        public Builder confidentialNoteEvents(Boolean confidentialNoteEvents) {
            return confidentialNoteEvents(Output.of(confidentialNoteEvents));
        }

        /**
         * @param issueChannel The name of the channel to receive issue events notifications.
         * 
         * @return builder
         * 
         */
        public Builder issueChannel(@Nullable Output<String> issueChannel) {
            $.issueChannel = issueChannel;
            return this;
        }

        /**
         * @param issueChannel The name of the channel to receive issue events notifications.
         * 
         * @return builder
         * 
         */
        public Builder issueChannel(String issueChannel) {
            return issueChannel(Output.of(issueChannel));
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
         * @param jobEvents Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to [https://github.com/xanzy/go-gitlab/issues/1354](https://github.com/xanzy/go-gitlab/issues/1354).
         * 
         * @return builder
         * 
         */
        public Builder jobEvents(@Nullable Output<Boolean> jobEvents) {
            $.jobEvents = jobEvents;
            return this;
        }

        /**
         * @param jobEvents Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to [https://github.com/xanzy/go-gitlab/issues/1354](https://github.com/xanzy/go-gitlab/issues/1354).
         * 
         * @return builder
         * 
         */
        public Builder jobEvents(Boolean jobEvents) {
            return jobEvents(Output.of(jobEvents));
        }

        /**
         * @param mergeRequestChannel The name of the channel to receive merge request events notifications.
         * 
         * @return builder
         * 
         */
        public Builder mergeRequestChannel(@Nullable Output<String> mergeRequestChannel) {
            $.mergeRequestChannel = mergeRequestChannel;
            return this;
        }

        /**
         * @param mergeRequestChannel The name of the channel to receive merge request events notifications.
         * 
         * @return builder
         * 
         */
        public Builder mergeRequestChannel(String mergeRequestChannel) {
            return mergeRequestChannel(Output.of(mergeRequestChannel));
        }

        /**
         * @param mergeRequestsEvents Enable notifications for merge requests events.
         * 
         * @return builder
         * 
         */
        public Builder mergeRequestsEvents(@Nullable Output<Boolean> mergeRequestsEvents) {
            $.mergeRequestsEvents = mergeRequestsEvents;
            return this;
        }

        /**
         * @param mergeRequestsEvents Enable notifications for merge requests events.
         * 
         * @return builder
         * 
         */
        public Builder mergeRequestsEvents(Boolean mergeRequestsEvents) {
            return mergeRequestsEvents(Output.of(mergeRequestsEvents));
        }

        /**
         * @param noteChannel The name of the channel to receive note events notifications.
         * 
         * @return builder
         * 
         */
        public Builder noteChannel(@Nullable Output<String> noteChannel) {
            $.noteChannel = noteChannel;
            return this;
        }

        /**
         * @param noteChannel The name of the channel to receive note events notifications.
         * 
         * @return builder
         * 
         */
        public Builder noteChannel(String noteChannel) {
            return noteChannel(Output.of(noteChannel));
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
         * @param notifyOnlyBrokenPipelines Send notifications for broken pipelines.
         * 
         * @return builder
         * 
         */
        public Builder notifyOnlyBrokenPipelines(@Nullable Output<Boolean> notifyOnlyBrokenPipelines) {
            $.notifyOnlyBrokenPipelines = notifyOnlyBrokenPipelines;
            return this;
        }

        /**
         * @param notifyOnlyBrokenPipelines Send notifications for broken pipelines.
         * 
         * @return builder
         * 
         */
        public Builder notifyOnlyBrokenPipelines(Boolean notifyOnlyBrokenPipelines) {
            return notifyOnlyBrokenPipelines(Output.of(notifyOnlyBrokenPipelines));
        }

        /**
         * @param notifyOnlyDefaultBranch This parameter has been replaced with `branches_to_be_notified`.
         * 
         * @return builder
         * 
         * @deprecated
         * use &#39;branches_to_be_notified&#39; argument instead
         * 
         */
        @Deprecated /* use 'branches_to_be_notified' argument instead */
        public Builder notifyOnlyDefaultBranch(@Nullable Output<Boolean> notifyOnlyDefaultBranch) {
            $.notifyOnlyDefaultBranch = notifyOnlyDefaultBranch;
            return this;
        }

        /**
         * @param notifyOnlyDefaultBranch This parameter has been replaced with `branches_to_be_notified`.
         * 
         * @return builder
         * 
         * @deprecated
         * use &#39;branches_to_be_notified&#39; argument instead
         * 
         */
        @Deprecated /* use 'branches_to_be_notified' argument instead */
        public Builder notifyOnlyDefaultBranch(Boolean notifyOnlyDefaultBranch) {
            return notifyOnlyDefaultBranch(Output.of(notifyOnlyDefaultBranch));
        }

        /**
         * @param pipelineChannel The name of the channel to receive pipeline events notifications.
         * 
         * @return builder
         * 
         */
        public Builder pipelineChannel(@Nullable Output<String> pipelineChannel) {
            $.pipelineChannel = pipelineChannel;
            return this;
        }

        /**
         * @param pipelineChannel The name of the channel to receive pipeline events notifications.
         * 
         * @return builder
         * 
         */
        public Builder pipelineChannel(String pipelineChannel) {
            return pipelineChannel(Output.of(pipelineChannel));
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
         * @param pushChannel The name of the channel to receive push events notifications.
         * 
         * @return builder
         * 
         */
        public Builder pushChannel(@Nullable Output<String> pushChannel) {
            $.pushChannel = pushChannel;
            return this;
        }

        /**
         * @param pushChannel The name of the channel to receive push events notifications.
         * 
         * @return builder
         * 
         */
        public Builder pushChannel(String pushChannel) {
            return pushChannel(Output.of(pushChannel));
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
         * @param tagPushChannel The name of the channel to receive tag push events notifications.
         * 
         * @return builder
         * 
         */
        public Builder tagPushChannel(@Nullable Output<String> tagPushChannel) {
            $.tagPushChannel = tagPushChannel;
            return this;
        }

        /**
         * @param tagPushChannel The name of the channel to receive tag push events notifications.
         * 
         * @return builder
         * 
         */
        public Builder tagPushChannel(String tagPushChannel) {
            return tagPushChannel(Output.of(tagPushChannel));
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
         * @param username Username to use.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Username to use.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        /**
         * @param webhook Webhook URL (Example, [https://hooks.slack.com/services/...)](https://hooks.slack.com/services/...)). This value cannot be imported.
         * 
         * @return builder
         * 
         */
        public Builder webhook(@Nullable Output<String> webhook) {
            $.webhook = webhook;
            return this;
        }

        /**
         * @param webhook Webhook URL (Example, [https://hooks.slack.com/services/...)](https://hooks.slack.com/services/...)). This value cannot be imported.
         * 
         * @return builder
         * 
         */
        public Builder webhook(String webhook) {
            return webhook(Output.of(webhook));
        }

        /**
         * @param wikiPageChannel The name of the channel to receive wiki page events notifications.
         * 
         * @return builder
         * 
         */
        public Builder wikiPageChannel(@Nullable Output<String> wikiPageChannel) {
            $.wikiPageChannel = wikiPageChannel;
            return this;
        }

        /**
         * @param wikiPageChannel The name of the channel to receive wiki page events notifications.
         * 
         * @return builder
         * 
         */
        public Builder wikiPageChannel(String wikiPageChannel) {
            return wikiPageChannel(Output.of(wikiPageChannel));
        }

        /**
         * @param wikiPageEvents Enable notifications for wiki page events.
         * 
         * @return builder
         * 
         */
        public Builder wikiPageEvents(@Nullable Output<Boolean> wikiPageEvents) {
            $.wikiPageEvents = wikiPageEvents;
            return this;
        }

        /**
         * @param wikiPageEvents Enable notifications for wiki page events.
         * 
         * @return builder
         * 
         */
        public Builder wikiPageEvents(Boolean wikiPageEvents) {
            return wikiPageEvents(Output.of(wikiPageEvents));
        }

        public IntegrationSlackState build() {
            return $;
        }
    }

}
