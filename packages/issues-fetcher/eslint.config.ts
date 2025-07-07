import antfu from '@antfu/eslint-config'

export default antfu({
  typescript: true,
  markdown: true,
  rules: {
    'no-console': 'off',
  },
})
