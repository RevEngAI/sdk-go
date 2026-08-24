# DataTypesAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3CopyFunctionSignatures**](DataTypesAPI.md#V3CopyFunctionSignatures) | **Post** /v3/analyses/{analysis_id}/signatures/copy | Copy function signatures
[**V3CreateAnalysisDataTypes**](DataTypesAPI.md#V3CreateAnalysisDataTypes) | **Post** /v3/analyses/{analysis_id}/data-types | Create an analysis&#39;s data types
[**V3GetAnalysisDataType**](DataTypesAPI.md#V3GetAnalysisDataType) | **Get** /v3/analyses/{analysis_id}/data-types/{data_type_id} | Get one of an analysis&#39;s data types
[**V3GetAnalysisDataTypeHistory**](DataTypesAPI.md#V3GetAnalysisDataTypeHistory) | **Get** /v3/analyses/{analysis_id}/data-types/{data_type_id}/history | Get a data type&#39;s edit history
[**V3GetFunctionSignature**](DataTypesAPI.md#V3GetFunctionSignature) | **Get** /v3/analyses/{analysis_id}/functions/{function_id}/signature | Get a function&#39;s signature
[**V3GetFunctionSignatureHistory**](DataTypesAPI.md#V3GetFunctionSignatureHistory) | **Get** /v3/analyses/{analysis_id}/functions/{function_id}/signature/history | Get a function signature&#39;s edit history
[**V3ListAnalysisDataTypes**](DataTypesAPI.md#V3ListAnalysisDataTypes) | **Get** /v3/analyses/{analysis_id}/data-types | List an analysis&#39;s data types
[**V3ListDataTypeFunctions**](DataTypesAPI.md#V3ListDataTypeFunctions) | **Get** /v3/analyses/{analysis_id}/data-types/{data_type_id}/functions | List the functions using a data type
[**V3ListFunctionSignatures**](DataTypesAPI.md#V3ListFunctionSignatures) | **Get** /v3/functions/signatures | Get signatures for many functions
[**V3UpdateAnalysisDataTypes**](DataTypesAPI.md#V3UpdateAnalysisDataTypes) | **Put** /v3/analyses/{analysis_id}/data-types | Update an analysis&#39;s data types
[**V3UpdateFunctionSignature**](DataTypesAPI.md#V3UpdateFunctionSignature) | **Put** /v3/analyses/{analysis_id}/functions/{function_id}/signature | Update a function&#39;s signature



## V3CopyFunctionSignatures

> CopyFunctionSignaturesOutputBody V3CopyFunctionSignatures(ctx, analysisId).CopyFunctionSignaturesInputBody(copyFunctionSignaturesInputBody).Execute()

Copy function signatures



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
	analysisId := int64(789) // int64 | Analysis ID
	copyFunctionSignaturesInputBody := *revengai.NewCopyFunctionSignaturesInputBody([]revengai.CopySignatureItem{*revengai.NewCopySignatureItem(int64(123), int64(123))}) // CopyFunctionSignaturesInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3CopyFunctionSignatures(context.Background(), analysisId).CopyFunctionSignaturesInputBody(copyFunctionSignaturesInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3CopyFunctionSignatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3CopyFunctionSignatures`: CopyFunctionSignaturesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3CopyFunctionSignatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3CopyFunctionSignaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **copyFunctionSignaturesInputBody** | [**CopyFunctionSignaturesInputBody**](CopyFunctionSignaturesInputBody.md) |  | 

### Return type

[**CopyFunctionSignaturesOutputBody**](CopyFunctionSignaturesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3CreateAnalysisDataTypes

> AnalysisDataTypesOutputBody V3CreateAnalysisDataTypes(ctx, analysisId).CreateAnalysisDataTypesInputBody(createAnalysisDataTypesInputBody).Execute()

Create an analysis's data types



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
	analysisId := int64(789) // int64 | Analysis ID
	createAnalysisDataTypesInputBody := *revengai.NewCreateAnalysisDataTypesInputBody([]revengai.CreateDataTypeEntry{revengai.CreateDataTypeEntry{CreateArrayDataType: revengai.NewCreateArrayDataType(*revengai.NewArrayDefinition(), "Kind_example", "Name_example")}}) // CreateAnalysisDataTypesInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3CreateAnalysisDataTypes(context.Background(), analysisId).CreateAnalysisDataTypesInputBody(createAnalysisDataTypesInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3CreateAnalysisDataTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3CreateAnalysisDataTypes`: AnalysisDataTypesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3CreateAnalysisDataTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3CreateAnalysisDataTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAnalysisDataTypesInputBody** | [**CreateAnalysisDataTypesInputBody**](CreateAnalysisDataTypesInputBody.md) |  | 

### Return type

[**AnalysisDataTypesOutputBody**](AnalysisDataTypesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetAnalysisDataType

> DataTypeEntry V3GetAnalysisDataType(ctx, analysisId, dataTypeId).Execute()

Get one of an analysis's data types



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
	analysisId := int64(789) // int64 | Analysis ID
	dataTypeId := int64(789) // int64 | Data type ID, as returned by the data types list for this analysis. 0 is a valid id.

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3GetAnalysisDataType(context.Background(), analysisId, dataTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3GetAnalysisDataType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetAnalysisDataType`: DataTypeEntry
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3GetAnalysisDataType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**dataTypeId** | **int64** | Data type ID, as returned by the data types list for this analysis. 0 is a valid id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetAnalysisDataTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataTypeEntry**](DataTypeEntry.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetAnalysisDataTypeHistory

> GetDataTypeHistoryBody V3GetAnalysisDataTypeHistory(ctx, analysisId, dataTypeId).Execute()

Get a data type's edit history



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
	analysisId := int64(789) // int64 | Analysis ID
	dataTypeId := int64(789) // int64 | Data type ID, as returned by the data types list for this analysis. 0 is a valid id.

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3GetAnalysisDataTypeHistory(context.Background(), analysisId, dataTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3GetAnalysisDataTypeHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetAnalysisDataTypeHistory`: GetDataTypeHistoryBody
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3GetAnalysisDataTypeHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**dataTypeId** | **int64** | Data type ID, as returned by the data types list for this analysis. 0 is a valid id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetAnalysisDataTypeHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDataTypeHistoryBody**](GetDataTypeHistoryBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetFunctionSignature

> FunctionSignatureBody V3GetFunctionSignature(ctx, analysisId, functionId).IncludeDataTypes(includeDataTypes).Execute()

Get a function's signature



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
	analysisId := int64(789) // int64 | Analysis ID
	functionId := int64(789) // int64 | Function ID
	includeDataTypes := true // bool | Include the data types the signature names in the response. (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3GetFunctionSignature(context.Background(), analysisId, functionId).IncludeDataTypes(includeDataTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3GetFunctionSignature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetFunctionSignature`: FunctionSignatureBody
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3GetFunctionSignature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetFunctionSignatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeDataTypes** | **bool** | Include the data types the signature names in the response. | 

### Return type

[**FunctionSignatureBody**](FunctionSignatureBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetFunctionSignatureHistory

> GetFunctionSignatureHistoryBody V3GetFunctionSignatureHistory(ctx, analysisId, functionId).Execute()

Get a function signature's edit history



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
	analysisId := int64(789) // int64 | Analysis ID
	functionId := int64(789) // int64 | Function ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3GetFunctionSignatureHistory(context.Background(), analysisId, functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3GetFunctionSignatureHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetFunctionSignatureHistory`: GetFunctionSignatureHistoryBody
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3GetFunctionSignatureHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetFunctionSignatureHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFunctionSignatureHistoryBody**](GetFunctionSignatureHistoryBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListAnalysisDataTypes

> ListAnalysisDataTypesOutputBody V3ListAnalysisDataTypes(ctx, analysisId).Offset(offset).Limit(limit).Kind(kind).Namespace(namespace).Search(search).SourceType(sourceType).OrderBy(orderBy).Order(order).Execute()

List an analysis's data types



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
	analysisId := int64(789) // int64 | Analysis ID
	offset := int64(789) // int64 | Pagination offset. (optional) (default to 0)
	limit := int64(789) // int64 | Page size. (optional) (default to 100)
	kind := []string{"Kind_example"} // []string | Only return types of these kinds. Repeat for more than one; empty means no filter. (optional)
	namespace := []*string{"Inner_example"} // []*string | Only return types in these namespaces, matched exactly. Omit for no filter; pass an empty value (namespace=) for the binary's own types, which have no namespace. (optional)
	search := "search_example" // string | Only return types whose name contains this term. Wildcards in the term are matched literally. (optional)
	sourceType := []string{"SourceType_example"} // []string | Only return types from these sources. Empty means no filter. (optional)
	orderBy := "orderBy_example" // string | Field to order by. name orders by namespace, then name, then kind; size orders by size with types of unknown size last, then by namespace, name and kind. (optional) (default to "name")
	order := "order_example" // string | Sort direction. (optional) (default to "ASC")

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3ListAnalysisDataTypes(context.Background(), analysisId).Offset(offset).Limit(limit).Kind(kind).Namespace(namespace).Search(search).SourceType(sourceType).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3ListAnalysisDataTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListAnalysisDataTypes`: ListAnalysisDataTypesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3ListAnalysisDataTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ListAnalysisDataTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int64** | Pagination offset. | [default to 0]
 **limit** | **int64** | Page size. | [default to 100]
 **kind** | **[]string** | Only return types of these kinds. Repeat for more than one; empty means no filter. | 
 **namespace** | **[]string** | Only return types in these namespaces, matched exactly. Omit for no filter; pass an empty value (namespace&#x3D;) for the binary&#39;s own types, which have no namespace. | 
 **search** | **string** | Only return types whose name contains this term. Wildcards in the term are matched literally. | 
 **sourceType** | **[]string** | Only return types from these sources. Empty means no filter. | 
 **orderBy** | **string** | Field to order by. name orders by namespace, then name, then kind; size orders by size with types of unknown size last, then by namespace, name and kind. | [default to &quot;name&quot;]
 **order** | **string** | Sort direction. | [default to &quot;ASC&quot;]

### Return type

[**ListAnalysisDataTypesOutputBody**](ListAnalysisDataTypesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListDataTypeFunctions

> ListDataTypeFunctionsBody V3ListDataTypeFunctions(ctx, analysisId, dataTypeId).PageSize(pageSize).AfterFunctionId(afterFunctionId).Execute()

List the functions using a data type



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
	analysisId := int64(789) // int64 | Analysis ID
	dataTypeId := int64(789) // int64 | Data type ID, as returned by the data types list for this analysis. 0 is a valid id.
	pageSize := int64(789) // int64 | Page size. (optional) (default to 50)
	afterFunctionId := int64(789) // int64 | Return functions with an ID greater than this. Pass the previous page's next_after_function_id; 0 starts at the first function. (optional) (default to 0)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3ListDataTypeFunctions(context.Background(), analysisId, dataTypeId).PageSize(pageSize).AfterFunctionId(afterFunctionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3ListDataTypeFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListDataTypeFunctions`: ListDataTypeFunctionsBody
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3ListDataTypeFunctions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**dataTypeId** | **int64** | Data type ID, as returned by the data types list for this analysis. 0 is a valid id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ListDataTypeFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int64** | Page size. | [default to 50]
 **afterFunctionId** | **int64** | Return functions with an ID greater than this. Pass the previous page&#39;s next_after_function_id; 0 starts at the first function. | [default to 0]

### Return type

[**ListDataTypeFunctionsBody**](ListDataTypeFunctionsBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListFunctionSignatures

> ListFunctionSignaturesOutputBody V3ListFunctionSignatures(ctx).FunctionIds(functionIds).IncludeDataTypes(includeDataTypes).Execute()

Get signatures for many functions



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
	functionIds := []int64{int64(123)} // []int64 | Function IDs to fetch signatures for.
	includeDataTypes := true // bool | Include the data types the signatures name in the response. (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3ListFunctionSignatures(context.Background()).FunctionIds(functionIds).IncludeDataTypes(includeDataTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3ListFunctionSignatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListFunctionSignatures`: ListFunctionSignaturesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3ListFunctionSignatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ListFunctionSignaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionIds** | **[]int64** | Function IDs to fetch signatures for. | 
 **includeDataTypes** | **bool** | Include the data types the signatures name in the response. | 

### Return type

[**ListFunctionSignaturesOutputBody**](ListFunctionSignaturesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UpdateAnalysisDataTypes

> AnalysisDataTypesOutputBody V3UpdateAnalysisDataTypes(ctx, analysisId).UpdateAnalysisDataTypesInputBody(updateAnalysisDataTypesInputBody).Execute()

Update an analysis's data types



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
	analysisId := int64(789) // int64 | Analysis ID
	updateAnalysisDataTypesInputBody := *revengai.NewUpdateAnalysisDataTypesInputBody([]revengai.UpdateDataTypeEntry{revengai.UpdateDataTypeEntry{UpdateArrayDataType: revengai.NewUpdateArrayDataType(int64(123), *revengai.NewArrayDefinition(), "Kind_example", "Name_example")}}) // UpdateAnalysisDataTypesInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3UpdateAnalysisDataTypes(context.Background(), analysisId).UpdateAnalysisDataTypesInputBody(updateAnalysisDataTypesInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3UpdateAnalysisDataTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3UpdateAnalysisDataTypes`: AnalysisDataTypesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3UpdateAnalysisDataTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3UpdateAnalysisDataTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAnalysisDataTypesInputBody** | [**UpdateAnalysisDataTypesInputBody**](UpdateAnalysisDataTypesInputBody.md) |  | 

### Return type

[**AnalysisDataTypesOutputBody**](AnalysisDataTypesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UpdateFunctionSignature

> FunctionSignatureEntry V3UpdateFunctionSignature(ctx, analysisId, functionId).UpdateFunctionSignatureInputBody(updateFunctionSignatureInputBody).Execute()

Update a function's signature



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
	analysisId := int64(789) // int64 | Analysis ID
	functionId := int64(789) // int64 | Function ID
	updateFunctionSignatureInputBody := *revengai.NewUpdateFunctionSignatureInputBody([]revengai.SignatureParameterInput{*revengai.NewSignatureParameterInput(int64(123))}) // UpdateFunctionSignatureInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTypesAPI.V3UpdateFunctionSignature(context.Background(), analysisId, functionId).UpdateFunctionSignatureInputBody(updateFunctionSignatureInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTypesAPI.V3UpdateFunctionSignature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3UpdateFunctionSignature`: FunctionSignatureEntry
	fmt.Fprintf(os.Stdout, "Response from `DataTypesAPI.V3UpdateFunctionSignature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3UpdateFunctionSignatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFunctionSignatureInputBody** | [**UpdateFunctionSignatureInputBody**](UpdateFunctionSignatureInputBody.md) |  | 

### Return type

[**FunctionSignatureEntry**](FunctionSignatureEntry.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

