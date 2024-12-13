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

// checks if the WritableRackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableRackRequest{}

// WritableRackRequest Adds support for custom fields and tags.
type WritableRackRequest struct {
	Name       string                            `json:"name"`
	FacilityId NullableString                    `json:"facility_id,omitempty"`
	Site       BriefSiteRequest                  `json:"site"`
	Location   NullableBriefLocationRequest      `json:"location,omitempty"`
	Tenant     NullableBriefTenantRequest        `json:"tenant,omitempty"`
	Status     *PatchedWritableRackRequestStatus `json:"status,omitempty"`
	Role       NullableBriefRackRoleRequest      `json:"role,omitempty"`
	Serial     *string                           `json:"serial,omitempty"`
	// A unique tag used to identify this rack
	AssetTag   NullableString                               `json:"asset_tag,omitempty"`
	RackType   NullableBriefRackTypeRequest                 `json:"rack_type,omitempty"`
	FormFactor NullablePatchedWritableRackRequestFormFactor `json:"form_factor,omitempty"`
	Width      *PatchedWritableRackRequestWidth             `json:"width,omitempty"`
	// Height in rack units
	UHeight *int32 `json:"u_height,omitempty"`
	// Starting unit for rack
	StartingUnit *int32          `json:"starting_unit,omitempty"`
	Weight       NullableFloat64 `json:"weight,omitempty"`
	// Maximum load capacity for the rack
	MaxWeight  NullableInt32                       `json:"max_weight,omitempty"`
	WeightUnit NullableDeviceTypeRequestWeightUnit `json:"weight_unit,omitempty"`
	// Units are numbered top-to-bottom
	DescUnits *bool `json:"desc_units,omitempty"`
	// Outer dimension of rack (width)
	OuterWidth NullableInt32 `json:"outer_width,omitempty"`
	// Outer dimension of rack (depth)
	OuterDepth NullableInt32                               `json:"outer_depth,omitempty"`
	OuterUnit  NullablePatchedWritableRackRequestOuterUnit `json:"outer_unit,omitempty"`
	// Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails.
	MountingDepth        NullableInt32                             `json:"mounting_depth,omitempty"`
	Airflow              NullablePatchedWritableRackRequestAirflow `json:"airflow,omitempty"`
	Description          *string                                   `json:"description,omitempty"`
	Comments             *string                                   `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                        `json:"tags,omitempty"`
	CustomFields         map[string]interface{}                    `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableRackRequest WritableRackRequest

// NewWritableRackRequest instantiates a new WritableRackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableRackRequest(name string, site BriefSiteRequest) *WritableRackRequest {
	this := WritableRackRequest{}
	this.Name = name
	this.Site = site
	return &this
}

// NewWritableRackRequestWithDefaults instantiates a new WritableRackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableRackRequestWithDefaults() *WritableRackRequest {
	this := WritableRackRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableRackRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableRackRequest) SetName(v string) {
	o.Name = v
}

// GetFacilityId returns the FacilityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetFacilityId() string {
	if o == nil || IsNil(o.FacilityId.Get()) {
		var ret string
		return ret
	}
	return *o.FacilityId.Get()
}

// GetFacilityIdOk returns a tuple with the FacilityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetFacilityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacilityId.Get(), o.FacilityId.IsSet()
}

// HasFacilityId returns a boolean if a field has been set.
func (o *WritableRackRequest) HasFacilityId() bool {
	if o != nil && o.FacilityId.IsSet() {
		return true
	}

	return false
}

// SetFacilityId gets a reference to the given NullableString and assigns it to the FacilityId field.
func (o *WritableRackRequest) SetFacilityId(v string) {
	o.FacilityId.Set(&v)
}

// SetFacilityIdNil sets the value for FacilityId to be an explicit nil
func (o *WritableRackRequest) SetFacilityIdNil() {
	o.FacilityId.Set(nil)
}

// UnsetFacilityId ensures that no value is present for FacilityId, not even an explicit nil
func (o *WritableRackRequest) UnsetFacilityId() {
	o.FacilityId.Unset()
}

