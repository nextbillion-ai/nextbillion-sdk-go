// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apijson"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/nextbillion-sdk-go/option"
	"github.com/stainless-sdks/nextbillion-sdk-go/packages/param"
	"github.com/stainless-sdks/nextbillion-sdk-go/packages/respjson"
)

// FleetifyRouteService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFleetifyRouteService] method instead.
type FleetifyRouteService struct {
	Options []option.RequestOption
	Steps   FleetifyRouteStepService
}

// NewFleetifyRouteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFleetifyRouteService(opts ...option.RequestOption) (r FleetifyRouteService) {
	r = FleetifyRouteService{}
	r.Options = opts
	r.Steps = NewFleetifyRouteStepService(opts...)
	return
}

// Dispatch a new route
func (r *FleetifyRouteService) New(ctx context.Context, params FleetifyRouteNewParams, opts ...option.RequestOption) (res *FleetifyRouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "fleetify/routes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Re-dispatch route
func (r *FleetifyRouteService) Redispatch(ctx context.Context, routeID string, params FleetifyRouteRedispatchParams, opts ...option.RequestOption) (res *FleetifyRouteRedispatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if routeID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	path := fmt.Sprintf("fleetify/routes/%s/redispatch", routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// An object returning the routing characteristics that are used to generate the
// route and turn-by-turn navigation steps for the dispatched route. The route and
// navigation steps are available when driver uses the in-app navigation.
//
// Please note the routing characteristics returned here are the same as those
// configured in the input request. The fields which were not specified in the
// input will be returned as blanks.
type Routing struct {
	// Returns the configuration of approaches for each step, that is used when
	// generating the route to help the driver with turn-by-turn navigation.
	Approaches string `json:"approaches"`
	// Returns the objects and maneuvers that will be avoided in the route that is
	// built when driver starts the in-app turn-by-turn navigation.
	Avoid string `json:"avoid"`
	// Returns the hazardous cargo type that the truck is carrying. The hazardous cargo
	// type is used to determine the compliant routes that the driver can take while
	// navigating the dispatched route.
	HazmatType string `json:"hazmat_type"`
	// Returns the driving mode that is used to build the route when driver starts the
	// in-app turn-by-turn navigation.
	Mode string `json:"mode"`
	// Returns the total load per axle of the truck, in tonnes, used to determine
	// compliant routes that the driver can take when he starts navigating the
	// dispatched route.
	TruckAxleLoad string `json:"truck_axle_load"`
	// Returns the truck dimensions, in centimeters, used to determine compliant routes
	// that the driver can take when he starts navigating the dispatched route.
	TruckSize string `json:"truck_size"`
	// Returns the truck weight that will determine compliant routes that can be used
	// by the driver during navigation.
	TruckWeight string `json:"truck_weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Approaches    respjson.Field
		Avoid         respjson.Field
		HazmatType    respjson.Field
		Mode          respjson.Field
		TruckAxleLoad respjson.Field
		TruckSize     respjson.Field
		TruckWeight   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Routing) RawJSON() string { return r.JSON.raw }
func (r *Routing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyRouteNewResponse struct {
	// An array of objects containing the details of each step in the dispatched route.
	// Each object represents a single step.
	Data FleetifyRouteNewResponseData `json:"data"`
	// Returns the status code of the response.
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyRouteNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An array of objects containing the details of each step in the dispatched route.
// Each object represents a single step.
type FleetifyRouteNewResponseData struct {
	// Returns the unique ID of the dispatched route.
	ID string `json:"id"`
	// Returns the UNIX timestamp, in seconds precision, at which this route dispatch
	// request was created.
	CreatedAt int64 `json:"created_at"`
	// Returns the total route distance, in meters, for informative display in the
	// driver app. It is the same as the value provided for `distance` field in the
	// input request.
	Distance int64 `json:"distance"`
	// Returns the details of the document that was specified in the input for
	// collecting the proof-of-completion for all steps in the dispatched routes. Each
	// object represents a new field in the document.
	DocumentSnapshot []any `json:"document_snapshot"`
	// An object returning the details of the driver to whom the route was dispatched.
	Driver FleetifyRouteNewResponseDataDriver `json:"driver"`
	// Returns the route optimization request ID which was used to dispatch the route.
	// An empty string is returned if the corresponding input was not provided.
	RoRequestID string `json:"ro_request_id"`
	// An object returning the routing characteristics that are used to generate the
	// route and turn-by-turn navigation steps for the dispatched route. The route and
	// navigation steps are available when driver uses the in-app navigation.
	//
	// Please note the routing characteristics returned here are the same as those
	// configured in the input request. The fields which were not specified in the
	// input will be returned as blanks.
	Routing Routing `json:"routing"`
	// Returns a shorter unique ID of the dispatched route for easier referencing and
	// displaying purposes.
	ShortID string `json:"short_id"`
	// An array of objects containing the details of all steps to be performed as part
	// of the dispatched route. Each object represents a single step during the route.
	Steps []RouteStepsResponse `json:"steps"`
	// Returns the total number of steps in the dispatched route.
	TotalSteps int64 `json:"total_steps"`
	// Returns the UNIX timestamp, in seconds precision, at which this route dispatch
	// request was updated.
	UpdatedAt int64 `json:"updated_at"`
	// Returns the ID of the vehicle to which the route was dispatched. The vehicle ID
	// returned here is the same as the one used in the route optimization request for
	// the given vehicle. An empty string is returned if the `ro_request_id` was not
	// provided in the input.
	VehicleID string `json:"vehicle_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Distance         respjson.Field
		DocumentSnapshot respjson.Field
		Driver           respjson.Field
		RoRequestID      respjson.Field
		Routing          respjson.Field
		ShortID          respjson.Field
		Steps            respjson.Field
		TotalSteps       respjson.Field
		UpdatedAt        respjson.Field
		VehicleID        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyRouteNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning the details of the driver to whom the route was dispatched.
type FleetifyRouteNewResponseDataDriver struct {
	// Returns the ID of the driver as specified in the
	// [NextBillion.ai](http://NextBillion.ai) Cloud Console.
	ID string `json:"id"`
	// Returns the email of the driver as specified in the
	// [NextBillion.ai](http://NextBillion.ai) Cloud Console.
	Email string `json:"email"`
	// Returns the full name of the driver as specified in
	// [NextBillion.ai](http://NextBillion.ai) Cloud Console.
	Fullname string `json:"fullname"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Fullname    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyRouteNewResponseDataDriver) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteNewResponseDataDriver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyRouteRedispatchResponse struct {
	// An array of objects containing the details of each step in the dispatched route.
	// Each object represents a single step.
	Data FleetifyRouteRedispatchResponseData `json:"data"`
	// Returns the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Message string `json:"message"`
	// Returns the status code of the response.
	Status int64 `json:"status"`
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
func (r FleetifyRouteRedispatchResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteRedispatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An array of objects containing the details of each step in the dispatched route.
// Each object represents a single step.
type FleetifyRouteRedispatchResponseData struct {
	// Returns the unique ID of the route.
	ID string `json:"id"`
	// Returns the number of steps already completed in the route.
	CompletedSteps int64 `json:"completed_steps"`
	// Returns the completion status of the route.
	Completion FleetifyRouteRedispatchResponseDataCompletion `json:"completion"`
	// Returns the UNIX timestamp, in seconds precision, at which this route dispatch
	// request was created.
	CreatedAt int64 `json:"created_at"`
	// Returns the total route distance, in meters, for informative display in the
	// driver app. It is the same as the value provided for `distance` field in the
	// input request.
	Distance int64 `json:"distance"`
	// Returns the details of the document that was specified in the input for
	// collecting the proof-of-completion for all steps in the dispatched routes. Each
	// object represents a new field in the document.
	DocumentSnapshot []any `json:"document_snapshot"`
	// An object returning the details of the driver to whom the route was dispatched.
	Driver FleetifyRouteRedispatchResponseDataDriver `json:"driver"`
	// Returns the route optimization request ID which was used to dispatch the route.
	// An empty string is returned if the corresponding input was not provided.
	RoRequestID string `json:"ro_request_id"`
	// An object returning the routing characteristics that are used to generate the
	// route and turn-by-turn navigation steps for the dispatched route. The route and
	// navigation steps are available when driver uses the in-app navigation.
	//
	// Please note the routing characteristics returned here are the same as those
	// configured in the input request. The fields which were not specified in the
	// input will be returned as blanks.
	Routing Routing `json:"routing"`
	// Returns a shorter unique ID of the route for easier referencing and displaying
	// purposes.
	ShortID string                                   `json:"short_id"`
	Steps   FleetifyRouteRedispatchResponseDataSteps `json:"steps"`
	// Returns the total number of steps in the dispatched route.
	TotalSteps int64 `json:"total_steps"`
	// Returns the UNIX timestamp, in seconds precision, at which this route dispatch
	// request was updated.
	UpdatedAt int64 `json:"updated_at"`
	// Returns the ID of the vehicle to which the route was dispatched. The vehicle ID
	// returned here is the same as the one used in the route optimization request for
	// the given vehicle. An empty string is returned if the `ro_request_id` was not
	// provided in the input.
	VehicleID string `json:"vehicle_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CompletedSteps   respjson.Field
		Completion       respjson.Field
		CreatedAt        respjson.Field
		Distance         respjson.Field
		DocumentSnapshot respjson.Field
		Driver           respjson.Field
		RoRequestID      respjson.Field
		Routing          respjson.Field
		ShortID          respjson.Field
		Steps            respjson.Field
		TotalSteps       respjson.Field
		UpdatedAt        respjson.Field
		VehicleID        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyRouteRedispatchResponseData) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteRedispatchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the completion status of the route.
type FleetifyRouteRedispatchResponseDataCompletion struct {
	// Returns the status of the route. It can take one of the following values -
	// "scheduled", "completed".
	//
	// Any of "`scheduled`", "`completed`".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyRouteRedispatchResponseDataCompletion) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteRedispatchResponseDataCompletion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning the details of the driver to whom the route was dispatched.
type FleetifyRouteRedispatchResponseDataDriver struct {
	// Returns the ID of the driver as specified in the
	// [NextBillion.ai](http://NextBillion.ai) Cloud Console.
	ID string `json:"id"`
	// Returns the email of the driver as specified in the
	// [NextBillion.ai](http://NextBillion.ai) Cloud Console.
	Email string `json:"email"`
	// Returns the full name of the driver as specified in
	// [NextBillion.ai](http://NextBillion.ai) Cloud Console.
	Fullname string `json:"fullname"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Fullname    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyRouteRedispatchResponseDataDriver) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteRedispatchResponseDataDriver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyRouteRedispatchResponseDataSteps struct {
	// Returns the unique ID of the step.
	ID string `json:"id"`
	// Returns the postal address where the step is executed. Its value is the same as
	// that specified in the input request when configuring the step details.
	Address string `json:"address"`
	// Returns the scheduled arrival time of the driver at the step as an UNIX
	// timestamp, in seconds precision. It is the same as that specified in the input
	// request while configuring the step details.
	//
	// The timestamp returned here is only for informative display on the driver's app
	// and it does not impact or get affected by the route generated.
	Arrival int64 `json:"arrival"`
	// Returns the completion status of the step.
	Completion FleetifyRouteRedispatchResponseDataStepsCompletion `json:"completion"`
	// Returns the UNIX timestamp, in seconds precision, at which this step was
	// created.
	CreatedAt int64 `json:"created_at"`
	// Returns the details of the document that was used for collecting the proof of
	// completion for the step. In case no document template ID was provided for the
	// given step, then a `null` value is returned. Each object represents a new field
	// in the document.
	DocumentSnapshot []any `json:"document_snapshot"`
	// Returns the duration for `layover` or `break` type steps.
	Duration int64 `json:"duration"`
	// Returns the location coordinates where the step is executed.
	Location []float64 `json:"location"`
	// An object returning custom details about the step that were configured in the
	// input request while configuring the step details. The information returned here
	// will be available for display on the Driver's app under step details.
	Meta FleetifyRouteRedispatchResponseDataStepsMeta `json:"meta"`
	// Returns a unique short ID of the step for easier referencing and displaying
	// purposes.
	ShortID string `json:"short_id"`
	// Returns the step type. It can belong to one of the following: `start`, `job` ,
	// `pickup`, `delivery`, `break`, `layover` , and `end`. For any given step, it
	// would be the same as that specified in the input request while configuring the
	// step details.
	Type string `json:"type"`
	// Returns the UNIX timestamp, in seconds precision, at which this step was last
	// updated.
	UpdatedAt int64 `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Address          respjson.Field
		Arrival          respjson.Field
		Completion       respjson.Field
		CreatedAt        respjson.Field
		DocumentSnapshot respjson.Field
		Duration         respjson.Field
		Location         respjson.Field
		Meta             respjson.Field
		ShortID          respjson.Field
		Type             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyRouteRedispatchResponseDataSteps) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteRedispatchResponseDataSteps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the completion status of the step.
type FleetifyRouteRedispatchResponseDataStepsCompletion struct {
	// Returns the status of the step. It can take one of the following values -
	// "scheduled", "completed".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyRouteRedispatchResponseDataStepsCompletion) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteRedispatchResponseDataStepsCompletion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning custom details about the step that were configured in the
// input request while configuring the step details. The information returned here
// will be available for display on the Driver's app under step details.
type FleetifyRouteRedispatchResponseDataStepsMeta struct {
	// Returns the customer name associated with the step. It can configured in the
	// input request using the `metadata` attribute of the step.
	CustomerName string `json:"customer_name"`
	// Returns the customer's phone number associated with the step. It can configured
	// in the input request using the `metadata` attribute of the step.
	CustomerPhoneNumber string `json:"customer_phone_number"`
	// Returns the custom instructions to carry out while performing the task. These
	// instructions can be provided at the time of configuring the step details in the
	// input request.
	Instructions string `json:"instructions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerName        respjson.Field
		CustomerPhoneNumber respjson.Field
		Instructions        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyRouteRedispatchResponseDataStepsMeta) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteRedispatchResponseDataStepsMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyRouteNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	// Specify the e-mail address of the driver who should receive the route. The
	// e-mail address must be registered in
	// [NextBillion.ai Cloud Console](https://console.nextbillion.ai/).
	DriverEmail string `json:"driver_email,required"`
	// An array of objects to collect the details about the intermediate steps in the
	// route to be dispatched. Each object corresponds to a single step. The array must
	// begin with a start-type step and end with an end-type step, to form a valid
	// route.
	Steps []RouteStepsRequestParam `json:"steps,omitzero,required"`
	// Specify the total distance, in meters, for an informative display in Driver's
	// app. The distance specified here has no effect on the actual route that the
	// service generates.
	Distance param.Opt[int64] `json:"distance,omitzero"`
	// Specify the ID of the document template that should be used to collect proof of
	// completion for all steps in the route. In order to complete each route step, the
	// driver will need to submit a form generated by the rules defined in the given
	// document template. Use the
	// [Documents API](https://docs.nextbillion.ai/docs/dispatches/documents-api) to
	// create, read and manage document templates.
	//
	// Please note that the document template ID assigned to a route does not apply to
	// following step types - `start`, `end`, `break`, `layover`.
	DocumentTemplateID param.Opt[string] `json:"document_template_id,omitzero"`
	// Specify the Route Optimization request ID. When this ID is provided, all other
	// fields will be ignored (including the required fields) and the route
	// optimization result will be used to form the routes and corresponding steps.
	//
	// Please note that:
	//
	//   - The driver's email ID must be provided in input `vehicle.metadata` as
	//     `user_email` such that the route optimization result must contain a valid
	//     driver email, step's arrival time, etc., to make a successful dispatch.
	//   - Document Template for collecting proof of delivery or completion can not be
	//     specified when using this field to dispatch a route.
	//   - In case of an error at any part among the routes, the API will immediately
	//     return the error with the index of the specific route or route step.
	//   - On a successful dispatch, the API returns the last route, if there are many,
	//     in the response payload.
	RoRequestID param.Opt[string] `json:"ro_request_id,omitzero"`
	// The `routing` object allows defining the routing characteristics that should be
	// used to generate a route when the Driver uses the in-app navigation. Only `car`
	// mode is supported currently.
	Routing FleetifyRouteNewParamsRouting `json:"routing,omitzero"`
	paramObj
}

func (r FleetifyRouteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyRouteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyRouteNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FleetifyRouteNewParams]'s query parameters as `url.Values`.
func (r FleetifyRouteNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The `routing` object allows defining the routing characteristics that should be
// used to generate a route when the Driver uses the in-app navigation. Only `car`
// mode is supported currently.
type FleetifyRouteNewParamsRouting struct {
	// Specify the total load per axle (including the weight of trailers and shipped
	// goods) of the truck, in tonnes. When specified, the dispatched route uses only
	// those roads which can be used by a truck to carry the specified load per axle.
	//
	// Please note this parameter is effective only when `mode=truck`.
	TruckAxleLoad param.Opt[int64] `json:"truck_axle_load,omitzero"`
	// Specify the dimensions of a truck, in centimeters (cm), in the format of
	// <height, width, length>. When specified, the dispatched route uses only those
	// roads which allow trucks with specified dimensions.
	//
	// Please note this parameter is effective only when `mode=truck`. Also, the
	// maximum dimensions that can be specified are as follows:
	//
	// Height = 1000 cm
	// Width = 5000 cm
	// Length = 5000 cm
	TruckSize param.Opt[string] `json:"truck_size,omitzero"`
	// Specify the weight of the truck, including trailers and shipped goods, in
	// kilograms (kg). When specified, the dispatched route uses only those roads which
	// allow trucks with specified weight.
	//
	// Please note this parameter is effective only when `mode=truck`. Also, the
	// maximum weight that can be specified for a truck is 100,000 kgs.
	TruckWeight param.Opt[int64] `json:"truck_weight,omitzero"`
	// Specify the side of the road from which the route should approach the step
	// location. When set to `unrestricted` a route can arrive at the step location
	// from either side of the road and when set to `curb` the route will arrive at the
	// step location only from the driving side of the region. Use a semi-colon `;` to
	// specify approach configurations for multiple steps.
	//
	// Any of "`unrestricted`", "`curb`".
	Approaches string `json:"approaches,omitzero"`
	// Setting this will ensure the generated route avoids the object(s) specified in
	// the input. Multiple values should be separated by a pipe (|). If `none` is
	// provided along with other values, an error is returned as a valid route is not
	// feasible.
	//
	// Any of "`toll`", "`highway`", "`ferry`", "`sharp_turn`", "`uturn`",
	// "`left_turn`", "`right_turn`", "`service_road`", "`none`".
	Avoid string `json:"avoid,omitzero"`
	// Specify the type of hazardous material being carried and the dispatch service
	// will avoid roads which are not suitable for the type of goods specified.
	// Multiple values can be separated using a pipe operator `|` .
	//
	// Please note that this parameter is effective only when `mode=truck`.
	//
	// Any of "`general`", "`circumstantial`", "`explosive`", "`harmful_to_water`".
	HazmatType string `json:"hazmat_type,omitzero"`
	// Specify the driving mode that the service should use to determine a route
	//
	// Any of "`car`".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r FleetifyRouteNewParamsRouting) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyRouteNewParamsRouting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyRouteNewParamsRouting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FleetifyRouteNewParamsRouting](
		"approaches", "`unrestricted`", "`curb`",
	)
	apijson.RegisterFieldValidator[FleetifyRouteNewParamsRouting](
		"avoid", "`toll`", "`highway`", "`ferry`", "`sharp_turn`", "`uturn`", "`left_turn`", "`right_turn`", "`service_road`", "`none`",
	)
	apijson.RegisterFieldValidator[FleetifyRouteNewParamsRouting](
		"hazmat_type", "`general`", "`circumstantial`", "`explosive`", "`harmful_to_water`",
	)
	apijson.RegisterFieldValidator[FleetifyRouteNewParamsRouting](
		"mode", "`car`",
	)
}

type FleetifyRouteRedispatchParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	// A collection of objects with details of the steps to be modified. Each object
	// corresponds to a single step.
	Operations []FleetifyRouteRedispatchParamsOperation `json:"operations,omitzero,required"`
	// Specify the distance of the route.
	Distance param.Opt[float64] `json:"distance,omitzero"`
	paramObj
}

func (r FleetifyRouteRedispatchParams) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyRouteRedispatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyRouteRedispatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FleetifyRouteRedispatchParams]'s query parameters as
// `url.Values`.
func (r FleetifyRouteRedispatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The properties Data, Operation are required.
type FleetifyRouteRedispatchParamsOperation struct {
	Data FleetifyRouteRedispatchParamsOperationData `json:"data,omitzero,required"`
	// Specify the type of operation to be performed for the step.
	//
	// Any of "`create`", "`update`", "`delete`".
	Operation string `json:"operation,omitzero,required"`
	paramObj
}

func (r FleetifyRouteRedispatchParamsOperation) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyRouteRedispatchParamsOperation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyRouteRedispatchParamsOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FleetifyRouteRedispatchParamsOperation](
		"operation", "`create`", "`update`", "`delete`",
	)
}

type FleetifyRouteRedispatchParamsOperationData struct {
	// Specify the ID of the document template to be used for collecting proof of
	// completion for the step. It would be applied to step which not be bind to
	// document template. Use the
	// [Documents API](https://docs.nextbillion.ai/docs/dispatches/documents-api) to
	// create, read and manage the document templates.
	//
	// Please note that the document template ID can not be assigned to following step
	// types - `start`, `end`, `break`, `layover`.
	DocumentTemplateID param.Opt[string] `json:"document_template_id,omitzero"`
	// Specify the ID of the step to be updated or deleted. Either one of `id` or
	// `short_id` of the step can be provided. This input will be ignored when
	// `operation: create` .
	StepID param.Opt[string] `json:"step_id,omitzero"`
	// Specify the mode of completion to be used for the step. Currently, following
	// values are allowed:
	//
	//   - `manual`: Steps must be marked as completed manually through the Driver App.
	//   - `geofence`: Steps are marked as completed automatically based on the entry
	//     conditions and geofence specified.
	//   - `geofence_manual_fallback`: Steps will be marked as completed automatically
	//     based on geofence and entry condition configurations but there will also be a
	//     provision for manually updating the status in case, geofence detection fails.
	//
	// Any of "`manual`", "`geofence`", "`geofence_manual_fallback`".
	CompletionMode RouteStepCompletionMode `json:"completion_mode,omitzero"`
	Step           RouteStepsRequestParam  `json:"step,omitzero"`
	paramObj
}

func (r FleetifyRouteRedispatchParamsOperationData) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyRouteRedispatchParamsOperationData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyRouteRedispatchParamsOperationData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
