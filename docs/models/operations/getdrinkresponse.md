# GetDrinkResponse

## Example Usage

```typescript
import { GetDrinkResponse } from "speakeasy-bar/models/operations";

let value: GetDrinkResponse = {
  headers: {
    "key": [
      "<value 1>",
      "<value 2>",
    ],
    "key1": [
      "<value 1>",
    ],
  },
  result: {
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

| Field                               | Type                                | Required                            | Description                         |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `headers`                           | Record<string, *string*[]>          | :heavy_check_mark:                  | N/A                                 |
| `result`                            | *operations.GetDrinkResponseResult* | :heavy_check_mark:                  | N/A                                 |