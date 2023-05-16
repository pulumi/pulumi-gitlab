# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetProjectIssueResult',
    'AwaitableGetProjectIssueResult',
    'get_project_issue',
    'get_project_issue_output',
]

@pulumi.output_type
class GetProjectIssueResult:
    """
    A collection of values returned by getProjectIssue.
    """
    def __init__(__self__, assignee_ids=None, author_id=None, closed_at=None, closed_by_user_id=None, confidential=None, created_at=None, description=None, discussion_locked=None, discussion_to_resolve=None, downvotes=None, due_date=None, epic_id=None, epic_issue_id=None, external_id=None, human_time_estimate=None, human_total_time_spent=None, id=None, iid=None, issue_id=None, issue_link_id=None, issue_type=None, labels=None, links=None, merge_request_to_resolve_discussions_of=None, merge_requests_count=None, milestone_id=None, moved_to_id=None, project=None, references=None, state=None, subscribed=None, task_completion_statuses=None, time_estimate=None, title=None, total_time_spent=None, updated_at=None, upvotes=None, user_notes_count=None, web_url=None, weight=None):
        if assignee_ids and not isinstance(assignee_ids, list):
            raise TypeError("Expected argument 'assignee_ids' to be a list")
        pulumi.set(__self__, "assignee_ids", assignee_ids)
        if author_id and not isinstance(author_id, int):
            raise TypeError("Expected argument 'author_id' to be a int")
        pulumi.set(__self__, "author_id", author_id)
        if closed_at and not isinstance(closed_at, str):
            raise TypeError("Expected argument 'closed_at' to be a str")
        pulumi.set(__self__, "closed_at", closed_at)
        if closed_by_user_id and not isinstance(closed_by_user_id, int):
            raise TypeError("Expected argument 'closed_by_user_id' to be a int")
        pulumi.set(__self__, "closed_by_user_id", closed_by_user_id)
        if confidential and not isinstance(confidential, bool):
            raise TypeError("Expected argument 'confidential' to be a bool")
        pulumi.set(__self__, "confidential", confidential)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if discussion_locked and not isinstance(discussion_locked, bool):
            raise TypeError("Expected argument 'discussion_locked' to be a bool")
        pulumi.set(__self__, "discussion_locked", discussion_locked)
        if discussion_to_resolve and not isinstance(discussion_to_resolve, str):
            raise TypeError("Expected argument 'discussion_to_resolve' to be a str")
        pulumi.set(__self__, "discussion_to_resolve", discussion_to_resolve)
        if downvotes and not isinstance(downvotes, int):
            raise TypeError("Expected argument 'downvotes' to be a int")
        pulumi.set(__self__, "downvotes", downvotes)
        if due_date and not isinstance(due_date, str):
            raise TypeError("Expected argument 'due_date' to be a str")
        pulumi.set(__self__, "due_date", due_date)
        if epic_id and not isinstance(epic_id, int):
            raise TypeError("Expected argument 'epic_id' to be a int")
        pulumi.set(__self__, "epic_id", epic_id)
        if epic_issue_id and not isinstance(epic_issue_id, int):
            raise TypeError("Expected argument 'epic_issue_id' to be a int")
        pulumi.set(__self__, "epic_issue_id", epic_issue_id)
        if external_id and not isinstance(external_id, str):
            raise TypeError("Expected argument 'external_id' to be a str")
        pulumi.set(__self__, "external_id", external_id)
        if human_time_estimate and not isinstance(human_time_estimate, str):
            raise TypeError("Expected argument 'human_time_estimate' to be a str")
        pulumi.set(__self__, "human_time_estimate", human_time_estimate)
        if human_total_time_spent and not isinstance(human_total_time_spent, str):
            raise TypeError("Expected argument 'human_total_time_spent' to be a str")
        pulumi.set(__self__, "human_total_time_spent", human_total_time_spent)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iid and not isinstance(iid, int):
            raise TypeError("Expected argument 'iid' to be a int")
        pulumi.set(__self__, "iid", iid)
        if issue_id and not isinstance(issue_id, int):
            raise TypeError("Expected argument 'issue_id' to be a int")
        pulumi.set(__self__, "issue_id", issue_id)
        if issue_link_id and not isinstance(issue_link_id, int):
            raise TypeError("Expected argument 'issue_link_id' to be a int")
        pulumi.set(__self__, "issue_link_id", issue_link_id)
        if issue_type and not isinstance(issue_type, str):
            raise TypeError("Expected argument 'issue_type' to be a str")
        pulumi.set(__self__, "issue_type", issue_type)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if merge_request_to_resolve_discussions_of and not isinstance(merge_request_to_resolve_discussions_of, int):
            raise TypeError("Expected argument 'merge_request_to_resolve_discussions_of' to be a int")
        pulumi.set(__self__, "merge_request_to_resolve_discussions_of", merge_request_to_resolve_discussions_of)
        if merge_requests_count and not isinstance(merge_requests_count, int):
            raise TypeError("Expected argument 'merge_requests_count' to be a int")
        pulumi.set(__self__, "merge_requests_count", merge_requests_count)
        if milestone_id and not isinstance(milestone_id, int):
            raise TypeError("Expected argument 'milestone_id' to be a int")
        pulumi.set(__self__, "milestone_id", milestone_id)
        if moved_to_id and not isinstance(moved_to_id, int):
            raise TypeError("Expected argument 'moved_to_id' to be a int")
        pulumi.set(__self__, "moved_to_id", moved_to_id)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if references and not isinstance(references, dict):
            raise TypeError("Expected argument 'references' to be a dict")
        pulumi.set(__self__, "references", references)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if subscribed and not isinstance(subscribed, bool):
            raise TypeError("Expected argument 'subscribed' to be a bool")
        pulumi.set(__self__, "subscribed", subscribed)
        if task_completion_statuses and not isinstance(task_completion_statuses, list):
            raise TypeError("Expected argument 'task_completion_statuses' to be a list")
        pulumi.set(__self__, "task_completion_statuses", task_completion_statuses)
        if time_estimate and not isinstance(time_estimate, int):
            raise TypeError("Expected argument 'time_estimate' to be a int")
        pulumi.set(__self__, "time_estimate", time_estimate)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if total_time_spent and not isinstance(total_time_spent, int):
            raise TypeError("Expected argument 'total_time_spent' to be a int")
        pulumi.set(__self__, "total_time_spent", total_time_spent)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if upvotes and not isinstance(upvotes, int):
            raise TypeError("Expected argument 'upvotes' to be a int")
        pulumi.set(__self__, "upvotes", upvotes)
        if user_notes_count and not isinstance(user_notes_count, int):
            raise TypeError("Expected argument 'user_notes_count' to be a int")
        pulumi.set(__self__, "user_notes_count", user_notes_count)
        if web_url and not isinstance(web_url, str):
            raise TypeError("Expected argument 'web_url' to be a str")
        pulumi.set(__self__, "web_url", web_url)
        if weight and not isinstance(weight, int):
            raise TypeError("Expected argument 'weight' to be a int")
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="assigneeIds")
    def assignee_ids(self) -> Sequence[int]:
        """
        The IDs of the users to assign the issue to.
        """
        return pulumi.get(self, "assignee_ids")

    @property
    @pulumi.getter(name="authorId")
    def author_id(self) -> int:
        """
        The ID of the author of the issue. Use `User` data source to get more information about the user.
        """
        return pulumi.get(self, "author_id")

    @property
    @pulumi.getter(name="closedAt")
    def closed_at(self) -> str:
        """
        When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        """
        return pulumi.get(self, "closed_at")

    @property
    @pulumi.getter(name="closedByUserId")
    def closed_by_user_id(self) -> int:
        """
        The ID of the user that closed the issue. Use `User` data source to get more information about the user.
        """
        return pulumi.get(self, "closed_by_user_id")

    @property
    @pulumi.getter
    def confidential(self) -> bool:
        """
        Set an issue to be confidential.
        """
        return pulumi.get(self, "confidential")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of an issue. Limited to 1,048,576 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="discussionLocked")
    def discussion_locked(self) -> bool:
        """
        Whether the issue is locked for discussions or not.
        """
        return pulumi.get(self, "discussion_locked")

    @property
    @pulumi.getter(name="discussionToResolve")
    def discussion_to_resolve(self) -> str:
        """
        The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
        """
        return pulumi.get(self, "discussion_to_resolve")

    @property
    @pulumi.getter
    def downvotes(self) -> int:
        """
        The number of downvotes the issue has received.
        """
        return pulumi.get(self, "downvotes")

    @property
    @pulumi.getter(name="dueDate")
    def due_date(self) -> str:
        """
        The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        """
        return pulumi.get(self, "due_date")

    @property
    @pulumi.getter(name="epicId")
    def epic_id(self) -> int:
        """
        ID of the epic to add the issue to. Valid values are greater than or equal to 0.
        """
        return pulumi.get(self, "epic_id")

    @property
    @pulumi.getter(name="epicIssueId")
    def epic_issue_id(self) -> int:
        """
        The ID of the epic issue.
        """
        return pulumi.get(self, "epic_issue_id")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> str:
        """
        The external ID of the issue.
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter(name="humanTimeEstimate")
    def human_time_estimate(self) -> str:
        """
        The human-readable time estimate of the issue.
        """
        return pulumi.get(self, "human_time_estimate")

    @property
    @pulumi.getter(name="humanTotalTimeSpent")
    def human_total_time_spent(self) -> str:
        """
        The human-readable total time spent of the issue.
        """
        return pulumi.get(self, "human_total_time_spent")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def iid(self) -> int:
        """
        The internal ID of the project's issue.
        """
        return pulumi.get(self, "iid")

    @property
    @pulumi.getter(name="issueId")
    def issue_id(self) -> int:
        """
        The instance-wide ID of the issue.
        """
        return pulumi.get(self, "issue_id")

    @property
    @pulumi.getter(name="issueLinkId")
    def issue_link_id(self) -> int:
        """
        The ID of the issue link.
        """
        return pulumi.get(self, "issue_link_id")

    @property
    @pulumi.getter(name="issueType")
    def issue_type(self) -> str:
        """
        The type of issue. Valid values are: `issue`, `incident`, `test_case`.
        """
        return pulumi.get(self, "issue_type")

    @property
    @pulumi.getter
    def labels(self) -> Sequence[str]:
        """
        The labels of an issue.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def links(self) -> Mapping[str, str]:
        """
        The links of the issue.
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter(name="mergeRequestToResolveDiscussionsOf")
    def merge_request_to_resolve_discussions_of(self) -> int:
        """
        The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
        """
        return pulumi.get(self, "merge_request_to_resolve_discussions_of")

    @property
    @pulumi.getter(name="mergeRequestsCount")
    def merge_requests_count(self) -> int:
        """
        The number of merge requests associated with the issue.
        """
        return pulumi.get(self, "merge_requests_count")

    @property
    @pulumi.getter(name="milestoneId")
    def milestone_id(self) -> int:
        """
        The global ID of a milestone to assign issue. To find the milestone_id associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
        """
        return pulumi.get(self, "milestone_id")

    @property
    @pulumi.getter(name="movedToId")
    def moved_to_id(self) -> int:
        """
        The ID of the issue that was moved to.
        """
        return pulumi.get(self, "moved_to_id")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The name or ID of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def references(self) -> Mapping[str, str]:
        """
        The references of the issue.
        """
        return pulumi.get(self, "references")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the issue. Valid values are: `opened`, `closed`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def subscribed(self) -> bool:
        """
        Whether the authenticated user is subscribed to the issue or not.
        """
        return pulumi.get(self, "subscribed")

    @property
    @pulumi.getter(name="taskCompletionStatuses")
    def task_completion_statuses(self) -> Sequence['outputs.GetProjectIssueTaskCompletionStatusResult']:
        """
        The task completion status. It's always a one element list.
        """
        return pulumi.get(self, "task_completion_statuses")

    @property
    @pulumi.getter(name="timeEstimate")
    def time_estimate(self) -> int:
        """
        The time estimate of the issue.
        """
        return pulumi.get(self, "time_estimate")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        The title of the issue.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter(name="totalTimeSpent")
    def total_time_spent(self) -> int:
        """
        The total time spent of the issue.
        """
        return pulumi.get(self, "total_time_spent")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def upvotes(self) -> int:
        """
        The number of upvotes the issue has received.
        """
        return pulumi.get(self, "upvotes")

    @property
    @pulumi.getter(name="userNotesCount")
    def user_notes_count(self) -> int:
        """
        The number of user notes on the issue.
        """
        return pulumi.get(self, "user_notes_count")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> str:
        """
        The web URL of the issue.
        """
        return pulumi.get(self, "web_url")

    @property
    @pulumi.getter
    def weight(self) -> int:
        """
        The weight of the issue. Valid values are greater than or equal to 0.
        """
        return pulumi.get(self, "weight")


class AwaitableGetProjectIssueResult(GetProjectIssueResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectIssueResult(
            assignee_ids=self.assignee_ids,
            author_id=self.author_id,
            closed_at=self.closed_at,
            closed_by_user_id=self.closed_by_user_id,
            confidential=self.confidential,
            created_at=self.created_at,
            description=self.description,
            discussion_locked=self.discussion_locked,
            discussion_to_resolve=self.discussion_to_resolve,
            downvotes=self.downvotes,
            due_date=self.due_date,
            epic_id=self.epic_id,
            epic_issue_id=self.epic_issue_id,
            external_id=self.external_id,
            human_time_estimate=self.human_time_estimate,
            human_total_time_spent=self.human_total_time_spent,
            id=self.id,
            iid=self.iid,
            issue_id=self.issue_id,
            issue_link_id=self.issue_link_id,
            issue_type=self.issue_type,
            labels=self.labels,
            links=self.links,
            merge_request_to_resolve_discussions_of=self.merge_request_to_resolve_discussions_of,
            merge_requests_count=self.merge_requests_count,
            milestone_id=self.milestone_id,
            moved_to_id=self.moved_to_id,
            project=self.project,
            references=self.references,
            state=self.state,
            subscribed=self.subscribed,
            task_completion_statuses=self.task_completion_statuses,
            time_estimate=self.time_estimate,
            title=self.title,
            total_time_spent=self.total_time_spent,
            updated_at=self.updated_at,
            upvotes=self.upvotes,
            user_notes_count=self.user_notes_count,
            web_url=self.web_url,
            weight=self.weight)


def get_project_issue(iid: Optional[int] = None,
                      project: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectIssueResult:
    """
    The `ProjectIssue` data source allows to retrieve details about an issue in a project.

    **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/issues.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    foo = gitlab.get_project(id="foo/bar/baz")
    welcome_issue = gitlab.get_project_issue(project=foo.id,
        iid=1)
    pulumi.export("welcomeIssueWebUrl", data["gitlab_project_issue"]["web_url"])
    ```


    :param int iid: The internal ID of the project's issue.
    :param str project: The name or ID of the project.
    """
    __args__ = dict()
    __args__['iid'] = iid
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjectIssue:getProjectIssue', __args__, opts=opts, typ=GetProjectIssueResult).value

    return AwaitableGetProjectIssueResult(
        assignee_ids=__ret__.assignee_ids,
        author_id=__ret__.author_id,
        closed_at=__ret__.closed_at,
        closed_by_user_id=__ret__.closed_by_user_id,
        confidential=__ret__.confidential,
        created_at=__ret__.created_at,
        description=__ret__.description,
        discussion_locked=__ret__.discussion_locked,
        discussion_to_resolve=__ret__.discussion_to_resolve,
        downvotes=__ret__.downvotes,
        due_date=__ret__.due_date,
        epic_id=__ret__.epic_id,
        epic_issue_id=__ret__.epic_issue_id,
        external_id=__ret__.external_id,
        human_time_estimate=__ret__.human_time_estimate,
        human_total_time_spent=__ret__.human_total_time_spent,
        id=__ret__.id,
        iid=__ret__.iid,
        issue_id=__ret__.issue_id,
        issue_link_id=__ret__.issue_link_id,
        issue_type=__ret__.issue_type,
        labels=__ret__.labels,
        links=__ret__.links,
        merge_request_to_resolve_discussions_of=__ret__.merge_request_to_resolve_discussions_of,
        merge_requests_count=__ret__.merge_requests_count,
        milestone_id=__ret__.milestone_id,
        moved_to_id=__ret__.moved_to_id,
        project=__ret__.project,
        references=__ret__.references,
        state=__ret__.state,
        subscribed=__ret__.subscribed,
        task_completion_statuses=__ret__.task_completion_statuses,
        time_estimate=__ret__.time_estimate,
        title=__ret__.title,
        total_time_spent=__ret__.total_time_spent,
        updated_at=__ret__.updated_at,
        upvotes=__ret__.upvotes,
        user_notes_count=__ret__.user_notes_count,
        web_url=__ret__.web_url,
        weight=__ret__.weight)


@_utilities.lift_output_func(get_project_issue)
def get_project_issue_output(iid: Optional[pulumi.Input[int]] = None,
                             project: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectIssueResult]:
    """
    The `ProjectIssue` data source allows to retrieve details about an issue in a project.

    **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/issues.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    foo = gitlab.get_project(id="foo/bar/baz")
    welcome_issue = gitlab.get_project_issue(project=foo.id,
        iid=1)
    pulumi.export("welcomeIssueWebUrl", data["gitlab_project_issue"]["web_url"])
    ```


    :param int iid: The internal ID of the project's issue.
    :param str project: The name or ID of the project.
    """
    ...
