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

__all__ = [
    'GetGroupServiceAccountResult',
    'AwaitableGetGroupServiceAccountResult',
    'get_group_service_account',
    'get_group_service_account_output',
]

@pulumi.output_type
class GetGroupServiceAccountResult:
    """
    A collection of values returned by getGroupServiceAccount.
    """
    def __init__(__self__, group=None, id=None, name=None, service_account_id=None, username=None):
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        pulumi.set(__self__, "group", group)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_account_id and not isinstance(service_account_id, str):
            raise TypeError("Expected argument 'service_account_id' to be a str")
        pulumi.set(__self__, "service_account_id", service_account_id)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def group(self) -> str:
        """
        The ID or URL-encoded path of the target group. Must be a top-level group.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the user. If not specified, the default Service account user name is used.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> str:
        """
        The service account id.
        """
        return pulumi.get(self, "service_account_id")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        """
        The username of the user. If not specified, it's automatically generated.
        """
        return pulumi.get(self, "username")


class AwaitableGetGroupServiceAccountResult(GetGroupServiceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupServiceAccountResult(
            group=self.group,
            id=self.id,
            name=self.name,
            service_account_id=self.service_account_id,
            username=self.username)


def get_group_service_account(group: Optional[str] = None,
                              name: Optional[str] = None,
                              service_account_id: Optional[str] = None,
                              username: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupServiceAccountResult:
    """
    The `GroupServiceAccount` data source retrieves information about a gitlab service account for a group.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_service_accounts.html#list-service-account-users)


    :param str group: The ID or URL-encoded path of the target group. Must be a top-level group.
    :param str name: The name of the user. If not specified, the default Service account user name is used.
    :param str service_account_id: The service account id.
    :param str username: The username of the user. If not specified, it's automatically generated.
    """
    __args__ = dict()
    __args__['group'] = group
    __args__['name'] = name
    __args__['serviceAccountId'] = service_account_id
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getGroupServiceAccount:getGroupServiceAccount', __args__, opts=opts, typ=GetGroupServiceAccountResult).value

    return AwaitableGetGroupServiceAccountResult(
        group=pulumi.get(__ret__, 'group'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        service_account_id=pulumi.get(__ret__, 'service_account_id'),
        username=pulumi.get(__ret__, 'username'))
def get_group_service_account_output(group: Optional[pulumi.Input[str]] = None,
                                     name: Optional[pulumi.Input[Optional[str]]] = None,
                                     service_account_id: Optional[pulumi.Input[str]] = None,
                                     username: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGroupServiceAccountResult]:
    """
    The `GroupServiceAccount` data source retrieves information about a gitlab service account for a group.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_service_accounts.html#list-service-account-users)


    :param str group: The ID or URL-encoded path of the target group. Must be a top-level group.
    :param str name: The name of the user. If not specified, the default Service account user name is used.
    :param str service_account_id: The service account id.
    :param str username: The username of the user. If not specified, it's automatically generated.
    """
    __args__ = dict()
    __args__['group'] = group
    __args__['name'] = name
    __args__['serviceAccountId'] = service_account_id
    __args__['username'] = username
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getGroupServiceAccount:getGroupServiceAccount', __args__, opts=opts, typ=GetGroupServiceAccountResult)
    return __ret__.apply(lambda __response__: GetGroupServiceAccountResult(
        group=pulumi.get(__response__, 'group'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        service_account_id=pulumi.get(__response__, 'service_account_id'),
        username=pulumi.get(__response__, 'username')))
