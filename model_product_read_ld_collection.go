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

// ProductReadLdCollection Product entity
type ProductReadLdCollection struct {
	// The product barcode
	Barcode string `json:"barcode"`
	// The product channel reference (DEPRECATED)
	ChannelReference *string `json:"channelReference,omitempty"`
	// The product cost price (Excluding taxes)
	CostPrice *float32 `json:"costPrice,omitempty"`
	// The product dimension X in cm
	DimensionX *float32 `json:"dimensionX,omitempty"`
	// The product dimension Y in cm
	DimensionY *float32 `json:"dimensionY,omitempty"`
	// The product dimension Z in cm
	DimensionZ *float32 `json:"dimensionZ,omitempty"`
	// The product HS code
	HsCode *string `json:"hsCode,omitempty"`
	// The product ID
	Id *int32 `json:"id,omitempty"`
	// The product name
	Name string `json:"name"`
	// The physical available product quantity
	QuantityAvailable *int32 `json:"quantityAvailable,omitempty"`
	// The incoming product quantity
	QuantityIncoming *int32 `json:"quantityIncoming,omitempty"`
	// The product quantity on hand
	QuantityOnHand *int32 `json:"quantityOnHand,omitempty"`
	// The outgoing product quantity
	QuantityOutgoing *int32 `json:"quantityOutgoing,omitempty"`
	// The sold product quantity
	QuantitySold *int32 `json:"quantitySold,omitempty"`
	// Your product reference (This could be your product ID)
	Reference *string `json:"reference,omitempty"`
	// The product selling price (Excluding taxes)
	SellingPrice *float32 `json:"sellingPrice,omitempty"`
	// The product warehouse SKU
	Sku *string `json:"sku,omitempty"`
	// The product volume in L (Calculated using dimensions)
	Volume *float32 `json:"volume,omitempty"`
	// The product weight in kg
	Weight *float32 `json:"weight,omitempty"`
}

// NewProductReadLdCollection instantiates a new ProductReadLdCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductReadLdCollection(barcode string, name string) *ProductReadLdCollection {
	this := ProductReadLdCollection{}
	this.Barcode = barcode
	this.Name = name
	return &this
}

// NewProductReadLdCollectionWithDefaults instantiates a new ProductReadLdCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductReadLdCollectionWithDefaults() *ProductReadLdCollection {
	this := ProductReadLdCollection{}
	return &this
}

// GetBarcode returns the Barcode field value
func (o *ProductReadLdCollection) GetBarcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetBarcodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Barcode, true
}

// SetBarcode sets field value
func (o *ProductReadLdCollection) SetBarcode(v string) {
	o.Barcode = v
}

// GetChannelReference returns the ChannelReference field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetChannelReference() string {
	if o == nil || o.ChannelReference == nil {
		var ret string
		return ret
	}
	return *o.ChannelReference
}

// GetChannelReferenceOk returns a tuple with the ChannelReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetChannelReferenceOk() (*string, bool) {
	if o == nil || o.ChannelReference == nil {
		return nil, false
	}
	return o.ChannelReference, true
}

// HasChannelReference returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasChannelReference() bool {
	if o != nil && o.ChannelReference != nil {
		return true
	}

	return false
}

// SetChannelReference gets a reference to the given string and assigns it to the ChannelReference field.
func (o *ProductReadLdCollection) SetChannelReference(v string) {
	o.ChannelReference = &v
}

// GetCostPrice returns the CostPrice field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetCostPrice() float32 {
	if o == nil || o.CostPrice == nil {
		var ret float32
		return ret
	}
	return *o.CostPrice
}

// GetCostPriceOk returns a tuple with the CostPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetCostPriceOk() (*float32, bool) {
	if o == nil || o.CostPrice == nil {
		return nil, false
	}
	return o.CostPrice, true
}

// HasCostPrice returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasCostPrice() bool {
	if o != nil && o.CostPrice != nil {
		return true
	}

	return false
}

// SetCostPrice gets a reference to the given float32 and assigns it to the CostPrice field.
func (o *ProductReadLdCollection) SetCostPrice(v float32) {
	o.CostPrice = &v
}

// GetDimensionX returns the DimensionX field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetDimensionX() float32 {
	if o == nil || o.DimensionX == nil {
		var ret float32
		return ret
	}
	return *o.DimensionX
}

// GetDimensionXOk returns a tuple with the DimensionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetDimensionXOk() (*float32, bool) {
	if o == nil || o.DimensionX == nil {
		return nil, false
	}
	return o.DimensionX, true
}

// HasDimensionX returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasDimensionX() bool {
	if o != nil && o.DimensionX != nil {
		return true
	}

	return false
}

// SetDimensionX gets a reference to the given float32 and assigns it to the DimensionX field.
func (o *ProductReadLdCollection) SetDimensionX(v float32) {
	o.DimensionX = &v
}

