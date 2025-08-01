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

__all__ = ['SystemHookArgs', 'SystemHook']

@pulumi.input_type
class SystemHookArgs:
    def __init__(__self__, *,
                 url: pulumi.Input[_builtins.str],
                 enable_ssl_verification: Optional[pulumi.Input[_builtins.bool]] = None,
                 merge_requests_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 push_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 repository_update_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 tag_push_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 token: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a SystemHook resource.
        :param pulumi.Input[_builtins.str] url: The hook URL.
        :param pulumi.Input[_builtins.bool] enable_ssl_verification: Do SSL verification when triggering the hook.
        :param pulumi.Input[_builtins.bool] merge_requests_events: Trigger hook on merge requests events.
        :param pulumi.Input[_builtins.bool] push_events: When true, the hook fires on push events.
        :param pulumi.Input[_builtins.bool] repository_update_events: Trigger hook on repository update events.
        :param pulumi.Input[_builtins.bool] tag_push_events: When true, the hook fires on new tags being pushed.
        :param pulumi.Input[_builtins.str] token: Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        """
        pulumi.set(__self__, "url", url)
        if enable_ssl_verification is not None:
            pulumi.set(__self__, "enable_ssl_verification", enable_ssl_verification)
        if merge_requests_events is not None:
            pulumi.set(__self__, "merge_requests_events", merge_requests_events)
        if push_events is not None:
            pulumi.set(__self__, "push_events", push_events)
        if repository_update_events is not None:
            pulumi.set(__self__, "repository_update_events", repository_update_events)
        if tag_push_events is not None:
            pulumi.set(__self__, "tag_push_events", tag_push_events)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @_builtins.property
    @pulumi.getter
    def url(self) -> pulumi.Input[_builtins.str]:
        """
        The hook URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "url", value)

    @_builtins.property
    @pulumi.getter(name="enableSslVerification")
    def enable_ssl_verification(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Do SSL verification when triggering the hook.
        """
        return pulumi.get(self, "enable_ssl_verification")

    @enable_ssl_verification.setter
    def enable_ssl_verification(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "enable_ssl_verification", value)

    @_builtins.property
    @pulumi.getter(name="mergeRequestsEvents")
    def merge_requests_events(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Trigger hook on merge requests events.
        """
        return pulumi.get(self, "merge_requests_events")

    @merge_requests_events.setter
    def merge_requests_events(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "merge_requests_events", value)

    @_builtins.property
    @pulumi.getter(name="pushEvents")
    def push_events(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        When true, the hook fires on push events.
        """
        return pulumi.get(self, "push_events")

    @push_events.setter
    def push_events(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "push_events", value)

    @_builtins.property
    @pulumi.getter(name="repositoryUpdateEvents")
    def repository_update_events(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Trigger hook on repository update events.
        """
        return pulumi.get(self, "repository_update_events")

    @repository_update_events.setter
    def repository_update_events(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "repository_update_events", value)

    @_builtins.property
    @pulumi.getter(name="tagPushEvents")
    def tag_push_events(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        When true, the hook fires on new tags being pushed.
        """
        return pulumi.get(self, "tag_push_events")

    @tag_push_events.setter
    def tag_push_events(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "tag_push_events", value)

    @_builtins.property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "token", value)


@pulumi.input_type
class _SystemHookState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[_builtins.str]] = None,
                 enable_ssl_verification: Optional[pulumi.Input[_builtins.bool]] = None,
                 merge_requests_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 push_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 repository_update_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 tag_push_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 token: Optional[pulumi.Input[_builtins.str]] = None,
                 url: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering SystemHook resources.
        :param pulumi.Input[_builtins.str] created_at: The date and time the hook was created in ISO8601 format.
        :param pulumi.Input[_builtins.bool] enable_ssl_verification: Do SSL verification when triggering the hook.
        :param pulumi.Input[_builtins.bool] merge_requests_events: Trigger hook on merge requests events.
        :param pulumi.Input[_builtins.bool] push_events: When true, the hook fires on push events.
        :param pulumi.Input[_builtins.bool] repository_update_events: Trigger hook on repository update events.
        :param pulumi.Input[_builtins.bool] tag_push_events: When true, the hook fires on new tags being pushed.
        :param pulumi.Input[_builtins.str] token: Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        :param pulumi.Input[_builtins.str] url: The hook URL.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if enable_ssl_verification is not None:
            pulumi.set(__self__, "enable_ssl_verification", enable_ssl_verification)
        if merge_requests_events is not None:
            pulumi.set(__self__, "merge_requests_events", merge_requests_events)
        if push_events is not None:
            pulumi.set(__self__, "push_events", push_events)
        if repository_update_events is not None:
            pulumi.set(__self__, "repository_update_events", repository_update_events)
        if tag_push_events is not None:
            pulumi.set(__self__, "tag_push_events", tag_push_events)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @_builtins.property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The date and time the hook was created in ISO8601 format.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "created_at", value)

    @_builtins.property
    @pulumi.getter(name="enableSslVerification")
    def enable_ssl_verification(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Do SSL verification when triggering the hook.
        """
        return pulumi.get(self, "enable_ssl_verification")

    @enable_ssl_verification.setter
    def enable_ssl_verification(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "enable_ssl_verification", value)

    @_builtins.property
    @pulumi.getter(name="mergeRequestsEvents")
    def merge_requests_events(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Trigger hook on merge requests events.
        """
        return pulumi.get(self, "merge_requests_events")

    @merge_requests_events.setter
    def merge_requests_events(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "merge_requests_events", value)

    @_builtins.property
    @pulumi.getter(name="pushEvents")
    def push_events(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        When true, the hook fires on push events.
        """
        return pulumi.get(self, "push_events")

    @push_events.setter
    def push_events(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "push_events", value)

    @_builtins.property
    @pulumi.getter(name="repositoryUpdateEvents")
    def repository_update_events(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Trigger hook on repository update events.
        """
        return pulumi.get(self, "repository_update_events")

    @repository_update_events.setter
    def repository_update_events(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "repository_update_events", value)

    @_builtins.property
    @pulumi.getter(name="tagPushEvents")
    def tag_push_events(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        When true, the hook fires on new tags being pushed.
        """
        return pulumi.get(self, "tag_push_events")

    @tag_push_events.setter
    def tag_push_events(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "tag_push_events", value)

    @_builtins.property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "token", value)

    @_builtins.property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The hook URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "url", value)


@pulumi.type_token("gitlab:index/systemHook:SystemHook")
class SystemHook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_ssl_verification: Optional[pulumi.Input[_builtins.bool]] = None,
                 merge_requests_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 push_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 repository_update_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 tag_push_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 token: Optional[pulumi.Input[_builtins.str]] = None,
                 url: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        The `SystemHook` resource allows to manage the lifecycle of a system hook.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/system_hooks/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.SystemHook("example",
            url="https://example.com/hook-%d",
            token="secret-token",
            push_events=True,
            tag_push_events=True,
            merge_requests_events=True,
            repository_update_events=True,
            enable_ssl_verification=True)
        ```

        ## Import

        Starting in Terraform v1.5.0, you can use an import block to import `gitlab_system_hook`. For example:

        terraform

        import {

          to = gitlab_system_hook.example

          id = "see CLI command below for ID"

        }

        Importing using the CLI is supported with the following syntax:

        You can import a system hook using the hook id `{hook-id}`, e.g.

        ```sh
        $ pulumi import gitlab:index/systemHook:SystemHook example 42
        ```

        NOTE: the `token` attribute won't be available for imported resources.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.bool] enable_ssl_verification: Do SSL verification when triggering the hook.
        :param pulumi.Input[_builtins.bool] merge_requests_events: Trigger hook on merge requests events.
        :param pulumi.Input[_builtins.bool] push_events: When true, the hook fires on push events.
        :param pulumi.Input[_builtins.bool] repository_update_events: Trigger hook on repository update events.
        :param pulumi.Input[_builtins.bool] tag_push_events: When true, the hook fires on new tags being pushed.
        :param pulumi.Input[_builtins.str] token: Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        :param pulumi.Input[_builtins.str] url: The hook URL.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemHookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `SystemHook` resource allows to manage the lifecycle of a system hook.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/system_hooks/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.SystemHook("example",
            url="https://example.com/hook-%d",
            token="secret-token",
            push_events=True,
            tag_push_events=True,
            merge_requests_events=True,
            repository_update_events=True,
            enable_ssl_verification=True)
        ```

        ## Import

        Starting in Terraform v1.5.0, you can use an import block to import `gitlab_system_hook`. For example:

        terraform

        import {

          to = gitlab_system_hook.example

          id = "see CLI command below for ID"

        }

        Importing using the CLI is supported with the following syntax:

        You can import a system hook using the hook id `{hook-id}`, e.g.

        ```sh
        $ pulumi import gitlab:index/systemHook:SystemHook example 42
        ```

        NOTE: the `token` attribute won't be available for imported resources.

        :param str resource_name: The name of the resource.
        :param SystemHookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemHookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_ssl_verification: Optional[pulumi.Input[_builtins.bool]] = None,
                 merge_requests_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 push_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 repository_update_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 tag_push_events: Optional[pulumi.Input[_builtins.bool]] = None,
                 token: Optional[pulumi.Input[_builtins.str]] = None,
                 url: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemHookArgs.__new__(SystemHookArgs)

            __props__.__dict__["enable_ssl_verification"] = enable_ssl_verification
            __props__.__dict__["merge_requests_events"] = merge_requests_events
            __props__.__dict__["push_events"] = push_events
            __props__.__dict__["repository_update_events"] = repository_update_events
            __props__.__dict__["tag_push_events"] = tag_push_events
            __props__.__dict__["token"] = None if token is None else pulumi.Output.secret(token)
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["created_at"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SystemHook, __self__).__init__(
            'gitlab:index/systemHook:SystemHook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[_builtins.str]] = None,
            enable_ssl_verification: Optional[pulumi.Input[_builtins.bool]] = None,
            merge_requests_events: Optional[pulumi.Input[_builtins.bool]] = None,
            push_events: Optional[pulumi.Input[_builtins.bool]] = None,
            repository_update_events: Optional[pulumi.Input[_builtins.bool]] = None,
            tag_push_events: Optional[pulumi.Input[_builtins.bool]] = None,
            token: Optional[pulumi.Input[_builtins.str]] = None,
            url: Optional[pulumi.Input[_builtins.str]] = None) -> 'SystemHook':
        """
        Get an existing SystemHook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] created_at: The date and time the hook was created in ISO8601 format.
        :param pulumi.Input[_builtins.bool] enable_ssl_verification: Do SSL verification when triggering the hook.
        :param pulumi.Input[_builtins.bool] merge_requests_events: Trigger hook on merge requests events.
        :param pulumi.Input[_builtins.bool] push_events: When true, the hook fires on push events.
        :param pulumi.Input[_builtins.bool] repository_update_events: Trigger hook on repository update events.
        :param pulumi.Input[_builtins.bool] tag_push_events: When true, the hook fires on new tags being pushed.
        :param pulumi.Input[_builtins.str] token: Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        :param pulumi.Input[_builtins.str] url: The hook URL.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemHookState.__new__(_SystemHookState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["enable_ssl_verification"] = enable_ssl_verification
        __props__.__dict__["merge_requests_events"] = merge_requests_events
        __props__.__dict__["push_events"] = push_events
        __props__.__dict__["repository_update_events"] = repository_update_events
        __props__.__dict__["tag_push_events"] = tag_push_events
        __props__.__dict__["token"] = token
        __props__.__dict__["url"] = url
        return SystemHook(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[_builtins.str]:
        """
        The date and time the hook was created in ISO8601 format.
        """
        return pulumi.get(self, "created_at")

    @_builtins.property
    @pulumi.getter(name="enableSslVerification")
    def enable_ssl_verification(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        Do SSL verification when triggering the hook.
        """
        return pulumi.get(self, "enable_ssl_verification")

    @_builtins.property
    @pulumi.getter(name="mergeRequestsEvents")
    def merge_requests_events(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        Trigger hook on merge requests events.
        """
        return pulumi.get(self, "merge_requests_events")

    @_builtins.property
    @pulumi.getter(name="pushEvents")
    def push_events(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        When true, the hook fires on push events.
        """
        return pulumi.get(self, "push_events")

    @_builtins.property
    @pulumi.getter(name="repositoryUpdateEvents")
    def repository_update_events(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        Trigger hook on repository update events.
        """
        return pulumi.get(self, "repository_update_events")

    @_builtins.property
    @pulumi.getter(name="tagPushEvents")
    def tag_push_events(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        When true, the hook fires on new tags being pushed.
        """
        return pulumi.get(self, "tag_push_events")

    @_builtins.property
    @pulumi.getter
    def token(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @_builtins.property
    @pulumi.getter
    def url(self) -> pulumi.Output[_builtins.str]:
        """
        The hook URL.
        """
        return pulumi.get(self, "url")

