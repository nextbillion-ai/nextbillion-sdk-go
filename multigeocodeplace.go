// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apiquery"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// MultigeocodePlaceService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMultigeocodePlaceService] method instead.
type MultigeocodePlaceService struct {
	Options []option.RequestOption
}

// NewMultigeocodePlaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMultigeocodePlaceService(opts ...option.RequestOption) (r MultigeocodePlaceService) {
	r = MultigeocodePlaceService{}
	r.Options = opts
	return
}

// The "Add Place" method allows users to create custom places
//
// Add place method provides the flexibility to create custom places in a way that
// suits your business needs. The newly created place and its attributes can be
// added to custom (proprietary) dataset - to the effect of building your own
// places dataset (s) - or, to a default dataset. Overcome inaccurate ‘POI’ details
// from default search provider by creating custom, highly accurate ‘POIs’.
func (r *MultigeocodePlaceService) New(ctx context.Context, params MultigeocodePlaceNewParams, opts ...option.RequestOption) (res *MultigeocodePlaceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "multigeocode/place"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Use this method to get the details of previously created custom places using its
// NextBillion ID.
func (r *MultigeocodePlaceService) Get(ctx context.Context, docID string, query MultigeocodePlaceGetParams, opts ...option.RequestOption) (res *MultigeocodePlaceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if docID == "" {
		err = errors.New("missing required docId parameter")
		return
	}
	path := fmt.Sprintf("multigeocode/place/%s", docID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The "Update Place" method allows businesses to update the attributes of an
// existing place.
//
// This method allows you to update the attributes of custom places. In effect,
// updating a place replaces the current information in search results with the
// updated information associated with the specific docID. Use this method to
// enhance the accuracy/usability of your search results with respect to the
// default dataset to suit your business needs.
//
// If you want to prioritize a particular result in your search results, update the
// ‘score’ of that specific place.
// Alternatively, you can block places which are no longer needed by setting their
// status: ‘disable’.
func (r *MultigeocodePlaceService) Update(ctx context.Context, docID string, params MultigeocodePlaceUpdateParams, opts ...option.RequestOption) (res *MultigeocodePlaceUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if docID == "" {
		err = errors.New("missing required docId parameter")
		return
	}
	path := fmt.Sprintf("multigeocode/place/%s", docID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// The "Delete Place" method enables businesses to delete a previously created
// place
//
// Use this method to delete a previously created place. Please note that the place
// associated with the specified docID only would be deleted. As a result, once a
// place is deleted, the search API can still return valid results from the default
// datasets or others, if present.
func (r *MultigeocodePlaceService) Delete(ctx context.Context, docID string, body MultigeocodePlaceDeleteParams, opts ...option.RequestOption) (res *MultigeocodePlaceDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if docID == "" {
		err = errors.New("missing required docId parameter")
		return
	}
	path := fmt.Sprintf("multigeocode/place/%s", docID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type PlaceItem struct {
	// This parameter represents the complete address of the place, including the
	// street, city, state, postal code and country.
	Address string `json:"address"`
	// This parameter represents additional building information if applicable.
	Building string `json:"building"`
	// This parameter represents the city or town of the place.
	City string `json:"city"`
	// This parameter represents the country of the place.
	Country string `json:"country"`
	// This parameter represents the district of the place.
	District string `json:"district"`
	// This parameter represents the geographical coordinates of the place. It includes
	// the latitude and longitude values.
	Geopoint PlaceItemGeopoint `json:"geopoint"`
	// This parameter represents the house or building number of the place.
	House string `json:"house"`
	// This parameter represents a point of interest within the place. A Point of
	// Interest (POI) refers to a specific location or area that is of interest to
	// individuals for various reasons. It could be a landmark, tourist attraction,
	// business, or any other location that people might find important or intriguing.
	Poi PlaceItemPoi `json:"poi"`
	// This parameter represents the postal code or ZIP code of the place.
	PostalCode string `json:"postalCode"`
	// This parameter represents the state or region of the place.
	State string `json:"state"`
	// This parameter represents the street name of the place.
	Street string `json:"street"`
	// This parameter represents the sub-district or locality of the place.
	SubDistrict string `json:"subDistrict"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Building    respjson.Field
		City        respjson.Field
		Country     respjson.Field
		District    respjson.Field
		Geopoint    respjson.Field
		House       respjson.Field
		Poi         respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		SubDistrict respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaceItem) RawJSON() string { return r.JSON.raw }
func (r *PlaceItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PlaceItem to a PlaceItemParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PlaceItemParam.Overrides()
func (r PlaceItem) ToParam() PlaceItemParam {
	return param.Override[PlaceItemParam](json.RawMessage(r.RawJSON()))
}

// This parameter represents the geographical coordinates of the place. It includes
// the latitude and longitude values.
type PlaceItemGeopoint struct {
	// This parameter represents the latitude value of the place.
	Lat float64 `json:"lat"`
	// This parameter represents the longitude value of the place.
	Lng float64 `json:"lng"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lng         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaceItemGeopoint) RawJSON() string { return r.JSON.raw }
func (r *PlaceItemGeopoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This parameter represents a point of interest within the place. A Point of
// Interest (POI) refers to a specific location or area that is of interest to
// individuals for various reasons. It could be a landmark, tourist attraction,
// business, or any other location that people might find important or intriguing.
type PlaceItemPoi struct {
	// A title that describes the point of interest.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaceItemPoi) RawJSON() string { return r.JSON.raw }
func (r *PlaceItemPoi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaceItemParam struct {
	// This parameter represents the complete address of the place, including the
	// street, city, state, postal code and country.
	Address param.Opt[string] `json:"address,omitzero"`
	// This parameter represents additional building information if applicable.
	Building param.Opt[string] `json:"building,omitzero"`
	// This parameter represents the city or town of the place.
	City param.Opt[string] `json:"city,omitzero"`
	// This parameter represents the country of the place.
	Country param.Opt[string] `json:"country,omitzero"`
	// This parameter represents the district of the place.
	District param.Opt[string] `json:"district,omitzero"`
	// This parameter represents the house or building number of the place.
	House param.Opt[string] `json:"house,omitzero"`
	// This parameter represents the postal code or ZIP code of the place.
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// This parameter represents the state or region of the place.
	State param.Opt[string] `json:"state,omitzero"`
	// This parameter represents the street name of the place.
	Street param.Opt[string] `json:"street,omitzero"`
	// This parameter represents the sub-district or locality of the place.
	SubDistrict param.Opt[string] `json:"subDistrict,omitzero"`
	// This parameter represents the geographical coordinates of the place. It includes
	// the latitude and longitude values.
	Geopoint PlaceItemGeopointParam `json:"geopoint,omitzero"`
	// This parameter represents a point of interest within the place. A Point of
	// Interest (POI) refers to a specific location or area that is of interest to
	// individuals for various reasons. It could be a landmark, tourist attraction,
	// business, or any other location that people might find important or intriguing.
	Poi PlaceItemPoiParam `json:"poi,omitzero"`
	paramObj
}

func (r PlaceItemParam) MarshalJSON() (data []byte, err error) {
	type shadow PlaceItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaceItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This parameter represents the geographical coordinates of the place. It includes
// the latitude and longitude values.
type PlaceItemGeopointParam struct {
	// This parameter represents the latitude value of the place.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// This parameter represents the longitude value of the place.
	Lng param.Opt[float64] `json:"lng,omitzero"`
	paramObj
}

func (r PlaceItemGeopointParam) MarshalJSON() (data []byte, err error) {
	type shadow PlaceItemGeopointParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaceItemGeopointParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This parameter represents a point of interest within the place. A Point of
// Interest (POI) refers to a specific location or area that is of interest to
// individuals for various reasons. It could be a landmark, tourist attraction,
// business, or any other location that people might find important or intriguing.
type PlaceItemPoiParam struct {
	// A title that describes the point of interest.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r PlaceItemPoiParam) MarshalJSON() (data []byte, err error) {
	type shadow PlaceItemPoiParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaceItemPoiParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MultigeocodePlaceNewResponse struct {
	// A unique NextBillion DocID will be created for the POI. Use this ID to search
	// this place through the “Get Place” method, to update attributes or ‘status’
	// through the “Update Place” method or delete it using the “Delete Place” method.
	DocID string `json:"docId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultigeocodePlaceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MultigeocodePlaceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MultigeocodePlaceGetResponse struct {
	// It displays the information about the current source and current status of the
	// place. Use the “Update Place” method to change these values, as needed.
	DataSorce MultigeocodePlaceGetResponseDataSorce `json:"dataSorce"`
	// The unique NextBillion ID for the result item.
	DocID string `json:"docId"`
	// This parameter represents the place details, including geographical information,
	// address and other related information.
	Place []PlaceItem `json:"place"`
	// It returns the system calculated weighted score of the place. It depends on how
	// ‘richly’ the place was defined at the time of creation. In order to modify the
	// score, use “Update Place” method and update information for parameters which are
	// not set currently. As an alternative, you can directly update the score to a
	// custom value.
	Score int64 `json:"score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSorce   respjson.Field
		DocID       respjson.Field
		Place       respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultigeocodePlaceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MultigeocodePlaceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// It displays the information about the current source and current status of the
// place. Use the “Update Place” method to change these values, as needed.
type MultigeocodePlaceGetResponseDataSorce struct {
	// This parameter represents the unique reference ID associated with the data
	// source.
	RefID string `json:"refId"`
	// This parameter represents the current dataset source of the information returned
	// in the result.
	Source string `json:"source"`
	// This parameter indicates if a place is currently discoverable by search API or
	// not.
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
func (r MultigeocodePlaceGetResponseDataSorce) RawJSON() string { return r.JSON.raw }
func (r *MultigeocodePlaceGetResponseDataSorce) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MultigeocodePlaceUpdateResponse struct {
	// This could be “Ok” representing success or “not found” representing error in
	// processing the request.
	Msg string `json:"msg"`
	// Represents the status of the response.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultigeocodePlaceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MultigeocodePlaceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MultigeocodePlaceDeleteResponse struct {
	// This could be “Ok” representing success or “not found” representing error in
	// processing the request.
	Msg string `json:"msg"`
	// Represents the status of the response.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultigeocodePlaceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MultigeocodePlaceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MultigeocodePlaceNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// This parameter represents the place details, including geographical information,
	// address and other related information.
	Place []MultigeocodePlaceNewParamsPlace `json:"place,omitzero,required"`
	// When 2 places are located within 100 meters of each other and have more than 90%
	// of matching attributes (at least 11 out of 12 attributes in the “place” object),
	// they will be considered duplicates and any requests to add such a new place
	// would be rejected. Set force=true to override this duplicate check. You can use
	// this to create closely located POIs. For instance, places inside a mall,
	// university or a government building etc.
	Force param.Opt[bool] `json:"force,omitzero"`
	// Search score of the place. This is calculated based on how ‘richly’ the place is
	// defined. For instance, a place with - street name, city, state and country
	// attributes set might be ranked lower than a place which has values of - house,
	// building, street name, city, state and country attributes set. The score
	// determines the rank of the place among search results. You can also use this
	// field to set a custom score as per its relevance to rank it among the search
	// results from multiple data sources.
	Score param.Opt[int64] `json:"score,omitzero"`
	// It contains information about the dataset that returns the specific result
	DataSource MultigeocodePlaceNewParamsDataSource `json:"dataSource,omitzero"`
	paramObj
}

func (r MultigeocodePlaceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MultigeocodePlaceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultigeocodePlaceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MultigeocodePlaceNewParams]'s query parameters as
// `url.Values`.
func (r MultigeocodePlaceNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property Geopoint is required.
type MultigeocodePlaceNewParamsPlace struct {
	// This parameter represents the geographical coordinates of the place. It includes
	// the latitude and longitude values.
	Geopoint MultigeocodePlaceNewParamsPlaceGeopoint `json:"geopoint,omitzero,required"`
	// This parameter represents the complete address of the place, including the
	// street, city, state, postal code and country.
	Address param.Opt[string] `json:"address,omitzero"`
	// This parameter represents additional building information if applicable.
	Building param.Opt[string] `json:"building,omitzero"`
	// This parameter represents the city or town of the place.
	City param.Opt[string] `json:"city,omitzero"`
	// Country of the search context provided as comma-separated
	// [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) country
	// codes.
	//
	// Note: Country codes should be provided in uppercase.
	Country param.Opt[string] `json:"country,omitzero"`
	// This parameter represents the district of the place.
	District param.Opt[string] `json:"district,omitzero"`
	// This parameter represents the house or building number of the place.
	House param.Opt[string] `json:"house,omitzero"`
	// This parameter represents the postal code or ZIP code of the place.
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// This parameter represents the state or region of the place.
	State param.Opt[string] `json:"state,omitzero"`
	// This parameter represents the street name of the place.
	Street param.Opt[string] `json:"street,omitzero"`
	// This parameter represents the sub-district or locality of the place.
	SubDistrict param.Opt[string] `json:"subDistrict,omitzero"`
	// This parameter represents a point of interest within the place. A Point of
	// Interest (POI) refers to a specific location or area that is of interest to
	// individuals for various reasons. It could be a landmark, tourist attraction,
	// business, or any other location that people might find important or intriguing.
	Poi MultigeocodePlaceNewParamsPlacePoi `json:"poi,omitzero"`
	paramObj
}

func (r MultigeocodePlaceNewParamsPlace) MarshalJSON() (data []byte, err error) {
	type shadow MultigeocodePlaceNewParamsPlace
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultigeocodePlaceNewParamsPlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This parameter represents the geographical coordinates of the place. It includes
// the latitude and longitude values.
type MultigeocodePlaceNewParamsPlaceGeopoint struct {
	// This parameter represents the latitude value of the place.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// This parameter represents the longitude value of the place.
	Lng param.Opt[float64] `json:"lng,omitzero"`
	paramObj
}

func (r MultigeocodePlaceNewParamsPlaceGeopoint) MarshalJSON() (data []byte, err error) {
	type shadow MultigeocodePlaceNewParamsPlaceGeopoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultigeocodePlaceNewParamsPlaceGeopoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This parameter represents a point of interest within the place. A Point of
// Interest (POI) refers to a specific location or area that is of interest to
// individuals for various reasons. It could be a landmark, tourist attraction,
// business, or any other location that people might find important or intriguing.
type MultigeocodePlaceNewParamsPlacePoi struct {
	// A title that describes the point of interest.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r MultigeocodePlaceNewParamsPlacePoi) MarshalJSON() (data []byte, err error) {
	type shadow MultigeocodePlaceNewParamsPlacePoi
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultigeocodePlaceNewParamsPlacePoi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// It contains information about the dataset that returns the specific result
type MultigeocodePlaceNewParamsDataSource struct {
	// This parameter represents the unique reference ID associated with the data
	// source.
	RefID param.Opt[string] `json:"refId,omitzero"`
	// This parameter represents the source of the data.
	Source param.Opt[string] `json:"source,omitzero"`
	// This parameter indicates if a place is searchable.
	//
	// Any of "enable", "disable".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r MultigeocodePlaceNewParamsDataSource) MarshalJSON() (data []byte, err error) {
	type shadow MultigeocodePlaceNewParamsDataSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultigeocodePlaceNewParamsDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MultigeocodePlaceNewParamsDataSource](
		"status", "enable", "disable",
	)
}

type MultigeocodePlaceGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [MultigeocodePlaceGetParams]'s query parameters as
// `url.Values`.
func (r MultigeocodePlaceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MultigeocodePlaceUpdateParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Search score of the place. This is calculated based on how ‘richly’ the place is
	// defined. For instance, a place with street name, city, state and country
	// attributes set might be ranked lower than a place which has values of house,
	// building, street name, city, state and country attributes set. The score
	// determines the rank of the place among search results. You can also use this
	// field to set a custom score as per its relevance to rank it among the search
	// results from multiple data sources.
	Score param.Opt[int64] `json:"score,omitzero"`
	// dataSource values can be updated to enhance or prioritize the search results to
	// better suit specific business use cases.
	DataSource MultigeocodePlaceUpdateParamsDataSource `json:"dataSource,omitzero"`
	// This parameter represents the place details, including geographical information,
	// address and other related information.
	Place []PlaceItemParam `json:"place,omitzero"`
	paramObj
}

func (r MultigeocodePlaceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MultigeocodePlaceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultigeocodePlaceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MultigeocodePlaceUpdateParams]'s query parameters as
// `url.Values`.
func (r MultigeocodePlaceUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// dataSource values can be updated to enhance or prioritize the search results to
// better suit specific business use cases.
type MultigeocodePlaceUpdateParamsDataSource struct {
	// This parameter represents the unique reference ID associated with the data
	// source.
	RefID param.Opt[string] `json:"refId,omitzero"`
	//  1. Move the place to a new dataset by setting the value to a unique dataset
	//     name. You can also move the place to an existing dataset by using an existing
	//     dataset name other than the current one. In both cases, the current
	//     datasource will be replaced for the specified docID.
	//
	//  2. Update the place in an existing dataset by setting the name to the current
	//     value.
	Source param.Opt[string] `json:"source,omitzero"`
	// Set this to either enable or disable to allow the place to be retrieved by a
	// search API or block it respectively.
	//
	// Any of "enable", "disable".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r MultigeocodePlaceUpdateParamsDataSource) MarshalJSON() (data []byte, err error) {
	type shadow MultigeocodePlaceUpdateParamsDataSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultigeocodePlaceUpdateParamsDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MultigeocodePlaceUpdateParamsDataSource](
		"status", "enable", "disable",
	)
}

type MultigeocodePlaceDeleteParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [MultigeocodePlaceDeleteParams]'s query parameters as
// `url.Values`.
func (r MultigeocodePlaceDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
