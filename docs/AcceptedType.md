# AcceptedType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **bool** | False when the analysis already held a type under this name and kind, which this request resolved to rather than replaced. | 
**DataTypeId** | **int64** | The data type the suggestion is now stored as. | 
**Key** | **string** | The suggestion this entry answers. | 
**SkippedMembers** | **int64** | Members left out of the stored type because no offset or width was established for them. | 

## Methods

### NewAcceptedType

`func NewAcceptedType(created bool, dataTypeId int64, key string, skippedMembers int64, ) *AcceptedType`

NewAcceptedType instantiates a new AcceptedType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptedTypeWithDefaults

`func NewAcceptedTypeWithDefaults() *AcceptedType`

NewAcceptedTypeWithDefaults instantiates a new AcceptedType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AcceptedType) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AcceptedType) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AcceptedType) SetCreated(v bool)`

SetCreated sets Created field to given value.


### GetDataTypeId

`func (o *AcceptedType) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *AcceptedType) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *AcceptedType) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.


### GetKey

`func (o *AcceptedType) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AcceptedType) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AcceptedType) SetKey(v string)`

SetKey sets Key field to given value.


### GetSkippedMembers

`func (o *AcceptedType) GetSkippedMembers() int64`

GetSkippedMembers returns the SkippedMembers field if non-nil, zero value otherwise.

### GetSkippedMembersOk

`func (o *AcceptedType) GetSkippedMembersOk() (*int64, bool)`

GetSkippedMembersOk returns a tuple with the SkippedMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedMembers

`func (o *AcceptedType) SetSkippedMembers(v int64)`

SetSkippedMembers sets SkippedMembers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


