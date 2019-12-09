## HEAD (Unreleased)
* Namespace names in .NET SDK are adjusted to PascalCase
([#25](https://github.com/pulumi/pulumi-gitlab/pull/25)).

---

## 1.3.0 (2019-12-09)
* Upgrade to v2.5.0 of the Gitlab Terraform Provider

## 1.2.0 (2019-12-04)
* Upgrade to support go 1.13.x
* Upgrade to pulumi-terraform-bridge v1.4.3
* Upgrade to v2.4.0 of the Gitlab Terraform Provider

## 1.1.0 (2019-11-15)
* Add support for DotNet SDK Generation

## 1.0.0 (2019-10-17)
* Upgrade to v2.3.0 of the Gitlab Terraform Provider

## 0.18.8 (2019-09-05)
* Upgrade to Pulumi v1.0.0

## 0.18.7 (2019-08-29)
* Upgrade pulumi-terraform to 3f206601e7

## 0.18.6 (2019-08-20)
* Depend on latest pulumi package

## 0.18.5 (2019-08-09)
* Upgrade to pulumi-terraform@9db2fc93cd

## 0.18.4 (2019-08-08)
* Update to pulumi-terraform@013b95b1c8

## 0.18.3 (2019-07-09)
* Fix detailed diffs with nested computed values.

## 0.18.2 (2019-07-08)
* Communicate detailed information about the difference between a resource's desired and actual state during a Pulumi update

## 0.18.1 (2019-06-21)
* Update to pulumi-terraform@3635bed3a5 which stops maps containing `.` being treated as nested maps.

## 0.18.0 (2019-06-17)
* Initial release of `pulumi-gitlab`, based on `v2.2.0` of the GitLab Terraform provider.
