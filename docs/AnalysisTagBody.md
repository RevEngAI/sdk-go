# AnalysisTagBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **NullableInt64** | Collection this tag maps to, or null | 
**Name** | **string** | Tag name | 
**Origin** | **string** | Origin of the tag | 

## Methods

### NewAnalysisTagBody

`func NewAnalysisTagBody(collectionId NullableInt64, name string, origin string, ) *AnalysisTagBody`

NewAnalysisTagBody instantiates a new AnalysisTagBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisTagBodyWithDefaults

`func NewAnalysisTagBodyWithDefaults() *AnalysisTagBody`

NewAnalysisTagBodyWithDefaults instantiates a new AnalysisTagBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *AnalysisTagBody) GetCollectionId() int64`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *AnalysisTagBody) GetCollectionIdOk() (*int64, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *AnalysisTagBody) SetCollectionId(v int64)`

SetCollectionId sets CollectionId field to given value.


### SetCollectionIdNil

`func (o *AnalysisTagBody) SetCollectionIdNil(b bool)`

 SetCollectionIdNil sets the value for CollectionId to be an explicit nil

### UnsetCollectionId
`func (o *AnalysisTagBody) UnsetCollectionId()`

UnsetCollectionId ensures that no value is present for CollectionId, not even an explicit nil
### GetName

`func (o *AnalysisTagBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalysisTagBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnalysisTagBody) SetName(v string)`

SetName sets Name field to given value.


### GetOrigin

`func (o *AnalysisTagBody) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *AnalysisTagBody) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *AnalysisTagBody) SetOrigin(v string)`

SetOrigin sets Origin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


