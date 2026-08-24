# GetAnalysisLogsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**[]AnalysisLogEntry**](AnalysisLogEntry.md) | Analysis log lines, oldest first | 

## Methods

### NewGetAnalysisLogsOutputBody

`func NewGetAnalysisLogsOutputBody(entries []AnalysisLogEntry, ) *GetAnalysisLogsOutputBody`

NewGetAnalysisLogsOutputBody instantiates a new GetAnalysisLogsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalysisLogsOutputBodyWithDefaults

`func NewGetAnalysisLogsOutputBodyWithDefaults() *GetAnalysisLogsOutputBody`

NewGetAnalysisLogsOutputBodyWithDefaults instantiates a new GetAnalysisLogsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *GetAnalysisLogsOutputBody) GetEntries() []AnalysisLogEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *GetAnalysisLogsOutputBody) GetEntriesOk() (*[]AnalysisLogEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *GetAnalysisLogsOutputBody) SetEntries(v []AnalysisLogEntry)`

SetEntries sets Entries field to given value.


### SetEntriesNil

`func (o *GetAnalysisLogsOutputBody) SetEntriesNil(b bool)`

 SetEntriesNil sets the value for Entries to be an explicit nil

### UnsetEntries
`func (o *GetAnalysisLogsOutputBody) UnsetEntries()`

UnsetEntries ensures that no value is present for Entries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


