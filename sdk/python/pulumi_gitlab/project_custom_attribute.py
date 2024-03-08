# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectCustomAttributeArgs', 'ProjectCustomAttribute']

@pulumi.input_type
class ProjectCustomAttributeArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 project: pulumi.Input[int],
                 value: pulumi.Input[str]):
        """
        The set of arguments for constructing a ProjectCustomAttribute resource.
        :param pulumi.Input[str] key: Key for the Custom Attribute.
        :param pulumi.Input[int] project: The id of the project.
        :param pulumi.Input[str] value: Value for the Custom Attribute.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Key for the Custom Attribute.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[int]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[int]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Value for the Custom Attribute.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class _ProjectCustomAttributeState:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[int]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectCustomAttribute resources.
        :param pulumi.Input[str] key: Key for the Custom Attribute.
        :param pulumi.Input[int] project: The id of the project.
        :param pulumi.Input[str] value: Value for the Custom Attribute.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Key for the Custom Attribute.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Value for the Custom Attribute.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class ProjectCustomAttribute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[int]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `ProjectCustomAttribute` resource allows to manage custom attributes for a project.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/custom_attributes.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        attr = gitlab.ProjectCustomAttribute("attr",
            key="location",
            project=42,
            value="Greenland")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        You can import a project custom attribute using an id made up of `{project-id}:{key}`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectCustomAttribute:ProjectCustomAttribute attr 42:location
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: Key for the Custom Attribute.
        :param pulumi.Input[int] project: The id of the project.
        :param pulumi.Input[str] value: Value for the Custom Attribute.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectCustomAttributeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectCustomAttribute` resource allows to manage custom attributes for a project.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/custom_attributes.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        attr = gitlab.ProjectCustomAttribute("attr",
            key="location",
            project=42,
            value="Greenland")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        You can import a project custom attribute using an id made up of `{project-id}:{key}`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectCustomAttribute:ProjectCustomAttribute attr 42:location
        ```

        :param str resource_name: The name of the resource.
        :param ProjectCustomAttributeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectCustomAttributeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[int]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectCustomAttributeArgs.__new__(ProjectCustomAttributeArgs)

            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
        super(ProjectCustomAttribute, __self__).__init__(
            'gitlab:index/projectCustomAttribute:ProjectCustomAttribute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[int]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'ProjectCustomAttribute':
        """
        Get an existing ProjectCustomAttribute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: Key for the Custom Attribute.
        :param pulumi.Input[int] project: The id of the project.
        :param pulumi.Input[str] value: Value for the Custom Attribute.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectCustomAttributeState.__new__(_ProjectCustomAttributeState)

        __props__.__dict__["key"] = key
        __props__.__dict__["project"] = project
        __props__.__dict__["value"] = value
        return ProjectCustomAttribute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        Key for the Custom Attribute.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[int]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        Value for the Custom Attribute.
        """
        return pulumi.get(self, "value")

