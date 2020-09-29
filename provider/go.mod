module github.com/pulumi/pulumi-gitlab/provider/v3

go 1.14

require (
	github.com/gitlabhq/terraform-provider-gitlab v1.3.1-0.20200924015720-21ad328f81b1
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.8.0
	github.com/pulumi/pulumi/sdk/v2 v2.10.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
