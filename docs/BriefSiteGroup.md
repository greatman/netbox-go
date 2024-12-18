# BriefSiteGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**SiteCount** | **int64** |  | [readonly] [default to 0]
**Depth** | **int64** |  | [readonly] 

## Methods

### NewBriefSiteGroup

`func NewBriefSiteGroup(id int64, url string, display string, name string, slug string, siteCount int64, depth int64, ) *BriefSiteGroup`

NewBriefSiteGroup instantiates a new BriefSiteGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefSiteGroupWithDefaults

`func NewBriefSiteGroupWithDefaults() *BriefSiteGroup`

NewBriefSiteGroupWithDefaults instantiates a new BriefSiteGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefSiteGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefSiteGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefSiteGroup) SetId(v int64)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefSiteGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefSiteGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefSiteGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefSiteGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefSiteGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefSiteGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefSiteGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefSiteGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefSiteGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefSiteGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefSiteGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefSiteGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefSiteGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefSiteGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefSiteGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefSiteGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSiteCount

`func (o *BriefSiteGroup) GetSiteCount() int64`

GetSiteCount returns the SiteCount field if non-nil, zero value otherwise.

### GetSiteCountOk

`func (o *BriefSiteGroup) GetSiteCountOk() (*int64, bool)`

GetSiteCountOk returns a tuple with the SiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCount

`func (o *BriefSiteGroup) SetSiteCount(v int64)`

SetSiteCount sets SiteCount field to given value.


### GetDepth

`func (o *BriefSiteGroup) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *BriefSiteGroup) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *BriefSiteGroup) SetDepth(v int64)`

SetDepth sets Depth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


