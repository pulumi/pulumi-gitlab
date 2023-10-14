# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

baseUrl: Optional[str]
"""
This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
"""

cacertFile: Optional[str]
"""
This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
"""

clientCert: Optional[str]
"""
File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
"""

clientKey: Optional[str]
"""
File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
`client_cert` is set.
"""

earlyAuthCheck: Optional[bool]
"""
(Experimental) By default the provider does a dummy request to get the current user in order to verify that the provider
configuration is correct and the GitLab API is reachable. Set this to `false` to skip this check. This may be useful if
the GitLab instance does not yet exist and is created within the same terraform module. It may be sourced from the
`GITLAB_EARLY_AUTH_CHECK`. This is an experimental feature and may change in the future. Please make sure to always keep
backups of your state.
"""

insecure: Optional[bool]
"""
When set to true this disables SSL verification of the connection to the GitLab instance.
"""

token: Optional[str]
"""
The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
used in this provider for authentication (using Bearer authorization token). See
https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
variable.
"""

