package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	googleWorkspace "github.com/hashicorp/terraform-provider-googleworkspace/internal/provider"
)

func NewProvider(version string) *schema.Provider {
	return googleWorkspace.New(version)()
}
