[![Actions Status](https://github.com/pulumi/pulumi-gitlab/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-gitlab/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fgitlab.svg)](https://www.npmjs.com/package/@pulumi/gitlab)
[![Python version](https://badge.fury.io/py/pulumi-gitlab.svg)](https://pypi.org/project/pulumi-gitlab)
[![NuGet version](https://badge.fury.io/nu/pulumi.gitlab.svg)](https://badge.fury.io/nu/pulumi.gitlab)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-gitlab/sdk/v6/go)](https://pkg.go.dev/github.com/pulumi/pulumi-gitlab/sdk/v6/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-gitlab/blob/master/LICENSE)

# GitLab provider

The GitLab resource provider for Pulumi lets you use GitLab resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/gitlab

or `yarn`:

    $ yarn add @pulumi/gitlab

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_gitlab

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-gitlab/sdk/v7

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Gitlab

## Concepts

The `@pulumi/gitlab` package provides a strongly-typed means to create cloud applications that create and interact
closely with GitLab resources.

## Configuration

The following configuration points are available:

* token (Optional) - This is the GitLab personal access token. It must be provided but can also be sourced via `GITLAB_TOKEN`.

* baseUrl (Optional) - This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab Enterprise e.g. https://my.gitlab.server/api/v4/. It is optional to provide this value and it can also be sourced from the GITLAB_BASE_URL environment variable. The value must end with a slash.

* cacertFile (Optional) - This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.

* insecure (Optional) - When set to true this disables SSL verification of the connection to the GitLab instance. Defaults to `false`.


## Reference

For further information, please visit [the GitLab provider docs](https://www.pulumi.com/docs/intro/cloud-providers/gitlab) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/gitlab).
