# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['InstanceVariableArgs', 'InstanceVariable']

@pulumi.input_type
class InstanceVariableArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str],
                 masked: Optional[pulumi.Input[bool]] = None,
                 protected: Optional[pulumi.Input[bool]] = None,
                 raw: Optional[pulumi.Input[bool]] = None,
                 variable_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceVariable resource.
        :param pulumi.Input[str] key: The name of the variable.
        :param pulumi.Input[str] value: The value of the variable.
        :param pulumi.Input[bool] masked: If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        :param pulumi.Input[bool] protected: If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        :param pulumi.Input[bool] raw: Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        :param pulumi.Input[str] variable_type: The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        """
        InstanceVariableArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            key=key,
            value=value,
            masked=masked,
            protected=protected,
            raw=raw,
            variable_type=variable_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             key: pulumi.Input[str],
             value: pulumi.Input[str],
             masked: Optional[pulumi.Input[bool]] = None,
             protected: Optional[pulumi.Input[bool]] = None,
             raw: Optional[pulumi.Input[bool]] = None,
             variable_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("key", key)
        _setter("value", value)
        if masked is not None:
            _setter("masked", masked)
        if protected is not None:
            _setter("protected", protected)
        if raw is not None:
            _setter("raw", raw)
        if variable_type is not None:
            _setter("variable_type", variable_type)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The name of the variable.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value of the variable.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def masked(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        """
        return pulumi.get(self, "masked")

    @masked.setter
    def masked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "masked", value)

    @property
    @pulumi.getter
    def protected(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        """
        return pulumi.get(self, "protected")

    @protected.setter
    def protected(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "protected", value)

    @property
    @pulumi.getter
    def raw(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        """
        return pulumi.get(self, "raw")

    @raw.setter
    def raw(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "raw", value)

    @property
    @pulumi.getter(name="variableType")
    def variable_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        """
        return pulumi.get(self, "variable_type")

    @variable_type.setter
    def variable_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variable_type", value)


@pulumi.input_type
class _InstanceVariableState:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 masked: Optional[pulumi.Input[bool]] = None,
                 protected: Optional[pulumi.Input[bool]] = None,
                 raw: Optional[pulumi.Input[bool]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 variable_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceVariable resources.
        :param pulumi.Input[str] key: The name of the variable.
        :param pulumi.Input[bool] masked: If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        :param pulumi.Input[bool] protected: If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        :param pulumi.Input[bool] raw: Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        :param pulumi.Input[str] value: The value of the variable.
        :param pulumi.Input[str] variable_type: The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        """
        _InstanceVariableState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            key=key,
            masked=masked,
            protected=protected,
            raw=raw,
            value=value,
            variable_type=variable_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             key: Optional[pulumi.Input[str]] = None,
             masked: Optional[pulumi.Input[bool]] = None,
             protected: Optional[pulumi.Input[bool]] = None,
             raw: Optional[pulumi.Input[bool]] = None,
             value: Optional[pulumi.Input[str]] = None,
             variable_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if key is not None:
            _setter("key", key)
        if masked is not None:
            _setter("masked", masked)
        if protected is not None:
            _setter("protected", protected)
        if raw is not None:
            _setter("raw", raw)
        if value is not None:
            _setter("value", value)
        if variable_type is not None:
            _setter("variable_type", variable_type)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the variable.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def masked(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        """
        return pulumi.get(self, "masked")

    @masked.setter
    def masked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "masked", value)

    @property
    @pulumi.getter
    def protected(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        """
        return pulumi.get(self, "protected")

    @protected.setter
    def protected(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "protected", value)

    @property
    @pulumi.getter
    def raw(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        """
        return pulumi.get(self, "raw")

    @raw.setter
    def raw(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "raw", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the variable.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="variableType")
    def variable_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        """
        return pulumi.get(self, "variable_type")

    @variable_type.setter
    def variable_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variable_type", value)


class InstanceVariable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 masked: Optional[pulumi.Input[bool]] = None,
                 protected: Optional[pulumi.Input[bool]] = None,
                 raw: Optional[pulumi.Input[bool]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 variable_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `InstanceVariable` resource allows to manage the lifecycle of an instance-level CI/CD variable.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/instance_level_ci_variables.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.InstanceVariable("example",
            key="instance_variable_key",
            masked=False,
            protected=False,
            value="instance_variable_value")
        ```

        ## Import

        GitLab instance variables can be imported using an id made up of `variablename`, e.g.

        ```sh
         $ pulumi import gitlab:index/instanceVariable:InstanceVariable example instance_variable_key
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: The name of the variable.
        :param pulumi.Input[bool] masked: If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        :param pulumi.Input[bool] protected: If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        :param pulumi.Input[bool] raw: Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        :param pulumi.Input[str] value: The value of the variable.
        :param pulumi.Input[str] variable_type: The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceVariableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `InstanceVariable` resource allows to manage the lifecycle of an instance-level CI/CD variable.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/instance_level_ci_variables.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.InstanceVariable("example",
            key="instance_variable_key",
            masked=False,
            protected=False,
            value="instance_variable_value")
        ```

        ## Import

        GitLab instance variables can be imported using an id made up of `variablename`, e.g.

        ```sh
         $ pulumi import gitlab:index/instanceVariable:InstanceVariable example instance_variable_key
        ```

        :param str resource_name: The name of the resource.
        :param InstanceVariableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceVariableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            InstanceVariableArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 masked: Optional[pulumi.Input[bool]] = None,
                 protected: Optional[pulumi.Input[bool]] = None,
                 raw: Optional[pulumi.Input[bool]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 variable_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceVariableArgs.__new__(InstanceVariableArgs)

            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["masked"] = masked
            __props__.__dict__["protected"] = protected
            __props__.__dict__["raw"] = raw
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            __props__.__dict__["variable_type"] = variable_type
        super(InstanceVariable, __self__).__init__(
            'gitlab:index/instanceVariable:InstanceVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key: Optional[pulumi.Input[str]] = None,
            masked: Optional[pulumi.Input[bool]] = None,
            protected: Optional[pulumi.Input[bool]] = None,
            raw: Optional[pulumi.Input[bool]] = None,
            value: Optional[pulumi.Input[str]] = None,
            variable_type: Optional[pulumi.Input[str]] = None) -> 'InstanceVariable':
        """
        Get an existing InstanceVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: The name of the variable.
        :param pulumi.Input[bool] masked: If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        :param pulumi.Input[bool] protected: If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        :param pulumi.Input[bool] raw: Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        :param pulumi.Input[str] value: The value of the variable.
        :param pulumi.Input[str] variable_type: The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceVariableState.__new__(_InstanceVariableState)

        __props__.__dict__["key"] = key
        __props__.__dict__["masked"] = masked
        __props__.__dict__["protected"] = protected
        __props__.__dict__["raw"] = raw
        __props__.__dict__["value"] = value
        __props__.__dict__["variable_type"] = variable_type
        return InstanceVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The name of the variable.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def masked(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        """
        return pulumi.get(self, "masked")

    @property
    @pulumi.getter
    def protected(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        """
        return pulumi.get(self, "protected")

    @property
    @pulumi.getter
    def raw(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        """
        return pulumi.get(self, "raw")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The value of the variable.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="variableType")
    def variable_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        """
        return pulumi.get(self, "variable_type")

