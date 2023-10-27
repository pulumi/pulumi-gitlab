# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IntegrationEmailsOnPushArgs', 'IntegrationEmailsOnPush']

@pulumi.input_type
class IntegrationEmailsOnPushArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 recipients: pulumi.Input[str],
                 branches_to_be_notified: Optional[pulumi.Input[str]] = None,
                 disable_diffs: Optional[pulumi.Input[bool]] = None,
                 push_events: Optional[pulumi.Input[bool]] = None,
                 send_from_committer_email: Optional[pulumi.Input[bool]] = None,
                 tag_push_events: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a IntegrationEmailsOnPush resource.
        :param pulumi.Input[str] project: ID or full-path of the project you want to activate integration on.
        :param pulumi.Input[str] recipients: Emails separated by whitespace.
        :param pulumi.Input[str] branches_to_be_notified: Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
        :param pulumi.Input[bool] disable_diffs: Disable code diffs.
        :param pulumi.Input[bool] push_events: Enable notifications for push events.
        :param pulumi.Input[bool] send_from_committer_email: Send from committer.
        :param pulumi.Input[bool] tag_push_events: Enable notifications for tag push events.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "recipients", recipients)
        if branches_to_be_notified is not None:
            pulumi.set(__self__, "branches_to_be_notified", branches_to_be_notified)
        if disable_diffs is not None:
            pulumi.set(__self__, "disable_diffs", disable_diffs)
        if push_events is not None:
            pulumi.set(__self__, "push_events", push_events)
        if send_from_committer_email is not None:
            pulumi.set(__self__, "send_from_committer_email", send_from_committer_email)
        if tag_push_events is not None:
            pulumi.set(__self__, "tag_push_events", tag_push_events)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        ID or full-path of the project you want to activate integration on.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def recipients(self) -> pulumi.Input[str]:
        """
        Emails separated by whitespace.
        """
        return pulumi.get(self, "recipients")

    @recipients.setter
    def recipients(self, value: pulumi.Input[str]):
        pulumi.set(self, "recipients", value)

    @property
    @pulumi.getter(name="branchesToBeNotified")
    def branches_to_be_notified(self) -> Optional[pulumi.Input[str]]:
        """
        Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
        """
        return pulumi.get(self, "branches_to_be_notified")

    @branches_to_be_notified.setter
    def branches_to_be_notified(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branches_to_be_notified", value)

    @property
    @pulumi.getter(name="disableDiffs")
    def disable_diffs(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable code diffs.
        """
        return pulumi.get(self, "disable_diffs")

    @disable_diffs.setter
    def disable_diffs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_diffs", value)

    @property
    @pulumi.getter(name="pushEvents")
    def push_events(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable notifications for push events.
        """
        return pulumi.get(self, "push_events")

    @push_events.setter
    def push_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "push_events", value)

    @property
    @pulumi.getter(name="sendFromCommitterEmail")
    def send_from_committer_email(self) -> Optional[pulumi.Input[bool]]:
        """
        Send from committer.
        """
        return pulumi.get(self, "send_from_committer_email")

    @send_from_committer_email.setter
    def send_from_committer_email(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_from_committer_email", value)

    @property
    @pulumi.getter(name="tagPushEvents")
    def tag_push_events(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable notifications for tag push events.
        """
        return pulumi.get(self, "tag_push_events")

    @tag_push_events.setter
    def tag_push_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tag_push_events", value)


@pulumi.input_type
class _IntegrationEmailsOnPushState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 branches_to_be_notified: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 disable_diffs: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 push_events: Optional[pulumi.Input[bool]] = None,
                 recipients: Optional[pulumi.Input[str]] = None,
                 send_from_committer_email: Optional[pulumi.Input[bool]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 tag_push_events: Optional[pulumi.Input[bool]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IntegrationEmailsOnPush resources.
        :param pulumi.Input[bool] active: Whether the integration is active.
        :param pulumi.Input[str] branches_to_be_notified: Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
        :param pulumi.Input[str] created_at: The ISO8601 date/time that this integration was activated at in UTC.
        :param pulumi.Input[bool] disable_diffs: Disable code diffs.
        :param pulumi.Input[str] project: ID or full-path of the project you want to activate integration on.
        :param pulumi.Input[bool] push_events: Enable notifications for push events.
        :param pulumi.Input[str] recipients: Emails separated by whitespace.
        :param pulumi.Input[bool] send_from_committer_email: Send from committer.
        :param pulumi.Input[str] slug: The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        :param pulumi.Input[bool] tag_push_events: Enable notifications for tag push events.
        :param pulumi.Input[str] title: Title of the integration.
        :param pulumi.Input[str] updated_at: The ISO8601 date/time that this integration was last updated at in UTC.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if branches_to_be_notified is not None:
            pulumi.set(__self__, "branches_to_be_notified", branches_to_be_notified)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if disable_diffs is not None:
            pulumi.set(__self__, "disable_diffs", disable_diffs)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if push_events is not None:
            pulumi.set(__self__, "push_events", push_events)
        if recipients is not None:
            pulumi.set(__self__, "recipients", recipients)
        if send_from_committer_email is not None:
            pulumi.set(__self__, "send_from_committer_email", send_from_committer_email)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if tag_push_events is not None:
            pulumi.set(__self__, "tag_push_events", tag_push_events)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the integration is active.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter(name="branchesToBeNotified")
    def branches_to_be_notified(self) -> Optional[pulumi.Input[str]]:
        """
        Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
        """
        return pulumi.get(self, "branches_to_be_notified")

    @branches_to_be_notified.setter
    def branches_to_be_notified(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branches_to_be_notified", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO8601 date/time that this integration was activated at in UTC.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="disableDiffs")
    def disable_diffs(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable code diffs.
        """
        return pulumi.get(self, "disable_diffs")

    @disable_diffs.setter
    def disable_diffs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_diffs", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        ID or full-path of the project you want to activate integration on.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="pushEvents")
    def push_events(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable notifications for push events.
        """
        return pulumi.get(self, "push_events")

    @push_events.setter
    def push_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "push_events", value)

    @property
    @pulumi.getter
    def recipients(self) -> Optional[pulumi.Input[str]]:
        """
        Emails separated by whitespace.
        """
        return pulumi.get(self, "recipients")

    @recipients.setter
    def recipients(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recipients", value)

    @property
    @pulumi.getter(name="sendFromCommitterEmail")
    def send_from_committer_email(self) -> Optional[pulumi.Input[bool]]:
        """
        Send from committer.
        """
        return pulumi.get(self, "send_from_committer_email")

    @send_from_committer_email.setter
    def send_from_committer_email(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_from_committer_email", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        """
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter(name="tagPushEvents")
    def tag_push_events(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable notifications for tag push events.
        """
        return pulumi.get(self, "tag_push_events")

    @tag_push_events.setter
    def tag_push_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tag_push_events", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Title of the integration.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO8601 date/time that this integration was last updated at in UTC.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class IntegrationEmailsOnPush(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branches_to_be_notified: Optional[pulumi.Input[str]] = None,
                 disable_diffs: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 push_events: Optional[pulumi.Input[bool]] = None,
                 recipients: Optional[pulumi.Input[str]] = None,
                 send_from_committer_email: Optional[pulumi.Input[bool]] = None,
                 tag_push_events: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The `IntegrationEmailsOnPush` resource allows to manage the lifecycle of a project integration with Emails on Push Service.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#emails-on-push)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        awesome_project = gitlab.Project("awesomeProject",
            description="My awesome project.",
            visibility_level="public")
        emails = gitlab.IntegrationEmailsOnPush("emails",
            project=awesome_project.id,
            recipients="myrecipient@example.com myotherrecipient@example.com")
        ```

        ## Import

        You can import a gitlab_integration_emails_on_push state using the project ID, e.g.

        ```sh
         $ pulumi import gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush emails 1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branches_to_be_notified: Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
        :param pulumi.Input[bool] disable_diffs: Disable code diffs.
        :param pulumi.Input[str] project: ID or full-path of the project you want to activate integration on.
        :param pulumi.Input[bool] push_events: Enable notifications for push events.
        :param pulumi.Input[str] recipients: Emails separated by whitespace.
        :param pulumi.Input[bool] send_from_committer_email: Send from committer.
        :param pulumi.Input[bool] tag_push_events: Enable notifications for tag push events.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationEmailsOnPushArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `IntegrationEmailsOnPush` resource allows to manage the lifecycle of a project integration with Emails on Push Service.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#emails-on-push)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        awesome_project = gitlab.Project("awesomeProject",
            description="My awesome project.",
            visibility_level="public")
        emails = gitlab.IntegrationEmailsOnPush("emails",
            project=awesome_project.id,
            recipients="myrecipient@example.com myotherrecipient@example.com")
        ```

        ## Import

        You can import a gitlab_integration_emails_on_push state using the project ID, e.g.

        ```sh
         $ pulumi import gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush emails 1
        ```

        :param str resource_name: The name of the resource.
        :param IntegrationEmailsOnPushArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntegrationEmailsOnPushArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branches_to_be_notified: Optional[pulumi.Input[str]] = None,
                 disable_diffs: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 push_events: Optional[pulumi.Input[bool]] = None,
                 recipients: Optional[pulumi.Input[str]] = None,
                 send_from_committer_email: Optional[pulumi.Input[bool]] = None,
                 tag_push_events: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntegrationEmailsOnPushArgs.__new__(IntegrationEmailsOnPushArgs)

            __props__.__dict__["branches_to_be_notified"] = branches_to_be_notified
            __props__.__dict__["disable_diffs"] = disable_diffs
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["push_events"] = push_events
            if recipients is None and not opts.urn:
                raise TypeError("Missing required property 'recipients'")
            __props__.__dict__["recipients"] = recipients
            __props__.__dict__["send_from_committer_email"] = send_from_committer_email
            __props__.__dict__["tag_push_events"] = tag_push_events
            __props__.__dict__["active"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["slug"] = None
            __props__.__dict__["title"] = None
            __props__.__dict__["updated_at"] = None
        super(IntegrationEmailsOnPush, __self__).__init__(
            'gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            branches_to_be_notified: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            disable_diffs: Optional[pulumi.Input[bool]] = None,
            project: Optional[pulumi.Input[str]] = None,
            push_events: Optional[pulumi.Input[bool]] = None,
            recipients: Optional[pulumi.Input[str]] = None,
            send_from_committer_email: Optional[pulumi.Input[bool]] = None,
            slug: Optional[pulumi.Input[str]] = None,
            tag_push_events: Optional[pulumi.Input[bool]] = None,
            title: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'IntegrationEmailsOnPush':
        """
        Get an existing IntegrationEmailsOnPush resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Whether the integration is active.
        :param pulumi.Input[str] branches_to_be_notified: Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
        :param pulumi.Input[str] created_at: The ISO8601 date/time that this integration was activated at in UTC.
        :param pulumi.Input[bool] disable_diffs: Disable code diffs.
        :param pulumi.Input[str] project: ID or full-path of the project you want to activate integration on.
        :param pulumi.Input[bool] push_events: Enable notifications for push events.
        :param pulumi.Input[str] recipients: Emails separated by whitespace.
        :param pulumi.Input[bool] send_from_committer_email: Send from committer.
        :param pulumi.Input[str] slug: The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        :param pulumi.Input[bool] tag_push_events: Enable notifications for tag push events.
        :param pulumi.Input[str] title: Title of the integration.
        :param pulumi.Input[str] updated_at: The ISO8601 date/time that this integration was last updated at in UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IntegrationEmailsOnPushState.__new__(_IntegrationEmailsOnPushState)

        __props__.__dict__["active"] = active
        __props__.__dict__["branches_to_be_notified"] = branches_to_be_notified
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["disable_diffs"] = disable_diffs
        __props__.__dict__["project"] = project
        __props__.__dict__["push_events"] = push_events
        __props__.__dict__["recipients"] = recipients
        __props__.__dict__["send_from_committer_email"] = send_from_committer_email
        __props__.__dict__["slug"] = slug
        __props__.__dict__["tag_push_events"] = tag_push_events
        __props__.__dict__["title"] = title
        __props__.__dict__["updated_at"] = updated_at
        return IntegrationEmailsOnPush(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[bool]:
        """
        Whether the integration is active.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter(name="branchesToBeNotified")
    def branches_to_be_notified(self) -> pulumi.Output[Optional[str]]:
        """
        Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
        """
        return pulumi.get(self, "branches_to_be_notified")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The ISO8601 date/time that this integration was activated at in UTC.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="disableDiffs")
    def disable_diffs(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable code diffs.
        """
        return pulumi.get(self, "disable_diffs")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        ID or full-path of the project you want to activate integration on.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="pushEvents")
    def push_events(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable notifications for push events.
        """
        return pulumi.get(self, "push_events")

    @property
    @pulumi.getter
    def recipients(self) -> pulumi.Output[str]:
        """
        Emails separated by whitespace.
        """
        return pulumi.get(self, "recipients")

    @property
    @pulumi.getter(name="sendFromCommitterEmail")
    def send_from_committer_email(self) -> pulumi.Output[Optional[bool]]:
        """
        Send from committer.
        """
        return pulumi.get(self, "send_from_committer_email")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        """
        The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        """
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter(name="tagPushEvents")
    def tag_push_events(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable notifications for tag push events.
        """
        return pulumi.get(self, "tag_push_events")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        Title of the integration.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The ISO8601 date/time that this integration was last updated at in UTC.
        """
        return pulumi.get(self, "updated_at")

