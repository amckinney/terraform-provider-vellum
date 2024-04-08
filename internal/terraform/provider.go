package terraform

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/vellum-ai/terraform-provider-vellum/internal/terraform/base"
)

// Ensure Provider satisfies the terraform provider interface(s).
var _ provider.Provider = (*Provider)(nil)

// Provider defines the provider implementation.
type Provider struct {
	provider.Provider

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// NewProvider returns a new vellum terraform provider.
func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &Provider{
			Provider: base.NewProvider(version),
			version:  version,
		}
	}
}
