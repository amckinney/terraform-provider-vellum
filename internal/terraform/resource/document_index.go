package resource

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	baseresource "github.com/vellum-ai/terraform-provider-vellum/internal/terraform/base/resource"
)

type DocumentIndex struct {
	resource.Resource
}

var _ resource.Resource = (*DocumentIndex)(nil)

func NewDocumentIndex() resource.Resource {
	return &DocumentIndex{
		Resource: baseresource.NewDocumentIndex(),
	}
}
