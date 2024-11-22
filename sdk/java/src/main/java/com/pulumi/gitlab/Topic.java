// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.TopicArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.TopicState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.Topic` resource allows to manage the lifecycle of topics that are then assignable to projects.
 * 
 * &gt; Topics are the successors for project tags. Aside from avoiding terminology collisions with Git tags, they are more descriptive and better searchable.
 * 
 * &gt; Deleting a topic was implemented in GitLab 14.9. For older versions of GitLab set `soft_destroy = true` to empty out a topic instead of deleting it.
 * 
 * **Upstream API**: [GitLab REST API docs for topics](https://docs.gitlab.com/ee/api/topics.html)
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_topic`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_topic.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * ```sh
 * $ pulumi import gitlab:index/topic:Topic You can import a topic to terraform state using `&lt;resource&gt; &lt;id&gt;`.
 * ```
 * 
 * The `id` must be an integer for the id of the topic you want to import,
 * 
 * for example:
 * 
 * ```sh
 * $ pulumi import gitlab:index/topic:Topic functional_programming 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/topic:Topic")
public class Topic extends com.pulumi.resources.CustomResource {
    /**
     * A local path to the avatar image to upload. **Note**: not available for imported resources.
     * 
     */
    @Export(name="avatar", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> avatar;

    /**
     * @return A local path to the avatar image to upload. **Note**: not available for imported resources.
     * 
     */
    public Output<Optional<String>> avatar() {
        return Codegen.optional(this.avatar);
    }
    /**
     * The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
     * 
     */
    @Export(name="avatarHash", refs={String.class}, tree="[0]")
    private Output<String> avatarHash;

    /**
     * @return The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
     * 
     */
    public Output<String> avatarHash() {
        return this.avatarHash;
    }
    /**
     * The URL of the avatar image.
     * 
     */
    @Export(name="avatarUrl", refs={String.class}, tree="[0]")
    private Output<String> avatarUrl;

    /**
     * @return The URL of the avatar image.
     * 
     */
    public Output<String> avatarUrl() {
        return this.avatarUrl;
    }
    /**
     * A text describing the topic.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A text describing the topic.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The topic&#39;s name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The topic&#39;s name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Empty the topics fields instead of deleting it.
     * 
     * @deprecated
     * GitLab 14.9 introduced the proper deletion of topics. This field is no longer needed.
     * 
     */
    @Deprecated /* GitLab 14.9 introduced the proper deletion of topics. This field is no longer needed. */
    @Export(name="softDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> softDestroy;

    /**
     * @return Empty the topics fields instead of deleting it.
     * 
     */
    public Output<Optional<Boolean>> softDestroy() {
        return Codegen.optional(this.softDestroy);
    }
    /**
     * The topic&#39;s description. Requires at least GitLab 15.0 for which it&#39;s a required argument.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> title;

    /**
     * @return The topic&#39;s description. Requires at least GitLab 15.0 for which it&#39;s a required argument.
     * 
     */
    public Output<Optional<String>> title() {
        return Codegen.optional(this.title);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Topic(java.lang.String name) {
        this(name, TopicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Topic(java.lang.String name, @Nullable TopicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Topic(java.lang.String name, @Nullable TopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/topic:Topic", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Topic(java.lang.String name, Output<java.lang.String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/topic:Topic", name, state, makeResourceOptions(options, id), false);
    }

    private static TopicArgs makeArgs(@Nullable TopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TopicArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Topic get(java.lang.String name, Output<java.lang.String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Topic(name, id, state, options);
    }
}
