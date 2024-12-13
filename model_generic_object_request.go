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

// checks if the GenericObjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericObjectRequest{}

// GenericObjectRequest Minimal representation of some generic object identified by ContentType and PK.
type GenericObjectRequest struct {
	ObjectType           string `json:"object_type"`
	ObjectId             int32  `json:"object_id"`
	AdditionalProperties map[string]interface{}
}

type _GenericObjectRequest GenericObjectRequest

// NewGenericObjectRequest instantiates a new GenericObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericObjectRequest(objectType string, objectId int32) *GenericObjectRequest {
	this := GenericObjectRequest{}
	this.ObjectType = objectType
	this.ObjectId = objectId
	return &this
}

// NewGenericObjectRequestWithDefaults instantiates a new GenericObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericObjectRequestWithDefaults() *GenericObjectRequest {
	this := GenericObjectRequest{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *GenericObjectRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *GenericObjectRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *GenericObjectRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *GenericObjectRequest) GetObjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *GenericObjectRequest) GetObjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *GenericObjectRequest) SetObjectId(v int32) {
	o.ObjectId = v
}

func (o GenericObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericObjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenericObjectRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_type",
		"object_id",
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

	varGenericObjectRequest := _GenericObjectRequest{}

	err = json.Unmarshal(data, &varGenericObjectRequest)

	if err != nil {
		return err
	}

	*o = GenericObjectRequest(varGenericObjectRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericObjectRequest struct {
	value *GenericObjectRequest
	isSet bool
}

func (v NullableGenericObjectRequest) Get() *GenericObjectRequest {
	return v.value
}

func (v *NullableGenericObjectRequest) Set(val *GenericObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericObjectRequest(val *GenericObjectRequest) *NullableGenericObjectRequest {
	return &NullableGenericObjectRequest{value: val, isSet: true}
}

func (v NullableGenericObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
