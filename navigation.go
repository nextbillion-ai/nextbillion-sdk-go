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
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// NavigationService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNavigationService] method instead.
type NavigationService struct {
	Options []option.RequestOption
}

// NewNavigationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNavigationService(opts ...option.RequestOption) (r NavigationService) {
	r = NavigationService{}
	r.Options = opts
	return
}

// Nextbillion.ai’s Navigation API is a service that computes a route between 2
// places, and also returns detailed turn by turn instructions for the route.
//
// The Navigation API can be used as an input into your Navigation app.
// Alternatively, you can directly use Nextbillion.ai’s Navigation SDK for a
// complete turn by turn navigation experience.
func (r *NavigationService) GetRoute(ctx context.Context, query NavigationGetRouteParams, opts ...option.RequestOption) (res *NavigationGetRouteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "navigation/json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NavigationGetRouteResponse struct {
	// Displays the error message in case of a failed request or operation. Please note
	// that this parameter is not returned in the response in case of a successful
	// request.
	Msg string `json:"msg"`
	// An array of objects describing different possible routes from the starting
	// location to the destination. Each object represents one route.
	Routes []NavigationGetRouteResponseRoute `json:"routes"`
	// A string indicating the state of the response. On normal responses, the value
	// will be Ok. Indicative HTTP error codes are returned for different errors. See
	// the [API Errors Codes](#api-error-codes) section below for more information.
	Status string `json:"status"`
	// warning when facing unexpected behaviour
	Warning []string `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		Routes      respjson.Field
		Status      respjson.Field
		Warning     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponse) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NavigationGetRouteResponseRoute struct {
	// The distance, in meters, of the complete trip.
	Distance float64 `json:"distance"`
	// The total distance of the route, including additional details such as extra
	// maneuvers or loops, in meters.
	DistanceFull float64 `json:"distance_full"`
	// The duration, in seconds, of the complete trip.
	Duration int64 `json:"duration"`
	// Location coordinates of the point where the route ends.
	EndLocation NavigationGetRouteResponseRouteEndLocation `json:"end_location"`
	// The GeoJSON representation of the route.
	Geojson NavigationGetRouteResponseRouteGeojson `json:"geojson"`
	// Encoded geometry of the returned route as per the selected format in geometry
	// and specified overview verbosity. Please note the overview will always be full
	// when original_shape parameter is used in the input request.
	Geometry string `json:"geometry"`
	// An array of objects returning the details about each leg of the route. waypoints
	// split the route into legs.
	Legs []NavigationGetRouteResponseRouteLeg `json:"legs"`
	// The predicted duration of the route based on real-time traffic conditions.
	PredictedDuration float64 `json:"predicted_duration"`
	// The raw estimated duration of the route in seconds.
	RawDuration float64 `json:"raw_duration"`
	// Special geospatial objects or landmarks crossed along the route.
	SpecialObjects any `json:"special_objects"`
	// Location coordinates of the point where the route starts.
	StartLocation NavigationGetRouteResponseRouteStartLocation `json:"start_location"`
	// A weight value associated with the route or leg.
	Weight float64 `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance          respjson.Field
		DistanceFull      respjson.Field
		Duration          respjson.Field
		EndLocation       respjson.Field
		Geojson           respjson.Field
		Geometry          respjson.Field
		Legs              respjson.Field
		PredictedDuration respjson.Field
		RawDuration       respjson.Field
		SpecialObjects    respjson.Field
		StartLocation     respjson.Field
		Weight            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRoute) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the route ends.
