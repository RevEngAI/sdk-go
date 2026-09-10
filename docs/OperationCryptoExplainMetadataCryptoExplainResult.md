# OperationCryptoExplainMetadataCryptoExplainResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | **bool** | Whether the operation has reached a terminal state. | 
**Error** | Pointer to [**Status**](Status.md) | Failure detail, populated only when done is true and the operation failed. | [optional] 
**Metadata** | Pointer to [**CryptoExplainMetadata**](CryptoExplainMetadata.md) | In-flight information and details. | [optional] 
**Name** | **string** | API resource name. | 
**Response** | Pointer to [**CryptoExplainResult**](CryptoExplainResult.md) | Result, set only when done is true and the operation succeeded. | [optional] 

## Methods

### NewOperationCryptoExplainMetadataCryptoExplainResult

`func NewOperationCryptoExplainMetadataCryptoExplainResult(done bool, name string, ) *OperationCryptoExplainMetadataCryptoExplainResult`

NewOperationCryptoExplainMetadataCryptoExplainResult instantiates a new OperationCryptoExplainMetadataCryptoExplainResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationCryptoExplainMetadataCryptoExplainResultWithDefaults

`func NewOperationCryptoExplainMetadataCryptoExplainResultWithDefaults() *OperationCryptoExplainMetadataCryptoExplainResult`

NewOperationCryptoExplainMetadataCryptoExplainResultWithDefaults instantiates a new OperationCryptoExplainMetadataCryptoExplainResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) SetDone(v bool)`

SetDone sets Done field to given value.


### GetError

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) GetError() Status`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) GetErrorOk() (*Status, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) SetError(v Status)`

SetError sets Error field to given value.

### HasError

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMetadata

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) GetMetadata() CryptoExplainMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) GetMetadataOk() (*CryptoExplainMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) SetMetadata(v CryptoExplainMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) SetName(v string)`

SetName sets Name field to given value.


### GetResponse

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) GetResponse() CryptoExplainResult`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) GetResponseOk() (*CryptoExplainResult, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) SetResponse(v CryptoExplainResult)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *OperationCryptoExplainMetadataCryptoExplainResult) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


