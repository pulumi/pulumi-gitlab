# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'GetProjectsResult',
    'AwaitableGetProjectsResult',
    'get_projects',
    'get_projects_output',
]

@pulumi.output_type
class GetProjectsResult:
    """
    A collection of values returned by getProjects.
    """
    def __init__(__self__, archived=None, group_id=None, id=None, include_subgroups=None, max_queryable_pages=None, membership=None, min_access_level=None, order_by=None, owned=None, page=None, per_page=None, projects=None, search=None, simple=None, sort=None, starred=None, statistics=None, topics=None, visibility=None, with_custom_attributes=None, with_issues_enabled=None, with_merge_requests_enabled=None, with_programming_language=None, with_shared=None):
        if archived and not isinstance(archived, bool):
            raise TypeError("Expected argument 'archived' to be a bool")
        pulumi.set(__self__, "archived", archived)
        if group_id and not isinstance(group_id, int):
            raise TypeError("Expected argument 'group_id' to be a int")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_subgroups and not isinstance(include_subgroups, bool):
            raise TypeError("Expected argument 'include_subgroups' to be a bool")
        pulumi.set(__self__, "include_subgroups", include_subgroups)
        if max_queryable_pages and not isinstance(max_queryable_pages, int):
            raise TypeError("Expected argument 'max_queryable_pages' to be a int")
        pulumi.set(__self__, "max_queryable_pages", max_queryable_pages)
        if membership and not isinstance(membership, bool):
            raise TypeError("Expected argument 'membership' to be a bool")
        pulumi.set(__self__, "membership", membership)
        if min_access_level and not isinstance(min_access_level, int):
            raise TypeError("Expected argument 'min_access_level' to be a int")
        pulumi.set(__self__, "min_access_level", min_access_level)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if owned and not isinstance(owned, bool):
            raise TypeError("Expected argument 'owned' to be a bool")
        pulumi.set(__self__, "owned", owned)
        if page and not isinstance(page, int):
            raise TypeError("Expected argument 'page' to be a int")
        pulumi.set(__self__, "page", page)
        if per_page and not isinstance(per_page, int):
            raise TypeError("Expected argument 'per_page' to be a int")
        pulumi.set(__self__, "per_page", per_page)
        if projects and not isinstance(projects, list):
            raise TypeError("Expected argument 'projects' to be a list")
        pulumi.set(__self__, "projects", projects)
        if search and not isinstance(search, str):
            raise TypeError("Expected argument 'search' to be a str")
        pulumi.set(__self__, "search", search)
        if simple and not isinstance(simple, bool):
            raise TypeError("Expected argument 'simple' to be a bool")
        pulumi.set(__self__, "simple", simple)
        if sort and not isinstance(sort, str):
            raise TypeError("Expected argument 'sort' to be a str")
        pulumi.set(__self__, "sort", sort)
        if starred and not isinstance(starred, bool):
            raise TypeError("Expected argument 'starred' to be a bool")
        pulumi.set(__self__, "starred", starred)
        if statistics and not isinstance(statistics, bool):
            raise TypeError("Expected argument 'statistics' to be a bool")
        pulumi.set(__self__, "statistics", statistics)
        if topics and not isinstance(topics, list):
            raise TypeError("Expected argument 'topics' to be a list")
        pulumi.set(__self__, "topics", topics)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)
        if with_custom_attributes and not isinstance(with_custom_attributes, bool):
            raise TypeError("Expected argument 'with_custom_attributes' to be a bool")
        pulumi.set(__self__, "with_custom_attributes", with_custom_attributes)
        if with_issues_enabled and not isinstance(with_issues_enabled, bool):
            raise TypeError("Expected argument 'with_issues_enabled' to be a bool")
        pulumi.set(__self__, "with_issues_enabled", with_issues_enabled)
        if with_merge_requests_enabled and not isinstance(with_merge_requests_enabled, bool):
            raise TypeError("Expected argument 'with_merge_requests_enabled' to be a bool")
        pulumi.set(__self__, "with_merge_requests_enabled", with_merge_requests_enabled)
        if with_programming_language and not isinstance(with_programming_language, str):
            raise TypeError("Expected argument 'with_programming_language' to be a str")
        pulumi.set(__self__, "with_programming_language", with_programming_language)
        if with_shared and not isinstance(with_shared, bool):
            raise TypeError("Expected argument 'with_shared' to be a bool")
        pulumi.set(__self__, "with_shared", with_shared)

    @property
    @pulumi.getter
    def archived(self) -> Optional[builtins.bool]:
        """
        Limit by archived status.
        """
        return pulumi.get(self, "archived")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[builtins.int]:
        """
        The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `min_access_level`, `with_programming_language` or `statistics`.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeSubgroups")
    def include_subgroups(self) -> Optional[builtins.bool]:
        """
        Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
        """
        return pulumi.get(self, "include_subgroups")

    @property
    @pulumi.getter(name="maxQueryablePages")
    def max_queryable_pages(self) -> Optional[builtins.int]:
        """
        The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
        """
        return pulumi.get(self, "max_queryable_pages")

    @property
    @pulumi.getter
    def membership(self) -> Optional[builtins.bool]:
        """
        Limit by projects that the current user is a member of.
        """
        return pulumi.get(self, "membership")

    @property
    @pulumi.getter(name="minAccessLevel")
    def min_access_level(self) -> Optional[builtins.int]:
        """
        Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/api/members/) for values. Cannot be used with `group_id`.
        """
        return pulumi.get(self, "min_access_level")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[builtins.str]:
        """
        Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/api/projects/#list-all-projects) for details.
        """
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter
    def owned(self) -> Optional[builtins.bool]:
        """
        Limit by projects owned by the current user.
        """
        return pulumi.get(self, "owned")

    @property
    @pulumi.getter
    def page(self) -> Optional[builtins.int]:
        """
        The first page to begin the query on.
        """
        return pulumi.get(self, "page")

    @property
    @pulumi.getter(name="perPage")
    def per_page(self) -> Optional[builtins.int]:
        """
        The number of results to return per page.
        """
        return pulumi.get(self, "per_page")

    @property
    @pulumi.getter
    def projects(self) -> Sequence['outputs.GetProjectsProjectResult']:
        """
        A list containing the projects matching the supplied arguments
        """
        return pulumi.get(self, "projects")

    @property
    @pulumi.getter
    def search(self) -> Optional[builtins.str]:
        """
        Return list of authorized projects matching the search criteria.
        """
        return pulumi.get(self, "search")

    @property
    @pulumi.getter
    def simple(self) -> Optional[builtins.bool]:
        """
        Return only the ID, URL, name, and path of each project.
        """
        return pulumi.get(self, "simple")

    @property
    @pulumi.getter
    def sort(self) -> Optional[builtins.str]:
        """
        Return projects sorted in `asc` or `desc` order. Default is `desc`.
        """
        return pulumi.get(self, "sort")

    @property
    @pulumi.getter
    def starred(self) -> Optional[builtins.bool]:
        """
        Limit by projects starred by the current user.
        """
        return pulumi.get(self, "starred")

    @property
    @pulumi.getter
    def statistics(self) -> Optional[builtins.bool]:
        """
        Include project statistics. Cannot be used with `group_id`.
        """
        return pulumi.get(self, "statistics")

    @property
    @pulumi.getter
    def topics(self) -> Optional[Sequence[builtins.str]]:
        """
        Limit by projects that have all of the given topics.
        """
        return pulumi.get(self, "topics")

    @property
    @pulumi.getter
    def visibility(self) -> Optional[builtins.str]:
        """
        Limit by visibility `public`, `internal`, or `private`.
        """
        return pulumi.get(self, "visibility")

    @property
    @pulumi.getter(name="withCustomAttributes")
    def with_custom_attributes(self) -> Optional[builtins.bool]:
        """
        Include custom attributes in response *(admins only)*.
        """
        return pulumi.get(self, "with_custom_attributes")

    @property
    @pulumi.getter(name="withIssuesEnabled")
    def with_issues_enabled(self) -> Optional[builtins.bool]:
        """
        Limit by projects with issues feature enabled. Default is `false`.
        """
        return pulumi.get(self, "with_issues_enabled")

    @property
    @pulumi.getter(name="withMergeRequestsEnabled")
    def with_merge_requests_enabled(self) -> Optional[builtins.bool]:
        """
        Limit by projects with merge requests feature enabled. Default is `false`.
        """
        return pulumi.get(self, "with_merge_requests_enabled")

    @property
    @pulumi.getter(name="withProgrammingLanguage")
    def with_programming_language(self) -> Optional[builtins.str]:
        """
        Limit by projects which use the given programming language. Cannot be used with `group_id`.
        """
        return pulumi.get(self, "with_programming_language")

    @property
    @pulumi.getter(name="withShared")
    def with_shared(self) -> Optional[builtins.bool]:
        """
        Include projects shared to this group. Default is `true`. Needs `group_id`.
        """
        return pulumi.get(self, "with_shared")


class AwaitableGetProjectsResult(GetProjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectsResult(
            archived=self.archived,
            group_id=self.group_id,
            id=self.id,
            include_subgroups=self.include_subgroups,
            max_queryable_pages=self.max_queryable_pages,
            membership=self.membership,
            min_access_level=self.min_access_level,
            order_by=self.order_by,
            owned=self.owned,
            page=self.page,
            per_page=self.per_page,
            projects=self.projects,
            search=self.search,
            simple=self.simple,
            sort=self.sort,
            starred=self.starred,
            statistics=self.statistics,
            topics=self.topics,
            visibility=self.visibility,
            with_custom_attributes=self.with_custom_attributes,
            with_issues_enabled=self.with_issues_enabled,
            with_merge_requests_enabled=self.with_merge_requests_enabled,
            with_programming_language=self.with_programming_language,
            with_shared=self.with_shared)


def get_projects(archived: Optional[builtins.bool] = None,
                 group_id: Optional[builtins.int] = None,
                 include_subgroups: Optional[builtins.bool] = None,
                 max_queryable_pages: Optional[builtins.int] = None,
                 membership: Optional[builtins.bool] = None,
                 min_access_level: Optional[builtins.int] = None,
                 order_by: Optional[builtins.str] = None,
                 owned: Optional[builtins.bool] = None,
                 page: Optional[builtins.int] = None,
                 per_page: Optional[builtins.int] = None,
                 search: Optional[builtins.str] = None,
                 simple: Optional[builtins.bool] = None,
                 sort: Optional[builtins.str] = None,
                 starred: Optional[builtins.bool] = None,
                 statistics: Optional[builtins.bool] = None,
                 topics: Optional[Sequence[builtins.str]] = None,
                 visibility: Optional[builtins.str] = None,
                 with_custom_attributes: Optional[builtins.bool] = None,
                 with_issues_enabled: Optional[builtins.bool] = None,
                 with_merge_requests_enabled: Optional[builtins.bool] = None,
                 with_programming_language: Optional[builtins.str] = None,
                 with_shared: Optional[builtins.bool] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectsResult:
    """
    The `get_projects` data source allows details of multiple projects to be retrieved. Optionally filtered by the set attributes.

    > This data source supports all available filters exposed by the [client-go](https://gitlab.com/gitlab-org/api/client-go) package, which might not expose all available filters exposed by the GitLab APIs.

    > The owner sub-attributes are only populated if the GitLab token used has an administrator scope.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/projects/#list-all-projects)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    # List projects within a group tree
    mygroup = gitlab.get_group(full_path="mygroup")
    group_projects = gitlab.get_projects(group_id=mygroup.id,
        order_by="name",
        include_subgroups=True,
        with_shared=False)
    # List projects using the search syntax
    projects = gitlab.get_projects(search="postgresql",
        visibility="private")
    ```


    :param builtins.bool archived: Limit by archived status.
    :param builtins.int group_id: The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `min_access_level`, `with_programming_language` or `statistics`.
    :param builtins.bool include_subgroups: Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
    :param builtins.int max_queryable_pages: The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
    :param builtins.bool membership: Limit by projects that the current user is a member of.
    :param builtins.int min_access_level: Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/api/members/) for values. Cannot be used with `group_id`.
    :param builtins.str order_by: Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/api/projects/#list-all-projects) for details.
    :param builtins.bool owned: Limit by projects owned by the current user.
    :param builtins.int page: The first page to begin the query on.
    :param builtins.int per_page: The number of results to return per page.
    :param builtins.str search: Return list of authorized projects matching the search criteria.
    :param builtins.bool simple: Return only the ID, URL, name, and path of each project.
    :param builtins.str sort: Return projects sorted in `asc` or `desc` order. Default is `desc`.
    :param builtins.bool starred: Limit by projects starred by the current user.
    :param builtins.bool statistics: Include project statistics. Cannot be used with `group_id`.
    :param Sequence[builtins.str] topics: Limit by projects that have all of the given topics.
    :param builtins.str visibility: Limit by visibility `public`, `internal`, or `private`.
    :param builtins.bool with_custom_attributes: Include custom attributes in response *(admins only)*.
    :param builtins.bool with_issues_enabled: Limit by projects with issues feature enabled. Default is `false`.
    :param builtins.bool with_merge_requests_enabled: Limit by projects with merge requests feature enabled. Default is `false`.
    :param builtins.str with_programming_language: Limit by projects which use the given programming language. Cannot be used with `group_id`.
    :param builtins.bool with_shared: Include projects shared to this group. Default is `true`. Needs `group_id`.
    """
    __args__ = dict()
    __args__['archived'] = archived
    __args__['groupId'] = group_id
    __args__['includeSubgroups'] = include_subgroups
    __args__['maxQueryablePages'] = max_queryable_pages
    __args__['membership'] = membership
    __args__['minAccessLevel'] = min_access_level
    __args__['orderBy'] = order_by
    __args__['owned'] = owned
    __args__['page'] = page
    __args__['perPage'] = per_page
    __args__['search'] = search
    __args__['simple'] = simple
    __args__['sort'] = sort
    __args__['starred'] = starred
    __args__['statistics'] = statistics
    __args__['topics'] = topics
    __args__['visibility'] = visibility
    __args__['withCustomAttributes'] = with_custom_attributes
    __args__['withIssuesEnabled'] = with_issues_enabled
    __args__['withMergeRequestsEnabled'] = with_merge_requests_enabled
    __args__['withProgrammingLanguage'] = with_programming_language
    __args__['withShared'] = with_shared
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjects:getProjects', __args__, opts=opts, typ=GetProjectsResult).value

    return AwaitableGetProjectsResult(
        archived=pulumi.get(__ret__, 'archived'),
        group_id=pulumi.get(__ret__, 'group_id'),
        id=pulumi.get(__ret__, 'id'),
        include_subgroups=pulumi.get(__ret__, 'include_subgroups'),
        max_queryable_pages=pulumi.get(__ret__, 'max_queryable_pages'),
        membership=pulumi.get(__ret__, 'membership'),
        min_access_level=pulumi.get(__ret__, 'min_access_level'),
        order_by=pulumi.get(__ret__, 'order_by'),
        owned=pulumi.get(__ret__, 'owned'),
        page=pulumi.get(__ret__, 'page'),
        per_page=pulumi.get(__ret__, 'per_page'),
        projects=pulumi.get(__ret__, 'projects'),
        search=pulumi.get(__ret__, 'search'),
        simple=pulumi.get(__ret__, 'simple'),
        sort=pulumi.get(__ret__, 'sort'),
        starred=pulumi.get(__ret__, 'starred'),
        statistics=pulumi.get(__ret__, 'statistics'),
        topics=pulumi.get(__ret__, 'topics'),
        visibility=pulumi.get(__ret__, 'visibility'),
        with_custom_attributes=pulumi.get(__ret__, 'with_custom_attributes'),
        with_issues_enabled=pulumi.get(__ret__, 'with_issues_enabled'),
        with_merge_requests_enabled=pulumi.get(__ret__, 'with_merge_requests_enabled'),
        with_programming_language=pulumi.get(__ret__, 'with_programming_language'),
        with_shared=pulumi.get(__ret__, 'with_shared'))
def get_projects_output(archived: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        group_id: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                        include_subgroups: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        max_queryable_pages: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                        membership: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        min_access_level: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                        order_by: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        owned: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        page: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                        per_page: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                        search: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        simple: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        sort: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        starred: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        statistics: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        topics: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                        visibility: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        with_custom_attributes: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        with_issues_enabled: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        with_merge_requests_enabled: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        with_programming_language: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        with_shared: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetProjectsResult]:
    """
    The `get_projects` data source allows details of multiple projects to be retrieved. Optionally filtered by the set attributes.

    > This data source supports all available filters exposed by the [client-go](https://gitlab.com/gitlab-org/api/client-go) package, which might not expose all available filters exposed by the GitLab APIs.

    > The owner sub-attributes are only populated if the GitLab token used has an administrator scope.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/projects/#list-all-projects)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    # List projects within a group tree
    mygroup = gitlab.get_group(full_path="mygroup")
    group_projects = gitlab.get_projects(group_id=mygroup.id,
        order_by="name",
        include_subgroups=True,
        with_shared=False)
    # List projects using the search syntax
    projects = gitlab.get_projects(search="postgresql",
        visibility="private")
    ```


    :param builtins.bool archived: Limit by archived status.
    :param builtins.int group_id: The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `min_access_level`, `with_programming_language` or `statistics`.
    :param builtins.bool include_subgroups: Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
    :param builtins.int max_queryable_pages: The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
    :param builtins.bool membership: Limit by projects that the current user is a member of.
    :param builtins.int min_access_level: Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/api/members/) for values. Cannot be used with `group_id`.
    :param builtins.str order_by: Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/api/projects/#list-all-projects) for details.
    :param builtins.bool owned: Limit by projects owned by the current user.
    :param builtins.int page: The first page to begin the query on.
    :param builtins.int per_page: The number of results to return per page.
    :param builtins.str search: Return list of authorized projects matching the search criteria.
    :param builtins.bool simple: Return only the ID, URL, name, and path of each project.
    :param builtins.str sort: Return projects sorted in `asc` or `desc` order. Default is `desc`.
    :param builtins.bool starred: Limit by projects starred by the current user.
    :param builtins.bool statistics: Include project statistics. Cannot be used with `group_id`.
    :param Sequence[builtins.str] topics: Limit by projects that have all of the given topics.
    :param builtins.str visibility: Limit by visibility `public`, `internal`, or `private`.
    :param builtins.bool with_custom_attributes: Include custom attributes in response *(admins only)*.
    :param builtins.bool with_issues_enabled: Limit by projects with issues feature enabled. Default is `false`.
    :param builtins.bool with_merge_requests_enabled: Limit by projects with merge requests feature enabled. Default is `false`.
    :param builtins.str with_programming_language: Limit by projects which use the given programming language. Cannot be used with `group_id`.
    :param builtins.bool with_shared: Include projects shared to this group. Default is `true`. Needs `group_id`.
    """
    __args__ = dict()
    __args__['archived'] = archived
    __args__['groupId'] = group_id
    __args__['includeSubgroups'] = include_subgroups
    __args__['maxQueryablePages'] = max_queryable_pages
    __args__['membership'] = membership
    __args__['minAccessLevel'] = min_access_level
    __args__['orderBy'] = order_by
    __args__['owned'] = owned
    __args__['page'] = page
    __args__['perPage'] = per_page
    __args__['search'] = search
    __args__['simple'] = simple
    __args__['sort'] = sort
    __args__['starred'] = starred
    __args__['statistics'] = statistics
    __args__['topics'] = topics
    __args__['visibility'] = visibility
    __args__['withCustomAttributes'] = with_custom_attributes
    __args__['withIssuesEnabled'] = with_issues_enabled
    __args__['withMergeRequestsEnabled'] = with_merge_requests_enabled
    __args__['withProgrammingLanguage'] = with_programming_language
    __args__['withShared'] = with_shared
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getProjects:getProjects', __args__, opts=opts, typ=GetProjectsResult)
    return __ret__.apply(lambda __response__: GetProjectsResult(
        archived=pulumi.get(__response__, 'archived'),
        group_id=pulumi.get(__response__, 'group_id'),
        id=pulumi.get(__response__, 'id'),
        include_subgroups=pulumi.get(__response__, 'include_subgroups'),
        max_queryable_pages=pulumi.get(__response__, 'max_queryable_pages'),
        membership=pulumi.get(__response__, 'membership'),
        min_access_level=pulumi.get(__response__, 'min_access_level'),
        order_by=pulumi.get(__response__, 'order_by'),
        owned=pulumi.get(__response__, 'owned'),
        page=pulumi.get(__response__, 'page'),
        per_page=pulumi.get(__response__, 'per_page'),
        projects=pulumi.get(__response__, 'projects'),
        search=pulumi.get(__response__, 'search'),
        simple=pulumi.get(__response__, 'simple'),
        sort=pulumi.get(__response__, 'sort'),
        starred=pulumi.get(__response__, 'starred'),
        statistics=pulumi.get(__response__, 'statistics'),
        topics=pulumi.get(__response__, 'topics'),
        visibility=pulumi.get(__response__, 'visibility'),
        with_custom_attributes=pulumi.get(__response__, 'with_custom_attributes'),
        with_issues_enabled=pulumi.get(__response__, 'with_issues_enabled'),
        with_merge_requests_enabled=pulumi.get(__response__, 'with_merge_requests_enabled'),
        with_programming_language=pulumi.get(__response__, 'with_programming_language'),
        with_shared=pulumi.get(__response__, 'with_shared')))
