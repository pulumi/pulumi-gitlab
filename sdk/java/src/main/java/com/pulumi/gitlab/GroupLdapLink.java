// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupLdapLinkArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupLdapLinkState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupLdapLink` resource allows to manage the lifecycle of an LDAP integration with a group.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#ldap-group-links)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitLab group ldap links can be imported using an id made up of `group_id:ldap_provider:cn:filter`. CN and Filter are mutually exclusive, so one will be missing.
 * 
 * If using the CN for the group link, the ID will end with a blank filter (&#34;:&#34;). e.g.,
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test &#34;12345:ldapmain:testcn:&#34;
 * ```
 * 
 * If using the Filter for the group link, the ID will have two &#34;::&#34; in the middle due to having a blank CN. e.g.,
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test &#34;12345:ldapmain::testfilter&#34;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/groupLdapLink:GroupLdapLink")
public class GroupLdapLink extends com.pulumi.resources.CustomResource {
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     * 
     * @deprecated
     * Use `group_access` instead of the `access_level` attribute.
     * 
     */
    @Deprecated /* Use `group_access` instead of the `access_level` attribute. */
    @Export(name="accessLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessLevel;

    /**
     * @return Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     * 
     */
    public Output<Optional<String>> accessLevel() {
        return Codegen.optional(this.accessLevel);
    }
    /**
     * The CN of the LDAP group to link with. Required if `filter` is not provided.
     * 
     */
    @Export(name="cn", refs={String.class}, tree="[0]")
    private Output<String> cn;

    /**
     * @return The CN of the LDAP group to link with. Required if `filter` is not provided.
     * 
     */
    public Output<String> cn() {
        return this.cn;
    }
    /**
     * The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
     * 
     */
    @Export(name="filter", refs={String.class}, tree="[0]")
    private Output<String> filter;

    /**
     * @return The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
     * 
     */
    public Output<String> filter() {
        return this.filter;
    }
    /**
     * If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
     * 
     */
    @Export(name="force", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> force;

    /**
     * @return If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
     * 
     */
    public Output<Optional<Boolean>> force() {
        return Codegen.optional(this.force);
    }
    /**
     * The ID or URL-encoded path of the group
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The ID or URL-encoded path of the group
     * 
     */
    public Output<String> group() {
        return this.group;
    }
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     * 
     */
    @Export(name="groupAccess", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupAccess;

    /**
     * @return Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     * 
     */
    public Output<Optional<String>> groupAccess() {
        return Codegen.optional(this.groupAccess);
    }
    /**
     * The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
     * 
     */
    @Export(name="ldapProvider", refs={String.class}, tree="[0]")
    private Output<String> ldapProvider;

    /**
     * @return The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
     * 
     */
    public Output<String> ldapProvider() {
        return this.ldapProvider;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupLdapLink(String name) {
        this(name, GroupLdapLinkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupLdapLink(String name, GroupLdapLinkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupLdapLink(String name, GroupLdapLinkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupLdapLink:GroupLdapLink", name, args == null ? GroupLdapLinkArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupLdapLink(String name, Output<String> id, @Nullable GroupLdapLinkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupLdapLink:GroupLdapLink", name, state, makeResourceOptions(options, id));
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
    public static GroupLdapLink get(String name, Output<String> id, @Nullable GroupLdapLinkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupLdapLink(name, id, state, options);
    }
}
