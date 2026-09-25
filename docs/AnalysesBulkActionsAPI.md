# AnalysesBulkActionsAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkAddAnalysisTags**](AnalysesBulkActionsAPI.md#BulkAddAnalysisTags) | **Patch** /v2/analyses/tags/add | Bulk Add Analysis Tags
[**BulkDeleteAnalyses**](AnalysesBulkActionsAPI.md#BulkDeleteAnalyses) | **Patch** /v2/analyses/delete | Bulk Delete Analyses
[**V3BatchAddAnalysisTags**](AnalysesBulkActionsAPI.md#V3BatchAddAnalysisTags) | **Post** /v3/analyses:batchAddTags | Add tags to multiple analyses.
[**V3BatchDeleteAnalyses**](AnalysesBulkActionsAPI.md#V3BatchDeleteAnalyses) | **Post** /v3/analyses:batchDelete | Delete multiple analyses.



## BulkAddAnalysisTags

> BaseResponseAnalysisBulkAddTagsResponse BulkAddAnalysisTags(ctx).AnalysisBulkAddTagsRequest(analysisBulkAddTagsRequest).Execute()

Bulk Add Analysis Tags



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisBulkAddTagsRequest := *revengai.NewAnalysisBulkAddTagsRequest([]string{"Tags_example"}, []int32{int32(123)}) // AnalysisBulkAddTagsRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesBulkActionsAPI.BulkAddAnalysisTags(context.Background()).AnalysisBulkAddTagsRequest(analysisBulkAddTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesBulkActionsAPI.BulkAddAnalysisTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkAddAnalysisTags`: BaseResponseAnalysisBulkAddTagsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesBulkActionsAPI.BulkAddAnalysisTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAddAnalysisTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analysisBulkAddTagsRequest** | [**AnalysisBulkAddTagsRequest**](AnalysisBulkAddTagsRequest.md) |  | 

### Return type

[**BaseResponseAnalysisBulkAddTagsResponse**](BaseResponseAnalysisBulkAddTagsResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteAnalyses

> BaseResponseDict BulkDeleteAnalyses(ctx).BulkDeleteAnalysesRequest(bulkDeleteAnalysesRequest).Execute()

Bulk Delete Analyses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	bulkDeleteAnalysesRequest := *revengai.NewBulkDeleteAnalysesRequest([]int32{int32(123)}) // BulkDeleteAnalysesRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesBulkActionsAPI.BulkDeleteAnalyses(context.Background()).BulkDeleteAnalysesRequest(bulkDeleteAnalysesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesBulkActionsAPI.BulkDeleteAnalyses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkDeleteAnalyses`: BaseResponseDict
	fmt.Fprintf(os.Stdout, "Response from `AnalysesBulkActionsAPI.BulkDeleteAnalyses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteAnalysesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkDeleteAnalysesRequest** | [**BulkDeleteAnalysesRequest**](BulkDeleteAnalysesRequest.md) |  | 

### Return type

[**BaseResponseDict**](BaseResponseDict.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BatchAddAnalysisTags

> BulkAddTagsOutputBody V3BatchAddAnalysisTags(ctx).BulkAddTagsInputBody(bulkAddTagsInputBody).Execute()

Add tags to multiple analyses.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	bulkAddTagsInputBody := *revengai.NewBulkAddTagsInputBody([]int64{int64(123)}, []string{"Tags_example"}) // BulkAddTagsInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesBulkActionsAPI.V3BatchAddAnalysisTags(context.Background()).BulkAddTagsInputBody(bulkAddTagsInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesBulkActionsAPI.V3BatchAddAnalysisTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3BatchAddAnalysisTags`: BulkAddTagsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesBulkActionsAPI.V3BatchAddAnalysisTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3BatchAddAnalysisTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkAddTagsInputBody** | [**BulkAddTagsInputBody**](BulkAddTagsInputBody.md) |  | 

### Return type

[**BulkAddTagsOutputBody**](BulkAddTagsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BatchDeleteAnalyses

> V3BatchDeleteAnalyses(ctx).BulkDeleteAnalysesInputBody(bulkDeleteAnalysesInputBody).Execute()

Delete multiple analyses.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	bulkDeleteAnalysesInputBody := *revengai.NewBulkDeleteAnalysesInputBody([]int64{int64(123)}) // BulkDeleteAnalysesInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	r, err := apiClient.AnalysesBulkActionsAPI.V3BatchDeleteAnalyses(context.Background()).BulkDeleteAnalysesInputBody(bulkDeleteAnalysesInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesBulkActionsAPI.V3BatchDeleteAnalyses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3BatchDeleteAnalysesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkDeleteAnalysesInputBody** | [**BulkDeleteAnalysesInputBody**](BulkDeleteAnalysesInputBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

