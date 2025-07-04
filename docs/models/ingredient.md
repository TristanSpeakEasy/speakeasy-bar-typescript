# Ingredient

## Example Usage

```typescript
import { Ingredient } from "speakeasy-bar/models";

let value: Ingredient = {
  name: "Sugar Syrup",
  type: "fresh",
  stock: 10,
  productCode: "AC-A2DF3",
  photo: "https://speakeasy.bar/ingredients/sugar_syrup.jpg",
};
```

## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `name`                                                                             | *string*                                                                           | :heavy_check_mark:                                                                 | The name of the ingredient.                                                        | Sugar Syrup                                                                        |
| `type`                                                                             | [models.IngredientType](../models/ingredienttype.md)                               | :heavy_check_mark:                                                                 | The type of ingredient.                                                            |                                                                                    |
| `stock`                                                                            | *number*                                                                           | :heavy_minus_sign:                                                                 | The number of units of the ingredient in stock, only available when authenticated. | 10                                                                                 |
| `productCode`                                                                      | *string*                                                                           | :heavy_minus_sign:                                                                 | The product code of an ingredient, only available when authenticated.              | AC-A2DF3                                                                           |
| `photo`                                                                            | *string*                                                                           | :heavy_minus_sign:                                                                 | A photo of the ingredient.                                                         | https://speakeasy.bar/ingredients/sugar_syrup.jpg                                  |