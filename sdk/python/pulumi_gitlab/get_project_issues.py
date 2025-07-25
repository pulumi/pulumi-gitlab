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
    def __init__(__self__, assignee_id=None, assignee_username=None, author_id=None, confidential=None, created_after=None, created_before=None, due_date=None, id=None, iids=None, issue_type=None, issues=None, labels=None, milestone=None, my_reaction_emoji=None, not_assignee_id=None, not_author_id=None, not_labels=None, not_milestone=None, not_my_reaction_emoji=None, order_by=None, project=None, scope=None, search=None, sort=None, updated_after=None, updated_before=None, weight=None, with_labels_details=None):
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
        if not_assignee_id and not isinstance(not_assignee_id, int):
            raise TypeError("Expected argument 'not_assignee_id' to be a int")
        pulumi.set(__self__, "not_assignee_id", not_assignee_id)
        if not_author_id and not isinstance(not_author_id, int):
            raise TypeError("Expected argument 'not_author_id' to be a int")
        pulumi.set(__self__, "not_author_id", not_author_id)
        if not_labels and not isinstance(not_labels, list):
            raise TypeError("Expected argument 'not_labels' to be a list")
        pulumi.set(__self__, "not_labels", not_labels)
        if not_milestone and not isinstance(not_milestone, str):
            raise TypeError("Expected argument 'not_milestone' to be a str")
        pulumi.set(__self__, "not_milestone", not_milestone)
        if not_my_reaction_emoji and not isinstance(not_my_reaction_emoji, str):
            raise TypeError("Expected argument 'not_my_reaction_emoji' to be a str")
        pulumi.set(__self__, "not_my_reaction_emoji", not_my_reaction_emoji)
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

    @_builtins.property
    @pulumi.getter(name="assigneeId")
    def assignee_id(self) -> Optional[_builtins.int]:
        """
        Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
        """
        return pulumi.get(self, "assignee_id")

    @_builtins.property
    @pulumi.getter(name="assigneeUsername")
    def assignee_username(self) -> Optional[_builtins.str]:
        """
        Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
        """
        return pulumi.get(self, "assignee_username")

    @_builtins.property
    @pulumi.getter(name="authorId")
    def author_id(self) -> Optional[_builtins.int]:
        """
        Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
        """
        return pulumi.get(self, "author_id")

    @_builtins.property
    @pulumi.getter
    def confidential(self) -> Optional[_builtins.bool]:
        """
        Filter confidential or public issues.
        """
        return pulumi.get(self, "confidential")

    @_builtins.property
    @pulumi.getter(name="createdAfter")
    def created_after(self) -> Optional[_builtins.str]:
        """
        Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        """
        return pulumi.get(self, "created_after")

    @_builtins.property
    @pulumi.getter(name="createdBefore")
    def created_before(self) -> Optional[_builtins.str]:
        """
        Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        """
        return pulumi.get(self, "created_before")

    @_builtins.property
    @pulumi.getter(name="dueDate")
    def due_date(self) -> Optional[_builtins.str]:
        """
        Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
        """
        return pulumi.get(self, "due_date")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def iids(self) -> Optional[Sequence[_builtins.int]]:
        """
        Return only the issues having the given iid
        """
        return pulumi.get(self, "iids")

    @_builtins.property
    @pulumi.getter(name="issueType")
    def issue_type(self) -> Optional[_builtins.str]:
        """
        Filter to a given type of issue. Valid values are [issue incident test_case].
        """
        return pulumi.get(self, "issue_type")

    @_builtins.property
    @pulumi.getter
    def issues(self) -> Sequence['outputs.GetProjectIssuesIssueResult']:
        """
        The list of issues returned by the search.
        """
        return pulumi.get(self, "issues")

    @_builtins.property
    @pulumi.getter
    def labels(self) -> Optional[Sequence[_builtins.str]]:
        """
        Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
        """
        return pulumi.get(self, "labels")

    @_builtins.property
    @pulumi.getter
    def milestone(self) -> Optional[_builtins.str]:
        """
        The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
        """
        return pulumi.get(self, "milestone")

    @_builtins.property
    @pulumi.getter(name="myReactionEmoji")
    def my_reaction_emoji(self) -> Optional[_builtins.str]:
        """
        Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
        """
        return pulumi.get(self, "my_reaction_emoji")

    @_builtins.property
    @pulumi.getter(name="notAssigneeId")
    def not_assignee_id(self) -> Optional[_builtins.int]:
        """
        Return issues that do not match the assignee id.
        """
        return pulumi.get(self, "not_assignee_id")

    @_builtins.property
    @pulumi.getter(name="notAuthorId")
    def not_author_id(self) -> Optional[_builtins.int]:
        """
        Return issues that do not match the author id.
        """
        return pulumi.get(self, "not_author_id")

    @_builtins.property
    @pulumi.getter(name="notLabels")
    def not_labels(self) -> Optional[Sequence[_builtins.str]]:
        """
        Return issues that do not match the labels.
        """
        return pulumi.get(self, "not_labels")

    @_builtins.property
    @pulumi.getter(name="notMilestone")
    def not_milestone(self) -> Optional[_builtins.str]:
        """
        Return issues that do not match the milestone.
        """
        return pulumi.get(self, "not_milestone")

    @_builtins.property
    @pulumi.getter(name="notMyReactionEmoji")
    def not_my_reaction_emoji(self) -> Optional[_builtins.str]:
        """
        Return issues not reacted by the authenticated user by the given emoji.
        """
        return pulumi.get(self, "not_my_reaction_emoji")

    @_builtins.property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[_builtins.str]:
        """
        Return issues ordered by. Valid values are `created_at`, `updated_at`, `priority`, `due_date`, `relative_position`, `label_priority`, `milestone_due`, `popularity`, `weight`. Default is created_at
        """
        return pulumi.get(self, "order_by")

    @_builtins.property
    @pulumi.getter
    def project(self) -> _builtins.str:
        """
        The name or id of the project.
        """
        return pulumi.get(self, "project")

    @_builtins.property
    @pulumi.getter
    def scope(self) -> Optional[_builtins.str]:
        """
        Return issues for the given scope. Valid values are `created_by_me`, `assigned_to_me`, `all`. Defaults to all.
        """
        return pulumi.get(self, "scope")

    @_builtins.property
    @pulumi.getter
    def search(self) -> Optional[_builtins.str]:
        """
        Search project issues against their title and description
        """
        return pulumi.get(self, "search")

    @_builtins.property
    @pulumi.getter
    def sort(self) -> Optional[_builtins.str]:
        """
        Return issues sorted in asc or desc order. Default is desc
        """
        return pulumi.get(self, "sort")

    @_builtins.property
    @pulumi.getter(name="updatedAfter")
    def updated_after(self) -> Optional[_builtins.str]:
        """
        Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        """
        return pulumi.get(self, "updated_after")

    @_builtins.property
    @pulumi.getter(name="updatedBefore")
    def updated_before(self) -> Optional[_builtins.str]:
        """
        Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        """
        return pulumi.get(self, "updated_before")

    @_builtins.property
    @pulumi.getter
    def weight(self) -> Optional[_builtins.int]:
        """
        Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
        """
        return pulumi.get(self, "weight")

    @_builtins.property
    @pulumi.getter(name="withLabelsDetails")
    def with_labels_details(self) -> Optional[_builtins.bool]:
        """
        If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false.
        """
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
            not_assignee_id=self.not_assignee_id,
            not_author_id=self.not_author_id,
            not_labels=self.not_labels,
            not_milestone=self.not_milestone,
            not_my_reaction_emoji=self.not_my_reaction_emoji,
            order_by=self.order_by,
            project=self.project,
            scope=self.scope,
            search=self.search,
            sort=self.sort,
            updated_after=self.updated_after,
            updated_before=self.updated_before,
            weight=self.weight,
            with_labels_details=self.with_labels_details)


