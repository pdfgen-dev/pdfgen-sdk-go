# \PdfApi

All URIs are relative to *https://api.pdfgen.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PdfPost**](PdfApi.md#V1PdfPost) | **Post** /v1/pdf | Generate PDF



## V1PdfPost

> map[string]interface{} V1PdfPost(ctx).Req(req).Execute()

Generate PDF



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
    req := *openapiclient.NewGithubComTcrisseappPdfgenInternalHandlersRestapiUtilsGeneratePDFRequest("Html_example") // GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsGeneratePDFRequest | HTML and variables to be converted to PDF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PdfApi.V1PdfPost(context.Background()).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PdfApi.V1PdfPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PdfPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PdfApi.V1PdfPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PdfPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **req** | [**GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsGeneratePDFRequest**](GithubComTcrisseappPdfgenInternalHandlersRestapiUtilsGeneratePDFRequest.md) | HTML and variables to be converted to PDF | 

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

