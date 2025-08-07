// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apiquery"
	shimjson "github.com/nextbillion-ai/nextbillion-sdk-go/internal/encoding/json"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// RestrictionService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestrictionService] method instead.
type RestrictionService struct {
	Options []option.RequestOption
}

// NewRestrictionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRestrictionService(opts ...option.RequestOption) (r RestrictionService) {
	r = RestrictionService{}
	r.Options = opts
	return
}

// Create a new restriction
func (r *RestrictionService) New(ctx context.Context, restrictionType RestrictionNewParamsRestrictionType, params RestrictionNewParams, opts ...option.RequestOption) (res *RichGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("restrictions/%v", restrictionType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a restriction by id
func (r *RestrictionService) Get(ctx context.Context, id int64, query RestrictionGetParams, opts ...option.RequestOption) (res *RichGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("restrictions/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a restriction
func (r *RestrictionService) Update(ctx context.Context, id int64, params RestrictionUpdateParams, opts ...option.RequestOption) (res *RichGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("restrictions/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get the paginated list of restrictions
func (r *RestrictionService) List(ctx context.Context, query RestrictionListParams, opts ...option.RequestOption) (res *RestrictionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "restrictions/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a restriction by ID
func (r *RestrictionService) Delete(ctx context.Context, id int64, body RestrictionDeleteParams, opts ...option.RequestOption) (res *RestrictionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("restrictions/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Get restrictions by bbox
func (r *RestrictionService) ListByBbox(ctx context.Context, query RestrictionListByBboxParams, opts ...option.RequestOption) (res *[]RichGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "restrictions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Set the state of a restriction by ID
func (r *RestrictionService) SetState(ctx context.Context, id int64, params RestrictionSetStateParams, opts ...option.RequestOption) (res *RichGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("restrictions/%v/state", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// The properties Area, Name are required.
type RichGroupRequestParam struct {
	// Specify the area name. It represents a region where restrictions can be applied.
	// This is a custom field and it is recommended for the users to check with
	// [NextBillion.ai](www.nextbillion.ai) support for the right value. Alternatively,
	// users can invoke the _[Areas](#supported-areas)_ method to get a list of
	// available areas for them.
	Area string `json:"area,required"`
	// Specify a custom, descriptive name for the restriction.
	Name string `json:"name,required"`
	// Use this parameter to add any custom information about the restriction being
	// created.
	Comment param.Opt[string] `json:"comment,omitzero"`
	// Provide a UNIX epoch timestamp in seconds, representing the time when the
	// restriction should cease to be in-effect.
	EndTime param.Opt[float64] `json:"end_time,omitzero"`
	// Specify the maximum truck height, in centimeter, that will be allowed under the
	// restriction. A value of 0 indicates no limit.
	//
	// Please note this parameter is effective only when restriction_type is truck. At
	// least one of truck parameters - weight, height, width and truck - needs to be
	// provided when restriction type is truck.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Specify the maximum truck length, in centimeter, that will be allowed under the
	// restriction. A value of 0 indicates no limit.
	//
	// Please note this parameter is effective only when restriction_type is truck. At
	// least one of truck parameters - weight, height, width and truck - needs to be
	// provided when restriction type is truck.
	Length param.Opt[int64] `json:"length,omitzero"`
	// It represents the days and times when the restriction is in effect. Users can
	// use this property to set recurring or one-time restrictions as per the
	// [given format](https://wiki.openstreetmap.org/wiki/Key:opening_hours) for
	// specifying the recurring schedule of the restriction.
	//
	// Please provided values as per the local time of the region where the restriction
	// is being applied.
	RepeatOn param.Opt[string] `json:"repeat_on,omitzero"`
	// Provide the the fixed speed of the segment where the restriction needs to be
	// applied. Please note that this parameter is mandatory when the restrictionType
	// is fixedspeed.
	Speed param.Opt[float64] `json:"speed,omitzero"`
	// Provide the the maximum speed of the segment where the restriction needs to be
	// applied. Please note that this parameter is mandatory when the restrictionType
	// is maxspeed.
	SpeedLimit param.Opt[float64] `json:"speed_limit,omitzero"`
	// Provide a UNIX epoch timestamp in seconds, representing the start time for the
	// restriction to be in-effect.
	StartTime param.Opt[float64] `json:"start_time,omitzero"`
	// Specify the maximum truck weight, in kilograms, that the restriction will allow.
	// A value of 0 indicates no limit.
	//
	// Please note this parameter is effective only when restriction_type is truck. At
	// least one of truck parameters - weight, height, width and truck - needs to be
	// provided for is truck restriction type.
	Weight param.Opt[int64] `json:"weight,omitzero"`
	// Specify the maximum truck width, in centimeter, that will be allowed under the
	// restriction. A value of 0 indicates no limit.
	//
	// Please note this parameter is effective only when restriction_type is truck. At
	// least one of truck parameters - weight, height, width and truck - needs to be
	// provided when restriction type is truck.
	Width param.Opt[int64] `json:"width,omitzero"`
	// Represents the traffic direction on the segments to which the restriction will
	// be applied.
	//
	// Any of "forward", "backward", "both".
	Direction RichGroupRequestDirection `json:"direction,omitzero"`
	// An array of coordinates denoting the boundary of an area in which the
	// restrictions are to be applied. The format in which coordinates should be listed
	// is defined by the latlon field.
	//
	// Geofences can be used to create all restriction types, except for a turn type
	// restriction. Please note that segments is not required when using geofence to
	// create restrictions.
	Geofence [][]float64 `json:"geofence,omitzero"`
	// Provide the driving modes for which the restriction should be effective. If the
	// value is an empty array or if it is not provided then the restriction would be
	// applied for all modes.
	//
	// Any of "0w", "2w", "3w", "4w", "6w".
	Mode []string `json:"mode,omitzero"`
	// An array of objects to collect the details of the segments of a road on which
	// the restriction has to be applied. Each object corresponds to a new segment.
	//
	// Please note that segments is mandatory for all restrtiction_type except turn.
	Segments []RichGroupRequestSegmentParam `json:"segments,omitzero"`
	// Specify a sequence of coordinates (track) where the restriction is to be
	// applied. The coordinates will be snapped to nearest road. Please note when using
	// tracks, segments and turns are not required.
	Tracks [][]float64 `json:"tracks,omitzero"`
	// An array of objects to collect the details of the turns of a road on which the
	// restriction has to be applied. Each object corresponds to a new turn.
	//
	// Please note that turns is mandatory for when restrtiction_type=turn.
	Turns []RichGroupRequestTurnParam `json:"turns,omitzero"`
	paramObj
}

func (r RichGroupRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RichGroupRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RichGroupRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the traffic direction on the segments to which the restriction will
// be applied.
type RichGroupRequestDirection string

const (
	RichGroupRequestDirectionForward  RichGroupRequestDirection = "forward"
	RichGroupRequestDirectionBackward RichGroupRequestDirection = "backward"
	RichGroupRequestDirectionBoth     RichGroupRequestDirection = "both"
)

type RichGroupRequestSegmentParam struct {
	// An integer value representing the ID of the starting node of the segment.
	From param.Opt[float64] `json:"from,omitzero"`
	// An integer value representing the ID of the ending node of the segment.
	To param.Opt[float64] `json:"to,omitzero"`
	paramObj
}

func (r RichGroupRequestSegmentParam) MarshalJSON() (data []byte, err error) {
	type shadow RichGroupRequestSegmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RichGroupRequestSegmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RichGroupRequestTurnParam struct {
	// An integer value that represents the ID of the starting node of the turn.
	From param.Opt[int64] `json:"from,omitzero"`
	// An integer value that represents the ID of the ending node of the turn.
	To param.Opt[int64] `json:"to,omitzero"`
	// An integer value that represents the ID of a node connecting from and to nodes
	// of the turn.
	Via param.Opt[int64] `json:"via,omitzero"`
	paramObj
}

func (r RichGroupRequestTurnParam) MarshalJSON() (data []byte, err error) {
	type shadow RichGroupRequestTurnParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RichGroupRequestTurnParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RichGroupResponse struct {
	// Returns the unique ID of the restriction. This ID can be used for update,
	// delete, get operations on the restriction using the available API methods.
	ID int64 `json:"id"`
	// Returns the area to which the restriction belongs to.
	Area string `json:"area"`
	// Returns the details of the bounding box containing the restriction.
	Bbox any `json:"bbox"`
	// Returns the comments that were provided for the restriction at the time of
	// creating or updating the request.
	Comment string `json:"comment"`
	// The timestamp at which the restriction was created.
	CreateAt time.Time `json:"create_at" format:"date-time"`
	// Returns the direction of travel on the segments to which the restriction
	// applies.
	//
	// Any of "forward", "backward", "both".
	Direction RichGroupResponseDirection `json:"direction"`
	// The time when the restriction ceases to be in-effect. It is a UNIX timestamp.
	EndTime float64 `json:"end_time"`
	// Returns the list of coordinates representing the area that was used to apply the
	// given restriction. The geofence returned is same as that provided while creating
	// or updating the restriction.
	Geofence any `json:"geofence"`
	// Returns the highway information on which the restriction applies to. If no
	// highway is impacted by the restriction, then this field is not present in the
	// response.
	Highway string `json:"highway"`
	// Returns an array denoting all the traveling modes the restriction applies on.
	Mode []string `json:"mode"`
	// Returns the name of the restriction. This value is same as that provided at the
	// time of creating or updating the restriction.
	Name string `json:"name"`
	// Returns the time periods during which this restriction active or repeats on. The
	// time values follow a
	// [given format](https://wiki.openstreetmap.org/wiki/Key:opening_hours).
	RepeatOn string `json:"repeat_on"`
	// Returns the type of restriction. This is the same value as provided when
	// creating or updating the restriction.
	//
	// Any of "closure", "maxspeed", "fixedspeed", "parking", "turn", "truck".
	RestrictionType RichGroupResponseRestrictionType `json:"restriction_type"`
	// Returns the fixed speed of segments. This field is not present in the response
	// if the restriction type is not fixedspeed
	Speed float64 `json:"speed"`
	// Returns the maximum speed of segments. This field is not present in the response
	// if the restriction type is not maxspeed
	SpeedLimit float64 `json:"speed_limit"`
	// The time when the restriction starts to be in-effect. It is a UNIX timestamp.
	StartTime float64 `json:"start_time"`
	// Returns the state of the "restriction" itself - enabled, disabled or deleted. It
	// does not denote if the restriction is actually in effect or not.
	//
	// Any of "enabled", "disabled", "deleted".
	State RichGroupResponseState `json:"state"`
	// Returns the status of the restriction at the time of making the request i.e.
	// whether the restriction is in force or not. It will have one of the following
	// values: active or inactive.
	//
	// Please note that this field can not be directly influenced by the users. It will
	// always be calculated using the start_time, end_time and repeat_on parameters.
	//
	// Any of "active", "inactive".
	Status RichGroupResponseStatus `json:"status"`
	// The timestamp at which the restriction was updated.
	UpdateAt time.Time `json:"update_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Area            respjson.Field
		Bbox            respjson.Field
		Comment         respjson.Field
		CreateAt        respjson.Field
		Direction       respjson.Field
		EndTime         respjson.Field
		Geofence        respjson.Field
		Highway         respjson.Field
		Mode            respjson.Field
		Name            respjson.Field
		RepeatOn        respjson.Field
		RestrictionType respjson.Field
		Speed           respjson.Field
		SpeedLimit      respjson.Field
		StartTime       respjson.Field
		State           respjson.Field
		Status          respjson.Field
		UpdateAt        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RichGroupResponse) RawJSON() string { return r.JSON.raw }
func (r *RichGroupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the direction of travel on the segments to which the restriction
// applies.
type RichGroupResponseDirection string

const (
	RichGroupResponseDirectionForward  RichGroupResponseDirection = "forward"
	RichGroupResponseDirectionBackward RichGroupResponseDirection = "backward"
	RichGroupResponseDirectionBoth     RichGroupResponseDirection = "both"
)

// Returns the type of restriction. This is the same value as provided when
// creating or updating the restriction.
type RichGroupResponseRestrictionType string

const (
	RichGroupResponseRestrictionTypeClosure    RichGroupResponseRestrictionType = "closure"
	RichGroupResponseRestrictionTypeMaxspeed   RichGroupResponseRestrictionType = "maxspeed"
	RichGroupResponseRestrictionTypeFixedspeed RichGroupResponseRestrictionType = "fixedspeed"
	RichGroupResponseRestrictionTypeParking    RichGroupResponseRestrictionType = "parking"
	RichGroupResponseRestrictionTypeTurn       RichGroupResponseRestrictionType = "turn"
	RichGroupResponseRestrictionTypeTruck      RichGroupResponseRestrictionType = "truck"
)

// Returns the state of the "restriction" itself - enabled, disabled or deleted. It
// does not denote if the restriction is actually in effect or not.
type RichGroupResponseState string

const (
	RichGroupResponseStateEnabled  RichGroupResponseState = "enabled"
	RichGroupResponseStateDisabled RichGroupResponseState = "disabled"
	RichGroupResponseStateDeleted  RichGroupResponseState = "deleted"
)

// Returns the status of the restriction at the time of making the request i.e.
// whether the restriction is in force or not. It will have one of the following
// values: active or inactive.
//
// Please note that this field can not be directly influenced by the users. It will
// always be calculated using the start_time, end_time and repeat_on parameters.
type RichGroupResponseStatus string

const (
	RichGroupResponseStatusActive   RichGroupResponseStatus = "active"
	RichGroupResponseStatusInactive RichGroupResponseStatus = "inactive"
)

type RestrictionListResponse struct {
	// An array of objects containing the details of the restrictions returned. Each
	// object represents one restriction.
	Data []RichGroupResponse         `json:"data"`
	Meta RestrictionListResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestrictionListResponse) RawJSON() string { return r.JSON.raw }
func (r *RestrictionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestrictionListResponseMeta struct {
	// An integer value indicating the maximum number of items retrieved per "page".
	// This is the same number as provided for the limit parameter in input.
	Limit int64 `json:"limit"`
	// An integer value indicating the number of items in the collection that were
	// skipped to display the current response. Please note that the offset starts from
	// zero.
	Offset int64 `json:"offset"`
	// An integer value indicating the total number of items available in the data set.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Offset      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestrictionListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *RestrictionListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestrictionDeleteResponse struct {
	// It is the unique ID of the restriction.
	ID float64 `json:"id"`
	// Returns the state of the restriction. It would always be deleted.
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestrictionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RestrictionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestrictionNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key              string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	RichGroupRequest RichGroupRequestParam
	// Use this parameter to decide the format for specifying the geofence coordinates.
	// If true, then the coordinates of geofence can be specified as
	// "latitude,longitude" format, otherwise they should be specified in
	// "longitude,latitude" format.
	Latlon param.Opt[bool] `query:"latlon,omitzero" json:"-"`
	paramObj
}

func (r RestrictionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RichGroupRequest)
}
func (r *RestrictionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.RichGroupRequest)
}

// URLQuery serializes [RestrictionNewParams]'s query parameters as `url.Values`.
func (r RestrictionNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RestrictionNewParamsRestrictionType string

const (
	RestrictionNewParamsRestrictionTypeTurn       RestrictionNewParamsRestrictionType = "turn"
	RestrictionNewParamsRestrictionTypeParking    RestrictionNewParamsRestrictionType = "parking"
	RestrictionNewParamsRestrictionTypeFixedspeed RestrictionNewParamsRestrictionType = "fixedspeed"
	RestrictionNewParamsRestrictionTypeMaxspeed   RestrictionNewParamsRestrictionType = "maxspeed"
	RestrictionNewParamsRestrictionTypeClosure    RestrictionNewParamsRestrictionType = "closure"
	RestrictionNewParamsRestrictionTypeTruck      RestrictionNewParamsRestrictionType = "truck"
)

type RestrictionGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// a internal parameter
	Transform param.Opt[bool] `query:"transform,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RestrictionGetParams]'s query parameters as `url.Values`.
func (r RestrictionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RestrictionUpdateParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key              string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	RichGroupRequest RichGroupRequestParam
	// Use this parameter to decide the format for specifying the geofence coordinates.
	// If true, then the coordinates of geofence can be specified as
	// "latitude,longitude" format, otherwise they should be specified in
	// "longitude,latitude" format.
	Latlon param.Opt[bool] `query:"latlon,omitzero" json:"-"`
	paramObj
}

func (r RestrictionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RichGroupRequest)
}
func (r *RestrictionUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.RichGroupRequest)
}

// URLQuery serializes [RestrictionUpdateParams]'s query parameters as
// `url.Values`.
func (r RestrictionUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RestrictionListParams struct {
	// Specify the area name. It represents a region where restrictions can be applied.
	//
	// _The area it belongs to. See Area API_
	Area string `query:"area,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// The number of restrictions to be returned in the response. Please note that if
	// the limit is set to a number more than the total number of available
	// restrictions, then all restrictions would be returned together.
	Limit int64 `query:"limit,required" json:"-"`
	// An integer value indicating the number of items in the collection that need to
	// be skipped in the response. Please note that the offset starts from 0, so the
	// first item returned in the result would be the item at (offset + 1) position in
	// collection.
	//
	// Users can use offset along with limit to implement paginated result.
	Offset int64 `query:"offset,required" json:"-"`
	// The name of the restriction. This should be same as that provided while creating
	// or updating the restriction.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// a internal parameter
	Transform param.Opt[bool] `query:"transform,omitzero" json:"-"`
	// Specify the modes of travel that the restriction pertains to.
	//
	// Any of "0w", "2w", "3w", "4w", "6w".
	Mode RestrictionListParamsMode `query:"mode,omitzero" json:"-"`
	// Specify the type of restrictions to fetch.
	//
	// Any of "turn", "parking", "fixedspeed", "maxspeed", "closure", "truck".
	RestrictionType RestrictionListParamsRestrictionType `query:"restriction_type,omitzero" json:"-"`
	// It represents where it comes from, currently the possible values include "rrt",
	// "xsm"
	//
	// Any of "rrt", "pbf".
	Source RestrictionListParamsSource `query:"source,omitzero" json:"-"`
	// This parameter is used to filter restrictions based on their state i.e. whether
	// the restriction is currently enabled, disabled, or deleted. For example, users
	// can retrieve a list of all the deleted restrictions by setting state=deleted.
	//
	// Any of "enabled", "disabled", "deleted".
	State RestrictionListParamsState `query:"state,omitzero" json:"-"`
	// Restrictions can be active or inactive at a given time by virtue of their
	// nature. For example, maximum speed limits can be active on the roads leading to
	// schools during school hours and be inactive afterwards or certain road closure
	// restrictions be active during holidays/concerts and be inactive otherwise.
	//
	// Use this parameter to filter the restrictions based on their status at the time
	// of making the request i.e. whether they are in force or not.
	//
	// Any of "active", "inactive".
	Status RestrictionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RestrictionListParams]'s query parameters as `url.Values`.
func (r RestrictionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the modes of travel that the restriction pertains to.
type RestrictionListParamsMode string

const (
	RestrictionListParamsMode0w RestrictionListParamsMode = "0w"
	RestrictionListParamsMode2w RestrictionListParamsMode = "2w"
	RestrictionListParamsMode3w RestrictionListParamsMode = "3w"
	RestrictionListParamsMode4w RestrictionListParamsMode = "4w"
	RestrictionListParamsMode6w RestrictionListParamsMode = "6w"
)

// Specify the type of restrictions to fetch.
type RestrictionListParamsRestrictionType string

const (
	RestrictionListParamsRestrictionTypeTurn       RestrictionListParamsRestrictionType = "turn"
	RestrictionListParamsRestrictionTypeParking    RestrictionListParamsRestrictionType = "parking"
	RestrictionListParamsRestrictionTypeFixedspeed RestrictionListParamsRestrictionType = "fixedspeed"
	RestrictionListParamsRestrictionTypeMaxspeed   RestrictionListParamsRestrictionType = "maxspeed"
	RestrictionListParamsRestrictionTypeClosure    RestrictionListParamsRestrictionType = "closure"
	RestrictionListParamsRestrictionTypeTruck      RestrictionListParamsRestrictionType = "truck"
)

// It represents where it comes from, currently the possible values include "rrt",
// "xsm"
type RestrictionListParamsSource string

const (
	RestrictionListParamsSourceRrt RestrictionListParamsSource = "rrt"
	RestrictionListParamsSourcePbf RestrictionListParamsSource = "pbf"
)

// This parameter is used to filter restrictions based on their state i.e. whether
// the restriction is currently enabled, disabled, or deleted. For example, users
// can retrieve a list of all the deleted restrictions by setting state=deleted.
type RestrictionListParamsState string

const (
	RestrictionListParamsStateEnabled  RestrictionListParamsState = "enabled"
	RestrictionListParamsStateDisabled RestrictionListParamsState = "disabled"
	RestrictionListParamsStateDeleted  RestrictionListParamsState = "deleted"
)

// Restrictions can be active or inactive at a given time by virtue of their
// nature. For example, maximum speed limits can be active on the roads leading to
// schools during school hours and be inactive afterwards or certain road closure
// restrictions be active during holidays/concerts and be inactive otherwise.
//
// Use this parameter to filter the restrictions based on their status at the time
// of making the request i.e. whether they are in force or not.
type RestrictionListParamsStatus string

const (
	RestrictionListParamsStatusActive   RestrictionListParamsStatus = "active"
	RestrictionListParamsStatusInactive RestrictionListParamsStatus = "inactive"
)

type RestrictionDeleteParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [RestrictionDeleteParams]'s query parameters as
// `url.Values`.
func (r RestrictionDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RestrictionListByBboxParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Specifies the maximum latitude value for the bounding box.
	MaxLat float64 `query:"max_lat,required" json:"-"`
	// Specifies the maximum longitude value for the bounding box.
	MaxLon float64 `query:"max_lon,required" json:"-"`
	// Specifies the minimum latitude value for the bounding box.
	MinLat float64 `query:"min_lat,required" json:"-"`
	// Specifies the minimum longitude value for the bounding box.
	MinLon float64 `query:"min_lon,required" json:"-"`
	// This is internal parameter with a default value as false.
	Transform param.Opt[bool] `query:"transform,omitzero" json:"-"`
	// Specify the modes of travel that the restriction pertains to.
	//
	// Any of "0w", "2w", "3w", "4w", "6w".
	Mode []string `query:"mode,omitzero" json:"-"`
	// Specify the type of restrictions to fetch.
	//
	// Any of "turn", "parking", "fixedspeed", "maxspeed", "closure", "truck".
	RestrictionType RestrictionListByBboxParamsRestrictionType `query:"restriction_type,omitzero" json:"-"`
	// This parameter represents where the restriction comes from and cannot be
	// modified by clients sending requests to the API endpoint.
	//
	// For example, an API endpoint that returns a list of restrictions could include
	// the source parameter to indicate where each item comes from. This parameter can
	// be useful for filtering, sorting, or grouping the results based on their source.
	//
	// Any of "rrt", "pbf".
	Source RestrictionListByBboxParamsSource `query:"source,omitzero" json:"-"`
	// This parameter is used to filter restrictions based on their state i.e. whether
	// the restriction is currently enabled, disabled, or deleted. For example, users
	// can retrieve a list of all the deleted restrictions by setting state=deleted.
	//
	// Any of "enabled", "disabled", "deleted".
	State RestrictionListByBboxParamsState `query:"state,omitzero" json:"-"`
	// Restrictions can be active or inactive at a given time by virtue of their
	// nature. For example, maximum speed limits can be active on the roads leading to
	// schools during school hours and be inactive afterwards or certain road closure
	// restrictions be active during holidays/concerts and be inactive otherwise.
	//
	// Use this parameter to filter the restrictions based on their status at the time
	// of making the request i.e. whether they are in force or not.
	//
	// Any of "active", "inactive".
	Status RestrictionListByBboxParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RestrictionListByBboxParams]'s query parameters as
// `url.Values`.
func (r RestrictionListByBboxParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the type of restrictions to fetch.
type RestrictionListByBboxParamsRestrictionType string

const (
	RestrictionListByBboxParamsRestrictionTypeTurn       RestrictionListByBboxParamsRestrictionType = "turn"
	RestrictionListByBboxParamsRestrictionTypeParking    RestrictionListByBboxParamsRestrictionType = "parking"
	RestrictionListByBboxParamsRestrictionTypeFixedspeed RestrictionListByBboxParamsRestrictionType = "fixedspeed"
	RestrictionListByBboxParamsRestrictionTypeMaxspeed   RestrictionListByBboxParamsRestrictionType = "maxspeed"
	RestrictionListByBboxParamsRestrictionTypeClosure    RestrictionListByBboxParamsRestrictionType = "closure"
	RestrictionListByBboxParamsRestrictionTypeTruck      RestrictionListByBboxParamsRestrictionType = "truck"
)

// This parameter represents where the restriction comes from and cannot be
// modified by clients sending requests to the API endpoint.
//
// For example, an API endpoint that returns a list of restrictions could include
// the source parameter to indicate where each item comes from. This parameter can
// be useful for filtering, sorting, or grouping the results based on their source.
type RestrictionListByBboxParamsSource string

const (
	RestrictionListByBboxParamsSourceRrt RestrictionListByBboxParamsSource = "rrt"
	RestrictionListByBboxParamsSourcePbf RestrictionListByBboxParamsSource = "pbf"
)

// This parameter is used to filter restrictions based on their state i.e. whether
// the restriction is currently enabled, disabled, or deleted. For example, users
// can retrieve a list of all the deleted restrictions by setting state=deleted.
type RestrictionListByBboxParamsState string

const (
	RestrictionListByBboxParamsStateEnabled  RestrictionListByBboxParamsState = "enabled"
	RestrictionListByBboxParamsStateDisabled RestrictionListByBboxParamsState = "disabled"
	RestrictionListByBboxParamsStateDeleted  RestrictionListByBboxParamsState = "deleted"
)

// Restrictions can be active or inactive at a given time by virtue of their
// nature. For example, maximum speed limits can be active on the roads leading to
// schools during school hours and be inactive afterwards or certain road closure
// restrictions be active during holidays/concerts and be inactive otherwise.
//
// Use this parameter to filter the restrictions based on their status at the time
// of making the request i.e. whether they are in force or not.
type RestrictionListByBboxParamsStatus string

const (
	RestrictionListByBboxParamsStatusActive   RestrictionListByBboxParamsStatus = "active"
	RestrictionListByBboxParamsStatusInactive RestrictionListByBboxParamsStatus = "inactive"
)

type RestrictionSetStateParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Use this field to specify the new state of the restriction. Please note that
	// this method cannot update the state of restrictions that are currently in
	// 'deleted' state.
	//
	// Any of "enabled", "disabled", "deleted".
	State RestrictionSetStateParamsState `json:"state,omitzero,required"`
	paramObj
}

func (r RestrictionSetStateParams) MarshalJSON() (data []byte, err error) {
	type shadow RestrictionSetStateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestrictionSetStateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RestrictionSetStateParams]'s query parameters as
// `url.Values`.
func (r RestrictionSetStateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Use this field to specify the new state of the restriction. Please note that
// this method cannot update the state of restrictions that are currently in
// 'deleted' state.
type RestrictionSetStateParamsState string

const (
	RestrictionSetStateParamsStateEnabled  RestrictionSetStateParamsState = "enabled"
	RestrictionSetStateParamsStateDisabled RestrictionSetStateParamsState = "disabled"
	RestrictionSetStateParamsStateDeleted  RestrictionSetStateParamsState = "deleted"
)
