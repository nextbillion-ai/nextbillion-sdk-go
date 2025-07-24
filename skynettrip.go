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

// SkynetTripService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetTripService] method instead.
type SkynetTripService struct {
	Options []option.RequestOption
}

// NewSkynetTripService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSkynetTripService(opts ...option.RequestOption) (r SkynetTripService) {
	r = SkynetTripService{}
	r.Options = opts
	return
}

// Retrieves detailed information about a specific trip.
func (r *SkynetTripService) Get(ctx context.Context, id string, query SkynetTripGetParams, opts ...option.RequestOption) (res *SkynetTripGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/trip/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates the data of a specified trip with the provided data.
func (r *SkynetTripService) Update(ctx context.Context, id string, params SkynetTripUpdateParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/trip/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Deletes a specified trip from the system.
func (r *SkynetTripService) Delete(ctx context.Context, id string, body SkynetTripDeleteParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/trip/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// End a trip
func (r *SkynetTripService) End(ctx context.Context, params SkynetTripEndParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/trip/end"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get summary of an ended trip
func (r *SkynetTripService) GetSummary(ctx context.Context, id string, query SkynetTripGetSummaryParams, opts ...option.RequestOption) (res *SkynetTripGetSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/trip/%s/summary", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Add a new trip to the system with the provided data.
func (r *SkynetTripService) Start(ctx context.Context, params SkynetTripStartParams, opts ...option.RequestOption) (res *SkynetTripStartResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/trip/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// An object with details of the `asset` properties.
type Asset struct {
	// ID of the `asset`. This is the same ID that was generated/provided at the time
	// of creating the `asset`.
	ID string `json:"id"`
	// A string dictionary object containing `attributes` of the `asset`. These
	// `attributes` were associated with the `asset` at the time of creating or
	// updating it.
	//
	// `attributes` can be added to an `asset` using the _Update Asset Attributes_
	// method.
	Attributes any `json:"attributes"`
	// A UNIX epoch timestamp in seconds representing the time at which the `asset` was
	// created.
	CreatedAt int64 `json:"created_at"`
	// Description of the `asset`. The value would be the same as that provided for the
	// `description` parameter at the time of creating or updating the `asset`.
	Description string `json:"description"`
	// ID of the `device` that is linked to this asset. Please note that there can be
	// multiple `device_id` linked to a single `asset`. An empty response is returned
	// if no devices are linked to the `asset`.
	//
	// User can link a device to an `asset` using the _Bind Asset to Device_ method.
	DeviceID string `json:"device_id"`
	// An object with details of the last tracked location of the asset.
	LatestLocation AssetLatestLocation `json:"latest_location"`
	// Any valid json object data. Can be used to save customized data. Max size is
	// 65kb.
	MetaData MetaData `json:"meta_data"`
	// Name of the `asset`. The value would be the same as that provided for the `name`
	// parameter at the time of creating or updating the `asset`.
	Name string `json:"name"`
	// State of the asset. It will be "active" when the asset is in use or available
	// for use, and it will be "deleted" in case the asset has been deleted.
	State string `json:"state"`
	// **This parameter will be deprecated soon! Please move existing `tags` to
	// `attributes` parameter.**
	//
	// Tags of the asset. These were associated with the `asset` when it was created or
	// updated. `tags` can be used for filtering assets in operations like _Get Asset
	// List_ and asset **Search** methods. They can also be used for monitoring of
	// assets using **Monitor** methods after linking `tags` and `asset`.
	Tags []string `json:"tags"`
	// A UNIX epoch timestamp in seconds representing the last time when the `asset`
	// was tracked.
	TrackedAt int64 `json:"tracked_at"`
	// A UNIX epoch timestamp in seconds representing the time at which the `asset` was
	// last updated.
	UpdatedAt int64 `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Attributes     respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		DeviceID       respjson.Field
		LatestLocation respjson.Field
		MetaData       respjson.Field
		Name           respjson.Field
		State          respjson.Field
		Tags           respjson.Field
		TrackedAt      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Asset) RawJSON() string { return r.JSON.raw }
func (r *Asset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the last tracked location of the asset.
type AssetLatestLocation struct {
	// If available, this property returns the accuracy of the GPS information received
	// at the last tracked location. It is represented as an estimated horizontal
	// accuracy radius, in meters, at the 68th percentile confidence level.
	Accuracy float64 `json:"accuracy"`
	// If available in the GPS information, this property returns the altitude of the
	// `asset` at the last tracked location. It is represented as height, in meters,
	// above the WGS84 reference ellipsoid.
	Altitude float64 `json:"altitude"`
	// If available in the GPS information, this property returns the heading of the
	// `asset` calculated from true north in clockwise direction at the last tracked
	// location. Please note that the bearing is not affected by the device
	// orientation.
	//
	// The bearing will always be in the range of [0, 360).
	Bearing float64 `json:"bearing"`
	// An object with the coordinates of the last tracked location.
	Location AssetLatestLocationLocation `json:"location"`
	// If available in the GPS information, this property returns the speed of the
	// `asset`, in meters per second, at the last tracked location.
	Speed float64 `json:"speed"`
	// A UNIX epoch timestamp in milliseconds, representing the time at which the
	// location was tracked.
	Timestamp int64 `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accuracy    respjson.Field
		Altitude    respjson.Field
		Bearing     respjson.Field
		Location    respjson.Field
		Speed       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetLatestLocation) RawJSON() string { return r.JSON.raw }
func (r *AssetLatestLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with the coordinates of the last tracked location.
type AssetLatestLocationLocation struct {
	// Latitude of the tracked location of the `asset`.
	Lat float64 `json:"lat"`
	// Longitude of the tracked location of the `asset`.
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
func (r AssetLatestLocationLocation) RawJSON() string { return r.JSON.raw }
func (r *AssetLatestLocationLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TripStop struct {
	// Returns the ID of the geofence that was used to indicate the area to make a
	// stop.
	GeofenceID string `json:"geofence_id"`
	// Returns any meta data that was added to provide additional information about the
	// stop.
	MetaData any `json:"meta_data"`
	// Returns the name of the stop that was provided when configuring this stop for
	// the trip.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GeofenceID  respjson.Field
		MetaData    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TripStop) RawJSON() string { return r.JSON.raw }
func (r *TripStop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetTripGetResponse struct {
	// An container for the trip returned by the service.
	Data SkynetTripGetResponseData `json:"data"`
	// Displays the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Message string `json:"message"`
	// A string indicating the state of the response. On successful responses, the
	// value will be `Ok`. Indicative error messages are returned for different errors.
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
func (r SkynetTripGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetTripGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An container for the trip returned by the service.
type SkynetTripGetResponseData struct {
	// An object containing the returned trip details.
	Trip SkynetTripGetResponseDataTrip `json:"trip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Trip        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetTripGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetTripGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the returned trip details.
type SkynetTripGetResponseDataTrip struct {
	// Returns the unique identifier of the trip.
	ID string `json:"id"`
	// Returns the ID of the asset linked to the trip when the trip was started or
	// updated.
	AssetID string `json:"asset_id"`
	// Returns the `attributes` provided for the trip at the time of starting or
	// updating it.
	Attributes any `json:"attributes"`
	// Returns the time, expressed as UNIX epoch timestamp in milliseconds, when the
	// trip was created.
	CreatedAt int64 `json:"created_at"`
	// Returns the custom description for the trip as provided at the time of starting
	// or updating the trip.
	Description string `json:"description"`
	// Returns the time, expressed as UNIX epoch timestamp in milliseconds, when the
	// trip ended.
	EndedAt int64 `json:"ended_at"`
	// Returns the metadata containing additional information about the trip as
	// provided at the time of starting or updating the trip.
	MetaData any `json:"meta_data"`
	// Returns the name for the trip as provided at the time of starting or updating
	// the trip.
	Name string `json:"name"`
	// An array of object returning the details of the locations tracked for the asset
	// during the trip which has ended. Each object represents a single location that
	// was tracked.
	//
	// Please note that this attribute will not be present in the response if no
	// locations were tracked/uploaded during the trip.
	Route []TrackLocation `json:"route"`
	// Returns the time, expressed as UNIX epoch timestamp in milliseconds, when the
	// trip started.
	StartedAt int64 `json:"started_at"`
	// Returns the current state of the trip. The value will be "active" if the trip is
	// still going on, otherwise the value returned would be "ended".
	State string `json:"state"`
	// An array of objects returning the details of the stops made during the trip.
	// Each object represents a single stop.
	Stops []TripStop `json:"stops"`
	// Returns the timeme, expressed as UNIX epoch timestamp in milliseconds, when the
	// trip was last updated.
	UpdatedAt int64 `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssetID     respjson.Field
		Attributes  respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		EndedAt     respjson.Field
		MetaData    respjson.Field
		Name        respjson.Field
		Route       respjson.Field
		StartedAt   respjson.Field
		State       respjson.Field
		Stops       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetTripGetResponseDataTrip) RawJSON() string { return r.JSON.raw }
func (r *SkynetTripGetResponseDataTrip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetTripGetSummaryResponse struct {
	// An container for the trip returned by the service.
	Data SkynetTripGetSummaryResponseData `json:"data"`
	// Displays the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Message string `json:"message"`
	// A string indicating the state of the response. On successful responses, the
	// value will be `Ok`. Indicative error messages are returned for different errors.
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
func (r SkynetTripGetSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetTripGetSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An container for the trip returned by the service.
type SkynetTripGetSummaryResponseData struct {
	// An object containing the returned trip summary.
	Trip SkynetTripGetSummaryResponseDataTrip `json:"trip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Trip        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetTripGetSummaryResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetTripGetSummaryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the returned trip summary.
type SkynetTripGetSummaryResponseDataTrip struct {
	// Returns the unique identifier of the trip.
	ID string `json:"id"`
	// An object with details of the `asset` properties.
	Asset Asset `json:"asset"`
	// Returns the ID of the asset linked to the trip when the trip was started or
	// updated.
	AssetID string `json:"asset_id"`
	// Returns the `attributes` provided for the trip at the time of starting or
	// updating it.
	Attributes any `json:"attributes"`
	// Returns the time, expressed as UNIX epoch timestamp in milliseconds, when the
	// trip was created.
	CreatedAt int64 `json:"created_at"`
	// Returns the custom description for the trip as provided at the time of starting
	// or updating the trip.
	Description string `json:"description"`
	// Returns the total distance covered during the trip, in meters. Please note that
	// this field will be available in the response only if a minimum of 3 locations
	// were tracked during the trip.
	Distance float64 `json:"distance"`
	// Returns the total duration elapsed during the trip, in seconds. Please note that
	// this field will be available in the response only if a minimum of 3 locations
	// were tracked during the trip.
	Duration float64 `json:"duration"`
	// Returns the time, expressed as UNIX epoch timestamp in milliseconds, when the
	// trip ended.
	EndedAt int64 `json:"ended_at"`
	// Returns the geometry of the route that was taken during the trip, encoded in
	// polyline format. Please note that this field will be available in the response
	// only if a minimum of 3 locations were tracked during the trip.
	Geometry []string `json:"geometry"`
	// Returns the metadata containing additional information about the trip as
	// provided at the time of starting or updating the trip.
	MetaData any `json:"meta_data"`
	// Returns the name for the trip as provided at the time of starting or updating
	// the trip.
	Name string `json:"name"`
	// An array of object returning the details of the locations tracked for the asset
	// during the trip which has ended. Each object represents a single location that
	// was tracked.
	//
	// Please note that this attribute will not be present in the response if no
	// locations were tracked/uploaded during the trip.
	Route []TrackLocation `json:"route"`
	// Returns the time, expressed as UNIX epoch timestamp in milliseconds, when the
	// trip started.
	StartedAt int64 `json:"started_at"`
	// Returns the current state of the trip. The value will be "active" if the trip is
	// still going on, otherwise the value returned would be "ended".
	State string `json:"state"`
	// An array of objects returning the details of the stops made during the trip.
	// Each object represents a single stop.
	Stops []TripStop `json:"stops"`
	// Returns the timeme, expressed as UNIX epoch timestamp in milliseconds, when the
	// trip was last updated.
	UpdatedAt int64 `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Asset       respjson.Field
		AssetID     respjson.Field
		Attributes  respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Distance    respjson.Field
		Duration    respjson.Field
		EndedAt     respjson.Field
		Geometry    respjson.Field
		MetaData    respjson.Field
		Name        respjson.Field
		Route       respjson.Field
		StartedAt   respjson.Field
		State       respjson.Field
		Stops       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetTripGetSummaryResponseDataTrip) RawJSON() string { return r.JSON.raw }
func (r *SkynetTripGetSummaryResponseDataTrip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetTripStartResponse struct {
	Data SkynetTripStartResponseData `json:"data"`
	// Displays the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Message string `json:"message"`
	// A string indicating the state of the response. On successful responses, the
	// value will be `Ok`. Indicative error messages are returned for different errors.
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
func (r SkynetTripStartResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetTripStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetTripStartResponseData struct {
	// Returns the ID of the newly created trip. It will be same as the `custom_id` if
	// that input was provided in the input request. Use this ID to manage this trip
	// using other available Trip methods.
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetTripStartResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetTripStartResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetTripGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetTripGetParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetTripGetParams]'s query parameters as `url.Values`.
func (r SkynetTripGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetTripGetParamsCluster string

const (
	SkynetTripGetParamsClusterAmerica SkynetTripGetParamsCluster = "america"
)

type SkynetTripUpdateParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Use this param to update the ID of the asset which made this trip. Please be
	// cautious when using this field as providing an ID other than what was provided
	// at the time of starting the trip, will link a new asset to the trip and un-link
	// the original asset, even if the trip is still active.
	AssetID string `json:"asset_id,required"`
	// Use this parameter to update the custom description of the trip.
	Description param.Opt[string] `json:"description,omitzero"`
	// Use this property to update the name of the trip.
	Name param.Opt[string] `json:"name,omitzero"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetTripUpdateParamsCluster `query:"cluster,omitzero" json:"-"`
	// Use this field to update the `attributes` of the trip. Please note that when
	// updating the `attributes` field, previously added information will be
	// overwritten.
	Attributes any `json:"attributes,omitzero"`
	// Use this JSON object to update additional details about the trip. This property
	// is used to add any custom information / context about the trip.
	//
	// Please note that updating the `meta_data` field will overwrite the previously
	// added information.
	MetaData any `json:"meta_data,omitzero"`
	// Use this object to update the details of the stops made during the trip. Each
	// object represents a single stop.
	//
	// Please note that when updating this field, the new stops will overwrite any
	// existing stops configured for the trip.
	Stops []SkynetTripUpdateParamsStop `json:"stops,omitzero"`
	paramObj
}

func (r SkynetTripUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetTripUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetTripUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetTripUpdateParams]'s query parameters as `url.Values`.
func (r SkynetTripUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetTripUpdateParamsCluster string

const (
	SkynetTripUpdateParamsClusterAmerica SkynetTripUpdateParamsCluster = "america"
)

type SkynetTripUpdateParamsStop struct {
	// Use this parameter to update the ID of the geofence indicating the area where
	// the asset needs to make a stop, during the trip. Only the IDs of geofences
	// created using
	// [NextBillion.ai's Geofence API](https://docs.nextbillion.ai/docs/tracking/api/geofence#create-a-geofence)
	// are accepted.
	//
	// Please note that updating this field will overwrite the previously added
	// information.
	GeofenceID param.Opt[string] `json:"geofence_id,omitzero"`
	// Use this filed to update the name of the stop.
	Name param.Opt[string] `json:"name,omitzero"`
	// Use this JSON object to update additional details about the stop. This property
	// is used to add any custom information / context about the stop.
	//
	// Please note that updating the `meta_data` field will overwrite the previously
	// added information.
	MetaData any `json:"meta_data,omitzero"`
	paramObj
}

func (r SkynetTripUpdateParamsStop) MarshalJSON() (data []byte, err error) {
	type shadow SkynetTripUpdateParamsStop
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetTripUpdateParamsStop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetTripDeleteParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetTripDeleteParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetTripDeleteParams]'s query parameters as `url.Values`.
func (r SkynetTripDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetTripDeleteParamsCluster string

const (
	SkynetTripDeleteParamsClusterAmerica SkynetTripDeleteParamsCluster = "america"
)

type SkynetTripEndParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Specify the ID of the trip to be ended.
	ID string `json:"id,required"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetTripEndParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

func (r SkynetTripEndParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetTripEndParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetTripEndParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetTripEndParams]'s query parameters as `url.Values`.
func (r SkynetTripEndParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetTripEndParamsCluster string

const (
	SkynetTripEndParamsClusterAmerica SkynetTripEndParamsCluster = "america"
)

type SkynetTripGetSummaryParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetTripGetSummaryParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetTripGetSummaryParams]'s query parameters as
// `url.Values`.
func (r SkynetTripGetSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetTripGetSummaryParamsCluster string

const (
	SkynetTripGetSummaryParamsClusterAmerica SkynetTripGetSummaryParamsCluster = "america"
)

type SkynetTripStartParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Specify the ID of the asset which is making this trip. The asset will be linked
	// to this trip.
	AssetID string `json:"asset_id,required"`
	// Set a unique ID for the new `trip`. If not provided, an ID will be automatically
	// generated in UUID format. A valid `custom_id` can contain letters, numbers, “-”,
	// & “\_” only.
	//
	// Please note that the ID of a `trip` can not be changed once it is created.
	CustomID param.Opt[string] `json:"custom_id,omitzero"`
	// Add a custom description for the trip.
	Description param.Opt[string] `json:"description,omitzero"`
	// Specify a name for the trip.
	Name param.Opt[string] `json:"name,omitzero"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetTripStartParamsCluster `query:"cluster,omitzero" json:"-"`
	// `attributes` can be used to store custom information about a trip in
	// `key`:`value` format. Use `attributes` to add any useful information or context
	// to your trips like the driver name, destination etc.
	//
	// Please note that the maximum number of `key`:`value` pairs that can be added to
	// an `attributes` object is 100. Also, the overall size of `attributes` object
	// should not exceed 65kb.
	Attributes any `json:"attributes,omitzero"`
	// An JSON object to collect additional details about the trip. Use this property
	// to add any custom information / context about the trip. The input will be passed
	// on to the response as-is and can be used to display useful information on, for
	// example, a UI app.
	MetaData any `json:"meta_data,omitzero"`
	// An array of objects to collect the details about all the stops that need to be
	// made before the trip ends. Each object represents one stop.
	Stops []SkynetTripStartParamsStop `json:"stops,omitzero"`
	paramObj
}

func (r SkynetTripStartParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetTripStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetTripStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetTripStartParams]'s query parameters as `url.Values`.
func (r SkynetTripStartParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetTripStartParamsCluster string

const (
	SkynetTripStartParamsClusterAmerica SkynetTripStartParamsCluster = "america"
)

type SkynetTripStartParamsStop struct {
	// Specify the ID of the geofence indicating the area where the asset needs to make
	// a stop, during the trip. Only the IDs of geofences created using
	// [NextBillion.ai's Geofence API](https://docs.nextbillion.ai/docs/tracking/api/geofence#create-a-geofence)
	// are accepted.
	GeofenceID param.Opt[string] `json:"geofence_id,omitzero"`
	// Specify a custom name for the stop.
	Name param.Opt[string] `json:"name,omitzero"`
	// An JSON object to collect additional details about the stop. Use this property
	// to add any custom information / context about the stop. The input will be passed
	// on to the response as-is and can be used to display useful information on, for
	// example, a UI app.
	MetaData any `json:"meta_data,omitzero"`
	paramObj
}

func (r SkynetTripStartParamsStop) MarshalJSON() (data []byte, err error) {
	type shadow SkynetTripStartParamsStop
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetTripStartParamsStop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
