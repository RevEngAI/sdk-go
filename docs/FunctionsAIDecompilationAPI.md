# FunctionsAIDecompilationAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAiDecompilation**](FunctionsAIDecompilationAPI.md#CreateAiDecompilation) | **Post** /v3/functions/{function_id}/ai-decompilation | Start AI decompilation
[**CreateAiDecompilationComment**](FunctionsAIDecompilationAPI.md#CreateAiDecompilationComment) | **Post** /v2/functions/{function_id}/ai-decompilation/comments | Create a comment for this function
[**CreateAiDecompilationTask**](FunctionsAIDecompilationAPI.md#CreateAiDecompilationTask) | **Post** /v2/functions/{function_id}/ai-decompilation | Begins AI Decompilation Process
[**DeleteAiDecompilationComment**](FunctionsAIDecompilationAPI.md#DeleteAiDecompilationComment) | **Delete** /v2/functions/{function_id}/ai-decompilation/comments/{comment_id} | Delete a comment
[**GetAiDecompilation**](FunctionsAIDecompilationAPI.md#GetAiDecompilation) | **Get** /v3/functions/{function_id}/ai-decompilation | Get AI decompilation result
[**GetAiDecompilationComments**](FunctionsAIDecompilationAPI.md#GetAiDecompilationComments) | **Get** /v2/functions/{function_id}/ai-decompilation/comments | Get comments for this function
[**GetAiDecompilationInlineComments**](FunctionsAIDecompilationAPI.md#GetAiDecompilationInlineComments) | **Get** /v3/functions/{function_id}/ai-decompilation/inline-comments | Get AI decompilation inline comments
[**GetAiDecompilationInlineCommentsStatus**](FunctionsAIDecompilationAPI.md#GetAiDecompilationInlineCommentsStatus) | **Get** /v3/functions/{function_id}/ai-decompilation/inline-comments/status | Get inline comments generation workflow status
[**GetAiDecompilationRating**](FunctionsAIDecompilationAPI.md#GetAiDecompilationRating) | **Get** /v2/functions/{function_id}/ai-decompilation/rating | Get rating for AI decompilation
[**GetAiDecompilationStatus**](FunctionsAIDecompilationAPI.md#GetAiDecompilationStatus) | **Get** /v3/functions/{function_id}/ai-decompilation/status | Get AI decompilation workflow status
[**GetAiDecompilationSummary**](FunctionsAIDecompilationAPI.md#GetAiDecompilationSummary) | **Get** /v3/functions/{function_id}/ai-decompilation/summary | Get AI decompilation summary
[**GetAiDecompilationSummaryStatus**](FunctionsAIDecompilationAPI.md#GetAiDecompilationSummaryStatus) | **Get** /v3/functions/{function_id}/ai-decompilation/summary/status | Get summary generation workflow status
[**GetAiDecompilationTaskResult**](FunctionsAIDecompilationAPI.md#GetAiDecompilationTaskResult) | **Get** /v2/functions/{function_id}/ai-decompilation | Polls AI Decompilation Process
[**GetAiDecompilationTaskStatus**](FunctionsAIDecompilationAPI.md#GetAiDecompilationTaskStatus) | **Get** /v2/functions/{function_id}/ai-decompilation/status | Check the status of a function ai decompilation
[**GetAiDecompilationTokenised**](FunctionsAIDecompilationAPI.md#GetAiDecompilationTokenised) | **Get** /v3/functions/{function_id}/ai-decompilation/tokenised | Get tokenised AI decompilation with function mapping
[**RegenerateAiDecompilationInlineComments**](FunctionsAIDecompilationAPI.md#RegenerateAiDecompilationInlineComments) | **Post** /v3/functions/{function_id}/ai-decompilation/inline-comments | Regenerate AI decompilation inline comments
[**RegenerateAiDecompilationSummary**](FunctionsAIDecompilationAPI.md#RegenerateAiDecompilationSummary) | **Post** /v3/functions/{function_id}/ai-decompilation/summary | Regenerate AI decompilation summary
[**UpdateAiDecompilationComment**](FunctionsAIDecompilationAPI.md#UpdateAiDecompilationComment) | **Patch** /v2/functions/{function_id}/ai-decompilation/comments/{comment_id} | Update a comment
[**UpsertAiDecompilationOverrides**](FunctionsAIDecompilationAPI.md#UpsertAiDecompilationOverrides) | **Patch** /v3/functions/{function_id}/ai-decompilation/overrides | Upsert variable/function name overrides
[**UpsertAiDecompilationRating**](FunctionsAIDecompilationAPI.md#UpsertAiDecompilationRating) | **Patch** /v2/functions/{function_id}/ai-decompilation/rating | Upsert rating for AI decompilation



## CreateAiDecompilation

> CreateAIDecompOutputBody CreateAiDecompilation(ctx, functionId).ContextAware(contextAware).Temperature(temperature).Execute()

Start AI decompilation



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
	contextAware := true // bool | Use context-aware decompilation (optional) (default to false)
	temperature := float64(1.2) // float64 | LLM temperature (0.0-1.0). Overrides the server default when set. Omit or set to -1 to use the server default. (optional) (default to -1)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.CreateAiDecompilation(context.Background(), functionId).ContextAware(contextAware).Temperature(temperature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.CreateAiDecompilation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAiDecompilation`: CreateAIDecompOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.CreateAiDecompilation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiDecompilationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contextAware** | **bool** | Use context-aware decompilation | [default to false]
 **temperature** | **float64** | LLM temperature (0.0-1.0). Overrides the server default when set. Omit or set to -1 to use the server default. | [default to -1]

### Return type

[**CreateAIDecompOutputBody**](CreateAIDecompOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiDecompilationComment

> BaseResponseCommentResponse CreateAiDecompilationComment(ctx, functionId).FunctionCommentCreateRequest(functionCommentCreateRequest).Execute()

Create a comment for this function



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
	functionCommentCreateRequest := *revengai.NewFunctionCommentCreateRequest("Content_example") // FunctionCommentCreateRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.CreateAiDecompilationComment(context.Background(), functionId).FunctionCommentCreateRequest(functionCommentCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.CreateAiDecompilationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAiDecompilationComment`: BaseResponseCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.CreateAiDecompilationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiDecompilationCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **functionCommentCreateRequest** | [**FunctionCommentCreateRequest**](FunctionCommentCreateRequest.md) |  | 

### Return type

[**BaseResponseCommentResponse**](BaseResponseCommentResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiDecompilationTask

> BaseResponse CreateAiDecompilationTask(ctx, functionId).Execute()

Begins AI Decompilation Process



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
	functionId := int64(789) // int64 | The ID of the function for which we are creating the decompilation task

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.CreateAiDecompilationTask(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.CreateAiDecompilationTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAiDecompilationTask`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.CreateAiDecompilationTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | The ID of the function for which we are creating the decompilation task | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiDecompilationTaskRequest struct via the builder pattern


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


## DeleteAiDecompilationComment

> BaseResponseBool DeleteAiDecompilationComment(ctx, commentId, functionId).Execute()

Delete a comment



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
	commentId := int32(56) // int32 | 
	functionId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.DeleteAiDecompilationComment(context.Background(), commentId, functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.DeleteAiDecompilationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiDecompilationComment`: BaseResponseBool
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.DeleteAiDecompilationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** |  | 
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiDecompilationCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseResponseBool**](BaseResponseBool.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilation

> DecompilationData GetAiDecompilation(ctx, functionId).Execute()

Get AI decompilation result



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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilation(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilation`: DecompilationData
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecompilationData**](DecompilationData.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationComments

> BaseResponseListCommentResponse GetAiDecompilationComments(ctx, functionId).Execute()

Get comments for this function



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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationComments(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationComments`: BaseResponseListCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseListCommentResponse**](BaseResponseListCommentResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationInlineComments

> CommentsData GetAiDecompilationInlineComments(ctx, functionId).Execute()

Get AI decompilation inline comments



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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationInlineComments(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationInlineComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationInlineComments`: CommentsData
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationInlineComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationInlineCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommentsData**](CommentsData.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationInlineCommentsStatus

> WorkflowProgress GetAiDecompilationInlineCommentsStatus(ctx, functionId).Execute()

Get inline comments generation workflow status



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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationInlineCommentsStatus(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationInlineCommentsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationInlineCommentsStatus`: WorkflowProgress
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationInlineCommentsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationInlineCommentsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowProgress**](WorkflowProgress.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationRating

> BaseResponseUnionGetAiDecompilationRatingResponseNoneType GetAiDecompilationRating(ctx, functionId).Execute()

Get rating for AI decompilation

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
	functionId := int64(789) // int64 | The ID of the function for which to get the rating

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationRating(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationRating``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationRating`: BaseResponseUnionGetAiDecompilationRatingResponseNoneType
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | The ID of the function for which to get the rating | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseUnionGetAiDecompilationRatingResponseNoneType**](BaseResponseUnionGetAiDecompilationRatingResponseNoneType.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationStatus

> WorkflowProgress GetAiDecompilationStatus(ctx, functionId).Execute()

Get AI decompilation workflow status



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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationStatus(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationStatus`: WorkflowProgress
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowProgress**](WorkflowProgress.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationSummary

> SummaryData GetAiDecompilationSummary(ctx, functionId).Execute()

Get AI decompilation summary



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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationSummary(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationSummary`: SummaryData
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SummaryData**](SummaryData.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationSummaryStatus

> WorkflowProgress GetAiDecompilationSummaryStatus(ctx, functionId).Execute()

Get summary generation workflow status



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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationSummaryStatus(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationSummaryStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationSummaryStatus`: WorkflowProgress
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationSummaryStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationSummaryStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowProgress**](WorkflowProgress.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationTaskResult

> BaseResponseGetAiDecompilationTask GetAiDecompilationTaskResult(ctx, functionId).Summarise(summarise).GenerateInlineComments(generateInlineComments).ForceRegenerate(forceRegenerate).Execute()

Polls AI Decompilation Process



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
	functionId := int64(789) // int64 | The ID of the function being decompiled
	summarise := true // bool | Generate a summary for the decompilation (optional) (default to true)
	generateInlineComments := true // bool | Generate inline comments for the decompilation (optional) (default to true)
	forceRegenerate := []revengai.RegenerateTarget{revengai.RegenerateTarget("summary")} // []RegenerateTarget | Force regeneration of summary and/or comments. (optional) (default to {})

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationTaskResult(context.Background(), functionId).Summarise(summarise).GenerateInlineComments(generateInlineComments).ForceRegenerate(forceRegenerate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationTaskResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationTaskResult`: BaseResponseGetAiDecompilationTask
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationTaskResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | The ID of the function being decompiled | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationTaskResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **summarise** | **bool** | Generate a summary for the decompilation | [default to true]
 **generateInlineComments** | **bool** | Generate inline comments for the decompilation | [default to true]
 **forceRegenerate** | [**[]RegenerateTarget**](RegenerateTarget.md) | Force regeneration of summary and/or comments. | [default to {}]

### Return type

[**BaseResponseGetAiDecompilationTask**](BaseResponseGetAiDecompilationTask.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationTaskStatus

> BaseResponseFunctionTaskResponse GetAiDecompilationTaskStatus(ctx, functionId).Execute()

Check the status of a function ai decompilation

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
	functionId := int64(789) // int64 | The ID of the function being checked

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationTaskStatus(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationTaskStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationTaskStatus`: BaseResponseFunctionTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationTaskStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | The ID of the function being checked | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseFunctionTaskResponse**](BaseResponseFunctionTaskResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationTokenised

> TokenisedData GetAiDecompilationTokenised(ctx, functionId).Execute()

Get tokenised AI decompilation with function mapping



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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationTokenised(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationTokenised``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationTokenised`: TokenisedData
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationTokenised`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationTokenisedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenisedData**](TokenisedData.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateAiDecompilationInlineComments

> RegenerateOutputBody RegenerateAiDecompilationInlineComments(ctx, functionId).Execute()

Regenerate AI decompilation inline comments



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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.RegenerateAiDecompilationInlineComments(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.RegenerateAiDecompilationInlineComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateAiDecompilationInlineComments`: RegenerateOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.RegenerateAiDecompilationInlineComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateAiDecompilationInlineCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegenerateOutputBody**](RegenerateOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateAiDecompilationSummary

> RegenerateOutputBody RegenerateAiDecompilationSummary(ctx, functionId).Execute()

Regenerate AI decompilation summary



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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.RegenerateAiDecompilationSummary(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.RegenerateAiDecompilationSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateAiDecompilationSummary`: RegenerateOutputBody
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.RegenerateAiDecompilationSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateAiDecompilationSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegenerateOutputBody**](RegenerateOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAiDecompilationComment

> BaseResponseCommentResponse UpdateAiDecompilationComment(ctx, commentId, functionId).CommentUpdateRequest(commentUpdateRequest).Execute()

Update a comment



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
	commentId := int32(56) // int32 | 
	functionId := int32(56) // int32 | 
	commentUpdateRequest := *revengai.NewCommentUpdateRequest("Content_example") // CommentUpdateRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.UpdateAiDecompilationComment(context.Background(), commentId, functionId).CommentUpdateRequest(commentUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.UpdateAiDecompilationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAiDecompilationComment`: BaseResponseCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.UpdateAiDecompilationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** |  | 
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAiDecompilationCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commentUpdateRequest** | [**CommentUpdateRequest**](CommentUpdateRequest.md) |  | 

### Return type

[**BaseResponseCommentResponse**](BaseResponseCommentResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertAiDecompilationOverrides

> UpsertOverridesData UpsertAiDecompilationOverrides(ctx, functionId).UpsertOverridesInputBody(upsertOverridesInputBody).Execute()

Upsert variable/function name overrides



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
	upsertOverridesInputBody := *revengai.NewUpsertOverridesInputBody(map[string]string{"key": "Inner_example"}) // UpsertOverridesInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.UpsertAiDecompilationOverrides(context.Background(), functionId).UpsertOverridesInputBody(upsertOverridesInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.UpsertAiDecompilationOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertAiDecompilationOverrides`: UpsertOverridesData
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.UpsertAiDecompilationOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertAiDecompilationOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsertOverridesInputBody** | [**UpsertOverridesInputBody**](UpsertOverridesInputBody.md) |  | 

### Return type

[**UpsertOverridesData**](UpsertOverridesData.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertAiDecompilationRating

> BaseResponse UpsertAiDecompilationRating(ctx, functionId).UpsertAiDecomplationRatingRequest(upsertAiDecomplationRatingRequest).Execute()

Upsert rating for AI decompilation

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
	functionId := int64(789) // int64 | The ID of the function being rated
	upsertAiDecomplationRatingRequest := *revengai.NewUpsertAiDecomplationRatingRequest(revengai.AiDecompilationRating("POSITIVE"), "Reason_example") // UpsertAiDecomplationRatingRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.UpsertAiDecompilationRating(context.Background(), functionId).UpsertAiDecomplationRatingRequest(upsertAiDecomplationRatingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.UpsertAiDecompilationRating``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertAiDecompilationRating`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.UpsertAiDecompilationRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | The ID of the function being rated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertAiDecompilationRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsertAiDecomplationRatingRequest** | [**UpsertAiDecomplationRatingRequest**](UpsertAiDecomplationRatingRequest.md) |  | 

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

