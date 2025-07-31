// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

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

// FleetifyDocumentTemplateService contains methods and other services that help
// with interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFleetifyDocumentTemplateService] method instead.
type FleetifyDocumentTemplateService struct {
	Options []option.RequestOption
}

// NewFleetifyDocumentTemplateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFleetifyDocumentTemplateService(opts ...option.RequestOption) (r FleetifyDocumentTemplateService) {
	r = FleetifyDocumentTemplateService{}
	r.Options = opts
	return
}

// Create Document template
func (r *FleetifyDocumentTemplateService) New(ctx context.Context, params FleetifyDocumentTemplateNewParams, opts ...option.RequestOption) (res *FleetifyDocumentTemplateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "fleetify/document_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve template by ID
func (r *FleetifyDocumentTemplateService) Get(ctx context.Context, id string, query FleetifyDocumentTemplateGetParams, opts ...option.RequestOption) (res *FleetifyDocumentTemplateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fleetify/document_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a document template
func (r *FleetifyDocumentTemplateService) Update(ctx context.Context, id string, params FleetifyDocumentTemplateUpdateParams, opts ...option.RequestOption) (res *FleetifyDocumentTemplateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fleetify/document_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Get all document templates
func (r *FleetifyDocumentTemplateService) List(ctx context.Context, query FleetifyDocumentTemplateListParams, opts ...option.RequestOption) (res *FleetifyDocumentTemplateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "fleetify/document_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a document template
func (r *FleetifyDocumentTemplateService) Delete(ctx context.Context, id string, body FleetifyDocumentTemplateDeleteParams, opts ...option.RequestOption) (res *FleetifyDocumentTemplateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fleetify/document_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// An object to collect the details of form fields - data structures, validation
// rules - for collecting required information after successfully executing a route
// step.
//
// The properties Label, Type are required.
type DocumentTemplateContentRequestParam struct {
	// Specify the label or the name of the field. The `label` specified here can be
	// used as field name when rendering the document in the Driver app.
	Label string `json:"label,required"`
	// Specify the data type of the field. It corresponds to the type of information
	// that the driver needs to collect.
	//
	// Any of "`string`", "`number`", "`date_time`", "`photos`", "`multi_choices`",
	// "`signature`", "`barcode`", "`single_choice`".
	Type DocumentTemplateContentRequestType `json:"type,omitzero,required"`
	// Specify the name of the document field. A field's`name` can be used for internal
	// references to the document field.
	Name param.Opt[string] `json:"name,omitzero"`
	// Specify if it is mandatory to fill the field. Default value is false.
	Required param.Opt[bool] `json:"required,omitzero"`
	// An object to define additional information required for `single_choice` or
	// `multi_choices` type document items.
	Meta DocumentTemplateContentRequestMetaParam `json:"meta,omitzero"`
	// Specify the validation rules for the field. This can be used to enforce data
	// quality and integrity checks. For example, if the field is a number type,
	// `validation` can define constraints like minimum / maximum number values.
	Validation DocumentTemplateContentRequestValidationParam `json:"validation,omitzero"`
	paramObj
}

func (r DocumentTemplateContentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTemplateContentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTemplateContentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specify the data type of the field. It corresponds to the type of information
// that the driver needs to collect.
type DocumentTemplateContentRequestType string

const (
	DocumentTemplateContentRequestTypeString       DocumentTemplateContentRequestType = "`string`"
	DocumentTemplateContentRequestTypeNumber       DocumentTemplateContentRequestType = "`number`"
	DocumentTemplateContentRequestTypeDateTime     DocumentTemplateContentRequestType = "`date_time`"
	DocumentTemplateContentRequestTypePhotos       DocumentTemplateContentRequestType = "`photos`"
	DocumentTemplateContentRequestTypeMultiChoices DocumentTemplateContentRequestType = "`multi_choices`"
	DocumentTemplateContentRequestTypeSignature    DocumentTemplateContentRequestType = "`signature`"
	DocumentTemplateContentRequestTypeBarcode      DocumentTemplateContentRequestType = "`barcode`"
	DocumentTemplateContentRequestTypeSingleChoice DocumentTemplateContentRequestType = "`single_choice`"
)

// An object to define additional information required for `single_choice` or
// `multi_choices` type document items.
//
// The property Options is required.
type DocumentTemplateContentRequestMetaParam struct {
	// An array of objects to define options for a `multi_choices` or `single_choice`
	// type document field. Each object represents one option.
	Options []DocumentTemplateContentRequestMetaOptionParam `json:"options,omitzero,required"`
	paramObj
}

func (r DocumentTemplateContentRequestMetaParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTemplateContentRequestMetaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTemplateContentRequestMetaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Label, Value are required.
type DocumentTemplateContentRequestMetaOptionParam struct {
	// Specify the label or name for the option.
	Label string `json:"label,required"`
	// Specify the value associated with the option. This value will be submitted when
	// the option is checked in the Driver app.
	Value string `json:"value,required"`
	paramObj
}

func (r DocumentTemplateContentRequestMetaOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTemplateContentRequestMetaOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTemplateContentRequestMetaOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specify the validation rules for the field. This can be used to enforce data
// quality and integrity checks. For example, if the field is a number type,
// `validation` can define constraints like minimum / maximum number values.
type DocumentTemplateContentRequestValidationParam struct {
	// Specifies the maximum allowed value for `number` type document field. Input
	// values must be less than or equal to this threshold.
	Max param.Opt[int64] `json:"max,omitzero"`
	// Specifies the maximum number of items for `multi_choices`, `photos` type
	// document fields. The number of provided input items must be less than or equal
	// to this threshold.
	MaxItems param.Opt[int64] `json:"max_items,omitzero"`
	// Specifies the minimum allowed value for `number` type document field. Input
	// values must be greater than or equal to this threshold.
	Min param.Opt[int64] `json:"min,omitzero"`
	// Specifies the minimum number of items for `multi_choices`, `photos` type
	// document fields. The number of provided input items must be greater than or
	// equal to this threshold.
	MinItems param.Opt[int64] `json:"min_items,omitzero"`
	paramObj
}

func (r DocumentTemplateContentRequestValidationParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTemplateContentRequestValidationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTemplateContentRequestValidationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An array of objects returning the details of data structures and validation
// rules and other properties of all document fields. Each object represents one
// document field.
type DocumentTemplateContentResponse struct {
	// Returns the label of the document field.
	Label string `json:"label"`
	// Returns the options configured for `single_choice` or `multi_choices` type
	// document items.
	Meta DocumentTemplateContentResponseMeta `json:"meta"`
	// Returns the name of the document field.
	Name string `json:"name"`
	// Indicates if the document field is mandatory or not.
	Required bool `json:"required"`
	// Returns the data type of the document field. It will always belong to one of
	// `string`, `number`, `date_time`, `photos`, `multi_choices`, `signature`,
	// `barcode`, and `single_choice.`
	Type string `json:"type"`
	// Returns the validation rules for `number` , `multi_choices` , and `photos`
	// document field types.
	Validation DocumentTemplateContentResponseValidation `json:"validation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Meta        respjson.Field
		Name        respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Validation  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentTemplateContentResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentTemplateContentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the options configured for `single_choice` or `multi_choices` type
// document items.
type DocumentTemplateContentResponseMeta struct {
	// An array of objects returning the options for `multi_choices` or `single_choice`
	// type document field. Each object represents one configured option.
	Options []DocumentTemplateContentResponseMetaOption `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentTemplateContentResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *DocumentTemplateContentResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTemplateContentResponseMetaOption struct {
	// Returns the label for the option.
	Label string `json:"label"`
	// Returns the value associated with the option. This value gets submitted when the
	// option is checked in the Driver app.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentTemplateContentResponseMetaOption) RawJSON() string { return r.JSON.raw }
func (r *DocumentTemplateContentResponseMetaOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the validation rules for `number` , `multi_choices` , and `photos`
// document field types.
type DocumentTemplateContentResponseValidation struct {
	// Returns the maximum allowed value for `number` type document item, as specified
	// at the time of configuring the field. This parameter is not present in the
	// response if it was not provided in the input.
	Max int64 `json:"max"`
	// Returns the maximum number of items required for `multi_choices`, `photos` type
	// document items. This parameter will not be present in the response if it was not
	// provided in the input.
	MaxItems string `json:"max_items"`
	// Returns the minimum allowed value for `number` type document item, as specified
	// at the time of configuring the field. This parameter is not present in the
	// response if it was not provided in the input.
	Min int64 `json:"min"`
	// Returns the minimum number of items required for `multi_choices`, `photos` type
	// document items. This parameter will not be present in the response if it was not
	// provided in the input.
	MinItems string `json:"min_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		MaxItems    respjson.Field
		Min         respjson.Field
		MinItems    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentTemplateContentResponseValidation) RawJSON() string { return r.JSON.raw }
func (r *DocumentTemplateContentResponseValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyDocumentTemplateNewResponse struct {
	// An object returning the details of the document template created.
	Data FleetifyDocumentTemplateNewResponseData `json:"data"`
	// Returns the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Msg string `json:"msg"`
	// Returns the HTTP response code.
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyDocumentTemplateNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyDocumentTemplateNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning the details of the document template created.
type FleetifyDocumentTemplateNewResponseData struct {
	// Returns the unique ID of the document template created.
	ID string `json:"id"`
	// An array of objects returning the details of data structures and validation
	// rules and other properties of all document fields. Each object represents one
	// document field.
	Content []DocumentTemplateContentResponse `json:"content"`
	// Returns the name of the document template as specified in the input.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyDocumentTemplateNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *FleetifyDocumentTemplateNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyDocumentTemplateGetResponse struct {
	// An object returning the details of the requested document template.
	Data FleetifyDocumentTemplateGetResponseData `json:"data"`
	// Returns the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Msg string `json:"msg"`
	// Returns the HTTP response code.
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyDocumentTemplateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyDocumentTemplateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning the details of the requested document template.
type FleetifyDocumentTemplateGetResponseData struct {
	// Returns the unique identifier of the document template.
	ID string `json:"id"`
	// An array of objects returning the details of data structures and validation
	// rules and other properties of all document fields. Each object represents one
	// document field.
	Content []DocumentTemplateContentResponse `json:"content"`
	// Returns the name of the document template as specified at the time of creating
	// the template.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyDocumentTemplateGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *FleetifyDocumentTemplateGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyDocumentTemplateUpdateResponse struct {
	// An object returning the details of the updated document template.
	Data FleetifyDocumentTemplateUpdateResponseData `json:"data"`
	// Returns the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Msg string `json:"msg"`
	// Returns the HTTP response code.
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyDocumentTemplateUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyDocumentTemplateUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object returning the details of the updated document template.
type FleetifyDocumentTemplateUpdateResponseData struct {
	// Returns the unique ID of the document template.
	ID string `json:"id"`
	// An array of object returning the details of updated data structures and
	// validation rules for document fields. Each object represents one document field.
	Content []DocumentTemplateContentResponse `json:"content"`
	// Returns the updated name of the document template.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyDocumentTemplateUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *FleetifyDocumentTemplateUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyDocumentTemplateListResponse struct {
	// An array of objects returning the details of each document template associated
	// with the specified API key. Each object represents one document template. In
	// case there are no templates associated with the given key, a blank array is
	// returned.
	Data []FleetifyDocumentTemplateListResponseData `json:"data"`
	// Returns the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Msg string `json:"msg"`
	// Returns the HTTP response code.
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyDocumentTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyDocumentTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An array of objects returning the details of each document template associated
// with the specified API key. Each object represents one document template. In
// case there are no templates associated with the given key, a blank array is
// returned.
type FleetifyDocumentTemplateListResponseData struct {
	// Returns the unique ID of the document template.
	ID string `json:"id"`
	// An array of objects returning the details of data structures and validation
	// rules and other properties of all document fields. Each object represents one
	// document field.
	Content []DocumentTemplateContentResponse `json:"content"`
	// Returns the name of the document template.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyDocumentTemplateListResponseData) RawJSON() string { return r.JSON.raw }
func (r *FleetifyDocumentTemplateListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyDocumentTemplateDeleteResponse struct {
	// Returns the error message in case of a failed request. If the request is
	// successful, this field is not present in the response.
	Msg string `json:"msg"`
	// Returns the HTTP response code.
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FleetifyDocumentTemplateDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FleetifyDocumentTemplateDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FleetifyDocumentTemplateNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	// A form field that drivers must complete when executing a route step. Defines the
	// data structure and validation rules for collecting required information during
	// route execution.
	Content []DocumentTemplateContentRequestParam `json:"content,omitzero,required"`
	// Specify a name for the document template to be created.
	Name string `json:"name,required"`
	paramObj
}

func (r FleetifyDocumentTemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyDocumentTemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyDocumentTemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FleetifyDocumentTemplateNewParams]'s query parameters as
// `url.Values`.
func (r FleetifyDocumentTemplateNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FleetifyDocumentTemplateGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	paramObj
}

// URLQuery serializes [FleetifyDocumentTemplateGetParams]'s query parameters as
// `url.Values`.
func (r FleetifyDocumentTemplateGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FleetifyDocumentTemplateUpdateParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	// Specify the document template name to be updated.
	Name param.Opt[string] `json:"name,omitzero"`
	// An object to collect the details of form fields to be updated - data structures,
	// validation rules. Please note that the details provided here will overwrite any
	// existing document fields in the given template.
	Content []DocumentTemplateContentRequestParam `json:"content,omitzero"`
	paramObj
}

func (r FleetifyDocumentTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FleetifyDocumentTemplateUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FleetifyDocumentTemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FleetifyDocumentTemplateUpdateParams]'s query parameters as
// `url.Values`.
func (r FleetifyDocumentTemplateUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FleetifyDocumentTemplateListParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	paramObj
}

// URLQuery serializes [FleetifyDocumentTemplateListParams]'s query parameters as
// `url.Values`.
func (r FleetifyDocumentTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FleetifyDocumentTemplateDeleteParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" json:"-"`
	paramObj
}

// URLQuery serializes [FleetifyDocumentTemplateDeleteParams]'s query parameters as
// `url.Values`.
func (r FleetifyDocumentTemplateDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
