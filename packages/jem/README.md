# Just Enough Models

Your universal model catalog, strong typed, everything, everywhere, all at once.

## Usage

Install the package:

```bash
npm install @moeru-ai/jem
```

Query a model has capabilities:

```typescript
import { hasCapabilities } from "@moeru-ai/jem";

const has = hasCapabilities("openai", "gpt-4o", ["tool-call"]);
```

If you don't know what model you will use, you can use `as` keyword to suppress the type error.

```typescript
const has = hasCapabilities("unknown" as ProviderNames, "unknown" as ModelNames<ProviderNames>, ["unknown"] as Capabilities<ProviderNames>[]);
```
