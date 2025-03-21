# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['GroupShareGroupArgs', 'GroupShareGroup']

@pulumi.input_type
class GroupShareGroupArgs:
    def __init__(__self__, *,
                 group_access: pulumi.Input[str],
                 group_id: pulumi.Input[str],
                 share_group_id: pulumi.Input[int],
                 expires_at: Optional[pulumi.Input[str]] = None,
                 member_role_id: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a GroupShareGroup resource.
        :param pulumi.Input[str] group_access: The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[str] group_id: The id of the main group to be shared.
        :param pulumi.Input[int] share_group_id: The id of the additional group with which the main group will be shared.
        :param pulumi.Input[str] expires_at: Share expiration date. Format: `YYYY-MM-DD`
        :param pulumi.Input[int] member_role_id: The ID of a custom member role. Only available for Ultimate instances.
        """
        pulumi.set(__self__, "group_access", group_access)
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "share_group_id", share_group_id)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if member_role_id is not None:
            pulumi.set(__self__, "member_role_id", member_role_id)

    @property
    @pulumi.getter(name="groupAccess")
    def group_access(self) -> pulumi.Input[str]:
        """
        The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        """
        return pulumi.get(self, "group_access")

    @group_access.setter
    def group_access(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_access", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        The id of the main group to be shared.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="shareGroupId")
    def share_group_id(self) -> pulumi.Input[int]:
        """
        The id of the additional group with which the main group will be shared.
        """
        return pulumi.get(self, "share_group_id")

    @share_group_id.setter
    def share_group_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "share_group_id", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        Share expiration date. Format: `YYYY-MM-DD`
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter(name="memberRoleId")
    def member_role_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of a custom member role. Only available for Ultimate instances.
        """
        return pulumi.get(self, "member_role_id")

    @member_role_id.setter
    def member_role_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "member_role_id", value)


@pulumi.input_type
class _GroupShareGroupState:
    def __init__(__self__, *,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 group_access: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 member_role_id: Optional[pulumi.Input[int]] = None,
                 share_group_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering GroupShareGroup resources.
        :param pulumi.Input[str] expires_at: Share expiration date. Format: `YYYY-MM-DD`
        :param pulumi.Input[str] group_access: The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[str] group_id: The id of the main group to be shared.
        :param pulumi.Input[int] member_role_id: The ID of a custom member role. Only available for Ultimate instances.
        :param pulumi.Input[int] share_group_id: The id of the additional group with which the main group will be shared.
        """
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if group_access is not None:
            pulumi.set(__self__, "group_access", group_access)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if member_role_id is not None:
            pulumi.set(__self__, "member_role_id", member_role_id)
        if share_group_id is not None:
            pulumi.set(__self__, "share_group_id", share_group_id)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        Share expiration date. Format: `YYYY-MM-DD`
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter(name="groupAccess")
    def group_access(self) -> Optional[pulumi.Input[str]]:
        """
        The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        """
        return pulumi.get(self, "group_access")

    @group_access.setter
    def group_access(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_access", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the main group to be shared.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="memberRoleId")
    def member_role_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of a custom member role. Only available for Ultimate instances.
        """
        return pulumi.get(self, "member_role_id")

    @member_role_id.setter
    def member_role_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "member_role_id", value)

    @property
    @pulumi.getter(name="shareGroupId")
    def share_group_id(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the additional group with which the main group will be shared.
        """
        return pulumi.get(self, "share_group_id")

    @share_group_id.setter
    def share_group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "share_group_id", value)


class GroupShareGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 group_access: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 member_role_id: Optional[pulumi.Input[int]] = None,
                 share_group_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The `GroupShareGroup` resource allows managing the lifecycle of a group shared with another group.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#share-groups-with-groups)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        test = gitlab.GroupShareGroup("test",
            group_id=foo["id"],
            share_group_id=bar["id"],
            group_access="guest",
            expires_at="2099-01-01")
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_share_group`. For example:

        terraform

        import {

          to = gitlab_group_share_group.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        GitLab group shares can be imported using an id made up of `mainGroupId:shareGroupId`, e.g.

        ```sh
        $ pulumi import gitlab:index/groupShareGroup:GroupShareGroup test 12345:1337
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expires_at: Share expiration date. Format: `YYYY-MM-DD`
        :param pulumi.Input[str] group_access: The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[str] group_id: The id of the main group to be shared.
        :param pulumi.Input[int] member_role_id: The ID of a custom member role. Only available for Ultimate instances.
        :param pulumi.Input[int] share_group_id: The id of the additional group with which the main group will be shared.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupShareGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupShareGroup` resource allows managing the lifecycle of a group shared with another group.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#share-groups-with-groups)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        test = gitlab.GroupShareGroup("test",
            group_id=foo["id"],
            share_group_id=bar["id"],
            group_access="guest",
            expires_at="2099-01-01")
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_share_group`. For example:

        terraform

        import {

          to = gitlab_group_share_group.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        GitLab group shares can be imported using an id made up of `mainGroupId:shareGroupId`, e.g.

        ```sh
        $ pulumi import gitlab:index/groupShareGroup:GroupShareGroup test 12345:1337
        ```

        :param str resource_name: The name of the resource.
        :param GroupShareGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupShareGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 group_access: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 member_role_id: Optional[pulumi.Input[int]] = None,
                 share_group_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupShareGroupArgs.__new__(GroupShareGroupArgs)

            __props__.__dict__["expires_at"] = expires_at
            if group_access is None and not opts.urn:
                raise TypeError("Missing required property 'group_access'")
            __props__.__dict__["group_access"] = group_access
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["member_role_id"] = member_role_id
            if share_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'share_group_id'")
            __props__.__dict__["share_group_id"] = share_group_id
        super(GroupShareGroup, __self__).__init__(
            'gitlab:index/groupShareGroup:GroupShareGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            group_access: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            member_role_id: Optional[pulumi.Input[int]] = None,
            share_group_id: Optional[pulumi.Input[int]] = None) -> 'GroupShareGroup':
        """
        Get an existing GroupShareGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expires_at: Share expiration date. Format: `YYYY-MM-DD`
        :param pulumi.Input[str] group_access: The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[str] group_id: The id of the main group to be shared.
        :param pulumi.Input[int] member_role_id: The ID of a custom member role. Only available for Ultimate instances.
        :param pulumi.Input[int] share_group_id: The id of the additional group with which the main group will be shared.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupShareGroupState.__new__(_GroupShareGroupState)

        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["group_access"] = group_access
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["member_role_id"] = member_role_id
        __props__.__dict__["share_group_id"] = share_group_id
        return GroupShareGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        """
        Share expiration date. Format: `YYYY-MM-DD`
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter(name="groupAccess")
    def group_access(self) -> pulumi.Output[str]:
        """
        The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        """
        return pulumi.get(self, "group_access")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        The id of the main group to be shared.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="memberRoleId")
    def member_role_id(self) -> pulumi.Output[int]:
        """
        The ID of a custom member role. Only available for Ultimate instances.
        """
        return pulumi.get(self, "member_role_id")

    @property
    @pulumi.getter(name="shareGroupId")
    def share_group_id(self) -> pulumi.Output[int]:
        """
        The id of the additional group with which the main group will be shared.
        """
        return pulumi.get(self, "share_group_id")

