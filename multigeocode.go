// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai

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

// MultigeocodeService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMultigeocodeService] method instead.
type MultigeocodeService struct {
	Options []option.RequestOption
	Place   MultigeocodePlaceService
}

// NewMultigeocodeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMultigeocodeService(opts ...option.RequestOption) (r MultigeocodeService) {
	r = MultigeocodeService{}
	r.Options = opts
	r.Place = NewMultigeocodePlaceService(opts...)
	return
}

// The method enables searching for known places from multiple data sources
//
// Use this method to find known places in default or your own custom (proprietary)
// dataset and get a combined search result. It accepts free-form, partially
// correct or even incomplete search texts. Results would be ranked based on the
// search score of a place.
func (r *MultigeocodeService) Search(ctx context.Context, params MultigeocodeSearchParams, opts ...option.RequestOption) (res *MultigeocodeSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "multigeocode/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type MultigeocodeSearchResponse struct {
	// An array of objects containing the search result response. Each object
	// represents one place returned in the search response. An empty array would be
	// returned if no matching place is found.
	Entities []MultigeocodeSearchResponseEntity `json:"entities"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultigeocodeSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *MultigeocodeSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MultigeocodeSearchResponseEntity struct {
	// It contains information about the dataset that returns the specific result
	DataSource MultigeocodeSearchResponseEntityDataSource `json:"dataSource"`
	// The unique NextBillion ID for the result item. This ID can be used as input in
	// “Get Place”, “Update Place”, “Delete Place” methods.
	DocID string `json:"docId"`
	// This parameter represents the place details, including geographical information,
	// address and other related information.
	Place []PlaceItem `json:"place"`
	// Integer value representing how good the result is. Higher score indicates a
	// better match between the search query and the result. This can be used to accept
	// or reject the results depending on how “relevant” a result is, for a given use
	// case
	Score int64 `json:"score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSource  respjson.Field
		DocID       respjson.Field
		Place       respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultigeocodeSearchResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MultigeocodeSearchResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// It contains information about the dataset that returns the specific result
type MultigeocodeSearchResponseEntityDataSource struct {
	// This parameter represents the unique reference ID associated with the data
	// source.
	RefID string `json:"refId"`
	// This parameter represents the source of the data.
	Source string `json:"source"`
	// This parameter indicates if a place is searchable.
	//
	// Any of "enable", "disable".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RefID       respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultigeocodeSearchResponseEntityDataSource) RawJSON() string { return r.JSON.raw }
func (r *MultigeocodeSearchResponseEntityDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MultigeocodeSearchParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Specify the center of the search context expressed as coordinates.
	At MultigeocodeSearchParamsAt `json:"at,omitzero,required"`
	// A free-form, complete or incomplete string to be searched. It allows searching
	// for places using keywords or names.
	Query string `json:"query,required"`
	// Specifies the primary city of the place.
	City param.Opt[string] `json:"city,omitzero"`
	// Country of the search context provided as comma-separated
	// [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) country
	// codes.
	// Note: Country codes should be provided in uppercase.
	Country param.Opt[string] `json:"country,omitzero"`
	// Specifies the district of the search place.
	District param.Opt[string] `json:"district,omitzero"`
	// Sets the maximum number of results to be returned.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Filters the results to places within the specified radius from the 'at'
	// location.
	//
	// Note: Supports 'meter' (m) and 'kilometer' (km) units. If no radius is given,
	// the search method returns as many results as specified in limit.
	Radius param.Opt[string] `json:"radius,omitzero"`
	// Specifies the state of the search place.
	State param.Opt[string] `json:"state,omitzero"`
	// Specifies the street name of the search place.
	Street param.Opt[string] `json:"street,omitzero"`
	// Specifies the subDistrict of the search place.
	SubDistrict param.Opt[string] `json:"subDistrict,omitzero"`
	paramObj
}

func (r MultigeocodeSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow MultigeocodeSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultigeocodeSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MultigeocodeSearchParams]'s query parameters as
// `url.Values`.
func (r MultigeocodeSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the center of the search context expressed as coordinates.
//
// The properties Lat, Lng are required.
type MultigeocodeSearchParamsAt struct {
	// Latitude coordinate of the location
	Lat float64 `json:"lat,required"`
	// Longitude coordinate of the location.
	Lng float64 `json:"lng,required"`
	paramObj
}

func (r MultigeocodeSearchParamsAt) MarshalJSON() (data []byte, err error) {
	type shadow MultigeocodeSearchParamsAt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultigeocodeSearchParamsAt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
