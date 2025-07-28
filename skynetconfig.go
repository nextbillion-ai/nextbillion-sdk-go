// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"net/http"
	"net/url"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apiquery"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// SkynetConfigService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetConfigService] method instead.
type SkynetConfigService struct {
	Options []option.RequestOption
}

// NewSkynetConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSkynetConfigService(opts ...option.RequestOption) (r SkynetConfigService) {
	r = SkynetConfigService{}
	r.Options = opts
	return
}

// Update webhook configuration
func (r *SkynetConfigService) New(ctx context.Context, params SkynetConfigNewParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Get webhook configuration
func (r *SkynetConfigService) List(ctx context.Context, query SkynetConfigListParams, opts ...option.RequestOption) (res *SkynetConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Test webhook configurations
func (r *SkynetConfigService) Testwebhook(ctx context.Context, body SkynetConfigTestwebhookParams, opts ...option.RequestOption) (res *SkynetConfigTestwebhookResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/config/testwebhook"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SkynetConfigListResponse struct {
	// A data object containing the `config` response.
	Data SkynetConfigListResponseData `json:"data"`
	// Displays the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Message string `json:"message"`
	// A string indicating the state of the response. On successful responses, the
	// value will be `Ok`. Indicative error messages are returned for different errors.
	// See the [API Error Codes](#api-error-codes) section below for more information.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetConfigListResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetConfigListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A data object containing the `config` response.
type SkynetConfigListResponseData struct {
	Config SkynetConfigListResponseDataConfig `json:"config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetConfigListResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetConfigListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetConfigListResponseDataConfig struct {
	// An array of strings representing the list of webhooks. Webhooks are used to
	// receive information, through POST requests, whenever any event is triggered.
	Webhook []string `json:"webhook"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Webhook     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetConfigListResponseDataConfig) RawJSON() string { return r.JSON.raw }
func (r *SkynetConfigListResponseDataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetConfigTestwebhookResponse struct {
	// A string indicating the state of the response. Please note this value will
	// always be `Ok`.
	//
	// The sample event information will be received on the webhook, if they were
	// successfully configured. If no event information is received by the webhook,
	// please reconfigure the webhook or else reach out to
	// [support@nextbillion.ai](mailto:support@nextbillion.ai) for help.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetConfigTestwebhookResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetConfigTestwebhookResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetConfigNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetConfigNewParamsCluster `query:"cluster,omitzero" json:"-"`
	// Use this array to update information about the webhooks. Please note that the
	// webhooks will be overwritten every time this method is used.
	Webhook []string `json:"webhook,omitzero"`
	paramObj
}

func (r SkynetConfigNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetConfigNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetConfigNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetConfigNewParams]'s query parameters as `url.Values`.
func (r SkynetConfigNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetConfigNewParamsCluster string

const (
	SkynetConfigNewParamsClusterAmerica SkynetConfigNewParamsCluster = "america"
)

type SkynetConfigListParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetConfigListParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetConfigListParams]'s query parameters as `url.Values`.
func (r SkynetConfigListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetConfigListParamsCluster string

const (
	SkynetConfigListParamsClusterAmerica SkynetConfigListParamsCluster = "america"
)

type SkynetConfigTestwebhookParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetConfigTestwebhookParams]'s query parameters as
// `url.Values`.
func (r SkynetConfigTestwebhookParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
