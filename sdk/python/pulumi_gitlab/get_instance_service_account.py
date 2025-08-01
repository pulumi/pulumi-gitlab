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

__all__ = [
    'GetInstanceServiceAccountResult',
    'AwaitableGetInstanceServiceAccountResult',
    'get_instance_service_account',
    'get_instance_service_account_output',
]

@pulumi.output_type
class GetInstanceServiceAccountResult:
    """
    A collection of values returned by getInstanceServiceAccount.
    """
    def __init__(__self__, email=None, id=None, name=None, service_account_id=None, username=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
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

    @_builtins.property
    @pulumi.getter
    def email(self) -> _builtins.str:
        """
        The email of the user.
        """
        return pulumi.get(self, "email")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def name(self) -> _builtins.str:
        """
        The name of the user.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> _builtins.str:
        """
        The service account id.
        """
        return pulumi.get(self, "service_account_id")

    @_builtins.property
    @pulumi.getter
    def username(self) -> _builtins.str:
        """
        The username of the user.
        """
        return pulumi.get(self, "username")


class AwaitableGetInstanceServiceAccountResult(GetInstanceServiceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceServiceAccountResult(
            email=self.email,
            id=self.id,
            name=self.name,
            service_account_id=self.service_account_id,
            username=self.username)


def get_instance_service_account(service_account_id: Optional[_builtins.str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceServiceAccountResult:
    """
    The `InstanceServiceAccount` data source retrieves information about a gitlab service account.

    > In order for a user to create a user account, they must have admin privileges at the instance level. This makes this feature unavailable on `gitlab.com`

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/user_service_accounts/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_instance_service_account(service_account_id="123")
    ```


    :param _builtins.str service_account_id: The service account id.
    """
    __args__ = dict()
    __args__['serviceAccountId'] = service_account_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getInstanceServiceAccount:getInstanceServiceAccount', __args__, opts=opts, typ=GetInstanceServiceAccountResult).value

    return AwaitableGetInstanceServiceAccountResult(
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        service_account_id=pulumi.get(__ret__, 'service_account_id'),
        username=pulumi.get(__ret__, 'username'))
def get_instance_service_account_output(service_account_id: Optional[pulumi.Input[_builtins.str]] = None,
                                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetInstanceServiceAccountResult]:
    """
    The `InstanceServiceAccount` data source retrieves information about a gitlab service account.

    > In order for a user to create a user account, they must have admin privileges at the instance level. This makes this feature unavailable on `gitlab.com`

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/user_service_accounts/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_instance_service_account(service_account_id="123")
    ```


    :param _builtins.str service_account_id: The service account id.
    """
    __args__ = dict()
    __args__['serviceAccountId'] = service_account_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getInstanceServiceAccount:getInstanceServiceAccount', __args__, opts=opts, typ=GetInstanceServiceAccountResult)
    return __ret__.apply(lambda __response__: GetInstanceServiceAccountResult(
        email=pulumi.get(__response__, 'email'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        service_account_id=pulumi.get(__response__, 'service_account_id'),
        username=pulumi.get(__response__, 'username')))
