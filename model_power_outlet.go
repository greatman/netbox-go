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

// checks if the PowerOutlet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerOutlet{}

// PowerOutlet Adds support for custom fields and tags.
type PowerOutlet struct {
	Id         int64               `json:"id"`
	Url        string              `json:"url"`
	DisplayUrl string              `json:"display_url"`
	Display    string              `json:"display"`
	Device     BriefDevice         `json:"device"`
	Module     NullableBriefModule `json:"module,omitempty"`
	Name       string              `json:"name"`
	// Physical label
	Label       *string                    `json:"label,omitempty"`
	Type        NullablePowerOutletType    `json:"type,omitempty"`
	Color       *string                    `json:"color,omitempty" validate:"regexp=^[0-9a-f]{6}$"`
	PowerPort   NullableBriefPowerPort     `json:"power_port,omitempty"`
	FeedLeg     NullablePowerOutletFeedLeg `json:"feed_leg,omitempty"`
	Description *string                    `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected *bool              `json:"mark_connected,omitempty"`
	Cable         NullableBriefCable `json:"cable"`
	CableEnd      string             `json:"cable_end"`
	LinkPeers     []interface{}      `json:"link_peers"`
	// Return the type of the peer link terminations, or None.
	LinkPeersType               NullableString         `json:"link_peers_type"`
	ConnectedEndpoints          []interface{}          `json:"connected_endpoints"`
	ConnectedEndpointsType      NullableString         `json:"connected_endpoints_type"`
	ConnectedEndpointsReachable bool                   `json:"connected_endpoints_reachable"`
	Tags                        []NestedTag            `json:"tags,omitempty"`
	CustomFields                map[string]interface{} `json:"custom_fields,omitempty"`
	Created                     NullableTime           `json:"created"`
	LastUpdated                 NullableTime           `json:"last_updated"`
	Occupied                    bool                   `json:"_occupied"`
	AdditionalProperties        map[string]interface{}
}

type _PowerOutlet PowerOutlet

// NewPowerOutlet instantiates a new PowerOutlet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerOutlet(id int64, url string, displayUrl string, display string, device BriefDevice, name string, cable NullableBriefCable, cableEnd string, linkPeers []interface{}, linkPeersType NullableString, connectedEndpoints []interface{}, connectedEndpointsType NullableString, connectedEndpointsReachable bool, created NullableTime, lastUpdated NullableTime, occupied bool) *PowerOutlet {
	this := PowerOutlet{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Device = device
	this.Name = name
	this.Cable = cable
	this.CableEnd = cableEnd
	this.LinkPeers = linkPeers
	this.LinkPeersType = linkPeersType
	this.ConnectedEndpoints = connectedEndpoints
	this.ConnectedEndpointsType = connectedEndpointsType
	this.ConnectedEndpointsReachable = connectedEndpointsReachable
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Occupied = occupied
	return &this
}

// NewPowerOutletWithDefaults instantiates a new PowerOutlet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerOutletWithDefaults() *PowerOutlet {
	this := PowerOutlet{}
	return &this
}

// GetId returns the Id field value
func (o *PowerOutlet) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PowerOutlet) SetId(v int64) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *PowerOutlet) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PowerOutlet) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *PowerOutlet) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *PowerOutlet) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *PowerOutlet) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *PowerOutlet) SetDisplay(v string) {
	o.Display = v
}

// GetDevice returns the Device field value
func (o *PowerOutlet) GetDevice() BriefDevice {
	if o == nil {
		var ret BriefDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetDeviceOk() (*BriefDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *PowerOutlet) SetDevice(v BriefDevice) {
	o.Device = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutlet) GetModule() BriefModule {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BriefModule
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutlet) GetModuleOk() (*BriefModule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *PowerOutlet) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBriefModule and assigns it to the Module field.
func (o *PowerOutlet) SetModule(v BriefModule) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *PowerOutlet) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *PowerOutlet) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value
func (o *PowerOutlet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PowerOutlet) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerOutlet) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerOutlet) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PowerOutlet) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutlet) GetType() PowerOutletType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret PowerOutletType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutlet) GetTypeOk() (*PowerOutletType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PowerOutlet) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullablePowerOutletType and assigns it to the Type field.
func (o *PowerOutlet) SetType(v PowerOutletType) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *PowerOutlet) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PowerOutlet) UnsetType() {
	o.Type.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *PowerOutlet) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *PowerOutlet) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *PowerOutlet) SetColor(v string) {
	o.Color = &v
}

// GetPowerPort returns the PowerPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutlet) GetPowerPort() BriefPowerPort {
	if o == nil || IsNil(o.PowerPort.Get()) {
		var ret BriefPowerPort
		return ret
	}
	return *o.PowerPort.Get()
}

// GetPowerPortOk returns a tuple with the PowerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutlet) GetPowerPortOk() (*BriefPowerPort, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerPort.Get(), o.PowerPort.IsSet()
}

// HasPowerPort returns a boolean if a field has been set.
func (o *PowerOutlet) HasPowerPort() bool {
	if o != nil && o.PowerPort.IsSet() {
		return true
	}

	return false
}

// SetPowerPort gets a reference to the given NullableBriefPowerPort and assigns it to the PowerPort field.
func (o *PowerOutlet) SetPowerPort(v BriefPowerPort) {
	o.PowerPort.Set(&v)
}

// SetPowerPortNil sets the value for PowerPort to be an explicit nil
func (o *PowerOutlet) SetPowerPortNil() {
	o.PowerPort.Set(nil)
}

// UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
func (o *PowerOutlet) UnsetPowerPort() {
	o.PowerPort.Unset()
}

// GetFeedLeg returns the FeedLeg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutlet) GetFeedLeg() PowerOutletFeedLeg {
	if o == nil || IsNil(o.FeedLeg.Get()) {
		var ret PowerOutletFeedLeg
		return ret
	}
	return *o.FeedLeg.Get()
}

// GetFeedLegOk returns a tuple with the FeedLeg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutlet) GetFeedLegOk() (*PowerOutletFeedLeg, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeedLeg.Get(), o.FeedLeg.IsSet()
}

// HasFeedLeg returns a boolean if a field has been set.
func (o *PowerOutlet) HasFeedLeg() bool {
	if o != nil && o.FeedLeg.IsSet() {
		return true
	}

	return false
}

// SetFeedLeg gets a reference to the given NullablePowerOutletFeedLeg and assigns it to the FeedLeg field.
func (o *PowerOutlet) SetFeedLeg(v PowerOutletFeedLeg) {
	o.FeedLeg.Set(&v)
}

// SetFeedLegNil sets the value for FeedLeg to be an explicit nil
func (o *PowerOutlet) SetFeedLegNil() {
	o.FeedLeg.Set(nil)
}

// UnsetFeedLeg ensures that no value is present for FeedLeg, not even an explicit nil
func (o *PowerOutlet) UnsetFeedLeg() {
	o.FeedLeg.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PowerOutlet) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PowerOutlet) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PowerOutlet) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *PowerOutlet) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *PowerOutlet) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *PowerOutlet) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetCable returns the Cable field value
// If the value is explicit nil, the zero value for BriefCable will be returned
func (o *PowerOutlet) GetCable() BriefCable {
	if o == nil || o.Cable.Get() == nil {
		var ret BriefCable
		return ret
	}

	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutlet) GetCableOk() (*BriefCable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// SetCable sets field value
func (o *PowerOutlet) SetCable(v BriefCable) {
	o.Cable.Set(&v)
}

// GetCableEnd returns the CableEnd field value
func (o *PowerOutlet) GetCableEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CableEnd
}

// GetCableEndOk returns a tuple with the CableEnd field value
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetCableEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CableEnd, true
}

// SetCableEnd sets field value
func (o *PowerOutlet) SetCableEnd(v string) {
	o.CableEnd = v
}

// GetLinkPeers returns the LinkPeers field value
func (o *PowerOutlet) GetLinkPeers() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.LinkPeers
}

// GetLinkPeersOk returns a tuple with the LinkPeers field value
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetLinkPeersOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkPeers, true
}

// SetLinkPeers sets field value
func (o *PowerOutlet) SetLinkPeers(v []interface{}) {
	o.LinkPeers = v
}

// GetLinkPeersType returns the LinkPeersType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PowerOutlet) GetLinkPeersType() string {
	if o == nil || o.LinkPeersType.Get() == nil {
		var ret string
		return ret
	}

	return *o.LinkPeersType.Get()
}

// GetLinkPeersTypeOk returns a tuple with the LinkPeersType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutlet) GetLinkPeersTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkPeersType.Get(), o.LinkPeersType.IsSet()
}

// SetLinkPeersType sets field value
func (o *PowerOutlet) SetLinkPeersType(v string) {
	o.LinkPeersType.Set(&v)
}

// GetConnectedEndpoints returns the ConnectedEndpoints field value
// If the value is explicit nil, the zero value for []interface{} will be returned
func (o *PowerOutlet) GetConnectedEndpoints() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.ConnectedEndpoints
}

// GetConnectedEndpointsOk returns a tuple with the ConnectedEndpoints field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutlet) GetConnectedEndpointsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ConnectedEndpoints) {
		return nil, false
	}
	return o.ConnectedEndpoints, true
}

// SetConnectedEndpoints sets field value
func (o *PowerOutlet) SetConnectedEndpoints(v []interface{}) {
	o.ConnectedEndpoints = v
}

// GetConnectedEndpointsType returns the ConnectedEndpointsType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PowerOutlet) GetConnectedEndpointsType() string {
	if o == nil || o.ConnectedEndpointsType.Get() == nil {
		var ret string
		return ret
	}

	return *o.ConnectedEndpointsType.Get()
}

// GetConnectedEndpointsTypeOk returns a tuple with the ConnectedEndpointsType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutlet) GetConnectedEndpointsTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointsType.Get(), o.ConnectedEndpointsType.IsSet()
}

// SetConnectedEndpointsType sets field value
func (o *PowerOutlet) SetConnectedEndpointsType(v string) {
	o.ConnectedEndpointsType.Set(&v)
}

// GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable field value
func (o *PowerOutlet) GetConnectedEndpointsReachable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ConnectedEndpointsReachable
}

// GetConnectedEndpointsReachableOk returns a tuple with the ConnectedEndpointsReachable field value
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetConnectedEndpointsReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectedEndpointsReachable, true
}

// SetConnectedEndpointsReachable sets field value
func (o *PowerOutlet) SetConnectedEndpointsReachable(v bool) {
	o.ConnectedEndpointsReachable = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PowerOutlet) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PowerOutlet) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *PowerOutlet) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PowerOutlet) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PowerOutlet) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PowerOutlet) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PowerOutlet) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutlet) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *PowerOutlet) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PowerOutlet) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutlet) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *PowerOutlet) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetOccupied returns the Occupied field value
func (o *PowerOutlet) GetOccupied() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value
// and a boolean to check if the value has been set.
func (o *PowerOutlet) GetOccupiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occupied, true
}

// SetOccupied sets field value
func (o *PowerOutlet) SetOccupied(v bool) {
	o.Occupied = v
}

func (o PowerOutlet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerOutlet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["device"] = o.Device
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if o.PowerPort.IsSet() {
		toSerialize["power_port"] = o.PowerPort.Get()
	}
	if o.FeedLeg.IsSet() {
		toSerialize["feed_leg"] = o.FeedLeg.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
	}
	toSerialize["cable"] = o.Cable.Get()
	toSerialize["cable_end"] = o.CableEnd
	toSerialize["link_peers"] = o.LinkPeers
	toSerialize["link_peers_type"] = o.LinkPeersType.Get()
	if o.ConnectedEndpoints != nil {
		toSerialize["connected_endpoints"] = o.ConnectedEndpoints
	}
	toSerialize["connected_endpoints_type"] = o.ConnectedEndpointsType.Get()
	toSerialize["connected_endpoints_reachable"] = o.ConnectedEndpointsReachable
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["_occupied"] = o.Occupied

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PowerOutlet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"device",
		"name",
		"cable",
		"cable_end",
		"link_peers",
		"link_peers_type",
		"connected_endpoints",
		"connected_endpoints_type",
		"connected_endpoints_reachable",
		"created",
		"last_updated",
		"_occupied",
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

	varPowerOutlet := _PowerOutlet{}

	err = json.Unmarshal(data, &varPowerOutlet)

	if err != nil {
		return err
	}

	*o = PowerOutlet(varPowerOutlet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "color")
		delete(additionalProperties, "power_port")
		delete(additionalProperties, "feed_leg")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "cable")
		delete(additionalProperties, "cable_end")
		delete(additionalProperties, "link_peers")
		delete(additionalProperties, "link_peers_type")
		delete(additionalProperties, "connected_endpoints")
		delete(additionalProperties, "connected_endpoints_type")
		delete(additionalProperties, "connected_endpoints_reachable")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "_occupied")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerOutlet struct {
	value *PowerOutlet
	isSet bool
}

func (v NullablePowerOutlet) Get() *PowerOutlet {
	return v.value
}

func (v *NullablePowerOutlet) Set(val *PowerOutlet) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutlet) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutlet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutlet(val *PowerOutlet) *NullablePowerOutlet {
	return &NullablePowerOutlet{value: val, isSet: true}
}

func (v NullablePowerOutlet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutlet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
