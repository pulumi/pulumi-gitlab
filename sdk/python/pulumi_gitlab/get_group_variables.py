# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetGroupVariablesResult',
    'AwaitableGetGroupVariablesResult',
    'get_group_variables',
    'get_group_variables_output',
]

@pulumi.output_type
class GetGroupVariablesResult:
    """
    A collection of values returned by getGroupVariables.
    """
    def __init__(__self__, environment_scope=None, group=None, id=None, variables=None):
        if environment_scope and not isinstance(environment_scope, str):
            raise TypeError("Expected argument 'environment_scope' to be a str")
        pulumi.set(__self__, "environment_scope", environment_scope)
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        pulumi.set(__self__, "group", group)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if variables and not isinstance(variables, list):
            raise TypeError("Expected argument 'variables' to be a list")
        pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter(name="environmentScope")
    def environment_scope(self) -> Optional[str]:
        """
        The environment scope of the variable. Defaults to all environment (`*`).
        """
        return pulumi.get(self, "environment_scope")

    @property
    @pulumi.getter
    def group(self) -> str:
        """
        The name or id of the group.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def variables(self) -> Sequence['outputs.GetGroupVariablesVariableResult']:
        """
        The list of variables returned by the search
        """
        return pulumi.get(self, "variables")


class AwaitableGetGroupVariablesResult(GetGroupVariablesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupVariablesResult(
            environment_scope=self.environment_scope,
            group=self.group,
            id=self.id,
            variables=self.variables)


def get_group_variables(environment_scope: Optional[str] = None,
                        group: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupVariablesResult:
    """
    The `get_group_variables` data source allows to retrieve all group-level CI/CD variables.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_level_variables.html)
    """
    __args__ = dict()
    __args__['environmentScope'] = environment_scope
    __args__['group'] = group
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getGroupVariables:getGroupVariables', __args__, opts=opts, typ=GetGroupVariablesResult).value

    return AwaitableGetGroupVariablesResult(
        environment_scope=pulumi.get(__ret__, 'environment_scope'),
        group=pulumi.get(__ret__, 'group'),
        id=pulumi.get(__ret__, 'id'),
        variables=pulumi.get(__ret__, 'variables'))


@_utilities.lift_output_func(get_group_variables)
def get_group_variables_output(environment_scope: Optional[pulumi.Input[Optional[str]]] = None,
                               group: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupVariablesResult]:
    """
    The `get_group_variables` data source allows to retrieve all group-level CI/CD variables.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_level_variables.html)
    """
    ...
