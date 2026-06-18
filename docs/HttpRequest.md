# HttpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesReceived** | Pointer to **int64** |  | [optional] 
**BytesSent** | Pointer to **int64** |  | [optional] 
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 
**ExtraHeaders** | Pointer to **[]string** |  | [optional] 
**Flags** | Pointer to **int64** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**PcapStreamId** | Pointer to **int64** |  | [optional] 
**PostData** | Pointer to **string** |  | [optional] 
**Proxy** | Pointer to **string** |  | [optional] 
**ProxyBypass** | Pointer to **string** |  | [optional] 
**Referer** | Pointer to **string** |  | [optional] 
**RequestBody** | Pointer to [**PcapBodyInfo**](PcapBodyInfo.md) |  | [optional] 
**ResponseBody** | Pointer to [**PcapBodyInfo**](PcapBodyInfo.md) |  | [optional] 
**ResponseStatus** | Pointer to **int64** |  | [optional] 
**ServerIp** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**ServerPort** | Pointer to **int64** |  | [optional] 
**Service** | Pointer to **int64** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Verb** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpRequest

`func NewHttpRequest() *HttpRequest`

NewHttpRequest instantiates a new HttpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpRequestWithDefaults

`func NewHttpRequestWithDefaults() *HttpRequest`

NewHttpRequestWithDefaults instantiates a new HttpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesReceived

`func (o *HttpRequest) GetBytesReceived() int64`

GetBytesReceived returns the BytesReceived field if non-nil, zero value otherwise.

### GetBytesReceivedOk

`func (o *HttpRequest) GetBytesReceivedOk() (*int64, bool)`

GetBytesReceivedOk returns a tuple with the BytesReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesReceived

`func (o *HttpRequest) SetBytesReceived(v int64)`

SetBytesReceived sets BytesReceived field to given value.

### HasBytesReceived

`func (o *HttpRequest) HasBytesReceived() bool`

HasBytesReceived returns a boolean if a field has been set.

### GetBytesSent

`func (o *HttpRequest) GetBytesSent() int64`

GetBytesSent returns the BytesSent field if non-nil, zero value otherwise.

### GetBytesSentOk

`func (o *HttpRequest) GetBytesSentOk() (*int64, bool)`

GetBytesSentOk returns a tuple with the BytesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesSent

`func (o *HttpRequest) SetBytesSent(v int64)`

SetBytesSent sets BytesSent field to given value.

### HasBytesSent

`func (o *HttpRequest) HasBytesSent() bool`

HasBytesSent returns a boolean if a field has been set.

### GetEvents

