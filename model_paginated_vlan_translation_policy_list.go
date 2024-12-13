/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.2-beta1 (4.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginatedVLANTranslationPolicyList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedVLANTranslationPolicyList{}

// PaginatedVLANTranslationPolicyList struct for PaginatedVLANTranslationPolicyList
type PaginatedVLANTranslationPolicyList struct {
	Count                int64                   `json:"count"`
	Next                 NullableString          `json:"next,omitempty"`
	Previous             NullableString          `json:"previous,omitempty"`
	Results              []VLANTranslationPolicy `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedVLANTranslationPolicyList PaginatedVLANTranslationPolicyList

// NewPaginatedVLANTranslationPolicyList instantiates a new PaginatedVLANTranslationPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedVLANTranslationPolicyList(count int64, results []VLANTranslationPolicy) *PaginatedVLANTranslationPolicyList {
	this := PaginatedVLANTranslationPolicyList{}
	this.Count = count
	this.Results = results
	return &this
}

// NewPaginatedVLANTranslationPolicyListWithDefaults instantiates a new PaginatedVLANTranslationPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedVLANTranslationPolicyListWithDefaults() *PaginatedVLANTranslationPolicyList {
	this := PaginatedVLANTranslationPolicyList{}
	return &this
}

// GetCount returns the Count field value
func (o *PaginatedVLANTranslationPolicyList) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginatedVLANTranslationPolicyList) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginatedVLANTranslationPolicyList) SetCount(v int64) {
	o.Count = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedVLANTranslationPolicyList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedVLANTranslationPolicyList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedVLANTranslationPolicyList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedVLANTranslationPolicyList) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedVLANTranslationPolicyList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedVLANTranslationPolicyList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedVLANTranslationPolicyList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedVLANTranslationPolicyList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedVLANTranslationPolicyList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedVLANTranslationPolicyList) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedVLANTranslationPolicyList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedVLANTranslationPolicyList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *PaginatedVLANTranslationPolicyList) GetResults() []VLANTranslationPolicy {
	if o == nil {
		var ret []VLANTranslationPolicy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedVLANTranslationPolicyList) GetResultsOk() ([]VLANTranslationPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedVLANTranslationPolicyList) SetResults(v []VLANTranslationPolicy) {
	o.Results = v
}

func (o PaginatedVLANTranslationPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedVLANTranslationPolicyList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedVLANTranslationPolicyList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"results",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaginatedVLANTranslationPolicyList := _PaginatedVLANTranslationPolicyList{}

	err = json.Unmarshal(data, &varPaginatedVLANTranslationPolicyList)

	if err != nil {
		return err
	}

	*o = PaginatedVLANTranslationPolicyList(varPaginatedVLANTranslationPolicyList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedVLANTranslationPolicyList struct {
	value *PaginatedVLANTranslationPolicyList
	isSet bool
}

func (v NullablePaginatedVLANTranslationPolicyList) Get() *PaginatedVLANTranslationPolicyList {
	return v.value
}

func (v *NullablePaginatedVLANTranslationPolicyList) Set(val *PaginatedVLANTranslationPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedVLANTranslationPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedVLANTranslationPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedVLANTranslationPolicyList(val *PaginatedVLANTranslationPolicyList) *NullablePaginatedVLANTranslationPolicyList {
	return &NullablePaginatedVLANTranslationPolicyList{value: val, isSet: true}
}

func (v NullablePaginatedVLANTranslationPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedVLANTranslationPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
