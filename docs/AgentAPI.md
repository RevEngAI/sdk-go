# AgentAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet**](AgentAPI.md#CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/capabilities/status | Check the status of a capabilities analysis workflow
[**CheckProtocolsTaskStatusV2AnalysesAnalysisIdAgentProtocolsStatusGet**](AgentAPI.md#CheckProtocolsTaskStatusV2AnalysesAnalysisIdAgentProtocolsStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/protocols/status | Check the status of a protocols discovery workflow
[**CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet**](AgentAPI.md#CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/remediation/status | Check the status of a remediation analysis workflow
[**CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet**](AgentAPI.md#CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/report-analysis/status | Check the status of a report analysis workflow
[**CheckSecretsTaskStatusV2AnalysesAnalysisIdAgentSecretsStatusGet**](AgentAPI.md#CheckSecretsTaskStatusV2AnalysesAnalysisIdAgentSecretsStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/secrets/status | Check the status of a secrets discovery workflow
[**CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet**](AgentAPI.md#CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet) | **Get** /v2/analyses/{analysis_id}/agent/triage/status | Check the status of a triage analysis workflow
[**CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost**](AgentAPI.md#CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost) | **Post** /v2/analyses/{analysis_id}/agent/capabilities | Queues a capabilities analysis workflow process
[**CreateProtocolsTaskV2AnalysesAnalysisIdAgentProtocolsPost**](AgentAPI.md#CreateProtocolsTaskV2AnalysesAnalysisIdAgentProtocolsPost) | **Post** /v2/analyses/{analysis_id}/agent/protocols | Queues a protocols discovery workflow process
[**CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost**](AgentAPI.md#CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost) | **Post** /v2/analyses/{analysis_id}/agent/remediation | Queues a remediation analysis workflow process
[**CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost**](AgentAPI.md#CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost) | **Post** /v2/analyses/{analysis_id}/agent/report-analysis | Queues a combined report analysis workflow process
[**CreateSecretsTaskV2AnalysesAnalysisIdAgentSecretsPost**](AgentAPI.md#CreateSecretsTaskV2AnalysesAnalysisIdAgentSecretsPost) | **Post** /v2/analyses/{analysis_id}/agent/secrets | Queues a secrets discovery workflow process
[**CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost**](AgentAPI.md#CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost) | **Post** /v2/analyses/{analysis_id}/agent/triage | Queues a triage analysis workflow process
[**GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet**](AgentAPI.md#GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet) | **Get** /v2/analyses/{analysis_id}/agent/capabilities | Get Capabilities Result
[**GetProtocolsResultV2AnalysesAnalysisIdAgentProtocolsGet**](AgentAPI.md#GetProtocolsResultV2AnalysesAnalysisIdAgentProtocolsGet) | **Get** /v2/analyses/{analysis_id}/agent/protocols | Get Protocols Result
[**GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet**](AgentAPI.md#GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet) | **Get** /v2/analyses/{analysis_id}/agent/remediation | Get Remediation Result
[**GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet**](AgentAPI.md#GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet) | **Get** /v2/analyses/{analysis_id}/agent/report-analysis | Get Report Analysis Result
[**GetSecretsResultV2AnalysesAnalysisIdAgentSecretsGet**](AgentAPI.md#GetSecretsResultV2AnalysesAnalysisIdAgentSecretsGet) | **Get** /v2/analyses/{analysis_id}/agent/secrets | Get Secrets Result
[**GetTriageResultV2AnalysesAnalysisIdAgentTriageGet**](AgentAPI.md#GetTriageResultV2AnalysesAnalysisIdAgentTriageGet) | **Get** /v2/analyses/{analysis_id}/agent/triage | Get Triage Result
[**V3CancelRenameUnnamedFunctions**](AgentAPI.md#V3CancelRenameUnnamedFunctions) | **Post** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions/cancel | Cancel the rename-unnamed-functions agent.
[**V3CancelSecurityScanOperation**](AgentAPI.md#V3CancelSecurityScanOperation) | **Post** /v3/operations/security-scan/{analysis_id}:cancel | Cancel a security-scan operation.
[**V3GetBinaryAgentFeedback**](AgentAPI.md#V3GetBinaryAgentFeedback) | **Get** /v3/analyses/{analysis_id}/agents/{agent}/feedback | Get the caller&#39;s feedback on an agent&#39;s output.
[**V3GetCapabilitiesOperation**](AgentAPI.md#V3GetCapabilitiesOperation) | **Get** /v3/operations/capabilities/{analysis_id} | Get a capabilities operation.
[**V3GetCryptoExplainOperation**](AgentAPI.md#V3GetCryptoExplainOperation) | **Get** /v3/operations/crypto-explain/{function_id} | Get a crypto-explain operation.
[**V3GetCryptoScanOperation**](AgentAPI.md#V3GetCryptoScanOperation) | **Get** /v3/operations/crypto-scan/{analysis_id} | Get a crypto-scan operation.
[**V3GetExecutionExplainOperation**](AgentAPI.md#V3GetExecutionExplainOperation) | **Get** /v3/operations/execution-explain/{function_id} | Get an execution-explain operation.
[**V3GetExecutionScanOperation**](AgentAPI.md#V3GetExecutionScanOperation) | **Get** /v3/operations/execution-scan/{analysis_id} | Get an execution-scan operation.
[**V3GetFilesystemAnalyseOperation**](AgentAPI.md#V3GetFilesystemAnalyseOperation) | **Get** /v3/operations/filesystem-analyse/{function_id} | Get a filesystem-analyse operation.
[**V3GetFilesystemScanOperation**](AgentAPI.md#V3GetFilesystemScanOperation) | **Get** /v3/operations/filesystem-scan/{analysis_id} | Get a filesystem-scan operation.
[**V3GetNetworkingExplainOperation**](AgentAPI.md#V3GetNetworkingExplainOperation) | **Get** /v3/operations/networking-explain/{function_id} | Get a networking-explain operation.
[**V3GetNetworkingScanOperation**](AgentAPI.md#V3GetNetworkingScanOperation) | **Get** /v3/operations/networking-scan/{analysis_id} | Get a networking-scan operation.
[**V3GetProtocolsOperation**](AgentAPI.md#V3GetProtocolsOperation) | **Get** /v3/operations/protocols/{analysis_id} | Get a protocols operation.
[**V3GetRemediationOperation**](AgentAPI.md#V3GetRemediationOperation) | **Get** /v3/operations/remediation/{analysis_id} | Get a remediation operation.
[**V3GetRenameUnnamedFunctionsResult**](AgentAPI.md#V3GetRenameUnnamedFunctionsResult) | **Get** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions | Get rename-unnamed-functions agent result.
[**V3GetRenameUnnamedFunctionsStatus**](AgentAPI.md#V3GetRenameUnnamedFunctionsStatus) | **Get** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions/status | Get rename-unnamed-functions agent status.
[**V3GetReportAnalysisOperation**](AgentAPI.md#V3GetReportAnalysisOperation) | **Get** /v3/operations/report-analysis/{analysis_id} | Get a report-analysis operation.
[**V3GetSecretsOperation**](AgentAPI.md#V3GetSecretsOperation) | **Get** /v3/operations/secrets/{analysis_id} | Get a secrets operation.
[**V3GetSecurityScanOperation**](AgentAPI.md#V3GetSecurityScanOperation) | **Get** /v3/operations/security-scan/{analysis_id} | Get a security-scan operation.
[**V3GetTriageOperation**](AgentAPI.md#V3GetTriageOperation) | **Get** /v3/operations/triage/{analysis_id} | Get a triage operation.
[**V3RunCapabilities**](AgentAPI.md#V3RunCapabilities) | **Post** /v3/analyses/{analysis_id}/capabilities:run | Run the capabilities agent.
[**V3RunCryptoExplain**](AgentAPI.md#V3RunCryptoExplain) | **Post** /v3/functions/{function_id}/crypto-explain:run | Run the crypto-explain agent.
[**V3RunCryptoScan**](AgentAPI.md#V3RunCryptoScan) | **Post** /v3/analyses/{analysis_id}/crypto-scan:run | Run the crypto-scan agent.
[**V3RunExecutionExplain**](AgentAPI.md#V3RunExecutionExplain) | **Post** /v3/functions/{function_id}/execution-explain:run | Run the execution-explain agent.
[**V3RunExecutionScan**](AgentAPI.md#V3RunExecutionScan) | **Post** /v3/analyses/{analysis_id}/execution-scan:run | Run the execution-scan agent.
[**V3RunFilesystemAnalyse**](AgentAPI.md#V3RunFilesystemAnalyse) | **Post** /v3/functions/{function_id}/filesystem-analyse:run | Run the filesystem-analyse agent.
[**V3RunFilesystemScan**](AgentAPI.md#V3RunFilesystemScan) | **Post** /v3/analyses/{analysis_id}/filesystem-scan:run | Run the filesystem-scan agent.
[**V3RunNetworkingExplain**](AgentAPI.md#V3RunNetworkingExplain) | **Post** /v3/functions/{function_id}/networking-explain:run | Run the networking-explain agent.
[**V3RunNetworkingScan**](AgentAPI.md#V3RunNetworkingScan) | **Post** /v3/analyses/{analysis_id}/networking-scan:run | Run the networking-scan agent.
[**V3RunProtocols**](AgentAPI.md#V3RunProtocols) | **Post** /v3/analyses/{analysis_id}/protocols:run | Run the protocols agent.
[**V3RunRemediation**](AgentAPI.md#V3RunRemediation) | **Post** /v3/analyses/{analysis_id}/remediation:run | Run the remediation agent.
[**V3RunReportAnalysis**](AgentAPI.md#V3RunReportAnalysis) | **Post** /v3/analyses/{analysis_id}/report-analysis:run | Run the report-analysis agent.
[**V3RunSecrets**](AgentAPI.md#V3RunSecrets) | **Post** /v3/analyses/{analysis_id}/secrets:run | Run the secrets agent.
[**V3RunSecurityScan**](AgentAPI.md#V3RunSecurityScan) | **Post** /v3/analyses/{analysis_id}/security-scan:run | Run the security-scan agent.
[**V3RunTriage**](AgentAPI.md#V3RunTriage) | **Post** /v3/analyses/{analysis_id}/triage:run | Run the triage agent.
[**V3TriggerRenameUnnamedFunctions**](AgentAPI.md#V3TriggerRenameUnnamedFunctions) | **Post** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions | Run the rename-unnamed-functions agent.
[**V3UpsertBinaryAgentFeedback**](AgentAPI.md#V3UpsertBinaryAgentFeedback) | **Put** /v3/analyses/{analysis_id}/agents/{agent}/feedback | Record feedback on an agent&#39;s output.



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


## CheckProtocolsTaskStatusV2AnalysesAnalysisIdAgentProtocolsStatusGet

> TaskStatusResponse CheckProtocolsTaskStatusV2AnalysesAnalysisIdAgentProtocolsStatusGet(ctx, analysisId).Execute()

Check the status of a protocols discovery workflow

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
	resp, r, err := apiClient.AgentAPI.CheckProtocolsTaskStatusV2AnalysesAnalysisIdAgentProtocolsStatusGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CheckProtocolsTaskStatusV2AnalysesAnalysisIdAgentProtocolsStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckProtocolsTaskStatusV2AnalysesAnalysisIdAgentProtocolsStatusGet`: TaskStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CheckProtocolsTaskStatusV2AnalysesAnalysisIdAgentProtocolsStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckProtocolsTaskStatusV2AnalysesAnalysisIdAgentProtocolsStatusGetRequest struct via the builder pattern


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


## CheckSecretsTaskStatusV2AnalysesAnalysisIdAgentSecretsStatusGet

> TaskStatusResponse CheckSecretsTaskStatusV2AnalysesAnalysisIdAgentSecretsStatusGet(ctx, analysisId).Execute()

Check the status of a secrets discovery workflow

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
	resp, r, err := apiClient.AgentAPI.CheckSecretsTaskStatusV2AnalysesAnalysisIdAgentSecretsStatusGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CheckSecretsTaskStatusV2AnalysesAnalysisIdAgentSecretsStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckSecretsTaskStatusV2AnalysesAnalysisIdAgentSecretsStatusGet`: TaskStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CheckSecretsTaskStatusV2AnalysesAnalysisIdAgentSecretsStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckSecretsTaskStatusV2AnalysesAnalysisIdAgentSecretsStatusGetRequest struct via the builder pattern


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


## CreateProtocolsTaskV2AnalysesAnalysisIdAgentProtocolsPost

> BaseResponseQueuedWorkflowTaskResponse CreateProtocolsTaskV2AnalysesAnalysisIdAgentProtocolsPost(ctx, analysisId).Execute()

Queues a protocols discovery workflow process

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
	resp, r, err := apiClient.AgentAPI.CreateProtocolsTaskV2AnalysesAnalysisIdAgentProtocolsPost(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CreateProtocolsTaskV2AnalysesAnalysisIdAgentProtocolsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProtocolsTaskV2AnalysesAnalysisIdAgentProtocolsPost`: BaseResponseQueuedWorkflowTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CreateProtocolsTaskV2AnalysesAnalysisIdAgentProtocolsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProtocolsTaskV2AnalysesAnalysisIdAgentProtocolsPostRequest struct via the builder pattern


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


## CreateSecretsTaskV2AnalysesAnalysisIdAgentSecretsPost

> BaseResponseQueuedWorkflowTaskResponse CreateSecretsTaskV2AnalysesAnalysisIdAgentSecretsPost(ctx, analysisId).Execute()

Queues a secrets discovery workflow process

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
	resp, r, err := apiClient.AgentAPI.CreateSecretsTaskV2AnalysesAnalysisIdAgentSecretsPost(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CreateSecretsTaskV2AnalysesAnalysisIdAgentSecretsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecretsTaskV2AnalysesAnalysisIdAgentSecretsPost`: BaseResponseQueuedWorkflowTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CreateSecretsTaskV2AnalysesAnalysisIdAgentSecretsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretsTaskV2AnalysesAnalysisIdAgentSecretsPostRequest struct via the builder pattern


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


## GetProtocolsResultV2AnalysesAnalysisIdAgentProtocolsGet

> BaseResponseProtocolsAgentResponse GetProtocolsResultV2AnalysesAnalysisIdAgentProtocolsGet(ctx, analysisId).Execute()

Get Protocols Result



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
	resp, r, err := apiClient.AgentAPI.GetProtocolsResultV2AnalysesAnalysisIdAgentProtocolsGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetProtocolsResultV2AnalysesAnalysisIdAgentProtocolsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtocolsResultV2AnalysesAnalysisIdAgentProtocolsGet`: BaseResponseProtocolsAgentResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetProtocolsResultV2AnalysesAnalysisIdAgentProtocolsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtocolsResultV2AnalysesAnalysisIdAgentProtocolsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseProtocolsAgentResponse**](BaseResponseProtocolsAgentResponse.md)

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


## GetSecretsResultV2AnalysesAnalysisIdAgentSecretsGet

> BaseResponseSecretsAgentResponse GetSecretsResultV2AnalysesAnalysisIdAgentSecretsGet(ctx, analysisId).Execute()

Get Secrets Result



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
	resp, r, err := apiClient.AgentAPI.GetSecretsResultV2AnalysesAnalysisIdAgentSecretsGet(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetSecretsResultV2AnalysesAnalysisIdAgentSecretsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretsResultV2AnalysesAnalysisIdAgentSecretsGet`: BaseResponseSecretsAgentResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetSecretsResultV2AnalysesAnalysisIdAgentSecretsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsResultV2AnalysesAnalysisIdAgentSecretsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseSecretsAgentResponse**](BaseResponseSecretsAgentResponse.md)

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


## V3GetBinaryAgentFeedback

> FeedbackOutputBody V3GetBinaryAgentFeedback(ctx, analysisId, agent).Execute()

Get the caller's feedback on an agent's output.



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
	agent := "agent_example" // string | Which agent's output the feedback is about

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3GetBinaryAgentFeedback(context.Background(), analysisId, agent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetBinaryAgentFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetBinaryAgentFeedback`: FeedbackOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetBinaryAgentFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**agent** | **string** | Which agent&#39;s output the feedback is about | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBinaryAgentFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FeedbackOutputBody**](FeedbackOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetCapabilitiesOperation

> OperationMetadataCapabilitiesResult V3GetCapabilitiesOperation(ctx, analysisId).Execute()

Get a capabilities operation.



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
	resp, r, err := apiClient.AgentAPI.V3GetCapabilitiesOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetCapabilitiesOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetCapabilitiesOperation`: OperationMetadataCapabilitiesResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetCapabilitiesOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetCapabilitiesOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataCapabilitiesResult**](OperationMetadataCapabilitiesResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetCryptoExplainOperation

> OperationCryptoExplainMetadataCryptoExplainResult V3GetCryptoExplainOperation(ctx, functionId).Execute()

Get a crypto-explain operation.



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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3GetCryptoExplainOperation(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetCryptoExplainOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetCryptoExplainOperation`: OperationCryptoExplainMetadataCryptoExplainResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetCryptoExplainOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetCryptoExplainOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationCryptoExplainMetadataCryptoExplainResult**](OperationCryptoExplainMetadataCryptoExplainResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetCryptoScanOperation

> OperationCryptoScanMetadataCryptoScanResult V3GetCryptoScanOperation(ctx, analysisId).Execute()

Get a crypto-scan operation.



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
	resp, r, err := apiClient.AgentAPI.V3GetCryptoScanOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetCryptoScanOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetCryptoScanOperation`: OperationCryptoScanMetadataCryptoScanResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetCryptoScanOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetCryptoScanOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationCryptoScanMetadataCryptoScanResult**](OperationCryptoScanMetadataCryptoScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetExecutionExplainOperation

> OperationExecutionExplainMetadataExecutionExplainResult V3GetExecutionExplainOperation(ctx, functionId).Execute()

Get an execution-explain operation.



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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3GetExecutionExplainOperation(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetExecutionExplainOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetExecutionExplainOperation`: OperationExecutionExplainMetadataExecutionExplainResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetExecutionExplainOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetExecutionExplainOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationExecutionExplainMetadataExecutionExplainResult**](OperationExecutionExplainMetadataExecutionExplainResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetExecutionScanOperation

> OperationExecutionScanMetadataExecutionScanResult V3GetExecutionScanOperation(ctx, analysisId).Execute()

Get an execution-scan operation.



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
	resp, r, err := apiClient.AgentAPI.V3GetExecutionScanOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetExecutionScanOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetExecutionScanOperation`: OperationExecutionScanMetadataExecutionScanResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetExecutionScanOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetExecutionScanOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationExecutionScanMetadataExecutionScanResult**](OperationExecutionScanMetadataExecutionScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetFilesystemAnalyseOperation

> OperationFilesystemAnalyseMetadataFilesystemAnalyseResult V3GetFilesystemAnalyseOperation(ctx, functionId).Execute()

Get a filesystem-analyse operation.



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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3GetFilesystemAnalyseOperation(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetFilesystemAnalyseOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetFilesystemAnalyseOperation`: OperationFilesystemAnalyseMetadataFilesystemAnalyseResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetFilesystemAnalyseOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetFilesystemAnalyseOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationFilesystemAnalyseMetadataFilesystemAnalyseResult**](OperationFilesystemAnalyseMetadataFilesystemAnalyseResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetFilesystemScanOperation

> OperationFilesystemScanMetadataFilesystemScanResult V3GetFilesystemScanOperation(ctx, analysisId).Execute()

Get a filesystem-scan operation.



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
	resp, r, err := apiClient.AgentAPI.V3GetFilesystemScanOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetFilesystemScanOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetFilesystemScanOperation`: OperationFilesystemScanMetadataFilesystemScanResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetFilesystemScanOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetFilesystemScanOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationFilesystemScanMetadataFilesystemScanResult**](OperationFilesystemScanMetadataFilesystemScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetNetworkingExplainOperation

> OperationNetworkingExplainMetadataNetworkingExplainResult V3GetNetworkingExplainOperation(ctx, functionId).Execute()

Get a networking-explain operation.



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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3GetNetworkingExplainOperation(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetNetworkingExplainOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetNetworkingExplainOperation`: OperationNetworkingExplainMetadataNetworkingExplainResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetNetworkingExplainOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetNetworkingExplainOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationNetworkingExplainMetadataNetworkingExplainResult**](OperationNetworkingExplainMetadataNetworkingExplainResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetNetworkingScanOperation

> OperationNetworkingScanMetadataNetworkingScanResult V3GetNetworkingScanOperation(ctx, analysisId).Execute()

Get a networking-scan operation.



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
	resp, r, err := apiClient.AgentAPI.V3GetNetworkingScanOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetNetworkingScanOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetNetworkingScanOperation`: OperationNetworkingScanMetadataNetworkingScanResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetNetworkingScanOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetNetworkingScanOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationNetworkingScanMetadataNetworkingScanResult**](OperationNetworkingScanMetadataNetworkingScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetProtocolsOperation

> OperationMetadataReportResult V3GetProtocolsOperation(ctx, analysisId).Execute()

Get a protocols operation.



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
	resp, r, err := apiClient.AgentAPI.V3GetProtocolsOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetProtocolsOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetProtocolsOperation`: OperationMetadataReportResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetProtocolsOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetProtocolsOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataReportResult**](OperationMetadataReportResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetRemediationOperation

> OperationMetadataRemediationResult V3GetRemediationOperation(ctx, analysisId).Execute()

Get a remediation operation.



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
	resp, r, err := apiClient.AgentAPI.V3GetRemediationOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetRemediationOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetRemediationOperation`: OperationMetadataRemediationResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetRemediationOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetRemediationOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataRemediationResult**](OperationMetadataRemediationResult.md)

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


## V3GetReportAnalysisOperation

> OperationMetadataThreatReportResult V3GetReportAnalysisOperation(ctx, analysisId).Execute()

Get a report-analysis operation.



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
	resp, r, err := apiClient.AgentAPI.V3GetReportAnalysisOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetReportAnalysisOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetReportAnalysisOperation`: OperationMetadataThreatReportResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetReportAnalysisOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetReportAnalysisOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataThreatReportResult**](OperationMetadataThreatReportResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetSecretsOperation

> OperationMetadataReportResult V3GetSecretsOperation(ctx, analysisId).Execute()

Get a secrets operation.



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
	resp, r, err := apiClient.AgentAPI.V3GetSecretsOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetSecretsOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetSecretsOperation`: OperationMetadataReportResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetSecretsOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetSecretsOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataReportResult**](OperationMetadataReportResult.md)

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


## V3GetTriageOperation

> OperationMetadataTriageResult V3GetTriageOperation(ctx, analysisId).Execute()

Get a triage operation.



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
	resp, r, err := apiClient.AgentAPI.V3GetTriageOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3GetTriageOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetTriageOperation`: OperationMetadataTriageResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3GetTriageOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetTriageOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataTriageResult**](OperationMetadataTriageResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunCapabilities

> OperationMetadataCapabilitiesResult V3RunCapabilities(ctx, analysisId).Execute()

Run the capabilities agent.



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
	resp, r, err := apiClient.AgentAPI.V3RunCapabilities(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunCapabilities`: OperationMetadataCapabilitiesResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunCapabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataCapabilitiesResult**](OperationMetadataCapabilitiesResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunCryptoExplain

> OperationCryptoExplainMetadataCryptoExplainResult V3RunCryptoExplain(ctx, functionId).Execute()

Run the crypto-explain agent.



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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3RunCryptoExplain(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunCryptoExplain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunCryptoExplain`: OperationCryptoExplainMetadataCryptoExplainResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunCryptoExplain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunCryptoExplainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationCryptoExplainMetadataCryptoExplainResult**](OperationCryptoExplainMetadataCryptoExplainResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunCryptoScan

> OperationCryptoScanMetadataCryptoScanResult V3RunCryptoScan(ctx, analysisId).TriggerCryptoScanInputBody(triggerCryptoScanInputBody).Execute()

Run the crypto-scan agent.



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
	triggerCryptoScanInputBody := *revengai.NewTriggerCryptoScanInputBody() // TriggerCryptoScanInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3RunCryptoScan(context.Background(), analysisId).TriggerCryptoScanInputBody(triggerCryptoScanInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunCryptoScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunCryptoScan`: OperationCryptoScanMetadataCryptoScanResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunCryptoScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunCryptoScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerCryptoScanInputBody** | [**TriggerCryptoScanInputBody**](TriggerCryptoScanInputBody.md) |  | 

### Return type

[**OperationCryptoScanMetadataCryptoScanResult**](OperationCryptoScanMetadataCryptoScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunExecutionExplain

> OperationExecutionExplainMetadataExecutionExplainResult V3RunExecutionExplain(ctx, functionId).TriggerExecutionExplainInputBody(triggerExecutionExplainInputBody).Execute()

Run the execution-explain agent.



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
	functionId := int64(789) // int64 | Function ID
	triggerExecutionExplainInputBody := *revengai.NewTriggerExecutionExplainInputBody("Category_example") // TriggerExecutionExplainInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3RunExecutionExplain(context.Background(), functionId).TriggerExecutionExplainInputBody(triggerExecutionExplainInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunExecutionExplain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunExecutionExplain`: OperationExecutionExplainMetadataExecutionExplainResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunExecutionExplain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunExecutionExplainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerExecutionExplainInputBody** | [**TriggerExecutionExplainInputBody**](TriggerExecutionExplainInputBody.md) |  | 

### Return type

[**OperationExecutionExplainMetadataExecutionExplainResult**](OperationExecutionExplainMetadataExecutionExplainResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunExecutionScan

> OperationExecutionScanMetadataExecutionScanResult V3RunExecutionScan(ctx, analysisId).TriggerExecutionScanInputBody(triggerExecutionScanInputBody).Execute()

Run the execution-scan agent.



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
	triggerExecutionScanInputBody := *revengai.NewTriggerExecutionScanInputBody() // TriggerExecutionScanInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3RunExecutionScan(context.Background(), analysisId).TriggerExecutionScanInputBody(triggerExecutionScanInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunExecutionScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunExecutionScan`: OperationExecutionScanMetadataExecutionScanResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunExecutionScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunExecutionScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerExecutionScanInputBody** | [**TriggerExecutionScanInputBody**](TriggerExecutionScanInputBody.md) |  | 

### Return type

[**OperationExecutionScanMetadataExecutionScanResult**](OperationExecutionScanMetadataExecutionScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunFilesystemAnalyse

> OperationFilesystemAnalyseMetadataFilesystemAnalyseResult V3RunFilesystemAnalyse(ctx, functionId).TriggerFilesystemAnalyseInputBody(triggerFilesystemAnalyseInputBody).Execute()

Run the filesystem-analyse agent.



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
	functionId := int64(789) // int64 | Function ID
	triggerFilesystemAnalyseInputBody := *revengai.NewTriggerFilesystemAnalyseInputBody("Category_example") // TriggerFilesystemAnalyseInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3RunFilesystemAnalyse(context.Background(), functionId).TriggerFilesystemAnalyseInputBody(triggerFilesystemAnalyseInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunFilesystemAnalyse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunFilesystemAnalyse`: OperationFilesystemAnalyseMetadataFilesystemAnalyseResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunFilesystemAnalyse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunFilesystemAnalyseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerFilesystemAnalyseInputBody** | [**TriggerFilesystemAnalyseInputBody**](TriggerFilesystemAnalyseInputBody.md) |  | 

### Return type

[**OperationFilesystemAnalyseMetadataFilesystemAnalyseResult**](OperationFilesystemAnalyseMetadataFilesystemAnalyseResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunFilesystemScan

> OperationFilesystemScanMetadataFilesystemScanResult V3RunFilesystemScan(ctx, analysisId).TriggerFilesystemScanInputBody(triggerFilesystemScanInputBody).Execute()

Run the filesystem-scan agent.



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
	triggerFilesystemScanInputBody := *revengai.NewTriggerFilesystemScanInputBody() // TriggerFilesystemScanInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3RunFilesystemScan(context.Background(), analysisId).TriggerFilesystemScanInputBody(triggerFilesystemScanInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunFilesystemScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunFilesystemScan`: OperationFilesystemScanMetadataFilesystemScanResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunFilesystemScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunFilesystemScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerFilesystemScanInputBody** | [**TriggerFilesystemScanInputBody**](TriggerFilesystemScanInputBody.md) |  | 

### Return type

[**OperationFilesystemScanMetadataFilesystemScanResult**](OperationFilesystemScanMetadataFilesystemScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunNetworkingExplain

> OperationNetworkingExplainMetadataNetworkingExplainResult V3RunNetworkingExplain(ctx, functionId).TriggerNetworkingExplainInputBody(triggerNetworkingExplainInputBody).Execute()

Run the networking-explain agent.



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
	functionId := int64(789) // int64 | Function ID
	triggerNetworkingExplainInputBody := *revengai.NewTriggerNetworkingExplainInputBody("Category_example") // TriggerNetworkingExplainInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3RunNetworkingExplain(context.Background(), functionId).TriggerNetworkingExplainInputBody(triggerNetworkingExplainInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunNetworkingExplain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunNetworkingExplain`: OperationNetworkingExplainMetadataNetworkingExplainResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunNetworkingExplain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunNetworkingExplainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerNetworkingExplainInputBody** | [**TriggerNetworkingExplainInputBody**](TriggerNetworkingExplainInputBody.md) |  | 

### Return type

[**OperationNetworkingExplainMetadataNetworkingExplainResult**](OperationNetworkingExplainMetadataNetworkingExplainResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunNetworkingScan

> OperationNetworkingScanMetadataNetworkingScanResult V3RunNetworkingScan(ctx, analysisId).TriggerNetworkingScanInputBody(triggerNetworkingScanInputBody).Execute()

Run the networking-scan agent.



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
	triggerNetworkingScanInputBody := *revengai.NewTriggerNetworkingScanInputBody() // TriggerNetworkingScanInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.V3RunNetworkingScan(context.Background(), analysisId).TriggerNetworkingScanInputBody(triggerNetworkingScanInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunNetworkingScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunNetworkingScan`: OperationNetworkingScanMetadataNetworkingScanResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunNetworkingScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunNetworkingScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerNetworkingScanInputBody** | [**TriggerNetworkingScanInputBody**](TriggerNetworkingScanInputBody.md) |  | 

### Return type

[**OperationNetworkingScanMetadataNetworkingScanResult**](OperationNetworkingScanMetadataNetworkingScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunProtocols

> OperationMetadataReportResult V3RunProtocols(ctx, analysisId).Execute()

Run the protocols agent.



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
	resp, r, err := apiClient.AgentAPI.V3RunProtocols(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunProtocols``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunProtocols`: OperationMetadataReportResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunProtocols`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunProtocolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataReportResult**](OperationMetadataReportResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunRemediation

> OperationMetadataRemediationResult V3RunRemediation(ctx, analysisId).Execute()

Run the remediation agent.



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
	resp, r, err := apiClient.AgentAPI.V3RunRemediation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunRemediation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunRemediation`: OperationMetadataRemediationResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunRemediation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunRemediationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataRemediationResult**](OperationMetadataRemediationResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunReportAnalysis

> OperationMetadataThreatReportResult V3RunReportAnalysis(ctx, analysisId).Execute()

Run the report-analysis agent.



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
	resp, r, err := apiClient.AgentAPI.V3RunReportAnalysis(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunReportAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunReportAnalysis`: OperationMetadataThreatReportResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunReportAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunReportAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataThreatReportResult**](OperationMetadataThreatReportResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunSecrets

> OperationMetadataReportResult V3RunSecrets(ctx, analysisId).Execute()

Run the secrets agent.



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
	resp, r, err := apiClient.AgentAPI.V3RunSecrets(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunSecrets`: OperationMetadataReportResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataReportResult**](OperationMetadataReportResult.md)

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


## V3RunTriage

> OperationMetadataTriageResult V3RunTriage(ctx, analysisId).Execute()

Run the triage agent.



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
	resp, r, err := apiClient.AgentAPI.V3RunTriage(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3RunTriage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunTriage`: OperationMetadataTriageResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.V3RunTriage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunTriageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationMetadataTriageResult**](OperationMetadataTriageResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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


## V3UpsertBinaryAgentFeedback

> V3UpsertBinaryAgentFeedback(ctx, analysisId, agent).SubmitFeedbackInputBody(submitFeedbackInputBody).Execute()

Record feedback on an agent's output.



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
	agent := "agent_example" // string | Which agent's output the feedback is about
	submitFeedbackInputBody := *revengai.NewSubmitFeedbackInputBody("Sentiment_example") // SubmitFeedbackInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.V3UpsertBinaryAgentFeedback(context.Background(), analysisId, agent).SubmitFeedbackInputBody(submitFeedbackInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.V3UpsertBinaryAgentFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**agent** | **string** | Which agent&#39;s output the feedback is about | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3UpsertBinaryAgentFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **submitFeedbackInputBody** | [**SubmitFeedbackInputBody**](SubmitFeedbackInputBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

