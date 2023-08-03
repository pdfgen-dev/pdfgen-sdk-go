# \TemplatesApi

All URIs are relative to *https://api.pdfgen.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplate**](TemplatesApi.md#CreateTemplate) | **Post** /v1/templates | Create Template
[**DeleteTemplate**](TemplatesApi.md#DeleteTemplate) | **Delete** /v1/templates/{id} | Delete Template
[**GeneratePdfFromTemplate**](TemplatesApi.md#GeneratePdfFromTemplate) | **Post** /v1/templates/{id}/pdf | Generate PDF from Template
[**GetTemplate**](TemplatesApi.md#GetTemplate) | **Get** /v1/templates/{id} | Get Templates
[**ListTemplates**](TemplatesApi.md#ListTemplates) | **Get** /v1/templates | List all Templates
[**PreviewTemplate**](TemplatesApi.md#PreviewTemplate) | **Post** /v1/templates/{id}/preview | Previews PDF with given variables
[**UpdateTemplate**](TemplatesApi.md#UpdateTemplate) | **Put** /v1/templates/{id} | Update Template



## CreateTemplate

> TemplateResponse CreateTemplate(ctx).Req(req).Execute()

Create Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pdfgen-dev/pdfgen-sdk-go"
)

func main() {
    req := *openapiclient.NewCreateTemplateRequest("Alias_example", "Html_example", "Name_example") // CreateTemplateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.CreateTemplate(context.Background()).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.CreateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTemplate`: TemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **req** | [**CreateTemplateRequest**](CreateTemplateRequest.md) |  | 

### Return type

[**TemplateResponse**](TemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplate

> string DeleteTemplate(ctx, id).Execute()

Delete Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pdfgen-dev/pdfgen-sdk-go"
)

func main() {
    id := "id_example" // string | Template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.DeleteTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.DeleteTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTemplate`: string
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.DeleteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneratePdfFromTemplate

> map[string]interface{} GeneratePdfFromTemplate(ctx, id).Variables(variables).Execute()

Generate PDF from Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pdfgen-dev/pdfgen-sdk-go"
)

func main() {
    id := "id_example" // string | Template ID
    variables := map[string]interface{}{ ... } // map[string]interface{} | Variables used for the template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.GeneratePdfFromTemplate(context.Background(), id).Variables(variables).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.GeneratePdfFromTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneratePdfFromTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.GeneratePdfFromTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePdfFromTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | **map[string]interface{}** | Variables used for the template | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplate

> TemplateResponse GetTemplate(ctx, id).Execute()

Get Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pdfgen-dev/pdfgen-sdk-go"
)

func main() {
    id := "id_example" // string | Template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.GetTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.GetTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplate`: TemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.GetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateResponse**](TemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplates

> ListTemplatesResponse ListTemplates(ctx).Execute()

List all Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pdfgen-dev/pdfgen-sdk-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.ListTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.ListTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTemplates`: ListTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.ListTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplatesRequest struct via the builder pattern


### Return type

[**ListTemplatesResponse**](ListTemplatesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewTemplate

> map[string]interface{} PreviewTemplate(ctx, id).Variables(variables).Execute()

Previews PDF with given variables



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pdfgen-dev/pdfgen-sdk-go"
)

func main() {
    id := "id_example" // string | Template ID or alias
    variables := map[string]interface{}{ ... } // map[string]interface{} | Variables used for the template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.PreviewTemplate(context.Background(), id).Variables(variables).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.PreviewTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.PreviewTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template ID or alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variables** | **map[string]interface{}** | Variables used for the template | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplate

> TemplateResponse UpdateTemplate(ctx, id).Req(req).Execute()

Update Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pdfgen-dev/pdfgen-sdk-go"
)

func main() {
    id := "id_example" // string | Template ID
    req := *openapiclient.NewUpdateTemplateRequest() // UpdateTemplateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.UpdateTemplate(context.Background(), id).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.UpdateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplate`: TemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.UpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **req** | [**UpdateTemplateRequest**](UpdateTemplateRequest.md) |  | 

### Return type

[**TemplateResponse**](TemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

