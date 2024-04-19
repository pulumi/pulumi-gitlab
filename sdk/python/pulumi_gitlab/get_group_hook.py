# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetGroupHookResult',
    'AwaitableGetGroupHookResult',
    'get_group_hook',
    'get_group_hook_output',
]

@pulumi.output_type
class GetGroupHookResult:
    """
    A collection of values returned by getGroupHook.
    """
    def __init__(__self__, confidential_issues_events=None, confidential_note_events=None, custom_webhook_template=None, deployment_events=None, enable_ssl_verification=None, group=None, group_id=None, hook_id=None, id=None, issues_events=None, job_events=None, merge_requests_events=None, note_events=None, pipeline_events=None, push_events=None, push_events_branch_filter=None, releases_events=None, subgroup_events=None, tag_push_events=None, token=None, url=None, wiki_page_events=None):
        if confidential_issues_events and not isinstance(confidential_issues_events, bool):
            raise TypeError("Expected argument 'confidential_issues_events' to be a bool")
        pulumi.set(__self__, "confidential_issues_events", confidential_issues_events)
        if confidential_note_events and not isinstance(confidential_note_events, bool):
            raise TypeError("Expected argument 'confidential_note_events' to be a bool")
        pulumi.set(__self__, "confidential_note_events", confidential_note_events)
        if custom_webhook_template and not isinstance(custom_webhook_template, str):
            raise TypeError("Expected argument 'custom_webhook_template' to be a str")
        pulumi.set(__self__, "custom_webhook_template", custom_webhook_template)
        if deployment_events and not isinstance(deployment_events, bool):
            raise TypeError("Expected argument 'deployment_events' to be a bool")
        pulumi.set(__self__, "deployment_events", deployment_events)
        if enable_ssl_verification and not isinstance(enable_ssl_verification, bool):
            raise TypeError("Expected argument 'enable_ssl_verification' to be a bool")
        pulumi.set(__self__, "enable_ssl_verification", enable_ssl_verification)
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        pulumi.set(__self__, "group", group)
        if group_id and not isinstance(group_id, int):
            raise TypeError("Expected argument 'group_id' to be a int")
        pulumi.set(__self__, "group_id", group_id)
        if hook_id and not isinstance(hook_id, int):
            raise TypeError("Expected argument 'hook_id' to be a int")
        pulumi.set(__self__, "hook_id", hook_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if issues_events and not isinstance(issues_events, bool):
            raise TypeError("Expected argument 'issues_events' to be a bool")
        pulumi.set(__self__, "issues_events", issues_events)
        if job_events and not isinstance(job_events, bool):
            raise TypeError("Expected argument 'job_events' to be a bool")
        pulumi.set(__self__, "job_events", job_events)
        if merge_requests_events and not isinstance(merge_requests_events, bool):
            raise TypeError("Expected argument 'merge_requests_events' to be a bool")
        pulumi.set(__self__, "merge_requests_events", merge_requests_events)
        if note_events and not isinstance(note_events, bool):
            raise TypeError("Expected argument 'note_events' to be a bool")
        pulumi.set(__self__, "note_events", note_events)
        if pipeline_events and not isinstance(pipeline_events, bool):
            raise TypeError("Expected argument 'pipeline_events' to be a bool")
        pulumi.set(__self__, "pipeline_events", pipeline_events)
        if push_events and not isinstance(push_events, bool):
            raise TypeError("Expected argument 'push_events' to be a bool")
        pulumi.set(__self__, "push_events", push_events)
        if push_events_branch_filter and not isinstance(push_events_branch_filter, str):
            raise TypeError("Expected argument 'push_events_branch_filter' to be a str")
        pulumi.set(__self__, "push_events_branch_filter", push_events_branch_filter)
        if releases_events and not isinstance(releases_events, bool):
            raise TypeError("Expected argument 'releases_events' to be a bool")
        pulumi.set(__self__, "releases_events", releases_events)
        if subgroup_events and not isinstance(subgroup_events, bool):
            raise TypeError("Expected argument 'subgroup_events' to be a bool")
        pulumi.set(__self__, "subgroup_events", subgroup_events)
        if tag_push_events and not isinstance(tag_push_events, bool):
            raise TypeError("Expected argument 'tag_push_events' to be a bool")
        pulumi.set(__self__, "tag_push_events", tag_push_events)
        if token and not isinstance(token, str):
            raise TypeError("Expected argument 'token' to be a str")
        pulumi.set(__self__, "token", token)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if wiki_page_events and not isinstance(wiki_page_events, bool):
            raise TypeError("Expected argument 'wiki_page_events' to be a bool")
        pulumi.set(__self__, "wiki_page_events", wiki_page_events)

    @property
    @pulumi.getter(name="confidentialIssuesEvents")
    def confidential_issues_events(self) -> bool:
        """
        Invoke the hook for confidential issues events.
        """
        return pulumi.get(self, "confidential_issues_events")

    @property
    @pulumi.getter(name="confidentialNoteEvents")
    def confidential_note_events(self) -> bool:
        """
        Invoke the hook for confidential notes events.
        """
        return pulumi.get(self, "confidential_note_events")

    @property
    @pulumi.getter(name="customWebhookTemplate")
    def custom_webhook_template(self) -> str:
        """
        Set a custom webhook template.
        """
        return pulumi.get(self, "custom_webhook_template")

    @property
    @pulumi.getter(name="deploymentEvents")
    def deployment_events(self) -> bool:
        """
        Invoke the hook for deployment events.
        """
        return pulumi.get(self, "deployment_events")

    @property
    @pulumi.getter(name="enableSslVerification")
    def enable_ssl_verification(self) -> bool:
        """
        Enable ssl verification when invoking the hook.
        """
        return pulumi.get(self, "enable_ssl_verification")

    @property
    @pulumi.getter
    def group(self) -> str:
        """
        The ID or full path of the group.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> int:
        """
        The id of the group for the hook.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="hookId")
    def hook_id(self) -> int:
        """
        The id of the group hook.
        """
        return pulumi.get(self, "hook_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="issuesEvents")
    def issues_events(self) -> bool:
        """
        Invoke the hook for issues events.
        """
        return pulumi.get(self, "issues_events")

    @property
    @pulumi.getter(name="jobEvents")
    def job_events(self) -> bool:
        """
        Invoke the hook for job events.
        """
        return pulumi.get(self, "job_events")

    @property
    @pulumi.getter(name="mergeRequestsEvents")
    def merge_requests_events(self) -> bool:
        """
        Invoke the hook for merge requests.
        """
        return pulumi.get(self, "merge_requests_events")

    @property
    @pulumi.getter(name="noteEvents")
    def note_events(self) -> bool:
        """
        Invoke the hook for notes events.
        """
        return pulumi.get(self, "note_events")

    @property
    @pulumi.getter(name="pipelineEvents")
    def pipeline_events(self) -> bool:
        """
        Invoke the hook for pipeline events.
        """
        return pulumi.get(self, "pipeline_events")

    @property
    @pulumi.getter(name="pushEvents")
    def push_events(self) -> bool:
        """
        Invoke the hook for push events.
        """
        return pulumi.get(self, "push_events")

    @property
    @pulumi.getter(name="pushEventsBranchFilter")
    def push_events_branch_filter(self) -> str:
        """
        Invoke the hook for push events on matching branches only.
        """
        return pulumi.get(self, "push_events_branch_filter")

    @property
    @pulumi.getter(name="releasesEvents")
    def releases_events(self) -> bool:
        """
        Invoke the hook for releases events.
        """
        return pulumi.get(self, "releases_events")

    @property
    @pulumi.getter(name="subgroupEvents")
    def subgroup_events(self) -> bool:
        """
        Invoke the hook for subgroup events.
        """
        return pulumi.get(self, "subgroup_events")

    @property
    @pulumi.getter(name="tagPushEvents")
    def tag_push_events(self) -> bool:
        """
        Invoke the hook for tag push events.
        """
        return pulumi.get(self, "tag_push_events")

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        A token to present when invoking the hook. The token is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The url of the hook to invoke.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="wikiPageEvents")
    def wiki_page_events(self) -> bool:
        """
        Invoke the hook for wiki page events.
        """
        return pulumi.get(self, "wiki_page_events")


