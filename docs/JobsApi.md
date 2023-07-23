# \JobsApi

All URIs are relative to *https://api.pdfgen.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JobsGet**](JobsApi.md#V1JobsGet) | **Get** /v1/jobs | List all Jobs
[**V1JobsIdDownloadGet**](JobsApi.md#V1JobsIdDownloadGet) | **Get** /v1/jobs/{id}/download | Download file



## V1JobsGet

> ListJobsResponse V1JobsGet(ctx).Execute()

List all Jobs



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
    resp, r, err := apiClient.JobsApi.V1JobsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.V1JobsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JobsGet`: ListJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.V1JobsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JobsGetRequest struct via the builder pattern


### Return type

[**ListJobsResponse**](ListJobsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JobsIdDownloadGet

> string V1JobsIdDownloadGet(ctx, id).Execute()

Download file



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
    id := "id_example" // string | Job ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.V1JobsIdDownloadGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.V1JobsIdDownloadGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JobsIdDownloadGet`: string
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.V1JobsIdDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JobsIdDownloadGetRequest struct via the builder pattern


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

