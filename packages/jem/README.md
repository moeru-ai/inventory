# Just Enough Models

Your universal model catalog, strong typed, everything, everywhere, all at once.

## Usage

Install the package:

```bash
npm install @proj-airi/jem
```

Query a model has capabilities:

```typescript
import { hasCapabilities } from '@proj-airi/jem'

const has = hasCapabilities('openai', 'gpt-4o', ['tool-call'])
```

If you don't know what model you will use, you can use `as` keyword to suppress the type error.

```typescript
import { hasCapabilities } from '@proj-airi/jem'
import type { ProviderNames, ModelNames, Capabilities } from '@proj-airi/jem'

const has = hasCapabilities("unknown" as ProviderNames, "unknown" as ModelNames<ProviderNames>, ["unknown"] as Capabilities<ProviderNames>[]);
```
