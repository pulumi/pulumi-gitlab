# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ReleaseLinkArgs', 'ReleaseLink']

@pulumi.input_type
class ReleaseLinkArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 tag_name: pulumi.Input[str],
                 url: pulumi.Input[str],
                 filepath: Optional[pulumi.Input[str]] = None,
                 link_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReleaseLink resource.
        :param pulumi.Input[str] project: The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        :param pulumi.Input[str] tag_name: The tag associated with the Release.
        :param pulumi.Input[str] url: The URL of the link. Link URLs must be unique within the release.
        :param pulumi.Input[str] filepath: Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        :param pulumi.Input[str] link_type: The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
        :param pulumi.Input[str] name: The name of the link. Link names must be unique within the release.
        """
        ReleaseLinkArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            project=project,
            tag_name=tag_name,
            url=url,
            filepath=filepath,
            link_type=link_type,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             project: Optional[pulumi.Input[str]] = None,
             tag_name: Optional[pulumi.Input[str]] = None,
             url: Optional[pulumi.Input[str]] = None,
             filepath: Optional[pulumi.Input[str]] = None,
             link_type: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if project is None:
            raise TypeError("Missing 'project' argument")
        if tag_name is None and 'tagName' in kwargs:
            tag_name = kwargs['tagName']
        if tag_name is None:
            raise TypeError("Missing 'tag_name' argument")
        if url is None:
            raise TypeError("Missing 'url' argument")
        if link_type is None and 'linkType' in kwargs:
            link_type = kwargs['linkType']

        _setter("project", project)
        _setter("tag_name", tag_name)
        _setter("url", url)
        if filepath is not None:
            _setter("filepath", filepath)
        if link_type is not None:
            _setter("link_type", link_type)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="tagName")
    def tag_name(self) -> pulumi.Input[str]:
        """
        The tag associated with the Release.
        """
        return pulumi.get(self, "tag_name")

    @tag_name.setter
    def tag_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "tag_name", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The URL of the link. Link URLs must be unique within the release.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def filepath(self) -> Optional[pulumi.Input[str]]:
        """
        Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        """
        return pulumi.get(self, "filepath")

    @filepath.setter
    def filepath(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filepath", value)

    @property
    @pulumi.getter(name="linkType")
    def link_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
        """
        return pulumi.get(self, "link_type")

    @link_type.setter
    def link_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the link. Link names must be unique within the release.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ReleaseLinkState:
    def __init__(__self__, *,
                 direct_asset_url: Optional[pulumi.Input[str]] = None,
                 external: Optional[pulumi.Input[bool]] = None,
                 filepath: Optional[pulumi.Input[str]] = None,
                 link_id: Optional[pulumi.Input[int]] = None,
                 link_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tag_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReleaseLink resources.
        :param pulumi.Input[str] direct_asset_url: Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        :param pulumi.Input[bool] external: External or internal link.
        :param pulumi.Input[str] filepath: Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        :param pulumi.Input[int] link_id: The ID of the link.
        :param pulumi.Input[str] link_type: The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
        :param pulumi.Input[str] name: The name of the link. Link names must be unique within the release.
        :param pulumi.Input[str] project: The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        :param pulumi.Input[str] tag_name: The tag associated with the Release.
        :param pulumi.Input[str] url: The URL of the link. Link URLs must be unique within the release.
        """
        _ReleaseLinkState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            direct_asset_url=direct_asset_url,
            external=external,
            filepath=filepath,
            link_id=link_id,
            link_type=link_type,
            name=name,
            project=project,
            tag_name=tag_name,
            url=url,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             direct_asset_url: Optional[pulumi.Input[str]] = None,
             external: Optional[pulumi.Input[bool]] = None,
             filepath: Optional[pulumi.Input[str]] = None,
             link_id: Optional[pulumi.Input[int]] = None,
             link_type: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             tag_name: Optional[pulumi.Input[str]] = None,
             url: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if direct_asset_url is None and 'directAssetUrl' in kwargs:
            direct_asset_url = kwargs['directAssetUrl']
        if link_id is None and 'linkId' in kwargs:
            link_id = kwargs['linkId']
        if link_type is None and 'linkType' in kwargs:
            link_type = kwargs['linkType']
        if tag_name is None and 'tagName' in kwargs:
            tag_name = kwargs['tagName']

        if direct_asset_url is not None:
            _setter("direct_asset_url", direct_asset_url)
        if external is not None:
            _setter("external", external)
        if filepath is not None:
            _setter("filepath", filepath)
        if link_id is not None:
            _setter("link_id", link_id)
        if link_type is not None:
            _setter("link_type", link_type)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if tag_name is not None:
            _setter("tag_name", tag_name)
        if url is not None:
            _setter("url", url)

    @property
    @pulumi.getter(name="directAssetUrl")
    def direct_asset_url(self) -> Optional[pulumi.Input[str]]:
        """
        Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        """
        return pulumi.get(self, "direct_asset_url")

    @direct_asset_url.setter
    def direct_asset_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direct_asset_url", value)

    @property
    @pulumi.getter
    def external(self) -> Optional[pulumi.Input[bool]]:
        """
        External or internal link.
        """
        return pulumi.get(self, "external")

    @external.setter
    def external(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external", value)

    @property
    @pulumi.getter
    def filepath(self) -> Optional[pulumi.Input[str]]:
        """
        Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        """
        return pulumi.get(self, "filepath")

    @filepath.setter
    def filepath(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filepath", value)

    @property
    @pulumi.getter(name="linkId")
    def link_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the link.
        """
        return pulumi.get(self, "link_id")

    @link_id.setter
    def link_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "link_id", value)

    @property
    @pulumi.getter(name="linkType")
    def link_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
        """
        return pulumi.get(self, "link_type")

    @link_type.setter
    def link_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the link. Link names must be unique within the release.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="tagName")
    def tag_name(self) -> Optional[pulumi.Input[str]]:
        """
        The tag associated with the Release.
        """
        return pulumi.get(self, "tag_name")

    @tag_name.setter
    def tag_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the link. Link URLs must be unique within the release.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ReleaseLink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filepath: Optional[pulumi.Input[str]] = None,
                 link_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tag_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `ReleaseLink` resource allows to manage the lifecycle of a release link.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # Create a project
        example_project = gitlab.Project("exampleProject", description="An example project")
        # Can create release link only to a tag associated with a release
        example_release_link = gitlab.ReleaseLink("exampleReleaseLink",
            project=example_project.id,
            tag_name="tag_name_associated_with_release",
            url="https://test/")
        ```

        ## Import

        Gitlab release link can be imported with a key composed of `<project>:<tag_name>:<link_id>`, e.g.

        ```sh
         $ pulumi import gitlab:index/releaseLink:ReleaseLink example "12345:test:2"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] filepath: Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        :param pulumi.Input[str] link_type: The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
        :param pulumi.Input[str] name: The name of the link. Link names must be unique within the release.
        :param pulumi.Input[str] project: The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        :param pulumi.Input[str] tag_name: The tag associated with the Release.
        :param pulumi.Input[str] url: The URL of the link. Link URLs must be unique within the release.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReleaseLinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `ReleaseLink` resource allows to manage the lifecycle of a release link.

        **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        # Create a project
        example_project = gitlab.Project("exampleProject", description="An example project")
        # Can create release link only to a tag associated with a release
        example_release_link = gitlab.ReleaseLink("exampleReleaseLink",
            project=example_project.id,
            tag_name="tag_name_associated_with_release",
            url="https://test/")
        ```

        ## Import

        Gitlab release link can be imported with a key composed of `<project>:<tag_name>:<link_id>`, e.g.

        ```sh
         $ pulumi import gitlab:index/releaseLink:ReleaseLink example "12345:test:2"
        ```

        :param str resource_name: The name of the resource.
        :param ReleaseLinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReleaseLinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ReleaseLinkArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filepath: Optional[pulumi.Input[str]] = None,
                 link_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tag_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReleaseLinkArgs.__new__(ReleaseLinkArgs)

            __props__.__dict__["filepath"] = filepath
            __props__.__dict__["link_type"] = link_type
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if tag_name is None and not opts.urn:
                raise TypeError("Missing required property 'tag_name'")
            __props__.__dict__["tag_name"] = tag_name
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["direct_asset_url"] = None
            __props__.__dict__["external"] = None
            __props__.__dict__["link_id"] = None
        super(ReleaseLink, __self__).__init__(
            'gitlab:index/releaseLink:ReleaseLink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            direct_asset_url: Optional[pulumi.Input[str]] = None,
            external: Optional[pulumi.Input[bool]] = None,
            filepath: Optional[pulumi.Input[str]] = None,
            link_id: Optional[pulumi.Input[int]] = None,
            link_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            tag_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ReleaseLink':
        """
        Get an existing ReleaseLink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] direct_asset_url: Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        :param pulumi.Input[bool] external: External or internal link.
        :param pulumi.Input[str] filepath: Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        :param pulumi.Input[int] link_id: The ID of the link.
        :param pulumi.Input[str] link_type: The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
        :param pulumi.Input[str] name: The name of the link. Link names must be unique within the release.
        :param pulumi.Input[str] project: The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        :param pulumi.Input[str] tag_name: The tag associated with the Release.
        :param pulumi.Input[str] url: The URL of the link. Link URLs must be unique within the release.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReleaseLinkState.__new__(_ReleaseLinkState)

        __props__.__dict__["direct_asset_url"] = direct_asset_url
        __props__.__dict__["external"] = external
        __props__.__dict__["filepath"] = filepath
        __props__.__dict__["link_id"] = link_id
        __props__.__dict__["link_type"] = link_type
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["tag_name"] = tag_name
        __props__.__dict__["url"] = url
        return ReleaseLink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="directAssetUrl")
    def direct_asset_url(self) -> pulumi.Output[str]:
        """
        Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        """
        return pulumi.get(self, "direct_asset_url")

    @property
    @pulumi.getter
    def external(self) -> pulumi.Output[bool]:
        """
        External or internal link.
        """
        return pulumi.get(self, "external")

    @property
    @pulumi.getter
    def filepath(self) -> pulumi.Output[Optional[str]]:
        """
        Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
        """
        return pulumi.get(self, "filepath")

    @property
    @pulumi.getter(name="linkId")
    def link_id(self) -> pulumi.Output[int]:
        """
        The ID of the link.
        """
        return pulumi.get(self, "link_id")

    @property
    @pulumi.getter(name="linkType")
    def link_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
        """
        return pulumi.get(self, "link_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the link. Link names must be unique within the release.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="tagName")
    def tag_name(self) -> pulumi.Output[str]:
        """
        The tag associated with the Release.
        """
        return pulumi.get(self, "tag_name")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the link. Link URLs must be unique within the release.
        """
        return pulumi.get(self, "url")

