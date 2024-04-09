package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	baseresource "github.com/vellum-ai/terraform-provider-vellum/internal/terraform/base/resource"
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
	d.base.Create(ctx, req, resp)
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
