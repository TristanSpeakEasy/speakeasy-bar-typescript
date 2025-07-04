// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"fmt"
	"log"
	"mockserver/internal/handler/assert"
	"mockserver/internal/logging"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/models/operations"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"mockserver/internal/tracking"
	"net/http"
)

func pathGetDrinks(dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")
		instanceID := req.Header.Get("x-speakeasy-test-instance-id")

		count := rt.GetRequestCount(test, instanceID)

		switch fmt.Sprintf("%s[%d]", test, count) {
		case "listDrinks-unauthenticated_drinks[0]":
			dir.HandlerFunc("listDrinks", testListDrinksListDrinksUnauthenticatedDrinks0)(w, req)
		case "listDrinks-authenticated_drinks[0]":
			dir.HandlerFunc("listDrinks", testListDrinksListDrinksAuthenticatedDrinks0)(w, req)
		default:
			http.Error(w, fmt.Sprintf("Unknown test: %s[%d]", test, count), http.StatusBadRequest)
		}
	}
}

func testListDrinksListDrinksUnauthenticatedDrinks0(w http.ResponseWriter, req *http.Request) {
	if err := assert.SecurityAuthorizationHeader(req, true, ""); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := assert.AcceptHeader(req, []string{"application/json"}); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.HeaderExists(req, "User-Agent"); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var respBody []operations.ResponseBody = []operations.ResponseBody{
		operations.CreateResponseBodyUnauthenticated(
			components.PublicDrink{
				Name:      "Old Fashioned",
				Type:      components.DrinkTypeCocktail.ToPointer(),
				Price:     1000,
				Photo:     types.String("https://speakeasy.bar/drinks/old_fashioned.jpg"),
				DataLevel: components.PublicDrinkDataLevelUnauthenticated.ToPointer(),
			},
		),
		operations.CreateResponseBodyUnauthenticated(
			components.PublicDrink{
				Name:      "Manhattan",
				Type:      components.DrinkTypeCocktail.ToPointer(),
				Price:     1200,
				Photo:     types.String("https://speakeasy.bar/drinks/manhattan.jpg"),
				DataLevel: components.PublicDrinkDataLevelUnauthenticated.ToPointer(),
			},
		),
		operations.CreateResponseBodyUnauthenticated(
			components.PublicDrink{
				Name:      "Negroni",
				Type:      components.DrinkTypeCocktail.ToPointer(),
				Price:     1500,
				Photo:     types.String("https://speakeasy.bar/drinks/negroni.jpg"),
				DataLevel: components.PublicDrinkDataLevelUnauthenticated.ToPointer(),
			},
		),
	}
	respBodyBytes, err := utils.MarshalJSON(respBody, "", true)

	if err != nil {
		http.Error(
			w,
			"Unable to encode response body as JSON: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBodyBytes)
}

func testListDrinksListDrinksAuthenticatedDrinks0(w http.ResponseWriter, req *http.Request) {
	if err := assert.SecurityAuthorizationHeader(req, true, ""); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := assert.AcceptHeader(req, []string{"application/json"}); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.HeaderExists(req, "User-Agent"); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var respBody []operations.ResponseBody = []operations.ResponseBody{
		operations.CreateResponseBodyAuthenticated(
			components.Drink{
				Name:        "Old Fashioned",
				Type:        components.DrinkTypeCocktail.ToPointer(),
				Price:       1000,
				Photo:       types.String("https://speakeasy.bar/drinks/old_fashioned.jpg"),
				DataLevel:   components.DrinkDataLevelAuthenticated.ToPointer(),
				Stock:       types.Int64(100),
				ProductCode: "AC-A2DF3",
			},
		),
		operations.CreateResponseBodyAuthenticated(
			components.Drink{
				Name:        "Manhattan",
				Type:        components.DrinkTypeCocktail.ToPointer(),
				Price:       1200,
				Photo:       types.String("https://speakeasy.bar/drinks/manhattan.jpg"),
				DataLevel:   components.DrinkDataLevelAuthenticated.ToPointer(),
				Stock:       types.Int64(50),
				ProductCode: "AC-A2DF4",
			},
		),
		operations.CreateResponseBodyAuthenticated(
			components.Drink{
				Name:        "Negroni",
				Type:        components.DrinkTypeCocktail.ToPointer(),
				Price:       1500,
				Photo:       types.String("https://speakeasy.bar/drinks/negroni.jpg"),
				DataLevel:   components.DrinkDataLevelAuthenticated.ToPointer(),
				Stock:       types.Int64(0),
				ProductCode: "AC-A2DF5",
			},
		),
	}
	respBodyBytes, err := utils.MarshalJSON(respBody, "", true)

	if err != nil {
		http.Error(
			w,
			"Unable to encode response body as JSON: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBodyBytes)
}
