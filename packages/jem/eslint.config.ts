import antfu from '@antfu/eslint-config'

export default antfu({
  ignores: ['./src/providers/**/*'],
  typescript: true,
  markdown: true,
})
