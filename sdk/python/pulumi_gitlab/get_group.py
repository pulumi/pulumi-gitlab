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
    'GetGroupResult',
    'AwaitableGetGroupResult',
    'get_group',
    'get_group_output',
]

@pulumi.output_type
class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, default_branch_protection=None, description=None, extra_shared_runners_minutes_limit=None, full_name=None, full_path=None, group_id=None, id=None, lfs_enabled=None, membership_lock=None, name=None, parent_id=None, path=None, prevent_forking_outside_group=None, request_access_enabled=None, runners_token=None, shared_runners_minutes_limit=None, shared_runners_setting=None, shared_with_groups=None, visibility_level=None, web_url=None, wiki_access_level=None):
        if default_branch_protection and not isinstance(default_branch_protection, int):
            raise TypeError("Expected argument 'default_branch_protection' to be a int")
        pulumi.set(__self__, "default_branch_protection", default_branch_protection)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if extra_shared_runners_minutes_limit and not isinstance(extra_shared_runners_minutes_limit, int):
            raise TypeError("Expected argument 'extra_shared_runners_minutes_limit' to be a int")
        pulumi.set(__self__, "extra_shared_runners_minutes_limit", extra_shared_runners_minutes_limit)
        if full_name and not isinstance(full_name, str):
            raise TypeError("Expected argument 'full_name' to be a str")
        pulumi.set(__self__, "full_name", full_name)
        if full_path and not isinstance(full_path, str):
            raise TypeError("Expected argument 'full_path' to be a str")
        pulumi.set(__self__, "full_path", full_path)
        if group_id and not isinstance(group_id, int):
            raise TypeError("Expected argument 'group_id' to be a int")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lfs_enabled and not isinstance(lfs_enabled, bool):
            raise TypeError("Expected argument 'lfs_enabled' to be a bool")
        pulumi.set(__self__, "lfs_enabled", lfs_enabled)
        if membership_lock and not isinstance(membership_lock, bool):
            raise TypeError("Expected argument 'membership_lock' to be a bool")
        pulumi.set(__self__, "membership_lock", membership_lock)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_id and not isinstance(parent_id, int):
            raise TypeError("Expected argument 'parent_id' to be a int")
        pulumi.set(__self__, "parent_id", parent_id)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if prevent_forking_outside_group and not isinstance(prevent_forking_outside_group, bool):
            raise TypeError("Expected argument 'prevent_forking_outside_group' to be a bool")
        pulumi.set(__self__, "prevent_forking_outside_group", prevent_forking_outside_group)
        if request_access_enabled and not isinstance(request_access_enabled, bool):
            raise TypeError("Expected argument 'request_access_enabled' to be a bool")
        pulumi.set(__self__, "request_access_enabled", request_access_enabled)
        if runners_token and not isinstance(runners_token, str):
            raise TypeError("Expected argument 'runners_token' to be a str")
        pulumi.set(__self__, "runners_token", runners_token)
        if shared_runners_minutes_limit and not isinstance(shared_runners_minutes_limit, int):
            raise TypeError("Expected argument 'shared_runners_minutes_limit' to be a int")
        pulumi.set(__self__, "shared_runners_minutes_limit", shared_runners_minutes_limit)
        if shared_runners_setting and not isinstance(shared_runners_setting, str):
            raise TypeError("Expected argument 'shared_runners_setting' to be a str")
        pulumi.set(__self__, "shared_runners_setting", shared_runners_setting)
        if shared_with_groups and not isinstance(shared_with_groups, list):
            raise TypeError("Expected argument 'shared_with_groups' to be a list")
        pulumi.set(__self__, "shared_with_groups", shared_with_groups)
        if visibility_level and not isinstance(visibility_level, str):
            raise TypeError("Expected argument 'visibility_level' to be a str")
        pulumi.set(__self__, "visibility_level", visibility_level)
        if web_url and not isinstance(web_url, str):
            raise TypeError("Expected argument 'web_url' to be a str")
        pulumi.set(__self__, "web_url", web_url)
        if wiki_access_level and not isinstance(wiki_access_level, str):
            raise TypeError("Expected argument 'wiki_access_level' to be a str")
        pulumi.set(__self__, "wiki_access_level", wiki_access_level)

    @property
    @pulumi.getter(name="defaultBranchProtection")
    def default_branch_protection(self) -> int:
        """
        Whether developers and maintainers can push to the applicable default branch.
        """
        return pulumi.get(self, "default_branch_protection")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="extraSharedRunnersMinutesLimit")
    def extra_shared_runners_minutes_limit(self) -> int:
        """
        Can be set by administrators only. Additional CI/CD minutes for this group.
        """
        return pulumi.get(self, "extra_shared_runners_minutes_limit")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> str:
        """
        The full name of the group.
        """
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> str:
        """
        The full path of the group.
        """
        return pulumi.get(self, "full_path")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> int:
        """
        The ID of the group.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lfsEnabled")
    def lfs_enabled(self) -> bool:
        """
        Boolean, is LFS enabled for projects in this group.
        """
        return pulumi.get(self, "lfs_enabled")

    @property
    @pulumi.getter(name="membershipLock")
    def membership_lock(self) -> bool:
        """
        Users cannot be added to projects in this group.
        """
        return pulumi.get(self, "membership_lock")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of this group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> int:
        """
        Integer, ID of the parent group.
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        The path of the group.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="preventForkingOutsideGroup")
    def prevent_forking_outside_group(self) -> bool:
        """
        When enabled, users can not fork projects from this group to external namespaces.
        """
        return pulumi.get(self, "prevent_forking_outside_group")

    @property
    @pulumi.getter(name="requestAccessEnabled")
    def request_access_enabled(self) -> bool:
        """
        Boolean, is request for access enabled to the group.
        """
        return pulumi.get(self, "request_access_enabled")

    @property
    @pulumi.getter(name="runnersToken")
    def runners_token(self) -> str:
        """
        The group level registration token to use during runner setup.
        """
        return pulumi.get(self, "runners_token")

    @property
    @pulumi.getter(name="sharedRunnersMinutesLimit")
    def shared_runners_minutes_limit(self) -> int:
        """
        Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
        """
        return pulumi.get(self, "shared_runners_minutes_limit")

    @property
    @pulumi.getter(name="sharedRunnersSetting")
    def shared_runners_setting(self) -> str:
        """
        Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
        """
        return pulumi.get(self, "shared_runners_setting")

    @property
    @pulumi.getter(name="sharedWithGroups")
    def shared_with_groups(self) -> Sequence['outputs.GetGroupSharedWithGroupResult']:
        """
        Describes groups which have access shared to this group.
        """
        return pulumi.get(self, "shared_with_groups")

    @property
    @pulumi.getter(name="visibilityLevel")
    def visibility_level(self) -> str:
        """
        Visibility level of the group. Possible values are `private`, `internal`, `public`.
        """
        return pulumi.get(self, "visibility_level")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> str:
        """
        Web URL of the group.
        """
        return pulumi.get(self, "web_url")

    @property
    @pulumi.getter(name="wikiAccessLevel")
    def wiki_access_level(self) -> str:
        """
        The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
        """
        return pulumi.get(self, "wiki_access_level")


