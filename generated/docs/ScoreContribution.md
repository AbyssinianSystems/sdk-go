# ScoreContribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**RuleId** | **string** |  | 
**Category** | **string** |  | 
**Label** | **string** |  | 
**Value** | **float32** |  | 
**AppliedCap** | Pointer to **float32** |  | [optional] 
**AppliedFloor** | Pointer to **float32** |  | [optional] 
**Evidence** | Pointer to [**[]EvidenceRef**](EvidenceRef.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 

## Methods

### NewScoreContribution

`func NewScoreContribution(key string, ruleId string, category string, label string, value float32, ) *ScoreContribution`

NewScoreContribution instantiates a new ScoreContribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreContributionWithDefaults

`func NewScoreContributionWithDefaults() *ScoreContribution`

NewScoreContributionWithDefaults instantiates a new ScoreContribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ScoreContribution) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScoreContribution) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScoreContribution) SetKey(v string)`

SetKey sets Key field to given value.


### GetRuleId

`func (o *ScoreContribution) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ScoreContribution) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ScoreContribution) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetCategory

`func (o *ScoreContribution) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ScoreContribution) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ScoreContribution) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetLabel

`func (o *ScoreContribution) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ScoreContribution) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ScoreContribution) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *ScoreContribution) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScoreContribution) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScoreContribution) SetValue(v float32)`

SetValue sets Value field to given value.


### GetAppliedCap

`func (o *ScoreContribution) GetAppliedCap() float32`

GetAppliedCap returns the AppliedCap field if non-nil, zero value otherwise.

### GetAppliedCapOk

`func (o *ScoreContribution) GetAppliedCapOk() (*float32, bool)`

GetAppliedCapOk returns a tuple with the AppliedCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCap

`func (o *ScoreContribution) SetAppliedCap(v float32)`

SetAppliedCap sets AppliedCap field to given value.

### HasAppliedCap

`func (o *ScoreContribution) HasAppliedCap() bool`

HasAppliedCap returns a boolean if a field has been set.

### GetAppliedFloor

`func (o *ScoreContribution) GetAppliedFloor() float32`

GetAppliedFloor returns the AppliedFloor field if non-nil, zero value otherwise.

### GetAppliedFloorOk

`func (o *ScoreContribution) GetAppliedFloorOk() (*float32, bool)`

GetAppliedFloorOk returns a tuple with the AppliedFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedFloor

`func (o *ScoreContribution) SetAppliedFloor(v float32)`

SetAppliedFloor sets AppliedFloor field to given value.

### HasAppliedFloor

`func (o *ScoreContribution) HasAppliedFloor() bool`

HasAppliedFloor returns a boolean if a field has been set.

### GetEvidence

`func (o *ScoreContribution) GetEvidence() []EvidenceRef`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *ScoreContribution) GetEvidenceOk() (*[]EvidenceRef, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *ScoreContribution) SetEvidence(v []EvidenceRef)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *ScoreContribution) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetNotes

`func (o *ScoreContribution) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ScoreContribution) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ScoreContribution) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ScoreContribution) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


