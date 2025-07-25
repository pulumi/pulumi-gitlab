# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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
    'GetProjectHooksResult',
    'AwaitableGetProjectHooksResult',
    'get_project_hooks',
    'get_project_hooks_output',
]

@pulumi.output_type
class GetProjectHooksResult:
    """
    A collection of values returned by getProjectHooks.
    """
    def __init__(__self__, hooks=None, id=None, project=None):
        if hooks and not isinstance(hooks, list):
            raise TypeError("Expected argument 'hooks' to be a list")
        pulumi.set(__self__, "hooks", hooks)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)

    @_builtins.property
    @pulumi.getter
    def hooks(self) -> Sequence['outputs.GetProjectHooksHookResult']:
        """
        The list of hooks.
        """
        return pulumi.get(self, "hooks")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def project(self) -> _builtins.str:
        """
        The name or id of the project.
        """
        return pulumi.get(self, "project")


class AwaitableGetProjectHooksResult(GetProjectHooksResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectHooksResult(
            hooks=self.hooks,
            id=self.id,
            project=self.project)


def get_project_hooks(project: Optional[_builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectHooksResult:
    """
    The `get_project_hooks` data source allows to retrieve details about hooks in a project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_webhooks/#list-webhooks-for-a-project)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_project(id="foo/bar/baz")
    examples = gitlab.get_project_hooks(project=example.id)
    ```


    :param _builtins.str project: The name or id of the project.
    """
    __args__ = dict()
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjectHooks:getProjectHooks', __args__, opts=opts, typ=GetProjectHooksResult).value

    return AwaitableGetProjectHooksResult(
        hooks=pulumi.get(__ret__, 'hooks'),
        id=pulumi.get(__ret__, 'id'),
        project=pulumi.get(__ret__, 'project'))
def get_project_hooks_output(project: Optional[pulumi.Input[_builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetProjectHooksResult]:
    """
    The `get_project_hooks` data source allows to retrieve details about hooks in a project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_webhooks/#list-webhooks-for-a-project)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_project(id="foo/bar/baz")
    examples = gitlab.get_project_hooks(project=example.id)
    ```


    :param _builtins.str project: The name or id of the project.
    """
    __args__ = dict()
    __args__['project'] = project
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getProjectHooks:getProjectHooks', __args__, opts=opts, typ=GetProjectHooksResult)
    return __ret__.apply(lambda __response__: GetProjectHooksResult(
        hooks=pulumi.get(__response__, 'hooks'),
        id=pulumi.get(__response__, 'id'),
        project=pulumi.get(__response__, 'project')))
