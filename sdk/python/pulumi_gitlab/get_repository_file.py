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
    'GetRepositoryFileResult',
    'AwaitableGetRepositoryFileResult',
    'get_repository_file',
    'get_repository_file_output',
]

@pulumi.output_type
class GetRepositoryFileResult:
    """
    A collection of values returned by getRepositoryFile.
    """
    def __init__(__self__, blob_id=None, commit_id=None, content=None, content_sha256=None, encoding=None, execute_filemode=None, file_name=None, file_path=None, id=None, last_commit_id=None, project=None, ref=None, size=None):
        if blob_id and not isinstance(blob_id, str):
            raise TypeError("Expected argument 'blob_id' to be a str")
        pulumi.set(__self__, "blob_id", blob_id)
        if commit_id and not isinstance(commit_id, str):
            raise TypeError("Expected argument 'commit_id' to be a str")
        pulumi.set(__self__, "commit_id", commit_id)
        if content and not isinstance(content, str):
            raise TypeError("Expected argument 'content' to be a str")
        pulumi.set(__self__, "content", content)
        if content_sha256 and not isinstance(content_sha256, str):
            raise TypeError("Expected argument 'content_sha256' to be a str")
        pulumi.set(__self__, "content_sha256", content_sha256)
        if encoding and not isinstance(encoding, str):
            raise TypeError("Expected argument 'encoding' to be a str")
        pulumi.set(__self__, "encoding", encoding)
        if execute_filemode and not isinstance(execute_filemode, bool):
            raise TypeError("Expected argument 'execute_filemode' to be a bool")
        pulumi.set(__self__, "execute_filemode", execute_filemode)
        if file_name and not isinstance(file_name, str):
            raise TypeError("Expected argument 'file_name' to be a str")
        pulumi.set(__self__, "file_name", file_name)
        if file_path and not isinstance(file_path, str):
            raise TypeError("Expected argument 'file_path' to be a str")
        pulumi.set(__self__, "file_path", file_path)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_commit_id and not isinstance(last_commit_id, str):
            raise TypeError("Expected argument 'last_commit_id' to be a str")
        pulumi.set(__self__, "last_commit_id", last_commit_id)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if ref and not isinstance(ref, str):
            raise TypeError("Expected argument 'ref' to be a str")
        pulumi.set(__self__, "ref", ref)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter(name="blobId")
    def blob_id(self) -> str:
        """
        The blob id.
        """
        return pulumi.get(self, "blob_id")

    @property
    @pulumi.getter(name="commitId")
    def commit_id(self) -> str:
        """
        The commit id.
        """
        return pulumi.get(self, "commit_id")

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        File content.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentSha256")
    def content_sha256(self) -> str:
        """
        File content sha256 digest.
        """
        return pulumi.get(self, "content_sha256")

    @property
    @pulumi.getter
    def encoding(self) -> str:
        """
        The file content encoding.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="executeFilemode")
    def execute_filemode(self) -> bool:
        """
        Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
        """
        return pulumi.get(self, "execute_filemode")

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> str:
        """
        The filename.
        """
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> str:
        """
        The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
        """
        return pulumi.get(self, "file_path")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastCommitId")
    def last_commit_id(self) -> str:
        """
        The last known commit id.
        """
        return pulumi.get(self, "last_commit_id")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The name or ID of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def ref(self) -> str:
        """
        The name of branch, tag or commit.
        """
        return pulumi.get(self, "ref")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The file size.
        """
        return pulumi.get(self, "size")


class AwaitableGetRepositoryFileResult(GetRepositoryFileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryFileResult(
            blob_id=self.blob_id,
            commit_id=self.commit_id,
            content=self.content,
            content_sha256=self.content_sha256,
            encoding=self.encoding,
            execute_filemode=self.execute_filemode,
            file_name=self.file_name,
            file_path=self.file_path,
            id=self.id,
            last_commit_id=self.last_commit_id,
            project=self.project,
            ref=self.ref,
            size=self.size)


def get_repository_file(file_path: Optional[str] = None,
                        project: Optional[str] = None,
                        ref: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryFileResult:
    """
    The `RepositoryFile` data source allows details of a file in a repository to be retrieved.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/repository_files.html)

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_repository_file(project="example",
        ref="main",
        file_path="README.md")
    ```
    <!--End PulumiCodeChooser -->


    :param str file_path: The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
    :param str project: The name or ID of the project.
    :param str ref: The name of branch, tag or commit.
    """
    __args__ = dict()
    __args__['filePath'] = file_path
    __args__['project'] = project
    __args__['ref'] = ref
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gitlab:index/getRepositoryFile:getRepositoryFile', __args__, opts=opts, typ=GetRepositoryFileResult).value

    return AwaitableGetRepositoryFileResult(
        blob_id=pulumi.get(__ret__, 'blob_id'),
        commit_id=pulumi.get(__ret__, 'commit_id'),
        content=pulumi.get(__ret__, 'content'),
        content_sha256=pulumi.get(__ret__, 'content_sha256'),
        encoding=pulumi.get(__ret__, 'encoding'),
        execute_filemode=pulumi.get(__ret__, 'execute_filemode'),
        file_name=pulumi.get(__ret__, 'file_name'),
        file_path=pulumi.get(__ret__, 'file_path'),
        id=pulumi.get(__ret__, 'id'),
        last_commit_id=pulumi.get(__ret__, 'last_commit_id'),
        project=pulumi.get(__ret__, 'project'),
        ref=pulumi.get(__ret__, 'ref'),
        size=pulumi.get(__ret__, 'size'))


@_utilities.lift_output_func(get_repository_file)
def get_repository_file_output(file_path: Optional[pulumi.Input[str]] = None,
                               project: Optional[pulumi.Input[str]] = None,
                               ref: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryFileResult]:
    """
    The `RepositoryFile` data source allows details of a file in a repository to be retrieved.

    **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/repository_files.html)

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_repository_file(project="example",
        ref="main",
        file_path="README.md")
    ```
    <!--End PulumiCodeChooser -->


    :param str file_path: The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
    :param str project: The name or ID of the project.
    :param str ref: The name of branch, tag or commit.
    """
    ...
