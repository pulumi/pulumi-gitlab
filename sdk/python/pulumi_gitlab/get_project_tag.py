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
    'GetProjectTagResult',
    'AwaitableGetProjectTagResult',
    'get_project_tag',
    'get_project_tag_output',
]

@pulumi.output_type
class GetProjectTagResult:
    """
    A collection of values returned by getProjectTag.
    """
    def __init__(__self__, commits=None, id=None, message=None, name=None, project=None, protected=None, releases=None, target=None):
        if commits and not isinstance(commits, list):
            raise TypeError("Expected argument 'commits' to be a list")
        pulumi.set(__self__, "commits", commits)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if message and not isinstance(message, str):
            raise TypeError("Expected argument 'message' to be a str")
        pulumi.set(__self__, "message", message)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if protected and not isinstance(protected, bool):
            raise TypeError("Expected argument 'protected' to be a bool")
        pulumi.set(__self__, "protected", protected)
        if releases and not isinstance(releases, list):
            raise TypeError("Expected argument 'releases' to be a list")
        pulumi.set(__self__, "releases", releases)
        if target and not isinstance(target, str):
            raise TypeError("Expected argument 'target' to be a str")
        pulumi.set(__self__, "target", target)

    @_builtins.property
    @pulumi.getter
    def commits(self) -> Sequence['outputs.GetProjectTagCommitResult']:
        """
        The commit associated with the tag.
        """
        return pulumi.get(self, "commits")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def message(self) -> _builtins.str:
        """
        The message of the annotated tag.
        """
        return pulumi.get(self, "message")

    @_builtins.property
    @pulumi.getter
    def name(self) -> _builtins.str:
        """
        The name of a tag.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def project(self) -> _builtins.str:
        """
        The ID or URL-encoded path of the project owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @_builtins.property
    @pulumi.getter
    def protected(self) -> _builtins.bool:
        """
        Bool, true if tag has tag protection.
        """
        return pulumi.get(self, "protected")

    @_builtins.property
    @pulumi.getter
    def releases(self) -> Sequence['outputs.GetProjectTagReleaseResult']:
        """
        The release associated with the tag.
        """
        return pulumi.get(self, "releases")

    @_builtins.property
    @pulumi.getter
    def target(self) -> _builtins.str:
        """
        The unique id assigned to the commit by Gitlab.
        """
        return pulumi.get(self, "target")


class AwaitableGetProjectTagResult(GetProjectTagResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectTagResult(
            commits=self.commits,
            id=self.id,
            message=self.message,
            name=self.name,
            project=self.project,
            protected=self.protected,
            releases=self.releases,
            target=self.target)


def get_project_tag(name: Optional[_builtins.str] = None,
                    project: Optional[_builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectTagResult:
    """
    The `ProjectTag` data source allows details of a project tag to be retrieved by its name.

    **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/tags/)


    :param _builtins.str name: The name of a tag.
    :param _builtins.str project: The ID or URL-encoded path of the project owned by the authenticated user.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjectTag:getProjectTag', __args__, opts=opts, typ=GetProjectTagResult).value

    return AwaitableGetProjectTagResult(
        commits=pulumi.get(__ret__, 'commits'),
        id=pulumi.get(__ret__, 'id'),
        message=pulumi.get(__ret__, 'message'),
        name=pulumi.get(__ret__, 'name'),
        project=pulumi.get(__ret__, 'project'),
        protected=pulumi.get(__ret__, 'protected'),
        releases=pulumi.get(__ret__, 'releases'),
        target=pulumi.get(__ret__, 'target'))
def get_project_tag_output(name: Optional[pulumi.Input[_builtins.str]] = None,
                           project: Optional[pulumi.Input[_builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetProjectTagResult]:
    """
    The `ProjectTag` data source allows details of a project tag to be retrieved by its name.

    **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/tags/)


    :param _builtins.str name: The name of a tag.
    :param _builtins.str project: The ID or URL-encoded path of the project owned by the authenticated user.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getProjectTag:getProjectTag', __args__, opts=opts, typ=GetProjectTagResult)
    return __ret__.apply(lambda __response__: GetProjectTagResult(
        commits=pulumi.get(__response__, 'commits'),
        id=pulumi.get(__response__, 'id'),
        message=pulumi.get(__response__, 'message'),
        name=pulumi.get(__response__, 'name'),
        project=pulumi.get(__response__, 'project'),
        protected=pulumi.get(__response__, 'protected'),
        releases=pulumi.get(__response__, 'releases'),
        target=pulumi.get(__response__, 'target')))
