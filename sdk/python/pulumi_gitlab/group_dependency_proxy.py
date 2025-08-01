# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['GroupDependencyProxyArgs', 'GroupDependencyProxy']

@pulumi.input_type
class GroupDependencyProxyArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[_builtins.str],
                 enabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 identity: Optional[pulumi.Input[_builtins.str]] = None,
                 secret: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a GroupDependencyProxy resource.
        :param pulumi.Input[_builtins.str] group: The ID or URL-encoded path of the group.
        :param pulumi.Input[_builtins.bool] enabled: Indicates whether the proxy is enabled.
        :param pulumi.Input[_builtins.str] identity: Identity credential used to authenticate with Docker Hub when pulling images. Can be a username (for password or personal access token (PAT)) or organization name (for organization access token (OAT)).
        :param pulumi.Input[_builtins.str] secret: Secret credential used to authenticate with Docker Hub when pulling images. Can be a password, personal access token (PAT), or organization access token (OAT). Cannot be imported.
        """
        pulumi.set(__self__, "group", group)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @_builtins.property
    @pulumi.getter
    def group(self) -> pulumi.Input[_builtins.str]:
        """
        The ID or URL-encoded path of the group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "group", value)

    @_builtins.property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Indicates whether the proxy is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @_builtins.property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Identity credential used to authenticate with Docker Hub when pulling images. Can be a username (for password or personal access token (PAT)) or organization name (for organization access token (OAT)).
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "identity", value)

    @_builtins.property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Secret credential used to authenticate with Docker Hub when pulling images. Can be a password, personal access token (PAT), or organization access token (OAT). Cannot be imported.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "secret", value)


@pulumi.input_type
class _GroupDependencyProxyState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 group: Optional[pulumi.Input[_builtins.str]] = None,
                 identity: Optional[pulumi.Input[_builtins.str]] = None,
                 secret: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering GroupDependencyProxy resources.
        :param pulumi.Input[_builtins.bool] enabled: Indicates whether the proxy is enabled.
        :param pulumi.Input[_builtins.str] group: The ID or URL-encoded path of the group.
        :param pulumi.Input[_builtins.str] identity: Identity credential used to authenticate with Docker Hub when pulling images. Can be a username (for password or personal access token (PAT)) or organization name (for organization access token (OAT)).
        :param pulumi.Input[_builtins.str] secret: Secret credential used to authenticate with Docker Hub when pulling images. Can be a password, personal access token (PAT), or organization access token (OAT). Cannot be imported.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @_builtins.property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Indicates whether the proxy is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @_builtins.property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID or URL-encoded path of the group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "group", value)

    @_builtins.property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Identity credential used to authenticate with Docker Hub when pulling images. Can be a username (for password or personal access token (PAT)) or organization name (for organization access token (OAT)).
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "identity", value)

    @_builtins.property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Secret credential used to authenticate with Docker Hub when pulling images. Can be a password, personal access token (PAT), or organization access token (OAT). Cannot be imported.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "secret", value)


