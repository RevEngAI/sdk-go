# ConversationsAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRun**](ConversationsAPI.md#CancelRun) | **Post** /v2/conversations/{id}/cancel | Cancel an active run
[**ConfirmTool**](ConversationsAPI.md#ConfirmTool) | **Post** /v2/conversations/{id}/confirm | Approve or reject a pending tool confirmation
[**CreateConversation**](ConversationsAPI.md#CreateConversation) | **Post** /v2/conversations | Create a new conversation
[**GetConversation**](ConversationsAPI.md#GetConversation) | **Get** /v2/conversations/{id} | Get a conversation with its events
[**ListConversations**](ConversationsAPI.md#ListConversations) | **Get** /v2/conversations | List conversations for the authenticated user
[**SendMessage**](ConversationsAPI.md#SendMessage) | **Post** /v2/conversations/{id}/messages | Send a message and start an agentic run
[**StreamEvents**](ConversationsAPI.md#StreamEvents) | **Get** /v2/conversations/{id}/events | Stream conversation events (SSE)



## CancelRun

> StatusResponse CancelRun(ctx, id).Execute()

Cancel an active run



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Conversation UUID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.CancelRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.CancelRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelRun`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.CancelRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Conversation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmTool

> StatusResponse ConfirmTool(ctx, id).ConfirmToolInputBody(confirmToolInputBody).Execute()

Approve or reject a pending tool confirmation



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Conversation UUID
	confirmToolInputBody := *revengai.NewConfirmToolInputBody(false) // ConfirmToolInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.ConfirmTool(context.Background(), id).ConfirmToolInputBody(confirmToolInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.ConfirmTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmTool`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.ConfirmTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Conversation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmToolInputBody** | [**ConfirmToolInputBody**](ConfirmToolInputBody.md) |  | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConversation

> Conversation CreateConversation(ctx).CreateConversationRequest(createConversationRequest).Execute()

Create a new conversation



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
	createConversationRequest := *revengai.NewCreateConversationRequest() // CreateConversationRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.CreateConversation(context.Background()).CreateConversationRequest(createConversationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.CreateConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConversation`: Conversation
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.CreateConversation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConversationRequest** | [**CreateConversationRequest**](CreateConversationRequest.md) |  | 

### Return type

[**Conversation**](Conversation.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversation

> ConversationWithEvents GetConversation(ctx, id).Execute()

Get a conversation with its events



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Conversation UUID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversation`: ConversationWithEvents
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Conversation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConversationWithEvents**](ConversationWithEvents.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversations

> []Conversation ListConversations(ctx).Execute()

List conversations for the authenticated user



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

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.ListConversations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.ListConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConversations`: []Conversation
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.ListConversations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConversationsRequest struct via the builder pattern


### Return type

[**[]Conversation**](Conversation.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> StatusResponse SendMessage(ctx, id).SendMessageRequest(sendMessageRequest).Execute()

Send a message and start an agentic run



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Conversation UUID
	sendMessageRequest := *revengai.NewSendMessageRequest("Content_example") // SendMessageRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.SendMessage(context.Background(), id).SendMessageRequest(sendMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.SendMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMessage`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.SendMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Conversation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendMessageRequest** | [**SendMessageRequest**](SendMessageRequest.md) |  | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamEvents

> []ServerSentEventsInner StreamEvents(ctx, id).LastEventId(lastEventId).Execute()

Stream conversation events (SSE)



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Conversation UUID
	lastEventId := int64(789) // int64 | Replay events after this ID (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.StreamEvents(context.Background(), id).LastEventId(lastEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.StreamEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamEvents`: []ServerSentEventsInner
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.StreamEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Conversation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastEventId** | **int64** | Replay events after this ID | 

### Return type

[**[]ServerSentEventsInner**](ServerSentEventsInner.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

