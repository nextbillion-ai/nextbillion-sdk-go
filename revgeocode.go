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

// RevgeocodeService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRevgeocodeService] method instead.
type RevgeocodeService struct {
	Options []option.RequestOption
}

// NewRevgeocodeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRevgeocodeService(opts ...option.RequestOption) (r RevgeocodeService) {
	r = RevgeocodeService{}
	r.Options = opts
	return
}

// Reverse Geocode
func (r *RevgeocodeService) Get(ctx context.Context, query RevgeocodeGetParams, opts ...option.RequestOption) (res *RevgeocodeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "revgeocode"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RevgeocodeGetResponse struct {
	// The results are presented as a JSON list of candidates in ranked order
	// (most-likely to least-likely) based on the matched location criteria.
	Items []RevgeocodeGetResponseItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RevgeocodeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RevgeocodeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RevgeocodeGetResponseItem struct {
	// The unique identifier for the result item.
	ID string `json:"id"`
	// An array returning the location coordinates of all the access points of the
	// search result.
	Access Access `json:"access"`
	// Postal address of the result item.
	Address Address `json:"address"`
	// The list of categories assigned to this place.
	Categories []Categories `json:"categories"`
	// Contact information like phone, email or website.
	Contacts []Contacts `json:"contacts"`
	// The distance "as the crow flies" from the search center to this result item in
	// meters.
	Distance int64 `json:"distance"`
	// The bounding box enclosing the geometric shape (area or line) that an individual
	// result covers. `place` typed results have no `mapView`.
	MapView MapView `json:"mapView"`
	// Returns the operating hours of the place, if available.
	OpeningHours RevgeocodeGetResponseItemOpeningHours `json:"openingHours"`
	// Returns the location coordinates of the result.
	Position Position `json:"position"`
	// Score of the result. A higher score indicates a closer match.
	Scoring RevgeocodeGetResponseItemScoring `json:"scoring"`
	// The localized display name of this result item.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Access       respjson.Field
		Address      respjson.Field
		Categories   respjson.Field
		Contacts     respjson.Field
		Distance     respjson.Field
		MapView      respjson.Field
		OpeningHours respjson.Field
		Position     respjson.Field
		Scoring      respjson.Field
		Title        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RevgeocodeGetResponseItem) RawJSON() string { return r.JSON.raw }
func (r *RevgeocodeGetResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the operating hours of the place, if available.
type RevgeocodeGetResponseItemOpeningHours struct {
	// A collection of attributes with details about the opening and closing hours for
	// each day of the week.
	TimeRanges []RevgeocodeGetResponseItemOpeningHoursTimeRange `json:"timeRanges"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TimeRanges  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RevgeocodeGetResponseItemOpeningHours) RawJSON() string { return r.JSON.raw }
func (r *RevgeocodeGetResponseItemOpeningHours) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RevgeocodeGetResponseItemOpeningHoursTimeRange struct {
	// Returns the closing time details.
	EndTime RevgeocodeGetResponseItemOpeningHoursTimeRangeEndTime `json:"endTime"`
	// Returns the open time details.
	StartTime RevgeocodeGetResponseItemOpeningHoursTimeRangeStartTime `json:"startTime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RevgeocodeGetResponseItemOpeningHoursTimeRange) RawJSON() string { return r.JSON.raw }
func (r *RevgeocodeGetResponseItemOpeningHoursTimeRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the closing time details.
type RevgeocodeGetResponseItemOpeningHoursTimeRangeEndTime struct {
	// The date to which the subsequent closing time details belong to.
	Date string `json:"date"`
	// The hour of the day when the place closes.
	Hour int64 `json:"hour"`
	// The minute of the hour when the place closes.
	Minute int64 `json:"minute"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Hour        respjson.Field
		Minute      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RevgeocodeGetResponseItemOpeningHoursTimeRangeEndTime) RawJSON() string { return r.JSON.raw }
func (r *RevgeocodeGetResponseItemOpeningHoursTimeRangeEndTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the open time details.
type RevgeocodeGetResponseItemOpeningHoursTimeRangeStartTime struct {
	// The date to which the subsequent open time details belong to.
	Date string `json:"date"`
	// The hour of the day when the place opens.
	Hour int64 `json:"hour"`
	// The minute of the hour when the place opens.
	Minute int64 `json:"minute"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Hour        respjson.Field
		Minute      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RevgeocodeGetResponseItemOpeningHoursTimeRangeStartTime) RawJSON() string { return r.JSON.raw }
func (r *RevgeocodeGetResponseItemOpeningHoursTimeRangeStartTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Score of the result. A higher score indicates a closer match.
type RevgeocodeGetResponseItemScoring struct {
	// A breakdown of how closely individual field of the result matched with the
	// provided query `q`.
	FieldScore any `json:"fieldScore"`
	// A score, out of 1, indicating how closely the result matches with the provided
	// query `q` .
	QueryScore float64 `json:"queryScore"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldScore  respjson.Field
		QueryScore  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RevgeocodeGetResponseItemScoring) RawJSON() string { return r.JSON.raw }
func (r *RevgeocodeGetResponseItemScoring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RevgeocodeGetParams struct {
	// Specify the center of the search context expressed as coordinates.
	//
	// Please note that one of "at", "in=circle" or "in=bbox" should be provided for
	// relevant results.
	At string `query:"at,required" format:"latitude,longitude" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Search within a geographic area. This is a hard filter. Results will be returned
	// if they are located within the specified area.
	//
	// # A geographic area can be
	//
	//   - a country (or multiple countries), provided as comma-separated
	//     [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) country
	//     codes
	//
	//     The country codes are to be provided in all uppercase.
	//
	//     Format: `countryCode:{countryCode}[,{countryCode}]`
	//
	//   - a circular area, provided as latitude, longitude, and radius (an integer with
	//     meters as unit)
	//
	//     Format: `circle:{latitude},{longitude};r={radius}`
	//
	//   - a bounding box, provided as _west longitude_, _south latitude_, _east
	//     longitude_, _north latitude_
	//
	//     Format:
	//     `bbox:{west longitude},{south latitude},{east longitude},{north latitude}`
	In param.Opt[string] `query:"in,omitzero" json:"-"`
	// Select the language to be used for result rendering from a list of
	// [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language
	// codes.
	Lang param.Opt[string] `query:"lang,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RevgeocodeGetParams]'s query parameters as `url.Values`.
func (r RevgeocodeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
