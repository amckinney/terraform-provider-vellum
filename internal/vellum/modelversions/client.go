// This file was auto-generated by Fern from our API Definition.

package modelversions

import (
	context "context"
	fmt "fmt"
	vellum "github.com/vellum-ai/terraform-provider-vellum/internal/vellum"
	core "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/core"
	option "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Deprecated. Use the `deployments/provider-payload` endpoint to fetch information that we send to Model providers.
func (c *Client) Retrieve(
	ctx context.Context,
	// A UUID string identifying this model version.
	id string,
	opts ...option.RequestOption,
) (*vellum.ModelVersionRead, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/model-versions/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellum.ModelVersionRead
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}