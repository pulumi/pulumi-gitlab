// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupLdapLinkArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupLdapLinkArgs Empty = new GroupLdapLinkArgs();

    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     * @deprecated
     * Use `group_access` instead of the `access_level` attribute.
     * 
     */
    @Deprecated /* Use `group_access` instead of the `access_level` attribute. */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     * @deprecated
     * Use `group_access` instead of the `access_level` attribute.
     * 
     */
    @Deprecated /* Use `group_access` instead of the `access_level` attribute. */
    public Optional<Output<String>> accessLevel() {
        return Optional.ofNullable(this.accessLevel);
    }

    /**
     * The CN of the LDAP group to link with. Required if `filter` is not provided.
     * 
     */
    @Import(name="cn")
    private @Nullable Output<String> cn;

    /**
     * @return The CN of the LDAP group to link with. Required if `filter` is not provided.
     * 
     */
    public Optional<Output<String>> cn() {
        return Optional.ofNullable(this.cn);
    }

    /**
     * The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
     * 
     */
    @Import(name="filter")
    private @Nullable Output<String> filter;

    /**
     * @return The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
     * 
     */
    public Optional<Output<String>> filter() {
        return Optional.ofNullable(this.filter);
    }

    /**
     * If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
     * 
     */
    @Import(name="force")
    private @Nullable Output<Boolean> force;

    /**
     * @return If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
     * 
     */
    public Optional<Output<Boolean>> force() {
        return Optional.ofNullable(this.force);
    }

    /**
     * The ID or URL-encoded path of the group
     * 
     */
    @Import(name="group", required=true)
    private Output<String> group;

    /**
     * @return The ID or URL-encoded path of the group
     * 
     */
    public Output<String> group() {
        return this.group;
    }

    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     */
    @Import(name="groupAccess")
    private @Nullable Output<String> groupAccess;

    /**
     * @return Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     */
    public Optional<Output<String>> groupAccess() {
        return Optional.ofNullable(this.groupAccess);
    }

    /**
     * The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/administration/raketasks/ldap/#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
     * 
     */
    @Import(name="ldapProvider", required=true)
    private Output<String> ldapProvider;

    /**
     * @return The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/administration/raketasks/ldap/#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
     * 
     */
    public Output<String> ldapProvider() {
        return this.ldapProvider;
    }

    /**
     * The ID of a custom member role. Only available for Ultimate instances. When using a custom role, the `group_access` must match the base role used to create the custom role.
     * 
     */
    @Import(name="memberRoleId")
    private @Nullable Output<Integer> memberRoleId;

    /**
     * @return The ID of a custom member role. Only available for Ultimate instances. When using a custom role, the `group_access` must match the base role used to create the custom role.
     * 
     */
    public Optional<Output<Integer>> memberRoleId() {
        return Optional.ofNullable(this.memberRoleId);
    }

    private GroupLdapLinkArgs() {}

    private GroupLdapLinkArgs(GroupLdapLinkArgs $) {
        this.accessLevel = $.accessLevel;
        this.cn = $.cn;
        this.filter = $.filter;
        this.force = $.force;
        this.group = $.group;
        this.groupAccess = $.groupAccess;
        this.ldapProvider = $.ldapProvider;
        this.memberRoleId = $.memberRoleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupLdapLinkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupLdapLinkArgs $;

        public Builder() {
            $ = new GroupLdapLinkArgs();
        }

        public Builder(GroupLdapLinkArgs defaults) {
            $ = new GroupLdapLinkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
         * 
         * @return builder
         * 
         * @deprecated
         * Use `group_access` instead of the `access_level` attribute.
         * 
         */
        @Deprecated /* Use `group_access` instead of the `access_level` attribute. */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
         * 
         * @return builder
         * 
         * @deprecated
         * Use `group_access` instead of the `access_level` attribute.
         * 
         */
        @Deprecated /* Use `group_access` instead of the `access_level` attribute. */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
        }

        /**
         * @param cn The CN of the LDAP group to link with. Required if `filter` is not provided.
         * 
         * @return builder
         * 
         */
        public Builder cn(@Nullable Output<String> cn) {
            $.cn = cn;
            return this;
        }

        /**
         * @param cn The CN of the LDAP group to link with. Required if `filter` is not provided.
         * 
         * @return builder
         * 
         */
        public Builder cn(String cn) {
            return cn(Output.of(cn));
        }

        /**
         * @param filter The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
         * 
         * @return builder
         * 
         */
        public Builder filter(@Nullable Output<String> filter) {
            $.filter = filter;
            return this;
        }

        /**
         * @param filter The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
         * 
         * @return builder
         * 
         */
        public Builder filter(String filter) {
            return filter(Output.of(filter));
        }

        /**
         * @param force If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
         * 
         * @return builder
         * 
         */
        public Builder force(@Nullable Output<Boolean> force) {
            $.force = force;
            return this;
        }

        /**
         * @param force If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
         * 
         * @return builder
         * 
         */
        public Builder force(Boolean force) {
            return force(Output.of(force));
        }

        /**
         * @param group The ID or URL-encoded path of the group
         * 
         * @return builder
         * 
         */
        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The ID or URL-encoded path of the group
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param groupAccess Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
         * 
         * @return builder
         * 
         */
        public Builder groupAccess(@Nullable Output<String> groupAccess) {
            $.groupAccess = groupAccess;
            return this;
        }

        /**
         * @param groupAccess Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
         * 
         * @return builder
         * 
         */
        public Builder groupAccess(String groupAccess) {
            return groupAccess(Output.of(groupAccess));
        }

        /**
         * @param ldapProvider The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/administration/raketasks/ldap/#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
         * 
         * @return builder
         * 
         */
        public Builder ldapProvider(Output<String> ldapProvider) {
            $.ldapProvider = ldapProvider;
            return this;
        }

        /**
         * @param ldapProvider The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/administration/raketasks/ldap/#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
         * 
         * @return builder
         * 
         */
        public Builder ldapProvider(String ldapProvider) {
            return ldapProvider(Output.of(ldapProvider));
        }

        /**
         * @param memberRoleId The ID of a custom member role. Only available for Ultimate instances. When using a custom role, the `group_access` must match the base role used to create the custom role.
         * 
         * @return builder
         * 
         */
        public Builder memberRoleId(@Nullable Output<Integer> memberRoleId) {
            $.memberRoleId = memberRoleId;
            return this;
        }

        /**
         * @param memberRoleId The ID of a custom member role. Only available for Ultimate instances. When using a custom role, the `group_access` must match the base role used to create the custom role.
         * 
         * @return builder
         * 
         */
        public Builder memberRoleId(Integer memberRoleId) {
            return memberRoleId(Output.of(memberRoleId));
        }

        public GroupLdapLinkArgs build() {
            if ($.group == null) {
                throw new MissingRequiredPropertyException("GroupLdapLinkArgs", "group");
            }
            if ($.ldapProvider == null) {
                throw new MissingRequiredPropertyException("GroupLdapLinkArgs", "ldapProvider");
            }
            return $;
        }
    }

}
