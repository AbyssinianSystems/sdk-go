# RawSignalEventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | **string** |  | 
**Producer** | **string** |  | 
**ProducerEventId** | Pointer to **string** |  | [optional] 
**SignalType** | **string** | Canonical or legacy alias value accepted by ingestion | 
**OccurredAt** | **time.Time** |  | 
**Severity** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **float32** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**CorrelationIds** | Pointer to **[]string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewRawSignalEventPayload

`func NewRawSignalEventPayload(subjectId string, producer string, signalType string, occurredAt time.Time, ) *RawSignalEventPayload`

NewRawSignalEventPayload instantiates a new RawSignalEventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawSignalEventPayloadWithDefaults

`func NewRawSignalEventPayloadWithDefaults() *RawSignalEventPayload`

NewRawSignalEventPayloadWithDefaults instantiates a new RawSignalEventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *RawSignalEventPayload) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *RawSignalEventPayload) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *RawSignalEventPayload) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetProducer

`func (o *RawSignalEventPayload) GetProducer() string`

GetProducer returns the Producer field if non-nil, zero value otherwise.

### GetProducerOk

`func (o *RawSignalEventPayload) GetProducerOk() (*string, bool)`

GetProducerOk returns a tuple with the Producer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducer

`func (o *RawSignalEventPayload) SetProducer(v string)`

SetProducer sets Producer field to given value.


### GetProducerEventId

`func (o *RawSignalEventPayload) GetProducerEventId() string`

GetProducerEventId returns the ProducerEventId field if non-nil, zero value otherwise.

### GetProducerEventIdOk

`func (o *RawSignalEventPayload) GetProducerEventIdOk() (*string, bool)`

GetProducerEventIdOk returns a tuple with the ProducerEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEventId

`func (o *RawSignalEventPayload) SetProducerEventId(v string)`

SetProducerEventId sets ProducerEventId field to given value.

### HasProducerEventId

`func (o *RawSignalEventPayload) HasProducerEventId() bool`

HasProducerEventId returns a boolean if a field has been set.

### GetSignalType

`func (o *RawSignalEventPayload) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *RawSignalEventPayload) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *RawSignalEventPayload) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetOccurredAt

`func (o *RawSignalEventPayload) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *RawSignalEventPayload) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *RawSignalEventPayload) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetSeverity

`func (o *RawSignalEventPayload) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RawSignalEventPayload) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RawSignalEventPayload) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RawSignalEventPayload) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetConfidence

`func (o *RawSignalEventPayload) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RawSignalEventPayload) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RawSignalEventPayload) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RawSignalEventPayload) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetValue

`func (o *RawSignalEventPayload) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RawSignalEventPayload) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RawSignalEventPayload) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *RawSignalEventPayload) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnit

`func (o *RawSignalEventPayload) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RawSignalEventPayload) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RawSignalEventPayload) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RawSignalEventPayload) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCorrelationIds

`func (o *RawSignalEventPayload) GetCorrelationIds() []string`

GetCorrelationIds returns the CorrelationIds field if non-nil, zero value otherwise.

### GetCorrelationIdsOk

`func (o *RawSignalEventPayload) GetCorrelationIdsOk() (*[]string, bool)`

GetCorrelationIdsOk returns a tuple with the CorrelationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIds

`func (o *RawSignalEventPayload) SetCorrelationIds(v []string)`

SetCorrelationIds sets CorrelationIds field to given value.

### HasCorrelationIds

`func (o *RawSignalEventPayload) HasCorrelationIds() bool`

HasCorrelationIds returns a boolean if a field has been set.

### GetAttributes

`func (o *RawSignalEventPayload) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RawSignalEventPayload) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RawSignalEventPayload) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RawSignalEventPayload) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


