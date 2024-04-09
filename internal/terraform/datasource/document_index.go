package datasource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	basedatasource "github.com/vellum-ai/terraform-provider-vellum/internal/terraform/base/datasource"
)

type DocumentIndex struct {
	base *basedatasource.DocumentIndex
}

var _ datasource.DataSource = (*DocumentIndex)(nil)

func NewDocumentIndex() datasource.DataSource {
	return &DocumentIndex{
		base: basedatasource.NewDocumentIndex().(*basedatasource.DocumentIndex),
	}
}

func (d *DocumentIndex) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	d.base.Metadata(ctx, req, resp)
}

func (d *DocumentIndex) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	d.base.Schema(ctx, req, resp)
}

func (d *DocumentIndex) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.base.Configure(ctx, req, resp)
}

func (d *DocumentIndex) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.base.Read(ctx, req, resp)
}
