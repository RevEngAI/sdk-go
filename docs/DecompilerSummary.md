# DecompilerSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | [**Subject**](Subject.md) |  | 
**Summary** | **string** |  | 
**Reachability** | Pointer to [**NullableCallChain**](CallChain.md) |  | [optional] 

## Methods

### NewDecompilerSummary

`func NewDecompilerSummary(subject Subject, summary string, ) *DecompilerSummary`

NewDecompilerSummary instantiates a new DecompilerSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecompilerSummaryWithDefaults

`func NewDecompilerSummaryWithDefaults() *DecompilerSummary`

NewDecompilerSummaryWithDefaults instantiates a new DecompilerSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *DecompilerSummary) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *DecompilerSummary) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *DecompilerSummary) SetSubject(v Subject)`

SetSubject sets Subject field to given value.


### GetSummary

`func (o *DecompilerSummary) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DecompilerSummary) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DecompilerSummary) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetReachability

`func (o *DecompilerSummary) GetReachability() CallChain`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *DecompilerSummary) GetReachabilityOk() (*CallChain, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *DecompilerSummary) SetReachability(v CallChain)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *DecompilerSummary) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### SetReachabilityNil

`func (o *DecompilerSummary) SetReachabilityNil(b bool)`

 SetReachabilityNil sets the value for Reachability to be an explicit nil

### UnsetReachability
`func (o *DecompilerSummary) UnsetReachability()`

UnsetReachability ensures that no value is present for Reachability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


