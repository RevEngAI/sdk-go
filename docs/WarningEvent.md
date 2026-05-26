# WarningEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | **int32** |  | 
**Identifiers** | Pointer to **[]string** |  | [optional] 
**Kind** | **string** |  | 
**Message** | **string** |  | 
**Seq** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewWarningEvent

`func NewWarningEvent(attempt int32, kind string, message string, seq int32, type_ string, ) *WarningEvent`

NewWarningEvent instantiates a new WarningEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningEventWithDefaults

`func NewWarningEventWithDefaults() *WarningEvent`

NewWarningEventWithDefaults instantiates a new WarningEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *WarningEvent) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *WarningEvent) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *WarningEvent) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.


### GetIdentifiers

`func (o *WarningEvent) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *WarningEvent) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *WarningEvent) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *WarningEvent) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiersNil

`func (o *WarningEvent) SetIdentifiersNil(b bool)`

 SetIdentifiersNil sets the value for Identifiers to be an explicit nil

### UnsetIdentifiers
`func (o *WarningEvent) UnsetIdentifiers()`

UnsetIdentifiers ensures that no value is present for Identifiers, not even an explicit nil
### GetKind

`func (o *WarningEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WarningEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WarningEvent) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMessage

`func (o *WarningEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WarningEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WarningEvent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSeq

`func (o *WarningEvent) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *WarningEvent) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *WarningEvent) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetType

`func (o *WarningEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WarningEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WarningEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


