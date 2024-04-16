# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PagesDomainArgs', 'PagesDomain']

@pulumi.input_type
class PagesDomainArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 project: pulumi.Input[str],
                 auto_ssl_enabled: Optional[pulumi.Input[bool]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 expired: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PagesDomain resource.
        :param pulumi.Input[str] domain: The custom domain indicated by the user.
        :param pulumi.Input[str] project: The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
        :param pulumi.Input[bool] auto_ssl_enabled: Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
        :param pulumi.Input[str] certificate: The certificate in PEM format with intermediates following in most specific to least specific order.
        :param pulumi.Input[bool] expired: Whether the certificate is expired.
        :param pulumi.Input[str] key: The certificate key in PEM format.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "project", project)
        if auto_ssl_enabled is not None:
            pulumi.set(__self__, "auto_ssl_enabled", auto_ssl_enabled)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if expired is not None:
            pulumi.set(__self__, "expired", expired)
        if key is not None:
            pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The custom domain indicated by the user.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="autoSslEnabled")
    def auto_ssl_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
        """
        return pulumi.get(self, "auto_ssl_enabled")

    @auto_ssl_enabled.setter
    def auto_ssl_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_ssl_enabled", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate in PEM format with intermediates following in most specific to least specific order.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def expired(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the certificate is expired.
        """
        return pulumi.get(self, "expired")

    @expired.setter
    def expired(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "expired", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate key in PEM format.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)


@pulumi.input_type
class _PagesDomainState:
    def __init__(__self__, *,
                 auto_ssl_enabled: Optional[pulumi.Input[bool]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 expired: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 verification_code: Optional[pulumi.Input[str]] = None,
                 verified: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering PagesDomain resources.
        :param pulumi.Input[bool] auto_ssl_enabled: Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
        :param pulumi.Input[str] certificate: The certificate in PEM format with intermediates following in most specific to least specific order.
        :param pulumi.Input[str] domain: The custom domain indicated by the user.
        :param pulumi.Input[bool] expired: Whether the certificate is expired.
        :param pulumi.Input[str] key: The certificate key in PEM format.
        :param pulumi.Input[str] project: The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
        :param pulumi.Input[str] url: The URL for the given domain.
        :param pulumi.Input[str] verification_code: The verification code for the domain.
        :param pulumi.Input[bool] verified: The certificate data.
        """
        if auto_ssl_enabled is not None:
            pulumi.set(__self__, "auto_ssl_enabled", auto_ssl_enabled)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if expired is not None:
            pulumi.set(__self__, "expired", expired)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if verification_code is not None:
            pulumi.set(__self__, "verification_code", verification_code)
        if verified is not None:
            pulumi.set(__self__, "verified", verified)

    @property
    @pulumi.getter(name="autoSslEnabled")
    def auto_ssl_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
        """
        return pulumi.get(self, "auto_ssl_enabled")

    @auto_ssl_enabled.setter
    def auto_ssl_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_ssl_enabled", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate in PEM format with intermediates following in most specific to least specific order.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The custom domain indicated by the user.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def expired(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the certificate is expired.
        """
        return pulumi.get(self, "expired")

    @expired.setter
    def expired(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "expired", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate key in PEM format.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL for the given domain.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="verificationCode")
    def verification_code(self) -> Optional[pulumi.Input[str]]:
        """
        The verification code for the domain.
        """
        return pulumi.get(self, "verification_code")

    @verification_code.setter
    def verification_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "verification_code", value)

    @property
    @pulumi.getter
    def verified(self) -> Optional[pulumi.Input[bool]]:
        """
        The certificate data.
        """
        return pulumi.get(self, "verified")

    @verified.setter
    def verified(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verified", value)


class PagesDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_ssl_enabled: Optional[pulumi.Input[bool]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 expired: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `PagesDomain` resource allows connecting custom domains and TLS certificates in GitLab Pages.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pages_domains.html)

        ## Import

        GitLab pages domain can be imported using an id made up of `projectId:domain` _without_ the http protocol, e.g.

        ```sh
        $ pulumi import gitlab:index/pagesDomain:PagesDomain this 123:example.com
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_ssl_enabled: Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
        :param pulumi.Input[str] certificate: The certificate in PEM format with intermediates following in most specific to least specific order.
        :param pulumi.Input[str] domain: The custom domain indicated by the user.
        :param pulumi.Input[bool] expired: Whether the certificate is expired.
        :param pulumi.Input[str] key: The certificate key in PEM format.
        :param pulumi.Input[str] project: The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PagesDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `PagesDomain` resource allows connecting custom domains and TLS certificates in GitLab Pages.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pages_domains.html)

        ## Import

        GitLab pages domain can be imported using an id made up of `projectId:domain` _without_ the http protocol, e.g.

        ```sh
        $ pulumi import gitlab:index/pagesDomain:PagesDomain this 123:example.com
        ```

        :param str resource_name: The name of the resource.
        :param PagesDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PagesDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_ssl_enabled: Optional[pulumi.Input[bool]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 expired: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PagesDomainArgs.__new__(PagesDomainArgs)

            __props__.__dict__["auto_ssl_enabled"] = auto_ssl_enabled
            __props__.__dict__["certificate"] = certificate
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["expired"] = expired
            __props__.__dict__["key"] = key
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["url"] = None
            __props__.__dict__["verification_code"] = None
            __props__.__dict__["verified"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["verificationCode"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(PagesDomain, __self__).__init__(
            'gitlab:index/pagesDomain:PagesDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_ssl_enabled: Optional[pulumi.Input[bool]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            expired: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            verification_code: Optional[pulumi.Input[str]] = None,
            verified: Optional[pulumi.Input[bool]] = None) -> 'PagesDomain':
        """
        Get an existing PagesDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_ssl_enabled: Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
        :param pulumi.Input[str] certificate: The certificate in PEM format with intermediates following in most specific to least specific order.
        :param pulumi.Input[str] domain: The custom domain indicated by the user.
        :param pulumi.Input[bool] expired: Whether the certificate is expired.
        :param pulumi.Input[str] key: The certificate key in PEM format.
        :param pulumi.Input[str] project: The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
        :param pulumi.Input[str] url: The URL for the given domain.
        :param pulumi.Input[str] verification_code: The verification code for the domain.
        :param pulumi.Input[bool] verified: The certificate data.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PagesDomainState.__new__(_PagesDomainState)

        __props__.__dict__["auto_ssl_enabled"] = auto_ssl_enabled
        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["domain"] = domain
        __props__.__dict__["expired"] = expired
        __props__.__dict__["key"] = key
        __props__.__dict__["project"] = project
        __props__.__dict__["url"] = url
        __props__.__dict__["verification_code"] = verification_code
        __props__.__dict__["verified"] = verified
        return PagesDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoSslEnabled")
    def auto_ssl_enabled(self) -> pulumi.Output[bool]:
        """
        Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
        """
        return pulumi.get(self, "auto_ssl_enabled")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[str]:
        """
        The certificate in PEM format with intermediates following in most specific to least specific order.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The custom domain indicated by the user.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def expired(self) -> pulumi.Output[bool]:
        """
        Whether the certificate is expired.
        """
        return pulumi.get(self, "expired")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[Optional[str]]:
        """
        The certificate key in PEM format.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL for the given domain.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="verificationCode")
    def verification_code(self) -> pulumi.Output[str]:
        """
        The verification code for the domain.
        """
        return pulumi.get(self, "verification_code")

    @property
    @pulumi.getter
    def verified(self) -> pulumi.Output[bool]:
        """
        The certificate data.
        """
        return pulumi.get(self, "verified")

