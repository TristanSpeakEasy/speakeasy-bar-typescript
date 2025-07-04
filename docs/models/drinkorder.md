# DrinkOrder

An order for a drink or ingredient.

## Example Usage

```typescript
import { DrinkOrder } from "speakeasy-bar/models";

let value: DrinkOrder = {
  orderType: "drink",
  productCode: "AC-A2DF3",
  quantity: 531514,
  status: "pending",
};
```

## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `orderType`                                                    | [models.DrinkOrderOrderType](../models/drinkorderordertype.md) | :heavy_check_mark:                                             | N/A                                                            |                                                                |
| `productCode`                                                  | *string*                                                       | :heavy_check_mark:                                             | The product code of the drink or ingredient.                   | AC-A2DF3                                                       |
| `quantity`                                                     | *number*                                                       | :heavy_check_mark:                                             | The number of units of the drink or ingredient to order.       |                                                                |
| `status`                                                       | [models.DrinkOrderStatus](../models/drinkorderstatus.md)       | :heavy_check_mark:                                             | The status of the order.                                       |                                                                |
| `barCounter`                                                   | [models.BarCounter](../models/barcounter.md)                   | :heavy_minus_sign:                                             | The bar counter to collect the drink from.                     |                                                                |