# SecurityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aslr** | **bool** |  | 
**Dep** | **bool** |  | 
**Cfg** | **bool** |  | 
**DriverModel** | **bool** |  | 
**AppContainer** | **bool** |  | 
**TerminalServerAware** | **bool** |  | 
**ImageIsolation** | **bool** |  | 
**CodeIntegrity** | **bool** |  | 
**HighEntropy** | **bool** |  | 
**Seh** | **bool** |  | 
**BoundImage** | **bool** |  | 

## Methods

### NewSecurityModel

`func NewSecurityModel(aslr bool, dep bool, cfg bool, driverModel bool, appContainer bool, terminalServerAware bool, imageIsolation bool, codeIntegrity bool, highEntropy bool, seh bool, boundImage bool, ) *SecurityModel`

NewSecurityModel instantiates a new SecurityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityModelWithDefaults

`func NewSecurityModelWithDefaults() *SecurityModel`

NewSecurityModelWithDefaults instantiates a new SecurityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAslr

`func (o *SecurityModel) GetAslr() bool`

GetAslr returns the Aslr field if non-nil, zero value otherwise.

### GetAslrOk

`func (o *SecurityModel) GetAslrOk() (*bool, bool)`

GetAslrOk returns a tuple with the Aslr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAslr

`func (o *SecurityModel) SetAslr(v bool)`

SetAslr sets Aslr field to given value.


### GetDep

`func (o *SecurityModel) GetDep() bool`

GetDep returns the Dep field if non-nil, zero value otherwise.

### GetDepOk

`func (o *SecurityModel) GetDepOk() (*bool, bool)`

GetDepOk returns a tuple with the Dep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDep

`func (o *SecurityModel) SetDep(v bool)`

SetDep sets Dep field to given value.


### GetCfg

`func (o *SecurityModel) GetCfg() bool`

GetCfg returns the Cfg field if non-nil, zero value otherwise.

### GetCfgOk

`func (o *SecurityModel) GetCfgOk() (*bool, bool)`

GetCfgOk returns a tuple with the Cfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfg

`func (o *SecurityModel) SetCfg(v bool)`

SetCfg sets Cfg field to given value.


### GetDriverModel

`func (o *SecurityModel) GetDriverModel() bool`

GetDriverModel returns the DriverModel field if non-nil, zero value otherwise.

### GetDriverModelOk

`func (o *SecurityModel) GetDriverModelOk() (*bool, bool)`

GetDriverModelOk returns a tuple with the DriverModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverModel

`func (o *SecurityModel) SetDriverModel(v bool)`

SetDriverModel sets DriverModel field to given value.


### GetAppContainer

`func (o *SecurityModel) GetAppContainer() bool`

GetAppContainer returns the AppContainer field if non-nil, zero value otherwise.

### GetAppContainerOk

`func (o *SecurityModel) GetAppContainerOk() (*bool, bool)`

GetAppContainerOk returns a tuple with the AppContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppContainer

`func (o *SecurityModel) SetAppContainer(v bool)`

SetAppContainer sets AppContainer field to given value.


### GetTerminalServerAware

`func (o *SecurityModel) GetTerminalServerAware() bool`

GetTerminalServerAware returns the TerminalServerAware field if non-nil, zero value otherwise.

### GetTerminalServerAwareOk

`func (o *SecurityModel) GetTerminalServerAwareOk() (*bool, bool)`

GetTerminalServerAwareOk returns a tuple with the TerminalServerAware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalServerAware

`func (o *SecurityModel) SetTerminalServerAware(v bool)`

SetTerminalServerAware sets TerminalServerAware field to given value.


### GetImageIsolation

`func (o *SecurityModel) GetImageIsolation() bool`

GetImageIsolation returns the ImageIsolation field if non-nil, zero value otherwise.

### GetImageIsolationOk

`func (o *SecurityModel) GetImageIsolationOk() (*bool, bool)`

GetImageIsolationOk returns a tuple with the ImageIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageIsolation

`func (o *SecurityModel) SetImageIsolation(v bool)`

SetImageIsolation sets ImageIsolation field to given value.


### GetCodeIntegrity

`func (o *SecurityModel) GetCodeIntegrity() bool`

GetCodeIntegrity returns the CodeIntegrity field if non-nil, zero value otherwise.

### GetCodeIntegrityOk

`func (o *SecurityModel) GetCodeIntegrityOk() (*bool, bool)`

GetCodeIntegrityOk returns a tuple with the CodeIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrity

`func (o *SecurityModel) SetCodeIntegrity(v bool)`

SetCodeIntegrity sets CodeIntegrity field to given value.


### GetHighEntropy

`func (o *SecurityModel) GetHighEntropy() bool`

GetHighEntropy returns the HighEntropy field if non-nil, zero value otherwise.

### GetHighEntropyOk

`func (o *SecurityModel) GetHighEntropyOk() (*bool, bool)`

GetHighEntropyOk returns a tuple with the HighEntropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighEntropy

`func (o *SecurityModel) SetHighEntropy(v bool)`

SetHighEntropy sets HighEntropy field to given value.


### GetSeh

`func (o *SecurityModel) GetSeh() bool`

GetSeh returns the Seh field if non-nil, zero value otherwise.

### GetSehOk

`func (o *SecurityModel) GetSehOk() (*bool, bool)`

GetSehOk returns a tuple with the Seh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeh

`func (o *SecurityModel) SetSeh(v bool)`

SetSeh sets Seh field to given value.


### GetBoundImage

`func (o *SecurityModel) GetBoundImage() bool`

GetBoundImage returns the BoundImage field if non-nil, zero value otherwise.

### GetBoundImageOk

`func (o *SecurityModel) GetBoundImageOk() (*bool, bool)`

GetBoundImageOk returns a tuple with the BoundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundImage

`func (o *SecurityModel) SetBoundImage(v bool)`

SetBoundImage sets BoundImage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


