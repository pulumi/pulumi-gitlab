# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
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
    def __init__(__self__, avatar_url=None, bio=None, can_create_group=None, can_create_project=None, color_scheme_id=None, created_at=None, current_sign_in_at=None, email=None, extern_uid=None, external=None, id=None, is_admin=None, last_sign_in_at=None, linkedin=None, location=None, name=None, note=None, organization=None, projects_limit=None, skype=None, state=None, theme_id=None, twitter=None, two_factor_enabled=None, user_id=None, user_provider=None, username=None, website_url=None):
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
        return pulumi.get(self, "avatar_url")

    @property
    @pulumi.getter
    def bio(self) -> str:
        return pulumi.get(self, "bio")

    @property
    @pulumi.getter(name="canCreateGroup")
    def can_create_group(self) -> bool:
        return pulumi.get(self, "can_create_group")

    @property
    @pulumi.getter(name="canCreateProject")
    def can_create_project(self) -> bool:
        return pulumi.get(self, "can_create_project")

    @property
    @pulumi.getter(name="colorSchemeId")
    def color_scheme_id(self) -> int:
        return pulumi.get(self, "color_scheme_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currentSignInAt")
    def current_sign_in_at(self) -> str:
        return pulumi.get(self, "current_sign_in_at")

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="externUid")
    def extern_uid(self) -> str:
        return pulumi.get(self, "extern_uid")

    @property
    @pulumi.getter
    def external(self) -> bool:
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
        return pulumi.get(self, "is_admin")

    @property
    @pulumi.getter(name="lastSignInAt")
    def last_sign_in_at(self) -> str:
        return pulumi.get(self, "last_sign_in_at")

    @property
    @pulumi.getter
    def linkedin(self) -> str:
        return pulumi.get(self, "linkedin")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def note(self) -> str:
        return pulumi.get(self, "note")

    @property
    @pulumi.getter
    def organization(self) -> str:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="projectsLimit")
    def projects_limit(self) -> int:
        return pulumi.get(self, "projects_limit")

    @property
    @pulumi.getter
    def skype(self) -> str:
        return pulumi.get(self, "skype")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="themeId")
    def theme_id(self) -> int:
        return pulumi.get(self, "theme_id")

    @property
    @pulumi.getter
    def twitter(self) -> str:
        return pulumi.get(self, "twitter")

    @property
    @pulumi.getter(name="twoFactorEnabled")
    def two_factor_enabled(self) -> bool:
        return pulumi.get(self, "two_factor_enabled")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> int:
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userProvider")
    def user_provider(self) -> str:
        return pulumi.get(self, "user_provider")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="websiteUrl")
    def website_url(self) -> str:
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
            last_sign_in_at=self.last_sign_in_at,
            linkedin=self.linkedin,
            location=self.location,
            name=self.name,
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
             user_id: Optional[int] = None,
             username: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    The `User` data source allows details of a user to be retrieved by either the user ID, username or email address.

    > Some attributes might not be returned depending on if you're an admin or not.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#single-user)
    """
    __args__ = dict()
    __args__['email'] = email
    __args__['userId'] = user_id
    __args__['username'] = username
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gitlab:index/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        avatar_url=__ret__.avatar_url,
        bio=__ret__.bio,
        can_create_group=__ret__.can_create_group,
        can_create_project=__ret__.can_create_project,
        color_scheme_id=__ret__.color_scheme_id,
        created_at=__ret__.created_at,
        current_sign_in_at=__ret__.current_sign_in_at,
        email=__ret__.email,
        extern_uid=__ret__.extern_uid,
        external=__ret__.external,
        id=__ret__.id,
        is_admin=__ret__.is_admin,
        last_sign_in_at=__ret__.last_sign_in_at,
        linkedin=__ret__.linkedin,
        location=__ret__.location,
        name=__ret__.name,
        note=__ret__.note,
        organization=__ret__.organization,
        projects_limit=__ret__.projects_limit,
        skype=__ret__.skype,
        state=__ret__.state,
        theme_id=__ret__.theme_id,
        twitter=__ret__.twitter,
        two_factor_enabled=__ret__.two_factor_enabled,
        user_id=__ret__.user_id,
        user_provider=__ret__.user_provider,
        username=__ret__.username,
        website_url=__ret__.website_url)


@_utilities.lift_output_func(get_user)
def get_user_output(email: Optional[pulumi.Input[Optional[str]]] = None,
                    user_id: Optional[pulumi.Input[Optional[int]]] = None,
                    username: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    The `User` data source allows details of a user to be retrieved by either the user ID, username or email address.

    > Some attributes might not be returned depending on if you're an admin or not.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#single-user)
    """
    ...
