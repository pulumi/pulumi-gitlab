# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectSecurityPolicyAttachmentArgs', 'ProjectSecurityPolicyAttachment']

@pulumi.input_type
class ProjectSecurityPolicyAttachmentArgs:
    def __init__(__self__, *,
                 policy_project: pulumi.Input[str],
                 project: pulumi.Input[str]):
        """
        The set of arguments for constructing a ProjectSecurityPolicyAttachment resource.
        :param pulumi.Input[str] policy_project: The ID or Full Path of the security policy project.
        :param pulumi.Input[str] project: The ID or Full Path of the project which will have the security policy project assigned to it.
        """
        pulumi.set(__self__, "policy_project", policy_project)
        pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="policyProject")
    def policy_project(self) -> pulumi.Input[str]:
        """
        The ID or Full Path of the security policy project.
        """
        return pulumi.get(self, "policy_project")

    @policy_project.setter
    def policy_project(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_project", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or Full Path of the project which will have the security policy project assigned to it.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _ProjectSecurityPolicyAttachmentState:
    def __init__(__self__, *,
                 policy_project: Optional[pulumi.Input[str]] = None,
                 policy_project_graphql_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 project_graphql_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectSecurityPolicyAttachment resources.
        :param pulumi.Input[str] policy_project: The ID or Full Path of the security policy project.
        :param pulumi.Input[str] policy_project_graphql_id: The GraphQL ID of the security policy project.
        :param pulumi.Input[str] project: The ID or Full Path of the project which will have the security policy project assigned to it.
        :param pulumi.Input[str] project_graphql_id: The GraphQL ID of the project to which the security policty project will be attached.
        """
        if policy_project is not None:
            pulumi.set(__self__, "policy_project", policy_project)
        if policy_project_graphql_id is not None:
            pulumi.set(__self__, "policy_project_graphql_id", policy_project_graphql_id)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if project_graphql_id is not None:
            pulumi.set(__self__, "project_graphql_id", project_graphql_id)

    @property
    @pulumi.getter(name="policyProject")
    def policy_project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or Full Path of the security policy project.
        """
        return pulumi.get(self, "policy_project")

    @policy_project.setter
    def policy_project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_project", value)

    @property
    @pulumi.getter(name="policyProjectGraphqlId")
    def policy_project_graphql_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GraphQL ID of the security policy project.
        """
        return pulumi.get(self, "policy_project_graphql_id")

    @policy_project_graphql_id.setter
    def policy_project_graphql_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_project_graphql_id", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or Full Path of the project which will have the security policy project assigned to it.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="projectGraphqlId")
    def project_graphql_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GraphQL ID of the project to which the security policty project will be attached.
        """
        return pulumi.get(self, "project_graphql_id")

    @project_graphql_id.setter
    def project_graphql_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_graphql_id", value)


class ProjectSecurityPolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_project: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `ProjectSecurityPolicyAttachment` resource allows to attach a security policy project to a project.

        **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/index.html#mutationsecuritypolicyprojectassign)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # This resource can be used to attach a security policy to a pre-existing project
        foo = gitlab.ProjectSecurityPolicyAttachment("foo",
            project="1234",
            policy_project="4567")
        # Or you can use Terraform to create a new project, add a policy to that project,
        # then attach that policy project to other projects.
        my_policy_project = gitlab.Project("my-policy-project", name="security-policy-project")
        policy_yml = gitlab.RepositoryFile("policy-yml",
            project=my_policy_project.id,
            file_path=".gitlab/security-policies/my-policy.yml",
            branch="master",
            encoding="text",
            content=\"\"\"---
        approval_policy:
        - name: test
        description: test
        enabled: true
        rules:
        - type: any_merge_request
            branch_type: protected
            commits: any
        approval_settings:
            block_branch_modification: true
            prevent_pushing_and_force_pushing: true
            prevent_approval_by_author: true
            prevent_approval_by_commit_author: true
            remove_approvals_with_new_commit: true
            require_password_to_approve: false
        fallback_behavior:
            fail: closed
        actions:
        - type: send_bot_message
            enabled: true
        \"\"\")
        my_policy = gitlab.ProjectSecurityPolicyAttachment("my-policy",
            project="1234",
            policy_project=my_policy_project.id)
        ```

        ## Import

        GitLab project security policy attachments can be imported using an id made up of `project:policy_project_id` where the policy project ID is the project ID of the policy project, e.g.

        ```sh
        $ pulumi import gitlab:index/projectSecurityPolicyAttachment:ProjectSecurityPolicyAttachment foo 1:2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_project: The ID or Full Path of the security policy project.
        :param pulumi.Input[str] project: The ID or Full Path of the project which will have the security policy project assigned to it.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectSecurityPolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectSecurityPolicyAttachment` resource allows to attach a security policy project to a project.

        **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/index.html#mutationsecuritypolicyprojectassign)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # This resource can be used to attach a security policy to a pre-existing project
        foo = gitlab.ProjectSecurityPolicyAttachment("foo",
            project="1234",
            policy_project="4567")
        # Or you can use Terraform to create a new project, add a policy to that project,
        # then attach that policy project to other projects.
        my_policy_project = gitlab.Project("my-policy-project", name="security-policy-project")
        policy_yml = gitlab.RepositoryFile("policy-yml",
            project=my_policy_project.id,
            file_path=".gitlab/security-policies/my-policy.yml",
            branch="master",
            encoding="text",
            content=\"\"\"---
        approval_policy:
        - name: test
        description: test
        enabled: true
        rules:
        - type: any_merge_request
            branch_type: protected
            commits: any
        approval_settings:
            block_branch_modification: true
            prevent_pushing_and_force_pushing: true
            prevent_approval_by_author: true
            prevent_approval_by_commit_author: true
            remove_approvals_with_new_commit: true
            require_password_to_approve: false
        fallback_behavior:
            fail: closed
        actions:
        - type: send_bot_message
            enabled: true
        \"\"\")
        my_policy = gitlab.ProjectSecurityPolicyAttachment("my-policy",
            project="1234",
            policy_project=my_policy_project.id)
        ```

        ## Import

        GitLab project security policy attachments can be imported using an id made up of `project:policy_project_id` where the policy project ID is the project ID of the policy project, e.g.

        ```sh
        $ pulumi import gitlab:index/projectSecurityPolicyAttachment:ProjectSecurityPolicyAttachment foo 1:2
        ```

        :param str resource_name: The name of the resource.
        :param ProjectSecurityPolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectSecurityPolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_project: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectSecurityPolicyAttachmentArgs.__new__(ProjectSecurityPolicyAttachmentArgs)

            if policy_project is None and not opts.urn:
                raise TypeError("Missing required property 'policy_project'")
            __props__.__dict__["policy_project"] = policy_project
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["policy_project_graphql_id"] = None
            __props__.__dict__["project_graphql_id"] = None
        super(ProjectSecurityPolicyAttachment, __self__).__init__(
            'gitlab:index/projectSecurityPolicyAttachment:ProjectSecurityPolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy_project: Optional[pulumi.Input[str]] = None,
            policy_project_graphql_id: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            project_graphql_id: Optional[pulumi.Input[str]] = None) -> 'ProjectSecurityPolicyAttachment':
        """
        Get an existing ProjectSecurityPolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_project: The ID or Full Path of the security policy project.
        :param pulumi.Input[str] policy_project_graphql_id: The GraphQL ID of the security policy project.
        :param pulumi.Input[str] project: The ID or Full Path of the project which will have the security policy project assigned to it.
        :param pulumi.Input[str] project_graphql_id: The GraphQL ID of the project to which the security policty project will be attached.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectSecurityPolicyAttachmentState.__new__(_ProjectSecurityPolicyAttachmentState)

        __props__.__dict__["policy_project"] = policy_project
        __props__.__dict__["policy_project_graphql_id"] = policy_project_graphql_id
        __props__.__dict__["project"] = project
        __props__.__dict__["project_graphql_id"] = project_graphql_id
        return ProjectSecurityPolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policyProject")
    def policy_project(self) -> pulumi.Output[str]:
        """
        The ID or Full Path of the security policy project.
        """
        return pulumi.get(self, "policy_project")

    @property
    @pulumi.getter(name="policyProjectGraphqlId")
    def policy_project_graphql_id(self) -> pulumi.Output[str]:
        """
        The GraphQL ID of the security policy project.
        """
        return pulumi.get(self, "policy_project_graphql_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or Full Path of the project which will have the security policy project assigned to it.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="projectGraphqlId")
    def project_graphql_id(self) -> pulumi.Output[str]:
        """
        The GraphQL ID of the project to which the security policty project will be attached.
        """
        return pulumi.get(self, "project_graphql_id")

