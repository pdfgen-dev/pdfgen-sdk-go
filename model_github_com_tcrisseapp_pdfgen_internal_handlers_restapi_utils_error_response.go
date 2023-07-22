/*
PDFGen API

The PDFGen API for HTML to PDF generation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse{}

// GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse struct for GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse
type GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse struct {
	ErrorCode *int32 `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse instantiates a new GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse() *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse {
	this := GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse{}
	return &this
}

// NewGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponseWithDefaults instantiates a new GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponseWithDefaults() *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse {
	this := GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) SetMessage(v string) {
	o.Message = &v
}

func (o GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse struct {
	value *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse
	isSet bool
}

func (v NullableGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) Get() *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse {
	return v.value
}

func (v *NullableGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) Set(val *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse(val *GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) *NullableGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse {
	return &NullableGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse{value: val, isSet: true}
}

func (v NullableGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


