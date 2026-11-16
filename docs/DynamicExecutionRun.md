# DynamicExecutionRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**DynamicExecutionRunConfig**](DynamicExecutionRunConfig.md) | Parameters the run was started with | 
**CreatedAt** | **time.Time** | When the run was started | 
**Id** | **string** | Run ID | 
**Status** | **string** | Run status | 

## Methods

### NewDynamicExecutionRun

`func NewDynamicExecutionRun(config DynamicExecutionRunConfig, createdAt time.Time, id string, status string, ) *DynamicExecutionRun`

NewDynamicExecutionRun instantiates a new DynamicExecutionRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicExecutionRunWithDefaults

`func NewDynamicExecutionRunWithDefaults() *DynamicExecutionRun`

NewDynamicExecutionRunWithDefaults instantiates a new DynamicExecutionRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *DynamicExecutionRun) GetConfig() DynamicExecutionRunConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DynamicExecutionRun) GetConfigOk() (*DynamicExecutionRunConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DynamicExecutionRun) SetConfig(v DynamicExecutionRunConfig)`

SetConfig sets Config field to given value.


### GetCreatedAt

`func (o *DynamicExecutionRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DynamicExecutionRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DynamicExecutionRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *DynamicExecutionRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DynamicExecutionRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DynamicExecutionRun) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *DynamicExecutionRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DynamicExecutionRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DynamicExecutionRun) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


