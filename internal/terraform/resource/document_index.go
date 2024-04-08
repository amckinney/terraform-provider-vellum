package resource

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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

var _ resource.ResourceWithConfigure = (*DocumentIndex)(nil)
var _ resource.ResourceWithImportState = (*DocumentIndex)(nil)

func NewDocumentIndex() resource.Resource {
	return &DocumentIndex{}
}

func (d *DocumentIndex) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_document_index"
}

func (d *DocumentIndex) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Document Index resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true, // TODO: It looks like the 'name' is meant to be the primary thing?
				Description:         "The Document Index's ID",
				MarkdownDescription: "The Document Index's ID",
			},
			"name": schema.StringAttribute{
				Required:            true, // TODO: Should this be optional?
				Description:         "A name that uniquely identifies this index within its workspace",
				MarkdownDescription: "A name that uniquely identifies this index within its workspace",
			},
			"label": schema.StringAttribute{
				Required:            true, // TODO: Should this be optional?
				Description:         "A human-readable label for the document index",
				MarkdownDescription: "A human-readable label for the document index",
			},
			"environment": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The environment this document index is used in\n\n* `DEVELOPMENT` - Development\n* `STAGING` - Staging\n* `PRODUCTION` - Production",
				MarkdownDescription: "The environment this document index is used in\n\n* `DEVELOPMENT` - Development\n* `STAGING` - Staging\n* `PRODUCTION` - Production",
			},
			"status": schema.StringAttribute{
				Optional:            true,
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

func (d *DocumentIndex) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

func (d *DocumentIndex) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model *DocumentIndexModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	request, err := d.modelToCreateRequest(model)
	if err != nil {
		resp.Diagnostics.AddError(
			"Create document index error",
			fmt.Sprintf("Unable to create document index request: %v", err),
		)
		return
	}

	response, err := d.client.DocumentIndexes.Create(
		ctx,
		request,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Create document index error",
			fmt.Sprintf("Unable to create document index, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(
			ctx,
			d.createResponseToModel(response),
		)...,
	)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *DocumentIndex) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var documentIndexState TfDocumentIndexModel
	var err error
	resp.Diagnostics.Append(req.State.Get(ctx, &documentIndexState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	documentIndex, err := d.client.DocumentIndexes.Retrieve(ctx, documentIndexState.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read document index, got error: %s", err))
		return
	}

	documentIndexModel, diagnostic := NewTfDocumentIndexModel(ctx, &documentIndexState, documentIndex)
	resp.Diagnostics.Append(diagnostic...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &documentIndexModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *DocumentIndex) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var documentIndexPlan *TfDocumentIndexModel
	var documentIndexState *TfDocumentIndexModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &documentIndexPlan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &documentIndexState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := documentIndexState.Id.ValueString()
	label := documentIndexPlan.Label.ValueString()

	var status *vellum.EntityStatus
	if documentIndexPlan.Status.ValueString() != "" {
		s, _ := vellum.NewEntityStatusFromString(documentIndexPlan.Status.ValueString())
		status = &s
	}

	var environment *vellum.EnvironmentEnum
	if documentIndexPlan.Environment.ValueString() != "" {
		env, _ := vellum.NewEnvironmentEnumFromString(documentIndexPlan.Environment.ValueString())
		environment = &env
	}

	documentIndex, err := d.client.DocumentIndexes.PartialUpdate(ctx,
		id,
		&vellum.PatchedDocumentIndexUpdateRequest{
			Label:       &label,
			Status:      status,
			Environment: environment,
		})

	if err != nil {
		resp.Diagnostics.AddError("error during document index update", err.Error())
		return
	}

	documentIndexModel, diagnostic := NewTfDocumentIndexModel(ctx, documentIndexPlan, documentIndex)
	resp.Diagnostics.Append(diagnostic...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &documentIndexModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *DocumentIndex) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var documentIndexState *TfDocumentIndexModel

	resp.Diagnostics.Append(req.State.Get(ctx, &documentIndexState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := d.client.DocumentIndexes.Destroy(
		ctx,
		documentIndexState.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("error when destroying the document index resource", err.Error())
		return
	}
}

func (d *DocumentIndex) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (d *DocumentIndex) modelToCreateRequest(model *DocumentIndexModel) (*vellum.DocumentIndexCreateRequest, error) {
	var status *vellum.EntityStatus
	if !model.Status.IsNull() {
		value, err := vellum.NewEntityStatusFromString(model.Status.ValueString())
		if err != nil {
			return nil, err
		}
		status = value.Ptr()
	}

	var environment *vellum.EnvironmentEnum
	if !model.Environment.IsNull() {
		value, err := vellum.NewEnvironmentEnumFromString(model.Environment.ValueString())
		if err != nil {
			return nil, err
		}
		environment = value.Ptr()
	}

	return &vellum.DocumentIndexCreateRequest{
		Label:       model.Label.ValueString(),
		Name:        model.Name.ValueString(),
		Status:      status,
		Environment: environment,
	}, nil
}

func (d *DocumentIndex) createResponseToModel(response *vellum.DocumentIndexRead) *DocumentIndexModel {
	return &DocumentIndexModel{
		Id:          types.StringValue(response.Id),
		Name:        types.StringValue(response.Name),
		Created:     types.StringValue(response.Created.Format(time.RFC3339)),
		Environment: types.StringValue(string(*response.Environment)),
		Label:       types.StringValue(response.Label),
		Status:      types.StringValue(string(*response.Status)),
	}
}
