/*
E-fulfilment Shop API

The E-fulfilment Shop API allows you to integrate your service with our warehouse.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package efulfilmentshop

import (
	"encoding/json"
)

// SaleLineWriteLd Sale Line entity
type SaleLineWriteLd struct {
	LdContext *string `json:"@context,omitempty"`
	LdId      *string `json:"@id,omitempty"`
	LdType    *string `json:"@type,omitempty"`
	// The sale line description
	Description string `json:"description"`
	// The product ID
	ProductId int32 `json:"productId"`
	// The sale line quantity
	Quantity int32 `json:"quantity"`
	// The sale ID
	SaleId int32 `json:"saleId"`
}

// NewSaleLineWriteLd instantiates a new SaleLineWriteLd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaleLineWriteLd(description string, productId int32, quantity int32, saleId int32) *SaleLineWriteLd {
	this := SaleLineWriteLd{}
	this.Description = description
	this.ProductId = productId
	this.Quantity = quantity
	this.SaleId = saleId
	return &this
}

// NewSaleLineWriteLdWithDefaults instantiates a new SaleLineWriteLd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaleLineWriteLdWithDefaults() *SaleLineWriteLd {
	this := SaleLineWriteLd{}
	return &this
}

// GetLdContext returns the LdContext field value if set, zero value otherwise.
func (o *SaleLineWriteLd) GetLdContext() string {
	if o == nil || o.LdContext == nil {
		var ret string
		return ret
	}
	return *o.LdContext
}

// GetLdContextOk returns a tuple with the LdContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleLineWriteLd) GetLdContextOk() (*string, bool) {
	if o == nil || o.LdContext == nil {
		return nil, false
	}
	return o.LdContext, true
}

// HasLdContext returns a boolean if a field has been set.
func (o *SaleLineWriteLd) HasLdContext() bool {
	if o != nil && o.LdContext != nil {
		return true
	}

	return false
}

// SetLdContext gets a reference to the given string and assigns it to the LdContext field.
func (o *SaleLineWriteLd) SetLdContext(v string) {
	o.LdContext = &v
}

// GetLdId returns the LdId field value if set, zero value otherwise.
func (o *SaleLineWriteLd) GetLdId() string {
	if o == nil || o.LdId == nil {
		var ret string
		return ret
	}
	return *o.LdId
}

// GetLdIdOk returns a tuple with the LdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleLineWriteLd) GetLdIdOk() (*string, bool) {
	if o == nil || o.LdId == nil {
		return nil, false
	}
	return o.LdId, true
}

// HasLdId returns a boolean if a field has been set.
func (o *SaleLineWriteLd) HasLdId() bool {
	if o != nil && o.LdId != nil {
		return true
	}

	return false
}

// SetLdId gets a reference to the given string and assigns it to the LdId field.
func (o *SaleLineWriteLd) SetLdId(v string) {
	o.LdId = &v
}

// GetLdType returns the LdType field value if set, zero value otherwise.
func (o *SaleLineWriteLd) GetLdType() string {
	if o == nil || o.LdType == nil {
		var ret string
		return ret
	}
	return *o.LdType
}

// GetLdTypeOk returns a tuple with the LdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleLineWriteLd) GetLdTypeOk() (*string, bool) {
	if o == nil || o.LdType == nil {
		return nil, false
	}
	return o.LdType, true
}

// HasLdType returns a boolean if a field has been set.
func (o *SaleLineWriteLd) HasLdType() bool {
	if o != nil && o.LdType != nil {
		return true
	}

	return false
}

// SetLdType gets a reference to the given string and assigns it to the LdType field.
func (o *SaleLineWriteLd) SetLdType(v string) {
	o.LdType = &v
}

// GetDescription returns the Description field value
func (o *SaleLineWriteLd) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SaleLineWriteLd) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SaleLineWriteLd) SetDescription(v string) {
	o.Description = v
}

// GetProductId returns the ProductId field value
func (o *SaleLineWriteLd) GetProductId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *SaleLineWriteLd) GetProductIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *SaleLineWriteLd) SetProductId(v int32) {
	o.ProductId = v
}

// GetQuantity returns the Quantity field value
func (o *SaleLineWriteLd) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *SaleLineWriteLd) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *SaleLineWriteLd) SetQuantity(v int32) {
	o.Quantity = v
}

// GetSaleId returns the SaleId field value
func (o *SaleLineWriteLd) GetSaleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SaleId
}

// GetSaleIdOk returns a tuple with the SaleId field value
// and a boolean to check if the value has been set.
func (o *SaleLineWriteLd) GetSaleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SaleId, true
}

// SetSaleId sets field value
func (o *SaleLineWriteLd) SetSaleId(v int32) {
	o.SaleId = v
}

func (o SaleLineWriteLd) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["productId"] = o.ProductId
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["saleId"] = o.SaleId
	}
	return json.Marshal(toSerialize)
}

type NullableSaleLineWriteLd struct {
	value *SaleLineWriteLd
	isSet bool
}

func (v NullableSaleLineWriteLd) Get() *SaleLineWriteLd {
	return v.value
}

func (v *NullableSaleLineWriteLd) Set(val *SaleLineWriteLd) {
	v.value = val
	v.isSet = true
}

func (v NullableSaleLineWriteLd) IsSet() bool {
	return v.isSet
}

func (v *NullableSaleLineWriteLd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaleLineWriteLd(val *SaleLineWriteLd) *NullableSaleLineWriteLd {
	return &NullableSaleLineWriteLd{value: val, isSet: true}
}

func (v NullableSaleLineWriteLd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaleLineWriteLd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
