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

// checks if the BookmarkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkRequest{}

// BookmarkRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BookmarkRequest struct {
	ObjectType           string           `json:"object_type"`
	ObjectId             int64            `json:"object_id"`
	User                 BriefUserRequest `json:"user"`
	AdditionalProperties map[string]interface{}
}

type _BookmarkRequest BookmarkRequest

// NewBookmarkRequest instantiates a new BookmarkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkRequest(objectType string, objectId int64, user BriefUserRequest) *BookmarkRequest {
	this := BookmarkRequest{}
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.User = user
	return &this
}

// NewBookmarkRequestWithDefaults instantiates a new BookmarkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkRequestWithDefaults() *BookmarkRequest {
	this := BookmarkRequest{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *BookmarkRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BookmarkRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BookmarkRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *BookmarkRequest) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *BookmarkRequest) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *BookmarkRequest) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetUser returns the User field value
func (o *BookmarkRequest) GetUser() BriefUserRequest {
	if o == nil {
		var ret BriefUserRequest
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *BookmarkRequest) GetUserOk() (*BriefUserRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *BookmarkRequest) SetUser(v BriefUserRequest) {
	o.User = v
}

func (o BookmarkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BookmarkRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_type",
		"object_id",
		"user",
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

	varBookmarkRequest := _BookmarkRequest{}

	err = json.Unmarshal(data, &varBookmarkRequest)

	if err != nil {
		return err
	}

	*o = BookmarkRequest(varBookmarkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBookmarkRequest struct {
	value *BookmarkRequest
	isSet bool
}

func (v NullableBookmarkRequest) Get() *BookmarkRequest {
	return v.value
}

func (v *NullableBookmarkRequest) Set(val *BookmarkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkRequest(val *BookmarkRequest) *NullableBookmarkRequest {
	return &NullableBookmarkRequest{value: val, isSet: true}
}

func (v NullableBookmarkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
