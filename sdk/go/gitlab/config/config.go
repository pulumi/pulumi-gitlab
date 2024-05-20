// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-gitlab/sdk/v7/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
// Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
// the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
func GetBaseUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "gitlab:baseUrl")
}

// This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
// CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
func GetCacertFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "gitlab:cacertFile")
}

// File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
func GetClientCert(ctx *pulumi.Context) string {
	return config.Get(ctx, "gitlab:clientCert")
}

// File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
// `clientCert` is set.
func GetClientKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "gitlab:clientKey")
}
func GetEarlyAuthCheck(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "gitlab:earlyAuthCheck")
}

// When set to true this disables SSL verification of the connection to the GitLab instance.
func GetInsecure(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "gitlab:insecure")
}

// The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
// used in this provider for authentication (using Bearer authorization token). See
// https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
// variable.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "gitlab:token")
}
