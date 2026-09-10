# OperationWorkflowProgressResultBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | **bool** | Whether the operation has reached a terminal state. | 
**Error** | Pointer to [**Status**](Status.md) | Failure detail, populated only when done is true and the operation failed. | [optional] 
**Metadata** | Pointer to [**WorkflowProgress**](WorkflowProgress.md) | In-flight information and details. | [optional] 
**Name** | **string** | API resource name. | 
**Response** | Pointer to [**ResultBody**](ResultBody.md) | Result, set only when done is true and the operation succeeded. | [optional] 

## Methods

### NewOperationWorkflowProgressResultBody

`func NewOperationWorkflowProgressResultBody(done bool, name string, ) *OperationWorkflowProgressResultBody`

NewOperationWorkflowProgressResultBody instantiates a new OperationWorkflowProgressResultBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWorkflowProgressResultBodyWithDefaults

`func NewOperationWorkflowProgressResultBodyWithDefaults() *OperationWorkflowProgressResultBody`

NewOperationWorkflowProgressResultBodyWithDefaults instantiates a new OperationWorkflowProgressResultBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *OperationWorkflowProgressResultBody) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *OperationWorkflowProgressResultBody) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *OperationWorkflowProgressResultBody) SetDone(v bool)`

SetDone sets Done field to given value.


### GetError

`func (o *OperationWorkflowProgressResultBody) GetError() Status`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OperationWorkflowProgressResultBody) GetErrorOk() (*Status, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OperationWorkflowProgressResultBody) SetError(v Status)`

SetError sets Error field to given value.

### HasError

`func (o *OperationWorkflowProgressResultBody) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMetadata

`func (o *OperationWorkflowProgressResultBody) GetMetadata() WorkflowProgress`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OperationWorkflowProgressResultBody) GetMetadataOk() (*WorkflowProgress, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OperationWorkflowProgressResultBody) SetMetadata(v WorkflowProgress)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OperationWorkflowProgressResultBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OperationWorkflowProgressResultBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperationWorkflowProgressResultBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperationWorkflowProgressResultBody) SetName(v string)`

SetName sets Name field to given value.


### GetResponse

`func (o *OperationWorkflowProgressResultBody) GetResponse() ResultBody`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *OperationWorkflowProgressResultBody) GetResponseOk() (*ResultBody, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *OperationWorkflowProgressResultBody) SetResponse(v ResultBody)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *OperationWorkflowProgressResultBody) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


