# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GroupLdapLinkArgs', 'GroupLdapLink']

@pulumi.input_type
class GroupLdapLinkArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 ldap_provider: pulumi.Input[str],
                 access_level: Optional[pulumi.Input[str]] = None,
                 cn: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 group_access: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GroupLdapLink resource.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group
        :param pulumi.Input[str] ldap_provider: The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
        :param pulumi.Input[str] access_level: Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] cn: The CN of the LDAP group to link with. Required if `filter` is not provided.
        :param pulumi.Input[str] filter: The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
        :param pulumi.Input[bool] force: If true, then delete and replace an existing LDAP link if one exists.
        :param pulumi.Input[str] group_access: Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        """
        GroupLdapLinkArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            group=group,
            ldap_provider=ldap_provider,
            access_level=access_level,
            cn=cn,
            filter=filter,
            force=force,
            group_access=group_access,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             group: Optional[pulumi.Input[str]] = None,
             ldap_provider: Optional[pulumi.Input[str]] = None,
             access_level: Optional[pulumi.Input[str]] = None,
             cn: Optional[pulumi.Input[str]] = None,
             filter: Optional[pulumi.Input[str]] = None,
             force: Optional[pulumi.Input[bool]] = None,
             group_access: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if group is None:
            raise TypeError("Missing 'group' argument")
        if ldap_provider is None and 'ldapProvider' in kwargs:
            ldap_provider = kwargs['ldapProvider']
        if ldap_provider is None:
            raise TypeError("Missing 'ldap_provider' argument")
        if access_level is None and 'accessLevel' in kwargs:
            access_level = kwargs['accessLevel']
        if group_access is None and 'groupAccess' in kwargs:
            group_access = kwargs['groupAccess']

        _setter("group", group)
        _setter("ldap_provider", ldap_provider)
        if access_level is not None:
            warnings.warn("""Use `group_access` instead of the `access_level` attribute.""", DeprecationWarning)
            pulumi.log.warn("""access_level is deprecated: Use `group_access` instead of the `access_level` attribute.""")
        if access_level is not None:
            _setter("access_level", access_level)
        if cn is not None:
            _setter("cn", cn)
        if filter is not None:
            _setter("filter", filter)
        if force is not None:
            _setter("force", force)
        if group_access is not None:
            _setter("group_access", group_access)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        The ID or URL-encoded path of the group
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="ldapProvider")
    def ldap_provider(self) -> pulumi.Input[str]:
        """
        The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
        """
        return pulumi.get(self, "ldap_provider")

    @ldap_provider.setter
    def ldap_provider(self, value: pulumi.Input[str]):
        pulumi.set(self, "ldap_provider", value)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        """
        warnings.warn("""Use `group_access` instead of the `access_level` attribute.""", DeprecationWarning)
        pulumi.log.warn("""access_level is deprecated: Use `group_access` instead of the `access_level` attribute.""")

        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter
    def cn(self) -> Optional[pulumi.Input[str]]:
        """
        The CN of the LDAP group to link with. Required if `filter` is not provided.
        """
        return pulumi.get(self, "cn")

    @cn.setter
    def cn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cn", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def force(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, then delete and replace an existing LDAP link if one exists.
        """
        return pulumi.get(self, "force")

    @force.setter
    def force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force", value)

    @property
    @pulumi.getter(name="groupAccess")
    def group_access(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        """
        return pulumi.get(self, "group_access")

    @group_access.setter
    def group_access(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_access", value)


@pulumi.input_type
class _GroupLdapLinkState:
    def __init__(__self__, *,
                 access_level: Optional[pulumi.Input[str]] = None,
                 cn: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 group_access: Optional[pulumi.Input[str]] = None,
                 ldap_provider: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupLdapLink resources.
        :param pulumi.Input[str] access_level: Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] cn: The CN of the LDAP group to link with. Required if `filter` is not provided.
        :param pulumi.Input[str] filter: The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
        :param pulumi.Input[bool] force: If true, then delete and replace an existing LDAP link if one exists.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group
        :param pulumi.Input[str] group_access: Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] ldap_provider: The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
        """
        _GroupLdapLinkState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            access_level=access_level,
            cn=cn,
            filter=filter,
            force=force,
            group=group,
            group_access=group_access,
            ldap_provider=ldap_provider,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             access_level: Optional[pulumi.Input[str]] = None,
             cn: Optional[pulumi.Input[str]] = None,
             filter: Optional[pulumi.Input[str]] = None,
             force: Optional[pulumi.Input[bool]] = None,
             group: Optional[pulumi.Input[str]] = None,
             group_access: Optional[pulumi.Input[str]] = None,
             ldap_provider: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if access_level is None and 'accessLevel' in kwargs:
            access_level = kwargs['accessLevel']
        if group_access is None and 'groupAccess' in kwargs:
            group_access = kwargs['groupAccess']
        if ldap_provider is None and 'ldapProvider' in kwargs:
            ldap_provider = kwargs['ldapProvider']

        if access_level is not None:
            warnings.warn("""Use `group_access` instead of the `access_level` attribute.""", DeprecationWarning)
            pulumi.log.warn("""access_level is deprecated: Use `group_access` instead of the `access_level` attribute.""")
        if access_level is not None:
            _setter("access_level", access_level)
        if cn is not None:
            _setter("cn", cn)
        if filter is not None:
            _setter("filter", filter)
        if force is not None:
            _setter("force", force)
        if group is not None:
            _setter("group", group)
        if group_access is not None:
            _setter("group_access", group_access)
        if ldap_provider is not None:
            _setter("ldap_provider", ldap_provider)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        """
        warnings.warn("""Use `group_access` instead of the `access_level` attribute.""", DeprecationWarning)
        pulumi.log.warn("""access_level is deprecated: Use `group_access` instead of the `access_level` attribute.""")

        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter
    def cn(self) -> Optional[pulumi.Input[str]]:
        """
        The CN of the LDAP group to link with. Required if `filter` is not provided.
        """
        return pulumi.get(self, "cn")

    @cn.setter
    def cn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cn", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def force(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, then delete and replace an existing LDAP link if one exists.
        """
        return pulumi.get(self, "force")

    @force.setter
    def force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or URL-encoded path of the group
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="groupAccess")
    def group_access(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        """
        return pulumi.get(self, "group_access")

    @group_access.setter
    def group_access(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_access", value)

    @property
    @pulumi.getter(name="ldapProvider")
    def ldap_provider(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
        """
        return pulumi.get(self, "ldap_provider")

    @ldap_provider.setter
    def ldap_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_provider", value)


class GroupLdapLink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 cn: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 group_access: Optional[pulumi.Input[str]] = None,
                 ldap_provider: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `GroupLdapLink` resource allows to manage the lifecycle of an LDAP integration with a group.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#ldap-group-links)

        ## Import

        GitLab group ldap links can be imported using an id made up of `group_id:ldap_provider:cn:filter`. CN and Filter are mutually exclusive, so one will be missing. If using the CN for the group link, the ID will end with a blank filter (":"). e.g.,

        ```sh
         $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain:testcn:"
        ```

         If using the Filter for the group link, the ID will have two "::" in the middle due to having a blank CN. e.g.,

        ```sh
         $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain::testfilter"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] cn: The CN of the LDAP group to link with. Required if `filter` is not provided.
        :param pulumi.Input[str] filter: The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
        :param pulumi.Input[bool] force: If true, then delete and replace an existing LDAP link if one exists.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group
        :param pulumi.Input[str] group_access: Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] ldap_provider: The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupLdapLinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupLdapLink` resource allows to manage the lifecycle of an LDAP integration with a group.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#ldap-group-links)

        ## Import

        GitLab group ldap links can be imported using an id made up of `group_id:ldap_provider:cn:filter`. CN and Filter are mutually exclusive, so one will be missing. If using the CN for the group link, the ID will end with a blank filter (":"). e.g.,

        ```sh
         $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain:testcn:"
        ```

         If using the Filter for the group link, the ID will have two "::" in the middle due to having a blank CN. e.g.,

        ```sh
         $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain::testfilter"
        ```

        :param str resource_name: The name of the resource.
        :param GroupLdapLinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupLdapLinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            GroupLdapLinkArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 cn: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 group_access: Optional[pulumi.Input[str]] = None,
                 ldap_provider: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupLdapLinkArgs.__new__(GroupLdapLinkArgs)

            __props__.__dict__["access_level"] = access_level
            __props__.__dict__["cn"] = cn
            __props__.__dict__["filter"] = filter
            __props__.__dict__["force"] = force
            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["group_access"] = group_access
            if ldap_provider is None and not opts.urn:
                raise TypeError("Missing required property 'ldap_provider'")
            __props__.__dict__["ldap_provider"] = ldap_provider
        super(GroupLdapLink, __self__).__init__(
            'gitlab:index/groupLdapLink:GroupLdapLink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_level: Optional[pulumi.Input[str]] = None,
            cn: Optional[pulumi.Input[str]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            force: Optional[pulumi.Input[bool]] = None,
            group: Optional[pulumi.Input[str]] = None,
            group_access: Optional[pulumi.Input[str]] = None,
            ldap_provider: Optional[pulumi.Input[str]] = None) -> 'GroupLdapLink':
        """
        Get an existing GroupLdapLink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] cn: The CN of the LDAP group to link with. Required if `filter` is not provided.
        :param pulumi.Input[str] filter: The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
        :param pulumi.Input[bool] force: If true, then delete and replace an existing LDAP link if one exists.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group
        :param pulumi.Input[str] group_access: Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] ldap_provider: The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupLdapLinkState.__new__(_GroupLdapLinkState)

        __props__.__dict__["access_level"] = access_level
        __props__.__dict__["cn"] = cn
        __props__.__dict__["filter"] = filter
        __props__.__dict__["force"] = force
        __props__.__dict__["group"] = group
        __props__.__dict__["group_access"] = group_access
        __props__.__dict__["ldap_provider"] = ldap_provider
        return GroupLdapLink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> pulumi.Output[Optional[str]]:
        """
        Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        """
        warnings.warn("""Use `group_access` instead of the `access_level` attribute.""", DeprecationWarning)
        pulumi.log.warn("""access_level is deprecated: Use `group_access` instead of the `access_level` attribute.""")

        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter
    def cn(self) -> pulumi.Output[str]:
        """
        The CN of the LDAP group to link with. Required if `filter` is not provided.
        """
        return pulumi.get(self, "cn")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[str]:
        """
        The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def force(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, then delete and replace an existing LDAP link if one exists.
        """
        return pulumi.get(self, "force")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of the group
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="groupAccess")
    def group_access(self) -> pulumi.Output[Optional[str]]:
        """
        Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        """
        return pulumi.get(self, "group_access")

    @property
    @pulumi.getter(name="ldapProvider")
    def ldap_provider(self) -> pulumi.Output[str]:
        """
        The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
        """
        return pulumi.get(self, "ldap_provider")

