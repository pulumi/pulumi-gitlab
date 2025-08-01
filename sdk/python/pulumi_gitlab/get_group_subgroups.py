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
    'GetGroupSubgroupsResult',
    'AwaitableGetGroupSubgroupsResult',
    'get_group_subgroups',
    'get_group_subgroups_output',
]

@pulumi.output_type
class GetGroupSubgroupsResult:
    """
    A collection of values returned by getGroupSubgroups.
    """
    def __init__(__self__, all_available=None, group_id=None, id=None, min_access_level=None, order_by=None, owned=None, search=None, skip_groups=None, sort=None, statistics=None, subgroups=None, with_custom_attributes=None):
        if all_available and not isinstance(all_available, bool):
            raise TypeError("Expected argument 'all_available' to be a bool")
        pulumi.set(__self__, "all_available", all_available)
        if group_id and not isinstance(group_id, int):
            raise TypeError("Expected argument 'group_id' to be a int")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if min_access_level and not isinstance(min_access_level, str):
            raise TypeError("Expected argument 'min_access_level' to be a str")
        pulumi.set(__self__, "min_access_level", min_access_level)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if owned and not isinstance(owned, bool):
            raise TypeError("Expected argument 'owned' to be a bool")
        pulumi.set(__self__, "owned", owned)
        if search and not isinstance(search, str):
            raise TypeError("Expected argument 'search' to be a str")
        pulumi.set(__self__, "search", search)
        if skip_groups and not isinstance(skip_groups, list):
            raise TypeError("Expected argument 'skip_groups' to be a list")
        pulumi.set(__self__, "skip_groups", skip_groups)
        if sort and not isinstance(sort, str):
            raise TypeError("Expected argument 'sort' to be a str")
        pulumi.set(__self__, "sort", sort)
        if statistics and not isinstance(statistics, bool):
            raise TypeError("Expected argument 'statistics' to be a bool")
        pulumi.set(__self__, "statistics", statistics)
        if subgroups and not isinstance(subgroups, list):
            raise TypeError("Expected argument 'subgroups' to be a list")
        pulumi.set(__self__, "subgroups", subgroups)
        if with_custom_attributes and not isinstance(with_custom_attributes, bool):
            raise TypeError("Expected argument 'with_custom_attributes' to be a bool")
        pulumi.set(__self__, "with_custom_attributes", with_custom_attributes)

    @_builtins.property
    @pulumi.getter(name="allAvailable")
    def all_available(self) -> _builtins.bool:
        """
        Show all the groups you have access to.
        """
        return pulumi.get(self, "all_available")

    @_builtins.property
    @pulumi.getter(name="groupId")
    def group_id(self) -> _builtins.int:
        """
        The ID of the group.
        """
        return pulumi.get(self, "group_id")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter(name="minAccessLevel")
    def min_access_level(self) -> _builtins.str:
        """
        Limit to groups where current user has at least this access level.
        """
        return pulumi.get(self, "min_access_level")

    @_builtins.property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> _builtins.str:
        """
        Order groups by name, path or id.
        """
        return pulumi.get(self, "order_by")

    @_builtins.property
    @pulumi.getter
    def owned(self) -> _builtins.bool:
        """
        Limit to groups explicitly owned by the current user.
        """
        return pulumi.get(self, "owned")

    @_builtins.property
    @pulumi.getter
    def search(self) -> _builtins.str:
        """
        Return the list of authorized groups matching the search criteria.
        """
        return pulumi.get(self, "search")

    @_builtins.property
    @pulumi.getter(name="skipGroups")
    def skip_groups(self) -> Sequence[_builtins.int]:
        """
        Skip the group IDs passed.
        """
        return pulumi.get(self, "skip_groups")

    @_builtins.property
    @pulumi.getter
    def sort(self) -> _builtins.str:
        """
        Order groups in asc or desc order.
        """
        return pulumi.get(self, "sort")

    @_builtins.property
    @pulumi.getter
    def statistics(self) -> _builtins.bool:
        """
        Include group statistics (administrators only).
        """
        return pulumi.get(self, "statistics")

    @_builtins.property
    @pulumi.getter
    def subgroups(self) -> Sequence['outputs.GetGroupSubgroupsSubgroupResult']:
        """
        Subgroups of the parent group.
        """
        return pulumi.get(self, "subgroups")

    @_builtins.property
    @pulumi.getter(name="withCustomAttributes")
    def with_custom_attributes(self) -> _builtins.bool:
        """
        Include custom attributes in response (administrators only).
        """
        return pulumi.get(self, "with_custom_attributes")


