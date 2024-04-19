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

__all__ = ['ProjectAccessTokenArgs', 'ProjectAccessToken']

@pulumi.input_type
class ProjectAccessTokenArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 scopes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 access_level: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rotation_configuration: Optional[pulumi.Input['ProjectAccessTokenRotationConfigurationArgs']] = None):
        """
        The set of arguments for constructing a ProjectAccessToken resource.
        :param pulumi.Input[str] project: The ID or full path of the project.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        :param pulumi.Input[str] access_level: The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
        :param pulumi.Input[str] expires_at: When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
        :param pulumi.Input[str] name: The name of the project access token.
        :param pulumi.Input['ProjectAccessTokenRotationConfigurationArgs'] rotation_configuration: The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "scopes", scopes)
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rotation_configuration is not None:
            pulumi.set(__self__, "rotation_configuration", rotation_configuration)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or full path of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the project access token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rotationConfiguration")
    def rotation_configuration(self) -> Optional[pulumi.Input['ProjectAccessTokenRotationConfigurationArgs']]:
        """
        The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        """
        return pulumi.get(self, "rotation_configuration")

    @rotation_configuration.setter
    def rotation_configuration(self, value: Optional[pulumi.Input['ProjectAccessTokenRotationConfigurationArgs']]):
        pulumi.set(self, "rotation_configuration", value)


@pulumi.input_type
class _ProjectAccessTokenState:
    def __init__(__self__, *,
                 access_level: Optional[pulumi.Input[str]] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 revoked: Optional[pulumi.Input[bool]] = None,
                 rotation_configuration: Optional[pulumi.Input['ProjectAccessTokenRotationConfigurationArgs']] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ProjectAccessToken resources.
        :param pulumi.Input[str] access_level: The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
        :param pulumi.Input[bool] active: True if the token is active.
        :param pulumi.Input[str] created_at: Time the token has been created, RFC3339 format.
        :param pulumi.Input[str] expires_at: When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
        :param pulumi.Input[str] name: The name of the project access token.
        :param pulumi.Input[str] project: The ID or full path of the project.
        :param pulumi.Input[bool] revoked: True if the token is revoked.
        :param pulumi.Input['ProjectAccessTokenRotationConfigurationArgs'] rotation_configuration: The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        :param pulumi.Input[str] token: The token of the project access token. **Note**: the token is not available for imported resources.
        :param pulumi.Input[int] user_id: The user_id associated to the token.
        """
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if active is not None:
            pulumi.set(__self__, "active", active)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if revoked is not None:
            pulumi.set(__self__, "revoked", revoked)
        if rotation_configuration is not None:
            pulumi.set(__self__, "rotation_configuration", rotation_configuration)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        True if the token is active.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Time the token has been created, RFC3339 format.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the project access token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or full path of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def revoked(self) -> Optional[pulumi.Input[bool]]:
        """
        True if the token is revoked.
        """
        return pulumi.get(self, "revoked")

    @revoked.setter
    def revoked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "revoked", value)

    @property
    @pulumi.getter(name="rotationConfiguration")
    def rotation_configuration(self) -> Optional[pulumi.Input['ProjectAccessTokenRotationConfigurationArgs']]:
        """
        The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        """
        return pulumi.get(self, "rotation_configuration")

    @rotation_configuration.setter
    def rotation_configuration(self, value: Optional[pulumi.Input['ProjectAccessTokenRotationConfigurationArgs']]):
        pulumi.set(self, "rotation_configuration", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The token of the project access token. **Note**: the token is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[int]]:
        """
        The user_id associated to the token.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "user_id", value)


class ProjectAccessToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rotation_configuration: Optional[pulumi.Input[pulumi.InputType['ProjectAccessTokenRotationConfigurationArgs']]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        The `ProjectAccessToken` resource allows to manage the lifecycle of a project access token.

        > Observability scopes are in beta and may not work on all instances. See more details in [the documentation](https://docs.gitlab.com/ee/operations/tracing.html)

        > Use `rotation_configuration` to automatically rotate tokens instead of using `timestamp()` as timestamp will cause changes with every plan. `pulumi up` must still be run to rotate the token.

        > Due to [Automatic reuse detection](https://docs.gitlab.com/ee/api/project_access_tokens.html#automatic-reuse-detection) it's possible that a new Project Access Token will immediately be revoked. Check if an old process using the old token is running if this happens.

        **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/project_access_tokens.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.ProjectAccessToken("example",
            project="25",
            name="Example project access token",
            expires_at="2020-03-14",
            access_level="reporter",
            scopes=["api"])
        example_project_variable = gitlab.ProjectVariable("example",
            project=example_gitlab_project["id"],
            key="pat",
            value=example.token)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        A GitLab Project Access Token can be imported using a key composed of `<project-id>:<token-id>`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectAccessToken:ProjectAccessToken example "12345:1"
        ```

        NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
        :param pulumi.Input[str] expires_at: When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
        :param pulumi.Input[str] name: The name of the project access token.
        :param pulumi.Input[str] project: The ID or full path of the project.
        :param pulumi.Input[pulumi.InputType['ProjectAccessTokenRotationConfigurationArgs']] rotation_configuration: The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectAccessTokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectAccessToken` resource allows to manage the lifecycle of a project access token.

        > Observability scopes are in beta and may not work on all instances. See more details in [the documentation](https://docs.gitlab.com/ee/operations/tracing.html)

        > Use `rotation_configuration` to automatically rotate tokens instead of using `timestamp()` as timestamp will cause changes with every plan. `pulumi up` must still be run to rotate the token.

        > Due to [Automatic reuse detection](https://docs.gitlab.com/ee/api/project_access_tokens.html#automatic-reuse-detection) it's possible that a new Project Access Token will immediately be revoked. Check if an old process using the old token is running if this happens.

        **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/project_access_tokens.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.ProjectAccessToken("example",
            project="25",
            name="Example project access token",
            expires_at="2020-03-14",
            access_level="reporter",
            scopes=["api"])
        example_project_variable = gitlab.ProjectVariable("example",
            project=example_gitlab_project["id"],
            key="pat",
            value=example.token)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        A GitLab Project Access Token can be imported using a key composed of `<project-id>:<token-id>`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectAccessToken:ProjectAccessToken example "12345:1"
        ```

        NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.

        :param str resource_name: The name of the resource.
        :param ProjectAccessTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectAccessTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rotation_configuration: Optional[pulumi.Input[pulumi.InputType['ProjectAccessTokenRotationConfigurationArgs']]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectAccessTokenArgs.__new__(ProjectAccessTokenArgs)

            __props__.__dict__["access_level"] = access_level
            __props__.__dict__["expires_at"] = expires_at
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["rotation_configuration"] = rotation_configuration
            if scopes is None and not opts.urn:
                raise TypeError("Missing required property 'scopes'")
            __props__.__dict__["scopes"] = scopes
            __props__.__dict__["active"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["revoked"] = None
            __props__.__dict__["token"] = None
            __props__.__dict__["user_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ProjectAccessToken, __self__).__init__(
            'gitlab:index/projectAccessToken:ProjectAccessToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_level: Optional[pulumi.Input[str]] = None,
            active: Optional[pulumi.Input[bool]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            revoked: Optional[pulumi.Input[bool]] = None,
            rotation_configuration: Optional[pulumi.Input[pulumi.InputType['ProjectAccessTokenRotationConfigurationArgs']]] = None,
            scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            token: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[int]] = None) -> 'ProjectAccessToken':
        """
        Get an existing ProjectAccessToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
        :param pulumi.Input[bool] active: True if the token is active.
        :param pulumi.Input[str] created_at: Time the token has been created, RFC3339 format.
        :param pulumi.Input[str] expires_at: When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
        :param pulumi.Input[str] name: The name of the project access token.
        :param pulumi.Input[str] project: The ID or full path of the project.
        :param pulumi.Input[bool] revoked: True if the token is revoked.
        :param pulumi.Input[pulumi.InputType['ProjectAccessTokenRotationConfigurationArgs']] rotation_configuration: The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        :param pulumi.Input[str] token: The token of the project access token. **Note**: the token is not available for imported resources.
        :param pulumi.Input[int] user_id: The user_id associated to the token.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectAccessTokenState.__new__(_ProjectAccessTokenState)

        __props__.__dict__["access_level"] = access_level
        __props__.__dict__["active"] = active
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["revoked"] = revoked
        __props__.__dict__["rotation_configuration"] = rotation_configuration
        __props__.__dict__["scopes"] = scopes
        __props__.__dict__["token"] = token
        __props__.__dict__["user_id"] = user_id
        return ProjectAccessToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> pulumi.Output[str]:
        """
        The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
        """
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[bool]:
        """
        True if the token is active.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Time the token has been created, RFC3339 format.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[str]:
        """
        When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the project access token.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or full path of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def revoked(self) -> pulumi.Output[bool]:
        """
        True if the token is revoked.
        """
        return pulumi.get(self, "revoked")

    @property
    @pulumi.getter(name="rotationConfiguration")
    def rotation_configuration(self) -> pulumi.Output[Optional['outputs.ProjectAccessTokenRotationConfiguration']]:
        """
        The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        """
        return pulumi.get(self, "rotation_configuration")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Sequence[str]]:
        """
        The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        The token of the project access token. **Note**: the token is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[int]:
        """
        The user_id associated to the token.
        """
        return pulumi.get(self, "user_id")

