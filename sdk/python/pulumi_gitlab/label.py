# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Label(pulumi.CustomResource):
    color: pulumi.Output[str]
    """
    The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
    """
    description: pulumi.Output[str]
    """
    The description of the label.
    """
    name: pulumi.Output[str]
    """
    The name of the label.
    """
    project: pulumi.Output[str]
    """
    The name or id of the project to add the label to.
    """
    def __init__(__self__, resource_name, opts=None, color=None, description=None, name=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        This resource allows you to create and manage labels for your GitLab projects.
        For further information on labels, consult the [gitlab
        documentation](https://docs.gitlab.com/ee/user/project/labels.html#project-labels).




        > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/label.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        :param pulumi.Input[str] description: The description of the label.
        :param pulumi.Input[str] name: The name of the label.
        :param pulumi.Input[str] project: The name or id of the project to add the label to.
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

            if color is None:
                raise TypeError("Missing required property 'color'")
            __props__['color'] = color
            __props__['description'] = description
            __props__['name'] = name
            if project is None:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
        super(Label, __self__).__init__(
            'gitlab:index/label:Label',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, color=None, description=None, name=None, project=None):
        """
        Get an existing Label resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        :param pulumi.Input[str] description: The description of the label.
        :param pulumi.Input[str] name: The name of the label.
        :param pulumi.Input[str] project: The name or id of the project to add the label to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["color"] = color
        __props__["description"] = description
        __props__["name"] = name
        __props__["project"] = project
        return Label(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

