# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectMembershipArgs', 'ProjectMembership']

@pulumi.input_type
class ProjectMembershipArgs:
    def __init__(__self__, *,
                 access_level: pulumi.Input[str],
                 project: pulumi.Input[str],
                 user_id: pulumi.Input[int],
                 expires_at: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectMembership resource.
        :param pulumi.Input[str] access_level: The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        :param pulumi.Input[int] user_id: The id of the user.
        :param pulumi.Input[str] expires_at: Expiration date for the project membership. Format: `YYYY-MM-DD`
        """
        pulumi.set(__self__, "access_level", access_level)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "user_id", user_id)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> pulumi.Input[str]:
        """
        The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_level", value)

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
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[int]:
        """
        The id of the user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration date for the project membership. Format: `YYYY-MM-DD`
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)


@pulumi.input_type
class _ProjectMembershipState:
    def __init__(__self__, *,
                 access_level: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ProjectMembership resources.
        :param pulumi.Input[str] access_level: The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] expires_at: Expiration date for the project membership. Format: `YYYY-MM-DD`
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        :param pulumi.Input[int] user_id: The id of the user.
        """
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration date for the project membership. Format: `YYYY-MM-DD`
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

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

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "user_id", value)


class ProjectMembership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The `ProjectMembership` resource allows to manage the lifecycle of a users project membership.

        > If a project should grant membership to an entire group use the `ProjectShareGroup` resource instead.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        test = gitlab.ProjectMembership("test",
            access_level="guest",
            project="12345",
            user_id=1337)
        example = gitlab.ProjectMembership("example",
            access_level="guest",
            expires_at="2022-12-31",
            project="67890",
            user_id=1234)
        ```

        ## Import

        GitLab project membership can be imported using an id made up of `project_id:user_id`, e.g.

        ```sh
         $ pulumi import gitlab:index/projectMembership:ProjectMembership test "12345:1337"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] expires_at: Expiration date for the project membership. Format: `YYYY-MM-DD`
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        :param pulumi.Input[int] user_id: The id of the user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectMembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectMembership` resource allows to manage the lifecycle of a users project membership.

        > If a project should grant membership to an entire group use the `ProjectShareGroup` resource instead.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        test = gitlab.ProjectMembership("test",
            access_level="guest",
            project="12345",
            user_id=1337)
        example = gitlab.ProjectMembership("example",
            access_level="guest",
            expires_at="2022-12-31",
            project="67890",
            user_id=1234)
        ```

        ## Import

        GitLab project membership can be imported using an id made up of `project_id:user_id`, e.g.

        ```sh
         $ pulumi import gitlab:index/projectMembership:ProjectMembership test "12345:1337"
        ```

        :param str resource_name: The name of the resource.
        :param ProjectMembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectMembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectMembershipArgs.__new__(ProjectMembershipArgs)

            if access_level is None and not opts.urn:
                raise TypeError("Missing required property 'access_level'")
            __props__.__dict__["access_level"] = access_level
            __props__.__dict__["expires_at"] = expires_at
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(ProjectMembership, __self__).__init__(
            'gitlab:index/projectMembership:ProjectMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_level: Optional[pulumi.Input[str]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[int]] = None) -> 'ProjectMembership':
        """
        Get an existing ProjectMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        :param pulumi.Input[str] expires_at: Expiration date for the project membership. Format: `YYYY-MM-DD`
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        :param pulumi.Input[int] user_id: The id of the user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectMembershipState.__new__(_ProjectMembershipState)

        __props__.__dict__["access_level"] = access_level
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["project"] = project
        __props__.__dict__["user_id"] = user_id
        return ProjectMembership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> pulumi.Output[str]:
        """
        The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        """
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        """
        Expiration date for the project membership. Format: `YYYY-MM-DD`
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[int]:
        """
        The id of the user.
        """
        return pulumi.get(self, "user_id")

