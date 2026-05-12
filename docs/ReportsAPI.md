# ReportsAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePdfReport**](ReportsAPI.md#CreatePdfReport) | **Post** /v3/analysis/{analysis_id}/pdf | Start PDF report generation
[**DownloadPdfReport**](ReportsAPI.md#DownloadPdfReport) | **Get** /v3/analysis/{analysis_id}/pdf/{task_id} | Download generated PDF report
[**GetPdfReportStatus**](ReportsAPI.md#GetPdfReportStatus) | **Get** /v3/analysis/{analysis_id}/pdf/{task_id}/status | Get PDF report workflow status



## CreatePdfReport

> GeneratePDFOutputBody CreatePdfReport(ctx, analysisId).Execute()

Start PDF report generation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v3"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.CreatePdfReport(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.CreatePdfReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePdfReport`: GeneratePDFOutputBody
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.CreatePdfReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePdfReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GeneratePDFOutputBody**](GeneratePDFOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadPdfReport

> DownloadPdfReport(ctx, analysisId, taskId).Execute()

Download generated PDF report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v3"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID
	taskId := "taskId_example" // string | Task ID returned by the create endpoint

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.DownloadPdfReport(context.Background(), analysisId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.DownloadPdfReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**taskId** | **string** | Task ID returned by the create endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadPdfReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPdfReportStatus

> WorkflowProgress GetPdfReportStatus(ctx, analysisId, taskId).Execute()

Get PDF report workflow status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v3"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID
	taskId := "taskId_example" // string | Task ID returned by the create endpoint

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetPdfReportStatus(context.Background(), analysisId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetPdfReportStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPdfReportStatus`: WorkflowProgress
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetPdfReportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**taskId** | **string** | Task ID returned by the create endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPdfReportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkflowProgress**](WorkflowProgress.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

