# IngestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**CanonicalEventRef** | Pointer to [**ArtifactRef**](ArtifactRef.md) |  | [optional] 
**Event** | Pointer to [**SignalEvent**](SignalEvent.md) |  | [optional] 
**Evaluation** | Pointer to [**EvaluationResult**](EvaluationResult.md) |  | [optional] 
**RejectionReasons** | Pointer to [**[]RejectionReason**](RejectionReason.md) |  | [optional] 
**CorrelationId** | Pointer to **string** | Request correlation identifier echoed in the &#x60;X-Correlation-ID&#x60; header and JSON error payloads. | [optional] 

## Methods

### NewIngestResult

`func NewIngestResult(status string, ) *IngestResult`

NewIngestResult instantiates a new IngestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestResultWithDefaults

`func NewIngestResultWithDefaults() *IngestResult`

NewIngestResultWithDefaults instantiates a new IngestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IngestResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IngestResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IngestResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *IngestResult) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IngestResult) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IngestResult) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IngestResult) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCanonicalEventRef

`func (o *IngestResult) GetCanonicalEventRef() ArtifactRef`

GetCanonicalEventRef returns the CanonicalEventRef field if non-nil, zero value otherwise.

### GetCanonicalEventRefOk

`func (o *IngestResult) GetCanonicalEventRefOk() (*ArtifactRef, bool)`

GetCanonicalEventRefOk returns a tuple with the CanonicalEventRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalEventRef

`func (o *IngestResult) SetCanonicalEventRef(v ArtifactRef)`

SetCanonicalEventRef sets CanonicalEventRef field to given value.

### HasCanonicalEventRef

`func (o *IngestResult) HasCanonicalEventRef() bool`

HasCanonicalEventRef returns a boolean if a field has been set.

### GetEvent

`func (o *IngestResult) GetEvent() SignalEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *IngestResult) GetEventOk() (*SignalEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *IngestResult) SetEvent(v SignalEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *IngestResult) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEvaluation

`func (o *IngestResult) GetEvaluation() EvaluationResult`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *IngestResult) GetEvaluationOk() (*EvaluationResult, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *IngestResult) SetEvaluation(v EvaluationResult)`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *IngestResult) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### GetRejectionReasons

`func (o *IngestResult) GetRejectionReasons() []RejectionReason`

GetRejectionReasons returns the RejectionReasons field if non-nil, zero value otherwise.

### GetRejectionReasonsOk

`func (o *IngestResult) GetRejectionReasonsOk() (*[]RejectionReason, bool)`

GetRejectionReasonsOk returns a tuple with the RejectionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReasons

`func (o *IngestResult) SetRejectionReasons(v []RejectionReason)`

SetRejectionReasons sets RejectionReasons field to given value.

### HasRejectionReasons

`func (o *IngestResult) HasRejectionReasons() bool`

HasRejectionReasons returns a boolean if a field has been set.

### GetCorrelationId

`func (o *IngestResult) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *IngestResult) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *IngestResult) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *IngestResult) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


