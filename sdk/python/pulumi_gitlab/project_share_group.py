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

__all__ = ['ProjectShareGroupArgs', 'ProjectShareGroup']

@pulumi.input_type
class ProjectShareGroupArgs:
    def __init__(__self__, *,
                 group_id: pulumi.Input[int],
                 project: pulumi.Input[str],
                 access_level: Optional[pulumi.Input[str]] = None,
                 group_access: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectShareGroup resource.
        :param pulumi.Input[int] group_id: The id of the group.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        :param pulumi.Input[str] access_level: The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[str] group_access: The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        """
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "project", project)
        if access_level is not None:
            warnings.warn("""Use `group_access` instead of the `access_level` attribute.""", DeprecationWarning)
            pulumi.log.warn("""access_level is deprecated: Use `group_access` instead of the `access_level` attribute.""")
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if group_access is not None:
            pulumi.set(__self__, "group_access", group_access)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[int]:
        """
        The id of the group.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or URL-encoded path of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="accessLevel")
    @_utilities.deprecated("""Use `group_access` instead of the `access_level` attribute.""")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter(name="groupAccess")
    def group_access(self) -> Optional[pulumi.Input[str]]:
        """
        The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        """
        return pulumi.get(self, "group_access")

    @group_access.setter
    def group_access(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_access", value)


@pulumi.input_type
class _ProjectShareGroupState:
    def __init__(__self__, *,
                 access_level: Optional[pulumi.Input[str]] = None,
                 group_access: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectShareGroup resources.
        :param pulumi.Input[str] access_level: The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[str] group_access: The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[int] group_id: The id of the group.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        """
        if access_level is not None:
            warnings.warn("""Use `group_access` instead of the `access_level` attribute.""", DeprecationWarning)
            pulumi.log.warn("""access_level is deprecated: Use `group_access` instead of the `access_level` attribute.""")
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if group_access is not None:
            pulumi.set(__self__, "group_access", group_access)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="accessLevel")
    @_utilities.deprecated("""Use `group_access` instead of the `access_level` attribute.""")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter(name="groupAccess")
    def group_access(self) -> Optional[pulumi.Input[str]]:
        """
        The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        """
        return pulumi.get(self, "group_access")

    @group_access.setter
    def group_access(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_access", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the group.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or URL-encoded path of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class ProjectShareGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 group_access: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `ProjectShareGroup` resource allows to manage the lifecycle of project shared with a group.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#share-project-with-group)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        test = gitlab.ProjectShareGroup("test",
            project="12345",
            group_id=1337,
            group_access="guest")
        ```

        ## Import

        GitLab project group shares can be imported using an id made up of `projectid:groupid`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectShareGroup:ProjectShareGroup test 12345:1337
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[str] group_access: The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[int] group_id: The id of the group.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectShareGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectShareGroup` resource allows to manage the lifecycle of project shared with a group.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#share-project-with-group)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        test = gitlab.ProjectShareGroup("test",
            project="12345",
            group_id=1337,
            group_access="guest")
        ```

        ## Import

        GitLab project group shares can be imported using an id made up of `projectid:groupid`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectShareGroup:ProjectShareGroup test 12345:1337
        ```

        :param str resource_name: The name of the resource.
        :param ProjectShareGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectShareGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 group_access: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectShareGroupArgs.__new__(ProjectShareGroupArgs)

            __props__.__dict__["access_level"] = access_level
            __props__.__dict__["group_access"] = group_access
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
        super(ProjectShareGroup, __self__).__init__(
            'gitlab:index/projectShareGroup:ProjectShareGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_level: Optional[pulumi.Input[str]] = None,
            group_access: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[int]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'ProjectShareGroup':
        """
        Get an existing ProjectShareGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[str] group_access: The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        :param pulumi.Input[int] group_id: The id of the group.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectShareGroupState.__new__(_ProjectShareGroupState)

        __props__.__dict__["access_level"] = access_level
        __props__.__dict__["group_access"] = group_access
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["project"] = project
        return ProjectShareGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessLevel")
    @_utilities.deprecated("""Use `group_access` instead of the `access_level` attribute.""")
    def access_level(self) -> pulumi.Output[Optional[str]]:
        """
        The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        """
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter(name="groupAccess")
    def group_access(self) -> pulumi.Output[Optional[str]]:
        """
        The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
        """
        return pulumi.get(self, "group_access")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[int]:
        """
        The id of the group.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of the project.
        """
        return pulumi.get(self, "project")

