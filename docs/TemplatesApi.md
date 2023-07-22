# \TemplatesApi

All URIs are relative to *https://api.pdfgen.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TemplatesGet**](TemplatesApi.md#V1TemplatesGet) | **Get** /v1/templates | List all Templates
[**V1TemplatesIdDelete**](TemplatesApi.md#V1TemplatesIdDelete) | **Delete** /v1/templates/{id} | Delete Template
[**V1TemplatesIdGet**](TemplatesApi.md#V1TemplatesIdGet) | **Get** /v1/templates/{id} | Get Templates
[**V1TemplatesIdPdfPost**](TemplatesApi.md#V1TemplatesIdPdfPost) | **Post** /v1/templates/{id}/pdf | Generate PDF from Template
[**V1TemplatesIdPreviewPost**](TemplatesApi.md#V1TemplatesIdPreviewPost) | **Post** /v1/templates/{id}/preview | Previews PDF with given variables
[**V1TemplatesIdPut**](TemplatesApi.md#V1TemplatesIdPut) | **Put** /v1/templates/{id} | Update Template
[**V1TemplatesPost**](TemplatesApi.md#V1TemplatesPost) | **Post** /v1/templates | Create Template



## V1TemplatesGet

> ListTemplatesResponse V1TemplatesGet(ctx).Execute()

List all Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.V1TemplatesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.V1TemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TemplatesGet`: ListTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.V1TemplatesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TemplatesGetRequest struct via the builder pattern


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


## V1TemplatesIdDelete

> string V1TemplatesIdDelete(ctx, id).Execute()

Delete Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.V1TemplatesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.V1TemplatesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TemplatesIdDelete`: string
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.V1TemplatesIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TemplatesIdDeleteRequest struct via the builder pattern


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


## V1TemplatesIdGet

> TemplateResponse V1TemplatesIdGet(ctx, id).Execute()

Get Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.V1TemplatesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.V1TemplatesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TemplatesIdGet`: TemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.V1TemplatesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TemplatesIdGetRequest struct via the builder pattern


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


## V1TemplatesIdPdfPost

> map[string]interface{} V1TemplatesIdPdfPost(ctx, id).Variables(variables).Execute()

Generate PDF from Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Template ID
    variables := map[string]interface{}{ ... } // map[string]interface{} | Variables used for the template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.V1TemplatesIdPdfPost(context.Background(), id).Variables(variables).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.V1TemplatesIdPdfPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TemplatesIdPdfPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.V1TemplatesIdPdfPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TemplatesIdPdfPostRequest struct via the builder pattern


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


## V1TemplatesIdPreviewPost

> map[string]interface{} V1TemplatesIdPreviewPost(ctx, id).Variables(variables).Execute()

Previews PDF with given variables



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Template ID or alias
    variables := map[string]interface{}{ ... } // map[string]interface{} | Variables used for the template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.V1TemplatesIdPreviewPost(context.Background(), id).Variables(variables).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.V1TemplatesIdPreviewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TemplatesIdPreviewPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.V1TemplatesIdPreviewPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template ID or alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TemplatesIdPreviewPostRequest struct via the builder pattern


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


## V1TemplatesIdPut

> TemplateResponse V1TemplatesIdPut(ctx, id).Req(req).Execute()

Update Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Template ID
    req := *openapiclient.NewGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsUpdateTemplateRequest() // GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsUpdateTemplateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.V1TemplatesIdPut(context.Background(), id).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.V1TemplatesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TemplatesIdPut`: TemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.V1TemplatesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TemplatesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **req** | [**GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsUpdateTemplateRequest**](GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsUpdateTemplateRequest.md) |  | 

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


## V1TemplatesPost

> TemplateResponse V1TemplatesPost(ctx).Req(req).Execute()

Create Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    req := *openapiclient.NewGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsCreateTemplateRequest("Alias_example", "Html_example", "Name_example") // GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsCreateTemplateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.V1TemplatesPost(context.Background()).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.V1TemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TemplatesPost`: TemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.V1TemplatesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **req** | [**GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsCreateTemplateRequest**](GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsCreateTemplateRequest.md) |  | 

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

