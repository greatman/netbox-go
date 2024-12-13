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

// checks if the CircuitCircuitTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircuitCircuitTerminationRequest{}

// CircuitCircuitTerminationRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type CircuitCircuitTerminationRequest struct {
	TerminationType NullableString `json:"termination_type,omitempty"`
	TerminationId   NullableInt64  `json:"termination_id,omitempty"`
	// Physical circuit speed
	PortSpeed NullableInt64 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed NullableInt64 `json:"upstream_speed,omitempty"`
	// ID of the local cross-connect
	XconnectId           *string `json:"xconnect_id,omitempty"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CircuitCircuitTerminationRequest CircuitCircuitTerminationRequest

// NewCircuitCircuitTerminationRequest instantiates a new CircuitCircuitTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitCircuitTerminationRequest() *CircuitCircuitTerminationRequest {
	this := CircuitCircuitTerminationRequest{}
	return &this
}

// NewCircuitCircuitTerminationRequestWithDefaults instantiates a new CircuitCircuitTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitCircuitTerminationRequestWithDefaults() *CircuitCircuitTerminationRequest {
	this := CircuitCircuitTerminationRequest{}
	return &this
}

// GetTerminationType returns the TerminationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitCircuitTerminationRequest) GetTerminationType() string {
	if o == nil || IsNil(o.TerminationType.Get()) {
		var ret string
		return ret
	}
	return *o.TerminationType.Get()
}

// GetTerminationTypeOk returns a tuple with the TerminationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTerminationRequest) GetTerminationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationType.Get(), o.TerminationType.IsSet()
}

// HasTerminationType returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationRequest) HasTerminationType() bool {
	if o != nil && o.TerminationType.IsSet() {
		return true
	}

	return false
}

// SetTerminationType gets a reference to the given NullableString and assigns it to the TerminationType field.
func (o *CircuitCircuitTerminationRequest) SetTerminationType(v string) {
	o.TerminationType.Set(&v)
}

// SetTerminationTypeNil sets the value for TerminationType to be an explicit nil
func (o *CircuitCircuitTerminationRequest) SetTerminationTypeNil() {
	o.TerminationType.Set(nil)
}

// UnsetTerminationType ensures that no value is present for TerminationType, not even an explicit nil
func (o *CircuitCircuitTerminationRequest) UnsetTerminationType() {
	o.TerminationType.Unset()
}

// GetTerminationId returns the TerminationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitCircuitTerminationRequest) GetTerminationId() int64 {
	if o == nil || IsNil(o.TerminationId.Get()) {
		var ret int64
		return ret
	}
	return *o.TerminationId.Get()
}

// GetTerminationIdOk returns a tuple with the TerminationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTerminationRequest) GetTerminationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationId.Get(), o.TerminationId.IsSet()
}

// HasTerminationId returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationRequest) HasTerminationId() bool {
	if o != nil && o.TerminationId.IsSet() {
		return true
	}

	return false
}

// SetTerminationId gets a reference to the given NullableInt64 and assigns it to the TerminationId field.
func (o *CircuitCircuitTerminationRequest) SetTerminationId(v int64) {
	o.TerminationId.Set(&v)
}

// SetTerminationIdNil sets the value for TerminationId to be an explicit nil
func (o *CircuitCircuitTerminationRequest) SetTerminationIdNil() {
	o.TerminationId.Set(nil)
}

// UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
func (o *CircuitCircuitTerminationRequest) UnsetTerminationId() {
	o.TerminationId.Unset()
}

// GetPortSpeed returns the PortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitCircuitTerminationRequest) GetPortSpeed() int64 {
	if o == nil || IsNil(o.PortSpeed.Get()) {
		var ret int64
		return ret
	}
	return *o.PortSpeed.Get()
}

// GetPortSpeedOk returns a tuple with the PortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTerminationRequest) GetPortSpeedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortSpeed.Get(), o.PortSpeed.IsSet()
}

// HasPortSpeed returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationRequest) HasPortSpeed() bool {
	if o != nil && o.PortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPortSpeed gets a reference to the given NullableInt64 and assigns it to the PortSpeed field.
func (o *CircuitCircuitTerminationRequest) SetPortSpeed(v int64) {
	o.PortSpeed.Set(&v)
}

// SetPortSpeedNil sets the value for PortSpeed to be an explicit nil
func (o *CircuitCircuitTerminationRequest) SetPortSpeedNil() {
	o.PortSpeed.Set(nil)
}

// UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
func (o *CircuitCircuitTerminationRequest) UnsetPortSpeed() {
	o.PortSpeed.Unset()
}

// GetUpstreamSpeed returns the UpstreamSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitCircuitTerminationRequest) GetUpstreamSpeed() int64 {
	if o == nil || IsNil(o.UpstreamSpeed.Get()) {
		var ret int64
		return ret
	}
	return *o.UpstreamSpeed.Get()
}

// GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTerminationRequest) GetUpstreamSpeedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamSpeed.Get(), o.UpstreamSpeed.IsSet()
}

// HasUpstreamSpeed returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationRequest) HasUpstreamSpeed() bool {
	if o != nil && o.UpstreamSpeed.IsSet() {
		return true
	}

	return false
}

// SetUpstreamSpeed gets a reference to the given NullableInt64 and assigns it to the UpstreamSpeed field.
func (o *CircuitCircuitTerminationRequest) SetUpstreamSpeed(v int64) {
	o.UpstreamSpeed.Set(&v)
}

// SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil
func (o *CircuitCircuitTerminationRequest) SetUpstreamSpeedNil() {
	o.UpstreamSpeed.Set(nil)
}

// UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
func (o *CircuitCircuitTerminationRequest) UnsetUpstreamSpeed() {
	o.UpstreamSpeed.Unset()
}

// GetXconnectId returns the XconnectId field value if set, zero value otherwise.
func (o *CircuitCircuitTerminationRequest) GetXconnectId() string {
	if o == nil || IsNil(o.XconnectId) {
		var ret string
		return ret
	}
	return *o.XconnectId
}

// GetXconnectIdOk returns a tuple with the XconnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTerminationRequest) GetXconnectIdOk() (*string, bool) {
	if o == nil || IsNil(o.XconnectId) {
		return nil, false
	}
	return o.XconnectId, true
}

// HasXconnectId returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationRequest) HasXconnectId() bool {
	if o != nil && !IsNil(o.XconnectId) {
		return true
	}

	return false
}

// SetXconnectId gets a reference to the given string and assigns it to the XconnectId field.
func (o *CircuitCircuitTerminationRequest) SetXconnectId(v string) {
	o.XconnectId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CircuitCircuitTerminationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTerminationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CircuitCircuitTerminationRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CircuitCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircuitCircuitTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CircuitCircuitTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	varCircuitCircuitTerminationRequest := _CircuitCircuitTerminationRequest{}

	err = json.Unmarshal(data, &varCircuitCircuitTerminationRequest)

	if err != nil {
		return err
	}

	*o = CircuitCircuitTerminationRequest(varCircuitCircuitTerminationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "termination_type")
		delete(additionalProperties, "termination_id")
		delete(additionalProperties, "port_speed")
		delete(additionalProperties, "upstream_speed")
		delete(additionalProperties, "xconnect_id")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCircuitCircuitTerminationRequest struct {
	value *CircuitCircuitTerminationRequest
	isSet bool
}

func (v NullableCircuitCircuitTerminationRequest) Get() *CircuitCircuitTerminationRequest {
	return v.value
}

func (v *NullableCircuitCircuitTerminationRequest) Set(val *CircuitCircuitTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitCircuitTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitCircuitTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitCircuitTerminationRequest(val *CircuitCircuitTerminationRequest) *NullableCircuitCircuitTerminationRequest {
	return &NullableCircuitCircuitTerminationRequest{value: val, isSet: true}
}

func (v NullableCircuitCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitCircuitTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
