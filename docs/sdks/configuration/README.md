# Configuration
(*configuration*)

## Overview

The configuration endpoints.

### Available Operations

* [subscribeToWebhooks](#subscribetowebhooks) - Subscribe to webhooks.

## subscribeToWebhooks

Subscribe to webhooks.

### Example Usage

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const result = await speakeasyBar.configuration.subscribeToWebhooks([]);

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { SpeakeasyBarCore } from "speakeasy-bar/core.js";
import { configurationSubscribeToWebhooks } from "speakeasy-bar/funcs/configurationSubscribeToWebhooks.js";

// Use `SpeakeasyBarCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const speakeasyBar = new SpeakeasyBarCore({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const res = await configurationSubscribeToWebhooks(speakeasyBar, []);
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("configurationSubscribeToWebhooks failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.RequestBody[]](../../models/.md)                                                                                                                                   | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
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