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

// checks if the WritableCustomFieldChoiceSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableCustomFieldChoiceSetRequest{}

// WritableCustomFieldChoiceSetRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableCustomFieldChoiceSetRequest struct {
	Name         string                                                        `json:"name"`
	Description  *string                                                       `json:"description,omitempty"`
	BaseChoices  NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices `json:"base_choices,omitempty"`
	ExtraChoices [][]interface{}                                               `json:"extra_choices"`
	// Choices are automatically ordered alphabetically
	OrderAlphabetically  *bool `json:"order_alphabetically,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableCustomFieldChoiceSetRequest WritableCustomFieldChoiceSetRequest

// NewWritableCustomFieldChoiceSetRequest instantiates a new WritableCustomFieldChoiceSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableCustomFieldChoiceSetRequest(name string, extraChoices [][]interface{}) *WritableCustomFieldChoiceSetRequest {
	this := WritableCustomFieldChoiceSetRequest{}
	this.Name = name
	this.ExtraChoices = extraChoices
	return &this
}

// NewWritableCustomFieldChoiceSetRequestWithDefaults instantiates a new WritableCustomFieldChoiceSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableCustomFieldChoiceSetRequestWithDefaults() *WritableCustomFieldChoiceSetRequest {
	this := WritableCustomFieldChoiceSetRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableCustomFieldChoiceSetRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldChoiceSetRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableCustomFieldChoiceSetRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableCustomFieldChoiceSetRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldChoiceSetRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableCustomFieldChoiceSetRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableCustomFieldChoiceSetRequest) SetDescription(v string) {
	o.Description = &v
}

// GetBaseChoices returns the BaseChoices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCustomFieldChoiceSetRequest) GetBaseChoices() PatchedWritableCustomFieldChoiceSetRequestBaseChoices {
	if o == nil || IsNil(o.BaseChoices.Get()) {
		var ret PatchedWritableCustomFieldChoiceSetRequestBaseChoices
		return ret
	}
	return *o.BaseChoices.Get()
}

// GetBaseChoicesOk returns a tuple with the BaseChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCustomFieldChoiceSetRequest) GetBaseChoicesOk() (*PatchedWritableCustomFieldChoiceSetRequestBaseChoices, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseChoices.Get(), o.BaseChoices.IsSet()
}

// HasBaseChoices returns a boolean if a field has been set.
func (o *WritableCustomFieldChoiceSetRequest) HasBaseChoices() bool {
	if o != nil && o.BaseChoices.IsSet() {
		return true
	}

	return false
}

// SetBaseChoices gets a reference to the given NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices and assigns it to the BaseChoices field.
func (o *WritableCustomFieldChoiceSetRequest) SetBaseChoices(v PatchedWritableCustomFieldChoiceSetRequestBaseChoices) {
	o.BaseChoices.Set(&v)
}

// SetBaseChoicesNil sets the value for BaseChoices to be an explicit nil
func (o *WritableCustomFieldChoiceSetRequest) SetBaseChoicesNil() {
	o.BaseChoices.Set(nil)
}

// UnsetBaseChoices ensures that no value is present for BaseChoices, not even an explicit nil
func (o *WritableCustomFieldChoiceSetRequest) UnsetBaseChoices() {
	o.BaseChoices.Unset()
}

// GetExtraChoices returns the ExtraChoices field value
func (o *WritableCustomFieldChoiceSetRequest) GetExtraChoices() [][]interface{} {
	if o == nil {
		var ret [][]interface{}
		return ret
	}

	return o.ExtraChoices
}

// GetExtraChoicesOk returns a tuple with the ExtraChoices field value
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldChoiceSetRequest) GetExtraChoicesOk() ([][]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraChoices, true
}

// SetExtraChoices sets field value
func (o *WritableCustomFieldChoiceSetRequest) SetExtraChoices(v [][]interface{}) {
	o.ExtraChoices = v
}

// GetOrderAlphabetically returns the OrderAlphabetically field value if set, zero value otherwise.
func (o *WritableCustomFieldChoiceSetRequest) GetOrderAlphabetically() bool {
	if o == nil || IsNil(o.OrderAlphabetically) {
		var ret bool
		return ret
	}
	return *o.OrderAlphabetically
}

// GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCustomFieldChoiceSetRequest) GetOrderAlphabeticallyOk() (*bool, bool) {
	if o == nil || IsNil(o.OrderAlphabetically) {
		return nil, false
	}
	return o.OrderAlphabetically, true
}

// HasOrderAlphabetically returns a boolean if a field has been set.
func (o *WritableCustomFieldChoiceSetRequest) HasOrderAlphabetically() bool {
	if o != nil && !IsNil(o.OrderAlphabetically) {
		return true
	}

	return false
}

// SetOrderAlphabetically gets a reference to the given bool and assigns it to the OrderAlphabetically field.
func (o *WritableCustomFieldChoiceSetRequest) SetOrderAlphabetically(v bool) {
	o.OrderAlphabetically = &v
}

func (o WritableCustomFieldChoiceSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableCustomFieldChoiceSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.BaseChoices.IsSet() {
		toSerialize["base_choices"] = o.BaseChoices.Get()
	}
	toSerialize["extra_choices"] = o.ExtraChoices
	if !IsNil(o.OrderAlphabetically) {
		toSerialize["order_alphabetically"] = o.OrderAlphabetically
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableCustomFieldChoiceSetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"extra_choices",
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

	varWritableCustomFieldChoiceSetRequest := _WritableCustomFieldChoiceSetRequest{}

	err = json.Unmarshal(data, &varWritableCustomFieldChoiceSetRequest)

	if err != nil {
		return err
	}

	*o = WritableCustomFieldChoiceSetRequest(varWritableCustomFieldChoiceSetRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "base_choices")
		delete(additionalProperties, "extra_choices")
		delete(additionalProperties, "order_alphabetically")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableCustomFieldChoiceSetRequest struct {
	value *WritableCustomFieldChoiceSetRequest
	isSet bool
}

func (v NullableWritableCustomFieldChoiceSetRequest) Get() *WritableCustomFieldChoiceSetRequest {
	return v.value
}

func (v *NullableWritableCustomFieldChoiceSetRequest) Set(val *WritableCustomFieldChoiceSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableCustomFieldChoiceSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableCustomFieldChoiceSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableCustomFieldChoiceSetRequest(val *WritableCustomFieldChoiceSetRequest) *NullableWritableCustomFieldChoiceSetRequest {
	return &NullableWritableCustomFieldChoiceSetRequest{value: val, isSet: true}
}

func (v NullableWritableCustomFieldChoiceSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableCustomFieldChoiceSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
