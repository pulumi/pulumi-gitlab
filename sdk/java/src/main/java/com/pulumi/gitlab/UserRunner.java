// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.UserRunnerArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.UserRunnerState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.UserRunner` resource allows creating a GitLab runner using the new [GitLab Runner Registration Flow](https://docs.gitlab.com/ee/ci/runners/new_creation_workflow.html).
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#create-a-runner)
 * 
 */
@ResourceType(type="gitlab:index/userRunner:UserRunner")
public class UserRunner extends com.pulumi.resources.CustomResource {
    /**
     * The access level of the runner. Valid values are: `not_protected`, `ref_protected`.
     * 
     */
    @Export(name="accessLevel", refs={String.class}, tree="[0]")
    private Output<String> accessLevel;

    /**
     * @return The access level of the runner. Valid values are: `not_protected`, `ref_protected`.
     * 
     */
    public Output<String> accessLevel() {
        return this.accessLevel;
    }
    /**
     * Description of the runner.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Description of the runner.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The ID of the group that the runner is created in. Required if runner*type is group*type.
     * 
     */
    @Export(name="groupId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> groupId;

    /**
     * @return The ID of the group that the runner is created in. Required if runner*type is group*type.
     * 
     */
    public Output<Optional<Integer>> groupId() {
        return Codegen.optional(this.groupId);
    }
    /**
     * Specifies if the runner should be locked for the current project.
     * 
     */
    @Export(name="locked", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> locked;

    /**
     * @return Specifies if the runner should be locked for the current project.
     * 
     */
    public Output<Boolean> locked() {
        return this.locked;
    }
    /**
     * Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
     * 
     */
    @Export(name="maximumTimeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> maximumTimeout;

    /**
     * @return Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
     * 
     */
    public Output<Integer> maximumTimeout() {
        return this.maximumTimeout;
    }
    /**
     * Specifies if the runner should ignore new jobs.
     * 
     */
    @Export(name="paused", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> paused;

    /**
     * @return Specifies if the runner should ignore new jobs.
     * 
     */
    public Output<Boolean> paused() {
        return this.paused;
    }
    /**
     * The ID of the project that the runner is created in. Required if runner*type is project*type.
     * 
     */
    @Export(name="projectId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> projectId;

    /**
     * @return The ID of the project that the runner is created in. Required if runner*type is project*type.
     * 
     */
    public Output<Optional<Integer>> projectId() {
        return Codegen.optional(this.projectId);
    }
    /**
     * The scope of the runner. Valid values are: `instance_type`, `group_type`, `project_type`.
     * 
     */
    @Export(name="runnerType", refs={String.class}, tree="[0]")
    private Output<String> runnerType;

    /**
     * @return The scope of the runner. Valid values are: `instance_type`, `group_type`, `project_type`.
     * 
     */
    public Output<String> runnerType() {
        return this.runnerType;
    }
    /**
     * A list of runner tags.
     * 
     */
    @Export(name="tagLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> tagLists;

    /**
     * @return A list of runner tags.
     * 
     */
    public Output<List<String>> tagLists() {
        return this.tagLists;
    }
    /**
     * The authentication token to use when setting up a new runner with this configuration. This value cannot be imported.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return The authentication token to use when setting up a new runner with this configuration. This value cannot be imported.
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * Specifies if the runner should handle untagged jobs.
     * 
     */
    @Export(name="untagged", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> untagged;

    /**
     * @return Specifies if the runner should handle untagged jobs.
     * 
     */
    public Output<Boolean> untagged() {
        return this.untagged;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserRunner(String name) {
        this(name, UserRunnerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserRunner(String name, UserRunnerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserRunner(String name, UserRunnerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/userRunner:UserRunner", name, args == null ? UserRunnerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserRunner(String name, Output<String> id, @Nullable UserRunnerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/userRunner:UserRunner", name, state, makeResourceOptions(options, id));
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
    public static UserRunner get(String name, Output<String> id, @Nullable UserRunnerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserRunner(name, id, state, options);
    }
}