module github.com/pulumi/pulumi-gitlab/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.13.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.8.0
	github.com/pulumi/pulumi/sdk/v2 v2.10.0
	github.com/terraform-providers/terraform-provider-gitlab v1.3.1-0.20200724102830-d48f13a9bfb5
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
