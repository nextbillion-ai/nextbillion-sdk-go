// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apijson"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/nextbillion-sdk-go/option"
	"github.com/stainless-sdks/nextbillion-sdk-go/packages/param"
	"github.com/stainless-sdks/nextbillion-sdk-go/packages/respjson"
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
	// test
	Destination string `json:"destination,required"`
	// test
	Origin    string            `json:"origin,required"`
	Waypoints param.Opt[string] `json:"waypoints,omitzero"`
	paramObj
}

func (r DirectionComputeRouteParams) MarshalJSON() (data []byte, err error) {
	type shadow DirectionComputeRouteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectionComputeRouteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
