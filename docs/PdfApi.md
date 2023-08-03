# \PdfApi

All URIs are relative to *https://api.pdfgen.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GeneratePdf**](PdfApi.md#GeneratePdf) | **Post** /v1/pdf | Generate PDF



## GeneratePdf

> map[string]interface{} GeneratePdf(ctx).Req(req).Execute()

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
    req := *openapiclient.NewGeneratePDFRequest("Html_example") // GeneratePDFRequest | HTML and variables to be converted to PDF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PdfApi.GeneratePdf(context.Background()).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PdfApi.GeneratePdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneratePdf`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PdfApi.GeneratePdf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **req** | [**GeneratePDFRequest**](GeneratePDFRequest.md) | HTML and variables to be converted to PDF | 

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

