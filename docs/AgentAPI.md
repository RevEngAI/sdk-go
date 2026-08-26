# AgentAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet**](AgentAPI.md#CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/capabilities/status | Check the status of a capabilities analysis workflow
[**CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet**](AgentAPI.md#CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/remediation/status | Check the status of a remediation analysis workflow
[**CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet**](AgentAPI.md#CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/report-analysis/status | Check the status of a report analysis workflow
[**CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet**](AgentAPI.md#CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/triage/status | Check the status of a triage analysis workflow
[**CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost**](AgentAPI.md#CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost) | **Post** /v2/analyses/{analysis_id}/agent/capabilities | Queues a capabilities analysis workflow process
[**CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost**](AgentAPI.md#CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost) | **Post** /v2/analyses/{analysis_id}/agent/remediation | Queues a remediation analysis workflow process
[**CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost**](AgentAPI.md#CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost) | **Post** /v2/analyses/{analysis_id}/agent/report-analysis | Queues a combined report analysis workflow process
[**CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost**](AgentAPI.md#CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost) | **Post** /v2/analyses/{analysis_id}/agent/triage | Queues a triage analysis workflow process
[**GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet**](AgentAPI.md#GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet) | **Get** /v2/analyses/{analysis_id}/agent/capabilities | Get Capabilities Result
[**GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet**](AgentAPI.md#GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet) | **Get** /v2/analyses/{analysis_id}/agent/remediation | Get Remediation Result
[**GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet**](AgentAPI.md#GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet) | **Get** /v2/analyses/{analysis_id}/agent/report-analysis | Get Report Analysis Result
[**GetTriageResultV2AnalysesAnalysisIdAgentTriageGet**](AgentAPI.md#GetTriageResultV2AnalysesAnalysisIdAgentTriageGet) | **Get** /v2/analyses/{analysis_id}/agent/triage | Get Triage Result
[**V3CancelRenameUnnamedFunctions**](AgentAPI.md#V3CancelRenameUnnamedFunctions) | **Post** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions/cancel | Cancel the rename-unnamed-functions agent.
[**V3CancelSecurityScanOperation**](AgentAPI.md#V3CancelSecurityScanOperation) | **Post** /v3/operations/security-scan/{analysis_id}:cancel | Cancel a security-scan operation.
[**V3GetRenameUnnamedFunctionsResult**](AgentAPI.md#V3GetRenameUnnamedFunctionsResult) | **Get** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions | Get rename-unnamed-functions agent result.
[**V3GetRenameUnnamedFunctionsStatus**](AgentAPI.md#V3GetRenameUnnamedFunctionsStatus) | **Get** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions/status | Get rename-unnamed-functions agent status.
[**V3GetSecurityScanOperation**](AgentAPI.md#V3GetSecurityScanOperation) | **Get** /v3/operations/security-scan/{analysis_id} | Get a security-scan operation.
[**V3RunSecurityScan**](AgentAPI.md#V3RunSecurityScan) | **Post** /v3/analyses/{analysis_id}/security-scan:run | Run the security-scan agent.
[**V3TriggerRenameUnnamedFunctions**](AgentAPI.md#V3TriggerRenameUnnamedFunctions) | **Post** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions | Run the rename-unnamed-functions agent.



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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet

> TaskStatusResponse CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet(ctx, analysisId).Execute()

Check the status of a remediation analysis workflow

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet`: TaskStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskStatusResponse**](TaskStatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost

> BaseResponseQueuedWorkflowTaskResponse CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost(ctx, analysisId).Execute()

Queues a remediation analysis workflow process

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost`: BaseResponseQueuedWorkflowTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseQueuedWorkflowTaskResponse**](BaseResponseQueuedWorkflowTaskResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet

> BaseResponseRemediationAgentResponse GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet(ctx, analysisId).Execute()

Get Remediation Result



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet`: BaseResponseRemediationAgentResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemediationResultV2AnalysesAnalysisIdAgentRemediationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseRemediationAgentResponse**](BaseResponseRemediationAgentResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3CancelRenameUnnamedFunctions

> V3CancelRenameUnnamedFunctions(ctx, analysisId).Execute()

Cancel the rename-unnamed-functions agent.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.V3CancelRenameUnnamedFunctions(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3CancelRenameUnnamedFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3CancelRenameUnnamedFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3CancelSecurityScanOperation

> V3CancelSecurityScanOperation(ctx, analysisId).Execute()

Cancel a security-scan operation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.V3CancelSecurityScanOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3CancelSecurityScanOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3CancelSecurityScanOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetRenameUnnamedFunctionsResult

> RenameUnnamedFunctionsResult V3GetRenameUnnamedFunctionsResult(ctx, analysisId).Execute()

Get rename-unnamed-functions agent result.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3GetRenameUnnamedFunctionsResult(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetRenameUnnamedFunctionsResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetRenameUnnamedFunctionsResult`: RenameUnnamedFunctionsResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetRenameUnnamedFunctionsResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetRenameUnnamedFunctionsResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RenameUnnamedFunctionsResult**](RenameUnnamedFunctionsResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetRenameUnnamedFunctionsStatus

> StatusBody V3GetRenameUnnamedFunctionsStatus(ctx, analysisId).Execute()

Get rename-unnamed-functions agent status.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3GetRenameUnnamedFunctionsStatus(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetRenameUnnamedFunctionsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetRenameUnnamedFunctionsStatus`: StatusBody
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetRenameUnnamedFunctionsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetRenameUnnamedFunctionsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusBody**](StatusBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetSecurityScanOperation

> OperationSecurityScanMetadataSecurityScanResult V3GetSecurityScanOperation(ctx, analysisId).Execute()

Get a security-scan operation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3GetSecurityScanOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetSecurityScanOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetSecurityScanOperation`: OperationSecurityScanMetadataSecurityScanResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetSecurityScanOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetSecurityScanOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationSecurityScanMetadataSecurityScanResult**](OperationSecurityScanMetadataSecurityScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunSecurityScan

> OperationSecurityScanMetadataSecurityScanResult V3RunSecurityScan(ctx, analysisId).TriggerSecurityScanInputBody(triggerSecurityScanInputBody).Execute()

Run the security-scan agent.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID
	triggerSecurityScanInputBody := *revengai.NewTriggerSecurityScanInputBody() // TriggerSecurityScanInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3RunSecurityScan(context.Background(), analysisId).TriggerSecurityScanInputBody(triggerSecurityScanInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunSecurityScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunSecurityScan`: OperationSecurityScanMetadataSecurityScanResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunSecurityScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunSecurityScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerSecurityScanInputBody** | [**TriggerSecurityScanInputBody**](TriggerSecurityScanInputBody.md) |  | 

### Return type

[**OperationSecurityScanMetadataSecurityScanResult**](OperationSecurityScanMetadataSecurityScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3TriggerRenameUnnamedFunctions

> StatusBody V3TriggerRenameUnnamedFunctions(ctx, analysisId).TriggerRenameUnnamedFunctionsInputBody(triggerRenameUnnamedFunctionsInputBody).Execute()

Run the rename-unnamed-functions agent.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID
	triggerRenameUnnamedFunctionsInputBody := *revengai.NewTriggerRenameUnnamedFunctionsInputBody() // TriggerRenameUnnamedFunctionsInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3TriggerRenameUnnamedFunctions(context.Background(), analysisId).TriggerRenameUnnamedFunctionsInputBody(triggerRenameUnnamedFunctionsInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3TriggerRenameUnnamedFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3TriggerRenameUnnamedFunctions`: StatusBody
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3TriggerRenameUnnamedFunctions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3TriggerRenameUnnamedFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerRenameUnnamedFunctionsInputBody** | [**TriggerRenameUnnamedFunctionsInputBody**](TriggerRenameUnnamedFunctionsInputBody.md) |  | 

### Return type

[**StatusBody**](StatusBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