type NavigationGetRouteResponseRouteEndLocation struct {
	// Latitude of the end_location.
	Latitude float64 `json:"latitude"`
	// Longitude of the end_location.
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
func (r NavigationGetRouteResponseRouteEndLocation) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteEndLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The GeoJSON representation of the route.
type NavigationGetRouteResponseRouteGeojson struct {
	Geometry   string `json:"geometry"`
	Properties string `json:"properties"`
	// Any of "Point", "MultiPoint", "LineString", "MultiLineString", "Polygon",
	// "MultiPolygon", "GeometryCollection", "Feature", "FeatureCollection", "Link".
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
func (r NavigationGetRouteResponseRouteGeojson) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteGeojson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NavigationGetRouteResponseRouteLeg struct {
	// An object containing leg distance value, in meters.
	Distance NavigationGetRouteResponseRouteLegDistance `json:"distance"`
	// An object containing leg duration value, in seconds.
	Duration NavigationGetRouteResponseRouteLegDuration `json:"duration"`
	// Location coordinates of the point where the leg ends.
	EndLocation NavigationGetRouteResponseRouteLegEndLocation `json:"end_location"`
	// The raw estimated duration of the leg in seconds.
	RawDuration any `json:"raw_duration"`
	// Location coordinates of the point where the leg starts.
	StartLocation NavigationGetRouteResponseRouteLegStartLocation `json:"start_location"`
	// An array of step objects containing turn-by-turn guidance for easy navigation.
	Steps []NavigationGetRouteResponseRouteLegStep `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance      respjson.Field
		Duration      respjson.Field
		EndLocation   respjson.Field
		RawDuration   respjson.Field
		StartLocation respjson.Field
		Steps         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLeg) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing leg distance value, in meters.
type NavigationGetRouteResponseRouteLegDistance struct {
	Value int64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegDistance) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegDistance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing leg duration value, in seconds.
type NavigationGetRouteResponseRouteLegDuration struct {
	Value int64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegDuration) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the leg ends.
type NavigationGetRouteResponseRouteLegEndLocation struct {
	// Latitude of end_location of the leg.
	Latitude float64 `json:"latitude"`
	// Longitude of end_location of the leg.
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
func (r NavigationGetRouteResponseRouteLegEndLocation) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegEndLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the leg starts.
type NavigationGetRouteResponseRouteLegStartLocation struct {
	// Latitude of start_location of the leg.
	Latitude float64 `json:"latitude"`
	// Longitude of start_location of the leg.
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
func (r NavigationGetRouteResponseRouteLegStartLocation) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStartLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NavigationGetRouteResponseRouteLegStep struct {
	// An object containing step distance value, in meters.
	Distance NavigationGetRouteResponseRouteLegStepDistance `json:"distance"`
	// Indicates the driving side of the road in case bidirectional traffic is allowed
	// on the given segment. It can have two values: "left" & "right".
	DrivingSide string `json:"driving_side"`
	// An object containing step duration value, in seconds.
	Duration NavigationGetRouteResponseRouteLegStepDuration `json:"duration"`
	// Location coordinates of the point where the step ends.
	EndLocation NavigationGetRouteResponseRouteLegStepEndLocation `json:"end_location"`
	// The GeoJSON representation of the step.
	Geojson NavigationGetRouteResponseRouteLegStepGeojson `json:"geojson"`
	// Encoded geometry of the step in the selected format.
	Geometry string `json:"geometry"`
	// An array of objects representing intersections (or cross-way) that the route
	// passes by along the step. For every step, the very first intersection
	// corresponds to the location of the maneuver. All intersections until the next
	// maneuver are listed in this object.
	Intersections []NavigationGetRouteResponseRouteLegStepIntersection `json:"intersections"`
	// An object with maneuver details for the step.
	Maneuver NavigationGetRouteResponseRouteLegStepManeuver `json:"maneuver"`
	// The name of the step.
	Name string `json:"name"`
	// A reference for the step.
	Reference string `json:"reference"`
	// An object containing road shield information.
	RoadShieldType NavigationGetRouteResponseRouteLegStepRoadShieldType `json:"road_shield_type"`
	// Location coordinates of the point where the step starts.
	StartLocation NavigationGetRouteResponseRouteLegStepStartLocation `json:"start_location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance       respjson.Field
		DrivingSide    respjson.Field
		Duration       respjson.Field
		EndLocation    respjson.Field
		Geojson        respjson.Field
		Geometry       respjson.Field
		Intersections  respjson.Field
		Maneuver       respjson.Field
		Name           respjson.Field
		Reference      respjson.Field
		RoadShieldType respjson.Field
		StartLocation  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStep) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing step distance value, in meters.
type NavigationGetRouteResponseRouteLegStepDistance struct {
	Value int64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStepDistance) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStepDistance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing step duration value, in seconds.
type NavigationGetRouteResponseRouteLegStepDuration struct {
	Value int64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStepDuration) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStepDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the step ends.
type NavigationGetRouteResponseRouteLegStepEndLocation struct {
	// Latitude of the end_location of the step.
	Latitude float64 `json:"latitude"`
	// Longitude of the end_location of the step.
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
func (r NavigationGetRouteResponseRouteLegStepEndLocation) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStepEndLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The GeoJSON representation of the step.
type NavigationGetRouteResponseRouteLegStepGeojson struct {
	Geometry string `json:"geometry"`
	Type     string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geometry    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStepGeojson) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStepGeojson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NavigationGetRouteResponseRouteLegStepIntersection struct {
	// A list of bearing values (e.g. [0,90,180,270]) that are available at the
	// intersection. The bearings describe all available roads at the intersection.
	Bearings []int64 `json:"bearings"`
	// An array of strings representing the classes or types of roads or paths at the
	// intersection. The classes can indicate the road hierarchy, such as a motorway,
	// primary road, secondary road, etc.
	Classes []string `json:"classes"`
	// A value of true indicates that the respective road could be entered on a valid
	// route. false indicates that the turn onto the respective road would violate a
	// restriction. Each entry value corresponds to the bearing angle at the same
	// index.
	Entry []bool `json:"entry"`
	// The number of incoming roads or paths at the intersection.
	IntersectionIn int64 `json:"intersection_in"`
	// The number of outgoing roads or paths from the intersection.
	IntersectionOut int64 `json:"intersection_out"`
	// An array of lane objects representing the lanes available at the intersection.
	// If no lane information is available for an intersection, the lanes property will
	// not be present.
	Lanes []NavigationGetRouteResponseRouteLegStepIntersectionLane `json:"lanes"`
	// A [longitude, latitude] pair describing the location of the intersection.
	Location NavigationGetRouteResponseRouteLegStepIntersectionLocation `json:"location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bearings        respjson.Field
		Classes         respjson.Field
		Entry           respjson.Field
		IntersectionIn  respjson.Field
		IntersectionOut respjson.Field
		Lanes           respjson.Field
		Location        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStepIntersection) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStepIntersection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NavigationGetRouteResponseRouteLegStepIntersectionLane struct {
	// It represents actions associated with the lane. These indications describe the
	// permitted maneuvers or directions that can be taken from the lane. Common
	// indications include "turn left," "turn right," "go straight," "merge," "exit,"
	// etc.
	Indications []string `json:"indications"`
	// This indicates the validity of the lane. It specifies whether the lane is
	// considered valid for making the indicated maneuver or if there are any
	// restrictions or limitations associated with it.
	Valid bool `json:"valid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indications respjson.Field
		Valid       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStepIntersectionLane) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStepIntersectionLane) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A [longitude, latitude] pair describing the location of the intersection.
type NavigationGetRouteResponseRouteLegStepIntersectionLocation struct {
	// The latitude coordinate of the intersection.
	Latitude float64 `json:"latitude"`
	// The longitude coordinate of the intersection.
	Longitude float64 `json:"longitude"`
	// The name or description of the intersection.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStepIntersectionLocation) RawJSON() string {
	return r.JSON.raw
}
func (r *NavigationGetRouteResponseRouteLegStepIntersectionLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with maneuver details for the step.
type NavigationGetRouteResponseRouteLegStepManeuver struct {
	// The clockwise angle from true north to the direction of travel immediately after
	// the maneuver. Range of values is between 0-359.
	BearingAfter float64 `json:"bearing_after"`
	// The clockwise angle from true north to the direction of travel immediately
	// before the maneuver. Range of values is between 0-359.
	BearingBefore float64 `json:"bearing_before"`
	// A coordinate pair describing the location of the maneuver.
	Coordinate NavigationGetRouteResponseRouteLegStepManeuverCoordinate `json:"coordinate"`
	// A text instruction describing the maneuver to be performed. It provides guidance
	// on the action to take at the maneuver location, such as "Turn left," "Go
	// straight," "Exit the roundabout," etc.
	Instruction string `json:"instruction"`
	// A string indicating the type of maneuver.
	ManeuverType string `json:"maneuver_type"`
	// A boolean value indicating whether the voice instruction for the maneuver should
	// be muted or not.
	Muted bool `json:"muted"`
	// The number of roundabouts encountered so far during the route. This parameter is
	// specific to roundabout maneuvers and indicates the count of roundabouts before
	// the current one.
	RoundaboutCount int64 `json:"roundabout_count"`
	// An array of voice instruction objects associated with the maneuver. Each object
	// provides additional details about the voice instruction, including the distance
	// along the geometry where the instruction applies, the instruction text, and the
	// unit of measurement.
	VoiceInstruction []NavigationGetRouteResponseRouteLegStepManeuverVoiceInstruction `json:"voice_instruction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BearingAfter     respjson.Field
		BearingBefore    respjson.Field
		Coordinate       respjson.Field
		Instruction      respjson.Field
		ManeuverType     respjson.Field
		Muted            respjson.Field
		RoundaboutCount  respjson.Field
		VoiceInstruction respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStepManeuver) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStepManeuver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A coordinate pair describing the location of the maneuver.
type NavigationGetRouteResponseRouteLegStepManeuverCoordinate struct {
	// The latitude coordinate of the maneuver.
	Latitude float64 `json:"latitude"`
	// The longitude coordinate of the maneuver.
	Longitude float64 `json:"longitude"`
	// The name or description of the maneuver location.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStepManeuverCoordinate) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStepManeuverCoordinate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NavigationGetRouteResponseRouteLegStepManeuverVoiceInstruction struct {
	DistanceAlongGeometry int64 `json:"distance_along_geometry"`
	// The guidance instructions for the upcoming maneuver
	Instruction string `json:"instruction"`
	// Unit of the distance_along_geometry metric
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DistanceAlongGeometry respjson.Field
		Instruction           respjson.Field
		Unit                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStepManeuverVoiceInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *NavigationGetRouteResponseRouteLegStepManeuverVoiceInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing road shield information.
type NavigationGetRouteResponseRouteLegStepRoadShieldType struct {
	// The URL to fetch the road shield image.
	ImageURL string `json:"image_url"`
	// A label identifying the inscription on the road shield, such as containing the
	// road number.
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigationGetRouteResponseRouteLegStepRoadShieldType) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStepRoadShieldType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the step starts.
type NavigationGetRouteResponseRouteLegStepStartLocation struct {
	// Latitude of start_location of the step.
	Latitude float64 `json:"latitude"`
	// Longitude of start_location of the step.
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
func (r NavigationGetRouteResponseRouteLegStepStartLocation) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteLegStepStartLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Location coordinates of the point where the route starts.
type NavigationGetRouteResponseRouteStartLocation struct {
	// Latitude of thestart_location.
	Latitude float64 `json:"latitude"`
	// Longitude of the start_location.
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
func (r NavigationGetRouteResponseRouteStartLocation) RawJSON() string { return r.JSON.raw }
func (r *NavigationGetRouteResponseRouteStartLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NavigationGetRouteParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Sets the number of alternative routes to return. It is effective only when
	// "alternatives" is "true". Please note that adding alternative route count does
	// not guarantee matching number of routes to be returned if potential alternative
	// routes do not exist.
	Altcount param.Opt[int64] `query:"altcount,omitzero" json:"-"`
	// When "true" the API will return alternate routes. The "alternatives" is
	// effective only when there are no "waypoints" included in the request. You can
	// set the number of alternate routes to be returned in the "altcount" property.
	Alternatives param.Opt[bool] `query:"alternatives,omitzero" json:"-"`
	// Limits the search to road segments with given bearing, in degrees, towards true
	// north in clockwise direction. Each "bearings" should be in the format of
	// "degree,range", where the "degree" should be a value between \[0, 360\] and
	// "range" should be a value between \[0, 180\]. Please note that the number of
	// "bearings" should be two more than the number of "waypoints". This is to account
	// for the bearing of "origin" and "destination". If a route can approach a
	// "waypoint" or the "destination" from any direction, the bearing for that point
	// can be specified as "0,180".
	Bearings param.Opt[string] `query:"bearings,omitzero" format:"degree0,range0;degree1,range1;..." json:"-"`
	// "destination" is the ending point of your route. Ensure that the "destination"
	// is a routable land location. Please note that this parameter is mandatory if the
	// "original_shape" parameter is not given.
	Destination param.Opt[string] `query:"destination,omitzero" format:"latitude,longitude" json:"-"`
	// Select the language to be used for result rendering from a list of
	// [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language
	// codes.
	Lang param.Opt[string] `query:"lang,omitzero" json:"-"`
	// "origin" is the starting point of your route. Ensure that "origin" is a routable
	// land location. Please note that this parameter is mandatory if the geometry
	// parameter is not given.
	Origin param.Opt[string] `query:"origin,omitzero" format:"latitude,longitude" json:"-"`
	// Takes a route geometry as input and returns the navigation information for that
	// route. Accepts "polyline" and "polyline6" encoded geometry as input.
	// "original_shape_type" becomes mandatory when "original_shape" is used. If this
	// parameter is provided, the only other parameters which will be considered are
	// "original_shape_type", "lang", "geometry". The rest of the parameters in the
	// input request will be ignored. Please note overview verbosity will always be
	// "full" when using this parameter.
	OriginalShape param.Opt[string] `query:"original_shape,omitzero" json:"-"`
	// "waypoints" are coordinates along the route between the "origin" and
	// "destination". It is a pipe-separated list of coordinate pairs. Please note that
	// the route returned will arrive at the "waypoints" in the sequence they are
	// provided in the input request. Please note that the maximum number of waypoints
	// that can be provided in a single request is 50 when using GET method and 200
	// with POST method.
	Waypoints param.Opt[string] `query:"waypoints,omitzero" format:"latitude_1,longitude_1|latitude_2,longitude_2|..." json:"-"`
	// A semicolon-separated list indicating the side of the road from which to
	// approach "waypoints" in a requested route. When set to "unrestricted" a route
	// can arrive at the waypoint from either side of the road and when set to "curb"
	// the route will arrive at the waypoint on the driving side of the region. Please
	// note the number of values provided must be one more than the number of
	// "waypoints". The last value of "approaches" will determine the approach for the
	// "destination". However, you can skip a coordinate and show its position in the
	// list with the ";" separator.
	//
	// Any of "unrestricted", "curb".
	Approaches NavigationGetRouteParamsApproaches `query:"approaches,omitzero" json:"-"`
	// Setting this will ensure the route avoids ferries, tolls, highways or nothing.
	// Multiple values should be separated by a pipe (|). If "none" is provided along
	// with other values, an error is returned as a valid route is not feasible. Please
	// note that when this parameter is not provided in the input, ferries are set to
	// be avoided by default. When this parameter is provided, only the mentioned
	// objects are avoided.
	//
	// Any of "toll", "ferry", "highway", "none".
	Avoid NavigationGetRouteParamsAvoid `query:"avoid,omitzero" json:"-"`
	// Sets the output format of the route geometry in the response. On providing
	// “polyline“ and “polyline6“ as input, respective encoded geometry is returned.
	// However, when “geojson“ is provided as the input value, “polyline“ encoded
	// geometry is returned in the response along with the geojson details of the
	// route.
	//
	// Any of "polyline", "polyline6", "geojson".
	Geometry NavigationGetRouteParamsGeometry `query:"geometry,omitzero" json:"-"`
	// Set which driving mode the service should use to determine a route. For example,
	// if you use "car", the API will return a route that a car can take. Using "truck"
	// will return a route a truck can use, taking into account appropriate truck
	// routing restrictions.
	//
	// When "mode=truck", following are the default dimensions that are used:
	//
	// \- truck_height = 214 centimeters
	//
	// \- truck_width = 183 centimeters
	//
	// \- truck_length = 519 centimeters
	//
	// \- truck_weight = 5000 kg
	//
	// Please use the Navigation Flexible version if you want to use custom truck
	// dimensions.
	//
	// Note: Only the "car" profile is enabled by default. Please note that customized
	// profiles (including "truck") might not be available for all regions. Please
	// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
	// representative or reach out at
	// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
	// additional profiles.
	//
	// Any of "car", "truck".
	Mode NavigationGetRouteParamsMode `query:"mode,omitzero" json:"-"`
	// Specify the encoding format of route geometry provided in the request using
	// "original_shape" parameter. Please note that an error is returned when this
	// parameter is not specified while an input is added to "original_shape"
	// parameter.
	//
	// Any of "polyline", "polyline6".
	OriginalShapeType NavigationGetRouteParamsOriginalShapeType `query:"original_shape_type,omitzero" json:"-"`
	// Specify the verbosity of route geometry. When set to "full", the most detailed
	// geometry available is returned. When set to "simplified", a simplified version
	// of the full geometry is returned. No overview geometry is returned when set to
	// "false".
	//
	// Any of "full", "simplified", "false".
	Overview NavigationGetRouteParamsOverview `query:"overview,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NavigationGetRouteParams]'s query parameters as
// `url.Values`.
func (r NavigationGetRouteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A semicolon-separated list indicating the side of the road from which to
// approach "waypoints" in a requested route. When set to "unrestricted" a route
// can arrive at the waypoint from either side of the road and when set to "curb"
// the route will arrive at the waypoint on the driving side of the region. Please
// note the number of values provided must be one more than the number of
// "waypoints". The last value of "approaches" will determine the approach for the
// "destination". However, you can skip a coordinate and show its position in the
// list with the ";" separator.
type NavigationGetRouteParamsApproaches string

const (
	NavigationGetRouteParamsApproachesUnrestricted NavigationGetRouteParamsApproaches = "unrestricted"
	NavigationGetRouteParamsApproachesCurb         NavigationGetRouteParamsApproaches = "curb"
)

// Setting this will ensure the route avoids ferries, tolls, highways or nothing.
// Multiple values should be separated by a pipe (|). If "none" is provided along
// with other values, an error is returned as a valid route is not feasible. Please
// note that when this parameter is not provided in the input, ferries are set to
// be avoided by default. When this parameter is provided, only the mentioned
// objects are avoided.
type NavigationGetRouteParamsAvoid string

const (
	NavigationGetRouteParamsAvoidToll    NavigationGetRouteParamsAvoid = "toll"
	NavigationGetRouteParamsAvoidFerry   NavigationGetRouteParamsAvoid = "ferry"
	NavigationGetRouteParamsAvoidHighway NavigationGetRouteParamsAvoid = "highway"
	NavigationGetRouteParamsAvoidNone    NavigationGetRouteParamsAvoid = "none"
)

// Sets the output format of the route geometry in the response. On providing
// “polyline“ and “polyline6“ as input, respective encoded geometry is returned.
// However, when “geojson“ is provided as the input value, “polyline“ encoded
// geometry is returned in the response along with the geojson details of the
// route.
type NavigationGetRouteParamsGeometry string

const (
	NavigationGetRouteParamsGeometryPolyline  NavigationGetRouteParamsGeometry = "polyline"
	NavigationGetRouteParamsGeometryPolyline6 NavigationGetRouteParamsGeometry = "polyline6"
	NavigationGetRouteParamsGeometryGeojson   NavigationGetRouteParamsGeometry = "geojson"
)

// Set which driving mode the service should use to determine a route. For example,
// if you use "car", the API will return a route that a car can take. Using "truck"
// will return a route a truck can use, taking into account appropriate truck
// routing restrictions.
//
// When "mode=truck", following are the default dimensions that are used:
//
// \- truck_height = 214 centimeters
//
// \- truck_width = 183 centimeters
//
// \- truck_length = 519 centimeters
//
// \- truck_weight = 5000 kg
//
// Please use the Navigation Flexible version if you want to use custom truck
// dimensions.
//
// Note: Only the "car" profile is enabled by default. Please note that customized
// profiles (including "truck") might not be available for all regions. Please
// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
// representative or reach out at
// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
// additional profiles.
type NavigationGetRouteParamsMode string

const (
	NavigationGetRouteParamsModeCar   NavigationGetRouteParamsMode = "car"
	NavigationGetRouteParamsModeTruck NavigationGetRouteParamsMode = "truck"
)

// Specify the encoding format of route geometry provided in the request using
// "original_shape" parameter. Please note that an error is returned when this
// parameter is not specified while an input is added to "original_shape"
// parameter.
type NavigationGetRouteParamsOriginalShapeType string

const (
	NavigationGetRouteParamsOriginalShapeTypePolyline  NavigationGetRouteParamsOriginalShapeType = "polyline"
	NavigationGetRouteParamsOriginalShapeTypePolyline6 NavigationGetRouteParamsOriginalShapeType = "polyline6"
)

// Specify the verbosity of route geometry. When set to "full", the most detailed
// geometry available is returned. When set to "simplified", a simplified version
// of the full geometry is returned. No overview geometry is returned when set to
// "false".
type NavigationGetRouteParamsOverview string

const (
	NavigationGetRouteParamsOverviewFull       NavigationGetRouteParamsOverview = "full"
	NavigationGetRouteParamsOverviewSimplified NavigationGetRouteParamsOverview = "simplified"
	NavigationGetRouteParamsOverviewFalse      NavigationGetRouteParamsOverview = "false"
)
