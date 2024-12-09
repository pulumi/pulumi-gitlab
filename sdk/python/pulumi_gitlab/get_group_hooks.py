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
    'GetGroupHooksResult',
    'AwaitableGetGroupHooksResult',
    'get_group_hooks',
    'get_group_hooks_output',
]

@pulumi.output_type
class GetGroupHooksResult:
    """
    A collection of values returned by getGroupHooks.
    """
    def __init__(__self__, group=None, hooks=None, id=None):
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        pulumi.set(__self__, "group", group)
        if hooks and not isinstance(hooks, list):
            raise TypeError("Expected argument 'hooks' to be a list")
        pulumi.set(__self__, "hooks", hooks)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def group(self) -> str:
        """
        The ID or full path of the group.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def hooks(self) -> Sequence['outputs.GetGroupHooksHookResult']:
        """
        The list of hooks.
        """
        return pulumi.get(self, "hooks")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetGroupHooksResult(GetGroupHooksResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupHooksResult(
            group=self.group,
            hooks=self.hooks,
            id=self.id)


def get_group_hooks(group: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupHooksResult:
    """
    The `get_group_hooks` data source allows to retrieve details about hooks in a group.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#list-group-hooks)


    :param str group: The ID or full path of the group.
    """
    __args__ = dict()
    __args__['group'] = group
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getGroupHooks:getGroupHooks', __args__, opts=opts, typ=GetGroupHooksResult).value

    return AwaitableGetGroupHooksResult(
        group=pulumi.get(__ret__, 'group'),
        hooks=pulumi.get(__ret__, 'hooks'),
        id=pulumi.get(__ret__, 'id'))
def get_group_hooks_output(group: Optional[pulumi.Input[str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGroupHooksResult]:
    """
    The `get_group_hooks` data source allows to retrieve details about hooks in a group.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#list-group-hooks)


    :param str group: The ID or full path of the group.
    """
    __args__ = dict()
    __args__['group'] = group
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getGroupHooks:getGroupHooks', __args__, opts=opts, typ=GetGroupHooksResult)
    return __ret__.apply(lambda __response__: GetGroupHooksResult(
        group=pulumi.get(__response__, 'group'),
        hooks=pulumi.get(__response__, 'hooks'),
        id=pulumi.get(__response__, 'id')))
