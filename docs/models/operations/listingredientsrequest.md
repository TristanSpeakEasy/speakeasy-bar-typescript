# ListIngredientsRequest

## Example Usage

```typescript
import { ListIngredientsRequest } from "speakeasy-bar/models/operations";

let value: ListIngredientsRequest = {};
```

## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ingredients`                                                                               | *string*[]                                                                                  | :heavy_minus_sign:                                                                          | A list of ingredients to filter by. If not provided all ingredients will be returned.       |
| `drinkIngredients`                                                                          | [operations.DrinkIngredient](../../models/operations/drinkingredient.md)[]                  | :heavy_minus_sign:                                                                          | A list of drink ingredients to filter by. If not provided all ingredients will be returned. |