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
from . import outputs
from ._inputs import *

__all__ = [
    'GetReleaseResult',
    'AwaitableGetReleaseResult',
    'get_release',
    'get_release_output',
]

@pulumi.output_type
class GetReleaseResult:
    """
    A collection of values returned by getRelease.
    """
    def __init__(__self__, assets=None, created_at=None, description=None, id=None, name=None, project_id=None, released_at=None, tag_name=None):
        if assets and not isinstance(assets, dict):
            raise TypeError("Expected argument 'assets' to be a dict")
        pulumi.set(__self__, "assets", assets)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if released_at and not isinstance(released_at, str):
            raise TypeError("Expected argument 'released_at' to be a str")
        pulumi.set(__self__, "released_at", released_at)
        if tag_name and not isinstance(tag_name, str):
            raise TypeError("Expected argument 'tag_name' to be a str")
        pulumi.set(__self__, "tag_name", tag_name)

    @property
    @pulumi.getter
    def assets(self) -> Optional['outputs.GetReleaseAssetsResult']:
        """
        The assets for a release
        """
        return pulumi.get(self, "assets")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date the release was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An HTML rendered description of the release.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the release.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The ID or URL-encoded path of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="releasedAt")
    def released_at(self) -> str:
        """
        The date the release was created.
        """
        return pulumi.get(self, "released_at")

    @property
    @pulumi.getter(name="tagName")
    def tag_name(self) -> str:
        """
        The Git tag the release is associated with.
        """
        return pulumi.get(self, "tag_name")


class AwaitableGetReleaseResult(GetReleaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReleaseResult(
            assets=self.assets,
            created_at=self.created_at,
            description=self.description,
            id=self.id,
            name=self.name,
            project_id=self.project_id,
            released_at=self.released_at,
            tag_name=self.tag_name)


def get_release(assets: Optional[Union['GetReleaseAssetsArgs', 'GetReleaseAssetsArgsDict']] = None,
                project_id: Optional[str] = None,
                tag_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReleaseResult:
    """
    The `Release` data source retrieves information about a gitlab release for a project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/releases/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    # By project ID and tag_name
    example = gitlab.get_release(project_id="1234",
        tag_name="v1.0")
    ```


    :param Union['GetReleaseAssetsArgs', 'GetReleaseAssetsArgsDict'] assets: The assets for a release
    :param str project_id: The ID or URL-encoded path of the project.
    :param str tag_name: The Git tag the release is associated with.
    """
    __args__ = dict()
    __args__['assets'] = assets
    __args__['projectId'] = project_id
    __args__['tagName'] = tag_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getRelease:getRelease', __args__, opts=opts, typ=GetReleaseResult).value

    return AwaitableGetReleaseResult(
        assets=pulumi.get(__ret__, 'assets'),
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        released_at=pulumi.get(__ret__, 'released_at'),
        tag_name=pulumi.get(__ret__, 'tag_name'))
def get_release_output(assets: Optional[pulumi.Input[Optional[Union['GetReleaseAssetsArgs', 'GetReleaseAssetsArgsDict']]]] = None,
                       project_id: Optional[pulumi.Input[str]] = None,
                       tag_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetReleaseResult]:
    """
    The `Release` data source retrieves information about a gitlab release for a project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/releases/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    # By project ID and tag_name
    example = gitlab.get_release(project_id="1234",
        tag_name="v1.0")
    ```


    :param Union['GetReleaseAssetsArgs', 'GetReleaseAssetsArgsDict'] assets: The assets for a release
    :param str project_id: The ID or URL-encoded path of the project.
    :param str tag_name: The Git tag the release is associated with.
    """
    __args__ = dict()
    __args__['assets'] = assets
    __args__['projectId'] = project_id
    __args__['tagName'] = tag_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getRelease:getRelease', __args__, opts=opts, typ=GetReleaseResult)
    return __ret__.apply(lambda __response__: GetReleaseResult(
        assets=pulumi.get(__response__, 'assets'),
        created_at=pulumi.get(__response__, 'created_at'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id'),
        released_at=pulumi.get(__response__, 'released_at'),
        tag_name=pulumi.get(__response__, 'tag_name')))
