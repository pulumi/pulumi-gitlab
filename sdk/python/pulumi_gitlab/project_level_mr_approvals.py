# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectLevelMrApprovalsArgs', 'ProjectLevelMrApprovals']

@pulumi.input_type
class ProjectLevelMrApprovalsArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 disable_overriding_approvers_per_merge_request: Optional[pulumi.Input[bool]] = None,
                 merge_requests_author_approval: Optional[pulumi.Input[bool]] = None,
                 merge_requests_disable_committers_approval: Optional[pulumi.Input[bool]] = None,
                 require_password_to_approve: Optional[pulumi.Input[bool]] = None,
                 reset_approvals_on_push: Optional[pulumi.Input[bool]] = None,
                 selective_code_owner_removals: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ProjectLevelMrApprovals resource.
        :param pulumi.Input[str] project: The ID or URL-encoded path of a project to change MR approval configuration.
        :param pulumi.Input[bool] disable_overriding_approvers_per_merge_request: Set to `true` to disable overriding approvers per merge request.
        :param pulumi.Input[bool] merge_requests_author_approval: Set to `true` to allow merge requests authors to approve their own merge requests.
        :param pulumi.Input[bool] merge_requests_disable_committers_approval: Set to `true` to disable merge request committers from approving their own merge requests.
        :param pulumi.Input[bool] require_password_to_approve: Set to `true` to require authentication to approve merge requests.
        :param pulumi.Input[bool] reset_approvals_on_push: Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        :param pulumi.Input[bool] selective_code_owner_removals: Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
        """
        pulumi.set(__self__, "project", project)
        if disable_overriding_approvers_per_merge_request is not None:
            pulumi.set(__self__, "disable_overriding_approvers_per_merge_request", disable_overriding_approvers_per_merge_request)
        if merge_requests_author_approval is not None:
            pulumi.set(__self__, "merge_requests_author_approval", merge_requests_author_approval)
        if merge_requests_disable_committers_approval is not None:
            pulumi.set(__self__, "merge_requests_disable_committers_approval", merge_requests_disable_committers_approval)
        if require_password_to_approve is not None:
            pulumi.set(__self__, "require_password_to_approve", require_password_to_approve)
        if reset_approvals_on_push is not None:
            pulumi.set(__self__, "reset_approvals_on_push", reset_approvals_on_push)
        if selective_code_owner_removals is not None:
            pulumi.set(__self__, "selective_code_owner_removals", selective_code_owner_removals)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or URL-encoded path of a project to change MR approval configuration.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="disableOverridingApproversPerMergeRequest")
    def disable_overriding_approvers_per_merge_request(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to disable overriding approvers per merge request.
        """
        return pulumi.get(self, "disable_overriding_approvers_per_merge_request")

    @disable_overriding_approvers_per_merge_request.setter
    def disable_overriding_approvers_per_merge_request(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_overriding_approvers_per_merge_request", value)

    @property
    @pulumi.getter(name="mergeRequestsAuthorApproval")
    def merge_requests_author_approval(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to allow merge requests authors to approve their own merge requests.
        """
        return pulumi.get(self, "merge_requests_author_approval")

    @merge_requests_author_approval.setter
    def merge_requests_author_approval(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "merge_requests_author_approval", value)

    @property
    @pulumi.getter(name="mergeRequestsDisableCommittersApproval")
    def merge_requests_disable_committers_approval(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to disable merge request committers from approving their own merge requests.
        """
        return pulumi.get(self, "merge_requests_disable_committers_approval")

    @merge_requests_disable_committers_approval.setter
    def merge_requests_disable_committers_approval(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "merge_requests_disable_committers_approval", value)

    @property
    @pulumi.getter(name="requirePasswordToApprove")
    def require_password_to_approve(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to require authentication to approve merge requests.
        """
        return pulumi.get(self, "require_password_to_approve")

    @require_password_to_approve.setter
    def require_password_to_approve(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_password_to_approve", value)

    @property
    @pulumi.getter(name="resetApprovalsOnPush")
    def reset_approvals_on_push(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        """
        return pulumi.get(self, "reset_approvals_on_push")

    @reset_approvals_on_push.setter
    def reset_approvals_on_push(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reset_approvals_on_push", value)

    @property
    @pulumi.getter(name="selectiveCodeOwnerRemovals")
    def selective_code_owner_removals(self) -> Optional[pulumi.Input[bool]]:
        """
        Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
        """
        return pulumi.get(self, "selective_code_owner_removals")

    @selective_code_owner_removals.setter
    def selective_code_owner_removals(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "selective_code_owner_removals", value)


@pulumi.input_type
class _ProjectLevelMrApprovalsState:
    def __init__(__self__, *,
                 disable_overriding_approvers_per_merge_request: Optional[pulumi.Input[bool]] = None,
                 merge_requests_author_approval: Optional[pulumi.Input[bool]] = None,
                 merge_requests_disable_committers_approval: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 require_password_to_approve: Optional[pulumi.Input[bool]] = None,
                 reset_approvals_on_push: Optional[pulumi.Input[bool]] = None,
                 selective_code_owner_removals: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering ProjectLevelMrApprovals resources.
        :param pulumi.Input[bool] disable_overriding_approvers_per_merge_request: Set to `true` to disable overriding approvers per merge request.
        :param pulumi.Input[bool] merge_requests_author_approval: Set to `true` to allow merge requests authors to approve their own merge requests.
        :param pulumi.Input[bool] merge_requests_disable_committers_approval: Set to `true` to disable merge request committers from approving their own merge requests.
        :param pulumi.Input[str] project: The ID or URL-encoded path of a project to change MR approval configuration.
        :param pulumi.Input[bool] require_password_to_approve: Set to `true` to require authentication to approve merge requests.
        :param pulumi.Input[bool] reset_approvals_on_push: Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        :param pulumi.Input[bool] selective_code_owner_removals: Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
        """
        if disable_overriding_approvers_per_merge_request is not None:
            pulumi.set(__self__, "disable_overriding_approvers_per_merge_request", disable_overriding_approvers_per_merge_request)
        if merge_requests_author_approval is not None:
            pulumi.set(__self__, "merge_requests_author_approval", merge_requests_author_approval)
        if merge_requests_disable_committers_approval is not None:
            pulumi.set(__self__, "merge_requests_disable_committers_approval", merge_requests_disable_committers_approval)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if require_password_to_approve is not None:
            pulumi.set(__self__, "require_password_to_approve", require_password_to_approve)
        if reset_approvals_on_push is not None:
            pulumi.set(__self__, "reset_approvals_on_push", reset_approvals_on_push)
        if selective_code_owner_removals is not None:
            pulumi.set(__self__, "selective_code_owner_removals", selective_code_owner_removals)

    @property
    @pulumi.getter(name="disableOverridingApproversPerMergeRequest")
    def disable_overriding_approvers_per_merge_request(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to disable overriding approvers per merge request.
        """
        return pulumi.get(self, "disable_overriding_approvers_per_merge_request")

    @disable_overriding_approvers_per_merge_request.setter
    def disable_overriding_approvers_per_merge_request(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_overriding_approvers_per_merge_request", value)

    @property
    @pulumi.getter(name="mergeRequestsAuthorApproval")
    def merge_requests_author_approval(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to allow merge requests authors to approve their own merge requests.
        """
        return pulumi.get(self, "merge_requests_author_approval")

    @merge_requests_author_approval.setter
    def merge_requests_author_approval(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "merge_requests_author_approval", value)

    @property
    @pulumi.getter(name="mergeRequestsDisableCommittersApproval")
    def merge_requests_disable_committers_approval(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to disable merge request committers from approving their own merge requests.
        """
        return pulumi.get(self, "merge_requests_disable_committers_approval")

    @merge_requests_disable_committers_approval.setter
    def merge_requests_disable_committers_approval(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "merge_requests_disable_committers_approval", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or URL-encoded path of a project to change MR approval configuration.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="requirePasswordToApprove")
    def require_password_to_approve(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to require authentication to approve merge requests.
        """
        return pulumi.get(self, "require_password_to_approve")

    @require_password_to_approve.setter
    def require_password_to_approve(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_password_to_approve", value)

    @property
    @pulumi.getter(name="resetApprovalsOnPush")
    def reset_approvals_on_push(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        """
        return pulumi.get(self, "reset_approvals_on_push")

    @reset_approvals_on_push.setter
    def reset_approvals_on_push(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reset_approvals_on_push", value)

    @property
    @pulumi.getter(name="selectiveCodeOwnerRemovals")
    def selective_code_owner_removals(self) -> Optional[pulumi.Input[bool]]:
        """
        Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
        """
        return pulumi.get(self, "selective_code_owner_removals")

    @selective_code_owner_removals.setter
    def selective_code_owner_removals(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "selective_code_owner_removals", value)


class ProjectLevelMrApprovals(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_overriding_approvers_per_merge_request: Optional[pulumi.Input[bool]] = None,
                 merge_requests_author_approval: Optional[pulumi.Input[bool]] = None,
                 merge_requests_disable_committers_approval: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 require_password_to_approve: Optional[pulumi.Input[bool]] = None,
                 reset_approvals_on_push: Optional[pulumi.Input[bool]] = None,
                 selective_code_owner_removals: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The `gitlab_project_level_mr_approval_rule` resource allows to manage the lifecycle of a Merge Request-level approval rule.

        > This resource requires a GitLab Enterprise instance.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/merge_request_approvals.html#merge-request-level-mr-approvals)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo_project = gitlab.Project("fooProject", description="My example project")
        foo_project_level_mr_approvals = gitlab.ProjectLevelMrApprovals("fooProjectLevelMrApprovals",
            project=foo_project.id,
            reset_approvals_on_push=True,
            disable_overriding_approvers_per_merge_request=False,
            merge_requests_author_approval=False,
            merge_requests_disable_committers_approval=True)
        ```

        ## Import

        ```sh
        $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals You can import an approval configuration state using `<resource> <project_id>`.
        ```

        # 

        For example:

        ```sh
        $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals foo 1234
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_overriding_approvers_per_merge_request: Set to `true` to disable overriding approvers per merge request.
        :param pulumi.Input[bool] merge_requests_author_approval: Set to `true` to allow merge requests authors to approve their own merge requests.
        :param pulumi.Input[bool] merge_requests_disable_committers_approval: Set to `true` to disable merge request committers from approving their own merge requests.
        :param pulumi.Input[str] project: The ID or URL-encoded path of a project to change MR approval configuration.
        :param pulumi.Input[bool] require_password_to_approve: Set to `true` to require authentication to approve merge requests.
        :param pulumi.Input[bool] reset_approvals_on_push: Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        :param pulumi.Input[bool] selective_code_owner_removals: Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectLevelMrApprovalsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `gitlab_project_level_mr_approval_rule` resource allows to manage the lifecycle of a Merge Request-level approval rule.

        > This resource requires a GitLab Enterprise instance.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/merge_request_approvals.html#merge-request-level-mr-approvals)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo_project = gitlab.Project("fooProject", description="My example project")
        foo_project_level_mr_approvals = gitlab.ProjectLevelMrApprovals("fooProjectLevelMrApprovals",
            project=foo_project.id,
            reset_approvals_on_push=True,
            disable_overriding_approvers_per_merge_request=False,
            merge_requests_author_approval=False,
            merge_requests_disable_committers_approval=True)
        ```

        ## Import

        ```sh
        $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals You can import an approval configuration state using `<resource> <project_id>`.
        ```

        # 

        For example:

        ```sh
        $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals foo 1234
        ```

        :param str resource_name: The name of the resource.
        :param ProjectLevelMrApprovalsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectLevelMrApprovalsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_overriding_approvers_per_merge_request: Optional[pulumi.Input[bool]] = None,
                 merge_requests_author_approval: Optional[pulumi.Input[bool]] = None,
                 merge_requests_disable_committers_approval: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 require_password_to_approve: Optional[pulumi.Input[bool]] = None,
                 reset_approvals_on_push: Optional[pulumi.Input[bool]] = None,
                 selective_code_owner_removals: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectLevelMrApprovalsArgs.__new__(ProjectLevelMrApprovalsArgs)

            __props__.__dict__["disable_overriding_approvers_per_merge_request"] = disable_overriding_approvers_per_merge_request
            __props__.__dict__["merge_requests_author_approval"] = merge_requests_author_approval
            __props__.__dict__["merge_requests_disable_committers_approval"] = merge_requests_disable_committers_approval
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["require_password_to_approve"] = require_password_to_approve
            __props__.__dict__["reset_approvals_on_push"] = reset_approvals_on_push
            __props__.__dict__["selective_code_owner_removals"] = selective_code_owner_removals
        super(ProjectLevelMrApprovals, __self__).__init__(
            'gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disable_overriding_approvers_per_merge_request: Optional[pulumi.Input[bool]] = None,
            merge_requests_author_approval: Optional[pulumi.Input[bool]] = None,
            merge_requests_disable_committers_approval: Optional[pulumi.Input[bool]] = None,
            project: Optional[pulumi.Input[str]] = None,
            require_password_to_approve: Optional[pulumi.Input[bool]] = None,
            reset_approvals_on_push: Optional[pulumi.Input[bool]] = None,
            selective_code_owner_removals: Optional[pulumi.Input[bool]] = None) -> 'ProjectLevelMrApprovals':
        """
        Get an existing ProjectLevelMrApprovals resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_overriding_approvers_per_merge_request: Set to `true` to disable overriding approvers per merge request.
        :param pulumi.Input[bool] merge_requests_author_approval: Set to `true` to allow merge requests authors to approve their own merge requests.
        :param pulumi.Input[bool] merge_requests_disable_committers_approval: Set to `true` to disable merge request committers from approving their own merge requests.
        :param pulumi.Input[str] project: The ID or URL-encoded path of a project to change MR approval configuration.
        :param pulumi.Input[bool] require_password_to_approve: Set to `true` to require authentication to approve merge requests.
        :param pulumi.Input[bool] reset_approvals_on_push: Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        :param pulumi.Input[bool] selective_code_owner_removals: Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectLevelMrApprovalsState.__new__(_ProjectLevelMrApprovalsState)

        __props__.__dict__["disable_overriding_approvers_per_merge_request"] = disable_overriding_approvers_per_merge_request
        __props__.__dict__["merge_requests_author_approval"] = merge_requests_author_approval
        __props__.__dict__["merge_requests_disable_committers_approval"] = merge_requests_disable_committers_approval
        __props__.__dict__["project"] = project
        __props__.__dict__["require_password_to_approve"] = require_password_to_approve
        __props__.__dict__["reset_approvals_on_push"] = reset_approvals_on_push
        __props__.__dict__["selective_code_owner_removals"] = selective_code_owner_removals
        return ProjectLevelMrApprovals(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableOverridingApproversPerMergeRequest")
    def disable_overriding_approvers_per_merge_request(self) -> pulumi.Output[bool]:
        """
        Set to `true` to disable overriding approvers per merge request.
        """
        return pulumi.get(self, "disable_overriding_approvers_per_merge_request")

    @property
    @pulumi.getter(name="mergeRequestsAuthorApproval")
    def merge_requests_author_approval(self) -> pulumi.Output[bool]:
        """
        Set to `true` to allow merge requests authors to approve their own merge requests.
        """
        return pulumi.get(self, "merge_requests_author_approval")

    @property
    @pulumi.getter(name="mergeRequestsDisableCommittersApproval")
    def merge_requests_disable_committers_approval(self) -> pulumi.Output[bool]:
        """
        Set to `true` to disable merge request committers from approving their own merge requests.
        """
        return pulumi.get(self, "merge_requests_disable_committers_approval")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of a project to change MR approval configuration.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="requirePasswordToApprove")
    def require_password_to_approve(self) -> pulumi.Output[bool]:
        """
        Set to `true` to require authentication to approve merge requests.
        """
        return pulumi.get(self, "require_password_to_approve")

    @property
    @pulumi.getter(name="resetApprovalsOnPush")
    def reset_approvals_on_push(self) -> pulumi.Output[bool]:
        """
        Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        """
        return pulumi.get(self, "reset_approvals_on_push")

    @property
    @pulumi.getter(name="selectiveCodeOwnerRemovals")
    def selective_code_owner_removals(self) -> pulumi.Output[bool]:
        """
        Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
        """
        return pulumi.get(self, "selective_code_owner_removals")

