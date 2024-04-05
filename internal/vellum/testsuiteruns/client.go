// This file was auto-generated by Fern from our API Definition.

package testsuiteruns

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

// Trigger a Test Suite and create a new Test Suite Run
func (c *Client) Create(
	ctx context.Context,
	request *vellum.TestSuiteRunCreateRequest,
	opts ...option.RequestOption,
) (*vellum.TestSuiteRunRead, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/test-suite-runs"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellum.TestSuiteRunRead
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve a specific Test Suite Run by ID
func (c *Client) Retrieve(
	ctx context.Context,
	// A UUID string identifying this test suite run.
	id string,
	opts ...option.RequestOption,
) (*vellum.TestSuiteRunRead, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/test-suite-runs/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellum.TestSuiteRunRead
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

func (c *Client) ListExecutions(
	ctx context.Context,
	// A UUID string identifying this test suite run.
	id string,
	request *vellum.TestSuiteRunsListExecutionsRequest,
	opts ...option.RequestOption,
) (*vellum.PaginatedTestSuiteRunExecutionList, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/test-suite-runs/%v/executions", id)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellum.PaginatedTestSuiteRunExecutionList
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