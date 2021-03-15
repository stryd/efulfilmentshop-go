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
	"time"
)

// ShipmentRead Shipment entity
type ShipmentRead struct {
	// The shipment carrier ID
	CarrierId *int32 `json:"carrierId,omitempty"`
	// The shipment ID
	Id *int32 `json:"id,omitempty"`
	// The sale ID
	SaleId *int32 `json:"saleId,omitempty"`
	// The shipment date
	ShippedAt *time.Time `json:"shippedAt,omitempty"`
	// The shipment tracking codes
	TrackingCodes *[]string `json:"trackingCodes,omitempty"`
}

// NewShipmentRead instantiates a new ShipmentRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentRead() *ShipmentRead {
	this := ShipmentRead{}
	return &this
}

// NewShipmentReadWithDefaults instantiates a new ShipmentRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentReadWithDefaults() *ShipmentRead {
	this := ShipmentRead{}
	return &this
}

// GetCarrierId returns the CarrierId field value if set, zero value otherwise.
func (o *ShipmentRead) GetCarrierId() int32 {
	if o == nil || o.CarrierId == nil {
		var ret int32
		return ret
	}
	return *o.CarrierId
}

// GetCarrierIdOk returns a tuple with the CarrierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentRead) GetCarrierIdOk() (*int32, bool) {
	if o == nil || o.CarrierId == nil {
		return nil, false
	}
	return o.CarrierId, true
}

// HasCarrierId returns a boolean if a field has been set.
func (o *ShipmentRead) HasCarrierId() bool {
	if o != nil && o.CarrierId != nil {
		return true
	}

	return false
}

// SetCarrierId gets a reference to the given int32 and assigns it to the CarrierId field.
func (o *ShipmentRead) SetCarrierId(v int32) {
	o.CarrierId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShipmentRead) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentRead) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShipmentRead) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ShipmentRead) SetId(v int32) {
	o.Id = &v
}

// GetSaleId returns the SaleId field value if set, zero value otherwise.
func (o *ShipmentRead) GetSaleId() int32 {
	if o == nil || o.SaleId == nil {
		var ret int32
		return ret
	}
	return *o.SaleId
}

// GetSaleIdOk returns a tuple with the SaleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentRead) GetSaleIdOk() (*int32, bool) {
	if o == nil || o.SaleId == nil {
		return nil, false
	}
	return o.SaleId, true
}

// HasSaleId returns a boolean if a field has been set.
func (o *ShipmentRead) HasSaleId() bool {
	if o != nil && o.SaleId != nil {
		return true
	}

	return false
}

// SetSaleId gets a reference to the given int32 and assigns it to the SaleId field.
func (o *ShipmentRead) SetSaleId(v int32) {
	o.SaleId = &v
}

// GetShippedAt returns the ShippedAt field value if set, zero value otherwise.
func (o *ShipmentRead) GetShippedAt() time.Time {
	if o == nil || o.ShippedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ShippedAt
}

// GetShippedAtOk returns a tuple with the ShippedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentRead) GetShippedAtOk() (*time.Time, bool) {
	if o == nil || o.ShippedAt == nil {
		return nil, false
	}
	return o.ShippedAt, true
}

// HasShippedAt returns a boolean if a field has been set.
func (o *ShipmentRead) HasShippedAt() bool {
	if o != nil && o.ShippedAt != nil {
		return true
	}

	return false
}

// SetShippedAt gets a reference to the given time.Time and assigns it to the ShippedAt field.
func (o *ShipmentRead) SetShippedAt(v time.Time) {
	o.ShippedAt = &v
}

// GetTrackingCodes returns the TrackingCodes field value if set, zero value otherwise.
func (o *ShipmentRead) GetTrackingCodes() []string {
	if o == nil || o.TrackingCodes == nil {
		var ret []string
		return ret
	}
	return *o.TrackingCodes
}

// GetTrackingCodesOk returns a tuple with the TrackingCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentRead) GetTrackingCodesOk() (*[]string, bool) {
	if o == nil || o.TrackingCodes == nil {
		return nil, false
	}
	return o.TrackingCodes, true
}

// HasTrackingCodes returns a boolean if a field has been set.
func (o *ShipmentRead) HasTrackingCodes() bool {
	if o != nil && o.TrackingCodes != nil {
		return true
	}

	return false
}

// SetTrackingCodes gets a reference to the given []string and assigns it to the TrackingCodes field.
func (o *ShipmentRead) SetTrackingCodes(v []string) {
	o.TrackingCodes = &v
}

func (o ShipmentRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CarrierId != nil {
		toSerialize["carrierId"] = o.CarrierId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SaleId != nil {
		toSerialize["saleId"] = o.SaleId
	}
	if o.ShippedAt != nil {
		toSerialize["shippedAt"] = o.ShippedAt
	}
	if o.TrackingCodes != nil {
		toSerialize["trackingCodes"] = o.TrackingCodes
	}
	return json.Marshal(toSerialize)
}

type NullableShipmentRead struct {
	value *ShipmentRead
	isSet bool
}

func (v NullableShipmentRead) Get() *ShipmentRead {
	return v.value
}

func (v *NullableShipmentRead) Set(val *ShipmentRead) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentRead) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentRead(val *ShipmentRead) *NullableShipmentRead {
	return &NullableShipmentRead{value: val, isSet: true}
}

func (v NullableShipmentRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


