# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['DeployKey']


class DeployKey(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_push: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## # gitlab\_deploy\_key

        This resource allows you to create and manage deploy keys for your GitLab projects.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.DeployKey("example",
            key="ssh-rsa AAAA...",
            project="example/deploying",
            title="Example deploy key")
        ```

        ## Import

        GitLab deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.

        ```sh
         $ pulumi import gitlab:index/deployKey:DeployKey test 1:3
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_push: Allow this deploy key to be used to push changes to the project.  Defaults to `false`. **NOTE::** this cannot currently be managed.
        :param pulumi.Input[str] key: The public ssh key body.
        :param pulumi.Input[str] project: The name or id of the project to add the deploy key to.
        :param pulumi.Input[str] title: A title to describe the deploy key with.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['can_push'] = can_push
            if key is None:
                raise TypeError("Missing required property 'key'")
            __props__['key'] = key
            if project is None:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
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
            key: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None) -> 'DeployKey':
        """
        Get an existing DeployKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_push: Allow this deploy key to be used to push changes to the project.  Defaults to `false`. **NOTE::** this cannot currently be managed.
        :param pulumi.Input[str] key: The public ssh key body.
        :param pulumi.Input[str] project: The name or id of the project to add the deploy key to.
        :param pulumi.Input[str] title: A title to describe the deploy key with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["can_push"] = can_push
        __props__["key"] = key
        __props__["project"] = project
        __props__["title"] = title
        return DeployKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canPush")
    def can_push(self) -> pulumi.Output[Optional[bool]]:
        """
        Allow this deploy key to be used to push changes to the project.  Defaults to `false`. **NOTE::** this cannot currently be managed.
        """
        return pulumi.get(self, "can_push")

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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

