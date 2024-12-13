/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.2-beta1 (4.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedCircuitTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCircuitTerminationRequest{}

// PatchedCircuitTerminationRequest Adds support for custom fields and tags.
type PatchedCircuitTerminationRequest struct {
	Circuit         *BriefCircuitRequest `json:"circuit,omitempty"`
	TermSide        *TerminationSide1    `json:"term_side,omitempty"`
	TerminationType NullableString       `json:"termination_type,omitempty"`
	TerminationId   NullableInt64        `json:"termination_id,omitempty"`
	// Physical circuit speed
	PortSpeed NullableInt64 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed NullableInt64 `json:"upstream_speed,omitempty"`
	// ID of the local cross-connect
	XconnectId *string `json:"xconnect_id,omitempty"`
	// Patch panel ID and port number(s)
	PpInfo      *string `json:"pp_info,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected        *bool                  `json:"mark_connected,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedCircuitTerminationRequest PatchedCircuitTerminationRequest

// NewPatchedCircuitTerminationRequest instantiates a new PatchedCircuitTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCircuitTerminationRequest() *PatchedCircuitTerminationRequest {
	this := PatchedCircuitTerminationRequest{}
	return &this
}

// NewPatchedCircuitTerminationRequestWithDefaults instantiates a new PatchedCircuitTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCircuitTerminationRequestWithDefaults() *PatchedCircuitTerminationRequest {
	this := PatchedCircuitTerminationRequest{}
	return &this
}

// GetCircuit returns the Circuit field value if set, zero value otherwise.
func (o *PatchedCircuitTerminationRequest) GetCircuit() BriefCircuitRequest {
	if o == nil || IsNil(o.Circuit) {
		var ret BriefCircuitRequest
		return ret
	}
	return *o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitTerminationRequest) GetCircuitOk() (*BriefCircuitRequest, bool) {
	if o == nil || IsNil(o.Circuit) {
		return nil, false
	}
	return o.Circuit, true
}

// HasCircuit returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasCircuit() bool {
	if o != nil && !IsNil(o.Circuit) {
		return true
	}

	return false
}

// SetCircuit gets a reference to the given BriefCircuitRequest and assigns it to the Circuit field.
func (o *PatchedCircuitTerminationRequest) SetCircuit(v BriefCircuitRequest) {
	o.Circuit = &v
}

// GetTermSide returns the TermSide field value if set, zero value otherwise.
func (o *PatchedCircuitTerminationRequest) GetTermSide() TerminationSide1 {
	if o == nil || IsNil(o.TermSide) {
		var ret TerminationSide1
		return ret
	}
	return *o.TermSide
}

// GetTermSideOk returns a tuple with the TermSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitTerminationRequest) GetTermSideOk() (*TerminationSide1, bool) {
	if o == nil || IsNil(o.TermSide) {
		return nil, false
	}
	return o.TermSide, true
}

// HasTermSide returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasTermSide() bool {
	if o != nil && !IsNil(o.TermSide) {
		return true
	}

	return false
}

// SetTermSide gets a reference to the given TerminationSide1 and assigns it to the TermSide field.
func (o *PatchedCircuitTerminationRequest) SetTermSide(v TerminationSide1) {
	o.TermSide = &v
}

// GetTerminationType returns the TerminationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCircuitTerminationRequest) GetTerminationType() string {
	if o == nil || IsNil(o.TerminationType.Get()) {
		var ret string
		return ret
	}
	return *o.TerminationType.Get()
}

// GetTerminationTypeOk returns a tuple with the TerminationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCircuitTerminationRequest) GetTerminationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationType.Get(), o.TerminationType.IsSet()
}

// HasTerminationType returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasTerminationType() bool {
	if o != nil && o.TerminationType.IsSet() {
		return true
	}

	return false
}

// SetTerminationType gets a reference to the given NullableString and assigns it to the TerminationType field.
func (o *PatchedCircuitTerminationRequest) SetTerminationType(v string) {
	o.TerminationType.Set(&v)
}

