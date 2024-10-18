// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupServiceAccountAccessTokenArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupServiceAccountAccessTokenState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupServiceAccountAccessToken` resource allows to manage the lifecycle of a group service account access token.
 * 
 * &gt; Use of the `timestamp()` function with expires_at will cause the resource to be re-created with every apply, it&#39;s recommended to use `plantimestamp()` or a static value instead.
 * 
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/group_service_accounts.html#create-a-personal-access-token-for-a-service-account-user)
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
 * import com.pulumi.gitlab.Group;
 * import com.pulumi.gitlab.GroupArgs;
 * import com.pulumi.gitlab.GroupServiceAccount;
 * import com.pulumi.gitlab.GroupServiceAccountArgs;
 * import com.pulumi.gitlab.GroupServiceAccountAccessToken;
 * import com.pulumi.gitlab.GroupServiceAccountAccessTokenArgs;
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
 *         var example = new Group("example", GroupArgs.builder()
 *             .name("example")
 *             .path("example")
 *             .description("An example group")
 *             .build());
 * 
 *         var example_sa = new GroupServiceAccount("example-sa", GroupServiceAccountArgs.builder()
 *             .group(example.id())
 *             .name("example-name")
 *             .username("example-username")
 *             .build());
 * 
 *         var example_sa_token = new GroupServiceAccountAccessToken("example-sa-token", GroupServiceAccountAccessTokenArgs.builder()
 *             .group(example.id())
 *             .userId(example_sa.id())
 *             .name("Example personal access token")
 *             .expiresAt("2020-03-14")
 *             .scopes("api")
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
 * ```sh
 * $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken You can import a service account access token using `&lt;resource&gt; &lt;id&gt;`. The
 * ```
 * 
 * `id` is in the form of &lt;group_id&gt;:&lt;service_account_id&gt;:&lt;access_token_id&gt;
 * 
 * Importing an access token does not import the access token value.
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken example 1:2:3
 * ```
 * 
 */
@ResourceType(type="gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken")
public class GroupServiceAccountAccessToken extends com.pulumi.resources.CustomResource {
    /**
     * True if the token is active.
     * 
     */
    @Export(name="active", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> active;

    /**
     * @return True if the token is active.
     * 
     */
    public Output<Boolean> active() {
        return this.active;
    }
    /**
     * Time the token has been created, RFC3339 format.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Time the token has been created, RFC3339 format.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
     * 
     */
    @Export(name="expiresAt", refs={String.class}, tree="[0]")
    private Output<String> expiresAt;

    /**
     * @return The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
     * 
     */
    public Output<String> expiresAt() {
        return this.expiresAt;
    }
    /**
     * The ID or URL-encoded path of the group containing the service account. Must be a top level group.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The ID or URL-encoded path of the group containing the service account. Must be a top level group.
     * 
     */
    public Output<String> group() {
        return this.group;
    }
    /**
     * The name of the personal access token.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the personal access token.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * True if the token is revoked.
     * 
     */
    @Export(name="revoked", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> revoked;

    /**
     * @return True if the token is revoked.
     * 
     */
    public Output<Boolean> revoked() {
        return this.revoked;
    }
    /**
     * The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    @Export(name="scopes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> scopes;

    /**
     * @return The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }
    /**
     * The token of the group service account access token. **Note**: the token is not available for imported resources.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return The token of the group service account access token. **Note**: the token is not available for imported resources.
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * The ID of a service account user.
     * 
     */
    @Export(name="userId", refs={Integer.class}, tree="[0]")
    private Output<Integer> userId;

    /**
     * @return The ID of a service account user.
     * 
     */
    public Output<Integer> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupServiceAccountAccessToken(java.lang.String name) {
        this(name, GroupServiceAccountAccessTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupServiceAccountAccessToken(java.lang.String name, GroupServiceAccountAccessTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupServiceAccountAccessToken(java.lang.String name, GroupServiceAccountAccessTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GroupServiceAccountAccessToken(java.lang.String name, Output<java.lang.String> id, @Nullable GroupServiceAccountAccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupServiceAccountAccessTokenArgs makeArgs(GroupServiceAccountAccessTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupServiceAccountAccessTokenArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static GroupServiceAccountAccessToken get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupServiceAccountAccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupServiceAccountAccessToken(name, id, state, options);
    }
}
