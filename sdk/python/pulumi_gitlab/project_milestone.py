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

__all__ = ['ProjectMilestoneArgs', 'ProjectMilestone']

@pulumi.input_type
class ProjectMilestoneArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 due_date: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectMilestone resource.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[str] title: The title of a milestone.
        :param pulumi.Input[str] description: The description of the milestone.
        :param pulumi.Input[str] due_date: The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        :param pulumi.Input[str] start_date: The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        :param pulumi.Input[str] state: The state of the milestone. Valid values are: `active`, `closed`.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if due_date is not None:
            pulumi.set(__self__, "due_date", due_date)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or URL-encoded path of the project owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        The title of a milestone.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the milestone.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dueDate")
    def due_date(self) -> Optional[pulumi.Input[str]]:
        """
        The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        """
        return pulumi.get(self, "due_date")

    @due_date.setter
    def due_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "due_date", value)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[str]]:
        """
        The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the milestone. Valid values are: `active`, `closed`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


@pulumi.input_type
class _ProjectMilestoneState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 due_date: Optional[pulumi.Input[str]] = None,
                 expired: Optional[pulumi.Input[bool]] = None,
                 iid: Optional[pulumi.Input[int]] = None,
                 milestone_id: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 web_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectMilestone resources.
        :param pulumi.Input[str] created_at: The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        :param pulumi.Input[str] description: The description of the milestone.
        :param pulumi.Input[str] due_date: The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        :param pulumi.Input[bool] expired: Bool, true if milestone expired.
        :param pulumi.Input[int] iid: The ID of the project's milestone.
        :param pulumi.Input[int] milestone_id: The instance-wide ID of the project’s milestone.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[int] project_id: The project ID of milestone.
        :param pulumi.Input[str] start_date: The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        :param pulumi.Input[str] state: The state of the milestone. Valid values are: `active`, `closed`.
        :param pulumi.Input[str] title: The title of a milestone.
        :param pulumi.Input[str] updated_at: The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        :param pulumi.Input[str] web_url: The web URL of the milestone.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if due_date is not None:
            pulumi.set(__self__, "due_date", due_date)
        if expired is not None:
            pulumi.set(__self__, "expired", expired)
        if iid is not None:
            pulumi.set(__self__, "iid", iid)
        if milestone_id is not None:
            pulumi.set(__self__, "milestone_id", milestone_id)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if web_url is not None:
            pulumi.set(__self__, "web_url", web_url)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the milestone.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dueDate")
    def due_date(self) -> Optional[pulumi.Input[str]]:
        """
        The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        """
        return pulumi.get(self, "due_date")

    @due_date.setter
    def due_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "due_date", value)

    @property
    @pulumi.getter
    def expired(self) -> Optional[pulumi.Input[bool]]:
        """
        Bool, true if milestone expired.
        """
        return pulumi.get(self, "expired")

    @expired.setter
    def expired(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "expired", value)

    @property
    @pulumi.getter
    def iid(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the project's milestone.
        """
        return pulumi.get(self, "iid")

    @iid.setter
    def iid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iid", value)

    @property
    @pulumi.getter(name="milestoneId")
    def milestone_id(self) -> Optional[pulumi.Input[int]]:
        """
        The instance-wide ID of the project’s milestone.
        """
        return pulumi.get(self, "milestone_id")

    @milestone_id.setter
    def milestone_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "milestone_id", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or URL-encoded path of the project owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        The project ID of milestone.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[str]]:
        """
        The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the milestone. Valid values are: `active`, `closed`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        The title of a milestone.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> Optional[pulumi.Input[str]]:
        """
        The web URL of the milestone.
        """
        return pulumi.get(self, "web_url")

    @web_url.setter
    def web_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "web_url", value)


class ProjectMilestone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 due_date: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `ProjectMilestone` resource allows to manage the lifecycle of a project milestone.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # Create a project for the milestone to use
        example = gitlab.Project("example",
            name="example",
            description="An example project",
            namespace_id=example_gitlab_group["id"])
        example_project_milestone = gitlab.ProjectMilestone("example",
            project=example.id,
            title="example")
        ```

        ## Import

        Gitlab project milestone can be imported with a key composed of `<project>:<milestone_id>`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectMilestone:ProjectMilestone example "12345:11"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the milestone.
        :param pulumi.Input[str] due_date: The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[str] start_date: The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        :param pulumi.Input[str] state: The state of the milestone. Valid values are: `active`, `closed`.
        :param pulumi.Input[str] title: The title of a milestone.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectMilestoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectMilestone` resource allows to manage the lifecycle of a project milestone.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # Create a project for the milestone to use
        example = gitlab.Project("example",
            name="example",
            description="An example project",
            namespace_id=example_gitlab_group["id"])
        example_project_milestone = gitlab.ProjectMilestone("example",
            project=example.id,
            title="example")
        ```

        ## Import

        Gitlab project milestone can be imported with a key composed of `<project>:<milestone_id>`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectMilestone:ProjectMilestone example "12345:11"
        ```

        :param str resource_name: The name of the resource.
        :param ProjectMilestoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectMilestoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 due_date: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectMilestoneArgs.__new__(ProjectMilestoneArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["due_date"] = due_date
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["start_date"] = start_date
            __props__.__dict__["state"] = state
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__.__dict__["title"] = title
            __props__.__dict__["created_at"] = None
            __props__.__dict__["expired"] = None
            __props__.__dict__["iid"] = None
            __props__.__dict__["milestone_id"] = None
            __props__.__dict__["project_id"] = None
            __props__.__dict__["updated_at"] = None
            __props__.__dict__["web_url"] = None
        super(ProjectMilestone, __self__).__init__(
            'gitlab:index/projectMilestone:ProjectMilestone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            due_date: Optional[pulumi.Input[str]] = None,
            expired: Optional[pulumi.Input[bool]] = None,
            iid: Optional[pulumi.Input[int]] = None,
            milestone_id: Optional[pulumi.Input[int]] = None,
            project: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            start_date: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            web_url: Optional[pulumi.Input[str]] = None) -> 'ProjectMilestone':
        """
        Get an existing ProjectMilestone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        :param pulumi.Input[str] description: The description of the milestone.
        :param pulumi.Input[str] due_date: The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        :param pulumi.Input[bool] expired: Bool, true if milestone expired.
        :param pulumi.Input[int] iid: The ID of the project's milestone.
        :param pulumi.Input[int] milestone_id: The instance-wide ID of the project’s milestone.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[int] project_id: The project ID of milestone.
        :param pulumi.Input[str] start_date: The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        :param pulumi.Input[str] state: The state of the milestone. Valid values are: `active`, `closed`.
        :param pulumi.Input[str] title: The title of a milestone.
        :param pulumi.Input[str] updated_at: The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        :param pulumi.Input[str] web_url: The web URL of the milestone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectMilestoneState.__new__(_ProjectMilestoneState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["due_date"] = due_date
        __props__.__dict__["expired"] = expired
        __props__.__dict__["iid"] = iid
        __props__.__dict__["milestone_id"] = milestone_id
        __props__.__dict__["project"] = project
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["start_date"] = start_date
        __props__.__dict__["state"] = state
        __props__.__dict__["title"] = title
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["web_url"] = web_url
        return ProjectMilestone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the milestone.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dueDate")
    def due_date(self) -> pulumi.Output[Optional[str]]:
        """
        The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        """
        return pulumi.get(self, "due_date")

    @property
    @pulumi.getter
    def expired(self) -> pulumi.Output[bool]:
        """
        Bool, true if milestone expired.
        """
        return pulumi.get(self, "expired")

    @property
    @pulumi.getter
    def iid(self) -> pulumi.Output[int]:
        """
        The ID of the project's milestone.
        """
        return pulumi.get(self, "iid")

    @property
    @pulumi.getter(name="milestoneId")
    def milestone_id(self) -> pulumi.Output[int]:
        """
        The instance-wide ID of the project’s milestone.
        """
        return pulumi.get(self, "milestone_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of the project owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        The project ID of milestone.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> pulumi.Output[Optional[str]]:
        """
        The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        The state of the milestone. Valid values are: `active`, `closed`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        The title of a milestone.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> pulumi.Output[str]:
        """
        The web URL of the milestone.
        """
        return pulumi.get(self, "web_url")

