# AnalysesXRefsAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetXrefByVaddr**](AnalysesXRefsAPI.md#GetXrefByVaddr) | **Get** /v2/analyses/{analysis_id}/xrefs/{vaddr} | [Beta] Look up xrefs by virtual address



## GetXrefByVaddr

> BaseResponseXrefResponse GetXrefByVaddr(ctx, analysisId, vaddr).Execute()

[Beta] Look up xrefs by virtual address



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
	vaddr := int32(56) // int32 | Virtual address to match against xrefs

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesXRefsAPI.GetXrefByVaddr(context.Background(), analysisId, vaddr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesXRefsAPI.GetXrefByVaddr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetXrefByVaddr`: BaseResponseXrefResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesXRefsAPI.GetXrefByVaddr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 
**vaddr** | **int32** | Virtual address to match against xrefs | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXrefByVaddrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseResponseXrefResponse**](BaseResponseXrefResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

