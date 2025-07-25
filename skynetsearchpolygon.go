// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apijson"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/nextbillion-sdk-go/option"
	"github.com/stainless-sdks/nextbillion-sdk-go/packages/param"
)

// SkynetSearchPolygonService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetSearchPolygonService] method instead.
type SkynetSearchPolygonService struct {
	Options []option.RequestOption
}

// NewSkynetSearchPolygonService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSkynetSearchPolygonService(opts ...option.RequestOption) (r SkynetSearchPolygonService) {
	r = SkynetSearchPolygonService{}
	r.Options = opts
	return
}

// Polygon Search
func (r *SkynetSearchPolygonService) New(ctx context.Context, params SkynetSearchPolygonNewParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/search/polygon"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Polygon Search
func (r *SkynetSearchPolygonService) Get(ctx context.Context, query SkynetSearchPolygonGetParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/search/polygon"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SkynetSearchPolygonNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// An object to collect geoJSON details of a custom polygon. Please ensure that:
	//
	//   - the `polygon` provided is enclosed. This can be achieved by making the last
	//     location coordinate in the list equal to the first location coordinate of the
	//     list.
	//
	// - the 'polygon' provided does not contain multiple rings.
	//
	// The contents of this object follow the
	// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	//
	// Please note that the maximum area of the search polygon allowed is 3000
	// km<sup>2</sup>.
	Polygon SkynetSearchPolygonNewParamsPolygon `json:"polygon,omitzero,required"`
	// **`tags` parameter will be deprecated soon! Please use the
	// `include_any_of_attributes` or `include_all_of_attributes` parameters to match
	// assets based on their labels or markers.**
	//
	// Use this parameter to filter the assets found inside the specified area by their
	// `tag`. Multiple `tag` can be separated using comma (`,`).
	//
	// Please note the tags are case sensitive.
	Filter param.Opt[string] `json:"filter,omitzero"`
	// if ture, can get 16x bigger limitation in search.
	MaxSearchLimit param.Opt[bool] `json:"max_search_limit,omitzero"`
	// Denotes page number. Use this along with the `ps` parameter to implement
	// pagination for your searched results. This parameter does not have a maximum
	// limit but would return an empty response in case a higher value is provided when
	// the result-set itself is smaller.
	Pn param.Opt[int64] `json:"pn,omitzero"`
	// Denotes number of search results per page. Use this along with the `pn`
	// parameter to implement pagination for your searched results. Please note that
	// `ps` has a default value of 20 and accepts integers only in the range of [1,
	// 100].
	Ps param.Opt[int64] `json:"ps,omitzero"`
	// An object to define the `attributes` which will be used to filter the assets
	// found within the `polygon`.
	MatchFilter SkynetSearchPolygonNewParamsMatchFilter `json:"match_filter,omitzero"`
	Sort        SkynetSearchPolygonNewParamsSort        `json:"sort,omitzero"`
	paramObj
}

func (r SkynetSearchPolygonNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetSearchPolygonNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetSearchPolygonNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetSearchPolygonNewParams]'s query parameters as
// `url.Values`.
func (r SkynetSearchPolygonNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// An object to collect geoJSON details of a custom polygon. Please ensure that:
//
//   - the `polygon` provided is enclosed. This can be achieved by making the last
//     location coordinate in the list equal to the first location coordinate of the
//     list.
//
// - the 'polygon' provided does not contain multiple rings.
//
// The contents of this object follow the
// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
//
// Please note that the maximum area of the search polygon allowed is 3000
// km<sup>2</sup>.
//
// The properties Coordinates, Type are required.
type SkynetSearchPolygonNewParamsPolygon struct {
	// An array of coordinates in the [longitude, latitude] format, representing the
	// polygon boundary.
	Coordinates []float64 `json:"coordinates,omitzero,required"`
	// Type of the geoJSON geometry. Should always be `polygon`.
	Type string `json:"type,required"`
	paramObj
}

func (r SkynetSearchPolygonNewParamsPolygon) MarshalJSON() (data []byte, err error) {
	type shadow SkynetSearchPolygonNewParamsPolygon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetSearchPolygonNewParamsPolygon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object to define the `attributes` which will be used to filter the assets
// found within the `polygon`.
type SkynetSearchPolygonNewParamsMatchFilter struct {
	// Use this parameter to filter the assets found inside the specified area by their
	// `attributes`. Only the assets having all the `attributes` that are added to this
	// parameter, will be returned in the search results. Multiple `attributes` can be
	// separated using commas (`,`).
	//
	// Please note the attributes are case sensitive. Also, this parameter can not be
	// used in conjunction with `include_any_of_attributes` parameter.
	IncludeAllOfAttributes param.Opt[string] `json:"include_all_of_attributes,omitzero"`
	// Use this parameter to filter the assets found inside the specified area by their
	// `attributes`. Assets having at least one of the `attributes` added to this
	// parameter, will be returned in the search results. Multiple `attributes` can be
	// separated using commas (`,`).
	//
	// Please note the attributes are case sensitive. Also, this parameter can not be
	// used in conjunction with `include_all_of_attributes` parameter.
	IncludeAnyOfAttributes param.Opt[string] `json:"include_any_of_attributes,omitzero"`
	paramObj
}

func (r SkynetSearchPolygonNewParamsMatchFilter) MarshalJSON() (data []byte, err error) {
	type shadow SkynetSearchPolygonNewParamsMatchFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetSearchPolygonNewParamsMatchFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetSearchPolygonNewParamsSort struct {
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
	SortBy string `json:"sort_by,omitzero"`
	// Specifies the location coordinates of the point which acts as destination for
	// sorting the assets in the search results. The service will sort each asset based
	// on the driving distance or travel time to this destination, from its current
	// location. Use the `sort_by` parameter to configure the metric that should be
	// used for sorting the assets. Please note that `sort_destination` is required
	// when `sort_by` is provided.
	SortDestination SkynetSearchPolygonNewParamsSortSortDestination `json:"sort_destination,omitzero"`
	// Specifies the driving mode to be used for determining travel duration or driving
	// distance for sorting the assets in search result.
	//
	// Any of "`car`", "`truck`".
	SortDrivingMode string `json:"sort_driving_mode,omitzero"`
	paramObj
}

func (r SkynetSearchPolygonNewParamsSort) MarshalJSON() (data []byte, err error) {
	type shadow SkynetSearchPolygonNewParamsSort
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetSearchPolygonNewParamsSort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SkynetSearchPolygonNewParamsSort](
		"sort_by", "`distance`", "`duration`", "`straight_distance`",
	)
	apijson.RegisterFieldValidator[SkynetSearchPolygonNewParamsSort](
		"sort_driving_mode", "`car`", "`truck`",
	)
}

// Specifies the location coordinates of the point which acts as destination for
// sorting the assets in the search results. The service will sort each asset based
// on the driving distance or travel time to this destination, from its current
// location. Use the `sort_by` parameter to configure the metric that should be
// used for sorting the assets. Please note that `sort_destination` is required
// when `sort_by` is provided.
//
// The properties Lat, Lon are required.
type SkynetSearchPolygonNewParamsSortSortDestination struct {
	// Latitude of the destination location
	Lat float64 `json:"lat,required"`
	// Longitude of the destination location
	Lon float64 `json:"lon,required"`
	paramObj
}

func (r SkynetSearchPolygonNewParamsSortSortDestination) MarshalJSON() (data []byte, err error) {
	type shadow SkynetSearchPolygonNewParamsSortSortDestination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetSearchPolygonNewParamsSortSortDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetSearchPolygonGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Define a custom polygon enclosing the area to be searched. It should be a pipe
	// (`|`) delimited list of location coordinates.
	//
	// Please ensure that the `polygon` provided is enclosed. This can be achieved by
	// making the last location coordinate in the list equal to the first location
	// coordinate of the list.
	//
	// Please note that the maximum area of the search polygon allowed is 3000
	// km<sup>2</sup>.
	Polygon string `query:"polygon,required" format:"latitude_1,longitude_1|latitude_2,longitude_2|..." json:"-"`
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
	SortBy SkynetSearchPolygonGetParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Specifies the driving mode to be used for determining travel duration or driving
	// distance for sorting the assets in search result.
	//
	// Any of "`car`", "`truck`".
	SortDrivingMode SkynetSearchPolygonGetParamsSortDrivingMode `query:"sort_driving_mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetSearchPolygonGetParams]'s query parameters as
// `url.Values`.
func (r SkynetSearchPolygonGetParams) URLQuery() (v url.Values, err error) {
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
type SkynetSearchPolygonGetParamsSortBy string

const (
	SkynetSearchPolygonGetParamsSortByDistance         SkynetSearchPolygonGetParamsSortBy = "`distance`"
	SkynetSearchPolygonGetParamsSortByDuration         SkynetSearchPolygonGetParamsSortBy = "`duration`"
	SkynetSearchPolygonGetParamsSortByStraightDistance SkynetSearchPolygonGetParamsSortBy = "`straight_distance`"
)

// Specifies the driving mode to be used for determining travel duration or driving
// distance for sorting the assets in search result.
type SkynetSearchPolygonGetParamsSortDrivingMode string

const (
	SkynetSearchPolygonGetParamsSortDrivingModeCar   SkynetSearchPolygonGetParamsSortDrivingMode = "`car`"
	SkynetSearchPolygonGetParamsSortDrivingModeTruck SkynetSearchPolygonGetParamsSortDrivingMode = "`truck`"
)
