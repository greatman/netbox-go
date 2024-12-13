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
	"time"
)

// checks if the ObjectChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectChange{}

// ObjectChange struct for ObjectChange
type ObjectChange struct {
	Id                   int32              `json:"id"`
	Url                  string             `json:"url"`
	DisplayUrl           string             `json:"display_url"`
	Display              string             `json:"display"`
	Time                 time.Time          `json:"time"`
	User                 BriefUser          `json:"user"`
	UserName             string             `json:"user_name"`
	RequestId            string             `json:"request_id"`
	Action               ObjectChangeAction `json:"action"`
	ChangedObjectType    string             `json:"changed_object_type"`
	ChangedObjectId      int64              `json:"changed_object_id"`
	ChangedObject        interface{}        `json:"changed_object"`
	PrechangeData        interface{}        `json:"prechange_data"`
	PostchangeData       interface{}        `json:"postchange_data"`
	AdditionalProperties map[string]interface{}
}

type _ObjectChange ObjectChange

// NewObjectChange instantiates a new ObjectChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectChange(id int32, url string, displayUrl string, display string, time time.Time, user BriefUser, userName string, requestId string, action ObjectChangeAction, changedObjectType string, changedObjectId int64, changedObject interface{}, prechangeData interface{}, postchangeData interface{}) *ObjectChange {
	this := ObjectChange{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Time = time
	this.User = user
	this.UserName = userName
	this.RequestId = requestId
	this.Action = action
	this.ChangedObjectType = changedObjectType
	this.ChangedObjectId = changedObjectId
	this.ChangedObject = changedObject
	this.PrechangeData = prechangeData
	this.PostchangeData = postchangeData
	return &this
}

// NewObjectChangeWithDefaults instantiates a new ObjectChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectChangeWithDefaults() *ObjectChange {
	this := ObjectChange{}
	return &this
}

// GetId returns the Id field value
func (o *ObjectChange) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ObjectChange) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ObjectChange) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ObjectChange) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *ObjectChange) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *ObjectChange) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *ObjectChange) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ObjectChange) SetDisplay(v string) {
	o.Display = v
}

// GetTime returns the Time field value
func (o *ObjectChange) GetTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *ObjectChange) SetTime(v time.Time) {
	o.Time = v
}

// GetUser returns the User field value
func (o *ObjectChange) GetUser() BriefUser {
	if o == nil {
		var ret BriefUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetUserOk() (*BriefUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ObjectChange) SetUser(v BriefUser) {
	o.User = v
}

// GetUserName returns the UserName field value
func (o *ObjectChange) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *ObjectChange) SetUserName(v string) {
	o.UserName = v
}

// GetRequestId returns the RequestId field value
func (o *ObjectChange) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ObjectChange) SetRequestId(v string) {
	o.RequestId = v
}

// GetAction returns the Action field value
func (o *ObjectChange) GetAction() ObjectChangeAction {
	if o == nil {
		var ret ObjectChangeAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetActionOk() (*ObjectChangeAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ObjectChange) SetAction(v ObjectChangeAction) {
	o.Action = v
}

// GetChangedObjectType returns the ChangedObjectType field value
func (o *ObjectChange) GetChangedObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangedObjectType
}

// GetChangedObjectTypeOk returns a tuple with the ChangedObjectType field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetChangedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangedObjectType, true
}

// SetChangedObjectType sets field value
func (o *ObjectChange) SetChangedObjectType(v string) {
	o.ChangedObjectType = v
}

// GetChangedObjectId returns the ChangedObjectId field value
func (o *ObjectChange) GetChangedObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ChangedObjectId
}

// GetChangedObjectIdOk returns a tuple with the ChangedObjectId field value
// and a boolean to check if the value has been set.
func (o *ObjectChange) GetChangedObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangedObjectId, true
}

// SetChangedObjectId sets field value
func (o *ObjectChange) SetChangedObjectId(v int64) {
	o.ChangedObjectId = v
}

// GetChangedObject returns the ChangedObject field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ObjectChange) GetChangedObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ChangedObject
}

// GetChangedObjectOk returns a tuple with the ChangedObject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectChange) GetChangedObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ChangedObject) {
		return nil, false
	}
	return &o.ChangedObject, true
}

// SetChangedObject sets field value
func (o *ObjectChange) SetChangedObject(v interface{}) {
	o.ChangedObject = v
}

// GetPrechangeData returns the PrechangeData field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ObjectChange) GetPrechangeData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PrechangeData
}

// GetPrechangeDataOk returns a tuple with the PrechangeData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectChange) GetPrechangeDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PrechangeData) {
		return nil, false
	}
	return &o.PrechangeData, true
}

// SetPrechangeData sets field value
func (o *ObjectChange) SetPrechangeData(v interface{}) {
	o.PrechangeData = v
}

// GetPostchangeData returns the PostchangeData field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ObjectChange) GetPostchangeData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PostchangeData
}

// GetPostchangeDataOk returns a tuple with the PostchangeData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectChange) GetPostchangeDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PostchangeData) {
		return nil, false
	}
	return &o.PostchangeData, true
}

// SetPostchangeData sets field value
func (o *ObjectChange) SetPostchangeData(v interface{}) {
	o.PostchangeData = v
}

func (o ObjectChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["time"] = o.Time
	toSerialize["user"] = o.User
	toSerialize["user_name"] = o.UserName
	toSerialize["request_id"] = o.RequestId
	toSerialize["action"] = o.Action
	toSerialize["changed_object_type"] = o.ChangedObjectType
	toSerialize["changed_object_id"] = o.ChangedObjectId
	if o.ChangedObject != nil {
		toSerialize["changed_object"] = o.ChangedObject
	}
	if o.PrechangeData != nil {
		toSerialize["prechange_data"] = o.PrechangeData
	}
	if o.PostchangeData != nil {
		toSerialize["postchange_data"] = o.PostchangeData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ObjectChange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"time",
		"user",
		"user_name",
		"request_id",
		"action",
		"changed_object_type",
		"changed_object_id",
		"changed_object",
		"prechange_data",
		"postchange_data",
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

	varObjectChange := _ObjectChange{}

	err = json.Unmarshal(data, &varObjectChange)

	if err != nil {
		return err
	}

	*o = ObjectChange(varObjectChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "time")
		delete(additionalProperties, "user")
		delete(additionalProperties, "user_name")
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "action")
		delete(additionalProperties, "changed_object_type")
		delete(additionalProperties, "changed_object_id")
		delete(additionalProperties, "changed_object")
		delete(additionalProperties, "prechange_data")
		delete(additionalProperties, "postchange_data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObjectChange struct {
	value *ObjectChange
	isSet bool
}

func (v NullableObjectChange) Get() *ObjectChange {
	return v.value
}

func (v *NullableObjectChange) Set(val *ObjectChange) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectChange) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectChange(val *ObjectChange) *NullableObjectChange {
	return &NullableObjectChange{value: val, isSet: true}
}

func (v NullableObjectChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}