class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            default_branch_protection=self.default_branch_protection,
            description=self.description,
            extra_shared_runners_minutes_limit=self.extra_shared_runners_minutes_limit,
            full_name=self.full_name,
            full_path=self.full_path,
            group_id=self.group_id,
            id=self.id,
            lfs_enabled=self.lfs_enabled,
            membership_lock=self.membership_lock,
            name=self.name,
            parent_id=self.parent_id,
            path=self.path,
            prevent_forking_outside_group=self.prevent_forking_outside_group,
            request_access_enabled=self.request_access_enabled,
            runners_token=self.runners_token,
            shared_runners_minutes_limit=self.shared_runners_minutes_limit,
            shared_runners_setting=self.shared_runners_setting,
            shared_with_groups=self.shared_with_groups,
            visibility_level=self.visibility_level,
            web_url=self.web_url,
            wiki_access_level=self.wiki_access_level)


def get_group(full_path: Optional[str] = None,
              group_id: Optional[int] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupResult:
    """
    The `Group` data source allows details of a group to be retrieved by its id or full path.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#details-of-a-group)


    :param str full_path: The full path of the group.
    :param int group_id: The ID of the group.
    """
    __args__ = dict()
    __args__['fullPath'] = full_path
    __args__['groupId'] = group_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult).value

    return AwaitableGetGroupResult(
        default_branch_protection=pulumi.get(__ret__, 'default_branch_protection'),
        description=pulumi.get(__ret__, 'description'),
        extra_shared_runners_minutes_limit=pulumi.get(__ret__, 'extra_shared_runners_minutes_limit'),
        full_name=pulumi.get(__ret__, 'full_name'),
        full_path=pulumi.get(__ret__, 'full_path'),
        group_id=pulumi.get(__ret__, 'group_id'),
        id=pulumi.get(__ret__, 'id'),
        lfs_enabled=pulumi.get(__ret__, 'lfs_enabled'),
        membership_lock=pulumi.get(__ret__, 'membership_lock'),
        name=pulumi.get(__ret__, 'name'),
        parent_id=pulumi.get(__ret__, 'parent_id'),
        path=pulumi.get(__ret__, 'path'),
        prevent_forking_outside_group=pulumi.get(__ret__, 'prevent_forking_outside_group'),
        request_access_enabled=pulumi.get(__ret__, 'request_access_enabled'),
        runners_token=pulumi.get(__ret__, 'runners_token'),
        shared_runners_minutes_limit=pulumi.get(__ret__, 'shared_runners_minutes_limit'),
        shared_runners_setting=pulumi.get(__ret__, 'shared_runners_setting'),
        shared_with_groups=pulumi.get(__ret__, 'shared_with_groups'),
        visibility_level=pulumi.get(__ret__, 'visibility_level'),
        web_url=pulumi.get(__ret__, 'web_url'),
        wiki_access_level=pulumi.get(__ret__, 'wiki_access_level'))
