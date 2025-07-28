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

// OptimizationService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOptimizationService] method instead.
type OptimizationService struct {
	Options          []option.RequestOption
	DriverAssignment OptimizationDriverAssignmentService
	V2               OptimizationV2Service
}

// NewOptimizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOptimizationService(opts ...option.RequestOption) (r OptimizationService) {
	r = OptimizationService{}
	r.Options = opts
	r.DriverAssignment = NewOptimizationDriverAssignmentService(opts...)
	r.V2 = NewOptimizationV2Service(opts...)
	return
}

// Nextbillion.ai Optimization API computes and returns an optimized route between
// an origin and destination which have multiple stop points in between. With
// NextBillion.ai's Route Optimization API you get.
//
// # Optimized routing between way points
//
// # Highly accurate ETAs with customized routes
//
// # Roundtrip optimization with customized destinations
//
// A list of all parameters is specified in the next section.
func (r *OptimizationService) Compute(ctx context.Context, query OptimizationComputeParams, opts ...option.RequestOption) (res *OptimizationComputeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "optimization/json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Re-optimization
func (r *OptimizationService) ReOptimize(ctx context.Context, params OptimizationReOptimizeParams, opts ...option.RequestOption) (res *PostResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "optimization/re_optimization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PostResponse struct {
	// A unique ID which can be used in the Optimization GET method to retrieve the
	// result of optimization.
	ID string `json:"id"`
	// Displays an acknowledgement message once the job is submitted.
	Message string `json:"message"`
	// A string indicating the state of the response. On successful responses, the
	// value will be `Ok`. Indicative error messages/codes are returned in case of
	// errors. See the [API Error Codes](#api-error-codes) section below for more
	// information.
	Status string `json:"status"`
	// Display the warnings for the given input parameters, values and constraints.
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostResponse) RawJSON() string { return r.JSON.raw }
func (r *PostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationComputeResponse struct {
	// A string indicating the state of the response. This is a separate code than the
	// HTTP status code. On normal valid responses, the value will be `Ok`.
	Code string `json:"code"`
	// Contains the latitude and longitude of a location
	Location OptimizationComputeResponseLocation `json:"location"`
	// An array of 0 or 1 trip objects. Each object has the following schema.
	Trips []OptimizationComputeResponseTrip `json:"trips"`
	// Each waypoint is an input coordinate snapped to the road and path network.
	Waypoints []OptimizationComputeResponseWaypoint `json:"waypoints"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Location    respjson.Field
		Trips       respjson.Field
		Waypoints   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationComputeResponse) RawJSON() string { return r.JSON.raw }
func (r *OptimizationComputeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contains the latitude and longitude of a location
type OptimizationComputeResponseLocation struct {
	// Latitude coordinate of the location.
	Latitude float64 `json:"latitude"`
	// Longitude coordinate of the location.
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
func (r OptimizationComputeResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *OptimizationComputeResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationComputeResponseTrip struct {
	// Distance of the trip in meters.
	Distance float64 `json:"distance"`
	// Duration of the trip in seconds
	Duration float64 `json:"duration"`
	// The GeoJSON representation of the route.
	Geojson OptimizationComputeResponseTripGeojson `json:"geojson"`
	// `polyline` or `polyline6` format of route geometry.
	Geometry string                               `json:"geometry"`
	Legs     []OptimizationComputeResponseTripLeg `json:"legs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Duration    respjson.Field
		Geojson     respjson.Field
		Geometry    respjson.Field
		Legs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationComputeResponseTrip) RawJSON() string { return r.JSON.raw }
func (r *OptimizationComputeResponseTrip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The GeoJSON representation of the route.
type OptimizationComputeResponseTripGeojson struct {
	// The encoded geometry of the geojson in the `trip`.
	Geometry string `json:"geometry"`
	// Additional properties associated with the `trip`.
	Properties string `json:"properties"`
	// The type of the GeoJSON object.
	//
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
func (r OptimizationComputeResponseTripGeojson) RawJSON() string { return r.JSON.raw }
func (r *OptimizationComputeResponseTripGeojson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationComputeResponseTripLeg struct {
	// Distance of leg in metres.
	Distance float64 `json:"distance"`
	// Duration of leg in seconds.
	Duration float64 `json:"duration"`
	// An array of step objects.
	Steps []OptimizationComputeResponseTripLegStep `json:"steps"`
	// Summary of the `leg` object.
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Duration    respjson.Field
		Steps       respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationComputeResponseTripLeg) RawJSON() string { return r.JSON.raw }
func (r *OptimizationComputeResponseTripLeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationComputeResponseTripLegStep struct {
	// Distance of the `step` object in meters.
	Distance float64 `json:"distance"`
	// Duration of the `step` object in seconds.
	Duration float64 `json:"duration"`
	// The GeoJSON representation of the `step`.
	Geojson OptimizationComputeResponseTripLegStepGeojson `json:"geojson"`
	// Encoded geometry of the `step` in the selected format.
	Geometry string `json:"geometry"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Duration    respjson.Field
		Geojson     respjson.Field
		Geometry    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationComputeResponseTripLegStep) RawJSON() string { return r.JSON.raw }
func (r *OptimizationComputeResponseTripLegStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The GeoJSON representation of the `step`.
type OptimizationComputeResponseTripLegStepGeojson struct {
	// The encoded geometry of the geojson in the `step`.
	Geometry string `json:"geometry"`
	// Additional properties associated with the `step`.
	Properties string `json:"properties"`
	// The type of the GeoJSON object.
	//
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
func (r OptimizationComputeResponseTripLegStepGeojson) RawJSON() string { return r.JSON.raw }
func (r *OptimizationComputeResponseTripLegStepGeojson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationComputeResponseWaypoint struct {
	// Describes the location of the waypoint.
	Location OptimizationComputeResponseWaypointLocation `json:"location"`
	// Name of the waypoint.
	Name string `json:"name"`
	// Denotes the ID of a trip. Starts with 0.
	TripsIndex int64 `json:"trips_index"`
	// Denotes the id of a waypoint. The first waypoint is denoted with 0. And onwards
	// with 1, 2 etc.
	WaypointIndex int64 `json:"waypoint_index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Location      respjson.Field
		Name          respjson.Field
		TripsIndex    respjson.Field
		WaypointIndex respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationComputeResponseWaypoint) RawJSON() string { return r.JSON.raw }
func (r *OptimizationComputeResponseWaypoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the location of the waypoint.
type OptimizationComputeResponseWaypointLocation struct {
	// Latitude coordinate of the waypoint.
	Latitude float64 `json:"latitude"`
	// Longitude coordinate of the waypoint.
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
func (r OptimizationComputeResponseWaypointLocation) RawJSON() string { return r.JSON.raw }
func (r *OptimizationComputeResponseWaypointLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationComputeParams struct {
	// This is a pipe-separated list of coordinates.
	//
	// Minimum 3 pairs of coordinates and Maximum 12 pairs of coordinates are allowed.
	Coordinates string `query:"coordinates,required" format:"latitude,longitude" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Indicates whether the returned route is a roundtrip.
	Roundtrip param.Opt[bool] `query:"roundtrip,omitzero" json:"-"`
	// Indicates whether the return geometry should be computed or not.
	WithGeometry param.Opt[bool] `query:"with_geometry,omitzero" json:"-"`
	// A semicolon-separated list indicating the side of the road from which to
	// approach waypoints in a requested route. If provided, the number of `approaches`
	// must be the same as the number of `coordinates`. However, you can skip a
	// coordinate and show its position in the list with the `;` separator.
	//
	// Any of "`unrestricted`", "`curb`".
	Approaches OptimizationComputeParamsApproaches `query:"approaches,omitzero" json:"-"`
	// Specify the destination coordinate of the returned route. If the input is
	// `last`, the last coordinate will be the destination.
	//
	// Any of "`any`", "`last`".
	Destination OptimizationComputeParamsDestination `query:"destination,omitzero" json:"-"`
	// Sets the output format of the route geometry in the response.
	//
	// On providing `polyline` and `polyline6` as input, respective encoded geometry is
	// returned. However, when `geojson` is provided as the input value, `polyline`
	// encoded geometry is returned in the response along with a geojson details of the
	// route.
	//
	// Any of "`polyline`", "`polyline6`", "`geojson`".
	Geometries OptimizationComputeParamsGeometries `query:"geometries,omitzero" json:"-"`
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
	// Please use the Directions Flexible version if you want to use custom truck
	// dimensions.
	//
	// Note: Only the "car" profile is enabled by default. Please note that customized
	// profiles (including "truck") might not be available for all regions. Please
	// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
	// representative or reach out at
	// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
	// additional profiles.
	//
	// Any of "`car`", "`truck`".
	Mode OptimizationComputeParamsMode `query:"mode,omitzero" json:"-"`
	// The coordinate at which to start the returned route. If this is not configured,
	// the return route’s destination will be the first coordinate.
	//
	// Any of "`any`", "`first`".
	Source OptimizationComputeParamsSource `query:"source,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OptimizationComputeParams]'s query parameters as
// `url.Values`.
func (r OptimizationComputeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A semicolon-separated list indicating the side of the road from which to
// approach waypoints in a requested route. If provided, the number of `approaches`
// must be the same as the number of `coordinates`. However, you can skip a
// coordinate and show its position in the list with the `;` separator.
type OptimizationComputeParamsApproaches string

const (
	OptimizationComputeParamsApproachesUnrestricted OptimizationComputeParamsApproaches = "`unrestricted`"
	OptimizationComputeParamsApproachesCurb         OptimizationComputeParamsApproaches = "`curb`"
)

// Specify the destination coordinate of the returned route. If the input is
// `last`, the last coordinate will be the destination.
type OptimizationComputeParamsDestination string

const (
	OptimizationComputeParamsDestinationAny  OptimizationComputeParamsDestination = "`any`"
	OptimizationComputeParamsDestinationLast OptimizationComputeParamsDestination = "`last`"
)

// Sets the output format of the route geometry in the response.
//
// On providing `polyline` and `polyline6` as input, respective encoded geometry is
// returned. However, when `geojson` is provided as the input value, `polyline`
// encoded geometry is returned in the response along with a geojson details of the
// route.
type OptimizationComputeParamsGeometries string

const (
	OptimizationComputeParamsGeometriesPolyline  OptimizationComputeParamsGeometries = "`polyline`"
	OptimizationComputeParamsGeometriesPolyline6 OptimizationComputeParamsGeometries = "`polyline6`"
	OptimizationComputeParamsGeometriesGeojson   OptimizationComputeParamsGeometries = "`geojson`"
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
// Please use the Directions Flexible version if you want to use custom truck
// dimensions.
//
// Note: Only the "car" profile is enabled by default. Please note that customized
// profiles (including "truck") might not be available for all regions. Please
// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
// representative or reach out at
// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
// additional profiles.
type OptimizationComputeParamsMode string

const (
	OptimizationComputeParamsModeCar   OptimizationComputeParamsMode = "`car`"
	OptimizationComputeParamsModeTruck OptimizationComputeParamsMode = "`truck`"
)

// The coordinate at which to start the returned route. If this is not configured,
// the return route’s destination will be the first coordinate.
type OptimizationComputeParamsSource string

const (
	OptimizationComputeParamsSourceAny   OptimizationComputeParamsSource = "`any`"
	OptimizationComputeParamsSourceFirst OptimizationComputeParamsSource = "`first`"
)

type OptimizationReOptimizeParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Specify the unique request ID that needs to be re-optimized.
	ExistingRequestID string `json:"existing_request_id,required"`
	// This section gathers information on modifications to the number of jobs or their
	// individual requirements for re-optimization. Any job from the original solution
	// not specified here will be re-planned without alteration during the
	// re-optimization process.
	JobChanges OptimizationReOptimizeParamsJobChanges `json:"job_changes,omitzero"`
	// Provide the list of locations to be used during re-optimization process. Please
	// note that
	//
	//   - Providing the location input overwrites the list of locations used in the
	//     original request.
	//   - The location_indexes associated with all tasks and vehicles (both from the
	//     original and new re-optimization input requests) will follow the updated list
	//     of locations.
	//
	// As a best practice:
	//
	//  1. Don't provide the `locations` input when re-optimizing, if the original set
	//     contains all the required location coordinates.
	//
	//  2. If any new location coordinates are required for re-optimization, copy the
	//     full, original location list and update it in the following manner before
	//     adding it to the re-optimization input:
	//
	//  1. Ensure to not update the indexes of locations which just need to be
	//     "modified".
	//
	//  2. Add new location coordinates towards the end of the list.
	Locations []string `json:"locations,omitzero"`
	// This section gathers information on modifications to the number of shipments or
	// their individual requirements for re-optimization. Any shipment from the
	// original solution not specified here will be re-planned without alteration
	// during the re-optimization process.
	ShipmentChanges OptimizationReOptimizeParamsShipmentChanges `json:"shipment_changes,omitzero"`
	// This section gathers information on modifications to the number of vehicles or
	// individual vehicle configurations for re-optimizing an existing solution. Any
	// vehicle from the original solution not specified here will be reused without
	// alteration during the re-optimization process.
	VehicleChanges OptimizationReOptimizeParamsVehicleChanges `json:"vehicle_changes,omitzero"`
	paramObj
}

func (r OptimizationReOptimizeParams) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationReOptimizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationReOptimizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OptimizationReOptimizeParams]'s query parameters as
// `url.Values`.
func (r OptimizationReOptimizeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// This section gathers information on modifications to the number of jobs or their
// individual requirements for re-optimization. Any job from the original solution
// not specified here will be re-planned without alteration during the
// re-optimization process.
type OptimizationReOptimizeParamsJobChanges struct {
	// An array of objects to collect the details of the new jobs to be added during
	// re-optimization. Each object represents one job. Please make sure the IDs
	// provided for new jobs are unique with respect to the IDs of the jobs in the
	// original request.
	Add []JobParam `json:"add,omitzero"`
	// An array of objects to collect the modified details of existing jobs used in the
	// original request. Each object represents one job. Please make sure all the job
	// IDs provided here are same as the ones in the original request.
	Modify []JobParam `json:"modify,omitzero"`
	// An array of job IDs to be removed when during re-optimization. All job IDs
	// provided must have been part of the original request.
	Remove []string `json:"remove,omitzero"`
	paramObj
}

func (r OptimizationReOptimizeParamsJobChanges) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationReOptimizeParamsJobChanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationReOptimizeParamsJobChanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This section gathers information on modifications to the number of shipments or
// their individual requirements for re-optimization. Any shipment from the
// original solution not specified here will be re-planned without alteration
// during the re-optimization process.
type OptimizationReOptimizeParamsShipmentChanges struct {
	// An array of objects to collect the details of the new shipments to be added
	// during re-optimization. Each object represents one shipment. Please make sure
	// the IDs provided for new shipments are unique with respect to the IDs of the
	// shipments in the original request.
	Add []ShipmentParam `json:"add,omitzero"`
	// An array of objects to collect the modified details of existing shipments used
	// in the original request. Each object represents one shipment. Please make sure
	// all the shipment IDs provided here are same as the ones in the original request.
	Modify []ShipmentParam `json:"modify,omitzero"`
	// An array of shipment IDs to be removed when during re-optimization. All shipment
	// IDs provided must have been part of the original request.
	Remove []string `json:"remove,omitzero"`
	paramObj
}

func (r OptimizationReOptimizeParamsShipmentChanges) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationReOptimizeParamsShipmentChanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationReOptimizeParamsShipmentChanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This section gathers information on modifications to the number of vehicles or
// individual vehicle configurations for re-optimizing an existing solution. Any
// vehicle from the original solution not specified here will be reused without
// alteration during the re-optimization process.
type OptimizationReOptimizeParamsVehicleChanges struct {
	// An array of objects to collect the details of the new vehicles to be added for
	// re-optimization. Each object represents one vehicle. Please make sure the IDs
	// provided for new vehicles are unique with respect to the IDs of the vehicles in
	// the original request.
	Add    []VehicleParam `json:"add,omitzero"`
	Modify VehicleParam   `json:"modify,omitzero"`
	// An array of vehicle IDs to be removed when during re-optimization. All vehicle
	// IDs provided must have been part of the original request.
	Remove []string `json:"remove,omitzero"`
	paramObj
}

func (r OptimizationReOptimizeParamsVehicleChanges) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationReOptimizeParamsVehicleChanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationReOptimizeParamsVehicleChanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
