# CollectionsAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCollection**](CollectionsAPI.md#CreateCollection) | **Post** /v2/collections | Creates new collection information
[**DeleteCollection**](CollectionsAPI.md#DeleteCollection) | **Delete** /v2/collections/{collection_id} | Deletes a collection
[**GetCollection**](CollectionsAPI.md#GetCollection) | **Get** /v2/collections/{collection_id} | Returns a collection
[**ListCollections**](CollectionsAPI.md#ListCollections) | **Get** /v2/collections | Gets basic collections information
[**UpdateCollection**](CollectionsAPI.md#UpdateCollection) | **Patch** /v2/collections/{collection_id} | Updates a collection
[**UpdateCollectionBinaries**](CollectionsAPI.md#UpdateCollectionBinaries) | **Patch** /v2/collections/{collection_id}/binaries | Updates a collection binaries
[**UpdateCollectionTags**](CollectionsAPI.md#UpdateCollectionTags) | **Patch** /v2/collections/{collection_id}/tags | Updates a collection tags
[**V3CreateCollection**](CollectionsAPI.md#V3CreateCollection) | **Post** /v3/collections | Create a collection.
[**V3GetCollection**](CollectionsAPI.md#V3GetCollection) | **Get** /v3/collections/{collection_id} | Get a collection.
[**V3ListCollections**](CollectionsAPI.md#V3ListCollections) | **Get** /v3/collections | List collections.



## CreateCollection

> BaseResponseCollectionResponse CreateCollection(ctx).CollectionCreateRequest(collectionCreateRequest).Execute()

Creates new collection information



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
	collectionCreateRequest := *revengai.NewCollectionCreateRequest("CollectionName_example", "Description_example", int32(123)) // CollectionCreateRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CreateCollection(context.Background()).CollectionCreateRequest(collectionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CreateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCollection`: BaseResponseCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CreateCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectionCreateRequest** | [**CollectionCreateRequest**](CollectionCreateRequest.md) |  | 

### Return type

[**BaseResponseCollectionResponse**](BaseResponseCollectionResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> BaseResponseBool DeleteCollection(ctx, collectionId).Execute()

Deletes a collection



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
	collectionId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.DeleteCollection(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.DeleteCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCollection`: BaseResponseBool
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.DeleteCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionRequest struct via the builder pattern


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


## GetCollection

> BaseResponseCollectionResponse GetCollection(ctx, collectionId).IncludeTags(includeTags).IncludeBinaries(includeBinaries).PageSize(pageSize).PageNumber(pageNumber).BinarySearchStr(binarySearchStr).Execute()

Returns a collection



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
	collectionId := int32(56) // int32 | 
	includeTags := true // bool |  (optional) (default to false)
	includeBinaries := true // bool |  (optional) (default to false)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	binarySearchStr := "binarySearchStr_example" // string |  (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.GetCollection(context.Background(), collectionId).IncludeTags(includeTags).IncludeBinaries(includeBinaries).PageSize(pageSize).PageNumber(pageNumber).BinarySearchStr(binarySearchStr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.GetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollection`: BaseResponseCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.GetCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTags** | **bool** |  | [default to false]
 **includeBinaries** | **bool** |  | [default to false]
 **pageSize** | **int32** |  | [default to 10]
 **pageNumber** | **int32** |  | [default to 1]
 **binarySearchStr** | **string** |  | 

### Return type

[**BaseResponseCollectionResponse**](BaseResponseCollectionResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollections

> BaseResponseListCollectionResults ListCollections(ctx).SearchTerm(searchTerm).Filters(filters).Limit(limit).Offset(offset).OrderBy(orderBy).Order(order).Execute()

Gets basic collections information



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
	searchTerm := "searchTerm_example" // string |  (optional) (default to "")
	filters := []revengai.Filters{revengai.Filters("official_only")} // []Filters |  (optional) (default to {})
	limit := int32(56) // int32 |  (optional) (default to 20)
	offset := int32(56) // int32 |  (optional) (default to 0)
	orderBy := revengai.app__api__rest__v2__collections__enums__OrderBy("created") // AppApiRestV2CollectionsEnumsOrderBy |  (optional) (default to "collection")
	order := revengai.Order("ASC") // Order |  (optional) (default to "ASC")

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.ListCollections(context.Background()).SearchTerm(searchTerm).Filters(filters).Limit(limit).Offset(offset).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.ListCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollections`: BaseResponseListCollectionResults
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.ListCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** |  | [default to &quot;&quot;]
 **filters** | [**[]Filters**](Filters.md) |  | [default to {}]
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]
 **orderBy** | [**AppApiRestV2CollectionsEnumsOrderBy**](AppApiRestV2CollectionsEnumsOrderBy.md) |  | [default to &quot;collection&quot;]
 **order** | [**Order**](Order.md) |  | [default to &quot;ASC&quot;]

### Return type

[**BaseResponseListCollectionResults**](BaseResponseListCollectionResults.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> BaseResponseCollectionResponse UpdateCollection(ctx, collectionId).CollectionUpdateRequest(collectionUpdateRequest).Execute()

Updates a collection



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
	collectionId := int32(56) // int32 | 
	collectionUpdateRequest := *revengai.NewCollectionUpdateRequest() // CollectionUpdateRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.UpdateCollection(context.Background(), collectionId).CollectionUpdateRequest(collectionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.UpdateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCollection`: BaseResponseCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.UpdateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionUpdateRequest** | [**CollectionUpdateRequest**](CollectionUpdateRequest.md) |  | 

### Return type

[**BaseResponseCollectionResponse**](BaseResponseCollectionResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollectionBinaries

> BaseResponseCollectionBinariesUpdateResponse UpdateCollectionBinaries(ctx, collectionId).CollectionBinariesUpdateRequest(collectionBinariesUpdateRequest).Execute()

Updates a collection binaries



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
	collectionId := int32(56) // int32 | 
	collectionBinariesUpdateRequest := *revengai.NewCollectionBinariesUpdateRequest([]int32{int32(123)}) // CollectionBinariesUpdateRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.UpdateCollectionBinaries(context.Background(), collectionId).CollectionBinariesUpdateRequest(collectionBinariesUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.UpdateCollectionBinaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCollectionBinaries`: BaseResponseCollectionBinariesUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.UpdateCollectionBinaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionBinariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionBinariesUpdateRequest** | [**CollectionBinariesUpdateRequest**](CollectionBinariesUpdateRequest.md) |  | 

### Return type

[**BaseResponseCollectionBinariesUpdateResponse**](BaseResponseCollectionBinariesUpdateResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollectionTags

> BaseResponseCollectionTagsUpdateResponse UpdateCollectionTags(ctx, collectionId).CollectionTagsUpdateRequest(collectionTagsUpdateRequest).Execute()

Updates a collection tags



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
	collectionId := int32(56) // int32 | 
	collectionTagsUpdateRequest := *revengai.NewCollectionTagsUpdateRequest([]string{"Tags_example"}) // CollectionTagsUpdateRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.UpdateCollectionTags(context.Background(), collectionId).CollectionTagsUpdateRequest(collectionTagsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.UpdateCollectionTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCollectionTags`: BaseResponseCollectionTagsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.UpdateCollectionTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionTagsUpdateRequest** | [**CollectionTagsUpdateRequest**](CollectionTagsUpdateRequest.md) |  | 

### Return type

[**BaseResponseCollectionTagsUpdateResponse**](BaseResponseCollectionTagsUpdateResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3CreateCollection

> CreateCollectionOutputBody V3CreateCollection(ctx).CreateCollectionInputBody(createCollectionInputBody).Execute()

Create a collection.



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
	createCollectionInputBody := *revengai.NewCreateCollectionInputBody("CollectionName_example", "CollectionScope_example", "Description_example", int64(123)) // CreateCollectionInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.V3CreateCollection(context.Background()).CreateCollectionInputBody(createCollectionInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.V3CreateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3CreateCollection`: CreateCollectionOutputBody
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.V3CreateCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3CreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCollectionInputBody** | [**CreateCollectionInputBody**](CreateCollectionInputBody.md) |  | 

### Return type

[**CreateCollectionOutputBody**](CreateCollectionOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetCollection

> GetCollectionOutputBody V3GetCollection(ctx, collectionId).IncludeTags(includeTags).IncludeBinaries(includeBinaries).PageSize(pageSize).PageNumber(pageNumber).BinarySearchStr(binarySearchStr).Execute()

Get a collection.



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
	collectionId := int64(789) // int64 | 
	includeTags := true // bool |  (optional)
	includeBinaries := true // bool |  (optional)
	pageSize := int64(789) // int64 |  (optional) (default to 10)
	pageNumber := int64(789) // int64 |  (optional) (default to 1)
	binarySearchStr := "binarySearchStr_example" // string |  (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.V3GetCollection(context.Background(), collectionId).IncludeTags(includeTags).IncludeBinaries(includeBinaries).PageSize(pageSize).PageNumber(pageNumber).BinarySearchStr(binarySearchStr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.V3GetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetCollection`: GetCollectionOutputBody
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.V3GetCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTags** | **bool** |  | 
 **includeBinaries** | **bool** |  | 
 **pageSize** | **int64** |  | [default to 10]
 **pageNumber** | **int64** |  | [default to 1]
 **binarySearchStr** | **string** |  | 

### Return type

[**GetCollectionOutputBody**](GetCollectionOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListCollections

> ListCollectionsOutputBody V3ListCollections(ctx).SearchTerm(searchTerm).Filters(filters).Limit(limit).Offset(offset).OrderBy(orderBy).Order(order).Execute()

List collections.



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
	searchTerm := "searchTerm_example" // string |  (optional)
	filters := []string{"Filters_example"} // []string |  (optional)
	limit := int64(789) // int64 |  (optional) (default to 20)
	offset := int64(789) // int64 |  (optional) (default to 0)
	orderBy := "orderBy_example" // string |  (optional) (default to "collection")
	order := "order_example" // string |  (optional) (default to "ASC")

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.V3ListCollections(context.Background()).SearchTerm(searchTerm).Filters(filters).Limit(limit).Offset(offset).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.V3ListCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListCollections`: ListCollectionsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.V3ListCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ListCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** |  | 
 **filters** | **[]string** |  | 
 **limit** | **int64** |  | [default to 20]
 **offset** | **int64** |  | [default to 0]
 **orderBy** | **string** |  | [default to &quot;collection&quot;]
 **order** | **string** |  | [default to &quot;ASC&quot;]

### Return type

[**ListCollectionsOutputBody**](ListCollectionsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

