# \JobsApi

All URIs are relative to *https://api.pdfgen.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadFile**](JobsApi.md#DownloadFile) | **Get** /v1/jobs/{id}/download | Download file
[**ListJobs**](JobsApi.md#ListJobs) | **Get** /v1/jobs | List all Jobs



## DownloadFile

> string DownloadFile(ctx, id).Execute()

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
    resp, r, err := apiClient.JobsApi.DownloadFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.DownloadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFile`: string
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.DownloadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileRequest struct via the builder pattern


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


## ListJobs

> ListJobsResponse ListJobs(ctx).Execute()

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
    resp, r, err := apiClient.JobsApi.ListJobs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.ListJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobs`: ListJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.ListJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListJobsRequest struct via the builder pattern


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

