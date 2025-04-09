# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['IntegrationGithubArgs', 'IntegrationGithub']

@pulumi.input_type
class IntegrationGithubArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[builtins.str],
                 repository_url: pulumi.Input[builtins.str],
                 token: pulumi.Input[builtins.str],
                 static_context: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a IntegrationGithub resource.
        :param pulumi.Input[builtins.str] project: ID of the project you want to activate integration on.
        :param pulumi.Input[builtins.str] token: A GitHub personal access token with at least `repo:status` scope.
        :param pulumi.Input[builtins.bool] static_context: Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "repository_url", repository_url)
        pulumi.set(__self__, "token", token)
        if static_context is not None:
            pulumi.set(__self__, "static_context", static_context)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[builtins.str]:
        """
        ID of the project you want to activate integration on.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "repository_url")

    @repository_url.setter
    def repository_url(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "repository_url", value)

    @property
    @pulumi.getter
    def token(self) -> pulumi.Input[builtins.str]:
        """
        A GitHub personal access token with at least `repo:status` scope.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="staticContext")
    def static_context(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
        """
        return pulumi.get(self, "static_context")

    @static_context.setter
    def static_context(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "static_context", value)


@pulumi.input_type
class _IntegrationGithubState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[builtins.bool]] = None,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 repository_url: Optional[pulumi.Input[builtins.str]] = None,
                 static_context: Optional[pulumi.Input[builtins.bool]] = None,
                 title: Optional[pulumi.Input[builtins.str]] = None,
                 token: Optional[pulumi.Input[builtins.str]] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering IntegrationGithub resources.
        :param pulumi.Input[builtins.bool] active: Whether the integration is active.
        :param pulumi.Input[builtins.str] created_at: Create time.
        :param pulumi.Input[builtins.str] project: ID of the project you want to activate integration on.
        :param pulumi.Input[builtins.bool] static_context: Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
        :param pulumi.Input[builtins.str] title: Title.
        :param pulumi.Input[builtins.str] token: A GitHub personal access token with at least `repo:status` scope.
        :param pulumi.Input[builtins.str] updated_at: Update time.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if repository_url is not None:
            pulumi.set(__self__, "repository_url", repository_url)
        if static_context is not None:
            pulumi.set(__self__, "static_context", static_context)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether the integration is active.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Create time.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the project you want to activate integration on.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "repository_url")

    @repository_url.setter
    def repository_url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repository_url", value)

    @property
    @pulumi.getter(name="staticContext")
    def static_context(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
        """
        return pulumi.get(self, "static_context")

    @static_context.setter
    def static_context(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "static_context", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Title.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A GitHub personal access token with at least `repo:status` scope.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Update time.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)


class IntegrationGithub(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 repository_url: Optional[pulumi.Input[builtins.str]] = None,
                 static_context: Optional[pulumi.Input[builtins.bool]] = None,
                 token: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `IntegrationGithub` resource allows to manage the lifecycle of a project integration with GitHub.

        > This resource requires a GitLab Enterprise instance.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/integrations/#github)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        awesome_project = gitlab.Project("awesome_project",
            name="awesome_project",
            description="My awesome project.",
            visibility_level="public")
        github = gitlab.IntegrationGithub("github",
            project=awesome_project.id,
            token="REDACTED",
            repository_url="https://github.com/gitlabhq/terraform-provider-gitlab")
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_github`. For example:

        terraform

        import {

          to = gitlab_integration_github.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        ```sh
        $ pulumi import gitlab:index/integrationGithub:IntegrationGithub You can import a gitlab_integration_github state using `<resource> <project_id>`:
        ```

        ```sh
        $ pulumi import gitlab:index/integrationGithub:IntegrationGithub github 1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] project: ID of the project you want to activate integration on.
        :param pulumi.Input[builtins.bool] static_context: Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
        :param pulumi.Input[builtins.str] token: A GitHub personal access token with at least `repo:status` scope.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationGithubArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `IntegrationGithub` resource allows to manage the lifecycle of a project integration with GitHub.

        > This resource requires a GitLab Enterprise instance.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/integrations/#github)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        awesome_project = gitlab.Project("awesome_project",
            name="awesome_project",
            description="My awesome project.",
            visibility_level="public")
        github = gitlab.IntegrationGithub("github",
            project=awesome_project.id,
            token="REDACTED",
            repository_url="https://github.com/gitlabhq/terraform-provider-gitlab")
        ```

        ## Import

        Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_github`. For example:

        terraform

        import {

          to = gitlab_integration_github.example

          id = "see CLI command below for ID"

        }

        Import using the CLI is supported using the following syntax:

        ```sh
        $ pulumi import gitlab:index/integrationGithub:IntegrationGithub You can import a gitlab_integration_github state using `<resource> <project_id>`:
        ```

        ```sh
        $ pulumi import gitlab:index/integrationGithub:IntegrationGithub github 1
        ```

        :param str resource_name: The name of the resource.
        :param IntegrationGithubArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntegrationGithubArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[builtins.str]] = None,
                 repository_url: Optional[pulumi.Input[builtins.str]] = None,
                 static_context: Optional[pulumi.Input[builtins.bool]] = None,
                 token: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntegrationGithubArgs.__new__(IntegrationGithubArgs)

            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if repository_url is None and not opts.urn:
                raise TypeError("Missing required property 'repository_url'")
            __props__.__dict__["repository_url"] = repository_url
            __props__.__dict__["static_context"] = static_context
            if token is None and not opts.urn:
                raise TypeError("Missing required property 'token'")
            __props__.__dict__["token"] = None if token is None else pulumi.Output.secret(token)
            __props__.__dict__["active"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["title"] = None
            __props__.__dict__["updated_at"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(IntegrationGithub, __self__).__init__(
            'gitlab:index/integrationGithub:IntegrationGithub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[builtins.bool]] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            project: Optional[pulumi.Input[builtins.str]] = None,
            repository_url: Optional[pulumi.Input[builtins.str]] = None,
            static_context: Optional[pulumi.Input[builtins.bool]] = None,
            title: Optional[pulumi.Input[builtins.str]] = None,
            token: Optional[pulumi.Input[builtins.str]] = None,
            updated_at: Optional[pulumi.Input[builtins.str]] = None) -> 'IntegrationGithub':
        """
        Get an existing IntegrationGithub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] active: Whether the integration is active.
        :param pulumi.Input[builtins.str] created_at: Create time.
        :param pulumi.Input[builtins.str] project: ID of the project you want to activate integration on.
        :param pulumi.Input[builtins.bool] static_context: Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
        :param pulumi.Input[builtins.str] title: Title.
        :param pulumi.Input[builtins.str] token: A GitHub personal access token with at least `repo:status` scope.
        :param pulumi.Input[builtins.str] updated_at: Update time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IntegrationGithubState.__new__(_IntegrationGithubState)

        __props__.__dict__["active"] = active
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["project"] = project
        __props__.__dict__["repository_url"] = repository_url
        __props__.__dict__["static_context"] = static_context
        __props__.__dict__["title"] = title
        __props__.__dict__["token"] = token
        __props__.__dict__["updated_at"] = updated_at
        return IntegrationGithub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[builtins.bool]:
        """
        Whether the integration is active.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        Create time.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[builtins.str]:
        """
        ID of the project you want to activate integration on.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "repository_url")

    @property
    @pulumi.getter(name="staticContext")
    def static_context(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
        """
        return pulumi.get(self, "static_context")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[builtins.str]:
        """
        Title.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[builtins.str]:
        """
        A GitHub personal access token with at least `repo:status` scope.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        Update time.
        """
        return pulumi.get(self, "updated_at")

