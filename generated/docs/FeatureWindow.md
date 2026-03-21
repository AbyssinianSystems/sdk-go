# FeatureWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**StartAt** | **time.Time** |  | 
**EndAt** | **time.Time** |  | 
**EventIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFeatureWindow

`func NewFeatureWindow(name string, startAt time.Time, endAt time.Time, ) *FeatureWindow`

NewFeatureWindow instantiates a new FeatureWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureWindowWithDefaults

`func NewFeatureWindowWithDefaults() *FeatureWindow`

NewFeatureWindowWithDefaults instantiates a new FeatureWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureWindow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureWindow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureWindow) SetName(v string)`

SetName sets Name field to given value.


### GetStartAt

`func (o *FeatureWindow) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *FeatureWindow) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *FeatureWindow) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.


### GetEndAt

`func (o *FeatureWindow) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *FeatureWindow) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *FeatureWindow) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.


### GetEventIds

`func (o *FeatureWindow) GetEventIds() []string`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *FeatureWindow) GetEventIdsOk() (*[]string, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *FeatureWindow) SetEventIds(v []string)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *FeatureWindow) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


