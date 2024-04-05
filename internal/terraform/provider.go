package terraform

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/vellum-ai/terraform-provider-vellum/internal/terraform/document_index"
	vellumclient "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/client"
	"github.com/vellum-ai/terraform-provider-vellum/internal/vellum/option"
)

// Ensure VellumProvider satisfies various provider interfaces.
var _ provider.Provider = &VellumProvider{}
var _ provider.ProviderWithFunctions = &VellumProvider{}

// VellumProvider defines the provider implementation.
type VellumProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// VellumProviderModel describes the provider data model.
type VellumProviderModel struct {
	APIKey types.String `tfsdk:"api_key"`
}

func (p *VellumProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "vellum"
	resp.Version = p.version
}

func (p *VellumProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				MarkdownDescription: "API Key to authenticate with the Vellum API",
				Optional:            true,
			},
		},
	}
}

func (p *VellumProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data VellumProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := vellumclient.NewClient(
		option.WithApiKey(
			os.Getenv("VELLUM_API_KEY"),
		),
	)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *VellumProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		document_index.Resource,
	}
}

func (p *VellumProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		document_index.DataSource,
	}
}

func (p *VellumProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &VellumProvider{
			version: version,
		}
	}
}
