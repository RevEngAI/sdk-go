# FunctionsCoreAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFunctionCallee**](FunctionsCoreAPI.md#AddFunctionCallee) | **Post** /v3/functions/{function_id}/callees | Add a callee to a function
[**AddUserStringToFunction**](FunctionsCoreAPI.md#AddUserStringToFunction) | **Post** /v3/functions/{function_id}/user-provided-strings | Add a user-provided string to a function.
[**AiUnstrip**](FunctionsCoreAPI.md#AiUnstrip) | **Post** /v2/analyses/{analysis_id}/functions/ai-unstrip | Performs matching and auto-unstrip for an analysis and its functions
[**AnalysisFunctionMatching**](FunctionsCoreAPI.md#AnalysisFunctionMatching) | **Post** /v2/analyses/{analysis_id}/functions/matches | Perform matching for the functions of an analysis
[**AutoUnstrip**](FunctionsCoreAPI.md#AutoUnstrip) | **Post** /v2/analyses/{analysis_id}/functions/auto-unstrip | Performs matching and auto-unstrip for an analysis and its functions
[**BatchFunctionMatching**](FunctionsCoreAPI.md#BatchFunctionMatching) | **Post** /v2/functions/matches | Perform function matching for an arbitrary batch of functions, binaries or collections
[**CancelAiUnstrip**](FunctionsCoreAPI.md#CancelAiUnstrip) | **Delete** /v2/analyses/{analysis_id}/functions/ai-unstrip/cancel | Cancels a running ai-unstrip
[**CancelAutoUnstrip**](FunctionsCoreAPI.md#CancelAutoUnstrip) | **Delete** /v2/analyses/{analysis_id}/functions/unstrip/cancel | Cancels a running auto-unstrip
[**GetAnalysisStrings**](FunctionsCoreAPI.md#GetAnalysisStrings) | **Get** /v2/analyses/{analysis_id}/functions/strings | Get string information found in the Analysis
[**GetAnalysisStringsStatus**](FunctionsCoreAPI.md#GetAnalysisStringsStatus) | **Get** /v2/analyses/{analysis_id}/functions/strings/status | Get string processing state for the Analysis
[**GetFunctionBlocks**](FunctionsCoreAPI.md#GetFunctionBlocks) | **Get** /v2/functions/{function_id}/blocks | Get disassembly blocks related to the function
[**GetFunctionBlocks_0**](FunctionsCoreAPI.md#GetFunctionBlocks_0) | **Get** /v3/functions/{function_id}/blocks | Get function disassembly
[**GetFunctionCalleesCallers**](FunctionsCoreAPI.md#GetFunctionCalleesCallers) | **Get** /v2/functions/{function_id}/callees_callers | Get list of functions that call or are called by the specified function
[**GetFunctionCalleesCallersBulk**](FunctionsCoreAPI.md#GetFunctionCalleesCallersBulk) | **Get** /v2/functions/callees_callers | Get list of functions that call or are called for a list of functions
[**GetFunctionCalleesCallers_0**](FunctionsCoreAPI.md#GetFunctionCalleesCallers_0) | **Get** /v3/functions/{function_id}/callees-callers | Get callees and callers for a function
[**GetFunctionCapabilities**](FunctionsCoreAPI.md#GetFunctionCapabilities) | **Get** /v2/functions/{function_id}/capabilities | Retrieve a functions capabilities
[**GetFunctionCapabilities_0**](FunctionsCoreAPI.md#GetFunctionCapabilities_0) | **Get** /v3/functions/{function_id}/capabilities | Get capabilities for a function
[**GetFunctionDetails**](FunctionsCoreAPI.md#GetFunctionDetails) | **Get** /v2/functions/{function_id} | Get function details
[**GetFunctionDetails_0**](FunctionsCoreAPI.md#GetFunctionDetails_0) | **Get** /v3/functions/{function_id} | Get function details
[**GetFunctionIndirectCallSites**](FunctionsCoreAPI.md#GetFunctionIndirectCallSites) | **Get** /v3/functions/{function_id}/indirect-call-sites | Get indirect call sites for a function
[**GetFunctionStrings**](FunctionsCoreAPI.md#GetFunctionStrings) | **Get** /v2/functions/{function_id}/strings | Get string information found in the function
[**GetFunctionStrings_0**](FunctionsCoreAPI.md#GetFunctionStrings_0) | **Get** /v3/functions/{function_id}/strings | List strings for a function.
[**GetFunctionsCalleesCallers**](FunctionsCoreAPI.md#GetFunctionsCalleesCallers) | **Get** /v3/functions/callees-callers | Get callees and callers for many functions
[**GetFunctionsMatches**](FunctionsCoreAPI.md#GetFunctionsMatches) | **Get** /v3/functions/matches | Get function-matching results for an explicit set of functions
[**GetFunctionsMatchingStatus**](FunctionsCoreAPI.md#GetFunctionsMatchingStatus) | **Get** /v3/functions/matches/status | Get function-matching status for an explicit set of functions
[**GetImportedFunction**](FunctionsCoreAPI.md#GetImportedFunction) | **Get** /v3/analyses/{analysis_id}/imported-functions/{imported_function_id} | Get an imported function with its callers
[**ListAnalysisFunctions**](FunctionsCoreAPI.md#ListAnalysisFunctions) | **Get** /v3/analyses/{analysis_id}/functions | List functions in an analysis
[**ListImportedFunctions**](FunctionsCoreAPI.md#ListImportedFunctions) | **Get** /v3/analyses/{analysis_id}/imported-functions | List imported functions in an analysis
[**StartFunctionsMatching**](FunctionsCoreAPI.md#StartFunctionsMatching) | **Post** /v3/functions/matches | Start function matching for an explicit set of functions
[**V3CanonicalizeFunctionNames**](FunctionsCoreAPI.md#V3CanonicalizeFunctionNames) | **Post** /v3/functions/canonical-names | Canonicalize a batch of function names



## AddFunctionCallee

> map[string]interface{} AddFunctionCallee(ctx, functionId).AddCalleeInputBody(addCalleeInputBody).Execute()

Add a callee to a function



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
	functionId := int64(789) // int64 | Function ID
	addCalleeInputBody := *revengai.NewAddCalleeInputBody(int64(123), false) // AddCalleeInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.AddFunctionCallee(context.Background(), functionId).AddCalleeInputBody(addCalleeInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.AddFunctionCallee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFunctionCallee`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.AddFunctionCallee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFunctionCalleeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addCalleeInputBody** | [**AddCalleeInputBody**](AddCalleeInputBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserStringToFunction

> map[string]interface{} AddUserStringToFunction(ctx, functionId).AddUserStringToFunctionInputBody(addUserStringToFunctionInputBody).Execute()

Add a user-provided string to a function.



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
	functionId := int64(789) // int64 | Function ID
	addUserStringToFunctionInputBody := *revengai.NewAddUserStringToFunctionInputBody("String_example", int64(123)) // AddUserStringToFunctionInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.AddUserStringToFunction(context.Background(), functionId).AddUserStringToFunctionInputBody(addUserStringToFunctionInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.AddUserStringToFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserStringToFunction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.AddUserStringToFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserStringToFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addUserStringToFunctionInputBody** | [**AddUserStringToFunctionInputBody**](AddUserStringToFunctionInputBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisStrings

> BaseResponseAnalysisStringsResponse GetAnalysisStrings(ctx, analysisId).Page(page).PageSize(pageSize).Search(search).FunctionSearch(functionSearch).OrderBy(orderBy).SortOrder(sortOrder).Execute()

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
	orderBy := "orderBy_example" // string | Order by field (optional) (default to "value")
	sortOrder := "sortOrder_example" // string | Sort order for the results (optional) (default to "ASC")

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetAnalysisStrings(context.Background(), analysisId).Page(page).PageSize(pageSize).Search(search).FunctionSearch(functionSearch).OrderBy(orderBy).SortOrder(sortOrder).Execute()
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
 **orderBy** | **string** | Order by field | [default to &quot;value&quot;]
 **sortOrder** | **string** | Sort order for the results | [default to &quot;ASC&quot;]

### Return type

[**BaseResponseAnalysisStringsResponse**](BaseResponseAnalysisStringsResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionBlocks_0

> DisassemblyOutputBody GetFunctionBlocks_0(ctx, functionId).Execute()

Get function disassembly



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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionBlocks_0(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionBlocks_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionBlocks_0`: DisassemblyOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionBlocks_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionBlocks_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisassemblyOutputBody**](DisassemblyOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionCalleesCallersBulk

> BaseResponseListCalleesCallerFunctionsResponse GetFunctionCalleesCallersBulk(ctx).FunctionIds(functionIds).Execute()

Get list of functions that call or are called for a list of functions

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
	functionIds := []*int32{int32(123)} // []*int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionCalleesCallersBulk(context.Background()).FunctionIds(functionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionCalleesCallersBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionCalleesCallersBulk`: BaseResponseListCalleesCallerFunctionsResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionCalleesCallersBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionCalleesCallersBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionIds** | **[]int32** |  | 

### Return type

[**BaseResponseListCalleesCallerFunctionsResponse**](BaseResponseListCalleesCallerFunctionsResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionCalleesCallers_0

> CallEdgesOutputBody GetFunctionCalleesCallers_0(ctx, functionId).Execute()

Get callees and callers for a function



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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionCalleesCallers_0(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionCalleesCallers_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionCalleesCallers_0`: CallEdgesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionCalleesCallers_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionCalleesCallers_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CallEdgesOutputBody**](CallEdgesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionCapabilities_0

> CapabilitiesOutputBody GetFunctionCapabilities_0(ctx, functionId).Execute()

Get capabilities for a function



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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionCapabilities_0(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionCapabilities_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionCapabilities_0`: CapabilitiesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionCapabilities_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionCapabilities_3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CapabilitiesOutputBody**](CapabilitiesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionDetails_0

> FunctionDetailsOutputBody GetFunctionDetails_0(ctx, functionId).Execute()

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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionDetails_0(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionDetails_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionDetails_0`: FunctionDetailsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionDetails_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionDetails_4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionDetailsOutputBody**](FunctionDetailsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionIndirectCallSites

> IndirectCallSitesOutputBody GetFunctionIndirectCallSites(ctx, functionId).Execute()

Get indirect call sites for a function



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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionIndirectCallSites(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionIndirectCallSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionIndirectCallSites`: IndirectCallSitesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionIndirectCallSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionIndirectCallSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndirectCallSitesOutputBody**](IndirectCallSitesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionStrings_0

> ListFunctionStringsOutputBody GetFunctionStrings_0(ctx, functionId).Page(page).PageSize(pageSize).Search(search).Execute()

List strings for a function.



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
	functionId := int64(789) // int64 | Function ID
	page := int64(789) // int64 | Page number (1-indexed). (optional) (default to 1)
	pageSize := int64(789) // int64 | Number of results per page. (optional) (default to 100)
	search := "search_example" // string | Filter by string value (case-insensitive substring match). (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionStrings_0(context.Background(), functionId).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionStrings_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionStrings_0`: ListFunctionStringsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionStrings_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionStrings_5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Page number (1-indexed). | [default to 1]
 **pageSize** | **int64** | Number of results per page. | [default to 100]
 **search** | **string** | Filter by string value (case-insensitive substring match). | 

### Return type

[**ListFunctionStringsOutputBody**](ListFunctionStringsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsCalleesCallers

> CallEdgesOutputBody GetFunctionsCalleesCallers(ctx).FunctionIds(functionIds).Execute()

Get callees and callers for many functions



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
	functionIds := []int64{int64(123)} // []int64 | Function IDs to fetch edges for.

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionsCalleesCallers(context.Background()).FunctionIds(functionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionsCalleesCallers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsCalleesCallers`: CallEdgesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionsCalleesCallers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsCalleesCallersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionIds** | **[]int64** | Function IDs to fetch edges for. | 

### Return type

[**CallEdgesOutputBody**](CallEdgesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsMatches

> GetMatchesOutputBody GetFunctionsMatches(ctx).MatchId(matchId).FunctionIds(functionIds).Execute()

Get function-matching results for an explicit set of functions



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
	matchId := "matchId_example" // string | Opaque token from a start-matching response. When supplied, returns that specific run instead of the latest. (optional)
	functionIds := []int64{int64(123)} // []int64 | Source function IDs whose matches to fetch. Required unless match_id is supplied. (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionsMatches(context.Background()).MatchId(matchId).FunctionIds(functionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionsMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsMatches`: GetMatchesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionsMatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **matchId** | **string** | Opaque token from a start-matching response. When supplied, returns that specific run instead of the latest. | 
 **functionIds** | **[]int64** | Source function IDs whose matches to fetch. Required unless match_id is supplied. | 

### Return type

[**GetMatchesOutputBody**](GetMatchesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsMatchingStatus

> GetMatchesStatusOutputBody GetFunctionsMatchingStatus(ctx).MatchId(matchId).FunctionIds(functionIds).Execute()

Get function-matching status for an explicit set of functions



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
	matchId := "matchId_example" // string | Opaque token from a start-matching response. When supplied, returns that specific run instead of the latest. (optional)
	functionIds := []int64{int64(123)} // []int64 | Source function IDs whose matches to fetch. Required unless match_id is supplied. (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetFunctionsMatchingStatus(context.Background()).MatchId(matchId).FunctionIds(functionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetFunctionsMatchingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsMatchingStatus`: GetMatchesStatusOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetFunctionsMatchingStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsMatchingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **matchId** | **string** | Opaque token from a start-matching response. When supplied, returns that specific run instead of the latest. | 
 **functionIds** | **[]int64** | Source function IDs whose matches to fetch. Required unless match_id is supplied. | 

### Return type

[**GetMatchesStatusOutputBody**](GetMatchesStatusOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportedFunction

> ImportedFunctionDetailOutputBody GetImportedFunction(ctx, analysisId, importedFunctionId).Execute()

Get an imported function with its callers



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
	importedFunctionId := int64(789) // int64 | Imported function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.GetImportedFunction(context.Background(), analysisId, importedFunctionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.GetImportedFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportedFunction`: ImportedFunctionDetailOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.GetImportedFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**importedFunctionId** | **int64** | Imported function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportedFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ImportedFunctionDetailOutputBody**](ImportedFunctionDetailOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalysisFunctions

> ListAnalysisFunctionsOutputBody ListAnalysisFunctions(ctx, analysisId).Offset(offset).Limit(limit).Execute()

List functions in an analysis



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
	offset := int64(789) // int64 | Pagination offset. Defaults to 0. (optional)
	limit := int64(789) // int64 | Page size. Defaults to 100. (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.ListAnalysisFunctions(context.Background(), analysisId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.ListAnalysisFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnalysisFunctions`: ListAnalysisFunctionsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.ListAnalysisFunctions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAnalysisFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int64** | Pagination offset. Defaults to 0. | 
 **limit** | **int64** | Page size. Defaults to 100. | 

### Return type

[**ListAnalysisFunctionsOutputBody**](ListAnalysisFunctionsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImportedFunctions

> ListImportedFunctionsOutputBody ListImportedFunctions(ctx, analysisId).Offset(offset).Limit(limit).Execute()

List imported functions in an analysis



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
	offset := int64(789) // int64 | Pagination offset. Defaults to 0. (optional)
	limit := int64(789) // int64 | Page size. Defaults to 100. (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.ListImportedFunctions(context.Background(), analysisId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.ListImportedFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImportedFunctions`: ListImportedFunctionsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.ListImportedFunctions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImportedFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int64** | Pagination offset. Defaults to 0. | 
 **limit** | **int64** | Page size. Defaults to 100. | 

### Return type

[**ListImportedFunctionsOutputBody**](ListImportedFunctionsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartFunctionsMatching

> StartMatchingOutputBody StartFunctionsMatching(ctx).StartMatchingForFunctionsInputBody(startMatchingForFunctionsInputBody).Execute()

Start function matching for an explicit set of functions



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
	startMatchingForFunctionsInputBody := *revengai.NewStartMatchingForFunctionsInputBody([]int64{int64(123)}) // StartMatchingForFunctionsInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.StartFunctionsMatching(context.Background()).StartMatchingForFunctionsInputBody(startMatchingForFunctionsInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.StartFunctionsMatching``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartFunctionsMatching`: StartMatchingOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.StartFunctionsMatching`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartFunctionsMatchingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startMatchingForFunctionsInputBody** | [**StartMatchingForFunctionsInputBody**](StartMatchingForFunctionsInputBody.md) |  | 

### Return type

[**StartMatchingOutputBody**](StartMatchingOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3CanonicalizeFunctionNames

> CanonicalizeNamesOutputBody V3CanonicalizeFunctionNames(ctx).CanonicalizeNamesInputBody(canonicalizeNamesInputBody).Execute()

Canonicalize a batch of function names



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
	canonicalizeNamesInputBody := *revengai.NewCanonicalizeNamesInputBody([]string{"Names_example"}) // CanonicalizeNamesInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsCoreAPI.V3CanonicalizeFunctionNames(context.Background()).CanonicalizeNamesInputBody(canonicalizeNamesInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsCoreAPI.V3CanonicalizeFunctionNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3CanonicalizeFunctionNames`: CanonicalizeNamesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsCoreAPI.V3CanonicalizeFunctionNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3CanonicalizeFunctionNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canonicalizeNamesInputBody** | [**CanonicalizeNamesInputBody**](CanonicalizeNamesInputBody.md) |  | 

### Return type

[**CanonicalizeNamesOutputBody**](CanonicalizeNamesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

