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

// GeofenceBatchService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeofenceBatchService] method instead.
type GeofenceBatchService struct {
	Options []option.RequestOption
}

// NewGeofenceBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGeofenceBatchService(opts ...option.RequestOption) (r GeofenceBatchService) {
	r = GeofenceBatchService{}
	r.Options = opts
	return
}

// Batch Creation of Geofence
func (r *GeofenceBatchService) New(ctx context.Context, params GeofenceBatchNewParams, opts ...option.RequestOption) (res *GeofenceBatchNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "geofence/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete Batch Geofence
func (r *GeofenceBatchService) Delete(ctx context.Context, params GeofenceBatchDeleteParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = append(r.Options[:], opts...)
	path := "geofence/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Batch Query of Geofence
func (r *GeofenceBatchService) Query(ctx context.Context, query GeofenceBatchQueryParams, opts ...option.RequestOption) (res *GeofenceBatchQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "geofence/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type GeofenceBatchNewResponse struct {
	// A data object containing the IDs of the geofences created.
	Data GeofenceBatchNewResponseData `json:"data"`
	// A string indicating the state of the response. On successful responses, the
	// value will be `Ok`. Indicative error messages are returned for different errors.
	// See the [API Error Codes](#api-error-codes) section below for more information.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceBatchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GeofenceBatchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A data object containing the IDs of the geofences created.
type GeofenceBatchNewResponseData struct {
	IDs []string `json:"ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IDs         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceBatchNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *GeofenceBatchNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceBatchQueryResponse struct {
	Data GeofenceBatchQueryResponseData `json:"data,required"`
	// A string indicating the state of the response. On successful responses, the
	// value will be `Ok`. Indicative error messages are returned for different errors.
	// See the [API Error Codes](#api-error-codes) section below for more information.
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceBatchQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *GeofenceBatchQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceBatchQueryResponseData struct {
	// An array of objects containing the details of the geofences returned matching
	// the IDs provided. Each object represents one geofence.
	List []Geofence `json:"list,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		List        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceBatchQueryResponseData) RawJSON() string { return r.JSON.raw }
func (r *GeofenceBatchQueryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceBatchNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// An array of objects to collect the details of the multiple geofences that need
	// to be created.
	Geofences []GeofenceEntityCreateParam `json:"geofences,omitzero"`
	paramObj
}

func (r GeofenceBatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceBatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GeofenceBatchNewParams]'s query parameters as `url.Values`.
func (r GeofenceBatchNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeofenceBatchDeleteParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// An array IDs of the geofence to be deleted. These are the IDs that were
	// generated/provided at the time of creating the respective geofences.
	IDs []string `json:"ids,omitzero"`
	paramObj
}

func (r GeofenceBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceBatchDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GeofenceBatchDeleteParams]'s query parameters as
// `url.Values`.
func (r GeofenceBatchDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeofenceBatchQueryParams struct {
	// Comma(`,`) separated list of IDs of the geofences to be searched.
	IDs string `query:"ids,required" format:"ID_1,ID_2,ID_3,...." json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [GeofenceBatchQueryParams]'s query parameters as
// `url.Values`.
func (r GeofenceBatchQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
