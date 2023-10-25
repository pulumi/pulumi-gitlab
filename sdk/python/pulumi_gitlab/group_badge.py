# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GroupBadgeArgs', 'GroupBadge']

@pulumi.input_type
class GroupBadgeArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 image_url: pulumi.Input[str],
                 link_url: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GroupBadge resource.
        :param pulumi.Input[str] group: The id of the group to add the badge to.
        :param pulumi.Input[str] image_url: The image url which will be presented on group overview.
        :param pulumi.Input[str] link_url: The url linked with the badge.
        :param pulumi.Input[str] name: The name of the badge.
        """
        GroupBadgeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            group=group,
            image_url=image_url,
            link_url=link_url,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             group: Optional[pulumi.Input[str]] = None,
             image_url: Optional[pulumi.Input[str]] = None,
             link_url: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if group is None:
            raise TypeError("Missing 'group' argument")
        if image_url is None and 'imageUrl' in kwargs:
            image_url = kwargs['imageUrl']
        if image_url is None:
            raise TypeError("Missing 'image_url' argument")
        if link_url is None and 'linkUrl' in kwargs:
            link_url = kwargs['linkUrl']
        if link_url is None:
            raise TypeError("Missing 'link_url' argument")

        _setter("group", group)
        _setter("image_url", image_url)
        _setter("link_url", link_url)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        The id of the group to add the badge to.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="imageUrl")
    def image_url(self) -> pulumi.Input[str]:
        """
        The image url which will be presented on group overview.
        """
        return pulumi.get(self, "image_url")

    @image_url.setter
    def image_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_url", value)

    @property
    @pulumi.getter(name="linkUrl")
    def link_url(self) -> pulumi.Input[str]:
        """
        The url linked with the badge.
        """
        return pulumi.get(self, "link_url")

    @link_url.setter
    def link_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "link_url", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the badge.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _GroupBadgeState:
    def __init__(__self__, *,
                 group: Optional[pulumi.Input[str]] = None,
                 image_url: Optional[pulumi.Input[str]] = None,
                 link_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rendered_image_url: Optional[pulumi.Input[str]] = None,
                 rendered_link_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupBadge resources.
        :param pulumi.Input[str] group: The id of the group to add the badge to.
        :param pulumi.Input[str] image_url: The image url which will be presented on group overview.
        :param pulumi.Input[str] link_url: The url linked with the badge.
        :param pulumi.Input[str] name: The name of the badge.
        :param pulumi.Input[str] rendered_image_url: The image_url argument rendered (in case of use of placeholders).
        :param pulumi.Input[str] rendered_link_url: The link_url argument rendered (in case of use of placeholders).
        """
        _GroupBadgeState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            group=group,
            image_url=image_url,
            link_url=link_url,
            name=name,
            rendered_image_url=rendered_image_url,
            rendered_link_url=rendered_link_url,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             group: Optional[pulumi.Input[str]] = None,
             image_url: Optional[pulumi.Input[str]] = None,
             link_url: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             rendered_image_url: Optional[pulumi.Input[str]] = None,
             rendered_link_url: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if image_url is None and 'imageUrl' in kwargs:
            image_url = kwargs['imageUrl']
        if link_url is None and 'linkUrl' in kwargs:
            link_url = kwargs['linkUrl']
        if rendered_image_url is None and 'renderedImageUrl' in kwargs:
            rendered_image_url = kwargs['renderedImageUrl']
        if rendered_link_url is None and 'renderedLinkUrl' in kwargs:
            rendered_link_url = kwargs['renderedLinkUrl']

        if group is not None:
            _setter("group", group)
        if image_url is not None:
            _setter("image_url", image_url)
        if link_url is not None:
            _setter("link_url", link_url)
        if name is not None:
            _setter("name", name)
        if rendered_image_url is not None:
            _setter("rendered_image_url", rendered_image_url)
        if rendered_link_url is not None:
            _setter("rendered_link_url", rendered_link_url)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the group to add the badge to.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="imageUrl")
    def image_url(self) -> Optional[pulumi.Input[str]]:
        """
        The image url which will be presented on group overview.
        """
        return pulumi.get(self, "image_url")

    @image_url.setter
    def image_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_url", value)

    @property
    @pulumi.getter(name="linkUrl")
    def link_url(self) -> Optional[pulumi.Input[str]]:
        """
        The url linked with the badge.
        """
        return pulumi.get(self, "link_url")

    @link_url.setter
    def link_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_url", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the badge.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="renderedImageUrl")
    def rendered_image_url(self) -> Optional[pulumi.Input[str]]:
        """
        The image_url argument rendered (in case of use of placeholders).
        """
        return pulumi.get(self, "rendered_image_url")

    @rendered_image_url.setter
    def rendered_image_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rendered_image_url", value)

    @property
    @pulumi.getter(name="renderedLinkUrl")
    def rendered_link_url(self) -> Optional[pulumi.Input[str]]:
        """
        The link_url argument rendered (in case of use of placeholders).
        """
        return pulumi.get(self, "rendered_link_url")

    @rendered_link_url.setter
    def rendered_link_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rendered_link_url", value)


