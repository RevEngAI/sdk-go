# FunctionsBlockCommentsAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateBlockCommentsForBlockInFunction**](FunctionsBlockCommentsAPI.md#GenerateBlockCommentsForBlockInFunction) | **Post** /v2/functions/{function_id}/block-comments/single | Generate block comments for a specific block in a function
[**GenerateBlockCommentsForFunction**](FunctionsBlockCommentsAPI.md#GenerateBlockCommentsForFunction) | **Post** /v2/functions/{function_id}/block-comments | Generate block comments for a function
[**GenerateOverviewCommentForFunction**](FunctionsBlockCommentsAPI.md#GenerateOverviewCommentForFunction) | **Post** /v2/functions/{function_id}/block-comments/overview | Generate overview comment for a function



## GenerateBlockCommentsForBlockInFunction

> BaseResponseBlockCommentsGenerationForFunctionResponse GenerateBlockCommentsForBlockInFunction(ctx, functionId).Block(block).Execute()

Generate block comments for a specific block in a function

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
	block := *revengai.NewBlock(int32(123)) // Block | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsBlockCommentsAPI.GenerateBlockCommentsForBlockInFunction(context.Background(), functionId).Block(block).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsBlockCommentsAPI.GenerateBlockCommentsForBlockInFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateBlockCommentsForBlockInFunction`: BaseResponseBlockCommentsGenerationForFunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsBlockCommentsAPI.GenerateBlockCommentsForBlockInFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateBlockCommentsForBlockInFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **block** | [**Block**](Block.md) |  | 

### Return type

[**BaseResponseBlockCommentsGenerationForFunctionResponse**](BaseResponseBlockCommentsGenerationForFunctionResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateBlockCommentsForFunction

> BaseResponseBlockCommentsGenerationForFunctionResponse GenerateBlockCommentsForFunction(ctx, functionId).Execute()

Generate block comments for a function

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
	resp, r, err := apiClient.FunctionsBlockCommentsAPI.GenerateBlockCommentsForFunction(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsBlockCommentsAPI.GenerateBlockCommentsForFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateBlockCommentsForFunction`: BaseResponseBlockCommentsGenerationForFunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsBlockCommentsAPI.GenerateBlockCommentsForFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateBlockCommentsForFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseBlockCommentsGenerationForFunctionResponse**](BaseResponseBlockCommentsGenerationForFunctionResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateOverviewCommentForFunction

> BaseResponseBlockCommentsOverviewGenerationResponse GenerateOverviewCommentForFunction(ctx, functionId).Execute()

Generate overview comment for a function

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
	resp, r, err := apiClient.FunctionsBlockCommentsAPI.GenerateOverviewCommentForFunction(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsBlockCommentsAPI.GenerateOverviewCommentForFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateOverviewCommentForFunction`: BaseResponseBlockCommentsOverviewGenerationResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsBlockCommentsAPI.GenerateOverviewCommentForFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateOverviewCommentForFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseBlockCommentsOverviewGenerationResponse**](BaseResponseBlockCommentsOverviewGenerationResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

