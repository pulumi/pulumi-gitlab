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

__all__ = ['BranchArgs', 'Branch']

@pulumi.input_type
class BranchArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 ref: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Branch resource.
        :param pulumi.Input[str] project: The ID or full path of the project which the branch is created against.
        :param pulumi.Input[str] ref: The ref which the branch is created from.
        :param pulumi.Input[str] name: The name for this branch.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "ref", ref)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or full path of the project which the branch is created against.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def ref(self) -> pulumi.Input[str]:
        """
        The ref which the branch is created from.
        """
        return pulumi.get(self, "ref")

    @ref.setter
    def ref(self, value: pulumi.Input[str]):
        pulumi.set(self, "ref", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for this branch.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _BranchState:
    def __init__(__self__, *,
                 can_push: Optional[pulumi.Input[bool]] = None,
                 commits: Optional[pulumi.Input[Sequence[pulumi.Input['BranchCommitArgs']]]] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 developer_can_merge: Optional[pulumi.Input[bool]] = None,
                 developer_can_push: Optional[pulumi.Input[bool]] = None,
                 merged: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 protected: Optional[pulumi.Input[bool]] = None,
                 ref: Optional[pulumi.Input[str]] = None,
                 web_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Branch resources.
        :param pulumi.Input[bool] can_push: Bool, true if you can push to the branch.
        :param pulumi.Input[Sequence[pulumi.Input['BranchCommitArgs']]] commits: The commit associated with the branch ref.
        :param pulumi.Input[bool] default: Bool, true if branch is the default branch for the project.
        :param pulumi.Input[bool] developer_can_merge: Bool, true if developer level access allows to merge branch.
        :param pulumi.Input[bool] developer_can_push: Bool, true if developer level access allows git push.
        :param pulumi.Input[bool] merged: Bool, true if the branch has been merged into it's parent.
        :param pulumi.Input[str] name: The name for this branch.
        :param pulumi.Input[str] project: The ID or full path of the project which the branch is created against.
        :param pulumi.Input[bool] protected: Bool, true if branch has branch protection.
        :param pulumi.Input[str] ref: The ref which the branch is created from.
        :param pulumi.Input[str] web_url: The url of the created branch (https).
        """
        if can_push is not None:
            pulumi.set(__self__, "can_push", can_push)
        if commits is not None:
            pulumi.set(__self__, "commits", commits)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if developer_can_merge is not None:
            pulumi.set(__self__, "developer_can_merge", developer_can_merge)
        if developer_can_push is not None:
            pulumi.set(__self__, "developer_can_push", developer_can_push)
        if merged is not None:
            pulumi.set(__self__, "merged", merged)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if protected is not None:
            pulumi.set(__self__, "protected", protected)
        if ref is not None:
            pulumi.set(__self__, "ref", ref)
        if web_url is not None:
            pulumi.set(__self__, "web_url", web_url)

    @property
    @pulumi.getter(name="canPush")
    def can_push(self) -> Optional[pulumi.Input[bool]]:
        """
        Bool, true if you can push to the branch.
        """
        return pulumi.get(self, "can_push")

    @can_push.setter
    def can_push(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_push", value)

    @property
    @pulumi.getter
    def commits(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BranchCommitArgs']]]]:
        """
        The commit associated with the branch ref.
        """
        return pulumi.get(self, "commits")

    @commits.setter
    def commits(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BranchCommitArgs']]]]):
        pulumi.set(self, "commits", value)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[bool]]:
        """
        Bool, true if branch is the default branch for the project.
        """
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default", value)

    @property
    @pulumi.getter(name="developerCanMerge")
    def developer_can_merge(self) -> Optional[pulumi.Input[bool]]:
        """
        Bool, true if developer level access allows to merge branch.
        """
        return pulumi.get(self, "developer_can_merge")

    @developer_can_merge.setter
    def developer_can_merge(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "developer_can_merge", value)

    @property
    @pulumi.getter(name="developerCanPush")
    def developer_can_push(self) -> Optional[pulumi.Input[bool]]:
        """
        Bool, true if developer level access allows git push.
        """
        return pulumi.get(self, "developer_can_push")

    @developer_can_push.setter
    def developer_can_push(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "developer_can_push", value)

    @property
    @pulumi.getter
    def merged(self) -> Optional[pulumi.Input[bool]]:
        """
        Bool, true if the branch has been merged into it's parent.
        """
        return pulumi.get(self, "merged")

    @merged.setter
    def merged(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "merged", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for this branch.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or full path of the project which the branch is created against.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def protected(self) -> Optional[pulumi.Input[bool]]:
        """
        Bool, true if branch has branch protection.
        """
        return pulumi.get(self, "protected")

    @protected.setter
    def protected(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "protected", value)

    @property
    @pulumi.getter
    def ref(self) -> Optional[pulumi.Input[str]]:
        """
        The ref which the branch is created from.
        """
        return pulumi.get(self, "ref")

    @ref.setter
    def ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ref", value)

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> Optional[pulumi.Input[str]]:
        """
        The url of the created branch (https).
        """
        return pulumi.get(self, "web_url")

    @web_url.setter
    def web_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "web_url", value)


class Branch(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ref: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `Branch` resource allows to manage the lifecycle of a repository branch.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/branches.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # Create a project for the branch to use
        example_project = gitlab.Project("exampleProject",
            description="An example project",
            namespace_id=gitlab_group["example"]["id"])
        example_branch = gitlab.Branch("exampleBranch",
            ref="main",
            project=example_project.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Gitlab branches can be imported with a key composed of `<project_id>:<branch_name>`, e.g.

        ```sh
        $ pulumi import gitlab:index/branch:Branch example "12345:develop"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name for this branch.
        :param pulumi.Input[str] project: The ID or full path of the project which the branch is created against.
        :param pulumi.Input[str] ref: The ref which the branch is created from.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BranchArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `Branch` resource allows to manage the lifecycle of a repository branch.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/branches.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # Create a project for the branch to use
        example_project = gitlab.Project("exampleProject",
            description="An example project",
            namespace_id=gitlab_group["example"]["id"])
        example_branch = gitlab.Branch("exampleBranch",
            ref="main",
            project=example_project.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Gitlab branches can be imported with a key composed of `<project_id>:<branch_name>`, e.g.

        ```sh
        $ pulumi import gitlab:index/branch:Branch example "12345:develop"
        ```

        :param str resource_name: The name of the resource.
        :param BranchArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BranchArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
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
            __props__ = BranchArgs.__new__(BranchArgs)

            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if ref is None and not opts.urn:
                raise TypeError("Missing required property 'ref'")
            __props__.__dict__["ref"] = ref
            __props__.__dict__["can_push"] = None
            __props__.__dict__["commits"] = None
            __props__.__dict__["default"] = None
            __props__.__dict__["developer_can_merge"] = None
            __props__.__dict__["developer_can_push"] = None
            __props__.__dict__["merged"] = None
            __props__.__dict__["protected"] = None
            __props__.__dict__["web_url"] = None
        super(Branch, __self__).__init__(
            'gitlab:index/branch:Branch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            can_push: Optional[pulumi.Input[bool]] = None,
            commits: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchCommitArgs']]]]] = None,
            default: Optional[pulumi.Input[bool]] = None,
            developer_can_merge: Optional[pulumi.Input[bool]] = None,
            developer_can_push: Optional[pulumi.Input[bool]] = None,
            merged: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            protected: Optional[pulumi.Input[bool]] = None,
            ref: Optional[pulumi.Input[str]] = None,
            web_url: Optional[pulumi.Input[str]] = None) -> 'Branch':
        """
        Get an existing Branch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_push: Bool, true if you can push to the branch.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchCommitArgs']]]] commits: The commit associated with the branch ref.
        :param pulumi.Input[bool] default: Bool, true if branch is the default branch for the project.
        :param pulumi.Input[bool] developer_can_merge: Bool, true if developer level access allows to merge branch.
        :param pulumi.Input[bool] developer_can_push: Bool, true if developer level access allows git push.
        :param pulumi.Input[bool] merged: Bool, true if the branch has been merged into it's parent.
        :param pulumi.Input[str] name: The name for this branch.
        :param pulumi.Input[str] project: The ID or full path of the project which the branch is created against.
        :param pulumi.Input[bool] protected: Bool, true if branch has branch protection.
        :param pulumi.Input[str] ref: The ref which the branch is created from.
        :param pulumi.Input[str] web_url: The url of the created branch (https).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BranchState.__new__(_BranchState)

        __props__.__dict__["can_push"] = can_push
        __props__.__dict__["commits"] = commits
        __props__.__dict__["default"] = default
        __props__.__dict__["developer_can_merge"] = developer_can_merge
        __props__.__dict__["developer_can_push"] = developer_can_push
        __props__.__dict__["merged"] = merged
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["protected"] = protected
        __props__.__dict__["ref"] = ref
        __props__.__dict__["web_url"] = web_url
        return Branch(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canPush")
    def can_push(self) -> pulumi.Output[bool]:
        """
        Bool, true if you can push to the branch.
        """
        return pulumi.get(self, "can_push")

    @property
    @pulumi.getter
    def commits(self) -> pulumi.Output[Sequence['outputs.BranchCommit']]:
        """
        The commit associated with the branch ref.
        """
        return pulumi.get(self, "commits")

    @property
    @pulumi.getter
    def default(self) -> pulumi.Output[bool]:
        """
        Bool, true if branch is the default branch for the project.
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter(name="developerCanMerge")
    def developer_can_merge(self) -> pulumi.Output[bool]:
        """
        Bool, true if developer level access allows to merge branch.
        """
        return pulumi.get(self, "developer_can_merge")

    @property
    @pulumi.getter(name="developerCanPush")
    def developer_can_push(self) -> pulumi.Output[bool]:
        """
        Bool, true if developer level access allows git push.
        """
        return pulumi.get(self, "developer_can_push")

    @property
    @pulumi.getter
    def merged(self) -> pulumi.Output[bool]:
        """
        Bool, true if the branch has been merged into it's parent.
        """
        return pulumi.get(self, "merged")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for this branch.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or full path of the project which the branch is created against.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def protected(self) -> pulumi.Output[bool]:
        """
        Bool, true if branch has branch protection.
        """
        return pulumi.get(self, "protected")

    @property
    @pulumi.getter
    def ref(self) -> pulumi.Output[str]:
        """
        The ref which the branch is created from.
        """
        return pulumi.get(self, "ref")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> pulumi.Output[str]:
        """
        The url of the created branch (https).
        """
        return pulumi.get(self, "web_url")

