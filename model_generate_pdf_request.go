/*
PDFGen API

The PDFGen API for HTML to PDF generation.

API version: 1.0
Contact: support@pdfgen.dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdfgen

import (
	"encoding/json"
)

// checks if the GeneratePDFRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneratePDFRequest{}

// GeneratePDFRequest struct for GeneratePDFRequest
type GeneratePDFRequest struct {
	Html string `json:"html"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// NewGeneratePDFRequest instantiates a new GeneratePDFRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneratePDFRequest(html string) *GeneratePDFRequest {
	this := GeneratePDFRequest{}
	this.Html = html
	return &this
}

// NewGeneratePDFRequestWithDefaults instantiates a new GeneratePDFRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneratePDFRequestWithDefaults() *GeneratePDFRequest {
	this := GeneratePDFRequest{}
	return &this
}

// GetHtml returns the Html field value
func (o *GeneratePDFRequest) GetHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Html
}

// GetHtmlOk returns a tuple with the Html field value
// and a boolean to check if the value has been set.
func (o *GeneratePDFRequest) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Html, true
}

// SetHtml sets field value
func (o *GeneratePDFRequest) SetHtml(v string) {
	o.Html = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *GeneratePDFRequest) GetVariables() map[string]interface{} {
	if o == nil || IsNil(o.Variables) {
		var ret map[string]interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePDFRequest) GetVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Variables) {
		return map[string]interface{}{}, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *GeneratePDFRequest) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]interface{} and assigns it to the Variables field.
func (o *GeneratePDFRequest) SetVariables(v map[string]interface{}) {
	o.Variables = v
}

func (o GeneratePDFRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneratePDFRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["html"] = o.Html
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

type NullableGeneratePDFRequest struct {
	value *GeneratePDFRequest
	isSet bool
}

func (v NullableGeneratePDFRequest) Get() *GeneratePDFRequest {
	return v.value
}

func (v *NullableGeneratePDFRequest) Set(val *GeneratePDFRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratePDFRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratePDFRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratePDFRequest(val *GeneratePDFRequest) *NullableGeneratePDFRequest {
	return &NullableGeneratePDFRequest{value: val, isSet: true}
}

func (v NullableGeneratePDFRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneratePDFRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

