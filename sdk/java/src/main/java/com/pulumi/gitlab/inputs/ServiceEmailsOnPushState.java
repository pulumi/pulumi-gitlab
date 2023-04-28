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


public final class ServiceEmailsOnPushState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEmailsOnPushState Empty = new ServiceEmailsOnPushState();

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
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
     * 
     */
    @Import(name="branchesToBeNotified")
    private @Nullable Output<String> branchesToBeNotified;

    /**
     * @return Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
     * 
     */
    public Optional<Output<String>> branchesToBeNotified() {
        return Optional.ofNullable(this.branchesToBeNotified);
    }

    /**
     * The ISO8601 date/time that this integration was activated at in UTC.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return The ISO8601 date/time that this integration was activated at in UTC.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Disable code diffs.
     * 
     */
    @Import(name="disableDiffs")
    private @Nullable Output<Boolean> disableDiffs;

    /**
     * @return Disable code diffs.
     * 
     */
    public Optional<Output<Boolean>> disableDiffs() {
        return Optional.ofNullable(this.disableDiffs);
    }

    /**
     * ID or full-path of the project you want to activate integration on.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return ID or full-path of the project you want to activate integration on.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
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
     * Emails separated by whitespace.
     * 
     */
    @Import(name="recipients")
    private @Nullable Output<String> recipients;

    /**
     * @return Emails separated by whitespace.
     * 
     */
    public Optional<Output<String>> recipients() {
        return Optional.ofNullable(this.recipients);
    }

    /**
     * Send from committer.
     * 
     */
    @Import(name="sendFromCommitterEmail")
    private @Nullable Output<Boolean> sendFromCommitterEmail;

    /**
     * @return Send from committer.
     * 
     */
    public Optional<Output<Boolean>> sendFromCommitterEmail() {
        return Optional.ofNullable(this.sendFromCommitterEmail);
    }

    /**
     * The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     * 
     */
    @Import(name="slug")
    private @Nullable Output<String> slug;

    /**
     * @return The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     * 
     */
    public Optional<Output<String>> slug() {
        return Optional.ofNullable(this.slug);
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
     * Title of the integration.
     * 
     */
    @Import(name="title")
    private @Nullable Output<String> title;

    /**
     * @return Title of the integration.
     * 
     */
    public Optional<Output<String>> title() {
        return Optional.ofNullable(this.title);
    }

    /**
     * The ISO8601 date/time that this integration was last updated at in UTC.
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return The ISO8601 date/time that this integration was last updated at in UTC.
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    private ServiceEmailsOnPushState() {}

    private ServiceEmailsOnPushState(ServiceEmailsOnPushState $) {
        this.active = $.active;
        this.branchesToBeNotified = $.branchesToBeNotified;
        this.createdAt = $.createdAt;
        this.disableDiffs = $.disableDiffs;
        this.project = $.project;
        this.pushEvents = $.pushEvents;
        this.recipients = $.recipients;
        this.sendFromCommitterEmail = $.sendFromCommitterEmail;
        this.slug = $.slug;
        this.tagPushEvents = $.tagPushEvents;
        this.title = $.title;
        this.updatedAt = $.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEmailsOnPushState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEmailsOnPushState $;

        public Builder() {
            $ = new ServiceEmailsOnPushState();
        }

        public Builder(ServiceEmailsOnPushState defaults) {
            $ = new ServiceEmailsOnPushState(Objects.requireNonNull(defaults));
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
         * @param branchesToBeNotified Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
         * 
         * @return builder
         * 
         */
        public Builder branchesToBeNotified(@Nullable Output<String> branchesToBeNotified) {
            $.branchesToBeNotified = branchesToBeNotified;
            return this;
        }

        /**
         * @param branchesToBeNotified Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
         * 
         * @return builder
         * 
         */
        public Builder branchesToBeNotified(String branchesToBeNotified) {
            return branchesToBeNotified(Output.of(branchesToBeNotified));
        }

        /**
         * @param createdAt The ISO8601 date/time that this integration was activated at in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt The ISO8601 date/time that this integration was activated at in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param disableDiffs Disable code diffs.
         * 
         * @return builder
         * 
         */
        public Builder disableDiffs(@Nullable Output<Boolean> disableDiffs) {
            $.disableDiffs = disableDiffs;
            return this;
        }

        /**
         * @param disableDiffs Disable code diffs.
         * 
         * @return builder
         * 
         */
        public Builder disableDiffs(Boolean disableDiffs) {
            return disableDiffs(Output.of(disableDiffs));
        }

        /**
         * @param project ID or full-path of the project you want to activate integration on.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project ID or full-path of the project you want to activate integration on.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
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
         * @param recipients Emails separated by whitespace.
         * 
         * @return builder
         * 
         */
        public Builder recipients(@Nullable Output<String> recipients) {
            $.recipients = recipients;
            return this;
        }

        /**
         * @param recipients Emails separated by whitespace.
         * 
         * @return builder
         * 
         */
        public Builder recipients(String recipients) {
            return recipients(Output.of(recipients));
        }

        /**
         * @param sendFromCommitterEmail Send from committer.
         * 
         * @return builder
         * 
         */
        public Builder sendFromCommitterEmail(@Nullable Output<Boolean> sendFromCommitterEmail) {
            $.sendFromCommitterEmail = sendFromCommitterEmail;
            return this;
        }

        /**
         * @param sendFromCommitterEmail Send from committer.
         * 
         * @return builder
         * 
         */
        public Builder sendFromCommitterEmail(Boolean sendFromCommitterEmail) {
            return sendFromCommitterEmail(Output.of(sendFromCommitterEmail));
        }

        /**
         * @param slug The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
         * 
         * @return builder
         * 
         */
        public Builder slug(@Nullable Output<String> slug) {
            $.slug = slug;
            return this;
        }

        /**
         * @param slug The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
         * 
         * @return builder
         * 
         */
        public Builder slug(String slug) {
            return slug(Output.of(slug));
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
         * @param title Title of the integration.
         * 
         * @return builder
         * 
         */
        public Builder title(@Nullable Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title Title of the integration.
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        /**
         * @param updatedAt The ISO8601 date/time that this integration was last updated at in UTC.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt The ISO8601 date/time that this integration was last updated at in UTC.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        public ServiceEmailsOnPushState build() {
            return $;
        }
    }

}