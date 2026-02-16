# AnalysesDynamicExecutionAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDynamicExecutionStatus**](AnalysesDynamicExecutionAPI.md#GetDynamicExecutionStatus) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/status | Get the status of a dynamic execution task
[**GetNetworkOverview**](AnalysesDynamicExecutionAPI.md#GetNetworkOverview) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/network-overview | Get the dynamic execution results for network overview
[**GetProcessDump**](AnalysesDynamicExecutionAPI.md#GetProcessDump) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/process-dumps/{dump_name} | Get the dynamic execution results for a specific process dump
[**GetProcessDumps**](AnalysesDynamicExecutionAPI.md#GetProcessDumps) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/process-dumps | Get the dynamic execution results for process dumps
[**GetProcessRegistry**](AnalysesDynamicExecutionAPI.md#GetProcessRegistry) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/process-registry | Get the dynamic execution results for process registry
[**GetProcessTree**](AnalysesDynamicExecutionAPI.md#GetProcessTree) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/process-tree | Get the dynamic execution results for process tree
[**GetTtps**](AnalysesDynamicExecutionAPI.md#GetTtps) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/ttps | Get the dynamic execution results for ttps



## GetDynamicExecutionStatus

> BaseResponseDynamicExecutionStatus GetDynamicExecutionStatus(ctx, analysisId).Execute()

Get the status of a dynamic execution task

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
	resp, r, err := apiClient.AnalysesDynamicExecutionAPI.GetDynamicExecutionStatus(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesDynamicExecutionAPI.GetDynamicExecutionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicExecutionStatus`: BaseResponseDynamicExecutionStatus
	fmt.Fprintf(os.Stdout, "Response from `AnalysesDynamicExecutionAPI.GetDynamicExecutionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicExecutionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseDynamicExecutionStatus**](BaseResponseDynamicExecutionStatus.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkOverview

> BaseResponseNetworkOverviewResponse GetNetworkOverview(ctx, analysisId).Execute()

Get the dynamic execution results for network overview

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
	resp, r, err := apiClient.AnalysesDynamicExecutionAPI.GetNetworkOverview(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesDynamicExecutionAPI.GetNetworkOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkOverview`: BaseResponseNetworkOverviewResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesDynamicExecutionAPI.GetNetworkOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseNetworkOverviewResponse**](BaseResponseNetworkOverviewResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessDump

> interface{} GetProcessDump(ctx, analysisId, dumpName).Execute()

Get the dynamic execution results for a specific process dump

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
	dumpName := "dumpName_example" // string | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesDynamicExecutionAPI.GetProcessDump(context.Background(), analysisId, dumpName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesDynamicExecutionAPI.GetProcessDump``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDump`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalysesDynamicExecutionAPI.GetProcessDump`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 
**dumpName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDumpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessDumps

> BaseResponseProcessDumps GetProcessDumps(ctx, analysisId).Execute()

Get the dynamic execution results for process dumps

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
	resp, r, err := apiClient.AnalysesDynamicExecutionAPI.GetProcessDumps(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesDynamicExecutionAPI.GetProcessDumps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDumps`: BaseResponseProcessDumps
	fmt.Fprintf(os.Stdout, "Response from `AnalysesDynamicExecutionAPI.GetProcessDumps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDumpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseProcessDumps**](BaseResponseProcessDumps.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessRegistry

> BaseResponseProcessRegistry GetProcessRegistry(ctx, analysisId).Execute()

Get the dynamic execution results for process registry

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
	resp, r, err := apiClient.AnalysesDynamicExecutionAPI.GetProcessRegistry(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesDynamicExecutionAPI.GetProcessRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessRegistry`: BaseResponseProcessRegistry
	fmt.Fprintf(os.Stdout, "Response from `AnalysesDynamicExecutionAPI.GetProcessRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseProcessRegistry**](BaseResponseProcessRegistry.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessTree

> BaseResponseProcessTree GetProcessTree(ctx, analysisId).Execute()

Get the dynamic execution results for process tree

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
	resp, r, err := apiClient.AnalysesDynamicExecutionAPI.GetProcessTree(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesDynamicExecutionAPI.GetProcessTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessTree`: BaseResponseProcessTree
	fmt.Fprintf(os.Stdout, "Response from `AnalysesDynamicExecutionAPI.GetProcessTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseProcessTree**](BaseResponseProcessTree.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTtps

> BaseResponseTTPS GetTtps(ctx, analysisId).Execute()

Get the dynamic execution results for ttps

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
	resp, r, err := apiClient.AnalysesDynamicExecutionAPI.GetTtps(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesDynamicExecutionAPI.GetTtps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTtps`: BaseResponseTTPS
	fmt.Fprintf(os.Stdout, "Response from `AnalysesDynamicExecutionAPI.GetTtps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTtpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseTTPS**](BaseResponseTTPS.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

