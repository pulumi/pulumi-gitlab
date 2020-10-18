# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['Group']


class Group(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_devops_enabled: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 emails_disabled: Optional[pulumi.Input[bool]] = None,
                 lfs_enabled: Optional[pulumi.Input[bool]] = None,
                 mentions_disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[int]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 project_creation_level: Optional[pulumi.Input[str]] = None,
                 request_access_enabled: Optional[pulumi.Input[bool]] = None,
                 require_two_factor_authentication: Optional[pulumi.Input[bool]] = None,
                 share_with_group_lock: Optional[pulumi.Input[bool]] = None,
                 subgroup_creation_level: Optional[pulumi.Input[str]] = None,
                 two_factor_grace_period: Optional[pulumi.Input[int]] = None,
                 visibility_level: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Group resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_devops_enabled: Boolean, defaults to false.  Default to Auto
               DevOps pipeline for all projects within this group.
        :param pulumi.Input[str] description: The description of the group.
        :param pulumi.Input[bool] emails_disabled: Boolean, defaults to false.  Disable email notifications
        :param pulumi.Input[bool] lfs_enabled: Boolean, defaults to true.  Whether to enable LFS
               support for projects in this group.
        :param pulumi.Input[bool] mentions_disabled: Boolean, defaults to false.  Disable the capability
               of a group from getting mentioned
        :param pulumi.Input[str] name: The name of this group.
        :param pulumi.Input[int] parent_id: Integer, id of the parent group (creates a nested group).
        :param pulumi.Input[str] path: The path of the group.
        :param pulumi.Input[str] project_creation_level: , defaults to Maintainer.
               Determine if developers can create projects
               in the group. Can be noone (No one), maintainer (Maintainers),
               or developer (Developers + Maintainers).
        :param pulumi.Input[bool] request_access_enabled: Boolean, defaults to false.  Whether to
               enable users to request access to the group.
        :param pulumi.Input[bool] require_two_factor_authentication: Boolean, defaults to false.
               equire all users in this group to setup Two-factor authentication.
        :param pulumi.Input[bool] share_with_group_lock: Boolean, defaults to false.  Prevent sharing
               a project with another group within this group.
        :param pulumi.Input[str] subgroup_creation_level: , defaults to Owner.
               Allowed to create subgroups.
               Can be owner (Owners), or maintainer (Maintainers).
        :param pulumi.Input[int] two_factor_grace_period: Int, defaults to 48.
               Time before Two-factor authentication is enforced (in hours).
        :param pulumi.Input[str] visibility_level: The group's visibility. Can be `private`, `internal`, or `public`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['auto_devops_enabled'] = auto_devops_enabled
            __props__['description'] = description
            __props__['emails_disabled'] = emails_disabled
            __props__['lfs_enabled'] = lfs_enabled
            __props__['mentions_disabled'] = mentions_disabled
            __props__['name'] = name
            __props__['parent_id'] = parent_id
            if path is None:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
            __props__['project_creation_level'] = project_creation_level
            __props__['request_access_enabled'] = request_access_enabled
            __props__['require_two_factor_authentication'] = require_two_factor_authentication
            __props__['share_with_group_lock'] = share_with_group_lock
            __props__['subgroup_creation_level'] = subgroup_creation_level
            __props__['two_factor_grace_period'] = two_factor_grace_period
            __props__['visibility_level'] = visibility_level
            __props__['full_name'] = None
            __props__['full_path'] = None
            __props__['runners_token'] = None
            __props__['web_url'] = None
        super(Group, __self__).__init__(
            'gitlab:index/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_devops_enabled: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            emails_disabled: Optional[pulumi.Input[bool]] = None,
            full_name: Optional[pulumi.Input[str]] = None,
            full_path: Optional[pulumi.Input[str]] = None,
            lfs_enabled: Optional[pulumi.Input[bool]] = None,
            mentions_disabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_id: Optional[pulumi.Input[int]] = None,
            path: Optional[pulumi.Input[str]] = None,
            project_creation_level: Optional[pulumi.Input[str]] = None,
            request_access_enabled: Optional[pulumi.Input[bool]] = None,
            require_two_factor_authentication: Optional[pulumi.Input[bool]] = None,
            runners_token: Optional[pulumi.Input[str]] = None,
            share_with_group_lock: Optional[pulumi.Input[bool]] = None,
            subgroup_creation_level: Optional[pulumi.Input[str]] = None,
            two_factor_grace_period: Optional[pulumi.Input[int]] = None,
            visibility_level: Optional[pulumi.Input[str]] = None,
            web_url: Optional[pulumi.Input[str]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_devops_enabled: Boolean, defaults to false.  Default to Auto
               DevOps pipeline for all projects within this group.
        :param pulumi.Input[str] description: The description of the group.
        :param pulumi.Input[bool] emails_disabled: Boolean, defaults to false.  Disable email notifications
        :param pulumi.Input[str] full_name: The full name of the group.
        :param pulumi.Input[str] full_path: The full path of the group.
        :param pulumi.Input[bool] lfs_enabled: Boolean, defaults to true.  Whether to enable LFS
               support for projects in this group.
        :param pulumi.Input[bool] mentions_disabled: Boolean, defaults to false.  Disable the capability
               of a group from getting mentioned
        :param pulumi.Input[str] name: The name of this group.
        :param pulumi.Input[int] parent_id: Integer, id of the parent group (creates a nested group).
        :param pulumi.Input[str] path: The path of the group.
        :param pulumi.Input[str] project_creation_level: , defaults to Maintainer.
               Determine if developers can create projects
               in the group. Can be noone (No one), maintainer (Maintainers),
               or developer (Developers + Maintainers).
        :param pulumi.Input[bool] request_access_enabled: Boolean, defaults to false.  Whether to
               enable users to request access to the group.
        :param pulumi.Input[bool] require_two_factor_authentication: Boolean, defaults to false.
               equire all users in this group to setup Two-factor authentication.
        :param pulumi.Input[str] runners_token: The group level registration token to use during runner setup.
        :param pulumi.Input[bool] share_with_group_lock: Boolean, defaults to false.  Prevent sharing
               a project with another group within this group.
        :param pulumi.Input[str] subgroup_creation_level: , defaults to Owner.
               Allowed to create subgroups.
               Can be owner (Owners), or maintainer (Maintainers).
        :param pulumi.Input[int] two_factor_grace_period: Int, defaults to 48.
               Time before Two-factor authentication is enforced (in hours).
        :param pulumi.Input[str] visibility_level: The group's visibility. Can be `private`, `internal`, or `public`.
        :param pulumi.Input[str] web_url: Web URL of the group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_devops_enabled"] = auto_devops_enabled
        __props__["description"] = description
        __props__["emails_disabled"] = emails_disabled
        __props__["full_name"] = full_name
        __props__["full_path"] = full_path
        __props__["lfs_enabled"] = lfs_enabled
        __props__["mentions_disabled"] = mentions_disabled
        __props__["name"] = name
        __props__["parent_id"] = parent_id
        __props__["path"] = path
        __props__["project_creation_level"] = project_creation_level
        __props__["request_access_enabled"] = request_access_enabled
        __props__["require_two_factor_authentication"] = require_two_factor_authentication
        __props__["runners_token"] = runners_token
        __props__["share_with_group_lock"] = share_with_group_lock
        __props__["subgroup_creation_level"] = subgroup_creation_level
        __props__["two_factor_grace_period"] = two_factor_grace_period
        __props__["visibility_level"] = visibility_level
        __props__["web_url"] = web_url
        return Group(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoDevopsEnabled")
    def auto_devops_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, defaults to false.  Default to Auto
        DevOps pipeline for all projects within this group.
        """
        return pulumi.get(self, "auto_devops_enabled")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="emailsDisabled")
    def emails_disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, defaults to false.  Disable email notifications
        """
        return pulumi.get(self, "emails_disabled")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> pulumi.Output[str]:
        """
        The full name of the group.
        """
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> pulumi.Output[str]:
        """
        The full path of the group.
        """
        return pulumi.get(self, "full_path")

    @property
    @pulumi.getter(name="lfsEnabled")
    def lfs_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, defaults to true.  Whether to enable LFS
        support for projects in this group.
        """
        return pulumi.get(self, "lfs_enabled")

    @property
    @pulumi.getter(name="mentionsDisabled")
    def mentions_disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, defaults to false.  Disable the capability
        of a group from getting mentioned
        """
        return pulumi.get(self, "mentions_disabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[Optional[int]]:
        """
        Integer, id of the parent group (creates a nested group).
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path of the group.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="projectCreationLevel")
    def project_creation_level(self) -> pulumi.Output[Optional[str]]:
        """
        , defaults to Maintainer.
        Determine if developers can create projects
        in the group. Can be noone (No one), maintainer (Maintainers),
        or developer (Developers + Maintainers).
        """
        return pulumi.get(self, "project_creation_level")

    @property
    @pulumi.getter(name="requestAccessEnabled")
    def request_access_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, defaults to false.  Whether to
        enable users to request access to the group.
        """
        return pulumi.get(self, "request_access_enabled")

    @property
    @pulumi.getter(name="requireTwoFactorAuthentication")
    def require_two_factor_authentication(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, defaults to false.
        equire all users in this group to setup Two-factor authentication.
        """
        return pulumi.get(self, "require_two_factor_authentication")

    @property
    @pulumi.getter(name="runnersToken")
    def runners_token(self) -> pulumi.Output[str]:
        """
        The group level registration token to use during runner setup.
        """
        return pulumi.get(self, "runners_token")

    @property
    @pulumi.getter(name="shareWithGroupLock")
    def share_with_group_lock(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, defaults to false.  Prevent sharing
        a project with another group within this group.
        """
        return pulumi.get(self, "share_with_group_lock")

    @property
    @pulumi.getter(name="subgroupCreationLevel")
    def subgroup_creation_level(self) -> pulumi.Output[Optional[str]]:
        """
        , defaults to Owner.
        Allowed to create subgroups.
        Can be owner (Owners), or maintainer (Maintainers).
        """
        return pulumi.get(self, "subgroup_creation_level")

    @property
    @pulumi.getter(name="twoFactorGracePeriod")
    def two_factor_grace_period(self) -> pulumi.Output[Optional[int]]:
        """
        Int, defaults to 48.
        Time before Two-factor authentication is enforced (in hours).
        """
        return pulumi.get(self, "two_factor_grace_period")

    @property
    @pulumi.getter(name="visibilityLevel")
    def visibility_level(self) -> pulumi.Output[str]:
        """
        The group's visibility. Can be `private`, `internal`, or `public`.
        """
        return pulumi.get(self, "visibility_level")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> pulumi.Output[str]:
        """
        Web URL of the group.
        """
        return pulumi.get(self, "web_url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

