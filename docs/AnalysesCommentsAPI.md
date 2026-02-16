# AnalysesCommentsAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnalysisComment**](AnalysesCommentsAPI.md#CreateAnalysisComment) | **Post** /v2/analyses/{analysis_id}/comments | Create a comment for this analysis
[**DeleteAnalysisComment**](AnalysesCommentsAPI.md#DeleteAnalysisComment) | **Delete** /v2/analyses/{analysis_id}/comments/{comment_id} | Delete a comment
[**GetAnalysisComments**](AnalysesCommentsAPI.md#GetAnalysisComments) | **Get** /v2/analyses/{analysis_id}/comments | Get comments for this analysis
[**UpdateAnalysisComment**](AnalysesCommentsAPI.md#UpdateAnalysisComment) | **Patch** /v2/analyses/{analysis_id}/comments/{comment_id} | Update a comment



## CreateAnalysisComment

> BaseResponseCommentResponse CreateAnalysisComment(ctx, analysisId).CommentBase(commentBase).Execute()

Create a comment for this analysis



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
	commentBase := *revengai.NewCommentBase("Content_example") // CommentBase | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCommentsAPI.CreateAnalysisComment(context.Background(), analysisId).CommentBase(commentBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCommentsAPI.CreateAnalysisComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnalysisComment`: BaseResponseCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCommentsAPI.CreateAnalysisComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalysisCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commentBase** | [**CommentBase**](CommentBase.md) |  | 

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


## DeleteAnalysisComment

> BaseResponseBool DeleteAnalysisComment(ctx, commentId, analysisId).Execute()

Delete a comment



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
	commentId := int32(56) // int32 | 
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCommentsAPI.DeleteAnalysisComment(context.Background(), commentId, analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCommentsAPI.DeleteAnalysisComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAnalysisComment`: BaseResponseBool
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCommentsAPI.DeleteAnalysisComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** |  | 
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnalysisCommentRequest struct via the builder pattern


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


## GetAnalysisComments

> BaseResponseListCommentResponse GetAnalysisComments(ctx, analysisId).Execute()

Get comments for this analysis



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
	resp, r, err := apiClient.AnalysesCommentsAPI.GetAnalysisComments(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCommentsAPI.GetAnalysisComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisComments`: BaseResponseListCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCommentsAPI.GetAnalysisComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisCommentsRequest struct via the builder pattern


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


## UpdateAnalysisComment

> BaseResponseCommentResponse UpdateAnalysisComment(ctx, commentId, analysisId).CommentUpdateRequest(commentUpdateRequest).Execute()

Update a comment



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
	commentId := int32(56) // int32 | 
	analysisId := int32(56) // int32 | 
	commentUpdateRequest := *revengai.NewCommentUpdateRequest("Content_example") // CommentUpdateRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCommentsAPI.UpdateAnalysisComment(context.Background(), commentId, analysisId).CommentUpdateRequest(commentUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCommentsAPI.UpdateAnalysisComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAnalysisComment`: BaseResponseCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCommentsAPI.UpdateAnalysisComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** |  | 
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnalysisCommentRequest struct via the builder pattern


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

