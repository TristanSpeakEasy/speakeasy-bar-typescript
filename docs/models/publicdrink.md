# PublicDrink

## Example Usage

```typescript
import { PublicDrink } from "speakeasy-bar/models";

let value: PublicDrink = {
  name: "Old Fashioned",
  price: 1000,
  photo: "https://speakeasy.bar/drinks/old_fashioned.jpg",
};
```

## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `name`                                                           | *string*                                                         | :heavy_check_mark:                                               | The name of the drink.                                           | Old Fashioned                                                    |
| `type`                                                           | [models.DrinkType](../models/drinktype.md)                       | :heavy_minus_sign:                                               | The type of drink.                                               |                                                                  |
| `price`                                                          | *number*                                                         | :heavy_check_mark:                                               | The price of one unit of the drink in US cents.                  | 1000                                                             |
| `photo`                                                          | *string*                                                         | :heavy_minus_sign:                                               | A photo of the drink.                                            | https://speakeasy.bar/drinks/old_fashioned.jpg                   |
| `dataLevel`                                                      | [models.PublicDrinkDataLevel](../models/publicdrinkdatalevel.md) | :heavy_minus_sign:                                               | The level of data included in the response.                      |                                                                  |