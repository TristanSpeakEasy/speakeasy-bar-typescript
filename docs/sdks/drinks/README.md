# Drinks
(*drinks*)

## Overview

The drinks endpoints.

### Available Operations

* [get](#get) - Get a list of drinks.
* [getById](#getbyid) - Get a drink.
* [update](#update) - Update a drink.
* [updateString](#updatestring) - Update a drink.
* [updateRaw](#updateraw) - Update a drink.
* [updateMultipart](#updatemultipart) - Update a drink.
* [delete](#delete) - Delete a drink.
* [search](#search) - Search for drinks.
* [stockUpdate](#stockupdate) - Receive stock updates.

## get

Get a list of drinks, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar();

async function run() {
  const result = await speakeasyBar.drinks.get();

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SpeakeasyBarCore } from "speakeasy-bar/core.js";
import { drinksGet } from "speakeasy-bar/funcs/drinksGet.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore();

async function run() {
  const res = await drinksGet(speakeasyBar);
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("drinksGet failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.ListDrinksRequest](../../models/operations/listdrinksrequest.md)                                                                                                   | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `security`                                                                                                                                                                     | [operations.ListDrinksSecurity](../../models/operations/listdrinkssecurity.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                             | The security requirements to use for the request.                                                                                                                              |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[operations.ListDrinksResponse](../../models/operations/listdrinksresponse.md)\>**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| errors.APIError                 | 5XX                             | application/json                |
| errors.SpeakeasyBarDefaultError | 4XX                             | \*/\*                           |

## getById

Get a drink by product code. Only available when authenticated.

### Example Usage

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const result = await speakeasyBar.drinks.getById({
    productCode: "AC-A2DF3",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SpeakeasyBarCore } from "speakeasy-bar/core.js";
import { drinksGetById } from "speakeasy-bar/funcs/drinksGetById.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const res = await drinksGetById(speakeasyBar, {
    productCode: "AC-A2DF3",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("drinksGetById failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.GetDrinkRequest](../../models/operations/getdrinkrequest.md)                                                                                                       | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[operations.GetDrinkResponse](../../models/operations/getdrinkresponse.md)\>**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| errors.APIError                 | 5XX                             | application/json                |
| errors.SpeakeasyBarDefaultError | 4XX                             | \*/\*                           |

## update

Update a drink. Only available when authenticated.

### Example Usage

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const result = await speakeasyBar.drinks.update({
    productCode: "AC-A2DF3",
    drink: {
      name: "Old Fashioned",
      type: "cocktail",
      price: 1000,
      photo: "https://speakeasy.bar/drinks/old_fashioned.jpg",
      productCode: "AC-A2DF3",
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
import { drinksUpdate } from "speakeasy-bar/funcs/drinksUpdate.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const res = await drinksUpdate(speakeasyBar, {
    productCode: "AC-A2DF3",
    drink: {
      name: "Old Fashioned",
      type: "cocktail",
      price: 1000,
      photo: "https://speakeasy.bar/drinks/old_fashioned.jpg",
      productCode: "AC-A2DF3",
    },
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("drinksUpdate failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.UpdateDrinkRequest](../../models/operations/updatedrinkrequest.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[operations.UpdateDrinkResponse](../../models/operations/updatedrinkresponse.md)\>**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| errors.APIError                 | 5XX                             | application/json                |
| errors.SpeakeasyBarDefaultError | 4XX                             | \*/\*                           |

## updateString

Update a drink. Only available when authenticated.

### Example Usage

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const result = await speakeasyBar.drinks.updateString({
    productCode: "AC-A2DF3",
    requestBody: "[object Object]",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SpeakeasyBarCore } from "speakeasy-bar/core.js";
import { drinksUpdateString } from "speakeasy-bar/funcs/drinksUpdateString.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const res = await drinksUpdateString(speakeasyBar, {
    productCode: "AC-A2DF3",
    requestBody: "[object Object]",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("drinksUpdateString failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.UpdateDrinkStringRequest](../../models/operations/updatedrinkstringrequest.md)                                                                                     | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[operations.UpdateDrinkStringResponse](../../models/operations/updatedrinkstringresponse.md)\>**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| errors.APIError                 | 5XX                             | application/json                |
| errors.SpeakeasyBarDefaultError | 4XX                             | \*/\*                           |

## updateRaw

Update a drink. Only available when authenticated.

### Example Usage

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const result = await speakeasyBar.drinks.updateRaw({
    productCode: "AC-A2DF3",
    drink: bytesToStream(new TextEncoder().encode("{\"name\":\"Old Fashioned\",\"type\":\"cocktail\",\"price\":1000,\"photo\":\"https://speakeasy.bar/drinks/old_fashioned.jpg\",\"productCode\":\"AC-A2DF3\"}")),
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SpeakeasyBarCore } from "speakeasy-bar/core.js";
import { drinksUpdateRaw } from "speakeasy-bar/funcs/drinksUpdateRaw.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const res = await drinksUpdateRaw(speakeasyBar, {
    productCode: "AC-A2DF3",
    drink: bytesToStream(new TextEncoder().encode("{\"name\":\"Old Fashioned\",\"type\":\"cocktail\",\"price\":1000,\"photo\":\"https://speakeasy.bar/drinks/old_fashioned.jpg\",\"productCode\":\"AC-A2DF3\"}")),
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("drinksUpdateRaw failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.UpdateDrinkRawRequest](../../models/operations/updatedrinkrawrequest.md)                                                                                           | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[operations.UpdateDrinkRawResponse](../../models/operations/updatedrinkrawresponse.md)\>**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| errors.APIError                 | 5XX                             | application/json                |
| errors.SpeakeasyBarDefaultError | 4XX                             | \*/\*                           |

## updateMultipart

Update a drink. Only available when authenticated.

### Example Usage

```typescript
import { openAsBlob } from "node:fs";
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const result = await speakeasyBar.drinks.updateMultipart({
    productCode: "AC-A2DF3",
    requestBody: {
      name: "Old Fashioned",
      type: "cocktail",
      price: 1000,
      photo: await openAsBlob("example.file"),
    },
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { openAsBlob } from "node:fs";
import { SpeakeasyBarCore } from "speakeasy-bar/core.js";
import { drinksUpdateMultipart } from "speakeasy-bar/funcs/drinksUpdateMultipart.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const res = await drinksUpdateMultipart(speakeasyBar, {
    productCode: "AC-A2DF3",
    requestBody: {
      name: "Old Fashioned",
      type: "cocktail",
      price: 1000,
      photo: await openAsBlob("example.file"),
    },
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("drinksUpdateMultipart failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.UpdateDrinkMultipartRequest](../../models/operations/updatedrinkmultipartrequest.md)                                                                               | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[operations.UpdateDrinkMultipartResponse](../../models/operations/updatedrinkmultipartresponse.md)\>**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| errors.APIError                 | 5XX                             | application/json                |
| errors.SpeakeasyBarDefaultError | 4XX                             | \*/\*                           |

## delete

Delete a drink. Only available when authenticated.

### Example Usage

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const result = await speakeasyBar.drinks.delete({
    productCode: "AC-A2DF3",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SpeakeasyBarCore } from "speakeasy-bar/core.js";
import { drinksDelete } from "speakeasy-bar/funcs/drinksDelete.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const res = await drinksDelete(speakeasyBar, {
    productCode: "AC-A2DF3",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("drinksDelete failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.DeleteDrinkRequest](../../models/operations/deletedrinkrequest.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
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

## search

Search for drinks, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar();

async function run() {
  const result = await speakeasyBar.drinks.search({
    query: "Old Fashioned",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SpeakeasyBarCore } from "speakeasy-bar/core.js";
import { drinksSearch } from "speakeasy-bar/funcs/drinksSearch.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore();

async function run() {
  const res = await drinksSearch(speakeasyBar, {
    query: "Old Fashioned",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("drinksSearch failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.SearchDrinksRequest](../../models/operations/searchdrinksrequest.md)                                                                                               | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `security`                                                                                                                                                                     | [operations.SearchDrinksSecurity](../../models/operations/searchdrinkssecurity.md)                                                                                             | :heavy_check_mark:                                                                                                                                                             | The security requirements to use for the request.                                                                                                                              |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[operations.SearchDrinksResponse](../../models/operations/searchdrinksresponse.md)\>**

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
  const result = await speakeasyBar.drinks.stockUpdate({
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
import { drinksStockUpdate } from "speakeasy-bar/funcs/drinksStockUpdate.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore();

async function run() {
  const res = await drinksStockUpdate(speakeasyBar, {
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
    console.log("drinksStockUpdate failed:", res.error);
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