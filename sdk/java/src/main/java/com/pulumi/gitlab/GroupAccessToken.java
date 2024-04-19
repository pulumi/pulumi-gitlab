// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupAccessTokenArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupAccessTokenState;
import com.pulumi.gitlab.outputs.GroupAccessTokenRotationConfiguration;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupAccessToken`resource allows to manage the lifecycle of a group access token.
 * 
 * &gt; Observability scopes are in beta and may not work on all instances. See more details in [the documentation](https://docs.gitlab.com/ee/operations/tracing.html)
 * 
 * &gt; Use `rotation_configuration` to automatically rotate tokens instead of using `timestamp()` as timestamp will cause changes with every plan. `pulumi up` must still be run to rotate the token.
 * 
 * &gt; Due to [Automatic reuse detection](https://docs.gitlab.com/ee/api/group_access_tokens.html#automatic-reuse-detection) it&#39;s possible that a new Group Access Token will immediately be revoked. Check if an old process using the old token is running if this happens.
 * 
 * **Upstream API**: [GitLab REST API](https://docs.gitlab.com/ee/api/group_access_tokens.html)
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
 * import com.pulumi.gitlab.GroupAccessToken;
 * import com.pulumi.gitlab.GroupAccessTokenArgs;
 * import com.pulumi.gitlab.GroupVariable;
 * import com.pulumi.gitlab.GroupVariableArgs;
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
 *         var example = new GroupAccessToken(&#34;example&#34;, GroupAccessTokenArgs.builder()        
 *             .group(&#34;25&#34;)
 *             .name(&#34;Example project access token&#34;)
 *             .expiresAt(&#34;2020-03-14&#34;)
 *             .accessLevel(&#34;developer&#34;)
 *             .scopes(&#34;api&#34;)
 *             .build());
 * 
 *         var exampleGroupVariable = new GroupVariable(&#34;exampleGroupVariable&#34;, GroupVariableArgs.builder()        
 *             .group(&#34;25&#34;)
 *             .key(&#34;gat&#34;)
 *             .value(example.token())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * A GitLab Group Access Token can be imported using a key composed of `&lt;group-id&gt;:&lt;token-id&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupAccessToken:GroupAccessToken example &#34;12345:1&#34;
 * ```
 * 
 * ATTENTION: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
 * 
 */
@ResourceType(type="gitlab:index/groupAccessToken:GroupAccessToken")
public class GroupAccessToken extends com.pulumi.resources.CustomResource {
    /**
     * The access level for the group access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
     * 
     */
    @Export(name="accessLevel", refs={String.class}, tree="[0]")
    private Output<String> accessLevel;

    /**
     * @return The access level for the group access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
     * 
     */
    public Output<String> accessLevel() {
        return this.accessLevel;
    }
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
     * When the token will expire, YYYY-MM-DD format.
     * 
     */
    @Export(name="expiresAt", refs={String.class}, tree="[0]")
    private Output<String> expiresAt;

    /**
     * @return When the token will expire, YYYY-MM-DD format.
     * 
     */
    public Output<String> expiresAt() {
        return this.expiresAt;
    }
    /**
     * The ID or full path of the group.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The ID or full path of the group.
     * 
     */
    public Output<String> group() {
        return this.group;
    }
    /**
     * The name of the group access token.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the group access token.
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
     * The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
     * 
     */
    @Export(name="rotationConfiguration", refs={GroupAccessTokenRotationConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ GroupAccessTokenRotationConfiguration> rotationConfiguration;

    /**
     * @return The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
     * 
     */
    public Output<Optional<GroupAccessTokenRotationConfiguration>> rotationConfiguration() {
        return Codegen.optional(this.rotationConfiguration);
    }
    /**
     * The scopes of the group access token. Valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    @Export(name="scopes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> scopes;

    /**
     * @return The scopes of the group access token. Valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }
    /**
     * The token of the group access token. **Note**: the token is not available for imported resources.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return The token of the group access token. **Note**: the token is not available for imported resources.
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * The user_id associated to the token.
     * 
     */
    @Export(name="userId", refs={Integer.class}, tree="[0]")
    private Output<Integer> userId;

    /**
     * @return The user_id associated to the token.
     * 
     */
    public Output<Integer> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupAccessToken(String name) {
        this(name, GroupAccessTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupAccessToken(String name, GroupAccessTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupAccessToken(String name, GroupAccessTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupAccessToken:GroupAccessToken", name, args == null ? GroupAccessTokenArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupAccessToken(String name, Output<String> id, @Nullable GroupAccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupAccessToken:GroupAccessToken", name, state, makeResourceOptions(options, id));
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
    public static GroupAccessToken get(String name, Output<String> id, @Nullable GroupAccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupAccessToken(name, id, state, options);
    }
}
