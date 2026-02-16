# AnalysesSecurityChecksAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScurityChecksTask**](AnalysesSecurityChecksAPI.md#CreateScurityChecksTask) | **Post** /v2/analyses/{analysis_id}/security-checks | Queues a security check process
[**GetSecurityChecks**](AnalysesSecurityChecksAPI.md#GetSecurityChecks) | **Get** /v2/analyses/{analysis_id}/security-checks | Get Security Checks
[**GetSecurityChecksTaskStatus**](AnalysesSecurityChecksAPI.md#GetSecurityChecksTaskStatus) | **Get** /v2/analyses/{analysis_id}/security-checks/status | Check the status of a security check process



## CreateScurityChecksTask

> QueuedSecurityChecksTaskResponse CreateScurityChecksTask(ctx, analysisId).Execute()

Queues a security check process

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesSecurityChecksAPI.CreateScurityChecksTask(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesSecurityChecksAPI.CreateScurityChecksTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScurityChecksTask`: QueuedSecurityChecksTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesSecurityChecksAPI.CreateScurityChecksTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScurityChecksTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueuedSecurityChecksTaskResponse**](QueuedSecurityChecksTaskResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityChecks

> BaseResponseSecurityChecksResponse GetSecurityChecks(ctx, analysisId).Page(page).PageSize(pageSize).Execute()

Get Security Checks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 
	page := int32(56) // int32 | The page number to retrieve.
	pageSize := int32(56) // int32 | Number of items per page.

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesSecurityChecksAPI.GetSecurityChecks(context.Background(), analysisId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesSecurityChecksAPI.GetSecurityChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityChecks`: BaseResponseSecurityChecksResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesSecurityChecksAPI.GetSecurityChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number to retrieve. | 
 **pageSize** | **int32** | Number of items per page. | 

### Return type

[**BaseResponseSecurityChecksResponse**](BaseResponseSecurityChecksResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityChecksTaskStatus

> CheckSecurityChecksTaskResponse GetSecurityChecksTaskStatus(ctx, analysisId).Execute()

Check the status of a security check process

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesSecurityChecksAPI.GetSecurityChecksTaskStatus(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesSecurityChecksAPI.GetSecurityChecksTaskStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityChecksTaskStatus`: CheckSecurityChecksTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesSecurityChecksAPI.GetSecurityChecksTaskStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityChecksTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckSecurityChecksTaskResponse**](CheckSecurityChecksTaskResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