class GroupBadge(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 image_url: Optional[pulumi.Input[str]] = None,
                 link_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `GroupBadge` resource allows to manage the lifecycle of group badges.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/badges.html#group-badges)

        ## Import

        GitLab group badges can be imported using an id made up of `{group_id}:{badge_id}`, e.g.

        ```sh
         $ pulumi import gitlab:index/groupBadge:GroupBadge foo 1:3
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The id of the group to add the badge to.
        :param pulumi.Input[str] image_url: The image url which will be presented on group overview.
        :param pulumi.Input[str] link_url: The url linked with the badge.
        :param pulumi.Input[str] name: The name of the badge.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupBadgeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupBadge` resource allows to manage the lifecycle of group badges.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/badges.html#group-badges)

        ## Import

        GitLab group badges can be imported using an id made up of `{group_id}:{badge_id}`, e.g.

        ```sh
         $ pulumi import gitlab:index/groupBadge:GroupBadge foo 1:3
        ```

        :param str resource_name: The name of the resource.
        :param GroupBadgeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupBadgeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            GroupBadgeArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 image_url: Optional[pulumi.Input[str]] = None,
                 link_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupBadgeArgs.__new__(GroupBadgeArgs)

            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            if image_url is None and not opts.urn:
                raise TypeError("Missing required property 'image_url'")
            __props__.__dict__["image_url"] = image_url
            if link_url is None and not opts.urn:
                raise TypeError("Missing required property 'link_url'")
            __props__.__dict__["link_url"] = link_url
            __props__.__dict__["name"] = name
            __props__.__dict__["rendered_image_url"] = None
            __props__.__dict__["rendered_link_url"] = None
        super(GroupBadge, __self__).__init__(
            'gitlab:index/groupBadge:GroupBadge',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group: Optional[pulumi.Input[str]] = None,
            image_url: Optional[pulumi.Input[str]] = None,
            link_url: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rendered_image_url: Optional[pulumi.Input[str]] = None,
            rendered_link_url: Optional[pulumi.Input[str]] = None) -> 'GroupBadge':
        """
        Get an existing GroupBadge resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The id of the group to add the badge to.
        :param pulumi.Input[str] image_url: The image url which will be presented on group overview.
        :param pulumi.Input[str] link_url: The url linked with the badge.
        :param pulumi.Input[str] name: The name of the badge.
        :param pulumi.Input[str] rendered_image_url: The image_url argument rendered (in case of use of placeholders).
        :param pulumi.Input[str] rendered_link_url: The link_url argument rendered (in case of use of placeholders).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupBadgeState.__new__(_GroupBadgeState)

        __props__.__dict__["group"] = group
        __props__.__dict__["image_url"] = image_url
        __props__.__dict__["link_url"] = link_url
        __props__.__dict__["name"] = name
        __props__.__dict__["rendered_image_url"] = rendered_image_url
        __props__.__dict__["rendered_link_url"] = rendered_link_url
        return GroupBadge(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The id of the group to add the badge to.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="imageUrl")
    def image_url(self) -> pulumi.Output[str]:
        """
        The image url which will be presented on group overview.
        """
        return pulumi.get(self, "image_url")

    @property
    @pulumi.getter(name="linkUrl")
    def link_url(self) -> pulumi.Output[str]:
        """
        The url linked with the badge.
        """
        return pulumi.get(self, "link_url")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the badge.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="renderedImageUrl")
    def rendered_image_url(self) -> pulumi.Output[str]:
        """
        The image_url argument rendered (in case of use of placeholders).
        """
        return pulumi.get(self, "rendered_image_url")

    @property
    @pulumi.getter(name="renderedLinkUrl")
    def rendered_link_url(self) -> pulumi.Output[str]:
        """
        The link_url argument rendered (in case of use of placeholders).
        """
        return pulumi.get(self, "rendered_link_url")