class AwaitableGetGroupHookResult(GetGroupHookResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupHookResult(
            confidential_issues_events=self.confidential_issues_events,
            confidential_note_events=self.confidential_note_events,
            custom_webhook_template=self.custom_webhook_template,
            deployment_events=self.deployment_events,
            enable_ssl_verification=self.enable_ssl_verification,
            group=self.group,
            group_id=self.group_id,
            hook_id=self.hook_id,
            id=self.id,
            issues_events=self.issues_events,
            job_events=self.job_events,
            merge_requests_events=self.merge_requests_events,
            note_events=self.note_events,
            pipeline_events=self.pipeline_events,
            push_events=self.push_events,
            push_events_branch_filter=self.push_events_branch_filter,
            releases_events=self.releases_events,
            subgroup_events=self.subgroup_events,
            tag_push_events=self.tag_push_events,
            token=self.token,
            url=self.url,
            wiki_page_events=self.wiki_page_events)


def get_group_hook(group: Optional[str] = None,
                   hook_id: Optional[int] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupHookResult:
    """
    The `GroupHook` data source allows to retrieve details about a hook in a group.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#get-group-hook)


    :param str group: The ID or full path of the group.
    :param int hook_id: The id of the group hook.
    """
    __args__ = dict()
    __args__['group'] = group
    __args__['hookId'] = hook_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getGroupHook:getGroupHook', __args__, opts=opts, typ=GetGroupHookResult).value

    return AwaitableGetGroupHookResult(
        confidential_issues_events=pulumi.get(__ret__, 'confidential_issues_events'),
        confidential_note_events=pulumi.get(__ret__, 'confidential_note_events'),
        custom_webhook_template=pulumi.get(__ret__, 'custom_webhook_template'),
        deployment_events=pulumi.get(__ret__, 'deployment_events'),
        enable_ssl_verification=pulumi.get(__ret__, 'enable_ssl_verification'),
        group=pulumi.get(__ret__, 'group'),
        group_id=pulumi.get(__ret__, 'group_id'),
        hook_id=pulumi.get(__ret__, 'hook_id'),
        id=pulumi.get(__ret__, 'id'),
        issues_events=pulumi.get(__ret__, 'issues_events'),
        job_events=pulumi.get(__ret__, 'job_events'),
        merge_requests_events=pulumi.get(__ret__, 'merge_requests_events'),
        note_events=pulumi.get(__ret__, 'note_events'),
        pipeline_events=pulumi.get(__ret__, 'pipeline_events'),
        push_events=pulumi.get(__ret__, 'push_events'),
        push_events_branch_filter=pulumi.get(__ret__, 'push_events_branch_filter'),
        releases_events=pulumi.get(__ret__, 'releases_events'),
        subgroup_events=pulumi.get(__ret__, 'subgroup_events'),
        tag_push_events=pulumi.get(__ret__, 'tag_push_events'),
        token=pulumi.get(__ret__, 'token'),
        url=pulumi.get(__ret__, 'url'),
        wiki_page_events=pulumi.get(__ret__, 'wiki_page_events'))


@_utilities.lift_output_func(get_group_hook)
def get_group_hook_output(group: Optional[pulumi.Input[str]] = None,
                          hook_id: Optional[pulumi.Input[int]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupHookResult]:
    """
    The `GroupHook` data source allows to retrieve details about a hook in a group.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#get-group-hook)


    :param str group: The ID or full path of the group.
    :param int hook_id: The id of the group hook.
    """
    ...
