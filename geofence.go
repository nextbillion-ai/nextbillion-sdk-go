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

// GeofenceService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeofenceService] method instead.
type GeofenceService struct {
	Options []option.RequestOption
	Console GeofenceConsoleService
	Batch   GeofenceBatchService
}

// NewGeofenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGeofenceService(opts ...option.RequestOption) (r GeofenceService) {
	r = GeofenceService{}
	r.Options = opts
	r.Console = NewGeofenceConsoleService(opts...)
	r.Batch = NewGeofenceBatchService(opts...)
	return
}

// Create a geofence
func (r *GeofenceService) New(ctx context.Context, params GeofenceNewParams, opts ...option.RequestOption) (res *GeofenceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "geofence"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a Geofence
func (r *GeofenceService) Get(ctx context.Context, id string, query GeofenceGetParams, opts ...option.RequestOption) (res *GeofenceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("geofence/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a Geofence
func (r *GeofenceService) Update(ctx context.Context, id string, params GeofenceUpdateParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("geofence/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Get Geofence List
func (r *GeofenceService) List(ctx context.Context, query GeofenceListParams, opts ...option.RequestOption) (res *GeofenceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "geofence/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a Geofence
func (r *GeofenceService) Delete(ctx context.Context, id string, body GeofenceDeleteParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("geofence/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Geofence Contains
func (r *GeofenceService) Contains(ctx context.Context, query GeofenceContainsParams, opts ...option.RequestOption) (res *GeofenceContainsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "geofence/contain"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// An object with details of the geofence.
type Geofence struct {
	// ID of the geofence provided/generated at the time of creating the geofence.
	ID           string               `json:"id"`
	CircleCenter GeofenceCircleCenter `json:"circle_center"`
	// When the `type` of the geofence is `circle`, this property returns the radius of
	// the geofence in meters (m).
	CircleRadius float64 `json:"circle_radius"`
	// Time at which the geofence was created, expressed as a UNIX timestamp in
	// seconds.
	CreatedAt int64 `json:"created_at"`
	// An object with geoJSON details of the geofence. The contents of this object
	// follow the [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Geojson PolygonGeojson `json:"geojson"`
	// For a geofence based on isochrone contour determined using a specific driving
	// distance, this property returns the duration value, in meters.
	//
	// The value would be the same as that provided for the `contours_meter` parameter
	// at the time of creating or updating the geofence.
	IcContoursMeter int64 `json:"ic_contours_meter"`
	// For a geofence based on isochrone contour determined using a specific driving
	// duration, this property returns the duration value, in minutes. The value would
	// be the same as the value provided for the `contours_minute` parameter at the
	// time of creating or updating the geofence.
	IcContoursMinute int64 `json:"ic_contours_minute"`
	// For a geofence based on isochrone contour, this property returns the coordinates
	// of the location, in [latitude,longitude] format, which was used as the starting
	// point to identify the geofence boundary.
	//
	// The value would be the same as that provided for the `coordinates` parameter at
	// the time of creating or updating the geofence.
	IcCoordinates string `json:"ic_coordinates" format:"latitude,longitude"`
	// For a geofence based on isochrone contour, this property returns the denoise
	// value which would be the same as that provided for the `denoise` parameter at
	// the time of creating or updating the geofence.
	IcDenoise float64 `json:"ic_denoise"`
	// For a geofence based on isochrone contour, this property returns the departure
	// time, as a UNIX epoch timestamp in seconds, which was used to determine the
	// geofence boundary after taking into account the traffic conditions at the time.
	//
	// The value would be the same as that provided for the `departure_time` parameter
	// at the time of creating or updating the geofence.
	IcDepartureTime int64 `json:"ic_departure_time"`
	// For a geofence based on isochrone contour, this property returns the driving
	// mode used to determine the geofence boundary.
	//
	// The value would be the same as that provided for the `mode` parameter at the
	// time of creating or updating the geofence.
	IcMode float64 `json:"ic_mode"`
	// Metadata of the geofence added at the time of creating or updating it.
	MetaData any `json:"meta_data"`
	// Name of the geofence added at the time of creating or updating it.
	Name string `json:"name"`
	// An array of strings representing the `tags` associated with the geofence added
	// at the time of creating or updating it.
	Tags []string `json:"tags"`
	// Type of the geofence.
	//
	// Any of "`circle`", "`polygon`", "`isochrone`".
	Type GeofenceType `json:"type"`
	// Time at which the geofence was last updated, expressed as a UNIX timestamp in
	// seconds.
	UpdatedAt int64 `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CircleCenter     respjson.Field
		CircleRadius     respjson.Field
		CreatedAt        respjson.Field
		Geojson          respjson.Field
		IcContoursMeter  respjson.Field
		IcContoursMinute respjson.Field
		IcCoordinates    respjson.Field
		IcDenoise        respjson.Field
		IcDepartureTime  respjson.Field
		IcMode           respjson.Field
		MetaData         respjson.Field
		Name             respjson.Field
		Tags             respjson.Field
		Type             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Geofence) RawJSON() string { return r.JSON.raw }
func (r *Geofence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceCircleCenter struct {
	// Latitude of the location.
	Lat float64 `json:"lat"`
	// Longitude of the location.
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
func (r GeofenceCircleCenter) RawJSON() string { return r.JSON.raw }
func (r *GeofenceCircleCenter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the geofence.
type GeofenceType string

const (
	GeofenceTypeCircle    GeofenceType = "`circle`"
	GeofenceTypePolygon   GeofenceType = "`polygon`"
	GeofenceTypeIsochrone GeofenceType = "`isochrone`"
)

// The property Type is required.
type GeofenceEntityCreateParam struct {
	// Specify the type of the geofence that is being created.
	//
	// Any of "`circle`", "`polygon`", "`isochrone`".
	Type GeofenceEntityCreateType `json:"type,omitzero,required"`
	// Set an unique ID for the new geofence. If not provided, an ID will be
	// automatically generated in UUID format. A valid `custom_id` can contain letters,
	// numbers, "-", & "\_" only.
	//
	// Please note that the ID of a geofence can not be changed once it is created.
	CustomID param.Opt[string] `json:"custom_id,omitzero"`
	// Name of the geofence. Use this field to assign a meaningful, custom name to the
	// geofence being created.
	Name param.Opt[string] `json:"name,omitzero"`
	// Provide the details to create a circular geofence. Please note that this object
	// is mandatory when `type` is `circle`. When the `type` is not `circle`, the
	// properties of this object will be ignored while creating the geofence.
	Circle GeofenceEntityCreateCircleParam `json:"circle,omitzero"`
	// Provide the details to create an isochrone based geofence. Use this object when
	// `type` is `isochrone`. When the `type` is not `isochrone`, the properties of
	// this object will be ignored while creating the geofence.
	Isochrone GeofenceEntityCreateIsochroneParam `json:"isochrone,omitzero"`
	// Metadata of the geofence. Use this field to define custom attributes that
	// provide more context and information about the geofence being created like
	// country, group ID etc.
	//
	// The data being added should be in valid JSON object format (i.e. `key` and
	// `value` pairs). Max size allowed for the object is 65kb.
	MetaData any `json:"meta_data,omitzero"`
	// Provide the details to create a custom polygon type of geofence. Please note
	// that this object is mandatory when `type` is `polygon`. When the `type` is not
	// `polygon`, the properties of this object will be ignored while creating the
	// geofence.
	//
	// Self-intersecting polygons or polygons containing other polygons are invalid and
	// will be removed while processing the request.
	//
	// Area of the polygon should be less than 2000 km<sup>2</sup>.
	Polygon GeofenceEntityCreatePolygonParam `json:"polygon,omitzero"`
	// An array of strings to associate multiple tags to the geofence. `tags` can be
	// used to search or filter geofences (using `Get Geofence List` method).
	//
	// Create valid `tags` using a string consisting of alphanumeric characters (A-Z,
	// a-z, 0-9) along with the underscore ('\_') and hyphen ('-') symbols.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r GeofenceEntityCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceEntityCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceEntityCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specify the type of the geofence that is being created.
type GeofenceEntityCreateType string

const (
	GeofenceEntityCreateTypeCircle    GeofenceEntityCreateType = "`circle`"
	GeofenceEntityCreateTypePolygon   GeofenceEntityCreateType = "`polygon`"
	GeofenceEntityCreateTypeIsochrone GeofenceEntityCreateType = "`isochrone`"
)

// Provide the details to create a circular geofence. Please note that this object
// is mandatory when `type` is `circle`. When the `type` is not `circle`, the
// properties of this object will be ignored while creating the geofence.
//
// The properties Center, Radius are required.
type GeofenceEntityCreateCircleParam struct {
	// Coordinate of the location which will act as the center of a circular geofence.
	Center GeofenceEntityCreateCircleCenterParam `json:"center,omitzero,required"`
	// Radius of the circular geofence, in meters. Maximum value allowed is 50000
	// meters.
	Radius float64 `json:"radius,required"`
	paramObj
}

func (r GeofenceEntityCreateCircleParam) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceEntityCreateCircleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceEntityCreateCircleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coordinate of the location which will act as the center of a circular geofence.
//
// The properties Lat, Lon are required.
type GeofenceEntityCreateCircleCenterParam struct {
	// Latitude of the `center` location.
	Lat float64 `json:"lat,required"`
	// Longitude of the `center` location.
	Lon float64 `json:"lon,required"`
	paramObj
}

func (r GeofenceEntityCreateCircleCenterParam) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceEntityCreateCircleCenterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceEntityCreateCircleCenterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provide the details to create an isochrone based geofence. Use this object when
// `type` is `isochrone`. When the `type` is not `isochrone`, the properties of
// this object will be ignored while creating the geofence.
//
// The property Coordinates is required.
type GeofenceEntityCreateIsochroneParam struct {
	// Coordinates of the location, in [latitude,longitude] format, which would act as
	// the starting point for identifying the isochrone polygon or the boundary of
	// reachable area. This parameter is mandatory when `type` is `isochrone`.
	Coordinates string `json:"coordinates,required"`
	// The distance, in meters, for which an isochrone polygon needs to be determined.
	// When provided, the API would create a geofence representing the area that can be
	// reached after driving the given number of meters starting from the point
	// specified in `coordinates`.
	//
	// The maximum distance that can be specified is 60000 meters (60km).
	//
	// At least one of `contours_meter` or `contours_minute` is mandatory when `type`
	// is `isochrone`.
	ContoursMeter param.Opt[int64] `json:"contours_meter,omitzero"`
	// The duration, in minutes, for which an isochrone polygon needs to be determined.
	// When provided, the API would create a geofence representing the area that can be
	// reached after driving for the given number of minutes starting from the point
	// specified in `coordinates`.
	//
	// The maximum duration that can be specified is 40 minutes.
	//
	// At least one of `contours_meter` or `contours_minute` is mandatory when `type`
	// is `isochrone`.
	ContoursMinute param.Opt[int64] `json:"contours_minute,omitzero"`
	// A floating point value from 0.0 to 1.0 that can be used to remove smaller
	// contours. A value of 1.0 will only return the largest contour for a given value.
	// A value of 0.5 drops any contours that are less than half the area of the
	// largest contour in the set of contours for that same value.
	Denoise param.Opt[float64] `json:"denoise,omitzero"`
	// A UNIX epoch timestamp in seconds format that can be used to set the departure
	// time. The isochrone boundary will be determined based on the typical traffic
	// conditions at the given time. If no input is provided for this parameter then
	// the traffic conditions at the time of making the request are considered
	DepartureTime param.Opt[int64] `json:"departure_time,omitzero"`
	// Set which driving mode the service should use to determine the isochrone line.
	//
	// For example, if you use `car`, the API will return an isochrone polygon that a
	// car can cover within the specified time or after driving the specified distance.
	// Using `truck` will return an isochrone that a truck can reach after taking into
	// account appropriate truck routing restrictions.
	//
	// Any of "`car`", "`truck`".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r GeofenceEntityCreateIsochroneParam) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceEntityCreateIsochroneParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceEntityCreateIsochroneParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GeofenceEntityCreateIsochroneParam](
		"mode", "`car`", "`truck`",
	)
}

// Provide the details to create a custom polygon type of geofence. Please note
// that this object is mandatory when `type` is `polygon`. When the `type` is not
// `polygon`, the properties of this object will be ignored while creating the
// geofence.
//
// Self-intersecting polygons or polygons containing other polygons are invalid and
// will be removed while processing the request.
//
// Area of the polygon should be less than 2000 km<sup>2</sup>.
//
// The property Geojson is required.
type GeofenceEntityCreatePolygonParam struct {
	// An object to collect geoJSON details of the geofence. The contents of this
	// object follow the
	// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Geojson GeofenceEntityCreatePolygonGeojsonParam `json:"geojson,omitzero,required"`
	paramObj
}

func (r GeofenceEntityCreatePolygonParam) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceEntityCreatePolygonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceEntityCreatePolygonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object to collect geoJSON details of the geofence. The contents of this
// object follow the
// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
//
// The properties Coordinates, Type are required.
type GeofenceEntityCreatePolygonGeojsonParam struct {
	// An array of coordinates in the [longitude, latitude] format, representing the
	// geofence boundary.
	Coordinates [][]float64 `json:"coordinates,omitzero,required"`
	// Type of the geoJSON geometry. Should always be `Polygon`.
	Type string `json:"type,required"`
	paramObj
}

func (r GeofenceEntityCreatePolygonGeojsonParam) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceEntityCreatePolygonGeojsonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceEntityCreatePolygonGeojsonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceNewResponse struct {
	// A data object containing the ID of the geofence created.
	Data GeofenceNewResponseData `json:"data"`
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
func (r GeofenceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GeofenceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A data object containing the ID of the geofence created.
type GeofenceNewResponseData struct {
	// Unique ID of the geofence created. It will be the same as `custom_id`, if
	// provided. Else it will be an auto generated UUID. Please note this ID cannot be
	// updated.
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *GeofenceNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceGetResponse struct {
	Data GeofenceGetResponseData `json:"data"`
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
func (r GeofenceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GeofenceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceGetResponseData struct {
	// An object with details of the geofence.
	Geofence Geofence `json:"geofence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geofence    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *GeofenceGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceListResponse struct {
	Data GeofenceListResponseData `json:"data"`
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
func (r GeofenceListResponse) RawJSON() string { return r.JSON.raw }
func (r *GeofenceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceListResponseData struct {
	List []Geofence `json:"list"`
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
func (r GeofenceListResponseData) RawJSON() string { return r.JSON.raw }
func (r *GeofenceListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceContainsResponse struct {
	Data GeofenceContainsResponseData `json:"data"`
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
func (r GeofenceContainsResponse) RawJSON() string { return r.JSON.raw }
func (r *GeofenceContainsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceContainsResponseData struct {
	// An array of objects containing each of the geofences provided in the `geofences`
	// input. If `geofences` in not provided then the array will return all the
	// geofences associated with the `key`
	ResultList []GeofenceContainsResponseDataResultList `json:"result_list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResultList  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceContainsResponseData) RawJSON() string { return r.JSON.raw }
func (r *GeofenceContainsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceContainsResponseDataResultList struct {
	// An object with details of the geofence.
	GeofenceDetail Geofence `json:"geofence_detail"`
	// ID of the geofence provided/generated at the time of creating the geofence.
	GeofenceID string `json:"geofence_id"`
	// An array of objects with results of the contains check for each of the
	// coordinate points in `locations` against the geofence represented by
	// `geofence_id`.
	Result []GeofenceContainsResponseDataResultListResult `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GeofenceDetail respjson.Field
		GeofenceID     respjson.Field
		Result         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceContainsResponseDataResultList) RawJSON() string { return r.JSON.raw }
func (r *GeofenceContainsResponseDataResultList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceContainsResponseDataResultListResult struct {
	// `true` when a coordinate point in `locations` is contained by this geofence.
	Contain bool `json:"contain"`
	// Index of the coordinate point in the input `locations`.
	LocationIndex int64 `json:"location_index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contain       respjson.Field
		LocationIndex respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeofenceContainsResponseDataResultListResult) RawJSON() string { return r.JSON.raw }
func (r *GeofenceContainsResponseDataResultListResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeofenceNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key                  string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	GeofenceEntityCreate GeofenceEntityCreateParam
	paramObj
}

func (r GeofenceNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.GeofenceEntityCreate)
}
func (r *GeofenceNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.GeofenceEntityCreate)
}

// URLQuery serializes [GeofenceNewParams]'s query parameters as `url.Values`.
func (r GeofenceNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeofenceGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [GeofenceGetParams]'s query parameters as `url.Values`.
func (r GeofenceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeofenceUpdateParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Use this parameter to update the `name` of a geofence. Users can assign
	// meaningful custom names to their geofences.
	Name param.Opt[string] `json:"name,omitzero"`
	// Use this object to update details of a circular geofence. Please note that this
	// object is mandatory only when `type` is `circle`. When the `type` is not
	// `circle`, the properties of this object will be ignored while creating the
	// geofence.
	Circle GeofenceUpdateParamsCircle `json:"circle,omitzero"`
	// Use this object to update details of an isochrone based geofence. Please note
	// that this object is mandatory only when `type` is `isochrone`. When the `type`
	// is not `isochrone`, the properties of this object will be ignored while creating
	// the geofence.
	Isochrone GeofenceUpdateParamsIsochrone `json:"isochrone,omitzero"`
	// Updated the `meta_data` associated with a geofence. Use this field to define
	// custom attributes that provide more context and information about the geofence
	// being updated like country, group ID etc.
	//
	// The data being added should be in valid JSON object format (i.e. `key` and
	// `value` pairs). Max size allowed for the object is 65kb.
	MetaData any `json:"meta_data,omitzero"`
	// Use this object to update details of a custom polygon geofence. Please note that
	// this object is mandatory only when `type` is `polygon`. When the `type` is not
	// `polygon`, the properties of this object will be ignored while creating the
	// geofence.
	//
	// Self-intersecting polygons or polygons containing other polygons are invalid and
	// will be removed while processing the request.
	//
	// Area of the polygon should be less than 2000 km<sup>2</sup>.
	Polygon GeofenceUpdateParamsPolygon `json:"polygon,omitzero"`
	// Use this parameter to add/modify one or multiple `tags` of a geofence. `tags`
	// can be used to search or filter geofences (using `Get Geofence List` method).
	//
	// Valid values for updating `tags` consist of alphanumeric characters (A-Z, a-z,
	// 0-9) along with the underscore ('\_') and hyphen ('-') symbols.
	Tags []string `json:"tags,omitzero"`
	// Use this parameter to update the `type` of a geofence. Please note that you will
	// need to provide required details for creating a geofence of the new `type`.
	// Check other parameters of this method to know more.
	//
	// Any of "`circle`", "`polygon`", "`isochrone`".
	Type GeofenceUpdateParamsType `json:"type,omitzero"`
	paramObj
}

func (r GeofenceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GeofenceUpdateParams]'s query parameters as `url.Values`.
func (r GeofenceUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Use this object to update details of a circular geofence. Please note that this
// object is mandatory only when `type` is `circle`. When the `type` is not
// `circle`, the properties of this object will be ignored while creating the
// geofence.
//
// The property Center is required.
type GeofenceUpdateParamsCircle struct {
	// Use this parameter to update the coordinate of the location which will act as
	// the center of a circular geofence.
	Center GeofenceUpdateParamsCircleCenter `json:"center,omitzero,required"`
	// Use this parameter to update the radius of the circular geofence, in meters.
	// Maximum value allowed is 50000 meters.
	Radius param.Opt[float64] `json:"radius,omitzero"`
	paramObj
}

func (r GeofenceUpdateParamsCircle) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceUpdateParamsCircle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceUpdateParamsCircle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use this parameter to update the coordinate of the location which will act as
// the center of a circular geofence.
type GeofenceUpdateParamsCircleCenter struct {
	// Latitude of the `center` location.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Longitude of the `center` location.
	Lon param.Opt[float64] `json:"lon,omitzero"`
	paramObj
}

func (r GeofenceUpdateParamsCircleCenter) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceUpdateParamsCircleCenter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceUpdateParamsCircleCenter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use this object to update details of an isochrone based geofence. Please note
// that this object is mandatory only when `type` is `isochrone`. When the `type`
// is not `isochrone`, the properties of this object will be ignored while creating
// the geofence.
type GeofenceUpdateParamsIsochrone struct {
	// Use this parameter to update the distance, in meters, for which an isochrone
	// polygon needs to be determined. When provided, the API would create a geofence
	// representing the area that can be reached after driving the given number of
	// meters starting from the point specified in `coordinates`.
	//
	// The maximum distance that can be specified is 60000 meters (60km).
	//
	// At least one of `contours_meter` or `contours_minute` is mandatory when `type`
	// is `isochrone`.
	ContoursMeter param.Opt[int64] `json:"contours_meter,omitzero"`
	// Use this parameter to update the duration, in minutes, for which an isochrone
	// polygon needs to be determined. When provided, the API would create a geofence
	// representing the area that can be reached after driving for the given number of
	// minutes starting from the point specified in `coordinates`.
	//
	// The maximum duration that can be specified is 40 minutes.
	//
	// At least one of `contours_meter` or `contours_minute` is mandatory when `type`
	// is `isochrone`.
	ContoursMinute param.Opt[int64] `json:"contours_minute,omitzero"`
	// Use this parameter to update the coordinates of the location, in
	// [latitude,longitude] format, which would act as the starting point for
	// identifying the isochrone polygon or the boundary of reachable area.
	Coordinates param.Opt[string] `json:"coordinates,omitzero"`
	// A floating point value from 0.0 to 1.0 that can be used to remove smaller
	// contours. A value of 1.0 will only return the largest contour for a given value.
	// A value of 0.5 drops any contours that are less than half the area of the
	// largest contour in the set of contours for that same value.
	//
	// Use this parameter to update the denoise value of the isochrone geofence.
	Denoise param.Opt[float64] `json:"denoise,omitzero"`
	// Use this parameter to update the `departure_time`, expressed as UNIX epoch
	// timestamp in seconds. The isochrone boundary will be determined based on the
	// typical traffic conditions at the given time.
	//
	// If no input is provided for this parameter then, the traffic conditions at the
	// time of making the request are considered by default. Please note that because
	// of this behavior the geofence boundaries may change even if the `departure_time`
	// was not specifically provided at the time of updating the geofence.
	DepartureTime param.Opt[int64] `json:"departure_time,omitzero"`
	// Use this parameter to update the driving mode that the service should use to
	// determine the isochrone line. For example, if you use `car`, the API will return
	// an isochrone polygon that a car can cover within the specified time or after
	// driving the specified distance. Using `truck` will return an isochrone that a
	// truck can reach after taking into account appropriate truck routing
	// restrictions.
	Mode param.Opt[string] `json:"mode,omitzero"`
	paramObj
}

func (r GeofenceUpdateParamsIsochrone) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceUpdateParamsIsochrone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceUpdateParamsIsochrone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use this object to update details of a custom polygon geofence. Please note that
// this object is mandatory only when `type` is `polygon`. When the `type` is not
// `polygon`, the properties of this object will be ignored while creating the
// geofence.
//
// Self-intersecting polygons or polygons containing other polygons are invalid and
// will be removed while processing the request.
//
// Area of the polygon should be less than 2000 km<sup>2</sup>.
type GeofenceUpdateParamsPolygon struct {
	// An object to collect geoJSON details of the `polygon` geofence. The contents of
	// this object follow the
	// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Geojson GeofenceUpdateParamsPolygonGeojson `json:"geojson,omitzero"`
	paramObj
}

func (r GeofenceUpdateParamsPolygon) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceUpdateParamsPolygon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceUpdateParamsPolygon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object to collect geoJSON details of the `polygon` geofence. The contents of
// this object follow the
// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
type GeofenceUpdateParamsPolygonGeojson struct {
	// Type of the geoJSON geometry. Should always be `Polygon`.
	Type param.Opt[string] `json:"type,omitzero"`
	// An array of coordinates in the [longitude, latitude] format, representing the
	// geofence boundary.
	Geometry [][]float64 `json:"geometry,omitzero"`
	paramObj
}

func (r GeofenceUpdateParamsPolygonGeojson) MarshalJSON() (data []byte, err error) {
	type shadow GeofenceUpdateParamsPolygonGeojson
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeofenceUpdateParamsPolygonGeojson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use this parameter to update the `type` of a geofence. Please note that you will
// need to provide required details for creating a geofence of the new `type`.
// Check other parameters of this method to know more.
type GeofenceUpdateParamsType string

const (
	GeofenceUpdateParamsTypeCircle    GeofenceUpdateParamsType = "`circle`"
	GeofenceUpdateParamsTypePolygon   GeofenceUpdateParamsType = "`polygon`"
	GeofenceUpdateParamsTypeIsochrone GeofenceUpdateParamsType = "`isochrone`"
)

type GeofenceListParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Denotes page number. Use this along with the `ps` parameter to implement
	// pagination for your searched results. This parameter does not have a maximum
	// limit but would return an empty response in case a higher value is provided when
	// the result-set itself is smaller.
	Pn param.Opt[int64] `query:"pn,omitzero" json:"-"`
	// Denotes number of search results per page. Use this along with the `pn`
	// parameter to implement pagination for your searched results.
	Ps param.Opt[int64] `query:"ps,omitzero" json:"-"`
	// Comma (`,`) separated list of `tags` which will be used to filter the geofences.
	//
	// Please note only the geofences which have all the `tags` added to this parameter
	// will be included in the result. This parameter can accept a string with a
	// maximum length of 256 characters.
	Tags param.Opt[string] `query:"tags,omitzero" format:"value_1,value_2,..." json:"-"`
	paramObj
}

// URLQuery serializes [GeofenceListParams]'s query parameters as `url.Values`.
func (r GeofenceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeofenceDeleteParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [GeofenceDeleteParams]'s query parameters as `url.Values`.
func (r GeofenceDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeofenceContainsParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Pipe (|) separated coordinates, in [latitude,longitude] format, of the locations
	// to be searched against the geofences.
	Locations string `query:"locations,required" json:"-"`
	// A `,` separated list geofence IDs against which the `locations` will be
	// searched. If not provided, then the 'locations' will be searched against all
	// your existing geofences.
	//
	// Maximum length of the string can be 256 characters.
	Geofences param.Opt[string] `query:"geofences,omitzero" format:"geofenceID_1,geofenceID_2,...." json:"-"`
	// When `true`, an array with detailed information of geofences is returned. When
	// `false`, an array containing only the IDs of the geofences is returned.
	Verbose param.Opt[string] `query:"verbose,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GeofenceContainsParams]'s query parameters as `url.Values`.
func (r GeofenceContainsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
