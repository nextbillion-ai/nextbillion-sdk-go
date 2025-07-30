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

// SkynetSearchService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetSearchService] method instead.
type SkynetSearchService struct {
	Options []option.RequestOption
	Polygon SkynetSearchPolygonService
}

// NewSkynetSearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSkynetSearchService(opts ...option.RequestOption) (r SkynetSearchService) {
	r = SkynetSearchService{}
	r.Options = opts
	r.Polygon = NewSkynetSearchPolygonService(opts...)
	return
}

// Around Search
func (r *SkynetSearchService) Around(ctx context.Context, query SkynetSearchAroundParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/search/around"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Bound Search
func (r *SkynetSearchService) Bound(ctx context.Context, query SkynetSearchBoundParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/search/bound"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SearchResponse struct {
	// A data object containing the search result.
	Data SearchResponseData `json:"data"`
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
func (r SearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A data object containing the search result.
type SearchResponseData struct {
	// An array of objects with details of the asset(s) returned in the search result.
	// Each object represents one `asset`
	Assets []SearchResponseDataAsset `json:"assets"`
	// An object with pagination details of the search results. Use this object to
	// implement pagination in your application.
	Page Pagination `json:"page"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets      respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchResponseDataAsset struct {
	// ID of `asset` which was last located inside the specified area in the input
	// request. This is the same ID that was generated/provided at the time of creating
	// the `asset`.
	ID string `json:"id"`
	// A UNIX timestamp in seconds representing the time at which the `asset` was
	// created.
	CreatedAt int64 `json:"created_at"`
	// Description of the `asset`. The value would be the same as that provided for the
	// `description` parameter at the time of creating or updating the `asset`.
	Description string `json:"description"`
	// An object with details of the tracked location. Please note that if there are no
	// tracking records for an asset, no location data will be returned.
	LatestLocation TrackLocation `json:"latest_location"`
	// Any valid json object data. Can be used to save customized data. Max size is
	// 65kb.
	MetaData MetaData `json:"meta_data"`
	// Name of `asset`. The value would be the same as that provided for the `name`
	// parameter at the time of creating or updating the `asset`.
	Name string `json:"name"`
	// An object returning the sorting details of the asset as per the configuration
	// specified in the input.
	RankingInfo SearchResponseDataAssetRankingInfo `json:"ranking_info"`
	// **This parameter will be deprecated soon! Please move existing `tags` to
	// `attributes` parameter.**
	//
	// Tags associated with the `asset`.
	Tags []string `json:"tags"`
	// A UNIX epoch timestamp in seconds representing the last time when the `asset`
	// was tracked.
	TrackedAt int64 `json:"tracked_at"`
	// A UNIX timestamp in seconds representing the time at which the `asset` was last
	// updated.
	UpdatedAt int64 `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		LatestLocation respjson.Field
		MetaData       respjson.Field
		Name           respjson.Field
		RankingInfo    respjson.Field
		Tags           respjson.Field
		TrackedAt      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponseDataAsset) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseDataAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning the sorting details of the asset as per the configuration
// specified in the input.
type SearchResponseDataAssetRankingInfo struct {
	// Driving distance between the asset and the `sort_destination`.
	Distance float64 `json:"distance"`
	// Driving duration between the asset and the `sort_destination`. Please note this
	// field in not returned in the response when `sort_by = straight_distance` .
	Duration float64 `json:"duration"`
	// Index of the ranked asset. The index value starts from 0.
	Index int64 `json:"index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Duration    respjson.Field
		Index       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponseDataAssetRankingInfo) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseDataAssetRankingInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetSearchAroundParams struct {
	// Location coordinates of the point which would act as the center of the circular
	// area to be searched.
	Center string `query:"center,required" format:"latitude,longitude" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Radius, in meters, of the circular area to be searched.
	Radius float64 `query:"radius,required" json:"-"`
	// **`tags` parameter will be deprecated soon! Please use the
	// `include_any_of_attributes` or `include_all_of_attributes` parameters to match
	// assets based on their labels or markers.**
	//
	// Use this parameter to filter the assets found inside the specified area by their
	// `tags`. Multiple `tags` can be separated using commas (`,`).
	//
	// Please note the tags are case sensitive.
	Filter param.Opt[string] `query:"filter,omitzero" format:"filter=tag:value_1,value_2..." json:"-"`
	// Use this parameter to filter the assets found inside the specified area by their
	// `attributes`. Only the assets having all the `attributes` that are added to this
	// parameter, will be returned in the search results. Multiple `attributes` can be
	// separated using pipes (`|`).
	//
	// Please note the attributes are case sensitive. Also, this parameter can not be
	// used in conjunction with `include_any_of_attributes` parameter.
	IncludeAllOfAttributes param.Opt[string] `query:"include_all_of_attributes,omitzero" format:"key_1:value_1|key_2:value_2" json:"-"`
	// Use this parameter to filter the assets found inside the specified area by their
	// `attributes`. Assets having at least one of the `attributes` added to this
	// parameter, will be returned in the search results. Multiple `attributes` can be
	// separated using pipes (`|`).
	//
	// Please note the attributes are case sensitive. Also, this parameter can not be
	// used in conjunction with `include_all_of_attributes` parameter.
	IncludeAnyOfAttributes param.Opt[string] `query:"include_any_of_attributes,omitzero" format:"key1:value1|key2:value2|..." json:"-"`
	// When true, the maximum limit is 20Km for around search API and 48000 Km2 for
	// other search methods.
	MaxSearchLimit param.Opt[bool] `query:"max_search_limit,omitzero" json:"-"`
	// Denotes page number. Use this along with the `ps` parameter to implement
	// pagination for your searched results. This parameter does not have a maximum
	// limit but would return an empty response in case a higher value is provided when
	// the result-set itself is smaller.
	Pn param.Opt[int64] `query:"pn,omitzero" json:"-"`
	// Denotes number of search results per page. Use this along with the `pn`
	// parameter to implement pagination for your searched results.
	Ps param.Opt[int64] `query:"ps,omitzero" json:"-"`
	// Specifies the location coordinates of the point which acts as destination for
	// sorting the assets in the search results. The service will sort each asset based
	// on the driving distance or travel time to this destination, from its current
	// location. Use the `sort_by` parameter to configure the metric that should be
	// used for sorting the assets. Please note that `sort_destination` is required
	// when `sort_by` is provided.
	SortDestination param.Opt[string] `query:"sort_destination,omitzero" format:"latitude,lontitude" json:"-"`
	// Specify the metric to sort the assets returned in the search result. The valid
	// values are:
	//
	//   - **distance** : Sorts the assets by driving distance to the given
	//     `sort_destination` .
	//   - **duration** : Sorts the assets by travel time to the given `sort_destination`
	//     .
	//   - **straight_distance** : Sort the assets by straight-line distance to the given
	//     `sort-destination` .
	//
	// Any of "`distance`", "`duration`", "`straight_distance`".
	SortBy SkynetSearchAroundParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Specifies the driving mode to be used for determining travel duration or driving
	// distance for sorting the assets in search result.
	//
	// Any of "`car`", "`truck`".
	SortDrivingMode SkynetSearchAroundParamsSortDrivingMode `query:"sort_driving_mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetSearchAroundParams]'s query parameters as
// `url.Values`.
func (r SkynetSearchAroundParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the metric to sort the assets returned in the search result. The valid
// values are:
//
//   - **distance** : Sorts the assets by driving distance to the given
//     `sort_destination` .
//   - **duration** : Sorts the assets by travel time to the given `sort_destination`
//     .
//   - **straight_distance** : Sort the assets by straight-line distance to the given
//     `sort-destination` .
type SkynetSearchAroundParamsSortBy string

const (
	SkynetSearchAroundParamsSortByDistance         SkynetSearchAroundParamsSortBy = "`distance`"
	SkynetSearchAroundParamsSortByDuration         SkynetSearchAroundParamsSortBy = "`duration`"
	SkynetSearchAroundParamsSortByStraightDistance SkynetSearchAroundParamsSortBy = "`straight_distance`"
)

// Specifies the driving mode to be used for determining travel duration or driving
// distance for sorting the assets in search result.
type SkynetSearchAroundParamsSortDrivingMode string

const (
	SkynetSearchAroundParamsSortDrivingModeCar   SkynetSearchAroundParamsSortDrivingMode = "`car`"
	SkynetSearchAroundParamsSortDrivingModeTruck SkynetSearchAroundParamsSortDrivingMode = "`truck`"
)

type SkynetSearchBoundParams struct {
	// Specify two, pipe (|) delimited location coordinates which would act as corners
	// of the bounding box area to be searched. The first one should be the southwest
	// coordinate of the `bounds` and the second one should be the northeast coordinate
	// of the `bounds`.
	Bound string `query:"bound,required" format:"latitude_1,longitude_2|latitude_1,longitude_2" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// **`tags` parameter will be deprecated soon! Please use the
	// `include_any_of_attributes` or `include_all_of_attributes` parameters to match
	// assets based on their labels or markers.**
	//
	// Use this parameter to filter the assets found inside the specified area by their
	// `tags`. Multiple `tags` can be separated using commas (`,`).
	//
	// Please note the tags are case sensitive.
	Filter param.Opt[string] `query:"filter,omitzero" format:"filter=tag:value_1,value_2..." json:"-"`
	// Use this parameter to filter the assets found inside the specified area by their
	// `attributes`. Only the assets having all the `attributes` that are added to this
	// parameter, will be returned in the search results. Multiple `attributes` can be
	// separated using pipes (`|`).
	//
	// Please note the attributes are case sensitive. Also, this parameter can not be
	// used in conjunction with `include_any_of_attributes` parameter.
	IncludeAllOfAttributes param.Opt[string] `query:"include_all_of_attributes,omitzero" format:"key_1:value_1|key_2:value_2" json:"-"`
	// Use this parameter to filter the assets found inside the specified area by their
	// `attributes`. Assets having at least one of the `attributes` added to this
	// parameter, will be returned in the search results. Multiple `attributes` can be
	// separated using pipes (`|`).
	//
	// Please note the attributes are case sensitive. Also, this parameter can not be
	// used in conjunction with `include_all_of_attributes` parameter.
	IncludeAnyOfAttributes param.Opt[string] `query:"include_any_of_attributes,omitzero" format:"key1:value1|key2:value2|..." json:"-"`
	// When true, the maximum limit is 20Km for around search API and 48000 Km2 for
	// other search methods.
	MaxSearchLimit param.Opt[bool] `query:"max_search_limit,omitzero" json:"-"`
	// Denotes page number. Use this along with the `ps` parameter to implement
	// pagination for your searched results. This parameter does not have a maximum
	// limit but would return an empty response in case a higher value is provided when
	// the result-set itself is smaller.
	Pn param.Opt[int64] `query:"pn,omitzero" json:"-"`
	// Denotes number of search results per page. Use this along with the `pn`
	// parameter to implement pagination for your searched results.
	Ps param.Opt[int64] `query:"ps,omitzero" json:"-"`
	// Specifies the location coordinates of the point which acts as destination for
	// sorting the assets in the search results. The service will sort each asset based
	// on the driving distance or travel time to this destination, from its current
	// location. Use the `sort_by` parameter to configure the metric that should be
	// used for sorting the assets. Please note that `sort_destination` is required
	// when `sort_by` is provided.
	SortDestination param.Opt[string] `query:"sort_destination,omitzero" format:"latitude,lontitude" json:"-"`
	// Specify the metric to sort the assets returned in the search result. The valid
	// values are:
	//
	//   - **distance** : Sorts the assets by driving distance to the given
	//     `sort_destination` .
	//   - **duration** : Sorts the assets by travel time to the given `sort_destination`
	//     .
	//   - **straight_distance** : Sort the assets by straight-line distance to the given
	//     `sort-destination` .
	//
	// Any of "`distance`", "`duration`", "`straight_distance`".
	SortBy SkynetSearchBoundParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Specifies the driving mode to be used for determining travel duration or driving
	// distance for sorting the assets in search result.
	//
	// Any of "`car`", "`truck`".
	SortDrivingMode SkynetSearchBoundParamsSortDrivingMode `query:"sort_driving_mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetSearchBoundParams]'s query parameters as
// `url.Values`.
func (r SkynetSearchBoundParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the metric to sort the assets returned in the search result. The valid
// values are:
//
//   - **distance** : Sorts the assets by driving distance to the given
//     `sort_destination` .
//   - **duration** : Sorts the assets by travel time to the given `sort_destination`
//     .
//   - **straight_distance** : Sort the assets by straight-line distance to the given
//     `sort-destination` .
type SkynetSearchBoundParamsSortBy string

const (
	SkynetSearchBoundParamsSortByDistance         SkynetSearchBoundParamsSortBy = "`distance`"
	SkynetSearchBoundParamsSortByDuration         SkynetSearchBoundParamsSortBy = "`duration`"
	SkynetSearchBoundParamsSortByStraightDistance SkynetSearchBoundParamsSortBy = "`straight_distance`"
)

// Specifies the driving mode to be used for determining travel duration or driving
// distance for sorting the assets in search result.
type SkynetSearchBoundParamsSortDrivingMode string

const (
	SkynetSearchBoundParamsSortDrivingModeCar   SkynetSearchBoundParamsSortDrivingMode = "`car`"
	SkynetSearchBoundParamsSortDrivingModeTruck SkynetSearchBoundParamsSortDrivingMode = "`truck`"
)
