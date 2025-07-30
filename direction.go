// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai

import (
	"context"
	"net/http"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// DirectionService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirectionService] method instead.
type DirectionService struct {
	Options []option.RequestOption
}

// NewDirectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDirectionService(opts ...option.RequestOption) (r DirectionService) {
	r = DirectionService{}
	r.Options = opts
	return
}

// Directions API is a service that computes a route with given coordinates.
func (r *DirectionService) ComputeRoute(ctx context.Context, body DirectionComputeRouteParams, opts ...option.RequestOption) (res *DirectionComputeRouteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "directions/json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DirectionComputeRouteResponse struct {
	// Displays the error message in case of a failed request or operation. Please note
	// that this parameter is not returned in the response in case of a successful
	// request.
	Msg string `json:"msg"`
	// An object containing details about the returned route. Will contain multiple
	// objects if more than one routes are present in the response.
	Route DirectionComputeRouteResponseRoute `json:"route"`
	// A string indicating the state of the response. On normal responses, the value
	// will be `Ok`. Indicative HTTP error codes are returned for different errors. See
	// the [API Errors Codes](#api-error-codes) section below for more information.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		Route       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponse) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing details about the returned route. Will contain multiple
// objects if more than one routes are present in the response.
type DirectionComputeRouteResponseRoute struct {
	// The distance, in meters, for the complete trip.
	Distance float64 `json:"distance"`
	// The duration, in seconds, of the complete trip.
	Duration float64 `json:"duration"`
	// Location coordinates of the point where the route ends. It is the same as the
	// `destination` in the input request. Returned only when `steps` is true in the
	// input request.
	EndLocation DirectionComputeRouteResponseRouteEndLocation `json:"end_location"`
	// An object with geoJSON details of the route. This object is returned when the
	// `geometry` field is set to `geojson` in the input request, otherwise it is not
	// present in the response. The contents of this object follow the
	// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Geojson DirectionComputeRouteResponseRouteGeojson `json:"geojson"`
	// Encoded geometry of the returned route in the selected format and specified
	// `overview` verbosity. This parameter is configured in the input request.
	Geometry string `json:"geometry"`
	// An array of objects returning the details about each `leg` of the route.
	// `waypoints` split the route into legs.
	Legs []DirectionComputeRouteResponseRouteLeg `json:"legs"`
	// Location coordinates of the point where the route starts. It is the same as the
	// `origin` in the input request. Returned only when `steps` is true in the input
	// request.
	StartLocation DirectionComputeRouteResponseRouteStartLocation `json:"start_location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance      respjson.Field
		Duration      respjson.Field
		EndLocation   respjson.Field
		Geojson       respjson.Field
		Geometry      respjson.Field
		Legs          respjson.Field
		StartLocation respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRoute) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the route ends. It is the same as the
// `destination` in the input request. Returned only when `steps` is true in the
// input request.
type DirectionComputeRouteResponseRouteEndLocation struct {
	// latitude of the `start_location`.
	Latitude float64 `json:"latitude"`
	// longitude of the `start_location`.
	Longitude float64 `json:"longitude"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteEndLocation) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteEndLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with geoJSON details of the route. This object is returned when the
// `geometry` field is set to `geojson` in the input request, otherwise it is not
// present in the response. The contents of this object follow the
// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
type DirectionComputeRouteResponseRouteGeojson struct {
	// An object with details of the geoJSON geometry of the route.
	Geometry DirectionComputeRouteResponseRouteGeojsonGeometry `json:"geometry"`
	// Property associated with the geoJSON shape.
	Properties string `json:"properties"`
	// Type of the geoJSON object.
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
func (r DirectionComputeRouteResponseRouteGeojson) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteGeojson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the geoJSON geometry of the route.
type DirectionComputeRouteResponseRouteGeojsonGeometry struct {
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
func (r DirectionComputeRouteResponseRouteGeojsonGeometry) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteGeojsonGeometry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectionComputeRouteResponseRouteLeg struct {
	// An object containing leg distance value, in meters.
	Distance DirectionComputeRouteResponseRouteLegDistance `json:"distance"`
	// An object containing leg duration value, in seconds.
	Duration DirectionComputeRouteResponseRouteLegDuration `json:"duration"`
	// Location coordinates of the point where the leg ends. Returned only when `steps`
	// is true in the input request.
	EndLocation DirectionComputeRouteResponseRouteLegEndLocation `json:"end_location"`
	// Location coordinates of the point where the leg starts. Returned only when
	// `steps` is true in the input request.
	StartLocation DirectionComputeRouteResponseRouteLegStartLocation `json:"start_location"`
	// An array of objects with details of each step of the `legs`. Returned only when
	// `steps` is `true` in the input request. An empty array is returned when `steps`
	// is `false` in the input request.
	Steps []DirectionComputeRouteResponseRouteLegStep `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance      respjson.Field
		Duration      respjson.Field
		EndLocation   respjson.Field
		StartLocation respjson.Field
		Steps         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLeg) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing leg distance value, in meters.
type DirectionComputeRouteResponseRouteLegDistance struct {
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegDistance) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegDistance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing leg duration value, in seconds.
type DirectionComputeRouteResponseRouteLegDuration struct {
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegDuration) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the leg ends. Returned only when `steps`
// is true in the input request.
type DirectionComputeRouteResponseRouteLegEndLocation struct {
	// Latitude of the `end_location` of the `leg`.
	Latitude float64 `json:"latitude"`
	// Longitude of the `end_location` of the `leg`.
	Longitude float64 `json:"longitude"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegEndLocation) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegEndLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the leg starts. Returned only when
// `steps` is true in the input request.
type DirectionComputeRouteResponseRouteLegStartLocation struct {
	// Latitude of the `start_location` of the `leg`.
	Latitude float64 `json:"latitude"`
	// Longitude of the `start_location` of the `leg`.
	Longitude float64 `json:"longitude"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegStartLocation) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegStartLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectionComputeRouteResponseRouteLegStep struct {
	// An object containing step distance value, in meters.
	Distance DirectionComputeRouteResponseRouteLegStepDistance `json:"distance"`
	// An object containing step duration value, in seconds.
	Duration DirectionComputeRouteResponseRouteLegStepDuration `json:"duration"`
	// Location coordinates of the point where the `step` ends.
	EndLocation DirectionComputeRouteResponseRouteLegStepEndLocation `json:"end_location"`
	// An object with geoJSON details of the `step`.This object is returned when the
	// `geometry` field is set to `geojson` in the input request, otherwise it is not
	// present in the response. The contents of this object follow the
	// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Geojson DirectionComputeRouteResponseRouteLegStepGeojson `json:"geojson"`
	// Encoded geometry of the `step` in the selected format.
	Geometry string `json:"geometry"`
	// An object with maneuver details for the `step`.
	Maneuver DirectionComputeRouteResponseRouteLegStepManeuver `json:"maneuver"`
	// Location coordinates of the point where the `step` starts.
	StartLocation DirectionComputeRouteResponseRouteLegStepStartLocation `json:"start_location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance      respjson.Field
		Duration      respjson.Field
		EndLocation   respjson.Field
		Geojson       respjson.Field
		Geometry      respjson.Field
		Maneuver      respjson.Field
		StartLocation respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegStep) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing step distance value, in meters.
type DirectionComputeRouteResponseRouteLegStepDistance struct {
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegStepDistance) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegStepDistance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing step duration value, in seconds.
type DirectionComputeRouteResponseRouteLegStepDuration struct {
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegStepDuration) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegStepDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the `step` ends.
type DirectionComputeRouteResponseRouteLegStepEndLocation struct {
	// Latitude of the `end_location` of the `step`.
	Latitude float64 `json:"latitude"`
	// Longitude of the `end_location` of the `step`.
	Longitude float64 `json:"longitude"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegStepEndLocation) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegStepEndLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with geoJSON details of the `step`.This object is returned when the
// `geometry` field is set to `geojson` in the input request, otherwise it is not
// present in the response. The contents of this object follow the
// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
type DirectionComputeRouteResponseRouteLegStepGeojson struct {
	// An object with details of the geoJSON geometry of the `step`.
	Geometry DirectionComputeRouteResponseRouteLegStepGeojsonGeometry `json:"geometry"`
	// Property associated with the geoJSON shape.
	Properties string `json:"properties"`
	// Type of the geoJSON object.
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
func (r DirectionComputeRouteResponseRouteLegStepGeojson) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegStepGeojson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the geoJSON geometry of the `step`.
type DirectionComputeRouteResponseRouteLegStepGeojsonGeometry struct {
	// An array of coordinates in the [longitude, latitude] format, representing the
	// `step` geometry.
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
func (r DirectionComputeRouteResponseRouteLegStepGeojsonGeometry) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegStepGeojsonGeometry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with maneuver details for the `step`.
type DirectionComputeRouteResponseRouteLegStepManeuver struct {
	// The clockwise angle from true north to the direction of travel immediately after
	// the `maneuver`. Range of values is between 0-359.
	BearingAfter int64 `json:"bearing_after"`
	// The clockwise angle from true north to the direction of travel immediately
	// before the `maneuver`. Range of values is between 0-359.
	BearingBefore int64 `json:"bearing_before"`
	// A coordinate pair describing the location of the `maneuver`.
	Coordinate DirectionComputeRouteResponseRouteLegStepManeuverCoordinate `json:"coordinate"`
	// A string indicating the type of `maneuver`.
	ManeuverType string `json:"maneuver_type"`
	// Modifier associated with `maneuver_type`.
	Modifier string `json:"modifier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BearingAfter  respjson.Field
		BearingBefore respjson.Field
		Coordinate    respjson.Field
		ManeuverType  respjson.Field
		Modifier      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegStepManeuver) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegStepManeuver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A coordinate pair describing the location of the `maneuver`.
type DirectionComputeRouteResponseRouteLegStepManeuverCoordinate struct {
	// Latitude of the `maneuver` location.
	Latitude float64 `json:"latitude"`
	// Longitude of the `maneuver` location.
	Longitude float64 `json:"longitude"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegStepManeuverCoordinate) RawJSON() string {
	return r.JSON.raw
}
func (r *DirectionComputeRouteResponseRouteLegStepManeuverCoordinate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the `step` starts.
type DirectionComputeRouteResponseRouteLegStepStartLocation struct {
	// Latitude of the `start_location` of the `step`.
	Latitude float64 `json:"latitude"`
	// Longitude of the `start_location` of the `step`.
	Longitude float64 `json:"longitude"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteLegStepStartLocation) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteLegStepStartLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the route starts. It is the same as the
// `origin` in the input request. Returned only when `steps` is true in the input
// request.
type DirectionComputeRouteResponseRouteStartLocation struct {
	// Latitude of the `start_location`.
	Latitude float64 `json:"latitude"`
	// Longitude of the `start_location`.
	Longitude float64 `json:"longitude"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectionComputeRouteResponseRouteStartLocation) RawJSON() string { return r.JSON.raw }
func (r *DirectionComputeRouteResponseRouteStartLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectionComputeRouteParams struct {
	Destination string `json:"destination,required"`
	Origin      string `json:"origin,required"`
	// Sets the number of alternative routes to return. It is effective only when
	// alternatives=true. Default to 3.
	//
	// Please note that adding alternative route count does not guarantee matching
	// number of routes to be returned if potential alternative routes do not exist.
	Altcount param.Opt[int64] `json:"altcount,omitzero"`
	// When true the API will return alternate routes.
	//
	// The alternatives is effective only when there are no waypoints included in the
	// request.
	//
	// You can set the number of alternate routes to be returned in the altcount
	// property.
	Alternatives param.Opt[bool] `json:"alternatives,omitzero"`
	// A semicolon-separated list indicating the side of the road from which to
	// approach waypoints in a requested route.
	//
	// When set to unrestricted a route can arrive at the waypoint from either side of
	// the road and when set to curb the route will arrive at the waypoint on the
	// driving side of the region.
	//
	// Please note the number of values provided must be one more than the number of
	// waypoints. The last value of approaches will determine the approach for the
	// destination. However, you can skip a coordinate and show its position in the
	// list with the ; separator.
	Approaches param.Opt[string] `json:"approaches,omitzero"`
	// Limits the search to road segments with given bearing, in degrees, towards true
	// north in clockwise direction. Each bearings should be in the format of
	// degree,range, where the degree should be a value between \[0, 360\] and range
	// should be a value between \[0, 180\].
	//
	// Please note that the number of bearings should be two more than the number of
	// waypoints. This is to account for the bearing of origin and destination. If a
	// route can approach a waypoint or the destination from any direction, the bearing
	// for that point can be specified as "0,180".
	Bearings param.Opt[string] `json:"bearings,omitzero"`
	// Requires option=flexible.
	//
	// Specify if crossing an international border is expected for operations near
	// border areas. When set to false, the API will prohibit routes crossing the
	// borders. When set to true, the service will return routes which cross the
	// borders between countries, if required for the given set destination and
	// waypoints.
	//
	// This feature is available in North America region only. Please get in touch with
	// [support@nextbillion.ai](mailto:support@nextbillion.ai) to enquire/enable other
	// areas.
	CrossBorder param.Opt[bool] `json:"cross_border,omitzero"`
	// Requires option=flexible.
	//
	// Use this parameter to set a departure time, expressed as UNIX epoch timestamp in
	// seconds, for calculating the isochrone contour. The response will consider the
	// typical traffic conditions at the given time and return a contour which can be
	// reached under those traffic conditions.
	//
	// Please note that if no input is provided for this parameter then the traffic
	// conditions at the time of making the request are considered.
	DepartureTime param.Opt[int64] `json:"departure_time,omitzero"`
	// Requires option=flexible.
	//
	// An array of durations, in seconds, for which the driver can drive continuously
	// before taking a rest. Multiple drive time limits can be separated by a comma
	// character ",". After driving for the given duration the driver will take a rest
	// for a fixed period, specified in rest_times . Once the rest duration is over,
	// the subsequent driving duration starts and the process continues until all drive
	// times and rest periods are exhausted or if the driver reaches the destination.
	// This feature is useful in complying with Hours of Service regulations and
	// calculates actual ETAs with regulated driving periods.
	//
	// As an example, a drive_time_limits=\[500, 300\] means that driver can drive for
	// 500 seconds before the first rest period and then drive for another 300 seconds
	// before taking a rest next time.
	//
	// \- If the trip duration is smaller than the first input of drive_time_limits,
	// then there will be no rest actions scheduled by the service.
	// \- If the trip duration is larger than the scheduled time, then a "warning" is
	// returned in the response - along with details of last leg of the trip - to
	// indicate the same.
	DriveTimeLimits param.Opt[string] `json:"drive_time_limits,omitzero"`
	// Requires option=flexible.
	//
	// An array of durations, in seconds, for which the driver should rest after
	// completing the corresponding continuous driving interval (provided in
	// drive_time_limits). Multiple rest times can be separated by a comma character
	// ",". Ideally, the number of rest_times provided should be equal to the number of
	// drive_time_limits provided for proper scheduling of driver breaks.
	//
	// As an example, a rest_times=\[500, 300\] means that driver can rest for 500
	// seconds after the first continuous driving session and rest for 300 seconds
	// after the next continuous driving session.
	//
	// \- If the number of rest_times provided are less than the number of
	// drive_time_limits, the service will schedule a rest period of "0" seconds after
	// each such drive time period which does not have a corresponding entry in
	// rest_times.
	// \- If the number of rest_times provided is more than the number of drive times
	// provided, the additional rest times are never applied.
	RestTimes param.Opt[string] `json:"rest_times,omitzero"`
	// Set this to true to receive additional details about the routes and each of its
	// legs (details of geometry, start & end locations) in the response.
	Steps param.Opt[bool] `json:"steps,omitzero"`
	// Requires option=flexible.
	//
	// Specify the total load per axle (including the weight of trailers and shipped
	// goods) of the truck, in tonnes. When used, the service will return routes which
	// are legally allowed to carry the load specified per axle.
	//
	// Please note this parameter is effective only when mode=truck.
	TruckAxleLoad param.Opt[float64] `json:"truck_axle_load,omitzero"`
	// Requires option=flexible.
	//
	// This defines the dimensions of a truck in centimeters (CM). This parameter is
	// effective only when the mode=truck. Maximum dimensions are as follows:
	//
	// \- Height = 1000 cm
	// \- Width = 5000 cm
	// \- Length = 5000 cm
	TruckSize param.Opt[string] `json:"truck_size,omitzero" format:"height,width,length"`
	// Requires option=flexible.
	//
	// This parameter defines the weight of the truck including trailers and shipped
	// goods in kilograms (KG). This parameter is effective only when mode=truck.
	TruckWeight param.Opt[int64] `json:"truck_weight,omitzero"`
	// Requires option=flexible.
	//
	// Specify the turn angles that can be taken safely by the vehicle. The permissible
	// turn angles are calculated as \[0 + turn_angle_range , 360 - turn_angle_range\].
	// Please note that this parameter is effective only when avoid=sharp_turn.
	//
	// It is worth highlighting here that providing smaller angles might lead to 4xx
	// errors as route engine might not be able find routes satisfying the smaller turn
	// angle criteria for all turns in the route.
	TurnAngleRange param.Opt[int64] `json:"turn_angle_range,omitzero"`
	// Pipe-separated list of coordinate pairs
	Waypoints param.Opt[string] `json:"waypoints,omitzero"`
	// When option=fast (by default):
	//
	// Setting this will ensure the route avoids ferries, tolls, highways or nothing.
	// Multiple values should be separated by a pipe "|". If none is provided along
	// with other values, an error is returned as a valid route is not feasible.
	//
	// Please note that when this parameter is not provided in the input, ferries are
	// set to be avoided by default. When this parameter is provided, only the
	// mentioned objects are avoided.
	//
	// When option=flexible:
	//
	// Set this parameter to find alternative routes that bypass specified objects. Use
	// a pipe "|" to separate multiple values. This is a flexible filter; if no
	// alternative routes exist, the service will still provide a route that includes
	// the objects. For a strict filter, consider using the exclude parameter.
	//
	// \- This parameter is effective only when route_type=fastest.
	// \- Following objects are exceptions to the flexible filtering behavior of avoid
	// parameter: bbox, tunnel and geofence_id. When used, the service will return an
	// error in case there are no alternative routes available.
	// \- When using avoid=bbox users also need to specify the boundaries of the
	// bounding box to be avoided. Multiple bounding boxes can be specified
	// simultaneously. The perimeter of a bounding box can not exceed 500 KM. Format:
	// bbox=min_latitude,min_longtitude,max_latitude,max_longitude. Example:
	// avoid=bbox: 34.0635,-118.2547, 34.0679,-118.2478 | bbox: 34.0521,-118.2342,
	// 34.0478,-118.2437
	// \- When using avoid=sharp_turn, default range of permissible turn angles is
	// \[120,240\] in the clockwise direction from the current road. In order to
	// override default range, please use turn_angle_range parameter.
	// \- When using avoid=geofence_id, only the the geofences created using
	// [NextBillion.ai](http://NextBillion.ai) Geofence API are valid.
	// \- When this parameter is not provided in the input, ferry routes are set to be
	// avoided by default. When this parameter is provided, only the mentioned
	// object(s) are avoided.
	// \- If none is provided along with other values, an error is returned as a valid
	// route is not feasible.
	//
	// Any of "toll", "ferry", "highway", "none", "sharp_turn", "uturn",
	// "service_road", "left_turn", "right_turn", "bbox", "geofence_id", "tunnel".
	Avoid DirectionComputeRouteParamsAvoid `json:"avoid,omitzero"`
	// Requires option=flexible.
	//
	// Specify the emission class to which the vehicle (engine) belongs to. The service
	// will use this setting to generate routes that are permissible for that engine
	// class. Only the emission classifications in the EU regions are supported
	// currently. Please reach out to
	// [support@nextbillion.ai](mailto:support@nextbillion.ai) to enable for your
	// region.
	//
	// Any of "euro0", "euro1", "euro2", "euro3", "euro4", "euro5", "euro6", "euro7",
	// "euro8", "euro9".
	EmissionClass DirectionComputeRouteParamsEmissionClass `json:"emission_class,omitzero"`
	// Requires option=flexible.
	//
	// This parameter serves as a mandatory filter, ensuring the service returns only
	// those routes that strictly avoid the object(s) indicated. Multiple values should
	// be separated by a pipe `|`). If no routes can be found that exclude the
	// specified object(s), the service will return an error. For a less strict
	// filtering approach, consider using the avoid parameter.
	//
	// \- This parameter is effective only when route_type=fastest.
	// \- When using exclude=sharp_turn, default range of permissible turn angles is
	// \[120,240\]. In order to override default range, please use turn_angle_range
	// parameter.
	// \- If none is provided along with other values, an error is returned as a valid
	// route is not feasible.
	//
	// Any of "toll", "ferry", "highway", "service_road", "uturn", "sharp_turn",
	// "left_turn", "right_turn", "none".
	Exclude DirectionComputeRouteParamsExclude `json:"exclude,omitzero"`
	// Sets the output format of the route geometry in the response.
	//
	// On providing polyline and polyline6 as input, respective encoded geometry is
	// returned. However, when geojson is provided as the input value, polyline encoded
	// geometry is returned in the response along with the geojson details of the
	// route.
	//
	// Any of "polyline", "polyline6", "geojson".
	Geometry DirectionComputeRouteParamsGeometry `json:"geometry,omitzero"`
	// Requires option=flexible.
	//
	// Specify the type of hazardous material being carried and the service will avoid
	// roads which are not suitable for the type of goods specified. Multiple values
	// can be separated using a pipe operator "|".
	//
	// Please note that this parameter is effective only when mode=truck.
	//
	// Any of "general", "circumstantial", "explosive", "harmful_to_water".
	HazmatType DirectionComputeRouteParamsHazmatType `json:"hazmat_type,omitzero"`
	// Set the driving mode the service should use to determine a route. In "car" mode,
	// the API will return a route that a car can take. Using "truck" mode will return
	// a route a truck can use, taking into account appropriate truck routing
	// restrictions.
	//
	// When mode=truck, following are the default dimensions that are used:
	//
	// \- truck_height = 214 centimeters
	// \- truck_width = 183 centimeters
	// \- truck_length = 519 centimeters
	// \- truck_weight = 5000 kg
	//
	// When option=flexible, you can use custom truck dimensions with truck_weight and
	// truck_size parameters.
	//
	// Note: Only the car profile is enabled by default. Please note that customized
	// profiles (including truck) might not be available for all regions. Please
	// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
	// representative or reach out at
	// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
	// additional profiles.
	//
	// Any of "car", "truck".
	Mode DirectionComputeRouteParamsMode `json:"mode,omitzero"`
	// The option parameter specifies the version of the directions service to use.
	// Setting option=flexible activates the Flexible API, which supports advanced
	// features like truck routing, time-based routing, route type selection
	// (fastest/shortest), and segment-wise speed limits. If not set, the API defaults
	// to the Fast version for real-time routing.
	//
	// Any of "fast", "flexible".
	Option DirectionComputeRouteParamsOption `json:"option,omitzero"`
	// Specify the verbosity of route geometry.
	//
	// When set to full, the most detailed geometry available is returned. When set to
	// simplified, a simplified version of the full geometry is returned. No overview
	// geometry is returned when set to false.
	//
	// Any of "full", "simplified", "false".
	Overview DirectionComputeRouteParamsOverview `json:"overview,omitzero"`
	// Requires option=flexible.
	//
	// Use this parameter to receive additional information about the road segments
	// returned in the response. Currently, following inputs are supported:
	//
	// \- max_speed : segment-wise maximum speed information of roads in the
	// response.
	// \- toll_distance : returns the total distance travelled on the road segments
	// having tolls.
	// \- toll_cost: returns the range of toll charges, in local currency, that can be
	// incurred for the suggested route.
	//
	// Any of "max_speed", "toll_distance", "toll_cost".
	RoadInfo DirectionComputeRouteParamsRoadInfo `json:"road_info,omitzero"`
	// Requires option=flexible.
	//
	// Set the route type that needs to be returned.
	//
	// Any of "fastest", "shortest".
	RouteType DirectionComputeRouteParamsRouteType `json:"route_type,omitzero"`
	paramObj
}

func (r DirectionComputeRouteParams) MarshalJSON() (data []byte, err error) {
	type shadow DirectionComputeRouteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectionComputeRouteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When option=fast (by default):
//
// Setting this will ensure the route avoids ferries, tolls, highways or nothing.
// Multiple values should be separated by a pipe "|". If none is provided along
// with other values, an error is returned as a valid route is not feasible.
//
// Please note that when this parameter is not provided in the input, ferries are
// set to be avoided by default. When this parameter is provided, only the
// mentioned objects are avoided.
//
// When option=flexible:
//
// Set this parameter to find alternative routes that bypass specified objects. Use
// a pipe "|" to separate multiple values. This is a flexible filter; if no
// alternative routes exist, the service will still provide a route that includes
// the objects. For a strict filter, consider using the exclude parameter.
//
// \- This parameter is effective only when route_type=fastest.
// \- Following objects are exceptions to the flexible filtering behavior of avoid
// parameter: bbox, tunnel and geofence_id. When used, the service will return an
// error in case there are no alternative routes available.
// \- When using avoid=bbox users also need to specify the boundaries of the
// bounding box to be avoided. Multiple bounding boxes can be specified
// simultaneously. The perimeter of a bounding box can not exceed 500 KM. Format:
// bbox=min_latitude,min_longtitude,max_latitude,max_longitude. Example:
// avoid=bbox: 34.0635,-118.2547, 34.0679,-118.2478 | bbox: 34.0521,-118.2342,
// 34.0478,-118.2437
// \- When using avoid=sharp_turn, default range of permissible turn angles is
// \[120,240\] in the clockwise direction from the current road. In order to
// override default range, please use turn_angle_range parameter.
// \- When using avoid=geofence_id, only the the geofences created using
// [NextBillion.ai](http://NextBillion.ai) Geofence API are valid.
// \- When this parameter is not provided in the input, ferry routes are set to be
// avoided by default. When this parameter is provided, only the mentioned
// object(s) are avoided.
// \- If none is provided along with other values, an error is returned as a valid
// route is not feasible.
type DirectionComputeRouteParamsAvoid string

const (
	DirectionComputeRouteParamsAvoidToll        DirectionComputeRouteParamsAvoid = "toll"
	DirectionComputeRouteParamsAvoidFerry       DirectionComputeRouteParamsAvoid = "ferry"
	DirectionComputeRouteParamsAvoidHighway     DirectionComputeRouteParamsAvoid = "highway"
	DirectionComputeRouteParamsAvoidNone        DirectionComputeRouteParamsAvoid = "none"
	DirectionComputeRouteParamsAvoidSharpTurn   DirectionComputeRouteParamsAvoid = "sharp_turn"
	DirectionComputeRouteParamsAvoidUturn       DirectionComputeRouteParamsAvoid = "uturn"
	DirectionComputeRouteParamsAvoidServiceRoad DirectionComputeRouteParamsAvoid = "service_road"
	DirectionComputeRouteParamsAvoidLeftTurn    DirectionComputeRouteParamsAvoid = "left_turn"
	DirectionComputeRouteParamsAvoidRightTurn   DirectionComputeRouteParamsAvoid = "right_turn"
	DirectionComputeRouteParamsAvoidBbox        DirectionComputeRouteParamsAvoid = "bbox"
	DirectionComputeRouteParamsAvoidGeofenceID  DirectionComputeRouteParamsAvoid = "geofence_id"
	DirectionComputeRouteParamsAvoidTunnel      DirectionComputeRouteParamsAvoid = "tunnel"
)

// Requires option=flexible.
//
// Specify the emission class to which the vehicle (engine) belongs to. The service
// will use this setting to generate routes that are permissible for that engine
// class. Only the emission classifications in the EU regions are supported
// currently. Please reach out to
// [support@nextbillion.ai](mailto:support@nextbillion.ai) to enable for your
// region.
type DirectionComputeRouteParamsEmissionClass string

const (
	DirectionComputeRouteParamsEmissionClassEuro0 DirectionComputeRouteParamsEmissionClass = "euro0"
	DirectionComputeRouteParamsEmissionClassEuro1 DirectionComputeRouteParamsEmissionClass = "euro1"
	DirectionComputeRouteParamsEmissionClassEuro2 DirectionComputeRouteParamsEmissionClass = "euro2"
	DirectionComputeRouteParamsEmissionClassEuro3 DirectionComputeRouteParamsEmissionClass = "euro3"
	DirectionComputeRouteParamsEmissionClassEuro4 DirectionComputeRouteParamsEmissionClass = "euro4"
	DirectionComputeRouteParamsEmissionClassEuro5 DirectionComputeRouteParamsEmissionClass = "euro5"
	DirectionComputeRouteParamsEmissionClassEuro6 DirectionComputeRouteParamsEmissionClass = "euro6"
	DirectionComputeRouteParamsEmissionClassEuro7 DirectionComputeRouteParamsEmissionClass = "euro7"
	DirectionComputeRouteParamsEmissionClassEuro8 DirectionComputeRouteParamsEmissionClass = "euro8"
	DirectionComputeRouteParamsEmissionClassEuro9 DirectionComputeRouteParamsEmissionClass = "euro9"
)

// Requires option=flexible.
//
// This parameter serves as a mandatory filter, ensuring the service returns only
// those routes that strictly avoid the object(s) indicated. Multiple values should
// be separated by a pipe `|`). If no routes can be found that exclude the
// specified object(s), the service will return an error. For a less strict
// filtering approach, consider using the avoid parameter.
//
// \- This parameter is effective only when route_type=fastest.
// \- When using exclude=sharp_turn, default range of permissible turn angles is
// \[120,240\]. In order to override default range, please use turn_angle_range
// parameter.
// \- If none is provided along with other values, an error is returned as a valid
// route is not feasible.
type DirectionComputeRouteParamsExclude string

const (
	DirectionComputeRouteParamsExcludeToll        DirectionComputeRouteParamsExclude = "toll"
	DirectionComputeRouteParamsExcludeFerry       DirectionComputeRouteParamsExclude = "ferry"
	DirectionComputeRouteParamsExcludeHighway     DirectionComputeRouteParamsExclude = "highway"
	DirectionComputeRouteParamsExcludeServiceRoad DirectionComputeRouteParamsExclude = "service_road"
	DirectionComputeRouteParamsExcludeUturn       DirectionComputeRouteParamsExclude = "uturn"
	DirectionComputeRouteParamsExcludeSharpTurn   DirectionComputeRouteParamsExclude = "sharp_turn"
	DirectionComputeRouteParamsExcludeLeftTurn    DirectionComputeRouteParamsExclude = "left_turn"
	DirectionComputeRouteParamsExcludeRightTurn   DirectionComputeRouteParamsExclude = "right_turn"
	DirectionComputeRouteParamsExcludeNone        DirectionComputeRouteParamsExclude = "none"
)

// Sets the output format of the route geometry in the response.
//
// On providing polyline and polyline6 as input, respective encoded geometry is
// returned. However, when geojson is provided as the input value, polyline encoded
// geometry is returned in the response along with the geojson details of the
// route.
type DirectionComputeRouteParamsGeometry string

const (
	DirectionComputeRouteParamsGeometryPolyline  DirectionComputeRouteParamsGeometry = "polyline"
	DirectionComputeRouteParamsGeometryPolyline6 DirectionComputeRouteParamsGeometry = "polyline6"
	DirectionComputeRouteParamsGeometryGeojson   DirectionComputeRouteParamsGeometry = "geojson"
)

// Requires option=flexible.
//
// Specify the type of hazardous material being carried and the service will avoid
// roads which are not suitable for the type of goods specified. Multiple values
// can be separated using a pipe operator "|".
//
// Please note that this parameter is effective only when mode=truck.
type DirectionComputeRouteParamsHazmatType string

const (
	DirectionComputeRouteParamsHazmatTypeGeneral        DirectionComputeRouteParamsHazmatType = "general"
	DirectionComputeRouteParamsHazmatTypeCircumstantial DirectionComputeRouteParamsHazmatType = "circumstantial"
	DirectionComputeRouteParamsHazmatTypeExplosive      DirectionComputeRouteParamsHazmatType = "explosive"
	DirectionComputeRouteParamsHazmatTypeHarmfulToWater DirectionComputeRouteParamsHazmatType = "harmful_to_water"
)

// Set the driving mode the service should use to determine a route. In "car" mode,
// the API will return a route that a car can take. Using "truck" mode will return
// a route a truck can use, taking into account appropriate truck routing
// restrictions.
//
// When mode=truck, following are the default dimensions that are used:
//
// \- truck_height = 214 centimeters
// \- truck_width = 183 centimeters
// \- truck_length = 519 centimeters
// \- truck_weight = 5000 kg
//
// When option=flexible, you can use custom truck dimensions with truck_weight and
// truck_size parameters.
//
// Note: Only the car profile is enabled by default. Please note that customized
// profiles (including truck) might not be available for all regions. Please
// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
// representative or reach out at
// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
// additional profiles.
type DirectionComputeRouteParamsMode string

const (
	DirectionComputeRouteParamsModeCar   DirectionComputeRouteParamsMode = "car"
	DirectionComputeRouteParamsModeTruck DirectionComputeRouteParamsMode = "truck"
)

// The option parameter specifies the version of the directions service to use.
// Setting option=flexible activates the Flexible API, which supports advanced
// features like truck routing, time-based routing, route type selection
// (fastest/shortest), and segment-wise speed limits. If not set, the API defaults
// to the Fast version for real-time routing.
type DirectionComputeRouteParamsOption string

const (
	DirectionComputeRouteParamsOptionFast     DirectionComputeRouteParamsOption = "fast"
	DirectionComputeRouteParamsOptionFlexible DirectionComputeRouteParamsOption = "flexible"
)

// Specify the verbosity of route geometry.
//
// When set to full, the most detailed geometry available is returned. When set to
// simplified, a simplified version of the full geometry is returned. No overview
// geometry is returned when set to false.
type DirectionComputeRouteParamsOverview string

const (
	DirectionComputeRouteParamsOverviewFull       DirectionComputeRouteParamsOverview = "full"
	DirectionComputeRouteParamsOverviewSimplified DirectionComputeRouteParamsOverview = "simplified"
	DirectionComputeRouteParamsOverviewFalse      DirectionComputeRouteParamsOverview = "false"
)

// Requires option=flexible.
//
// Use this parameter to receive additional information about the road segments
// returned in the response. Currently, following inputs are supported:
//
// \- max_speed : segment-wise maximum speed information of roads in the
// response.
// \- toll_distance : returns the total distance travelled on the road segments
// having tolls.
// \- toll_cost: returns the range of toll charges, in local currency, that can be
// incurred for the suggested route.
type DirectionComputeRouteParamsRoadInfo string

const (
	DirectionComputeRouteParamsRoadInfoMaxSpeed     DirectionComputeRouteParamsRoadInfo = "max_speed"
	DirectionComputeRouteParamsRoadInfoTollDistance DirectionComputeRouteParamsRoadInfo = "toll_distance"
	DirectionComputeRouteParamsRoadInfoTollCost     DirectionComputeRouteParamsRoadInfo = "toll_cost"
)

// Requires option=flexible.
//
// Set the route type that needs to be returned.
type DirectionComputeRouteParamsRouteType string

const (
	DirectionComputeRouteParamsRouteTypeFastest  DirectionComputeRouteParamsRouteType = "fastest"
	DirectionComputeRouteParamsRouteTypeShortest DirectionComputeRouteParamsRouteType = "shortest"
)
