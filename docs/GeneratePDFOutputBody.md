# GeneratePDFOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyRunning** | Pointer to **bool** | True when an existing PDF generation is in progress for this analysis and user | [optional] 
**TaskId** | **string** | Workflow task ID — use to poll status and download the PDF | 

## Methods

### NewGeneratePDFOutputBody

`func NewGeneratePDFOutputBody(taskId string, ) *GeneratePDFOutputBody`

NewGeneratePDFOutputBody instantiates a new GeneratePDFOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratePDFOutputBodyWithDefaults

`func NewGeneratePDFOutputBodyWithDefaults() *GeneratePDFOutputBody`

NewGeneratePDFOutputBodyWithDefaults instantiates a new GeneratePDFOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyRunning

`func (o *GeneratePDFOutputBody) GetAlreadyRunning() bool`

GetAlreadyRunning returns the AlreadyRunning field if non-nil, zero value otherwise.

### GetAlreadyRunningOk

`func (o *GeneratePDFOutputBody) GetAlreadyRunningOk() (*bool, bool)`

GetAlreadyRunningOk returns a tuple with the AlreadyRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyRunning

`func (o *GeneratePDFOutputBody) SetAlreadyRunning(v bool)`

SetAlreadyRunning sets AlreadyRunning field to given value.

### HasAlreadyRunning

`func (o *GeneratePDFOutputBody) HasAlreadyRunning() bool`

HasAlreadyRunning returns a boolean if a field has been set.

### GetTaskId

`func (o *GeneratePDFOutputBody) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *GeneratePDFOutputBody) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *GeneratePDFOutputBody) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


