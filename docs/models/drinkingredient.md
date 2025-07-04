# DrinkIngredient

An ingredient required to make a drink.

## Example Usage

```typescript
import { DrinkIngredient } from "speakeasy-bar/models";

let value: DrinkIngredient = {
  ingredientProductCode: "AC-A2DF3",
  quantity: 920472,
};
```

## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ingredientProductCode`                                               | *string*                                                              | :heavy_check_mark:                                                    | The product code of an ingredient, only available when authenticated. | AC-A2DF3                                                              |
| `quantity`                                                            | *number*                                                              | :heavy_check_mark:                                                    | The number of units of the ingredient required to make the drink.     |                                                                       |