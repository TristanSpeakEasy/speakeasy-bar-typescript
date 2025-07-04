# CreateOrderResponseBody

The order was created successfully.


## Supported Types

### `models.DrinkOrder`

```typescript
const value: models.DrinkOrder = {
  orderType: "drink",
  productCode: "AC-A2DF3",
  quantity: 531514,
  status: "pending",
};
```

### `models.IngredientOrder`

```typescript
const value: models.IngredientOrder = {
  orderType: "ingredient",
  productCode: "AC-A2DF3",
  quantity: 83527,
  status: "processing",
  deliveryAddress: "123 Main St, Anytown, USA",
};
```

