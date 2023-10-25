# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetReleaseLinkResult',
    'AwaitableGetReleaseLinkResult',
    'get_release_link',
    'get_release_link_output',
]

@pulumi.output_type
class GetReleaseLinkResult:
    """
    A collection of values returned by getReleaseLink.
    """
    def __init__(__self__, direct_asset_url=None, external=None, filepath=None, id=None, link_id=None, link_type=None, name=None, project=None, tag_name=None, url=None):
        if direct_asset_url and not isinstance(direct_asset_url, str):
            raise TypeError("Expected argument 'direct_asset_url' to be a str")
        pulumi.set(__self__, "direct_asset_url", direct_asset_url)
        if external and not isinstance(external, bool):
            raise TypeError("Expected argument 'external' to be a bool")
        pulumi.set(__self__, "external", external)
        if filepath and not isinstance(filepath, str):
            raise TypeError("Expected argument 'filepath' to be a str")
        pulumi.set(__self__, "filepath", filepath)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if link_id and not isinstance(link_id, int):
            raise TypeError("Expected argument 'link_id' to be a int")
        pulumi.set(__self__, "link_id", link_id)
        if link_type and not isinstance(link_type, str):
            raise TypeError("Expected argument 'link_type' to be a str")
        pulumi.set(__self__, "link_type", link_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if tag_name and not isinstance(tag_name, str):
            raise TypeError("Expected argument 'tag_name' to be a str")
        pulumi.set(__self__, "tag_name", tag_name)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="directAssetUrl")
    def direct_asset_url(self) -> str:
        """
        Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        """
        return pulumi.get(self, "direct_asset_url")

    @property
    @pulumi.getter
    def external(self) -> bool:
        """
        External or internal link.
        """
        return pulumi.get(self, "external")

    @property
    @pulumi.getter
    def filepath(self) -> str:
        """
        Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        """
        return pulumi.get(self, "filepath")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="linkId")
    def link_id(self) -> int:
        """
        The ID of the link.
        """
        return pulumi.get(self, "link_id")

    @property
    @pulumi.getter(name="linkType")
    def link_type(self) -> str:
        """
        The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
        """
        return pulumi.get(self, "link_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the link. Link names must be unique within the release.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="tagName")
    def tag_name(self) -> str:
        """
        The tag associated with the Release.
        """
        return pulumi.get(self, "tag_name")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The URL of the link. Link URLs must be unique within the release.
        """
        return pulumi.get(self, "url")


class AwaitableGetReleaseLinkResult(GetReleaseLinkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReleaseLinkResult(
            direct_asset_url=self.direct_asset_url,
            external=self.external,
            filepath=self.filepath,
            id=self.id,
            link_id=self.link_id,
            link_type=self.link_type,
            name=self.name,
            project=self.project,
            tag_name=self.tag_name,
            url=self.url)


def get_release_link(link_id: Optional[int] = None,
                     project: Optional[str] = None,
                     tag_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReleaseLinkResult:
    """
    The `ReleaseLink` data source allows get details of a release link.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_release_link(link_id=11,
        project="foo/bar",
        tag_name="v1.0.1")
    ```


    :param int link_id: The ID of the link.
    :param str project: The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
    :param str tag_name: The tag associated with the Release.
    """
    __args__ = dict()
    __args__['linkId'] = link_id
    __args__['project'] = project
    __args__['tagName'] = tag_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getReleaseLink:getReleaseLink', __args__, opts=opts, typ=GetReleaseLinkResult).value

    return AwaitableGetReleaseLinkResult(
        direct_asset_url=pulumi.get(__ret__, 'direct_asset_url'),
        external=pulumi.get(__ret__, 'external'),
        filepath=pulumi.get(__ret__, 'filepath'),
        id=pulumi.get(__ret__, 'id'),
        link_id=pulumi.get(__ret__, 'link_id'),
        link_type=pulumi.get(__ret__, 'link_type'),
        name=pulumi.get(__ret__, 'name'),
        project=pulumi.get(__ret__, 'project'),
        tag_name=pulumi.get(__ret__, 'tag_name'),
        url=pulumi.get(__ret__, 'url'))


@_utilities.lift_output_func(get_release_link)
def get_release_link_output(link_id: Optional[pulumi.Input[int]] = None,
                            project: Optional[pulumi.Input[str]] = None,
                            tag_name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReleaseLinkResult]:
    """
    The `ReleaseLink` data source allows get details of a release link.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_release_link(link_id=11,
        project="foo/bar",
        tag_name="v1.0.1")
    ```


    :param int link_id: The ID of the link.
    :param str project: The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
    :param str tag_name: The tag associated with the Release.
    """
    ...