// GetDimensionY returns the DimensionY field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetDimensionY() float32 {
	if o == nil || o.DimensionY == nil {
		var ret float32
		return ret
	}
	return *o.DimensionY
}

// GetDimensionYOk returns a tuple with the DimensionY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetDimensionYOk() (*float32, bool) {
	if o == nil || o.DimensionY == nil {
		return nil, false
	}
	return o.DimensionY, true
}

// HasDimensionY returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasDimensionY() bool {
	if o != nil && o.DimensionY != nil {
		return true
	}

	return false
}

// SetDimensionY gets a reference to the given float32 and assigns it to the DimensionY field.
func (o *ProductReadLdCollection) SetDimensionY(v float32) {
	o.DimensionY = &v
}

// GetDimensionZ returns the DimensionZ field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetDimensionZ() float32 {
	if o == nil || o.DimensionZ == nil {
		var ret float32
		return ret
	}
	return *o.DimensionZ
}

// GetDimensionZOk returns a tuple with the DimensionZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetDimensionZOk() (*float32, bool) {
	if o == nil || o.DimensionZ == nil {
		return nil, false
	}
	return o.DimensionZ, true
}

// HasDimensionZ returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasDimensionZ() bool {
	if o != nil && o.DimensionZ != nil {
		return true
	}

	return false
}

// SetDimensionZ gets a reference to the given float32 and assigns it to the DimensionZ field.
func (o *ProductReadLdCollection) SetDimensionZ(v float32) {
	o.DimensionZ = &v
}

// GetHsCode returns the HsCode field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetHsCode() string {
	if o == nil || o.HsCode == nil {
		var ret string
		return ret
	}
	return *o.HsCode
}

// GetHsCodeOk returns a tuple with the HsCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetHsCodeOk() (*string, bool) {
	if o == nil || o.HsCode == nil {
		return nil, false
	}
	return o.HsCode, true
}

// HasHsCode returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasHsCode() bool {
	if o != nil && o.HsCode != nil {
		return true
	}

	return false
}

// SetHsCode gets a reference to the given string and assigns it to the HsCode field.
func (o *ProductReadLdCollection) SetHsCode(v string) {
	o.HsCode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProductReadLdCollection) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ProductReadLdCollection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductReadLdCollection) SetName(v string) {
	o.Name = v
}

// GetQuantityAvailable returns the QuantityAvailable field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetQuantityAvailable() int32 {
	if o == nil || o.QuantityAvailable == nil {
		var ret int32
		return ret
	}
	return *o.QuantityAvailable
}

// GetQuantityAvailableOk returns a tuple with the QuantityAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetQuantityAvailableOk() (*int32, bool) {
	if o == nil || o.QuantityAvailable == nil {
		return nil, false
	}
	return o.QuantityAvailable, true
}

// HasQuantityAvailable returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasQuantityAvailable() bool {
	if o != nil && o.QuantityAvailable != nil {
		return true
	}

	return false
}

// SetQuantityAvailable gets a reference to the given int32 and assigns it to the QuantityAvailable field.
func (o *ProductReadLdCollection) SetQuantityAvailable(v int32) {
	o.QuantityAvailable = &v
}

// GetQuantityIncoming returns the QuantityIncoming field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetQuantityIncoming() int32 {
	if o == nil || o.QuantityIncoming == nil {
		var ret int32
		return ret
	}
	return *o.QuantityIncoming
}

// GetQuantityIncomingOk returns a tuple with the QuantityIncoming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetQuantityIncomingOk() (*int32, bool) {
	if o == nil || o.QuantityIncoming == nil {
		return nil, false
	}
	return o.QuantityIncoming, true
}

// HasQuantityIncoming returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasQuantityIncoming() bool {
	if o != nil && o.QuantityIncoming != nil {
		return true
	}

	return false
}

// SetQuantityIncoming gets a reference to the given int32 and assigns it to the QuantityIncoming field.
func (o *ProductReadLdCollection) SetQuantityIncoming(v int32) {
	o.QuantityIncoming = &v
}

// GetQuantityOnHand returns the QuantityOnHand field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetQuantityOnHand() int32 {
	if o == nil || o.QuantityOnHand == nil {
		var ret int32
		return ret
	}
	return *o.QuantityOnHand
}

// GetQuantityOnHandOk returns a tuple with the QuantityOnHand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetQuantityOnHandOk() (*int32, bool) {
	if o == nil || o.QuantityOnHand == nil {
		return nil, false
	}
	return o.QuantityOnHand, true
}

// HasQuantityOnHand returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasQuantityOnHand() bool {
	if o != nil && o.QuantityOnHand != nil {
		return true
	}

	return false
}

// SetQuantityOnHand gets a reference to the given int32 and assigns it to the QuantityOnHand field.
func (o *ProductReadLdCollection) SetQuantityOnHand(v int32) {
	o.QuantityOnHand = &v
}

