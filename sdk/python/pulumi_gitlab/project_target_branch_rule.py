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

__all__ = ['ProjectTargetBranchRuleArgs', 'ProjectTargetBranchRule']

@pulumi.input_type
class ProjectTargetBranchRuleArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 source_branch_pattern: pulumi.Input[str],
                 target_branch_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a ProjectTargetBranchRule resource.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        :param pulumi.Input[str] source_branch_pattern: A pattern matching the branch name for which the merge request should have a default target branch configured.
        :param pulumi.Input[str] target_branch_name: The name of the branch to which the merge request should be addressed.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "source_branch_pattern", source_branch_pattern)
        pulumi.set(__self__, "target_branch_name", target_branch_name)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or URL-encoded path of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="sourceBranchPattern")
    def source_branch_pattern(self) -> pulumi.Input[str]:
        """
        A pattern matching the branch name for which the merge request should have a default target branch configured.
        """
        return pulumi.get(self, "source_branch_pattern")

    @source_branch_pattern.setter
    def source_branch_pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_branch_pattern", value)

    @property
    @pulumi.getter(name="targetBranchName")
    def target_branch_name(self) -> pulumi.Input[str]:
        """
        The name of the branch to which the merge request should be addressed.
        """
        return pulumi.get(self, "target_branch_name")

    @target_branch_name.setter
    def target_branch_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_branch_name", value)


@pulumi.input_type
class _ProjectTargetBranchRuleState:
    def __init__(__self__, *,
                 project: Optional[pulumi.Input[str]] = None,
                 source_branch_pattern: Optional[pulumi.Input[str]] = None,
                 target_branch_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectTargetBranchRule resources.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        :param pulumi.Input[str] source_branch_pattern: A pattern matching the branch name for which the merge request should have a default target branch configured.
        :param pulumi.Input[str] target_branch_name: The name of the branch to which the merge request should be addressed.
        """
        if project is not None:
            pulumi.set(__self__, "project", project)
        if source_branch_pattern is not None:
            pulumi.set(__self__, "source_branch_pattern", source_branch_pattern)
        if target_branch_name is not None:
            pulumi.set(__self__, "target_branch_name", target_branch_name)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or URL-encoded path of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="sourceBranchPattern")
    def source_branch_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        A pattern matching the branch name for which the merge request should have a default target branch configured.
        """
        return pulumi.get(self, "source_branch_pattern")

    @source_branch_pattern.setter
    def source_branch_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_branch_pattern", value)

    @property
    @pulumi.getter(name="targetBranchName")
    def target_branch_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the branch to which the merge request should be addressed.
        """
        return pulumi.get(self, "target_branch_name")

    @target_branch_name.setter
    def target_branch_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_branch_name", value)


class ProjectTargetBranchRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 source_branch_pattern: Optional[pulumi.Input[str]] = None,
                 target_branch_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `ProjectTargetBranchRule` resource allows to configure default target branch rules when creating a merge request.

        **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#mutationprojecttargetbranchrulecreate)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        :param pulumi.Input[str] source_branch_pattern: A pattern matching the branch name for which the merge request should have a default target branch configured.
        :param pulumi.Input[str] target_branch_name: The name of the branch to which the merge request should be addressed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectTargetBranchRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectTargetBranchRule` resource allows to configure default target branch rules when creating a merge request.

        **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#mutationprojecttargetbranchrulecreate)

        :param str resource_name: The name of the resource.
        :param ProjectTargetBranchRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectTargetBranchRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 source_branch_pattern: Optional[pulumi.Input[str]] = None,
                 target_branch_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectTargetBranchRuleArgs.__new__(ProjectTargetBranchRuleArgs)

            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if source_branch_pattern is None and not opts.urn:
                raise TypeError("Missing required property 'source_branch_pattern'")
            __props__.__dict__["source_branch_pattern"] = source_branch_pattern
            if target_branch_name is None and not opts.urn:
                raise TypeError("Missing required property 'target_branch_name'")
            __props__.__dict__["target_branch_name"] = target_branch_name
        super(ProjectTargetBranchRule, __self__).__init__(
            'gitlab:index/projectTargetBranchRule:ProjectTargetBranchRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            project: Optional[pulumi.Input[str]] = None,
            source_branch_pattern: Optional[pulumi.Input[str]] = None,
            target_branch_name: Optional[pulumi.Input[str]] = None) -> 'ProjectTargetBranchRule':
        """
        Get an existing ProjectTargetBranchRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project.
        :param pulumi.Input[str] source_branch_pattern: A pattern matching the branch name for which the merge request should have a default target branch configured.
        :param pulumi.Input[str] target_branch_name: The name of the branch to which the merge request should be addressed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectTargetBranchRuleState.__new__(_ProjectTargetBranchRuleState)

        __props__.__dict__["project"] = project
        __props__.__dict__["source_branch_pattern"] = source_branch_pattern
        __props__.__dict__["target_branch_name"] = target_branch_name
        return ProjectTargetBranchRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="sourceBranchPattern")
    def source_branch_pattern(self) -> pulumi.Output[str]:
        """
        A pattern matching the branch name for which the merge request should have a default target branch configured.
        """
        return pulumi.get(self, "source_branch_pattern")

    @property
    @pulumi.getter(name="targetBranchName")
    def target_branch_name(self) -> pulumi.Output[str]:
        """
        The name of the branch to which the merge request should be addressed.
        """
        return pulumi.get(self, "target_branch_name")

