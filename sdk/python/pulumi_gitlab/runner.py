# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RunnerArgs', 'Runner']

@pulumi.input_type
class RunnerArgs:
    def __init__(__self__, *,
                 registration_token: pulumi.Input[str],
                 access_level: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 maximum_timeout: Optional[pulumi.Input[int]] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 run_untagged: Optional[pulumi.Input[bool]] = None,
                 tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Runner resource.
        :param pulumi.Input[str] registration_token: The registration token used to register the runner.
        :param pulumi.Input[str] access_level: The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
        :param pulumi.Input[str] description: The runner's description.
        :param pulumi.Input[bool] locked: Whether the runner should be locked for current project.
        :param pulumi.Input[int] maximum_timeout: Maximum timeout set when this runner handles the job.
        :param pulumi.Input[bool] paused: Whether the runner should ignore new jobs.
        :param pulumi.Input[bool] run_untagged: Whether the runner should handle untagged jobs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tag_lists: List of runner’s tags.
        """
        pulumi.set(__self__, "registration_token", registration_token)
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if maximum_timeout is not None:
            pulumi.set(__self__, "maximum_timeout", maximum_timeout)
        if paused is not None:
            pulumi.set(__self__, "paused", paused)
        if run_untagged is not None:
            pulumi.set(__self__, "run_untagged", run_untagged)
        if tag_lists is not None:
            pulumi.set(__self__, "tag_lists", tag_lists)

    @property
    @pulumi.getter(name="registrationToken")
    def registration_token(self) -> pulumi.Input[str]:
        """
        The registration token used to register the runner.
        """
        return pulumi.get(self, "registration_token")

    @registration_token.setter
    def registration_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "registration_token", value)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The runner's description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def locked(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the runner should be locked for current project.
        """
        return pulumi.get(self, "locked")

    @locked.setter
    def locked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "locked", value)

    @property
    @pulumi.getter(name="maximumTimeout")
    def maximum_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum timeout set when this runner handles the job.
        """
        return pulumi.get(self, "maximum_timeout")

    @maximum_timeout.setter
    def maximum_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_timeout", value)

    @property
    @pulumi.getter
    def paused(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the runner should ignore new jobs.
        """
        return pulumi.get(self, "paused")

    @paused.setter
    def paused(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "paused", value)

    @property
    @pulumi.getter(name="runUntagged")
    def run_untagged(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the runner should handle untagged jobs.
        """
        return pulumi.get(self, "run_untagged")

    @run_untagged.setter
    def run_untagged(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "run_untagged", value)

    @property
    @pulumi.getter(name="tagLists")
    def tag_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of runner’s tags.
        """
        return pulumi.get(self, "tag_lists")

    @tag_lists.setter
    def tag_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tag_lists", value)


@pulumi.input_type
class _RunnerState:
    def __init__(__self__, *,
                 access_level: Optional[pulumi.Input[str]] = None,
                 authentication_token: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 maximum_timeout: Optional[pulumi.Input[int]] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 registration_token: Optional[pulumi.Input[str]] = None,
                 run_untagged: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Runner resources.
        :param pulumi.Input[str] access_level: The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
        :param pulumi.Input[str] authentication_token: The authentication token used for building a config.toml file. This value is not present when imported.
        :param pulumi.Input[str] description: The runner's description.
        :param pulumi.Input[bool] locked: Whether the runner should be locked for current project.
        :param pulumi.Input[int] maximum_timeout: Maximum timeout set when this runner handles the job.
        :param pulumi.Input[bool] paused: Whether the runner should ignore new jobs.
        :param pulumi.Input[str] registration_token: The registration token used to register the runner.
        :param pulumi.Input[bool] run_untagged: Whether the runner should handle untagged jobs.
        :param pulumi.Input[str] status: The status of runners to show, one of: online and offline. active and paused are also possible values
               			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tag_lists: List of runner’s tags.
        """
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if authentication_token is not None:
            pulumi.set(__self__, "authentication_token", authentication_token)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if maximum_timeout is not None:
            pulumi.set(__self__, "maximum_timeout", maximum_timeout)
        if paused is not None:
            pulumi.set(__self__, "paused", paused)
        if registration_token is not None:
            pulumi.set(__self__, "registration_token", registration_token)
        if run_untagged is not None:
            pulumi.set(__self__, "run_untagged", run_untagged)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tag_lists is not None:
            pulumi.set(__self__, "tag_lists", tag_lists)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication token used for building a config.toml file. This value is not present when imported.
        """
        return pulumi.get(self, "authentication_token")

    @authentication_token.setter
    def authentication_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication_token", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The runner's description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def locked(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the runner should be locked for current project.
        """
        return pulumi.get(self, "locked")

    @locked.setter
    def locked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "locked", value)

    @property
    @pulumi.getter(name="maximumTimeout")
    def maximum_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum timeout set when this runner handles the job.
        """
        return pulumi.get(self, "maximum_timeout")

    @maximum_timeout.setter
    def maximum_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_timeout", value)

    @property
    @pulumi.getter
    def paused(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the runner should ignore new jobs.
        """
        return pulumi.get(self, "paused")

    @paused.setter
    def paused(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "paused", value)

    @property
    @pulumi.getter(name="registrationToken")
    def registration_token(self) -> Optional[pulumi.Input[str]]:
        """
        The registration token used to register the runner.
        """
        return pulumi.get(self, "registration_token")

    @registration_token.setter
    def registration_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registration_token", value)

    @property
    @pulumi.getter(name="runUntagged")
    def run_untagged(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the runner should handle untagged jobs.
        """
        return pulumi.get(self, "run_untagged")

    @run_untagged.setter
    def run_untagged(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "run_untagged", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of runners to show, one of: online and offline. active and paused are also possible values
        			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="tagLists")
    def tag_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of runner’s tags.
        """
        return pulumi.get(self, "tag_lists")

    @tag_lists.setter
    def tag_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tag_lists", value)


class Runner(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 maximum_timeout: Optional[pulumi.Input[int]] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 registration_token: Optional[pulumi.Input[str]] = None,
                 run_untagged: Optional[pulumi.Input[bool]] = None,
                 tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        The `Runner` resource allows to manage the lifecycle of a runner.

        A runner can either be registered at an instance level or group level.
        The runner will be registered at a group level if the token used is from a group, or at an instance level if the token used is for the instance.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/runners.html#register-a-new-runner)

        ## Import

        A GitLab Runner can be imported using the runner's ID, eg

        ```sh
         $ pulumi import gitlab:index/runner:Runner this 1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
        :param pulumi.Input[str] description: The runner's description.
        :param pulumi.Input[bool] locked: Whether the runner should be locked for current project.
        :param pulumi.Input[int] maximum_timeout: Maximum timeout set when this runner handles the job.
        :param pulumi.Input[bool] paused: Whether the runner should ignore new jobs.
        :param pulumi.Input[str] registration_token: The registration token used to register the runner.
        :param pulumi.Input[bool] run_untagged: Whether the runner should handle untagged jobs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tag_lists: List of runner’s tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RunnerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `Runner` resource allows to manage the lifecycle of a runner.

        A runner can either be registered at an instance level or group level.
        The runner will be registered at a group level if the token used is from a group, or at an instance level if the token used is for the instance.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/runners.html#register-a-new-runner)

        ## Import

        A GitLab Runner can be imported using the runner's ID, eg

        ```sh
         $ pulumi import gitlab:index/runner:Runner this 1
        ```

        :param str resource_name: The name of the resource.
        :param RunnerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RunnerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 maximum_timeout: Optional[pulumi.Input[int]] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 registration_token: Optional[pulumi.Input[str]] = None,
                 run_untagged: Optional[pulumi.Input[bool]] = None,
                 tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RunnerArgs.__new__(RunnerArgs)

            __props__.__dict__["access_level"] = access_level
            __props__.__dict__["description"] = description
            __props__.__dict__["locked"] = locked
            __props__.__dict__["maximum_timeout"] = maximum_timeout
            __props__.__dict__["paused"] = paused
            if registration_token is None and not opts.urn:
                raise TypeError("Missing required property 'registration_token'")
            __props__.__dict__["registration_token"] = None if registration_token is None else pulumi.Output.secret(registration_token)
            __props__.__dict__["run_untagged"] = run_untagged
            __props__.__dict__["tag_lists"] = tag_lists
            __props__.__dict__["authentication_token"] = None
            __props__.__dict__["status"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["authenticationToken", "registrationToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Runner, __self__).__init__(
            'gitlab:index/runner:Runner',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_level: Optional[pulumi.Input[str]] = None,
            authentication_token: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            locked: Optional[pulumi.Input[bool]] = None,
            maximum_timeout: Optional[pulumi.Input[int]] = None,
            paused: Optional[pulumi.Input[bool]] = None,
            registration_token: Optional[pulumi.Input[str]] = None,
            run_untagged: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Runner':
        """
        Get an existing Runner resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
        :param pulumi.Input[str] authentication_token: The authentication token used for building a config.toml file. This value is not present when imported.
        :param pulumi.Input[str] description: The runner's description.
        :param pulumi.Input[bool] locked: Whether the runner should be locked for current project.
        :param pulumi.Input[int] maximum_timeout: Maximum timeout set when this runner handles the job.
        :param pulumi.Input[bool] paused: Whether the runner should ignore new jobs.
        :param pulumi.Input[str] registration_token: The registration token used to register the runner.
        :param pulumi.Input[bool] run_untagged: Whether the runner should handle untagged jobs.
        :param pulumi.Input[str] status: The status of runners to show, one of: online and offline. active and paused are also possible values
               			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tag_lists: List of runner’s tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RunnerState.__new__(_RunnerState)

        __props__.__dict__["access_level"] = access_level
        __props__.__dict__["authentication_token"] = authentication_token
        __props__.__dict__["description"] = description
        __props__.__dict__["locked"] = locked
        __props__.__dict__["maximum_timeout"] = maximum_timeout
        __props__.__dict__["paused"] = paused
        __props__.__dict__["registration_token"] = registration_token
        __props__.__dict__["run_untagged"] = run_untagged
        __props__.__dict__["status"] = status
        __props__.__dict__["tag_lists"] = tag_lists
        return Runner(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> pulumi.Output[str]:
        """
        The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
        """
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter(name="authenticationToken")
    def authentication_token(self) -> pulumi.Output[str]:
        """
        The authentication token used for building a config.toml file. This value is not present when imported.
        """
        return pulumi.get(self, "authentication_token")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The runner's description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def locked(self) -> pulumi.Output[bool]:
        """
        Whether the runner should be locked for current project.
        """
        return pulumi.get(self, "locked")

    @property
    @pulumi.getter(name="maximumTimeout")
    def maximum_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum timeout set when this runner handles the job.
        """
        return pulumi.get(self, "maximum_timeout")

    @property
    @pulumi.getter
    def paused(self) -> pulumi.Output[bool]:
        """
        Whether the runner should ignore new jobs.
        """
        return pulumi.get(self, "paused")

    @property
    @pulumi.getter(name="registrationToken")
    def registration_token(self) -> pulumi.Output[str]:
        """
        The registration token used to register the runner.
        """
        return pulumi.get(self, "registration_token")

    @property
    @pulumi.getter(name="runUntagged")
    def run_untagged(self) -> pulumi.Output[bool]:
        """
        Whether the runner should handle untagged jobs.
        """
        return pulumi.get(self, "run_untagged")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of runners to show, one of: online and offline. active and paused are also possible values
        			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tagLists")
    def tag_lists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of runner’s tags.
        """
        return pulumi.get(self, "tag_lists")
