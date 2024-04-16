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
    'GetUsersResult',
    'AwaitableGetUsersResult',
    'get_users',
    'get_users_output',
]

@pulumi.output_type
class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, active=None, blocked=None, created_after=None, created_before=None, extern_provider=None, extern_uid=None, id=None, order_by=None, search=None, sort=None, users=None):
        if active and not isinstance(active, bool):
            raise TypeError("Expected argument 'active' to be a bool")
        pulumi.set(__self__, "active", active)
        if blocked and not isinstance(blocked, bool):
            raise TypeError("Expected argument 'blocked' to be a bool")
        pulumi.set(__self__, "blocked", blocked)
        if created_after and not isinstance(created_after, str):
            raise TypeError("Expected argument 'created_after' to be a str")
        pulumi.set(__self__, "created_after", created_after)
        if created_before and not isinstance(created_before, str):
            raise TypeError("Expected argument 'created_before' to be a str")
        pulumi.set(__self__, "created_before", created_before)
        if extern_provider and not isinstance(extern_provider, str):
            raise TypeError("Expected argument 'extern_provider' to be a str")
        pulumi.set(__self__, "extern_provider", extern_provider)
        if extern_uid and not isinstance(extern_uid, str):
            raise TypeError("Expected argument 'extern_uid' to be a str")
        pulumi.set(__self__, "extern_uid", extern_uid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if search and not isinstance(search, str):
            raise TypeError("Expected argument 'search' to be a str")
        pulumi.set(__self__, "search", search)
        if sort and not isinstance(sort, str):
            raise TypeError("Expected argument 'sort' to be a str")
        pulumi.set(__self__, "sort", sort)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def active(self) -> Optional[bool]:
        """
        Filter users that are active.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter
    def blocked(self) -> Optional[bool]:
        """
        Filter users that are blocked.
        """
        return pulumi.get(self, "blocked")

    @property
    @pulumi.getter(name="createdAfter")
    def created_after(self) -> Optional[str]:
        """
        Search for users created after a specific date. (Requires administrator privileges)
        """
        return pulumi.get(self, "created_after")

    @property
    @pulumi.getter(name="createdBefore")
    def created_before(self) -> Optional[str]:
        """
        Search for users created before a specific date. (Requires administrator privileges)
        """
        return pulumi.get(self, "created_before")

    @property
    @pulumi.getter(name="externProvider")
    def extern_provider(self) -> Optional[str]:
        """
        Lookup users by external provider. (Requires administrator privileges)
        """
        return pulumi.get(self, "extern_provider")

    @property
    @pulumi.getter(name="externUid")
    def extern_uid(self) -> Optional[str]:
        """
        Lookup users by external UID. (Requires administrator privileges)
        """
        return pulumi.get(self, "extern_uid")

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
        Order the users' list by `id`, `name`, `username`, `created_at` or `updated_at`. (Requires administrator privileges)
        """
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter
    def search(self) -> Optional[str]:
        """
        Search users by username, name or email.
        """
        return pulumi.get(self, "search")

    @property
    @pulumi.getter
    def sort(self) -> Optional[str]:
        """
        Sort users' list in asc or desc order. (Requires administrator privileges)
        """
        return pulumi.get(self, "sort")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetUsersUserResult']:
        """
        The list of users.
        """
        return pulumi.get(self, "users")


class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            active=self.active,
            blocked=self.blocked,
            created_after=self.created_after,
            created_before=self.created_before,
            extern_provider=self.extern_provider,
            extern_uid=self.extern_uid,
            id=self.id,
            order_by=self.order_by,
            search=self.search,
            sort=self.sort,
            users=self.users)


