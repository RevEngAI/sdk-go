# FunctionsRenamingHistoryAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchRenameFunction**](FunctionsRenamingHistoryAPI.md#BatchRenameFunction) | **Post** /v2/functions/rename/batch | Batch Rename Functions
[**BatchRenameFunctions**](FunctionsRenamingHistoryAPI.md#BatchRenameFunctions) | **Post** /v3/functions/rename | Batch rename functions
[**GetFunctionHistory**](FunctionsRenamingHistoryAPI.md#GetFunctionHistory) | **Get** /v3/functions/{function_id}/history | Get function name history
[**GetFunctionNameHistory**](FunctionsRenamingHistoryAPI.md#GetFunctionNameHistory) | **Get** /v2/functions/history/{function_id} | Get Function Name History
[**RenameFunction**](FunctionsRenamingHistoryAPI.md#RenameFunction) | **Post** /v3/functions/{function_id}/rename | Rename a function
[**RenameFunctionId**](FunctionsRenamingHistoryAPI.md#RenameFunctionId) | **Post** /v2/functions/rename/{function_id} | Rename Function
[**RevertFunctionName**](FunctionsRenamingHistoryAPI.md#RevertFunctionName) | **Post** /v2/functions/history/{function_id}/{history_id} | Revert the function name
[**RevertFunctionName_0**](FunctionsRenamingHistoryAPI.md#RevertFunctionName_0) | **Post** /v3/functions/{function_id}/history/{history_id}/revert | Revert function name



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
	revengai "github.com/RevEngAI/sdk-go/v3"
)

func main() {
	functionsListRename := *revengai.NewFunctionsListRename([]revengai.FunctionRenameMap{*revengai.NewFunctionRenameMap(int64(123), "NewName_example", "NewMangledName_example")}) // FunctionsListRename | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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


## BatchRenameFunctions

> BatchRenameOutputBody BatchRenameFunctions(ctx).BatchRenameInputBody(batchRenameInputBody).Execute()

Batch rename functions



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
	batchRenameInputBody := *revengai.NewBatchRenameInputBody([]revengai.BatchRenameItem{*revengai.NewBatchRenameItem(int64(123), "NewName_example")}) // BatchRenameInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsRenamingHistoryAPI.BatchRenameFunctions(context.Background()).BatchRenameInputBody(batchRenameInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsRenamingHistoryAPI.BatchRenameFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchRenameFunctions`: BatchRenameOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsRenamingHistoryAPI.BatchRenameFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchRenameFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRenameInputBody** | [**BatchRenameInputBody**](BatchRenameInputBody.md) |  | 

### Return type

[**BatchRenameOutputBody**](BatchRenameOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionHistory

> []HistoryEntry GetFunctionHistory(ctx, functionId).Execute()

Get function name history



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
	resp, r, err := apiClient.FunctionsRenamingHistoryAPI.GetFunctionHistory(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsRenamingHistoryAPI.GetFunctionHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionHistory`: []HistoryEntry
	fmt.Fprintf(os.Stdout, "Response from `FunctionsRenamingHistoryAPI.GetFunctionHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HistoryEntry**](HistoryEntry.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
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
	revengai "github.com/RevEngAI/sdk-go/v3"
)

func main() {
	functionId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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


## RenameFunction

> RenameOutputBody RenameFunction(ctx, functionId).RenameInputBody(renameInputBody).Execute()

Rename a function



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
	renameInputBody := *revengai.NewRenameInputBody("NewName_example") // RenameInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsRenamingHistoryAPI.RenameFunction(context.Background(), functionId).RenameInputBody(renameInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsRenamingHistoryAPI.RenameFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameFunction`: RenameOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsRenamingHistoryAPI.RenameFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renameInputBody** | [**RenameInputBody**](RenameInputBody.md) |  | 

### Return type

[**RenameOutputBody**](RenameOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
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
	revengai "github.com/RevEngAI/sdk-go/v3"
)

func main() {
	functionId := int32(56) // int32 | 
	functionRename := *revengai.NewFunctionRename("NewName_example", "NewMangledName_example") // FunctionRename | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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
	revengai "github.com/RevEngAI/sdk-go/v3"
)

func main() {
	functionId := int32(56) // int32 | 
	historyId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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


## RevertFunctionName_0

> RevertOutputBody RevertFunctionName_0(ctx, functionId, historyId).Execute()

Revert function name



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
	historyId := int64(789) // int64 | History ID to revert to

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsRenamingHistoryAPI.RevertFunctionName_0(context.Background(), functionId, historyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsRenamingHistoryAPI.RevertFunctionName_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevertFunctionName_0`: RevertOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsRenamingHistoryAPI.RevertFunctionName_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 
**historyId** | **int64** | History ID to revert to | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertFunctionName_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RevertOutputBody**](RevertOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

