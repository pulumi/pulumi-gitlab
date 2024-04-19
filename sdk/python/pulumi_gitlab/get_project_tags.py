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
    'GetProjectTagsResult',
    'AwaitableGetProjectTagsResult',
    'get_project_tags',
    'get_project_tags_output',
]

@pulumi.output_type
class GetProjectTagsResult:
    """
    A collection of values returned by getProjectTags.
    """
    def __init__(__self__, id=None, order_by=None, project=None, search=None, sort=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if search and not isinstance(search, str):
            raise TypeError("Expected argument 'search' to be a str")
        pulumi.set(__self__, "search", search)
        if sort and not isinstance(sort, str):
            raise TypeError("Expected argument 'sort' to be a str")
        pulumi.set(__self__, "sort", sort)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        """
        Return tags ordered by `name` or `updated` fields. Default is `updated`.
        """
        return pulumi.get(self, "order_by")

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
        Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
        """
        return pulumi.get(self, "search")

    @property
    @pulumi.getter
    def sort(self) -> Optional[str]:
        """
        Return tags sorted in `asc` or `desc` order. Default is `desc`.
        """
        return pulumi.get(self, "sort")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GetProjectTagsTagResult']:
        """
        List of repository tags from a project.
        """
        return pulumi.get(self, "tags")


class AwaitableGetProjectTagsResult(GetProjectTagsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectTagsResult(
            id=self.id,
            order_by=self.order_by,
            project=self.project,
            search=self.search,
            sort=self.sort,
            tags=self.tags)


def get_project_tags(order_by: Optional[str] = None,
                     project: Optional[str] = None,
                     search: Optional[str] = None,
                     sort: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectTagsResult:
    """
    The `get_project_tags` data source allows details of project tags to be retrieved by some search criteria.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/tags.html#list-project-repository-tags)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_project_tags(project="foo/bar")
    ```


    :param str order_by: Return tags ordered by `name` or `updated` fields. Default is `updated`.
    :param str project: The ID or URL-encoded path of the project owned by the authenticated user.
    :param str search: Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
    :param str sort: Return tags sorted in `asc` or `desc` order. Default is `desc`.
    """
    __args__ = dict()
    __args__['orderBy'] = order_by
    __args__['project'] = project
    __args__['search'] = search
    __args__['sort'] = sort
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjectTags:getProjectTags', __args__, opts=opts, typ=GetProjectTagsResult).value

    return AwaitableGetProjectTagsResult(
        id=pulumi.get(__ret__, 'id'),
        order_by=pulumi.get(__ret__, 'order_by'),
        project=pulumi.get(__ret__, 'project'),
        search=pulumi.get(__ret__, 'search'),
        sort=pulumi.get(__ret__, 'sort'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_project_tags)
def get_project_tags_output(order_by: Optional[pulumi.Input[Optional[str]]] = None,
                            project: Optional[pulumi.Input[str]] = None,
                            search: Optional[pulumi.Input[Optional[str]]] = None,
                            sort: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectTagsResult]:
    """
    The `get_project_tags` data source allows details of project tags to be retrieved by some search criteria.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/tags.html#list-project-repository-tags)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_project_tags(project="foo/bar")
    ```


    :param str order_by: Return tags ordered by `name` or `updated` fields. Default is `updated`.
    :param str project: The ID or URL-encoded path of the project owned by the authenticated user.
    :param str search: Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
    :param str sort: Return tags sorted in `asc` or `desc` order. Default is `desc`.
    """
    ...
