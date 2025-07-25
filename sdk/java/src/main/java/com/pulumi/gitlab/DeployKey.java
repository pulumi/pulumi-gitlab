// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.DeployKeyArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.DeployKeyState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.DeployKey` resource manages the lifecycle of a project deploy key.
 * 
 * &gt; To enable an already existing deploy key for another project use the `gitlab.DeployKeyEnable` resource.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/deploy_keys/)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.DeployKey;
 * import com.pulumi.gitlab.DeployKeyArgs;
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
 *         // No expiry
 *         var example = new DeployKey("example", DeployKeyArgs.builder()
 *             .project("example/deploying")
 *             .title("Example deploy key")
 *             .key("ssh-ed25519 AAAA...")
 *             .build());
 * 
 *         // With expiry
 *         var exampleExpires = new DeployKey("exampleExpires", DeployKeyArgs.builder()
 *             .project("example/deploying")
 *             .title("Example deploy key")
 *             .key("ssh-ed25519 AAAA...")
 *             .expiresAt("2025-01-21T00:00:00Z")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_deploy_key`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_deploy_key.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Importing using the CLI is supported with the following syntax:
 * 
 * GitLab deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.
 * 
 * `project_id` can be whatever the [get single project api][get_single_project] takes for
 * 
 * its `:id` value, so for example:
 * 
 * ```sh
 * $ pulumi import gitlab:index/deployKey:DeployKey test 1:3
 * ```
 * 
 * ```sh
 * $ pulumi import gitlab:index/deployKey:DeployKey test richardc/example:3
 * ```
 * 
 */
@ResourceType(type="gitlab:index/deployKey:DeployKey")
public class DeployKey extends com.pulumi.resources.CustomResource {
    /**
     * Allow this deploy key to be used to push changes to the project. Defaults to `false`.
     * 
     */
    @Export(name="canPush", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> canPush;

    /**
     * @return Allow this deploy key to be used to push changes to the project. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> canPush() {
        return Codegen.optional(this.canPush);
    }
    /**
     * The id of the project deploy key.
     * 
     */
    @Export(name="deployKeyId", refs={Integer.class}, tree="[0]")
    private Output<Integer> deployKeyId;

    /**
     * @return The id of the project deploy key.
     * 
     */
    public Output<Integer> deployKeyId() {
        return this.deployKeyId;
    }
    /**
     * Expiration date for the deploy key. Does not expire if no value is provided. Expected in RFC3339 format `(2019-03-15T08:00:00Z)`
     * 
     */
    @Export(name="expiresAt", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expiresAt;

    /**
     * @return Expiration date for the deploy key. Does not expire if no value is provided. Expected in RFC3339 format `(2019-03-15T08:00:00Z)`
     * 
     */
    public Output<Optional<String>> expiresAt() {
        return Codegen.optional(this.expiresAt);
    }
    /**
     * The public ssh key body.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The public ssh key body.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The name or id of the project to add the deploy key to.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The name or id of the project to add the deploy key to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * A title to describe the deploy key with.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return A title to describe the deploy key with.
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DeployKey(java.lang.String name) {
        this(name, DeployKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DeployKey(java.lang.String name, DeployKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DeployKey(java.lang.String name, DeployKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/deployKey:DeployKey", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DeployKey(java.lang.String name, Output<java.lang.String> id, @Nullable DeployKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/deployKey:DeployKey", name, state, makeResourceOptions(options, id), false);
    }

    private static DeployKeyArgs makeArgs(DeployKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DeployKeyArgs.Empty : args;
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
    public static DeployKey get(java.lang.String name, Output<java.lang.String> id, @Nullable DeployKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DeployKey(name, id, state, options);
    }
}
