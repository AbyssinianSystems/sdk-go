# AuthenticationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Reason** | **string** |  | 
**RejectionReasons** | [**[]RejectionReason**](RejectionReason.md) |  | 
**CorrelationId** | Pointer to **string** | Request correlation identifier echoed in the &#x60;X-Correlation-ID&#x60; header and JSON error payloads. | [optional] 

## Methods

### NewAuthenticationError

`func NewAuthenticationError(status string, reason string, rejectionReasons []RejectionReason, ) *AuthenticationError`

NewAuthenticationError instantiates a new AuthenticationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationErrorWithDefaults

`func NewAuthenticationErrorWithDefaults() *AuthenticationError`

NewAuthenticationErrorWithDefaults instantiates a new AuthenticationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AuthenticationError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthenticationError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthenticationError) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *AuthenticationError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AuthenticationError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AuthenticationError) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRejectionReasons

`func (o *AuthenticationError) GetRejectionReasons() []RejectionReason`

GetRejectionReasons returns the RejectionReasons field if non-nil, zero value otherwise.

### GetRejectionReasonsOk

`func (o *AuthenticationError) GetRejectionReasonsOk() (*[]RejectionReason, bool)`

GetRejectionReasonsOk returns a tuple with the RejectionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReasons

`func (o *AuthenticationError) SetRejectionReasons(v []RejectionReason)`

SetRejectionReasons sets RejectionReasons field to given value.


### GetCorrelationId

`func (o *AuthenticationError) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *AuthenticationError) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *AuthenticationError) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *AuthenticationError) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


