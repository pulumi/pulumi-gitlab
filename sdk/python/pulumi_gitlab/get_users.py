# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, active=None, blocked=None, created_after=None, created_before=None, extern_provider=None, extern_uid=None, id=None, order_by=None, search=None, sort=None, users=None):
        if active and not isinstance(active, bool):
            raise TypeError("Expected argument 'active' to be a bool")
        __self__.active = active
        if blocked and not isinstance(blocked, bool):
            raise TypeError("Expected argument 'blocked' to be a bool")
        __self__.blocked = blocked
        if created_after and not isinstance(created_after, str):
            raise TypeError("Expected argument 'created_after' to be a str")
        __self__.created_after = created_after
        if created_before and not isinstance(created_before, str):
            raise TypeError("Expected argument 'created_before' to be a str")
        __self__.created_before = created_before
        if extern_provider and not isinstance(extern_provider, str):
            raise TypeError("Expected argument 'extern_provider' to be a str")
        __self__.extern_provider = extern_provider
        if extern_uid and not isinstance(extern_uid, str):
            raise TypeError("Expected argument 'extern_uid' to be a str")
        __self__.extern_uid = extern_uid
        """
        The external UID of the user.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        __self__.order_by = order_by
        if search and not isinstance(search, str):
            raise TypeError("Expected argument 'search' to be a str")
        __self__.search = search
        if sort and not isinstance(sort, str):
            raise TypeError("Expected argument 'sort' to be a str")
        __self__.sort = sort
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        __self__.users = users
        """
        The list of users.
        """
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

def get_users(active=None,blocked=None,created_after=None,created_before=None,extern_provider=None,extern_uid=None,order_by=None,search=None,sort=None,opts=None):
    """
    Provides details about a list of users in the gitlab provider. The results include id, username, email, name and more about the requested users. Users can also be sorted and filtered using several options.

    **NOTE**: Some of the available options require administrator privileges. Please visit [Gitlab API documentation][users_for_admins] for more information.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/d/users.html.markdown.


    :param bool active: Filter users that are active.
    :param bool blocked: Filter users that are blocked.
    :param str created_after: Search for users created after a specific date. (Requires administrator privileges)
    :param str created_before: Search for users created before a specific date. (Requires administrator privileges)
    :param str extern_provider: Lookup users by external provider. (Requires administrator privileges)
    :param str extern_uid: Lookup users by external UID. (Requires administrator privileges)
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
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gitlab:index/getUsers:getUsers', __args__, opts=opts).value

    return AwaitableGetUsersResult(
        active=__ret__.get('active'),
        blocked=__ret__.get('blocked'),
        created_after=__ret__.get('createdAfter'),
        created_before=__ret__.get('createdBefore'),
        extern_provider=__ret__.get('externProvider'),
        extern_uid=__ret__.get('externUid'),
        id=__ret__.get('id'),
        order_by=__ret__.get('orderBy'),
        search=__ret__.get('search'),
        sort=__ret__.get('sort'),
        users=__ret__.get('users'))
