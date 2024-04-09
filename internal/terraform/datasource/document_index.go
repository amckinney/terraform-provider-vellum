package datasource

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	basedatasource "github.com/vellum-ai/terraform-provider-vellum/internal/terraform/base/datasource"
)

type DocumentIndex struct {
	*basedatasource.DocumentIndex
}

var _ datasource.DataSource = (*DocumentIndex)(nil)

func NewDocumentIndex() datasource.DataSource {
	return &DocumentIndex{
		DocumentIndex: basedatasource.NewDocumentIndex().(*basedatasource.DocumentIndex),
	}
}
