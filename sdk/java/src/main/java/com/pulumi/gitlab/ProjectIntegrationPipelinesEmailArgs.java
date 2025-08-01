// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectIntegrationPipelinesEmailArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectIntegrationPipelinesEmailArgs Empty = new ProjectIntegrationPipelinesEmailArgs();

    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
     * 
     */
    @Import(name="branchesToBeNotified")
    private @Nullable Output<String> branchesToBeNotified;

    /**
     * @return Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
     * 
     */
    public Optional<Output<String>> branchesToBeNotified() {
        return Optional.ofNullable(this.branchesToBeNotified);
    }

    /**
     * Notify only broken pipelines. Default is true.
     * 
     */
    @Import(name="notifyOnlyBrokenPipelines")
    private @Nullable Output<Boolean> notifyOnlyBrokenPipelines;

    /**
     * @return Notify only broken pipelines. Default is true.
     * 
     */
    public Optional<Output<Boolean>> notifyOnlyBrokenPipelines() {
        return Optional.ofNullable(this.notifyOnlyBrokenPipelines);
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
     * ) email addresses where notifications are sent.
     * 
     */
    @Import(name="recipients", required=true)
    private Output<List<String>> recipients;

    /**
     * @return ) email addresses where notifications are sent.
     * 
     */
    public Output<List<String>> recipients() {
        return this.recipients;
    }

    private ProjectIntegrationPipelinesEmailArgs() {}

    private ProjectIntegrationPipelinesEmailArgs(ProjectIntegrationPipelinesEmailArgs $) {
        this.branchesToBeNotified = $.branchesToBeNotified;
        this.notifyOnlyBrokenPipelines = $.notifyOnlyBrokenPipelines;
        this.project = $.project;
        this.recipients = $.recipients;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectIntegrationPipelinesEmailArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectIntegrationPipelinesEmailArgs $;

        public Builder() {
            $ = new ProjectIntegrationPipelinesEmailArgs();
        }

        public Builder(ProjectIntegrationPipelinesEmailArgs defaults) {
            $ = new ProjectIntegrationPipelinesEmailArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param branchesToBeNotified Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
         * 
         * @return builder
         * 
         */
        public Builder branchesToBeNotified(@Nullable Output<String> branchesToBeNotified) {
            $.branchesToBeNotified = branchesToBeNotified;
            return this;
        }

        /**
         * @param branchesToBeNotified Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
         * 
         * @return builder
         * 
         */
        public Builder branchesToBeNotified(String branchesToBeNotified) {
            return branchesToBeNotified(Output.of(branchesToBeNotified));
        }

        /**
         * @param notifyOnlyBrokenPipelines Notify only broken pipelines. Default is true.
         * 
         * @return builder
         * 
         */
        public Builder notifyOnlyBrokenPipelines(@Nullable Output<Boolean> notifyOnlyBrokenPipelines) {
            $.notifyOnlyBrokenPipelines = notifyOnlyBrokenPipelines;
            return this;
        }

        /**
         * @param notifyOnlyBrokenPipelines Notify only broken pipelines. Default is true.
         * 
         * @return builder
         * 
         */
        public Builder notifyOnlyBrokenPipelines(Boolean notifyOnlyBrokenPipelines) {
            return notifyOnlyBrokenPipelines(Output.of(notifyOnlyBrokenPipelines));
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
         * @param recipients ) email addresses where notifications are sent.
         * 
         * @return builder
         * 
         */
        public Builder recipients(Output<List<String>> recipients) {
            $.recipients = recipients;
            return this;
        }

        /**
         * @param recipients ) email addresses where notifications are sent.
         * 
         * @return builder
         * 
         */
        public Builder recipients(List<String> recipients) {
            return recipients(Output.of(recipients));
        }

        /**
         * @param recipients ) email addresses where notifications are sent.
         * 
         * @return builder
         * 
         */
        public Builder recipients(String... recipients) {
            return recipients(List.of(recipients));
        }

        public ProjectIntegrationPipelinesEmailArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("ProjectIntegrationPipelinesEmailArgs", "project");
            }
            if ($.recipients == null) {
                throw new MissingRequiredPropertyException("ProjectIntegrationPipelinesEmailArgs", "recipients");
            }
            return $;
        }
    }

}
