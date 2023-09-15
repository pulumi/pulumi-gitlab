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
 * The `gitlab.DeployKey` resource allows to manage the lifecycle of a deploy key.
 * 
 * &gt; To enable an already existing deploy key for another project use the `gitlab_project_deploy_key` resource.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_keys.html)
 * 
 * ## Example Usage
 * ```java
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
 *         var example = new DeployKey(&#34;example&#34;, DeployKeyArgs.builder()        
 *             .key(&#34;ssh-ed25519 AAAA...&#34;)
 *             .project(&#34;example/deploying&#34;)
 *             .title(&#34;Example deploy key&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitLab deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g. `project_id` can be whatever the [get single project api][get_single_project] takes for its `:id` value, so for example
 * 
 * ```sh
 *  $ pulumi import gitlab:index/deployKey:DeployKey test 1:3
 * ```
 * 
 * ```sh
 *  $ pulumi import gitlab:index/deployKey:DeployKey test richardc/example:3
 * ```
 * 
 */
@ResourceType(type="gitlab:index/deployKey:DeployKey")
public class DeployKey extends com.pulumi.resources.CustomResource {
    /**
     * Allow this deploy key to be used to push changes to the project. Defaults to `false`.
     * 
     */
    @Export(name="canPush", type=Boolean.class, parameters={})
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
    @Export(name="deployKeyId", type=Integer.class, parameters={})
    private Output<Integer> deployKeyId;

    /**
     * @return The id of the project deploy key.
     * 
     */
    public Output<Integer> deployKeyId() {
        return this.deployKeyId;
    }
    /**
     * The public ssh key body.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
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
    @Export(name="project", type=String.class, parameters={})
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
    @Export(name="title", type=String.class, parameters={})
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
    public DeployKey(String name) {
        this(name, DeployKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DeployKey(String name, DeployKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DeployKey(String name, DeployKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/deployKey:DeployKey", name, args == null ? DeployKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DeployKey(String name, Output<String> id, @Nullable DeployKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/deployKey:DeployKey", name, state, makeResourceOptions(options, id));
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
    public static DeployKey get(String name, Output<String> id, @Nullable DeployKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DeployKey(name, id, state, options);
    }
}
