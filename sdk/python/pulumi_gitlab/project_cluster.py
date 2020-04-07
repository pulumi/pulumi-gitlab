# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class ProjectCluster(pulumi.CustomResource):
    cluster_type: pulumi.Output[str]
    created_at: pulumi.Output[str]
    domain: pulumi.Output[str]
    """
    The base domain of the cluster.
    """
    enabled: pulumi.Output[bool]
    """
    Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
    """
    environment_scope: pulumi.Output[str]
    """
    The associated environment to the cluster. Defaults to `*`.
    """
    kubernetes_api_url: pulumi.Output[str]
    """
    The URL to access the Kubernetes API.
    """
    kubernetes_authorization_type: pulumi.Output[str]
    """
    The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
    """
    kubernetes_ca_cert: pulumi.Output[str]
    """
    TLS certificate (needed if API is using a self-signed TLS certificate).
    """
    kubernetes_namespace: pulumi.Output[str]
    """
    The unique namespace related to the project.
    """
    kubernetes_token: pulumi.Output[str]
    """
    The token to authenticate against Kubernetes.
    """
    managed: pulumi.Output[bool]
    """
    Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
    """
    name: pulumi.Output[str]
    """
    The name of cluster.
    """
    platform_type: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The id of the project to add the cluster to.
    """
    provider_type: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, domain=None, enabled=None, environment_scope=None, kubernetes_api_url=None, kubernetes_authorization_type=None, kubernetes_ca_cert=None, kubernetes_namespace=None, kubernetes_token=None, managed=None, name=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        This resource allows you to create and manage project clusters for your GitLab projects.
        For further information on clusters, consult the [gitlab
        documentation](https://docs.gitlab.com/ce/user/project/clusters/index.html).

        > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/project_cluster.html.markdown.

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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, cluster_type=None, created_at=None, domain=None, enabled=None, environment_scope=None, kubernetes_api_url=None, kubernetes_authorization_type=None, kubernetes_ca_cert=None, kubernetes_namespace=None, kubernetes_token=None, managed=None, name=None, platform_type=None, project=None, provider_type=None):
        """
        Get an existing ProjectCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
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
        __props__["name"] = name
        __props__["platform_type"] = platform_type
        __props__["project"] = project
        __props__["provider_type"] = provider_type
        return ProjectCluster(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

