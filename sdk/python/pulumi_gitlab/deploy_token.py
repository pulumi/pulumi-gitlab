# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['DeployTokenArgs', 'DeployToken']

@pulumi.input_type
class DeployTokenArgs:
    def __init__(__self__, *,
                 scopes: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 expires_at: Optional[pulumi.Input[builtins.str]] = None,
                 group: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DeployToken resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] scopes: Valid values: `read_repository`, `read_registry`, `read_package_registry`, `write_registry`, `write_package_registry`.
        :param pulumi.Input[builtins.str] expires_at: Time the token will expire it, RFC3339 format. Will not expire per default.
        :param pulumi.Input[builtins.str] group: The name or id of the group to add the deploy token to.
        :param pulumi.Input[builtins.str] name: A name to describe the deploy token with.
        :param pulumi.Input[builtins.str] project: The name or id of the project to add the deploy token to.
        :param pulumi.Input[builtins.str] username: A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
        """
        pulumi.set(__self__, "scopes", scopes)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        Valid values: `read_repository`, `read_registry`, `read_package_registry`, `write_registry`, `write_package_registry`.
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Time the token will expire it, RFC3339 format. Will not expire per default.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name or id of the group to add the deploy token to.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A name to describe the deploy token with.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name or id of the project to add the deploy token to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _DeployTokenState:
    def __init__(__self__, *,
                 deploy_token_id: Optional[pulumi.Input[builtins.int]] = None,
                 expires_at: Optional[pulumi.Input[builtins.str]] = None,
                 group: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 token: Optional[pulumi.Input[builtins.str]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DeployToken resources.
        :param pulumi.Input[builtins.int] deploy_token_id: The id of the deploy token.
        :param pulumi.Input[builtins.str] expires_at: Time the token will expire it, RFC3339 format. Will not expire per default.
        :param pulumi.Input[builtins.str] group: The name or id of the group to add the deploy token to.
        :param pulumi.Input[builtins.str] name: A name to describe the deploy token with.
        :param pulumi.Input[builtins.str] project: The name or id of the project to add the deploy token to.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] scopes: Valid values: `read_repository`, `read_registry`, `read_package_registry`, `write_registry`, `write_package_registry`.
        :param pulumi.Input[builtins.str] token: The secret token. This is only populated when creating a new deploy token. **Note**: The token is not available for imported resources.
        :param pulumi.Input[builtins.str] username: A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
        """
        if deploy_token_id is not None:
            pulumi.set(__self__, "deploy_token_id", deploy_token_id)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="deployTokenId")
    def deploy_token_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The id of the deploy token.
        """
        return pulumi.get(self, "deploy_token_id")

    @deploy_token_id.setter
    def deploy_token_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "deploy_token_id", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Time the token will expire it, RFC3339 format. Will not expire per default.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name or id of the group to add the deploy token to.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A name to describe the deploy token with.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name or id of the project to add the deploy token to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Valid values: `read_repository`, `read_registry`, `read_package_registry`, `write_registry`, `write_package_registry`.
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The secret token. This is only populated when creating a new deploy token. **Note**: The token is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "username", value)


class DeployToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expires_at: Optional[pulumi.Input[builtins.str]] = None,
                 group: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `DeployToken` resource allows to manage the lifecycle of group and project deploy tokens.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/deploy_tokens/)

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_deploy_token`. For example:

        terraform

        import {

          to = gitlab_deploy_token.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        GitLab deploy tokens can be imported using an id made up of `{type}:{type_id}:{deploy_token_id}`, where type is one of: project, group.

        ```sh
        $ pulumi import gitlab:index/deployToken:DeployToken group_token group:1:3
        ```

        ```sh
        $ pulumi import gitlab:index/deployToken:DeployToken project_token project:1:4
        ```

        Note: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] expires_at: Time the token will expire it, RFC3339 format. Will not expire per default.
        :param pulumi.Input[builtins.str] group: The name or id of the group to add the deploy token to.
        :param pulumi.Input[builtins.str] name: A name to describe the deploy token with.
        :param pulumi.Input[builtins.str] project: The name or id of the project to add the deploy token to.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] scopes: Valid values: `read_repository`, `read_registry`, `read_package_registry`, `write_registry`, `write_package_registry`.
        :param pulumi.Input[builtins.str] username: A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeployTokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `DeployToken` resource allows to manage the lifecycle of group and project deploy tokens.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/deploy_tokens/)

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_deploy_token`. For example:

        terraform

        import {

          to = gitlab_deploy_token.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        GitLab deploy tokens can be imported using an id made up of `{type}:{type_id}:{deploy_token_id}`, where type is one of: project, group.

        ```sh
        $ pulumi import gitlab:index/deployToken:DeployToken group_token group:1:3
        ```

        ```sh
        $ pulumi import gitlab:index/deployToken:DeployToken project_token project:1:4
        ```

        Note: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.

        :param str resource_name: The name of the resource.
        :param DeployTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeployTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expires_at: Optional[pulumi.Input[builtins.str]] = None,
                 group: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeployTokenArgs.__new__(DeployTokenArgs)

            __props__.__dict__["expires_at"] = expires_at
            __props__.__dict__["group"] = group
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if scopes is None and not opts.urn:
                raise TypeError("Missing required property 'scopes'")
            __props__.__dict__["scopes"] = scopes
            __props__.__dict__["username"] = username
            __props__.__dict__["deploy_token_id"] = None
            __props__.__dict__["token"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(DeployToken, __self__).__init__(
            'gitlab:index/deployToken:DeployToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deploy_token_id: Optional[pulumi.Input[builtins.int]] = None,
            expires_at: Optional[pulumi.Input[builtins.str]] = None,
            group: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            project: Optional[pulumi.Input[builtins.str]] = None,
            scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            token: Optional[pulumi.Input[builtins.str]] = None,
            username: Optional[pulumi.Input[builtins.str]] = None) -> 'DeployToken':
        """
        Get an existing DeployToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] deploy_token_id: The id of the deploy token.
        :param pulumi.Input[builtins.str] expires_at: Time the token will expire it, RFC3339 format. Will not expire per default.
        :param pulumi.Input[builtins.str] group: The name or id of the group to add the deploy token to.
        :param pulumi.Input[builtins.str] name: A name to describe the deploy token with.
        :param pulumi.Input[builtins.str] project: The name or id of the project to add the deploy token to.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] scopes: Valid values: `read_repository`, `read_registry`, `read_package_registry`, `write_registry`, `write_package_registry`.
        :param pulumi.Input[builtins.str] token: The secret token. This is only populated when creating a new deploy token. **Note**: The token is not available for imported resources.
        :param pulumi.Input[builtins.str] username: A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeployTokenState.__new__(_DeployTokenState)

        __props__.__dict__["deploy_token_id"] = deploy_token_id
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["group"] = group
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["scopes"] = scopes
        __props__.__dict__["token"] = token
        __props__.__dict__["username"] = username
        return DeployToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deployTokenId")
    def deploy_token_id(self) -> pulumi.Output[builtins.int]:
        """
        The id of the deploy token.
        """
        return pulumi.get(self, "deploy_token_id")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Time the token will expire it, RFC3339 format. Will not expire per default.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name or id of the group to add the deploy token to.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        A name to describe the deploy token with.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name or id of the project to add the deploy token to.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        Valid values: `read_repository`, `read_registry`, `read_package_registry`, `write_registry`, `write_package_registry`.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[builtins.str]:
        """
        The secret token. This is only populated when creating a new deploy token. **Note**: The token is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[builtins.str]:
        """
        A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
        """
        return pulumi.get(self, "username")

