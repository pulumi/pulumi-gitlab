# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class ProjectMembership(pulumi.CustomResource):
    access_level: pulumi.Output[str]
    """
    One of five levels of access to the project.
    """
    project_id: pulumi.Output[str]
    """
    The id of the project.
    """
    user_id: pulumi.Output[float]
    """
    The id of the user.
    """
    def __init__(__self__, resource_name, opts=None, access_level=None, project_id=None, user_id=None, __props__=None, __name__=None, __opts__=None):
        """
        This resource allows you to add a current user to an existing project with a set access level.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/project_membership.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: One of five levels of access to the project.
        :param pulumi.Input[str] project_id: The id of the project.
        :param pulumi.Input[float] user_id: The id of the user.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if access_level is None:
                raise TypeError("Missing required property 'access_level'")
            __props__['access_level'] = access_level
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            if user_id is None:
                raise TypeError("Missing required property 'user_id'")
            __props__['user_id'] = user_id
        super(ProjectMembership, __self__).__init__(
            'gitlab:index/projectMembership:ProjectMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_level=None, project_id=None, user_id=None):
        """
        Get an existing ProjectMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: One of five levels of access to the project.
        :param pulumi.Input[str] project_id: The id of the project.
        :param pulumi.Input[float] user_id: The id of the user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_level"] = access_level
        __props__["project_id"] = project_id
        __props__["user_id"] = user_id
        return ProjectMembership(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

