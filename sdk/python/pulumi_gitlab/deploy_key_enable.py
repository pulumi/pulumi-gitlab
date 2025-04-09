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

__all__ = ['DeployKeyEnableArgs', 'DeployKeyEnable']

@pulumi.input_type
class DeployKeyEnableArgs:
    def __init__(__self__, *,
                 key_id: pulumi.Input[builtins.str],
                 project: pulumi.Input[builtins.str],
                 can_push: Optional[pulumi.Input[builtins.bool]] = None,
                 key: Optional[pulumi.Input[builtins.str]] = None,
                 title: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DeployKeyEnable resource.
        :param pulumi.Input[builtins.str] key_id: The Gitlab key id for the pre-existing deploy key
        :param pulumi.Input[builtins.str] project: The name or id of the project to add the deploy key to.
        :param pulumi.Input[builtins.bool] can_push: Can deploy key push to the project's repository.
        :param pulumi.Input[builtins.str] key: Deploy key.
        :param pulumi.Input[builtins.str] title: Deploy key's title.
        """
        pulumi.set(__self__, "key_id", key_id)
        pulumi.set(__self__, "project", project)
        if can_push is not None:
            pulumi.set(__self__, "can_push", can_push)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Input[builtins.str]:
        """
        The Gitlab key id for the pre-existing deploy key
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[builtins.str]:
        """
        The name or id of the project to add the deploy key to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="canPush")
    def can_push(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Can deploy key push to the project's repository.
        """
        return pulumi.get(self, "can_push")

    @can_push.setter
    def can_push(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "can_push", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Deploy key.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Deploy key's title.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "title", value)


@pulumi.input_type
class _DeployKeyEnableState:
    def __init__(__self__, *,
                 can_push: Optional[pulumi.Input[builtins.bool]] = None,
                 key: Optional[pulumi.Input[builtins.str]] = None,
                 key_id: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 title: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DeployKeyEnable resources.
        :param pulumi.Input[builtins.bool] can_push: Can deploy key push to the project's repository.
        :param pulumi.Input[builtins.str] key: Deploy key.
        :param pulumi.Input[builtins.str] key_id: The Gitlab key id for the pre-existing deploy key
        :param pulumi.Input[builtins.str] project: The name or id of the project to add the deploy key to.
        :param pulumi.Input[builtins.str] title: Deploy key's title.
        """
        if can_push is not None:
            pulumi.set(__self__, "can_push", can_push)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter(name="canPush")
    def can_push(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Can deploy key push to the project's repository.
        """
        return pulumi.get(self, "can_push")

    @can_push.setter
    def can_push(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "can_push", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Deploy key.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Gitlab key id for the pre-existing deploy key
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name or id of the project to add the deploy key to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Deploy key's title.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "title", value)


class DeployKeyEnable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_push: Optional[pulumi.Input[builtins.bool]] = None,
                 key: Optional[pulumi.Input[builtins.str]] = None,
                 key_id: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 title: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `DeployKeyEnable` resource allows to enable an already existing deploy key (see `DeployKey resource`) for a specific project.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/deploy_keys/#enable-a-deploy-key)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # A repo to host the deployment key
        parent = gitlab.Project("parent", name="parent_project")
        # A second repo to use the deployment key from the parent project
        foo = gitlab.Project("foo", name="foo_project")
        # Upload a deployment key for the parent repo
        parent_deploy_key = gitlab.DeployKey("parent",
            project=parent.id,
            title="Example deploy key",
            key="ssh-ed25519 AAAA...")
        # Enable the deployment key on the second repo
        foo_deploy_key_enable = gitlab.DeployKeyEnable("foo",
            project=foo.id,
            key_id=parent_deploy_key.deploy_key_id)
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_deploy_key_enable`. For example:

        terraform

        import {

          to = gitlab_deploy_key_enable.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        GitLab enabled deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.

        `project_id` can be whatever the [get single project api][get_single_project] takes for

        its `:id` value, so for example:

        ```sh
        $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example 12345:67890
        ```

        ```sh
        $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example richardc/example:67890
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] can_push: Can deploy key push to the project's repository.
        :param pulumi.Input[builtins.str] key: Deploy key.
        :param pulumi.Input[builtins.str] key_id: The Gitlab key id for the pre-existing deploy key
        :param pulumi.Input[builtins.str] project: The name or id of the project to add the deploy key to.
        :param pulumi.Input[builtins.str] title: Deploy key's title.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeployKeyEnableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `DeployKeyEnable` resource allows to enable an already existing deploy key (see `DeployKey resource`) for a specific project.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/deploy_keys/#enable-a-deploy-key)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # A repo to host the deployment key
        parent = gitlab.Project("parent", name="parent_project")
        # A second repo to use the deployment key from the parent project
        foo = gitlab.Project("foo", name="foo_project")
        # Upload a deployment key for the parent repo
        parent_deploy_key = gitlab.DeployKey("parent",
            project=parent.id,
            title="Example deploy key",
            key="ssh-ed25519 AAAA...")
        # Enable the deployment key on the second repo
        foo_deploy_key_enable = gitlab.DeployKeyEnable("foo",
            project=foo.id,
            key_id=parent_deploy_key.deploy_key_id)
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_deploy_key_enable`. For example:

        terraform

        import {

          to = gitlab_deploy_key_enable.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        GitLab enabled deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.

        `project_id` can be whatever the [get single project api][get_single_project] takes for

        its `:id` value, so for example:

        ```sh
        $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example 12345:67890
        ```

        ```sh
        $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example richardc/example:67890
        ```

        :param str resource_name: The name of the resource.
        :param DeployKeyEnableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeployKeyEnableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_push: Optional[pulumi.Input[builtins.bool]] = None,
                 key: Optional[pulumi.Input[builtins.str]] = None,
                 key_id: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 title: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeployKeyEnableArgs.__new__(DeployKeyEnableArgs)

            __props__.__dict__["can_push"] = can_push
            __props__.__dict__["key"] = key
            if key_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_id'")
            __props__.__dict__["key_id"] = key_id
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["title"] = title
        super(DeployKeyEnable, __self__).__init__(
            'gitlab:index/deployKeyEnable:DeployKeyEnable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            can_push: Optional[pulumi.Input[builtins.bool]] = None,
            key: Optional[pulumi.Input[builtins.str]] = None,
            key_id: Optional[pulumi.Input[builtins.str]] = None,
            project: Optional[pulumi.Input[builtins.str]] = None,
            title: Optional[pulumi.Input[builtins.str]] = None) -> 'DeployKeyEnable':
        """
        Get an existing DeployKeyEnable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] can_push: Can deploy key push to the project's repository.
        :param pulumi.Input[builtins.str] key: Deploy key.
        :param pulumi.Input[builtins.str] key_id: The Gitlab key id for the pre-existing deploy key
        :param pulumi.Input[builtins.str] project: The name or id of the project to add the deploy key to.
        :param pulumi.Input[builtins.str] title: Deploy key's title.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeployKeyEnableState.__new__(_DeployKeyEnableState)

        __props__.__dict__["can_push"] = can_push
        __props__.__dict__["key"] = key
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["project"] = project
        __props__.__dict__["title"] = title
        return DeployKeyEnable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canPush")
    def can_push(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Can deploy key push to the project's repository.
        """
        return pulumi.get(self, "can_push")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[builtins.str]:
        """
        Deploy key.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[builtins.str]:
        """
        The Gitlab key id for the pre-existing deploy key
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[builtins.str]:
        """
        The name or id of the project to add the deploy key to.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[builtins.str]:
        """
        Deploy key's title.
        """
        return pulumi.get(self, "title")