@pulumi.type_token("gitlab:index/groupDependencyProxy:GroupDependencyProxy")
class GroupDependencyProxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 group: Optional[pulumi.Input[_builtins.str]] = None,
                 identity: Optional[pulumi.Input[_builtins.str]] = None,
                 secret: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        The `GroupDependencyProxy` resource allows managing the group docker dependency proxy. More than one dependency proxy per group will conflict with each other.

        If you're looking to manage the project-level package dependency proxy, see the `gitlab_project_package_registry_proxy` resource instead.

        **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/api/graphql/reference/#mutationupdatedependencyproxysettings)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.GroupDependencyProxy("foo",
            group="1234",
            enabled=True,
            identity="newidentity",
            secret="somesecret")
        ```

        ## Import

        Starting in Terraform v1.5.0, you can use an import block to import `gitlab_group_dependency_proxy`. For example:

        terraform

        import {

          to = gitlab_group_dependency_proxy.example

          id = "see CLI command below for ID"

        }

        Importing using the CLI is supported with the following syntax:

        You can import a group dependency proxy using the group id. e.g. `{group-id}`

        "secret" will not populate when importing the dependency proxy, but will still

        be required in the configuration.

        ```sh
        $ pulumi import gitlab:index/groupDependencyProxy:GroupDependencyProxy foo 42
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.bool] enabled: Indicates whether the proxy is enabled.
        :param pulumi.Input[_builtins.str] group: The ID or URL-encoded path of the group.
        :param pulumi.Input[_builtins.str] identity: Identity credential used to authenticate with Docker Hub when pulling images. Can be a username (for password or personal access token (PAT)) or organization name (for organization access token (OAT)).
        :param pulumi.Input[_builtins.str] secret: Secret credential used to authenticate with Docker Hub when pulling images. Can be a password, personal access token (PAT), or organization access token (OAT). Cannot be imported.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupDependencyProxyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupDependencyProxy` resource allows managing the group docker dependency proxy. More than one dependency proxy per group will conflict with each other.

        If you're looking to manage the project-level package dependency proxy, see the `gitlab_project_package_registry_proxy` resource instead.

        **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/api/graphql/reference/#mutationupdatedependencyproxysettings)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.GroupDependencyProxy("foo",
            group="1234",
            enabled=True,
            identity="newidentity",
            secret="somesecret")
        ```

        ## Import

        Starting in Terraform v1.5.0, you can use an import block to import `gitlab_group_dependency_proxy`. For example:

        terraform

        import {

          to = gitlab_group_dependency_proxy.example

          id = "see CLI command below for ID"

        }

        Importing using the CLI is supported with the following syntax:

        You can import a group dependency proxy using the group id. e.g. `{group-id}`

        "secret" will not populate when importing the dependency proxy, but will still

        be required in the configuration.

        ```sh
        $ pulumi import gitlab:index/groupDependencyProxy:GroupDependencyProxy foo 42
        ```

        :param str resource_name: The name of the resource.
        :param GroupDependencyProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupDependencyProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 group: Optional[pulumi.Input[_builtins.str]] = None,
                 identity: Optional[pulumi.Input[_builtins.str]] = None,
                 secret: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupDependencyProxyArgs.__new__(GroupDependencyProxyArgs)

            __props__.__dict__["enabled"] = enabled
            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["identity"] = identity
            __props__.__dict__["secret"] = None if secret is None else pulumi.Output.secret(secret)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(GroupDependencyProxy, __self__).__init__(
            'gitlab:index/groupDependencyProxy:GroupDependencyProxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[_builtins.bool]] = None,
            group: Optional[pulumi.Input[_builtins.str]] = None,
            identity: Optional[pulumi.Input[_builtins.str]] = None,
            secret: Optional[pulumi.Input[_builtins.str]] = None) -> 'GroupDependencyProxy':
        """
        Get an existing GroupDependencyProxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.bool] enabled: Indicates whether the proxy is enabled.
        :param pulumi.Input[_builtins.str] group: The ID or URL-encoded path of the group.
        :param pulumi.Input[_builtins.str] identity: Identity credential used to authenticate with Docker Hub when pulling images. Can be a username (for password or personal access token (PAT)) or organization name (for organization access token (OAT)).
        :param pulumi.Input[_builtins.str] secret: Secret credential used to authenticate with Docker Hub when pulling images. Can be a password, personal access token (PAT), or organization access token (OAT). Cannot be imported.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupDependencyProxyState.__new__(_GroupDependencyProxyState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["group"] = group
        __props__.__dict__["identity"] = identity
        __props__.__dict__["secret"] = secret
        return GroupDependencyProxy(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[_builtins.bool]:
        """
        Indicates whether the proxy is enabled.
        """
        return pulumi.get(self, "enabled")

    @_builtins.property
    @pulumi.getter
    def group(self) -> pulumi.Output[_builtins.str]:
        """
        The ID or URL-encoded path of the group.
        """
        return pulumi.get(self, "group")

    @_builtins.property
    @pulumi.getter
    def identity(self) -> pulumi.Output[_builtins.str]:
        """
        Identity credential used to authenticate with Docker Hub when pulling images. Can be a username (for password or personal access token (PAT)) or organization name (for organization access token (OAT)).
        """
        return pulumi.get(self, "identity")

    @_builtins.property
    @pulumi.getter
    def secret(self) -> pulumi.Output[_builtins.str]:
        """
        Secret credential used to authenticate with Docker Hub when pulling images. Can be a password, personal access token (PAT), or organization access token (OAT). Cannot be imported.
        """
        return pulumi.get(self, "secret")

