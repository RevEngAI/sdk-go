# FunctionsDataTypesAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateFunctionDataTypesForAnalysis**](FunctionsDataTypesAPI.md#GenerateFunctionDataTypesForAnalysis) | **Post** /v2/analyses/{analysis_id}/functions/data_types | Generate Function Data Types
[**GenerateFunctionDataTypesForFunctions**](FunctionsDataTypesAPI.md#GenerateFunctionDataTypesForFunctions) | **Post** /v2/functions/data_types | Generate Function Data Types for an arbitrary list of functions
[**GetFunctionDataTypes**](FunctionsDataTypesAPI.md#GetFunctionDataTypes) | **Get** /v2/analyses/{analysis_id}/functions/{function_id}/data_types | Get Function Data Types
[**ListFunctionDataTypesForAnalysis**](FunctionsDataTypesAPI.md#ListFunctionDataTypesForAnalysis) | **Get** /v2/analyses/{analysis_id}/functions/data_types | List Function Data Types
[**ListFunctionDataTypesForFunctions**](FunctionsDataTypesAPI.md#ListFunctionDataTypesForFunctions) | **Get** /v2/functions/data_types | List Function Data Types
[**UpdateFunctionDataTypes**](FunctionsDataTypesAPI.md#UpdateFunctionDataTypes) | **Put** /v2/analyses/{analysis_id}/functions/{function_id}/data_types | Update Function Data Types



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
	revengai "github.com/RevEngAI/sdk-go"
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
	revengai "github.com/RevEngAI/sdk-go"
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
	revengai "github.com/RevEngAI/sdk-go"
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
	revengai "github.com/RevEngAI/sdk-go"
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
	revengai "github.com/RevEngAI/sdk-go"
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


## UpdateFunctionDataTypes

> BaseResponseFunctionDataTypes UpdateFunctionDataTypes(ctx, analysisId, functionId).UpdateFunctionDataTypes(updateFunctionDataTypes).Execute()

Update Function Data Types



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
	functionId := int32(56) // int32 | 
	updateFunctionDataTypes := *revengai.NewUpdateFunctionDataTypes(int32(123), *revengai.NewFunctionInfoInput([]revengai.FuncDepsInner{*revengai.NewFuncDepsInner("Name_example", int32(123), map[string]int32{"key": int32(123)}, "Type_example", int32(123))})) // UpdateFunctionDataTypes | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDataTypesAPI.UpdateFunctionDataTypes(context.Background(), analysisId, functionId).UpdateFunctionDataTypes(updateFunctionDataTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDataTypesAPI.UpdateFunctionDataTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFunctionDataTypes`: BaseResponseFunctionDataTypes
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDataTypesAPI.UpdateFunctionDataTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionDataTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFunctionDataTypes** | [**UpdateFunctionDataTypes**](UpdateFunctionDataTypes.md) |  | 

### Return type

[**BaseResponseFunctionDataTypes**](BaseResponseFunctionDataTypes.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

