# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'GetBranchResult',
    'AwaitableGetBranchResult',
    'get_branch',
    'get_branch_output',
]

@pulumi.output_type
class GetBranchResult:
    """
    A collection of values returned by getBranch.
    """
    def __init__(__self__, can_push=None, commits=None, default=None, developer_can_merge=None, developer_can_push=None, id=None, merged=None, name=None, project=None, protected=None, web_url=None):
        if can_push and not isinstance(can_push, bool):
            raise TypeError("Expected argument 'can_push' to be a bool")
        pulumi.set(__self__, "can_push", can_push)
        if commits and not isinstance(commits, list):
            raise TypeError("Expected argument 'commits' to be a list")
        pulumi.set(__self__, "commits", commits)
        if default and not isinstance(default, bool):
            raise TypeError("Expected argument 'default' to be a bool")
        pulumi.set(__self__, "default", default)
        if developer_can_merge and not isinstance(developer_can_merge, bool):
            raise TypeError("Expected argument 'developer_can_merge' to be a bool")
        pulumi.set(__self__, "developer_can_merge", developer_can_merge)
        if developer_can_push and not isinstance(developer_can_push, bool):
            raise TypeError("Expected argument 'developer_can_push' to be a bool")
        pulumi.set(__self__, "developer_can_push", developer_can_push)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if merged and not isinstance(merged, bool):
            raise TypeError("Expected argument 'merged' to be a bool")
        pulumi.set(__self__, "merged", merged)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if protected and not isinstance(protected, bool):
            raise TypeError("Expected argument 'protected' to be a bool")
        pulumi.set(__self__, "protected", protected)
        if web_url and not isinstance(web_url, str):
            raise TypeError("Expected argument 'web_url' to be a str")
        pulumi.set(__self__, "web_url", web_url)

    @property
    @pulumi.getter(name="canPush")
    def can_push(self) -> builtins.bool:
        """
        Bool, true if you can push to the branch.
        """
        return pulumi.get(self, "can_push")

    @property
    @pulumi.getter
    def commits(self) -> Sequence['outputs.GetBranchCommitResult']:
        """
        The commit associated with the branch ref.
        """
        return pulumi.get(self, "commits")

    @property
    @pulumi.getter
    def default(self) -> builtins.bool:
        """
        Bool, true if branch is the default branch for the project.
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter(name="developerCanMerge")
    def developer_can_merge(self) -> builtins.bool:
        """
        Bool, true if developer level access allows to merge branch.
        """
        return pulumi.get(self, "developer_can_merge")

    @property
    @pulumi.getter(name="developerCanPush")
    def developer_can_push(self) -> builtins.bool:
        """
        Bool, true if developer level access allows git push.
        """
        return pulumi.get(self, "developer_can_push")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def merged(self) -> builtins.bool:
        """
        Bool, true if the branch has been merged into it's parent.
        """
        return pulumi.get(self, "merged")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the branch.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> builtins.str:
        """
        The full path or id of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def protected(self) -> builtins.bool:
        """
        Bool, true if branch has branch protection.
        """
        return pulumi.get(self, "protected")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> builtins.str:
        """
        The url of the created branch (https.)
        """
        return pulumi.get(self, "web_url")


class AwaitableGetBranchResult(GetBranchResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBranchResult(
            can_push=self.can_push,
            commits=self.commits,
            default=self.default,
            developer_can_merge=self.developer_can_merge,
            developer_can_push=self.developer_can_push,
            id=self.id,
            merged=self.merged,
            name=self.name,
            project=self.project,
            protected=self.protected,
            web_url=self.web_url)


def get_branch(name: Optional[builtins.str] = None,
               project: Optional[builtins.str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBranchResult:
    """
    The `Branch` data source allows details of a repository branch to be retrieved by its name and project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/branches/#get-single-repository-branch)


    :param builtins.str name: The name of the branch.
    :param builtins.str project: The full path or id of the project.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getBranch:getBranch', __args__, opts=opts, typ=GetBranchResult).value

    return AwaitableGetBranchResult(
        can_push=pulumi.get(__ret__, 'can_push'),
        commits=pulumi.get(__ret__, 'commits'),
        default=pulumi.get(__ret__, 'default'),
        developer_can_merge=pulumi.get(__ret__, 'developer_can_merge'),
        developer_can_push=pulumi.get(__ret__, 'developer_can_push'),
        id=pulumi.get(__ret__, 'id'),
        merged=pulumi.get(__ret__, 'merged'),
        name=pulumi.get(__ret__, 'name'),
        project=pulumi.get(__ret__, 'project'),
        protected=pulumi.get(__ret__, 'protected'),
        web_url=pulumi.get(__ret__, 'web_url'))
def get_branch_output(name: Optional[pulumi.Input[builtins.str]] = None,
                      project: Optional[pulumi.Input[builtins.str]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetBranchResult]:
    """
    The `Branch` data source allows details of a repository branch to be retrieved by its name and project.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/branches/#get-single-repository-branch)


    :param builtins.str name: The name of the branch.
    :param builtins.str project: The full path or id of the project.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getBranch:getBranch', __args__, opts=opts, typ=GetBranchResult)
    return __ret__.apply(lambda __response__: GetBranchResult(
        can_push=pulumi.get(__response__, 'can_push'),
        commits=pulumi.get(__response__, 'commits'),
        default=pulumi.get(__response__, 'default'),
        developer_can_merge=pulumi.get(__response__, 'developer_can_merge'),
        developer_can_push=pulumi.get(__response__, 'developer_can_push'),
        id=pulumi.get(__response__, 'id'),
        merged=pulumi.get(__response__, 'merged'),
        name=pulumi.get(__response__, 'name'),
        project=pulumi.get(__response__, 'project'),
        protected=pulumi.get(__response__, 'protected'),
        web_url=pulumi.get(__response__, 'web_url')))