`func (o *HttpRequest) GetEvents() []ReportEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *HttpRequest) GetEventsOk() (*[]ReportEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *HttpRequest) SetEvents(v []ReportEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *HttpRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *HttpRequest) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *HttpRequest) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetExtraHeaders

`func (o *HttpRequest) GetExtraHeaders() []string`

GetExtraHeaders returns the ExtraHeaders field if non-nil, zero value otherwise.

### GetExtraHeadersOk

`func (o *HttpRequest) GetExtraHeadersOk() (*[]string, bool)`

GetExtraHeadersOk returns a tuple with the ExtraHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeaders

`func (o *HttpRequest) SetExtraHeaders(v []string)`

SetExtraHeaders sets ExtraHeaders field to given value.

### HasExtraHeaders

`func (o *HttpRequest) HasExtraHeaders() bool`

HasExtraHeaders returns a boolean if a field has been set.

### SetExtraHeadersNil

`func (o *HttpRequest) SetExtraHeadersNil(b bool)`

 SetExtraHeadersNil sets the value for ExtraHeaders to be an explicit nil

### UnsetExtraHeaders
`func (o *HttpRequest) UnsetExtraHeaders()`

UnsetExtraHeaders ensures that no value is present for ExtraHeaders, not even an explicit nil
### GetFlags

`func (o *HttpRequest) GetFlags() int64`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *HttpRequest) GetFlagsOk() (*int64, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *HttpRequest) SetFlags(v int64)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *HttpRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetPassword

`func (o *HttpRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HttpRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HttpRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HttpRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *HttpRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HttpRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HttpRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HttpRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPcapStreamId

`func (o *HttpRequest) GetPcapStreamId() int64`

GetPcapStreamId returns the PcapStreamId field if non-nil, zero value otherwise.

### GetPcapStreamIdOk

`func (o *HttpRequest) GetPcapStreamIdOk() (*int64, bool)`

GetPcapStreamIdOk returns a tuple with the PcapStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapStreamId

`func (o *HttpRequest) SetPcapStreamId(v int64)`

SetPcapStreamId sets PcapStreamId field to given value.

### HasPcapStreamId

`func (o *HttpRequest) HasPcapStreamId() bool`

HasPcapStreamId returns a boolean if a field has been set.

### GetPostData

`func (o *HttpRequest) GetPostData() string`

GetPostData returns the PostData field if non-nil, zero value otherwise.

### GetPostDataOk

`func (o *HttpRequest) GetPostDataOk() (*string, bool)`

GetPostDataOk returns a tuple with the PostData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostData

`func (o *HttpRequest) SetPostData(v string)`

SetPostData sets PostData field to given value.

### HasPostData

`func (o *HttpRequest) HasPostData() bool`

HasPostData returns a boolean if a field has been set.

### GetProxy

`func (o *HttpRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *HttpRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *HttpRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *HttpRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetProxyBypass

`func (o *HttpRequest) GetProxyBypass() string`

GetProxyBypass returns the ProxyBypass field if non-nil, zero value otherwise.

### GetProxyBypassOk

`func (o *HttpRequest) GetProxyBypassOk() (*string, bool)`

GetProxyBypassOk returns a tuple with the ProxyBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyBypass

`func (o *HttpRequest) SetProxyBypass(v string)`

SetProxyBypass sets ProxyBypass field to given value.

### HasProxyBypass

`func (o *HttpRequest) HasProxyBypass() bool`

HasProxyBypass returns a boolean if a field has been set.

### GetReferer

`func (o *HttpRequest) GetReferer() string`

GetReferer returns the Referer field if non-nil, zero value otherwise.

### GetRefererOk

`func (o *HttpRequest) GetRefererOk() (*string, bool)`

GetRefererOk returns a tuple with the Referer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferer

`func (o *HttpRequest) SetReferer(v string)`

SetReferer sets Referer field to given value.

### HasReferer

`func (o *HttpRequest) HasReferer() bool`

HasReferer returns a boolean if a field has been set.

### GetRequestBody

`func (o *HttpRequest) GetRequestBody() PcapBodyInfo`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *HttpRequest) GetRequestBodyOk() (*PcapBodyInfo, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *HttpRequest) SetRequestBody(v PcapBodyInfo)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *HttpRequest) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetResponseBody

`func (o *HttpRequest) GetResponseBody() PcapBodyInfo`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *HttpRequest) GetResponseBodyOk() (*PcapBodyInfo, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *HttpRequest) SetResponseBody(v PcapBodyInfo)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *HttpRequest) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseStatus

`func (o *HttpRequest) GetResponseStatus() int64`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *HttpRequest) GetResponseStatusOk() (*int64, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *HttpRequest) SetResponseStatus(v int64)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *HttpRequest) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### GetServerIp

`func (o *HttpRequest) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *HttpRequest) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *HttpRequest) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *HttpRequest) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetServerName

`func (o *HttpRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *HttpRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *HttpRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *HttpRequest) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerPort

`func (o *HttpRequest) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *HttpRequest) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *HttpRequest) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *HttpRequest) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetService

`func (o *HttpRequest) GetService() int64`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *HttpRequest) GetServiceOk() (*int64, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *HttpRequest) SetService(v int64)`

SetService sets Service field to given value.

### HasService

`func (o *HttpRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### GetUserAgent

`func (o *HttpRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *HttpRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *HttpRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *HttpRequest) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUsername

`func (o *HttpRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HttpRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HttpRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HttpRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVerb

`func (o *HttpRequest) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *HttpRequest) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *HttpRequest) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *HttpRequest) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### GetVersion

`func (o *HttpRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HttpRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HttpRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HttpRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