def get_users(active: Optional[bool] = None,
              blocked: Optional[bool] = None,
              created_after: Optional[str] = None,
              created_before: Optional[str] = None,
              extern_provider: Optional[str] = None,
              extern_uid: Optional[str] = None,
              order_by: Optional[str] = None,
              search: Optional[str] = None,
              sort: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUsersResult:
    """
    The `get_users` data source allows details of multiple users to be retrieved given some optional filter criteria.

    > Some attributes might not be returned depending on if you're an admin or not.

    > Some available options require administrator privileges.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ce/api/users.html#list-users)

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_users(sort="desc",
        order_by="name",
        created_before="2019-01-01")
    example_two = gitlab.get_users(search="username")
    ```
    <!--End PulumiCodeChooser -->


    :param bool active: Filter users that are active.
    :param bool blocked: Filter users that are blocked.
    :param str created_after: Search for users created after a specific date. (Requires administrator privileges)
    :param str created_before: Search for users created before a specific date. (Requires administrator privileges)
    :param str extern_provider: Lookup users by external provider. (Requires administrator privileges)
    :param str order_by: Order the users' list by `id`, `name`, `username`, `created_at` or `updated_at`. (Requires administrator privileges)
    :param str search: Search users by username, name or email.
    :param str sort: Sort users' list in asc or desc order. (Requires administrator privileges)
    """
    __args__ = dict()
    __args__['active'] = active
    __args__['blocked'] = blocked
    __args__['createdAfter'] = created_after
    __args__['createdBefore'] = created_before
    __args__['externProvider'] = extern_provider
    __args__['externUid'] = extern_uid
    __args__['orderBy'] = order_by
    __args__['search'] = search
    __args__['sort'] = sort
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult).value

    return AwaitableGetUsersResult(
        active=pulumi.get(__ret__, 'active'),
        blocked=pulumi.get(__ret__, 'blocked'),
        created_after=pulumi.get(__ret__, 'created_after'),
        created_before=pulumi.get(__ret__, 'created_before'),
        extern_provider=pulumi.get(__ret__, 'extern_provider'),
        extern_uid=pulumi.get(__ret__, 'extern_uid'),
        id=pulumi.get(__ret__, 'id'),
        order_by=pulumi.get(__ret__, 'order_by'),
        search=pulumi.get(__ret__, 'search'),
        sort=pulumi.get(__ret__, 'sort'),
        users=pulumi.get(__ret__, 'users'))


@_utilities.lift_output_func(get_users)
def get_users_output(active: Optional[pulumi.Input[Optional[bool]]] = None,
                     blocked: Optional[pulumi.Input[Optional[bool]]] = None,
                     created_after: Optional[pulumi.Input[Optional[str]]] = None,
                     created_before: Optional[pulumi.Input[Optional[str]]] = None,
                     extern_provider: Optional[pulumi.Input[Optional[str]]] = None,
                     extern_uid: Optional[pulumi.Input[Optional[str]]] = None,
                     order_by: Optional[pulumi.Input[Optional[str]]] = None,
                     search: Optional[pulumi.Input[Optional[str]]] = None,
                     sort: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUsersResult]:
    """
    The `get_users` data source allows details of multiple users to be retrieved given some optional filter criteria.

    > Some attributes might not be returned depending on if you're an admin or not.

    > Some available options require administrator privileges.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ce/api/users.html#list-users)

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_users(sort="desc",
        order_by="name",
        created_before="2019-01-01")
    example_two = gitlab.get_users(search="username")
    ```
    <!--End PulumiCodeChooser -->


    :param bool active: Filter users that are active.
    :param bool blocked: Filter users that are blocked.
    :param str created_after: Search for users created after a specific date. (Requires administrator privileges)
    :param str created_before: Search for users created before a specific date. (Requires administrator privileges)
    :param str extern_provider: Lookup users by external provider. (Requires administrator privileges)
    :param str order_by: Order the users' list by `id`, `name`, `username`, `created_at` or `updated_at`. (Requires administrator privileges)
    :param str search: Search users by username, name or email.
    :param str sort: Sort users' list in asc or desc order. (Requires administrator privileges)
    """
    ...
