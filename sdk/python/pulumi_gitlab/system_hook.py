# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemHookArgs', 'SystemHook']

@pulumi.input_type
class SystemHookArgs:
    def __init__(__self__, *,
                 url: pulumi.Input[str],
                 enable_ssl_verification: Optional[pulumi.Input[bool]] = None,
                 merge_requests_events: Optional[pulumi.Input[bool]] = None,
                 push_events: Optional[pulumi.Input[bool]] = None,
                 repository_update_events: Optional[pulumi.Input[bool]] = None,
                 tag_push_events: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemHook resource.
        :param pulumi.Input[str] url: The hook URL.
        :param pulumi.Input[bool] enable_ssl_verification: Do SSL verification when triggering the hook.
        :param pulumi.Input[bool] merge_requests_events: Trigger hook on merge requests events.
        :param pulumi.Input[bool] push_events: When true, the hook fires on push events.
        :param pulumi.Input[bool] repository_update_events: Trigger hook on repository update events.
        :param pulumi.Input[bool] tag_push_events: When true, the hook fires on new tags being pushed.
        :param pulumi.Input[str] token: Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
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

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The hook URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="enableSslVerification")
    def enable_ssl_verification(self) -> Optional[pulumi.Input[bool]]:
        """
        Do SSL verification when triggering the hook.
        """
        return pulumi.get(self, "enable_ssl_verification")

    @enable_ssl_verification.setter
    def enable_ssl_verification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ssl_verification", value)

    @property
    @pulumi.getter(name="mergeRequestsEvents")
    def merge_requests_events(self) -> Optional[pulumi.Input[bool]]:
        """
        Trigger hook on merge requests events.
        """
        return pulumi.get(self, "merge_requests_events")

    @merge_requests_events.setter
    def merge_requests_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "merge_requests_events", value)

    @property
    @pulumi.getter(name="pushEvents")
    def push_events(self) -> Optional[pulumi.Input[bool]]:
        """
        When true, the hook fires on push events.
        """
        return pulumi.get(self, "push_events")

    @push_events.setter
    def push_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "push_events", value)

    @property
    @pulumi.getter(name="repositoryUpdateEvents")
    def repository_update_events(self) -> Optional[pulumi.Input[bool]]:
        """
        Trigger hook on repository update events.
        """
        return pulumi.get(self, "repository_update_events")

    @repository_update_events.setter
    def repository_update_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "repository_update_events", value)

    @property
    @pulumi.getter(name="tagPushEvents")
    def tag_push_events(self) -> Optional[pulumi.Input[bool]]:
        """
        When true, the hook fires on new tags being pushed.
        """
        return pulumi.get(self, "tag_push_events")

    @tag_push_events.setter
    def tag_push_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tag_push_events", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


