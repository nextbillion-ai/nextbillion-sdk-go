// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai

import (
	"context"
	"errors"
	"fmt"
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

// SkynetAssetService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetAssetService] method instead.
type SkynetAssetService struct {
	Options  []option.RequestOption
	Event    SkynetAssetEventService
	Location SkynetAssetLocationService
}

// NewSkynetAssetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSkynetAssetService(opts ...option.RequestOption) (r SkynetAssetService) {
	r = SkynetAssetService{}
	r.Options = opts
	r.Event = NewSkynetAssetEventService(opts...)
	r.Location = NewSkynetAssetLocationService(opts...)
	return
}

// Create an Asset
func (r *SkynetAssetService) New(ctx context.Context, params SkynetAssetNewParams, opts ...option.RequestOption) (res *SkynetAssetNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skynet/asset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get an Asset
func (r *SkynetAssetService) Get(ctx context.Context, id string, query SkynetAssetGetParams, opts ...option.RequestOption) (res *SkynetAssetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/asset/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update an Asset
func (r *SkynetAssetService) Update(ctx context.Context, id string, params SkynetAssetUpdateParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/asset/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Get Asset List
func (r *SkynetAssetService) List(ctx context.Context, query SkynetAssetListParams, opts ...option.RequestOption) (res *SkynetAssetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skynet/asset/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an Asset
func (r *SkynetAssetService) Delete(ctx context.Context, id string, body SkynetAssetDeleteParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/asset/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Bind asset to device
func (r *SkynetAssetService) Bind(ctx context.Context, id string, params SkynetAssetBindParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/asset/%s/bind", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Upload track info
func (r *SkynetAssetService) Track(ctx context.Context, id string, params SkynetAssetTrackParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/asset/%s/track", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update asset attributes. (add)
func (r *SkynetAssetService) UpdateAttributes(ctx context.Context, id string, params SkynetAssetUpdateAttributesParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/asset/%s/attributes", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type MetaData = any

type SimpleResp struct {
	// Displays the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Message string `json:"message"`
	// A string indicating the state of the response. On successful responses, the
	// value will be Ok. Indicative error messages are returned for different errors.
	// See the [API Error Codes](#api-error-codes) section below for more information.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimpleResp) RawJSON() string { return r.JSON.raw }
func (r *SimpleResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetNewResponse struct {
	// An object containing the ID of the asset created.
	Data SkynetAssetNewResponseData `json:"data"`
	// Displays the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Message string `json:"message"`
	// A string indicating the state of the response. On successful responses, the
	// value will be Ok. Indicative error messages are returned for different errors.
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
func (r SkynetAssetNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the ID of the asset created.
type SkynetAssetNewResponseData struct {
	// Unique ID of the asset created. It will be the same as custom_id, if provided.
	// Else it will be an auto generated UUID. Please note this ID cannot be updated.
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetGetResponse struct {
	// An object containing the information about the asset returned.
	Data SkynetAssetGetResponseData `json:"data"`
	// Displays the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Message string `json:"message"`
	// A string indicating the state of the response. On successful responses, the
	// value will be Ok. Indicative error messages are returned for different errors.
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
func (r SkynetAssetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the information about the asset returned.
type SkynetAssetGetResponseData struct {
	// An object with details of the asset properties.
	Asset AssetDetails `json:"asset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetListResponse struct {
	// A data object containing the list of assets.
	Data SkynetAssetListResponseData `json:"data"`
	// Displays the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Message string `json:"message"`
	// A string indicating the state of the response. On successful responses, the
	// value will be Ok. Indicative error messages are returned for different errors.
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
func (r SkynetAssetListResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A data object containing the list of assets.
type SkynetAssetListResponseData struct {
	// An array of objects, with each object representing one asset.
	List []AssetDetails `json:"list"`
	// An object with pagination details of the search results. Use this object to
	// implement pagination in your application.
	Page Pagination `json:"page"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		List        respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetListResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// Set a unique ID for the new asset. If not provided, an ID will be automatically
	// generated in UUID format. A valid custom*id can contain letters, numbers, "-", &
	// "*" only.
	//
	// Please note that the ID of an asset can not be changed once it is created.
	CustomID param.Opt[string] `json:"custom_id,omitzero"`
	// Description for the asset.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the asset. Use this field to assign a meaningful, custom name to the
	// asset being created.
	Name param.Opt[string] `json:"name,omitzero"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetAssetNewParamsCluster `query:"cluster,omitzero" json:"-"`
	// attributes can be used to store custom information about an asset in key:value
	// format. Use attributes to add any useful information or context to your assets
	// like the vehicle type, shift timing etc. Moreover, these attributes can be used
	// to filter assets in **Search**, **Monitor**, and _Get Asset List_ queries.
	//
	// Please note that the maximum number of key:value pairs that can be added to an
	// attributes object is 100. Also, the overall size of attributes object should not
	// exceed 65kb.
	Attributes any `json:"attributes,omitzero"`
	// Any valid json object data. Can be used to save customized data. Max size is
	// 65kb.
	MetaData MetaData `json:"meta_data,omitzero"`
	// **This parameter will be deprecated soon! Please use the attributes parameter to
	// add labels or markers for the asset.**
	//
	// Tags of the asset. tags can be used for filtering assets in operations like _Get
	// Asset List_ and asset **Search** methods. They can also be used for monitoring
	// of assets using the **Monitor** methods after linking tags and asset.
	//
	// Valid tags are strings consisting of alphanumeric characters (A-Z, a-z, 0-9)
	// along with the underscore ('\_') and hyphen ('-') symbols.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SkynetAssetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetAssetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetAssetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetAssetNewParams]'s query parameters as `url.Values`.
func (r SkynetAssetNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetAssetNewParamsCluster string

const (
	SkynetAssetNewParamsClusterAmerica SkynetAssetNewParamsCluster = "america"
)

type SkynetAssetGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetAssetGetParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetAssetGetParams]'s query parameters as `url.Values`.
func (r SkynetAssetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetAssetGetParamsCluster string

const (
	SkynetAssetGetParamsClusterAmerica SkynetAssetGetParamsCluster = "america"
)

type SkynetAssetUpdateParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// Use this param to update the description of an asset.
	Description param.Opt[string] `json:"description,omitzero"`
	// Use this param to update the name of an asset. Users can assign meaningful
	// custom names to their assets.
	Name param.Opt[string] `json:"name,omitzero"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetAssetUpdateParamsCluster `query:"cluster,omitzero" json:"-"`
	// Use this param to update the attributes of an asset in key:value format. Users
	// can maintain any useful information or context about the assets by utilising
	// this parameter.
	//
	// Please be careful when using this parameter while updating an asset as the new
	// attributes object provided will completely overwrite the old attributes object.
	// Use the _Update Asset Attributes_ method to add new or modify existing
	// attributes.
	//
	// Another point to note is that the overall size of the attributes object cannot
	// exceed 65kb and the maximum number of key:value pairs that can be added to this
	// object is 100.
	Attributes any `json:"attributes,omitzero"`
	// Any valid json object data. Can be used to save customized data. Max size is
	// 65kb.
	MetaData MetaData `json:"meta_data,omitzero"`
	// **This parameter will be deprecated soon! Please use the attributes parameter to
	// add labels or markers for the asset.**
	//
	// Use this param to update the tags of an asset. tags can be used to filter asset
	// in _Get Asset List_, **Search** and **Monitor** queries.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SkynetAssetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetAssetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetAssetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetAssetUpdateParams]'s query parameters as
// `url.Values`.
func (r SkynetAssetUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetAssetUpdateParamsCluster string

const (
	SkynetAssetUpdateParamsClusterAmerica SkynetAssetUpdateParamsCluster = "america"
)

type SkynetAssetListParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// Use this parameter to filter the assets by their attributes. Only the assets
	// having all the attributes added to this parameter, will be returned in the
	// response. Multiple attributes can be separated using pipes (|).
	//
	// Please note the attributes are case sensitive. Also, this parameter can not be
	// used in conjunction with include_any_of_attributes parameter.
	IncludeAllOfAttributes param.Opt[string] `query:"include_all_of_attributes,omitzero" format:"key_1:value_1|key_2:value_2" json:"-"`
	// Use this parameter to filter the assets by their attributes. Assets having at
	// least one of the attributes added to this parameter, will be returned in the
	// response. Multiple attributes can be separated using pipes (|).
	//
	// Please note the attributes are case sensitive. Also, this parameter can not be
	// used in conjunction with include_all_of_attributes parameter.
	IncludeAnyOfAttributes param.Opt[string] `query:"include_any_of_attributes,omitzero" format:"key1:value1|key2:value2|..." json:"-"`
	// Denotes page number. Use this along with the ps parameter to implement
	// pagination for your searched results. This parameter does not have a maximum
	// limit but would return an empty response in case a higher value is provided when
	// the result-set itself is smaller.
	Pn param.Opt[int64] `query:"pn,omitzero" json:"-"`
	// Denotes number of search results per page. Use this along with the pn parameter
	// to implement pagination for your searched results.
	Ps param.Opt[int64] `query:"ps,omitzero" json:"-"`
	// Provide a single field to sort the results by. Only updated_at or created_at
	// fields can be selected for ordering the results.
	//
	// By default, the result is sorted by created_at field in the descending order.
	// Allowed values for specifying the order are asc for ascending order and desc for
	// descending order.
	Sort param.Opt[string] `query:"sort,omitzero" format:"field:order" json:"-"`
	// **This parameter will be deprecated soon! Please use the
	// include_all_of_attributes or include_any_of_attributes parameters to provide
	// labels or markers for the assets to be retrieved.**
	//
	// tags can be used to filter the assets. Only those assets which have all the tags
	// provided, will be included in the result. In case multiple tags need to be
	// specified, use , to separate them.
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetAssetListParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetAssetListParams]'s query parameters as `url.Values`.
func (r SkynetAssetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetAssetListParamsCluster string

const (
	SkynetAssetListParamsClusterAmerica SkynetAssetListParamsCluster = "america"
)

type SkynetAssetDeleteParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetAssetDeleteParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetAssetDeleteParams]'s query parameters as
// `url.Values`.
func (r SkynetAssetDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetAssetDeleteParamsCluster string

const (
	SkynetAssetDeleteParamsClusterAmerica SkynetAssetDeleteParamsCluster = "america"
)

type SkynetAssetBindParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// Device ID to be linked to the asset identified by id.
	//
	// Please note that the device needs to be linked to an asset before using it in
	// the _Upload locations of an Asset_ method for sending GPS information about the
	// asset.
	DeviceID string `json:"device_id" api:"required"`
	paramObj
}

func (r SkynetAssetBindParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetAssetBindParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetAssetBindParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetAssetBindParams]'s query parameters as `url.Values`.
func (r SkynetAssetBindParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SkynetAssetTrackParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// ID of the device used to upload the tracking information of the asset.
	//
	// Please note that the device_id used here must already be linked to the asset.
	// Use the _Bind Device to Asset_ method to link a device with your asset.
	DeviceID string `json:"device_id" api:"required"`
	// An array of objects to collect the location tracking information for an asset.
	// Each object must correspond to details of only one location.
	Locations SkynetAssetTrackParamsLocations `json:"locations,omitzero" api:"required"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetAssetTrackParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

func (r SkynetAssetTrackParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetAssetTrackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetAssetTrackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetAssetTrackParams]'s query parameters as `url.Values`.
func (r SkynetAssetTrackParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// An array of objects to collect the location tracking information for an asset.
// Each object must correspond to details of only one location.
//
// The properties Location, Timestamp are required.
type SkynetAssetTrackParamsLocations struct {
	// An object to collect the coordinate details of the tracked location. Please note
	// this field is mandatory when uploading locations for an asset.
	Location SkynetAssetTrackParamsLocationsLocation `json:"location,omitzero" api:"required"`
	// Use this parameter to provide the time, expressed as UNIX epoch timestamp in
	// milliseconds, when the location was tracked. Please note this field is mandatory
	// when uploading locations for an asset.
	Timestamp int64 `json:"timestamp" api:"required"`
	// Use this parameter to provide the accuracy of the GPS information at the tracked
	// location. It is the estimated horizontal accuracy radius, in meters.
	Accuracy param.Opt[float64] `json:"accuracy,omitzero"`
	// Use this parameter to provide the altitude, in meters, of the asset at the
	// tracked location.
	Altitude param.Opt[float64] `json:"altitude,omitzero"`
	// Use this parameter to provide the battery level of the GPS device, as a
	// percentage, when the location is tracked. It should have a minimum value of 0
	// and a maximum value of 100.
	BatteryLevel param.Opt[int64] `json:"battery_level,omitzero"`
	// Use this parameter to provide the heading of the asset, in radians, calculated
	// from true north in clockwise direction. This should always be in the range of
	// [0, 360).
	Bearing param.Opt[float64] `json:"bearing,omitzero"`
	// Use this parameter to provide the speed of the asset, in meters per second, at
	// the tracked location.
	Speed param.Opt[float64] `json:"speed,omitzero"`
	// NB tracking mode.
	TrackingMode param.Opt[string] `json:"tracking_mode,omitzero"`
	// Use this object to add any custom data about the location that is being
	// uploaded. Recommended to use the key:value format for adding the desired
	// information.
	//
	// Please note that the maximum size of meta_data object should not exceed 65Kb.
	MetaData any `json:"meta_data,omitzero"`
	paramObj
}

func (r SkynetAssetTrackParamsLocations) MarshalJSON() (data []byte, err error) {
	type shadow SkynetAssetTrackParamsLocations
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetAssetTrackParamsLocations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object to collect the coordinate details of the tracked location. Please note
// this field is mandatory when uploading locations for an asset.
//
// The properties Lat, Lon are required.
type SkynetAssetTrackParamsLocationsLocation struct {
	// Latitude of the tracked location of the asset.
	Lat float64 `json:"lat" api:"required"`
	// Longitude of the tracked location of the asset.
	Lon float64 `json:"lon" api:"required"`
	paramObj
}

func (r SkynetAssetTrackParamsLocationsLocation) MarshalJSON() (data []byte, err error) {
	type shadow SkynetAssetTrackParamsLocationsLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetAssetTrackParamsLocationsLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// the cluster of the region you want to use
type SkynetAssetTrackParamsCluster string

const (
	SkynetAssetTrackParamsClusterAmerica SkynetAssetTrackParamsCluster = "america"
)

type SkynetAssetUpdateAttributesParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// attributes can be used to add any useful information or context to your assets
	// like the vehicle type, shift timing etc. These attributes can also be used to
	// filter assets in **Search**, **Monitor**, and _Get Asset List_ queries.
	//
	// Provide the attributes to be added or updated, in key:value format. If an
	// existing key is provided in the input, then the value will be modified as per
	// the input value. If a new key is provided in the input, then the key would be
	// added to the existing set. The contents of any value field are neither altered
	// nor removed unless specifically referred to by its key in the input request.
	//
	// Please note that the maximum number of key:value pairs that can be added to an
	// attributes object is 100. Also, the overall size of attributes object should not
	// exceed 65kb.
	Attributes any `json:"attributes,omitzero" api:"required"`
	paramObj
}

func (r SkynetAssetUpdateAttributesParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetAssetUpdateAttributesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetAssetUpdateAttributesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetAssetUpdateAttributesParams]'s query parameters as
// `url.Values`.
func (r SkynetAssetUpdateAttributesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
