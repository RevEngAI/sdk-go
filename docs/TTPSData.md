# TTPSData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | **int32** |  | 
**Ttps** | [**[]TTPSElement**](TTPSElement.md) |  | 

## Methods

### NewTTPSData

`func NewTTPSData(score int32, ttps []TTPSElement, ) *TTPSData`

NewTTPSData instantiates a new TTPSData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTTPSDataWithDefaults

`func NewTTPSDataWithDefaults() *TTPSData`

NewTTPSDataWithDefaults instantiates a new TTPSData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *TTPSData) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *TTPSData) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *TTPSData) SetScore(v int32)`

SetScore sets Score field to given value.


### GetTtps

`func (o *TTPSData) GetTtps() []TTPSElement`

GetTtps returns the Ttps field if non-nil, zero value otherwise.

### GetTtpsOk

`func (o *TTPSData) GetTtpsOk() (*[]TTPSElement, bool)`

GetTtpsOk returns a tuple with the Ttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtps

`func (o *TTPSData) SetTtps(v []TTPSElement)`

SetTtps sets Ttps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


