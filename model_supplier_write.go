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

// SupplierWrite Supplier entity
type SupplierWrite struct {
	// The supplier city
	City *string `json:"city,omitempty"`
	// The supplier country code (ISO 3166-1 alpha-2 format)
	CountryCode *string `json:"countryCode,omitempty"`
	// The supplier email
	Email *string `json:"email,omitempty"`
	// The supplier house number
	HouseNumber *string `json:"houseNumber,omitempty"`
	// The supplier name
	Name string `json:"name"`
	// The supplier phone
	Phone *string `json:"phone,omitempty"`
	// The supplier street
	Street *string `json:"street,omitempty"`
	// The supplier street 2
	Street2 *string `json:"street2,omitempty"`
	// The supplier zip
	Zip *string `json:"zip,omitempty"`
}

// NewSupplierWrite instantiates a new SupplierWrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplierWrite(name string) *SupplierWrite {
	this := SupplierWrite{}
	this.Name = name
	return &this
}

// NewSupplierWriteWithDefaults instantiates a new SupplierWrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplierWriteWithDefaults() *SupplierWrite {
	this := SupplierWrite{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *SupplierWrite) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWrite) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *SupplierWrite) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *SupplierWrite) SetCity(v string) {
	o.City = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *SupplierWrite) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWrite) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *SupplierWrite) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *SupplierWrite) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SupplierWrite) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWrite) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SupplierWrite) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SupplierWrite) SetEmail(v string) {
	o.Email = &v
}

// GetHouseNumber returns the HouseNumber field value if set, zero value otherwise.
func (o *SupplierWrite) GetHouseNumber() string {
	if o == nil || o.HouseNumber == nil {
		var ret string
		return ret
	}
	return *o.HouseNumber
}

// GetHouseNumberOk returns a tuple with the HouseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWrite) GetHouseNumberOk() (*string, bool) {
	if o == nil || o.HouseNumber == nil {
		return nil, false
	}
	return o.HouseNumber, true
}

// HasHouseNumber returns a boolean if a field has been set.
func (o *SupplierWrite) HasHouseNumber() bool {
	if o != nil && o.HouseNumber != nil {
		return true
	}

	return false
}

// SetHouseNumber gets a reference to the given string and assigns it to the HouseNumber field.
func (o *SupplierWrite) SetHouseNumber(v string) {
	o.HouseNumber = &v
}

// GetName returns the Name field value
func (o *SupplierWrite) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SupplierWrite) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SupplierWrite) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *SupplierWrite) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWrite) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *SupplierWrite) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *SupplierWrite) SetPhone(v string) {
	o.Phone = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *SupplierWrite) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWrite) GetStreetOk() (*string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *SupplierWrite) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *SupplierWrite) SetStreet(v string) {
	o.Street = &v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *SupplierWrite) GetStreet2() string {
	if o == nil || o.Street2 == nil {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWrite) GetStreet2Ok() (*string, bool) {
	if o == nil || o.Street2 == nil {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *SupplierWrite) HasStreet2() bool {
	if o != nil && o.Street2 != nil {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *SupplierWrite) SetStreet2(v string) {
	o.Street2 = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *SupplierWrite) GetZip() string {
	if o == nil || o.Zip == nil {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWrite) GetZipOk() (*string, bool) {
	if o == nil || o.Zip == nil {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *SupplierWrite) HasZip() bool {
	if o != nil && o.Zip != nil {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *SupplierWrite) SetZip(v string) {
	o.Zip = &v
}

func (o SupplierWrite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.HouseNumber != nil {
		toSerialize["houseNumber"] = o.HouseNumber
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Street != nil {
		toSerialize["street"] = o.Street
	}
	if o.Street2 != nil {
		toSerialize["street2"] = o.Street2
	}
	if o.Zip != nil {
		toSerialize["zip"] = o.Zip
	}
	return json.Marshal(toSerialize)
}

type NullableSupplierWrite struct {
	value *SupplierWrite
	isSet bool
}

func (v NullableSupplierWrite) Get() *SupplierWrite {
	return v.value
}

func (v *NullableSupplierWrite) Set(val *SupplierWrite) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplierWrite) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplierWrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplierWrite(val *SupplierWrite) *NullableSupplierWrite {
	return &NullableSupplierWrite{value: val, isSet: true}
}

func (v NullableSupplierWrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplierWrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


