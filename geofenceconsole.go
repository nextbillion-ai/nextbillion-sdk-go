// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apiquery"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// GeofenceConsoleService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeofenceConsoleService] method instead.
type GeofenceConsoleService struct {
	Options []option.RequestOption
}

// NewGeofenceConsoleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGeofenceConsoleService(opts ...option.RequestOption) (r GeofenceConsoleService) {
	r = GeofenceConsoleService{}
	r.Options = opts
	return
}

// preview geofence geojson
func (r *GeofenceConsoleService) Preview(ctx context.Context, body GeofenceConsolePreviewParams, opts ...option.RequestOption) (res *GeofenceConsolePreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "geofence/console/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Console Geofence Search API
func (r *GeofenceConsoleService) Search(ctx context.Context, query GeofenceConsoleSearchParams, opts ...option.RequestOption) (res *GeofenceConsoleSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "geofence/console/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// An object with geoJSON details of the geofence. The contents of this object
// follow the [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
type PolygonGeojson struct {
	// An array of coordinates in the [longitude, latitude] format, representing the
	// geofence boundary.
	Coordinates [][]float64 `json:"coordinates"`
	// Type of the geoJSON geometry. Will always be `Polygon`.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolygonGeojson) RawJSON() string { return r.JSON.raw }
func (r *PolygonGeojson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceConsolePreviewResponse struct {
	Data    GeofenceConsolePreviewResponseData `json:"data"`
	Message string                             `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceConsolePreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *GeofenceConsolePreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceConsolePreviewResponseData struct {
	// An object with geoJSON details of the geofence. The contents of this object
	// follow the [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Geojson PolygonGeojson `json:"geojson"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geojson     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceConsolePreviewResponseData) RawJSON() string { return r.JSON.raw }
func (r *GeofenceConsolePreviewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceConsoleSearchResponse struct {
	Data GeofenceConsoleSearchResponseData `json:"data,required"`
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
func (r GeofenceConsoleSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *GeofenceConsoleSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceConsoleSearchResponseData struct {
	Result []GeofenceConsoleSearchResponseDataResult `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceConsoleSearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *GeofenceConsoleSearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceConsoleSearchResponseDataResult struct {
	// ID of goefence. Could be empty string.
	ID string `json:"id,required"`
	// Name of goefence. Could be empty string.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceConsoleSearchResponseDataResult) RawJSON() string { return r.JSON.raw }
func (r *GeofenceConsoleSearchResponseDataResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceConsolePreviewParams struct {
	GeofenceEntityCreate GeofenceEntityCreateParam
	paramObj
}

func (r GeofenceConsolePreviewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.GeofenceEntityCreate)
}
func (r *GeofenceConsolePreviewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.GeofenceEntityCreate)
}

type GeofenceConsoleSearchParams struct {
	// string to be searched, will used to match name or id of geofence.
	Query string `query:"query,required" format:"key1=value&key" json:"-"`
	paramObj
}

// URLQuery serializes [GeofenceConsoleSearchParams]'s query parameters as
// `url.Values`.
func (r GeofenceConsoleSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
