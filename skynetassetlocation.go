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

// SkynetAssetLocationService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetAssetLocationService] method instead.
type SkynetAssetLocationService struct {
	Options []option.RequestOption
}

// NewSkynetAssetLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSkynetAssetLocationService(opts ...option.RequestOption) (r SkynetAssetLocationService) {
	r = SkynetAssetLocationService{}
	r.Options = opts
	return
}

// Track locations of an asset
func (r *SkynetAssetLocationService) List(ctx context.Context, id string, query SkynetAssetLocationListParams, opts ...option.RequestOption) (res *SkynetAssetLocationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/asset/%s/location/list", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Track the last location of an asset
func (r *SkynetAssetLocationService) GetLast(ctx context.Context, id string, query SkynetAssetLocationGetLastParams, opts ...option.RequestOption) (res *SkynetAssetLocationGetLastResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/asset/%s/location/last", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// An object with details of the tracked location. Please note that if there are no
// tracking records for an asset, no location data will be returned.
type TrackLocation struct {
	// If available, this property returns the accuracy of the GPS information received
	// at the tracked location. It is represented as an estimated horizontal accuracy
	// radius, in meters, at the 68th percentile confidence level.
	Accuracy float64 `json:"accuracy"`
	// If available in the GPS information, this property returns the altitude of the
	// asset at the tracked location. It is represented as height, in meters, above the
	// WGS84 reference ellipsoid.
	Altitude float64 `json:"altitude"`
	// Returns the battery level of the GPS device, as a percentage, when the location
	// was tracked. It has a minimum value of 0 and a maximum value of 100.
	BatteryLevel int64 `json:"battery_level"`
	// If available in the GPS information, this property returns the heading of the
	// asset calculated from true north in clockwise direction at the tracked location.
	// Please note that the bearing is not affected by the device orientation.
	//
	// The bearing will always be in the range of [0, 360).
	Bearing float64 `json:"bearing"`
	// An object with the coordinates of the last tracked location.
	Location TrackLocationLocation `json:"location"`
	// Specifies the custom data about the location that was added when the location
	// was uploaded.
	MetaData any `json:"meta_data"`
	// If available in the GPS information, this property returns the speed of the
	// asset, in meters per second, at the tracked location.
	Speed float64 `json:"speed"`
	// A UNIX epoch timestamp in milliseconds, representing the time at which the
	// location was tracked.
	Timestamp int64 `json:"timestamp"`
	// Internal parameter for tracking mode.
	TrackingMode string `json:"tracking_mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accuracy     respjson.Field
		Altitude     respjson.Field
		BatteryLevel respjson.Field
		Bearing      respjson.Field
		Location     respjson.Field
		MetaData     respjson.Field
		Speed        respjson.Field
		Timestamp    respjson.Field
		TrackingMode respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackLocation) RawJSON() string { return r.JSON.raw }
func (r *TrackLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with the coordinates of the last tracked location.
type TrackLocationLocation struct {
	// Latitude of the tracked location of the asset.
	Lat float64 `json:"lat"`
	// Longitude of the tracked location of the asset.
	Lon float64 `json:"lon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lon         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackLocationLocation) RawJSON() string { return r.JSON.raw }
func (r *TrackLocationLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetLocationListResponse struct {
	Data SkynetAssetLocationListResponseData `json:"data"`
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
func (r SkynetAssetLocationListResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetLocationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetLocationListResponseData struct {
	// Distance of the path, in meters, formed by connecting all tracked locations
	// returned.
	//
	// Please note that distance is returned only when the mapmatch property of
	// correction parameter is set to 1.
	Distance float64 `json:"distance"`
	// An object with geoJSON details of the route. It is returned only when the
	// mapmatch property of the correction parameter is set to 1 and geometry_type is
	// geojson, otherwise it is not present in the response. The contents of this
	// object follow the
	// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Geojson SkynetAssetLocationListResponseDataGeojson `json:"geojson"`
	// Geometry of tracked locations in the requested format. It is returned only if
	// the mapmatch property of the ‘correction’ parameter is set to 1.
	Geometry []string `json:"geometry"`
	// An array of objects with details of the tracked locations of the asset. Each
	// object represents one tracked location.
	List []TrackLocation `json:"list"`
	// An object with pagination details of the search results. Use this object to
	// implement pagination in your application.
	Page Pagination `json:"page"`
	// An array of objects with details about the snapped points for each of the
	// tracked locations returned for the asset.
	//
	// Please note that this property is returned only when the mapmatch property of
	// correction parameter is set to 1.
	SnappedPoints []SkynetAssetLocationListResponseDataSnappedPoint `json:"snapped_points"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance      respjson.Field
		Geojson       respjson.Field
		Geometry      respjson.Field
		List          respjson.Field
		Page          respjson.Field
		SnappedPoints respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetLocationListResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetLocationListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with geoJSON details of the route. It is returned only when the
// mapmatch property of the correction parameter is set to 1 and geometry_type is
// geojson, otherwise it is not present in the response. The contents of this
// object follow the
// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
type SkynetAssetLocationListResponseDataGeojson struct {
	// An object with details of the geoJSON geometry of the route.
	Geometry SkynetAssetLocationListResponseDataGeojsonGeometry `json:"geometry"`
	// Type of the geoJSON object.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geometry    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetLocationListResponseDataGeojson) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetLocationListResponseDataGeojson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the geoJSON geometry of the route.
type SkynetAssetLocationListResponseDataGeojsonGeometry struct {
	// An array of coordinates in the [longitude, latitude] format, representing the
	// route geometry.
	Coordinates []float64 `json:"coordinates"`
	// Type of the geoJSON geometry.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetLocationListResponseDataGeojsonGeometry) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetLocationListResponseDataGeojsonGeometry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetLocationListResponseDataSnappedPoint struct {
	// The bearing angle of the snapped point from the original tracked location, in
	// radians. It indicates the direction of the snapped point.
	Bearing string `json:"bearing"`
	// The distance of the snapped point from the original tracked location, in meters.
	Distance float64 `json:"distance"`
	// The latitude and longitude coordinates of the snapped point.
	Location SkynetAssetLocationListResponseDataSnappedPointLocation `json:"location"`
	// The name of the street or road of the snapped point.
	Name string `json:"name"`
	// The index of the tracked location to which this snapped point corresponds to.
	OriginalIndex string `json:"originalIndex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bearing       respjson.Field
		Distance      respjson.Field
		Location      respjson.Field
		Name          respjson.Field
		OriginalIndex respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetLocationListResponseDataSnappedPoint) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetLocationListResponseDataSnappedPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The latitude and longitude coordinates of the snapped point.
type SkynetAssetLocationListResponseDataSnappedPointLocation struct {
	// Latitude of the snapped point.
	Lat float64 `json:"lat"`
	// Longitude of the snapped point.
	Lon float64 `json:"lon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lon         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetLocationListResponseDataSnappedPointLocation) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetLocationListResponseDataSnappedPointLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetLocationGetLastResponse struct {
	// An object containing the information about the last tracked location of the
	// requested asset.
	Data SkynetAssetLocationGetLastResponseData `json:"data"`
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
func (r SkynetAssetLocationGetLastResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetLocationGetLastResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the information about the last tracked location of the
// requested asset.
type SkynetAssetLocationGetLastResponseData struct {
	// An object with details of the tracked location. Please note that if there are no
	// tracking records for an asset, no location data will be returned.
	Location TrackLocation `json:"location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Location    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetLocationGetLastResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetLocationGetLastResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetLocationListParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Describe the geometry characteristics through a , separated list of properties.
	//
	// Setting mapmatch to 1 returns the geometry of the tracked points, snapped to the
	// nearest road.
	//
	// Setting interpolate to 1 smoothens the snapped geometry by adding more points,
	// as needed. Please note, mapmatch should be set to 1 for interpolate to be
	// effective.
	//
	// mode is used to set the transport mode for which the snapped route will be
	// determined. Allowed values for mode are car and truck.
	Correction param.Opt[string] `query:"correction,omitzero" format:"mapmatch=boolean,interpolate=boolean,mode=string" json:"-"`
	// Time until which the tracked locations of the asset need to be retrieved.
	EndTime param.Opt[int64] `query:"end_time,omitzero" json:"-"`
	// Denotes page number. Use this along with the ps parameter to implement
	// pagination for your searched results. This parameter does not have a maximum
	// limit but would return an empty response in case a higher value is provided when
	// the result-set itself is smaller.
	Pn param.Opt[int64] `query:"pn,omitzero" json:"-"`
	// Denotes number of search results per page. Use this along with the pn parameter
	// to implement pagination for your searched results.
	Ps param.Opt[int64] `query:"ps,omitzero" json:"-"`
	// Time after which the tracked locations of the asset need to be retrieved.
	StartTime param.Opt[int64] `query:"start_time,omitzero" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetAssetLocationListParamsCluster `query:"cluster,omitzero" json:"-"`
	// Set the geometry format to encode the path linking the tracked locations of the
	// asset.
	//
	// Please note that geometry_type is effective only when mapmatch property of
	// correction parameter is set to 1.
	//
	// Any of "polyline", "polyline6", "geojson".
	GeometryType SkynetAssetLocationListParamsGeometryType `query:"geometry_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetAssetLocationListParams]'s query parameters as
// `url.Values`.
func (r SkynetAssetLocationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetAssetLocationListParamsCluster string

const (
	SkynetAssetLocationListParamsClusterAmerica SkynetAssetLocationListParamsCluster = "america"
)

// Set the geometry format to encode the path linking the tracked locations of the
// asset.
//
// Please note that geometry_type is effective only when mapmatch property of
// correction parameter is set to 1.
type SkynetAssetLocationListParamsGeometryType string

const (
	SkynetAssetLocationListParamsGeometryTypePolyline  SkynetAssetLocationListParamsGeometryType = "polyline"
	SkynetAssetLocationListParamsGeometryTypePolyline6 SkynetAssetLocationListParamsGeometryType = "polyline6"
	SkynetAssetLocationListParamsGeometryTypeGeojson   SkynetAssetLocationListParamsGeometryType = "geojson"
)

type SkynetAssetLocationGetLastParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetAssetLocationGetLastParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetAssetLocationGetLastParams]'s query parameters as
// `url.Values`.
func (r SkynetAssetLocationGetLastParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetAssetLocationGetLastParamsCluster string

const (
	SkynetAssetLocationGetLastParamsClusterAmerica SkynetAssetLocationGetLastParamsCluster = "america"
)
