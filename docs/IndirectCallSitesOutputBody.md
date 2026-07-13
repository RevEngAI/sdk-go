# IndirectCallSitesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **int64** |  | 
**Sites** | [**[]IndirectCallSite**](IndirectCallSite.md) |  | 

## Methods

### NewIndirectCallSitesOutputBody

`func NewIndirectCallSitesOutputBody(functionId int64, sites []IndirectCallSite, ) *IndirectCallSitesOutputBody`

NewIndirectCallSitesOutputBody instantiates a new IndirectCallSitesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndirectCallSitesOutputBodyWithDefaults

`func NewIndirectCallSitesOutputBodyWithDefaults() *IndirectCallSitesOutputBody`

NewIndirectCallSitesOutputBodyWithDefaults instantiates a new IndirectCallSitesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *IndirectCallSitesOutputBody) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *IndirectCallSitesOutputBody) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *IndirectCallSitesOutputBody) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetSites

`func (o *IndirectCallSitesOutputBody) GetSites() []IndirectCallSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *IndirectCallSitesOutputBody) GetSitesOk() (*[]IndirectCallSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *IndirectCallSitesOutputBody) SetSites(v []IndirectCallSite)`

SetSites sets Sites field to given value.


### SetSitesNil

`func (o *IndirectCallSitesOutputBody) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *IndirectCallSitesOutputBody) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


