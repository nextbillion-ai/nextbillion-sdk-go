// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apijson"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/nextbillion-sdk-go/option"
	"github.com/stainless-sdks/nextbillion-sdk-go/packages/param"
	"github.com/stainless-sdks/nextbillion-sdk-go/packages/respjson"
)

// OptimizationDriverAssignmentService contains methods and other services that
// help with interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOptimizationDriverAssignmentService] method instead.
type OptimizationDriverAssignmentService struct {
	Options []option.RequestOption
}

// NewOptimizationDriverAssignmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOptimizationDriverAssignmentService(opts ...option.RequestOption) (r OptimizationDriverAssignmentService) {
	r = OptimizationDriverAssignmentService{}
	r.Options = opts
	return
}

// Assigns available drivers (vehicles) to open orders based on specified criteria
// and constraints.
func (r *OptimizationDriverAssignmentService) AssignDrivers(ctx context.Context, params OptimizationDriverAssignmentAssignDriversParams, opts ...option.RequestOption) (res *OptimizationDriverAssignmentAssignDriversResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "optimization/driver-assignment/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Location info.
type Location struct {
	// Latitude of location.
	Lat float64 `json:"lat,required"`
	// Longitude of location.
	Lon float64 `json:"lon,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lon         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Location) RawJSON() string { return r.JSON.raw }
func (r *Location) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Location to a LocationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LocationParam.Overrides()
func (r Location) ToParam() LocationParam {
	return param.Override[LocationParam](json.RawMessage(r.RawJSON()))
}

// Location info.
//
// The properties Lat, Lon are required.
type LocationParam struct {
	// Latitude of location.
	Lat float64 `json:"lat,required"`
	// Longitude of location.
	Lon float64 `json:"lon,required"`
	paramObj
}

func (r LocationParam) MarshalJSON() (data []byte, err error) {
	type shadow LocationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Location are required.
type VehicleParam struct {
	// Specify a unique ID for the vehicle.
	ID string `json:"id,required"`
	// Specify the location coordinates where the vehicle is currently located. This
	// input is mandatory for each vehicle.
	Location VehicleLocationParam `json:"location,omitzero,required"`
	// Specify the priority for this vehicle. A higher value indicates a higher
	// priority. When specified, it will override any priority score deduced from
	// `vehicle_attribute_priority_mappings` for this vehicle. Valid values are \[1,
	// 10\] and default is 0.
	Priority param.Opt[int64] `json:"priority,omitzero"`
	// Specify custom attributes for the vehicle. Each attribute should be created as a
	// `key:value` pair. These attributes can be used in the
	// `orders.vehicle_preferences` input to refine the search of vehicles for each
	// order.
	//
	// The maximum number of key:value pairs that can be specified under `attributes`
	// for a given vehicle, is limited to 30.
	Attributes any `json:"attributes,omitzero"`
	// An array of objects to collect the location coordinates of the stops remaining
	// on an ongoing trip of the vehicle. The service can assign new orders to the
	// vehicle if they are cost-effective. Once a new order is assigned, the vehicle
	// must complete all the steps in the ongoing trip before proceeding to pickup the
	// newly assigned order.
	//
	// Please note that a maximum of 10 waypoints can be specified for a given vehicle.
	RemainingWaypoints []LocationParam `json:"remaining_waypoints,omitzero"`
	paramObj
}

func (r VehicleParam) MarshalJSON() (data []byte, err error) {
	type shadow VehicleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VehicleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specify the location coordinates where the vehicle is currently located. This
// input is mandatory for each vehicle.
type VehicleLocationParam struct {
	// Latitude of the vehicle's current location.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Longitude of the vehicle's current location.
	Lng param.Opt[float64] `json:"lng,omitzero"`
	paramObj
}

func (r VehicleLocationParam) MarshalJSON() (data []byte, err error) {
	type shadow VehicleLocationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VehicleLocationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationDriverAssignmentAssignDriversResponse struct {
	// Displays indicative error message in case of a failed request or operation.
	// Please note that this parameter is not returned in the response in case of a
	// successful request.
	Message string `json:"message"`
	// An object containing the details of the assignments.
	Result OptimizationDriverAssignmentAssignDriversResponseResult `json:"result"`
	// An integer indicating the HTTP response code. See the
	// [API Error Handling](https://docs.nextbillion.ai/optimization/driver-assignment-api#api-error-handling)
	// section below for more information.
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Result      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationDriverAssignmentAssignDriversResponse) RawJSON() string { return r.JSON.raw }
func (r *OptimizationDriverAssignmentAssignDriversResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the details of the assignments.
type OptimizationDriverAssignmentAssignDriversResponseResult struct {
	// An array of objects containing the details of the potential, alternate vehicle
	// assignments for the orders in the input. This attribute will not be returned in
	// the response if the `alternate_assignments` was not provided in the input. Each
	// object represents alternate assignments for a single order.
	AlternateAssignments []OptimizationDriverAssignmentAssignDriversResponseResultAlternateAssignment `json:"alternate_assignments"`
	// A collection of vehicles IDs that were not assigned to any orders. A `null`
	// value is returned if there are no vehicles without an order assignment.
	AvailableVehicles []string `json:"available_vehicles"`
	// An collection of objects returning the trip details for each vehicle which was
	// assigned to an order. Each object corresponds to one vehicle.
	Trips []OptimizationDriverAssignmentAssignDriversResponseResultTrip `json:"trips"`
	// A collection of objects listing the details of orders which remained unassigned.
	// Each object represents a single order. A `null` value is returned if there are
	// no unassigned orders.
	UnassignedOrders []OptimizationDriverAssignmentAssignDriversResponseResultUnassignedOrder `json:"unassigned_orders"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AlternateAssignments respjson.Field
		AvailableVehicles    respjson.Field
		Trips                respjson.Field
		UnassignedOrders     respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationDriverAssignmentAssignDriversResponseResult) RawJSON() string { return r.JSON.raw }
func (r *OptimizationDriverAssignmentAssignDriversResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationDriverAssignmentAssignDriversResponseResultAlternateAssignment struct {
	// An array of objects containing the details of the alternate vehicle assignments.
	// Each object represents an alternate vehicle assignment.
	Assignments []OptimizationDriverAssignmentAssignDriversResponseResultAlternateAssignmentAssignment `json:"assignments"`
	// Returns the order ID associated with the alternate assignments.
	OrderID string `json:"order_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assignments respjson.Field
		OrderID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationDriverAssignmentAssignDriversResponseResultAlternateAssignment) RawJSON() string {
	return r.JSON.raw
}
func (r *OptimizationDriverAssignmentAssignDriversResponseResultAlternateAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationDriverAssignmentAssignDriversResponseResultAlternateAssignmentAssignment struct {
	// Returns the ETA to the order's pickup location for the given vehicle.
	PickupEta int64 `json:"pickup_eta"`
	// Returns the vehicle ID which could potentially be assigned to the given order.
	VehicleID string `json:"vehicle_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PickupEta   respjson.Field
		VehicleID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationDriverAssignmentAssignDriversResponseResultAlternateAssignmentAssignment) RawJSON() string {
	return r.JSON.raw
}
func (r *OptimizationDriverAssignmentAssignDriversResponseResultAlternateAssignmentAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationDriverAssignmentAssignDriversResponseResultTrip struct {
	// Returns a unique trip ID.
	TripID string `json:"trip_id"`
	// Returns the details of the vehicle, assigned order and the trip steps.
	Vehicle OptimizationDriverAssignmentAssignDriversResponseResultTripVehicle `json:"vehicle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TripID      respjson.Field
		Vehicle     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationDriverAssignmentAssignDriversResponseResultTrip) RawJSON() string {
	return r.JSON.raw
}
func (r *OptimizationDriverAssignmentAssignDriversResponseResultTrip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the details of the vehicle, assigned order and the trip steps.
type OptimizationDriverAssignmentAssignDriversResponseResultTripVehicle struct {
	// Returns the ID of the vehicle.
	ID string `json:"id"`
	// A collection of objects returning the sequence of steps that the vehicle needs
	// to perform for a trip.
	Steps OptimizationDriverAssignmentAssignDriversResponseResultTripVehicleSteps `json:"steps"`
	// Location info.
	VehicleCurrentLocation Location `json:"vehicle_current_location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Steps                  respjson.Field
		VehicleCurrentLocation respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationDriverAssignmentAssignDriversResponseResultTripVehicle) RawJSON() string {
	return r.JSON.raw
}
func (r *OptimizationDriverAssignmentAssignDriversResponseResultTripVehicle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of objects returning the sequence of steps that the vehicle needs
// to perform for a trip.
type OptimizationDriverAssignmentAssignDriversResponseResultTripVehicleSteps struct {
	// Returns the driving distance, in meters, to the step's location from previous
	// step's location. For the first step of a trip, `distance` indicates the driving
	// distance from `vehicle_current_location` to the step's location.
	Distance int64 `json:"distance"`
	// Returns the driving duration, in seconds, to the step's location from previous
	// step's location. For the first step of a trip, `eta` indicates the driving
	// duration from `vehicle_current_location` to the step's location.
	Eta int64 `json:"eta"`
	// Location info.
	Location Location `json:"location"`
	// Returns the ID of the order. In case the step type is **ongoing**, an empty
	// string is returned.
	OrderID string `json:"order_id"`
	// Returns the type of the step. Currently, it can take following values:
	//
	//   - **pickup:** Indicates the pickup step for an order
	//   - **dropoff:** Indicates the dropoff step for an order. It is returned only if
	//     `dropoff_details` was **true** in the input request.
	//   - **ongoing:** Indicates a step that the vehicle needs to complete on its
	//     current trip. This is returned in the response only when `remaining_waypoints`
	//     input was provided for the given vehicle.
	//   - **intermediate_waypoint:** Indicates an intermediate stop that the vehicle
	//     needs to complete in case multiple dropoffs are provided in the input.
	//
	// Any of "pickup", "dropoff", "ongoing".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Eta         respjson.Field
		Location    respjson.Field
		OrderID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationDriverAssignmentAssignDriversResponseResultTripVehicleSteps) RawJSON() string {
	return r.JSON.raw
}
func (r *OptimizationDriverAssignmentAssignDriversResponseResultTripVehicleSteps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationDriverAssignmentAssignDriversResponseResultUnassignedOrder struct {
	// Returns the ID of the order which remained unassigned.
	OrderID string `json:"order_id"`
	// Returns the most probable reason due to which the order remained unassigned.
	UnassignedReason string `json:"unassigned_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderID          respjson.Field
		UnassignedReason respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizationDriverAssignmentAssignDriversResponseResultUnassignedOrder) RawJSON() string {
	return r.JSON.raw
}
func (r *OptimizationDriverAssignmentAssignDriversResponseResultUnassignedOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationDriverAssignmentAssignDriversParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Specify the filtering criterion for the vehicles with respect to each order's
	// location. `filter` is a mandatory input for all requests.
	Filter OptimizationDriverAssignmentAssignDriversParamsFilter `json:"filter,omitzero,required"`
	// Collects the details of open orders to be fulfilled. Each object represents one
	// order. All requests must include `orders` as a mandatory input. A maximum of 200
	// orders is allowed per request.
	Orders []OptimizationDriverAssignmentAssignDriversParamsOrder `json:"orders,omitzero,required"`
	// Collects the details of vehicles available to fulfill the orders. Each object
	// represents one vehicle. All requests must include `vehicles` as a mandatory
	// input. A maximum of 100 vehicles is allowed per request.
	Vehicles []VehicleParam `json:"vehicles,omitzero,required"`
	// Configure the assignment constraints and response settings.
	Options OptimizationDriverAssignmentAssignDriversParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParams) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OptimizationDriverAssignmentAssignDriversParams]'s query
// parameters as `url.Values`.
func (r OptimizationDriverAssignmentAssignDriversParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the filtering criterion for the vehicles with respect to each order's
// location. `filter` is a mandatory input for all requests.
type OptimizationDriverAssignmentAssignDriversParamsFilter struct {
	// Defines a `driving_distance` filter, in meters. If a vehicle needs to drive
	// further than this distance to reach a pickup location, it will not be assigned
	// to that order. Valid range of values for this filter is \[1, 10000\].
	DrivingDistance param.Opt[float64] `json:"driving_distance,omitzero"`
	// Specify a duration, in seconds, which will be used to filter out ineligible
	// vehicles for each order. Any vehicle which would take more time than specified
	// here, to reach the pickup location of a given order, will be ruled out for
	// assignment for that particular order. Valid values for `pickup_eta` are \[1,
	// 3600\].
	PickupEta param.Opt[int64] `json:"pickup_eta,omitzero"`
	// Specify a radius, in meters, which will be used to filter out ineligible
	// vehicles for each order. The pickup location of an order will act as the center
	// of the circle when identifying eligible vehicles. Valid values for `radius` are
	// \[1, 10000\].
	Radius param.Opt[float64] `json:"radius,omitzero"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Pickup are required.
type OptimizationDriverAssignmentAssignDriversParamsOrder struct {
	// Specify a unique ID for the order.
	ID string `json:"id,required"`
	// Specify the location coordinates of the pickup location of the order. This input
	// is mandatory for each order.
	Pickup OptimizationDriverAssignmentAssignDriversParamsOrderPickup `json:"pickup,omitzero,required"`
	// Specify the priority for this order. A higher value indicates a higher priority.
	// When specified, it will override any priority score deduced from
	// `order_attribute_priority_mappings` for this order. Valid values are \[1, 10\]
	// and default is 0.
	Priority param.Opt[int64] `json:"priority,omitzero"`
	// Specify the service time, in seconds, for the order. Service time is the
	// duration that the driver is likely to wait at the pickup location after
	// arriving. The impact of the service time is realized in the ETA for the
	// "dropoff" type step.
	ServiceTime param.Opt[int64] `json:"service_time,omitzero"`
	// Specify custom attributes for the orders. Each attribute should be created as a
	// `key:value` pair. The **keys** provided can be used in
	// `options.order_attribute_priority_mappings` to assign a custom priority for this
	// order based on its attributes.
	//
	// The maximum number of key:value pairs that can be specified under `attributes`
	// for a given order, is limited to 30.
	Attributes any `json:"attributes,omitzero"`
	// Use this parameter to specify the location coordinates of the destination of the
	// trip or the intermediate stops to be completed before it.
	//
	// # Please note
	//
	// - The last location provided is treated as the destination of the trip.
	// - `dropoffs` is mandatory when `dropoff_details` is set to **true**.
	Dropoffs []OptimizationDriverAssignmentAssignDriversParamsOrderDropoff `json:"dropoffs,omitzero"`
	// Define custom preferences for task assignment based on vehicle's attributes. If
	// multiple criteria are provided, they are evaluated using an AND
	// condition—meaning all specified criteria must be met individually for a vehicle
	// to be considered.
	//
	// For example, if `required_all_of_attributes`, `required_any_of_attributes`, and
	// `exclude_all_of_attributes` are all provided, an eligible vehicle must satisfy
	// the following to be considered for assignments:
	//
	// 1.  Meet all conditions specified in `required_all_of_attributes`.
	// 2.  Meet at least one of the conditions listed in `required_any_of_attributes`.
	// 3.  Not meet any conditions mentioned in `exclude_all_of_attributes`.
	//
	// Consequently, a vehicle which does not have any `attributes` defined can't be
	// assigned to an order which has `vehicle_preferences` configured.
	VehiclePreferences OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferences `json:"vehicle_preferences,omitzero"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsOrder) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsOrder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specify the location coordinates of the pickup location of the order. This input
// is mandatory for each order.
type OptimizationDriverAssignmentAssignDriversParamsOrderPickup struct {
	// Latitude of the pickup location.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Longitude of the pickup location.
	Lng param.Opt[float64] `json:"lng,omitzero"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsOrderPickup) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsOrderPickup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsOrderPickup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizationDriverAssignmentAssignDriversParamsOrderDropoff struct {
	// Latitude of the stop location.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Longitude of the stop location.
	Lng param.Opt[float64] `json:"lng,omitzero"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsOrderDropoff) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsOrderDropoff
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsOrderDropoff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Define custom preferences for task assignment based on vehicle's attributes. If
// multiple criteria are provided, they are evaluated using an AND
// condition—meaning all specified criteria must be met individually for a vehicle
// to be considered.
//
// For example, if `required_all_of_attributes`, `required_any_of_attributes`, and
// `exclude_all_of_attributes` are all provided, an eligible vehicle must satisfy
// the following to be considered for assignments:
//
// 1.  Meet all conditions specified in `required_all_of_attributes`.
// 2.  Meet at least one of the conditions listed in `required_any_of_attributes`.
// 3.  Not meet any conditions mentioned in `exclude_all_of_attributes`.
//
// Consequently, a vehicle which does not have any `attributes` defined can't be
// assigned to an order which has `vehicle_preferences` configured.
type OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferences struct {
	// An array of objects to add exclusion requirements for the order. A vehicle must
	// **not meet any of the conditions** specified here to be considered for
	// assignment. Each object represents a single condition. Please note that a
	// maximum of 10 conditions can be added here for a given order.
	ExcludeAllOfAttributes []OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesExcludeAllOfAttribute `json:"exclude_all_of_attributes,omitzero"`
	// An array of objects to add mandatory requirements for the order. A vehicle must
	// **meet** **all conditions** specified here to be considered for assignment. Each
	// object represents a single condition. Please note that a maximum of 10
	// conditions can be added here for a given order.
	RequiredAllOfAttributes []OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAllOfAttribute `json:"required_all_of_attributes,omitzero"`
	// An array of objects to add optional requirements for the order. A vehicle must
	// **meet** **at least one of the conditions** specified here to be considered for
	// assignment. Each object represents a single condition. Please note that a
	// maximum of 10 conditions can be added here for a given order.
	RequiredAnyOfAttributes []OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAnyOfAttribute `json:"required_any_of_attributes,omitzero"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferences) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Attribute, Operator, Value are required.
type OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesExcludeAllOfAttribute struct {
	// Specify the name of the attribute. The `attribute` is compared to the keys (of
	// each `key:value` pair) in `vehicles.attributes` during evaluation.
	Attribute string `json:"attribute,required"`
	// Specify the operator to denote the relation between `attribute` and the `value`
	// specified above. The `attribute` , `operator` and `value` together constitute
	// the condition that a vehicle must meet to be eligible for assignment. Currently,
	// we support following operators currently:
	//
	// - Equal to (==)
	// - Less than (<)
	// - Less tha equal to (<=)
	// - Greater than (>)
	// - Greater than equal to (>=)
	// - Contains (contains)
	//
	// Please note that when using "contains" operator only one `value` can be
	// specified and the corresponding `attribute` must contain multiple values when
	// defined for a vehicle.
	Operator string `json:"operator,required"`
	// Specify the desired value of the attribute to be applied for this order. `value`
	// provided here is compared to the values (of each `key:value` pair) in
	// `vehicles.attributes` during evaluation.
	Value string `json:"value,required"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesExcludeAllOfAttribute) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesExcludeAllOfAttribute
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesExcludeAllOfAttribute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Attribute, Operator, Value are required.
type OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAllOfAttribute struct {
	// Specify the name of the attribute. The `attribute` is compared to the keys (of
	// each `key:value` pair) in `vehicles.attributes` during evaluation.
	Attribute string `json:"attribute,required"`
	// Specify the operator to denote the relation between `attribute` and the `value`
	// specified above. The `attribute` , `operator` and `value` together constitute
	// the condition that a vehicle must meet to be eligible for assignment. Currently,
	// we support following operators currently:
	//
	// - Equal to (==)
	// - Less than (<)
	// - Less tha equal to (<=)
	// - Greater than (>)
	// - Greater than equal to (>=)
	// - Contains (contains)
	//
	// Please note that when using "contains" operator only one `value` can be
	// specified and the corresponding `attribute` must contain multiple values when
	// defined for a vehicle.
	Operator string `json:"operator,required"`
	// Specify the desired value of the attribute to be applied for this order. `value`
	// provided here is compared to the values (of each `key:value` pair) in
	// `vehicles.attributes` during evaluation.
	Value string `json:"value,required"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAllOfAttribute) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAllOfAttribute
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAllOfAttribute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Attribute, Operator, Value are required.
type OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAnyOfAttribute struct {
	// Specify the name of the attribute. The `attribute` is compared to the keys (of
	// each `key:value` pair) in `vehicles.attributes` during evaluation.
	Attribute string `json:"attribute,required"`
	// Specify the operator to denote the relation between `attribute` and the `value`
	// specified above. The `attribute` , `operator` and `value` together constitute
	// the condition that a vehicle must meet to be eligible for assignment. Currently,
	// we support following operators currently:
	//
	// - Equal to (==)
	// - Less than (<)
	// - Less tha equal to (<=)
	// - Greater than (>)
	// - Greater than equal to (>=)
	// - Contains (contains)
	//
	// Please note that when using "contains" operator only one `value` can be
	// specified and the corresponding `attribute` must contain multiple values when
	// defined for a vehicle.
	Operator string `json:"operator,required"`
	// Specify the desired value of the attribute to be applied for this order. `value`
	// provided here is compared to the values (of each `key:value` pair) in
	// `vehicles.attributes` during evaluation.
	Value string `json:"value,required"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAnyOfAttribute) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAnyOfAttribute
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAnyOfAttribute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configure the assignment constraints and response settings.
type OptimizationDriverAssignmentAssignDriversParamsOptions struct {
	// Specify the maximum number of potential, alternate vehicle assignments to be
	// returned for each order, apart from the vehicle which was assigned as
	// recommended. Please note that:
	//
	//   - The maximum number of alternate assignments that can be requested are 3.
	//   - It is not necessary that the service will return the specified number of
	//     alternate assignments for each order. The number of alternate assignments
	//     returned will depend on the number of vehicles provided in the input.
	//   - Order which could not be assigned to any vehicles due to their `filter` or
	//     attribute matching criteria will not be eligible for alternate assignments as
	//     well.
	AlternateAssignments param.Opt[int64] `json:"alternate_assignments,omitzero"`
	// When **true**, the service returns the drop-off steps for each trip and related
	// details in the response. Defaults to **false**.
	DropoffDetails param.Opt[bool] `json:"dropoff_details,omitzero"`
	// Collection of rules for assigning custom priority to orders based on their
	// attributes. In case an order satisfies more than one rule, the highest priority
	// score from all the rules satisfied, would be the effective priority score for
	// such an order.
	OrderAttributePriorityMappings []OptimizationDriverAssignmentAssignDriversParamsOptionsOrderAttributePriorityMapping `json:"order_attribute_priority_mappings,omitzero"`
	// Choose a travel cost that will be used by the service for assigning vehicles
	// efficiently from a set of qualifying ones.
	//
	// Any of "driving_eta", "driving_distance", "straight_line_distance".
	TravelCost string `json:"travel_cost,omitzero"`
	// Collection of rules for assigning custom priority to vehicles based on their
	// attributes. In case a vehicle satisfies more than one rule, the highest priority
	// score from all the rules satisfied, would be the effective priority score for
	// such a vehicle.
	VehicleAttributePriorityMappings []OptimizationDriverAssignmentAssignDriversParamsOptionsVehicleAttributePriorityMapping `json:"vehicle_attribute_priority_mappings,omitzero"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[OptimizationDriverAssignmentAssignDriversParamsOptions](
		"travel_cost", "driving_eta", "driving_distance", "straight_line_distance",
	)
}

// The properties Attribute, Operator, Priority, Value are required.
type OptimizationDriverAssignmentAssignDriversParamsOptionsOrderAttributePriorityMapping struct {
	// Specify the name of the attribute. The `attribute` is compared to the keys (of
	// each `key:value` pair) in `orders.attributes` during evaluation.
	Attribute string `json:"attribute,required"`
	// Specify the operator to denote the relation between `attribute` and the `value`
	// specified above. The `attribute` , `operator` and `value` together constitute
	// the condition that an order must meet to assume the specified priority. We
	// support the following operators currently:
	//
	// - Equal to (==)
	// - Less than (<)
	// - Less tha equal to (<=)
	// - Greater than (>)
	// - Greater than equal to (>=)
	// - Contains (contains)
	//
	// Please note that when using "contains" operator only one `value` can be
	// specified and the corresponding `attribute` must contain multiple values when
	// defined for an order.
	Operator string `json:"operator,required"`
	// Specify the priority score that should be assigned when an order qualifies the
	// criteria specified above. A higher value indicates a higher priority. Valid
	// values are \[1,10\].
	Priority string `json:"priority,required"`
	// Specify the desired value of the attribute to be applied for this order. `value`
	// provided here is compared to the values (of each `key:value` pair) in
	// `orders.attributes` during evaluation.
	Value string `json:"value,required"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsOptionsOrderAttributePriorityMapping) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsOptionsOrderAttributePriorityMapping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsOptionsOrderAttributePriorityMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Attribute, Operator, Priority, Value are required.
type OptimizationDriverAssignmentAssignDriversParamsOptionsVehicleAttributePriorityMapping struct {
	// Specify the name of the attribute. The `attribute` is compared to the keys (of
	// each `key:value` pair) in `vehicles.attributes` during evaluation.
	Attribute string `json:"attribute,required"`
	// Specify the operator to denote the relation between `attribute` and the `value`
	// specified above. The `attribute` , `operator` and `value` together constitute
	// the condition that a vehicle must meet to assume the specified priority. We
	// support the following operators currently:
	//
	// - Equal to (==)
	// - Less than (<)
	// - Less tha equal to (<=)
	// - Greater than (>)
	// - Greater than equal to (>=)
	// - Contains (contains)
	//
	// Please note that when using "contains" operator only one `value` can be
	// specified and the corresponding `attribute` must contain multiple values when
	// defined for a vehicle.
	Operator string `json:"operator,required"`
	// Specify the priority score that should be assigned when a vehicle qualifies the
	// criteria specified above. A higher value indicates a higher priority. Valid
	// values are \[1,10\].
	Priority string `json:"priority,required"`
	// Specify the desired value of the attribute to be applied for this vehicle.
	// `value` provided here is compared to the values (of each `key:value` pair) in
	// `vehicles.attributes` during evaluation.
	Value string `json:"value,required"`
	paramObj
}

func (r OptimizationDriverAssignmentAssignDriversParamsOptionsVehicleAttributePriorityMapping) MarshalJSON() (data []byte, err error) {
	type shadow OptimizationDriverAssignmentAssignDriversParamsOptionsVehicleAttributePriorityMapping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizationDriverAssignmentAssignDriversParamsOptionsVehicleAttributePriorityMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