class AwaitableGetGroupSubgroupsResult(GetGroupSubgroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupSubgroupsResult(
            all_available=self.all_available,
            group_id=self.group_id,
            id=self.id,
            min_access_level=self.min_access_level,
            order_by=self.order_by,
            owned=self.owned,
            search=self.search,
            skip_groups=self.skip_groups,
            sort=self.sort,
            statistics=self.statistics,
            subgroups=self.subgroups,
            with_custom_attributes=self.with_custom_attributes)


def get_group_subgroups(all_available: Optional[_builtins.bool] = None,
                        group_id: Optional[_builtins.int] = None,
                        min_access_level: Optional[_builtins.str] = None,
                        order_by: Optional[_builtins.str] = None,
                        owned: Optional[_builtins.bool] = None,
                        search: Optional[_builtins.str] = None,
                        skip_groups: Optional[Sequence[_builtins.int]] = None,
                        sort: Optional[_builtins.str] = None,
                        statistics: Optional[_builtins.bool] = None,
                        with_custom_attributes: Optional[_builtins.bool] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupSubgroupsResult:
    """
    The `get_group_subgroups` data source allows to get subgroups of a group.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#list-subgroups)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    subgroups = gitlab.get_group_subgroups(group_id=123456)
    pulumi.export("subgroups", subgroups)
    ```


    :param _builtins.bool all_available: Show all the groups you have access to.
    :param _builtins.int group_id: The ID of the group.
    :param _builtins.str min_access_level: Limit to groups where current user has at least this access level.
    :param _builtins.str order_by: Order groups by name, path or id.
    :param _builtins.bool owned: Limit to groups explicitly owned by the current user.
    :param _builtins.str search: Return the list of authorized groups matching the search criteria.
    :param Sequence[_builtins.int] skip_groups: Skip the group IDs passed.
    :param _builtins.str sort: Order groups in asc or desc order.
    :param _builtins.bool statistics: Include group statistics (administrators only).
    :param _builtins.bool with_custom_attributes: Include custom attributes in response (administrators only).
    """
    __args__ = dict()
    __args__['allAvailable'] = all_available
    __args__['groupId'] = group_id
    __args__['minAccessLevel'] = min_access_level
    __args__['orderBy'] = order_by
    __args__['owned'] = owned
    __args__['search'] = search
    __args__['skipGroups'] = skip_groups
    __args__['sort'] = sort
    __args__['statistics'] = statistics
    __args__['withCustomAttributes'] = with_custom_attributes
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getGroupSubgroups:getGroupSubgroups', __args__, opts=opts, typ=GetGroupSubgroupsResult).value

    return AwaitableGetGroupSubgroupsResult(
        all_available=pulumi.get(__ret__, 'all_available'),
        group_id=pulumi.get(__ret__, 'group_id'),
        id=pulumi.get(__ret__, 'id'),
        min_access_level=pulumi.get(__ret__, 'min_access_level'),
        order_by=pulumi.get(__ret__, 'order_by'),
        owned=pulumi.get(__ret__, 'owned'),
        search=pulumi.get(__ret__, 'search'),
        skip_groups=pulumi.get(__ret__, 'skip_groups'),
        sort=pulumi.get(__ret__, 'sort'),
        statistics=pulumi.get(__ret__, 'statistics'),
        subgroups=pulumi.get(__ret__, 'subgroups'),
        with_custom_attributes=pulumi.get(__ret__, 'with_custom_attributes'))
def get_group_subgroups_output(all_available: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                               group_id: Optional[pulumi.Input[_builtins.int]] = None,
                               min_access_level: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                               order_by: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                               owned: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                               search: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                               skip_groups: Optional[pulumi.Input[Optional[Sequence[_builtins.int]]]] = None,
                               sort: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                               statistics: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                               with_custom_attributes: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGroupSubgroupsResult]:
    """
    The `get_group_subgroups` data source allows to get subgroups of a group.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#list-subgroups)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    subgroups = gitlab.get_group_subgroups(group_id=123456)
    pulumi.export("subgroups", subgroups)
    ```


    :param _builtins.bool all_available: Show all the groups you have access to.
    :param _builtins.int group_id: The ID of the group.
    :param _builtins.str min_access_level: Limit to groups where current user has at least this access level.
    :param _builtins.str order_by: Order groups by name, path or id.
    :param _builtins.bool owned: Limit to groups explicitly owned by the current user.
    :param _builtins.str search: Return the list of authorized groups matching the search criteria.
    :param Sequence[_builtins.int] skip_groups: Skip the group IDs passed.
    :param _builtins.str sort: Order groups in asc or desc order.
    :param _builtins.bool statistics: Include group statistics (administrators only).
    :param _builtins.bool with_custom_attributes: Include custom attributes in response (administrators only).
    """
    __args__ = dict()
    __args__['allAvailable'] = all_available
    __args__['groupId'] = group_id
    __args__['minAccessLevel'] = min_access_level
    __args__['orderBy'] = order_by
    __args__['owned'] = owned
    __args__['search'] = search
    __args__['skipGroups'] = skip_groups
    __args__['sort'] = sort
    __args__['statistics'] = statistics
    __args__['withCustomAttributes'] = with_custom_attributes
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getGroupSubgroups:getGroupSubgroups', __args__, opts=opts, typ=GetGroupSubgroupsResult)
    return __ret__.apply(lambda __response__: GetGroupSubgroupsResult(
        all_available=pulumi.get(__response__, 'all_available'),
        group_id=pulumi.get(__response__, 'group_id'),
        id=pulumi.get(__response__, 'id'),
        min_access_level=pulumi.get(__response__, 'min_access_level'),
        order_by=pulumi.get(__response__, 'order_by'),
        owned=pulumi.get(__response__, 'owned'),
        search=pulumi.get(__response__, 'search'),
        skip_groups=pulumi.get(__response__, 'skip_groups'),
        sort=pulumi.get(__response__, 'sort'),
        statistics=pulumi.get(__response__, 'statistics'),
        subgroups=pulumi.get(__response__, 'subgroups'),
        with_custom_attributes=pulumi.get(__response__, 'with_custom_attributes')))
