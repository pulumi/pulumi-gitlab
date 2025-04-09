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

__all__ = ['PipelineScheduleVariableArgs', 'PipelineScheduleVariable']

@pulumi.input_type
class PipelineScheduleVariableArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[builtins.str],
                 pipeline_schedule_id: pulumi.Input[builtins.int],
                 project: pulumi.Input[builtins.str],
                 value: pulumi.Input[builtins.str],
                 variable_type: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a PipelineScheduleVariable resource.
        :param pulumi.Input[builtins.str] key: Name of the variable.
        :param pulumi.Input[builtins.int] pipeline_schedule_id: The id of the pipeline schedule.
        :param pulumi.Input[builtins.str] project: The id of the project to add the schedule to.
        :param pulumi.Input[builtins.str] value: Value of the variable.
        :param pulumi.Input[builtins.str] variable_type: The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "pipeline_schedule_id", pipeline_schedule_id)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "value", value)
        if variable_type is not None:
            pulumi.set(__self__, "variable_type", variable_type)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[builtins.str]:
        """
        Name of the variable.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="pipelineScheduleId")
    def pipeline_schedule_id(self) -> pulumi.Input[builtins.int]:
        """
        The id of the pipeline schedule.
        """
        return pulumi.get(self, "pipeline_schedule_id")

    @pipeline_schedule_id.setter
    def pipeline_schedule_id(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "pipeline_schedule_id", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[builtins.str]:
        """
        The id of the project to add the schedule to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[builtins.str]:
        """
        Value of the variable.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="variableType")
    def variable_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
        """
        return pulumi.get(self, "variable_type")

    @variable_type.setter
    def variable_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "variable_type", value)


@pulumi.input_type
class _PipelineScheduleVariableState:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[builtins.str]] = None,
                 pipeline_schedule_id: Optional[pulumi.Input[builtins.int]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None,
                 variable_type: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering PipelineScheduleVariable resources.
        :param pulumi.Input[builtins.str] key: Name of the variable.
        :param pulumi.Input[builtins.int] pipeline_schedule_id: The id of the pipeline schedule.
        :param pulumi.Input[builtins.str] project: The id of the project to add the schedule to.
        :param pulumi.Input[builtins.str] value: Value of the variable.
        :param pulumi.Input[builtins.str] variable_type: The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if pipeline_schedule_id is not None:
            pulumi.set(__self__, "pipeline_schedule_id", pipeline_schedule_id)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if variable_type is not None:
            pulumi.set(__self__, "variable_type", variable_type)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the variable.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="pipelineScheduleId")
    def pipeline_schedule_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The id of the pipeline schedule.
        """
        return pulumi.get(self, "pipeline_schedule_id")

    @pipeline_schedule_id.setter
    def pipeline_schedule_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "pipeline_schedule_id", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The id of the project to add the schedule to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Value of the variable.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="variableType")
    def variable_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
        """
        return pulumi.get(self, "variable_type")

    @variable_type.setter
    def variable_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "variable_type", value)


class PipelineScheduleVariable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[builtins.str]] = None,
                 pipeline_schedule_id: Optional[pulumi.Input[builtins.int]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None,
                 variable_type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `PipelineScheduleVariable` resource allows to manage the lifecycle of a variable for a pipeline schedule.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/pipeline_schedules/#pipeline-schedule-variables)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.PipelineSchedule("example",
            project="12345",
            description="Used to schedule builds",
            ref="master",
            cron="0 1 * * *")
        example_pipeline_schedule_variable = gitlab.PipelineScheduleVariable("example",
            project=example.project,
            pipeline_schedule_id=example.pipeline_schedule_id,
            key="EXAMPLE_KEY",
            value="example")
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_pipeline_schedule_variable`. For example:

        terraform

        import {

          to = gitlab_pipeline_schedule_variable.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        Pipeline schedule variables can be imported using an id made up of `project_id:pipeline_schedule_id:key`, e.g.

        ```sh
        $ pulumi import gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable example 123456789:13:mykey
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] key: Name of the variable.
        :param pulumi.Input[builtins.int] pipeline_schedule_id: The id of the pipeline schedule.
        :param pulumi.Input[builtins.str] project: The id of the project to add the schedule to.
        :param pulumi.Input[builtins.str] value: Value of the variable.
        :param pulumi.Input[builtins.str] variable_type: The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineScheduleVariableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `PipelineScheduleVariable` resource allows to manage the lifecycle of a variable for a pipeline schedule.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/pipeline_schedules/#pipeline-schedule-variables)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.PipelineSchedule("example",
            project="12345",
            description="Used to schedule builds",
            ref="master",
            cron="0 1 * * *")
        example_pipeline_schedule_variable = gitlab.PipelineScheduleVariable("example",
            project=example.project,
            pipeline_schedule_id=example.pipeline_schedule_id,
            key="EXAMPLE_KEY",
            value="example")
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_pipeline_schedule_variable`. For example:

        terraform

        import {

          to = gitlab_pipeline_schedule_variable.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        Pipeline schedule variables can be imported using an id made up of `project_id:pipeline_schedule_id:key`, e.g.

        ```sh
        $ pulumi import gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable example 123456789:13:mykey
        ```

        :param str resource_name: The name of the resource.
        :param PipelineScheduleVariableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineScheduleVariableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[builtins.str]] = None,
                 pipeline_schedule_id: Optional[pulumi.Input[builtins.int]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None,
                 variable_type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PipelineScheduleVariableArgs.__new__(PipelineScheduleVariableArgs)

            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            if pipeline_schedule_id is None and not opts.urn:
                raise TypeError("Missing required property 'pipeline_schedule_id'")
            __props__.__dict__["pipeline_schedule_id"] = pipeline_schedule_id
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            __props__.__dict__["variable_type"] = variable_type
        super(PipelineScheduleVariable, __self__).__init__(
            'gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key: Optional[pulumi.Input[builtins.str]] = None,
            pipeline_schedule_id: Optional[pulumi.Input[builtins.int]] = None,
            project: Optional[pulumi.Input[builtins.str]] = None,
            value: Optional[pulumi.Input[builtins.str]] = None,
            variable_type: Optional[pulumi.Input[builtins.str]] = None) -> 'PipelineScheduleVariable':
        """
        Get an existing PipelineScheduleVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] key: Name of the variable.
        :param pulumi.Input[builtins.int] pipeline_schedule_id: The id of the pipeline schedule.
        :param pulumi.Input[builtins.str] project: The id of the project to add the schedule to.
        :param pulumi.Input[builtins.str] value: Value of the variable.
        :param pulumi.Input[builtins.str] variable_type: The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PipelineScheduleVariableState.__new__(_PipelineScheduleVariableState)

        __props__.__dict__["key"] = key
        __props__.__dict__["pipeline_schedule_id"] = pipeline_schedule_id
        __props__.__dict__["project"] = project
        __props__.__dict__["value"] = value
        __props__.__dict__["variable_type"] = variable_type
        return PipelineScheduleVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[builtins.str]:
        """
        Name of the variable.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="pipelineScheduleId")
    def pipeline_schedule_id(self) -> pulumi.Output[builtins.int]:
        """
        The id of the pipeline schedule.
        """
        return pulumi.get(self, "pipeline_schedule_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[builtins.str]:
        """
        The id of the project to add the schedule to.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[builtins.str]:
        """
        Value of the variable.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="variableType")
    def variable_type(self) -> pulumi.Output[builtins.str]:
        """
        The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
        """
        return pulumi.get(self, "variable_type")

