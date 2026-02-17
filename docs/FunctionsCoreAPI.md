# FunctionsCoreAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiUnstrip**](FunctionsCoreAPI.md#AiUnstrip) | **Post** /v2/analyses/{analysis_id}/functions/ai-unstrip | Performs matching and auto-unstrip for an analysis and its functions
[**AnalysisFunctionMatching**](FunctionsCoreAPI.md#AnalysisFunctionMatching) | **Post** /v2/analyses/{analysis_id}/functions/matches | Perform matching for the functions of an analysis
[**AutoUnstrip**](FunctionsCoreAPI.md#AutoUnstrip) | **Post** /v2/analyses/{analysis_id}/functions/auto-unstrip | Performs matching and auto-unstrip for an analysis and its functions
[**BatchFunctionMatching**](FunctionsCoreAPI.md#BatchFunctionMatching) | **Post** /v2/functions/matches | Perform function matching for an arbitrary batch of functions, binaries or collections
[**CancelAiUnstrip**](FunctionsCoreAPI.md#CancelAiUnstrip) | **Delete** /v2/analyses/{analysis_id}/functions/ai-unstrip/cancel | Cancels a running ai-unstrip
[**CancelAutoUnstrip**](FunctionsCoreAPI.md#CancelAutoUnstrip) | **Delete** /v2/analyses/{analysis_id}/functions/unstrip/cancel | Cancels a running auto-unstrip
[**GetAnalysisStrings**](FunctionsCoreAPI.md#GetAnalysisStrings) | **Get** /v2/analyses/{analysis_id}/functions/strings | Get string information found in the Analysis
[**GetAnalysisStringsStatus**](FunctionsCoreAPI.md#GetAnalysisStringsStatus) | **Get** /v2/analyses/{analysis_id}/functions/strings/status | Get string processing state for the Analysis
[**GetFunctionBlocks**](FunctionsCoreAPI.md#GetFunctionBlocks) | **Get** /v2/functions/{function_id}/blocks | Get disassembly blocks related to the function
[**GetFunctionCalleesCallers**](FunctionsCoreAPI.md#GetFunctionCalleesCallers) | **Get** /v2/functions/{function_id}/callees_callers | Get list of functions that call or are called by the specified function
[**GetFunctionCapabilities**](FunctionsCoreAPI.md#GetFunctionCapabilities) | **Get** /v2/functions/{function_id}/capabilities | Retrieve a functions capabilities
[**GetFunctionDetails**](FunctionsCoreAPI.md#GetFunctionDetails) | **Get** /v2/functions/{function_id} | Get function details
[**GetFunctionStrings**](FunctionsCoreAPI.md#GetFunctionStrings) | **Get** /v2/functions/{function_id}/strings | Get string information found in the function



## AiUnstrip

> AutoUnstripResponse AiUnstrip(ctx, analysisId).AiUnstripRequest(aiUnstripRequest).Execute()

Performs matching and auto-unstrip for an analysis and its functions



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
	aiUnstripRequest := *revengai.NewAiUnstripRequest() // AiUnstripRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.AiUnstrip(context.Background(), analysisId).AiUnstripRequest(aiUnstripRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.AiUnstrip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiUnstrip`: AutoUnstripResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.AiUnstrip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiUnstripRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aiUnstripRequest** | [**AiUnstripRequest**](AiUnstripRequest.md) |  | 

### Return type

[**AutoUnstripResponse**](AutoUnstripResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalysisFunctionMatching

> FunctionMatchingResponse AnalysisFunctionMatching(ctx, analysisId).AnalysisFunctionMatchingRequest(analysisFunctionMatchingRequest).Execute()

Perform matching for the functions of an analysis



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
	analysisFunctionMatchingRequest := *revengai.NewAnalysisFunctionMatchingRequest() // AnalysisFunctionMatchingRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.AnalysisFunctionMatching(context.Background(), analysisId).AnalysisFunctionMatchingRequest(analysisFunctionMatchingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.AnalysisFunctionMatching``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalysisFunctionMatching`: FunctionMatchingResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.AnalysisFunctionMatching`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalysisFunctionMatchingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analysisFunctionMatchingRequest** | [**AnalysisFunctionMatchingRequest**](AnalysisFunctionMatchingRequest.md) |  | 

### Return type

[**FunctionMatchingResponse**](FunctionMatchingResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoUnstrip

> AutoUnstripResponse AutoUnstrip(ctx, analysisId).AutoUnstripRequest(autoUnstripRequest).Execute()

Performs matching and auto-unstrip for an analysis and its functions



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
	autoUnstripRequest := *revengai.NewAutoUnstripRequest() // AutoUnstripRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.AutoUnstrip(context.Background(), analysisId).AutoUnstripRequest(autoUnstripRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.AutoUnstrip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoUnstrip`: AutoUnstripResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.AutoUnstrip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoUnstripRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoUnstripRequest** | [**AutoUnstripRequest**](AutoUnstripRequest.md) |  | 

### Return type

[**AutoUnstripResponse**](AutoUnstripResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchFunctionMatching

> FunctionMatchingResponse BatchFunctionMatching(ctx).FunctionMatchingRequest(functionMatchingRequest).Execute()

Perform function matching for an arbitrary batch of functions, binaries or collections



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
	functionMatchingRequest := *revengai.NewFunctionMatchingRequest(int32(123), []int64{int64(123)}) // FunctionMatchingRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.BatchFunctionMatching(context.Background()).FunctionMatchingRequest(functionMatchingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.BatchFunctionMatching``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchFunctionMatching`: FunctionMatchingResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.BatchFunctionMatching`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchFunctionMatchingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionMatchingRequest** | [**FunctionMatchingRequest**](FunctionMatchingRequest.md) |  | 

### Return type

[**FunctionMatchingResponse**](FunctionMatchingResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelAiUnstrip

> AutoUnstripResponse CancelAiUnstrip(ctx, analysisId).Execute()

Cancels a running ai-unstrip



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
	resp, r, err := apiClient.FunctionsCoreAPI.CancelAiUnstrip(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.CancelAiUnstrip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelAiUnstrip`: AutoUnstripResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.CancelAiUnstrip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelAiUnstripRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoUnstripResponse**](AutoUnstripResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelAutoUnstrip

> AutoUnstripResponse CancelAutoUnstrip(ctx, analysisId).Execute()

Cancels a running auto-unstrip



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
	resp, r, err := apiClient.FunctionsCoreAPI.CancelAutoUnstrip(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.CancelAutoUnstrip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelAutoUnstrip`: AutoUnstripResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.CancelAutoUnstrip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelAutoUnstripRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoUnstripResponse**](AutoUnstripResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisStrings

> BaseResponseAnalysisStringsResponse GetAnalysisStrings(ctx, analysisId).Page(page).PageSize(pageSize).Search(search).FunctionSearch(functionSearch).Execute()

Get string information found in the Analysis



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
	page := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page. (optional) (default to 100)
	search := "search_example" // string | Search is applied to string value (optional)
	functionSearch := "functionSearch_example" // string | Search is applied to function names (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetAnalysisStrings(context.Background(), analysisId).Page(page).PageSize(pageSize).Search(search).FunctionSearch(functionSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetAnalysisStrings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisStrings`: BaseResponseAnalysisStringsResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetAnalysisStrings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisStringsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | Number of items per page. | [default to 100]
 **search** | **string** | Search is applied to string value | 
 **functionSearch** | **string** | Search is applied to function names | 

### Return type

[**BaseResponseAnalysisStringsResponse**](BaseResponseAnalysisStringsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisStringsStatus

> BaseResponseAnalysisStringsStatusResponse GetAnalysisStringsStatus(ctx, analysisId).Execute()

Get string processing state for the Analysis



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
	resp, r, err := apiClient.FunctionsCoreAPI.GetAnalysisStringsStatus(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetAnalysisStringsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisStringsStatus`: BaseResponseAnalysisStringsStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetAnalysisStringsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisStringsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseAnalysisStringsStatusResponse**](BaseResponseAnalysisStringsStatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionBlocks

> BaseResponseFunctionBlocksResponse GetFunctionBlocks(ctx, functionId).Execute()

Get disassembly blocks related to the function



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
	functionId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionBlocks(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionBlocks`: BaseResponseFunctionBlocksResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseFunctionBlocksResponse**](BaseResponseFunctionBlocksResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionCalleesCallers

> BaseResponseCalleesCallerFunctionsResponse GetFunctionCalleesCallers(ctx, functionId).Execute()

Get list of functions that call or are called by the specified function

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
	functionId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionCalleesCallers(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionCalleesCallers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionCalleesCallers`: BaseResponseCalleesCallerFunctionsResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionCalleesCallers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionCalleesCallersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseCalleesCallerFunctionsResponse**](BaseResponseCalleesCallerFunctionsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionCapabilities

> BaseResponseFunctionCapabilityResponse GetFunctionCapabilities(ctx, functionId).Execute()

Retrieve a functions capabilities

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
	functionId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionCapabilities(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionCapabilities`: BaseResponseFunctionCapabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionCapabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseFunctionCapabilityResponse**](BaseResponseFunctionCapabilityResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionDetails

> BaseResponseFunctionsDetailResponse GetFunctionDetails(ctx, functionId).Execute()

Get function details

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
	functionId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionDetails(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionDetails`: BaseResponseFunctionsDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseFunctionsDetailResponse**](BaseResponseFunctionsDetailResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionStrings

> BaseResponseFunctionStringsResponse GetFunctionStrings(ctx, functionId).Page(page).PageSize(pageSize).Search(search).Execute()

Get string information found in the function



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
	functionId := int32(56) // int32 | 
	page := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page. (optional) (default to 100)
	search := "search_example" // string | Search is applied to string value (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionStrings(context.Background(), functionId).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionStrings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionStrings`: BaseResponseFunctionStringsResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionStrings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionStringsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | Number of items per page. | [default to 100]
 **search** | **string** | Search is applied to string value | 

### Return type

[**BaseResponseFunctionStringsResponse**](BaseResponseFunctionStringsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