// SetTerminationTypeNil sets the value for TerminationType to be an explicit nil
func (o *PatchedCircuitTerminationRequest) SetTerminationTypeNil() {
	o.TerminationType.Set(nil)
}

// UnsetTerminationType ensures that no value is present for TerminationType, not even an explicit nil
func (o *PatchedCircuitTerminationRequest) UnsetTerminationType() {
	o.TerminationType.Unset()
}

// GetTerminationId returns the TerminationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCircuitTerminationRequest) GetTerminationId() int64 {
	if o == nil || IsNil(o.TerminationId.Get()) {
		var ret int64
		return ret
	}
	return *o.TerminationId.Get()
}

// GetTerminationIdOk returns a tuple with the TerminationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCircuitTerminationRequest) GetTerminationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationId.Get(), o.TerminationId.IsSet()
}

// HasTerminationId returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasTerminationId() bool {
	if o != nil && o.TerminationId.IsSet() {
		return true
	}

	return false
}

// SetTerminationId gets a reference to the given NullableInt64 and assigns it to the TerminationId field.
func (o *PatchedCircuitTerminationRequest) SetTerminationId(v int64) {
	o.TerminationId.Set(&v)
}

// SetTerminationIdNil sets the value for TerminationId to be an explicit nil
func (o *PatchedCircuitTerminationRequest) SetTerminationIdNil() {
	o.TerminationId.Set(nil)
}

// UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
func (o *PatchedCircuitTerminationRequest) UnsetTerminationId() {
	o.TerminationId.Unset()
}

// GetPortSpeed returns the PortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCircuitTerminationRequest) GetPortSpeed() int64 {
	if o == nil || IsNil(o.PortSpeed.Get()) {
		var ret int64
		return ret
	}
	return *o.PortSpeed.Get()
}

// GetPortSpeedOk returns a tuple with the PortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCircuitTerminationRequest) GetPortSpeedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortSpeed.Get(), o.PortSpeed.IsSet()
}

// HasPortSpeed returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasPortSpeed() bool {
	if o != nil && o.PortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPortSpeed gets a reference to the given NullableInt64 and assigns it to the PortSpeed field.
func (o *PatchedCircuitTerminationRequest) SetPortSpeed(v int64) {
	o.PortSpeed.Set(&v)
}

// SetPortSpeedNil sets the value for PortSpeed to be an explicit nil
func (o *PatchedCircuitTerminationRequest) SetPortSpeedNil() {
	o.PortSpeed.Set(nil)
}

// UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
func (o *PatchedCircuitTerminationRequest) UnsetPortSpeed() {
	o.PortSpeed.Unset()
}

// GetUpstreamSpeed returns the UpstreamSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCircuitTerminationRequest) GetUpstreamSpeed() int64 {
	if o == nil || IsNil(o.UpstreamSpeed.Get()) {
		var ret int64
		return ret
	}
	return *o.UpstreamSpeed.Get()
}

// GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCircuitTerminationRequest) GetUpstreamSpeedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamSpeed.Get(), o.UpstreamSpeed.IsSet()
}

// HasUpstreamSpeed returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasUpstreamSpeed() bool {
	if o != nil && o.UpstreamSpeed.IsSet() {
		return true
	}

	return false
}

// SetUpstreamSpeed gets a reference to the given NullableInt64 and assigns it to the UpstreamSpeed field.
func (o *PatchedCircuitTerminationRequest) SetUpstreamSpeed(v int64) {
	o.UpstreamSpeed.Set(&v)
}

// SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil
func (o *PatchedCircuitTerminationRequest) SetUpstreamSpeedNil() {
	o.UpstreamSpeed.Set(nil)
}

// UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
func (o *PatchedCircuitTerminationRequest) UnsetUpstreamSpeed() {
	o.UpstreamSpeed.Unset()
}

// GetXconnectId returns the XconnectId field value if set, zero value otherwise.
func (o *PatchedCircuitTerminationRequest) GetXconnectId() string {
	if o == nil || IsNil(o.XconnectId) {
		var ret string
		return ret
	}
	return *o.XconnectId
}

