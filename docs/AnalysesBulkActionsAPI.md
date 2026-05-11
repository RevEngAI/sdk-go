# AnalysesBulkActionsAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkAddAnalysisTags**](AnalysesBulkActionsAPI.md#BulkAddAnalysisTags) | **Patch** /v2/analyses/tags/add | Bulk Add Analysis Tags
[**BulkDeleteAnalyses**](AnalysesBulkActionsAPI.md#BulkDeleteAnalyses) | **Patch** /v2/analyses/delete | Bulk Delete Analyses



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
	revengai "github.com/RevEngAI/sdk-go/v3"
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

[APIKey](../README.md#APIKey)

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
	revengai "github.com/RevEngAI/sdk-go/v3"
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

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

