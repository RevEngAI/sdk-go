# TTPSElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Attack** | [**[]TTPSAttack**](TTPSAttack.md) |  | 
**Occurrences** | [**[]TTPSOccurance**](TTPSOccurance.md) |  | 
**Score** | **int32** |  | 

## Methods

### NewTTPSElement

`func NewTTPSElement(name string, attack []TTPSAttack, occurrences []TTPSOccurance, score int32, ) *TTPSElement`

NewTTPSElement instantiates a new TTPSElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTTPSElementWithDefaults

`func NewTTPSElementWithDefaults() *TTPSElement`

NewTTPSElementWithDefaults instantiates a new TTPSElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TTPSElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TTPSElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TTPSElement) SetName(v string)`

SetName sets Name field to given value.


### GetAttack

`func (o *TTPSElement) GetAttack() []TTPSAttack`

GetAttack returns the Attack field if non-nil, zero value otherwise.

### GetAttackOk

`func (o *TTPSElement) GetAttackOk() (*[]TTPSAttack, bool)`

GetAttackOk returns a tuple with the Attack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttack

`func (o *TTPSElement) SetAttack(v []TTPSAttack)`

SetAttack sets Attack field to given value.


### GetOccurrences

`func (o *TTPSElement) GetOccurrences() []TTPSOccurance`

GetOccurrences returns the Occurrences field if non-nil, zero value otherwise.

### GetOccurrencesOk

`func (o *TTPSElement) GetOccurrencesOk() (*[]TTPSOccurance, bool)`

GetOccurrencesOk returns a tuple with the Occurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrences

`func (o *TTPSElement) SetOccurrences(v []TTPSOccurance)`

SetOccurrences sets Occurrences field to given value.


### GetScore

`func (o *TTPSElement) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *TTPSElement) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *TTPSElement) SetScore(v int32)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


