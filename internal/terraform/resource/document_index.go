package resource

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	baseresource "github.com/vellum-ai/terraform-provider-vellum/internal/terraform/base/resource"
	"github.com/vellum-ai/terraform-provider-vellum/internal/vellum"
)

type DocumentIndex struct {
	base *baseresource.DocumentIndex
}

var _ resource.Resource = (*DocumentIndex)(nil)

func NewDocumentIndex() resource.Resource {
	return &DocumentIndex{
		base: baseresource.NewDocumentIndex().(*baseresource.DocumentIndex),
	}
}

func (d *DocumentIndex) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	d.base.Metadata(ctx, req, resp)
}

func (d *DocumentIndex) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	d.base.Schema(ctx, req, resp)
}

func (d *DocumentIndex) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	d.base.Configure(ctx, req, resp)
}

func (d *DocumentIndex) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model *baseresource.DocumentIndexModel

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

	response, err := d.base.Vellum.DocumentIndexes.Create(
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
			&baseresource.DocumentIndexModel{
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

func (d *DocumentIndex) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	d.base.Read(ctx, req, resp)
}

func (d *DocumentIndex) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	d.base.Update(ctx, req, resp)
}

func (d *DocumentIndex) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	d.base.Delete(ctx, req, resp)
}

func (d *DocumentIndex) modelToCreateRequest(model *baseresource.DocumentIndexModel) (*vellum.DocumentIndexCreateRequest, error) {
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
		Label:          model.Label.ValueString(),
		Name:           model.Name.ValueString(),
		Status:         status,
		Environment:    environment,
		IndexingConfig: defaultIndexingConfig,
	}, nil
}

// TODO: Replace this with data.indexing_config, improve indexing_config param in vellum backend.
var defaultIndexingConfig = map[string]interface{}{
	"chunking": map[string]interface{}{
		"chunker_name": "sentence-chunker",
		"chunker_config": map[string]interface{}{
			"character_limit":   1000,
			"min_overlap_ratio": 0.5,
		},
	},
	"vectorizer": map[string]interface{}{
		"model_name": "hkunlp/instructor-xl",
		"config": map[string]interface{}{
			"instruction_domain":             "",
			"instruction_document_text_type": "plain_text",
			"instruction_query_text_type":    "plain_text",
		},
	},
}