def get_project_issues(assignee_id: Optional[_builtins.int] = None,
                       assignee_username: Optional[_builtins.str] = None,
                       author_id: Optional[_builtins.int] = None,
                       confidential: Optional[_builtins.bool] = None,
                       created_after: Optional[_builtins.str] = None,
                       created_before: Optional[_builtins.str] = None,
                       due_date: Optional[_builtins.str] = None,
                       iids: Optional[Sequence[_builtins.int]] = None,
                       issue_type: Optional[_builtins.str] = None,
                       labels: Optional[Sequence[_builtins.str]] = None,
                       milestone: Optional[_builtins.str] = None,
                       my_reaction_emoji: Optional[_builtins.str] = None,
                       not_assignee_id: Optional[_builtins.int] = None,
                       not_author_id: Optional[_builtins.int] = None,
                       not_labels: Optional[Sequence[_builtins.str]] = None,
                       not_milestone: Optional[_builtins.str] = None,
                       not_my_reaction_emoji: Optional[_builtins.str] = None,
                       order_by: Optional[_builtins.str] = None,
                       project: Optional[_builtins.str] = None,
                       scope: Optional[_builtins.str] = None,
                       search: Optional[_builtins.str] = None,
                       sort: Optional[_builtins.str] = None,
                       updated_after: Optional[_builtins.str] = None,
                       updated_before: Optional[_builtins.str] = None,
                       weight: Optional[_builtins.int] = None,
                       with_labels_details: Optional[_builtins.bool] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectIssuesResult:
    """
    The `get_project_issues` data source allows to retrieve details about issues in a project.

    **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/issues/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    foo = gitlab.get_project(id="foo/bar/baz")
    all_with_foo = gitlab.get_project_issues(project=foo.id,
        search="foo")
    ```


    :param _builtins.int assignee_id: Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
    :param _builtins.str assignee_username: Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
    :param _builtins.int author_id: Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
    :param _builtins.bool confidential: Filter confidential or public issues.
    :param _builtins.str created_after: Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
    :param _builtins.str created_before: Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
    :param _builtins.str due_date: Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
    :param Sequence[_builtins.int] iids: Return only the issues having the given iid
    :param _builtins.str issue_type: Filter to a given type of issue. Valid values are [issue incident test_case].
    :param Sequence[_builtins.str] labels: Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
    :param _builtins.str milestone: The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
    :param _builtins.str my_reaction_emoji: Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
    :param _builtins.int not_assignee_id: Return issues that do not match the assignee id.
    :param _builtins.int not_author_id: Return issues that do not match the author id.
    :param Sequence[_builtins.str] not_labels: Return issues that do not match the labels.
    :param _builtins.str not_milestone: Return issues that do not match the milestone.
    :param _builtins.str not_my_reaction_emoji: Return issues not reacted by the authenticated user by the given emoji.
    :param _builtins.str order_by: Return issues ordered by. Valid values are `created_at`, `updated_at`, `priority`, `due_date`, `relative_position`, `label_priority`, `milestone_due`, `popularity`, `weight`. Default is created_at
    :param _builtins.str project: The name or id of the project.
    :param _builtins.str scope: Return issues for the given scope. Valid values are `created_by_me`, `assigned_to_me`, `all`. Defaults to all.
    :param _builtins.str search: Search project issues against their title and description
    :param _builtins.str sort: Return issues sorted in asc or desc order. Default is desc
    :param _builtins.str updated_after: Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
    :param _builtins.str updated_before: Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
    :param _builtins.int weight: Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
    :param _builtins.bool with_labels_details: If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false.
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
    __args__['notAssigneeId'] = not_assignee_id
    __args__['notAuthorId'] = not_author_id
    __args__['notLabels'] = not_labels
    __args__['notMilestone'] = not_milestone
    __args__['notMyReactionEmoji'] = not_my_reaction_emoji
    __args__['orderBy'] = order_by
    __args__['project'] = project
    __args__['scope'] = scope
    __args__['search'] = search
    __args__['sort'] = sort
    __args__['updatedAfter'] = updated_after
    __args__['updatedBefore'] = updated_before
    __args__['weight'] = weight
    __args__['withLabelsDetails'] = with_labels_details
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjectIssues:getProjectIssues', __args__, opts=opts, typ=GetProjectIssuesResult).value

    return AwaitableGetProjectIssuesResult(
        assignee_id=pulumi.get(__ret__, 'assignee_id'),
        assignee_username=pulumi.get(__ret__, 'assignee_username'),
        author_id=pulumi.get(__ret__, 'author_id'),
        confidential=pulumi.get(__ret__, 'confidential'),
        created_after=pulumi.get(__ret__, 'created_after'),
        created_before=pulumi.get(__ret__, 'created_before'),
        due_date=pulumi.get(__ret__, 'due_date'),
        id=pulumi.get(__ret__, 'id'),
        iids=pulumi.get(__ret__, 'iids'),
        issue_type=pulumi.get(__ret__, 'issue_type'),
        issues=pulumi.get(__ret__, 'issues'),
        labels=pulumi.get(__ret__, 'labels'),
        milestone=pulumi.get(__ret__, 'milestone'),
        my_reaction_emoji=pulumi.get(__ret__, 'my_reaction_emoji'),
        not_assignee_id=pulumi.get(__ret__, 'not_assignee_id'),
        not_author_id=pulumi.get(__ret__, 'not_author_id'),
        not_labels=pulumi.get(__ret__, 'not_labels'),
        not_milestone=pulumi.get(__ret__, 'not_milestone'),
        not_my_reaction_emoji=pulumi.get(__ret__, 'not_my_reaction_emoji'),
        order_by=pulumi.get(__ret__, 'order_by'),
        project=pulumi.get(__ret__, 'project'),
        scope=pulumi.get(__ret__, 'scope'),
        search=pulumi.get(__ret__, 'search'),
        sort=pulumi.get(__ret__, 'sort'),
        updated_after=pulumi.get(__ret__, 'updated_after'),
        updated_before=pulumi.get(__ret__, 'updated_before'),
        weight=pulumi.get(__ret__, 'weight'),
        with_labels_details=pulumi.get(__ret__, 'with_labels_details'))
def get_project_issues_output(assignee_id: Optional[pulumi.Input[Optional[_builtins.int]]] = None,
                              assignee_username: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              author_id: Optional[pulumi.Input[Optional[_builtins.int]]] = None,
                              confidential: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                              created_after: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              created_before: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              due_date: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              iids: Optional[pulumi.Input[Optional[Sequence[_builtins.int]]]] = None,
                              issue_type: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              labels: Optional[pulumi.Input[Optional[Sequence[_builtins.str]]]] = None,
                              milestone: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              my_reaction_emoji: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              not_assignee_id: Optional[pulumi.Input[Optional[_builtins.int]]] = None,
                              not_author_id: Optional[pulumi.Input[Optional[_builtins.int]]] = None,
                              not_labels: Optional[pulumi.Input[Optional[Sequence[_builtins.str]]]] = None,
                              not_milestone: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              not_my_reaction_emoji: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              order_by: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              project: Optional[pulumi.Input[_builtins.str]] = None,
                              scope: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              search: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              sort: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              updated_after: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              updated_before: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              weight: Optional[pulumi.Input[Optional[_builtins.int]]] = None,
                              with_labels_details: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetProjectIssuesResult]:
    """
    The `get_project_issues` data source allows to retrieve details about issues in a project.

    **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/issues/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    foo = gitlab.get_project(id="foo/bar/baz")
    all_with_foo = gitlab.get_project_issues(project=foo.id,
        search="foo")
    ```


    :param _builtins.int assignee_id: Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
    :param _builtins.str assignee_username: Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
    :param _builtins.int author_id: Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
    :param _builtins.bool confidential: Filter confidential or public issues.
    :param _builtins.str created_after: Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
    :param _builtins.str created_before: Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
    :param _builtins.str due_date: Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
    :param Sequence[_builtins.int] iids: Return only the issues having the given iid
    :param _builtins.str issue_type: Filter to a given type of issue. Valid values are [issue incident test_case].
    :param Sequence[_builtins.str] labels: Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
    :param _builtins.str milestone: The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
    :param _builtins.str my_reaction_emoji: Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
    :param _builtins.int not_assignee_id: Return issues that do not match the assignee id.
    :param _builtins.int not_author_id: Return issues that do not match the author id.
    :param Sequence[_builtins.str] not_labels: Return issues that do not match the labels.
    :param _builtins.str not_milestone: Return issues that do not match the milestone.
    :param _builtins.str not_my_reaction_emoji: Return issues not reacted by the authenticated user by the given emoji.
    :param _builtins.str order_by: Return issues ordered by. Valid values are `created_at`, `updated_at`, `priority`, `due_date`, `relative_position`, `label_priority`, `milestone_due`, `popularity`, `weight`. Default is created_at
    :param _builtins.str project: The name or id of the project.
    :param _builtins.str scope: Return issues for the given scope. Valid values are `created_by_me`, `assigned_to_me`, `all`. Defaults to all.
    :param _builtins.str search: Search project issues against their title and description
    :param _builtins.str sort: Return issues sorted in asc or desc order. Default is desc
    :param _builtins.str updated_after: Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
    :param _builtins.str updated_before: Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
    :param _builtins.int weight: Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
    :param _builtins.bool with_labels_details: If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false.
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
    __args__['notAssigneeId'] = not_assignee_id
    __args__['notAuthorId'] = not_author_id
    __args__['notLabels'] = not_labels
    __args__['notMilestone'] = not_milestone
    __args__['notMyReactionEmoji'] = not_my_reaction_emoji
    __args__['orderBy'] = order_by
    __args__['project'] = project
    __args__['scope'] = scope
    __args__['search'] = search
    __args__['sort'] = sort
    __args__['updatedAfter'] = updated_after
    __args__['updatedBefore'] = updated_before
    __args__['weight'] = weight
    __args__['withLabelsDetails'] = with_labels_details
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getProjectIssues:getProjectIssues', __args__, opts=opts, typ=GetProjectIssuesResult)
    return __ret__.apply(lambda __response__: GetProjectIssuesResult(
        assignee_id=pulumi.get(__response__, 'assignee_id'),
        assignee_username=pulumi.get(__response__, 'assignee_username'),
        author_id=pulumi.get(__response__, 'author_id'),
        confidential=pulumi.get(__response__, 'confidential'),
        created_after=pulumi.get(__response__, 'created_after'),
        created_before=pulumi.get(__response__, 'created_before'),
        due_date=pulumi.get(__response__, 'due_date'),
        id=pulumi.get(__response__, 'id'),
        iids=pulumi.get(__response__, 'iids'),
        issue_type=pulumi.get(__response__, 'issue_type'),
        issues=pulumi.get(__response__, 'issues'),
        labels=pulumi.get(__response__, 'labels'),
        milestone=pulumi.get(__response__, 'milestone'),
        my_reaction_emoji=pulumi.get(__response__, 'my_reaction_emoji'),
        not_assignee_id=pulumi.get(__response__, 'not_assignee_id'),
        not_author_id=pulumi.get(__response__, 'not_author_id'),
        not_labels=pulumi.get(__response__, 'not_labels'),
        not_milestone=pulumi.get(__response__, 'not_milestone'),
        not_my_reaction_emoji=pulumi.get(__response__, 'not_my_reaction_emoji'),
        order_by=pulumi.get(__response__, 'order_by'),
        project=pulumi.get(__response__, 'project'),
        scope=pulumi.get(__response__, 'scope'),
        search=pulumi.get(__response__, 'search'),
        sort=pulumi.get(__response__, 'sort'),
        updated_after=pulumi.get(__response__, 'updated_after'),
        updated_before=pulumi.get(__response__, 'updated_before'),
        weight=pulumi.get(__response__, 'weight'),
        with_labels_details=pulumi.get(__response__, 'with_labels_details')))
