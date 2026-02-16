# \FunctionsRenamingHistoryAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchRenameFunction**](FunctionsRenamingHistoryAPI.md#BatchRenameFunction) | **Post** /v2/functions/rename/batch | Batch Rename Functions
[**GetFunctionNameHistory**](FunctionsRenamingHistoryAPI.md#GetFunctionNameHistory) | **Get** /v2/functions/history/{function_id} | Get Function Name History
[**RenameFunctionId**](FunctionsRenamingHistoryAPI.md#RenameFunctionId) | **Post** /v2/functions/rename/{function_id} | Rename Function
[**RevertFunctionName**](FunctionsRenamingHistoryAPI.md#RevertFunctionName) | **Post** /v2/functions/history/{function_id}/{history_id} | Revert the function name



## BatchRenameFunction

> BaseResponse BatchRenameFunction(ctx).FunctionsListRename(functionsListRename).Execute()

Batch Rename Functions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	functionsListRename := *openapiclient.NewFunctionsListRename([]openapiclient.FunctionRenameMap{*openapiclient.NewFunctionRenameMap(int64(123), "NewName_example", "NewMangledName_example")}) // FunctionsListRename | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsRenamingHistoryAPI.BatchRenameFunction(context.Background()).FunctionsListRename(functionsListRename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsRenamingHistoryAPI.BatchRenameFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchRenameFunction`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsRenamingHistoryAPI.BatchRenameFunction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchRenameFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionsListRename** | [**FunctionsListRename**](FunctionsListRename.md) |  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionNameHistory

> BaseResponseListFunctionNameHistory GetFunctionNameHistory(ctx, functionId).Execute()

Get Function Name History



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	functionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsRenamingHistoryAPI.GetFunctionNameHistory(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsRenamingHistoryAPI.GetFunctionNameHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionNameHistory`: BaseResponseListFunctionNameHistory
	fmt.Fprintf(os.Stdout, "Response from `FunctionsRenamingHistoryAPI.GetFunctionNameHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionNameHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseListFunctionNameHistory**](BaseResponseListFunctionNameHistory.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameFunctionId

> BaseResponse RenameFunctionId(ctx, functionId).FunctionRename(functionRename).Execute()

Rename Function



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	functionId := int32(56) // int32 | 
	functionRename := *openapiclient.NewFunctionRename("NewName_example", "NewMangledName_example") // FunctionRename | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsRenamingHistoryAPI.RenameFunctionId(context.Background(), functionId).FunctionRename(functionRename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsRenamingHistoryAPI.RenameFunctionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameFunctionId`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsRenamingHistoryAPI.RenameFunctionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameFunctionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **functionRename** | [**FunctionRename**](FunctionRename.md) |  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertFunctionName

> BaseResponse RevertFunctionName(ctx, functionId, historyId).Execute()

Revert the function name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	functionId := int32(56) // int32 | 
	historyId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsRenamingHistoryAPI.RevertFunctionName(context.Background(), functionId, historyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsRenamingHistoryAPI.RevertFunctionName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevertFunctionName`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsRenamingHistoryAPI.RevertFunctionName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 
**historyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertFunctionNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

