// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apiquery"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// SkynetAssetEventService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetAssetEventService] method instead.
type SkynetAssetEventService struct {
	Options []option.RequestOption
}

// NewSkynetAssetEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSkynetAssetEventService(opts ...option.RequestOption) (r SkynetAssetEventService) {
	r = SkynetAssetEventService{}
	r.Options = opts
	return
}

// Event History of an Asset
func (r *SkynetAssetEventService) List(ctx context.Context, id string, query SkynetAssetEventListParams, opts ...option.RequestOption) (res *SkynetAssetEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/asset/%s/event/list", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SkynetAssetEventListResponse struct {
	// An object containing the information about the event history for the requested
	// `asset`.
	Data SkynetAssetEventListResponseData `json:"data"`
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
func (r SkynetAssetEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the information about the event history for the requested
// `asset`.
type SkynetAssetEventListResponseData struct {
	// An array of objects with each object on the list representing one event.
	List []SkynetAssetEventListResponseDataList `json:"list"`
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
func (r SkynetAssetEventListResponseData) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetEventListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetEventListResponseDataList struct {
	// ID of the `asset`. This is the same ID that was generated/provided at the time
	// of creating the `asset`.
	AssetID string `json:"asset_id"`
	// Nature of the event triggered by the `asset`. It can have following values:
	//
	// - `enter`: When the `asset` enters a specific geofence
	//
	// - `exit`: When the `asset` moves out of a specific geofence.
	//
	// - `speeding`: When the `asset` exceeds the certain speed limit.
	//
	// - `idle`: When the `asset` exhibits idle or no activity.
	//
	// Any of "`enter`", "`exit`", "`speeding`", "`idle`".
	EventType string `json:"event_type"`
	// Additional information about the event. Currently, this object returns the speed
	// limit that was used to generate the over-speeding events, for a `speeding` type
	// event.
	//
	// It is worth highlighting that, when the `use_admin_speed_limit` is `true`, the
	// speed limit value will be obtained from the underlying road information.
	// Whereas, if the `use_admin_speed_limit` is `false`, the speed limit will be
	// equal to the `customer_speed_limit` value provided by the user when creating or
	// updating the `monitor`.
	Extra any `json:"extra"`
	// ID of the `geofence` associated with the event.
	GeofenceID string `json:"geofence_id"`
	// ID of the `monitor` associated with the event.
	MonitorID string `json:"monitor_id"`
	// Tags associated with the `monitor`.
	MonitorTags []string `json:"monitor_tags"`
	// An object with details of the `asset` at the last tracked location before the
	// event was triggered.
	PrevLocation SkynetAssetEventListResponseDataListPrevLocation `json:"prev_location"`
	// A UNIX epoch timestamp in milliseconds representing the time at which the event
	// was added/created.
	Timestamp int64 `json:"timestamp"`
	// An object with details of the `asset` at the location where the event was
	// triggered.
	TriggeredLocation SkynetAssetEventListResponseDataListTriggeredLocation `json:"triggered_location"`
	// A UNIX epoch timestamp in milliseconds representing the time at which the event
	// was triggered.
	TriggeredTimestamp int64 `json:"triggered_timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetID            respjson.Field
		EventType          respjson.Field
		Extra              respjson.Field
		GeofenceID         respjson.Field
		MonitorID          respjson.Field
		MonitorTags        respjson.Field
		PrevLocation       respjson.Field
		Timestamp          respjson.Field
		TriggeredLocation  respjson.Field
		TriggeredTimestamp respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetEventListResponseDataList) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetEventListResponseDataList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the `asset` at the last tracked location before the
// event was triggered.
type SkynetAssetEventListResponseDataListPrevLocation struct {
	// If available, this property returns the heading of the `asset` from true north
	// in clockwise direction, at the `prev_location` tracked for the `asset`.
	Bearing float64 `json:"bearing"`
	// `prev_location` information of the `asset`.
	Location SkynetAssetEventListResponseDataListPrevLocationLocation `json:"location"`
	// Returns the custom data added during the location information upload.
	MetaData any `json:"meta_data"`
	// If available, this property returns the speed of the `asset`, in meters per
	// second, at the `prev_location` of the `asset`.
	Speed float64 `json:"speed"`
	// A UNIX epoch timestamp in milliseconds representing the time at which the
	// `asset` was at the `prev_location`.
	Timestamp int64 `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bearing     respjson.Field
		Location    respjson.Field
		MetaData    respjson.Field
		Speed       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetEventListResponseDataListPrevLocation) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetEventListResponseDataListPrevLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `prev_location` information of the `asset`.
type SkynetAssetEventListResponseDataListPrevLocationLocation struct {
	// Latitude of the `prev_location` tracked for the `asset`.
	Lat float64 `json:"lat"`
	// Longitude of the `prev_location` tracked for the `asset`.
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
func (r SkynetAssetEventListResponseDataListPrevLocationLocation) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetEventListResponseDataListPrevLocationLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the `asset` at the location where the event was
// triggered.
type SkynetAssetEventListResponseDataListTriggeredLocation struct {
	// If available, this property returns the heading of the `asset` from true north
	// in clockwise direction, when the event was triggered.
	Bearing float64 `json:"bearing"`
	// An object with information about the location at which the event was triggered.
	Location SkynetAssetEventListResponseDataListTriggeredLocationLocation `json:"location"`
	// Returns the custom data added during the location information upload.
	MetaData any `json:"meta_data"`
	// If available, this property returns the speed of the `asset`, in meters per
	// second, when the event was triggered.
	Speed float64 `json:"speed"`
	// A UNIX epoch timestamp in milliseconds representing the time at which the
	// `asset` was at the `triggered_location`.
	Timestamp int64 `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bearing     respjson.Field
		Location    respjson.Field
		MetaData    respjson.Field
		Speed       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetAssetEventListResponseDataListTriggeredLocation) RawJSON() string { return r.JSON.raw }
func (r *SkynetAssetEventListResponseDataListTriggeredLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with information about the location at which the event was triggered.
type SkynetAssetEventListResponseDataListTriggeredLocationLocation struct {
	// Latitude of the `triggered_location` of the event.
	Lat float64 `json:"lat"`
	// Longitude of the `triggered_location` of the event.
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
func (r SkynetAssetEventListResponseDataListTriggeredLocationLocation) RawJSON() string {
	return r.JSON.raw
}
func (r *SkynetAssetEventListResponseDataListTriggeredLocationLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetAssetEventListParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Time before which the events triggered by the `asset` need to be retrieved.
	EndTime param.Opt[int64] `query:"end_time,omitzero" json:"-"`
	// Filter the events by `monitor_id`. When provided, only the events triggered by
	// the `monitor` will be returned in response.
	//
	// Please note that if the `attributes` of the asset identified by `id` and those
	// of the `monitor` do not match, then no events might be returned for this
	// `monitor_id`.
	MonitorID param.Opt[string] `query:"monitor_id,omitzero" json:"-"`
	// Denotes page number. Use this along with the `ps` parameter to implement
	// pagination for your searched results. This parameter does not have a maximum
	// limit but would return an empty response in case a higher value is provided when
	// the result-set itself is smaller.
	Pn param.Opt[int64] `query:"pn,omitzero" json:"-"`
	// Denotes number of search results per page. Use this along with the `pn`
	// parameter to implement pagination for your searched results.
	Ps param.Opt[int64] `query:"ps,omitzero" json:"-"`
	// Time after which the events triggered by the `asset` need to be retrieved.
	StartTime param.Opt[int64] `query:"start_time,omitzero" json:"-"`
	// the cluster of the region you want to use
	//
	// Any of "america".
	Cluster SkynetAssetEventListParamsCluster `query:"cluster,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetAssetEventListParams]'s query parameters as
// `url.Values`.
func (r SkynetAssetEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// the cluster of the region you want to use
type SkynetAssetEventListParamsCluster string

const (
	SkynetAssetEventListParamsClusterAmerica SkynetAssetEventListParamsCluster = "america"
)
