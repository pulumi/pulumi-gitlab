// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.SystemHookArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.SystemHookState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.SystemHook` resource allows to manage the lifecycle of a system hook.
 * 
 * &gt; This resource requires GitLab 14.9
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/system_hooks.html)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.SystemHook;
 * import com.pulumi.gitlab.SystemHookArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new SystemHook(&#34;example&#34;, SystemHookArgs.builder()        
 *             .enableSslVerification(true)
 *             .mergeRequestsEvents(true)
 *             .pushEvents(true)
 *             .repositoryUpdateEvents(true)
 *             .tagPushEvents(true)
 *             .token(&#34;secret-token&#34;)
 *             .url(&#34;https://example.com/hook-%d&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * You can import a system hook using the hook id `{hook-id}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/systemHook:SystemHook example 42
 * ```
 * 
 * NOTE: the `token` attribute won&#39;t be available for imported resources.
 * 
 */
@ResourceType(type="gitlab:index/systemHook:SystemHook")
public class SystemHook extends com.pulumi.resources.CustomResource {
    /**
     * The date and time the hook was created in ISO8601 format.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The date and time the hook was created in ISO8601 format.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Do SSL verification when triggering the hook.
     * 
     */
    @Export(name="enableSslVerification", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableSslVerification;

    /**
     * @return Do SSL verification when triggering the hook.
     * 
     */
    public Output<Optional<Boolean>> enableSslVerification() {
        return Codegen.optional(this.enableSslVerification);
    }
    /**
     * Trigger hook on merge requests events.
     * 
     */
    @Export(name="mergeRequestsEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mergeRequestsEvents;

    /**
     * @return Trigger hook on merge requests events.
     * 
     */
    public Output<Optional<Boolean>> mergeRequestsEvents() {
        return Codegen.optional(this.mergeRequestsEvents);
    }
    /**
     * When true, the hook fires on push events.
     * 
     */
    @Export(name="pushEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> pushEvents;

    /**
     * @return When true, the hook fires on push events.
     * 
     */
    public Output<Optional<Boolean>> pushEvents() {
        return Codegen.optional(this.pushEvents);
    }
    /**
     * Trigger hook on repository update events.
     * 
     */
    @Export(name="repositoryUpdateEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> repositoryUpdateEvents;

    /**
     * @return Trigger hook on repository update events.
     * 
     */
    public Output<Optional<Boolean>> repositoryUpdateEvents() {
        return Codegen.optional(this.repositoryUpdateEvents);
    }
    /**
     * When true, the hook fires on new tags being pushed.
     * 
     */
    @Export(name="tagPushEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> tagPushEvents;

    /**
     * @return When true, the hook fires on new tags being pushed.
     * 
     */
    public Output<Optional<Boolean>> tagPushEvents() {
        return Codegen.optional(this.tagPushEvents);
    }
    /**
     * Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> token;

    /**
     * @return Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
     * 
     */
    public Output<Optional<String>> token() {
        return Codegen.optional(this.token);
    }
    /**
     * The hook URL.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The hook URL.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SystemHook(String name) {
        this(name, SystemHookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SystemHook(String name, SystemHookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SystemHook(String name, SystemHookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/systemHook:SystemHook", name, args == null ? SystemHookArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SystemHook(String name, Output<String> id, @Nullable SystemHookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/systemHook:SystemHook", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "token"
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
    public static SystemHook get(String name, Output<String> id, @Nullable SystemHookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SystemHook(name, id, state, options);
    }
}
