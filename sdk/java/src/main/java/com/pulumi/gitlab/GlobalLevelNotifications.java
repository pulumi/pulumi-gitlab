// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GlobalLevelNotificationsArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GlobalLevelNotificationsState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.GlobalLevelNotifications` resource allows to manage global notifications.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/notification_settings/#group--project-level-notification-settings)
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_global_level_notifications`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_global_level_notifications.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * Note: You can import a global notification state using &#34;gitlab&#34; as the ID.
 * 
 * The ID will always be gitlab, because the global notificatio only exists
 * 
 * once per user
 * 
 * ```sh
 * $ pulumi import gitlab:index/globalLevelNotifications:GlobalLevelNotifications example gitlab
 * ```
 * 
 */
@ResourceType(type="gitlab:index/globalLevelNotifications:GlobalLevelNotifications")
public class GlobalLevelNotifications extends com.pulumi.resources.CustomResource {
    /**
     * Enable notifications for closed issues. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="closeIssue", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> closeIssue;

    /**
     * @return Enable notifications for closed issues. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> closeIssue() {
        return this.closeIssue;
    }
    /**
     * Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="closeMergeRequest", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> closeMergeRequest;

    /**
     * @return Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> closeMergeRequest() {
        return this.closeMergeRequest;
    }
    /**
     * Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="failedPipeline", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> failedPipeline;

    /**
     * @return Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> failedPipeline() {
        return this.failedPipeline;
    }
    /**
     * Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="fixedPipeline", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> fixedPipeline;

    /**
     * @return Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> fixedPipeline() {
        return this.fixedPipeline;
    }
    /**
     * Enable notifications for due issues. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="issueDue", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> issueDue;

    /**
     * @return Enable notifications for due issues. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> issueDue() {
        return this.issueDue;
    }
    /**
     * The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
     * 
     */
    @Export(name="level", refs={String.class}, tree="[0]")
    private Output<String> level;

    /**
     * @return The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
     * 
     */
    public Output<String> level() {
        return this.level;
    }
    /**
     * Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="mergeMergeRequest", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> mergeMergeRequest;

    /**
     * @return Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> mergeMergeRequest() {
        return this.mergeMergeRequest;
    }
    /**
     * Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="mergeWhenPipelineSucceeds", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> mergeWhenPipelineSucceeds;

    /**
     * @return Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> mergeWhenPipelineSucceeds() {
        return this.mergeWhenPipelineSucceeds;
    }
    /**
     * Enable notifications for moved projects. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="movedProject", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> movedProject;

    /**
     * @return Enable notifications for moved projects. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> movedProject() {
        return this.movedProject;
    }
    /**
     * Enable notifications for new issues. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="newIssue", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> newIssue;

    /**
     * @return Enable notifications for new issues. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> newIssue() {
        return this.newIssue;
    }
    /**
     * Enable notifications for new merge requests. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="newMergeRequest", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> newMergeRequest;

    /**
     * @return Enable notifications for new merge requests. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> newMergeRequest() {
        return this.newMergeRequest;
    }
    /**
     * Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="newNote", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> newNote;

    /**
     * @return Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> newNote() {
        return this.newNote;
    }
    /**
     * Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="pushToMergeRequest", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> pushToMergeRequest;

    /**
     * @return Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> pushToMergeRequest() {
        return this.pushToMergeRequest;
    }
    /**
     * Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="reassignIssue", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reassignIssue;

    /**
     * @return Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> reassignIssue() {
        return this.reassignIssue;
    }
    /**
     * Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="reassignMergeRequest", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reassignMergeRequest;

    /**
     * @return Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> reassignMergeRequest() {
        return this.reassignMergeRequest;
    }
    /**
     * Enable notifications for reopened issues. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="reopenIssue", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reopenIssue;

    /**
     * @return Enable notifications for reopened issues. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> reopenIssue() {
        return this.reopenIssue;
    }
    /**
     * Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="reopenMergeRequest", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reopenMergeRequest;

    /**
     * @return Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> reopenMergeRequest() {
        return this.reopenMergeRequest;
    }
    /**
     * Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
     * 
     */
    @Export(name="successPipeline", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> successPipeline;

    /**
     * @return Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
     * 
     */
    public Output<Boolean> successPipeline() {
        return this.successPipeline;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GlobalLevelNotifications(String name) {
        this(name, GlobalLevelNotificationsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GlobalLevelNotifications(String name, @Nullable GlobalLevelNotificationsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GlobalLevelNotifications(String name, @Nullable GlobalLevelNotificationsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/globalLevelNotifications:GlobalLevelNotifications", name, args == null ? GlobalLevelNotificationsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GlobalLevelNotifications(String name, Output<String> id, @Nullable GlobalLevelNotificationsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/globalLevelNotifications:GlobalLevelNotifications", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static GlobalLevelNotifications get(String name, Output<String> id, @Nullable GlobalLevelNotificationsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GlobalLevelNotifications(name, id, state, options);
    }
}
