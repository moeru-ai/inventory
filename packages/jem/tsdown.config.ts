import { defineConfig } from 'tsdown'

export default defineConfig({
  entry: './src/index.ts',
  format: ['esm'],
  outDir: './dist',
  dts: {
    sourcemap: true,
    isolatedDeclarations: true,
  },
})
