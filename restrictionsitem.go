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

// RestrictionsItemService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestrictionsItemService] method instead.
type RestrictionsItemService struct {
	Options []option.RequestOption
}

// NewRestrictionsItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRestrictionsItemService(opts ...option.RequestOption) (r RestrictionsItemService) {
	r = RestrictionsItemService{}
	r.Options = opts
	return
}

// Get restriction items by bbox
func (r *RestrictionsItemService) List(ctx context.Context, query RestrictionsItemListParams, opts ...option.RequestOption) (res *[]RestrictionsItemListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "restrictions_items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RestrictionsItemListResponse struct {
	ID         float64                                `json:"id" api:"required"`
	Area       string                                 `json:"area" api:"required"`
	Coordinate RestrictionsItemListResponseCoordinate `json:"coordinate" api:"required"`
	GroupID    float64                                `json:"group_id" api:"required"`
	// Any of "segment", "turn".
	GroupType RestrictionsItemListResponseGroupType `json:"group_type" api:"required"`
	// Any of "0w", "1w", "2w", "3w", "4w", "6w".
	Mode     []string `json:"mode" api:"required"`
	RepeatOn string   `json:"repeat_on" api:"required"`
	// Any of "closure", "fixedspeed", "maxspeed", "turn", "truck".
	RestrictionType RestrictionsItemListResponseRestrictionType `json:"restriction_type" api:"required"`
	// Any of "rrt", "pbf".
	Source RestrictionsItemListResponseSource `json:"source" api:"required"`
	// Any of "enabled", "disabled", "deleted".
	State RestrictionsItemListResponseState `json:"state" api:"required"`
	// Any of "active", "inactive".
	Status RestrictionsItemListResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Area            respjson.Field
		Coordinate      respjson.Field
		GroupID         respjson.Field
		GroupType       respjson.Field
		Mode            respjson.Field
		RepeatOn        respjson.Field
		RestrictionType respjson.Field
		Source          respjson.Field
		State           respjson.Field
		Status          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestrictionsItemListResponse) RawJSON() string { return r.JSON.raw }
func (r *RestrictionsItemListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestrictionsItemListResponseCoordinate struct {
	Lat float64 `json:"lat"`
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
func (r RestrictionsItemListResponseCoordinate) RawJSON() string { return r.JSON.raw }
func (r *RestrictionsItemListResponseCoordinate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestrictionsItemListResponseGroupType string

const (
	RestrictionsItemListResponseGroupTypeSegment RestrictionsItemListResponseGroupType = "segment"
	RestrictionsItemListResponseGroupTypeTurn    RestrictionsItemListResponseGroupType = "turn"
)

type RestrictionsItemListResponseRestrictionType string

const (
	RestrictionsItemListResponseRestrictionTypeClosure    RestrictionsItemListResponseRestrictionType = "closure"
	RestrictionsItemListResponseRestrictionTypeFixedspeed RestrictionsItemListResponseRestrictionType = "fixedspeed"
	RestrictionsItemListResponseRestrictionTypeMaxspeed   RestrictionsItemListResponseRestrictionType = "maxspeed"
	RestrictionsItemListResponseRestrictionTypeTurn       RestrictionsItemListResponseRestrictionType = "turn"
	RestrictionsItemListResponseRestrictionTypeTruck      RestrictionsItemListResponseRestrictionType = "truck"
)

type RestrictionsItemListResponseSource string

const (
	RestrictionsItemListResponseSourceRrt RestrictionsItemListResponseSource = "rrt"
	RestrictionsItemListResponseSourcePbf RestrictionsItemListResponseSource = "pbf"
)

type RestrictionsItemListResponseState string

const (
	RestrictionsItemListResponseStateEnabled  RestrictionsItemListResponseState = "enabled"
	RestrictionsItemListResponseStateDisabled RestrictionsItemListResponseState = "disabled"
	RestrictionsItemListResponseStateDeleted  RestrictionsItemListResponseState = "deleted"
)

type RestrictionsItemListResponseStatus string

const (
	RestrictionsItemListResponseStatusActive   RestrictionsItemListResponseStatus = "active"
	RestrictionsItemListResponseStatusInactive RestrictionsItemListResponseStatus = "inactive"
)

type RestrictionsItemListParams struct {
	MaxLat  float64            `query:"max_lat" api:"required" json:"-"`
	MaxLon  float64            `query:"max_lon" api:"required" json:"-"`
	MinLat  float64            `query:"min_lat" api:"required" json:"-"`
	MinLon  float64            `query:"min_lon" api:"required" json:"-"`
	GroupID param.Opt[float64] `query:"group_id,omitzero" json:"-"`
	Source  param.Opt[string]  `query:"source,omitzero" json:"-"`
	// Any of "0w", "1w", "2w", "3w", "4w", "6w".
	Mode RestrictionsItemListParamsMode `query:"mode,omitzero" json:"-"`
	// Any of "turn", "parking", "fixedspeed", "maxspeed", "closure", "truck".
	RestrictionType RestrictionsItemListParamsRestrictionType `query:"restriction_type,omitzero" json:"-"`
	// Any of "enabled", "disabled", "deleted".
	State RestrictionsItemListParamsState `query:"state,omitzero" json:"-"`
	// Any of "active", "inactive".
	Status RestrictionsItemListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RestrictionsItemListParams]'s query parameters as
// `url.Values`.
func (r RestrictionsItemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RestrictionsItemListParamsMode string

const (
	RestrictionsItemListParamsMode0w RestrictionsItemListParamsMode = "0w"
	RestrictionsItemListParamsMode1w RestrictionsItemListParamsMode = "1w"
	RestrictionsItemListParamsMode2w RestrictionsItemListParamsMode = "2w"
	RestrictionsItemListParamsMode3w RestrictionsItemListParamsMode = "3w"
	RestrictionsItemListParamsMode4w RestrictionsItemListParamsMode = "4w"
	RestrictionsItemListParamsMode6w RestrictionsItemListParamsMode = "6w"
)

type RestrictionsItemListParamsRestrictionType string

const (
	RestrictionsItemListParamsRestrictionTypeTurn       RestrictionsItemListParamsRestrictionType = "turn"
	RestrictionsItemListParamsRestrictionTypeParking    RestrictionsItemListParamsRestrictionType = "parking"
	RestrictionsItemListParamsRestrictionTypeFixedspeed RestrictionsItemListParamsRestrictionType = "fixedspeed"
	RestrictionsItemListParamsRestrictionTypeMaxspeed   RestrictionsItemListParamsRestrictionType = "maxspeed"
	RestrictionsItemListParamsRestrictionTypeClosure    RestrictionsItemListParamsRestrictionType = "closure"
	RestrictionsItemListParamsRestrictionTypeTruck      RestrictionsItemListParamsRestrictionType = "truck"
)

type RestrictionsItemListParamsState string

const (
	RestrictionsItemListParamsStateEnabled  RestrictionsItemListParamsState = "enabled"
	RestrictionsItemListParamsStateDisabled RestrictionsItemListParamsState = "disabled"
	RestrictionsItemListParamsStateDeleted  RestrictionsItemListParamsState = "deleted"
)

type RestrictionsItemListParamsStatus string

const (
	RestrictionsItemListParamsStatusActive   RestrictionsItemListParamsStatus = "active"
	RestrictionsItemListParamsStatusInactive RestrictionsItemListParamsStatus = "inactive"
)
