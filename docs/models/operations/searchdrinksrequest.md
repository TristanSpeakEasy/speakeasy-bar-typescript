# SearchDrinksRequest

## Example Usage

```typescript
import { SearchDrinksRequest } from "speakeasy-bar/models/operations";

let value: SearchDrinksRequest = {
  query: "<value>",
};
```

## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `query`                                                                      | *string*                                                                     | :heavy_check_mark:                                                           | The search query.                                                            |
| `type`                                                                       | [models.DrinkType](../../models/drinktype.md)                                | :heavy_minus_sign:                                                           | The type of drink to filter by. If not provided all drinks will be returned. |