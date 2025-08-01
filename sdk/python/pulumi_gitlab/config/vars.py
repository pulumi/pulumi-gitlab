# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

import types

__config__ = pulumi.Config('gitlab')


class _ExportableConfig(types.ModuleType):
    @_builtins.property
    def base_url(self) -> Optional[str]:
        """
        This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
        Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
        the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
        """
        return __config__.get('baseUrl')

    @_builtins.property
    def cacert_file(self) -> Optional[str]:
        """
        This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
        CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
        """
        return __config__.get('cacertFile')

    @_builtins.property
    def client_cert(self) -> Optional[str]:
        """
        File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
        """
        return __config__.get('clientCert')

    @_builtins.property
    def client_key(self) -> Optional[str]:
        """
        File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
        `client_cert` is set.
        """
        return __config__.get('clientKey')

    @_builtins.property
    def early_auth_check(self) -> Optional[bool]:
        return __config__.get_bool('earlyAuthCheck')

    @_builtins.property
    def headers(self) -> Optional[str]:
        """
        A map of headers to append to all API request to the GitLab instance.
        """
        return __config__.get('headers')

    @_builtins.property
    def insecure(self) -> Optional[bool]:
        """
        When set to true this disables SSL verification of the connection to the GitLab instance.
        """
        return __config__.get_bool('insecure')

    @_builtins.property
    def retries(self) -> Optional[int]:
        """
        The number of retries to execute when receiving a 429 Rate Limit error. Each retry will exponentially back off.
        """
        return __config__.get_int('retries')

    @_builtins.property
    def token(self) -> Optional[str]:
        """
        The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
        used in this provider for authentication (using Bearer authorization token). See
        https://docs.gitlab.com/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment variable.
        """
        return __config__.get('token')

