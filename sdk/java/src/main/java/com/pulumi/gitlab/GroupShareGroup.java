// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupShareGroupArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupShareGroupState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupShareGroup` resource allows to manage the lifecycle of group shared with another group.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#share-groups-with-groups)
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
 * import com.pulumi.gitlab.GroupShareGroup;
 * import com.pulumi.gitlab.GroupShareGroupArgs;
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
 *         var test = new GroupShareGroup(&#34;test&#34;, GroupShareGroupArgs.builder()        
 *             .groupId(foo.id())
 *             .shareGroupId(bar.id())
 *             .groupAccess(&#34;guest&#34;)
 *             .expiresAt(&#34;2099-01-01&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitLab group shares can be imported using an id made up of `mainGroupId:shareGroupId`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupShareGroup:GroupShareGroup test 12345:1337
 * ```
 * 
 */
@ResourceType(type="gitlab:index/groupShareGroup:GroupShareGroup")
public class GroupShareGroup extends com.pulumi.resources.CustomResource {
    /**
     * Share expiration date. Format: `YYYY-MM-DD`
     * 
     */
    @Export(name="expiresAt", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expiresAt;

    /**
     * @return Share expiration date. Format: `YYYY-MM-DD`
     * 
     */
    public Output<Optional<String>> expiresAt() {
        return Codegen.optional(this.expiresAt);
    }
    /**
     * The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     * 
     */
    @Export(name="groupAccess", refs={String.class}, tree="[0]")
    private Output<String> groupAccess;

    /**
     * @return The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     * 
     */
    public Output<String> groupAccess() {
        return this.groupAccess;
    }
    /**
     * The id of the main group to be shared.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return The id of the main group to be shared.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * The id of the additional group with which the main group will be shared.
     * 
     */
    @Export(name="shareGroupId", refs={Integer.class}, tree="[0]")
    private Output<Integer> shareGroupId;

    /**
     * @return The id of the additional group with which the main group will be shared.
     * 
     */
    public Output<Integer> shareGroupId() {
        return this.shareGroupId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupShareGroup(String name) {
        this(name, GroupShareGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupShareGroup(String name, GroupShareGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupShareGroup(String name, GroupShareGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupShareGroup:GroupShareGroup", name, args == null ? GroupShareGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupShareGroup(String name, Output<String> id, @Nullable GroupShareGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupShareGroup:GroupShareGroup", name, state, makeResourceOptions(options, id));
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
    public static GroupShareGroup get(String name, Output<String> id, @Nullable GroupShareGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupShareGroup(name, id, state, options);
    }
}
