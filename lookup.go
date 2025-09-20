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
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// LookupService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLookupService] method instead.
type LookupService struct {
	Options []option.RequestOption
}

// NewLookupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLookupService(opts ...option.RequestOption) (r LookupService) {
	r = LookupService{}
	r.Options = opts
	return
}

// Lookup By ID
func (r *LookupService) ByID(ctx context.Context, query LookupByIDParams, opts ...option.RequestOption) (res *LookupByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type LookupByIDResponse struct {
	// The results are presented as a JSON list of candidates in ranked order
	// (most-likely to least-likely) based on the matched location criteria.
	Items []LookupByIDResponseItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *LookupByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupByIDResponseItem struct {
	// The unique identifier for the result item.
	ID string `json:"id"`
	// An array returning the location coordinates of all the access points of the
	// search result.
	Access Access `json:"access"`
	// Postal address of the result item.
	Address Address `json:"address"`
	// The list of categories assigned to this place.
	Categories []Categories `json:"categories"`
	// Contact information like phone, email or website.
	Contacts []Contacts `json:"contacts"`
	// The bounding box enclosing the geometric shape (area or line) that an individual
	// result covers. place typed results have no mapView.
	MapView MapView `json:"mapView"`
	// Returns the location coordinates of the result.
	Position Position `json:"position"`
	// The localized display name of this result item.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Access      respjson.Field
		Address     respjson.Field
		Categories  respjson.Field
		Contacts    respjson.Field
		MapView     respjson.Field
		Position    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupByIDResponseItem) RawJSON() string { return r.JSON.raw }
func (r *LookupByIDResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LookupByIDParams struct {
	// Specify the unique identifier of a specific POI, Street, Geography, Point
	// Address or other entities to retrieve its details.
	ID string `query:"id,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [LookupByIDParams]'s query parameters as `url.Values`.
func (r LookupByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
