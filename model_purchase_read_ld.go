/*
E-fulfilment Shop API

The E-fulfilment Shop API allows you to integrate your service with our warehouse.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package efulfilmentshop

import (
	"encoding/json"
	"time"
)

// PurchaseReadLd Purchase entity
type PurchaseReadLd struct {
	LdContext *string `json:"@context,omitempty"`
	LdId      *string `json:"@id,omitempty"`
	LdType    *string `json:"@type,omitempty"`
	// The purchase ID
	Id *int32 `json:"id,omitempty"`
	// The purchase name
	Name *string `json:"name,omitempty"`
	// The purchase planned delivery date
	PlannedDate CustomTime `json:"plannedDate"`
	// The purchase supplier ID
	SupplierId int32 `json:"supplierId"`
}

// NewPurchaseReadLd instantiates a new PurchaseReadLd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseReadLd(plannedDate time.Time, supplierId int32) *PurchaseReadLd {
	this := PurchaseReadLd{}
	this.PlannedDate = CustomTime{plannedDate}
	this.SupplierId = supplierId
	return &this
}

// NewPurchaseReadLdWithDefaults instantiates a new PurchaseReadLd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseReadLdWithDefaults() *PurchaseReadLd {
	this := PurchaseReadLd{}
	return &this
}

// GetLdContext returns the LdContext field value if set, zero value otherwise.
func (o *PurchaseReadLd) GetLdContext() string {
	if o == nil || o.LdContext == nil {
		var ret string
		return ret
	}
	return *o.LdContext
}

// GetLdContextOk returns a tuple with the LdContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseReadLd) GetLdContextOk() (*string, bool) {
	if o == nil || o.LdContext == nil {
		return nil, false
	}
	return o.LdContext, true
}

// HasLdContext returns a boolean if a field has been set.
func (o *PurchaseReadLd) HasLdContext() bool {
	if o != nil && o.LdContext != nil {
		return true
	}

	return false
}

// SetLdContext gets a reference to the given string and assigns it to the LdContext field.
func (o *PurchaseReadLd) SetLdContext(v string) {
	o.LdContext = &v
}

// GetLdId returns the LdId field value if set, zero value otherwise.
func (o *PurchaseReadLd) GetLdId() string {
	if o == nil || o.LdId == nil {
		var ret string
		return ret
	}
	return *o.LdId
}

// GetLdIdOk returns a tuple with the LdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseReadLd) GetLdIdOk() (*string, bool) {
	if o == nil || o.LdId == nil {
		return nil, false
	}
	return o.LdId, true
}

// HasLdId returns a boolean if a field has been set.
func (o *PurchaseReadLd) HasLdId() bool {
	if o != nil && o.LdId != nil {
		return true
	}

	return false
}

// SetLdId gets a reference to the given string and assigns it to the LdId field.
func (o *PurchaseReadLd) SetLdId(v string) {
	o.LdId = &v
}

// GetLdType returns the LdType field value if set, zero value otherwise.
func (o *PurchaseReadLd) GetLdType() string {
	if o == nil || o.LdType == nil {
		var ret string
		return ret
	}
	return *o.LdType
}

// GetLdTypeOk returns a tuple with the LdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseReadLd) GetLdTypeOk() (*string, bool) {
	if o == nil || o.LdType == nil {
		return nil, false
	}
	return o.LdType, true
}

// HasLdType returns a boolean if a field has been set.
func (o *PurchaseReadLd) HasLdType() bool {
	if o != nil && o.LdType != nil {
		return true
	}

	return false
}

// SetLdType gets a reference to the given string and assigns it to the LdType field.
func (o *PurchaseReadLd) SetLdType(v string) {
	o.LdType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PurchaseReadLd) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseReadLd) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PurchaseReadLd) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PurchaseReadLd) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PurchaseReadLd) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseReadLd) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PurchaseReadLd) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PurchaseReadLd) SetName(v string) {
	o.Name = &v
}

// GetPlannedDate returns the PlannedDate field value
func (o *PurchaseReadLd) GetPlannedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PlannedDate.Time
}

// GetPlannedDateOk returns a tuple with the PlannedDate field value
// and a boolean to check if the value has been set.
func (o *PurchaseReadLd) GetPlannedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlannedDate.Time, true
}

// SetPlannedDate sets field value
func (o *PurchaseReadLd) SetPlannedDate(v time.Time) {
	o.PlannedDate = CustomTime{v}
}

// GetSupplierId returns the SupplierId field value
func (o *PurchaseReadLd) GetSupplierId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SupplierId
}

// GetSupplierIdOk returns a tuple with the SupplierId field value
// and a boolean to check if the value has been set.
func (o *PurchaseReadLd) GetSupplierIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupplierId, true
}

// SetSupplierId sets field value
func (o *PurchaseReadLd) SetSupplierId(v int32) {
	o.SupplierId = v
}

func (o PurchaseReadLd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LdContext != nil {
		toSerialize["ld-context"] = o.LdContext
	}
	if o.LdId != nil {
		toSerialize["ld-id"] = o.LdId
	}
	if o.LdType != nil {
		toSerialize["ld-type"] = o.LdType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["plannedDate"] = o.PlannedDate
	}
	if true {
		toSerialize["supplierId"] = o.SupplierId
	}
	return json.Marshal(toSerialize)
}

type NullablePurchaseReadLd struct {
	value *PurchaseReadLd
	isSet bool
}

func (v NullablePurchaseReadLd) Get() *PurchaseReadLd {
	return v.value
}

func (v *NullablePurchaseReadLd) Set(val *PurchaseReadLd) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseReadLd) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseReadLd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseReadLd(val *PurchaseReadLd) *NullablePurchaseReadLd {
	return &NullablePurchaseReadLd{value: val, isSet: true}
}

func (v NullablePurchaseReadLd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseReadLd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
