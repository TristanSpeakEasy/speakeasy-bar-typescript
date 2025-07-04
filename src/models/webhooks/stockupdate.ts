/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import * as z from "zod";
import { safeParse } from "../../lib/schemas.js";
import { Result as SafeParseResult } from "../../types/fp.js";
import { SDKValidationError } from "../errors/sdkvalidationerror.js";
import * as models from "../index.js";

export type StockUpdateRequest = {
  drink?: models.DrinkInput | undefined;
  ingredient?: models.IngredientInput | undefined;
};

/** @internal */
export const StockUpdateRequest$inboundSchema: z.ZodType<
  StockUpdateRequest,
  z.ZodTypeDef,
  unknown
> = z.object({
  drink: models.DrinkInput$inboundSchema.optional(),
  ingredient: models.IngredientInput$inboundSchema.optional(),
});

/** @internal */
export type StockUpdateRequest$Outbound = {
  drink?: models.DrinkInput$Outbound | undefined;
  ingredient?: models.IngredientInput$Outbound | undefined;
};

/** @internal */
export const StockUpdateRequest$outboundSchema: z.ZodType<
  StockUpdateRequest$Outbound,
  z.ZodTypeDef,
  StockUpdateRequest
> = z.object({
  drink: models.DrinkInput$outboundSchema.optional(),
  ingredient: models.IngredientInput$outboundSchema.optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace StockUpdateRequest$ {
  /** @deprecated use `StockUpdateRequest$inboundSchema` instead. */
  export const inboundSchema = StockUpdateRequest$inboundSchema;
  /** @deprecated use `StockUpdateRequest$outboundSchema` instead. */
  export const outboundSchema = StockUpdateRequest$outboundSchema;
  /** @deprecated use `StockUpdateRequest$Outbound` instead. */
  export type Outbound = StockUpdateRequest$Outbound;
}

export function stockUpdateRequestToJSON(
  stockUpdateRequest: StockUpdateRequest,
): string {
  return JSON.stringify(
    StockUpdateRequest$outboundSchema.parse(stockUpdateRequest),
  );
}

export function stockUpdateRequestFromJSON(
  jsonString: string,
): SafeParseResult<StockUpdateRequest, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => StockUpdateRequest$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'StockUpdateRequest' from JSON`,
  );
}
