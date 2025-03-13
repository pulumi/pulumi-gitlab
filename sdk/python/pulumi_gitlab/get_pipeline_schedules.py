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
from . import outputs

__all__ = [
    'GetPipelineSchedulesResult',
    'AwaitableGetPipelineSchedulesResult',
    'get_pipeline_schedules',
    'get_pipeline_schedules_output',
]

@pulumi.output_type
class GetPipelineSchedulesResult:
    """
    A collection of values returned by getPipelineSchedules.
    """
    def __init__(__self__, id=None, pipeline_schedules=None, project=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if pipeline_schedules and not isinstance(pipeline_schedules, list):
            raise TypeError("Expected argument 'pipeline_schedules' to be a list")
        pulumi.set(__self__, "pipeline_schedules", pipeline_schedules)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="pipelineSchedules")
    def pipeline_schedules(self) -> Sequence['outputs.GetPipelineSchedulesPipelineScheduleResult']:
        """
        The list of pipeline schedules.
        """
        return pulumi.get(self, "pipeline_schedules")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The name or id of the project to add the schedule to.
        """
        return pulumi.get(self, "project")


class AwaitableGetPipelineSchedulesResult(GetPipelineSchedulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPipelineSchedulesResult(
            id=self.id,
            pipeline_schedules=self.pipeline_schedules,
            project=self.project)


def get_pipeline_schedules(project: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPipelineSchedulesResult:
    """
    The `PipelineSchedule` data source retrieves information about a gitlab pipeline schedule for a project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/pipeline_schedules/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_pipeline_schedules(project="12345")
    ```


    :param str project: The name or id of the project to add the schedule to.
    """
    __args__ = dict()
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getPipelineSchedules:getPipelineSchedules', __args__, opts=opts, typ=GetPipelineSchedulesResult).value

    return AwaitableGetPipelineSchedulesResult(
        id=pulumi.get(__ret__, 'id'),
        pipeline_schedules=pulumi.get(__ret__, 'pipeline_schedules'),
        project=pulumi.get(__ret__, 'project'))
def get_pipeline_schedules_output(project: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPipelineSchedulesResult]:
    """
    The `PipelineSchedule` data source retrieves information about a gitlab pipeline schedule for a project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/pipeline_schedules/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_pipeline_schedules(project="12345")
    ```


    :param str project: The name or id of the project to add the schedule to.
    """
    __args__ = dict()
    __args__['project'] = project
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getPipelineSchedules:getPipelineSchedules', __args__, opts=opts, typ=GetPipelineSchedulesResult)
    return __ret__.apply(lambda __response__: GetPipelineSchedulesResult(
        id=pulumi.get(__response__, 'id'),
        pipeline_schedules=pulumi.get(__response__, 'pipeline_schedules'),
        project=pulumi.get(__response__, 'project')))
