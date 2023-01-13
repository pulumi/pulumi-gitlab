// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupMembershipArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupMembershipState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupMembership` resource allows to manage the lifecycle of a users group membersip.
 * 
 * &gt; If a group should grant membership to another group use the `gitlab.GroupShareGroup` resource instead.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.GroupMembership;
 * import com.pulumi.gitlab.GroupMembershipArgs;
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
 *         var test = new GroupMembership(&#34;test&#34;, GroupMembershipArgs.builder()        
 *             .accessLevel(&#34;guest&#34;)
 *             .expiresAt(&#34;2020-12-31&#34;)
 *             .groupId(&#34;12345&#34;)
 *             .userId(1337)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitLab group membership can be imported using an id made up of `group_id:user_id`, e.g.
 * 
 * ```sh
 *  $ pulumi import gitlab:index/groupMembership:GroupMembership test &#34;12345:1337&#34;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/groupMembership:GroupMembership")
public class GroupMembership extends com.pulumi.resources.CustomResource {
    /**
     * Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
     * 
     */
    @Export(name="accessLevel", type=String.class, parameters={})
    private Output<String> accessLevel;

    /**
     * @return Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
     * 
     */
    public Output<String> accessLevel() {
        return this.accessLevel;
    }
    /**
     * Expiration date for the group membership. Format: `YYYY-MM-DD`
     * 
     */
    @Export(name="expiresAt", type=String.class, parameters={})
    private Output</* @Nullable */ String> expiresAt;

    /**
     * @return Expiration date for the group membership. Format: `YYYY-MM-DD`
     * 
     */
    public Output<Optional<String>> expiresAt() {
        return Codegen.optional(this.expiresAt);
    }
    /**
     * The id of the group.
     * 
     */
    @Export(name="groupId", type=String.class, parameters={})
    private Output<String> groupId;

    /**
     * @return The id of the group.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
     * 
     */
    @Export(name="skipSubresourcesOnDestroy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> skipSubresourcesOnDestroy;

    /**
     * @return Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
     * 
     */
    public Output<Optional<Boolean>> skipSubresourcesOnDestroy() {
        return Codegen.optional(this.skipSubresourcesOnDestroy);
    }
    /**
     * Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
     * 
     */
    @Export(name="unassignIssuablesOnDestroy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> unassignIssuablesOnDestroy;

    /**
     * @return Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
     * 
     */
    public Output<Optional<Boolean>> unassignIssuablesOnDestroy() {
        return Codegen.optional(this.unassignIssuablesOnDestroy);
    }
    /**
     * The id of the user.
     * 
     */
    @Export(name="userId", type=Integer.class, parameters={})
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
    public GroupMembership(String name) {
        this(name, GroupMembershipArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupMembership(String name, GroupMembershipArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupMembership(String name, GroupMembershipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupMembership:GroupMembership", name, args == null ? GroupMembershipArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupMembership(String name, Output<String> id, @Nullable GroupMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupMembership:GroupMembership", name, state, makeResourceOptions(options, id));
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
    public static GroupMembership get(String name, Output<String> id, @Nullable GroupMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupMembership(name, id, state, options);
    }
}