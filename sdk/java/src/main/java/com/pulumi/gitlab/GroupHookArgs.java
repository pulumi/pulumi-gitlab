// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.inputs.GroupHookCustomHeaderArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupHookArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupHookArgs Empty = new GroupHookArgs();

    /**
     * Filter push events by branch. Valid values are: `wildcard`, `regex`, `all_branches`.
     * 
     */
    @Import(name="branchFilterStrategy")
    private @Nullable Output<String> branchFilterStrategy;

    /**
     * @return Filter push events by branch. Valid values are: `wildcard`, `regex`, `all_branches`.
     * 
     */
    public Optional<Output<String>> branchFilterStrategy() {
        return Optional.ofNullable(this.branchFilterStrategy);
    }

    /**
     * Invoke the hook for confidential issues events.
     * 
     */
    @Import(name="confidentialIssuesEvents")
    private @Nullable Output<Boolean> confidentialIssuesEvents;

    /**
     * @return Invoke the hook for confidential issues events.
     * 
     */
    public Optional<Output<Boolean>> confidentialIssuesEvents() {
        return Optional.ofNullable(this.confidentialIssuesEvents);
    }

    /**
     * Invoke the hook for confidential note events.
     * 
     */
    @Import(name="confidentialNoteEvents")
    private @Nullable Output<Boolean> confidentialNoteEvents;

    /**
     * @return Invoke the hook for confidential note events.
     * 
     */
    public Optional<Output<Boolean>> confidentialNoteEvents() {
        return Optional.ofNullable(this.confidentialNoteEvents);
    }

    /**
     * Custom headers for the project webhook.
     * 
     */
    @Import(name="customHeaders")
    private @Nullable Output<List<GroupHookCustomHeaderArgs>> customHeaders;

    /**
     * @return Custom headers for the project webhook.
     * 
     */
    public Optional<Output<List<GroupHookCustomHeaderArgs>>> customHeaders() {
        return Optional.ofNullable(this.customHeaders);
    }

    /**
     * Custom webhook template.
     * 
     */
    @Import(name="customWebhookTemplate")
    private @Nullable Output<String> customWebhookTemplate;

    /**
     * @return Custom webhook template.
     * 
     */
    public Optional<Output<String>> customWebhookTemplate() {
        return Optional.ofNullable(this.customWebhookTemplate);
    }

    /**
     * Invoke the hook for deployment events.
     * 
     */
    @Import(name="deploymentEvents")
    private @Nullable Output<Boolean> deploymentEvents;

    /**
     * @return Invoke the hook for deployment events.
     * 
     */
    public Optional<Output<Boolean>> deploymentEvents() {
        return Optional.ofNullable(this.deploymentEvents);
    }

    /**
     * Description of the group webhook.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the group webhook.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Enable SSL verification when invoking the hook.
     * 
     */
    @Import(name="enableSslVerification")
    private @Nullable Output<Boolean> enableSslVerification;

    /**
     * @return Enable SSL verification when invoking the hook.
     * 
     */
    public Optional<Output<Boolean>> enableSslVerification() {
        return Optional.ofNullable(this.enableSslVerification);
    }

    /**
     * Invoke the hook for feature flag events.
     * 
     */
    @Import(name="featureFlagEvents")
    private @Nullable Output<Boolean> featureFlagEvents;

    /**
     * @return Invoke the hook for feature flag events.
     * 
     */
    public Optional<Output<Boolean>> featureFlagEvents() {
        return Optional.ofNullable(this.featureFlagEvents);
    }

    /**
     * The full path or id of the group to add the hook to.
     * 
     */
    @Import(name="group", required=true)
    private Output<String> group;

    /**
     * @return The full path or id of the group to add the hook to.
     * 
     */
    public Output<String> group() {
        return this.group;
    }

    /**
     * Invoke the hook for issues events.
     * 
     */
    @Import(name="issuesEvents")
    private @Nullable Output<Boolean> issuesEvents;

    /**
     * @return Invoke the hook for issues events.
     * 
     */
    public Optional<Output<Boolean>> issuesEvents() {
        return Optional.ofNullable(this.issuesEvents);
    }

    /**
     * Invoke the hook for job events.
     * 
     */
    @Import(name="jobEvents")
    private @Nullable Output<Boolean> jobEvents;

    /**
     * @return Invoke the hook for job events.
     * 
     */
    public Optional<Output<Boolean>> jobEvents() {
        return Optional.ofNullable(this.jobEvents);
    }

    /**
     * Invoke the hook for merge requests events.
     * 
     */
    @Import(name="mergeRequestsEvents")
    private @Nullable Output<Boolean> mergeRequestsEvents;

    /**
     * @return Invoke the hook for merge requests events.
     * 
     */
    public Optional<Output<Boolean>> mergeRequestsEvents() {
        return Optional.ofNullable(this.mergeRequestsEvents);
    }

    /**
     * Name of the group webhook.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the group webhook.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Invoke the hook for note events.
     * 
     */
    @Import(name="noteEvents")
    private @Nullable Output<Boolean> noteEvents;

    /**
     * @return Invoke the hook for note events.
     * 
     */
    public Optional<Output<Boolean>> noteEvents() {
        return Optional.ofNullable(this.noteEvents);
    }

    /**
     * Invoke the hook for pipeline events.
     * 
     */
    @Import(name="pipelineEvents")
    private @Nullable Output<Boolean> pipelineEvents;

    /**
     * @return Invoke the hook for pipeline events.
     * 
     */
    public Optional<Output<Boolean>> pipelineEvents() {
        return Optional.ofNullable(this.pipelineEvents);
    }

    /**
     * Invoke the hook for push events.
     * 
     */
    @Import(name="pushEvents")
    private @Nullable Output<Boolean> pushEvents;

    /**
     * @return Invoke the hook for push events.
     * 
     */
    public Optional<Output<Boolean>> pushEvents() {
        return Optional.ofNullable(this.pushEvents);
    }

    /**
     * Invoke the hook for push events on matching branches only.
     * 
     */
    @Import(name="pushEventsBranchFilter")
    private @Nullable Output<String> pushEventsBranchFilter;

    /**
     * @return Invoke the hook for push events on matching branches only.
     * 
     */
    public Optional<Output<String>> pushEventsBranchFilter() {
        return Optional.ofNullable(this.pushEventsBranchFilter);
    }

    /**
     * Invoke the hook for release events.
     * 
     */
    @Import(name="releasesEvents")
    private @Nullable Output<Boolean> releasesEvents;

    /**
     * @return Invoke the hook for release events.
     * 
     */
    public Optional<Output<Boolean>> releasesEvents() {
        return Optional.ofNullable(this.releasesEvents);
    }

    /**
     * Invoke the hook for subgroup events.
     * 
     */
    @Import(name="subgroupEvents")
    private @Nullable Output<Boolean> subgroupEvents;

    /**
     * @return Invoke the hook for subgroup events.
     * 
     */
    public Optional<Output<Boolean>> subgroupEvents() {
        return Optional.ofNullable(this.subgroupEvents);
    }

    /**
     * Invoke the hook for tag push events.
     * 
     */
    @Import(name="tagPushEvents")
    private @Nullable Output<Boolean> tagPushEvents;

    /**
     * @return Invoke the hook for tag push events.
     * 
     */
    public Optional<Output<Boolean>> tagPushEvents() {
        return Optional.ofNullable(this.tagPushEvents);
    }

    /**
     * A token to present when invoking the hook. The token is not available for imported resources.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return A token to present when invoking the hook. The token is not available for imported resources.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * The url of the hook to invoke. Forces re-creation to preserve `token`.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The url of the hook to invoke. Forces re-creation to preserve `token`.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * Invoke the hook for wiki page events.
     * 
     */
    @Import(name="wikiPageEvents")
    private @Nullable Output<Boolean> wikiPageEvents;

    /**
     * @return Invoke the hook for wiki page events.
     * 
     */
    public Optional<Output<Boolean>> wikiPageEvents() {
        return Optional.ofNullable(this.wikiPageEvents);
    }

    private GroupHookArgs() {}

    private GroupHookArgs(GroupHookArgs $) {
        this.branchFilterStrategy = $.branchFilterStrategy;
        this.confidentialIssuesEvents = $.confidentialIssuesEvents;
        this.confidentialNoteEvents = $.confidentialNoteEvents;
        this.customHeaders = $.customHeaders;
        this.customWebhookTemplate = $.customWebhookTemplate;
        this.deploymentEvents = $.deploymentEvents;
        this.description = $.description;
        this.enableSslVerification = $.enableSslVerification;
        this.featureFlagEvents = $.featureFlagEvents;
        this.group = $.group;
        this.issuesEvents = $.issuesEvents;
        this.jobEvents = $.jobEvents;
        this.mergeRequestsEvents = $.mergeRequestsEvents;
        this.name = $.name;
        this.noteEvents = $.noteEvents;
        this.pipelineEvents = $.pipelineEvents;
        this.pushEvents = $.pushEvents;
        this.pushEventsBranchFilter = $.pushEventsBranchFilter;
        this.releasesEvents = $.releasesEvents;
        this.subgroupEvents = $.subgroupEvents;
        this.tagPushEvents = $.tagPushEvents;
        this.token = $.token;
        this.url = $.url;
        this.wikiPageEvents = $.wikiPageEvents;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupHookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupHookArgs $;

        public Builder() {
            $ = new GroupHookArgs();
        }

        public Builder(GroupHookArgs defaults) {
            $ = new GroupHookArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param branchFilterStrategy Filter push events by branch. Valid values are: `wildcard`, `regex`, `all_branches`.
         * 
         * @return builder
         * 
         */
        public Builder branchFilterStrategy(@Nullable Output<String> branchFilterStrategy) {
            $.branchFilterStrategy = branchFilterStrategy;
            return this;
        }

        /**
         * @param branchFilterStrategy Filter push events by branch. Valid values are: `wildcard`, `regex`, `all_branches`.
         * 
         * @return builder
         * 
         */
        public Builder branchFilterStrategy(String branchFilterStrategy) {
            return branchFilterStrategy(Output.of(branchFilterStrategy));
        }

        /**
         * @param confidentialIssuesEvents Invoke the hook for confidential issues events.
         * 
         * @return builder
         * 
         */
        public Builder confidentialIssuesEvents(@Nullable Output<Boolean> confidentialIssuesEvents) {
            $.confidentialIssuesEvents = confidentialIssuesEvents;
            return this;
        }

        /**
         * @param confidentialIssuesEvents Invoke the hook for confidential issues events.
         * 
         * @return builder
         * 
         */
        public Builder confidentialIssuesEvents(Boolean confidentialIssuesEvents) {
            return confidentialIssuesEvents(Output.of(confidentialIssuesEvents));
        }

        /**
         * @param confidentialNoteEvents Invoke the hook for confidential note events.
         * 
         * @return builder
         * 
         */
        public Builder confidentialNoteEvents(@Nullable Output<Boolean> confidentialNoteEvents) {
            $.confidentialNoteEvents = confidentialNoteEvents;
            return this;
        }

        /**
         * @param confidentialNoteEvents Invoke the hook for confidential note events.
         * 
         * @return builder
         * 
         */
        public Builder confidentialNoteEvents(Boolean confidentialNoteEvents) {
            return confidentialNoteEvents(Output.of(confidentialNoteEvents));
        }

        /**
         * @param customHeaders Custom headers for the project webhook.
         * 
         * @return builder
         * 
         */
        public Builder customHeaders(@Nullable Output<List<GroupHookCustomHeaderArgs>> customHeaders) {
            $.customHeaders = customHeaders;
            return this;
        }

        /**
         * @param customHeaders Custom headers for the project webhook.
         * 
         * @return builder
         * 
         */
        public Builder customHeaders(List<GroupHookCustomHeaderArgs> customHeaders) {
            return customHeaders(Output.of(customHeaders));
        }

        /**
         * @param customHeaders Custom headers for the project webhook.
         * 
         * @return builder
         * 
         */
        public Builder customHeaders(GroupHookCustomHeaderArgs... customHeaders) {
            return customHeaders(List.of(customHeaders));
        }

        /**
         * @param customWebhookTemplate Custom webhook template.
         * 
         * @return builder
         * 
         */
        public Builder customWebhookTemplate(@Nullable Output<String> customWebhookTemplate) {
            $.customWebhookTemplate = customWebhookTemplate;
            return this;
        }

        /**
         * @param customWebhookTemplate Custom webhook template.
         * 
         * @return builder
         * 
         */
        public Builder customWebhookTemplate(String customWebhookTemplate) {
            return customWebhookTemplate(Output.of(customWebhookTemplate));
        }

        /**
         * @param deploymentEvents Invoke the hook for deployment events.
         * 
         * @return builder
         * 
         */
        public Builder deploymentEvents(@Nullable Output<Boolean> deploymentEvents) {
            $.deploymentEvents = deploymentEvents;
            return this;
        }

        /**
         * @param deploymentEvents Invoke the hook for deployment events.
         * 
         * @return builder
         * 
         */
        public Builder deploymentEvents(Boolean deploymentEvents) {
            return deploymentEvents(Output.of(deploymentEvents));
        }

        /**
         * @param description Description of the group webhook.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the group webhook.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param enableSslVerification Enable SSL verification when invoking the hook.
         * 
         * @return builder
         * 
         */
        public Builder enableSslVerification(@Nullable Output<Boolean> enableSslVerification) {
            $.enableSslVerification = enableSslVerification;
            return this;
        }

        /**
         * @param enableSslVerification Enable SSL verification when invoking the hook.
         * 
         * @return builder
         * 
         */
        public Builder enableSslVerification(Boolean enableSslVerification) {
            return enableSslVerification(Output.of(enableSslVerification));
        }

        /**
         * @param featureFlagEvents Invoke the hook for feature flag events.
         * 
         * @return builder
         * 
         */
        public Builder featureFlagEvents(@Nullable Output<Boolean> featureFlagEvents) {
            $.featureFlagEvents = featureFlagEvents;
            return this;
        }

        /**
         * @param featureFlagEvents Invoke the hook for feature flag events.
         * 
         * @return builder
         * 
         */
        public Builder featureFlagEvents(Boolean featureFlagEvents) {
            return featureFlagEvents(Output.of(featureFlagEvents));
        }

        /**
         * @param group The full path or id of the group to add the hook to.
         * 
         * @return builder
         * 
         */
        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The full path or id of the group to add the hook to.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param issuesEvents Invoke the hook for issues events.
         * 
         * @return builder
         * 
         */
        public Builder issuesEvents(@Nullable Output<Boolean> issuesEvents) {
            $.issuesEvents = issuesEvents;
            return this;
        }

        /**
         * @param issuesEvents Invoke the hook for issues events.
         * 
         * @return builder
         * 
         */
        public Builder issuesEvents(Boolean issuesEvents) {
            return issuesEvents(Output.of(issuesEvents));
        }

        /**
         * @param jobEvents Invoke the hook for job events.
         * 
         * @return builder
         * 
         */
        public Builder jobEvents(@Nullable Output<Boolean> jobEvents) {
            $.jobEvents = jobEvents;
            return this;
        }

        /**
         * @param jobEvents Invoke the hook for job events.
         * 
         * @return builder
         * 
         */
        public Builder jobEvents(Boolean jobEvents) {
            return jobEvents(Output.of(jobEvents));
        }

        /**
         * @param mergeRequestsEvents Invoke the hook for merge requests events.
         * 
         * @return builder
         * 
         */
        public Builder mergeRequestsEvents(@Nullable Output<Boolean> mergeRequestsEvents) {
            $.mergeRequestsEvents = mergeRequestsEvents;
            return this;
        }

        /**
         * @param mergeRequestsEvents Invoke the hook for merge requests events.
         * 
         * @return builder
         * 
         */
        public Builder mergeRequestsEvents(Boolean mergeRequestsEvents) {
            return mergeRequestsEvents(Output.of(mergeRequestsEvents));
        }

        /**
         * @param name Name of the group webhook.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the group webhook.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param noteEvents Invoke the hook for note events.
         * 
         * @return builder
         * 
         */
        public Builder noteEvents(@Nullable Output<Boolean> noteEvents) {
            $.noteEvents = noteEvents;
            return this;
        }

        /**
         * @param noteEvents Invoke the hook for note events.
         * 
         * @return builder
         * 
         */
        public Builder noteEvents(Boolean noteEvents) {
            return noteEvents(Output.of(noteEvents));
        }

        /**
         * @param pipelineEvents Invoke the hook for pipeline events.
         * 
         * @return builder
         * 
         */
        public Builder pipelineEvents(@Nullable Output<Boolean> pipelineEvents) {
            $.pipelineEvents = pipelineEvents;
            return this;
        }

        /**
         * @param pipelineEvents Invoke the hook for pipeline events.
         * 
         * @return builder
         * 
         */
        public Builder pipelineEvents(Boolean pipelineEvents) {
            return pipelineEvents(Output.of(pipelineEvents));
        }

        /**
         * @param pushEvents Invoke the hook for push events.
         * 
         * @return builder
         * 
         */
        public Builder pushEvents(@Nullable Output<Boolean> pushEvents) {
            $.pushEvents = pushEvents;
            return this;
        }

        /**
         * @param pushEvents Invoke the hook for push events.
         * 
         * @return builder
         * 
         */
        public Builder pushEvents(Boolean pushEvents) {
            return pushEvents(Output.of(pushEvents));
        }

        /**
         * @param pushEventsBranchFilter Invoke the hook for push events on matching branches only.
         * 
         * @return builder
         * 
         */
        public Builder pushEventsBranchFilter(@Nullable Output<String> pushEventsBranchFilter) {
            $.pushEventsBranchFilter = pushEventsBranchFilter;
            return this;
        }

        /**
         * @param pushEventsBranchFilter Invoke the hook for push events on matching branches only.
         * 
         * @return builder
         * 
         */
        public Builder pushEventsBranchFilter(String pushEventsBranchFilter) {
            return pushEventsBranchFilter(Output.of(pushEventsBranchFilter));
        }

        /**
         * @param releasesEvents Invoke the hook for release events.
         * 
         * @return builder
         * 
         */
        public Builder releasesEvents(@Nullable Output<Boolean> releasesEvents) {
            $.releasesEvents = releasesEvents;
            return this;
        }

        /**
         * @param releasesEvents Invoke the hook for release events.
         * 
         * @return builder
         * 
         */
        public Builder releasesEvents(Boolean releasesEvents) {
            return releasesEvents(Output.of(releasesEvents));
        }

        /**
         * @param subgroupEvents Invoke the hook for subgroup events.
         * 
         * @return builder
         * 
         */
        public Builder subgroupEvents(@Nullable Output<Boolean> subgroupEvents) {
            $.subgroupEvents = subgroupEvents;
            return this;
        }

        /**
         * @param subgroupEvents Invoke the hook for subgroup events.
         * 
         * @return builder
         * 
         */
        public Builder subgroupEvents(Boolean subgroupEvents) {
            return subgroupEvents(Output.of(subgroupEvents));
        }

        /**
         * @param tagPushEvents Invoke the hook for tag push events.
         * 
         * @return builder
         * 
         */
        public Builder tagPushEvents(@Nullable Output<Boolean> tagPushEvents) {
            $.tagPushEvents = tagPushEvents;
            return this;
        }

        /**
         * @param tagPushEvents Invoke the hook for tag push events.
         * 
         * @return builder
         * 
         */
        public Builder tagPushEvents(Boolean tagPushEvents) {
            return tagPushEvents(Output.of(tagPushEvents));
        }

        /**
         * @param token A token to present when invoking the hook. The token is not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token A token to present when invoking the hook. The token is not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param url The url of the hook to invoke. Forces re-creation to preserve `token`.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The url of the hook to invoke. Forces re-creation to preserve `token`.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param wikiPageEvents Invoke the hook for wiki page events.
         * 
         * @return builder
         * 
         */
        public Builder wikiPageEvents(@Nullable Output<Boolean> wikiPageEvents) {
            $.wikiPageEvents = wikiPageEvents;
            return this;
        }

        /**
         * @param wikiPageEvents Invoke the hook for wiki page events.
         * 
         * @return builder
         * 
         */
        public Builder wikiPageEvents(Boolean wikiPageEvents) {
            return wikiPageEvents(Output.of(wikiPageEvents));
        }

        public GroupHookArgs build() {
            if ($.group == null) {
                throw new MissingRequiredPropertyException("GroupHookArgs", "group");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("GroupHookArgs", "url");
            }
            return $;
        }
    }

}
