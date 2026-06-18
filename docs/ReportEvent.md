# ReportEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiCalls** | Pointer to [**[]ApiCall**](ApiCall.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**DesiredAccess** | Pointer to **[]string** |  | [optional] 
**ProcessSeqid** | Pointer to **int64** |  | [optional] 
**TotalBytes** | Pointer to **int64** |  | [optional] 
**Type** | **string** |  | 
**Value** | Pointer to **string** |  | [optional] 
**ValueName** | Pointer to **string** |  | [optional] 

## Methods

### NewReportEvent

`func NewReportEvent(type_ string, ) *ReportEvent`

NewReportEvent instantiates a new ReportEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportEventWithDefaults

`func NewReportEventWithDefaults() *ReportEvent`

NewReportEventWithDefaults instantiates a new ReportEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiCalls

`func (o *ReportEvent) GetApiCalls() []ApiCall`

GetApiCalls returns the ApiCalls field if non-nil, zero value otherwise.

### GetApiCallsOk

`func (o *ReportEvent) GetApiCallsOk() (*[]ApiCall, bool)`

GetApiCallsOk returns a tuple with the ApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCalls

`func (o *ReportEvent) SetApiCalls(v []ApiCall)`

SetApiCalls sets ApiCalls field to given value.

### HasApiCalls

`func (o *ReportEvent) HasApiCalls() bool`

HasApiCalls returns a boolean if a field has been set.

### SetApiCallsNil

`func (o *ReportEvent) SetApiCallsNil(b bool)`

 SetApiCallsNil sets the value for ApiCalls to be an explicit nil

### UnsetApiCalls
`func (o *ReportEvent) UnsetApiCalls()`

UnsetApiCalls ensures that no value is present for ApiCalls, not even an explicit nil
### GetCount

`func (o *ReportEvent) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ReportEvent) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ReportEvent) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ReportEvent) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDesiredAccess

`func (o *ReportEvent) GetDesiredAccess() []string`

GetDesiredAccess returns the DesiredAccess field if non-nil, zero value otherwise.

### GetDesiredAccessOk

`func (o *ReportEvent) GetDesiredAccessOk() (*[]string, bool)`

GetDesiredAccessOk returns a tuple with the DesiredAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredAccess

`func (o *ReportEvent) SetDesiredAccess(v []string)`

SetDesiredAccess sets DesiredAccess field to given value.

### HasDesiredAccess

`func (o *ReportEvent) HasDesiredAccess() bool`

HasDesiredAccess returns a boolean if a field has been set.

### SetDesiredAccessNil

`func (o *ReportEvent) SetDesiredAccessNil(b bool)`

 SetDesiredAccessNil sets the value for DesiredAccess to be an explicit nil

### UnsetDesiredAccess
`func (o *ReportEvent) UnsetDesiredAccess()`

UnsetDesiredAccess ensures that no value is present for DesiredAccess, not even an explicit nil
### GetProcessSeqid

`func (o *ReportEvent) GetProcessSeqid() int64`

GetProcessSeqid returns the ProcessSeqid field if non-nil, zero value otherwise.

### GetProcessSeqidOk

`func (o *ReportEvent) GetProcessSeqidOk() (*int64, bool)`

GetProcessSeqidOk returns a tuple with the ProcessSeqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessSeqid

`func (o *ReportEvent) SetProcessSeqid(v int64)`

SetProcessSeqid sets ProcessSeqid field to given value.

### HasProcessSeqid

`func (o *ReportEvent) HasProcessSeqid() bool`

HasProcessSeqid returns a boolean if a field has been set.

### GetTotalBytes

`func (o *ReportEvent) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *ReportEvent) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *ReportEvent) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *ReportEvent) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetType

`func (o *ReportEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportEvent) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ReportEvent) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReportEvent) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReportEvent) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ReportEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueName

`func (o *ReportEvent) GetValueName() string`

GetValueName returns the ValueName field if non-nil, zero value otherwise.

### GetValueNameOk

`func (o *ReportEvent) GetValueNameOk() (*string, bool)`

GetValueNameOk returns a tuple with the ValueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueName

`func (o *ReportEvent) SetValueName(v string)`

SetValueName sets ValueName field to given value.

### HasValueName

`func (o *ReportEvent) HasValueName() bool`

HasValueName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


