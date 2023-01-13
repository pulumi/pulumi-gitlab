// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.RunnerArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.RunnerState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.Runner` resource allows to manage the lifecycle of a runner.
 * 
 * A runner can either be registered at an instance level or group level.
 * The runner will be registered at a group level if the token used is from a group, or at an instance level if the token used is for the instance.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/runners.html#register-a-new-runner)
 * 
 * ## Import
 * 
 * A GitLab Runner can be imported using the runner&#39;s ID, eg
 * 
 * ```sh
 *  $ pulumi import gitlab:index/runner:Runner this 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/runner:Runner")
public class Runner extends com.pulumi.resources.CustomResource {
    /**
     * The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
     * 
     */
    @Export(name="accessLevel", type=String.class, parameters={})
    private Output<String> accessLevel;

    /**
     * @return The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
     * 
     */
    public Output<String> accessLevel() {
        return this.accessLevel;
    }
    /**
     * The authentication token used for building a config.toml file. This value is not present when imported.
     * 
     */
    @Export(name="authenticationToken", type=String.class, parameters={})
    private Output<String> authenticationToken;

    /**
     * @return The authentication token used for building a config.toml file. This value is not present when imported.
     * 
     */
    public Output<String> authenticationToken() {
        return this.authenticationToken;
    }
    /**
     * The runner&#39;s description.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The runner&#39;s description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether the runner should be locked for current project.
     * 
     */
    @Export(name="locked", type=Boolean.class, parameters={})
    private Output<Boolean> locked;

    /**
     * @return Whether the runner should be locked for current project.
     * 
     */
    public Output<Boolean> locked() {
        return this.locked;
    }
    /**
     * Maximum timeout set when this runner handles the job.
     * 
     */
    @Export(name="maximumTimeout", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> maximumTimeout;

    /**
     * @return Maximum timeout set when this runner handles the job.
     * 
     */
    public Output<Optional<Integer>> maximumTimeout() {
        return Codegen.optional(this.maximumTimeout);
    }
    /**
     * Whether the runner should ignore new jobs.
     * 
     */
    @Export(name="paused", type=Boolean.class, parameters={})
    private Output<Boolean> paused;

    /**
     * @return Whether the runner should ignore new jobs.
     * 
     */
    public Output<Boolean> paused() {
        return this.paused;
    }
    /**
     * The registration token used to register the runner.
     * 
     */
    @Export(name="registrationToken", type=String.class, parameters={})
    private Output<String> registrationToken;

    /**
     * @return The registration token used to register the runner.
     * 
     */
    public Output<String> registrationToken() {
        return this.registrationToken;
    }
    /**
     * Whether the runner should handle untagged jobs.
     * 
     */
    @Export(name="runUntagged", type=Boolean.class, parameters={})
    private Output<Boolean> runUntagged;

    /**
     * @return Whether the runner should handle untagged jobs.
     * 
     */
    public Output<Boolean> runUntagged() {
        return this.runUntagged;
    }
    /**
     * The status of runners to show, one of: online and offline. active and paused are also possible values
     * 			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of runners to show, one of: online and offline. active and paused are also possible values
     * 			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * List of runner’s tags.
     * 
     */
    @Export(name="tagLists", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tagLists;

    /**
     * @return List of runner’s tags.
     * 
     */
    public Output<Optional<List<String>>> tagLists() {
        return Codegen.optional(this.tagLists);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Runner(String name) {
        this(name, RunnerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Runner(String name, RunnerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Runner(String name, RunnerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/runner:Runner", name, args == null ? RunnerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Runner(String name, Output<String> id, @Nullable RunnerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/runner:Runner", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "authenticationToken",
                "registrationToken"
            ))
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
    public static Runner get(String name, Output<String> id, @Nullable RunnerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Runner(name, id, state, options);
    }
}