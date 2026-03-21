# SignalEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DedupeKey** | **string** |  | 
**SchemaVersion** | **string** |  | 
**SubjectId** | **string** |  | 
**SubjectKind** | **string** |  | 
**ProducerEventId** | Pointer to **string** |  | [optional] 
**SignalType** | **string** |  | 
**Producer** | **string** |  | 
**OccurredAt** | **time.Time** |  | 
**IngestedAt** | **time.Time** |  | 
**Severity** | **string** |  | 
**Confidence** | Pointer to **float32** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**NormalizationVersion** | **string** |  | 
**CorrelationIds** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to [**[]MetadataField**](MetadataField.md) |  | [optional] 

## Methods

### NewSignalEvent

`func NewSignalEvent(id string, dedupeKey string, schemaVersion string, subjectId string, subjectKind string, signalType string, producer string, occurredAt time.Time, ingestedAt time.Time, severity string, normalizationVersion string, ) *SignalEvent`

NewSignalEvent instantiates a new SignalEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalEventWithDefaults

`func NewSignalEventWithDefaults() *SignalEvent`

NewSignalEventWithDefaults instantiates a new SignalEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignalEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignalEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignalEvent) SetId(v string)`

SetId sets Id field to given value.


### GetDedupeKey

`func (o *SignalEvent) GetDedupeKey() string`

GetDedupeKey returns the DedupeKey field if non-nil, zero value otherwise.

### GetDedupeKeyOk

`func (o *SignalEvent) GetDedupeKeyOk() (*string, bool)`

GetDedupeKeyOk returns a tuple with the DedupeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupeKey

`func (o *SignalEvent) SetDedupeKey(v string)`

SetDedupeKey sets DedupeKey field to given value.


### GetSchemaVersion

`func (o *SignalEvent) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *SignalEvent) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *SignalEvent) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetSubjectId

`func (o *SignalEvent) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *SignalEvent) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *SignalEvent) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetSubjectKind

`func (o *SignalEvent) GetSubjectKind() string`

GetSubjectKind returns the SubjectKind field if non-nil, zero value otherwise.

### GetSubjectKindOk

`func (o *SignalEvent) GetSubjectKindOk() (*string, bool)`

GetSubjectKindOk returns a tuple with the SubjectKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKind

`func (o *SignalEvent) SetSubjectKind(v string)`

SetSubjectKind sets SubjectKind field to given value.


### GetProducerEventId

`func (o *SignalEvent) GetProducerEventId() string`

GetProducerEventId returns the ProducerEventId field if non-nil, zero value otherwise.

### GetProducerEventIdOk

`func (o *SignalEvent) GetProducerEventIdOk() (*string, bool)`

GetProducerEventIdOk returns a tuple with the ProducerEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEventId

`func (o *SignalEvent) SetProducerEventId(v string)`

SetProducerEventId sets ProducerEventId field to given value.

### HasProducerEventId

`func (o *SignalEvent) HasProducerEventId() bool`

HasProducerEventId returns a boolean if a field has been set.

### GetSignalType

`func (o *SignalEvent) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *SignalEvent) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *SignalEvent) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetProducer

`func (o *SignalEvent) GetProducer() string`

GetProducer returns the Producer field if non-nil, zero value otherwise.

### GetProducerOk

`func (o *SignalEvent) GetProducerOk() (*string, bool)`

GetProducerOk returns a tuple with the Producer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducer

`func (o *SignalEvent) SetProducer(v string)`

SetProducer sets Producer field to given value.


### GetOccurredAt

`func (o *SignalEvent) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *SignalEvent) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *SignalEvent) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetIngestedAt

`func (o *SignalEvent) GetIngestedAt() time.Time`

GetIngestedAt returns the IngestedAt field if non-nil, zero value otherwise.

### GetIngestedAtOk

`func (o *SignalEvent) GetIngestedAtOk() (*time.Time, bool)`

GetIngestedAtOk returns a tuple with the IngestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedAt

`func (o *SignalEvent) SetIngestedAt(v time.Time)`

SetIngestedAt sets IngestedAt field to given value.


### GetSeverity

`func (o *SignalEvent) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SignalEvent) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SignalEvent) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetConfidence

`func (o *SignalEvent) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SignalEvent) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SignalEvent) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *SignalEvent) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetValue

`func (o *SignalEvent) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SignalEvent) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SignalEvent) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *SignalEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnit

`func (o *SignalEvent) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SignalEvent) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SignalEvent) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SignalEvent) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetNormalizationVersion

`func (o *SignalEvent) GetNormalizationVersion() string`

GetNormalizationVersion returns the NormalizationVersion field if non-nil, zero value otherwise.

### GetNormalizationVersionOk

`func (o *SignalEvent) GetNormalizationVersionOk() (*string, bool)`

GetNormalizationVersionOk returns a tuple with the NormalizationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizationVersion

`func (o *SignalEvent) SetNormalizationVersion(v string)`

SetNormalizationVersion sets NormalizationVersion field to given value.


### GetCorrelationIds

`func (o *SignalEvent) GetCorrelationIds() []string`

GetCorrelationIds returns the CorrelationIds field if non-nil, zero value otherwise.

### GetCorrelationIdsOk

`func (o *SignalEvent) GetCorrelationIdsOk() (*[]string, bool)`

GetCorrelationIdsOk returns a tuple with the CorrelationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIds

`func (o *SignalEvent) SetCorrelationIds(v []string)`

SetCorrelationIds sets CorrelationIds field to given value.

### HasCorrelationIds

`func (o *SignalEvent) HasCorrelationIds() bool`

HasCorrelationIds returns a boolean if a field has been set.

### GetMetadata

`func (o *SignalEvent) GetMetadata() []MetadataField`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SignalEvent) GetMetadataOk() (*[]MetadataField, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SignalEvent) SetMetadata(v []MetadataField)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SignalEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


