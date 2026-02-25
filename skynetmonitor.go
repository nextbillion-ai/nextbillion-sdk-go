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

// SkynetMonitorService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetMonitorService] method instead.
type SkynetMonitorService struct {
	Options []option.RequestOption
}

// NewSkynetMonitorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSkynetMonitorService(opts ...option.RequestOption) (r SkynetMonitorService) {
	r = SkynetMonitorService{}
	r.Options = opts
	return
}

// Create a Monitor
func (r *SkynetMonitorService) New(ctx context.Context, params SkynetMonitorNewParams, opts ...option.RequestOption) (res *SkynetMonitorNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skynet/monitor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a Monitor
func (r *SkynetMonitorService) Get(ctx context.Context, id string, query SkynetMonitorGetParams, opts ...option.RequestOption) (res *SkynetMonitorGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/monitor/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a Monitor
func (r *SkynetMonitorService) Update(ctx context.Context, id string, params SkynetMonitorUpdateParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/monitor/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Get Monitor List
func (r *SkynetMonitorService) List(ctx context.Context, query SkynetMonitorListParams, opts ...option.RequestOption) (res *SkynetMonitorListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skynet/monitor/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a Monitor
func (r *SkynetMonitorService) Delete(ctx context.Context, id string, body SkynetMonitorDeleteParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/monitor/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type Metadata = any

type Monitor struct {
	// Unique ID of the monitor. This is the same ID that was generated at the time of
	// creating the monitor.
	ID string `json:"id"`
	// A UNIX epoch timestamp in seconds representing the time at which the monitor was
	// created.
	CreatedAt int64 `json:"created_at"`
	// Description of the monitor. The value would be the same as that provided for the
	// description parameter at the time of creating or updating the monitor.
	Description string `json:"description"`
	// An object returning the details of the geofence that are associated with the
	// monitor for an enter, exit or enter_and_exit type of monitor.
	GeofenceConfig MonitorGeofenceConfig `json:"geofence_config"`
	// Geofence IDs that are linked to the monitor. These IDs were associated with the
	// monitor at the time of creating or updating it.
	//
	// The monitor uses the geofences mentioned here to create events of type nature
	// for the eligible asset(s).
	Geofences []string `json:"geofences"`
	// An object returning the details of the idle activity constraints for a idle type
	// of monitor.
	IdleConfig MonitorIdleConfig `json:"idle_config"`
	// Use this object to update the attributes of the monitor.
	MatchFilter MonitorMatchFilter `json:"match_filter"`
	// Any valid json object data. Can be used to save customized data. Max size is
	// 65kb.
	MetaData Metadata `json:"meta_data"`
	// Name of the monitor. The value would be the same as that provided for the name
	// parameter at the time of creating or updating the monitor.
	Name string `json:"name"`
	// An object returning the details of the over-speeding constraints for a speeding
	// type of monitor.
	SpeedingConfig MonitorSpeedingConfig `json:"speeding_config"`
	// Tags of the monitor. The values would be the same as that provided for the tags
	// parameter at the time of creating or updating the monitor.
	Tags []string `json:"tags"`
	// Type of the monitor. It represents the type of asset activity that the monitor
	// is configured to detect.
	//
	// Any of "enter", "exit", "enter_and_exit", "speeding", "idle".
	Type MonitorType `json:"type"`
	// A UNIX epoch timestamp in seconds representing the time at which the monitor was
	// last updated.
	UpdatedAt int64 `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		GeofenceConfig respjson.Field
		Geofences      respjson.Field
		IdleConfig     respjson.Field
		MatchFilter    respjson.Field
		MetaData       respjson.Field
		Name           respjson.Field
		SpeedingConfig respjson.Field
		Tags           respjson.Field
		Type           respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Monitor) RawJSON() string { return r.JSON.raw }
func (r *Monitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning the details of the geofence that are associated with the
// monitor for an enter, exit or enter_and_exit type of monitor.
type MonitorGeofenceConfig struct {
	// An array of geofence IDs that are linked to the monitor. Geofences are
	// geographic boundaries that can be used to trigger events based on an asset's
	// location.
	GeofenceIDs []string `json:"geofence_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GeofenceIDs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGeofenceConfig) RawJSON() string { return r.JSON.raw }
func (r *MonitorGeofenceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning the details of the idle activity constraints for a idle type
// of monitor.
type MonitorIdleConfig struct {
	// This parameter returns the distance threshold that was used to determine if the
	// asset was idle or not. The value returned for this parameter is the same as that
	// provided while creating or updating a idle type monitor.
	DistanceTolerance float64 `json:"distance_tolerance"`
	// This parameter returns the time duration for which the monitor tracks the
	// distance covered by an asset before triggering an idle event. The value returned
	// for this parameter is the same as that provided while creating or updating a
	// idle type monitor.
	TimeTolerance int64 `json:"time_tolerance"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DistanceTolerance respjson.Field
		TimeTolerance     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorIdleConfig) RawJSON() string { return r.JSON.raw }
func (r *MonitorIdleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use this object to update the attributes of the monitor.
type MonitorMatchFilter struct {
	// A string type dictionary object to specify the attributes which will be used to
	// identify the asset(s) on which the monitor would be applied. Please note that
	// using this parameter overwrites the existing attributes of the monitor.
	//
	// If the attributes added to a monitor do not match fully with the attributes
	// added to any asset, the monitor will be ineffective.
	//
	// Please note that the maximum number of key:value pairs that
	// 'include_all_of_attributes' can take is 100. Also, the overall size of the
	// match_filter object should not exceed 65kb.
	IncludeAllOfAttributes any `json:"include_all_of_attributes"`
	// A string dictionary object to specify the attributes, separated by a ,. Only the
	// assets with any one of the attributes added to this parameter will be linked to
	// this monitor. Once an asset and a monitor are linked, the monitor will be able
	// to create events for the asset when an activity specified in type is detected.
	//
	// If no input is provided for this object or if the attributes added here do not
	// match at least one of the attributes added to any asset, the monitor will be
	// ineffective.
	//
	// Please note that the maximum number of key:value pairs that
	// include_any_of_attributes can take is 100. Also, the overall size of
	// match_filter object should not exceed 65kb.
	IncludeAnyOfAttributes any `json:"include_any_of_attributes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludeAllOfAttributes respjson.Field
		IncludeAnyOfAttributes respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorMatchFilter) RawJSON() string { return r.JSON.raw }
func (r *MonitorMatchFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning the details of the over-speeding constraints for a speeding
// type of monitor.
type MonitorSpeedingConfig struct {
	// This property returns the actual speed limit that the monitor uses as a
	// threshold for generating a speed limit event. The value returned for this
	// parameter is the same as that provided while creating or updating a speeding
	// type monitor.
	CustomerSpeedLimit int64 `json:"customer_speed_limit"`
	// This property returns the time duration value, in milliseconds, for which the
	// monitor will track the speed of the asset. An event is triggered if the speed
	// remains higher than the specified limit for a duration more than the tolerance
	// value.
	//
	// The value returned for this parameter is the same as that provided while
	// creating or updating a speeding type monitor.
	TimeTolerance int64 `json:"time_tolerance"`
	// A boolean value denoting if the administrative speed limit of the road was used
	// as speed limit threshold for triggering events. The value returned for this
	// parameter is the same as that provided while creating or updating a speeding
	// type monitor.
	UseAdminSpeedLimit bool `json:"use_admin_speed_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerSpeedLimit respjson.Field
		TimeTolerance      respjson.Field
		UseAdminSpeedLimit respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorSpeedingConfig) RawJSON() string { return r.JSON.raw }
func (r *MonitorSpeedingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the monitor. It represents the type of asset activity that the monitor
// is configured to detect.
type MonitorType string

const (
	MonitorTypeEnter        MonitorType = "enter"
	MonitorTypeExit         MonitorType = "exit"
	MonitorTypeEnterAndExit MonitorType = "enter_and_exit"
	MonitorTypeSpeeding     MonitorType = "speeding"
	MonitorTypeIdle         MonitorType = "idle"
)

// An object with pagination details of the search results. Use this object to
// implement pagination in your application.
type Pagination struct {
	// A boolean value indicating whether there are more items available beyond the
	// current page.
	Hasmore bool `json:"hasmore"`
	// An integer value indicating the current page number (starting at 0).
	Page int64 `json:"page"`
	// An integer value indicating the maximum number of items retrieved per page.
	Size int64 `json:"size"`
	// An integer value indicating the total number of items available in the data set.
	// This parameter can be used to calculate the total number of pages available.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hasmore     respjson.Field
		Page        respjson.Field
		Size        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pagination) RawJSON() string { return r.JSON.raw }
func (r *Pagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetMonitorNewResponse struct {
	// A data object containing the ID of the monitor created.
	Data SkynetMonitorNewResponseData `json:"data"`
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
func (r SkynetMonitorNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetMonitorNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A data object containing the ID of the monitor created.
type SkynetMonitorNewResponseData struct {
	// Unique ID of the monitor created. Please note this ID cannot be updated.
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetMonitorNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetMonitorNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetMonitorGetResponse struct {
	// A data object containing the details of the monitor.
	Data SkynetMonitorGetResponseData `json:"data"`
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
func (r SkynetMonitorGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetMonitorGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A data object containing the details of the monitor.
type SkynetMonitorGetResponseData struct {
	Monitor Monitor `json:"monitor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Monitor     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetMonitorGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetMonitorGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetMonitorListResponse struct {
	// A data object containing the result.
	Data SkynetMonitorListResponseData `json:"data"`
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
func (r SkynetMonitorListResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetMonitorListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A data object containing the result.
type SkynetMonitorListResponseData struct {
	// An array of objects listing all the monitors. Each object represents one
	// monitor.
	List []Monitor `json:"list"`
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
func (r SkynetMonitorListResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetMonitorListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetMonitorNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// Use this parameter to add tags to the monitor. tags can be used for filtering
	// monitors in the _Get Monitor List_ operation. They can also be used for easy
	// identification of monitors.
	//
	// Please note that valid tags are strings, consisting of alphanumeric characters
	// (A-Z, a-z, 0-9) along with the underscore ('\_') and hyphen ('-') symbols.
	Tags []string `json:"tags,omitzero" api:"required"`
	// Specify the type of activity the monitor would detect.
	//
	// The monitor will be able to detect the specified type of activity and create
	// events for eligible asset. A monitor can detect following types of asset
	// activity:
	//
	//   - enter: The monitor will create an event when a linked asset enters into the
	//     specified geofence.
	//
	//   - exit: The monitor will create an event when a linked asset exits the specified
	//     geofence.
	//
	//   - enter_and_exit: The monitor will create an event when a linked asset either
	//     enters or exits the specified geofence.
	//
	//   - speeding: The monitor will create an event when a linked asset exceeds a given
	//     speed limit.
	//
	//   - idle: The monitor will create an event when a linked asset exhibits idle
	//     activity.
	//
	// Please note that assets and geofences can be linked to a monitor using the
	// match_filter and geofence_config attributes respectively.
	//
	// Any of "enter", "exit", "enter_and_exit", "speeding", "idle".
	Type SkynetMonitorNewParamsType `json:"type,omitzero" api:"required"`
	// Set a unique ID for the new monitor. If not provided, an ID will be
	// automatically generated in UUID format. A valid custom*id can contain letters,
	// numbers, "-", & "*" only.
	//
	// Please note that the ID of an monitor can not be changed once it is created.
	CustomID param.Opt[string] `json:"custom_id,omitzero"`
	// Add a description for your monitor using this parameter.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the monitor. Use this field to assign a meaningful, custom name to the
	// monitor being created.
	Name param.Opt[string] `json:"name,omitzero"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetMonitorNewParamsCluster `query:"cluster,omitzero" json:"-"`
	// Geofences are geographic boundaries surrounding an area of interest.
	// geofence_config is used to specify the geofences for creating enter or exit type
	// of events based on the asset's location. When an asset associated with the
	// monitor enters the given geofence, an enter type event is created, whereas when
	// the asset moves out of the geofence an exit type event is created.
	//
	// Please note that this object is mandatory when the monitor type belongs to one
	// of enter, exit or enter_and_exit.
	GeofenceConfig SkynetMonitorNewParamsGeofenceConfig `json:"geofence_config,omitzero"`
	// **Deprecated. Please use the geofence_config to specify the geofence_ids for
	// this monitor.**
	//
	// An array of strings to collect the geofence IDs that should be linked to the
	// monitor. Geofences are geographic boundaries that can be used to trigger events
	// based on an asset's location.
	GeofenceIDs []string `json:"geofence_ids,omitzero"`
	// idle_config is used to set up constraints for creating idle events. When an
	// asset associated with the monitor has not moved a given distance within a given
	// time, the Live Tracking API can create events to denote such instances. Please
	// note that this object is mandatory when the monitor type is idle.
	//
	// Let's look at the properties of this object.
	IdleConfig SkynetMonitorNewParamsIdleConfig `json:"idle_config,omitzero"`
	// This object is used to identify the asset(s) on which the monitor would be
	// applied.
	MatchFilter SkynetMonitorNewParamsMatchFilter `json:"match_filter,omitzero"`
	// Any valid json object data. Can be used to save customized data. Max size is
	// 65kb.
	MetaData Metadata `json:"meta_data,omitzero"`
	// speeding_config is used to set up constraints for creating over-speed events.
	// When an asset associated with a monitor is traveling at a speed above the given
	// limits, the Live Tracking API can create events to denote such instances. There
	// is also an option to set up a tolerance before creating an event. Please note
	// that this object is mandatory when type=speeding.
	//
	// Let's look at the properties of this object.
	SpeedingConfig SkynetMonitorNewParamsSpeedingConfig `json:"speeding_config,omitzero"`
	paramObj
}

func (r SkynetMonitorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetMonitorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetMonitorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetMonitorNewParams]'s query parameters as `url.Values`.
func (r SkynetMonitorNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the type of activity the monitor would detect.
//
// The monitor will be able to detect the specified type of activity and create
// events for eligible asset. A monitor can detect following types of asset
// activity:
//
//   - enter: The monitor will create an event when a linked asset enters into the
//     specified geofence.
//
//   - exit: The monitor will create an event when a linked asset exits the specified
//     geofence.
//
//   - enter_and_exit: The monitor will create an event when a linked asset either
//     enters or exits the specified geofence.
//
//   - speeding: The monitor will create an event when a linked asset exceeds a given
//     speed limit.
//
//   - idle: The monitor will create an event when a linked asset exhibits idle
//     activity.
//
// Please note that assets and geofences can be linked to a monitor using the
// match_filter and geofence_config attributes respectively.
type SkynetMonitorNewParamsType string

const (
	SkynetMonitorNewParamsTypeEnter        SkynetMonitorNewParamsType = "enter"
	SkynetMonitorNewParamsTypeExit         SkynetMonitorNewParamsType = "exit"
	SkynetMonitorNewParamsTypeEnterAndExit SkynetMonitorNewParamsType = "enter_and_exit"
	SkynetMonitorNewParamsTypeSpeeding     SkynetMonitorNewParamsType = "speeding"
	SkynetMonitorNewParamsTypeIdle         SkynetMonitorNewParamsType = "idle"
)

// the cluster of the region you want to use
type SkynetMonitorNewParamsCluster string

const (
	SkynetMonitorNewParamsClusterAmerica SkynetMonitorNewParamsCluster = "america"
)

// Geofences are geographic boundaries surrounding an area of interest.
// geofence_config is used to specify the geofences for creating enter or exit type
// of events based on the asset's location. When an asset associated with the
// monitor enters the given geofence, an enter type event is created, whereas when
// the asset moves out of the geofence an exit type event is created.
//
// Please note that this object is mandatory when the monitor type belongs to one
// of enter, exit or enter_and_exit.
//
// The property GeofenceIDs is required.
type SkynetMonitorNewParamsGeofenceConfig struct {
	// An array of strings to collect the geofence IDs that should be linked to the
	// monitor. Please note geofence_ids are mandatory when using the geofence_config
	// attribute.
	GeofenceIDs []string `json:"geofence_ids,omitzero" api:"required"`
	paramObj
}

func (r SkynetMonitorNewParamsGeofenceConfig) MarshalJSON() (data []byte, err error) {
	type shadow SkynetMonitorNewParamsGeofenceConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetMonitorNewParamsGeofenceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// idle_config is used to set up constraints for creating idle events. When an
// asset associated with the monitor has not moved a given distance within a given
// time, the Live Tracking API can create events to denote such instances. Please
// note that this object is mandatory when the monitor type is idle.
//
// Let's look at the properties of this object.
//
// The property DistanceTolerance is required.
type SkynetMonitorNewParamsIdleConfig struct {
	// Use this parameter to configure a distance threshold that will be used to
	// determine if the asset was idle or not. If the asset moves by a distance less
	// than the value of this parameter within a certain time period, the monitor would
	// create an idle event against the asset. The distance_tolerance should be
	// provided in meters.
	//
	// Users can set an appropriate value for this parameter, along with appropriate
	// time_tolerance value, to avoid triggering idle events when the asset is crossing
	// a busy intersection or waiting at the traffic lights.
	DistanceTolerance float64 `json:"distance_tolerance" api:"required"`
	// Use this parameter to configure a time duration for which the monitor would
	// track the distance covered by an asset before triggering an idle event. The
	// time_tolerance should be provided in milliseconds.
	//
	// If the distance covered by the asset during a time_tolerance is less than that
	// specified in distance_tolerance the asset will be assumed to be idle.
	//
	// Please observe that this attribute along with distance_tolerance parameter can
	// be used to control the "sensitivity" of the monitor with respect to idle alerts.
	// If the distance_tolerance is set a high value, then setting time_tolerance to a
	// low value may result in a situation where asset is always judged as idle. On the
	// contrary, it might never be judged as idle if distance_tolerance is set to a low
	// value but time_tolerance is set to a high value.
	//
	// It is recommended to use these properties with appropriate values to trigger
	// genuine idle events. The appropriate values might depend on the traffic
	// conditions, nature of operations that the asset is involved in, type of asset
	// and other factors.
	TimeTolerance param.Opt[int64] `json:"time_tolerance,omitzero"`
	paramObj
}

func (r SkynetMonitorNewParamsIdleConfig) MarshalJSON() (data []byte, err error) {
	type shadow SkynetMonitorNewParamsIdleConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetMonitorNewParamsIdleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This object is used to identify the asset(s) on which the monitor would be
// applied.
type SkynetMonitorNewParamsMatchFilter struct {
	// A string type dictionary object to specify the attributes. Only the assets
	// having all of the attributes added to this parameter will be linked to this
	// monitor. Once an asset is linked to a monitor, the monitor will be able to
	// create events for that asset whenever an activity specified in type is detected.
	// Multiple attributes should be separated by a comma ,.
	//
	// Please note that this parameter can not be used in conjunction with
	// include_any_of_attributes. Also, the maximum number of key:value pairs that this
	// parameter can take is 100 and the overall size of the match_filter object should
	// not exceed 65kb.
	IncludeAllOfAttributes any `json:"include_all_of_attributes,omitzero"`
	// A string type dictionary object to specify the attributes. The assets having at
	// least one of the attributes added to this parameter will be linked to this
	// monitor. Once an asset is linked to a monitor, the monitor will be able to
	// create events for that asset whenever an activity specified in type is detected.
	// Multiple attributes should be separated by a comma ,.
	//
	// Please note that this parameter can not be used in conjunction with
	// include_all_of_attributes. Also, the maximum number of key:value pairs that this
	// parameter can take is 100 and the overall size of the match_filter object should
	// not exceed 65kb.
	IncludeAnyOfAttributes any `json:"include_any_of_attributes,omitzero"`
	paramObj
}

func (r SkynetMonitorNewParamsMatchFilter) MarshalJSON() (data []byte, err error) {
	type shadow SkynetMonitorNewParamsMatchFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetMonitorNewParamsMatchFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// speeding_config is used to set up constraints for creating over-speed events.
// When an asset associated with a monitor is traveling at a speed above the given
// limits, the Live Tracking API can create events to denote such instances. There
// is also an option to set up a tolerance before creating an event. Please note
// that this object is mandatory when type=speeding.
//
// Let's look at the properties of this object.
type SkynetMonitorNewParamsSpeedingConfig struct {
	// Use this parameter to establish the speed limit that will allow the monitor to
	// create events, depending on the time_tolerance value, when an asset's tracked
	// speed exceeds it. The speed limit should be specified in meters per second.
	//
	// Please note that customer_speed_limit is mandatory when use_admin_speed_limit is
	// false. However, when use_admin_speed_limit is true, customer_speed_limit is
	// ineffective.
	CustomerSpeedLimit param.Opt[int64] `json:"customer_speed_limit,omitzero"`
	// Use this parameter to configure a time tolerance before triggering an event.
	// Adding a tolerance would make the Tracking service wait for the specified time
	// before triggering the event. Consequently, an event is triggered only when the
	// time for which the asset has been over-speeding continuously, exceeds the
	// configured tolerance time. The unit for this parameter is milliseconds.
	//
	// It can be seen that this attribute is used to control the "sensitivity" of the
	// monitor with respect to speed alerts. Higher the value of time_tolerance the
	// less sensitive the monitor would be to instances of over-speeding. Conversely,
	// if 'time_tolerance' is set to 0, the monitor will be extremely sensitive and
	// will create an event as soon as tracking information with a speed value greater
	// than the specified limit is received.
	TimeTolerance param.Opt[int64] `json:"time_tolerance,omitzero"`
	// A boolean attribute to indicate which speed limit values should be used by the
	// monitor. When use_admin_speed_limit is true, the administrative speed limit of
	// the road on which the asset is located, will be used to generate events when the
	// asset’s tracked speed exceeds it. Whereas, when use_admin_speed_limit is false,
	// the customer_speed_limit specified will be used to generate events when the
	// asset's tracked speed exceeds it.
	//
	// Please note that if use_admin_speed_limit is false, customer_speed_limit is
	// mandatory, however, when use_admin_speed_limit is true then customer_speed_limit
	// is ineffective.
	UseAdminSpeedLimit param.Opt[bool] `json:"use_admin_speed_limit,omitzero"`
	paramObj
}

func (r SkynetMonitorNewParamsSpeedingConfig) MarshalJSON() (data []byte, err error) {
	type shadow SkynetMonitorNewParamsSpeedingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetMonitorNewParamsSpeedingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetMonitorGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetMonitorGetParams]'s query parameters as `url.Values`.
func (r SkynetMonitorGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SkynetMonitorUpdateParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// Use this parameter to update the description of the monitor.
	Description param.Opt[string] `json:"description,omitzero"`
	// Use this parameter to update the name of the monitor. Users can add meaningful
	// names to the monitors like "warehouse_exit", "depot_entry" etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// geofence_config is used to update the set of geofences linked to the monitor for
	// creating enter or exit type of events based on the asset's location. Please note
	// that this object is mandatory when the monitor type belongs to one of enter,
	// exit or enter_and_exit.
	GeofenceConfig SkynetMonitorUpdateParamsGeofenceConfig `json:"geofence_config,omitzero"`
	// Use this parameter to update the geofences linked to the monitor by providing
	// the geofence id as , separated strings. Geofences are geographic boundaries that
	// can be used to trigger events based on an asset's location.
	GeofenceIDs []string `json:"geofence_ids,omitzero"`
	// idle_config is used to update the constraints for creating idle events. When an
	// asset associated with the monitor has not moved a given distance within a given
	// time, the Live Tracking API can create events to denote such instances.
	//
	// Please note that this object is mandatory when the monitor type is idle.
	IdleConfig SkynetMonitorUpdateParamsIdleConfig `json:"idle_config,omitzero"`
	// Use this object to update the attributes of the monitor. Please note that using
	// this property will overwrite the existing attributes that the monitor might be
	// using currently to match any asset(s).
	MatchFilter SkynetMonitorUpdateParamsMatchFilter `json:"match_filter,omitzero"`
	// Any valid json object data. Can be used to save customized data. Max size is
	// 65kb.
	MetaData Metadata `json:"meta_data,omitzero"`
	// speeding_config is used to update the tolerance values for creating over-speed
	// events. When an asset associated with a monitor is traveling at a speed above
	// the given limits, Live Tracking API creates events to indicate such instances.
	//
	// Please note that this object is mandatory when the monitor type is speeding.
	SpeedingConfig SkynetMonitorUpdateParamsSpeedingConfig `json:"speeding_config,omitzero"`
	// Use this parameter to update the tags of the monitor. tags can be used for
	// filtering monitors in the _Get Monitor List_ operation. They can also be used
	// for easy identification of monitors. Using this parameter overwrites the
	// existing tags of the monitor.
	//
	// Please note that valid tags are strings, consisting of alphanumeric characters
	// (A-Z, a-z, 0-9) along with the underscore ('\_') and hyphen ('-') symbols.
	Tags []string `json:"tags,omitzero"`
	// Use this parameter to update the type of the monitor. The monitor will be able
	// to detect the specified type of activity and create events for eligible asset. A
	// monitor can detect following types of asset activity:
	//
	//   - enter: The monitor will create an event when a linked asset enters into the
	//     specified geofence.
	//
	//   - exit: The monitor will create an event when a linked asset exits the specified
	//     geofence.
	//
	//   - enter_and_exit: The monitor will create an event when a linked asset either
	//     enters or exits the specified geofence.
	//
	//   - speeding: The monitor will create an event when a linked asset exceeds a given
	//     speed limit.
	//
	//   - idle: The monitor will create an event when a linked asset exhibits idle
	//     activity.
	//
	// Please note that assets and geofences can be linked to a monitor using the
	// match_filter and geofence_config attributes respectively.
	//
	// Any of "enter", "exit", "enter_and_exit", "speeding", "idle".
	Type SkynetMonitorUpdateParamsType `json:"type,omitzero"`
	paramObj
}

func (r SkynetMonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetMonitorUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetMonitorUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetMonitorUpdateParams]'s query parameters as
// `url.Values`.
func (r SkynetMonitorUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// geofence_config is used to update the set of geofences linked to the monitor for
// creating enter or exit type of events based on the asset's location. Please note
// that this object is mandatory when the monitor type belongs to one of enter,
// exit or enter_and_exit.
//
// The property GeofenceIDs is required.
type SkynetMonitorUpdateParamsGeofenceConfig struct {
	// Use this array to update the geofence IDs that should be linked to the monitor.
	// Please note geofence_ids are mandatory when using the geofence_config attribute.
	GeofenceIDs []string `json:"geofence_ids,omitzero" api:"required"`
	paramObj
}

func (r SkynetMonitorUpdateParamsGeofenceConfig) MarshalJSON() (data []byte, err error) {
	type shadow SkynetMonitorUpdateParamsGeofenceConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetMonitorUpdateParamsGeofenceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// idle_config is used to update the constraints for creating idle events. When an
// asset associated with the monitor has not moved a given distance within a given
// time, the Live Tracking API can create events to denote such instances.
//
// Please note that this object is mandatory when the monitor type is idle.
//
// The property DistanceTolerance is required.
type SkynetMonitorUpdateParamsIdleConfig struct {
	// Use this parameter to update the distance threshold that will be used to
	// determine if the asset was idle or not. When the asset, within time_tolerance
	// duration, moves less than the value for this parameter, the monitor creates an
	// idle event against the asset. The distance_tolerance should be provided in
	// meters.
	//
	// Please note distance_tolerance is mandatory when idle_config attribute is used.
	DistanceTolerance float64 `json:"distance_tolerance" api:"required"`
	// Use this parameter to update the time duration for which the monitor would track
	// the distance covered by an asset before triggering an idle event. The
	// time_tolerance should be provided in milliseconds.
	//
	// If the distance covered by the asset during a time_tolerance is less than that
	// specified in distance_tolerance the asset will be assumed to be idle.
	//
	// This attribute along with distance_tolerance parameter can be used to control
	// the "sensitivity" of the monitor with respect to idle alerts. It is recommended
	// to use these properties with appropriate values to trigger genuine idle events.
	// The appropriate values might depend on the traffic conditions, nature of
	// operations that the asset is involved in, type of asset and other factors.
	TimeTolerance param.Opt[int64] `json:"time_tolerance,omitzero"`
	paramObj
}

func (r SkynetMonitorUpdateParamsIdleConfig) MarshalJSON() (data []byte, err error) {
	type shadow SkynetMonitorUpdateParamsIdleConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetMonitorUpdateParamsIdleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use this object to update the attributes of the monitor. Please note that using
// this property will overwrite the existing attributes that the monitor might be
// using currently to match any asset(s).
type SkynetMonitorUpdateParamsMatchFilter struct {
	// A string type dictionary object to specify the attributes. Only the assets
	// having all of the attributes added to this parameter will be linked to this
	// monitor. Once an asset is linked to a monitor, the monitor will be able to
	// create events for that asset whenever an activity specified in type is detected.
	// Multiple attributes should be separated by a comma ,.
	//
	// Please note that this parameter can not be used in conjunction with
	// include_any_of_attributes. Also, the maximum number of key:value pairs that this
	// parameter can take is 100 and the overall size of the match_filter object should
	// not exceed 65kb.
	IncludeAllOfAttributes any `json:"include_all_of_attributes,omitzero"`
	// A string type dictionary object to specify the attributes. The assets having at
	// least one of the attributes added to this parameter will be linked to this
	// monitor. Once an asset is linked to a monitor, the monitor will be able to
	// create events for that asset whenever an activity specified in type is detected.
	// Multiple attributes should be separated by a comma ,.
	//
	// Please note that this parameter can not be used in conjunction with
	// include_all_of_attributes. Also, the maximum number of key:value pairs that this
	// parameter can take is 100 and the overall size of the match_filter object should
	// not exceed 65kb.
	IncludeAnyOfAttributes any `json:"include_any_of_attributes,omitzero"`
	paramObj
}

func (r SkynetMonitorUpdateParamsMatchFilter) MarshalJSON() (data []byte, err error) {
	type shadow SkynetMonitorUpdateParamsMatchFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetMonitorUpdateParamsMatchFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// speeding_config is used to update the tolerance values for creating over-speed
// events. When an asset associated with a monitor is traveling at a speed above
// the given limits, Live Tracking API creates events to indicate such instances.
//
// Please note that this object is mandatory when the monitor type is speeding.
type SkynetMonitorUpdateParamsSpeedingConfig struct {
	// Use this parameter to update the speed limit value that the monitor will use to
	// create events, depending on the time_tolerance value. The speed limit should be
	// specified in meters per second.
	//
	// Please note that customer_speed_limit is mandatory when use_admin_speed_limit is
	// false. However, when use_admin_speed_limit is true, customer_speed_limit is
	// ineffective.
	CustomerSpeedLimit param.Opt[string] `json:"customer_speed_limit,omitzero"`
	// Use this parameter to update the time tolerance before triggering an event.
	// Adding a tolerance would make the Tracking service wait for the specified time
	// before triggering the event. Consequently, an event is triggered only when the
	// time for which the asset has been over-speeding continuously, exceeds the
	// configured tolerance time. The unit for this parameter is milliseconds.
	//
	// It can be seen that this attribute is used to control the "sensitivity" of the
	// monitor with respect to speed alerts. Higher the value of time_tolerance the
	// less sensitive the monitor would be to instances of over-speeding. Conversely,
	// if 'time_tolerance' is set to 0, the monitor will be extremely sensitive and
	// will create an event as soon as tracking information with a speed value greater
	// than the specified limit is received.
	TimeTolerance param.Opt[int64] `json:"time_tolerance,omitzero"`
	// Use this attribute to update which speed limit values will be used by the
	// monitor. When use_admin_speed_limit is true, the administrative speed limit of
	// the road on which the asset is located, is used to generate events when the
	// asset’s tracked speed exceeds it. Whereas, when use_admin_speed_limit is false,
	// the customer_speed_limit specified will be used to generate events when the
	// asset's tracked speed exceeds it.
	//
	// Please note that if use_admin_speed_limit is false, customer_speed_limit is
	// mandatory, otherwise when use_admin_speed_limit is true then
	// customer_speed_limit is ineffective.
	UseAdminSpeedLimit param.Opt[bool] `json:"use_admin_speed_limit,omitzero"`
	paramObj
}

func (r SkynetMonitorUpdateParamsSpeedingConfig) MarshalJSON() (data []byte, err error) {
	type shadow SkynetMonitorUpdateParamsSpeedingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetMonitorUpdateParamsSpeedingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use this parameter to update the type of the monitor. The monitor will be able
// to detect the specified type of activity and create events for eligible asset. A
// monitor can detect following types of asset activity:
//
//   - enter: The monitor will create an event when a linked asset enters into the
//     specified geofence.
//
//   - exit: The monitor will create an event when a linked asset exits the specified
//     geofence.
//
//   - enter_and_exit: The monitor will create an event when a linked asset either
//     enters or exits the specified geofence.
//
//   - speeding: The monitor will create an event when a linked asset exceeds a given
//     speed limit.
//
//   - idle: The monitor will create an event when a linked asset exhibits idle
//     activity.
//
// Please note that assets and geofences can be linked to a monitor using the
// match_filter and geofence_config attributes respectively.
type SkynetMonitorUpdateParamsType string

const (
	SkynetMonitorUpdateParamsTypeEnter        SkynetMonitorUpdateParamsType = "enter"
	SkynetMonitorUpdateParamsTypeExit         SkynetMonitorUpdateParamsType = "exit"
	SkynetMonitorUpdateParamsTypeEnterAndExit SkynetMonitorUpdateParamsType = "enter_and_exit"
	SkynetMonitorUpdateParamsTypeSpeeding     SkynetMonitorUpdateParamsType = "speeding"
	SkynetMonitorUpdateParamsTypeIdle         SkynetMonitorUpdateParamsType = "idle"
)

type SkynetMonitorListParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
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
	// tags can be used to filter the monitors. Only those monitors which have all the
	// tags provided here, will be included in the search result. In case multiple tags
	// need to be specified, use , to separate them.
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetMonitorListParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetMonitorListParams]'s query parameters as
// `url.Values`.
func (r SkynetMonitorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetMonitorListParamsCluster string

const (
	SkynetMonitorListParamsClusterAmerica SkynetMonitorListParamsCluster = "america"
)

type SkynetMonitorDeleteParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetMonitorDeleteParams]'s query parameters as
// `url.Values`.
func (r SkynetMonitorDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
