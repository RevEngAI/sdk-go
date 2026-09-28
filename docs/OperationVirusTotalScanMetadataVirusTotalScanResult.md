# OperationVirusTotalScanMetadataVirusTotalScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | **bool** | Whether the operation has reached a terminal state. | 
**Error** | Pointer to [**Status**](Status.md) | Failure detail, populated only when done is true and the operation failed. | [optional] 
**Metadata** | Pointer to [**VirusTotalScanMetadata**](VirusTotalScanMetadata.md) | In-flight information and details. | [optional] 
**Name** | **string** | API resource name. | 
**Response** | Pointer to [**VirusTotalScanResult**](VirusTotalScanResult.md) | Result, set only when done is true and the operation succeeded. | [optional] 

## Methods

### NewOperationVirusTotalScanMetadataVirusTotalScanResult

`func NewOperationVirusTotalScanMetadataVirusTotalScanResult(done bool, name string, ) *OperationVirusTotalScanMetadataVirusTotalScanResult`

NewOperationVirusTotalScanMetadataVirusTotalScanResult instantiates a new OperationVirusTotalScanMetadataVirusTotalScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationVirusTotalScanMetadataVirusTotalScanResultWithDefaults

`func NewOperationVirusTotalScanMetadataVirusTotalScanResultWithDefaults() *OperationVirusTotalScanMetadataVirusTotalScanResult`

NewOperationVirusTotalScanMetadataVirusTotalScanResultWithDefaults instantiates a new OperationVirusTotalScanMetadataVirusTotalScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) SetDone(v bool)`

SetDone sets Done field to given value.


### GetError

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) GetError() Status`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) GetErrorOk() (*Status, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) SetError(v Status)`

SetError sets Error field to given value.

### HasError

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMetadata

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) GetMetadata() VirusTotalScanMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) GetMetadataOk() (*VirusTotalScanMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) SetMetadata(v VirusTotalScanMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) SetName(v string)`

SetName sets Name field to given value.


### GetResponse

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) GetResponse() VirusTotalScanResult`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) GetResponseOk() (*VirusTotalScanResult, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) SetResponse(v VirusTotalScanResult)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *OperationVirusTotalScanMetadataVirusTotalScanResult) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


