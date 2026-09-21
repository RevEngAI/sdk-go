# FunctionSimilarity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Similarity** | **float32** |  | 
**Subject** | [**Subject**](Subject.md) |  | 

## Methods

### NewFunctionSimilarity

`func NewFunctionSimilarity(name string, similarity float32, subject Subject, ) *FunctionSimilarity`

NewFunctionSimilarity instantiates a new FunctionSimilarity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionSimilarityWithDefaults

`func NewFunctionSimilarityWithDefaults() *FunctionSimilarity`

NewFunctionSimilarityWithDefaults instantiates a new FunctionSimilarity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionSimilarity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionSimilarity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionSimilarity) SetName(v string)`

SetName sets Name field to given value.


### GetSimilarity

`func (o *FunctionSimilarity) GetSimilarity() float32`

GetSimilarity returns the Similarity field if non-nil, zero value otherwise.

### GetSimilarityOk

`func (o *FunctionSimilarity) GetSimilarityOk() (*float32, bool)`

GetSimilarityOk returns a tuple with the Similarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarity

`func (o *FunctionSimilarity) SetSimilarity(v float32)`

SetSimilarity sets Similarity field to given value.


### GetSubject

`func (o *FunctionSimilarity) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *FunctionSimilarity) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *FunctionSimilarity) SetSubject(v Subject)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


