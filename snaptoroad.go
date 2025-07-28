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

// SnapToRoadService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSnapToRoadService] method instead.
type SnapToRoadService struct {
	Options []option.RequestOption
}

// NewSnapToRoadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSnapToRoadService(opts ...option.RequestOption) (r SnapToRoadService) {
	r = SnapToRoadService{}
	r.Options = opts
	return
}

// Nextbillion.ai Snap To Roads API takes a series of locations along a route, and
// returns the new locations on this route that are snapped to the best-matched
// roads where the trip took place. You can set various parameters, such as
// timestamps or radius, to optimize the result.
func (r *SnapToRoadService) Snap(ctx context.Context, query SnapToRoadSnapParams, opts ...option.RequestOption) (res *SnapToRoadSnapResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "snapToRoads/json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response Body
type SnapToRoadSnapResponse struct {
	// The total distance of the snapped path in meters.
	Distance int64 `json:"distance"`
	// A GeoJSON object with details of the snapped path. This object is returned when
	// the `geometry` field is set to `geojson` in the input request, otherwise it is
	// not present in the response. The contents of this object follow the
	// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Geojson SnapToRoadSnapResponseGeojson `json:"geojson"`
	// An array of strings containing the encoded geometries of snapped paths in
	// `polyline` or `polyline6` format.
	Geometry []string `json:"geometry"`
	// Displays the error message in case of a failed request or operation. Please note
	// that this parameter is not returned in the response in case of a successful
	// request.
	Msg string `json:"msg"`
	// An object containing the maximum speed information for each road segment present
	// in the route.
	RoadInfo SnapToRoadSnapResponseRoadInfo `json:"road_info"`
	// An array of objects. Each object provides the details of a `path` coordinate
	// point snapped to the nearest road.
	SnappedPoints []SnapToRoadSnapResponseSnappedPoint `json:"snappedPoints"`
	// A string indicating the state of the response. On normal responses, the value
	// will be `Ok`. Indicative HTTP error codes are returned for different errors. See
	// the [API Errors Codes](#api-error-codes) section below for more information.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance      respjson.Field
		Geojson       respjson.Field
		Geometry      respjson.Field
		Msg           respjson.Field
		RoadInfo      respjson.Field
		SnappedPoints respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapToRoadSnapResponse) RawJSON() string { return r.JSON.raw }
func (r *SnapToRoadSnapResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A GeoJSON object with details of the snapped path. This object is returned when
// the `geometry` field is set to `geojson` in the input request, otherwise it is
// not present in the response. The contents of this object follow the
// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
type SnapToRoadSnapResponseGeojson struct {
	// An object with details of the geoJSON geometry of the snapped path.
	Geometry SnapToRoadSnapResponseGeojsonGeometry `json:"geometry"`
	// Properties associated with the geoJSON shape of the snapped path.
	Properties string `json:"properties"`
	// Type of the GeoJSON object.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geometry    respjson.Field
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapToRoadSnapResponseGeojson) RawJSON() string { return r.JSON.raw }
func (r *SnapToRoadSnapResponseGeojson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the geoJSON geometry of the snapped path.
type SnapToRoadSnapResponseGeojsonGeometry struct {
	// An array of coordinates in the [longitude, latitude] format, representing the
	// snapped path geometry.
	Coordinates []float64 `json:"coordinates"`
	// Type of the geoJSON geometry
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
func (r SnapToRoadSnapResponseGeojsonGeometry) RawJSON() string { return r.JSON.raw }
func (r *SnapToRoadSnapResponseGeojsonGeometry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the maximum speed information for each road segment present
// in the route.
type SnapToRoadSnapResponseRoadInfo struct {
	// An array of objects containing maximum speed, in kilometers per hour, for each
	// segment of the route. Each object represents one road segment.
	MaxSpeed []SnapToRoadSnapResponseRoadInfoMaxSpeed `json:"max_speed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxSpeed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapToRoadSnapResponseRoadInfo) RawJSON() string { return r.JSON.raw }
func (r *SnapToRoadSnapResponseRoadInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapToRoadSnapResponseRoadInfoMaxSpeed struct {
	// `length` refers to a sequence of 'n' consecutive vertices in the route geometry
	// starting from the `offset`, forming a continuous section of route where the
	// maximum speed is the same and is indicated in `value`.
	Length int64 `json:"length"`
	// `offset` is the index value of the vertex of route geometry, which is the
	// starting point of the segment.
	Offset int64 `json:"offset"`
	// `value` denotes the maximum speed of this segment, in kilometers per hour.
	//
	//   - A value of "-1" indicates that the speed is unlimited for this road segment.
	//   - A value of "0" indicates that there is no information about the maximum speed
	//     for this road segment.
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Length      respjson.Field
		Offset      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapToRoadSnapResponseRoadInfoMaxSpeed) RawJSON() string { return r.JSON.raw }
func (r *SnapToRoadSnapResponseRoadInfoMaxSpeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapToRoadSnapResponseSnappedPoint struct {
	// The bearing, calculated as the angle from true north in clockwise direction, of
	// the route leading to the next snapped point from the current `snapped_point`, in
	// radians. In case of the last `snapped_point` of the route, the bearing indicates
	// the direction of the route to the previous `snapped_location`.
	Bearing float64 `json:"bearing,required"`
	// The distance of the snapped point from the original input coordinate in meters.
	Distance float64 `json:"distance,required"`
	// The latitude and longitude coordinates of the snapped point.
	Location SnapToRoadSnapResponseSnappedPointLocation `json:"location,required"`
	// The name of the street or road that the input coordinate snapped to.
	Name string `json:"name,required"`
	// The index of the input `path` coordinate point to which this snapped point
	// corresponds to.
	OriginalIndex int64 `json:"originalIndex,required"`
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
func (r SnapToRoadSnapResponseSnappedPoint) RawJSON() string { return r.JSON.raw }
func (r *SnapToRoadSnapResponseSnappedPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The latitude and longitude coordinates of the snapped point.
type SnapToRoadSnapResponseSnappedPointLocation struct {
	// Latitude of the snapped point.
	Latitude float64 `json:"latitude,required"`
	// Longitude of the snapped point.
	Longitude float64 `json:"longitude,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapToRoadSnapResponseSnappedPointLocation) RawJSON() string { return r.JSON.raw }
func (r *SnapToRoadSnapResponseSnappedPointLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapToRoadSnapParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Pipe-separated list of coordinate points along a path which would be snapped to
	// a road.
	Path string `query:"path,required" format:"latitude_1,longitude_1|latitude_2,longitude_2|..." json:"-"`
	// Pipe separated radiuses, in meters (m), up to which a coordinate point can be
	// snapped. Please note, if no valid road is available within the specified radius,
	// the API would snap the points to nearest, most viable road. When using this
	// parameter, it is recommended to specify as many radius values as the number of
	// points in "path" parameter. If the same number of "radiuses" are not provided,
	// the API will use the default radius value of 25 meters for all locations.
	Radiuses param.Opt[string] `query:"radiuses,omitzero" format:"radius_1|radius_2|..." json:"-"`
	// Pipe-separated UNIX epoch timestamp in seconds for each location. If used, the
	// number of timestamps must be equal to the number of coordinate points in the
	// "path" parameter. The "timestamps" must increase monotonically starting from the
	// first timestamp. This means that each subsequent timestamp should either be more
	// than or equal to the preceding one.
	Timestamps param.Opt[string] `query:"timestamps,omitzero" format:"timestamps_1|timestamps_2|..." json:"-"`
	// Enable it to ignore locations outside the service boundary. When "true", the
	// service would ignore "path" coordinates points falling outside the accessible
	// area, which otherwise would cause an error when this parameter is "false".
	TolerateOutlier param.Opt[bool] `query:"tolerate_outlier,omitzero" json:"-"`
	// A semicolon-separated list indicating the side of the road from which to
	// approach the locations on the snapped route. When set to "unrestricted" a route
	// can arrive at the snapped location from either side of the road and when set to
	// "curb" the route will arrive at the snapped location on the driving side of the
	// region. Please note the number of values provided must be equal to the number of
	// coordinate points provided in the "path" parameter. However, you can skip a
	// coordinate and show its position in the list with the ";" separator.
	//
	// Any of "`unrestricted`", "`curb`".
	Approaches SnapToRoadSnapParamsApproaches `query:"approaches,omitzero" json:"-"`
	// Setting this will ensure the route avoids ferries, tolls, highways or nothing.
	// Multiple values should be separated by a pipe (|). If "none" is provided along
	// with other values, an error is returned as a valid route is not feasible. Please
	// note that when this parameter is not provided in the input, ferries are set to
	// be avoided by default. When this parameter is provided, only the mentioned
	// objects are avoided.
	//
	// Any of "toll", "ferry", "highway", "none".
	Avoid SnapToRoadSnapParamsAvoid `query:"avoid,omitzero" json:"-"`
	// Sets the output format of the route geometry in the response. Only the
	// "polyline" or "polyline6" encoded "geometry" of the snapped path is returned in
	// the response depending on the value provided in the input. When "geojson" is
	// selected as the input value, "polyline6" encoded geometry of the snapped path is
	// returned along with a "geojson" object.
	//
	// Any of "`polyline`", "`polyline6`", "`geojson`".
	Geometry SnapToRoadSnapParamsGeometry `query:"geometry,omitzero" json:"-"`
	// Set which driving mode the service should use to determine a route. For example,
	// if you use "car", the API will return a route that a car can take. Using "truck"
	// will return a route a truck can use, taking into account appropriate truck
	// routing restrictions.
	//
	// Note: Only the "car" profile is enabled by default. Please note that customized
	// profiles (including "truck") might not be available for all regions. Please
	// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
	// representative or reach out at
	// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
	// additional profiles.
	//
	// Any of "`car`", "`truck`".
	Mode SnapToRoadSnapParamsMode `query:"mode,omitzero" json:"-"`
	// Include this parameter in the request to return segment-wise speed information
	// of the route returned in the response.
	//
	// Please note that returning speed information is a function of "road_info"
	// parameter, which is effective only when "option=flexible". However, the
	// resultant route might not contain all the locations provided in "path" input.
	//
	// Any of "`flexible`".
	Option SnapToRoadSnapParamsOption `query:"option,omitzero" json:"-"`
	// Use this parameter to receive segment-wise maximum speed information of the
	// route in the response. "max_speed" is the only allowed value.
	//
	// Any of "`max_speed`".
	RoadInfo SnapToRoadSnapParamsRoadInfo `query:"road_info,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SnapToRoadSnapParams]'s query parameters as `url.Values`.
func (r SnapToRoadSnapParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A semicolon-separated list indicating the side of the road from which to
// approach the locations on the snapped route. When set to "unrestricted" a route
// can arrive at the snapped location from either side of the road and when set to
// "curb" the route will arrive at the snapped location on the driving side of the
// region. Please note the number of values provided must be equal to the number of
// coordinate points provided in the "path" parameter. However, you can skip a
// coordinate and show its position in the list with the ";" separator.
type SnapToRoadSnapParamsApproaches string

const (
	SnapToRoadSnapParamsApproachesUnrestricted SnapToRoadSnapParamsApproaches = "`unrestricted`"
	SnapToRoadSnapParamsApproachesCurb         SnapToRoadSnapParamsApproaches = "`curb`"
)

// Setting this will ensure the route avoids ferries, tolls, highways or nothing.
// Multiple values should be separated by a pipe (|). If "none" is provided along
// with other values, an error is returned as a valid route is not feasible. Please
// note that when this parameter is not provided in the input, ferries are set to
// be avoided by default. When this parameter is provided, only the mentioned
// objects are avoided.
type SnapToRoadSnapParamsAvoid string

const (
	SnapToRoadSnapParamsAvoidToll    SnapToRoadSnapParamsAvoid = "toll"
	SnapToRoadSnapParamsAvoidFerry   SnapToRoadSnapParamsAvoid = "ferry"
	SnapToRoadSnapParamsAvoidHighway SnapToRoadSnapParamsAvoid = "highway"
	SnapToRoadSnapParamsAvoidNone    SnapToRoadSnapParamsAvoid = "none"
)

// Sets the output format of the route geometry in the response. Only the
// "polyline" or "polyline6" encoded "geometry" of the snapped path is returned in
// the response depending on the value provided in the input. When "geojson" is
// selected as the input value, "polyline6" encoded geometry of the snapped path is
// returned along with a "geojson" object.
type SnapToRoadSnapParamsGeometry string

const (
	SnapToRoadSnapParamsGeometryPolyline  SnapToRoadSnapParamsGeometry = "`polyline`"
	SnapToRoadSnapParamsGeometryPolyline6 SnapToRoadSnapParamsGeometry = "`polyline6`"
	SnapToRoadSnapParamsGeometryGeojson   SnapToRoadSnapParamsGeometry = "`geojson`"
)

// Set which driving mode the service should use to determine a route. For example,
// if you use "car", the API will return a route that a car can take. Using "truck"
// will return a route a truck can use, taking into account appropriate truck
// routing restrictions.
//
// Note: Only the "car" profile is enabled by default. Please note that customized
// profiles (including "truck") might not be available for all regions. Please
// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
// representative or reach out at
// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
// additional profiles.
type SnapToRoadSnapParamsMode string

const (
	SnapToRoadSnapParamsModeCar   SnapToRoadSnapParamsMode = "`car`"
	SnapToRoadSnapParamsModeTruck SnapToRoadSnapParamsMode = "`truck`"
)

// Include this parameter in the request to return segment-wise speed information
// of the route returned in the response.
//
// Please note that returning speed information is a function of "road_info"
// parameter, which is effective only when "option=flexible". However, the
// resultant route might not contain all the locations provided in "path" input.
type SnapToRoadSnapParamsOption string

const (
	SnapToRoadSnapParamsOptionFlexible SnapToRoadSnapParamsOption = "`flexible`"
)

// Use this parameter to receive segment-wise maximum speed information of the
// route in the response. "max_speed" is the only allowed value.
type SnapToRoadSnapParamsRoadInfo string

const (
	SnapToRoadSnapParamsRoadInfoMaxSpeed SnapToRoadSnapParamsRoadInfo = "`max_speed`"
)
