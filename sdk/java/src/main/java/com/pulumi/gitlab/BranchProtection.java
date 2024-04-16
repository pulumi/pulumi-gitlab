// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.BranchProtectionArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.BranchProtectionState;
import com.pulumi.gitlab.outputs.BranchProtectionAllowedToMerge;
import com.pulumi.gitlab.outputs.BranchProtectionAllowedToPush;
import com.pulumi.gitlab.outputs.BranchProtectionAllowedToUnprotect;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Gitlab protected branches can be imported with a key composed of `&lt;project_id&gt;:&lt;branch&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/branchProtection:BranchProtection BranchProtect &#34;12345:main&#34;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/branchProtection:BranchProtection")
public class BranchProtection extends com.pulumi.resources.CustomResource {
    /**
     * Can be set to true to allow users with push access to force push.
     * 
     */
    @Export(name="allowForcePush", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowForcePush;

    /**
     * @return Can be set to true to allow users with push access to force push.
     * 
     */
    public Output<Boolean> allowForcePush() {
        return this.allowForcePush;
    }
    /**
     * Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     * 
     */
    @Export(name="allowedToMerges", refs={List.class,BranchProtectionAllowedToMerge.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BranchProtectionAllowedToMerge>> allowedToMerges;

    /**
     * @return Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     * 
     */
    public Output<Optional<List<BranchProtectionAllowedToMerge>>> allowedToMerges() {
        return Codegen.optional(this.allowedToMerges);
    }
    /**
     * Array of access levels and user(s)/group(s) allowed to push to protected branch.
     * 
     */
    @Export(name="allowedToPushes", refs={List.class,BranchProtectionAllowedToPush.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BranchProtectionAllowedToPush>> allowedToPushes;

    /**
     * @return Array of access levels and user(s)/group(s) allowed to push to protected branch.
     * 
     */
    public Output<Optional<List<BranchProtectionAllowedToPush>>> allowedToPushes() {
        return Codegen.optional(this.allowedToPushes);
    }
    /**
     * Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
     * 
     */
    @Export(name="allowedToUnprotects", refs={List.class,BranchProtectionAllowedToUnprotect.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BranchProtectionAllowedToUnprotect>> allowedToUnprotects;

    /**
     * @return Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
     * 
     */
    public Output<Optional<List<BranchProtectionAllowedToUnprotect>>> allowedToUnprotects() {
        return Codegen.optional(this.allowedToUnprotects);
    }
    /**
     * Name of the branch.
     * 
     */
    @Export(name="branch", refs={String.class}, tree="[0]")
    private Output<String> branch;

    /**
     * @return Name of the branch.
     * 
     */
    public Output<String> branch() {
        return this.branch;
    }
    /**
     * The ID of the branch protection (not the branch name).
     * 
     */
    @Export(name="branchProtectionId", refs={Integer.class}, tree="[0]")
    private Output<Integer> branchProtectionId;

    /**
     * @return The ID of the branch protection (not the branch name).
     * 
     */
    public Output<Integer> branchProtectionId() {
        return this.branchProtectionId;
    }
    /**
     * Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
     * 
     */
    @Export(name="codeOwnerApprovalRequired", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> codeOwnerApprovalRequired;

    /**
     * @return Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
     * 
     */
    public Output<Boolean> codeOwnerApprovalRequired() {
        return this.codeOwnerApprovalRequired;
    }
    /**
     * Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
     * 
     */
    @Export(name="mergeAccessLevel", refs={String.class}, tree="[0]")
    private Output<String> mergeAccessLevel;

    /**
     * @return Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
     * 
     */
    public Output<String> mergeAccessLevel() {
        return this.mergeAccessLevel;
    }
    /**
     * The id of the project.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The id of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
     * 
     */
    @Export(name="pushAccessLevel", refs={String.class}, tree="[0]")
    private Output<String> pushAccessLevel;

    /**
     * @return Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
     * 
     */
    public Output<String> pushAccessLevel() {
        return this.pushAccessLevel;
    }
    /**
     * Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
     * 
     */
    @Export(name="unprotectAccessLevel", refs={String.class}, tree="[0]")
    private Output<String> unprotectAccessLevel;

    /**
     * @return Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
     * 
     */
    public Output<String> unprotectAccessLevel() {
        return this.unprotectAccessLevel;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BranchProtection(String name) {
        this(name, BranchProtectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BranchProtection(String name, BranchProtectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BranchProtection(String name, BranchProtectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/branchProtection:BranchProtection", name, args == null ? BranchProtectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BranchProtection(String name, Output<String> id, @Nullable BranchProtectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/branchProtection:BranchProtection", name, state, makeResourceOptions(options, id));
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
    public static BranchProtection get(String name, Output<String> id, @Nullable BranchProtectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BranchProtection(name, id, state, options);
    }
}
