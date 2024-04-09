package terraform

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/vellum-ai/terraform-provider-vellum/internal/terraform/base"
)

// Ensure Provider satisfies the terraform provider interface(s).
var _ provider.Provider = (*Provider)(nil)

// Provider defines the provider implementation.
type Provider struct {
	base *base.Provider

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// NewProvider returns a new vellum terraform provider.
func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &Provider{
			base:    base.NewProvider(version).(*base.Provider),
			version: version,
		}
	}
}

// ---

func (p *Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	p.base.Metadata(ctx, req, resp)
}

func (p *Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	p.base.Schema(ctx, req, resp)
}

func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	p.base.Configure(ctx, req, resp)
}

func (p *Provider) Resources(ctx context.Context) []func() resource.Resource {
	return p.base.Resources(ctx)
}

func (p *Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return p.base.DataSources(ctx)
}
