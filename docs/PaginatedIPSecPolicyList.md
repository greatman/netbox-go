# PaginatedIPSecPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]IPSecPolicy**](IPSecPolicy.md) |  | 

## Methods

### NewPaginatedIPSecPolicyList

`func NewPaginatedIPSecPolicyList(count int64, results []IPSecPolicy, ) *PaginatedIPSecPolicyList`

NewPaginatedIPSecPolicyList instantiates a new PaginatedIPSecPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedIPSecPolicyListWithDefaults

`func NewPaginatedIPSecPolicyListWithDefaults() *PaginatedIPSecPolicyList`

NewPaginatedIPSecPolicyListWithDefaults instantiates a new PaginatedIPSecPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedIPSecPolicyList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedIPSecPolicyList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedIPSecPolicyList) SetCount(v int64)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedIPSecPolicyList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedIPSecPolicyList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedIPSecPolicyList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedIPSecPolicyList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedIPSecPolicyList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedIPSecPolicyList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedIPSecPolicyList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedIPSecPolicyList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedIPSecPolicyList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedIPSecPolicyList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedIPSecPolicyList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedIPSecPolicyList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedIPSecPolicyList) GetResults() []IPSecPolicy`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedIPSecPolicyList) GetResultsOk() (*[]IPSecPolicy, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedIPSecPolicyList) SetResults(v []IPSecPolicy)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