@pulumi.input_type
class _SystemHookState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 enable_ssl_verification: Optional[pulumi.Input[bool]] = None,
                 merge_requests_events: Optional[pulumi.Input[bool]] = None,
                 push_events: Optional[pulumi.Input[bool]] = None,
                 repository_update_events: Optional[pulumi.Input[bool]] = None,
                 tag_push_events: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemHook resources.
        :param pulumi.Input[str] created_at: The date and time the hook was created in ISO8601 format.
        :param pulumi.Input[bool] enable_ssl_verification: Do SSL verification when triggering the hook.
        :param pulumi.Input[bool] merge_requests_events: Trigger hook on merge requests events.
        :param pulumi.Input[bool] push_events: When true, the hook fires on push events.
        :param pulumi.Input[bool] repository_update_events: Trigger hook on repository update events.
        :param pulumi.Input[bool] tag_push_events: When true, the hook fires on new tags being pushed.
        :param pulumi.Input[str] token: Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        :param pulumi.Input[str] url: The hook URL.
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

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the hook was created in ISO8601 format.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="enableSslVerification")
    def enable_ssl_verification(self) -> Optional[pulumi.Input[bool]]:
        """
        Do SSL verification when triggering the hook.
        """
        return pulumi.get(self, "enable_ssl_verification")

    @enable_ssl_verification.setter
    def enable_ssl_verification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ssl_verification", value)

    @property
    @pulumi.getter(name="mergeRequestsEvents")
    def merge_requests_events(self) -> Optional[pulumi.Input[bool]]:
        """
        Trigger hook on merge requests events.
        """
        return pulumi.get(self, "merge_requests_events")

    @merge_requests_events.setter
    def merge_requests_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "merge_requests_events", value)

    @property
    @pulumi.getter(name="pushEvents")
    def push_events(self) -> Optional[pulumi.Input[bool]]:
        """
        When true, the hook fires on push events.
        """
        return pulumi.get(self, "push_events")

    @push_events.setter
    def push_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "push_events", value)

    @property
    @pulumi.getter(name="repositoryUpdateEvents")
    def repository_update_events(self) -> Optional[pulumi.Input[bool]]:
        """
        Trigger hook on repository update events.
        """
        return pulumi.get(self, "repository_update_events")

    @repository_update_events.setter
    def repository_update_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "repository_update_events", value)

    @property
    @pulumi.getter(name="tagPushEvents")
    def tag_push_events(self) -> Optional[pulumi.Input[bool]]:
        """
        When true, the hook fires on new tags being pushed.
        """
        return pulumi.get(self, "tag_push_events")

    @tag_push_events.setter
    def tag_push_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tag_push_events", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The hook URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class SystemHook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_ssl_verification: Optional[pulumi.Input[bool]] = None,
                 merge_requests_events: Optional[pulumi.Input[bool]] = None,
                 push_events: Optional[pulumi.Input[bool]] = None,
                 repository_update_events: Optional[pulumi.Input[bool]] = None,
                 tag_push_events: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `SystemHook` resource allows to manage the lifecycle of a system hook.

        > This resource requires GitLab 14.9

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/system_hooks.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.SystemHook("example",
            enable_ssl_verification=True,
            merge_requests_events=True,
            push_events=True,
            repository_update_events=True,
            tag_push_events=True,
            token="secret-token",
            url="https://example.com/hook-%d")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        You can import a system hook using the hook id `{hook-id}`, e.g.

        ```sh
        $ pulumi import gitlab:index/systemHook:SystemHook example 42
        ```

        NOTE: the `token` attribute won't be available for imported resources.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_ssl_verification: Do SSL verification when triggering the hook.
        :param pulumi.Input[bool] merge_requests_events: Trigger hook on merge requests events.
        :param pulumi.Input[bool] push_events: When true, the hook fires on push events.
        :param pulumi.Input[bool] repository_update_events: Trigger hook on repository update events.
        :param pulumi.Input[bool] tag_push_events: When true, the hook fires on new tags being pushed.
        :param pulumi.Input[str] token: Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        :param pulumi.Input[str] url: The hook URL.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemHookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `SystemHook` resource allows to manage the lifecycle of a system hook.

        > This resource requires GitLab 14.9

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/system_hooks.html)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        example = gitlab.SystemHook("example",
            enable_ssl_verification=True,
            merge_requests_events=True,
            push_events=True,
            repository_update_events=True,
            tag_push_events=True,
            token="secret-token",
            url="https://example.com/hook-%d")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

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
                 enable_ssl_verification: Optional[pulumi.Input[bool]] = None,
                 merge_requests_events: Optional[pulumi.Input[bool]] = None,
                 push_events: Optional[pulumi.Input[bool]] = None,
                 repository_update_events: Optional[pulumi.Input[bool]] = None,
                 tag_push_events: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
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
            created_at: Optional[pulumi.Input[str]] = None,
            enable_ssl_verification: Optional[pulumi.Input[bool]] = None,
            merge_requests_events: Optional[pulumi.Input[bool]] = None,
            push_events: Optional[pulumi.Input[bool]] = None,
            repository_update_events: Optional[pulumi.Input[bool]] = None,
            tag_push_events: Optional[pulumi.Input[bool]] = None,
            token: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'SystemHook':
        """
        Get an existing SystemHook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date and time the hook was created in ISO8601 format.
        :param pulumi.Input[bool] enable_ssl_verification: Do SSL verification when triggering the hook.
        :param pulumi.Input[bool] merge_requests_events: Trigger hook on merge requests events.
        :param pulumi.Input[bool] push_events: When true, the hook fires on push events.
        :param pulumi.Input[bool] repository_update_events: Trigger hook on repository update events.
        :param pulumi.Input[bool] tag_push_events: When true, the hook fires on new tags being pushed.
        :param pulumi.Input[str] token: Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        :param pulumi.Input[str] url: The hook URL.
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

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time the hook was created in ISO8601 format.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="enableSslVerification")
    def enable_ssl_verification(self) -> pulumi.Output[Optional[bool]]:
        """
        Do SSL verification when triggering the hook.
        """
        return pulumi.get(self, "enable_ssl_verification")

    @property
    @pulumi.getter(name="mergeRequestsEvents")
    def merge_requests_events(self) -> pulumi.Output[Optional[bool]]:
        """
        Trigger hook on merge requests events.
        """
        return pulumi.get(self, "merge_requests_events")

    @property
    @pulumi.getter(name="pushEvents")
    def push_events(self) -> pulumi.Output[Optional[bool]]:
        """
        When true, the hook fires on push events.
        """
        return pulumi.get(self, "push_events")

    @property
    @pulumi.getter(name="repositoryUpdateEvents")
    def repository_update_events(self) -> pulumi.Output[Optional[bool]]:
        """
        Trigger hook on repository update events.
        """
        return pulumi.get(self, "repository_update_events")

    @property
    @pulumi.getter(name="tagPushEvents")
    def tag_push_events(self) -> pulumi.Output[Optional[bool]]:
        """
        When true, the hook fires on new tags being pushed.
        """
        return pulumi.get(self, "tag_push_events")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[Optional[str]]:
        """
        Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The hook URL.
        """
        return pulumi.get(self, "url")

