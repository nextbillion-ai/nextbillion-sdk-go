// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"encoding/json"
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

// FleetifyRouteStepService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFleetifyRouteStepService] method instead.
type FleetifyRouteStepService struct {
	Options []option.RequestOption
}

// NewFleetifyRouteStepService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFleetifyRouteStepService(opts ...option.RequestOption) (r FleetifyRouteStepService) {
	r = FleetifyRouteStepService{}
	r.Options = opts
	return
}

// Insert a new step
func (r *FleetifyRouteStepService) New(ctx context.Context, routeID string, params FleetifyRouteStepNewParams, opts ...option.RequestOption) (res *FleetifyRouteStepNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if routeID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	path := fmt.Sprintf("fleetify/routes/%s/steps", routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get step data
func (r *FleetifyRouteStepService) Get(ctx context.Context, stepID string, params FleetifyRouteStepGetParams, opts ...option.RequestOption) (res *FleetifyRouteStepGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.RouteID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	if stepID == "" {
		err = errors.New("missing required stepID parameter")
		return
	}
	path := fmt.Sprintf("fleetify/routes/%s/steps/%s", params.RouteID, stepID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update a step
func (r *FleetifyRouteStepService) Update(ctx context.Context, stepsID string, params FleetifyRouteStepUpdateParams, opts ...option.RequestOption) (res *FleetifyRouteStepUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.RouteID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	if stepsID == "" {
		err = errors.New("missing required stepsID parameter")
		return
	}
	path := fmt.Sprintf("fleetify/routes/%s/steps/%s", params.RouteID, stepsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete a step
func (r *FleetifyRouteStepService) Delete(ctx context.Context, stepsID string, params FleetifyRouteStepDeleteParams, opts ...option.RequestOption) (res *FleetifyRouteStepDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.RouteID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	if stepsID == "" {
		err = errors.New("missing required stepsID parameter")
		return
	}
	path := fmt.Sprintf("fleetify/routes/%s/steps/%s", params.RouteID, stepsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Complete a route step with document submission, or update the document of a
// completed route step.
//
// When all steps are completed, the encapsulating route’s status will change to
// `completed` automatically.
//
// Either `Session Token` must be provided to authenticate the request.
func (r *FleetifyRouteStepService) Complete(ctx context.Context, stepsID string, params FleetifyRouteStepCompleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.RouteID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	if stepsID == "" {
		err = errors.New("missing required stepsID parameter")
		return
	}
	path := fmt.Sprintf("fleetify/routes/%s/steps/%s", params.RouteID, stepsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return
}

type DocumentSubmission = any

// Specify the mode of completion to be used for the step. Currently, following
// values are allowed:
//
//   - `manual`: Steps must be marked as completed manually through the Driver App.
//   - `geofence`: Steps are marked as completed automatically based on the entry
//     conditions and geofence specified.
//   - `geofence_manual_fallback`: Steps will be marked as completed automatically
//     based on geofence and entry condition configurations but there will also be a
//     provision for manually updating the status in case, geofence detection fails.
type RouteStepCompletionMode string

const (
	RouteStepCompletionModeManual                 RouteStepCompletionMode = "`manual`"
	RouteStepCompletionModeGeofence               RouteStepCompletionMode = "`geofence`"
	RouteStepCompletionModeGeofenceManualFallback RouteStepCompletionMode = "`geofence_manual_fallback`"
)

// Specify the configurations of the geofence which will be used to detect presence
// of the driver and complete the tasks automatically. Please note that this
// attribute is required when `completion_mode` is either "geofence" or
// "geofence_manual_fallback".
type RouteStepGeofenceConfig struct {
	// Specify the radius of the cicular geofence, in meters. Once specified, the
	// service will create a geofence with task's location as the center of the circle
	// having the given radius. Valid values for `radius` are \[10, 5000\].
	Radius float64 `json:"radius"`
	// Specify the type of the geofence. Currently, `circle` is the only suppoeted
	// value.
	//
	// Any of "circle".
	Type RouteStepGeofenceConfigType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Radius      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteStepGeofenceConfig) RawJSON() string { return r.JSON.raw }
func (r *RouteStepGeofenceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RouteStepGeofenceConfig to a RouteStepGeofenceConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RouteStepGeofenceConfigParam.Overrides()
func (r RouteStepGeofenceConfig) ToParam() RouteStepGeofenceConfigParam {
	return param.Override[RouteStepGeofenceConfigParam](json.RawMessage(r.RawJSON()))
}

// Specify the type of the geofence. Currently, `circle` is the only suppoeted
// value.
type RouteStepGeofenceConfigType string

const (
	RouteStepGeofenceConfigTypeCircle RouteStepGeofenceConfigType = "circle"
)

// Specify the configurations of the geofence which will be used to detect presence
// of the driver and complete the tasks automatically. Please note that this
// attribute is required when `completion_mode` is either "geofence" or
// "geofence_manual_fallback".
type RouteStepGeofenceConfigParam struct {
	// Specify the radius of the cicular geofence, in meters. Once specified, the
	// service will create a geofence with task's location as the center of the circle
	// having the given radius. Valid values for `radius` are \[10, 5000\].
	Radius param.Opt[float64] `json:"radius,omitzero"`
	// Specify the type of the geofence. Currently, `circle` is the only suppoeted
	// value.
	//
	// Any of "circle".
	Type RouteStepGeofenceConfigType `json:"type,omitzero"`
	paramObj
}

func (r RouteStepGeofenceConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow RouteStepGeofenceConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RouteStepGeofenceConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Arrival, Location, Type are required.
type RouteStepsRequestParam struct {
	// Specify the scheduled arrival time of the driver, as an UNIX timestamp in
	// seconds, at the step. Please note that:
	//
	//   - Arrival time for each step should be equal to or greater than the previous
	//     step.
	//   - Past times can not be provided.
	//   - The time provided is used only for informative display on the driver app and
	//     it does not impact or get affected by the route generated.
	Arrival int64 `json:"arrival,required"`
	// Specify the location coordinates where the steps should be performed in
	// `[latitude, longitude]`.
	Location []float64 `json:"location,omitzero,required"`
	// Specify the step type. It can belong to one of the following: `start`, `job` ,
	// `pickup`, `delivery`, `end`. A `duration` is mandatory when the step type is
	// either `layover` or a `break`.
	//
	// Any of "`start`", "`job`", "`pickup`", "`delivery`", "`break`", "`layover`",
	// "`end`".
	Type RouteStepsRequestType `json:"type,omitzero,required"`
	// Specify the postal address for the step.
	Address param.Opt[string] `json:"address,omitzero"`
	// Specify the ID of the document template to be used for collecting proof of
	// completion for the step. If not specified, the document template specified at
	// the route level will be used for the step. Use the
	// [Documents API](https://docs.nextbillion.ai/docs/dispatches/documents-api) to
	// create, read and manage the document templates.
	//
	// Please note that the document template ID can not be assigned to following step
	// types - `start`, `end`, `break`, `layover`.
	DocumentTemplateID param.Opt[string] `json:"document_template_id,omitzero"`
	// Specify the duration of the `layover` or `break` type steps, in seconds. Please
	// note it is mandatory when step type is either "layover" or "break".
	Duration param.Opt[int64] `json:"duration,omitzero"`
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
	// Specify the configurations of the geofence which will be used to detect presence
	// of the driver and complete the tasks automatically. Please note that this
	// attribute is required when `completion_mode` is either "geofence" or
	// "geofence_manual_fallback".
	GeofenceConfig RouteStepGeofenceConfigParam `json:"geofence_config,omitzero"`
	// An object to specify any additional details about the task to be associated with
	// the step in the response. The information provided here will be available on the
	// Driver's app under step details. This attribute can be used to provide context
	// about or instructions to the driver for performing the task
	Meta RouteStepsRequestMetaParam `json:"meta,omitzero"`
	paramObj
}

func (r RouteStepsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RouteStepsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RouteStepsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specify the step type. It can belong to one of the following: `start`, `job` ,
// `pickup`, `delivery`, `end`. A `duration` is mandatory when the step type is
// either `layover` or a `break`.
type RouteStepsRequestType string

const (
	RouteStepsRequestTypeStart    RouteStepsRequestType = "`start`"
	RouteStepsRequestTypeJob      RouteStepsRequestType = "`job`"
	RouteStepsRequestTypePickup   RouteStepsRequestType = "`pickup`"
	RouteStepsRequestTypeDelivery RouteStepsRequestType = "`delivery`"
	RouteStepsRequestTypeBreak    RouteStepsRequestType = "`break`"
	RouteStepsRequestTypeLayover  RouteStepsRequestType = "`layover`"
	RouteStepsRequestTypeEnd      RouteStepsRequestType = "`end`"
)

// An object to specify any additional details about the task to be associated with
// the step in the response. The information provided here will be available on the
// Driver's app under step details. This attribute can be used to provide context
// about or instructions to the driver for performing the task
type RouteStepsRequestMetaParam struct {
	// Specify the name of the customer for which the step has to be performed.
	CustomerName param.Opt[string] `json:"customer_name,omitzero"`
	// Specify the phone number of the person to be contacted when at step location.
	CustomerPhoneNumber param.Opt[string] `json:"customer_phone_number,omitzero"`
	// Specify custom instructions to be carried out while performing the step.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r RouteStepsRequestMetaParam) MarshalJSON() (data []byte, err error) {
	type shadow RouteStepsRequestMetaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RouteStepsRequestMetaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteStepsResponse struct {
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
	Arrival    int64                        `json:"arrival"`
	Completion RouteStepsResponseCompletion `json:"completion"`
	// Represents the timestamp of the creation in seconds since the Unix epoch.
	// Example: `1738743999`.
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
	Meta RouteStepsResponseMeta `json:"meta"`
	// Returns a unique short ID of the step for easier referencing and displaying
	// purposes.
	ShortID string `json:"short_id"`
	// Returns the step type. It can belong to one of the following: `start`, `job` ,
	// `pickup`, `delivery`, `break`, `layover` , and `end`. For any given step, it
	// would be the same as that specified in the input request while configuring the
	// step details.
	Type string `json:"type"`
	// Represents the timestamp of the last update in seconds since the Unix epoch.
	// Example: `1738743999`.
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
func (r RouteStepsResponse) RawJSON() string { return r.JSON.raw }
func (r *RouteStepsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteStepsResponseCompletion struct {
	// Represents the timestamp of the completion in seconds since the Unix epoch.
	// Example: `1738743999`.
	CompletedAt int64 `json:"completed_at"`
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
	CompletedByMode RouteStepCompletionMode `json:"completed_by_mode"`
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
	CompletionMode RouteStepCompletionMode `json:"completion_mode"`
	// A key-value map storing form submission data, where keys correspond to field
	// labels and values can be of any type depend on the type of according document
	// item.
	Document DocumentSubmission `json:"document"`
	// Represents the timestamp of the last doc modification in seconds since the Unix
	// epoch. Example: `1738743999`.
	DocumentModifiedAt int64 `json:"document_modified_at"`
	// Specify the configurations of the geofence which will be used to detect presence
	// of the driver and complete the tasks automatically. Please note that this
	// attribute is required when `completion_mode` is either "geofence" or
	// "geofence_manual_fallback".
	GeofenceConfig RouteStepGeofenceConfig `json:"geofence_config"`
	// Status of the step.
	//
	// Any of "scheduled", "completed", "canceled".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt        respjson.Field
		CompletedByMode    respjson.Field
		CompletionMode     respjson.Field
		Document           respjson.Field
		DocumentModifiedAt respjson.Field
		GeofenceConfig     respjson.Field
		Status             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteStepsResponseCompletion) RawJSON() string { return r.JSON.raw }
func (r *RouteStepsResponseCompletion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning custom details about the step that were configured in the
// input request while configuring the step details. The information returned here
// will be available for display on the Driver's app under step details.
type RouteStepsResponseMeta struct {
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
func (r RouteStepsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *RouteStepsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyRouteStepNewResponse struct {
	Data RouteStepsResponse `json:"data"`
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
func (r FleetifyRouteStepNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteStepNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyRouteStepGetResponse struct {
	Data RouteStepsResponse `json:"data"`
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
func (r FleetifyRouteStepGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteStepGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyRouteStepUpdateResponse struct {
	Data RouteStepsResponse `json:"data"`
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
func (r FleetifyRouteStepUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteStepUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyRouteStepDeleteResponse struct {
	// Returns the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Message string `json:"message"`
	// Returns the status code of the response.
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyRouteStepDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyRouteStepDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyRouteStepNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	// Specify the scheduled arrival time of the driver, as an UNIX timestamp in
	// seconds, at the step. Please note that:
	//
	//   - Arrival time for each step should be equal to or greater than the previous
	//     step.
	//   - Past times can not be provided.
	//   - The time provided is used only for informative display on the driver app and
	//     it does not impact or get affected by the route generated.
	Arrival int64 `json:"arrival,required"`
	// Specify the location coordinates where the steps should be performed in
	// `[latitude, longitude]`.
	Location []float64 `json:"location,omitzero,required"`
	// Indicates the index at which to insert the step, starting from 0 up to the total
	// number of steps in the route.
	Position int64 `json:"position,required"`
	// Specify the step type. It can belong to one of the following: `start`, `job` ,
	// `pickup`, `delivery`, `end`. A `duration` is mandatory when the step type is
	// either `layover` or a `break`.
	//
	// Any of "`start`", "`job`", "`pickup`", "`delivery`", "`break`", "`layover`",
	// "`end`".
	Type FleetifyRouteStepNewParamsType `json:"type,omitzero,required"`
	// Specify the postal address for the step.
	Address param.Opt[string] `json:"address,omitzero"`
	// Specify the ID of the document template to be used for collecting proof of
	// completion for the step. If not specified, the document template specified at
	// the route level will be used for the step. Use the
	// [Documents API](https://docs.nextbillion.ai/docs/dispatches/documents-api) to
	// create, read and manage the document templates.
	//
	// Please note that the document template ID can not be assigned to following step
	// types - `start`, `end`, `break`, `layover`.
	DocumentTemplateID param.Opt[string] `json:"document_template_id,omitzero"`
	// Specify the duration of the `layover` or `break` type steps, in seconds. Please
	// note it is mandatory when step type is either "layover" or "break".
	Duration param.Opt[int64] `json:"duration,omitzero"`
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
	// Specify the configurations of the geofence which will be used to detect presence
	// of the driver and complete the tasks automatically. Please note that this
	// attribute is required when `completion_mode` is either "geofence" or
	// "geofence_manual_fallback".
	GeofenceConfig RouteStepGeofenceConfigParam `json:"geofence_config,omitzero"`
	// An object to specify any additional details about the task to be associated with
	// the step in the response. The information provided here will be available on the
	// Driver's app under step details. This attribute can be used to provide context
	// about or instructions to the driver for performing the task
	Meta FleetifyRouteStepNewParamsMeta `json:"meta,omitzero"`
	paramObj
}

func (r FleetifyRouteStepNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyRouteStepNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyRouteStepNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FleetifyRouteStepNewParams]'s query parameters as
// `url.Values`.
func (r FleetifyRouteStepNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the step type. It can belong to one of the following: `start`, `job` ,
// `pickup`, `delivery`, `end`. A `duration` is mandatory when the step type is
// either `layover` or a `break`.
type FleetifyRouteStepNewParamsType string

const (
	FleetifyRouteStepNewParamsTypeStart    FleetifyRouteStepNewParamsType = "`start`"
	FleetifyRouteStepNewParamsTypeJob      FleetifyRouteStepNewParamsType = "`job`"
	FleetifyRouteStepNewParamsTypePickup   FleetifyRouteStepNewParamsType = "`pickup`"
	FleetifyRouteStepNewParamsTypeDelivery FleetifyRouteStepNewParamsType = "`delivery`"
	FleetifyRouteStepNewParamsTypeBreak    FleetifyRouteStepNewParamsType = "`break`"
	FleetifyRouteStepNewParamsTypeLayover  FleetifyRouteStepNewParamsType = "`layover`"
	FleetifyRouteStepNewParamsTypeEnd      FleetifyRouteStepNewParamsType = "`end`"
)

// An object to specify any additional details about the task to be associated with
// the step in the response. The information provided here will be available on the
// Driver's app under step details. This attribute can be used to provide context
// about or instructions to the driver for performing the task
type FleetifyRouteStepNewParamsMeta struct {
	// Specify the name of the customer for which the step has to be performed.
	CustomerName param.Opt[string] `json:"customer_name,omitzero"`
	// Specify the phone number of the person to be contacted when at step location.
	CustomerPhoneNumber param.Opt[string] `json:"customer_phone_number,omitzero"`
	// Specify custom instructions to be carried out while performing the step.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r FleetifyRouteStepNewParamsMeta) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyRouteStepNewParamsMeta
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyRouteStepNewParamsMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyRouteStepGetParams struct {
	RouteID string `path:"routeID,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	paramObj
}

// URLQuery serializes [FleetifyRouteStepGetParams]'s query parameters as
// `url.Values`.
func (r FleetifyRouteStepGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FleetifyRouteStepUpdateParams struct {
	RouteID string `path:"routeID,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	// Specify the scheduled arrival time of the driver, as an UNIX timestamp in
	// seconds, at the step. Please note that:
	//
	//   - Arrival time for each step should be equal to or greater than the previous
	//     step.
	//   - Past times can not be provided.
	//   - The time provided is used only for informative display on the driver app and
	//     it does not impact or get affected by the route generated.
	Arrival int64 `json:"arrival,required"`
	// Specify the new position of the step. Provide a position different from the
	// current position of the step to update sequence in which the step get completed.
	Position int64 `json:"position,required"`
	// Specify the postal address for the step.
	Address param.Opt[string] `json:"address,omitzero"`
	// Update the ID of the document template to be used for collecting proof of
	// completion for the step. If an empty string "" is provided, the current document
	// template associated to the step will be removed.
	DocumentTemplateID param.Opt[string] `json:"document_template_id,omitzero"`
	// Specify the duration of the `layover` or `break` type steps, in seconds. Please
	// note it is mandatory when step type is either "layover" or "break".
	Duration param.Opt[int64] `json:"duration,omitzero"`
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
	// Specify the configurations of the geofence which will be used to detect presence
	// of the driver and complete the tasks automatically. Please note that this
	// attribute is required when `completion_mode` is either "geofence" or
	// "geofence_manual_fallback".
	GeofenceConfig RouteStepGeofenceConfigParam `json:"geofence_config,omitzero"`
	// Specify the location coordinates where the steps should be performed in
	// `[latitude, longitude]`.
	Location []float64 `json:"location,omitzero"`
	// An object to specify any additional details about the task to be associated with
	// the step in the response. The information provided here will be available on the
	// Driver's app under step details. This attribute can be used to provide context
	// about or instructions to the driver for performing the task
	Meta FleetifyRouteStepUpdateParamsMeta `json:"meta,omitzero"`
	// Specify the step type. It can belong to one of the following: `start`, `job` ,
	// `pickup`, `delivery`, `end`. A `duration` is mandatory when the step type is
	// either `layover` or a `break`.
	//
	// Any of "`start`", "`job`", "`pickup`", "`delivery`", "`break`", "`layover`",
	// "`end`".
	Type FleetifyRouteStepUpdateParamsType `json:"type,omitzero"`
	paramObj
}

func (r FleetifyRouteStepUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyRouteStepUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyRouteStepUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FleetifyRouteStepUpdateParams]'s query parameters as
// `url.Values`.
func (r FleetifyRouteStepUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// An object to specify any additional details about the task to be associated with
// the step in the response. The information provided here will be available on the
// Driver's app under step details. This attribute can be used to provide context
// about or instructions to the driver for performing the task
type FleetifyRouteStepUpdateParamsMeta struct {
	// Specify the name of the customer for which the step has to be performed.
	CustomerName param.Opt[string] `json:"customer_name,omitzero"`
	// Specify the phone number of the person to be contacted when at step location.
	CustomerPhoneNumber param.Opt[string] `json:"customer_phone_number,omitzero"`
	// Specify custom instructions to be carried out while performing the step.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r FleetifyRouteStepUpdateParamsMeta) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyRouteStepUpdateParamsMeta
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyRouteStepUpdateParamsMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specify the step type. It can belong to one of the following: `start`, `job` ,
// `pickup`, `delivery`, `end`. A `duration` is mandatory when the step type is
// either `layover` or a `break`.
type FleetifyRouteStepUpdateParamsType string

const (
	FleetifyRouteStepUpdateParamsTypeStart    FleetifyRouteStepUpdateParamsType = "`start`"
	FleetifyRouteStepUpdateParamsTypeJob      FleetifyRouteStepUpdateParamsType = "`job`"
	FleetifyRouteStepUpdateParamsTypePickup   FleetifyRouteStepUpdateParamsType = "`pickup`"
	FleetifyRouteStepUpdateParamsTypeDelivery FleetifyRouteStepUpdateParamsType = "`delivery`"
	FleetifyRouteStepUpdateParamsTypeBreak    FleetifyRouteStepUpdateParamsType = "`break`"
	FleetifyRouteStepUpdateParamsTypeLayover  FleetifyRouteStepUpdateParamsType = "`layover`"
	FleetifyRouteStepUpdateParamsTypeEnd      FleetifyRouteStepUpdateParamsType = "`end`"
)

type FleetifyRouteStepDeleteParams struct {
	RouteID string `path:"routeID,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	paramObj
}

// URLQuery serializes [FleetifyRouteStepDeleteParams]'s query parameters as
// `url.Values`.
func (r FleetifyRouteStepDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FleetifyRouteStepCompleteParams struct {
	RouteID string `path:"routeID,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	// Sets the status of the route step. Currently only `completed` is supported.
	//
	// Note: once marked `completed`, a step cannot transition to other statuses. You
	// can only update the document afterwards.
	Mode param.Opt[string] `json:"mode,omitzero"`
	// Sets the status of the route step. Currently only `completed` is supported.
	//
	// Note: once marked `completed`, a step cannot transition to other statuses. You
	// can only update the document afterwards.
	Status param.Opt[string] `json:"status,omitzero"`
	// A key-value map storing form submission data, where keys correspond to field
	// labels and values can be of any type depend on the type of according document
	// item.
	Document DocumentSubmission `json:"document,omitzero"`
	paramObj
}

func (r FleetifyRouteStepCompleteParams) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyRouteStepCompleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyRouteStepCompleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FleetifyRouteStepCompleteParams]'s query parameters as
// `url.Values`.
func (r FleetifyRouteStepCompleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
