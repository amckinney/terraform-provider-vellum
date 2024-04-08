package base

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	vellumdatasource "github.com/vellum-ai/terraform-provider-vellum/internal/terraform/datasource"
	vellumresource "github.com/vellum-ai/terraform-provider-vellum/internal/terraform/resource"
	vellumclient "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/client"
	"github.com/vellum-ai/terraform-provider-vellum/internal/vellum/option"
)

// Ensure Provider satisfies the terraform provider interface(s).
var _ provider.Provider = (*Provider)(nil)

// Provider defines the provider implementation.
type Provider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// NewProvider returns a new vellum terraform provider.
func NewProvider(version string) provider.Provider {
	return &Provider{
		version: version,
	}
}

// ProviderModel describes the provider data model.
type ProviderModel struct {
	ApiKey types.String `tfsdk:"api_key"`
}

func (p *Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "vellum"
	resp.Version = p.version
}

func (p *Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				MarkdownDescription: "API Key to authenticate with the Vellum API",
				Optional:            true,
			},
		},
	}
}

func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var model ProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiKey := os.Getenv("VELLUM_API_KEY")
	if !model.ApiKey.IsNull() {
		apiKey = model.ApiKey.ValueString()
	}

	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"An API key is required to use the vellum provider",
			"You must set a VELLUM_API_KEY or specify an api_key in the provider constructor",
		)
		return
	}

	client := vellumclient.NewClient(
		option.WithApiKey(
			apiKey,
		),
	)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *Provider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		vellumresource.NewDocumentIndex,
	}
}

func (p *Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		vellumdatasource.NewDocumentIndex,
	}
}
