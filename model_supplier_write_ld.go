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

// SupplierWriteLd Supplier entity
type SupplierWriteLd struct {
	LdContext *string `json:"@context,omitempty"`
	LdId      *string `json:"@id,omitempty"`
	LdType    *string `json:"@type,omitempty"`
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

// NewSupplierWriteLd instantiates a new SupplierWriteLd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplierWriteLd(name string) *SupplierWriteLd {
	this := SupplierWriteLd{}
	this.Name = name
	return &this
}

// NewSupplierWriteLdWithDefaults instantiates a new SupplierWriteLd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplierWriteLdWithDefaults() *SupplierWriteLd {
	this := SupplierWriteLd{}
	return &this
}

// GetLdContext returns the LdContext field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetLdContext() string {
	if o == nil || o.LdContext == nil {
		var ret string
		return ret
	}
	return *o.LdContext
}

// GetLdContextOk returns a tuple with the LdContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetLdContextOk() (*string, bool) {
	if o == nil || o.LdContext == nil {
		return nil, false
	}
	return o.LdContext, true
}

// HasLdContext returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasLdContext() bool {
	if o != nil && o.LdContext != nil {
		return true
	}

	return false
}

// SetLdContext gets a reference to the given string and assigns it to the LdContext field.
func (o *SupplierWriteLd) SetLdContext(v string) {
	o.LdContext = &v
}

// GetLdId returns the LdId field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetLdId() string {
	if o == nil || o.LdId == nil {
		var ret string
		return ret
	}
	return *o.LdId
}

// GetLdIdOk returns a tuple with the LdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetLdIdOk() (*string, bool) {
	if o == nil || o.LdId == nil {
		return nil, false
	}
	return o.LdId, true
}

// HasLdId returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasLdId() bool {
	if o != nil && o.LdId != nil {
		return true
	}

	return false
}

// SetLdId gets a reference to the given string and assigns it to the LdId field.
func (o *SupplierWriteLd) SetLdId(v string) {
	o.LdId = &v
}

// GetLdType returns the LdType field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetLdType() string {
	if o == nil || o.LdType == nil {
		var ret string
		return ret
	}
	return *o.LdType
}

// GetLdTypeOk returns a tuple with the LdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetLdTypeOk() (*string, bool) {
	if o == nil || o.LdType == nil {
		return nil, false
	}
	return o.LdType, true
}

// HasLdType returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasLdType() bool {
	if o != nil && o.LdType != nil {
		return true
	}

	return false
}

// SetLdType gets a reference to the given string and assigns it to the LdType field.
func (o *SupplierWriteLd) SetLdType(v string) {
	o.LdType = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *SupplierWriteLd) SetCity(v string) {
	o.City = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *SupplierWriteLd) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SupplierWriteLd) SetEmail(v string) {
	o.Email = &v
}

// GetHouseNumber returns the HouseNumber field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetHouseNumber() string {
	if o == nil || o.HouseNumber == nil {
		var ret string
		return ret
	}
	return *o.HouseNumber
}

// GetHouseNumberOk returns a tuple with the HouseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetHouseNumberOk() (*string, bool) {
	if o == nil || o.HouseNumber == nil {
		return nil, false
	}
	return o.HouseNumber, true
}

// HasHouseNumber returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasHouseNumber() bool {
	if o != nil && o.HouseNumber != nil {
		return true
	}

	return false
}

// SetHouseNumber gets a reference to the given string and assigns it to the HouseNumber field.
func (o *SupplierWriteLd) SetHouseNumber(v string) {
	o.HouseNumber = &v
}

// GetName returns the Name field value
func (o *SupplierWriteLd) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SupplierWriteLd) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *SupplierWriteLd) SetPhone(v string) {
	o.Phone = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetStreetOk() (*string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *SupplierWriteLd) SetStreet(v string) {
	o.Street = &v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetStreet2() string {
	if o == nil || o.Street2 == nil {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetStreet2Ok() (*string, bool) {
	if o == nil || o.Street2 == nil {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasStreet2() bool {
	if o != nil && o.Street2 != nil {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *SupplierWriteLd) SetStreet2(v string) {
	o.Street2 = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *SupplierWriteLd) GetZip() string {
	if o == nil || o.Zip == nil {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplierWriteLd) GetZipOk() (*string, bool) {
	if o == nil || o.Zip == nil {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *SupplierWriteLd) HasZip() bool {
	if o != nil && o.Zip != nil {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *SupplierWriteLd) SetZip(v string) {
	o.Zip = &v
}

func (o SupplierWriteLd) MarshalJSON() ([]byte, error) {
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

type NullableSupplierWriteLd struct {
	value *SupplierWriteLd
	isSet bool
}

func (v NullableSupplierWriteLd) Get() *SupplierWriteLd {
	return v.value
}

func (v *NullableSupplierWriteLd) Set(val *SupplierWriteLd) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplierWriteLd) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplierWriteLd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplierWriteLd(val *SupplierWriteLd) *NullableSupplierWriteLd {
	return &NullableSupplierWriteLd{value: val, isSet: true}
}

func (v NullableSupplierWriteLd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplierWriteLd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}