# HealthStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Store** | Pointer to **string** |  | [optional] 
**CorrelationId** | Pointer to **string** | Request correlation identifier echoed in the &#x60;X-Correlation-ID&#x60; header and JSON error payloads. | [optional] 

## Methods

### NewHealthStatus

`func NewHealthStatus(status string, ) *HealthStatus`

NewHealthStatus instantiates a new HealthStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthStatusWithDefaults

`func NewHealthStatusWithDefaults() *HealthStatus`

NewHealthStatusWithDefaults instantiates a new HealthStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStore

`func (o *HealthStatus) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *HealthStatus) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *HealthStatus) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *HealthStatus) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetCorrelationId

`func (o *HealthStatus) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *HealthStatus) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *HealthStatus) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *HealthStatus) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


