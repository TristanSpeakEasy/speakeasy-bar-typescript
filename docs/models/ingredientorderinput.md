# IngredientOrderInput

An order for a drink or ingredient.

## Example Usage

```typescript
import { IngredientOrderInput } from "speakeasy-bar/models";

let value: IngredientOrderInput = {
  orderType: "ingredient",
  productCode: "AC-A2DF3",
  quantity: 700871,
  deliveryAddress: "123 Main St, Anytown, USA",
};
```

## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `orderType`                                                              | [models.IngredientOrderOrderType](../models/ingredientorderordertype.md) | :heavy_check_mark:                                                       | N/A                                                                      |                                                                          |
| `productCode`                                                            | *string*                                                                 | :heavy_check_mark:                                                       | The product code of the drink or ingredient.                             | AC-A2DF3                                                                 |
| `quantity`                                                               | *number*                                                                 | :heavy_check_mark:                                                       | The number of units of the drink or ingredient to order.                 |                                                                          |
| `deliveryAddress`                                                        | *string*                                                                 | :heavy_minus_sign:                                                       | The address to deliver the ingredient to.                                | 123 Main St, Anytown, USA                                                |