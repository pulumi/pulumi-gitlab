package shim

import (
	"github.com/gitlabhq/terraform-provider-gitlab/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider() *schema.Provider {
	return provider.New("")()
}
