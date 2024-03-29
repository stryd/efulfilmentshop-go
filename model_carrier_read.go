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

// CarrierRead Carrier entity
type CarrierRead struct {
	// The carrier ID
	Id *int32 `json:"id,omitempty"`
	// The carrier name
	Name *string `json:"name,omitempty"`
}

// NewCarrierRead instantiates a new CarrierRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierRead() *CarrierRead {
	this := CarrierRead{}
	return &this
}

// NewCarrierReadWithDefaults instantiates a new CarrierRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierReadWithDefaults() *CarrierRead {
	this := CarrierRead{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CarrierRead) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierRead) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CarrierRead) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CarrierRead) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CarrierRead) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierRead) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CarrierRead) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CarrierRead) SetName(v string) {
	o.Name = &v
}

func (o CarrierRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCarrierRead struct {
	value *CarrierRead
	isSet bool
}

func (v NullableCarrierRead) Get() *CarrierRead {
	return v.value
}

func (v *NullableCarrierRead) Set(val *CarrierRead) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierRead) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierRead(val *CarrierRead) *NullableCarrierRead {
	return &NullableCarrierRead{value: val, isSet: true}
}

func (v NullableCarrierRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarrierRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


