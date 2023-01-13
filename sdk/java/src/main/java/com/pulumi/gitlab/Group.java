// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.Group` resource allows to manage the lifecycle of a group.
 * 
 * &gt; On GitLab SaaS, you must use the GitLab UI to create groups without a parent group. You cannot use this provider nor the API to do this.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.Group;
 * import com.pulumi.gitlab.GroupArgs;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectArgs;
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
 *         var exampleGroup = new Group(&#34;exampleGroup&#34;, GroupArgs.builder()        
 *             .path(&#34;example&#34;)
 *             .description(&#34;An example group&#34;)
 *             .build());
 * 
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .description(&#34;An example project&#34;)
 *             .namespaceId(exampleGroup.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import gitlab:index/group:Group You can import a group state using `&lt;resource&gt; &lt;id&gt;`. The
 * ```
 * 
 *  `id` can be whatever the [details of a group][details_of_a_group] api takes for its `:id` value, so for example
 * 
 * ```sh
 *  $ pulumi import gitlab:index/group:Group example example
 * ```
 * 
 */
@ResourceType(type="gitlab:index/group:Group")
public class Group extends com.pulumi.resources.CustomResource {
    /**
     * Defaults to false. Default to Auto DevOps pipeline for all projects within this group.
     * 
     */
    @Export(name="autoDevopsEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoDevopsEnabled;

    /**
     * @return Defaults to false. Default to Auto DevOps pipeline for all projects within this group.
     * 
     */
    public Output<Optional<Boolean>> autoDevopsEnabled() {
        return Codegen.optional(this.autoDevopsEnabled);
    }
    /**
     * Defaults to 2. See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection
     * 
     */
    @Export(name="defaultBranchProtection", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> defaultBranchProtection;

    /**
     * @return Defaults to 2. See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection
     * 
     */
    public Output<Optional<Integer>> defaultBranchProtection() {
        return Codegen.optional(this.defaultBranchProtection);
    }
    /**
     * The description of the group.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Defaults to false. Disable email notifications.
     * 
     */
    @Export(name="emailsDisabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> emailsDisabled;

    /**
     * @return Defaults to false. Disable email notifications.
     * 
     */
    public Output<Optional<Boolean>> emailsDisabled() {
        return Codegen.optional(this.emailsDisabled);
    }
    /**
     * The full name of the group.
     * 
     */
    @Export(name="fullName", type=String.class, parameters={})
    private Output<String> fullName;

    /**
     * @return The full name of the group.
     * 
     */
    public Output<String> fullName() {
        return this.fullName;
    }
    /**
     * The full path of the group.
     * 
     */
    @Export(name="fullPath", type=String.class, parameters={})
    private Output<String> fullPath;

    /**
     * @return The full path of the group.
     * 
     */
    public Output<String> fullPath() {
        return this.fullPath;
    }
    /**
     * Defaults to true. Enable/disable Large File Storage (LFS) for the projects in this group.
     * 
     */
    @Export(name="lfsEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> lfsEnabled;

    /**
     * @return Defaults to true. Enable/disable Large File Storage (LFS) for the projects in this group.
     * 
     */
    public Output<Optional<Boolean>> lfsEnabled() {
        return Codegen.optional(this.lfsEnabled);
    }
    /**
     * Defaults to false. Disable the capability of a group from getting mentioned.
     * 
     */
    @Export(name="mentionsDisabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> mentionsDisabled;

    /**
     * @return Defaults to false. Disable the capability of a group from getting mentioned.
     * 
     */
    public Output<Optional<Boolean>> mentionsDisabled() {
        return Codegen.optional(this.mentionsDisabled);
    }
    /**
     * The name of this group.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of this group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Id of the parent group (creates a nested group).
     * 
     */
    @Export(name="parentId", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> parentId;

    /**
     * @return Id of the parent group (creates a nested group).
     * 
     */
    public Output<Optional<Integer>> parentId() {
        return Codegen.optional(this.parentId);
    }
    /**
     * The path of the group.
     * 
     */
    @Export(name="path", type=String.class, parameters={})
    private Output<String> path;

    /**
     * @return The path of the group.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
     * 
     */
    @Export(name="preventForkingOutsideGroup", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> preventForkingOutsideGroup;

    /**
     * @return Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
     * 
     */
    public Output<Optional<Boolean>> preventForkingOutsideGroup() {
        return Codegen.optional(this.preventForkingOutsideGroup);
    }
    /**
     * Defaults to maintainer. Determine if developers can create projects in the group.
     * 
     */
    @Export(name="projectCreationLevel", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectCreationLevel;

    /**
     * @return Defaults to maintainer. Determine if developers can create projects in the group.
     * 
     */
    public Output<Optional<String>> projectCreationLevel() {
        return Codegen.optional(this.projectCreationLevel);
    }
    /**
     * Defaults to false. Allow users to request member access.
     * 
     */
    @Export(name="requestAccessEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requestAccessEnabled;

    /**
     * @return Defaults to false. Allow users to request member access.
     * 
     */
    public Output<Optional<Boolean>> requestAccessEnabled() {
        return Codegen.optional(this.requestAccessEnabled);
    }
    /**
     * Defaults to false. Require all users in this group to setup Two-factor authentication.
     * 
     */
    @Export(name="requireTwoFactorAuthentication", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requireTwoFactorAuthentication;

    /**
     * @return Defaults to false. Require all users in this group to setup Two-factor authentication.
     * 
     */
    public Output<Optional<Boolean>> requireTwoFactorAuthentication() {
        return Codegen.optional(this.requireTwoFactorAuthentication);
    }
    /**
     * The group level registration token to use during runner setup.
     * 
     */
    @Export(name="runnersToken", type=String.class, parameters={})
    private Output<String> runnersToken;

    /**
     * @return The group level registration token to use during runner setup.
     * 
     */
    public Output<String> runnersToken() {
        return this.runnersToken;
    }
    /**
     * Defaults to false. Prevent sharing a project with another group within this group.
     * 
     */
    @Export(name="shareWithGroupLock", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> shareWithGroupLock;

    /**
     * @return Defaults to false. Prevent sharing a project with another group within this group.
     * 
     */
    public Output<Optional<Boolean>> shareWithGroupLock() {
        return Codegen.optional(this.shareWithGroupLock);
    }
    /**
     * Defaults to owner. Allowed to create subgroups.
     * 
     */
    @Export(name="subgroupCreationLevel", type=String.class, parameters={})
    private Output</* @Nullable */ String> subgroupCreationLevel;

    /**
     * @return Defaults to owner. Allowed to create subgroups.
     * 
     */
    public Output<Optional<String>> subgroupCreationLevel() {
        return Codegen.optional(this.subgroupCreationLevel);
    }
    /**
     * Defaults to 48. Time before Two-factor authentication is enforced (in hours).
     * 
     */
    @Export(name="twoFactorGracePeriod", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> twoFactorGracePeriod;

    /**
     * @return Defaults to 48. Time before Two-factor authentication is enforced (in hours).
     * 
     */
    public Output<Optional<Integer>> twoFactorGracePeriod() {
        return Codegen.optional(this.twoFactorGracePeriod);
    }
    /**
     * The group&#39;s visibility. Can be `private`, `internal`, or `public`.
     * 
     */
    @Export(name="visibilityLevel", type=String.class, parameters={})
    private Output<String> visibilityLevel;

    /**
     * @return The group&#39;s visibility. Can be `private`, `internal`, or `public`.
     * 
     */
    public Output<String> visibilityLevel() {
        return this.visibilityLevel;
    }
    /**
     * Web URL of the group.
     * 
     */
    @Export(name="webUrl", type=String.class, parameters={})
    private Output<String> webUrl;

    /**
     * @return Web URL of the group.
     * 
     */
    public Output<String> webUrl() {
        return this.webUrl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Group(String name) {
        this(name, GroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Group(String name, GroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Group(String name, GroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/group:Group", name, args == null ? GroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Group(String name, Output<String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/group:Group", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "runnersToken"
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
    public static Group get(String name, Output<String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Group(name, id, state, options);
    }
}