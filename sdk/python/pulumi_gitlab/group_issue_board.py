# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GroupIssueBoardArgs', 'GroupIssueBoard']

@pulumi.input_type
class GroupIssueBoardArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input['GroupIssueBoardListArgs']]]] = None,
                 milestone_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GroupIssueBoard resource.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group owned by the authenticated user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: The list of label names which the board should be scoped to.
        :param pulumi.Input[Sequence[pulumi.Input['GroupIssueBoardListArgs']]] lists: The list of issue board lists.
        :param pulumi.Input[int] milestone_id: The milestone the board should be scoped to.
        :param pulumi.Input[str] name: The name of the board.
        """
        pulumi.set(__self__, "group", group)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if lists is not None:
            pulumi.set(__self__, "lists", lists)
        if milestone_id is not None:
            pulumi.set(__self__, "milestone_id", milestone_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        The ID or URL-encoded path of the group owned by the authenticated user.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of label names which the board should be scoped to.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupIssueBoardListArgs']]]]:
        """
        The list of issue board lists.
        """
        return pulumi.get(self, "lists")

    @lists.setter
    def lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupIssueBoardListArgs']]]]):
        pulumi.set(self, "lists", value)

    @property
    @pulumi.getter(name="milestoneId")
    def milestone_id(self) -> Optional[pulumi.Input[int]]:
        """
        The milestone the board should be scoped to.
        """
        return pulumi.get(self, "milestone_id")

    @milestone_id.setter
    def milestone_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "milestone_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the board.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _GroupIssueBoardState:
    def __init__(__self__, *,
                 group: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input['GroupIssueBoardListArgs']]]] = None,
                 milestone_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupIssueBoard resources.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group owned by the authenticated user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: The list of label names which the board should be scoped to.
        :param pulumi.Input[Sequence[pulumi.Input['GroupIssueBoardListArgs']]] lists: The list of issue board lists.
        :param pulumi.Input[int] milestone_id: The milestone the board should be scoped to.
        :param pulumi.Input[str] name: The name of the board.
        """
        if group is not None:
            pulumi.set(__self__, "group", group)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if lists is not None:
            pulumi.set(__self__, "lists", lists)
        if milestone_id is not None:
            pulumi.set(__self__, "milestone_id", milestone_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or URL-encoded path of the group owned by the authenticated user.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of label names which the board should be scoped to.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupIssueBoardListArgs']]]]:
        """
        The list of issue board lists.
        """
        return pulumi.get(self, "lists")

    @lists.setter
    def lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupIssueBoardListArgs']]]]):
        pulumi.set(self, "lists", value)

    @property
    @pulumi.getter(name="milestoneId")
    def milestone_id(self) -> Optional[pulumi.Input[int]]:
        """
        The milestone the board should be scoped to.
        """
        return pulumi.get(self, "milestone_id")

    @milestone_id.setter
    def milestone_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "milestone_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the board.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class GroupIssueBoard(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupIssueBoardListArgs']]]]] = None,
                 milestone_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `GroupIssueBoard` resource allows to manage the lifecycle of a issue board in a group.

        > Multiple issue boards on one group requires a GitLab Premium or above License.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_boards.html)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group owned by the authenticated user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: The list of label names which the board should be scoped to.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupIssueBoardListArgs']]]] lists: The list of issue board lists.
        :param pulumi.Input[int] milestone_id: The milestone the board should be scoped to.
        :param pulumi.Input[str] name: The name of the board.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupIssueBoardArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupIssueBoard` resource allows to manage the lifecycle of a issue board in a group.

        > Multiple issue boards on one group requires a GitLab Premium or above License.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_boards.html)

        :param str resource_name: The name of the resource.
        :param GroupIssueBoardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupIssueBoardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupIssueBoardListArgs']]]]] = None,
                 milestone_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupIssueBoardArgs.__new__(GroupIssueBoardArgs)

            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["labels"] = labels
            __props__.__dict__["lists"] = lists
            __props__.__dict__["milestone_id"] = milestone_id
            __props__.__dict__["name"] = name
        super(GroupIssueBoard, __self__).__init__(
            'gitlab:index/groupIssueBoard:GroupIssueBoard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupIssueBoardListArgs']]]]] = None,
            milestone_id: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'GroupIssueBoard':
        """
        Get an existing GroupIssueBoard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group owned by the authenticated user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: The list of label names which the board should be scoped to.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupIssueBoardListArgs']]]] lists: The list of issue board lists.
        :param pulumi.Input[int] milestone_id: The milestone the board should be scoped to.
        :param pulumi.Input[str] name: The name of the board.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupIssueBoardState.__new__(_GroupIssueBoardState)

        __props__.__dict__["group"] = group
        __props__.__dict__["labels"] = labels
        __props__.__dict__["lists"] = lists
        __props__.__dict__["milestone_id"] = milestone_id
        __props__.__dict__["name"] = name
        return GroupIssueBoard(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of the group owned by the authenticated user.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of label names which the board should be scoped to.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def lists(self) -> pulumi.Output[Optional[Sequence['outputs.GroupIssueBoardList']]]:
        """
        The list of issue board lists.
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="milestoneId")
    def milestone_id(self) -> pulumi.Output[Optional[int]]:
        """
        The milestone the board should be scoped to.
        """
        return pulumi.get(self, "milestone_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the board.
        """
        return pulumi.get(self, "name")

