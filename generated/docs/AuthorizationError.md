# AuthorizationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Reason** | **string** |  | 
**RejectionReasons** | [**[]RejectionReason**](RejectionReason.md) |  | 
**CorrelationId** | Pointer to **string** | Request correlation identifier echoed in the &#x60;X-Correlation-ID&#x60; header and JSON error payloads. | [optional] 

## Methods

### NewAuthorizationError

`func NewAuthorizationError(status string, reason string, rejectionReasons []RejectionReason, ) *AuthorizationError`

NewAuthorizationError instantiates a new AuthorizationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationErrorWithDefaults

`func NewAuthorizationErrorWithDefaults() *AuthorizationError`

NewAuthorizationErrorWithDefaults instantiates a new AuthorizationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AuthorizationError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorizationError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorizationError) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *AuthorizationError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AuthorizationError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AuthorizationError) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRejectionReasons

`func (o *AuthorizationError) GetRejectionReasons() []RejectionReason`

GetRejectionReasons returns the RejectionReasons field if non-nil, zero value otherwise.

### GetRejectionReasonsOk

`func (o *AuthorizationError) GetRejectionReasonsOk() (*[]RejectionReason, bool)`

GetRejectionReasonsOk returns a tuple with the RejectionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReasons

`func (o *AuthorizationError) SetRejectionReasons(v []RejectionReason)`

SetRejectionReasons sets RejectionReasons field to given value.


### GetCorrelationId

`func (o *AuthorizationError) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *AuthorizationError) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *AuthorizationError) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *AuthorizationError) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


