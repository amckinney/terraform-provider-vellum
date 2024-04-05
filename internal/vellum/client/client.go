// This file was auto-generated by Fern from our API Definition.

package client

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	vellum "github.com/vellum-ai/terraform-provider-vellum/internal/vellum"
	core "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/core"
	deployments "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/deployments"
	documentindexes "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/documentindexes"
	documents "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/documents"
	folderentities "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/folderentities"
	modelversions "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/modelversions"
	option "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/option"
	registeredprompts "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/registeredprompts"
	sandboxes "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/sandboxes"
	testsuiteruns "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/testsuiteruns"
	testsuites "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/testsuites"
	workflowdeployments "github.com/vellum-ai/terraform-provider-vellum/internal/vellum/workflowdeployments"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Deployments         *deployments.Client
	DocumentIndexes     *documentindexes.Client
	Documents           *documents.Client
	FolderEntities      *folderentities.Client
	ModelVersions       *modelversions.Client
	RegisteredPrompts   *registeredprompts.Client
	Sandboxes           *sandboxes.Client
	TestSuiteRuns       *testsuiteruns.Client
	TestSuites          *testsuites.Client
	WorkflowDeployments *workflowdeployments.Client
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
		header:              options.ToHeader(),
		Deployments:         deployments.NewClient(opts...),
		DocumentIndexes:     documentindexes.NewClient(opts...),
		Documents:           documents.NewClient(opts...),
		FolderEntities:      folderentities.NewClient(opts...),
		ModelVersions:       modelversions.NewClient(opts...),
		RegisteredPrompts:   registeredprompts.NewClient(opts...),
		Sandboxes:           sandboxes.NewClient(opts...),
		TestSuiteRuns:       testsuiteruns.NewClient(opts...),
		TestSuites:          testsuites.NewClient(opts...),
		WorkflowDeployments: workflowdeployments.NewClient(opts...),
	}
}

// Executes a deployed Prompt and returns the result.
func (c *Client) ExecutePrompt(
	ctx context.Context,
	request *vellum.ExecutePromptRequest,
	opts ...option.RequestOption,
) (*vellum.ExecutePromptResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/execute-prompt"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellum.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(vellum.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellum.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellum.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *vellum.ExecutePromptResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Executes a deployed Prompt and streams back the results.
func (c *Client) ExecutePromptStream(
	ctx context.Context,
	request *vellum.ExecutePromptStreamRequest,
	opts ...option.RequestOption,
) (*core.Stream[vellum.ExecutePromptEvent], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/execute-prompt-stream"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellum.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(vellum.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellum.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellum.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	streamer := core.NewStreamer[vellum.ExecutePromptEvent](c.caller)
	return streamer.Stream(
		ctx,
		&core.StreamParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			ErrorDecoder: errorDecoder,
		},
	)
}

// Executes a deployed Workflow and returns its outputs.
func (c *Client) ExecuteWorkflow(
	ctx context.Context,
	request *vellum.ExecuteWorkflowRequest,
	opts ...option.RequestOption,
) (*vellum.ExecuteWorkflowResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/execute-workflow"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellum.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellum.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellum.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *vellum.ExecuteWorkflowResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Executes a deployed Workflow and streams back its results.
func (c *Client) ExecuteWorkflowStream(
	ctx context.Context,
	request *vellum.ExecuteWorkflowStreamRequest,
	opts ...option.RequestOption,
) (*core.Stream[vellum.WorkflowStreamEvent], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/execute-workflow-stream"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellum.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellum.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellum.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	streamer := core.NewStreamer[vellum.WorkflowStreamEvent](c.caller)
	return streamer.Stream(
		ctx,
		&core.StreamParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			ErrorDecoder: errorDecoder,
		},
	)
}

// Generate a completion using a previously defined deployment.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) Generate(
	ctx context.Context,
	request *vellum.GenerateBodyRequest,
	opts ...option.RequestOption,
) (*vellum.GenerateResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/generate"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellum.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(vellum.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellum.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellum.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *vellum.GenerateResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Generate a stream of completions using a previously defined deployment.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) GenerateStream(
	ctx context.Context,
	request *vellum.GenerateStreamBodyRequest,
	opts ...option.RequestOption,
) (*core.Stream[vellum.GenerateStreamResponse], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/generate-stream"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellum.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(vellum.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellum.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellum.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	streamer := core.NewStreamer[vellum.GenerateStreamResponse](c.caller)
	return streamer.Stream(
		ctx,
		&core.StreamParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			ErrorDecoder: errorDecoder,
		},
	)
}

// Perform a search against a document index.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) Search(
	ctx context.Context,
	request *vellum.SearchRequestBodyRequest,
	opts ...option.RequestOption,
) (*vellum.SearchResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/search"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellum.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellum.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellum.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *vellum.SearchResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Used to submit feedback regarding the quality of previously generated completions.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) SubmitCompletionActuals(
	ctx context.Context,
	request *vellum.SubmitCompletionActualsRequest,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/submit-completion-actuals"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellum.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellum.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellum.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

// Used to submit feedback regarding the quality of previous workflow execution and its outputs.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) SubmitWorkflowExecutionActuals(
	ctx context.Context,
	request *vellum.SubmitWorkflowExecutionActualsRequest,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "v1/submit-workflow-execution-actuals"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
		},
	); err != nil {
		return err
	}
	return nil
}
