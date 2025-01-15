import withNuxt from './.nuxt/eslint.config.mjs'
import stylistic from '@stylistic/eslint-plugin'

export default withNuxt(
  {
    files: ['**/*.vue', '**/*.ts'],
    rules: {
      'vue/no-multiple-template-root': 'off',
      'no-console': 'error',
      '@typescript-eslint/no-explicit-any': 'off',
    },
  },
  {
    files: ['**/components/*.vue'],
    rules: {
      'vue/multi-word-component-names': 'error',
    },
  },
  {
    files: ['**/pages/*.vue'],
    rules: {
      'vue/require-v-for-key': 'error',
      'vue/no-use-v-if-with-v-for': 'error',
    },
  },
  stylistic.configs.customize({
    indent: 2,
    quotes: 'single',
    semi: false,
  }),
)
