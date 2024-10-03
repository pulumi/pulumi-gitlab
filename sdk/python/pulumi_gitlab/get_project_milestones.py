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
    'GetProjectMilestonesResult',
    'AwaitableGetProjectMilestonesResult',
    'get_project_milestones',
    'get_project_milestones_output',
]

@pulumi.output_type
class GetProjectMilestonesResult:
    """
    A collection of values returned by getProjectMilestones.
    """
    def __init__(__self__, id=None, iids=None, include_parent_milestones=None, milestones=None, project=None, search=None, state=None, title=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iids and not isinstance(iids, list):
            raise TypeError("Expected argument 'iids' to be a list")
        pulumi.set(__self__, "iids", iids)
        if include_parent_milestones and not isinstance(include_parent_milestones, bool):
            raise TypeError("Expected argument 'include_parent_milestones' to be a bool")
        pulumi.set(__self__, "include_parent_milestones", include_parent_milestones)
        if milestones and not isinstance(milestones, list):
            raise TypeError("Expected argument 'milestones' to be a list")
        pulumi.set(__self__, "milestones", milestones)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if search and not isinstance(search, str):
            raise TypeError("Expected argument 'search' to be a str")
        pulumi.set(__self__, "search", search)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def iids(self) -> Optional[Sequence[int]]:
        """
        Return only the milestones having the given `iid` (Note: ignored if `include_parent_milestones` is set as `true`).
        """
        return pulumi.get(self, "iids")

    @property
    @pulumi.getter(name="includeParentMilestones")
    def include_parent_milestones(self) -> Optional[bool]:
        """
        Include group milestones from parent group and its ancestors. Introduced in GitLab 13.4.
        """
        return pulumi.get(self, "include_parent_milestones")

    @property
    @pulumi.getter
    def milestones(self) -> Sequence['outputs.GetProjectMilestonesMilestoneResult']:
        """
        List of milestones from a project.
        """
        return pulumi.get(self, "milestones")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID or URL-encoded path of the project owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def search(self) -> Optional[str]:
        """
        Return only milestones with a title or description matching the provided string.
        """
        return pulumi.get(self, "search")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        Return only `active` or `closed` milestones.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def title(self) -> Optional[str]:
        """
        Return only the milestones having the given `title`.
        """
        return pulumi.get(self, "title")


class AwaitableGetProjectMilestonesResult(GetProjectMilestonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectMilestonesResult(
            id=self.id,
            iids=self.iids,
            include_parent_milestones=self.include_parent_milestones,
            milestones=self.milestones,
            project=self.project,
            search=self.search,
            state=self.state,
            title=self.title)


def get_project_milestones(iids: Optional[Sequence[int]] = None,
                           include_parent_milestones: Optional[bool] = None,
                           project: Optional[str] = None,
                           search: Optional[str] = None,
                           state: Optional[str] = None,
                           title: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectMilestonesResult:
    """
    The `get_project_milestones` data source allows get details of a project milestones.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)


    :param Sequence[int] iids: Return only the milestones having the given `iid` (Note: ignored if `include_parent_milestones` is set as `true`).
    :param bool include_parent_milestones: Include group milestones from parent group and its ancestors. Introduced in GitLab 13.4.
    :param str project: The ID or URL-encoded path of the project owned by the authenticated user.
    :param str search: Return only milestones with a title or description matching the provided string.
    :param str state: Return only `active` or `closed` milestones.
    :param str title: Return only the milestones having the given `title`.
    """
    __args__ = dict()
    __args__['iids'] = iids
    __args__['includeParentMilestones'] = include_parent_milestones
    __args__['project'] = project
    __args__['search'] = search
    __args__['state'] = state
    __args__['title'] = title
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjectMilestones:getProjectMilestones', __args__, opts=opts, typ=GetProjectMilestonesResult).value

    return AwaitableGetProjectMilestonesResult(
        id=pulumi.get(__ret__, 'id'),
        iids=pulumi.get(__ret__, 'iids'),
        include_parent_milestones=pulumi.get(__ret__, 'include_parent_milestones'),
        milestones=pulumi.get(__ret__, 'milestones'),
        project=pulumi.get(__ret__, 'project'),
        search=pulumi.get(__ret__, 'search'),
        state=pulumi.get(__ret__, 'state'),
        title=pulumi.get(__ret__, 'title'))
def get_project_milestones_output(iids: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                                  include_parent_milestones: Optional[pulumi.Input[Optional[bool]]] = None,
                                  project: Optional[pulumi.Input[str]] = None,
                                  search: Optional[pulumi.Input[Optional[str]]] = None,
                                  state: Optional[pulumi.Input[Optional[str]]] = None,
                                  title: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectMilestonesResult]:
    """
    The `get_project_milestones` data source allows get details of a project milestones.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)


    :param Sequence[int] iids: Return only the milestones having the given `iid` (Note: ignored if `include_parent_milestones` is set as `true`).
    :param bool include_parent_milestones: Include group milestones from parent group and its ancestors. Introduced in GitLab 13.4.
    :param str project: The ID or URL-encoded path of the project owned by the authenticated user.
    :param str search: Return only milestones with a title or description matching the provided string.
    :param str state: Return only `active` or `closed` milestones.
    :param str title: Return only the milestones having the given `title`.
    """
    __args__ = dict()
    __args__['iids'] = iids
    __args__['includeParentMilestones'] = include_parent_milestones
    __args__['project'] = project
    __args__['search'] = search
    __args__['state'] = state
    __args__['title'] = title
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getProjectMilestones:getProjectMilestones', __args__, opts=opts, typ=GetProjectMilestonesResult)
    return __ret__.apply(lambda __response__: GetProjectMilestonesResult(
        id=pulumi.get(__response__, 'id'),
        iids=pulumi.get(__response__, 'iids'),
        include_parent_milestones=pulumi.get(__response__, 'include_parent_milestones'),
        milestones=pulumi.get(__response__, 'milestones'),
        project=pulumi.get(__response__, 'project'),
        search=pulumi.get(__response__, 'search'),
        state=pulumi.get(__response__, 'state'),
        title=pulumi.get(__response__, 'title')))
