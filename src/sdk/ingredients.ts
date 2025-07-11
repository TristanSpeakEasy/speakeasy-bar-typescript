/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { ingredientsGet } from "../funcs/ingredientsGet.js";
import { ingredientsStockUpdate } from "../funcs/ingredientsStockUpdate.js";
import { ClientSDK, RequestOptions } from "../lib/sdks.js";
import * as models from "../models/index.js";
import * as operations from "../models/operations/index.js";
import * as webhooks from "../models/webhooks/index.js";
import { unwrapAsync } from "../types/fp.js";
import { WebhookRecipient } from "../types/webhooks.js";

export class Ingredients extends ClientSDK {
  /**
   * Get a list of ingredients.
   *
   * @remarks
   * Get a list of ingredients, if authenticated this will include stock levels and product codes otherwise it will only include public information.
   */
  async get(
    request?: operations.ListIngredientsRequest | undefined,
    options?: RequestOptions,
  ): Promise<operations.ListIngredientsResponse> {
    return unwrapAsync(ingredientsGet(
      this,
      request,
      options,
    ));
  }

  /**
   * Receive stock updates.
   *
   * @remarks
   * Receive stock updates from the bar, this will be called whenever the stock levels of a drink or ingredient changes.
   */
  async stockUpdate(
    recipient: WebhookRecipient,
    request: webhooks.StockUpdateRequest,
    options?: RequestOptions,
  ): Promise<models.ErrorT | undefined> {
    return unwrapAsync(ingredientsStockUpdate(
      this,
      recipient,
      request,
      options,
    ));
  }
}
