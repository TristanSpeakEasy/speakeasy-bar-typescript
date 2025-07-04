// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type UpdateDrinkRequest struct {
	ProductCode string                `pathParam:"style=simple,explode=false,name=productCode"`
	Drink       components.DrinkInput `request:"mediaType=application/json"`
}

func (o *UpdateDrinkRequest) GetProductCode() string {
	if o == nil {
		return ""
	}
	return o.ProductCode
}

func (o *UpdateDrinkRequest) GetDrink() components.DrinkInput {
	if o == nil {
		return components.DrinkInput{}
	}
	return o.Drink
}

type UpdateDrinkResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The drink was updated successfully.
	Drink *components.Drink
	Body  []byte
	// An unknown error occurred interacting with the API.
	Error *components.Error
}

func (o *UpdateDrinkResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateDrinkResponse) GetDrink() *components.Drink {
	if o == nil {
		return nil
	}
	return o.Drink
}

func (o *UpdateDrinkResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *UpdateDrinkResponse) GetError() *components.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
