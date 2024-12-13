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

// checks if the Site type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Site{}

// Site Adds support for custom fields and tags.
type Site struct {
	Id         int32  `json:"id"`
	Url        string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display    string `json:"display"`
	// Full name of the site
	Name   string                 `json:"name"`
	Slug   string                 `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Status *LocationStatus        `json:"status,omitempty"`
	Region NullableBriefRegion    `json:"region,omitempty"`
	Group  NullableBriefSiteGroup `json:"group,omitempty"`
	Tenant NullableBriefTenant    `json:"tenant,omitempty"`
	// Local facility ID or description
	Facility    *string        `json:"facility,omitempty"`
	TimeZone    NullableString `json:"time_zone,omitempty"`
	Description *string        `json:"description,omitempty"`
	// Physical location of the building
	PhysicalAddress *string `json:"physical_address,omitempty"`
	// If different from the physical address
	ShippingAddress *string `json:"shipping_address,omitempty"`
	// GPS coordinate in decimal format (xx.yyyyyy)
	Latitude NullableFloat64 `json:"latitude,omitempty"`
	// GPS coordinate in decimal format (xx.yyyyyy)
	Longitude            NullableFloat64        `json:"longitude,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Asns                 []ASN                  `json:"asns,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	CircuitCount         int64                  `json:"circuit_count"`
	DeviceCount          int64                  `json:"device_count"`
	PrefixCount          int64                  `json:"prefix_count"`
	RackCount            int64                  `json:"rack_count"`
	VirtualmachineCount  int64                  `json:"virtualmachine_count"`
	VlanCount            int64                  `json:"vlan_count"`
	AdditionalProperties map[string]interface{}
}

type _Site Site

// NewSite instantiates a new Site object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSite(id int32, url string, displayUrl string, display string, name string, slug string, created NullableTime, lastUpdated NullableTime, circuitCount int64, deviceCount int64, prefixCount int64, rackCount int64, virtualmachineCount int64, vlanCount int64) *Site {
	this := Site{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Created = created
	this.LastUpdated = lastUpdated
	this.CircuitCount = circuitCount
	this.DeviceCount = deviceCount
	this.PrefixCount = prefixCount
	this.RackCount = rackCount
	this.VirtualmachineCount = virtualmachineCount
	this.VlanCount = vlanCount
	return &this
}

// NewSiteWithDefaults instantiates a new Site object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteWithDefaults() *Site {
	this := Site{}
	return &this
}

// GetId returns the Id field value
func (o *Site) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Site) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Site) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Site) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Site) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Site) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *Site) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *Site) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *Site) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *Site) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Site) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Site) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *Site) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Site) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Site) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *Site) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Site) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Site) SetSlug(v string) {
	o.Slug = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Site) GetStatus() LocationStatus {
	if o == nil || IsNil(o.Status) {
		var ret LocationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetStatusOk() (*LocationStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Site) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given LocationStatus and assigns it to the Status field.
func (o *Site) SetStatus(v LocationStatus) {
	o.Status = &v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetRegion() BriefRegion {
	if o == nil || IsNil(o.Region.Get()) {
		var ret BriefRegion
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetRegionOk() (*BriefRegion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *Site) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableBriefRegion and assigns it to the Region field.
func (o *Site) SetRegion(v BriefRegion) {
	o.Region.Set(&v)
}

// SetRegionNil sets the value for Region to be an explicit nil
func (o *Site) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *Site) UnsetRegion() {
	o.Region.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetGroup() BriefSiteGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefSiteGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetGroupOk() (*BriefSiteGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *Site) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefSiteGroup and assigns it to the Group field.
func (o *Site) SetGroup(v BriefSiteGroup) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *Site) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *Site) UnsetGroup() {
	o.Group.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Site) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *Site) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Site) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Site) UnsetTenant() {
	o.Tenant.Unset()
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *Site) GetFacility() string {
	if o == nil || IsNil(o.Facility) {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetFacilityOk() (*string, bool) {
	if o == nil || IsNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *Site) HasFacility() bool {
	if o != nil && !IsNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *Site) SetFacility(v string) {
	o.Facility = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone.Get()) {
		var ret string
		return ret
	}
	return *o.TimeZone.Get()
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeZone.Get(), o.TimeZone.IsSet()
}

// HasTimeZone returns a boolean if a field has been set.
func (o *Site) HasTimeZone() bool {
	if o != nil && o.TimeZone.IsSet() {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given NullableString and assigns it to the TimeZone field.
func (o *Site) SetTimeZone(v string) {
	o.TimeZone.Set(&v)
}

// SetTimeZoneNil sets the value for TimeZone to be an explicit nil
func (o *Site) SetTimeZoneNil() {
	o.TimeZone.Set(nil)
}

// UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
func (o *Site) UnsetTimeZone() {
	o.TimeZone.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Site) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Site) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Site) SetDescription(v string) {
	o.Description = &v
}

// GetPhysicalAddress returns the PhysicalAddress field value if set, zero value otherwise.
func (o *Site) GetPhysicalAddress() string {
	if o == nil || IsNil(o.PhysicalAddress) {
		var ret string
		return ret
	}
	return *o.PhysicalAddress
}

// GetPhysicalAddressOk returns a tuple with the PhysicalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetPhysicalAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalAddress) {
		return nil, false
	}
	return o.PhysicalAddress, true
}

