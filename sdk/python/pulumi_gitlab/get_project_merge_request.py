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
    'GetProjectMergeRequestResult',
    'AwaitableGetProjectMergeRequestResult',
    'get_project_merge_request',
    'get_project_merge_request_output',
]

@pulumi.output_type
class GetProjectMergeRequestResult:
    """
    A collection of values returned by getProjectMergeRequest.
    """
    def __init__(__self__, assignee=None, assignees=None, author=None, blocking_discussions_resolved=None, changes_count=None, closed_at=None, closed_by=None, created_at=None, id=None, iid=None, project=None):
        if assignee and not isinstance(assignee, dict):
            raise TypeError("Expected argument 'assignee' to be a dict")
        pulumi.set(__self__, "assignee", assignee)
        if assignees and not isinstance(assignees, list):
            raise TypeError("Expected argument 'assignees' to be a list")
        pulumi.set(__self__, "assignees", assignees)
        if author and not isinstance(author, dict):
            raise TypeError("Expected argument 'author' to be a dict")
        pulumi.set(__self__, "author", author)
        if blocking_discussions_resolved and not isinstance(blocking_discussions_resolved, bool):
            raise TypeError("Expected argument 'blocking_discussions_resolved' to be a bool")
        pulumi.set(__self__, "blocking_discussions_resolved", blocking_discussions_resolved)
        if changes_count and not isinstance(changes_count, str):
            raise TypeError("Expected argument 'changes_count' to be a str")
        pulumi.set(__self__, "changes_count", changes_count)
        if closed_at and not isinstance(closed_at, str):
            raise TypeError("Expected argument 'closed_at' to be a str")
        pulumi.set(__self__, "closed_at", closed_at)
        if closed_by and not isinstance(closed_by, dict):
            raise TypeError("Expected argument 'closed_by' to be a dict")
        pulumi.set(__self__, "closed_by", closed_by)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, float):
            raise TypeError("Expected argument 'id' to be a float")
        pulumi.set(__self__, "id", id)
        if iid and not isinstance(iid, float):
            raise TypeError("Expected argument 'iid' to be a float")
        pulumi.set(__self__, "iid", iid)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def assignee(self) -> 'outputs.GetProjectMergeRequestAssigneeResult':
        """
        First assignee of the merge request.
        """
        return pulumi.get(self, "assignee")

    @property
    @pulumi.getter
    def assignees(self) -> Sequence['outputs.GetProjectMergeRequestAssigneeResult']:
        """
        Assignees of the merge request.
        """
        return pulumi.get(self, "assignees")

    @property
    @pulumi.getter
    def author(self) -> 'outputs.GetProjectMergeRequestAuthorResult':
        """
        User who created this merge request.
        """
        return pulumi.get(self, "author")

    @property
    @pulumi.getter(name="blockingDiscussionsResolved")
    def blocking_discussions_resolved(self) -> bool:
        """
        Indicates if all discussions are resolved only if all are
        required before merge request can be merged.
        """
        return pulumi.get(self, "blocking_discussions_resolved")

    @property
    @pulumi.getter(name="changesCount")
    def changes_count(self) -> str:
        """
        Number of changes made on the merge request. Empty when the
        merge request is created, and populates asynchronously.
        """
        return pulumi.get(self, "changes_count")

    @property
    @pulumi.getter(name="closedAt")
    def closed_at(self) -> str:
        """
        Timestamp of when the merge request was closed.
        """
        return pulumi.get(self, "closed_at")

    @property
    @pulumi.getter(name="closedBy")
    def closed_by(self) -> 'outputs.GetProjectMergeRequestClosedByResult':
        """
        User who closed this merge request.
        """
        return pulumi.get(self, "closed_by")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Timestamp of when the merge request was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> float:
        """
        The unique instance level ID of the merge request.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def iid(self) -> float:
        """
        The unique project level ID of the merge request.
        """
        return pulumi.get(self, "iid")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID or path of the project.
        """
        return pulumi.get(self, "project")


class AwaitableGetProjectMergeRequestResult(GetProjectMergeRequestResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectMergeRequestResult(
            assignee=self.assignee,
            assignees=self.assignees,
            author=self.author,
            blocking_discussions_resolved=self.blocking_discussions_resolved,
            changes_count=self.changes_count,
            closed_at=self.closed_at,
            closed_by=self.closed_by,
            created_at=self.created_at,
            id=self.id,
            iid=self.iid,
            project=self.project)


def get_project_merge_request(iid: Optional[float] = None,
                              project: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectMergeRequestResult:
    """
    The `get_project_merge_request` data source retrieves
    information about a single merge request related to a specific project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/merge_requests.html#get-single-mr)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    by_project_id = gitlab.get_project_merge_request(project="123",
        iid=456)
    by_project_name = gitlab.get_project_merge_request(project="company/group/project1",
        iid=3)
    ```


    :param float iid: The unique project level ID of the merge request.
    :param str project: The ID or path of the project.
    """
    __args__ = dict()
    __args__['iid'] = iid
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjectMergeRequest:getProjectMergeRequest', __args__, opts=opts, typ=GetProjectMergeRequestResult).value

    return AwaitableGetProjectMergeRequestResult(
        assignee=pulumi.get(__ret__, 'assignee'),
        assignees=pulumi.get(__ret__, 'assignees'),
        author=pulumi.get(__ret__, 'author'),
        blocking_discussions_resolved=pulumi.get(__ret__, 'blocking_discussions_resolved'),
        changes_count=pulumi.get(__ret__, 'changes_count'),
        closed_at=pulumi.get(__ret__, 'closed_at'),
        closed_by=pulumi.get(__ret__, 'closed_by'),
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        iid=pulumi.get(__ret__, 'iid'),
        project=pulumi.get(__ret__, 'project'))
def get_project_merge_request_output(iid: Optional[pulumi.Input[float]] = None,
                                     project: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectMergeRequestResult]:
    """
    The `get_project_merge_request` data source retrieves
    information about a single merge request related to a specific project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/merge_requests.html#get-single-mr)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    by_project_id = gitlab.get_project_merge_request(project="123",
        iid=456)
    by_project_name = gitlab.get_project_merge_request(project="company/group/project1",
        iid=3)
    ```


    :param float iid: The unique project level ID of the merge request.
    :param str project: The ID or path of the project.
    """
    __args__ = dict()
    __args__['iid'] = iid
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getProjectMergeRequest:getProjectMergeRequest', __args__, opts=opts, typ=GetProjectMergeRequestResult)
    return __ret__.apply(lambda __response__: GetProjectMergeRequestResult(
        assignee=pulumi.get(__response__, 'assignee'),
        assignees=pulumi.get(__response__, 'assignees'),
        author=pulumi.get(__response__, 'author'),
        blocking_discussions_resolved=pulumi.get(__response__, 'blocking_discussions_resolved'),
        changes_count=pulumi.get(__response__, 'changes_count'),
        closed_at=pulumi.get(__response__, 'closed_at'),
        closed_by=pulumi.get(__response__, 'closed_by'),
        created_at=pulumi.get(__response__, 'created_at'),
        id=pulumi.get(__response__, 'id'),
        iid=pulumi.get(__response__, 'iid'),
        project=pulumi.get(__response__, 'project')))