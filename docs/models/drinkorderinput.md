# DrinkOrderInput

An order for a drink or ingredient.

## Example Usage

```typescript
import { DrinkOrderInput } from "speakeasy-bar/models";

let value: DrinkOrderInput = {
  orderType: "drink",
  productCode: "AC-A2DF3",
  quantity: 707819,
};
```

## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `orderType`                                                    | [models.DrinkOrderOrderType](../models/drinkorderordertype.md) | :heavy_check_mark:                                             | N/A                                                            |                                                                |
| `productCode`                                                  | *string*                                                       | :heavy_check_mark:                                             | The product code of the drink or ingredient.                   | AC-A2DF3                                                       |
| `quantity`                                                     | *number*                                                       | :heavy_check_mark:                                             | The number of units of the drink or ingredient to order.       |                                                                |
| `barCounter`                                                   | [models.BarCounter](../models/barcounter.md)                   | :heavy_minus_sign:                                             | The bar counter to collect the drink from.                     |                                                                |