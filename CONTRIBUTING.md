# Terraform Provider

## Package Structure

The generated SDK is deposited in the `internal/vellum` package, whereas the Terraform-specific
code is generated in the `internal/terraform` package. The main files you will be interacting with
and/or reading are shown below:

```sh
.
├── internal
│   ├── terraform
│   │   ├── base
│   │   │   ├── datasource
│   │   │   │   └── ...
│   │   │   ├── provider.go
│   │   │   └── resource
│   │   │       └── ...
│   │   ├── datasource
│   │   │   └── ...
│   │   ├── provider.go
│   │   └── resource
│   │       └── ...
│   └── vellum
│       └── ...
└── main.go
```

## Overrides

The package structure is setup so that it's easy to extend the behavior of the generated code. Maintainers can
simply edit the types deposited as the root of the `internal/terraform` package, which includes the following:

- `internal/terraform/provider.go`
- `internal/terraform/datasource/...`
- `internal/terraform/resource/...`

Note that _none_ of the files generated in the `internal/terraform/base/...` directory are expected to be edited.

### Adding an override

Suppose that we need to edit the property used to retrieve a data source. By default, the terraform provider will
use the resource's primary key to retrieve the data source, but the data source might support multiple ways to be
resolved, such as by their unique ID or name.

For example, consider the `DocumentIndex` data source defined in the `internal/terraform/base/datasource/document_index.go` file.
The `name` property is required, so the data source is expressed in Terraform with the following configuration:

```terraform
terraform {
  required_providers {
    vellum = {
      source = "hashicorp.com/ai/vellum"
    }
  }
}

provider "vellum" {}

data "vellum_document_index" "reference" {
  name = "reference"
}
```

We can refactor the terraform provider to support additional data source parameters in a few simple steps.

#### 1. Inspect the base data source

Whenever we need to make an override, it's best to start with the genearted code as a base, then add and/or remove
only the pieces that we're concerned with. In this case, we're going to make a schema change to adapt how we
read the data source, so we will only need to understand the `Schemas` and `Read` methods shown below (notice the
`name` property is marked as _required_):

```go
type DocumentIndex struct {
	Vellum *vellumclient.Client
}

func (d *DocumentIndex) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
    ...
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
      ...
		},
	}
}

func (d *DocumentIndex) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model *DocumentIndexModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := d.Vellum.DocumentIndexes.Retrieve(
		ctx,
    	model.Name.ValueString(),
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
```

#### 2. Refactor the methods

Now we can actually refactor the implementation so that we support reading both the ID and name properties.
To do so, we'll need to make these properties _optional_, then interact with both properties before we call
Vellum's `client.DocumentIndexes.Retrieve` endpoint.

We could re-implement the entire method from scratch, or use the base implementation as much as we can to
reduce code duplication. Both approaches are shown below:

```go
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

	response, err := d.Vellum.DocumentIndexes.Retrieve(
		ctx,
    	retrieveID,
	)

  ...
}
```

That's it! The terraform provider already registers the top-level document index data source, so no
further changes are required. Plus, the rest of the document index data source's behavior is left unchanged.
