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

__all__ = ['IntegrationPipelinesEmailArgs', 'IntegrationPipelinesEmail']

@pulumi.input_type
class IntegrationPipelinesEmailArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[builtins.str],
                 recipients: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 branches_to_be_notified: Optional[pulumi.Input[builtins.str]] = None,
                 notify_only_broken_pipelines: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a IntegrationPipelinesEmail resource.
        :param pulumi.Input[builtins.str] project: ID of the project you want to activate integration on.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] recipients: ) email addresses where notifications are sent.
        :param pulumi.Input[builtins.str] branches_to_be_notified: Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
        :param pulumi.Input[builtins.bool] notify_only_broken_pipelines: Notify only broken pipelines. Default is true.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "recipients", recipients)
        if branches_to_be_notified is not None:
            pulumi.set(__self__, "branches_to_be_notified", branches_to_be_notified)
        if notify_only_broken_pipelines is not None:
            pulumi.set(__self__, "notify_only_broken_pipelines", notify_only_broken_pipelines)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[builtins.str]:
        """
        ID of the project you want to activate integration on.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def recipients(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        ) email addresses where notifications are sent.
        """
        return pulumi.get(self, "recipients")

    @recipients.setter
    def recipients(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "recipients", value)

    @property
    @pulumi.getter(name="branchesToBeNotified")
    def branches_to_be_notified(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
        """
        return pulumi.get(self, "branches_to_be_notified")

    @branches_to_be_notified.setter
    def branches_to_be_notified(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "branches_to_be_notified", value)

    @property
    @pulumi.getter(name="notifyOnlyBrokenPipelines")
    def notify_only_broken_pipelines(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Notify only broken pipelines. Default is true.
        """
        return pulumi.get(self, "notify_only_broken_pipelines")

    @notify_only_broken_pipelines.setter
    def notify_only_broken_pipelines(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "notify_only_broken_pipelines", value)


@pulumi.input_type
class _IntegrationPipelinesEmailState:
    def __init__(__self__, *,
                 branches_to_be_notified: Optional[pulumi.Input[builtins.str]] = None,
                 notify_only_broken_pipelines: Optional[pulumi.Input[builtins.bool]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering IntegrationPipelinesEmail resources.
        :param pulumi.Input[builtins.str] branches_to_be_notified: Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
        :param pulumi.Input[builtins.bool] notify_only_broken_pipelines: Notify only broken pipelines. Default is true.
        :param pulumi.Input[builtins.str] project: ID of the project you want to activate integration on.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] recipients: ) email addresses where notifications are sent.
        """
        if branches_to_be_notified is not None:
            pulumi.set(__self__, "branches_to_be_notified", branches_to_be_notified)
        if notify_only_broken_pipelines is not None:
            pulumi.set(__self__, "notify_only_broken_pipelines", notify_only_broken_pipelines)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if recipients is not None:
            pulumi.set(__self__, "recipients", recipients)

    @property
    @pulumi.getter(name="branchesToBeNotified")
    def branches_to_be_notified(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
        """
        return pulumi.get(self, "branches_to_be_notified")

    @branches_to_be_notified.setter
    def branches_to_be_notified(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "branches_to_be_notified", value)

    @property
    @pulumi.getter(name="notifyOnlyBrokenPipelines")
    def notify_only_broken_pipelines(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Notify only broken pipelines. Default is true.
        """
        return pulumi.get(self, "notify_only_broken_pipelines")

    @notify_only_broken_pipelines.setter
    def notify_only_broken_pipelines(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "notify_only_broken_pipelines", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the project you want to activate integration on.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def recipients(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        ) email addresses where notifications are sent.
        """
        return pulumi.get(self, "recipients")

    @recipients.setter
    def recipients(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "recipients", value)


class IntegrationPipelinesEmail(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branches_to_be_notified: Optional[pulumi.Input[builtins.str]] = None,
                 notify_only_broken_pipelines: Optional[pulumi.Input[builtins.bool]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        The `IntegrationPipelinesEmail` resource allows to manage the lifecycle of a project integration with Pipeline Emails Service.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/integrations/#pipeline-emails)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        awesome_project = gitlab.Project("awesome_project",
            name="awesome_project",
            description="My awesome project.",
            visibility_level="public")
        email = gitlab.IntegrationPipelinesEmail("email",
            project=awesome_project.id,
            recipients=["gitlab@user.create"],
            notify_only_broken_pipelines=True,
            branches_to_be_notified="all")
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_pipelines_email`. For example:

        terraform

        import {

          to = gitlab_integration_pipelines_email.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        You can import a gitlab_integration_pipelines_email state using the project ID, e.g.

        ```sh
        $ pulumi import gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail email 1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] branches_to_be_notified: Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
        :param pulumi.Input[builtins.bool] notify_only_broken_pipelines: Notify only broken pipelines. Default is true.
        :param pulumi.Input[builtins.str] project: ID of the project you want to activate integration on.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] recipients: ) email addresses where notifications are sent.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationPipelinesEmailArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `IntegrationPipelinesEmail` resource allows to manage the lifecycle of a project integration with Pipeline Emails Service.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/integrations/#pipeline-emails)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        awesome_project = gitlab.Project("awesome_project",
            name="awesome_project",
            description="My awesome project.",
            visibility_level="public")
        email = gitlab.IntegrationPipelinesEmail("email",
            project=awesome_project.id,
            recipients=["gitlab@user.create"],
            notify_only_broken_pipelines=True,
            branches_to_be_notified="all")
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_pipelines_email`. For example:

        terraform

        import {

          to = gitlab_integration_pipelines_email.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        You can import a gitlab_integration_pipelines_email state using the project ID, e.g.

        ```sh
        $ pulumi import gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail email 1
        ```

        :param str resource_name: The name of the resource.
        :param IntegrationPipelinesEmailArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntegrationPipelinesEmailArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branches_to_be_notified: Optional[pulumi.Input[builtins.str]] = None,
                 notify_only_broken_pipelines: Optional[pulumi.Input[builtins.bool]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntegrationPipelinesEmailArgs.__new__(IntegrationPipelinesEmailArgs)

            __props__.__dict__["branches_to_be_notified"] = branches_to_be_notified
            __props__.__dict__["notify_only_broken_pipelines"] = notify_only_broken_pipelines
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if recipients is None and not opts.urn:
                raise TypeError("Missing required property 'recipients'")
            __props__.__dict__["recipients"] = recipients
        super(IntegrationPipelinesEmail, __self__).__init__(
            'gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branches_to_be_notified: Optional[pulumi.Input[builtins.str]] = None,
            notify_only_broken_pipelines: Optional[pulumi.Input[builtins.bool]] = None,
            project: Optional[pulumi.Input[builtins.str]] = None,
            recipients: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'IntegrationPipelinesEmail':
        """
        Get an existing IntegrationPipelinesEmail resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] branches_to_be_notified: Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
        :param pulumi.Input[builtins.bool] notify_only_broken_pipelines: Notify only broken pipelines. Default is true.
        :param pulumi.Input[builtins.str] project: ID of the project you want to activate integration on.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] recipients: ) email addresses where notifications are sent.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IntegrationPipelinesEmailState.__new__(_IntegrationPipelinesEmailState)

        __props__.__dict__["branches_to_be_notified"] = branches_to_be_notified
        __props__.__dict__["notify_only_broken_pipelines"] = notify_only_broken_pipelines
        __props__.__dict__["project"] = project
        __props__.__dict__["recipients"] = recipients
        return IntegrationPipelinesEmail(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="branchesToBeNotified")
    def branches_to_be_notified(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
        """
        return pulumi.get(self, "branches_to_be_notified")

    @property
    @pulumi.getter(name="notifyOnlyBrokenPipelines")
    def notify_only_broken_pipelines(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Notify only broken pipelines. Default is true.
        """
        return pulumi.get(self, "notify_only_broken_pipelines")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[builtins.str]:
        """
        ID of the project you want to activate integration on.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def recipients(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        ) email addresses where notifications are sent.
        """
        return pulumi.get(self, "recipients")

