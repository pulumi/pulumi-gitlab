# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ProjectTagArgs', 'ProjectTag']

@pulumi.input_type
class ProjectTagArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 ref: pulumi.Input[str],
                 message: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectTag resource.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[str] ref: Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
        :param pulumi.Input[str] message: The message of the annotated tag.
        :param pulumi.Input[str] name: The name of a tag.
        """
        ProjectTagArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            project=project,
            ref=ref,
            message=message,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             project: Optional[pulumi.Input[str]] = None,
             ref: Optional[pulumi.Input[str]] = None,
             message: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if project is None:
            raise TypeError("Missing 'project' argument")
        if ref is None:
            raise TypeError("Missing 'ref' argument")

        _setter("project", project)
        _setter("ref", ref)
        if message is not None:
            _setter("message", message)
        if name is not None:
            _setter("name", name)

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
    def ref(self) -> pulumi.Input[str]:
        """
        Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "ref")

    @ref.setter
    def ref(self, value: pulumi.Input[str]):
        pulumi.set(self, "ref", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        The message of the annotated tag.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a tag.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ProjectTagState:
    def __init__(__self__, *,
                 commits: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectTagCommitArgs']]]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 protected: Optional[pulumi.Input[bool]] = None,
                 ref: Optional[pulumi.Input[str]] = None,
                 releases: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectTagReleaseArgs']]]] = None,
                 target: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectTag resources.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectTagCommitArgs']]] commits: The commit associated with the tag.
        :param pulumi.Input[str] message: The message of the annotated tag.
        :param pulumi.Input[str] name: The name of a tag.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[bool] protected: Bool, true if tag has tag protection.
        :param pulumi.Input[str] ref: Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectTagReleaseArgs']]] releases: The release associated with the tag.
        :param pulumi.Input[str] target: The unique id assigned to the commit by Gitlab.
        """
        _ProjectTagState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            commits=commits,
            message=message,
            name=name,
            project=project,
            protected=protected,
            ref=ref,
            releases=releases,
            target=target,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             commits: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectTagCommitArgs']]]] = None,
             message: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             protected: Optional[pulumi.Input[bool]] = None,
             ref: Optional[pulumi.Input[str]] = None,
             releases: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectTagReleaseArgs']]]] = None,
             target: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if commits is not None:
            _setter("commits", commits)
        if message is not None:
            _setter("message", message)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if protected is not None:
            _setter("protected", protected)
        if ref is not None:
            _setter("ref", ref)
        if releases is not None:
            _setter("releases", releases)
        if target is not None:
            _setter("target", target)

    @property
    @pulumi.getter
    def commits(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectTagCommitArgs']]]]:
        """
        The commit associated with the tag.
        """
        return pulumi.get(self, "commits")

    @commits.setter
    def commits(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectTagCommitArgs']]]]):
        pulumi.set(self, "commits", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        The message of the annotated tag.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a tag.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
    @pulumi.getter
    def protected(self) -> Optional[pulumi.Input[bool]]:
        """
        Bool, true if tag has tag protection.
        """
        return pulumi.get(self, "protected")

    @protected.setter
    def protected(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "protected", value)

    @property
    @pulumi.getter
    def ref(self) -> Optional[pulumi.Input[str]]:
        """
        Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "ref")

    @ref.setter
    def ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ref", value)

    @property
    @pulumi.getter
    def releases(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectTagReleaseArgs']]]]:
        """
        The release associated with the tag.
        """
        return pulumi.get(self, "releases")

    @releases.setter
    def releases(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectTagReleaseArgs']]]]):
        pulumi.set(self, "releases", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        """
        The unique id assigned to the commit by Gitlab.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)


class ProjectTag(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ref: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `ProjectTag` resource allows to manage the lifecycle of a tag in a project.

        **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/tags.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # Create a project for the tag to use
        example_project = gitlab.Project("exampleProject",
            description="An example project",
            namespace_id=gitlab_group["example"]["id"])
        example_project_tag = gitlab.ProjectTag("exampleProjectTag",
            ref="main",
            project=example_project.id)
        ```

        ## Import

        Gitlab project tags can be imported with a key composed of `<project_id>:<tag_name>`, e.g.

        ```sh
         $ pulumi import gitlab:index/projectTag:ProjectTag example "12345:develop"
        ```

         NOTEthe `ref` attribute won't be available for imported `gitlab_project_tag` resources.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] message: The message of the annotated tag.
        :param pulumi.Input[str] name: The name of a tag.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[str] ref: Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectTagArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectTag` resource allows to manage the lifecycle of a tag in a project.

        **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/tags.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # Create a project for the tag to use
        example_project = gitlab.Project("exampleProject",
            description="An example project",
            namespace_id=gitlab_group["example"]["id"])
        example_project_tag = gitlab.ProjectTag("exampleProjectTag",
            ref="main",
            project=example_project.id)
        ```

        ## Import

        Gitlab project tags can be imported with a key composed of `<project_id>:<tag_name>`, e.g.

        ```sh
         $ pulumi import gitlab:index/projectTag:ProjectTag example "12345:develop"
        ```

         NOTEthe `ref` attribute won't be available for imported `gitlab_project_tag` resources.

        :param str resource_name: The name of the resource.
        :param ProjectTagArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectTagArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ProjectTagArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ref: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectTagArgs.__new__(ProjectTagArgs)

            __props__.__dict__["message"] = message
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if ref is None and not opts.urn:
                raise TypeError("Missing required property 'ref'")
            __props__.__dict__["ref"] = ref
            __props__.__dict__["commits"] = None
            __props__.__dict__["protected"] = None
            __props__.__dict__["releases"] = None
            __props__.__dict__["target"] = None
        super(ProjectTag, __self__).__init__(
            'gitlab:index/projectTag:ProjectTag',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            commits: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectTagCommitArgs']]]]] = None,
            message: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            protected: Optional[pulumi.Input[bool]] = None,
            ref: Optional[pulumi.Input[str]] = None,
            releases: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectTagReleaseArgs']]]]] = None,
            target: Optional[pulumi.Input[str]] = None) -> 'ProjectTag':
        """
        Get an existing ProjectTag resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectTagCommitArgs']]]] commits: The commit associated with the tag.
        :param pulumi.Input[str] message: The message of the annotated tag.
        :param pulumi.Input[str] name: The name of a tag.
        :param pulumi.Input[str] project: The ID or URL-encoded path of the project owned by the authenticated user.
        :param pulumi.Input[bool] protected: Bool, true if tag has tag protection.
        :param pulumi.Input[str] ref: Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectTagReleaseArgs']]]] releases: The release associated with the tag.
        :param pulumi.Input[str] target: The unique id assigned to the commit by Gitlab.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectTagState.__new__(_ProjectTagState)

        __props__.__dict__["commits"] = commits
        __props__.__dict__["message"] = message
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["protected"] = protected
        __props__.__dict__["ref"] = ref
        __props__.__dict__["releases"] = releases
        __props__.__dict__["target"] = target
        return ProjectTag(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def commits(self) -> pulumi.Output[Sequence['outputs.ProjectTagCommit']]:
        """
        The commit associated with the tag.
        """
        return pulumi.get(self, "commits")

    @property
    @pulumi.getter
    def message(self) -> pulumi.Output[Optional[str]]:
        """
        The message of the annotated tag.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of a tag.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or URL-encoded path of the project owned by the authenticated user.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def protected(self) -> pulumi.Output[bool]:
        """
        Bool, true if tag has tag protection.
        """
        return pulumi.get(self, "protected")

    @property
    @pulumi.getter
    def ref(self) -> pulumi.Output[str]:
        """
        Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "ref")

    @property
    @pulumi.getter
    def releases(self) -> pulumi.Output[Sequence['outputs.ProjectTagRelease']]:
        """
        The release associated with the tag.
        """
        return pulumi.get(self, "releases")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[str]:
        """
        The unique id assigned to the commit by Gitlab.
        """
        return pulumi.get(self, "target")

