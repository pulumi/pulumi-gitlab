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
from . import outputs
from ._inputs import *

__all__ = ['GroupServiceAccountAccessTokenArgs', 'GroupServiceAccountAccessToken']

@pulumi.input_type
class GroupServiceAccountAccessTokenArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 scopes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 user_id: pulumi.Input[int],
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rotation_configuration: Optional[pulumi.Input['GroupServiceAccountAccessTokenRotationConfigurationArgs']] = None):
        """
        The set of arguments for constructing a GroupServiceAccountAccessToken resource.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group containing the service account. Must be a top level group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        :param pulumi.Input[int] user_id: The ID of a service account user.
        :param pulumi.Input[str] expires_at: The service account access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
        :param pulumi.Input[str] name: The name of the personal access token.
        :param pulumi.Input['GroupServiceAccountAccessTokenRotationConfigurationArgs'] rotation_configuration: The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        """
        pulumi.set(__self__, "group", group)
        pulumi.set(__self__, "scopes", scopes)
        pulumi.set(__self__, "user_id", user_id)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rotation_configuration is not None:
            pulumi.set(__self__, "rotation_configuration", rotation_configuration)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        The ID or URL-encoded path of the group containing the service account. Must be a top level group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[int]:
        """
        The ID of a service account user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        The service account access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the personal access token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rotationConfiguration")
    def rotation_configuration(self) -> Optional[pulumi.Input['GroupServiceAccountAccessTokenRotationConfigurationArgs']]:
        """
        The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        """
        return pulumi.get(self, "rotation_configuration")

    @rotation_configuration.setter
    def rotation_configuration(self, value: Optional[pulumi.Input['GroupServiceAccountAccessTokenRotationConfigurationArgs']]):
        pulumi.set(self, "rotation_configuration", value)


@pulumi.input_type
class _GroupServiceAccountAccessTokenState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 revoked: Optional[pulumi.Input[bool]] = None,
                 rotation_configuration: Optional[pulumi.Input['GroupServiceAccountAccessTokenRotationConfigurationArgs']] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering GroupServiceAccountAccessToken resources.
        :param pulumi.Input[bool] active: True if the token is active.
        :param pulumi.Input[str] created_at: Time the token has been created, RFC3339 format.
        :param pulumi.Input[str] expires_at: The service account access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group containing the service account. Must be a top level group.
        :param pulumi.Input[str] name: The name of the personal access token.
        :param pulumi.Input[bool] revoked: True if the token is revoked.
        :param pulumi.Input['GroupServiceAccountAccessTokenRotationConfigurationArgs'] rotation_configuration: The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        :param pulumi.Input[str] token: The token of the group service account access token. **Note**: the token is not available for imported resources.
        :param pulumi.Input[int] user_id: The ID of a service account user.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if name is not None:
            pulumi.set(__self__, "name", name)
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
        The service account access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or URL-encoded path of the group containing the service account. Must be a top level group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the personal access token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
    def rotation_configuration(self) -> Optional[pulumi.Input['GroupServiceAccountAccessTokenRotationConfigurationArgs']]:
        """
        The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        """
        return pulumi.get(self, "rotation_configuration")

    @rotation_configuration.setter
    def rotation_configuration(self, value: Optional[pulumi.Input['GroupServiceAccountAccessTokenRotationConfigurationArgs']]):
        pulumi.set(self, "rotation_configuration", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The token of the group service account access token. **Note**: the token is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of a service account user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "user_id", value)


class GroupServiceAccountAccessToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rotation_configuration: Optional[pulumi.Input[Union['GroupServiceAccountAccessTokenRotationConfigurationArgs', 'GroupServiceAccountAccessTokenRotationConfigurationArgsDict']]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The `GroupServiceAccountAccessToken` resource allows to manage the lifecycle of a group service account access token.

        > Use of the `timestamp()` function with expires_at will cause the resource to be re-created with every apply, it's recommended to use `plantimestamp()` or a static value instead.

        > Reading the access token status of a service account requires an admin token or a top-level group owner token on gitlab.com. As a result, this resource will ignore permission errors when attempting to read the token status, and will rely on the values in state instead. This can lead to apply-time failures if the token configured for the provider doesn't have permissions to rotate tokens for the service account.

        > Use `rotation_configuration` to automatically rotate tokens instead of using `timestamp()` as timestamp will cause changes with every plan. `pulumi up` must still be run to rotate the token.

        > Due to a limitation in the API, the `rotation_configuration` is unable to set the new expiry date before GitLab 17.9. Instead, when the resource is created, it will default the expiry date to 7 days in the future. On each subsequent apply, the new expiry will be 7 days from the date of the apply.

        **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/group_service_accounts/#create-a-personal-access-token-for-a-service-account-user)

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_service_account_access_token`. For example:

        terraform

        import {

          to = gitlab_group_service_account_access_token.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        ```sh
        $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken You can import a service account access token using `<resource> <id>`. The
        ```

        `id` is in the form of <group_id>:<service_account_id>:<access_token_id>

        Importing an access token does not import the access token value.

        ```sh
        $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken example 1:2:3
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expires_at: The service account access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group containing the service account. Must be a top level group.
        :param pulumi.Input[str] name: The name of the personal access token.
        :param pulumi.Input[Union['GroupServiceAccountAccessTokenRotationConfigurationArgs', 'GroupServiceAccountAccessTokenRotationConfigurationArgsDict']] rotation_configuration: The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        :param pulumi.Input[int] user_id: The ID of a service account user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupServiceAccountAccessTokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupServiceAccountAccessToken` resource allows to manage the lifecycle of a group service account access token.

        > Use of the `timestamp()` function with expires_at will cause the resource to be re-created with every apply, it's recommended to use `plantimestamp()` or a static value instead.

        > Reading the access token status of a service account requires an admin token or a top-level group owner token on gitlab.com. As a result, this resource will ignore permission errors when attempting to read the token status, and will rely on the values in state instead. This can lead to apply-time failures if the token configured for the provider doesn't have permissions to rotate tokens for the service account.

        > Use `rotation_configuration` to automatically rotate tokens instead of using `timestamp()` as timestamp will cause changes with every plan. `pulumi up` must still be run to rotate the token.

        > Due to a limitation in the API, the `rotation_configuration` is unable to set the new expiry date before GitLab 17.9. Instead, when the resource is created, it will default the expiry date to 7 days in the future. On each subsequent apply, the new expiry will be 7 days from the date of the apply.

        **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/group_service_accounts/#create-a-personal-access-token-for-a-service-account-user)

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_service_account_access_token`. For example:

        terraform

        import {

          to = gitlab_group_service_account_access_token.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        ```sh
        $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken You can import a service account access token using `<resource> <id>`. The
        ```

        `id` is in the form of <group_id>:<service_account_id>:<access_token_id>

        Importing an access token does not import the access token value.

        ```sh
        $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken example 1:2:3
        ```

        :param str resource_name: The name of the resource.
        :param GroupServiceAccountAccessTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupServiceAccountAccessTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rotation_configuration: Optional[pulumi.Input[Union['GroupServiceAccountAccessTokenRotationConfigurationArgs', 'GroupServiceAccountAccessTokenRotationConfigurationArgsDict']]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupServiceAccountAccessTokenArgs.__new__(GroupServiceAccountAccessTokenArgs)

            __props__.__dict__["expires_at"] = expires_at
            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["name"] = name
            __props__.__dict__["rotation_configuration"] = rotation_configuration
            if scopes is None and not opts.urn:
                raise TypeError("Missing required property 'scopes'")
            __props__.__dict__["scopes"] = scopes
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["active"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["revoked"] = None
            __props__.__dict__["token"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(GroupServiceAccountAccessToken, __self__).__init__(
            'gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            group: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            revoked: Optional[pulumi.Input[bool]] = None,
            rotation_configuration: Optional[pulumi.Input[Union['GroupServiceAccountAccessTokenRotationConfigurationArgs', 'GroupServiceAccountAccessTokenRotationConfigurationArgsDict']]] = None,
            scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            token: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[int]] = None) -> 'GroupServiceAccountAccessToken':
        """
        Get an existing GroupServiceAccountAccessToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: True if the token is active.
        :param pulumi.Input[str] created_at: Time the token has been created, RFC3339 format.
        :param pulumi.Input[str] expires_at: The service account access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
        :param pulumi.Input[str] group: The ID or URL-encoded path of the group containing the service account. Must be a top level group.
        :param pulumi.Input[str] name: The name of the personal access token.
        :param pulumi.Input[bool] revoked: True if the token is revoked.
        :param pulumi.Input[Union['GroupServiceAccountAccessTokenRotationConfigurationArgs', 'GroupServiceAccountAccessTokenRotationConfigurationArgsDict']] rotation_configuration: The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        :param pulumi.Input[str] token: The token of the group service account access token. **Note**: the token is not available for imported resources.
        :param pulumi.Input[int] user_id: The ID of a service account user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupServiceAccountAccessTokenState.__new__(_GroupServiceAccountAccessTokenState)

        __props__.__dict__["active"] = active
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["group"] = group
        __props__.__dict__["name"] = name
        __props__.__dict__["revoked"] = revoked
        __props__.__dict__["rotation_configuration"] = rotation_configuration
        __props__.__dict__["scopes"] = scopes
        __props__.__dict__["token"] = token
        __props__.__dict__["user_id"] = user_id
        return GroupServiceAccountAccessToken(resource_name, opts=opts, __props__=__props__)

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
        The service account access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of the group containing the service account. Must be a top level group.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the personal access token.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def revoked(self) -> pulumi.Output[bool]:
        """
        True if the token is revoked.
        """
        return pulumi.get(self, "revoked")

    @property
    @pulumi.getter(name="rotationConfiguration")
    def rotation_configuration(self) -> pulumi.Output[Optional['outputs.GroupServiceAccountAccessTokenRotationConfiguration']]:
        """
        The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        """
        return pulumi.get(self, "rotation_configuration")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Sequence[str]]:
        """
        The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        The token of the group service account access token. **Note**: the token is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[int]:
        """
        The ID of a service account user.
        """
        return pulumi.get(self, "user_id")