def get_group_output(full_path: Optional[pulumi.Input[Optional[str]]] = None,
                     group_id: Optional[pulumi.Input[Optional[int]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupResult]:
    """
    The `Group` data source allows details of a group to be retrieved by its id or full path.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#details-of-a-group)


    :param str full_path: The full path of the group.
    :param int group_id: The ID of the group.
    """
    __args__ = dict()
    __args__['fullPath'] = full_path
    __args__['groupId'] = group_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult)
    return __ret__.apply(lambda __response__: GetGroupResult(
        default_branch_protection=pulumi.get(__response__, 'default_branch_protection'),
        description=pulumi.get(__response__, 'description'),
        extra_shared_runners_minutes_limit=pulumi.get(__response__, 'extra_shared_runners_minutes_limit'),
        full_name=pulumi.get(__response__, 'full_name'),
        full_path=pulumi.get(__response__, 'full_path'),
        group_id=pulumi.get(__response__, 'group_id'),
        id=pulumi.get(__response__, 'id'),
        lfs_enabled=pulumi.get(__response__, 'lfs_enabled'),
        membership_lock=pulumi.get(__response__, 'membership_lock'),
        name=pulumi.get(__response__, 'name'),
        parent_id=pulumi.get(__response__, 'parent_id'),
        path=pulumi.get(__response__, 'path'),
        prevent_forking_outside_group=pulumi.get(__response__, 'prevent_forking_outside_group'),
        request_access_enabled=pulumi.get(__response__, 'request_access_enabled'),
        runners_token=pulumi.get(__response__, 'runners_token'),
        shared_runners_minutes_limit=pulumi.get(__response__, 'shared_runners_minutes_limit'),
        shared_runners_setting=pulumi.get(__response__, 'shared_runners_setting'),
        shared_with_groups=pulumi.get(__response__, 'shared_with_groups'),
        visibility_level=pulumi.get(__response__, 'visibility_level'),
        web_url=pulumi.get(__response__, 'web_url'),
        wiki_access_level=pulumi.get(__response__, 'wiki_access_level')))
