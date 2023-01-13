// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupLdapLinkArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupLdapLinkArgs Empty = new GroupLdapLinkArgs();

    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     * 
     * @deprecated
     * Use `group_access` instead of the `access_level` attribute.
     * 
     */
    @Deprecated /* Use `group_access` instead of the `access_level` attribute. */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
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
     * The CN of the LDAP group to link with.
     * 
     */
    @Import(name="cn", required=true)
    private Output<String> cn;

    /**
     * @return The CN of the LDAP group to link with.
     * 
     */
    public Output<String> cn() {
        return this.cn;
    }

    /**
     * If true, then delete and replace an existing LDAP link if one exists.
     * 
     */
    @Import(name="force")
    private @Nullable Output<Boolean> force;

    /**
     * @return If true, then delete and replace an existing LDAP link if one exists.
     * 
     */
    public Optional<Output<Boolean>> force() {
        return Optional.ofNullable(this.force);
    }

    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     * 
     */
    @Import(name="groupAccess")
    private @Nullable Output<String> groupAccess;

    /**
     * @return Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     * 
     */
    public Optional<Output<String>> groupAccess() {
        return Optional.ofNullable(this.groupAccess);
    }

    /**
     * The id of the GitLab group.
     * 
     */
    @Import(name="groupId", required=true)
    private Output<String> groupId;

    /**
     * @return The id of the GitLab group.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }

    /**
     * The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
     * 
     */
    @Import(name="ldapProvider", required=true)
    private Output<String> ldapProvider;

    /**
     * @return The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
     * 
     */
    public Output<String> ldapProvider() {
        return this.ldapProvider;
    }

    private GroupLdapLinkArgs() {}

    private GroupLdapLinkArgs(GroupLdapLinkArgs $) {
        this.accessLevel = $.accessLevel;
        this.cn = $.cn;
        this.force = $.force;
        this.groupAccess = $.groupAccess;
        this.groupId = $.groupId;
        this.ldapProvider = $.ldapProvider;
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
         * @param accessLevel Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
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
         * @param accessLevel Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
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
         * @param cn The CN of the LDAP group to link with.
         * 
         * @return builder
         * 
         */
        public Builder cn(Output<String> cn) {
            $.cn = cn;
            return this;
        }

        /**
         * @param cn The CN of the LDAP group to link with.
         * 
         * @return builder
         * 
         */
        public Builder cn(String cn) {
            return cn(Output.of(cn));
        }

        /**
         * @param force If true, then delete and replace an existing LDAP link if one exists.
         * 
         * @return builder
         * 
         */
        public Builder force(@Nullable Output<Boolean> force) {
            $.force = force;
            return this;
        }

        /**
         * @param force If true, then delete and replace an existing LDAP link if one exists.
         * 
         * @return builder
         * 
         */
        public Builder force(Boolean force) {
            return force(Output.of(force));
        }

        /**
         * @param groupAccess Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
         * 
         * @return builder
         * 
         */
        public Builder groupAccess(@Nullable Output<String> groupAccess) {
            $.groupAccess = groupAccess;
            return this;
        }

        /**
         * @param groupAccess Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
         * 
         * @return builder
         * 
         */
        public Builder groupAccess(String groupAccess) {
            return groupAccess(Output.of(groupAccess));
        }

        /**
         * @param groupId The id of the GitLab group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The id of the GitLab group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param ldapProvider The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
         * 
         * @return builder
         * 
         */
        public Builder ldapProvider(Output<String> ldapProvider) {
            $.ldapProvider = ldapProvider;
            return this;
        }

        /**
         * @param ldapProvider The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
         * 
         * @return builder
         * 
         */
        public Builder ldapProvider(String ldapProvider) {
            return ldapProvider(Output.of(ldapProvider));
        }

        public GroupLdapLinkArgs build() {
            $.cn = Objects.requireNonNull($.cn, "expected parameter 'cn' to be non-null");
            $.groupId = Objects.requireNonNull($.groupId, "expected parameter 'groupId' to be non-null");
            $.ldapProvider = Objects.requireNonNull($.ldapProvider, "expected parameter 'ldapProvider' to be non-null");
            return $;
        }
    }

}