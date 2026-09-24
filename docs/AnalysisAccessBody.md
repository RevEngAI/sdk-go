# AnalysisAccessBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **bool** | True when the caller owns this analysis | 
**Username** | **string** | Username of the analysis owner | 

## Methods

### NewAnalysisAccessBody

`func NewAnalysisAccessBody(owner bool, username string, ) *AnalysisAccessBody`

NewAnalysisAccessBody instantiates a new AnalysisAccessBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisAccessBodyWithDefaults

`func NewAnalysisAccessBodyWithDefaults() *AnalysisAccessBody`

NewAnalysisAccessBodyWithDefaults instantiates a new AnalysisAccessBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *AnalysisAccessBody) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AnalysisAccessBody) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AnalysisAccessBody) SetOwner(v bool)`

SetOwner sets Owner field to given value.


### GetUsername

`func (o *AnalysisAccessBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AnalysisAccessBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AnalysisAccessBody) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


