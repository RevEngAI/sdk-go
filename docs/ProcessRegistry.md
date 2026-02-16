# ProcessRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**map[string][]Registry**](array.md) |  | 

## Methods

### NewProcessRegistry

`func NewProcessRegistry(success bool, data map[string][]Registry, ) *ProcessRegistry`

NewProcessRegistry instantiates a new ProcessRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessRegistryWithDefaults

`func NewProcessRegistryWithDefaults() *ProcessRegistry`

NewProcessRegistryWithDefaults instantiates a new ProcessRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ProcessRegistry) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ProcessRegistry) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ProcessRegistry) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ProcessRegistry) GetData() map[string][]Registry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProcessRegistry) GetDataOk() (*map[string][]Registry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProcessRegistry) SetData(v map[string][]Registry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


