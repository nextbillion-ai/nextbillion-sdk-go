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
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// AreaService contains methods and other services that help with interacting with
// the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAreaService] method instead.
type AreaService struct {
	Options []option.RequestOption
}

// NewAreaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAreaService(opts ...option.RequestOption) (r AreaService) {
	r = AreaService{}
	r.Options = opts
	return
}

// Get available areas
func (r *AreaService) List(ctx context.Context, query AreaListParams, opts ...option.RequestOption) (res *[]AreaListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "areas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AreaListResponse struct {
	// Returns the code for the available area.
	Code string `json:"code"`
	// Returns available traveling modes for the given area.
	Modes []string `json:"modes"`
	// Returns the name of the available area.
	Name string `json:"name"`
	// Returns the offset from UTC time.
	Timezone float64 `json:"timezone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Modes       respjson.Field
		Name        respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AreaListResponse) RawJSON() string { return r.JSON.raw }
func (r *AreaListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AreaListParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [AreaListParams]'s query parameters as `url.Values`.
func (r AreaListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
