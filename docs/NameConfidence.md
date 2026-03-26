# NameConfidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | **float32** | Confidence score as a percentage | 
**Name** | **string** | The suggested function name | 

## Methods

### NewNameConfidence

`func NewNameConfidence(confidence float32, name string, ) *NameConfidence`

NewNameConfidence instantiates a new NameConfidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameConfidenceWithDefaults

`func NewNameConfidenceWithDefaults() *NameConfidence`

NewNameConfidenceWithDefaults instantiates a new NameConfidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *NameConfidence) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *NameConfidence) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *NameConfidence) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.


### GetName

`func (o *NameConfidence) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NameConfidence) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NameConfidence) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


