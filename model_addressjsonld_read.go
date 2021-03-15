/*
 * E-fulfilment Shop API
 *
 * The E-fulfilment Shop API allows you to integrate your service with our warehouse.
 *
 * API version: 1.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package efulfillment_go

import (
	"encoding/json"
)

// AddressjsonldRead Address entity
type AddressjsonldRead struct {
	JSONContext *string `json:"@context,omitempty"`
	JSONId *string `json:"@id,omitempty"`
	JSONType *string `json:"@type,omitempty"`
	// The address city
	City string `json:"city"`
	// The address company
	Company *string `json:"company,omitempty"`
	// The address country code (ISO 3166-1 alpha-2 format)
	CountryCode string `json:"countryCode"`
	// The address email
	Email *string `json:"email,omitempty"`
	// The address house number
	HouseNumber *string `json:"houseNumber,omitempty"`
	// The address ID
	Id *int32 `json:"id,omitempty"`
	// The address name
	Name string `json:"name"`
	// The address phone
	Phone *string `json:"phone,omitempty"`
	// The address street
	Street string `json:"street"`
	// The address street 2
	Street2 *string `json:"street2,omitempty"`
	// The address zip
	Zip string `json:"zip"`
}

// NewAddressjsonldRead instantiates a new AddressjsonldRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressjsonldRead(city string, countryCode string, name string, street string, zip string, ) *AddressjsonldRead {
	this := AddressjsonldRead{}
	this.City = city
	this.CountryCode = countryCode
	this.Name = name
	this.Street = street
	this.Zip = zip
	return &this
}

// NewAddressjsonldReadWithDefaults instantiates a new AddressjsonldRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressjsonldReadWithDefaults() *AddressjsonldRead {
	this := AddressjsonldRead{}
	return &this
}

// GetJSONContext returns the JSONContext field value if set, zero value otherwise.
func (o *AddressjsonldRead) GetJSONContext() string {
	if o == nil || o.JSONContext == nil {
		var ret string
		return ret
	}
	return *o.JSONContext
}

// GetJSONContextOk returns a tuple with the JSONContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetJSONContextOk() (*string, bool) {
	if o == nil || o.JSONContext == nil {
		return nil, false
	}
	return o.JSONContext, true
}

// HasJSONContext returns a boolean if a field has been set.
func (o *AddressjsonldRead) HasJSONContext() bool {
	if o != nil && o.JSONContext != nil {
		return true
	}

	return false
}

// SetJSONContext gets a reference to the given string and assigns it to the JSONContext field.
func (o *AddressjsonldRead) SetJSONContext(v string) {
	o.JSONContext = &v
}

// GetJSONId returns the JSONId field value if set, zero value otherwise.
func (o *AddressjsonldRead) GetJSONId() string {
	if o == nil || o.JSONId == nil {
		var ret string
		return ret
	}
	return *o.JSONId
}

// GetJSONIdOk returns a tuple with the JSONId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetJSONIdOk() (*string, bool) {
	if o == nil || o.JSONId == nil {
		return nil, false
	}
	return o.JSONId, true
}

// HasJSONId returns a boolean if a field has been set.
func (o *AddressjsonldRead) HasJSONId() bool {
	if o != nil && o.JSONId != nil {
		return true
	}

	return false
}

// SetJSONId gets a reference to the given string and assigns it to the JSONId field.
func (o *AddressjsonldRead) SetJSONId(v string) {
	o.JSONId = &v
}

// GetJSONType returns the JSONType field value if set, zero value otherwise.
func (o *AddressjsonldRead) GetJSONType() string {
	if o == nil || o.JSONType == nil {
		var ret string
		return ret
	}
	return *o.JSONType
}

// GetJSONTypeOk returns a tuple with the JSONType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetJSONTypeOk() (*string, bool) {
	if o == nil || o.JSONType == nil {
		return nil, false
	}
	return o.JSONType, true
}

// HasJSONType returns a boolean if a field has been set.
func (o *AddressjsonldRead) HasJSONType() bool {
	if o != nil && o.JSONType != nil {
		return true
	}

	return false
}

// SetJSONType gets a reference to the given string and assigns it to the JSONType field.
func (o *AddressjsonldRead) SetJSONType(v string) {
	o.JSONType = &v
}

// GetCity returns the City field value
func (o *AddressjsonldRead) GetCity() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *AddressjsonldRead) SetCity(v string) {
	o.City = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *AddressjsonldRead) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *AddressjsonldRead) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *AddressjsonldRead) SetCompany(v string) {
	o.Company = &v
}

// GetCountryCode returns the CountryCode field value
func (o *AddressjsonldRead) GetCountryCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *AddressjsonldRead) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AddressjsonldRead) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AddressjsonldRead) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AddressjsonldRead) SetEmail(v string) {
	o.Email = &v
}

// GetHouseNumber returns the HouseNumber field value if set, zero value otherwise.
func (o *AddressjsonldRead) GetHouseNumber() string {
	if o == nil || o.HouseNumber == nil {
		var ret string
		return ret
	}
	return *o.HouseNumber
}

// GetHouseNumberOk returns a tuple with the HouseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetHouseNumberOk() (*string, bool) {
	if o == nil || o.HouseNumber == nil {
		return nil, false
	}
	return o.HouseNumber, true
}

// HasHouseNumber returns a boolean if a field has been set.
func (o *AddressjsonldRead) HasHouseNumber() bool {
	if o != nil && o.HouseNumber != nil {
		return true
	}

	return false
}

// SetHouseNumber gets a reference to the given string and assigns it to the HouseNumber field.
func (o *AddressjsonldRead) SetHouseNumber(v string) {
	o.HouseNumber = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddressjsonldRead) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddressjsonldRead) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AddressjsonldRead) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *AddressjsonldRead) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddressjsonldRead) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *AddressjsonldRead) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *AddressjsonldRead) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *AddressjsonldRead) SetPhone(v string) {
	o.Phone = &v
}

// GetStreet returns the Street field value
func (o *AddressjsonldRead) GetStreet() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *AddressjsonldRead) SetStreet(v string) {
	o.Street = v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *AddressjsonldRead) GetStreet2() string {
	if o == nil || o.Street2 == nil {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetStreet2Ok() (*string, bool) {
	if o == nil || o.Street2 == nil {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *AddressjsonldRead) HasStreet2() bool {
	if o != nil && o.Street2 != nil {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *AddressjsonldRead) SetStreet2(v string) {
	o.Street2 = &v
}

// GetZip returns the Zip field value
func (o *AddressjsonldRead) GetZip() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Zip
}

// GetZipOk returns a tuple with the Zip field value
// and a boolean to check if the value has been set.
func (o *AddressjsonldRead) GetZipOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Zip, true
}

// SetZip sets field value
func (o *AddressjsonldRead) SetZip(v string) {
	o.Zip = v
}

func (o AddressjsonldRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JSONContext != nil {
		toSerialize["@context"] = o.JSONContext
	}
	if o.JSONId != nil {
		toSerialize["@id"] = o.JSONId
	}
	if o.JSONType != nil {
		toSerialize["@type"] = o.JSONType
	}
	if true {
		toSerialize["city"] = o.City
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if true {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.HouseNumber != nil {
		toSerialize["houseNumber"] = o.HouseNumber
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if true {
		toSerialize["street"] = o.Street
	}
	if o.Street2 != nil {
		toSerialize["street2"] = o.Street2
	}
	if true {
		toSerialize["zip"] = o.Zip
	}
	return json.Marshal(toSerialize)
}

type NullableAddressjsonldRead struct {
	value *AddressjsonldRead
	isSet bool
}

func (v NullableAddressjsonldRead) Get() *AddressjsonldRead {
	return v.value
}

func (v *NullableAddressjsonldRead) Set(val *AddressjsonldRead) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressjsonldRead) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressjsonldRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressjsonldRead(val *AddressjsonldRead) *NullableAddressjsonldRead {
	return &NullableAddressjsonldRead{value: val, isSet: true}
}

func (v NullableAddressjsonldRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressjsonldRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


