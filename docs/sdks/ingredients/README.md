# Ingredients
(*ingredients*)

## Overview

The ingredients endpoints.

### Available Operations

* [get](#get) - Get a list of ingredients.
* [stockUpdate](#stockupdate) - Receive stock updates.

## get

Get a list of ingredients, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const result = await speakeasyBar.ingredients.get();

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SpeakeasyBarCore } from "speakeasy-bar/core.js";
import { ingredientsGet } from "speakeasy-bar/funcs/ingredientsGet.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const res = await ingredientsGet(speakeasyBar);
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("ingredientsGet failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.ListIngredientsRequest](../../models/operations/listingredientsrequest.md)                                                                                         | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[operations.ListIngredientsResponse](../../models/operations/listingredientsresponse.md)\>**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| errors.APIError                 | 5XX                             | application/json                |
| errors.SpeakeasyBarDefaultError | 4XX                             | \*/\*                           |

## stockUpdate

Receive stock updates from the bar, this will be called whenever the stock levels of a drink or ingredient changes.

### Example Usage

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar();

async function run() {
  const result = await speakeasyBar.ingredients.stockUpdate({
    drink: {
      name: "Old Fashioned",
      price: 1000,
      photo: "https://speakeasy.bar/drinks/old_fashioned.jpg",
      productCode: "AC-A2DF3",
      ingredients: [
        {
          ingredientProductCode: "AC-A2DF3",
          quantity: 111009,
        },
      ],
    },
    ingredient: {
      name: "Sugar Syrup",
      type: "packaged",
      productCode: "AC-A2DF3",
      photo: "https://speakeasy.bar/ingredients/sugar_syrup.jpg",
    },
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SpeakeasyBarCore } from "speakeasy-bar/core.js";
import { ingredientsStockUpdate } from "speakeasy-bar/funcs/ingredientsStockUpdate.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore();

async function run() {
  const res = await ingredientsStockUpdate(speakeasyBar, {
    drink: {
      name: "Old Fashioned",
      price: 1000,
      photo: "https://speakeasy.bar/drinks/old_fashioned.jpg",
      productCode: "AC-A2DF3",
      ingredients: [
        {
          ingredientProductCode: "AC-A2DF3",
          quantity: 111009,
        },
      ],
    },
    ingredient: {
      name: "Sugar Syrup",
      type: "packaged",
      productCode: "AC-A2DF3",
      photo: "https://speakeasy.bar/ingredients/sugar_syrup.jpg",
    },
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("ingredientsStockUpdate failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [webhooks.StockUpdateRequest](../../models/webhooks/stockupdaterequest.md)                                                                                                     | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[models.ErrorT](../../models/errort.md)\>**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| errors.APIError                 | 5XX                             | application/json                |
| errors.SpeakeasyBarDefaultError | 4XX                             | \*/\*                           |