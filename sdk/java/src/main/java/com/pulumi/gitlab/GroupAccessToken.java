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
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab_group_access`token resource allows to manage the lifecycle of a group access token.
 * 
 * &gt; Group Access Token were introduced in GitLab 14.7
 * 
 * **Upstream API**: [GitLab REST API](https://docs.gitlab.com/ee/api/group_access_tokens.html)
 * 
 * ## Example Usage
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
 *         var exampleGroupAccessToken = new GroupAccessToken(&#34;exampleGroupAccessToken&#34;, GroupAccessTokenArgs.builder()        
 *             .group(&#34;25&#34;)
 *             .expiresAt(&#34;2020-03-14&#34;)
 *             .accessLevel(&#34;developer&#34;)
 *             .scopes(&#34;api&#34;)
 *             .build());
 * 
 *         var exampleGroupVariable = new GroupVariable(&#34;exampleGroupVariable&#34;, GroupVariableArgs.builder()        
 *             .group(&#34;25&#34;)
 *             .key(&#34;gat&#34;)
 *             .value(exampleGroupAccessToken.token())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * A GitLab Group Access Token can be imported using a key composed of `&lt;group-id&gt;:&lt;token-id&gt;`, e.g.
 * 
 * ```sh
 *  $ pulumi import gitlab:index/groupAccessToken:GroupAccessToken example &#34;12345:1&#34;
 * ```
 * 
 *  ATTENTIONthe `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
 * 
 */
@ResourceType(type="gitlab:index/groupAccessToken:GroupAccessToken")
public class GroupAccessToken extends com.pulumi.resources.CustomResource {
    /**
     * The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    @Export(name="accessLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessLevel;

    /**
     * @return The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    public Output<Optional<String>> accessLevel() {
        return Codegen.optional(this.accessLevel);
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
     * The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
     * 
     */
    @Export(name="expiresAt", refs={String.class}, tree="[0]")
    private Output<String> expiresAt;

    /**
     * @return The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
     * 
     */
    public Output<String> expiresAt() {
        return this.expiresAt;
    }
    /**
     * The ID or path of the group to add the group access token to.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The ID or path of the group to add the group access token to.
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
     * The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`.
     * 
     */
    @Export(name="scopes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> scopes;

    /**
     * @return The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`.
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }
    /**
     * The group access token. This is only populated when creating a new group access token. This attribute is not available for imported resources.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return The group access token. This is only populated when creating a new group access token. This attribute is not available for imported resources.
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * The user id associated to the token.
     * 
     */
    @Export(name="userId", refs={Integer.class}, tree="[0]")
    private Output<Integer> userId;

    /**
     * @return The user id associated to the token.
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
