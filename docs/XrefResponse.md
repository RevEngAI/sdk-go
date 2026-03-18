# XrefResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XrefToList** | [**[]XrefToResponse**](XrefToResponse.md) |  | 
**XrefFromList** | [**[]XrefFromResponse**](XrefFromResponse.md) |  | 

## Methods

### NewXrefResponse

`func NewXrefResponse(xrefToList []XrefToResponse, xrefFromList []XrefFromResponse, ) *XrefResponse`

NewXrefResponse instantiates a new XrefResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXrefResponseWithDefaults

`func NewXrefResponseWithDefaults() *XrefResponse`

NewXrefResponseWithDefaults instantiates a new XrefResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXrefToList

`func (o *XrefResponse) GetXrefToList() []XrefToResponse`

GetXrefToList returns the XrefToList field if non-nil, zero value otherwise.

### GetXrefToListOk

`func (o *XrefResponse) GetXrefToListOk() (*[]XrefToResponse, bool)`

GetXrefToListOk returns a tuple with the XrefToList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXrefToList

`func (o *XrefResponse) SetXrefToList(v []XrefToResponse)`

SetXrefToList sets XrefToList field to given value.


### GetXrefFromList

`func (o *XrefResponse) GetXrefFromList() []XrefFromResponse`

GetXrefFromList returns the XrefFromList field if non-nil, zero value otherwise.

### GetXrefFromListOk

`func (o *XrefResponse) GetXrefFromListOk() (*[]XrefFromResponse, bool)`

GetXrefFromListOk returns a tuple with the XrefFromList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXrefFromList

`func (o *XrefResponse) SetXrefFromList(v []XrefFromResponse)`

SetXrefFromList sets XrefFromList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


