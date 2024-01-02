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


public final class IntegrationEmailsOnPushArgs extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationEmailsOnPushArgs Empty = new IntegrationEmailsOnPushArgs();

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
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return ID or full-path of the project you want to activate integration on.
     * 
     */
    public Output<String> project() {
        return this.project;
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
    @Import(name="recipients", required=true)
    private Output<String> recipients;

    /**
     * @return Emails separated by whitespace.
     * 
     */
    public Output<String> recipients() {
        return this.recipients;
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

    private IntegrationEmailsOnPushArgs() {}

    private IntegrationEmailsOnPushArgs(IntegrationEmailsOnPushArgs $) {
        this.branchesToBeNotified = $.branchesToBeNotified;
        this.disableDiffs = $.disableDiffs;
        this.project = $.project;
        this.pushEvents = $.pushEvents;
        this.recipients = $.recipients;
        this.sendFromCommitterEmail = $.sendFromCommitterEmail;
        this.tagPushEvents = $.tagPushEvents;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationEmailsOnPushArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationEmailsOnPushArgs $;

        public Builder() {
            $ = new IntegrationEmailsOnPushArgs();
        }

        public Builder(IntegrationEmailsOnPushArgs defaults) {
            $ = new IntegrationEmailsOnPushArgs(Objects.requireNonNull(defaults));
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
        public Builder project(Output<String> project) {
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
        public Builder recipients(Output<String> recipients) {
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

        public IntegrationEmailsOnPushArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("IntegrationEmailsOnPushArgs", "project");
            }
            if ($.recipients == null) {
                throw new MissingRequiredPropertyException("IntegrationEmailsOnPushArgs", "recipients");
            }
            return $;
        }
    }

}
