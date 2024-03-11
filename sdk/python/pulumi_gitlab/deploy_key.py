# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DeployKeyArgs', 'DeployKey']

@pulumi.input_type
class DeployKeyArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 project: pulumi.Input[str],
                 title: pulumi.Input[str],
                 can_push: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a DeployKey resource.
        :param pulumi.Input[str] key: The public ssh key body.
        :param pulumi.Input[str] project: The name or id of the project to add the deploy key to.
        :param pulumi.Input[str] title: A title to describe the deploy key with.
        :param pulumi.Input[bool] can_push: Allow this deploy key to be used to push changes to the project. Defaults to `false`.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "title", title)
        if can_push is not None:
            pulumi.set(__self__, "can_push", can_push)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The public ssh key body.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The name or id of the project to add the deploy key to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        A title to describe the deploy key with.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter(name="canPush")
    def can_push(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow this deploy key to be used to push changes to the project. Defaults to `false`.
        """
        return pulumi.get(self, "can_push")

    @can_push.setter
    def can_push(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_push", value)


@pulumi.input_type
class _DeployKeyState:
    def __init__(__self__, *,
                 can_push: Optional[pulumi.Input[bool]] = None,
                 deploy_key_id: Optional[pulumi.Input[int]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DeployKey resources.
        :param pulumi.Input[bool] can_push: Allow this deploy key to be used to push changes to the project. Defaults to `false`.
        :param pulumi.Input[int] deploy_key_id: The id of the project deploy key.
        :param pulumi.Input[str] key: The public ssh key body.
        :param pulumi.Input[str] project: The name or id of the project to add the deploy key to.
        :param pulumi.Input[str] title: A title to describe the deploy key with.
        """
        if can_push is not None:
            pulumi.set(__self__, "can_push", can_push)
        if deploy_key_id is not None:
            pulumi.set(__self__, "deploy_key_id", deploy_key_id)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter(name="canPush")
    def can_push(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow this deploy key to be used to push changes to the project. Defaults to `false`.
        """
        return pulumi.get(self, "can_push")

    @can_push.setter
    def can_push(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_push", value)

    @property
    @pulumi.getter(name="deployKeyId")
    def deploy_key_id(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the project deploy key.
        """
        return pulumi.get(self, "deploy_key_id")

    @deploy_key_id.setter
    def deploy_key_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "deploy_key_id", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The public ssh key body.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The name or id of the project to add the deploy key to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        A title to describe the deploy key with.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


class DeployKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_push: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `DeployKey` resource allows to manage the lifecycle of a deploy key.

        > To enable an already existing deploy key for another project use the `gitlab_project_deploy_key` resource.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_keys.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.DeployKey("example",
            key="ssh-ed25519 AAAA...",
            project="example/deploying",
            title="Example deploy key")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitLab deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.

        `project_id` can be whatever the [get single project api][get_single_project] takes for

        its `:id` value, so for example:

        ```sh
        $ pulumi import gitlab:index/deployKey:DeployKey test 1:3
        ```

        ```sh
        $ pulumi import gitlab:index/deployKey:DeployKey test richardc/example:3
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_push: Allow this deploy key to be used to push changes to the project. Defaults to `false`.
        :param pulumi.Input[str] key: The public ssh key body.
        :param pulumi.Input[str] project: The name or id of the project to add the deploy key to.
        :param pulumi.Input[str] title: A title to describe the deploy key with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeployKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `DeployKey` resource allows to manage the lifecycle of a deploy key.

        > To enable an already existing deploy key for another project use the `gitlab_project_deploy_key` resource.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_keys.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.DeployKey("example",
            key="ssh-ed25519 AAAA...",
            project="example/deploying",
            title="Example deploy key")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitLab deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.

        `project_id` can be whatever the [get single project api][get_single_project] takes for

        its `:id` value, so for example:

        ```sh
        $ pulumi import gitlab:index/deployKey:DeployKey test 1:3
        ```

        ```sh
        $ pulumi import gitlab:index/deployKey:DeployKey test richardc/example:3
        ```

        :param str resource_name: The name of the resource.
        :param DeployKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeployKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_push: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeployKeyArgs.__new__(DeployKeyArgs)

            __props__.__dict__["can_push"] = can_push
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__.__dict__["title"] = title
            __props__.__dict__["deploy_key_id"] = None
        super(DeployKey, __self__).__init__(
            'gitlab:index/deployKey:DeployKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            can_push: Optional[pulumi.Input[bool]] = None,
            deploy_key_id: Optional[pulumi.Input[int]] = None,
            key: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None) -> 'DeployKey':
        """
        Get an existing DeployKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_push: Allow this deploy key to be used to push changes to the project. Defaults to `false`.
        :param pulumi.Input[int] deploy_key_id: The id of the project deploy key.
        :param pulumi.Input[str] key: The public ssh key body.
        :param pulumi.Input[str] project: The name or id of the project to add the deploy key to.
        :param pulumi.Input[str] title: A title to describe the deploy key with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeployKeyState.__new__(_DeployKeyState)

        __props__.__dict__["can_push"] = can_push
        __props__.__dict__["deploy_key_id"] = deploy_key_id
        __props__.__dict__["key"] = key
        __props__.__dict__["project"] = project
        __props__.__dict__["title"] = title
        return DeployKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canPush")
    def can_push(self) -> pulumi.Output[Optional[bool]]:
        """
        Allow this deploy key to be used to push changes to the project. Defaults to `false`.
        """
        return pulumi.get(self, "can_push")

    @property
    @pulumi.getter(name="deployKeyId")
    def deploy_key_id(self) -> pulumi.Output[int]:
        """
        The id of the project deploy key.
        """
        return pulumi.get(self, "deploy_key_id")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The public ssh key body.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The name or id of the project to add the deploy key to.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        A title to describe the deploy key with.
        """
        return pulumi.get(self, "title")

