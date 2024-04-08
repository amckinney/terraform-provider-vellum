package datasource

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	basedatasource "github.com/vellum-ai/terraform-provider-vellum/internal/terraform/base/datasource"
)

type DocumentIndex struct {
	datasource.DataSource
}

var _ datasource.DataSource = (*DocumentIndex)(nil)

func NewDocumentIndex() datasource.DataSource {
	return &DocumentIndex{
		DataSource: basedatasource.NewDocumentIndex(),
	}
}
