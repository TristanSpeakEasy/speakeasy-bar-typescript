# Drink

## Example Usage

```typescript
import { Drink } from "speakeasy-bar/models";

let value: Drink = {
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
};
```

## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `name`                                                                         | *string*                                                                       | :heavy_check_mark:                                                             | The name of the drink.                                                         | Old Fashioned                                                                  |
| `type`                                                                         | [models.DrinkType](../models/drinktype.md)                                     | :heavy_minus_sign:                                                             | The type of drink.                                                             |                                                                                |
| `price`                                                                        | *number*                                                                       | :heavy_check_mark:                                                             | The price of one unit of the drink in US cents.                                | 1000                                                                           |
| `photo`                                                                        | *string*                                                                       | :heavy_minus_sign:                                                             | A photo of the drink.                                                          | https://speakeasy.bar/drinks/old_fashioned.jpg                                 |
| `dataLevel`                                                                    | [models.DrinkDataLevel](../models/drinkdatalevel.md)                           | :heavy_minus_sign:                                                             | The level of data included in the response.                                    |                                                                                |
| `stock`                                                                        | *number*                                                                       | :heavy_minus_sign:                                                             | The number of units of the drink in stock, only available when authenticated.  |                                                                                |
| `productCode`                                                                  | *string*                                                                       | :heavy_check_mark:                                                             | The product code of a drink, only available when authenticated.                | AC-A2DF3                                                                       |
| `ingredients`                                                                  | [models.DrinkIngredient](../models/drinkingredient.md)[]                       | :heavy_minus_sign:                                                             | The ingredients required to make the drink, only available when authenticated. |                                                                                |