// GetSite returns the Site field value
func (o *WritableRackRequest) GetSite() BriefSiteRequest {
	if o == nil {
		var ret BriefSiteRequest
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetSiteOk() (*BriefSiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *WritableRackRequest) SetSite(v BriefSiteRequest) {
	o.Site = v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetLocation() BriefLocationRequest {
	if o == nil || IsNil(o.Location.Get()) {
		var ret BriefLocationRequest
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetLocationOk() (*BriefLocationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *WritableRackRequest) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableBriefLocationRequest and assigns it to the Location field.
func (o *WritableRackRequest) SetLocation(v BriefLocationRequest) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *WritableRackRequest) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *WritableRackRequest) UnsetLocation() {
	o.Location.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableRackRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *WritableRackRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableRackRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableRackRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableRackRequest) GetStatus() PatchedWritableRackRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableRackRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetStatusOk() (*PatchedWritableRackRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableRackRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableRackRequestStatus and assigns it to the Status field.
func (o *WritableRackRequest) SetStatus(v PatchedWritableRackRequestStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetRole() BriefRackRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefRackRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetRoleOk() (*BriefRackRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *WritableRackRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefRackRoleRequest and assigns it to the Role field.
func (o *WritableRackRequest) SetRole(v BriefRackRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *WritableRackRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *WritableRackRequest) UnsetRole() {
	o.Role.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *WritableRackRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *WritableRackRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *WritableRackRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *WritableRackRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *WritableRackRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}

// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *WritableRackRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *WritableRackRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetRackType returns the RackType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetRackType() BriefRackTypeRequest {
	if o == nil || IsNil(o.RackType.Get()) {
		var ret BriefRackTypeRequest
		return ret
	}
	return *o.RackType.Get()
}

// GetRackTypeOk returns a tuple with the RackType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetRackTypeOk() (*BriefRackTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.RackType.Get(), o.RackType.IsSet()
}

// HasRackType returns a boolean if a field has been set.
func (o *WritableRackRequest) HasRackType() bool {
	if o != nil && o.RackType.IsSet() {
		return true
	}

	return false
}

// SetRackType gets a reference to the given NullableBriefRackTypeRequest and assigns it to the RackType field.
func (o *WritableRackRequest) SetRackType(v BriefRackTypeRequest) {
	o.RackType.Set(&v)
}

// SetRackTypeNil sets the value for RackType to be an explicit nil
func (o *WritableRackRequest) SetRackTypeNil() {
	o.RackType.Set(nil)
}

// UnsetRackType ensures that no value is present for RackType, not even an explicit nil
func (o *WritableRackRequest) UnsetRackType() {
	o.RackType.Unset()
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetFormFactor() PatchedWritableRackRequestFormFactor {
	if o == nil || IsNil(o.FormFactor.Get()) {
		var ret PatchedWritableRackRequestFormFactor
		return ret
	}
	return *o.FormFactor.Get()
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetFormFactorOk() (*PatchedWritableRackRequestFormFactor, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormFactor.Get(), o.FormFactor.IsSet()
}

// HasFormFactor returns a boolean if a field has been set.
func (o *WritableRackRequest) HasFormFactor() bool {
	if o != nil && o.FormFactor.IsSet() {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given NullablePatchedWritableRackRequestFormFactor and assigns it to the FormFactor field.
func (o *WritableRackRequest) SetFormFactor(v PatchedWritableRackRequestFormFactor) {
	o.FormFactor.Set(&v)
}

// SetFormFactorNil sets the value for FormFactor to be an explicit nil
func (o *WritableRackRequest) SetFormFactorNil() {
	o.FormFactor.Set(nil)
}

// UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
func (o *WritableRackRequest) UnsetFormFactor() {
	o.FormFactor.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *WritableRackRequest) GetWidth() PatchedWritableRackRequestWidth {
	if o == nil || IsNil(o.Width) {
		var ret PatchedWritableRackRequestWidth
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetWidthOk() (*PatchedWritableRackRequestWidth, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *WritableRackRequest) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given PatchedWritableRackRequestWidth and assigns it to the Width field.
func (o *WritableRackRequest) SetWidth(v PatchedWritableRackRequestWidth) {
	o.Width = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *WritableRackRequest) GetUHeight() int32 {
	if o == nil || IsNil(o.UHeight) {
		var ret int32
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetUHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *WritableRackRequest) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given int32 and assigns it to the UHeight field.
func (o *WritableRackRequest) SetUHeight(v int32) {
	o.UHeight = &v
}

// GetStartingUnit returns the StartingUnit field value if set, zero value otherwise.
func (o *WritableRackRequest) GetStartingUnit() int32 {
	if o == nil || IsNil(o.StartingUnit) {
		var ret int32
		return ret
	}
	return *o.StartingUnit
}

// GetStartingUnitOk returns a tuple with the StartingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetStartingUnitOk() (*int32, bool) {
	if o == nil || IsNil(o.StartingUnit) {
		return nil, false
	}
	return o.StartingUnit, true
}

// HasStartingUnit returns a boolean if a field has been set.
func (o *WritableRackRequest) HasStartingUnit() bool {
	if o != nil && !IsNil(o.StartingUnit) {
		return true
	}

	return false
}

// SetStartingUnit gets a reference to the given int32 and assigns it to the StartingUnit field.
func (o *WritableRackRequest) SetStartingUnit(v int32) {
	o.StartingUnit = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *WritableRackRequest) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *WritableRackRequest) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil
func (o *WritableRackRequest) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *WritableRackRequest) UnsetWeight() {
	o.Weight.Unset()
}

// GetMaxWeight returns the MaxWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetMaxWeight() int32 {
	if o == nil || IsNil(o.MaxWeight.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxWeight.Get()
}

// GetMaxWeightOk returns a tuple with the MaxWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetMaxWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxWeight.Get(), o.MaxWeight.IsSet()
}

// HasMaxWeight returns a boolean if a field has been set.
func (o *WritableRackRequest) HasMaxWeight() bool {
	if o != nil && o.MaxWeight.IsSet() {
		return true
	}

	return false
}

// SetMaxWeight gets a reference to the given NullableInt32 and assigns it to the MaxWeight field.
func (o *WritableRackRequest) SetMaxWeight(v int32) {
	o.MaxWeight.Set(&v)
}

// SetMaxWeightNil sets the value for MaxWeight to be an explicit nil
func (o *WritableRackRequest) SetMaxWeightNil() {
	o.MaxWeight.Set(nil)
}

// UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
func (o *WritableRackRequest) UnsetMaxWeight() {
	o.MaxWeight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetWeightUnit() DeviceTypeRequestWeightUnit {
	if o == nil || IsNil(o.WeightUnit.Get()) {
		var ret DeviceTypeRequestWeightUnit
		return ret
	}
	return *o.WeightUnit.Get()
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetWeightUnitOk() (*DeviceTypeRequestWeightUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeightUnit.Get(), o.WeightUnit.IsSet()
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *WritableRackRequest) HasWeightUnit() bool {
	if o != nil && o.WeightUnit.IsSet() {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given NullableDeviceTypeRequestWeightUnit and assigns it to the WeightUnit field.
func (o *WritableRackRequest) SetWeightUnit(v DeviceTypeRequestWeightUnit) {
	o.WeightUnit.Set(&v)
}

// SetWeightUnitNil sets the value for WeightUnit to be an explicit nil
func (o *WritableRackRequest) SetWeightUnitNil() {
	o.WeightUnit.Set(nil)
}

// UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
func (o *WritableRackRequest) UnsetWeightUnit() {
	o.WeightUnit.Unset()
}

// GetDescUnits returns the DescUnits field value if set, zero value otherwise.
func (o *WritableRackRequest) GetDescUnits() bool {
	if o == nil || IsNil(o.DescUnits) {
		var ret bool
		return ret
	}
	return *o.DescUnits
}

// GetDescUnitsOk returns a tuple with the DescUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetDescUnitsOk() (*bool, bool) {
	if o == nil || IsNil(o.DescUnits) {
		return nil, false
	}
	return o.DescUnits, true
}

// HasDescUnits returns a boolean if a field has been set.
func (o *WritableRackRequest) HasDescUnits() bool {
	if o != nil && !IsNil(o.DescUnits) {
		return true
	}

	return false
}

// SetDescUnits gets a reference to the given bool and assigns it to the DescUnits field.
func (o *WritableRackRequest) SetDescUnits(v bool) {
	o.DescUnits = &v
}

// GetOuterWidth returns the OuterWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetOuterWidth() int32 {
	if o == nil || IsNil(o.OuterWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterWidth.Get()
}

// GetOuterWidthOk returns a tuple with the OuterWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetOuterWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterWidth.Get(), o.OuterWidth.IsSet()
}

// HasOuterWidth returns a boolean if a field has been set.
func (o *WritableRackRequest) HasOuterWidth() bool {
	if o != nil && o.OuterWidth.IsSet() {
		return true
	}

	return false
}

// SetOuterWidth gets a reference to the given NullableInt32 and assigns it to the OuterWidth field.
func (o *WritableRackRequest) SetOuterWidth(v int32) {
	o.OuterWidth.Set(&v)
}

// SetOuterWidthNil sets the value for OuterWidth to be an explicit nil
func (o *WritableRackRequest) SetOuterWidthNil() {
	o.OuterWidth.Set(nil)
}

// UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
func (o *WritableRackRequest) UnsetOuterWidth() {
	o.OuterWidth.Unset()
}

// GetOuterDepth returns the OuterDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetOuterDepth() int32 {
	if o == nil || IsNil(o.OuterDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterDepth.Get()
}

// GetOuterDepthOk returns a tuple with the OuterDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetOuterDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterDepth.Get(), o.OuterDepth.IsSet()
}

// HasOuterDepth returns a boolean if a field has been set.
func (o *WritableRackRequest) HasOuterDepth() bool {
	if o != nil && o.OuterDepth.IsSet() {
		return true
	}

	return false
}

// SetOuterDepth gets a reference to the given NullableInt32 and assigns it to the OuterDepth field.
func (o *WritableRackRequest) SetOuterDepth(v int32) {
	o.OuterDepth.Set(&v)
}

// SetOuterDepthNil sets the value for OuterDepth to be an explicit nil
func (o *WritableRackRequest) SetOuterDepthNil() {
	o.OuterDepth.Set(nil)
}

// UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
func (o *WritableRackRequest) UnsetOuterDepth() {
	o.OuterDepth.Unset()
}

// GetOuterUnit returns the OuterUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetOuterUnit() PatchedWritableRackRequestOuterUnit {
	if o == nil || IsNil(o.OuterUnit.Get()) {
		var ret PatchedWritableRackRequestOuterUnit
		return ret
	}
	return *o.OuterUnit.Get()
}

// GetOuterUnitOk returns a tuple with the OuterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetOuterUnitOk() (*PatchedWritableRackRequestOuterUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterUnit.Get(), o.OuterUnit.IsSet()
}

// HasOuterUnit returns a boolean if a field has been set.
func (o *WritableRackRequest) HasOuterUnit() bool {
	if o != nil && o.OuterUnit.IsSet() {
		return true
	}

	return false
}

// SetOuterUnit gets a reference to the given NullablePatchedWritableRackRequestOuterUnit and assigns it to the OuterUnit field.
func (o *WritableRackRequest) SetOuterUnit(v PatchedWritableRackRequestOuterUnit) {
	o.OuterUnit.Set(&v)
}

// SetOuterUnitNil sets the value for OuterUnit to be an explicit nil
func (o *WritableRackRequest) SetOuterUnitNil() {
	o.OuterUnit.Set(nil)
}

// UnsetOuterUnit ensures that no value is present for OuterUnit, not even an explicit nil
func (o *WritableRackRequest) UnsetOuterUnit() {
	o.OuterUnit.Unset()
}

// GetMountingDepth returns the MountingDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetMountingDepth() int32 {
	if o == nil || IsNil(o.MountingDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.MountingDepth.Get()
}

// GetMountingDepthOk returns a tuple with the MountingDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetMountingDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountingDepth.Get(), o.MountingDepth.IsSet()
}

// HasMountingDepth returns a boolean if a field has been set.
func (o *WritableRackRequest) HasMountingDepth() bool {
	if o != nil && o.MountingDepth.IsSet() {
		return true
	}

	return false
}

// SetMountingDepth gets a reference to the given NullableInt32 and assigns it to the MountingDepth field.
func (o *WritableRackRequest) SetMountingDepth(v int32) {
	o.MountingDepth.Set(&v)
}

// SetMountingDepthNil sets the value for MountingDepth to be an explicit nil
func (o *WritableRackRequest) SetMountingDepthNil() {
	o.MountingDepth.Set(nil)
}

// UnsetMountingDepth ensures that no value is present for MountingDepth, not even an explicit nil
func (o *WritableRackRequest) UnsetMountingDepth() {
	o.MountingDepth.Unset()
}

// GetAirflow returns the Airflow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRackRequest) GetAirflow() PatchedWritableRackRequestAirflow {
	if o == nil || IsNil(o.Airflow.Get()) {
		var ret PatchedWritableRackRequestAirflow
		return ret
	}
	return *o.Airflow.Get()
}

// GetAirflowOk returns a tuple with the Airflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRackRequest) GetAirflowOk() (*PatchedWritableRackRequestAirflow, bool) {
	if o == nil {
		return nil, false
	}
	return o.Airflow.Get(), o.Airflow.IsSet()
}

// HasAirflow returns a boolean if a field has been set.
func (o *WritableRackRequest) HasAirflow() bool {
	if o != nil && o.Airflow.IsSet() {
		return true
	}

	return false
}

// SetAirflow gets a reference to the given NullablePatchedWritableRackRequestAirflow and assigns it to the Airflow field.
func (o *WritableRackRequest) SetAirflow(v PatchedWritableRackRequestAirflow) {
	o.Airflow.Set(&v)
}

// SetAirflowNil sets the value for Airflow to be an explicit nil
func (o *WritableRackRequest) SetAirflowNil() {
	o.Airflow.Set(nil)
}

// UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
func (o *WritableRackRequest) UnsetAirflow() {
	o.Airflow.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableRackRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableRackRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableRackRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableRackRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableRackRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableRackRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableRackRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableRackRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableRackRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableRackRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRackRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableRackRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableRackRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableRackRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableRackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.FacilityId.IsSet() {
		toSerialize["facility_id"] = o.FacilityId.Get()
	}
	toSerialize["site"] = o.Site
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if o.RackType.IsSet() {
		toSerialize["rack_type"] = o.RackType.Get()
	}
	if o.FormFactor.IsSet() {
		toSerialize["form_factor"] = o.FormFactor.Get()
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.StartingUnit) {
		toSerialize["starting_unit"] = o.StartingUnit
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if o.MaxWeight.IsSet() {
		toSerialize["max_weight"] = o.MaxWeight.Get()
	}
	if o.WeightUnit.IsSet() {
		toSerialize["weight_unit"] = o.WeightUnit.Get()
	}
	if !IsNil(o.DescUnits) {
		toSerialize["desc_units"] = o.DescUnits
	}
	if o.OuterWidth.IsSet() {
		toSerialize["outer_width"] = o.OuterWidth.Get()
	}
	if o.OuterDepth.IsSet() {
		toSerialize["outer_depth"] = o.OuterDepth.Get()
	}
	if o.OuterUnit.IsSet() {
		toSerialize["outer_unit"] = o.OuterUnit.Get()
	}
	if o.MountingDepth.IsSet() {
		toSerialize["mounting_depth"] = o.MountingDepth.Get()
	}
	if o.Airflow.IsSet() {
		toSerialize["airflow"] = o.Airflow.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
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

func (o *WritableRackRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"site",
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

	varWritableRackRequest := _WritableRackRequest{}

	err = json.Unmarshal(data, &varWritableRackRequest)

	if err != nil {
		return err
	}

	*o = WritableRackRequest(varWritableRackRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "facility_id")
		delete(additionalProperties, "site")
		delete(additionalProperties, "location")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "rack_type")
		delete(additionalProperties, "form_factor")
		delete(additionalProperties, "width")
		delete(additionalProperties, "u_height")
		delete(additionalProperties, "starting_unit")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "max_weight")
		delete(additionalProperties, "weight_unit")
		delete(additionalProperties, "desc_units")
		delete(additionalProperties, "outer_width")
		delete(additionalProperties, "outer_depth")
		delete(additionalProperties, "outer_unit")
		delete(additionalProperties, "mounting_depth")
		delete(additionalProperties, "airflow")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableRackRequest struct {
	value *WritableRackRequest
	isSet bool
}

func (v NullableWritableRackRequest) Get() *WritableRackRequest {
	return v.value
}

func (v *NullableWritableRackRequest) Set(val *WritableRackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableRackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableRackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableRackRequest(val *WritableRackRequest) *NullableWritableRackRequest {
	return &NullableWritableRackRequest{value: val, isSet: true}
}

func (v NullableWritableRackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableRackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
