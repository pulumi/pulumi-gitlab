# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GroupSamlLinkArgs', 'GroupSamlLink']

@pulumi.input_type
class GroupSamlLinkArgs:
    def __init__(__self__, *,
                 access_level: pulumi.Input[str],
                 group: pulumi.Input[str],
                 saml_group_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a GroupSamlLink resource.
        :param pulumi.Input[str] access_level: Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        :param pulumi.Input[str] group: The ID or path of the group to add the SAML Group Link to.
        :param pulumi.Input[str] saml_group_name: The name of the SAML group.
        """
        pulumi.set(__self__, "access_level", access_level)
        pulumi.set(__self__, "group", group)
        pulumi.set(__self__, "saml_group_name", saml_group_name)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> pulumi.Input[str]:
        """
        Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        The ID or path of the group to add the SAML Group Link to.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="samlGroupName")
    def saml_group_name(self) -> pulumi.Input[str]:
        """
        The name of the SAML group.
        """
        return pulumi.get(self, "saml_group_name")

    @saml_group_name.setter
    def saml_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "saml_group_name", value)


@pulumi.input_type
class _GroupSamlLinkState:
    def __init__(__self__, *,
                 access_level: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 saml_group_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupSamlLink resources.
        :param pulumi.Input[str] access_level: Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        :param pulumi.Input[str] group: The ID or path of the group to add the SAML Group Link to.
        :param pulumi.Input[str] saml_group_name: The name of the SAML group.
        """
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if saml_group_name is not None:
            pulumi.set(__self__, "saml_group_name", saml_group_name)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or path of the group to add the SAML Group Link to.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="samlGroupName")
    def saml_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SAML group.
        """
        return pulumi.get(self, "saml_group_name")

    @saml_group_name.setter
    def saml_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "saml_group_name", value)


class GroupSamlLink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 saml_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `GroupSamlLink` resource allows to manage the lifecycle of an SAML integration with a group.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#saml-group-links)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        test = gitlab.GroupSamlLink("test",
            group="12345",
            access_level="developer",
            saml_group_name="samlgroupname1")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitLab group saml links can be imported using an id made up of `group_id:saml_group_name`, e.g.

        ```sh
        $ pulumi import gitlab:index/groupSamlLink:GroupSamlLink test "12345:samlgroupname1"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        :param pulumi.Input[str] group: The ID or path of the group to add the SAML Group Link to.
        :param pulumi.Input[str] saml_group_name: The name of the SAML group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupSamlLinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupSamlLink` resource allows to manage the lifecycle of an SAML integration with a group.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#saml-group-links)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        test = gitlab.GroupSamlLink("test",
            group="12345",
            access_level="developer",
            saml_group_name="samlgroupname1")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitLab group saml links can be imported using an id made up of `group_id:saml_group_name`, e.g.

        ```sh
        $ pulumi import gitlab:index/groupSamlLink:GroupSamlLink test "12345:samlgroupname1"
        ```

        :param str resource_name: The name of the resource.
        :param GroupSamlLinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupSamlLinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 saml_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupSamlLinkArgs.__new__(GroupSamlLinkArgs)

            if access_level is None and not opts.urn:
                raise TypeError("Missing required property 'access_level'")
            __props__.__dict__["access_level"] = access_level
            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            if saml_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'saml_group_name'")
            __props__.__dict__["saml_group_name"] = saml_group_name
        super(GroupSamlLink, __self__).__init__(
            'gitlab:index/groupSamlLink:GroupSamlLink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_level: Optional[pulumi.Input[str]] = None,
            group: Optional[pulumi.Input[str]] = None,
            saml_group_name: Optional[pulumi.Input[str]] = None) -> 'GroupSamlLink':
        """
        Get an existing GroupSamlLink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        :param pulumi.Input[str] group: The ID or path of the group to add the SAML Group Link to.
        :param pulumi.Input[str] saml_group_name: The name of the SAML group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupSamlLinkState.__new__(_GroupSamlLinkState)

        __props__.__dict__["access_level"] = access_level
        __props__.__dict__["group"] = group
        __props__.__dict__["saml_group_name"] = saml_group_name
        return GroupSamlLink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> pulumi.Output[str]:
        """
        Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        """
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The ID or path of the group to add the SAML Group Link to.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="samlGroupName")
    def saml_group_name(self) -> pulumi.Output[str]:
        """
        The name of the SAML group.
        """
        return pulumi.get(self, "saml_group_name")

