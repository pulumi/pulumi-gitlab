# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ProjectProtectedEnvironmentArgs', 'ProjectProtectedEnvironment']

@pulumi.input_type
class ProjectProtectedEnvironmentArgs:
    def __init__(__self__, *,
                 environment: pulumi.Input[str],
                 project: pulumi.Input[str],
                 approval_rules: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentApprovalRuleArgs']]]] = None,
                 deploy_access_levels: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]] = None,
                 required_approval_count: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ProjectProtectedEnvironment resource.
        :param pulumi.Input[str] environment: The name of the environment.
        :param pulumi.Input[str] project: The ID or full path of the project which the protected environment is created against.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentApprovalRuleArgs']]] approval_rules: Array of approval rules to deploy, with each described by a hash.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentDeployAccessLevelArgs']]] deploy_access_levels: Array of access levels allowed to deploy, with each described by a hash.
        :param pulumi.Input[int] required_approval_count: The number of approvals required to deploy to this environment.
        """
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "project", project)
        if approval_rules is not None:
            pulumi.set(__self__, "approval_rules", approval_rules)
        if deploy_access_levels is not None:
            pulumi.set(__self__, "deploy_access_levels", deploy_access_levels)
        if required_approval_count is not None:
            pulumi.set(__self__, "required_approval_count", required_approval_count)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Input[str]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or full path of the project which the protected environment is created against.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="approvalRules")
    def approval_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentApprovalRuleArgs']]]]:
        """
        Array of approval rules to deploy, with each described by a hash.
        """
        return pulumi.get(self, "approval_rules")

    @approval_rules.setter
    def approval_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentApprovalRuleArgs']]]]):
        pulumi.set(self, "approval_rules", value)

    @property
    @pulumi.getter(name="deployAccessLevels")
    def deploy_access_levels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]]:
        """
        Array of access levels allowed to deploy, with each described by a hash.
        """
        return pulumi.get(self, "deploy_access_levels")

    @deploy_access_levels.setter
    def deploy_access_levels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]]):
        pulumi.set(self, "deploy_access_levels", value)

    @property
    @pulumi.getter(name="requiredApprovalCount")
    def required_approval_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of approvals required to deploy to this environment.
        """
        return pulumi.get(self, "required_approval_count")

    @required_approval_count.setter
    def required_approval_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "required_approval_count", value)


@pulumi.input_type
class _ProjectProtectedEnvironmentState:
    def __init__(__self__, *,
                 approval_rules: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentApprovalRuleArgs']]]] = None,
                 deploy_access_levels: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 required_approval_count: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ProjectProtectedEnvironment resources.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentApprovalRuleArgs']]] approval_rules: Array of approval rules to deploy, with each described by a hash.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentDeployAccessLevelArgs']]] deploy_access_levels: Array of access levels allowed to deploy, with each described by a hash.
        :param pulumi.Input[str] environment: The name of the environment.
        :param pulumi.Input[str] project: The ID or full path of the project which the protected environment is created against.
        :param pulumi.Input[int] required_approval_count: The number of approvals required to deploy to this environment.
        """
        if approval_rules is not None:
            pulumi.set(__self__, "approval_rules", approval_rules)
        if deploy_access_levels is not None:
            pulumi.set(__self__, "deploy_access_levels", deploy_access_levels)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if required_approval_count is not None:
            pulumi.set(__self__, "required_approval_count", required_approval_count)

    @property
    @pulumi.getter(name="approvalRules")
    def approval_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentApprovalRuleArgs']]]]:
        """
        Array of approval rules to deploy, with each described by a hash.
        """
        return pulumi.get(self, "approval_rules")

    @approval_rules.setter
    def approval_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentApprovalRuleArgs']]]]):
        pulumi.set(self, "approval_rules", value)

    @property
    @pulumi.getter(name="deployAccessLevels")
    def deploy_access_levels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]]:
        """
        Array of access levels allowed to deploy, with each described by a hash.
        """
        return pulumi.get(self, "deploy_access_levels")

    @deploy_access_levels.setter
    def deploy_access_levels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]]):
        pulumi.set(self, "deploy_access_levels", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or full path of the project which the protected environment is created against.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="requiredApprovalCount")
    def required_approval_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of approvals required to deploy to this environment.
        """
        return pulumi.get(self, "required_approval_count")

    @required_approval_count.setter
    def required_approval_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "required_approval_count", value)


