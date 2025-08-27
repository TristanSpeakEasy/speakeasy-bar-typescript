# UpdateDrinkRequest

## Example Usage

```typescript
import { UpdateDrinkRequest } from "speakeasy-bar/models/operations";

let value: UpdateDrinkRequest = {
  productCode: "AC-A2DF3",
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
};
```

## Fields

| Field                                           | Type                                            | Required                                        | Description                                     | Example                                         |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `productCode`                                   | *string*                                        | :heavy_check_mark:                              | N/A                                             | AC-A2DF3                                        |
| `drink`                                         | [models.DrinkInput](../../models/drinkinput.md) | :heavy_check_mark:                              | N/A                                             |                                                 |