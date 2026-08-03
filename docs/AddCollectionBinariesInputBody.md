# AddCollectionBinariesInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binaries** | **[]int64** | Binary IDs to add to the collection. Binary IDs already linked to the collection are ignored. | 

## Methods

### NewAddCollectionBinariesInputBody

`func NewAddCollectionBinariesInputBody(binaries []int64, ) *AddCollectionBinariesInputBody`

NewAddCollectionBinariesInputBody instantiates a new AddCollectionBinariesInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCollectionBinariesInputBodyWithDefaults

`func NewAddCollectionBinariesInputBodyWithDefaults() *AddCollectionBinariesInputBody`

NewAddCollectionBinariesInputBodyWithDefaults instantiates a new AddCollectionBinariesInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaries

`func (o *AddCollectionBinariesInputBody) GetBinaries() []int64`

GetBinaries returns the Binaries field if non-nil, zero value otherwise.

### GetBinariesOk

`func (o *AddCollectionBinariesInputBody) GetBinariesOk() (*[]int64, bool)`

GetBinariesOk returns a tuple with the Binaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaries

`func (o *AddCollectionBinariesInputBody) SetBinaries(v []int64)`

SetBinaries sets Binaries field to given value.


### SetBinariesNil

`func (o *AddCollectionBinariesInputBody) SetBinariesNil(b bool)`

 SetBinariesNil sets the value for Binaries to be an explicit nil

### UnsetBinaries
`func (o *AddCollectionBinariesInputBody) UnsetBinaries()`

UnsetBinaries ensures that no value is present for Binaries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


