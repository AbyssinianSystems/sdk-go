# RejectionReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Field** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 

## Methods

### NewRejectionReason

`func NewRejectionReason(code string, message string, ) *RejectionReason`

NewRejectionReason instantiates a new RejectionReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectionReasonWithDefaults

`func NewRejectionReasonWithDefaults() *RejectionReason`

NewRejectionReasonWithDefaults instantiates a new RejectionReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RejectionReason) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RejectionReason) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RejectionReason) SetCode(v string)`

SetCode sets Code field to given value.


### GetField

`func (o *RejectionReason) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *RejectionReason) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *RejectionReason) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *RejectionReason) HasField() bool`

HasField returns a boolean if a field has been set.

### GetMessage

`func (o *RejectionReason) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RejectionReason) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RejectionReason) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


