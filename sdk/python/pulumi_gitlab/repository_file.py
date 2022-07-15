# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RepositoryFileArgs', 'RepositoryFile']

@pulumi.input_type
class RepositoryFileArgs:
    def __init__(__self__, *,
                 branch: pulumi.Input[str],
                 commit_message: pulumi.Input[str],
                 content: pulumi.Input[str],
                 file_path: pulumi.Input[str],
                 project: pulumi.Input[str],
                 author_email: Optional[pulumi.Input[str]] = None,
                 author_name: Optional[pulumi.Input[str]] = None,
                 execute_filemode: Optional[pulumi.Input[bool]] = None,
                 start_branch: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RepositoryFile resource.
        :param pulumi.Input[str] branch: Name of the branch to which to commit to.
        :param pulumi.Input[str] commit_message: Commit message.
        :param pulumi.Input[str] content: File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently
               supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
        :param pulumi.Input[str] file_path: The full path of the file. It must be relative to the root of the project without a leading slash `/`.
        :param pulumi.Input[str] project: The name or ID of the project.
        :param pulumi.Input[str] author_email: Email of the commit author.
        :param pulumi.Input[str] author_name: Name of the commit author.
        :param pulumi.Input[bool] execute_filemode: Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
        :param pulumi.Input[str] start_branch: Name of the branch to start the new commit from.
        """
        pulumi.set(__self__, "branch", branch)
        pulumi.set(__self__, "commit_message", commit_message)
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "file_path", file_path)
        pulumi.set(__self__, "project", project)
        if author_email is not None:
            pulumi.set(__self__, "author_email", author_email)
        if author_name is not None:
            pulumi.set(__self__, "author_name", author_name)
        if execute_filemode is not None:
            pulumi.set(__self__, "execute_filemode", execute_filemode)
        if start_branch is not None:
            pulumi.set(__self__, "start_branch", start_branch)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Input[str]:
        """
        Name of the branch to which to commit to.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: pulumi.Input[str]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter(name="commitMessage")
    def commit_message(self) -> pulumi.Input[str]:
        """
        Commit message.
        """
        return pulumi.get(self, "commit_message")

    @commit_message.setter
    def commit_message(self, value: pulumi.Input[str]):
        pulumi.set(self, "commit_message", value)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently
        supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> pulumi.Input[str]:
        """
        The full path of the file. It must be relative to the root of the project without a leading slash `/`.
        """
        return pulumi.get(self, "file_path")

    @file_path.setter
    def file_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_path", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The name or ID of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="authorEmail")
    def author_email(self) -> Optional[pulumi.Input[str]]:
        """
        Email of the commit author.
        """
        return pulumi.get(self, "author_email")

    @author_email.setter
    def author_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author_email", value)

    @property
    @pulumi.getter(name="authorName")
    def author_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the commit author.
        """
        return pulumi.get(self, "author_name")

    @author_name.setter
    def author_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author_name", value)

    @property
    @pulumi.getter(name="executeFilemode")
    def execute_filemode(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
        """
        return pulumi.get(self, "execute_filemode")

    @execute_filemode.setter
    def execute_filemode(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "execute_filemode", value)

    @property
    @pulumi.getter(name="startBranch")
    def start_branch(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the branch to start the new commit from.
        """
        return pulumi.get(self, "start_branch")

    @start_branch.setter
    def start_branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_branch", value)


@pulumi.input_type
class _RepositoryFileState:
    def __init__(__self__, *,
                 author_email: Optional[pulumi.Input[str]] = None,
                 author_name: Optional[pulumi.Input[str]] = None,
                 blob_id: Optional[pulumi.Input[str]] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_id: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_sha256: Optional[pulumi.Input[str]] = None,
                 encoding: Optional[pulumi.Input[str]] = None,
                 execute_filemode: Optional[pulumi.Input[bool]] = None,
                 file_name: Optional[pulumi.Input[str]] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 last_commit_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ref: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 start_branch: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RepositoryFile resources.
        :param pulumi.Input[str] author_email: Email of the commit author.
        :param pulumi.Input[str] author_name: Name of the commit author.
        :param pulumi.Input[str] blob_id: The blob id.
        :param pulumi.Input[str] branch: Name of the branch to which to commit to.
        :param pulumi.Input[str] commit_id: The commit id.
        :param pulumi.Input[str] commit_message: Commit message.
        :param pulumi.Input[str] content: File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently
               supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
        :param pulumi.Input[str] content_sha256: File content sha256 digest.
        :param pulumi.Input[str] encoding: The file content encoding.
        :param pulumi.Input[bool] execute_filemode: Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
        :param pulumi.Input[str] file_name: The filename.
        :param pulumi.Input[str] file_path: The full path of the file. It must be relative to the root of the project without a leading slash `/`.
        :param pulumi.Input[str] last_commit_id: The last known commit id.
        :param pulumi.Input[str] project: The name or ID of the project.
        :param pulumi.Input[str] ref: The name of branch, tag or commit.
        :param pulumi.Input[int] size: The file size.
        :param pulumi.Input[str] start_branch: Name of the branch to start the new commit from.
        """
        if author_email is not None:
            pulumi.set(__self__, "author_email", author_email)
        if author_name is not None:
            pulumi.set(__self__, "author_name", author_name)
        if blob_id is not None:
            pulumi.set(__self__, "blob_id", blob_id)
        if branch is not None:
            pulumi.set(__self__, "branch", branch)
        if commit_id is not None:
            pulumi.set(__self__, "commit_id", commit_id)
        if commit_message is not None:
            pulumi.set(__self__, "commit_message", commit_message)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if content_sha256 is not None:
            pulumi.set(__self__, "content_sha256", content_sha256)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if execute_filemode is not None:
            pulumi.set(__self__, "execute_filemode", execute_filemode)
        if file_name is not None:
            pulumi.set(__self__, "file_name", file_name)
        if file_path is not None:
            pulumi.set(__self__, "file_path", file_path)
        if last_commit_id is not None:
            pulumi.set(__self__, "last_commit_id", last_commit_id)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if ref is not None:
            pulumi.set(__self__, "ref", ref)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if start_branch is not None:
            pulumi.set(__self__, "start_branch", start_branch)

    @property
    @pulumi.getter(name="authorEmail")
    def author_email(self) -> Optional[pulumi.Input[str]]:
        """
        Email of the commit author.
        """
        return pulumi.get(self, "author_email")

    @author_email.setter
    def author_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author_email", value)

    @property
    @pulumi.getter(name="authorName")
    def author_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the commit author.
        """
        return pulumi.get(self, "author_name")

    @author_name.setter
    def author_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author_name", value)

    @property
    @pulumi.getter(name="blobId")
    def blob_id(self) -> Optional[pulumi.Input[str]]:
        """
        The blob id.
        """
        return pulumi.get(self, "blob_id")

    @blob_id.setter
    def blob_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "blob_id", value)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the branch to which to commit to.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter(name="commitId")
    def commit_id(self) -> Optional[pulumi.Input[str]]:
        """
        The commit id.
        """
        return pulumi.get(self, "commit_id")

    @commit_id.setter
    def commit_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_id", value)

    @property
    @pulumi.getter(name="commitMessage")
    def commit_message(self) -> Optional[pulumi.Input[str]]:
        """
        Commit message.
        """
        return pulumi.get(self, "commit_message")

    @commit_message.setter
    def commit_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_message", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently
        supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="contentSha256")
    def content_sha256(self) -> Optional[pulumi.Input[str]]:
        """
        File content sha256 digest.
        """
        return pulumi.get(self, "content_sha256")

    @content_sha256.setter
    def content_sha256(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_sha256", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        The file content encoding.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="executeFilemode")
    def execute_filemode(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
        """
        return pulumi.get(self, "execute_filemode")

    @execute_filemode.setter
    def execute_filemode(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "execute_filemode", value)

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> Optional[pulumi.Input[str]]:
        """
        The filename.
        """
        return pulumi.get(self, "file_name")

    @file_name.setter
    def file_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_name", value)

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> Optional[pulumi.Input[str]]:
        """
        The full path of the file. It must be relative to the root of the project without a leading slash `/`.
        """
        return pulumi.get(self, "file_path")

    @file_path.setter
    def file_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_path", value)

    @property
    @pulumi.getter(name="lastCommitId")
    def last_commit_id(self) -> Optional[pulumi.Input[str]]:
        """
        The last known commit id.
        """
        return pulumi.get(self, "last_commit_id")

    @last_commit_id.setter
    def last_commit_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_commit_id", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The name or ID of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def ref(self) -> Optional[pulumi.Input[str]]:
        """
        The name of branch, tag or commit.
        """
        return pulumi.get(self, "ref")

    @ref.setter
    def ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ref", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        The file size.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="startBranch")
    def start_branch(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the branch to start the new commit from.
        """
        return pulumi.get(self, "start_branch")

    @start_branch.setter
    def start_branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_branch", value)


class RepositoryFile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 author_email: Optional[pulumi.Input[str]] = None,
                 author_name: Optional[pulumi.Input[str]] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 execute_filemode: Optional[pulumi.Input[bool]] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 start_branch: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        # A Repository File can be imported using an id made up of `<project-id>:<branch-name>:<file-path>`, e.g.

        ```sh
         $ pulumi import gitlab:index/repositoryFile:RepositoryFile this 1:main:foo/bar.txt
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] author_email: Email of the commit author.
        :param pulumi.Input[str] author_name: Name of the commit author.
        :param pulumi.Input[str] branch: Name of the branch to which to commit to.
        :param pulumi.Input[str] commit_message: Commit message.
        :param pulumi.Input[str] content: File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently
               supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
        :param pulumi.Input[bool] execute_filemode: Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
        :param pulumi.Input[str] file_path: The full path of the file. It must be relative to the root of the project without a leading slash `/`.
        :param pulumi.Input[str] project: The name or ID of the project.
        :param pulumi.Input[str] start_branch: Name of the branch to start the new commit from.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryFileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        # A Repository File can be imported using an id made up of `<project-id>:<branch-name>:<file-path>`, e.g.

        ```sh
         $ pulumi import gitlab:index/repositoryFile:RepositoryFile this 1:main:foo/bar.txt
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryFileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryFileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 author_email: Optional[pulumi.Input[str]] = None,
                 author_name: Optional[pulumi.Input[str]] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 execute_filemode: Optional[pulumi.Input[bool]] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 start_branch: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryFileArgs.__new__(RepositoryFileArgs)

            __props__.__dict__["author_email"] = author_email
            __props__.__dict__["author_name"] = author_name
            if branch is None and not opts.urn:
                raise TypeError("Missing required property 'branch'")
            __props__.__dict__["branch"] = branch
            if commit_message is None and not opts.urn:
                raise TypeError("Missing required property 'commit_message'")
            __props__.__dict__["commit_message"] = commit_message
            if content is None and not opts.urn:
                raise TypeError("Missing required property 'content'")
            __props__.__dict__["content"] = content
            __props__.__dict__["execute_filemode"] = execute_filemode
            if file_path is None and not opts.urn:
                raise TypeError("Missing required property 'file_path'")
            __props__.__dict__["file_path"] = file_path
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["start_branch"] = start_branch
            __props__.__dict__["blob_id"] = None
            __props__.__dict__["commit_id"] = None
            __props__.__dict__["content_sha256"] = None
            __props__.__dict__["encoding"] = None
            __props__.__dict__["file_name"] = None
            __props__.__dict__["last_commit_id"] = None
            __props__.__dict__["ref"] = None
            __props__.__dict__["size"] = None
        super(RepositoryFile, __self__).__init__(
            'gitlab:index/repositoryFile:RepositoryFile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            author_email: Optional[pulumi.Input[str]] = None,
            author_name: Optional[pulumi.Input[str]] = None,
            blob_id: Optional[pulumi.Input[str]] = None,
            branch: Optional[pulumi.Input[str]] = None,
            commit_id: Optional[pulumi.Input[str]] = None,
            commit_message: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            content_sha256: Optional[pulumi.Input[str]] = None,
            encoding: Optional[pulumi.Input[str]] = None,
            execute_filemode: Optional[pulumi.Input[bool]] = None,
            file_name: Optional[pulumi.Input[str]] = None,
            file_path: Optional[pulumi.Input[str]] = None,
            last_commit_id: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            ref: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            start_branch: Optional[pulumi.Input[str]] = None) -> 'RepositoryFile':
        """
        Get an existing RepositoryFile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] author_email: Email of the commit author.
        :param pulumi.Input[str] author_name: Name of the commit author.
        :param pulumi.Input[str] blob_id: The blob id.
        :param pulumi.Input[str] branch: Name of the branch to which to commit to.
        :param pulumi.Input[str] commit_id: The commit id.
        :param pulumi.Input[str] commit_message: Commit message.
        :param pulumi.Input[str] content: File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently
               supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
        :param pulumi.Input[str] content_sha256: File content sha256 digest.
        :param pulumi.Input[str] encoding: The file content encoding.
        :param pulumi.Input[bool] execute_filemode: Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
        :param pulumi.Input[str] file_name: The filename.
        :param pulumi.Input[str] file_path: The full path of the file. It must be relative to the root of the project without a leading slash `/`.
        :param pulumi.Input[str] last_commit_id: The last known commit id.
        :param pulumi.Input[str] project: The name or ID of the project.
        :param pulumi.Input[str] ref: The name of branch, tag or commit.
        :param pulumi.Input[int] size: The file size.
        :param pulumi.Input[str] start_branch: Name of the branch to start the new commit from.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryFileState.__new__(_RepositoryFileState)

        __props__.__dict__["author_email"] = author_email
        __props__.__dict__["author_name"] = author_name
        __props__.__dict__["blob_id"] = blob_id
        __props__.__dict__["branch"] = branch
        __props__.__dict__["commit_id"] = commit_id
        __props__.__dict__["commit_message"] = commit_message
        __props__.__dict__["content"] = content
        __props__.__dict__["content_sha256"] = content_sha256
        __props__.__dict__["encoding"] = encoding
        __props__.__dict__["execute_filemode"] = execute_filemode
        __props__.__dict__["file_name"] = file_name
        __props__.__dict__["file_path"] = file_path
        __props__.__dict__["last_commit_id"] = last_commit_id
        __props__.__dict__["project"] = project
        __props__.__dict__["ref"] = ref
        __props__.__dict__["size"] = size
        __props__.__dict__["start_branch"] = start_branch
        return RepositoryFile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorEmail")
    def author_email(self) -> pulumi.Output[Optional[str]]:
        """
        Email of the commit author.
        """
        return pulumi.get(self, "author_email")

    @property
    @pulumi.getter(name="authorName")
    def author_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the commit author.
        """
        return pulumi.get(self, "author_name")

    @property
    @pulumi.getter(name="blobId")
    def blob_id(self) -> pulumi.Output[str]:
        """
        The blob id.
        """
        return pulumi.get(self, "blob_id")

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[str]:
        """
        Name of the branch to which to commit to.
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter(name="commitId")
    def commit_id(self) -> pulumi.Output[str]:
        """
        The commit id.
        """
        return pulumi.get(self, "commit_id")

    @property
    @pulumi.getter(name="commitMessage")
    def commit_message(self) -> pulumi.Output[str]:
        """
        Commit message.
        """
        return pulumi.get(self, "commit_message")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently
        supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentSha256")
    def content_sha256(self) -> pulumi.Output[str]:
        """
        File content sha256 digest.
        """
        return pulumi.get(self, "content_sha256")

    @property
    @pulumi.getter
    def encoding(self) -> pulumi.Output[str]:
        """
        The file content encoding.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="executeFilemode")
    def execute_filemode(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
        """
        return pulumi.get(self, "execute_filemode")

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> pulumi.Output[str]:
        """
        The filename.
        """
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> pulumi.Output[str]:
        """
        The full path of the file. It must be relative to the root of the project without a leading slash `/`.
        """
        return pulumi.get(self, "file_path")

    @property
    @pulumi.getter(name="lastCommitId")
    def last_commit_id(self) -> pulumi.Output[str]:
        """
        The last known commit id.
        """
        return pulumi.get(self, "last_commit_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The name or ID of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def ref(self) -> pulumi.Output[str]:
        """
        The name of branch, tag or commit.
        """
        return pulumi.get(self, "ref")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The file size.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="startBranch")
    def start_branch(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the branch to start the new commit from.
        """
        return pulumi.get(self, "start_branch")

