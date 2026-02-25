// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apiquery"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// BatchService contains methods and other services that help with interacting with
// the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatchService] method instead.
type BatchService struct {
	Options []option.RequestOption
}

// NewBatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBatchService(opts ...option.RequestOption) (r BatchService) {
	r = BatchService{}
	r.Options = opts
	return
}

// Create Batch Routing
func (r *BatchService) New(ctx context.Context, params BatchNewParams, opts ...option.RequestOption) (res *BatchNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get Batch Result
func (r *BatchService) Get(ctx context.Context, query BatchGetParams, opts ...option.RequestOption) (res *BatchGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BatchNewResponse struct {
	// Displays the error message in case of a failed request or operation. Please note
	// that this parameter is not returned in the response in case of a successful
	// request.
	Msg string `json:"msg"`
	// Returns the overall status of the API request. Its value will belong to one of
	// success, failed, and pending. It can also contain HTTP error codes in case of a
	// failed request or operation.
	Status string `json:"status"`
	// Returns the unique ID of the batch processing task. Use this ID using the GET
	// request to retrieve the solution once the task processing is completed.
	TrackID string `json:"track_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		Status      respjson.Field
		TrackID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponse struct {
	// Displays the error message in case of a failed request or operation. Please note
	// that this parameter is not returned in the response in case of a successful
	// request.
	Msg string `json:"msg"`
	// An array of objects returning the results of all the individual routing queries
	// specified in the input. Each object represents the solution to an individual
	// query in the input.
	Responses []BatchGetResponseResponse `json:"responses"`
	// Returns the overall status of the API request. Its value will always be one of
	// success, failed, and pending.
	Status string `json:"status"`
	// Returns the unique ID of the batch processing task.
	TrackID string `json:"track_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		Responses   respjson.Field
		Status      respjson.Field
		TrackID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponseResponse struct {
	// An object returning the routing solution of an individual query. The JSON format
	// and structure of the response would vary depending on the routing endpoint used
	// in each individual query. However, it will be consistent with standard response
	// for a given routing endpoint.
	Response any `json:"response"`
	// Returns the HTTP status code for the individual routing request. See the
	// [API Errors Codes](#api-error-codes) section below for more information.
	StatusCode int64 `json:"status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response    respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// An array of objects to collect the details of individual routing queries that
	// will form a batch.
	Requests []BatchNewParamsRequest `json:"requests,omitzero"`
	paramObj
}

func (r BatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BatchNewParams]'s query parameters as `url.Values`.
func (r BatchNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BatchNewParamsRequest struct {
	// Specify the routing query in the form of a string. The supported attributes and
	// their formats are consistent with the standard routing endpoint that is being
	// used as part of the batch. Check the
	// [Sample Request](https://docs.nextbillion.ai/docs/navigation/batch-routing-api#sample-request-1)
	// section for an example request.
	Query param.Opt[string] `json:"query,omitzero"`
	paramObj
}

func (r BatchNewParamsRequest) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParamsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParamsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// Specify the track ID of the batch that was returned in the response after
	// submitting a successful batch request.
	TrackID string `query:"track_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [BatchGetParams]'s query parameters as `url.Values`.
func (r BatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