class ProjectProtectedEnvironment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approval_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectProtectedEnvironmentApprovalRuleArgs']]]]] = None,
                 deploy_access_levels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 required_approval_count: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The `ProjectProtectedEnvironment` resource allows to manage the lifecycle of a protected environment in a project.

        > In order to use a user or group in the `deploy_access_levels` configuration,
           you need to make sure that users have access to the project and groups must have this project shared.
           You may use the `ProjectMembership` and `gitlab_project_shared_group` resources to achieve this.
           Unfortunately, the GitLab API does not complain about users and groups without access to the project and just ignores those.
           In case this happens you will get perpetual state diffs.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_environments.html)

        ## Import

        GitLab protected environments can be imported using an id made up of `projectId:environmentName`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment bar 123:production
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectProtectedEnvironmentApprovalRuleArgs']]]] approval_rules: Array of approval rules to deploy, with each described by a hash.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]] deploy_access_levels: Array of access levels allowed to deploy, with each described by a hash.
        :param pulumi.Input[str] environment: The name of the environment.
        :param pulumi.Input[str] project: The ID or full path of the project which the protected environment is created against.
        :param pulumi.Input[int] required_approval_count: The number of approvals required to deploy to this environment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectProtectedEnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectProtectedEnvironment` resource allows to manage the lifecycle of a protected environment in a project.

        > In order to use a user or group in the `deploy_access_levels` configuration,
           you need to make sure that users have access to the project and groups must have this project shared.
           You may use the `ProjectMembership` and `gitlab_project_shared_group` resources to achieve this.
           Unfortunately, the GitLab API does not complain about users and groups without access to the project and just ignores those.
           In case this happens you will get perpetual state diffs.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_environments.html)

        ## Import

        GitLab protected environments can be imported using an id made up of `projectId:environmentName`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment bar 123:production
        ```

        :param str resource_name: The name of the resource.
        :param ProjectProtectedEnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectProtectedEnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approval_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectProtectedEnvironmentApprovalRuleArgs']]]]] = None,
                 deploy_access_levels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 required_approval_count: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectProtectedEnvironmentArgs.__new__(ProjectProtectedEnvironmentArgs)

            __props__.__dict__["approval_rules"] = approval_rules
            __props__.__dict__["deploy_access_levels"] = deploy_access_levels
            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["required_approval_count"] = required_approval_count
        super(ProjectProtectedEnvironment, __self__).__init__(
            'gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            approval_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectProtectedEnvironmentApprovalRuleArgs']]]]] = None,
            deploy_access_levels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            required_approval_count: Optional[pulumi.Input[int]] = None) -> 'ProjectProtectedEnvironment':
        """
        Get an existing ProjectProtectedEnvironment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectProtectedEnvironmentApprovalRuleArgs']]]] approval_rules: Array of approval rules to deploy, with each described by a hash.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectProtectedEnvironmentDeployAccessLevelArgs']]]] deploy_access_levels: Array of access levels allowed to deploy, with each described by a hash.
        :param pulumi.Input[str] environment: The name of the environment.
        :param pulumi.Input[str] project: The ID or full path of the project which the protected environment is created against.
        :param pulumi.Input[int] required_approval_count: The number of approvals required to deploy to this environment.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectProtectedEnvironmentState.__new__(_ProjectProtectedEnvironmentState)

        __props__.__dict__["approval_rules"] = approval_rules
        __props__.__dict__["deploy_access_levels"] = deploy_access_levels
        __props__.__dict__["environment"] = environment
        __props__.__dict__["project"] = project
        __props__.__dict__["required_approval_count"] = required_approval_count
        return ProjectProtectedEnvironment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="approvalRules")
    def approval_rules(self) -> pulumi.Output[Sequence['outputs.ProjectProtectedEnvironmentApprovalRule']]:
        """
        Array of approval rules to deploy, with each described by a hash.
        """
        return pulumi.get(self, "approval_rules")

    @property
    @pulumi.getter(name="deployAccessLevels")
    def deploy_access_levels(self) -> pulumi.Output[Optional[Sequence['outputs.ProjectProtectedEnvironmentDeployAccessLevel']]]:
        """
        Array of access levels allowed to deploy, with each described by a hash.
        """
        return pulumi.get(self, "deploy_access_levels")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or full path of the project which the protected environment is created against.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="requiredApprovalCount")
    def required_approval_count(self) -> pulumi.Output[int]:
        """
        The number of approvals required to deploy to this environment.
        """
        return pulumi.get(self, "required_approval_count")

