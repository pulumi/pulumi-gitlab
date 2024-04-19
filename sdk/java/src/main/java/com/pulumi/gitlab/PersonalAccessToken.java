// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.PersonalAccessTokenArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.PersonalAccessTokenState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * The `gitlab.PersonalAccessToken` resource allows to manage the lifecycle of a personal access token for a specified user.
 * 
 * &gt; This resource requires administration privileges.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/personal_access_tokens.html)
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
 * import com.pulumi.gitlab.PersonalAccessToken;
 * import com.pulumi.gitlab.PersonalAccessTokenArgs;
 * import com.pulumi.gitlab.ProjectVariable;
 * import com.pulumi.gitlab.ProjectVariableArgs;
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
 *         var example = new PersonalAccessToken(&#34;example&#34;, PersonalAccessTokenArgs.builder()        
 *             .userId(&#34;25&#34;)
 *             .name(&#34;Example personal access token&#34;)
 *             .expiresAt(&#34;2020-03-14&#34;)
 *             .scopes(&#34;api&#34;)
 *             .build());
 * 
 *         var exampleProjectVariable = new ProjectVariable(&#34;exampleProjectVariable&#34;, ProjectVariableArgs.builder()        
 *             .project(exampleGitlabProject.id())
 *             .key(&#34;pat&#34;)
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
 * A GitLab Personal Access Token can be imported using a key composed of `&lt;user-id&gt;:&lt;token-id&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/personalAccessToken:PersonalAccessToken example &#34;12345:1&#34;
 * ```
 * 
 * NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
 * 
 */
@ResourceType(type="gitlab:index/personalAccessToken:PersonalAccessToken")
public class PersonalAccessToken extends com.pulumi.resources.CustomResource {
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
     * The scope for the personal access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `create_runner`.
     * 
     */
    @Export(name="scopes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> scopes;

    /**
     * @return The scope for the personal access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `create_runner`.
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }
    /**
     * The personal access token. This is only populated when creating a new personal access token. This attribute is not available for imported resources.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return The personal access token. This is only populated when creating a new personal access token. This attribute is not available for imported resources.
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * The id of the user.
     * 
     */
    @Export(name="userId", refs={Integer.class}, tree="[0]")
    private Output<Integer> userId;

    /**
     * @return The id of the user.
     * 
     */
    public Output<Integer> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PersonalAccessToken(String name) {
        this(name, PersonalAccessTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PersonalAccessToken(String name, PersonalAccessTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PersonalAccessToken(String name, PersonalAccessTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/personalAccessToken:PersonalAccessToken", name, args == null ? PersonalAccessTokenArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PersonalAccessToken(String name, Output<String> id, @Nullable PersonalAccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/personalAccessToken:PersonalAccessToken", name, state, makeResourceOptions(options, id));
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
    public static PersonalAccessToken get(String name, Output<String> id, @Nullable PersonalAccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PersonalAccessToken(name, id, state, options);
    }
}
