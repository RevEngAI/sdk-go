# AcceptTypeSuggestionsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | [**[]AcceptedType**](AcceptedType.md) | One entry per requested suggestion, in request order. | 
**DataTypes** | [**[]DataTypeEntry**](DataTypeEntry.md) | The type each requested suggestion resolved to, plus every type minted to satisfy one, ordered by data_type_id. | 

## Methods

### NewAcceptTypeSuggestionsOutputBody

`func NewAcceptTypeSuggestionsOutputBody(accepted []AcceptedType, dataTypes []DataTypeEntry, ) *AcceptTypeSuggestionsOutputBody`

NewAcceptTypeSuggestionsOutputBody instantiates a new AcceptTypeSuggestionsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptTypeSuggestionsOutputBodyWithDefaults

`func NewAcceptTypeSuggestionsOutputBodyWithDefaults() *AcceptTypeSuggestionsOutputBody`

NewAcceptTypeSuggestionsOutputBodyWithDefaults instantiates a new AcceptTypeSuggestionsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *AcceptTypeSuggestionsOutputBody) GetAccepted() []AcceptedType`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *AcceptTypeSuggestionsOutputBody) GetAcceptedOk() (*[]AcceptedType, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *AcceptTypeSuggestionsOutputBody) SetAccepted(v []AcceptedType)`

SetAccepted sets Accepted field to given value.


### SetAcceptedNil

`func (o *AcceptTypeSuggestionsOutputBody) SetAcceptedNil(b bool)`

 SetAcceptedNil sets the value for Accepted to be an explicit nil

### UnsetAccepted
`func (o *AcceptTypeSuggestionsOutputBody) UnsetAccepted()`

UnsetAccepted ensures that no value is present for Accepted, not even an explicit nil
### GetDataTypes

`func (o *AcceptTypeSuggestionsOutputBody) GetDataTypes() []DataTypeEntry`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *AcceptTypeSuggestionsOutputBody) GetDataTypesOk() (*[]DataTypeEntry, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *AcceptTypeSuggestionsOutputBody) SetDataTypes(v []DataTypeEntry)`

SetDataTypes sets DataTypes field to given value.


### SetDataTypesNil

`func (o *AcceptTypeSuggestionsOutputBody) SetDataTypesNil(b bool)`

 SetDataTypesNil sets the value for DataTypes to be an explicit nil

### UnsetDataTypes
`func (o *AcceptTypeSuggestionsOutputBody) UnsetDataTypes()`

UnsetDataTypes ensures that no value is present for DataTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


