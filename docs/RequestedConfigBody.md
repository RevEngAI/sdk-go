# RequestedConfigBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | **bool** | Whether the capabilities agent was requested | 
**Functions** | **bool** | Whether the functions pipeline was requested | 
**Sandbox** | **bool** | Whether dynamic execution (sandbox) was requested | 
**Scrape** | **bool** | Whether external-source scraping was requested | 
**Triage** | **bool** | Whether the triage agent was requested | 

## Methods

### NewRequestedConfigBody

`func NewRequestedConfigBody(capabilities bool, functions bool, sandbox bool, scrape bool, triage bool, ) *RequestedConfigBody`

NewRequestedConfigBody instantiates a new RequestedConfigBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedConfigBodyWithDefaults

`func NewRequestedConfigBodyWithDefaults() *RequestedConfigBody`

NewRequestedConfigBodyWithDefaults instantiates a new RequestedConfigBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *RequestedConfigBody) GetCapabilities() bool`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *RequestedConfigBody) GetCapabilitiesOk() (*bool, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *RequestedConfigBody) SetCapabilities(v bool)`

SetCapabilities sets Capabilities field to given value.


### GetFunctions

`func (o *RequestedConfigBody) GetFunctions() bool`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *RequestedConfigBody) GetFunctionsOk() (*bool, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *RequestedConfigBody) SetFunctions(v bool)`

SetFunctions sets Functions field to given value.


### GetSandbox

`func (o *RequestedConfigBody) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *RequestedConfigBody) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *RequestedConfigBody) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.


### GetScrape

`func (o *RequestedConfigBody) GetScrape() bool`

GetScrape returns the Scrape field if non-nil, zero value otherwise.

### GetScrapeOk

`func (o *RequestedConfigBody) GetScrapeOk() (*bool, bool)`

GetScrapeOk returns a tuple with the Scrape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrape

`func (o *RequestedConfigBody) SetScrape(v bool)`

SetScrape sets Scrape field to given value.


### GetTriage

`func (o *RequestedConfigBody) GetTriage() bool`

GetTriage returns the Triage field if non-nil, zero value otherwise.

### GetTriageOk

`func (o *RequestedConfigBody) GetTriageOk() (*bool, bool)`

GetTriageOk returns a tuple with the Triage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriage

`func (o *RequestedConfigBody) SetTriage(v bool)`

SetTriage sets Triage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


