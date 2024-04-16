// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.UserGpgKeyArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.UserGpgKeyState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.UserGpgKey` resource allows to manage the lifecycle of a GPG key assigned to the current user or a specific user.
 * 
 * &gt; Managing GPG keys for arbitrary users requires admin privileges.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#get-a-specific-gpg-key)
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
 * import com.pulumi.gitlab.GitlabFunctions;
 * import com.pulumi.gitlab.inputs.GetUserArgs;
 * import com.pulumi.gitlab.UserGpgKey;
 * import com.pulumi.gitlab.UserGpgKeyArgs;
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
 *         final var example = GitlabFunctions.getUser(GetUserArgs.builder()
 *             .username(&#34;example-user&#34;)
 *             .build());
 * 
 *         // Manages a GPG key for the specified user. An admin token is required if `user_id` is specified.
 *         var exampleUserGpgKey = new UserGpgKey(&#34;exampleUserGpgKey&#34;, UserGpgKeyArgs.builder()        
 *             .userId(example.applyValue(getUserResult -&gt; getUserResult.id()))
 *             .key(&#34;&#34;&#34;
 * -----BEGIN PGP PUBLIC KEY BLOCK-----
 * ...
 * -----END PGP PUBLIC KEY BLOCK-----            &#34;&#34;&#34;)
 *             .build());
 * 
 *         // Manages a GPG key for the current user
 *         var exampleUser = new UserGpgKey(&#34;exampleUser&#34;, UserGpgKeyArgs.builder()        
 *             .key(&#34;&#34;&#34;
 * -----BEGIN PGP PUBLIC KEY BLOCK-----
 * ...
 * -----END PGP PUBLIC KEY BLOCK-----            &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * You can import a GPG key for a specific user using an id made up of `{user-id}:{key}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/userGpgKey:UserGpgKey example 42:1
 * ```
 * 
 * Alternatively, you can import a GPG key for the current user using an id made up of `{key}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/userGpgKey:UserGpgKey example_user 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/userGpgKey:UserGpgKey")
public class UserGpgKey extends com.pulumi.resources.CustomResource {
    /**
     * The time when this key was created in GitLab.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The time when this key was created in GitLab.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The armored GPG public key.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The armored GPG public key.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The ID of the GPG key.
     * 
     */
    @Export(name="keyId", refs={Integer.class}, tree="[0]")
    private Output<Integer> keyId;

    /**
     * @return The ID of the GPG key.
     * 
     */
    public Output<Integer> keyId() {
        return this.keyId;
    }
    /**
     * The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the specified user, and an admin token is required.
     * 
     */
    @Export(name="userId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> userId;

    /**
     * @return The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the specified user, and an admin token is required.
     * 
     */
    public Output<Optional<Integer>> userId() {
        return Codegen.optional(this.userId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserGpgKey(String name) {
        this(name, UserGpgKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserGpgKey(String name, UserGpgKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserGpgKey(String name, UserGpgKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/userGpgKey:UserGpgKey", name, args == null ? UserGpgKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserGpgKey(String name, Output<String> id, @Nullable UserGpgKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/userGpgKey:UserGpgKey", name, state, makeResourceOptions(options, id));
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
    public static UserGpgKey get(String name, Output<String> id, @Nullable UserGpgKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserGpgKey(name, id, state, options);
    }
}
