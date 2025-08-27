# speakeasy-bar

Developer-friendly & type-safe Typescript SDK specifically catered to leverage *speakeasy-bar* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=speakeasy-bar&utm_campaign=typescript"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/speakeasy-self/speakeasy-self). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

The Speakeasy Bar: A bar that serves drinks.

A secret underground bar that serves drinks to those in the know.

For more information about the API: [The Speakeasy Bar Documentation.](https://docs.speakeasy.bar)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [speakeasy-bar](#speakeasy-bar)
  * [SDK Installation](#sdk-installation)
  * [Requirements](#requirements)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Standalone functions](#standalone-functions)
  * [File uploads](#file-uploads)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Debugging](#debugging)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

> [!TIP]
> To finish publishing your SDK to npm and others you must [run your first generation action](https://www.speakeasy.com/docs/github-setup#step-by-step-guide).


The SDK can be installed with either [npm](https://www.npmjs.com/), [pnpm](https://pnpm.io/), [bun](https://bun.sh/) or [yarn](https://classic.yarnpkg.com/en/) package managers.

### NPM

```bash
npm add https://github.com/TristanSpeakEasy/speakeasy-bar-typescript
```

### PNPM

```bash
pnpm add https://github.com/TristanSpeakEasy/speakeasy-bar-typescript
```

### Bun

```bash
bun add https://github.com/TristanSpeakEasy/speakeasy-bar-typescript
```

### Yarn

```bash
yarn add https://github.com/TristanSpeakEasy/speakeasy-bar-typescript zod

# Note that Yarn does not install peer dependencies automatically. You will need
# to install zod as shown above.
```

> [!NOTE]
> This package is published with CommonJS and ES Modules (ESM) support.


### Model Context Protocol (MCP) Server

This SDK is also an installable MCP server where the various SDK methods are
exposed as tools that can be invoked by AI applications.

> Node.js v20 or greater is required to run the MCP server from npm.

<details>
<summary>Claude installation steps</summary>

Add the following server definition to your `claude_desktop_config.json` file:

```json
{
  "mcpServers": {
    "SpeakeasyBar": {
      "command": "npx",
      "args": [
        "-y", "--package", "speakeasy-bar",
        "--",
        "mcp", "start",
        "--api-key", "..."
      ]
    }
  }
}
```

</details>

<details>
<summary>Cursor installation steps</summary>

Create a `.cursor/mcp.json` file in your project root with the following content:

```json
{
  "mcpServers": {
    "SpeakeasyBar": {
      "command": "npx",
      "args": [
        "-y", "--package", "speakeasy-bar",
        "--",
        "mcp", "start",
        "--api-key", "..."
      ]
    }
  }
}
```

</details>

You can also run MCP servers as a standalone binary with no additional dependencies. You must pull these binaries from available Github releases:

```bash
curl -L -o mcp-server \
    https://github.com/{org}/{repo}/releases/download/{tag}/mcp-server-bun-darwin-arm64 && \
chmod +x mcp-server
```

If the repo is a private repo you must add your Github PAT to download a release `-H "Authorization: Bearer {GITHUB_PAT}"`.


```json
{
  "mcpServers": {
    "Todos": {
      "command": "./DOWNLOAD/PATH/mcp-server",
      "args": [
        "start"
      ]
    }
  }
}
```

For a full list of server arguments, run:

```sh
npx -y --package speakeasy-bar -- mcp start --help
```
<!-- End SDK Installation [installation] -->

<!-- Start Requirements [requirements] -->
## Requirements

For supported JavaScript runtimes, please consult [RUNTIMES.md](RUNTIMES.md).
<!-- End Requirements [requirements] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar();

async function run() {
  const result = await speakeasyBar.authentication.authenticate({});

  console.log(result);
}

run();

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type   | Scheme  | Environment Variable   |
| -------- | ------ | ------- | ---------------------- |
| `apiKey` | apiKey | API key | `SPEAKEASYBAR_API_KEY` |

To authenticate with the API the `apiKey` parameter must be set when initializing the SDK client instance. For example:
```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const result = await speakeasyBar.authentication.authenticate({});

  console.log(result);
}

run();

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar();

async function run() {
  const result = await speakeasyBar.drinks.get({});

  console.log(result);
}

run();

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [authentication](docs/sdks/authentication/README.md)

* [authenticate](docs/sdks/authentication/README.md#authenticate) - Authenticate with the API by providing a username and password.

### [configuration](docs/sdks/configuration/README.md)

* [subscribeToWebhooks](docs/sdks/configuration/README.md#subscribetowebhooks) - Subscribe to webhooks.

### [drinks](docs/sdks/drinks/README.md)

* [get](docs/sdks/drinks/README.md#get) - Get a list of drinks.
* [getById](docs/sdks/drinks/README.md#getbyid) - Get a drink.
* [update](docs/sdks/drinks/README.md#update) - Update a drink.
* [updateString](docs/sdks/drinks/README.md#updatestring) - Update a drink.
* [updateRaw](docs/sdks/drinks/README.md#updateraw) - Update a drink.
* [updateMultipart](docs/sdks/drinks/README.md#updatemultipart) - Update a drink.
* [delete](docs/sdks/drinks/README.md#delete) - Delete a drink.
* [search](docs/sdks/drinks/README.md#search) - Search for drinks.
* [stockUpdate](docs/sdks/drinks/README.md#stockupdate) - Receive stock updates.

### [ingredients](docs/sdks/ingredients/README.md)

* [get](docs/sdks/ingredients/README.md#get) - Get a list of ingredients.
* [stockUpdate](docs/sdks/ingredients/README.md#stockupdate) - Receive stock updates.

### [orders](docs/sdks/orders/README.md)

* [create](docs/sdks/orders/README.md#create) - Create an order.


</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Standalone functions [standalone-funcs] -->
## Standalone functions

All the methods listed above are available as standalone functions. These
functions are ideal for use in applications running in the browser, serverless
runtimes or other environments where application bundle size is a primary
concern. When using a bundler to build your application, all unused
functionality will be either excluded from the final bundle or tree-shaken away.

To read more about standalone functions, check [FUNCTIONS.md](./FUNCTIONS.md).

<details>

<summary>Available standalone functions</summary>

- [`authenticationAuthenticate`](docs/sdks/authentication/README.md#authenticate) - Authenticate with the API by providing a username and password.
- [`configurationSubscribeToWebhooks`](docs/sdks/configuration/README.md#subscribetowebhooks) - Subscribe to webhooks.
- [`drinksDelete`](docs/sdks/drinks/README.md#delete) - Delete a drink.
- [`drinksGet`](docs/sdks/drinks/README.md#get) - Get a list of drinks.
- [`drinksGetById`](docs/sdks/drinks/README.md#getbyid) - Get a drink.
- [`drinksSearch`](docs/sdks/drinks/README.md#search) - Search for drinks.
- [`drinksStockUpdate`](docs/sdks/drinks/README.md#stockupdate) - Receive stock updates.
- [`drinksUpdate`](docs/sdks/drinks/README.md#update) - Update a drink.
- [`drinksUpdateMultipart`](docs/sdks/drinks/README.md#updatemultipart) - Update a drink.
- [`drinksUpdateRaw`](docs/sdks/drinks/README.md#updateraw) - Update a drink.
- [`drinksUpdateString`](docs/sdks/drinks/README.md#updatestring) - Update a drink.
- [`ingredientsGet`](docs/sdks/ingredients/README.md#get) - Get a list of ingredients.
- [`ingredientsStockUpdate`](docs/sdks/ingredients/README.md#stockupdate) - Receive stock updates.
- [`ordersCreate`](docs/sdks/orders/README.md#create) - Create an order.

</details>
<!-- End Standalone functions [standalone-funcs] -->

<!-- Start File uploads [file-upload] -->
## File uploads

Certain SDK methods accept files as part of a multi-part request. It is possible and typically recommended to upload files as a stream rather than reading the entire contents into memory. This avoids excessive memory consumption and potentially crashing with out-of-memory errors when working with very large files. The following example demonstrates how to attach a file stream to a request.

> [!TIP]
>
> Depending on your JavaScript runtime, there are convenient utilities that return a handle to a file without reading the entire contents into memory:
>
> - **Node.js v20+:** Since v20, Node.js comes with a native `openAsBlob` function in [`node:fs`](https://nodejs.org/docs/latest-v20.x/api/fs.html#fsopenasblobpath-options).
> - **Bun:** The native [`Bun.file`](https://bun.sh/docs/api/file-io#reading-files-bun-file) function produces a file handle that can be used for streaming file uploads.
> - **Browsers:** All supported browsers return an instance to a [`File`](https://developer.mozilla.org/en-US/docs/Web/API/File) when reading the value from an `<input type="file">` element.
> - **Node.js v18:** A file stream can be created using the `fileFrom` helper from [`fetch-blob/from.js`](https://www.npmjs.com/package/fetch-blob).

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  apiKey: process.env["SPEAKEASYBAR_API_KEY"] ?? "",
});

async function run() {
  const result = await speakeasyBar.drinks.updateRaw({
    productCode: "AC-A2DF3",
    drink: bytesToStream(
      new TextEncoder().encode(
        "{\"name\":\"Old Fashioned\",\"type\":\"cocktail\",\"price\":1000,\"photo\":\"https://speakeasy.bar/drinks/old_fashioned.jpg\",\"productCode\":\"AC-A2DF3\"}",
      ),
    ),
  });

  console.log(result);
}

run();

```
<!-- End File uploads [file-upload] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries.  If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API.  However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a retryConfig object to the call:
```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar();

async function run() {
  const result = await speakeasyBar.authentication.authenticate({}, {
    retries: {
      strategy: "backoff",
      backoff: {
        initialInterval: 1,
        maxInterval: 50,
        exponent: 1.1,
        maxElapsedTime: 100,
      },
      retryConnectionErrors: false,
    },
  });

  console.log(result);
}

run();

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  retryConfig: {
    strategy: "backoff",
    backoff: {
      initialInterval: 1,
      maxInterval: 50,
      exponent: 1.1,
      maxElapsedTime: 100,
    },
    retryConnectionErrors: false,
  },
});

async function run() {
  const result = await speakeasyBar.authentication.authenticate({});

  console.log(result);
}

run();

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

[`SpeakeasyBarError`](./src/models/errors/speakeasybarerror.ts) is the base class for all HTTP error responses. It has the following properties:

| Property            | Type       | Description                                                                             |
| ------------------- | ---------- | --------------------------------------------------------------------------------------- |
| `error.message`     | `string`   | Error message                                                                           |
| `error.statusCode`  | `number`   | HTTP response status code eg `404`                                                      |
| `error.headers`     | `Headers`  | HTTP response headers                                                                   |
| `error.body`        | `string`   | HTTP body. Can be empty string if no body is returned.                                  |
| `error.rawResponse` | `Response` | Raw HTTP response                                                                       |
| `error.data$`       |            | Optional. Some errors may contain structured data. [See Error Classes](#error-classes). |

### Example
```typescript
import { SpeakeasyBar } from "speakeasy-bar";
import * as errors from "speakeasy-bar/models/errors";

const speakeasyBar = new SpeakeasyBar();

async function run() {
  try {
    const result = await speakeasyBar.authentication.authenticate({});

    console.log(result);
  } catch (error) {
    // The base class for HTTP error responses
    if (error instanceof errors.SpeakeasyBarError) {
      console.log(error.message);
      console.log(error.statusCode);
      console.log(error.body);
      console.log(error.headers);

      // Depending on the method different errors may be thrown
      if (error instanceof errors.APIError) {
        console.log(error.data$.code); // string
        console.log(error.data$.message); // string
        console.log(error.data$.details); // { [k: string]: any }
      }
    }
  }
}

run();

```

### Error Classes
**Primary errors:**
* [`SpeakeasyBarError`](./src/models/errors/speakeasybarerror.ts): The base class for HTTP error responses.
  * [`APIError`](./src/models/errors/apierror.ts): An error occurred interacting with the API. Status code `5XX`. *

<details><summary>Less common errors (6)</summary>

<br />

**Network errors:**
* [`ConnectionError`](./src/models/errors/httpclienterrors.ts): HTTP client was unable to make a request to a server.
* [`RequestTimeoutError`](./src/models/errors/httpclienterrors.ts): HTTP request timed out due to an AbortSignal signal.
* [`RequestAbortedError`](./src/models/errors/httpclienterrors.ts): HTTP request was aborted by the client.
* [`InvalidRequestError`](./src/models/errors/httpclienterrors.ts): Any input used to create a request is invalid.
* [`UnexpectedClientError`](./src/models/errors/httpclienterrors.ts): Unrecognised or unexpected error.


**Inherit from [`SpeakeasyBarError`](./src/models/errors/speakeasybarerror.ts)**:
* [`ResponseValidationError`](./src/models/errors/responsevalidationerror.ts): Type mismatch between the data returned from the server and the structure expected by the SDK. See `error.rawValue` for the raw value and `error.pretty()` for a nicely formatted multi-line string.

</details>

\* Check [the method documentation](#available-resources-and-operations) to see if the error is applicable.
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally by passing a server index to the `serverIdx: number` optional parameter when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                                               | Variables                        | Description                                 |
| --- | ---------------------------------------------------- | -------------------------------- | ------------------------------------------- |
| 0   | `https://speakeasy.bar`                              |                                  | The production server.                      |
| 1   | `https://staging.speakeasy.bar`                      |                                  | The staging server.                         |
| 2   | `https://{organization}.{environment}.speakeasy.bar` | `organization`<br/>`environment` | A per-organization and per-environment API. |

If the selected server has variables, you may override its default values through the additional parameters made available in the SDK constructor:

| Variable       | Parameter                               | Supported Values                           | Default  | Description                                                   |
| -------------- | --------------------------------------- | ------------------------------------------ | -------- | ------------------------------------------------------------- |
| `organization` | `organization: string`                  | string                                     | `"api"`  | The organization name. Defaults to a generic organization.    |
| `environment`  | `environment: models.ServerEnvironment` | - `"prod"`<br/>- `"staging"`<br/>- `"dev"` | `"prod"` | The environment name. Defaults to the production environment. |

#### Example

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  serverIdx: 2,
  organization: "<value>",
  environment: "dev",
});

async function run() {
  const result = await speakeasyBar.authentication.authenticate({});

  console.log(result);
}

run();

```

### Override Server URL Per-Client

The default server can also be overridden globally by passing a URL to the `serverURL: string` optional parameter when initializing the SDK client instance. For example:
```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar({
  serverURL: "https://api.prod.speakeasy.bar",
});

async function run() {
  const result = await speakeasyBar.authentication.authenticate({});

  console.log(result);
}

run();

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The TypeScript SDK makes API calls using an `HTTPClient` that wraps the native
[Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API). This
client is a thin wrapper around `fetch` and provides the ability to attach hooks
around the request lifecycle that can be used to modify the request or handle
errors and response.

The `HTTPClient` constructor takes an optional `fetcher` argument that can be
used to integrate a third-party HTTP client or when writing tests to mock out
the HTTP client and feed in fixtures.

The following example shows how to use the `"beforeRequest"` hook to to add a
custom header and a timeout to requests and how to use the `"requestError"` hook
to log errors:

```typescript
import { SpeakeasyBar } from "speakeasy-bar";
import { HTTPClient } from "speakeasy-bar/lib/http";

const httpClient = new HTTPClient({
  // fetcher takes a function that has the same signature as native `fetch`.
  fetcher: (request) => {
    return fetch(request);
  }
});

httpClient.addHook("beforeRequest", (request) => {
  const nextRequest = new Request(request, {
    signal: request.signal || AbortSignal.timeout(5000)
  });

  nextRequest.headers.set("x-custom-header", "custom value");

  return nextRequest;
});

httpClient.addHook("requestError", (error, request) => {
  console.group("Request Error");
  console.log("Reason:", `${error}`);
  console.log("Endpoint:", `${request.method} ${request.url}`);
  console.groupEnd();
});

const sdk = new SpeakeasyBar({ httpClient });
```
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Debugging [debug] -->
## Debugging

You can setup your SDK to emit debug logs for SDK requests and responses.

You can pass a logger that matches `console`'s interface as an SDK option.

> [!WARNING]
> Beware that debug logging will reveal secrets, like API tokens in headers, in log messages printed to a console or files. It's recommended to use this feature only during local development and not in production.

```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const sdk = new SpeakeasyBar({ debugLogger: console });
```

You can also enable a default debug logger by setting an environment variable `SPEAKEASYBAR_DEBUG` to true.
<!-- End Debugging [debug] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=speakeasy-bar&utm_campaign=typescript)
