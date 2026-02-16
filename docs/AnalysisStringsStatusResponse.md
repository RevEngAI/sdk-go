# AnalysisStringsStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**BinariesTaskStatus**](BinariesTaskStatus.md) | The current status of the strings extraction task | 

## Methods

### NewAnalysisStringsStatusResponse

`func NewAnalysisStringsStatusResponse(status BinariesTaskStatus, ) *AnalysisStringsStatusResponse`

NewAnalysisStringsStatusResponse instantiates a new AnalysisStringsStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisStringsStatusResponseWithDefaults

`func NewAnalysisStringsStatusResponseWithDefaults() *AnalysisStringsStatusResponse`

NewAnalysisStringsStatusResponseWithDefaults instantiates a new AnalysisStringsStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AnalysisStringsStatusResponse) GetStatus() BinariesTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalysisStringsStatusResponse) GetStatusOk() (*BinariesTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalysisStringsStatusResponse) SetStatus(v BinariesTaskStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


