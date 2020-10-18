# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['ProjectCluster']


class ProjectCluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 environment_scope: Optional[pulumi.Input[str]] = None,
                 kubernetes_api_url: Optional[pulumi.Input[str]] = None,
                 kubernetes_authorization_type: Optional[pulumi.Input[str]] = None,
                 kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
                 kubernetes_namespace: Optional[pulumi.Input[str]] = None,
                 kubernetes_token: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[bool]] = None,
                 management_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## # gitlab\_project\_cluster

        This resource allows you to create and manage project clusters for your GitLab projects.
        For further information on clusters, consult the [gitlab
        documentation](https://docs.gitlab.com/ce/user/project/clusters/index.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.Project("foo")
        bar = gitlab.ProjectCluster("bar",
            domain="example.com",
            enabled=True,
            environment_scope="*",
            kubernetes_api_url="https://124.124.124",
            kubernetes_authorization_type="rbac",
            kubernetes_ca_cert="some-cert",
            kubernetes_namespace="namespace",
            kubernetes_token="some-token",
            management_project_id="123456",
            project=foo.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: The base domain of the cluster.
        :param pulumi.Input[bool] enabled: Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] environment_scope: The associated environment to the cluster. Defaults to `*`.
        :param pulumi.Input[str] kubernetes_api_url: The URL to access the Kubernetes API.
        :param pulumi.Input[str] kubernetes_authorization_type: The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        :param pulumi.Input[str] kubernetes_ca_cert: TLS certificate (needed if API is using a self-signed TLS certificate).
        :param pulumi.Input[str] kubernetes_namespace: The unique namespace related to the project.
        :param pulumi.Input[str] kubernetes_token: The token to authenticate against Kubernetes.
        :param pulumi.Input[bool] managed: Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] management_project_id: The ID of the management project for the cluster.
        :param pulumi.Input[str] name: The name of cluster.
        :param pulumi.Input[str] project: The id of the project to add the cluster to.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['domain'] = domain
            __props__['enabled'] = enabled
            __props__['environment_scope'] = environment_scope
            if kubernetes_api_url is None:
                raise TypeError("Missing required property 'kubernetes_api_url'")
            __props__['kubernetes_api_url'] = kubernetes_api_url
            __props__['kubernetes_authorization_type'] = kubernetes_authorization_type
            __props__['kubernetes_ca_cert'] = kubernetes_ca_cert
            __props__['kubernetes_namespace'] = kubernetes_namespace
            if kubernetes_token is None:
                raise TypeError("Missing required property 'kubernetes_token'")
            __props__['kubernetes_token'] = kubernetes_token
            __props__['managed'] = managed
            __props__['management_project_id'] = management_project_id
            __props__['name'] = name
            if project is None:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            __props__['cluster_type'] = None
            __props__['created_at'] = None
            __props__['platform_type'] = None
            __props__['provider_type'] = None
        super(ProjectCluster, __self__).__init__(
            'gitlab:index/projectCluster:ProjectCluster',
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
            kubernetes_api_url: Optional[pulumi.Input[str]] = None,
            kubernetes_authorization_type: Optional[pulumi.Input[str]] = None,
            kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
            kubernetes_namespace: Optional[pulumi.Input[str]] = None,
            kubernetes_token: Optional[pulumi.Input[str]] = None,
            managed: Optional[pulumi.Input[bool]] = None,
            management_project_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            platform_type: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            provider_type: Optional[pulumi.Input[str]] = None) -> 'ProjectCluster':
        """
        Get an existing ProjectCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: The base domain of the cluster.
        :param pulumi.Input[bool] enabled: Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] environment_scope: The associated environment to the cluster. Defaults to `*`.
        :param pulumi.Input[str] kubernetes_api_url: The URL to access the Kubernetes API.
        :param pulumi.Input[str] kubernetes_authorization_type: The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        :param pulumi.Input[str] kubernetes_ca_cert: TLS certificate (needed if API is using a self-signed TLS certificate).
        :param pulumi.Input[str] kubernetes_namespace: The unique namespace related to the project.
        :param pulumi.Input[str] kubernetes_token: The token to authenticate against Kubernetes.
        :param pulumi.Input[bool] managed: Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        :param pulumi.Input[str] management_project_id: The ID of the management project for the cluster.
        :param pulumi.Input[str] name: The name of cluster.
        :param pulumi.Input[str] project: The id of the project to add the cluster to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cluster_type"] = cluster_type
        __props__["created_at"] = created_at
        __props__["domain"] = domain
        __props__["enabled"] = enabled
        __props__["environment_scope"] = environment_scope
        __props__["kubernetes_api_url"] = kubernetes_api_url
        __props__["kubernetes_authorization_type"] = kubernetes_authorization_type
        __props__["kubernetes_ca_cert"] = kubernetes_ca_cert
        __props__["kubernetes_namespace"] = kubernetes_namespace
        __props__["kubernetes_token"] = kubernetes_token
        __props__["managed"] = managed
        __props__["management_project_id"] = management_project_id
        __props__["name"] = name
        __props__["platform_type"] = platform_type
        __props__["project"] = project
        __props__["provider_type"] = provider_type
        return ProjectCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
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
    @pulumi.getter(name="kubernetesNamespace")
    def kubernetes_namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The unique namespace related to the project.
        """
        return pulumi.get(self, "kubernetes_namespace")

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
        return pulumi.get(self, "platform_type")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The id of the project to add the cluster to.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="providerType")
    def provider_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "provider_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

