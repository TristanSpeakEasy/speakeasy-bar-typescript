# StockUpdateRequest

## Example Usage

```typescript
import { StockUpdateRequest } from "speakeasy-bar/models/webhooks";

let value: StockUpdateRequest = {
  drink: {
    name: "Old Fashioned",
    price: 1000,
    photo: "https://speakeasy.bar/drinks/old_fashioned.jpg",
    productCode: "AC-A2DF3",
    ingredients: [
      {
        ingredientProductCode: "AC-A2DF3",
        quantity: 832313,
      },
    ],
  },
  ingredient: {
    name: "Sugar Syrup",
    type: "packaged",
    productCode: "AC-A2DF3",
    photo: "https://speakeasy.bar/ingredients/sugar_syrup.jpg",
  },
};
```

## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `drink`                                                   | [models.DrinkInput](../../models/drinkinput.md)           | :heavy_minus_sign:                                        | N/A                                                       |
| `ingredient`                                              | [models.IngredientInput](../../models/ingredientinput.md) | :heavy_minus_sign:                                        | N/A                                                       |