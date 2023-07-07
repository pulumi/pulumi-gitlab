# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetCurrentUserResult',
    'AwaitableGetCurrentUserResult',
    'get_current_user',
]

@pulumi.output_type
class GetCurrentUserResult:
    """
    A collection of values returned by getCurrentUser.
    """
    def __init__(__self__, bot=None, global_id=None, global_namespace_id=None, group_count=None, id=None, name=None, namespace_id=None, public_email=None, username=None):
        if bot and not isinstance(bot, bool):
            raise TypeError("Expected argument 'bot' to be a bool")
        pulumi.set(__self__, "bot", bot)
        if global_id and not isinstance(global_id, str):
            raise TypeError("Expected argument 'global_id' to be a str")
        pulumi.set(__self__, "global_id", global_id)
        if global_namespace_id and not isinstance(global_namespace_id, str):
            raise TypeError("Expected argument 'global_namespace_id' to be a str")
        pulumi.set(__self__, "global_namespace_id", global_namespace_id)
        if group_count and not isinstance(group_count, int):
            raise TypeError("Expected argument 'group_count' to be a int")
        pulumi.set(__self__, "group_count", group_count)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace_id and not isinstance(namespace_id, str):
            raise TypeError("Expected argument 'namespace_id' to be a str")
        pulumi.set(__self__, "namespace_id", namespace_id)
        if public_email and not isinstance(public_email, str):
            raise TypeError("Expected argument 'public_email' to be a str")
        pulumi.set(__self__, "public_email", public_email)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def bot(self) -> bool:
        """
        Indicates if the user is a bot.
        """
        return pulumi.get(self, "bot")

    @property
    @pulumi.getter(name="globalId")
    def global_id(self) -> str:
        """
        Global ID of the user. This is in the form of a GraphQL globally unique ID.
        """
        return pulumi.get(self, "global_id")

    @property
    @pulumi.getter(name="globalNamespaceId")
    def global_namespace_id(self) -> str:
        """
        Personal namespace of the user. This is in the form of a GraphQL globally unique ID.
        """
        return pulumi.get(self, "global_namespace_id")

    @property
    @pulumi.getter(name="groupCount")
    def group_count(self) -> int:
        """
        Group count for the user.
        """
        return pulumi.get(self, "group_count")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the user.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Human-readable name of the user. Returns **** if the user is a project bot and the requester does not have permission to view the project.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> str:
        """
        Personal namespace of the user.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="publicEmail")
    def public_email(self) -> str:
        """
        User’s public email.
        """
        return pulumi.get(self, "public_email")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        Username of the user. Unique within this instance of GitLab.
        """
        return pulumi.get(self, "username")


class AwaitableGetCurrentUserResult(GetCurrentUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCurrentUserResult(
            bot=self.bot,
            global_id=self.global_id,
            global_namespace_id=self.global_namespace_id,
            group_count=self.group_count,
            id=self.id,
            name=self.name,
            namespace_id=self.namespace_id,
            public_email=self.public_email,
            username=self.username)


def get_current_user(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCurrentUserResult:
    """
    The `get_current_user` data source allows details of the current user (determined by `token` provider attribute) to be retrieved.

    **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/index.html#querycurrentuser)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_current_user()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getCurrentUser:getCurrentUser', __args__, opts=opts, typ=GetCurrentUserResult).value

    return AwaitableGetCurrentUserResult(
        bot=pulumi.get(__ret__, 'bot'),
        global_id=pulumi.get(__ret__, 'global_id'),
        global_namespace_id=pulumi.get(__ret__, 'global_namespace_id'),
        group_count=pulumi.get(__ret__, 'group_count'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        namespace_id=pulumi.get(__ret__, 'namespace_id'),
        public_email=pulumi.get(__ret__, 'public_email'),
        username=pulumi.get(__ret__, 'username'))
