module github.com/pulumi/pulumi-gitlab/provider/v4

go 1.16

require (
	github.com/gitlabhq/terraform-provider-gitlab v1.3.1-0.20210720003513-89de4bd5e56d
	github.com/hashicorp/terraform-plugin-sdk v1.16.1
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.11.0
	github.com/pulumi/pulumi/sdk/v3 v3.17.0
)

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-test => github.com/hashicorp/terraform-plugin-test v1.3.0
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
