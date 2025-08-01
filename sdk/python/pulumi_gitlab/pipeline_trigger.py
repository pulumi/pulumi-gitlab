# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['PipelineTriggerArgs', 'PipelineTrigger']

@pulumi.input_type
class PipelineTriggerArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[_builtins.str],
                 project: pulumi.Input[_builtins.str]):
        """
        The set of arguments for constructing a PipelineTrigger resource.
        :param pulumi.Input[_builtins.str] description: The description of the pipeline trigger.
        :param pulumi.Input[_builtins.str] project: The name or id of the project to add the trigger to.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "project", project)

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Input[_builtins.str]:
        """
        The description of the pipeline trigger.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def project(self) -> pulumi.Input[_builtins.str]:
        """
        The name or id of the project to add the trigger to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _PipelineTriggerState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 pipeline_trigger_id: Optional[pulumi.Input[_builtins.int]] = None,
                 project: Optional[pulumi.Input[_builtins.str]] = None,
                 token: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering PipelineTrigger resources.
        :param pulumi.Input[_builtins.str] description: The description of the pipeline trigger.
        :param pulumi.Input[_builtins.int] pipeline_trigger_id: The pipeline trigger id.
        :param pulumi.Input[_builtins.str] project: The name or id of the project to add the trigger to.
        :param pulumi.Input[_builtins.str] token: The pipeline trigger token. This value is not available during import.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if pipeline_trigger_id is not None:
            pulumi.set(__self__, "pipeline_trigger_id", pipeline_trigger_id)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The description of the pipeline trigger.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter(name="pipelineTriggerId")
    def pipeline_trigger_id(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The pipeline trigger id.
        """
        return pulumi.get(self, "pipeline_trigger_id")

    @pipeline_trigger_id.setter
    def pipeline_trigger_id(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "pipeline_trigger_id", value)

    @_builtins.property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name or id of the project to add the trigger to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "project", value)

    @_builtins.property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The pipeline trigger token. This value is not available during import.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "token", value)


@pulumi.type_token("gitlab:index/pipelineTrigger:PipelineTrigger")
class PipelineTrigger(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 project: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        The `PipelineTrigger` resource allows to manage the lifecycle of a pipeline trigger.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/pipeline_triggers/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.PipelineTrigger("example",
            project="12345",
            description="Used to trigger builds")
        ```

        ## Import

        Starting in Terraform v1.5.0, you can use an import block to import `gitlab_pipeline_trigger`. For example:

        terraform

        import {

          to = gitlab_pipeline_trigger.example

          id = "see CLI command below for ID"

        }

        Importing using the CLI is supported with the following syntax:

        GitLab pipeline triggers can be imported using an id made up of `{project_id}:{pipeline_trigger_id}`, e.g.

        ```sh
        $ pulumi import gitlab:index/pipelineTrigger:PipelineTrigger test 1:3
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] description: The description of the pipeline trigger.
        :param pulumi.Input[_builtins.str] project: The name or id of the project to add the trigger to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineTriggerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `PipelineTrigger` resource allows to manage the lifecycle of a pipeline trigger.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/pipeline_triggers/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.PipelineTrigger("example",
            project="12345",
            description="Used to trigger builds")
        ```

        ## Import

        Starting in Terraform v1.5.0, you can use an import block to import `gitlab_pipeline_trigger`. For example:

        terraform

        import {

          to = gitlab_pipeline_trigger.example

          id = "see CLI command below for ID"

        }

        Importing using the CLI is supported with the following syntax:

        GitLab pipeline triggers can be imported using an id made up of `{project_id}:{pipeline_trigger_id}`, e.g.

        ```sh
        $ pulumi import gitlab:index/pipelineTrigger:PipelineTrigger test 1:3
        ```

        :param str resource_name: The name of the resource.
        :param PipelineTriggerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineTriggerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 project: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PipelineTriggerArgs.__new__(PipelineTriggerArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["pipeline_trigger_id"] = None
            __props__.__dict__["token"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(PipelineTrigger, __self__).__init__(
            'gitlab:index/pipelineTrigger:PipelineTrigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[_builtins.str]] = None,
            pipeline_trigger_id: Optional[pulumi.Input[_builtins.int]] = None,
            project: Optional[pulumi.Input[_builtins.str]] = None,
            token: Optional[pulumi.Input[_builtins.str]] = None) -> 'PipelineTrigger':
        """
        Get an existing PipelineTrigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] description: The description of the pipeline trigger.
        :param pulumi.Input[_builtins.int] pipeline_trigger_id: The pipeline trigger id.
        :param pulumi.Input[_builtins.str] project: The name or id of the project to add the trigger to.
        :param pulumi.Input[_builtins.str] token: The pipeline trigger token. This value is not available during import.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PipelineTriggerState.__new__(_PipelineTriggerState)

        __props__.__dict__["description"] = description
        __props__.__dict__["pipeline_trigger_id"] = pipeline_trigger_id
        __props__.__dict__["project"] = project
        __props__.__dict__["token"] = token
        return PipelineTrigger(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Output[_builtins.str]:
        """
        The description of the pipeline trigger.
        """
        return pulumi.get(self, "description")

    @_builtins.property
    @pulumi.getter(name="pipelineTriggerId")
    def pipeline_trigger_id(self) -> pulumi.Output[_builtins.int]:
        """
        The pipeline trigger id.
        """
        return pulumi.get(self, "pipeline_trigger_id")

    @_builtins.property
    @pulumi.getter
    def project(self) -> pulumi.Output[_builtins.str]:
        """
        The name or id of the project to add the trigger to.
        """
        return pulumi.get(self, "project")

    @_builtins.property
    @pulumi.getter
    def token(self) -> pulumi.Output[_builtins.str]:
        """
        The pipeline trigger token. This value is not available during import.
        """
        return pulumi.get(self, "token")

