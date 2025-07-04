# UpdateDrinkRawResponse


## Supported Types

### `models.Drink`

```typescript
const value: models.Drink = {
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

### `Uint8Array`

```typescript
const value: Uint8Array = new TextEncoder().encode("0xdC7D3f4be0");
```

### `models.ErrorT`

```typescript
const value: models.ErrorT = {
  code: "<value>",
  message: "<value>",
};
```

