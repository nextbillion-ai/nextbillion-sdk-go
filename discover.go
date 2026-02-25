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

// DiscoverService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDiscoverService] method instead.
type DiscoverService struct {
	Options []option.RequestOption
}

// NewDiscoverService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDiscoverService(opts ...option.RequestOption) (r DiscoverService) {
	r = DiscoverService{}
	r.Options = opts
	return
}

// Discover matching places
func (r *DiscoverService) Get(ctx context.Context, query DiscoverGetParams, opts ...option.RequestOption) (res *DiscoverGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "discover"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DiscoverGetResponse struct {
	// The results are presented as a JSON list of candidates in ranked order
	// (most-likely to least-likely) based on the matched location criteria.
	Items []DiscoverGetResponseItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiscoverGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DiscoverGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DiscoverGetResponseItem struct {
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
	// result covers. place typed results have no mapView.
	MapView MapView `json:"mapView"`
	// Returns the operating hours of the place, if available.
	OpeningHours DiscoverGetResponseItemOpeningHours `json:"openingHours"`
	// Returns the location coordinates of the result.
	Position Position `json:"position"`
	// Score of the result. A higher score indicates a closer match.
	Scoring DiscoverGetResponseItemScoring `json:"scoring"`
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
func (r DiscoverGetResponseItem) RawJSON() string { return r.JSON.raw }
func (r *DiscoverGetResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the operating hours of the place, if available.
type DiscoverGetResponseItemOpeningHours struct {
	// A collection of attributes with details about the opening and closing hours for
	// each day of the week.
	TimeRanges []DiscoverGetResponseItemOpeningHoursTimeRange `json:"timeRanges"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TimeRanges  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiscoverGetResponseItemOpeningHours) RawJSON() string { return r.JSON.raw }
func (r *DiscoverGetResponseItemOpeningHours) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DiscoverGetResponseItemOpeningHoursTimeRange struct {
	// Returns the closing time details.
	EndTime DiscoverGetResponseItemOpeningHoursTimeRangeEndTime `json:"endTime"`
	// Returns the open time details.
	StartTime DiscoverGetResponseItemOpeningHoursTimeRangeStartTime `json:"startTime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiscoverGetResponseItemOpeningHoursTimeRange) RawJSON() string { return r.JSON.raw }
func (r *DiscoverGetResponseItemOpeningHoursTimeRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the closing time details.
type DiscoverGetResponseItemOpeningHoursTimeRangeEndTime struct {
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
func (r DiscoverGetResponseItemOpeningHoursTimeRangeEndTime) RawJSON() string { return r.JSON.raw }
func (r *DiscoverGetResponseItemOpeningHoursTimeRangeEndTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the open time details.
type DiscoverGetResponseItemOpeningHoursTimeRangeStartTime struct {
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
func (r DiscoverGetResponseItemOpeningHoursTimeRangeStartTime) RawJSON() string { return r.JSON.raw }
func (r *DiscoverGetResponseItemOpeningHoursTimeRangeStartTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Score of the result. A higher score indicates a closer match.
type DiscoverGetResponseItemScoring struct {
	// A breakdown of how closely individual field of the result matched with the
	// provided query q.
	FieldScore any `json:"fieldScore"`
	// A score, out of 1, indicating how closely the result matches with the provided
	// query q .
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
func (r DiscoverGetResponseItemScoring) RawJSON() string { return r.JSON.raw }
func (r *DiscoverGetResponseItemScoring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DiscoverGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// Specify the free-text search query.
	//
	// Please note that whitespace, urls, email addresses, or other out-of-scope
	// queries will yield no results.
	Q string `query:"q" api:"required" json:"-"`
	// Specify the center of the search context expressed as coordinates.
	//
	// Please note that one of "at", "in=circle" or "in=bbox" should be provided for
	// relevant results.
	At param.Opt[string] `query:"at,omitzero" format:"latitude,longitude" json:"-"`
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
	//     Format: countryCode:{countryCode}[,{countryCode}]
	//
	//   - a circular area, provided as latitude, longitude, and radius (an integer with
	//     meters as unit)
	//
	//     Format: circle:{latitude},{longitude};r={radius}
	//
	//   - a bounding box, provided as _west longitude_, _south latitude_, _east
	//     longitude_, _north latitude_
	//
	//     Format: bbox:{west longitude},{south latitude},{east longitude},{north
	//     latitude}
	//
	// Please provide one of 'at', 'in=circle' or 'in=bbox' input for a relevant
	// result.
	In param.Opt[string] `query:"in,omitzero" json:"-"`
	// Select the language to be used for result rendering from a list of
	// [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language
	// codes.
	Lang param.Opt[string] `query:"lang,omitzero" json:"-"`
	// Sets the maximum number of results to be returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DiscoverGetParams]'s query parameters as `url.Values`.
func (r DiscoverGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
