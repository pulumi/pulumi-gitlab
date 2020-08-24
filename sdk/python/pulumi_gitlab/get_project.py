# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = [
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
]

@pulumi.output_type
class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, archived=None, default_branch=None, description=None, http_url_to_repo=None, id=None, issues_enabled=None, lfs_enabled=None, merge_requests_enabled=None, name=None, namespace_id=None, path=None, pipelines_enabled=None, remove_source_branch_after_merge=None, request_access_enabled=None, runners_token=None, snippets_enabled=None, ssh_url_to_repo=None, visibility_level=None, web_url=None, wiki_enabled=None):
        if archived and not isinstance(archived, bool):
            raise TypeError("Expected argument 'archived' to be a bool")
        pulumi.set(__self__, "archived", archived)
        if default_branch and not isinstance(default_branch, str):
            raise TypeError("Expected argument 'default_branch' to be a str")
        pulumi.set(__self__, "default_branch", default_branch)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if http_url_to_repo and not isinstance(http_url_to_repo, str):
            raise TypeError("Expected argument 'http_url_to_repo' to be a str")
        pulumi.set(__self__, "http_url_to_repo", http_url_to_repo)
        if id and not isinstance(id, float):
            raise TypeError("Expected argument 'id' to be a float")
        pulumi.set(__self__, "id", id)
        if issues_enabled and not isinstance(issues_enabled, bool):
            raise TypeError("Expected argument 'issues_enabled' to be a bool")
        pulumi.set(__self__, "issues_enabled", issues_enabled)
        if lfs_enabled and not isinstance(lfs_enabled, bool):
            raise TypeError("Expected argument 'lfs_enabled' to be a bool")
        pulumi.set(__self__, "lfs_enabled", lfs_enabled)
        if merge_requests_enabled and not isinstance(merge_requests_enabled, bool):
            raise TypeError("Expected argument 'merge_requests_enabled' to be a bool")
        pulumi.set(__self__, "merge_requests_enabled", merge_requests_enabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace_id and not isinstance(namespace_id, float):
            raise TypeError("Expected argument 'namespace_id' to be a float")
        pulumi.set(__self__, "namespace_id", namespace_id)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if pipelines_enabled and not isinstance(pipelines_enabled, bool):
            raise TypeError("Expected argument 'pipelines_enabled' to be a bool")
        pulumi.set(__self__, "pipelines_enabled", pipelines_enabled)
        if remove_source_branch_after_merge and not isinstance(remove_source_branch_after_merge, bool):
            raise TypeError("Expected argument 'remove_source_branch_after_merge' to be a bool")
        pulumi.set(__self__, "remove_source_branch_after_merge", remove_source_branch_after_merge)
        if request_access_enabled and not isinstance(request_access_enabled, bool):
            raise TypeError("Expected argument 'request_access_enabled' to be a bool")
        pulumi.set(__self__, "request_access_enabled", request_access_enabled)
        if runners_token and not isinstance(runners_token, str):
            raise TypeError("Expected argument 'runners_token' to be a str")
        pulumi.set(__self__, "runners_token", runners_token)
        if snippets_enabled and not isinstance(snippets_enabled, bool):
            raise TypeError("Expected argument 'snippets_enabled' to be a bool")
        pulumi.set(__self__, "snippets_enabled", snippets_enabled)
        if ssh_url_to_repo and not isinstance(ssh_url_to_repo, str):
            raise TypeError("Expected argument 'ssh_url_to_repo' to be a str")
        pulumi.set(__self__, "ssh_url_to_repo", ssh_url_to_repo)
        if visibility_level and not isinstance(visibility_level, str):
            raise TypeError("Expected argument 'visibility_level' to be a str")
        pulumi.set(__self__, "visibility_level", visibility_level)
        if web_url and not isinstance(web_url, str):
            raise TypeError("Expected argument 'web_url' to be a str")
        pulumi.set(__self__, "web_url", web_url)
        if wiki_enabled and not isinstance(wiki_enabled, bool):
            raise TypeError("Expected argument 'wiki_enabled' to be a bool")
        pulumi.set(__self__, "wiki_enabled", wiki_enabled)

    @property
    @pulumi.getter
    def archived(self) -> bool:
        """
        Whether the project is in read-only mode (archived).
        """
        return pulumi.get(self, "archived")

    @property
    @pulumi.getter(name="defaultBranch")
    def default_branch(self) -> str:
        """
        The default branch for the project.
        """
        return pulumi.get(self, "default_branch")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the project.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="httpUrlToRepo")
    def http_url_to_repo(self) -> str:
        """
        URL that can be provided to `git clone` to clone the
        repository via HTTP.
        """
        return pulumi.get(self, "http_url_to_repo")

    @property
    @pulumi.getter
    def id(self) -> float:
        """
        Integer that uniquely identifies the project within the gitlab install.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="issuesEnabled")
    def issues_enabled(self) -> bool:
        """
        Enable issue tracking for the project.
        """
        return pulumi.get(self, "issues_enabled")

    @property
    @pulumi.getter(name="lfsEnabled")
    def lfs_enabled(self) -> bool:
        """
        Enable LFS for the project.
        """
        return pulumi.get(self, "lfs_enabled")

    @property
    @pulumi.getter(name="mergeRequestsEnabled")
    def merge_requests_enabled(self) -> bool:
        """
        Enable merge requests for the project.
        """
        return pulumi.get(self, "merge_requests_enabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> float:
        """
        The namespace (group or user) of the project. Defaults to your user.
        See `Group` for an example.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        The path of the repository.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="pipelinesEnabled")
    def pipelines_enabled(self) -> bool:
        """
        Enable pipelines for the project.
        """
        return pulumi.get(self, "pipelines_enabled")

    @property
    @pulumi.getter(name="removeSourceBranchAfterMerge")
    def remove_source_branch_after_merge(self) -> bool:
        """
        Enable `Delete source branch` option by default for all new merge requests
        """
        return pulumi.get(self, "remove_source_branch_after_merge")

    @property
    @pulumi.getter(name="requestAccessEnabled")
    def request_access_enabled(self) -> bool:
        """
        Allow users to request member access.
        """
        return pulumi.get(self, "request_access_enabled")

    @property
    @pulumi.getter(name="runnersToken")
    def runners_token(self) -> str:
        """
        Registration token to use during runner setup.
        """
        return pulumi.get(self, "runners_token")

    @property
    @pulumi.getter(name="snippetsEnabled")
    def snippets_enabled(self) -> bool:
        """
        Enable snippets for the project.
        """
        return pulumi.get(self, "snippets_enabled")

    @property
    @pulumi.getter(name="sshUrlToRepo")
    def ssh_url_to_repo(self) -> str:
        """
        URL that can be provided to `git clone` to clone the
        repository via SSH.
        """
        return pulumi.get(self, "ssh_url_to_repo")

    @property
    @pulumi.getter(name="visibilityLevel")
    def visibility_level(self) -> str:
        """
        Repositories are created as private by default.
        """
        return pulumi.get(self, "visibility_level")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> str:
        """
        URL that can be used to find the project in a browser.
        """
        return pulumi.get(self, "web_url")

    @property
    @pulumi.getter(name="wikiEnabled")
    def wiki_enabled(self) -> bool:
        """
        Enable wiki for the project.
        """
        return pulumi.get(self, "wiki_enabled")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            archived=self.archived,
            default_branch=self.default_branch,
            description=self.description,
            http_url_to_repo=self.http_url_to_repo,
            id=self.id,
            issues_enabled=self.issues_enabled,
            lfs_enabled=self.lfs_enabled,
            merge_requests_enabled=self.merge_requests_enabled,
            name=self.name,
            namespace_id=self.namespace_id,
            path=self.path,
            pipelines_enabled=self.pipelines_enabled,
            remove_source_branch_after_merge=self.remove_source_branch_after_merge,
            request_access_enabled=self.request_access_enabled,
            runners_token=self.runners_token,
            snippets_enabled=self.snippets_enabled,
            ssh_url_to_repo=self.ssh_url_to_repo,
            visibility_level=self.visibility_level,
            web_url=self.web_url,
            wiki_enabled=self.wiki_enabled)


