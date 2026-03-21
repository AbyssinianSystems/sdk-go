# OutcomeReasonCodeStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonCode** | [**RecommendationReasonCode**](RecommendationReasonCode.md) |  | 
**RecommendationCount** | **int32** |  | 
**ConfirmedCaseCount** | **int32** |  | 
**CleanCaseCount** | **int32** |  | 

## Methods

### NewOutcomeReasonCodeStat

`func NewOutcomeReasonCodeStat(reasonCode RecommendationReasonCode, recommendationCount int32, confirmedCaseCount int32, cleanCaseCount int32, ) *OutcomeReasonCodeStat`

NewOutcomeReasonCodeStat instantiates a new OutcomeReasonCodeStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcomeReasonCodeStatWithDefaults

`func NewOutcomeReasonCodeStatWithDefaults() *OutcomeReasonCodeStat`

NewOutcomeReasonCodeStatWithDefaults instantiates a new OutcomeReasonCodeStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonCode

`func (o *OutcomeReasonCodeStat) GetReasonCode() RecommendationReasonCode`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *OutcomeReasonCodeStat) GetReasonCodeOk() (*RecommendationReasonCode, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *OutcomeReasonCodeStat) SetReasonCode(v RecommendationReasonCode)`

SetReasonCode sets ReasonCode field to given value.


### GetRecommendationCount

`func (o *OutcomeReasonCodeStat) GetRecommendationCount() int32`

GetRecommendationCount returns the RecommendationCount field if non-nil, zero value otherwise.

### GetRecommendationCountOk

`func (o *OutcomeReasonCodeStat) GetRecommendationCountOk() (*int32, bool)`

GetRecommendationCountOk returns a tuple with the RecommendationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationCount

`func (o *OutcomeReasonCodeStat) SetRecommendationCount(v int32)`

SetRecommendationCount sets RecommendationCount field to given value.


### GetConfirmedCaseCount

`func (o *OutcomeReasonCodeStat) GetConfirmedCaseCount() int32`

GetConfirmedCaseCount returns the ConfirmedCaseCount field if non-nil, zero value otherwise.

### GetConfirmedCaseCountOk

`func (o *OutcomeReasonCodeStat) GetConfirmedCaseCountOk() (*int32, bool)`

GetConfirmedCaseCountOk returns a tuple with the ConfirmedCaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedCaseCount

`func (o *OutcomeReasonCodeStat) SetConfirmedCaseCount(v int32)`

SetConfirmedCaseCount sets ConfirmedCaseCount field to given value.


### GetCleanCaseCount

`func (o *OutcomeReasonCodeStat) GetCleanCaseCount() int32`

GetCleanCaseCount returns the CleanCaseCount field if non-nil, zero value otherwise.

### GetCleanCaseCountOk

`func (o *OutcomeReasonCodeStat) GetCleanCaseCountOk() (*int32, bool)`

GetCleanCaseCountOk returns a tuple with the CleanCaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanCaseCount

`func (o *OutcomeReasonCodeStat) SetCleanCaseCount(v int32)`

SetCleanCaseCount sets CleanCaseCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


