# Order

An order for a drink or ingredient.

## Example Usage

```typescript
import { Order } from "speakeasy-bar/models";

let value: Order = {
  orderType: "ingredient",
  productCode: "AC-A2DF3",
  quantity: 347953,
  status: "processing",
};
```

## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `orderType`                                              | [models.OrderOrderType](../models/orderordertype.md)     | :heavy_check_mark:                                       | The type of order.                                       |                                                          |
| `productCode`                                            | *string*                                                 | :heavy_check_mark:                                       | The product code of the drink or ingredient.             | AC-A2DF3                                                 |
| `quantity`                                               | *number*                                                 | :heavy_check_mark:                                       | The number of units of the drink or ingredient to order. |                                                          |
| `status`                                                 | [models.OrderStatus](../models/orderstatus.md)           | :heavy_check_mark:                                       | The status of the order.                                 |                                                          |