# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GroupClusterArgs', 'GroupCluster']

@pulumi.input_type
class GroupClusterArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 kubernetes_api_url: pulumi.Input[str],
                 kubernetes_token: pulumi.Input[str],
                 domain: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 environment_scope: Optional[pulumi.Input[str]] = None,
                 kubernetes_authorization_type: Optional[pulumi.Input[str]] = None,
                 kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[bool]] = None,
                 management_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GroupCluster resource.
        :param pulumi.Input[str] group: The id of the group to add the cluster to.
        :param pulumi.Input[str] kubernetes_api_url: The URL to access the Kubernetes API.
        :param pulumi.Input[str] kubernetes_token: The token to authenticate against Kubernetes.
        :param pulumi.Input[str] domain: The base domain of the cluster.
        :param pulumi.Input[bool] enabled: Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] environment_scope: The associated environment to the cluster. Defaults to `*`.
        :param pulumi.Input[str] kubernetes_authorization_type: The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        :param pulumi.Input[str] kubernetes_ca_cert: TLS certificate (needed if API is using a self-signed TLS certificate).
        :param pulumi.Input[bool] managed: Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] management_project_id: The ID of the management project for the cluster.
        :param pulumi.Input[str] name: The name of cluster.
        """
        GroupClusterArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            group=group,
            kubernetes_api_url=kubernetes_api_url,
            kubernetes_token=kubernetes_token,
            domain=domain,
            enabled=enabled,
            environment_scope=environment_scope,
            kubernetes_authorization_type=kubernetes_authorization_type,
            kubernetes_ca_cert=kubernetes_ca_cert,
            managed=managed,
            management_project_id=management_project_id,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             group: Optional[pulumi.Input[str]] = None,
             kubernetes_api_url: Optional[pulumi.Input[str]] = None,
             kubernetes_token: Optional[pulumi.Input[str]] = None,
             domain: Optional[pulumi.Input[str]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             environment_scope: Optional[pulumi.Input[str]] = None,
             kubernetes_authorization_type: Optional[pulumi.Input[str]] = None,
             kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
             managed: Optional[pulumi.Input[bool]] = None,
             management_project_id: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if group is None:
            raise TypeError("Missing 'group' argument")
        if kubernetes_api_url is None and 'kubernetesApiUrl' in kwargs:
            kubernetes_api_url = kwargs['kubernetesApiUrl']
        if kubernetes_api_url is None:
            raise TypeError("Missing 'kubernetes_api_url' argument")
        if kubernetes_token is None and 'kubernetesToken' in kwargs:
            kubernetes_token = kwargs['kubernetesToken']
        if kubernetes_token is None:
            raise TypeError("Missing 'kubernetes_token' argument")
        if environment_scope is None and 'environmentScope' in kwargs:
            environment_scope = kwargs['environmentScope']
        if kubernetes_authorization_type is None and 'kubernetesAuthorizationType' in kwargs:
            kubernetes_authorization_type = kwargs['kubernetesAuthorizationType']
        if kubernetes_ca_cert is None and 'kubernetesCaCert' in kwargs:
            kubernetes_ca_cert = kwargs['kubernetesCaCert']
        if management_project_id is None and 'managementProjectId' in kwargs:
            management_project_id = kwargs['managementProjectId']

        _setter("group", group)
        _setter("kubernetes_api_url", kubernetes_api_url)
        _setter("kubernetes_token", kubernetes_token)
        if domain is not None:
            _setter("domain", domain)
        if enabled is not None:
            _setter("enabled", enabled)
        if environment_scope is not None:
            _setter("environment_scope", environment_scope)
        if kubernetes_authorization_type is not None:
            _setter("kubernetes_authorization_type", kubernetes_authorization_type)
        if kubernetes_ca_cert is not None:
            _setter("kubernetes_ca_cert", kubernetes_ca_cert)
        if managed is not None:
            _setter("managed", managed)
        if management_project_id is not None:
            _setter("management_project_id", management_project_id)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        The id of the group to add the cluster to.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="kubernetesApiUrl")
    def kubernetes_api_url(self) -> pulumi.Input[str]:
        """
        The URL to access the Kubernetes API.
        """
        return pulumi.get(self, "kubernetes_api_url")

    @kubernetes_api_url.setter
    def kubernetes_api_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "kubernetes_api_url", value)

    @property
    @pulumi.getter(name="kubernetesToken")
    def kubernetes_token(self) -> pulumi.Input[str]:
        """
        The token to authenticate against Kubernetes.
        """
        return pulumi.get(self, "kubernetes_token")

    @kubernetes_token.setter
    def kubernetes_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "kubernetes_token", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The base domain of the cluster.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="environmentScope")
    def environment_scope(self) -> Optional[pulumi.Input[str]]:
        """
        The associated environment to the cluster. Defaults to `*`.
        """
        return pulumi.get(self, "environment_scope")

    @environment_scope.setter
    def environment_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment_scope", value)

    @property
    @pulumi.getter(name="kubernetesAuthorizationType")
    def kubernetes_authorization_type(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        """
        return pulumi.get(self, "kubernetes_authorization_type")

    @kubernetes_authorization_type.setter
    def kubernetes_authorization_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_authorization_type", value)

    @property
    @pulumi.getter(name="kubernetesCaCert")
    def kubernetes_ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        TLS certificate (needed if API is using a self-signed TLS certificate).
        """
        return pulumi.get(self, "kubernetes_ca_cert")

    @kubernetes_ca_cert.setter
    def kubernetes_ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_ca_cert", value)

    @property
    @pulumi.getter
    def managed(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        """
        return pulumi.get(self, "managed")

    @managed.setter
    def managed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "managed", value)

    @property
    @pulumi.getter(name="managementProjectId")
    def management_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the management project for the cluster.
        """
        return pulumi.get(self, "management_project_id")

    @management_project_id.setter
    def management_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "management_project_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _GroupClusterState:
    def __init__(__self__, *,
                 cluster_type: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 environment_scope: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 kubernetes_api_url: Optional[pulumi.Input[str]] = None,
                 kubernetes_authorization_type: Optional[pulumi.Input[str]] = None,
                 kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
                 kubernetes_token: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[bool]] = None,
                 management_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_type: Optional[pulumi.Input[str]] = None,
                 provider_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupCluster resources.
        :param pulumi.Input[str] cluster_type: Cluster type.
        :param pulumi.Input[str] created_at: Create time.
        :param pulumi.Input[str] domain: The base domain of the cluster.
        :param pulumi.Input[bool] enabled: Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] environment_scope: The associated environment to the cluster. Defaults to `*`.
        :param pulumi.Input[str] group: The id of the group to add the cluster to.
        :param pulumi.Input[str] kubernetes_api_url: The URL to access the Kubernetes API.
        :param pulumi.Input[str] kubernetes_authorization_type: The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        :param pulumi.Input[str] kubernetes_ca_cert: TLS certificate (needed if API is using a self-signed TLS certificate).
        :param pulumi.Input[str] kubernetes_token: The token to authenticate against Kubernetes.
        :param pulumi.Input[bool] managed: Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] management_project_id: The ID of the management project for the cluster.
        :param pulumi.Input[str] name: The name of cluster.
        :param pulumi.Input[str] platform_type: Platform type.
        :param pulumi.Input[str] provider_type: Provider type.
        """
        _GroupClusterState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            cluster_type=cluster_type,
            created_at=created_at,
            domain=domain,
            enabled=enabled,
            environment_scope=environment_scope,
            group=group,
            kubernetes_api_url=kubernetes_api_url,
            kubernetes_authorization_type=kubernetes_authorization_type,
            kubernetes_ca_cert=kubernetes_ca_cert,
            kubernetes_token=kubernetes_token,
            managed=managed,
            management_project_id=management_project_id,
            name=name,
            platform_type=platform_type,
            provider_type=provider_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             cluster_type: Optional[pulumi.Input[str]] = None,
             created_at: Optional[pulumi.Input[str]] = None,
             domain: Optional[pulumi.Input[str]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             environment_scope: Optional[pulumi.Input[str]] = None,
             group: Optional[pulumi.Input[str]] = None,
             kubernetes_api_url: Optional[pulumi.Input[str]] = None,
             kubernetes_authorization_type: Optional[pulumi.Input[str]] = None,
             kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
             kubernetes_token: Optional[pulumi.Input[str]] = None,
             managed: Optional[pulumi.Input[bool]] = None,
             management_project_id: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             platform_type: Optional[pulumi.Input[str]] = None,
             provider_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if cluster_type is None and 'clusterType' in kwargs:
            cluster_type = kwargs['clusterType']
        if created_at is None and 'createdAt' in kwargs:
            created_at = kwargs['createdAt']
        if environment_scope is None and 'environmentScope' in kwargs:
            environment_scope = kwargs['environmentScope']
        if kubernetes_api_url is None and 'kubernetesApiUrl' in kwargs:
            kubernetes_api_url = kwargs['kubernetesApiUrl']
        if kubernetes_authorization_type is None and 'kubernetesAuthorizationType' in kwargs:
            kubernetes_authorization_type = kwargs['kubernetesAuthorizationType']
        if kubernetes_ca_cert is None and 'kubernetesCaCert' in kwargs:
            kubernetes_ca_cert = kwargs['kubernetesCaCert']
        if kubernetes_token is None and 'kubernetesToken' in kwargs:
            kubernetes_token = kwargs['kubernetesToken']
        if management_project_id is None and 'managementProjectId' in kwargs:
            management_project_id = kwargs['managementProjectId']
        if platform_type is None and 'platformType' in kwargs:
            platform_type = kwargs['platformType']
        if provider_type is None and 'providerType' in kwargs:
            provider_type = kwargs['providerType']

        if cluster_type is not None:
            _setter("cluster_type", cluster_type)
        if created_at is not None:
            _setter("created_at", created_at)
        if domain is not None:
            _setter("domain", domain)
        if enabled is not None:
            _setter("enabled", enabled)
        if environment_scope is not None:
            _setter("environment_scope", environment_scope)
        if group is not None:
            _setter("group", group)
        if kubernetes_api_url is not None:
            _setter("kubernetes_api_url", kubernetes_api_url)
        if kubernetes_authorization_type is not None:
            _setter("kubernetes_authorization_type", kubernetes_authorization_type)
        if kubernetes_ca_cert is not None:
            _setter("kubernetes_ca_cert", kubernetes_ca_cert)
        if kubernetes_token is not None:
            _setter("kubernetes_token", kubernetes_token)
        if managed is not None:
            _setter("managed", managed)
        if management_project_id is not None:
            _setter("management_project_id", management_project_id)
        if name is not None:
            _setter("name", name)
        if platform_type is not None:
            _setter("platform_type", platform_type)
        if provider_type is not None:
            _setter("provider_type", provider_type)

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster type.
        """
        return pulumi.get(self, "cluster_type")

    @cluster_type.setter
    def cluster_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_type", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Create time.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The base domain of the cluster.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="environmentScope")
    def environment_scope(self) -> Optional[pulumi.Input[str]]:
        """
        The associated environment to the cluster. Defaults to `*`.
        """
        return pulumi.get(self, "environment_scope")

    @environment_scope.setter
    def environment_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment_scope", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the group to add the cluster to.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="kubernetesApiUrl")
    def kubernetes_api_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL to access the Kubernetes API.
        """
        return pulumi.get(self, "kubernetes_api_url")

    @kubernetes_api_url.setter
    def kubernetes_api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_api_url", value)

    @property
    @pulumi.getter(name="kubernetesAuthorizationType")
    def kubernetes_authorization_type(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        """
        return pulumi.get(self, "kubernetes_authorization_type")

    @kubernetes_authorization_type.setter
    def kubernetes_authorization_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_authorization_type", value)

    @property
    @pulumi.getter(name="kubernetesCaCert")
    def kubernetes_ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        TLS certificate (needed if API is using a self-signed TLS certificate).
        """
        return pulumi.get(self, "kubernetes_ca_cert")

    @kubernetes_ca_cert.setter
    def kubernetes_ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_ca_cert", value)

    @property
    @pulumi.getter(name="kubernetesToken")
    def kubernetes_token(self) -> Optional[pulumi.Input[str]]:
        """
        The token to authenticate against Kubernetes.
        """
        return pulumi.get(self, "kubernetes_token")

    @kubernetes_token.setter
    def kubernetes_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_token", value)

    @property
    @pulumi.getter
    def managed(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        """
        return pulumi.get(self, "managed")

    @managed.setter
    def managed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "managed", value)

    @property
    @pulumi.getter(name="managementProjectId")
    def management_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the management project for the cluster.
        """
        return pulumi.get(self, "management_project_id")

    @management_project_id.setter
    def management_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "management_project_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="platformType")
    def platform_type(self) -> Optional[pulumi.Input[str]]:
        """
        Platform type.
        """
        return pulumi.get(self, "platform_type")

    @platform_type.setter
    def platform_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "platform_type", value)

    @property
    @pulumi.getter(name="providerType")
    def provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        Provider type.
        """
        return pulumi.get(self, "provider_type")

    @provider_type.setter
    def provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_type", value)


class GroupCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 environment_scope: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 kubernetes_api_url: Optional[pulumi.Input[str]] = None,
                 kubernetes_authorization_type: Optional[pulumi.Input[str]] = None,
                 kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
                 kubernetes_token: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[bool]] = None,
                 management_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `GroupCluster` resource allows to manage the lifecycle of a group cluster.

        > This is deprecated GitLab feature since 14.5

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_clusters.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.Group("foo", path="foo-path")
        bar = gitlab.GroupCluster("bar",
            group=foo.id,
            domain="example.com",
            enabled=True,
            kubernetes_api_url="https://124.124.124",
            kubernetes_token="some-token",
            kubernetes_ca_cert="some-cert",
            kubernetes_authorization_type="rbac",
            environment_scope="*",
            management_project_id="123456")
        ```

        ## Import

        GitLab group clusters can be imported using an id made up of `groupid:clusterid`, e.g.

        ```sh
         $ pulumi import gitlab:index/groupCluster:GroupCluster bar 123:321
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: The base domain of the cluster.
        :param pulumi.Input[bool] enabled: Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] environment_scope: The associated environment to the cluster. Defaults to `*`.
        :param pulumi.Input[str] group: The id of the group to add the cluster to.
        :param pulumi.Input[str] kubernetes_api_url: The URL to access the Kubernetes API.
        :param pulumi.Input[str] kubernetes_authorization_type: The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        :param pulumi.Input[str] kubernetes_ca_cert: TLS certificate (needed if API is using a self-signed TLS certificate).
        :param pulumi.Input[str] kubernetes_token: The token to authenticate against Kubernetes.
        :param pulumi.Input[bool] managed: Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] management_project_id: The ID of the management project for the cluster.
        :param pulumi.Input[str] name: The name of cluster.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `GroupCluster` resource allows to manage the lifecycle of a group cluster.

        > This is deprecated GitLab feature since 14.5

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_clusters.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.Group("foo", path="foo-path")
        bar = gitlab.GroupCluster("bar",
            group=foo.id,
            domain="example.com",
            enabled=True,
            kubernetes_api_url="https://124.124.124",
            kubernetes_token="some-token",
            kubernetes_ca_cert="some-cert",
            kubernetes_authorization_type="rbac",
            environment_scope="*",
            management_project_id="123456")
        ```

        ## Import

        GitLab group clusters can be imported using an id made up of `groupid:clusterid`, e.g.

        ```sh
         $ pulumi import gitlab:index/groupCluster:GroupCluster bar 123:321
        ```

        :param str resource_name: The name of the resource.
        :param GroupClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            GroupClusterArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 environment_scope: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 kubernetes_api_url: Optional[pulumi.Input[str]] = None,
                 kubernetes_authorization_type: Optional[pulumi.Input[str]] = None,
                 kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
                 kubernetes_token: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[bool]] = None,
                 management_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupClusterArgs.__new__(GroupClusterArgs)

            __props__.__dict__["domain"] = domain
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["environment_scope"] = environment_scope
            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            if kubernetes_api_url is None and not opts.urn:
                raise TypeError("Missing required property 'kubernetes_api_url'")
            __props__.__dict__["kubernetes_api_url"] = kubernetes_api_url
            __props__.__dict__["kubernetes_authorization_type"] = kubernetes_authorization_type
            __props__.__dict__["kubernetes_ca_cert"] = kubernetes_ca_cert
            if kubernetes_token is None and not opts.urn:
                raise TypeError("Missing required property 'kubernetes_token'")
            __props__.__dict__["kubernetes_token"] = None if kubernetes_token is None else pulumi.Output.secret(kubernetes_token)
            __props__.__dict__["managed"] = managed
            __props__.__dict__["management_project_id"] = management_project_id
            __props__.__dict__["name"] = name
            __props__.__dict__["cluster_type"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["platform_type"] = None
            __props__.__dict__["provider_type"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["kubernetesToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(GroupCluster, __self__).__init__(
            'gitlab:index/groupCluster:GroupCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_type: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            environment_scope: Optional[pulumi.Input[str]] = None,
            group: Optional[pulumi.Input[str]] = None,
            kubernetes_api_url: Optional[pulumi.Input[str]] = None,
            kubernetes_authorization_type: Optional[pulumi.Input[str]] = None,
            kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
            kubernetes_token: Optional[pulumi.Input[str]] = None,
            managed: Optional[pulumi.Input[bool]] = None,
            management_project_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            platform_type: Optional[pulumi.Input[str]] = None,
            provider_type: Optional[pulumi.Input[str]] = None) -> 'GroupCluster':
        """
        Get an existing GroupCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_type: Cluster type.
        :param pulumi.Input[str] created_at: Create time.
        :param pulumi.Input[str] domain: The base domain of the cluster.
        :param pulumi.Input[bool] enabled: Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] environment_scope: The associated environment to the cluster. Defaults to `*`.
        :param pulumi.Input[str] group: The id of the group to add the cluster to.
        :param pulumi.Input[str] kubernetes_api_url: The URL to access the Kubernetes API.
        :param pulumi.Input[str] kubernetes_authorization_type: The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        :param pulumi.Input[str] kubernetes_ca_cert: TLS certificate (needed if API is using a self-signed TLS certificate).
        :param pulumi.Input[str] kubernetes_token: The token to authenticate against Kubernetes.
        :param pulumi.Input[bool] managed: Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] management_project_id: The ID of the management project for the cluster.
        :param pulumi.Input[str] name: The name of cluster.
        :param pulumi.Input[str] platform_type: Platform type.
        :param pulumi.Input[str] provider_type: Provider type.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupClusterState.__new__(_GroupClusterState)

        __props__.__dict__["cluster_type"] = cluster_type
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["domain"] = domain
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["environment_scope"] = environment_scope
        __props__.__dict__["group"] = group
        __props__.__dict__["kubernetes_api_url"] = kubernetes_api_url
        __props__.__dict__["kubernetes_authorization_type"] = kubernetes_authorization_type
        __props__.__dict__["kubernetes_ca_cert"] = kubernetes_ca_cert
        __props__.__dict__["kubernetes_token"] = kubernetes_token
        __props__.__dict__["managed"] = managed
        __props__.__dict__["management_project_id"] = management_project_id
        __props__.__dict__["name"] = name
        __props__.__dict__["platform_type"] = platform_type
        __props__.__dict__["provider_type"] = provider_type
        return GroupCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> pulumi.Output[str]:
        """
        Cluster type.
        """
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Create time.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[Optional[str]]:
        """
        The base domain of the cluster.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="environmentScope")
    def environment_scope(self) -> pulumi.Output[Optional[str]]:
        """
        The associated environment to the cluster. Defaults to `*`.
        """
        return pulumi.get(self, "environment_scope")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The id of the group to add the cluster to.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="kubernetesApiUrl")
    def kubernetes_api_url(self) -> pulumi.Output[str]:
        """
        The URL to access the Kubernetes API.
        """
        return pulumi.get(self, "kubernetes_api_url")

    @property
    @pulumi.getter(name="kubernetesAuthorizationType")
    def kubernetes_authorization_type(self) -> pulumi.Output[Optional[str]]:
        """
        The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        """
        return pulumi.get(self, "kubernetes_authorization_type")

    @property
    @pulumi.getter(name="kubernetesCaCert")
    def kubernetes_ca_cert(self) -> pulumi.Output[Optional[str]]:
        """
        TLS certificate (needed if API is using a self-signed TLS certificate).
        """
        return pulumi.get(self, "kubernetes_ca_cert")

    @property
    @pulumi.getter(name="kubernetesToken")
    def kubernetes_token(self) -> pulumi.Output[str]:
        """
        The token to authenticate against Kubernetes.
        """
        return pulumi.get(self, "kubernetes_token")

    @property
    @pulumi.getter
    def managed(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        """
        return pulumi.get(self, "managed")

    @property
    @pulumi.getter(name="managementProjectId")
    def management_project_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the management project for the cluster.
        """
        return pulumi.get(self, "management_project_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformType")
    def platform_type(self) -> pulumi.Output[str]:
        """
        Platform type.
        """
        return pulumi.get(self, "platform_type")

    @property
    @pulumi.getter(name="providerType")
    def provider_type(self) -> pulumi.Output[str]:
        """
        Provider type.
        """
        return pulumi.get(self, "provider_type")

