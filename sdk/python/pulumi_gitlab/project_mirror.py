# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectMirrorArgs', 'ProjectMirror']

@pulumi.input_type
class ProjectMirrorArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 url: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 keep_divergent_refs: Optional[pulumi.Input[bool]] = None,
                 only_protected_branches: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ProjectMirror resource.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] url: The URL of the remote repository to be mirrored.
        :param pulumi.Input[bool] enabled: Determines if the mirror is enabled.
        :param pulumi.Input[bool] keep_divergent_refs: Determines if divergent refs are skipped.
        :param pulumi.Input[bool] only_protected_branches: Determines if only protected branches are mirrored.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "url", url)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if keep_divergent_refs is not None:
            pulumi.set(__self__, "keep_divergent_refs", keep_divergent_refs)
        if only_protected_branches is not None:
            pulumi.set(__self__, "only_protected_branches", only_protected_branches)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The URL of the remote repository to be mirrored.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if the mirror is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="keepDivergentRefs")
    def keep_divergent_refs(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if divergent refs are skipped.
        """
        return pulumi.get(self, "keep_divergent_refs")

    @keep_divergent_refs.setter
    def keep_divergent_refs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_divergent_refs", value)

    @property
    @pulumi.getter(name="onlyProtectedBranches")
    def only_protected_branches(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if only protected branches are mirrored.
        """
        return pulumi.get(self, "only_protected_branches")

    @only_protected_branches.setter
    def only_protected_branches(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "only_protected_branches", value)


@pulumi.input_type
class _ProjectMirrorState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 keep_divergent_refs: Optional[pulumi.Input[bool]] = None,
                 mirror_id: Optional[pulumi.Input[int]] = None,
                 only_protected_branches: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectMirror resources.
        :param pulumi.Input[bool] enabled: Determines if the mirror is enabled.
        :param pulumi.Input[bool] keep_divergent_refs: Determines if divergent refs are skipped.
        :param pulumi.Input[int] mirror_id: Mirror ID.
        :param pulumi.Input[bool] only_protected_branches: Determines if only protected branches are mirrored.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] url: The URL of the remote repository to be mirrored.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if keep_divergent_refs is not None:
            pulumi.set(__self__, "keep_divergent_refs", keep_divergent_refs)
        if mirror_id is not None:
            pulumi.set(__self__, "mirror_id", mirror_id)
        if only_protected_branches is not None:
            pulumi.set(__self__, "only_protected_branches", only_protected_branches)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if the mirror is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="keepDivergentRefs")
    def keep_divergent_refs(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if divergent refs are skipped.
        """
        return pulumi.get(self, "keep_divergent_refs")

    @keep_divergent_refs.setter
    def keep_divergent_refs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_divergent_refs", value)

    @property
    @pulumi.getter(name="mirrorId")
    def mirror_id(self) -> Optional[pulumi.Input[int]]:
        """
        Mirror ID.
        """
        return pulumi.get(self, "mirror_id")

    @mirror_id.setter
    def mirror_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mirror_id", value)

    @property
    @pulumi.getter(name="onlyProtectedBranches")
    def only_protected_branches(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if only protected branches are mirrored.
        """
        return pulumi.get(self, "only_protected_branches")

    @only_protected_branches.setter
    def only_protected_branches(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "only_protected_branches", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the remote repository to be mirrored.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ProjectMirror(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 keep_divergent_refs: Optional[pulumi.Input[bool]] = None,
                 only_protected_branches: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `ProjectMirror` resource allows to manage the lifecycle of a project mirror.

        This is for *pushing* changes to a remote repository. *Pull Mirroring* can be configured using a combination of the
        import_url, mirror, and mirror_trigger_builds properties on the Project resource.

        > **Warning** By default, the provider sets the `keep_divergent_refs` argument to `True`.
           If you manually set `keep_divergent_refs` to `False`, GitLab mirroring removes branches in the target that aren't in the source.
           This action can result in unexpected branch deletions.

        > **Destroy Behavior** GitLab 14.10 introduced an API endpoint to delete a project mirror.
           Therefore, for GitLab 14.10 and newer the project mirror will be destroyed when the resource is destroyed.
           For older versions, the mirror will be disabled and the resource will be destroyed.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/remote_mirrors.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.ProjectMirror("foo",
            project="1",
            url="https://username:password@github.com/org/repository.git")
        ```

        ## Import

        GitLab project mirror can be imported using an id made up of `project_id:mirror_id`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectMirror:ProjectMirror foo "12345:1337"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Determines if the mirror is enabled.
        :param pulumi.Input[bool] keep_divergent_refs: Determines if divergent refs are skipped.
        :param pulumi.Input[bool] only_protected_branches: Determines if only protected branches are mirrored.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] url: The URL of the remote repository to be mirrored.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectMirrorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ProjectMirror` resource allows to manage the lifecycle of a project mirror.

        This is for *pushing* changes to a remote repository. *Pull Mirroring* can be configured using a combination of the
        import_url, mirror, and mirror_trigger_builds properties on the Project resource.

        > **Warning** By default, the provider sets the `keep_divergent_refs` argument to `True`.
           If you manually set `keep_divergent_refs` to `False`, GitLab mirroring removes branches in the target that aren't in the source.
           This action can result in unexpected branch deletions.

        > **Destroy Behavior** GitLab 14.10 introduced an API endpoint to delete a project mirror.
           Therefore, for GitLab 14.10 and newer the project mirror will be destroyed when the resource is destroyed.
           For older versions, the mirror will be disabled and the resource will be destroyed.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/remote_mirrors.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        foo = gitlab.ProjectMirror("foo",
            project="1",
            url="https://username:password@github.com/org/repository.git")
        ```

        ## Import

        GitLab project mirror can be imported using an id made up of `project_id:mirror_id`, e.g.

        ```sh
        $ pulumi import gitlab:index/projectMirror:ProjectMirror foo "12345:1337"
        ```

        :param str resource_name: The name of the resource.
        :param ProjectMirrorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectMirrorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 keep_divergent_refs: Optional[pulumi.Input[bool]] = None,
                 only_protected_branches: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectMirrorArgs.__new__(ProjectMirrorArgs)

            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["keep_divergent_refs"] = keep_divergent_refs
            __props__.__dict__["only_protected_branches"] = only_protected_branches
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = None if url is None else pulumi.Output.secret(url)
            __props__.__dict__["mirror_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["url"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ProjectMirror, __self__).__init__(
            'gitlab:index/projectMirror:ProjectMirror',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            keep_divergent_refs: Optional[pulumi.Input[bool]] = None,
            mirror_id: Optional[pulumi.Input[int]] = None,
            only_protected_branches: Optional[pulumi.Input[bool]] = None,
            project: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ProjectMirror':
        """
        Get an existing ProjectMirror resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Determines if the mirror is enabled.
        :param pulumi.Input[bool] keep_divergent_refs: Determines if divergent refs are skipped.
        :param pulumi.Input[int] mirror_id: Mirror ID.
        :param pulumi.Input[bool] only_protected_branches: Determines if only protected branches are mirrored.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] url: The URL of the remote repository to be mirrored.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectMirrorState.__new__(_ProjectMirrorState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["keep_divergent_refs"] = keep_divergent_refs
        __props__.__dict__["mirror_id"] = mirror_id
        __props__.__dict__["only_protected_branches"] = only_protected_branches
        __props__.__dict__["project"] = project
        __props__.__dict__["url"] = url
        return ProjectMirror(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if the mirror is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="keepDivergentRefs")
    def keep_divergent_refs(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if divergent refs are skipped.
        """
        return pulumi.get(self, "keep_divergent_refs")

    @property
    @pulumi.getter(name="mirrorId")
    def mirror_id(self) -> pulumi.Output[int]:
        """
        Mirror ID.
        """
        return pulumi.get(self, "mirror_id")

    @property
    @pulumi.getter(name="onlyProtectedBranches")
    def only_protected_branches(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if only protected branches are mirrored.
        """
        return pulumi.get(self, "only_protected_branches")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the remote repository to be mirrored.
        """
        return pulumi.get(self, "url")

