# FunctionsDataTypesAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchUpdateFunctionDataTypes**](FunctionsDataTypesAPI.md#BatchUpdateFunctionDataTypes) | **Put** /v3/analyses/{analysis_id}/functions/data-types | Batch update function data types
[**GenerateFunctionDataTypesForAnalysis**](FunctionsDataTypesAPI.md#GenerateFunctionDataTypesForAnalysis) | **Post** /v2/analyses/{analysis_id}/functions/data_types | Generate Function Data Types
[**GenerateFunctionDataTypesForFunctions**](FunctionsDataTypesAPI.md#GenerateFunctionDataTypesForFunctions) | **Post** /v2/functions/data_types | Generate Function Data Types for an arbitrary list of functions
[**GetFunctionDataTypes**](FunctionsDataTypesAPI.md#GetFunctionDataTypes) | **Get** /v2/analyses/{analysis_id}/functions/{function_id}/data_types | Get Function Data Types
[**GetFunctionDataTypes_0**](FunctionsDataTypesAPI.md#GetFunctionDataTypes_0) | **Get** /v3/analyses/{analysis_id}/functions/{function_id}/data-types | Get data types for a single function
[**ListAnalysisFunctionsDataTypes**](FunctionsDataTypesAPI.md#ListAnalysisFunctionsDataTypes) | **Get** /v3/analyses/{analysis_id}/functions/data-types | List data types for all functions in an analysis
[**ListFunctionDataTypesForAnalysis**](FunctionsDataTypesAPI.md#ListFunctionDataTypesForAnalysis) | **Get** /v2/analyses/{analysis_id}/functions/data_types | List Function Data Types
[**ListFunctionDataTypesForFunctions**](FunctionsDataTypesAPI.md#ListFunctionDataTypesForFunctions) | **Get** /v2/functions/data_types | List Function Data Types
[**ListFunctionsDataTypes**](FunctionsDataTypesAPI.md#ListFunctionsDataTypes) | **Get** /v3/functions/data-types | Get data types for many functions



## BatchUpdateFunctionDataTypes

> BatchUpdateDataTypesOutputBody BatchUpdateFunctionDataTypes(ctx, analysisId).BatchUpdateDataTypesInputBody(batchUpdateDataTypesInputBody).Execute()

Batch update function data types



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
	batchUpdateDataTypesInputBody := *revengai.NewBatchUpdateDataTypesInputBody([]revengai.BatchUpdateDataTypesItem{*revengai.NewBatchUpdateDataTypesItem(interface{}(123), int64(123), int64(123))}) // BatchUpdateDataTypesInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDataTypesAPI.BatchUpdateFunctionDataTypes(context.Background(), analysisId).BatchUpdateDataTypesInputBody(batchUpdateDataTypesInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDataTypesAPI.BatchUpdateFunctionDataTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchUpdateFunctionDataTypes`: BatchUpdateDataTypesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDataTypesAPI.BatchUpdateFunctionDataTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateFunctionDataTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchUpdateDataTypesInputBody** | [**BatchUpdateDataTypesInputBody**](BatchUpdateDataTypesInputBody.md) |  | 

### Return type

[**BatchUpdateDataTypesOutputBody**](BatchUpdateDataTypesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateFunctionDataTypesForAnalysis

> BaseResponseGenerateFunctionDataTypes GenerateFunctionDataTypesForAnalysis(ctx, analysisId).FunctionDataTypesParams(functionDataTypesParams).Execute()

Generate Function Data Types



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
	functionDataTypesParams := *revengai.NewFunctionDataTypesParams([]int64{int64(123)}) // FunctionDataTypesParams | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDataTypesAPI.GenerateFunctionDataTypesForAnalysis(context.Background(), analysisId).FunctionDataTypesParams(functionDataTypesParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDataTypesAPI.GenerateFunctionDataTypesForAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateFunctionDataTypesForAnalysis`: BaseResponseGenerateFunctionDataTypes
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDataTypesAPI.GenerateFunctionDataTypesForAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateFunctionDataTypesForAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **functionDataTypesParams** | [**FunctionDataTypesParams**](FunctionDataTypesParams.md) |  | 

### Return type

[**BaseResponseGenerateFunctionDataTypes**](BaseResponseGenerateFunctionDataTypes.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateFunctionDataTypesForFunctions

> BaseResponseGenerationStatusList GenerateFunctionDataTypesForFunctions(ctx).FunctionDataTypesParams(functionDataTypesParams).Execute()

Generate Function Data Types for an arbitrary list of functions



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
	functionDataTypesParams := *revengai.NewFunctionDataTypesParams([]int64{int64(123)}) // FunctionDataTypesParams | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDataTypesAPI.GenerateFunctionDataTypesForFunctions(context.Background()).FunctionDataTypesParams(functionDataTypesParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDataTypesAPI.GenerateFunctionDataTypesForFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateFunctionDataTypesForFunctions`: BaseResponseGenerationStatusList
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDataTypesAPI.GenerateFunctionDataTypesForFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateFunctionDataTypesForFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionDataTypesParams** | [**FunctionDataTypesParams**](FunctionDataTypesParams.md) |  | 

### Return type

[**BaseResponseGenerationStatusList**](BaseResponseGenerationStatusList.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionDataTypes

> BaseResponseFunctionDataTypes GetFunctionDataTypes(ctx, analysisId, functionId).Execute()

Get Function Data Types



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
	functionId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDataTypesAPI.GetFunctionDataTypes(context.Background(), analysisId, functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDataTypesAPI.GetFunctionDataTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionDataTypes`: BaseResponseFunctionDataTypes
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDataTypesAPI.GetFunctionDataTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionDataTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseResponseFunctionDataTypes**](BaseResponseFunctionDataTypes.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionDataTypes_0

> DataTypesEntry GetFunctionDataTypes_0(ctx, analysisId, functionId).Execute()

Get data types for a single function



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
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDataTypesAPI.GetFunctionDataTypes_0(context.Background(), analysisId, functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDataTypesAPI.GetFunctionDataTypes_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionDataTypes_0`: DataTypesEntry
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDataTypesAPI.GetFunctionDataTypes_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionDataTypes_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataTypesEntry**](DataTypesEntry.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalysisFunctionsDataTypes

> ListAnalysisFunctionsDataTypesOutputBody ListAnalysisFunctionsDataTypes(ctx, analysisId).Offset(offset).Limit(limit).Execute()

List data types for all functions in an analysis



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
	resp, r, err := apiClient.FunctionsDataTypesAPI.ListAnalysisFunctionsDataTypes(context.Background(), analysisId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDataTypesAPI.ListAnalysisFunctionsDataTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnalysisFunctionsDataTypes`: ListAnalysisFunctionsDataTypesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDataTypesAPI.ListAnalysisFunctionsDataTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAnalysisFunctionsDataTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int64** | Pagination offset. Defaults to 0. | 
 **limit** | **int64** | Page size. Defaults to 100. | 

### Return type

[**ListAnalysisFunctionsDataTypesOutputBody**](ListAnalysisFunctionsDataTypesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunctionDataTypesForAnalysis

> BaseResponseFunctionDataTypesList ListFunctionDataTypesForAnalysis(ctx, analysisId).FunctionIds(functionIds).Execute()

List Function Data Types



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
	functionIds := []*int32{int32(123)} // []*int32 |  (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDataTypesAPI.ListFunctionDataTypesForAnalysis(context.Background(), analysisId).FunctionIds(functionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDataTypesAPI.ListFunctionDataTypesForAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFunctionDataTypesForAnalysis`: BaseResponseFunctionDataTypesList
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDataTypesAPI.ListFunctionDataTypesForAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFunctionDataTypesForAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **functionIds** | **[]int32** |  | 

### Return type

[**BaseResponseFunctionDataTypesList**](BaseResponseFunctionDataTypesList.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunctionDataTypesForFunctions

> BaseResponseFunctionDataTypesList ListFunctionDataTypesForFunctions(ctx).FunctionIds(functionIds).Execute()

List Function Data Types



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
	functionIds := []*int32{int32(123)} // []*int32 |  (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDataTypesAPI.ListFunctionDataTypesForFunctions(context.Background()).FunctionIds(functionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDataTypesAPI.ListFunctionDataTypesForFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFunctionDataTypesForFunctions`: BaseResponseFunctionDataTypesList
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDataTypesAPI.ListFunctionDataTypesForFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFunctionDataTypesForFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionIds** | **[]int32** |  | 

### Return type

[**BaseResponseFunctionDataTypesList**](BaseResponseFunctionDataTypesList.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunctionsDataTypes

> ListFunctionsDataTypesOutputBody ListFunctionsDataTypes(ctx).FunctionIds(functionIds).Execute()

Get data types for many functions



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
	functionIds := []int64{int64(123)} // []int64 | Function IDs to fetch data-types for.

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDataTypesAPI.ListFunctionsDataTypes(context.Background()).FunctionIds(functionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDataTypesAPI.ListFunctionsDataTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFunctionsDataTypes`: ListFunctionsDataTypesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDataTypesAPI.ListFunctionsDataTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFunctionsDataTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionIds** | **[]int64** | Function IDs to fetch data-types for. | 

### Return type

[**ListFunctionsDataTypesOutputBody**](ListFunctionsDataTypesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

