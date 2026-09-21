# OperationMetadataCapabilitiesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | **bool** | Whether the operation has reached a terminal state. | 
**Error** | Pointer to [**Status**](Status.md) | Failure detail, populated only when done is true and the operation failed. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) | In-flight information and details. | [optional] 
**Name** | **string** | API resource name. | 
**Response** | Pointer to [**CapabilitiesResult**](CapabilitiesResult.md) | Result, set only when done is true and the operation succeeded. | [optional] 

## Methods

### NewOperationMetadataCapabilitiesResult

`func NewOperationMetadataCapabilitiesResult(done bool, name string, ) *OperationMetadataCapabilitiesResult`

NewOperationMetadataCapabilitiesResult instantiates a new OperationMetadataCapabilitiesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationMetadataCapabilitiesResultWithDefaults

`func NewOperationMetadataCapabilitiesResultWithDefaults() *OperationMetadataCapabilitiesResult`

NewOperationMetadataCapabilitiesResultWithDefaults instantiates a new OperationMetadataCapabilitiesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *OperationMetadataCapabilitiesResult) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *OperationMetadataCapabilitiesResult) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *OperationMetadataCapabilitiesResult) SetDone(v bool)`

SetDone sets Done field to given value.


### GetError

`func (o *OperationMetadataCapabilitiesResult) GetError() Status`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OperationMetadataCapabilitiesResult) GetErrorOk() (*Status, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OperationMetadataCapabilitiesResult) SetError(v Status)`

SetError sets Error field to given value.

### HasError

`func (o *OperationMetadataCapabilitiesResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMetadata

`func (o *OperationMetadataCapabilitiesResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OperationMetadataCapabilitiesResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OperationMetadataCapabilitiesResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OperationMetadataCapabilitiesResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OperationMetadataCapabilitiesResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperationMetadataCapabilitiesResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperationMetadataCapabilitiesResult) SetName(v string)`

SetName sets Name field to given value.


### GetResponse

`func (o *OperationMetadataCapabilitiesResult) GetResponse() CapabilitiesResult`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *OperationMetadataCapabilitiesResult) GetResponseOk() (*CapabilitiesResult, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *OperationMetadataCapabilitiesResult) SetResponse(v CapabilitiesResult)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *OperationMetadataCapabilitiesResult) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


