// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"io"
	"mockserver/internal/sdk/models/components"
)

type Photo struct {
	FileName string    `multipartForm:"name=fileName"`
	Content  io.Reader `multipartForm:"content"`
}

func (o *Photo) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *Photo) GetContent() io.Reader {
	if o == nil {
		return nil
	}
	return o.Content
}

type UpdateDrinkMultipartRequestBody struct {
	// The name of the drink.
	Name *string `multipartForm:"name=name"`
	// The type of drink.
	Type *components.DrinkType `multipartForm:"name=type"`
	// The price of one unit of the drink in US cents.
	Price *float64 `multipartForm:"name=price"`
	// A photo of the drink.
	Photo *Photo `multipartForm:"file"`
}

func (o *UpdateDrinkMultipartRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateDrinkMultipartRequestBody) GetType() *components.DrinkType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UpdateDrinkMultipartRequestBody) GetPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *UpdateDrinkMultipartRequestBody) GetPhoto() *Photo {
	if o == nil {
		return nil
	}
	return o.Photo
}

type UpdateDrinkMultipartRequest struct {
	ProductCode string                          `pathParam:"style=simple,explode=false,name=productCode"`
	RequestBody UpdateDrinkMultipartRequestBody `request:"mediaType=multipart/form-data"`
}

func (o *UpdateDrinkMultipartRequest) GetProductCode() string {
	if o == nil {
		return ""
	}
	return o.ProductCode
}

func (o *UpdateDrinkMultipartRequest) GetRequestBody() UpdateDrinkMultipartRequestBody {
	if o == nil {
		return UpdateDrinkMultipartRequestBody{}
	}
	return o.RequestBody
}

type UpdateDrinkMultipartResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The drink was updated successfully.
	Drink *components.Drink
	Body  []byte
	// An unknown error occurred interacting with the API.
	Error *components.Error
}

func (o *UpdateDrinkMultipartResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateDrinkMultipartResponse) GetDrink() *components.Drink {
	if o == nil {
		return nil
	}
	return o.Drink
}

func (o *UpdateDrinkMultipartResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *UpdateDrinkMultipartResponse) GetError() *components.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
