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

__all__ = [
    'GetInstanceVariableResult',
    'AwaitableGetInstanceVariableResult',
    'get_instance_variable',
    'get_instance_variable_output',
]

@pulumi.output_type
class GetInstanceVariableResult:
    """
    A collection of values returned by getInstanceVariable.
    """
    def __init__(__self__, description=None, id=None, key=None, masked=None, protected=None, raw=None, value=None, variable_type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if masked and not isinstance(masked, bool):
            raise TypeError("Expected argument 'masked' to be a bool")
        pulumi.set(__self__, "masked", masked)
        if protected and not isinstance(protected, bool):
            raise TypeError("Expected argument 'protected' to be a bool")
        pulumi.set(__self__, "protected", protected)
        if raw and not isinstance(raw, bool):
            raise TypeError("Expected argument 'raw' to be a bool")
        pulumi.set(__self__, "raw", raw)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)
        if variable_type and not isinstance(variable_type, str):
            raise TypeError("Expected argument 'variable_type' to be a str")
        pulumi.set(__self__, "variable_type", variable_type)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the variable. Maximum of 255 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The name of the variable.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def masked(self) -> bool:
        """
        If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        """
        return pulumi.get(self, "masked")

    @property
    @pulumi.getter
    def protected(self) -> bool:
        """
        If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        """
        return pulumi.get(self, "protected")

    @property
    @pulumi.getter
    def raw(self) -> bool:
        """
        Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        """
        return pulumi.get(self, "raw")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the variable.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="variableType")
    def variable_type(self) -> str:
        """
        The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        """
        return pulumi.get(self, "variable_type")


class AwaitableGetInstanceVariableResult(GetInstanceVariableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceVariableResult(
            description=self.description,
            id=self.id,
            key=self.key,
            masked=self.masked,
            protected=self.protected,
            raw=self.raw,
            value=self.value,
            variable_type=self.variable_type)


def get_instance_variable(key: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceVariableResult:
    """
    The `InstanceVariable` data source allows to retrieve details about an instance-level CI/CD variable.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/instance_level_ci_variables.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    foo = gitlab.get_instance_variable(key="foo")
    ```


    :param str key: The name of the variable.
    """
    __args__ = dict()
    __args__['key'] = key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getInstanceVariable:getInstanceVariable', __args__, opts=opts, typ=GetInstanceVariableResult).value

    return AwaitableGetInstanceVariableResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        key=pulumi.get(__ret__, 'key'),
        masked=pulumi.get(__ret__, 'masked'),
        protected=pulumi.get(__ret__, 'protected'),
        raw=pulumi.get(__ret__, 'raw'),
        value=pulumi.get(__ret__, 'value'),
        variable_type=pulumi.get(__ret__, 'variable_type'))
def get_instance_variable_output(key: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceVariableResult]:
    """
    The `InstanceVariable` data source allows to retrieve details about an instance-level CI/CD variable.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/instance_level_ci_variables.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    foo = gitlab.get_instance_variable(key="foo")
    ```


    :param str key: The name of the variable.
    """
    __args__ = dict()
    __args__['key'] = key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getInstanceVariable:getInstanceVariable', __args__, opts=opts, typ=GetInstanceVariableResult)
    return __ret__.apply(lambda __response__: GetInstanceVariableResult(
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        key=pulumi.get(__response__, 'key'),
        masked=pulumi.get(__response__, 'masked'),
        protected=pulumi.get(__response__, 'protected'),
        raw=pulumi.get(__response__, 'raw'),
        value=pulumi.get(__response__, 'value'),
        variable_type=pulumi.get(__response__, 'variable_type')))
