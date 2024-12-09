# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
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

__all__ = [
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, avatar_url=None, bio=None, can_create_group=None, can_create_project=None, color_scheme_id=None, created_at=None, current_sign_in_at=None, email=None, extern_uid=None, external=None, id=None, is_admin=None, is_bot=None, last_sign_in_at=None, linkedin=None, location=None, name=None, namespace_id=None, note=None, organization=None, projects_limit=None, skype=None, state=None, theme_id=None, twitter=None, two_factor_enabled=None, user_id=None, user_provider=None, username=None, website_url=None):
        if avatar_url and not isinstance(avatar_url, str):
            raise TypeError("Expected argument 'avatar_url' to be a str")
        pulumi.set(__self__, "avatar_url", avatar_url)
        if bio and not isinstance(bio, str):
            raise TypeError("Expected argument 'bio' to be a str")
        pulumi.set(__self__, "bio", bio)
        if can_create_group and not isinstance(can_create_group, bool):
            raise TypeError("Expected argument 'can_create_group' to be a bool")
        pulumi.set(__self__, "can_create_group", can_create_group)
        if can_create_project and not isinstance(can_create_project, bool):
            raise TypeError("Expected argument 'can_create_project' to be a bool")
        pulumi.set(__self__, "can_create_project", can_create_project)
        if color_scheme_id and not isinstance(color_scheme_id, int):
            raise TypeError("Expected argument 'color_scheme_id' to be a int")
        pulumi.set(__self__, "color_scheme_id", color_scheme_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if current_sign_in_at and not isinstance(current_sign_in_at, str):
            raise TypeError("Expected argument 'current_sign_in_at' to be a str")
        pulumi.set(__self__, "current_sign_in_at", current_sign_in_at)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if extern_uid and not isinstance(extern_uid, str):
            raise TypeError("Expected argument 'extern_uid' to be a str")
        pulumi.set(__self__, "extern_uid", extern_uid)
        if external and not isinstance(external, bool):
            raise TypeError("Expected argument 'external' to be a bool")
        pulumi.set(__self__, "external", external)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_admin and not isinstance(is_admin, bool):
            raise TypeError("Expected argument 'is_admin' to be a bool")
        pulumi.set(__self__, "is_admin", is_admin)
        if is_bot and not isinstance(is_bot, bool):
            raise TypeError("Expected argument 'is_bot' to be a bool")
        pulumi.set(__self__, "is_bot", is_bot)
        if last_sign_in_at and not isinstance(last_sign_in_at, str):
            raise TypeError("Expected argument 'last_sign_in_at' to be a str")
        pulumi.set(__self__, "last_sign_in_at", last_sign_in_at)
        if linkedin and not isinstance(linkedin, str):
            raise TypeError("Expected argument 'linkedin' to be a str")
        pulumi.set(__self__, "linkedin", linkedin)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace_id and not isinstance(namespace_id, int):
            raise TypeError("Expected argument 'namespace_id' to be a int")
        pulumi.set(__self__, "namespace_id", namespace_id)
        if note and not isinstance(note, str):
            raise TypeError("Expected argument 'note' to be a str")
        pulumi.set(__self__, "note", note)
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        pulumi.set(__self__, "organization", organization)
        if projects_limit and not isinstance(projects_limit, int):
            raise TypeError("Expected argument 'projects_limit' to be a int")
        pulumi.set(__self__, "projects_limit", projects_limit)
        if skype and not isinstance(skype, str):
            raise TypeError("Expected argument 'skype' to be a str")
        pulumi.set(__self__, "skype", skype)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if theme_id and not isinstance(theme_id, int):
            raise TypeError("Expected argument 'theme_id' to be a int")
        pulumi.set(__self__, "theme_id", theme_id)
        if twitter and not isinstance(twitter, str):
            raise TypeError("Expected argument 'twitter' to be a str")
        pulumi.set(__self__, "twitter", twitter)
        if two_factor_enabled and not isinstance(two_factor_enabled, bool):
            raise TypeError("Expected argument 'two_factor_enabled' to be a bool")
        pulumi.set(__self__, "two_factor_enabled", two_factor_enabled)
        if user_id and not isinstance(user_id, int):
            raise TypeError("Expected argument 'user_id' to be a int")
        pulumi.set(__self__, "user_id", user_id)
        if user_provider and not isinstance(user_provider, str):
            raise TypeError("Expected argument 'user_provider' to be a str")
        pulumi.set(__self__, "user_provider", user_provider)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if website_url and not isinstance(website_url, str):
            raise TypeError("Expected argument 'website_url' to be a str")
        pulumi.set(__self__, "website_url", website_url)

    @property
    @pulumi.getter(name="avatarUrl")
    def avatar_url(self) -> str:
        """
        The avatar URL of the user.
        """
        return pulumi.get(self, "avatar_url")

    @property
    @pulumi.getter
    def bio(self) -> str:
        """
        The bio of the user.
        """
        return pulumi.get(self, "bio")

    @property
    @pulumi.getter(name="canCreateGroup")
    def can_create_group(self) -> bool:
        """
        Whether the user can create groups.
        """
        return pulumi.get(self, "can_create_group")

    @property
    @pulumi.getter(name="canCreateProject")
    def can_create_project(self) -> bool:
        """
        Whether the user can create projects.
        """
        return pulumi.get(self, "can_create_project")

    @property
    @pulumi.getter(name="colorSchemeId")
    def color_scheme_id(self) -> int:
        """
        User's color scheme ID.
        """
        return pulumi.get(self, "color_scheme_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date the user was created at.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currentSignInAt")
    def current_sign_in_at(self) -> str:
        """
        Current user's sign-in date.
        """
        return pulumi.get(self, "current_sign_in_at")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        The public email address of the user. **Note**: before GitLab 14.8 the lookup was based on the users primary email address.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="externUid")
    def extern_uid(self) -> str:
        """
        The external UID of the user.
        """
        return pulumi.get(self, "extern_uid")

    @property
    @pulumi.getter
    def external(self) -> bool:
        """
        Whether the user is external.
        """
        return pulumi.get(self, "external")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> bool:
        """
        Whether the user is an admin.
        """
        return pulumi.get(self, "is_admin")

    @property
    @pulumi.getter(name="isBot")
    def is_bot(self) -> bool:
        """
        Whether the user is a bot.
        """
        return pulumi.get(self, "is_bot")

    @property
    @pulumi.getter(name="lastSignInAt")
    def last_sign_in_at(self) -> str:
        """
        Last user's sign-in date.
        """
        return pulumi.get(self, "last_sign_in_at")

    @property
    @pulumi.getter
    def linkedin(self) -> str:
        """
        LinkedIn profile of the user.
        """
        return pulumi.get(self, "linkedin")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the user.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> int:
        """
        The ID of the user's namespace. Requires admin token to access this field. Available since GitLab 14.10.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter
    def note(self) -> str:
        """
        Admin notes for this user.
        """
        return pulumi.get(self, "note")

    @property
    @pulumi.getter
    def organization(self) -> str:
        """
        The organization of the user.
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="projectsLimit")
    def projects_limit(self) -> int:
        """
        Number of projects the user can create.
        """
        return pulumi.get(self, "projects_limit")

    @property
    @pulumi.getter
    def skype(self) -> str:
        """
        Skype username of the user.
        """
        return pulumi.get(self, "skype")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Whether the user is active or blocked.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="themeId")
    def theme_id(self) -> int:
        """
        User's theme ID.
        """
        return pulumi.get(self, "theme_id")

    @property
    @pulumi.getter
    def twitter(self) -> str:
        """
        Twitter username of the user.
        """
        return pulumi.get(self, "twitter")

    @property
    @pulumi.getter(name="twoFactorEnabled")
    def two_factor_enabled(self) -> bool:
        """
        Whether user's two-factor auth is enabled.
        """
        return pulumi.get(self, "two_factor_enabled")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> int:
        """
        The ID of the user.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userProvider")
    def user_provider(self) -> str:
        """
        The UID provider of the user.
        """
        return pulumi.get(self, "user_provider")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        The username of the user.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="websiteUrl")
    def website_url(self) -> str:
        """
        User's website URL.
        """
        return pulumi.get(self, "website_url")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            avatar_url=self.avatar_url,
            bio=self.bio,
            can_create_group=self.can_create_group,
            can_create_project=self.can_create_project,
            color_scheme_id=self.color_scheme_id,
            created_at=self.created_at,
            current_sign_in_at=self.current_sign_in_at,
            email=self.email,
            extern_uid=self.extern_uid,
            external=self.external,
            id=self.id,
            is_admin=self.is_admin,
            is_bot=self.is_bot,
            last_sign_in_at=self.last_sign_in_at,
            linkedin=self.linkedin,
            location=self.location,
            name=self.name,
            namespace_id=self.namespace_id,
            note=self.note,
            organization=self.organization,
            projects_limit=self.projects_limit,
            skype=self.skype,
            state=self.state,
            theme_id=self.theme_id,
            twitter=self.twitter,
            two_factor_enabled=self.two_factor_enabled,
            user_id=self.user_id,
            user_provider=self.user_provider,
            username=self.username,
            website_url=self.website_url)


def get_user(email: Optional[str] = None,
             namespace_id: Optional[int] = None,
             user_id: Optional[int] = None,
             username: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    The `User` data source allows details of a user to be retrieved by either the user ID, username or email address.

    > Some attributes might not be returned depending on if you're an admin or not.

    > When using the `email` attribute, an exact match is not guaranteed. The most related match will be returned. Starting with GitLab 16.6,
    the most related match will prioritize an exact match if one is available.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#single-user)


    :param str email: The public email address of the user. **Note**: before GitLab 14.8 the lookup was based on the users primary email address.
    :param int namespace_id: The ID of the user's namespace. Requires admin token to access this field. Available since GitLab 14.10.
    :param int user_id: The ID of the user.
    :param str username: The username of the user.
    """
    __args__ = dict()
    __args__['email'] = email
    __args__['namespaceId'] = namespace_id
    __args__['userId'] = user_id
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        avatar_url=pulumi.get(__ret__, 'avatar_url'),
        bio=pulumi.get(__ret__, 'bio'),
        can_create_group=pulumi.get(__ret__, 'can_create_group'),
        can_create_project=pulumi.get(__ret__, 'can_create_project'),
        color_scheme_id=pulumi.get(__ret__, 'color_scheme_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        current_sign_in_at=pulumi.get(__ret__, 'current_sign_in_at'),
        email=pulumi.get(__ret__, 'email'),
        extern_uid=pulumi.get(__ret__, 'extern_uid'),
        external=pulumi.get(__ret__, 'external'),
        id=pulumi.get(__ret__, 'id'),
        is_admin=pulumi.get(__ret__, 'is_admin'),
        is_bot=pulumi.get(__ret__, 'is_bot'),
        last_sign_in_at=pulumi.get(__ret__, 'last_sign_in_at'),
        linkedin=pulumi.get(__ret__, 'linkedin'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        namespace_id=pulumi.get(__ret__, 'namespace_id'),
        note=pulumi.get(__ret__, 'note'),
        organization=pulumi.get(__ret__, 'organization'),
        projects_limit=pulumi.get(__ret__, 'projects_limit'),
        skype=pulumi.get(__ret__, 'skype'),
        state=pulumi.get(__ret__, 'state'),
        theme_id=pulumi.get(__ret__, 'theme_id'),
        twitter=pulumi.get(__ret__, 'twitter'),
        two_factor_enabled=pulumi.get(__ret__, 'two_factor_enabled'),
        user_id=pulumi.get(__ret__, 'user_id'),
        user_provider=pulumi.get(__ret__, 'user_provider'),
        username=pulumi.get(__ret__, 'username'),
        website_url=pulumi.get(__ret__, 'website_url'))
def get_user_output(email: Optional[pulumi.Input[Optional[str]]] = None,
                    namespace_id: Optional[pulumi.Input[Optional[int]]] = None,
                    user_id: Optional[pulumi.Input[Optional[int]]] = None,
                    username: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUserResult]:
    """
    The `User` data source allows details of a user to be retrieved by either the user ID, username or email address.

    > Some attributes might not be returned depending on if you're an admin or not.

    > When using the `email` attribute, an exact match is not guaranteed. The most related match will be returned. Starting with GitLab 16.6,
    the most related match will prioritize an exact match if one is available.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#single-user)


    :param str email: The public email address of the user. **Note**: before GitLab 14.8 the lookup was based on the users primary email address.
    :param int namespace_id: The ID of the user's namespace. Requires admin token to access this field. Available since GitLab 14.10.
    :param int user_id: The ID of the user.
    :param str username: The username of the user.
    """
    __args__ = dict()
    __args__['email'] = email
    __args__['namespaceId'] = namespace_id
    __args__['userId'] = user_id
    __args__['username'] = username
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('gitlab:index/getUser:getUser', __args__, opts=opts, typ=GetUserResult)
    return __ret__.apply(lambda __response__: GetUserResult(
        avatar_url=pulumi.get(__response__, 'avatar_url'),
        bio=pulumi.get(__response__, 'bio'),
        can_create_group=pulumi.get(__response__, 'can_create_group'),
        can_create_project=pulumi.get(__response__, 'can_create_project'),
        color_scheme_id=pulumi.get(__response__, 'color_scheme_id'),
        created_at=pulumi.get(__response__, 'created_at'),
        current_sign_in_at=pulumi.get(__response__, 'current_sign_in_at'),
        email=pulumi.get(__response__, 'email'),
        extern_uid=pulumi.get(__response__, 'extern_uid'),
        external=pulumi.get(__response__, 'external'),
        id=pulumi.get(__response__, 'id'),
        is_admin=pulumi.get(__response__, 'is_admin'),
        is_bot=pulumi.get(__response__, 'is_bot'),
        last_sign_in_at=pulumi.get(__response__, 'last_sign_in_at'),
        linkedin=pulumi.get(__response__, 'linkedin'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        namespace_id=pulumi.get(__response__, 'namespace_id'),
        note=pulumi.get(__response__, 'note'),
        organization=pulumi.get(__response__, 'organization'),
        projects_limit=pulumi.get(__response__, 'projects_limit'),
        skype=pulumi.get(__response__, 'skype'),
        state=pulumi.get(__response__, 'state'),
        theme_id=pulumi.get(__response__, 'theme_id'),
        twitter=pulumi.get(__response__, 'twitter'),
        two_factor_enabled=pulumi.get(__response__, 'two_factor_enabled'),
        user_id=pulumi.get(__response__, 'user_id'),
        user_provider=pulumi.get(__response__, 'user_provider'),
        username=pulumi.get(__response__, 'username'),
        website_url=pulumi.get(__response__, 'website_url')))
