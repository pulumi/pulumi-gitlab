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

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> str:
        """
        Internal GitLab application id.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def confidential(self) -> bool:
        """
        Indicates if the application is kept confidential.
        """
        return pulumi.get(self, "confidential")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the GitLab application.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="redirectUrl")
    def redirect_url(self) -> str:
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


def get_application(id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationResult:
    """
    The `Application` data source retrieves information about a gitlab application.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/applications.html)
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
def get_application_output(id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApplicationResult]:
    """
    The `Application` data source retrieves information about a gitlab application.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/applications.html)
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
