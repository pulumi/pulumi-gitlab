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

__all__ = ['GroupEpicBoardArgs', 'GroupEpicBoard']

@pulumi.input_type
class GroupEpicBoardArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input['GroupEpicBoardListArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GroupEpicBoard resource.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group owned by the authenticated user.
        :param pulumi.Input[Sequence[pulumi.Input['GroupEpicBoardListArgs']]] lists: The list of epic board lists.
        :param pulumi.Input[str] name: The name of the board.
        """
        pulumi.set(__self__, "group", group)
        if lists is not None:
            pulumi.set(__self__, "lists", lists)
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
    def lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupEpicBoardListArgs']]]]:
        """
        The list of epic board lists.
        """
        return pulumi.get(self, "lists")

    @lists.setter
    def lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupEpicBoardListArgs']]]]):
        pulumi.set(self, "lists", value)

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
class _GroupEpicBoardState:
    def __init__(__self__, *,
                 group: Optional[pulumi.Input[str]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input['GroupEpicBoardListArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupEpicBoard resources.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group owned by the authenticated user.
        :param pulumi.Input[Sequence[pulumi.Input['GroupEpicBoardListArgs']]] lists: The list of epic board lists.
        :param pulumi.Input[str] name: The name of the board.
        """
        if group is not None:
            pulumi.set(__self__, "group", group)
        if lists is not None:
            pulumi.set(__self__, "lists", lists)
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
    def lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupEpicBoardListArgs']]]]:
        """
        The list of epic board lists.
        """
        return pulumi.get(self, "lists")

    @lists.setter
    def lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupEpicBoardListArgs']]]]):
        pulumi.set(self, "lists", value)

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


class GroupEpicBoard(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupEpicBoardListArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `GroupEpicBoard` resource allows to manage the lifecycle of a epic board in a group.

        > Multiple epic boards on one group requires a GitLab Premium or above License.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_boards.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.Group("example",
            path="test_group",
            description="An example group")
        label1 = gitlab.GroupLabel("label1",
            group=example.id,
            color="#FF0000")
        label3 = gitlab.GroupLabel("label3",
            group=example.id,
            color="#003000")
        epic_board = gitlab.GroupEpicBoard("epicBoard",
            group=example.path,
            lists=[gitlab.GroupEpicBoardListArgs(
                label_id=label1.label_id,
            )])
        ```

        ## Import

        You can import this resource with an id made up of `{group-id}:{epic-board-id}`, e.g.

        ```sh
         $ pulumi import gitlab:index/groupEpicBoard:GroupEpicBoard agile 70:156
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group owned by the authenticated user.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupEpicBoardListArgs']]]] lists: The list of epic board lists.
        :param pulumi.Input[str] name: The name of the board.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupEpicBoardArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupEpicBoard` resource allows to manage the lifecycle of a epic board in a group.

        > Multiple epic boards on one group requires a GitLab Premium or above License.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_boards.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.Group("example",
            path="test_group",
            description="An example group")
        label1 = gitlab.GroupLabel("label1",
            group=example.id,
            color="#FF0000")
        label3 = gitlab.GroupLabel("label3",
            group=example.id,
            color="#003000")
        epic_board = gitlab.GroupEpicBoard("epicBoard",
            group=example.path,
            lists=[gitlab.GroupEpicBoardListArgs(
                label_id=label1.label_id,
            )])
        ```

        ## Import

        You can import this resource with an id made up of `{group-id}:{epic-board-id}`, e.g.

        ```sh
         $ pulumi import gitlab:index/groupEpicBoard:GroupEpicBoard agile 70:156
        ```

        :param str resource_name: The name of the resource.
        :param GroupEpicBoardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupEpicBoardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupEpicBoardListArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupEpicBoardArgs.__new__(GroupEpicBoardArgs)

            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["lists"] = lists
            __props__.__dict__["name"] = name
        super(GroupEpicBoard, __self__).__init__(
            'gitlab:index/groupEpicBoard:GroupEpicBoard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group: Optional[pulumi.Input[str]] = None,
            lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupEpicBoardListArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'GroupEpicBoard':
        """
        Get an existing GroupEpicBoard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group owned by the authenticated user.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupEpicBoardListArgs']]]] lists: The list of epic board lists.
        :param pulumi.Input[str] name: The name of the board.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupEpicBoardState.__new__(_GroupEpicBoardState)

        __props__.__dict__["group"] = group
        __props__.__dict__["lists"] = lists
        __props__.__dict__["name"] = name
        return GroupEpicBoard(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of the group owned by the authenticated user.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def lists(self) -> pulumi.Output[Optional[Sequence['outputs.GroupEpicBoardList']]]:
        """
        The list of epic board lists.
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the board.
        """
        return pulumi.get(self, "name")
