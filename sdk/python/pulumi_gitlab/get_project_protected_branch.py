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
    'GetProjectProtectedBranchResult',
    'AwaitableGetProjectProtectedBranchResult',
    'get_project_protected_branch',
    'get_project_protected_branch_output',
]

@pulumi.output_type
class GetProjectProtectedBranchResult:
    """
    A collection of values returned by getProjectProtectedBranch.
    """
    def __init__(__self__, allow_force_push=None, code_owner_approval_required=None, id=None, merge_access_levels=None, name=None, project_id=None, push_access_levels=None):
        if allow_force_push and not isinstance(allow_force_push, bool):
            raise TypeError("Expected argument 'allow_force_push' to be a bool")
        pulumi.set(__self__, "allow_force_push", allow_force_push)
        if code_owner_approval_required and not isinstance(code_owner_approval_required, bool):
            raise TypeError("Expected argument 'code_owner_approval_required' to be a bool")
        pulumi.set(__self__, "code_owner_approval_required", code_owner_approval_required)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if merge_access_levels and not isinstance(merge_access_levels, list):
            raise TypeError("Expected argument 'merge_access_levels' to be a list")
        pulumi.set(__self__, "merge_access_levels", merge_access_levels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if push_access_levels and not isinstance(push_access_levels, list):
            raise TypeError("Expected argument 'push_access_levels' to be a list")
        pulumi.set(__self__, "push_access_levels", push_access_levels)

    @property
    @pulumi.getter(name="allowForcePush")
    def allow_force_push(self) -> bool:
        """
        Whether force push is allowed.
        """
        return pulumi.get(self, "allow_force_push")

    @property
    @pulumi.getter(name="codeOwnerApprovalRequired")
    def code_owner_approval_required(self) -> bool:
        """
        Reject code pushes that change files listed in the CODEOWNERS file.
        """
        return pulumi.get(self, "code_owner_approval_required")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="mergeAccessLevels")
    def merge_access_levels(self) -> Optional[Sequence['outputs.GetProjectProtectedBranchMergeAccessLevelResult']]:
        """
        Array of access levels and user(s)/group(s) allowed to merge to protected branch.
        """
        return pulumi.get(self, "merge_access_levels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the protected branch.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The integer or path with namespace that uniquely identifies the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="pushAccessLevels")
    def push_access_levels(self) -> Optional[Sequence['outputs.GetProjectProtectedBranchPushAccessLevelResult']]:
        """
        Array of access levels and user(s)/group(s) allowed to push to protected branch.
        """
        return pulumi.get(self, "push_access_levels")


class AwaitableGetProjectProtectedBranchResult(GetProjectProtectedBranchResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectProtectedBranchResult(
            allow_force_push=self.allow_force_push,
            code_owner_approval_required=self.code_owner_approval_required,
            id=self.id,
            merge_access_levels=self.merge_access_levels,
            name=self.name,
            project_id=self.project_id,
            push_access_levels=self.push_access_levels)


def get_project_protected_branch(merge_access_levels: Optional[Sequence[Union['GetProjectProtectedBranchMergeAccessLevelArgs', 'GetProjectProtectedBranchMergeAccessLevelArgsDict']]] = None,
                                 name: Optional[str] = None,
                                 project_id: Optional[str] = None,
                                 push_access_levels: Optional[Sequence[Union['GetProjectProtectedBranchPushAccessLevelArgs', 'GetProjectProtectedBranchPushAccessLevelArgsDict']]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectProtectedBranchResult:
    """
    The `get_project_protected_branch` data source allows details of a protected branch to be retrieved by its name and the project it belongs to.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_branches.html#get-a-single-protected-branch-or-wildcard-protected-branch)


    :param Sequence[Union['GetProjectProtectedBranchMergeAccessLevelArgs', 'GetProjectProtectedBranchMergeAccessLevelArgsDict']] merge_access_levels: Array of access levels and user(s)/group(s) allowed to merge to protected branch.
    :param str name: The name of the protected branch.
    :param str project_id: The integer or path with namespace that uniquely identifies the project.
    :param Sequence[Union['GetProjectProtectedBranchPushAccessLevelArgs', 'GetProjectProtectedBranchPushAccessLevelArgsDict']] push_access_levels: Array of access levels and user(s)/group(s) allowed to push to protected branch.
    """
    __args__ = dict()
    __args__['mergeAccessLevels'] = merge_access_levels
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['pushAccessLevels'] = push_access_levels
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjectProtectedBranch:getProjectProtectedBranch', __args__, opts=opts, typ=GetProjectProtectedBranchResult).value

    return AwaitableGetProjectProtectedBranchResult(
        allow_force_push=pulumi.get(__ret__, 'allow_force_push'),
        code_owner_approval_required=pulumi.get(__ret__, 'code_owner_approval_required'),
        id=pulumi.get(__ret__, 'id'),
        merge_access_levels=pulumi.get(__ret__, 'merge_access_levels'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        push_access_levels=pulumi.get(__ret__, 'push_access_levels'))
def get_project_protected_branch_output(merge_access_levels: Optional[pulumi.Input[Optional[Sequence[Union['GetProjectProtectedBranchMergeAccessLevelArgs', 'GetProjectProtectedBranchMergeAccessLevelArgsDict']]]]] = None,
                                        name: Optional[pulumi.Input[str]] = None,
                                        project_id: Optional[pulumi.Input[str]] = None,
                                        push_access_levels: Optional[pulumi.Input[Optional[Sequence[Union['GetProjectProtectedBranchPushAccessLevelArgs', 'GetProjectProtectedBranchPushAccessLevelArgsDict']]]]] = None,
                                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetProjectProtectedBranchResult]:
    """
    The `get_project_protected_branch` data source allows details of a protected branch to be retrieved by its name and the project it belongs to.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_branches.html#get-a-single-protected-branch-or-wildcard-protected-branch)


    :param Sequence[Union['GetProjectProtectedBranchMergeAccessLevelArgs', 'GetProjectProtectedBranchMergeAccessLevelArgsDict']] merge_access_levels: Array of access levels and user(s)/group(s) allowed to merge to protected branch.
    :param str name: The name of the protected branch.
    :param str project_id: The integer or path with namespace that uniquely identifies the project.
    :param Sequence[Union['GetProjectProtectedBranchPushAccessLevelArgs', 'GetProjectProtectedBranchPushAccessLevelArgsDict']] push_access_levels: Array of access levels and user(s)/group(s) allowed to push to protected branch.
    """
    __args__ = dict()
    __args__['mergeAccessLevels'] = merge_access_levels
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['pushAccessLevels'] = push_access_levels
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getProjectProtectedBranch:getProjectProtectedBranch', __args__, opts=opts, typ=GetProjectProtectedBranchResult)
    return __ret__.apply(lambda __response__: GetProjectProtectedBranchResult(
        allow_force_push=pulumi.get(__response__, 'allow_force_push'),
        code_owner_approval_required=pulumi.get(__response__, 'code_owner_approval_required'),
        id=pulumi.get(__response__, 'id'),
        merge_access_levels=pulumi.get(__response__, 'merge_access_levels'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id'),
        push_access_levels=pulumi.get(__response__, 'push_access_levels')))
