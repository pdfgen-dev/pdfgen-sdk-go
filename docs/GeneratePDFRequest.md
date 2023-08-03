# GeneratePDFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Html** | **string** |  | 
**Variables** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGeneratePDFRequest

`func NewGeneratePDFRequest(html string, ) *GeneratePDFRequest`

NewGeneratePDFRequest instantiates a new GeneratePDFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratePDFRequestWithDefaults

`func NewGeneratePDFRequestWithDefaults() *GeneratePDFRequest`

NewGeneratePDFRequestWithDefaults instantiates a new GeneratePDFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHtml

`func (o *GeneratePDFRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *GeneratePDFRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *GeneratePDFRequest) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetVariables

`func (o *GeneratePDFRequest) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *GeneratePDFRequest) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *GeneratePDFRequest) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *GeneratePDFRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


