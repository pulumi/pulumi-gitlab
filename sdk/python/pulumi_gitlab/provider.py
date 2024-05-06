# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 base_url: Optional[pulumi.Input[str]] = None,
                 cacert_file: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 early_auth_check: Optional[pulumi.Input[bool]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] base_url: This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
               Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
               the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
        :param pulumi.Input[str] cacert_file: This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
               CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
        :param pulumi.Input[str] client_cert: File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
        :param pulumi.Input[str] client_key: File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
               `client_cert` is set.
        :param pulumi.Input[bool] insecure: When set to true this disables SSL verification of the connection to the GitLab instance.
        :param pulumi.Input[str] token: The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
               used in this provider for authentication (using Bearer authorization token). See
               https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
               variable.
        """
        if base_url is not None:
            pulumi.set(__self__, "base_url", base_url)
        if cacert_file is not None:
            pulumi.set(__self__, "cacert_file", cacert_file)
        if client_cert is not None:
            pulumi.set(__self__, "client_cert", client_cert)
        if client_key is not None:
            pulumi.set(__self__, "client_key", client_key)
        if early_auth_check is not None:
            pulumi.set(__self__, "early_auth_check", early_auth_check)
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> Optional[pulumi.Input[str]]:
        """
        This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
        Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
        the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
        """
        return pulumi.get(self, "base_url")

    @base_url.setter
    def base_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_url", value)

    @property
    @pulumi.getter(name="cacertFile")
    def cacert_file(self) -> Optional[pulumi.Input[str]]:
        """
        This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
        CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
        """
        return pulumi.get(self, "cacert_file")

    @cacert_file.setter
    def cacert_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cacert_file", value)

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> Optional[pulumi.Input[str]]:
        """
        File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
        """
        return pulumi.get(self, "client_cert")

    @client_cert.setter
    def client_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_cert", value)

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> Optional[pulumi.Input[str]]:
        """
        File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
        `client_cert` is set.
        """
        return pulumi.get(self, "client_key")

    @client_key.setter
    def client_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_key", value)

    @property
    @pulumi.getter(name="earlyAuthCheck")
    def early_auth_check(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "early_auth_check")

    @early_auth_check.setter
    def early_auth_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "early_auth_check", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to true this disables SSL verification of the connection to the GitLab instance.
        """
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
        used in this provider for authentication (using Bearer authorization token). See
        https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
        variable.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_url: Optional[pulumi.Input[str]] = None,
                 cacert_file: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 early_auth_check: Optional[pulumi.Input[bool]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the gitlab package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] base_url: This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
               Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
               the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
        :param pulumi.Input[str] cacert_file: This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
               CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
        :param pulumi.Input[str] client_cert: File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
        :param pulumi.Input[str] client_key: File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
               `client_cert` is set.
        :param pulumi.Input[bool] insecure: When set to true this disables SSL verification of the connection to the GitLab instance.
        :param pulumi.Input[str] token: The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
               used in this provider for authentication (using Bearer authorization token). See
               https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
               variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the gitlab package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_url: Optional[pulumi.Input[str]] = None,
                 cacert_file: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 early_auth_check: Optional[pulumi.Input[bool]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["base_url"] = base_url
            __props__.__dict__["cacert_file"] = cacert_file
            __props__.__dict__["client_cert"] = client_cert
            __props__.__dict__["client_key"] = client_key
            __props__.__dict__["early_auth_check"] = pulumi.Output.from_input(early_auth_check).apply(pulumi.runtime.to_json) if early_auth_check is not None else None
            __props__.__dict__["insecure"] = pulumi.Output.from_input(insecure).apply(pulumi.runtime.to_json) if insecure is not None else None
            __props__.__dict__["token"] = None if token is None else pulumi.Output.secret(token)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Provider, __self__).__init__(
            'gitlab',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> pulumi.Output[Optional[str]]:
        """
        This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
        Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
        the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
        """
        return pulumi.get(self, "base_url")

    @property
    @pulumi.getter(name="cacertFile")
    def cacert_file(self) -> pulumi.Output[Optional[str]]:
        """
        This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
        CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
        """
        return pulumi.get(self, "cacert_file")

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> pulumi.Output[Optional[str]]:
        """
        File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
        """
        return pulumi.get(self, "client_cert")

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> pulumi.Output[Optional[str]]:
        """
        File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
        `client_cert` is set.
        """
        return pulumi.get(self, "client_key")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[Optional[str]]:
        """
        The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
        used in this provider for authentication (using Bearer authorization token). See
        https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
        variable.
        """
        return pulumi.get(self, "token")

