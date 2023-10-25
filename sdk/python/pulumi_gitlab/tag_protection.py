# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TagProtectionArgs', 'TagProtection']

@pulumi.input_type
class TagProtectionArgs:
    def __init__(__self__, *,
                 create_access_level: pulumi.Input[str],
                 project: pulumi.Input[str],
                 tag: pulumi.Input[str],
                 allowed_to_creates: Optional[pulumi.Input[Sequence[pulumi.Input['TagProtectionAllowedToCreateArgs']]]] = None):
        """
        The set of arguments for constructing a TagProtection resource.
        :param pulumi.Input[str] create_access_level: Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] tag: Name of the tag or wildcard.
        :param pulumi.Input[Sequence[pulumi.Input['TagProtectionAllowedToCreateArgs']]] allowed_to_creates: User or group which are allowed to create.
        """
        TagProtectionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            create_access_level=create_access_level,
            project=project,
            tag=tag,
            allowed_to_creates=allowed_to_creates,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             create_access_level: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             tag: Optional[pulumi.Input[str]] = None,
             allowed_to_creates: Optional[pulumi.Input[Sequence[pulumi.Input['TagProtectionAllowedToCreateArgs']]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if create_access_level is None and 'createAccessLevel' in kwargs:
            create_access_level = kwargs['createAccessLevel']
        if create_access_level is None:
            raise TypeError("Missing 'create_access_level' argument")
        if project is None:
            raise TypeError("Missing 'project' argument")
        if tag is None:
            raise TypeError("Missing 'tag' argument")
        if allowed_to_creates is None and 'allowedToCreates' in kwargs:
            allowed_to_creates = kwargs['allowedToCreates']

        _setter("create_access_level", create_access_level)
        _setter("project", project)
        _setter("tag", tag)
        if allowed_to_creates is not None:
            _setter("allowed_to_creates", allowed_to_creates)

    @property
    @pulumi.getter(name="createAccessLevel")
    def create_access_level(self) -> pulumi.Input[str]:
        """
        Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
        """
        return pulumi.get(self, "create_access_level")

    @create_access_level.setter
    def create_access_level(self, value: pulumi.Input[str]):
        pulumi.set(self, "create_access_level", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Input[str]:
        """
        Name of the tag or wildcard.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: pulumi.Input[str]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter(name="allowedToCreates")
    def allowed_to_creates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TagProtectionAllowedToCreateArgs']]]]:
        """
        User or group which are allowed to create.
        """
        return pulumi.get(self, "allowed_to_creates")

    @allowed_to_creates.setter
    def allowed_to_creates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TagProtectionAllowedToCreateArgs']]]]):
        pulumi.set(self, "allowed_to_creates", value)


@pulumi.input_type
class _TagProtectionState:
    def __init__(__self__, *,
                 allowed_to_creates: Optional[pulumi.Input[Sequence[pulumi.Input['TagProtectionAllowedToCreateArgs']]]] = None,
                 create_access_level: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TagProtection resources.
        :param pulumi.Input[Sequence[pulumi.Input['TagProtectionAllowedToCreateArgs']]] allowed_to_creates: User or group which are allowed to create.
        :param pulumi.Input[str] create_access_level: Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] tag: Name of the tag or wildcard.
        """
        _TagProtectionState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            allowed_to_creates=allowed_to_creates,
            create_access_level=create_access_level,
            project=project,
            tag=tag,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             allowed_to_creates: Optional[pulumi.Input[Sequence[pulumi.Input['TagProtectionAllowedToCreateArgs']]]] = None,
             create_access_level: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             tag: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if allowed_to_creates is None and 'allowedToCreates' in kwargs:
            allowed_to_creates = kwargs['allowedToCreates']
        if create_access_level is None and 'createAccessLevel' in kwargs:
            create_access_level = kwargs['createAccessLevel']

        if allowed_to_creates is not None:
            _setter("allowed_to_creates", allowed_to_creates)
        if create_access_level is not None:
            _setter("create_access_level", create_access_level)
        if project is not None:
            _setter("project", project)
        if tag is not None:
            _setter("tag", tag)

    @property
    @pulumi.getter(name="allowedToCreates")
    def allowed_to_creates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TagProtectionAllowedToCreateArgs']]]]:
        """
        User or group which are allowed to create.
        """
        return pulumi.get(self, "allowed_to_creates")

    @allowed_to_creates.setter
    def allowed_to_creates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TagProtectionAllowedToCreateArgs']]]]):
        pulumi.set(self, "allowed_to_creates", value)

    @property
    @pulumi.getter(name="createAccessLevel")
    def create_access_level(self) -> Optional[pulumi.Input[str]]:
        """
        Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
        """
        return pulumi.get(self, "create_access_level")

    @create_access_level.setter
    def create_access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_access_level", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the tag or wildcard.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)


class TagProtection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_to_creates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagProtectionAllowedToCreateArgs']]]]] = None,
                 create_access_level: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        tag_protect = gitlab.TagProtection("tagProtect",
            allowed_to_creates=[
                gitlab.TagProtectionAllowedToCreateArgs(
                    user_id=42,
                ),
                gitlab.TagProtectionAllowedToCreateArgs(
                    group_id=43,
                ),
            ],
            create_access_level="developer",
            project="12345",
            tag="TagProtected")
        ```

        ## Import

        Tag protections can be imported using an id made up of `project_id:tag_name`, e.g.

        ```sh
         $ pulumi import gitlab:index/tagProtection:TagProtection example 123456789:v1.0.0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagProtectionAllowedToCreateArgs']]]] allowed_to_creates: User or group which are allowed to create.
        :param pulumi.Input[str] create_access_level: Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] tag: Name of the tag or wildcard.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TagProtectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        tag_protect = gitlab.TagProtection("tagProtect",
            allowed_to_creates=[
                gitlab.TagProtectionAllowedToCreateArgs(
                    user_id=42,
                ),
                gitlab.TagProtectionAllowedToCreateArgs(
                    group_id=43,
                ),
            ],
            create_access_level="developer",
            project="12345",
            tag="TagProtected")
        ```

        ## Import

        Tag protections can be imported using an id made up of `project_id:tag_name`, e.g.

        ```sh
         $ pulumi import gitlab:index/tagProtection:TagProtection example 123456789:v1.0.0
        ```

        :param str resource_name: The name of the resource.
        :param TagProtectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TagProtectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            TagProtectionArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_to_creates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagProtectionAllowedToCreateArgs']]]]] = None,
                 create_access_level: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TagProtectionArgs.__new__(TagProtectionArgs)

            __props__.__dict__["allowed_to_creates"] = allowed_to_creates
            if create_access_level is None and not opts.urn:
                raise TypeError("Missing required property 'create_access_level'")
            __props__.__dict__["create_access_level"] = create_access_level
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if tag is None and not opts.urn:
                raise TypeError("Missing required property 'tag'")
            __props__.__dict__["tag"] = tag
        super(TagProtection, __self__).__init__(
            'gitlab:index/tagProtection:TagProtection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_to_creates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagProtectionAllowedToCreateArgs']]]]] = None,
            create_access_level: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            tag: Optional[pulumi.Input[str]] = None) -> 'TagProtection':
        """
        Get an existing TagProtection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagProtectionAllowedToCreateArgs']]]] allowed_to_creates: User or group which are allowed to create.
        :param pulumi.Input[str] create_access_level: Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] tag: Name of the tag or wildcard.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TagProtectionState.__new__(_TagProtectionState)

        __props__.__dict__["allowed_to_creates"] = allowed_to_creates
        __props__.__dict__["create_access_level"] = create_access_level
        __props__.__dict__["project"] = project
        __props__.__dict__["tag"] = tag
        return TagProtection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedToCreates")
    def allowed_to_creates(self) -> pulumi.Output[Optional[Sequence['outputs.TagProtectionAllowedToCreate']]]:
        """
        User or group which are allowed to create.
        """
        return pulumi.get(self, "allowed_to_creates")

    @property
    @pulumi.getter(name="createAccessLevel")
    def create_access_level(self) -> pulumi.Output[str]:
        """
        Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
        """
        return pulumi.get(self, "create_access_level")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Output[str]:
        """
        Name of the tag or wildcard.
        """
        return pulumi.get(self, "tag")

