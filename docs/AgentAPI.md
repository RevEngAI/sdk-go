# AgentAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet**](AgentAPI.md#CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/capabilities/status | Check the status of a capabilities analysis workflow
[**CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet**](AgentAPI.md#CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/report-analysis/status | Check the status of a report analysis workflow
[**CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet**](AgentAPI.md#CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/triage/status | Check the status of a triage analysis workflow
[**CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost**](AgentAPI.md#CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost) | **Post** /v2/analyses/{analysis_id}/agent/capabilities | Queues a capabilities analysis workflow process
[**CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost**](AgentAPI.md#CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost) | **Post** /v2/analyses/{analysis_id}/agent/report-analysis | Queues a combined report analysis workflow process
[**CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost**](AgentAPI.md#CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost) | **Post** /v2/analyses/{analysis_id}/agent/triage | Queues a triage analysis workflow process
[**GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet**](AgentAPI.md#GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet) | **Get** /v2/analyses/{analysis_id}/agent/capabilities | Get Capabilities Result
[**GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet**](AgentAPI.md#GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet) | **Get** /v2/analyses/{analysis_id}/agent/report-analysis | Get Report Analysis Result
[**GetTriageResultV2AnalysesAnalysisIdAgentTriageGet**](AgentAPI.md#GetTriageResultV2AnalysesAnalysisIdAgentTriageGet) | **Get** /v2/analyses/{analysis_id}/agent/triage | Get Triage Result



## CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet

> TaskStatusResponse CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet(ctx, analysisId).Execute()

Check the status of a capabilities analysis workflow

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
	resp, r, err := apiClient.AgentAPI.CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet`: TaskStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskStatusResponse**](TaskStatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet

> TaskStatusResponse CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet(ctx, analysisId).Execute()

Check the status of a report analysis workflow

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
	resp, r, err := apiClient.AgentAPI.CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet`: TaskStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskStatusResponse**](TaskStatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet

> TaskStatusResponse CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet(ctx, analysisId).Execute()

Check the status of a triage analysis workflow

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
	resp, r, err := apiClient.AgentAPI.CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet`: TaskStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskStatusResponse**](TaskStatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost

> BaseResponseQueuedWorkflowTaskResponse CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost(ctx, analysisId).Execute()

Queues a capabilities analysis workflow process

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
	resp, r, err := apiClient.AgentAPI.CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost`: BaseResponseQueuedWorkflowTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseQueuedWorkflowTaskResponse**](BaseResponseQueuedWorkflowTaskResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost

> QueuedWorkflowTaskResponse CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost(ctx, analysisId).Execute()

Queues a combined report analysis workflow process

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
	resp, r, err := apiClient.AgentAPI.CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost`: QueuedWorkflowTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueuedWorkflowTaskResponse**](QueuedWorkflowTaskResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost

> BaseResponseQueuedWorkflowTaskResponse CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost(ctx, analysisId).Execute()

Queues a triage analysis workflow process

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
	resp, r, err := apiClient.AgentAPI.CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost`: BaseResponseQueuedWorkflowTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTriageTaskV2AnalysesAnalysisIdAgentTriagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseQueuedWorkflowTaskResponse**](BaseResponseQueuedWorkflowTaskResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet

> BaseResponseCapabilitiesAgentResponse GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet(ctx, analysisId).Execute()

Get Capabilities Result

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
	resp, r, err := apiClient.AgentAPI.GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet`: BaseResponseCapabilitiesAgentResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseCapabilitiesAgentResponse**](BaseResponseCapabilitiesAgentResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet

> BaseResponseReportAnalysisResponse GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet(ctx, analysisId).Execute()

Get Report Analysis Result



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
	resp, r, err := apiClient.AgentAPI.GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet`: BaseResponseReportAnalysisResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseReportAnalysisResponse**](BaseResponseReportAnalysisResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTriageResultV2AnalysesAnalysisIdAgentTriageGet

> BaseResponseTriageReportResponse GetTriageResultV2AnalysesAnalysisIdAgentTriageGet(ctx, analysisId).Execute()

Get Triage Result

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
	resp, r, err := apiClient.AgentAPI.GetTriageResultV2AnalysesAnalysisIdAgentTriageGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetTriageResultV2AnalysesAnalysisIdAgentTriageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTriageResultV2AnalysesAnalysisIdAgentTriageGet`: BaseResponseTriageReportResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetTriageResultV2AnalysesAnalysisIdAgentTriageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTriageResultV2AnalysesAnalysisIdAgentTriageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseTriageReportResponse**](BaseResponseTriageReportResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

