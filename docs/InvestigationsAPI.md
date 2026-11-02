# InvestigationsAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3CreateInvestigation**](InvestigationsAPI.md#V3CreateInvestigation) | **Post** /v3/investigations | Create an investigation.
[**V3CreateInvestigationConversation**](InvestigationsAPI.md#V3CreateInvestigationConversation) | **Post** /v3/investigations/{investigation_id}/conversations | Start a conversation in an investigation.
[**V3DeleteInvestigation**](InvestigationsAPI.md#V3DeleteInvestigation) | **Delete** /v3/investigations/{investigation_id} | Delete an investigation.
[**V3GetInvestigation**](InvestigationsAPI.md#V3GetInvestigation) | **Get** /v3/investigations/{investigation_id} | Get an investigation.
[**V3GetInvestigationFinding**](InvestigationsAPI.md#V3GetInvestigationFinding) | **Get** /v3/investigations/{investigation_id}/findings/{finding_id} | Get a finding of an investigation.
[**V3ListInvestigationAgentRuns**](InvestigationsAPI.md#V3ListInvestigationAgentRuns) | **Get** /v3/investigations/agent-runs | List agent runs on investigations.
[**V3ListInvestigationConversations**](InvestigationsAPI.md#V3ListInvestigationConversations) | **Get** /v3/investigations/{investigation_id}/conversations | List the conversations of an investigation.
[**V3ListInvestigationFindings**](InvestigationsAPI.md#V3ListInvestigationFindings) | **Get** /v3/investigations/{investigation_id}/findings | List the findings of an investigation.
[**V3ListInvestigationSpend**](InvestigationsAPI.md#V3ListInvestigationSpend) | **Get** /v3/investigations/{investigation_id}/spend | List what an investigation spent.
[**V3ListInvestigations**](InvestigationsAPI.md#V3ListInvestigations) | **Get** /v3/investigations | List investigations.
[**V3PatchInvestigation**](InvestigationsAPI.md#V3PatchInvestigation) | **Patch** /v3/investigations/{investigation_id} | Update an investigation.
[**V3SetInvestigationBudget**](InvestigationsAPI.md#V3SetInvestigationBudget) | **Put** /v3/investigations/{investigation_id}/budget | Set the budget of an investigation.



## V3CreateInvestigation

> InvestigationBody V3CreateInvestigation(ctx).CreateInvestigationInputBody(createInvestigationInputBody).Execute()

Create an investigation.



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
	createInvestigationInputBody := *revengai.NewCreateInvestigationInputBody("CaseType_example", "Name_example", "Severity_example") // CreateInvestigationInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3CreateInvestigation(context.Background()).CreateInvestigationInputBody(createInvestigationInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3CreateInvestigation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3CreateInvestigation`: InvestigationBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3CreateInvestigation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3CreateInvestigationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInvestigationInputBody** | [**CreateInvestigationInputBody**](CreateInvestigationInputBody.md) |  | 

### Return type

[**InvestigationBody**](InvestigationBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3CreateInvestigationConversation

> ConversationBody V3CreateInvestigationConversation(ctx, investigationId).CreateConversationInputBody(createConversationInputBody).Execute()

Start a conversation in an investigation.



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
	investigationId := int64(789) // int64 | 
	createConversationInputBody := *revengai.NewCreateConversationInputBody() // CreateConversationInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3CreateInvestigationConversation(context.Background(), investigationId).CreateConversationInputBody(createConversationInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3CreateInvestigationConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3CreateInvestigationConversation`: ConversationBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3CreateInvestigationConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3CreateInvestigationConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConversationInputBody** | [**CreateConversationInputBody**](CreateConversationInputBody.md) |  | 

### Return type

[**ConversationBody**](ConversationBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeleteInvestigation

> V3DeleteInvestigation(ctx, investigationId).Execute()

Delete an investigation.



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
	investigationId := int64(789) // int64 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	r, err := apiClient.InvestigationsAPI.V3DeleteInvestigation(context.Background(), investigationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3DeleteInvestigation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DeleteInvestigationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetInvestigation

> GetInvestigationOutputBody V3GetInvestigation(ctx, investigationId).Execute()

Get an investigation.



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
	investigationId := int64(789) // int64 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3GetInvestigation(context.Background(), investigationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3GetInvestigation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetInvestigation`: GetInvestigationOutputBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3GetInvestigation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetInvestigationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInvestigationOutputBody**](GetInvestigationOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetInvestigationFinding

> GetFindingOutputBody V3GetInvestigationFinding(ctx, investigationId, findingId).Execute()

Get a finding of an investigation.



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
	investigationId := int64(789) // int64 | 
	findingId := int64(789) // int64 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3GetInvestigationFinding(context.Background(), investigationId, findingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3GetInvestigationFinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetInvestigationFinding`: GetFindingOutputBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3GetInvestigationFinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **int64** |  | 
**findingId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetInvestigationFindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFindingOutputBody**](GetFindingOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListInvestigationAgentRuns

> ListAgentRunsOutputBody V3ListInvestigationAgentRuns(ctx).InvestigationId(investigationId).OwnerId(ownerId).Status(status).Severity(severity).Search(search).Limit(limit).Offset(offset).Execute()

List agent runs on investigations.



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
	investigationId := int64(789) // int64 | Only return runs on the files of this investigation. 0 returns runs across every investigation you can see. (optional) (default to 0)
	ownerId := []int64{int64(123)} // []int64 | Only return runs in investigations owned by one of these user IDs. (optional)
	status := []string{"Status_example"} // []string | Only return runs in investigations in one of these statuses. (optional)
	severity := []string{"Severity_example"} // []string | Only return runs in investigations with one of these severities. (optional)
	search := "search_example" // string | Only return runs in investigations whose name contains this text, ignoring case. (optional)
	limit := int64(789) // int64 |  (optional) (default to 20)
	offset := int64(789) // int64 |  (optional) (default to 0)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3ListInvestigationAgentRuns(context.Background()).InvestigationId(investigationId).OwnerId(ownerId).Status(status).Severity(severity).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3ListInvestigationAgentRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListInvestigationAgentRuns`: ListAgentRunsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3ListInvestigationAgentRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ListInvestigationAgentRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigationId** | **int64** | Only return runs on the files of this investigation. 0 returns runs across every investigation you can see. | [default to 0]
 **ownerId** | **[]int64** | Only return runs in investigations owned by one of these user IDs. | 
 **status** | **[]string** | Only return runs in investigations in one of these statuses. | 
 **severity** | **[]string** | Only return runs in investigations with one of these severities. | 
 **search** | **string** | Only return runs in investigations whose name contains this text, ignoring case. | 
 **limit** | **int64** |  | [default to 20]
 **offset** | **int64** |  | [default to 0]

### Return type

[**ListAgentRunsOutputBody**](ListAgentRunsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListInvestigationConversations

> ListConversationsOutputBody V3ListInvestigationConversations(ctx, investigationId).Limit(limit).Offset(offset).Execute()

List the conversations of an investigation.



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
	investigationId := int64(789) // int64 | 
	limit := int64(789) // int64 |  (optional) (default to 20)
	offset := int64(789) // int64 |  (optional) (default to 0)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3ListInvestigationConversations(context.Background(), investigationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3ListInvestigationConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListInvestigationConversations`: ListConversationsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3ListInvestigationConversations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ListInvestigationConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** |  | [default to 20]
 **offset** | **int64** |  | [default to 0]

### Return type

[**ListConversationsOutputBody**](ListConversationsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListInvestigationFindings

> ListFindingsOutputBody V3ListInvestigationFindings(ctx, investigationId).FindingType(findingType).Severity(severity).Limit(limit).Offset(offset).Execute()

List the findings of an investigation.



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
	investigationId := int64(789) // int64 | 
	findingType := []string{"FindingType_example"} // []string | Only return findings of one of these types. (optional)
	severity := []string{"Severity_example"} // []string | Only return findings with one of these severities. (optional)
	limit := int64(789) // int64 |  (optional) (default to 20)
	offset := int64(789) // int64 |  (optional) (default to 0)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3ListInvestigationFindings(context.Background(), investigationId).FindingType(findingType).Severity(severity).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3ListInvestigationFindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListInvestigationFindings`: ListFindingsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3ListInvestigationFindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ListInvestigationFindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **findingType** | **[]string** | Only return findings of one of these types. | 
 **severity** | **[]string** | Only return findings with one of these severities. | 
 **limit** | **int64** |  | [default to 20]
 **offset** | **int64** |  | [default to 0]

### Return type

[**ListFindingsOutputBody**](ListFindingsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListInvestigationSpend

> ListSpendOutputBody V3ListInvestigationSpend(ctx, investigationId).Limit(limit).Offset(offset).Execute()

List what an investigation spent.



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
	investigationId := int64(789) // int64 | 
	limit := int64(789) // int64 |  (optional) (default to 20)
	offset := int64(789) // int64 |  (optional) (default to 0)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3ListInvestigationSpend(context.Background(), investigationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3ListInvestigationSpend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListInvestigationSpend`: ListSpendOutputBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3ListInvestigationSpend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ListInvestigationSpendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** |  | [default to 20]
 **offset** | **int64** |  | [default to 0]

### Return type

[**ListSpendOutputBody**](ListSpendOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListInvestigations

> ListInvestigationsOutputBody V3ListInvestigations(ctx).Status(status).Severity(severity).CaseType(caseType).OwnerId(ownerId).Search(search).OrderBy(orderBy).Order(order).Limit(limit).Offset(offset).Execute()

List investigations.



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
	status := []string{"Status_example"} // []string | Only return investigations in one of these statuses. (optional)
	severity := []string{"Severity_example"} // []string | Only return investigations with one of these severities. (optional)
	caseType := []string{"CaseType_example"} // []string | Only return investigations of one of these case types. (optional)
	ownerId := []int64{int64(123)} // []int64 | Only return investigations owned by one of these user IDs. (optional)
	search := "search_example" // string | Only return investigations whose name contains this text, ignoring case. (optional)
	orderBy := "orderBy_example" // string |  (optional) (default to "updated")
	order := "order_example" // string |  (optional) (default to "DESC")
	limit := int64(789) // int64 |  (optional) (default to 20)
	offset := int64(789) // int64 |  (optional) (default to 0)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3ListInvestigations(context.Background()).Status(status).Severity(severity).CaseType(caseType).OwnerId(ownerId).Search(search).OrderBy(orderBy).Order(order).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3ListInvestigations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListInvestigations`: ListInvestigationsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3ListInvestigations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ListInvestigationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **[]string** | Only return investigations in one of these statuses. | 
 **severity** | **[]string** | Only return investigations with one of these severities. | 
 **caseType** | **[]string** | Only return investigations of one of these case types. | 
 **ownerId** | **[]int64** | Only return investigations owned by one of these user IDs. | 
 **search** | **string** | Only return investigations whose name contains this text, ignoring case. | 
 **orderBy** | **string** |  | [default to &quot;updated&quot;]
 **order** | **string** |  | [default to &quot;DESC&quot;]
 **limit** | **int64** |  | [default to 20]
 **offset** | **int64** |  | [default to 0]

### Return type

[**ListInvestigationsOutputBody**](ListInvestigationsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PatchInvestigation

> InvestigationBody V3PatchInvestigation(ctx, investigationId).PatchInvestigationInputBody(patchInvestigationInputBody).Execute()

Update an investigation.



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
	investigationId := int64(789) // int64 | 
	patchInvestigationInputBody := *revengai.NewPatchInvestigationInputBody() // PatchInvestigationInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3PatchInvestigation(context.Background(), investigationId).PatchInvestigationInputBody(patchInvestigationInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3PatchInvestigation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PatchInvestigation`: InvestigationBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3PatchInvestigation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PatchInvestigationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchInvestigationInputBody** | [**PatchInvestigationInputBody**](PatchInvestigationInputBody.md) |  | 

### Return type

[**InvestigationBody**](InvestigationBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SetInvestigationBudget

> BudgetBody V3SetInvestigationBudget(ctx, investigationId).SetBudgetInputBody(setBudgetInputBody).Execute()

Set the budget of an investigation.



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
	investigationId := int64(789) // int64 | 
	setBudgetInputBody := *revengai.NewSetBudgetInputBody(NullableInt64(123)) // SetBudgetInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.InvestigationsAPI.V3SetInvestigationBudget(context.Background(), investigationId).SetBudgetInputBody(setBudgetInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvestigationsAPI.V3SetInvestigationBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SetInvestigationBudget`: BudgetBody
	fmt.Fprintf(os.Stdout, "Response from `InvestigationsAPI.V3SetInvestigationBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SetInvestigationBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setBudgetInputBody** | [**SetBudgetInputBody**](SetBudgetInputBody.md) |  | 

### Return type

[**BudgetBody**](BudgetBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

