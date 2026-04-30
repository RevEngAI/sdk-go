# StatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**ConversationUuid** | **string** |  | 
**Status** | **string** |  | 
**TraceId** | Pointer to **string** | OpenTelemetry trace ID for this run. Use this to look up tool call spans in your trace backend. | [optional] 

## Methods

### NewStatusResponse

`func NewStatusResponse(conversationUuid string, status string, ) *StatusResponse`

NewStatusResponse instantiates a new StatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseWithDefaults

`func NewStatusResponseWithDefaults() *StatusResponse`

NewStatusResponseWithDefaults instantiates a new StatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *StatusResponse) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *StatusResponse) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *StatusResponse) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *StatusResponse) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetConversationUuid

`func (o *StatusResponse) GetConversationUuid() string`

GetConversationUuid returns the ConversationUuid field if non-nil, zero value otherwise.

### GetConversationUuidOk

`func (o *StatusResponse) GetConversationUuidOk() (*string, bool)`

GetConversationUuidOk returns a tuple with the ConversationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationUuid

`func (o *StatusResponse) SetConversationUuid(v string)`

SetConversationUuid sets ConversationUuid field to given value.


### GetStatus

`func (o *StatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTraceId

`func (o *StatusResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *StatusResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *StatusResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *StatusResponse) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


