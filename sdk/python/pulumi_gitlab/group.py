# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Group(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the group.
    """
    full_name: pulumi.Output[str]
    """
    The full name of the group.
    """
    full_path: pulumi.Output[str]
    """
    The full path of the group.
    """
    lfs_enabled: pulumi.Output[bool]
    """
    Boolean, defaults to true.  Whether to enable LFS
    support for projects in this group.
    """
    name: pulumi.Output[str]
    """
    The name of this group.
    """
    parent_id: pulumi.Output[float]
    """
    Integer, id of the parent group (creates a nested group).
    """
    path: pulumi.Output[str]
    """
    The path of the group.
    """
    request_access_enabled: pulumi.Output[bool]
    """
    Boolean, defaults to false.  Whether to
    enable users to request access to the group.
    """
    runners_token: pulumi.Output[str]
    """
    The group level registration token to use during runner setup.
    """
    visibility_level: pulumi.Output[str]
    """
    Set to `public` to create a public group.
    Valid values are `private`, `internal`, `public`.
    Groups are created as private by default.
    """
    web_url: pulumi.Output[str]
    """
    Web URL of the group.
    """
    def __init__(__self__, resource_name, opts=None, description=None, lfs_enabled=None, name=None, parent_id=None, path=None, request_access_enabled=None, visibility_level=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Group resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the group.
        :param pulumi.Input[bool] lfs_enabled: Boolean, defaults to true.  Whether to enable LFS
               support for projects in this group.
        :param pulumi.Input[str] name: The name of this group.
        :param pulumi.Input[float] parent_id: Integer, id of the parent group (creates a nested group).
        :param pulumi.Input[str] path: The path of the group.
        :param pulumi.Input[bool] request_access_enabled: Boolean, defaults to false.  Whether to
               enable users to request access to the group.
        :param pulumi.Input[str] visibility_level: Set to `public` to create a public group.
               Valid values are `private`, `internal`, `public`.
               Groups are created as private by default.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['lfs_enabled'] = lfs_enabled
            __props__['name'] = name
            __props__['parent_id'] = parent_id
            if path is None:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
            __props__['request_access_enabled'] = request_access_enabled
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
    def get(resource_name, id, opts=None, description=None, full_name=None, full_path=None, lfs_enabled=None, name=None, parent_id=None, path=None, request_access_enabled=None, runners_token=None, visibility_level=None, web_url=None):
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the group.
        :param pulumi.Input[str] full_name: The full name of the group.
        :param pulumi.Input[str] full_path: The full path of the group.
        :param pulumi.Input[bool] lfs_enabled: Boolean, defaults to true.  Whether to enable LFS
               support for projects in this group.
        :param pulumi.Input[str] name: The name of this group.
        :param pulumi.Input[float] parent_id: Integer, id of the parent group (creates a nested group).
        :param pulumi.Input[str] path: The path of the group.
        :param pulumi.Input[bool] request_access_enabled: Boolean, defaults to false.  Whether to
               enable users to request access to the group.
        :param pulumi.Input[str] runners_token: The group level registration token to use during runner setup.
        :param pulumi.Input[str] visibility_level: Set to `public` to create a public group.
               Valid values are `private`, `internal`, `public`.
               Groups are created as private by default.
        :param pulumi.Input[str] web_url: Web URL of the group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["full_name"] = full_name
        __props__["full_path"] = full_path
        __props__["lfs_enabled"] = lfs_enabled
        __props__["name"] = name
        __props__["parent_id"] = parent_id
        __props__["path"] = path
        __props__["request_access_enabled"] = request_access_enabled
        __props__["runners_token"] = runners_token
        __props__["visibility_level"] = visibility_level
        __props__["web_url"] = web_url
        return Group(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

