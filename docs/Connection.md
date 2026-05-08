# Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 
**LocalIp** | **string** |  | 
**LocalPort** | **interface{}** |  | 
**Protocol** | **string** |  | 
**RemoteIp** | **string** |  | 
**RemotePort** | **interface{}** |  | 

## Methods

### NewConnection

`func NewConnection(localIp string, localPort interface{}, protocol string, remoteIp string, remotePort interface{}, ) *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *Connection) GetEvents() []ReportEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Connection) GetEventsOk() (*[]ReportEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Connection) SetEvents(v []ReportEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Connection) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *Connection) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *Connection) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetLocalIp

`func (o *Connection) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *Connection) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *Connection) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.


### GetLocalPort

`func (o *Connection) GetLocalPort() interface{}`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *Connection) GetLocalPortOk() (*interface{}, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *Connection) SetLocalPort(v interface{})`

SetLocalPort sets LocalPort field to given value.


### SetLocalPortNil

`func (o *Connection) SetLocalPortNil(b bool)`

 SetLocalPortNil sets the value for LocalPort to be an explicit nil

### UnsetLocalPort
`func (o *Connection) UnsetLocalPort()`

UnsetLocalPort ensures that no value is present for LocalPort, not even an explicit nil
### GetProtocol

`func (o *Connection) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Connection) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Connection) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetRemoteIp

`func (o *Connection) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *Connection) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *Connection) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.


### GetRemotePort

`func (o *Connection) GetRemotePort() interface{}`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *Connection) GetRemotePortOk() (*interface{}, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *Connection) SetRemotePort(v interface{})`

SetRemotePort sets RemotePort field to given value.


### SetRemotePortNil

`func (o *Connection) SetRemotePortNil(b bool)`

 SetRemotePortNil sets the value for RemotePort to be an explicit nil

### UnsetRemotePort
`func (o *Connection) UnsetRemotePort()`

UnsetRemotePort ensures that no value is present for RemotePort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


