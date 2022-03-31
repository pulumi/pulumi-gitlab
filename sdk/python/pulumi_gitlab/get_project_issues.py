# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetProjectIssuesResult',
    'AwaitableGetProjectIssuesResult',
    'get_project_issues',
    'get_project_issues_output',
]

@pulumi.output_type
class GetProjectIssuesResult:
    """
    A collection of values returned by getProjectIssues.
    """
    def __init__(__self__, assignee_id=None, assignee_username=None, author_id=None, confidential=None, created_after=None, created_before=None, due_date=None, id=None, iids=None, issue_type=None, issues=None, labels=None, milestone=None, my_reaction_emoji=None, not_assignee_ids=None, not_author_ids=None, not_labels=None, not_milestones=None, not_my_reaction_emojis=None, order_by=None, project=None, scope=None, search=None, sort=None, updated_after=None, updated_before=None, weight=None, with_labels_details=None):
        if assignee_id and not isinstance(assignee_id, int):
            raise TypeError("Expected argument 'assignee_id' to be a int")
        pulumi.set(__self__, "assignee_id", assignee_id)
        if assignee_username and not isinstance(assignee_username, str):
            raise TypeError("Expected argument 'assignee_username' to be a str")
        pulumi.set(__self__, "assignee_username", assignee_username)
        if author_id and not isinstance(author_id, int):
            raise TypeError("Expected argument 'author_id' to be a int")
        pulumi.set(__self__, "author_id", author_id)
        if confidential and not isinstance(confidential, bool):
            raise TypeError("Expected argument 'confidential' to be a bool")
        pulumi.set(__self__, "confidential", confidential)
        if created_after and not isinstance(created_after, str):
            raise TypeError("Expected argument 'created_after' to be a str")
        pulumi.set(__self__, "created_after", created_after)
        if created_before and not isinstance(created_before, str):
            raise TypeError("Expected argument 'created_before' to be a str")
        pulumi.set(__self__, "created_before", created_before)
        if due_date and not isinstance(due_date, str):
            raise TypeError("Expected argument 'due_date' to be a str")
        pulumi.set(__self__, "due_date", due_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iids and not isinstance(iids, list):
            raise TypeError("Expected argument 'iids' to be a list")
        pulumi.set(__self__, "iids", iids)
        if issue_type and not isinstance(issue_type, str):
            raise TypeError("Expected argument 'issue_type' to be a str")
        pulumi.set(__self__, "issue_type", issue_type)
        if issues and not isinstance(issues, list):
            raise TypeError("Expected argument 'issues' to be a list")
        pulumi.set(__self__, "issues", issues)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if milestone and not isinstance(milestone, str):
            raise TypeError("Expected argument 'milestone' to be a str")
        pulumi.set(__self__, "milestone", milestone)
        if my_reaction_emoji and not isinstance(my_reaction_emoji, str):
            raise TypeError("Expected argument 'my_reaction_emoji' to be a str")
        pulumi.set(__self__, "my_reaction_emoji", my_reaction_emoji)
        if not_assignee_ids and not isinstance(not_assignee_ids, list):
            raise TypeError("Expected argument 'not_assignee_ids' to be a list")
        pulumi.set(__self__, "not_assignee_ids", not_assignee_ids)
        if not_author_ids and not isinstance(not_author_ids, list):
            raise TypeError("Expected argument 'not_author_ids' to be a list")
        pulumi.set(__self__, "not_author_ids", not_author_ids)
        if not_labels and not isinstance(not_labels, list):
            raise TypeError("Expected argument 'not_labels' to be a list")
        pulumi.set(__self__, "not_labels", not_labels)
        if not_milestones and not isinstance(not_milestones, list):
            raise TypeError("Expected argument 'not_milestones' to be a list")
        pulumi.set(__self__, "not_milestones", not_milestones)
        if not_my_reaction_emojis and not isinstance(not_my_reaction_emojis, list):
            raise TypeError("Expected argument 'not_my_reaction_emojis' to be a list")
        pulumi.set(__self__, "not_my_reaction_emojis", not_my_reaction_emojis)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        pulumi.set(__self__, "scope", scope)
        if search and not isinstance(search, str):
            raise TypeError("Expected argument 'search' to be a str")
        pulumi.set(__self__, "search", search)
        if sort and not isinstance(sort, str):
            raise TypeError("Expected argument 'sort' to be a str")
        pulumi.set(__self__, "sort", sort)
        if updated_after and not isinstance(updated_after, str):
            raise TypeError("Expected argument 'updated_after' to be a str")
        pulumi.set(__self__, "updated_after", updated_after)
        if updated_before and not isinstance(updated_before, str):
            raise TypeError("Expected argument 'updated_before' to be a str")
        pulumi.set(__self__, "updated_before", updated_before)
        if weight and not isinstance(weight, int):
            raise TypeError("Expected argument 'weight' to be a int")
        pulumi.set(__self__, "weight", weight)
        if with_labels_details and not isinstance(with_labels_details, bool):
            raise TypeError("Expected argument 'with_labels_details' to be a bool")
        pulumi.set(__self__, "with_labels_details", with_labels_details)

    @property
    @pulumi.getter(name="assigneeId")
    def assignee_id(self) -> Optional[int]:
        return pulumi.get(self, "assignee_id")

    @property
    @pulumi.getter(name="assigneeUsername")
    def assignee_username(self) -> Optional[str]:
        return pulumi.get(self, "assignee_username")

    @property
    @pulumi.getter(name="authorId")
    def author_id(self) -> Optional[int]:
        return pulumi.get(self, "author_id")

    @property
    @pulumi.getter
    def confidential(self) -> Optional[bool]:
        return pulumi.get(self, "confidential")

    @property
    @pulumi.getter(name="createdAfter")
    def created_after(self) -> Optional[str]:
        return pulumi.get(self, "created_after")

    @property
    @pulumi.getter(name="createdBefore")
    def created_before(self) -> Optional[str]:
        return pulumi.get(self, "created_before")

    @property
    @pulumi.getter(name="dueDate")
    def due_date(self) -> Optional[str]:
        return pulumi.get(self, "due_date")

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
        return pulumi.get(self, "iids")

    @property
    @pulumi.getter(name="issueType")
    def issue_type(self) -> Optional[str]:
        return pulumi.get(self, "issue_type")

    @property
    @pulumi.getter
    def issues(self) -> Sequence['outputs.GetProjectIssuesIssueResult']:
        return pulumi.get(self, "issues")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def milestone(self) -> Optional[str]:
        return pulumi.get(self, "milestone")

    @property
    @pulumi.getter(name="myReactionEmoji")
    def my_reaction_emoji(self) -> Optional[str]:
        return pulumi.get(self, "my_reaction_emoji")

    @property
    @pulumi.getter(name="notAssigneeIds")
    def not_assignee_ids(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "not_assignee_ids")

    @property
    @pulumi.getter(name="notAuthorIds")
    def not_author_ids(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "not_author_ids")

    @property
    @pulumi.getter(name="notLabels")
    def not_labels(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "not_labels")

    @property
    @pulumi.getter(name="notMilestones")
    def not_milestones(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "not_milestones")

    @property
    @pulumi.getter(name="notMyReactionEmojis")
    def not_my_reaction_emojis(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "not_my_reaction_emojis")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def scope(self) -> Optional[str]:
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def search(self) -> Optional[str]:
        return pulumi.get(self, "search")

    @property
    @pulumi.getter
    def sort(self) -> Optional[str]:
        return pulumi.get(self, "sort")

    @property
    @pulumi.getter(name="updatedAfter")
    def updated_after(self) -> Optional[str]:
        return pulumi.get(self, "updated_after")

    @property
    @pulumi.getter(name="updatedBefore")
    def updated_before(self) -> Optional[str]:
        return pulumi.get(self, "updated_before")

    @property
    @pulumi.getter
    def weight(self) -> Optional[int]:
        return pulumi.get(self, "weight")

    @property
    @pulumi.getter(name="withLabelsDetails")
    def with_labels_details(self) -> Optional[bool]:
        return pulumi.get(self, "with_labels_details")


class AwaitableGetProjectIssuesResult(GetProjectIssuesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectIssuesResult(
            assignee_id=self.assignee_id,
            assignee_username=self.assignee_username,
            author_id=self.author_id,
            confidential=self.confidential,
            created_after=self.created_after,
            created_before=self.created_before,
            due_date=self.due_date,
            id=self.id,
            iids=self.iids,
            issue_type=self.issue_type,
            issues=self.issues,
            labels=self.labels,
            milestone=self.milestone,
            my_reaction_emoji=self.my_reaction_emoji,
            not_assignee_ids=self.not_assignee_ids,
            not_author_ids=self.not_author_ids,
            not_labels=self.not_labels,
            not_milestones=self.not_milestones,
            not_my_reaction_emojis=self.not_my_reaction_emojis,
            order_by=self.order_by,
            project=self.project,
            scope=self.scope,
            search=self.search,
            sort=self.sort,
            updated_after=self.updated_after,
            updated_before=self.updated_before,
            weight=self.weight,
            with_labels_details=self.with_labels_details)


def get_project_issues(assignee_id: Optional[int] = None,
                       assignee_username: Optional[str] = None,
                       author_id: Optional[int] = None,
                       confidential: Optional[bool] = None,
                       created_after: Optional[str] = None,
                       created_before: Optional[str] = None,
                       due_date: Optional[str] = None,
                       iids: Optional[Sequence[int]] = None,
                       issue_type: Optional[str] = None,
                       labels: Optional[Sequence[str]] = None,
                       milestone: Optional[str] = None,
                       my_reaction_emoji: Optional[str] = None,
                       not_assignee_ids: Optional[Sequence[int]] = None,
                       not_author_ids: Optional[Sequence[int]] = None,
                       not_labels: Optional[Sequence[str]] = None,
                       not_milestones: Optional[Sequence[str]] = None,
                       not_my_reaction_emojis: Optional[Sequence[str]] = None,
                       order_by: Optional[str] = None,
                       project: Optional[str] = None,
                       scope: Optional[str] = None,
                       search: Optional[str] = None,
                       sort: Optional[str] = None,
                       updated_after: Optional[str] = None,
                       updated_before: Optional[str] = None,
                       weight: Optional[int] = None,
                       with_labels_details: Optional[bool] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectIssuesResult:
    """
    The `get_project_issues` data source allows to retrieve details about issues in a project.

    **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/issues.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    foo = gitlab.get_project(id="foo/bar/baz")
    all_with_foo = gitlab.get_project_issues(project=foo.id,
        search="foo")
    ```
    """
    __args__ = dict()
    __args__['assigneeId'] = assignee_id
    __args__['assigneeUsername'] = assignee_username
    __args__['authorId'] = author_id
    __args__['confidential'] = confidential
    __args__['createdAfter'] = created_after
    __args__['createdBefore'] = created_before
    __args__['dueDate'] = due_date
    __args__['iids'] = iids
    __args__['issueType'] = issue_type
    __args__['labels'] = labels
    __args__['milestone'] = milestone
    __args__['myReactionEmoji'] = my_reaction_emoji
    __args__['notAssigneeIds'] = not_assignee_ids
    __args__['notAuthorIds'] = not_author_ids
    __args__['notLabels'] = not_labels
    __args__['notMilestones'] = not_milestones
    __args__['notMyReactionEmojis'] = not_my_reaction_emojis
    __args__['orderBy'] = order_by
    __args__['project'] = project
    __args__['scope'] = scope
    __args__['search'] = search
    __args__['sort'] = sort
    __args__['updatedAfter'] = updated_after
    __args__['updatedBefore'] = updated_before
    __args__['weight'] = weight
    __args__['withLabelsDetails'] = with_labels_details
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjectIssues:getProjectIssues', __args__, opts=opts, typ=GetProjectIssuesResult).value

    return AwaitableGetProjectIssuesResult(
        assignee_id=__ret__.assignee_id,
        assignee_username=__ret__.assignee_username,
        author_id=__ret__.author_id,
        confidential=__ret__.confidential,
        created_after=__ret__.created_after,
        created_before=__ret__.created_before,
        due_date=__ret__.due_date,
        id=__ret__.id,
        iids=__ret__.iids,
        issue_type=__ret__.issue_type,
        issues=__ret__.issues,
        labels=__ret__.labels,
        milestone=__ret__.milestone,
        my_reaction_emoji=__ret__.my_reaction_emoji,
        not_assignee_ids=__ret__.not_assignee_ids,
        not_author_ids=__ret__.not_author_ids,
        not_labels=__ret__.not_labels,
        not_milestones=__ret__.not_milestones,
        not_my_reaction_emojis=__ret__.not_my_reaction_emojis,
        order_by=__ret__.order_by,
        project=__ret__.project,
        scope=__ret__.scope,
        search=__ret__.search,
        sort=__ret__.sort,
        updated_after=__ret__.updated_after,
        updated_before=__ret__.updated_before,
        weight=__ret__.weight,
        with_labels_details=__ret__.with_labels_details)


@_utilities.lift_output_func(get_project_issues)
def get_project_issues_output(assignee_id: Optional[pulumi.Input[Optional[int]]] = None,
                              assignee_username: Optional[pulumi.Input[Optional[str]]] = None,
                              author_id: Optional[pulumi.Input[Optional[int]]] = None,
                              confidential: Optional[pulumi.Input[Optional[bool]]] = None,
                              created_after: Optional[pulumi.Input[Optional[str]]] = None,
                              created_before: Optional[pulumi.Input[Optional[str]]] = None,
                              due_date: Optional[pulumi.Input[Optional[str]]] = None,
                              iids: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                              issue_type: Optional[pulumi.Input[Optional[str]]] = None,
                              labels: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              milestone: Optional[pulumi.Input[Optional[str]]] = None,
                              my_reaction_emoji: Optional[pulumi.Input[Optional[str]]] = None,
                              not_assignee_ids: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                              not_author_ids: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                              not_labels: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              not_milestones: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              not_my_reaction_emojis: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              order_by: Optional[pulumi.Input[Optional[str]]] = None,
                              project: Optional[pulumi.Input[str]] = None,
                              scope: Optional[pulumi.Input[Optional[str]]] = None,
                              search: Optional[pulumi.Input[Optional[str]]] = None,
                              sort: Optional[pulumi.Input[Optional[str]]] = None,
                              updated_after: Optional[pulumi.Input[Optional[str]]] = None,
                              updated_before: Optional[pulumi.Input[Optional[str]]] = None,
                              weight: Optional[pulumi.Input[Optional[int]]] = None,
                              with_labels_details: Optional[pulumi.Input[Optional[bool]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectIssuesResult]:
    """
    The `get_project_issues` data source allows to retrieve details about issues in a project.

    **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/issues.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    foo = gitlab.get_project(id="foo/bar/baz")
    all_with_foo = gitlab.get_project_issues(project=foo.id,
        search="foo")
    ```
    """
    ...
