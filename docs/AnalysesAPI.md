# AnalysesAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalysisQueuePosition**](AnalysesAPI.md#GetAnalysisQueuePosition) | **Get** /v2/analyses/{analysis_id}/queue-position | Get the queue position of an analysis



## GetAnalysisQueuePosition

> QueuePositionResponse GetAnalysisQueuePosition(ctx, analysisId).Execute()

Get the queue position of an analysis



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
	resp, r, err := apiClient.AnalysesAPI.GetAnalysisQueuePosition(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesAPI.GetAnalysisQueuePosition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisQueuePosition`: QueuePositionResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesAPI.GetAnalysisQueuePosition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisQueuePositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueuePositionResponse**](QueuePositionResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