// GetXconnectIdOk returns a tuple with the XconnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitTerminationRequest) GetXconnectIdOk() (*string, bool) {
	if o == nil || IsNil(o.XconnectId) {
		return nil, false
	}
	return o.XconnectId, true
}

// HasXconnectId returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasXconnectId() bool {
	if o != nil && !IsNil(o.XconnectId) {
		return true
	}

	return false
}

// SetXconnectId gets a reference to the given string and assigns it to the XconnectId field.
func (o *PatchedCircuitTerminationRequest) SetXconnectId(v string) {
	o.XconnectId = &v
}

// GetPpInfo returns the PpInfo field value if set, zero value otherwise.
func (o *PatchedCircuitTerminationRequest) GetPpInfo() string {
	if o == nil || IsNil(o.PpInfo) {
		var ret string
		return ret
	}
	return *o.PpInfo
}

// GetPpInfoOk returns a tuple with the PpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitTerminationRequest) GetPpInfoOk() (*string, bool) {
	if o == nil || IsNil(o.PpInfo) {
		return nil, false
	}
	return o.PpInfo, true
}

// HasPpInfo returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasPpInfo() bool {
	if o != nil && !IsNil(o.PpInfo) {
		return true
	}

	return false
}

// SetPpInfo gets a reference to the given string and assigns it to the PpInfo field.
func (o *PatchedCircuitTerminationRequest) SetPpInfo(v string) {
	o.PpInfo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedCircuitTerminationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitTerminationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedCircuitTerminationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *PatchedCircuitTerminationRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitTerminationRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *PatchedCircuitTerminationRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedCircuitTerminationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitTerminationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedCircuitTerminationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedCircuitTerminationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitTerminationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedCircuitTerminationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedCircuitTerminationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCircuitTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Circuit) {
		toSerialize["circuit"] = o.Circuit
	}
	if !IsNil(o.TermSide) {
		toSerialize["term_side"] = o.TermSide
	}
	if o.TerminationType.IsSet() {
		toSerialize["termination_type"] = o.TerminationType.Get()
	}
	if o.TerminationId.IsSet() {
		toSerialize["termination_id"] = o.TerminationId.Get()
	}
	if o.PortSpeed.IsSet() {
		toSerialize["port_speed"] = o.PortSpeed.Get()
	}
	if o.UpstreamSpeed.IsSet() {
		toSerialize["upstream_speed"] = o.UpstreamSpeed.Get()
	}
	if !IsNil(o.XconnectId) {
		toSerialize["xconnect_id"] = o.XconnectId
	}
	if !IsNil(o.PpInfo) {
		toSerialize["pp_info"] = o.PpInfo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedCircuitTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedCircuitTerminationRequest := _PatchedCircuitTerminationRequest{}

	err = json.Unmarshal(data, &varPatchedCircuitTerminationRequest)

	if err != nil {
		return err
	}

	*o = PatchedCircuitTerminationRequest(varPatchedCircuitTerminationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "circuit")
		delete(additionalProperties, "term_side")
		delete(additionalProperties, "termination_type")
		delete(additionalProperties, "termination_id")
		delete(additionalProperties, "port_speed")
		delete(additionalProperties, "upstream_speed")
		delete(additionalProperties, "xconnect_id")
		delete(additionalProperties, "pp_info")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedCircuitTerminationRequest struct {
	value *PatchedCircuitTerminationRequest
	isSet bool
}

func (v NullablePatchedCircuitTerminationRequest) Get() *PatchedCircuitTerminationRequest {
	return v.value
}

func (v *NullablePatchedCircuitTerminationRequest) Set(val *PatchedCircuitTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCircuitTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCircuitTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCircuitTerminationRequest(val *PatchedCircuitTerminationRequest) *NullablePatchedCircuitTerminationRequest {
	return &NullablePatchedCircuitTerminationRequest{value: val, isSet: true}
}

func (v NullablePatchedCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCircuitTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