def get_project(archived: Optional[bool] = None,
                default_branch: Optional[str] = None,
                description: Optional[str] = None,
                http_url_to_repo: Optional[str] = None,
                id: Optional[float] = None,
                issues_enabled: Optional[bool] = None,
                lfs_enabled: Optional[bool] = None,
                merge_requests_enabled: Optional[bool] = None,
                name: Optional[str] = None,
                namespace_id: Optional[float] = None,
                path: Optional[str] = None,
                pipelines_enabled: Optional[bool] = None,
                remove_source_branch_after_merge: Optional[bool] = None,
                request_access_enabled: Optional[bool] = None,
                runners_token: Optional[str] = None,
                snippets_enabled: Optional[bool] = None,
                ssh_url_to_repo: Optional[str] = None,
                visibility_level: Optional[str] = None,
                web_url: Optional[str] = None,
                wiki_enabled: Optional[bool] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    Provides details about a specific project in the gitlab provider. The results include the name of the project, path, description, default branch, etc.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_project(id=30)
    ```


    :param bool archived: Whether the project is in read-only mode (archived).
    :param str default_branch: The default branch for the project.
    :param str description: A description of the project.
    :param str http_url_to_repo: URL that can be provided to `git clone` to clone the
           repository via HTTP.
    :param float id: The integer that uniquely identifies the project within the gitlab install.
    :param bool issues_enabled: Enable issue tracking for the project.
    :param bool lfs_enabled: Enable LFS for the project.
    :param bool merge_requests_enabled: Enable merge requests for the project.
    :param float namespace_id: The namespace (group or user) of the project. Defaults to your user.
           See `Group` for an example.
    :param str path: The path of the repository.
    :param bool pipelines_enabled: Enable pipelines for the project.
    :param bool remove_source_branch_after_merge: Enable `Delete source branch` option by default for all new merge requests
    :param bool request_access_enabled: Allow users to request member access.
    :param str runners_token: Registration token to use during runner setup.
    :param bool snippets_enabled: Enable snippets for the project.
    :param str ssh_url_to_repo: URL that can be provided to `git clone` to clone the
           repository via SSH.
    :param str visibility_level: Repositories are created as private by default.
    :param str web_url: URL that can be used to find the project in a browser.
    :param bool wiki_enabled: Enable wiki for the project.
    """
    __args__ = dict()
    __args__['archived'] = archived
    __args__['defaultBranch'] = default_branch
    __args__['description'] = description
    __args__['httpUrlToRepo'] = http_url_to_repo
    __args__['id'] = id
    __args__['issuesEnabled'] = issues_enabled
    __args__['lfsEnabled'] = lfs_enabled
    __args__['mergeRequestsEnabled'] = merge_requests_enabled
    __args__['name'] = name
    __args__['namespaceId'] = namespace_id
    __args__['path'] = path
    __args__['pipelinesEnabled'] = pipelines_enabled
    __args__['removeSourceBranchAfterMerge'] = remove_source_branch_after_merge
    __args__['requestAccessEnabled'] = request_access_enabled
    __args__['runnersToken'] = runners_token
    __args__['snippetsEnabled'] = snippets_enabled
    __args__['sshUrlToRepo'] = ssh_url_to_repo
    __args__['visibilityLevel'] = visibility_level
    __args__['webUrl'] = web_url
    __args__['wikiEnabled'] = wiki_enabled
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProject:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        archived=__ret__.archived,
        default_branch=__ret__.default_branch,
        description=__ret__.description,
        http_url_to_repo=__ret__.http_url_to_repo,
        id=__ret__.id,
        issues_enabled=__ret__.issues_enabled,
        lfs_enabled=__ret__.lfs_enabled,
        merge_requests_enabled=__ret__.merge_requests_enabled,
        name=__ret__.name,
        namespace_id=__ret__.namespace_id,
        path=__ret__.path,
        pipelines_enabled=__ret__.pipelines_enabled,
        remove_source_branch_after_merge=__ret__.remove_source_branch_after_merge,
        request_access_enabled=__ret__.request_access_enabled,
        runners_token=__ret__.runners_token,
        snippets_enabled=__ret__.snippets_enabled,
        ssh_url_to_repo=__ret__.ssh_url_to_repo,
        visibility_level=__ret__.visibility_level,
        web_url=__ret__.web_url,
        wiki_enabled=__ret__.wiki_enabled)