// HasPhysicalAddress returns a boolean if a field has been set.
func (o *Site) HasPhysicalAddress() bool {
	if o != nil && !IsNil(o.PhysicalAddress) {
		return true
	}

	return false
}

// SetPhysicalAddress gets a reference to the given string and assigns it to the PhysicalAddress field.
func (o *Site) SetPhysicalAddress(v string) {
	o.PhysicalAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *Site) GetShippingAddress() string {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret string
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetShippingAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *Site) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given string and assigns it to the ShippingAddress field.
func (o *Site) SetShippingAddress(v string) {
	o.ShippingAddress = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *Site) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableFloat64 and assigns it to the Latitude field.
func (o *Site) SetLatitude(v float64) {
	o.Latitude.Set(&v)
}

// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *Site) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *Site) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *Site) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableFloat64 and assigns it to the Longitude field.
func (o *Site) SetLongitude(v float64) {
	o.Longitude.Set(&v)
}

// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *Site) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *Site) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Site) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Site) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Site) SetComments(v string) {
	o.Comments = &v
}

// GetAsns returns the Asns field value if set, zero value otherwise.
func (o *Site) GetAsns() []ASN {
	if o == nil || IsNil(o.Asns) {
		var ret []ASN
		return ret
	}
	return o.Asns
}

// GetAsnsOk returns a tuple with the Asns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetAsnsOk() ([]ASN, bool) {
	if o == nil || IsNil(o.Asns) {
		return nil, false
	}
	return o.Asns, true
}

// HasAsns returns a boolean if a field has been set.
func (o *Site) HasAsns() bool {
	if o != nil && !IsNil(o.Asns) {
		return true
	}

	return false
}

// SetAsns gets a reference to the given []ASN and assigns it to the Asns field.
func (o *Site) SetAsns(v []ASN) {
	o.Asns = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Site) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Site) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Site) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Site) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Site) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Site) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Site) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Site) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Site) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Site) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetCircuitCount returns the CircuitCount field value
func (o *Site) GetCircuitCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CircuitCount
}

// GetCircuitCountOk returns a tuple with the CircuitCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetCircuitCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CircuitCount, true
}

// SetCircuitCount sets field value
func (o *Site) SetCircuitCount(v int64) {
	o.CircuitCount = v
}

// GetDeviceCount returns the DeviceCount field value
func (o *Site) GetDeviceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetDeviceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *Site) SetDeviceCount(v int64) {
	o.DeviceCount = v
}

// GetPrefixCount returns the PrefixCount field value
func (o *Site) GetPrefixCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetPrefixCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixCount, true
}

// SetPrefixCount sets field value
func (o *Site) SetPrefixCount(v int64) {
	o.PrefixCount = v
}

// GetRackCount returns the RackCount field value
func (o *Site) GetRackCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetRackCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RackCount, true
}

// SetRackCount sets field value
func (o *Site) SetRackCount(v int64) {
	o.RackCount = v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value
func (o *Site) GetVirtualmachineCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetVirtualmachineCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualmachineCount, true
}

// SetVirtualmachineCount sets field value
func (o *Site) SetVirtualmachineCount(v int64) {
	o.VirtualmachineCount = v
}

// GetVlanCount returns the VlanCount field value
func (o *Site) GetVlanCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetVlanCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanCount, true
}

// SetVlanCount sets field value
func (o *Site) SetVlanCount(v int64) {
	o.VlanCount = v
}

func (o Site) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Site) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Facility) {
		toSerialize["facility"] = o.Facility
	}
	if o.TimeZone.IsSet() {
		toSerialize["time_zone"] = o.TimeZone.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PhysicalAddress) {
		toSerialize["physical_address"] = o.PhysicalAddress
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Asns) {
		toSerialize["asns"] = o.Asns
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["circuit_count"] = o.CircuitCount
	toSerialize["device_count"] = o.DeviceCount
	toSerialize["prefix_count"] = o.PrefixCount
	toSerialize["rack_count"] = o.RackCount
	toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	toSerialize["vlan_count"] = o.VlanCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Site) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"name",
		"slug",
		"created",
		"last_updated",
		"circuit_count",
		"device_count",
		"prefix_count",
		"rack_count",
		"virtualmachine_count",
		"vlan_count",
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

	varSite := _Site{}

	err = json.Unmarshal(data, &varSite)

	if err != nil {
		return err
	}

	*o = Site(varSite)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "status")
		delete(additionalProperties, "region")
		delete(additionalProperties, "group")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "facility")
		delete(additionalProperties, "time_zone")
		delete(additionalProperties, "description")
		delete(additionalProperties, "physical_address")
		delete(additionalProperties, "shipping_address")
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "asns")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "circuit_count")
		delete(additionalProperties, "device_count")
		delete(additionalProperties, "prefix_count")
		delete(additionalProperties, "rack_count")
		delete(additionalProperties, "virtualmachine_count")
		delete(additionalProperties, "vlan_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSite struct {
	value *Site
	isSet bool
}

func (v NullableSite) Get() *Site {
	return v.value
}

func (v *NullableSite) Set(val *Site) {
	v.value = val
	v.isSet = true
}

func (v NullableSite) IsSet() bool {
	return v.isSet
}

func (v *NullableSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSite(val *Site) *NullableSite {
	return &NullableSite{value: val, isSet: true}
}

func (v NullableSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
