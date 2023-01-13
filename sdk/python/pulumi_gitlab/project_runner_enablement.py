# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectRunnerEnablementArgs', 'ProjectRunnerEnablement']

@pulumi.input_type
class ProjectRunnerEnablementArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 runner_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a ProjectRunnerEnablement resource.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[int] runner_id: The ID of a runner to enable for the project.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "runner_id", runner_id)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or URL-encoded path of the project owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="runnerId")
    def runner_id(self) -> pulumi.Input[int]:
        """
        The ID of a runner to enable for the project.
        """
        return pulumi.get(self, "runner_id")

    @runner_id.setter
    def runner_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "runner_id", value)


@pulumi.input_type
class _ProjectRunnerEnablementState:
    def __init__(__self__, *,
                 project: Optional[pulumi.Input[str]] = None,
                 runner_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ProjectRunnerEnablement resources.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[int] runner_id: The ID of a runner to enable for the project.
        """
        if project is not None:
            pulumi.set(__self__, "project", project)
        if runner_id is not None:
            pulumi.set(__self__, "runner_id", runner_id)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or URL-encoded path of the project owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="runnerId")
    def runner_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of a runner to enable for the project.
        """
        return pulumi.get(self, "runner_id")

    @runner_id.setter
    def runner_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "runner_id", value)


class ProjectRunnerEnablement(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 runner_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The `ProjectRunnerEnablement` resource allows to enable a runner in a project.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/runners.html#enable-a-runner-in-project)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.ProjectRunnerEnablement("foo",
            project="5",
            runner_id=7)
        ```

        ## Import

        GitLab project runners can be imported using an id made up of `project:runner_id`, e.g.

        ```sh
         $ pulumi import gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement foo 5:7
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[int] runner_id: The ID of a runner to enable for the project.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectRunnerEnablementArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectRunnerEnablement` resource allows to enable a runner in a project.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/runners.html#enable-a-runner-in-project)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.ProjectRunnerEnablement("foo",
            project="5",
            runner_id=7)
        ```

        ## Import

        GitLab project runners can be imported using an id made up of `project:runner_id`, e.g.

        ```sh
         $ pulumi import gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement foo 5:7
        ```

        :param str resource_name: The name of the resource.
        :param ProjectRunnerEnablementArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectRunnerEnablementArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 runner_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectRunnerEnablementArgs.__new__(ProjectRunnerEnablementArgs)

            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if runner_id is None and not opts.urn:
                raise TypeError("Missing required property 'runner_id'")
            __props__.__dict__["runner_id"] = runner_id
        super(ProjectRunnerEnablement, __self__).__init__(
            'gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            project: Optional[pulumi.Input[str]] = None,
            runner_id: Optional[pulumi.Input[int]] = None) -> 'ProjectRunnerEnablement':
        """
        Get an existing ProjectRunnerEnablement resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[int] runner_id: The ID of a runner to enable for the project.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectRunnerEnablementState.__new__(_ProjectRunnerEnablementState)

        __props__.__dict__["project"] = project
        __props__.__dict__["runner_id"] = runner_id
        return ProjectRunnerEnablement(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of the project owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="runnerId")
    def runner_id(self) -> pulumi.Output[int]:
        """
        The ID of a runner to enable for the project.
        """
        return pulumi.get(self, "runner_id")
