# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceCustomIssueTrackerArgs', 'ServiceCustomIssueTracker']

@pulumi.input_type
class ServiceCustomIssueTrackerArgs:
    def __init__(__self__, *,
                 issues_url: pulumi.Input[str],
                 project: pulumi.Input[str],
                 project_url: pulumi.Input[str]):
        """
        The set of arguments for constructing a ServiceCustomIssueTracker resource.
        :param pulumi.Input[str] issues_url: The URL to view an issue in the external issue tracker. Must contain :id.
        :param pulumi.Input[str] project: The ID or full path of the project for the custom issue tracker.
        :param pulumi.Input[str] project_url: The URL to the project in the external issue tracker.
        """
        pulumi.set(__self__, "issues_url", issues_url)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "project_url", project_url)

    @property
    @pulumi.getter(name="issuesUrl")
    def issues_url(self) -> pulumi.Input[str]:
        """
        The URL to view an issue in the external issue tracker. Must contain :id.
        """
        return pulumi.get(self, "issues_url")

    @issues_url.setter
    def issues_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "issues_url", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or full path of the project for the custom issue tracker.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="projectUrl")
    def project_url(self) -> pulumi.Input[str]:
        """
        The URL to the project in the external issue tracker.
        """
        return pulumi.get(self, "project_url")

    @project_url.setter
    def project_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_url", value)


@pulumi.input_type
class _ServiceCustomIssueTrackerState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 issues_url: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 project_url: Optional[pulumi.Input[str]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceCustomIssueTracker resources.
        :param pulumi.Input[bool] active: Whether the integration is active.
        :param pulumi.Input[str] created_at: The ISO8601 date/time that this integration was activated at in UTC.
        :param pulumi.Input[str] issues_url: The URL to view an issue in the external issue tracker. Must contain :id.
        :param pulumi.Input[str] project: The ID or full path of the project for the custom issue tracker.
        :param pulumi.Input[str] project_url: The URL to the project in the external issue tracker.
        :param pulumi.Input[str] slug: The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        :param pulumi.Input[str] updated_at: The ISO8601 date/time that this integration was last updated at in UTC.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if issues_url is not None:
            pulumi.set(__self__, "issues_url", issues_url)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if project_url is not None:
            pulumi.set(__self__, "project_url", project_url)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
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
    @pulumi.getter(name="issuesUrl")
    def issues_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL to view an issue in the external issue tracker. Must contain :id.
        """
        return pulumi.get(self, "issues_url")

    @issues_url.setter
    def issues_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issues_url", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or full path of the project for the custom issue tracker.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="projectUrl")
    def project_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL to the project in the external issue tracker.
        """
        return pulumi.get(self, "project_url")

    @project_url.setter
    def project_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_url", value)

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
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO8601 date/time that this integration was last updated at in UTC.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class ServiceCustomIssueTracker(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 issues_url: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 project_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `ServiceCustomIssueTracker` resource allows to manage the lifecycle of a project integration with Custom Issue Tracker.

        > This resource is deprecated. use `IntegrationCustomIssueTracker`instead!

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#custom-issue-tracker)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        awesome_project = gitlab.Project("awesomeProject",
            description="My awesome project.",
            visibility_level="public")
        tracker = gitlab.ServiceCustomIssueTracker("tracker",
            project=awesome_project.id,
            project_url="https://customtracker.com/issues",
            issues_url="https://customtracker.com/TEST-:id")
        ```

        ## Import

        You can import a gitlab_service_custom_issue_tracker state using the project ID, e.g.

        ```sh
        $ pulumi import gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker tracker 1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] issues_url: The URL to view an issue in the external issue tracker. Must contain :id.
        :param pulumi.Input[str] project: The ID or full path of the project for the custom issue tracker.
        :param pulumi.Input[str] project_url: The URL to the project in the external issue tracker.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceCustomIssueTrackerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ServiceCustomIssueTracker` resource allows to manage the lifecycle of a project integration with Custom Issue Tracker.

        > This resource is deprecated. use `IntegrationCustomIssueTracker`instead!

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#custom-issue-tracker)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        awesome_project = gitlab.Project("awesomeProject",
            description="My awesome project.",
            visibility_level="public")
        tracker = gitlab.ServiceCustomIssueTracker("tracker",
            project=awesome_project.id,
            project_url="https://customtracker.com/issues",
            issues_url="https://customtracker.com/TEST-:id")
        ```

        ## Import

        You can import a gitlab_service_custom_issue_tracker state using the project ID, e.g.

        ```sh
        $ pulumi import gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker tracker 1
        ```

        :param str resource_name: The name of the resource.
        :param ServiceCustomIssueTrackerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceCustomIssueTrackerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 issues_url: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 project_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceCustomIssueTrackerArgs.__new__(ServiceCustomIssueTrackerArgs)

            if issues_url is None and not opts.urn:
                raise TypeError("Missing required property 'issues_url'")
            __props__.__dict__["issues_url"] = issues_url
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if project_url is None and not opts.urn:
                raise TypeError("Missing required property 'project_url'")
            __props__.__dict__["project_url"] = project_url
            __props__.__dict__["active"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["slug"] = None
            __props__.__dict__["updated_at"] = None
        super(ServiceCustomIssueTracker, __self__).__init__(
            'gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            issues_url: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            project_url: Optional[pulumi.Input[str]] = None,
            slug: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'ServiceCustomIssueTracker':
        """
        Get an existing ServiceCustomIssueTracker resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Whether the integration is active.
        :param pulumi.Input[str] created_at: The ISO8601 date/time that this integration was activated at in UTC.
        :param pulumi.Input[str] issues_url: The URL to view an issue in the external issue tracker. Must contain :id.
        :param pulumi.Input[str] project: The ID or full path of the project for the custom issue tracker.
        :param pulumi.Input[str] project_url: The URL to the project in the external issue tracker.
        :param pulumi.Input[str] slug: The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        :param pulumi.Input[str] updated_at: The ISO8601 date/time that this integration was last updated at in UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceCustomIssueTrackerState.__new__(_ServiceCustomIssueTrackerState)

        __props__.__dict__["active"] = active
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["issues_url"] = issues_url
        __props__.__dict__["project"] = project
        __props__.__dict__["project_url"] = project_url
        __props__.__dict__["slug"] = slug
        __props__.__dict__["updated_at"] = updated_at
        return ServiceCustomIssueTracker(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[bool]:
        """
        Whether the integration is active.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The ISO8601 date/time that this integration was activated at in UTC.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="issuesUrl")
    def issues_url(self) -> pulumi.Output[str]:
        """
        The URL to view an issue in the external issue tracker. Must contain :id.
        """
        return pulumi.get(self, "issues_url")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or full path of the project for the custom issue tracker.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="projectUrl")
    def project_url(self) -> pulumi.Output[str]:
        """
        The URL to the project in the external issue tracker.
        """
        return pulumi.get(self, "project_url")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        """
        The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        """
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The ISO8601 date/time that this integration was last updated at in UTC.
        """
        return pulumi.get(self, "updated_at")

