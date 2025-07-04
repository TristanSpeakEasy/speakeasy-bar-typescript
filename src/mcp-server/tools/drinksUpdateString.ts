/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { drinksUpdateString } from "../../funcs/drinksUpdateString.js";
import * as operations from "../../models/operations/index.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: operations.UpdateDrinkStringRequest$inboundSchema,
};

export const tool$drinksUpdateString: ToolDefinition<typeof args> = {
  name: "drinks-update-string",
  description: `Update a drink.

Update a drink. Only available when authenticated.`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await drinksUpdateString(
      client,
      args.request,
      { fetchOptions: { signal: ctx.signal } },
    ).$inspect();

    if (!result.ok) {
      return {
        content: [{ type: "text", text: result.error.message }],
        isError: true,
      };
    }

    const value = result.value;

    return formatResult(value, apiCall);
  },
};
