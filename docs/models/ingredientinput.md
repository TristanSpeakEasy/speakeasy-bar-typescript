# IngredientInput

## Example Usage

```typescript
import { IngredientInput } from "speakeasy-bar/models";

let value: IngredientInput = {
  name: "Sugar Syrup",
  type: "long-life",
  productCode: "AC-A2DF3",
  photo: "https://speakeasy.bar/ingredients/sugar_syrup.jpg",
};
```

## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `name`                                                                | *string*                                                              | :heavy_check_mark:                                                    | The name of the ingredient.                                           | Sugar Syrup                                                           |
| `type`                                                                | [models.IngredientType](../models/ingredienttype.md)                  | :heavy_check_mark:                                                    | The type of ingredient.                                               |                                                                       |
| `productCode`                                                         | *string*                                                              | :heavy_minus_sign:                                                    | The product code of an ingredient, only available when authenticated. | AC-A2DF3                                                              |
| `photo`                                                               | *string*                                                              | :heavy_minus_sign:                                                    | A photo of the ingredient.                                            | https://speakeasy.bar/ingredients/sugar_syrup.jpg                     |