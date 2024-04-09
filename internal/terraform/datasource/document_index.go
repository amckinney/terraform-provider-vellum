package datasource

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

func (d *DocumentIndex) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.base.Configure(ctx, req, resp)
}

func (d *DocumentIndex) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	// Apply the schema defined by the generated base.
	d.base.Schema(ctx, req, resp)

	// Override the 'id' attribute and make it optional.
	resp.Schema.Attributes["id"] = schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Description:         "The Document Index's ID",
		MarkdownDescription: "The Document Index's ID",
	}

	// Override the 'name' attribute and make it optional.
	resp.Schema.Attributes["name"] = schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Description:         "A name that uniquely identifies this index within its workspace",
		MarkdownDescription: "A name that uniquely identifies this index within its workspace",
	}
}

func (d *DocumentIndex) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model *basedatasource.DocumentIndexModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Add validation to prevent both fields from being set.
	if !model.Id.IsNull() && !model.Name.IsNull() {
		resp.Diagnostics.AddError(
			"Cannot read document index data source with multiple unique identifiers",
			"Either an 'id' or 'name' is required to read a document index data source, but both were set",
		)
		return
	}

	// Add validation to guarantee at least one ID was set.
	if model.Id.IsNull() && model.Name.IsNull() {
		resp.Diagnostics.AddError(
			"Cannot read document index data source without a unique identifier",
			"Either an 'id' or 'name' is required to read a document index data source",
		)
		return
	}

	// Resolve the ID from either the 'name' or the 'id'.
	retrieveID := model.Name.ValueString()
	if retrieveID == "" {
		retrieveID = model.Id.ValueString()
	}

	response, err := d.base.Vellum.DocumentIndexes.Retrieve(
		ctx,
		retrieveID,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Retrieve document index error",
			fmt.Sprintf("Unable to retrieve document index, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(
			ctx,
			&basedatasource.DocumentIndexModel{
				Id:          types.StringValue(response.Id),
				Name:        types.StringValue(response.Name),
				Created:     types.StringValue(response.Created.Format(time.RFC3339)),
				Environment: types.StringValue(string(*response.Environment)),
				Label:       types.StringValue(response.Label),
				Status:      types.StringValue(string(*response.Status)),
			},
		)...,
	)
}
