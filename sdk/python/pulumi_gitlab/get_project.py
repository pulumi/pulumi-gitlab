# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, archived=None, default_branch=None, description=None, http_url_to_repo=None, id=None, issues_enabled=None, merge_requests_enabled=None, name=None, namespace_id=None, path=None, runners_token=None, snippets_enabled=None, ssh_url_to_repo=None, visibility_level=None, web_url=None, wiki_enabled=None):
        if archived and not isinstance(archived, bool):
            raise TypeError("Expected argument 'archived' to be a bool")
        __self__.archived = archived
        """
        Whether the project is in read-only mode (archived).
        """
        if default_branch and not isinstance(default_branch, str):
            raise TypeError("Expected argument 'default_branch' to be a str")
        __self__.default_branch = default_branch
        """
        The default branch for the project.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        A description of the project.
        """
        if http_url_to_repo and not isinstance(http_url_to_repo, str):
            raise TypeError("Expected argument 'http_url_to_repo' to be a str")
        __self__.http_url_to_repo = http_url_to_repo
        """
        URL that can be provided to `git clone` to clone the
        repository via HTTP.
        """
        if id and not isinstance(id, float):
            raise TypeError("Expected argument 'id' to be a float")
        __self__.id = id
        """
        Integer that uniquely identifies the project within the gitlab install.
        """
        if issues_enabled and not isinstance(issues_enabled, bool):
            raise TypeError("Expected argument 'issues_enabled' to be a bool")
        __self__.issues_enabled = issues_enabled
        """
        Enable issue tracking for the project.
        """
        if merge_requests_enabled and not isinstance(merge_requests_enabled, bool):
            raise TypeError("Expected argument 'merge_requests_enabled' to be a bool")
        __self__.merge_requests_enabled = merge_requests_enabled
        """
        Enable merge requests for the project.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if namespace_id and not isinstance(namespace_id, float):
            raise TypeError("Expected argument 'namespace_id' to be a float")
        __self__.namespace_id = namespace_id
        """
        The namespace (group or user) of the project. Defaults to your user.
        See `gitlab_group` for an example.
        """
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
        """
        The path of the repository.
        """
        if runners_token and not isinstance(runners_token, str):
            raise TypeError("Expected argument 'runners_token' to be a str")
        __self__.runners_token = runners_token
        """
        Registration token to use during runner setup.
        """
        if snippets_enabled and not isinstance(snippets_enabled, bool):
            raise TypeError("Expected argument 'snippets_enabled' to be a bool")
        __self__.snippets_enabled = snippets_enabled
        """
        Enable snippets for the project.
        """
        if ssh_url_to_repo and not isinstance(ssh_url_to_repo, str):
            raise TypeError("Expected argument 'ssh_url_to_repo' to be a str")
        __self__.ssh_url_to_repo = ssh_url_to_repo
        """
        URL that can be provided to `git clone` to clone the
        repository via SSH.
        """
        if visibility_level and not isinstance(visibility_level, str):
            raise TypeError("Expected argument 'visibility_level' to be a str")
        __self__.visibility_level = visibility_level
        """
        Repositories are created as private by default.
        """
        if web_url and not isinstance(web_url, str):
            raise TypeError("Expected argument 'web_url' to be a str")
        __self__.web_url = web_url
        """
        URL that can be used to find the project in a browser.
        """
        if wiki_enabled and not isinstance(wiki_enabled, bool):
            raise TypeError("Expected argument 'wiki_enabled' to be a bool")
        __self__.wiki_enabled = wiki_enabled
        """
        Enable wiki for the project.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return self

    __iter__ = __await__

def get_project(archived=None,default_branch=None,description=None,http_url_to_repo=None,id=None,issues_enabled=None,merge_requests_enabled=None,name=None,namespace_id=None,path=None,runners_token=None,snippets_enabled=None,ssh_url_to_repo=None,visibility_level=None,web_url=None,wiki_enabled=None,opts=None):
    """
    Provides details about a specific project in the gitlab provider. The results include the name of the project, path, description, default branch, etc.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/d/project.html.markdown.
    """
    __args__ = dict()

    __args__['archived'] = archived
    __args__['defaultBranch'] = default_branch
    __args__['description'] = description
    __args__['httpUrlToRepo'] = http_url_to_repo
    __args__['id'] = id
    __args__['issuesEnabled'] = issues_enabled
    __args__['mergeRequestsEnabled'] = merge_requests_enabled
    __args__['name'] = name
    __args__['namespaceId'] = namespace_id
    __args__['path'] = path
    __args__['runnersToken'] = runners_token
    __args__['snippetsEnabled'] = snippets_enabled
    __args__['sshUrlToRepo'] = ssh_url_to_repo
    __args__['visibilityLevel'] = visibility_level
    __args__['webUrl'] = web_url
    __args__['wikiEnabled'] = wiki_enabled
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProject:getProject', __args__, opts=opts).value

    return GetProjectResult(
        archived=__ret__.get('archived'),
        default_branch=__ret__.get('defaultBranch'),
        description=__ret__.get('description'),
        http_url_to_repo=__ret__.get('httpUrlToRepo'),
        id=__ret__.get('id'),
        issues_enabled=__ret__.get('issuesEnabled'),
        merge_requests_enabled=__ret__.get('mergeRequestsEnabled'),
        name=__ret__.get('name'),
        namespace_id=__ret__.get('namespaceId'),
        path=__ret__.get('path'),
        runners_token=__ret__.get('runnersToken'),
        snippets_enabled=__ret__.get('snippetsEnabled'),
        ssh_url_to_repo=__ret__.get('sshUrlToRepo'),
        visibility_level=__ret__.get('visibilityLevel'),
        web_url=__ret__.get('webUrl'),
        wiki_enabled=__ret__.get('wikiEnabled'))
