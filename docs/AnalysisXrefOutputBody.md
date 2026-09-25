# AnalysisXrefOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XrefFromList** | [**[]XrefFromBody**](XrefFromBody.md) | Xrefs that originate at the queried vaddr, one per target address | 
**XrefToList** | [**[]XrefIntoBody**](XrefIntoBody.md) | Xrefs that target the queried vaddr, one per referencing address | 

## Methods

### NewAnalysisXrefOutputBody

`func NewAnalysisXrefOutputBody(xrefFromList []XrefFromBody, xrefToList []XrefIntoBody, ) *AnalysisXrefOutputBody`

NewAnalysisXrefOutputBody instantiates a new AnalysisXrefOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisXrefOutputBodyWithDefaults

`func NewAnalysisXrefOutputBodyWithDefaults() *AnalysisXrefOutputBody`

NewAnalysisXrefOutputBodyWithDefaults instantiates a new AnalysisXrefOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXrefFromList

`func (o *AnalysisXrefOutputBody) GetXrefFromList() []XrefFromBody`

GetXrefFromList returns the XrefFromList field if non-nil, zero value otherwise.

### GetXrefFromListOk

`func (o *AnalysisXrefOutputBody) GetXrefFromListOk() (*[]XrefFromBody, bool)`

GetXrefFromListOk returns a tuple with the XrefFromList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXrefFromList

`func (o *AnalysisXrefOutputBody) SetXrefFromList(v []XrefFromBody)`

SetXrefFromList sets XrefFromList field to given value.


### SetXrefFromListNil

`func (o *AnalysisXrefOutputBody) SetXrefFromListNil(b bool)`

 SetXrefFromListNil sets the value for XrefFromList to be an explicit nil

### UnsetXrefFromList
`func (o *AnalysisXrefOutputBody) UnsetXrefFromList()`

UnsetXrefFromList ensures that no value is present for XrefFromList, not even an explicit nil
### GetXrefToList

`func (o *AnalysisXrefOutputBody) GetXrefToList() []XrefIntoBody`

GetXrefToList returns the XrefToList field if non-nil, zero value otherwise.

### GetXrefToListOk

`func (o *AnalysisXrefOutputBody) GetXrefToListOk() (*[]XrefIntoBody, bool)`

GetXrefToListOk returns a tuple with the XrefToList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXrefToList

`func (o *AnalysisXrefOutputBody) SetXrefToList(v []XrefIntoBody)`

SetXrefToList sets XrefToList field to given value.


### SetXrefToListNil

`func (o *AnalysisXrefOutputBody) SetXrefToListNil(b bool)`

 SetXrefToListNil sets the value for XrefToList to be an explicit nil

### UnsetXrefToList
`func (o *AnalysisXrefOutputBody) UnsetXrefToList()`

UnsetXrefToList ensures that no value is present for XrefToList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


