package datasource

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/vellum-ai/terraform-provider-vellum/internal/vellum"
	vellumclient "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/client"
)

type DocumentIndexModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Label       types.String `tfsdk:"label"`
	Environment types.String `tfsdk:"environment"`
	Status      types.String `tfsdk:"status"`
	Created     types.String `tfsdk:"created"`
}

type DocumentIndex struct {
	client *vellumclient.Client
}

var _ datasource.DataSource = (*DocumentIndex)(nil)

// TODO: Add in override option parameter(s).
func NewDocumentIndex() datasource.DataSource {
	return &DocumentIndex{}
}

func (d *DocumentIndex) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_document_index"
}

func (d *DocumentIndex) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Document Index data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The Document Index's ID",
				MarkdownDescription: "The Document Index's ID",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "A name that uniquely identifies this index within its workspace",
				MarkdownDescription: "A name that uniquely identifies this index within its workspace",
			},
			"label": schema.StringAttribute{
				Computed:            true,
				Description:         "A human-readable label for the document index",
				MarkdownDescription: "A human-readable label for the document index",
			},
			"environment": schema.StringAttribute{
				Computed:            true,
				Description:         "The environment this document index is used in\n\n* `DEVELOPMENT` - Development\n* `STAGING` - Staging\n* `PRODUCTION` - Production",
				MarkdownDescription: "The environment this document index is used in\n\n* `DEVELOPMENT` - Development\n* `STAGING` - Staging\n* `PRODUCTION` - Production",
			},
			"status": schema.StringAttribute{
				Computed:            true,
				Description:         "The current status of the document index\n\n* `ACTIVE` - Active\n* `ARCHIVED` - Archived",
				MarkdownDescription: "The current status of the document index\n\n* `ACTIVE` - Active\n* `ARCHIVED` - Archived",
			},
			"created": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *DocumentIndex) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*vellumclient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *vellum.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *DocumentIndex) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model *DocumentIndexModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := d.client.DocumentIndexes.Retrieve(
		ctx,
		d.modelToRetrieveRequest(model),
	)
	if err != nil {
		resp.Diagnostics.AddError("error getting document index information", err.Error())
		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(
			ctx,
			d.retrieveResponseToModel(response),
		)...,
	)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *DocumentIndex) modelToRetrieveRequest(model *DocumentIndexModel) string {
	// Unlike the original solution, this approach _requires_ the name to be set.
	//
	// From the perspective of the Terraform provider, we should only expose a single
	// parameter to be set (i.e. 'name'), and set the 'id' as 'Computed'.
	//
	// By setting the parameter as 'Required', we get better error handling and a
	// simpler (yet more restrictive) user experience. Plus, the user is actually
	// free to set either ID or name in the terraform data source.
	return model.Name.ValueString()
}

func (d *DocumentIndex) retrieveResponseToModel(response *vellum.DocumentIndexRead) *DocumentIndexModel {
	return &DocumentIndexModel{
		Id:          types.StringValue(response.Id),
		Name:        types.StringValue(response.Name),
		Created:     types.StringValue(response.Created.Format(time.RFC3339)),
		Environment: types.StringValue(string(*response.Environment)),
		Label:       types.StringValue(response.Label),
		Status:      types.StringValue(string(*response.Status)),
	}
}
