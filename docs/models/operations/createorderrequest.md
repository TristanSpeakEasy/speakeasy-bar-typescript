# CreateOrderRequest

## Example Usage

```typescript
import { CreateOrderRequest } from "speakeasy-bar/models/operations";

let value: CreateOrderRequest = {
  requestBody: {
    orderType: "ingredient",
    productCode: "AC-A2DF3",
    quantity: 608897,
  },
};
```

## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `callbackUrl`                              | *string*                                   | :heavy_minus_sign:                         | The url to call when the order is updated. |
| `requestBody`                              | *operations.CreateOrderRequestBody*        | :heavy_check_mark:                         | N/A                                        |