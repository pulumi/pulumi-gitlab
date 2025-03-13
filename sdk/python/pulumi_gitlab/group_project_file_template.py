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

__all__ = ['GroupProjectFileTemplateArgs', 'GroupProjectFileTemplate']

@pulumi.input_type
class GroupProjectFileTemplateArgs:
    def __init__(__self__, *,
                 file_template_project_id: pulumi.Input[int],
                 group_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a GroupProjectFileTemplate resource.
        :param pulumi.Input[int] file_template_project_id: The ID of the project that will be used for file templates. This project must be the direct
               			child of the project defined by the group_id
        :param pulumi.Input[int] group_id: The ID of the group that will use the file template project. This group must be the direct
                           parent of the project defined by project_id
        """
        pulumi.set(__self__, "file_template_project_id", file_template_project_id)
        pulumi.set(__self__, "group_id", group_id)

    @property
    @pulumi.getter(name="fileTemplateProjectId")
    def file_template_project_id(self) -> pulumi.Input[int]:
        """
        The ID of the project that will be used for file templates. This project must be the direct
        			child of the project defined by the group_id
        """
        return pulumi.get(self, "file_template_project_id")

    @file_template_project_id.setter
    def file_template_project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "file_template_project_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[int]:
        """
        The ID of the group that will use the file template project. This group must be the direct
                    parent of the project defined by project_id
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "group_id", value)


@pulumi.input_type
class _GroupProjectFileTemplateState:
    def __init__(__self__, *,
                 file_template_project_id: Optional[pulumi.Input[int]] = None,
                 group_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering GroupProjectFileTemplate resources.
        :param pulumi.Input[int] file_template_project_id: The ID of the project that will be used for file templates. This project must be the direct
               			child of the project defined by the group_id
        :param pulumi.Input[int] group_id: The ID of the group that will use the file template project. This group must be the direct
                           parent of the project defined by project_id
        """
        if file_template_project_id is not None:
            pulumi.set(__self__, "file_template_project_id", file_template_project_id)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)

    @property
    @pulumi.getter(name="fileTemplateProjectId")
    def file_template_project_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the project that will be used for file templates. This project must be the direct
        			child of the project defined by the group_id
        """
        return pulumi.get(self, "file_template_project_id")

    @file_template_project_id.setter
    def file_template_project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "file_template_project_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the group that will use the file template project. This group must be the direct
                    parent of the project defined by project_id
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_id", value)


class GroupProjectFileTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_template_project_id: Optional[pulumi.Input[int]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The `GroupProjectFileTemplate` resource allows setting a project from which
        custom file templates will be loaded. In order to use this resource, the project selected must be a direct child of
        the group selected. After the resource has run, `gitlab_project_template.template_project_id` is available for use.
        For more information about which file types are available as templates, view
        [GitLab's documentation](https://docs.gitlab.com/user/group/custom_project_templates/)

        > This resource requires a GitLab Enterprise instance with a Premium license.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#update-group)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.Group("foo",
            name="group",
            path="group",
            description="An example group")
        bar = gitlab.Project("bar",
            name="template project",
            description="contains file templates",
            visibility_level="public",
            namespace_id=foo.id)
        template_link = gitlab.GroupProjectFileTemplate("template_link",
            group_id=foo.id,
            file_template_project_id=bar.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] file_template_project_id: The ID of the project that will be used for file templates. This project must be the direct
               			child of the project defined by the group_id
        :param pulumi.Input[int] group_id: The ID of the group that will use the file template project. This group must be the direct
                           parent of the project defined by project_id
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupProjectFileTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupProjectFileTemplate` resource allows setting a project from which
        custom file templates will be loaded. In order to use this resource, the project selected must be a direct child of
        the group selected. After the resource has run, `gitlab_project_template.template_project_id` is available for use.
        For more information about which file types are available as templates, view
        [GitLab's documentation](https://docs.gitlab.com/user/group/custom_project_templates/)

        > This resource requires a GitLab Enterprise instance with a Premium license.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#update-group)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.Group("foo",
            name="group",
            path="group",
            description="An example group")
        bar = gitlab.Project("bar",
            name="template project",
            description="contains file templates",
            visibility_level="public",
            namespace_id=foo.id)
        template_link = gitlab.GroupProjectFileTemplate("template_link",
            group_id=foo.id,
            file_template_project_id=bar.id)
        ```

        :param str resource_name: The name of the resource.
        :param GroupProjectFileTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupProjectFileTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_template_project_id: Optional[pulumi.Input[int]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupProjectFileTemplateArgs.__new__(GroupProjectFileTemplateArgs)

            if file_template_project_id is None and not opts.urn:
                raise TypeError("Missing required property 'file_template_project_id'")
            __props__.__dict__["file_template_project_id"] = file_template_project_id
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
        super(GroupProjectFileTemplate, __self__).__init__(
            'gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            file_template_project_id: Optional[pulumi.Input[int]] = None,
            group_id: Optional[pulumi.Input[int]] = None) -> 'GroupProjectFileTemplate':
        """
        Get an existing GroupProjectFileTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] file_template_project_id: The ID of the project that will be used for file templates. This project must be the direct
               			child of the project defined by the group_id
        :param pulumi.Input[int] group_id: The ID of the group that will use the file template project. This group must be the direct
                           parent of the project defined by project_id
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupProjectFileTemplateState.__new__(_GroupProjectFileTemplateState)

        __props__.__dict__["file_template_project_id"] = file_template_project_id
        __props__.__dict__["group_id"] = group_id
        return GroupProjectFileTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fileTemplateProjectId")
    def file_template_project_id(self) -> pulumi.Output[int]:
        """
        The ID of the project that will be used for file templates. This project must be the direct
        			child of the project defined by the group_id
        """
        return pulumi.get(self, "file_template_project_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[int]:
        """
        The ID of the group that will use the file template project. This group must be the direct
                    parent of the project defined by project_id
        """
        return pulumi.get(self, "group_id")

