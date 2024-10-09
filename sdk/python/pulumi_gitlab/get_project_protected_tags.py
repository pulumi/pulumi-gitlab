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

__all__ = [
    'GetProjectProtectedTagsResult',
    'AwaitableGetProjectProtectedTagsResult',
    'get_project_protected_tags',
    'get_project_protected_tags_output',
]

@pulumi.output_type
class GetProjectProtectedTagsResult:
    """
    A collection of values returned by getProjectProtectedTags.
    """
    def __init__(__self__, id=None, project=None, protected_tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if protected_tags and not isinstance(protected_tags, list):
            raise TypeError("Expected argument 'protected_tags' to be a list")
        pulumi.set(__self__, "protected_tags", protected_tags)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The integer or path with namespace that uniquely identifies the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="protectedTags")
    def protected_tags(self) -> Sequence['outputs.GetProjectProtectedTagsProtectedTagResult']:
        """
        A list of protected tags, as defined below.
        """
        return pulumi.get(self, "protected_tags")


class AwaitableGetProjectProtectedTagsResult(GetProjectProtectedTagsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectProtectedTagsResult(
            id=self.id,
            project=self.project,
            protected_tags=self.protected_tags)


def get_project_protected_tags(project: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectProtectedTagsResult:
    """
    The `get_project_protected_tags` data source allows details of the protected tags of a given project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_tags.html#list-protected-tags)


    :param str project: The integer or path with namespace that uniquely identifies the project.
    """
    __args__ = dict()
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjectProtectedTags:getProjectProtectedTags', __args__, opts=opts, typ=GetProjectProtectedTagsResult).value

    return AwaitableGetProjectProtectedTagsResult(
        id=pulumi.get(__ret__, 'id'),
        project=pulumi.get(__ret__, 'project'),
        protected_tags=pulumi.get(__ret__, 'protected_tags'))
def get_project_protected_tags_output(project: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectProtectedTagsResult]:
    """
    The `get_project_protected_tags` data source allows details of the protected tags of a given project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_tags.html#list-protected-tags)


    :param str project: The integer or path with namespace that uniquely identifies the project.
    """
    __args__ = dict()
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getProjectProtectedTags:getProjectProtectedTags', __args__, opts=opts, typ=GetProjectProtectedTagsResult)
    return __ret__.apply(lambda __response__: GetProjectProtectedTagsResult(
        id=pulumi.get(__response__, 'id'),
        project=pulumi.get(__response__, 'project'),
        protected_tags=pulumi.get(__response__, 'protected_tags')))
