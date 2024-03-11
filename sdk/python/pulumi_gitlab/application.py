# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 redirect_url: pulumi.Input[str],
                 scopes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 confidential: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Application resource.
        :param pulumi.Input[str] redirect_url: The URL gitlab should send the user to after authentication.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `read_api`, `read_user`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `openid`, `profile`, `email`.
               This is only populated when creating a new application. This attribute is not available for imported resources
        :param pulumi.Input[bool] confidential: The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
        :param pulumi.Input[str] name: Name of the application.
        """
        pulumi.set(__self__, "redirect_url", redirect_url)
        pulumi.set(__self__, "scopes", scopes)
        if confidential is not None:
            pulumi.set(__self__, "confidential", confidential)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="redirectUrl")
    def redirect_url(self) -> pulumi.Input[str]:
        """
        The URL gitlab should send the user to after authentication.
        """
        return pulumi.get(self, "redirect_url")

    @redirect_url.setter
    def redirect_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "redirect_url", value)

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `read_api`, `read_user`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `openid`, `profile`, `email`.
        This is only populated when creating a new application. This attribute is not available for imported resources
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter
    def confidential(self) -> Optional[pulumi.Input[bool]]:
        """
        The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
        """
        return pulumi.get(self, "confidential")

    @confidential.setter
    def confidential(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "confidential", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the application.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ApplicationState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 confidential: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 redirect_url: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Application resources.
        :param pulumi.Input[str] application_id: Internal name of the application.
        :param pulumi.Input[bool] confidential: The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
        :param pulumi.Input[str] name: Name of the application.
        :param pulumi.Input[str] redirect_url: The URL gitlab should send the user to after authentication.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `read_api`, `read_user`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `openid`, `profile`, `email`.
               This is only populated when creating a new application. This attribute is not available for imported resources
        :param pulumi.Input[str] secret: Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if confidential is not None:
            pulumi.set(__self__, "confidential", confidential)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if redirect_url is not None:
            pulumi.set(__self__, "redirect_url", redirect_url)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        Internal name of the application.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def confidential(self) -> Optional[pulumi.Input[bool]]:
        """
        The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
        """
        return pulumi.get(self, "confidential")

    @confidential.setter
    def confidential(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "confidential", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the application.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="redirectUrl")
    def redirect_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL gitlab should send the user to after authentication.
        """
        return pulumi.get(self, "redirect_url")

    @redirect_url.setter
    def redirect_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redirect_url", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `read_api`, `read_user`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `openid`, `profile`, `email`.
        This is only populated when creating a new application. This attribute is not available for imported resources
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


class Application(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 confidential: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 redirect_url: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        The `Application` resource allows to manage the lifecycle of applications in gitlab.

        > In order to use a user for a user to create an application, they must have admin privileges at the instance level.
        To create an OIDC application, a scope of "openid".

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/applications.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        oidc = gitlab.Application("oidc",
            confidential=True,
            redirect_url="https://mycompany.com",
            scopes=["openid"])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Gitlab applications can be imported with their id, e.g.

        ```sh
        $ pulumi import gitlab:index/application:Application example "1"
        ```

        NOTE: the secret and scopes cannot be imported

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] confidential: The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
        :param pulumi.Input[str] name: Name of the application.
        :param pulumi.Input[str] redirect_url: The URL gitlab should send the user to after authentication.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `read_api`, `read_user`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `openid`, `profile`, `email`.
               This is only populated when creating a new application. This attribute is not available for imported resources
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `Application` resource allows to manage the lifecycle of applications in gitlab.

        > In order to use a user for a user to create an application, they must have admin privileges at the instance level.
        To create an OIDC application, a scope of "openid".

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/applications.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        oidc = gitlab.Application("oidc",
            confidential=True,
            redirect_url="https://mycompany.com",
            scopes=["openid"])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Gitlab applications can be imported with their id, e.g.

        ```sh
        $ pulumi import gitlab:index/application:Application example "1"
        ```

        NOTE: the secret and scopes cannot be imported

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 confidential: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 redirect_url: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationArgs.__new__(ApplicationArgs)

            __props__.__dict__["confidential"] = confidential
            __props__.__dict__["name"] = name
            if redirect_url is None and not opts.urn:
                raise TypeError("Missing required property 'redirect_url'")
            __props__.__dict__["redirect_url"] = redirect_url
            if scopes is None and not opts.urn:
                raise TypeError("Missing required property 'scopes'")
            __props__.__dict__["scopes"] = scopes
            __props__.__dict__["application_id"] = None
            __props__.__dict__["secret"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Application, __self__).__init__(
            'gitlab:index/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            confidential: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            redirect_url: Optional[pulumi.Input[str]] = None,
            scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            secret: Optional[pulumi.Input[str]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: Internal name of the application.
        :param pulumi.Input[bool] confidential: The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
        :param pulumi.Input[str] name: Name of the application.
        :param pulumi.Input[str] redirect_url: The URL gitlab should send the user to after authentication.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `read_api`, `read_user`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `openid`, `profile`, `email`.
               This is only populated when creating a new application. This attribute is not available for imported resources
        :param pulumi.Input[str] secret: Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationState.__new__(_ApplicationState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["confidential"] = confidential
        __props__.__dict__["name"] = name
        __props__.__dict__["redirect_url"] = redirect_url
        __props__.__dict__["scopes"] = scopes
        __props__.__dict__["secret"] = secret
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        Internal name of the application.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def confidential(self) -> pulumi.Output[bool]:
        """
        The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
        """
        return pulumi.get(self, "confidential")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the application.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="redirectUrl")
    def redirect_url(self) -> pulumi.Output[str]:
        """
        The URL gitlab should send the user to after authentication.
        """
        return pulumi.get(self, "redirect_url")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Sequence[str]]:
        """
        Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `read_api`, `read_user`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `openid`, `profile`, `email`.
        This is only populated when creating a new application. This attribute is not available for imported resources
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        """
        Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "secret")

