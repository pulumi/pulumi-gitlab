[![Build Status](https://travis-ci.com/pulumi/pulumi-gitlab.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-gitlab)

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

    $ go get github.com/pulumi/pulumi-gitlab/sdk/go/...

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Gitlab

## Concepts

The `@pulumi/gitlab` package provides a strongly-typed means to create cloud applications that create and interact
closely with GitLab resources.

## Configuration

The following configuration points are available:

* token (Optional) - This is the GitLab personal access token. It must be provided but can also be sourced via `GITLAB_TOKEN`. 

* base_url (Optional) - This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab Enterprise e.g. https://my.gitlab.server/api/v4/. It is optional to provide this value and it can also be sourced from the GITLAB_BASE_URL environment variable. The value must end with a slash.

* cacert_file (Optional) - This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.

* insecure (Optional) - When set to true this disables SSL verification of the connection to the GitLab instance. Defaults to `false`.


## Reference

For further information, please visit [the GitLab provider docs](https://www.pulumi.com/docs/intro/cloud-providers/gitlab) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/gitlab).
