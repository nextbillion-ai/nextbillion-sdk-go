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

// PostalcodeService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostalcodeService] method instead.
type PostalcodeService struct {
	Options []option.RequestOption
}

// NewPostalcodeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPostalcodeService(opts ...option.RequestOption) (r PostalcodeService) {
	r = PostalcodeService{}
	r.Options = opts
	return
}

// Retrieve coordinates by postal code
func (r *PostalcodeService) GetCoordinates(ctx context.Context, params PostalcodeGetCoordinatesParams, opts ...option.RequestOption) (res *PostalcodeGetCoordinatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "postalcode"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PostalcodeGetCoordinatesResponse struct {
	// An object that contains details about the place that was provided in the input.
	Places PostalcodeGetCoordinatesResponsePlaces `json:"places"`
	// Returns a message, in case the input provided triggers any warnings.
	Warning []string `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Places      respjson.Field
		Warning     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostalcodeGetCoordinatesResponse) RawJSON() string { return r.JSON.raw }
func (r *PostalcodeGetCoordinatesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object that contains details about the place that was provided in the input.
type PostalcodeGetCoordinatesResponsePlaces struct {
	// Returns the address of the `postalcode` returned.
	Address string `json:"address"`
	// An object containing the boundary details of the postal code area. This object
	// will not be returned in case the boundary information of the postal code
	// provided is not available (only for selected countries).
	//
	// Please note the contents of this object will change based on the `format` field
	// in the input. When the `format` field is not present in the input this object
	// would contain `multipolygon` - `polygon` - `points` objects depending on the
	// boundary of the given postal code. When the `format` field is present in the
	// input, then the contents of this object would match the
	// [geojson format and standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Boundary PostalcodeGetCoordinatesResponsePlacesBoundary `json:"boundary"`
	// Name of the country containing the geographic coordinate point / postal code
	// provided in the input request.
	Country string `json:"country"`
	// Returns the [alpha-3 ISO code](https://www.iban.com/country-codes) of the
	// country containing the `postalcode` returned.
	CountryCode string `json:"countryCode"`
	// This property is returned only when the API is requested to fetch the postal
	// code containing the location coordinate provided in the `at` input parameter.
	// `distance` denotes the straight line distance, in meters, from the requested
	// location coordinate to the postal code centroid.
	Distance float64 `json:"distance"`
	// Name of the district or region containing the geographic coordinate point /
	// postal code provided in the input request.
	District string `json:"district"`
	// Refers to the geographic coordinate denoting the center of the postal code in
	// latitude, longitude format.
	Geopoint PostalcodeGetCoordinatesResponsePlacesGeopoint `json:"geopoint"`
	// Returns the postal code associated with the requested geographic coordinate
	// point or the postal code itself as provided in the input API request.
	Postalcode string `json:"postalcode"`
	// Name of the state or province containing the geographic coordinate point /
	// postal code provided in the input request.
	State string `json:"state"`
	// Name of the sub-district or sub-region containing the postal code or geographic
	// coordinate point / postal code provided in the input request
	Subdistrict string `json:"subdistrict"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Boundary    respjson.Field
		Country     respjson.Field
		CountryCode respjson.Field
		Distance    respjson.Field
		District    respjson.Field
		Geopoint    respjson.Field
		Postalcode  respjson.Field
		State       respjson.Field
		Subdistrict respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostalcodeGetCoordinatesResponsePlaces) RawJSON() string { return r.JSON.raw }
func (r *PostalcodeGetCoordinatesResponsePlaces) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the boundary details of the postal code area. This object
// will not be returned in case the boundary information of the postal code
// provided is not available (only for selected countries).
//
// Please note the contents of this object will change based on the `format` field
// in the input. When the `format` field is not present in the input this object
// would contain `multipolygon` - `polygon` - `points` objects depending on the
// boundary of the given postal code. When the `format` field is present in the
// input, then the contents of this object would match the
// [geojson format and standard](https://datatracker.ietf.org/doc/html/rfc7946).
type PostalcodeGetCoordinatesResponsePlacesBoundary struct {
	// An object with geoJSON details of the boundary. This object is returned when the
	// `format` field is set to `geojson` in the input request, otherwise it is not
	// present in the response. The contents of this object follow the
	// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Geometry PostalcodeGetCoordinatesResponsePlacesBoundaryGeometry `json:"geometry"`
	// An array of objects containing information about all the polygons forming the
	// postal code area. In case, the postal code area is formed by multiple polygons
	// not containing each other, a matching count of `polygon` objects will be
	// returned.
	//
	// Please note that this object is returned only when `format` field is not
	// specified in the input, otherwise it is not present in the response.
	Multipolygon []PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygon `json:"multipolygon"`
	// Property associated with the geoJSON shape.
	Properties string `json:"properties"`
	// Type of the geoJSON object. This parameter is returned when the `format` field
	// is set to `geojson` in the input request, otherwise it is not present in the
	// response. The contents of this object follow the
	// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geometry     respjson.Field
		Multipolygon respjson.Field
		Properties   respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostalcodeGetCoordinatesResponsePlacesBoundary) RawJSON() string { return r.JSON.raw }
func (r *PostalcodeGetCoordinatesResponsePlacesBoundary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with geoJSON details of the boundary. This object is returned when the
// `format` field is set to `geojson` in the input request, otherwise it is not
// present in the response. The contents of this object follow the
// [geoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
type PostalcodeGetCoordinatesResponsePlacesBoundaryGeometry struct {
	// An array of coordinates in the [longitude, latitude] format, representing the
	// coordinates points which lie on the boundary of the postal code area.
	Coordinates [][][]float64 `json:"coordinates"`
	// Type of the geoJSON geometry.
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
func (r PostalcodeGetCoordinatesResponsePlacesBoundaryGeometry) RawJSON() string { return r.JSON.raw }
func (r *PostalcodeGetCoordinatesResponsePlacesBoundaryGeometry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygon struct {
	// An object containing the details of a single polygon that is a part of the
	// postal code area. In case the postal code area contains other polygon(s), the
	// details of such polygon(s) would be returned through an array of `points`
	// object.
	Polygon []PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygonPolygon `json:"polygon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Polygon     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygon) RawJSON() string {
	return r.JSON.raw
}
func (r *PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygonPolygon struct {
	// Represents an array of geographic coordinates that define a `polygon` boundary.
	Points []PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygonPolygonPoint `json:"points"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Points      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygonPolygon) RawJSON() string {
	return r.JSON.raw
}
func (r *PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygonPolygon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygonPolygonPoint struct {
	// Latitude of the coordinate.
	Lat float64 `json:"lat"`
	// Longitude of the coordinate.
	Lng float64 `json:"lng"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lng         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygonPolygonPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *PostalcodeGetCoordinatesResponsePlacesBoundaryMultipolygonPolygonPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refers to the geographic coordinate denoting the center of the postal code in
// latitude, longitude format.
type PostalcodeGetCoordinatesResponsePlacesGeopoint struct {
	// Latitude of the location.
	Lat float64 `json:"lat"`
	// Longitude of the location.
	Lng float64 `json:"lng"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lng         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostalcodeGetCoordinatesResponsePlacesGeopoint) RawJSON() string { return r.JSON.raw }
func (r *PostalcodeGetCoordinatesResponsePlacesGeopoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostalcodeGetCoordinatesParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Country containing the postal code or the location. It is mandatory if
	// `postalcode` is provided in the request. [See this example](#note).
	//
	// Please check the [API Query Limits](#api-query-limits) section below for a list
	// of the countries covered by the Geocode Postcode API. Users can provide either
	// the name or the alpha-2/3 code as per the
	// [ISO 3166-1 standard](https://en.wikipedia.org/wiki/ISO_3166-1) of a country
	// covered by the API as input for this parameter.
	Country param.Opt[string] `json:"country,omitzero"`
	// Provide the postal code for which the information is needed. At least one of
	// (`postalcode` + `country`) or `at` needs to be provided. Please note that only 1
	// postal code can be requested. [See this example](#note).
	Postalcode param.Opt[string] `json:"postalcode,omitzero"`
	// Location coordinates that you want to get the postal code of. If not providing
	// `postalcode` in the request, `at` becomes mandatory. Please note that only 1
	// point can be requested. [See this example](#note).
	At PostalcodeGetCoordinatesParamsAt `json:"at,omitzero"`
	// Specify the format in which the boundary details of the post code will be
	// returned. When specified, the boundary details will be returned in the `geojson`
	// format. When not specified, the boundary details are returned in general format.
	//
	// Any of "`geojson`".
	Format PostalcodeGetCoordinatesParamsFormat `json:"format,omitzero"`
	paramObj
}

func (r PostalcodeGetCoordinatesParams) MarshalJSON() (data []byte, err error) {
	type shadow PostalcodeGetCoordinatesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostalcodeGetCoordinatesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PostalcodeGetCoordinatesParams]'s query parameters as
// `url.Values`.
func (r PostalcodeGetCoordinatesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Location coordinates that you want to get the postal code of. If not providing
// `postalcode` in the request, `at` becomes mandatory. Please note that only 1
// point can be requested. [See this example](#note).
type PostalcodeGetCoordinatesParamsAt struct {
	// Latitude of the location.
	Lat param.Opt[float64] `json:"lat,omitzero"`
	// Longitude of the location.
	Lng param.Opt[float64] `json:"lng,omitzero"`
	paramObj
}

func (r PostalcodeGetCoordinatesParamsAt) MarshalJSON() (data []byte, err error) {
	type shadow PostalcodeGetCoordinatesParamsAt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostalcodeGetCoordinatesParamsAt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specify the format in which the boundary details of the post code will be
// returned. When specified, the boundary details will be returned in the `geojson`
// format. When not specified, the boundary details are returned in general format.
type PostalcodeGetCoordinatesParamsFormat string

const (
	PostalcodeGetCoordinatesParamsFormatGeojson PostalcodeGetCoordinatesParamsFormat = "`geojson`"
)
