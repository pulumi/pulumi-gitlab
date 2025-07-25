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
    'GetApplicationResult',
    'AwaitableGetApplicationResult',
    'get_application',
    'get_application_output',
]

@pulumi.output_type
class GetApplicationResult:
    """
    A collection of values returned by getApplication.
    """
    def __init__(__self__, application_id=None, confidential=None, id=None, name=None, redirect_url=None):
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if confidential and not isinstance(confidential, bool):
            raise TypeError("Expected argument 'confidential' to be a bool")
        pulumi.set(__self__, "confidential", confidential)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if redirect_url and not isinstance(redirect_url, str):
            raise TypeError("Expected argument 'redirect_url' to be a str")
        pulumi.set(__self__, "redirect_url", redirect_url)

    @_builtins.property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> _builtins.str:
        """
        Internal GitLab application id.
        """
        return pulumi.get(self, "application_id")

    @_builtins.property
    @pulumi.getter
    def confidential(self) -> _builtins.bool:
        """
        Indicates if the application is kept confidential.
        """
        return pulumi.get(self, "confidential")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def name(self) -> _builtins.str:
        """
        The name of the GitLab application.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="redirectUrl")
    def redirect_url(self) -> _builtins.str:
        """
        The redirect url of the application.
        """
        return pulumi.get(self, "redirect_url")


class AwaitableGetApplicationResult(GetApplicationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationResult(
            application_id=self.application_id,
            confidential=self.confidential,
            id=self.id,
            name=self.name,
            redirect_url=self.redirect_url)


def get_application(id: Optional[_builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationResult:
    """
    The `Application` data source retrieves information about a gitlab application.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/applications/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    oidc = gitlab.get_application(id="1")
    ```
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getApplication:getApplication', __args__, opts=opts, typ=GetApplicationResult).value

    return AwaitableGetApplicationResult(
        application_id=pulumi.get(__ret__, 'application_id'),
        confidential=pulumi.get(__ret__, 'confidential'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        redirect_url=pulumi.get(__ret__, 'redirect_url'))
def get_application_output(id: Optional[pulumi.Input[_builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApplicationResult]:
    """
    The `Application` data source retrieves information about a gitlab application.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/applications/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    oidc = gitlab.get_application(id="1")
    ```
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getApplication:getApplication', __args__, opts=opts, typ=GetApplicationResult)
    return __ret__.apply(lambda __response__: GetApplicationResult(
        application_id=pulumi.get(__response__, 'application_id'),
        confidential=pulumi.get(__response__, 'confidential'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        redirect_url=pulumi.get(__response__, 'redirect_url')))
