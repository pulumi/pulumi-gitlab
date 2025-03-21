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

__all__ = ['ProjectEnvironmentArgs', 'ProjectEnvironment']

@pulumi.input_type
class ProjectEnvironmentArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 auto_stop_setting: Optional[pulumi.Input[str]] = None,
                 cluster_agent_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_url: Optional[pulumi.Input[str]] = None,
                 flux_resource_path: Optional[pulumi.Input[str]] = None,
                 kubernetes_namespace: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stop_before_destroy: Optional[pulumi.Input[bool]] = None,
                 tier: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectEnvironment resource.
        :param pulumi.Input[str] project: The ID or full path of the project to environment is created for.
        :param pulumi.Input[str] auto_stop_setting: The auto stop setting for the environment. Allowed values are `always`, `with_action`. If this is set to `with_action` and `stop_before_destroy` is `true`, the environment will be force-stopped.
        :param pulumi.Input[int] cluster_agent_id: The cluster agent to associate with this environment.
        :param pulumi.Input[str] description: The description of the environment.
        :param pulumi.Input[str] external_url: Place to link to for this environment.
        :param pulumi.Input[str] flux_resource_path: The Flux resource path to associate with this environment.
        :param pulumi.Input[str] kubernetes_namespace: The Kubernetes namespace to associate with this environment.
        :param pulumi.Input[str] name: The name of the environment.
        :param pulumi.Input[bool] stop_before_destroy: Determines whether the environment is attempted to be stopped before the environment is deleted. If `auto_stop_setting` is set to `with_action`, this will perform a force stop.
        :param pulumi.Input[str] tier: The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
        """
        pulumi.set(__self__, "project", project)
        if auto_stop_setting is not None:
            pulumi.set(__self__, "auto_stop_setting", auto_stop_setting)
        if cluster_agent_id is not None:
            pulumi.set(__self__, "cluster_agent_id", cluster_agent_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if external_url is not None:
            pulumi.set(__self__, "external_url", external_url)
        if flux_resource_path is not None:
            pulumi.set(__self__, "flux_resource_path", flux_resource_path)
        if kubernetes_namespace is not None:
            pulumi.set(__self__, "kubernetes_namespace", kubernetes_namespace)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if stop_before_destroy is not None:
            pulumi.set(__self__, "stop_before_destroy", stop_before_destroy)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or full path of the project to environment is created for.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="autoStopSetting")
    def auto_stop_setting(self) -> Optional[pulumi.Input[str]]:
        """
        The auto stop setting for the environment. Allowed values are `always`, `with_action`. If this is set to `with_action` and `stop_before_destroy` is `true`, the environment will be force-stopped.
        """
        return pulumi.get(self, "auto_stop_setting")

    @auto_stop_setting.setter
    def auto_stop_setting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_stop_setting", value)

    @property
    @pulumi.getter(name="clusterAgentId")
    def cluster_agent_id(self) -> Optional[pulumi.Input[int]]:
        """
        The cluster agent to associate with this environment.
        """
        return pulumi.get(self, "cluster_agent_id")

    @cluster_agent_id.setter
    def cluster_agent_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cluster_agent_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the environment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="externalUrl")
    def external_url(self) -> Optional[pulumi.Input[str]]:
        """
        Place to link to for this environment.
        """
        return pulumi.get(self, "external_url")

    @external_url.setter
    def external_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_url", value)

    @property
    @pulumi.getter(name="fluxResourcePath")
    def flux_resource_path(self) -> Optional[pulumi.Input[str]]:
        """
        The Flux resource path to associate with this environment.
        """
        return pulumi.get(self, "flux_resource_path")

    @flux_resource_path.setter
    def flux_resource_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flux_resource_path", value)

    @property
    @pulumi.getter(name="kubernetesNamespace")
    def kubernetes_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The Kubernetes namespace to associate with this environment.
        """
        return pulumi.get(self, "kubernetes_namespace")

    @kubernetes_namespace.setter
    def kubernetes_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_namespace", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="stopBeforeDestroy")
    def stop_before_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether the environment is attempted to be stopped before the environment is deleted. If `auto_stop_setting` is set to `with_action`, this will perform a force stop.
        """
        return pulumi.get(self, "stop_before_destroy")

    @stop_before_destroy.setter
    def stop_before_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "stop_before_destroy", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)


@pulumi.input_type
class _ProjectEnvironmentState:
    def __init__(__self__, *,
                 auto_stop_at: Optional[pulumi.Input[str]] = None,
                 auto_stop_setting: Optional[pulumi.Input[str]] = None,
                 cluster_agent_id: Optional[pulumi.Input[int]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_url: Optional[pulumi.Input[str]] = None,
                 flux_resource_path: Optional[pulumi.Input[str]] = None,
                 kubernetes_namespace: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 stop_before_destroy: Optional[pulumi.Input[bool]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectEnvironment resources.
        :param pulumi.Input[str] auto_stop_at: The ISO8601 date/time that this environment will be automatically stopped at in UTC.
        :param pulumi.Input[str] auto_stop_setting: The auto stop setting for the environment. Allowed values are `always`, `with_action`. If this is set to `with_action` and `stop_before_destroy` is `true`, the environment will be force-stopped.
        :param pulumi.Input[int] cluster_agent_id: The cluster agent to associate with this environment.
        :param pulumi.Input[str] created_at: The ISO8601 date/time that this environment was created at in UTC.
        :param pulumi.Input[str] description: The description of the environment.
        :param pulumi.Input[str] external_url: Place to link to for this environment.
        :param pulumi.Input[str] flux_resource_path: The Flux resource path to associate with this environment.
        :param pulumi.Input[str] kubernetes_namespace: The Kubernetes namespace to associate with this environment.
        :param pulumi.Input[str] name: The name of the environment.
        :param pulumi.Input[str] project: The ID or full path of the project to environment is created for.
        :param pulumi.Input[str] slug: The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        :param pulumi.Input[str] state: State the environment is in. Valid values are `available`, `stopped`.
        :param pulumi.Input[bool] stop_before_destroy: Determines whether the environment is attempted to be stopped before the environment is deleted. If `auto_stop_setting` is set to `with_action`, this will perform a force stop.
        :param pulumi.Input[str] tier: The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
        :param pulumi.Input[str] updated_at: The ISO8601 date/time that this environment was last updated at in UTC.
        """
        if auto_stop_at is not None:
            pulumi.set(__self__, "auto_stop_at", auto_stop_at)
        if auto_stop_setting is not None:
            pulumi.set(__self__, "auto_stop_setting", auto_stop_setting)
        if cluster_agent_id is not None:
            pulumi.set(__self__, "cluster_agent_id", cluster_agent_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if external_url is not None:
            pulumi.set(__self__, "external_url", external_url)
        if flux_resource_path is not None:
            pulumi.set(__self__, "flux_resource_path", flux_resource_path)
        if kubernetes_namespace is not None:
            pulumi.set(__self__, "kubernetes_namespace", kubernetes_namespace)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if stop_before_destroy is not None:
            pulumi.set(__self__, "stop_before_destroy", stop_before_destroy)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="autoStopAt")
    def auto_stop_at(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO8601 date/time that this environment will be automatically stopped at in UTC.
        """
        return pulumi.get(self, "auto_stop_at")

    @auto_stop_at.setter
    def auto_stop_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_stop_at", value)

    @property
    @pulumi.getter(name="autoStopSetting")
    def auto_stop_setting(self) -> Optional[pulumi.Input[str]]:
        """
        The auto stop setting for the environment. Allowed values are `always`, `with_action`. If this is set to `with_action` and `stop_before_destroy` is `true`, the environment will be force-stopped.
        """
        return pulumi.get(self, "auto_stop_setting")

    @auto_stop_setting.setter
    def auto_stop_setting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_stop_setting", value)

    @property
    @pulumi.getter(name="clusterAgentId")
    def cluster_agent_id(self) -> Optional[pulumi.Input[int]]:
        """
        The cluster agent to associate with this environment.
        """
        return pulumi.get(self, "cluster_agent_id")

    @cluster_agent_id.setter
    def cluster_agent_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cluster_agent_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO8601 date/time that this environment was created at in UTC.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the environment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="externalUrl")
    def external_url(self) -> Optional[pulumi.Input[str]]:
        """
        Place to link to for this environment.
        """
        return pulumi.get(self, "external_url")

    @external_url.setter
    def external_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_url", value)

    @property
    @pulumi.getter(name="fluxResourcePath")
    def flux_resource_path(self) -> Optional[pulumi.Input[str]]:
        """
        The Flux resource path to associate with this environment.
        """
        return pulumi.get(self, "flux_resource_path")

    @flux_resource_path.setter
    def flux_resource_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flux_resource_path", value)

    @property
    @pulumi.getter(name="kubernetesNamespace")
    def kubernetes_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The Kubernetes namespace to associate with this environment.
        """
        return pulumi.get(self, "kubernetes_namespace")

    @kubernetes_namespace.setter
    def kubernetes_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_namespace", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or full path of the project to environment is created for.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        """
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        State the environment is in. Valid values are `available`, `stopped`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="stopBeforeDestroy")
    def stop_before_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether the environment is attempted to be stopped before the environment is deleted. If `auto_stop_setting` is set to `with_action`, this will perform a force stop.
        """
        return pulumi.get(self, "stop_before_destroy")

    @stop_before_destroy.setter
    def stop_before_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "stop_before_destroy", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO8601 date/time that this environment was last updated at in UTC.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class ProjectEnvironment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_stop_setting: Optional[pulumi.Input[str]] = None,
                 cluster_agent_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_url: Optional[pulumi.Input[str]] = None,
                 flux_resource_path: Optional[pulumi.Input[str]] = None,
                 kubernetes_namespace: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 stop_before_destroy: Optional[pulumi.Input[bool]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        this = gitlab.Group("this",
            name="example",
            path="example",
            description="An example group")
        this_project = gitlab.Project("this",
            name="example",
            namespace_id=this.id,
            initialize_with_readme=True)
        this_project_environment = gitlab.ProjectEnvironment("this",
            project=this_project.id,
            name="example",
            external_url="www.example.com")
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_environment`. For example:

        terraform

        import {

          to = gitlab_project_environment.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        GitLab project environments can be imported using an id made up of `projectId:environmenId`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectEnvironment:ProjectEnvironment bar 123:321
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_stop_setting: The auto stop setting for the environment. Allowed values are `always`, `with_action`. If this is set to `with_action` and `stop_before_destroy` is `true`, the environment will be force-stopped.
        :param pulumi.Input[int] cluster_agent_id: The cluster agent to associate with this environment.
        :param pulumi.Input[str] description: The description of the environment.
        :param pulumi.Input[str] external_url: Place to link to for this environment.
        :param pulumi.Input[str] flux_resource_path: The Flux resource path to associate with this environment.
        :param pulumi.Input[str] kubernetes_namespace: The Kubernetes namespace to associate with this environment.
        :param pulumi.Input[str] name: The name of the environment.
        :param pulumi.Input[str] project: The ID or full path of the project to environment is created for.
        :param pulumi.Input[bool] stop_before_destroy: Determines whether the environment is attempted to be stopped before the environment is deleted. If `auto_stop_setting` is set to `with_action`, this will perform a force stop.
        :param pulumi.Input[str] tier: The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectEnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        this = gitlab.Group("this",
            name="example",
            path="example",
            description="An example group")
        this_project = gitlab.Project("this",
            name="example",
            namespace_id=this.id,
            initialize_with_readme=True)
        this_project_environment = gitlab.ProjectEnvironment("this",
            project=this_project.id,
            name="example",
            external_url="www.example.com")
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_environment`. For example:

        terraform

        import {

          to = gitlab_project_environment.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        GitLab project environments can be imported using an id made up of `projectId:environmenId`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectEnvironment:ProjectEnvironment bar 123:321
        ```

        :param str resource_name: The name of the resource.
        :param ProjectEnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectEnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_stop_setting: Optional[pulumi.Input[str]] = None,
                 cluster_agent_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_url: Optional[pulumi.Input[str]] = None,
                 flux_resource_path: Optional[pulumi.Input[str]] = None,
                 kubernetes_namespace: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 stop_before_destroy: Optional[pulumi.Input[bool]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectEnvironmentArgs.__new__(ProjectEnvironmentArgs)

            __props__.__dict__["auto_stop_setting"] = auto_stop_setting
            __props__.__dict__["cluster_agent_id"] = cluster_agent_id
            __props__.__dict__["description"] = description
            __props__.__dict__["external_url"] = external_url
            __props__.__dict__["flux_resource_path"] = flux_resource_path
            __props__.__dict__["kubernetes_namespace"] = kubernetes_namespace
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["stop_before_destroy"] = stop_before_destroy
            __props__.__dict__["tier"] = tier
            __props__.__dict__["auto_stop_at"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["slug"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["updated_at"] = None
        super(ProjectEnvironment, __self__).__init__(
            'gitlab:index/projectEnvironment:ProjectEnvironment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_stop_at: Optional[pulumi.Input[str]] = None,
            auto_stop_setting: Optional[pulumi.Input[str]] = None,
            cluster_agent_id: Optional[pulumi.Input[int]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            external_url: Optional[pulumi.Input[str]] = None,
            flux_resource_path: Optional[pulumi.Input[str]] = None,
            kubernetes_namespace: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            slug: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            stop_before_destroy: Optional[pulumi.Input[bool]] = None,
            tier: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'ProjectEnvironment':
        """
        Get an existing ProjectEnvironment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_stop_at: The ISO8601 date/time that this environment will be automatically stopped at in UTC.
        :param pulumi.Input[str] auto_stop_setting: The auto stop setting for the environment. Allowed values are `always`, `with_action`. If this is set to `with_action` and `stop_before_destroy` is `true`, the environment will be force-stopped.
        :param pulumi.Input[int] cluster_agent_id: The cluster agent to associate with this environment.
        :param pulumi.Input[str] created_at: The ISO8601 date/time that this environment was created at in UTC.
        :param pulumi.Input[str] description: The description of the environment.
        :param pulumi.Input[str] external_url: Place to link to for this environment.
        :param pulumi.Input[str] flux_resource_path: The Flux resource path to associate with this environment.
        :param pulumi.Input[str] kubernetes_namespace: The Kubernetes namespace to associate with this environment.
        :param pulumi.Input[str] name: The name of the environment.
        :param pulumi.Input[str] project: The ID or full path of the project to environment is created for.
        :param pulumi.Input[str] slug: The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        :param pulumi.Input[str] state: State the environment is in. Valid values are `available`, `stopped`.
        :param pulumi.Input[bool] stop_before_destroy: Determines whether the environment is attempted to be stopped before the environment is deleted. If `auto_stop_setting` is set to `with_action`, this will perform a force stop.
        :param pulumi.Input[str] tier: The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
        :param pulumi.Input[str] updated_at: The ISO8601 date/time that this environment was last updated at in UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectEnvironmentState.__new__(_ProjectEnvironmentState)

        __props__.__dict__["auto_stop_at"] = auto_stop_at
        __props__.__dict__["auto_stop_setting"] = auto_stop_setting
        __props__.__dict__["cluster_agent_id"] = cluster_agent_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["external_url"] = external_url
        __props__.__dict__["flux_resource_path"] = flux_resource_path
        __props__.__dict__["kubernetes_namespace"] = kubernetes_namespace
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["slug"] = slug
        __props__.__dict__["state"] = state
        __props__.__dict__["stop_before_destroy"] = stop_before_destroy
        __props__.__dict__["tier"] = tier
        __props__.__dict__["updated_at"] = updated_at
        return ProjectEnvironment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoStopAt")
    def auto_stop_at(self) -> pulumi.Output[str]:
        """
        The ISO8601 date/time that this environment will be automatically stopped at in UTC.
        """
        return pulumi.get(self, "auto_stop_at")

    @property
    @pulumi.getter(name="autoStopSetting")
    def auto_stop_setting(self) -> pulumi.Output[str]:
        """
        The auto stop setting for the environment. Allowed values are `always`, `with_action`. If this is set to `with_action` and `stop_before_destroy` is `true`, the environment will be force-stopped.
        """
        return pulumi.get(self, "auto_stop_setting")

    @property
    @pulumi.getter(name="clusterAgentId")
    def cluster_agent_id(self) -> pulumi.Output[Optional[int]]:
        """
        The cluster agent to associate with this environment.
        """
        return pulumi.get(self, "cluster_agent_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The ISO8601 date/time that this environment was created at in UTC.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the environment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalUrl")
    def external_url(self) -> pulumi.Output[Optional[str]]:
        """
        Place to link to for this environment.
        """
        return pulumi.get(self, "external_url")

    @property
    @pulumi.getter(name="fluxResourcePath")
    def flux_resource_path(self) -> pulumi.Output[Optional[str]]:
        """
        The Flux resource path to associate with this environment.
        """
        return pulumi.get(self, "flux_resource_path")

    @property
    @pulumi.getter(name="kubernetesNamespace")
    def kubernetes_namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The Kubernetes namespace to associate with this environment.
        """
        return pulumi.get(self, "kubernetes_namespace")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or full path of the project to environment is created for.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        """
        The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        """
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State the environment is in. Valid values are `available`, `stopped`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stopBeforeDestroy")
    def stop_before_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines whether the environment is attempted to be stopped before the environment is deleted. If `auto_stop_setting` is set to `with_action`, this will perform a force stop.
        """
        return pulumi.get(self, "stop_before_destroy")

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Output[str]:
        """
        The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The ISO8601 date/time that this environment was last updated at in UTC.
        """
        return pulumi.get(self, "updated_at")

