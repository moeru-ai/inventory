import { providers } from "./providers";
import { models } from "./models";

export type Providers = typeof providers;
export type ProviderNames = Providers[number]["name"];
export type Models = typeof models;
export type ModelIds = Models[number]["modelId"];

type Filter<T extends readonly any[], P> = T extends [infer First, ...infer Rest] ? First extends P ? [First, ...Filter<Rest, P>] : Filter<Rest, P> : [];

export type ModelsByProvider<N extends ProviderNames> = Filter<Models, { provider: N }>;
export type ModelIdsByProvider<N extends ProviderNames> = ModelsByProvider<N>[number]["modelId"];
export type Model<N extends ProviderNames, M extends ModelIdsByProvider<N>> = Filter<ModelsByProvider<N>, { modelId: M }>[number];

export type CapabilitiesByModel<
  N extends ProviderNames,
  M extends ModelIdsByProvider<N>,
> = Model<N, M>['capabilities'][number];

export function hasCapabilities<N extends ProviderNames, M extends ModelIdsByProvider<N>>(provider: N, modelId: M, capabilities: CapabilitiesByModel<N, M>[]): Record<CapabilitiesByModel<N, M>, boolean> {
  const model = models.find((model) => model.provider === provider && model.modelId === modelId);
  if (!model) {
    throw new Error(`Model ${modelId} not found`);
  }

  const result: Partial<Record<CapabilitiesByModel<N, M>, boolean>> = {};
  capabilities.forEach((capability) => {
    result[capability] = model.capabilities.includes(capability);
  });

  return result as Record<CapabilitiesByModel<N, M>, boolean>;
}
