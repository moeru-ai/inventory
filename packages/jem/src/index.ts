import * as providers from "./providers";

export type Providers = typeof providers;
export type ProviderNames = keyof Providers;
export type Models<N extends ProviderNames> = Providers[N]["models"];
export type ModelNames<N extends ProviderNames> = Models<N>[number]["modelId"];
export type Capabilities<N extends ProviderNames> = Models<N>[number]["capabilities"][number];

export function hasCapabilities<N extends ProviderNames, M extends ModelNames<N>>(provider: N, modelId: M, capabilities: Capabilities<N>[]) {
  const model = providers[provider].models.find((m) => m.modelId === modelId);
    if (!model) {
    throw new Error(`Model ${modelId} not found`);
  }

  const result: Partial<Record<Capabilities<N>, boolean>> = {};
  capabilities.forEach((capability) => {
    result[capability] = model.capabilities.includes(capability);
  });
  return result;
}
