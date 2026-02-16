# \SearchAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchBinaries**](SearchAPI.md#SearchBinaries) | **Get** /v2/search/binaries | Binaries search
[**SearchCollections**](SearchAPI.md#SearchCollections) | **Get** /v2/search/collections | Collections search
[**SearchFunctions**](SearchAPI.md#SearchFunctions) | **Get** /v2/search/functions | Functions search
[**SearchTags**](SearchAPI.md#SearchTags) | **Get** /v2/search/tags | Tags search



## SearchBinaries

> BaseResponseBinarySearchResponse SearchBinaries(ctx).Page(page).PageSize(pageSize).PartialName(partialName).PartialSha256(partialSha256).Tags(tags).ModelName(modelName).UserFilesOnly(userFilesOnly).Execute()

Binaries search



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
	page := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page. (optional) (default to 10)
	partialName := "partialName_example" // string | The partial or full name of the binary being searched (optional)
	partialSha256 := "partialSha256_example" // string | The partial or full sha256 of the binary being searched (optional)
	tags := []string{"Inner_example"} // []string | The tags to be searched for (optional)
	modelName := "modelName_example" // string | The name of the model used to analyze the binary the function belongs to (optional)
	userFilesOnly := true // bool | Whether to only search user's uploaded files (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchBinaries(context.Background()).Page(page).PageSize(pageSize).PartialName(partialName).PartialSha256(partialSha256).Tags(tags).ModelName(modelName).UserFilesOnly(userFilesOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchBinaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchBinaries`: BaseResponseBinarySearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchBinaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBinariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | Number of items per page. | [default to 10]
 **partialName** | **string** | The partial or full name of the binary being searched | 
 **partialSha256** | **string** | The partial or full sha256 of the binary being searched | 
 **tags** | **[]string** | The tags to be searched for | 
 **modelName** | **string** | The name of the model used to analyze the binary the function belongs to | 
 **userFilesOnly** | **bool** | Whether to only search user&#39;s uploaded files | [default to false]

### Return type

[**BaseResponseBinarySearchResponse**](BaseResponseBinarySearchResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCollections

> BaseResponseCollectionSearchResponse SearchCollections(ctx).Page(page).PageSize(pageSize).PartialCollectionName(partialCollectionName).PartialBinaryName(partialBinaryName).PartialBinarySha256(partialBinarySha256).Tags(tags).ModelName(modelName).Filters(filters).OrderBy(orderBy).OrderByDirection(orderByDirection).Execute()

Collections search



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
	page := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page. (optional) (default to 10)
	partialCollectionName := "partialCollectionName_example" // string | The partial or full name of the collection being searched (optional)
	partialBinaryName := "partialBinaryName_example" // string | The partial or full name of the binary belonging to the collection (optional)
	partialBinarySha256 := "partialBinarySha256_example" // string | The partial or full sha256 of the binary belonging to the collection (optional)
	tags := []string{"Inner_example"} // []string | The tags to be searched for (optional)
	modelName := "modelName_example" // string | The name of the model used to analyze the binary the function belongs to (optional)
	filters := []openapiclient.Filters{openapiclient.Filters("official_only")} // []Filters | The filters to be used for the search (optional)
	orderBy := openapiclient.app__api__rest__v2__collections__enums__OrderBy("created") // AppApiRestV2CollectionsEnumsOrderBy | The field to sort the order by in the results (optional) (default to "created")
	orderByDirection := openapiclient.Order("ASC") // Order | The order direction in which to return results (optional) (default to "DESC")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchCollections(context.Background()).Page(page).PageSize(pageSize).PartialCollectionName(partialCollectionName).PartialBinaryName(partialBinaryName).PartialBinarySha256(partialBinarySha256).Tags(tags).ModelName(modelName).Filters(filters).OrderBy(orderBy).OrderByDirection(orderByDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCollections`: BaseResponseCollectionSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | Number of items per page. | [default to 10]
 **partialCollectionName** | **string** | The partial or full name of the collection being searched | 
 **partialBinaryName** | **string** | The partial or full name of the binary belonging to the collection | 
 **partialBinarySha256** | **string** | The partial or full sha256 of the binary belonging to the collection | 
 **tags** | **[]string** | The tags to be searched for | 
 **modelName** | **string** | The name of the model used to analyze the binary the function belongs to | 
 **filters** | [**[]Filters**](Filters.md) | The filters to be used for the search | 
 **orderBy** | [**AppApiRestV2CollectionsEnumsOrderBy**](AppApiRestV2CollectionsEnumsOrderBy.md) | The field to sort the order by in the results | [default to &quot;created&quot;]
 **orderByDirection** | [**Order**](Order.md) | The order direction in which to return results | [default to &quot;DESC&quot;]

### Return type

[**BaseResponseCollectionSearchResponse**](BaseResponseCollectionSearchResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFunctions

> BaseResponseFunctionSearchResponse SearchFunctions(ctx).Page(page).PageSize(pageSize).PartialName(partialName).ModelName(modelName).Execute()

Functions search



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
	page := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page. (optional) (default to 10)
	partialName := "partialName_example" // string | The partial or full name of the function being searched (optional)
	modelName := "modelName_example" // string | The name of the model used to analyze the binary the function belongs to (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchFunctions(context.Background()).Page(page).PageSize(pageSize).PartialName(partialName).ModelName(modelName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFunctions`: BaseResponseFunctionSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | Number of items per page. | [default to 10]
 **partialName** | **string** | The partial or full name of the function being searched | 
 **modelName** | **string** | The name of the model used to analyze the binary the function belongs to | 

### Return type

[**BaseResponseFunctionSearchResponse**](BaseResponseFunctionSearchResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTags

> BaseResponseTagSearchResponse SearchTags(ctx).PartialName(partialName).Page(page).PageSize(pageSize).Execute()

Tags search



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
	partialName := "partialName_example" // string | The partial or full name of the tag to search for
	page := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchTags(context.Background()).PartialName(partialName).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTags`: BaseResponseTagSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partialName** | **string** | The partial or full name of the tag to search for | 
 **page** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | Number of items per page. | [default to 10]

### Return type

[**BaseResponseTagSearchResponse**](BaseResponseTagSearchResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

