<!-- Start SDK Example Usage [usage] -->
```typescript
import { SpeakeasyBar } from "speakeasy-bar";

const speakeasyBar = new SpeakeasyBar();

async function run() {
  const result = await speakeasyBar.authentication.authenticate({});

  console.log(result);
}

run();

```
<!-- End SDK Example Usage [usage] -->