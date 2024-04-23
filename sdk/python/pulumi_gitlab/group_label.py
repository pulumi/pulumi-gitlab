# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GroupLabelArgs', 'GroupLabel']

@pulumi.input_type
class GroupLabelArgs:
    def __init__(__self__, *,
                 color: pulumi.Input[str],
                 group: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GroupLabel resource.
        :param pulumi.Input[str] color: The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        :param pulumi.Input[str] group: The name or id of the group to add the label to.
        :param pulumi.Input[str] description: The description of the label.
        :param pulumi.Input[str] name: The name of the label.
        """
        pulumi.set(__self__, "color", color)
        pulumi.set(__self__, "group", group)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Input[str]:
        """
        The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: pulumi.Input[str]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        The name or id of the group to add the label to.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the label.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the label.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _GroupLabelState:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 label_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupLabel resources.
        :param pulumi.Input[str] color: The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        :param pulumi.Input[str] description: The description of the label.
        :param pulumi.Input[str] group: The name or id of the group to add the label to.
        :param pulumi.Input[int] label_id: The id of the group label.
        :param pulumi.Input[str] name: The name of the label.
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if label_id is not None:
            pulumi.set(__self__, "label_id", label_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[str]]:
        """
        The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the label.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The name or id of the group to add the label to.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="labelId")
    def label_id(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the group label.
        """
        return pulumi.get(self, "label_id")

    @label_id.setter
    def label_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "label_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the label.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class GroupLabel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `GroupLabel` resource allows to manage the lifecycle of labels within a group.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/labels.html#group-labels)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        fixme = gitlab.GroupLabel("fixme",
            group="example",
            name="fixme",
            description="issue with failing tests",
            color="#ffcc00")
        ```

        ## Import

        Gitlab group labels can be imported using an id made up of `{group_id}:{group_label_id}`, e.g.

        ```sh
        $ pulumi import gitlab:index/groupLabel:GroupLabel example 12345:fixme
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        :param pulumi.Input[str] description: The description of the label.
        :param pulumi.Input[str] group: The name or id of the group to add the label to.
        :param pulumi.Input[str] name: The name of the label.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupLabelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupLabel` resource allows to manage the lifecycle of labels within a group.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/labels.html#group-labels)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        fixme = gitlab.GroupLabel("fixme",
            group="example",
            name="fixme",
            description="issue with failing tests",
            color="#ffcc00")
        ```

        ## Import

        Gitlab group labels can be imported using an id made up of `{group_id}:{group_label_id}`, e.g.

        ```sh
        $ pulumi import gitlab:index/groupLabel:GroupLabel example 12345:fixme
        ```

        :param str resource_name: The name of the resource.
        :param GroupLabelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupLabelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupLabelArgs.__new__(GroupLabelArgs)

            if color is None and not opts.urn:
                raise TypeError("Missing required property 'color'")
            __props__.__dict__["color"] = color
            __props__.__dict__["description"] = description
            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["name"] = name
            __props__.__dict__["label_id"] = None
        super(GroupLabel, __self__).__init__(
            'gitlab:index/groupLabel:GroupLabel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            group: Optional[pulumi.Input[str]] = None,
            label_id: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'GroupLabel':
        """
        Get an existing GroupLabel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        :param pulumi.Input[str] description: The description of the label.
        :param pulumi.Input[str] group: The name or id of the group to add the label to.
        :param pulumi.Input[int] label_id: The id of the group label.
        :param pulumi.Input[str] name: The name of the label.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupLabelState.__new__(_GroupLabelState)

        __props__.__dict__["color"] = color
        __props__.__dict__["description"] = description
        __props__.__dict__["group"] = group
        __props__.__dict__["label_id"] = label_id
        __props__.__dict__["name"] = name
        return GroupLabel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[str]:
        """
        The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the label.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The name or id of the group to add the label to.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="labelId")
    def label_id(self) -> pulumi.Output[int]:
        """
        The id of the group label.
        """
        return pulumi.get(self, "label_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the label.
        """
        return pulumi.get(self, "name")

