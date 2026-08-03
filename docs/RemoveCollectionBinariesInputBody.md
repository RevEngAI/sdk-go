# RemoveCollectionBinariesInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binaries** | **[]int64** | Binary IDs to remove from the collection. Binary IDs not linked to the collection are ignored. | 

## Methods

### NewRemoveCollectionBinariesInputBody

`func NewRemoveCollectionBinariesInputBody(binaries []int64, ) *RemoveCollectionBinariesInputBody`

NewRemoveCollectionBinariesInputBody instantiates a new RemoveCollectionBinariesInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveCollectionBinariesInputBodyWithDefaults

`func NewRemoveCollectionBinariesInputBodyWithDefaults() *RemoveCollectionBinariesInputBody`

NewRemoveCollectionBinariesInputBodyWithDefaults instantiates a new RemoveCollectionBinariesInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaries

`func (o *RemoveCollectionBinariesInputBody) GetBinaries() []int64`

GetBinaries returns the Binaries field if non-nil, zero value otherwise.

### GetBinariesOk

`func (o *RemoveCollectionBinariesInputBody) GetBinariesOk() (*[]int64, bool)`

GetBinariesOk returns a tuple with the Binaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaries

`func (o *RemoveCollectionBinariesInputBody) SetBinaries(v []int64)`

SetBinaries sets Binaries field to given value.


### SetBinariesNil

`func (o *RemoveCollectionBinariesInputBody) SetBinariesNil(b bool)`

 SetBinariesNil sets the value for Binaries to be an explicit nil

### UnsetBinaries
`func (o *RemoveCollectionBinariesInputBody) UnsetBinaries()`

UnsetBinaries ensures that no value is present for Binaries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


