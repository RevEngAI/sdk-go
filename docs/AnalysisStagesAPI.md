# AnalysisStagesAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalysisStages**](AnalysisStagesAPI.md#GetAnalysisStages) | **Get** /v2/analysis-stages/{analysis_id} | Get Analysis Stages
[**GetPipelineStatus**](AnalysisStagesAPI.md#GetPipelineStatus) | **Get** /v2/analysis-stages/{analysis_id}/pipeline-status | Get Pipeline Status



## GetAnalysisStages

> BaseResponseAnalysisStagesResponse GetAnalysisStages(ctx, analysisId).Execute()

Get Analysis Stages



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
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisStagesAPI.GetAnalysisStages(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisStagesAPI.GetAnalysisStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisStages`: BaseResponseAnalysisStagesResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysisStagesAPI.GetAnalysisStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseAnalysisStagesResponse**](BaseResponseAnalysisStagesResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineStatus

> BaseResponsePipelineStatusResponse GetPipelineStatus(ctx, analysisId).Execute()

Get Pipeline Status



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
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisStagesAPI.GetPipelineStatus(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisStagesAPI.GetPipelineStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelineStatus`: BaseResponsePipelineStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysisStagesAPI.GetPipelineStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponsePipelineStatusResponse**](BaseResponsePipelineStatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

