/*
 * E-fulfilment Shop API
 *
 * The E-fulfilment Shop API allows you to integrate your service with our warehouse.
 *
 * API version: 1.3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package efulfilmentshop

import (
	"encoding/json"
	"time"
)

// PurchaseLinejsonldWrite Purchase Line entity
type PurchaseLinejsonldWrite struct {
	JsonContext *string `json:"json_context,omitempty"`
	JsonId *string `json:"json_id,omitempty"`
	JsonType *string `json:"json_type,omitempty"`
	// The purchase line description
	Description string `json:"description"`
	// The purchase line planned delivery date
	PlannedDate *time.Time `json:"plannedDate,omitempty"`
	// The product ID
	ProductId int32 `json:"productId"`
	// The purchase ID
	PurchaseId int32 `json:"purchaseId"`
	// The purchase line quantity
	Quantity int32 `json:"quantity"`
}

// NewPurchaseLinejsonldWrite instantiates a new PurchaseLinejsonldWrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseLinejsonldWrite(description string, productId int32, purchaseId int32, quantity int32, ) *PurchaseLinejsonldWrite {
	this := PurchaseLinejsonldWrite{}
	this.Description = description
	this.ProductId = productId
	this.PurchaseId = purchaseId
	this.Quantity = quantity
	return &this
}

// NewPurchaseLinejsonldWriteWithDefaults instantiates a new PurchaseLinejsonldWrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseLinejsonldWriteWithDefaults() *PurchaseLinejsonldWrite {
	this := PurchaseLinejsonldWrite{}
	return &this
}

// GetJsonContext returns the JsonContext field value if set, zero value otherwise.
func (o *PurchaseLinejsonldWrite) GetJsonContext() string {
	if o == nil || o.JsonContext == nil {
		var ret string
		return ret
	}
	return *o.JsonContext
}

// GetJsonContextOk returns a tuple with the JsonContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseLinejsonldWrite) GetJsonContextOk() (*string, bool) {
	if o == nil || o.JsonContext == nil {
		return nil, false
	}
	return o.JsonContext, true
}

// HasJsonContext returns a boolean if a field has been set.
func (o *PurchaseLinejsonldWrite) HasJsonContext() bool {
	if o != nil && o.JsonContext != nil {
		return true
	}

	return false
}

// SetJsonContext gets a reference to the given string and assigns it to the JsonContext field.
func (o *PurchaseLinejsonldWrite) SetJsonContext(v string) {
	o.JsonContext = &v
}

// GetJsonId returns the JsonId field value if set, zero value otherwise.
func (o *PurchaseLinejsonldWrite) GetJsonId() string {
	if o == nil || o.JsonId == nil {
		var ret string
		return ret
	}
	return *o.JsonId
}

// GetJsonIdOk returns a tuple with the JsonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseLinejsonldWrite) GetJsonIdOk() (*string, bool) {
	if o == nil || o.JsonId == nil {
		return nil, false
	}
	return o.JsonId, true
}

// HasJsonId returns a boolean if a field has been set.
func (o *PurchaseLinejsonldWrite) HasJsonId() bool {
	if o != nil && o.JsonId != nil {
		return true
	}

	return false
}

// SetJsonId gets a reference to the given string and assigns it to the JsonId field.
func (o *PurchaseLinejsonldWrite) SetJsonId(v string) {
	o.JsonId = &v
}

// GetJsonType returns the JsonType field value if set, zero value otherwise.
func (o *PurchaseLinejsonldWrite) GetJsonType() string {
	if o == nil || o.JsonType == nil {
		var ret string
		return ret
	}
	return *o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseLinejsonldWrite) GetJsonTypeOk() (*string, bool) {
	if o == nil || o.JsonType == nil {
		return nil, false
	}
	return o.JsonType, true
}

// HasJsonType returns a boolean if a field has been set.
func (o *PurchaseLinejsonldWrite) HasJsonType() bool {
	if o != nil && o.JsonType != nil {
		return true
	}

	return false
}

// SetJsonType gets a reference to the given string and assigns it to the JsonType field.
func (o *PurchaseLinejsonldWrite) SetJsonType(v string) {
	o.JsonType = &v
}

// GetDescription returns the Description field value
func (o *PurchaseLinejsonldWrite) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PurchaseLinejsonldWrite) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PurchaseLinejsonldWrite) SetDescription(v string) {
	o.Description = v
}

// GetPlannedDate returns the PlannedDate field value if set, zero value otherwise.
func (o *PurchaseLinejsonldWrite) GetPlannedDate() time.Time {
	if o == nil || o.PlannedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.PlannedDate
}

// GetPlannedDateOk returns a tuple with the PlannedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseLinejsonldWrite) GetPlannedDateOk() (*time.Time, bool) {
	if o == nil || o.PlannedDate == nil {
		return nil, false
	}
	return o.PlannedDate, true
}

// HasPlannedDate returns a boolean if a field has been set.
func (o *PurchaseLinejsonldWrite) HasPlannedDate() bool {
	if o != nil && o.PlannedDate != nil {
		return true
	}

	return false
}

// SetPlannedDate gets a reference to the given time.Time and assigns it to the PlannedDate field.
func (o *PurchaseLinejsonldWrite) SetPlannedDate(v time.Time) {
	o.PlannedDate = &v
}

// GetProductId returns the ProductId field value
func (o *PurchaseLinejsonldWrite) GetProductId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *PurchaseLinejsonldWrite) GetProductIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *PurchaseLinejsonldWrite) SetProductId(v int32) {
	o.ProductId = v
}

// GetPurchaseId returns the PurchaseId field value
func (o *PurchaseLinejsonldWrite) GetPurchaseId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.PurchaseId
}

// GetPurchaseIdOk returns a tuple with the PurchaseId field value
// and a boolean to check if the value has been set.
func (o *PurchaseLinejsonldWrite) GetPurchaseIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PurchaseId, true
}

// SetPurchaseId sets field value
func (o *PurchaseLinejsonldWrite) SetPurchaseId(v int32) {
	o.PurchaseId = v
}

// GetQuantity returns the Quantity field value
func (o *PurchaseLinejsonldWrite) GetQuantity() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *PurchaseLinejsonldWrite) GetQuantityOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *PurchaseLinejsonldWrite) SetQuantity(v int32) {
	o.Quantity = v
}

func (o PurchaseLinejsonldWrite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JsonContext != nil {
		toSerialize["json_context"] = o.JsonContext
	}
	if o.JsonId != nil {
		toSerialize["json_id"] = o.JsonId
	}
	if o.JsonType != nil {
		toSerialize["json_type"] = o.JsonType
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.PlannedDate != nil {
		toSerialize["plannedDate"] = o.PlannedDate
	}
	if true {
		toSerialize["productId"] = o.ProductId
	}
	if true {
		toSerialize["purchaseId"] = o.PurchaseId
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	return json.Marshal(toSerialize)
}

type NullablePurchaseLinejsonldWrite struct {
	value *PurchaseLinejsonldWrite
	isSet bool
}

func (v NullablePurchaseLinejsonldWrite) Get() *PurchaseLinejsonldWrite {
	return v.value
}

func (v *NullablePurchaseLinejsonldWrite) Set(val *PurchaseLinejsonldWrite) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseLinejsonldWrite) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseLinejsonldWrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseLinejsonldWrite(val *PurchaseLinejsonldWrite) *NullablePurchaseLinejsonldWrite {
	return &NullablePurchaseLinejsonldWrite{value: val, isSet: true}
}

func (v NullablePurchaseLinejsonldWrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseLinejsonldWrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


