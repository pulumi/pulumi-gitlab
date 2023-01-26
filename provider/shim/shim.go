package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/gitlab-org/terraform-provider-gitlab/internal/provider"
)

func NewProvider() *schema.Provider {
	return provider.New("")()
}