// GetQuantityOutgoing returns the QuantityOutgoing field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetQuantityOutgoing() int32 {
	if o == nil || o.QuantityOutgoing == nil {
		var ret int32
		return ret
	}
	return *o.QuantityOutgoing
}

// GetQuantityOutgoingOk returns a tuple with the QuantityOutgoing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetQuantityOutgoingOk() (*int32, bool) {
	if o == nil || o.QuantityOutgoing == nil {
		return nil, false
	}
	return o.QuantityOutgoing, true
}

// HasQuantityOutgoing returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasQuantityOutgoing() bool {
	if o != nil && o.QuantityOutgoing != nil {
		return true
	}

	return false
}

// SetQuantityOutgoing gets a reference to the given int32 and assigns it to the QuantityOutgoing field.
func (o *ProductReadLdCollection) SetQuantityOutgoing(v int32) {
	o.QuantityOutgoing = &v
}

// GetQuantitySold returns the QuantitySold field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetQuantitySold() int32 {
	if o == nil || o.QuantitySold == nil {
		var ret int32
		return ret
	}
	return *o.QuantitySold
}

// GetQuantitySoldOk returns a tuple with the QuantitySold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetQuantitySoldOk() (*int32, bool) {
	if o == nil || o.QuantitySold == nil {
		return nil, false
	}
	return o.QuantitySold, true
}

// HasQuantitySold returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasQuantitySold() bool {
	if o != nil && o.QuantitySold != nil {
		return true
	}

	return false
}

// SetQuantitySold gets a reference to the given int32 and assigns it to the QuantitySold field.
func (o *ProductReadLdCollection) SetQuantitySold(v int32) {
	o.QuantitySold = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ProductReadLdCollection) SetReference(v string) {
	o.Reference = &v
}

// GetSellingPrice returns the SellingPrice field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetSellingPrice() float32 {
	if o == nil || o.SellingPrice == nil {
		var ret float32
		return ret
	}
	return *o.SellingPrice
}

// GetSellingPriceOk returns a tuple with the SellingPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetSellingPriceOk() (*float32, bool) {
	if o == nil || o.SellingPrice == nil {
		return nil, false
	}
	return o.SellingPrice, true
}

// HasSellingPrice returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasSellingPrice() bool {
	if o != nil && o.SellingPrice != nil {
		return true
	}

	return false
}

// SetSellingPrice gets a reference to the given float32 and assigns it to the SellingPrice field.
func (o *ProductReadLdCollection) SetSellingPrice(v float32) {
	o.SellingPrice = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ProductReadLdCollection) SetSku(v string) {
	o.Sku = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetVolume() float32 {
	if o == nil || o.Volume == nil {
		var ret float32
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetVolumeOk() (*float32, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given float32 and assigns it to the Volume field.
func (o *ProductReadLdCollection) SetVolume(v float32) {
	o.Volume = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ProductReadLdCollection) GetWeight() float32 {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadLdCollection) GetWeightOk() (*float32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ProductReadLdCollection) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *ProductReadLdCollection) SetWeight(v float32) {
	o.Weight = &v
}

func (o ProductReadLdCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["barcode"] = o.Barcode
	}
	if o.ChannelReference != nil {
		toSerialize["channelReference"] = o.ChannelReference
	}
	if o.CostPrice != nil {
		toSerialize["costPrice"] = o.CostPrice
	}
	if o.DimensionX != nil {
		toSerialize["dimensionX"] = o.DimensionX
	}
	if o.DimensionY != nil {
		toSerialize["dimensionY"] = o.DimensionY
	}
	if o.DimensionZ != nil {
		toSerialize["dimensionZ"] = o.DimensionZ
	}
	if o.HsCode != nil {
		toSerialize["hsCode"] = o.HsCode
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.QuantityAvailable != nil {
		toSerialize["quantityAvailable"] = o.QuantityAvailable
	}
	if o.QuantityIncoming != nil {
		toSerialize["quantityIncoming"] = o.QuantityIncoming
	}
	if o.QuantityOnHand != nil {
		toSerialize["quantityOnHand"] = o.QuantityOnHand
	}
	if o.QuantityOutgoing != nil {
		toSerialize["quantityOutgoing"] = o.QuantityOutgoing
	}
	if o.QuantitySold != nil {
		toSerialize["quantitySold"] = o.QuantitySold
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.SellingPrice != nil {
		toSerialize["sellingPrice"] = o.SellingPrice
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableProductReadLdCollection struct {
	value *ProductReadLdCollection
	isSet bool
}

func (v NullableProductReadLdCollection) Get() *ProductReadLdCollection {
	return v.value
}

func (v *NullableProductReadLdCollection) Set(val *ProductReadLdCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableProductReadLdCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableProductReadLdCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductReadLdCollection(val *ProductReadLdCollection) *NullableProductReadLdCollection {
	return &NullableProductReadLdCollection{value: val, isSet: true}
}

func (v NullableProductReadLdCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductReadLdCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


