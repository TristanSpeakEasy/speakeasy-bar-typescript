# CreateOrderOrderUpdateRequest

## Example Usage

```typescript
import { CreateOrderOrderUpdateRequest } from "speakeasy-bar/models/callbacks";

let value: CreateOrderOrderUpdateRequest = {
  order: {
    orderType: "ingredient",
    productCode: "AC-A2DF3",
    quantity: 994963,
  },
};
```

## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `order`                                         | [models.OrderInput](../../models/orderinput.md) | :heavy_minus_sign:                              | An order for a drink or ingredient.             |