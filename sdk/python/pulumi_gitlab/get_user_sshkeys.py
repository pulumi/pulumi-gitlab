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
    'GetUserSshkeysResult',
    'AwaitableGetUserSshkeysResult',
    'get_user_sshkeys',
    'get_user_sshkeys_output',
]

@pulumi.output_type
class GetUserSshkeysResult:
    """
    A collection of values returned by getUserSshkeys.
    """
    def __init__(__self__, id=None, keys=None, user_id=None, username=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
        if user_id and not isinstance(user_id, int):
            raise TypeError("Expected argument 'user_id' to be a int")
        pulumi.set(__self__, "user_id", user_id)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def keys(self) -> Sequence['outputs.GetUserSshkeysKeyResult']:
        """
        The user's keys.
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> int:
        """
        ID of the user to get the SSH keys for.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        Username of the user to get the SSH keys for.
        """
        return pulumi.get(self, "username")


class AwaitableGetUserSshkeysResult(GetUserSshkeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserSshkeysResult(
            id=self.id,
            keys=self.keys,
            user_id=self.user_id,
            username=self.username)


def get_user_sshkeys(user_id: Optional[int] = None,
                     username: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserSshkeysResult:
    """
    The `get_user_sshkeys` data source allows a list of SSH keys to be retrieved by either the user ID or username.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#list-ssh-keys-for-user)


    :param str username: Username of the user to get the SSH keys for.
    """
    __args__ = dict()
    __args__['userId'] = user_id
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getUserSshkeys:getUserSshkeys', __args__, opts=opts, typ=GetUserSshkeysResult).value

    return AwaitableGetUserSshkeysResult(
        id=pulumi.get(__ret__, 'id'),
        keys=pulumi.get(__ret__, 'keys'),
        user_id=pulumi.get(__ret__, 'user_id'),
        username=pulumi.get(__ret__, 'username'))


@_utilities.lift_output_func(get_user_sshkeys)
def get_user_sshkeys_output(user_id: Optional[pulumi.Input[Optional[int]]] = None,
                            username: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserSshkeysResult]:
    """
    The `get_user_sshkeys` data source allows a list of SSH keys to be retrieved by either the user ID or username.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#list-ssh-keys-for-user)


    :param str username: Username of the user to get the SSH keys for.
    """
    ...
